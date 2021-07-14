package validate

const (
	minLenght = 1
	maxLenght = 1000
)

func Validate(arr []int) bool {
	if len(arr) < minLenght || len(arr) > maxLenght {
		return false
	}
	return true
}
