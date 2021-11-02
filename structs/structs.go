package structs

type Rectangle struct {
    Height float64
    Width float64
}

type Circle struct {
    Radius float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
    perimeter = 2 * (rectangle.Width + rectangle.Height)
    return
}

func Area(rectangle Rectangle) (area float64) {
    area = rectangle.Width * rectangle.Height
    return
}
