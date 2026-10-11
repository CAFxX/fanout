package main

import (
    "fmt"
)

func genFn0(p0 int8) int8 {
    p0 += p0
    p0 -= p0
    if true {
        var v5 int = -687
        _ = v5
    } else {
        var v6 float64 = (1e10 * 1.5)
        _ = v6
    }
    _ = p0
    return p0
}
func genFn1(p0 uint) int64 {
    p0 += (^(p0))
    var v7 int8 = 100
    v8 := []int{-9223372036854775808}
    _ = len(v8)
    switch "" {
    case "\n\t":
        if len(v8) > 0 {
            v8[0] = 42
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := []int{-2147483648, 1, 42, 128}
    v2 := []int{}
    v3 := []int{4294967295, 692, 2147483647}
    var v4 string = "foo bar"
    var v9 bool = (((-2.25 * (float64(0.1) * float64((((3.14159 - 1e10) - (0.0 * (0.0 - 3.14159))) / 3.14159)))) * ((1e10 / (float64(float32((float64(3.14159) * 1.5))) / 1e10)) * -2.25)) > -2.25)
    if (2 < 42) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (1.5 > float32(3.14159)) {
            v9 = v9
            var v10 float32 = -2.25
            var v11 uint16 = 358
            _ = v10
            _ = v11
        }
    }
    var v12 int16 = 207
    var v13 bool = (-462 == 9223372036854775807)
    switch 127 {
    case 389:
        var v14 int8 = 42
        _ = v14
    }
    v2 = append(v2, int(v12))
    v4 = (v4 + v4)
    switch 4294967295 {
    case -985:
        var v15 uint = 0
        _ = v15
    case 0:
        var v16 int64 = 241
        _ = v16
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch (v4 + "\n\t") {
    case "hello":
        v4 = v4
    case "\n\t":
        v13 = v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
