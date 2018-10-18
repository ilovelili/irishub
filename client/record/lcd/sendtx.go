package lcd

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"net/http"
	"os"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/irisnet/irishub/client/context"
	"github.com/irisnet/irishub/client/utils"
	"github.com/irisnet/irishub/modules/record"
)

type postRecordReq struct {
	BaseTx      context.BaseTx `json:"base_tx"`
	FilePath    string         `json:"file_path"`
	Data        string         `json:"data"`
	Submitter   string         `json:"submitter"` //  Address of the submitter
	PinedNode   string         `json:"pined_node"`
	Description string         `json:"description"`
}

func postRecordHandlerFn(cdc *wire.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req postRecordReq
		err := utils.ReadPostBody(w, r, cdc, &req)
		if err != nil {
			return
		}

		cliCtx = utils.InitRequestClictx(cliCtx, r, req.BaseTx.LocalAccountName, req.Submitter)
		txCtx, err := context.NewTxContextFromBaseTx(cliCtx, cdc, req.BaseTx)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		submitter, err := sdk.AccAddressFromBech32(req.Submitter)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		onchainData := req.Data
		filePath := req.FilePath

		var recordHash string
		var dataSize int64
		// --onchain-data has a high priority over --file-path
		if len(onchainData) != 0 {
			dataSize = int64(binary.Size([]byte(onchainData)))
			sum := sha256.Sum256([]byte(onchainData))
			recordHash = hex.EncodeToString(sum[:])
		} else if len(filePath) != 0 {

			var fileInfo os.FileInfo
			if fileInfo, err = os.Stat(req.FilePath); os.IsNotExist(err) {
				utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			dataSize = fileInfo.Size()

			//upload to ipfs
			sh := ipfs.NewShell(req.PinedNode)
			f, err := os.Open(req.FilePath)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			recordHash, err = sh.Add(f)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		} else {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		submitTime := time.Now().Unix()

		// create the message
		msg := record.NewMsgSubmitFile(req.Description, submitTime, submitter, recordHash, dataSize, onchainData)
		err = msg.ValidateBasic()
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.SendOrReturnUnsignedTx(w, cliCtx, txCtx, req.BaseTx, []sdk.Msg{msg})
	}
}
