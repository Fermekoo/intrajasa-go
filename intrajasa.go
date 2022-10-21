package intrajasa

type EnvType int8

const (
	_ EnvType = iota

	Sandbox

	Production
)

type VaType int8

const (
	_ VaType = iota
	OneTime
	Continues
	Phone
)

var BaseUrl = map[EnvType]string{
	Sandbox:    "https://sandbox-va.kliringindonesia.co.id",
	Production: "https://va.kliringindonesia.co.id",
}
