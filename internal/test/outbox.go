package test

import (
	"github.com/ssup2-playground/msa-event-market_service-auth/pkg/entity/uuid"
)

const (
	OutboxIDWrongFormat        = "aaaa-aaaa"
	OutboxAggregateTypeCorrect = "aggType"
	OutboxAggregateIDCorrect   = "aggID"
	OutboxEventTypeCorrect     = "eventID"
	OutboxPayloadCorrect       = "payload"
	OutboxSpanContextCorrect   = "spanContext"
)

var (
	OutboxIDCorrect = uuid.FromStringOrNil("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
)
