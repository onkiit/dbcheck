package registry

var model []interface{}

func RegisterModels(models ...interface{}) {
	model = append(model, models...)
}

func Model() []interface{} {
	return model
}
