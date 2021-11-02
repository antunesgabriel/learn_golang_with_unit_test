package structs

import "testing"

func TestPerimeter(t *testing.T) {
    feedback := func(expected, result *float64, t *testing.T) {
        if *expected != *result {
            t.Errorf("[FAILED] - Expected: %.2f - Result: %.2f", *expected, *result)
        }
    }

    t.Run("Rectangle", func(t *testing.T) {
        t.Run("calc perimeter", func(t *testing.T) {
            rectangle := Rectangle{
                10.0,
                10.0,
            }

            expected := 40.0
            result := Perimeter(rectangle)

            feedback(&expected, &result, t)
        })

        t.Run("calc area", func(t *testing.T) {
            rectangle := Rectangle{
                6.0,
                12.0,
            }

            result := Area(rectangle)
            expected := 72.0

            feedback(&expected, &result, t)
        })
    })

    t.Run("Circle", func(t *testing.T) {
        t.Run("calc area", func(t *testing.T) {
            circle := Circle{10}

            expected := 314.1592653589793
            result := Area(circle)

            feedback(&expected, &result, t)
        })
    })


}
