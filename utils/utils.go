package utils

import "github.com/google/uuid"

func GetShortUUID(UUID uuid.UUID) string {
	return UUID.String()[0:8]
}