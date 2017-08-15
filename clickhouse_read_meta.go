package clickhouse

import (
	"fmt"

	"github.com/kshvakov/clickhouse/internal/data"
	"github.com/kshvakov/clickhouse/internal/protocol"
)

func (ch *clickhouse) readMeta() (*data.Block, error) {
	packet, err := ch.decoder.Uvarint()
	if err != nil {
		return nil, err
	}
	switch packet {
	case protocol.ServerData:
		block, err := ch.readBlock()
		if err != nil {
			return nil, err
		}
		ch.logf("[read meta] <- data: packet=%d, columns=%d, rows=%d", packet, block.NumColumns, block.NumRows)
		return block, nil
	case protocol.ServerException:
		ch.logf("[read meta] <- exception")
		return nil, ch.exception()
	default:
		return nil, fmt.Errorf("unexpected packet [%d] from server", packet)
	}
}
