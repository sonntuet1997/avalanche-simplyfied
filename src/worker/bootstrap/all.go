// Code generated by cli. DO NOT EDIT.
package bootstrap

import "go.uber.org/fx"

func All() fx.Option {
	return fx.Options(
		GeneratedConfig(),
		Custom(),
	)
}
