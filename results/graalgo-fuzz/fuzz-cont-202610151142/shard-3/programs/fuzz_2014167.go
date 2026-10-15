package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64, p1 int32, p2 int32) uint64 {
    var v5 float64 = 3.14159
    if (!false) {
        fmt.Println(v5)
    }
    var v6 uint32 = 2
    if (false && false) {
        p1 += p2
        p0 += p0
    } else {
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(v6, p0, v5)
            var v8 string = "hello"
            _ = fi7
            _ = v8
        }
        fmt.Println(v6, p1)
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int16 = 1000
    v3 := make(map[string]uint32)
    var v4 string = "a"
    fmt.Println(v4)
    v4 = "hello"
    var v10 int8 = 42
    switch v4 {
    case "foo bar":
        fmt.Println(v10, v1)
    }
    v1 -= (v1 / ((0) | 1))
    if (127 <= 1) {
        v4 = (((v4 + v4) + v4) + "foo bar")
        if (!false) {
            var v11 string = v4
            _ = v11
        } else {
            if (0 >= 32767) {
                v10 = v10
                if (42 >= 0) {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        var v13 bool = false
                        _ = fi12
                        _ = v13
                    }
                    var v14 uint64 = (v1 >> ((751) & 63))
                    _ = v14
                } else {
                    var v15 uint16 = 452
                    _ = v15
                }
                if true {
                    var v16 int = 0
                    _ = v16
                } else {
                    var v17 string = (v4 + (v4 + (v4 + ("foo bar" + (v4 + "")))))
                    _ = v17
                }
            }
            var v18 float32 = float32(float64(0.0))
            if ((4294967295 == 65535) && (v1 > (v1 % ((v1) | 1)))) {
                v10 = (v10 - v10)
                fmt.Println(v1, v10, v2)
                v19 := []int{128, -290, 721, -1}
                _ = v19
            }
            _ = v18
        }
    }
    v20 := GenS1{F0: 783}
    for fi21 := 0; fi21 < 10; fi21++ {
        if (((v2 * (v2 & (v2 >> ((255) & 63)))) < v2) || (1 <= 2147483647)) {
            v22 := []int{}
            var v23 int64 = -902
            _ = v22
            _ = v23
        }
        _ = fi21
    }
    var v24 int = 286
    switch v24 {
    case -193:
        switch 536 {
        case 128:
            v3[(v4 + v4)] = 516
        default:
            var v25 float32 = -2.25
            _ = v25
        }
    }
    var v26 int8 = (v10 + (v10 << ((4294967295) & 63)))
    var v27 bool = ((((1e10 / -2.25) * 3.14159) - (float32(((1e10 * float64(1e10)) + 1.5)) / (float32(0.1) / (1.5 / (3.14159 / ((1e10 + 0.0) + 0.0)))))) != -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
