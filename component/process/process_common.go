//go:build !android

package process

import "github.com/metacubex/mihomo/constant"

func FindPackageName(metadata *constant.Metadata) (string, error) {
	return "", nil
}
