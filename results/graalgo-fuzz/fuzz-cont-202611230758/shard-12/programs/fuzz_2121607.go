package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0(p0 int32, p1 string) bool {
    var v6 float32 = ((1.5 + (0.0 - ((-2.25 - (1.5 / -2.25)) / (0.1 / 1.5)))) * 1.5)
    fmt.Println(v6, v6, v6)
    v6 = v6
    if (42 < 2) {
        p1 = "hello"
        v7 := GenS1{F0: 127, F1: 65535}
        _ = v7
    }
    switch 42 {
    case 949:
        p0 += (p0 + (p0 * p0))
    case 4294967295:
        p1 = (("" + "hello") + ("a" + "\n\t"))
    case 42:
        v8 := []int{32767, 175}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return (v6 <= 0.0)
}

func main() {
    v1 := []uint32{42, 2147483648}
    v2 := make(map[string]uint32)
    var v3 int32 = 2147483647
    v4 := make(map[int]bool)
    var v5 uint32 = 2
    v9 := GenS1{F0: 42, F1: 2}
    v10 := GenS1{F0: -78, F1: 732}
    if (!((18446744073709551615 >= 1) && (!false))) {
        v5 -= 699
        for fi11 := 0; fi11 < 4; fi11++ {
            switch "" {
            case "foo bar":
                _ = len(v2)
            case "\n\t":
                v5 -= (v5 / ((4294967295) | 1))
            case "":
                if (((3.14159 / (3.14159 + (((3.14159 + 0.1) + float32(1e10)) - ((float32(((1.5 * (0.1 - (0.0 * float64(1e10)))) + (0.1 / float64((1.5 * (float32(1e10) / 1.5)))))) - 3.14159) / (float32((0.1 / 1.0)) + 0.0))))) + float32((((float64(3.14159) + -2.25) - -2.25) - 1e10))) != 3.14159) {
                    v3 += v3
                }
            default:
                var v12 float32 = (3.14159 - ((((3.14159 * (3.14159 * 0.0)) * 1.5) - -2.25) + -2.25))
                _ = v12
            }
            if len(v1) > 0 {
                v1[0] = (v5 + (^(v5)))
            }
            _ = fi11
        }
    } else {
        v3 += v3
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        if (true || (float32(float64((1.5 - 0.0))) <= -2.25)) {
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS2{F0: (fi14 | fi14)}
                _ = fi14
                _ = v15
            }
            var v16 int8 = 1
            var v17 int8 = int8(v3)
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: 8, F1: 0}
        _ = fi13
        _ = v18
    }
    v3 = v3
    v19 := GenS1{F0: 46, F1: 445}
    v20 := GenS3{F0: 0, F1: 0}
    _ = len(v1)
    if ((1.5 - 0.0) <= (-2.25 - ((((3.14159 / (1e10 / (1.5 - (0.1 / 3.14159)))) + 3.14159) + (0.0 * 3.14159)) - (((float32(((0.1 + 0.0) + -2.25)) + ((float32(-2.25) + 1.5) - 3.14159)) - (float32(3.14159) * -2.25)) - float32(1.5))))) {
        var v21 float32 = ((3.14159 - 0.0) * -2.25)
        _ = v21
    } else {
        _ = len(v1)
        v5 = (v5 & 624)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
