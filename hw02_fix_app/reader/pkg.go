package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/DmitryFPS/go/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	res := data

	return res, nil
}
