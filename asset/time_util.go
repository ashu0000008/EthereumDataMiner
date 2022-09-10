package asset

import (
	"fmt"
	"time"
)

func GetDateNow() string {
	return fmt.Sprintf("%d-%d-%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
}
