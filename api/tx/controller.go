package tx

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/internal-tx/cache"
	"github.com/internal-tx/util"
)

// eth: https://go.getblock.io/8d7a9e5561c843508e6607243f8b9147
func GetInternalTx(c *fiber.Ctx) error {
	req := &QueryInternalTx{}
	if err := c.QueryParser(req); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	// get from cache
	key := fmt.Sprintf("%v-%v", req.JsonRPC, req.TxHash)
	v, ok := cache.Get(strings.ToLower(key))
	if ok {
		return c.Status(http.StatusOK).JSON(util.SuccessResponse(v))
	}

	ethClient, err := ethclient.DialContext(c.Context(), req.JsonRPC)
	if err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	result := map[string]interface{}{}
	if err := ethClient.Client().CallContext(c.Context(), &result, "debug_traceTransaction", req.TxHash, map[string]string{
		"tracer": "callTracer",
	}); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	cache.Add(strings.ToLower(key), result)

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(result))
}
