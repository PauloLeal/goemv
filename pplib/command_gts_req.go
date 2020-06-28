package pplib

import (
	"errors"
	"fmt"
	"strconv"
)

type GtsRequest struct {
	AcquirerIndex string `json:"acquirer_index"`
}

func (cmd *GtsRequest) GetName() string {
	return "GTS"
}

func (cmd *GtsRequest) Validate() error {
	if !isValidDataType(cmd.AcquirerIndex, NUMBER, 2) {
		return errors.New("invalid AcquirerIndex value")
	}
	return nil
}

func (cmd *GtsRequest) Parse(rawData string) error {
	pr := NewPositionalReader(rawData)

	cmdName := pr.Read(3)
	if cmdName != cmd.GetName() {
		return errors.New(fmt.Sprintf("cannot parse %s command", cmd.GetName()))
	}

	size, err := strconv.Atoi(pr.Read(3))
	if err != nil {
		return err
	}

	cmd.AcquirerIndex = pr.Read(size)

	return cmd.Validate()
}

func (cmd *GtsRequest) String() string {
	err := cmd.Validate()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s002%02s", cmd.GetName(), cmd.AcquirerIndex)
}
