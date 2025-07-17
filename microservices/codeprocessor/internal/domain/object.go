package domain

import domain "codeproc/microservices/codeprocessor/internal/domain/entity"

type Object interface {
	Do(domain.Object)
}
