package resource

import (
	"app.deploy/constant"
	"fmt"
)

func GetAppLabelKey() string {
	return fmt.Sprintf("%s/%s", constant.ControllerGroup, constant.ControllerVersionV1)
}
