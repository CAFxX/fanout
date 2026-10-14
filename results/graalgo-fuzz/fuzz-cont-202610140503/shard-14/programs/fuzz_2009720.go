package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0() int8 {
    switch "" {
    case "hello":
        var v5 float64 = ((0.0 * -2.25) / float64(((3.14159 * float32(float64((((1e10 + 0.1) / 0.1) - 3.14159)))) / (1.5 - (3.14159 * 1e10)))))
        _ = v5
    }
    switch ("foo bar" + ((("a" + "foo bar") + "a") + "foo bar")) {
    case "\n\t":
        if false {
            var v6 float32 = (0.1 - float32(0.1))
            _ = v6
        }
    case "hello":
        var v7 uint16 = 0
        _ = v7
    }
    if (true && (122 >= 885)) {
        var v8 uint32 = 2147483648
        _ = v8
    }
    return 1
}

func main() {
    v1 := []int{-934, -17, -9223372036854775808, 128}
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 float64 = 1.5
    v2[32767] = true
    var v9 uint16 = 255
    if true {
        v4 = 3.14159
    } else {
        v10 := GenS1{F0: -62, F1: (v9 * 65535)}
        _ = v10
    }
    var v11 uint8 = 79
    var v12 uint = 42
    switch -294 {
    case 42:
        v12 -= v12
    }
    v3["foo bar"] = (float32(((((v4 * v4) * (v4 + (v4 + (float64(float32(0.0)) + ((v4 / v4) - ((float64(3.14159) + (float64(((0.0 * ((3.14159 / 1.5) * 3.14159)) + ((1.5 * (-2.25 / 1e10)) * -2.25))) / float64(3.14159))) * float64((1.5 / 1.5)))))))) / v4) + (v4 + (v4 / v4)))) <= -2.25)
    v13 := v3[""]
    var v14 float32 = (((1.5 * 0.0) - (-2.25 * (1.5 / (float32(v4) - 3.14159)))) + float32(v4))
    v4 = float64(((v14 + 0.1) / (3.14159 - v14)))
    var v15 uint = 953
    v16 := GenS2{F0: (v11 % ((105) | 1)), F1: (v11 | (v11 ^ (v11 % ((v11) | 1))))}
    var v17 float64 = (((0.0 + (3.14159 - 1e10)) * ((-2.25 - v4) / v4)) * 1e10)
    v18 := GenS2{F0: (v11 + 127), F1: (v11 >> (((v15 / (((v15 >> ((v15) & 63))) | 1))) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
