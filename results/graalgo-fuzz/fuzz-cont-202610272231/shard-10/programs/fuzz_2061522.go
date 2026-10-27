package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint32 }
func genFn0() int64 {
    var v7 int64 = 9223372036854775807
    v8 := GenS1{F0: 4294967295, F1: (v7 >> ((4294967295) & 63)), F2: 2}
    v7 = v7
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{96, 401, 32768}
    var v3 uint16 = 65535
    var v4 float32 = -2.25
    var v5 uint16 = v3
    var v6 int = -2147483648
    switch v6 {
    case 4294967295:
        if true {
            v3 = 0
        } else {
            fmt.Println(v6)
            fmt.Println(v5, v3, v3)
        }
    case 981:
        switch (("\n\t" + (("a" + "foo bar") + (("\n\t" + "") + "a"))) + "a") {
        case "\n\t":
            fmt.Println(v4, v6, v6)
        case "foo bar":
            if true {
                v5 += v3
            }
        case "":
            var v9 string = (("\n\t" + (("a" + "foo bar") + "")) + "foo bar")
            _ = v9
        }
    default:
        var v10 int8 = 124
        _ = v10
    }
    v11 := GenS2{F0: v3, F1: 128}
    switch (v6 + 62) {
    case 2147483647:
        var v12 int16 = -32768
        _ = v12
    }
    fmt.Println(v3, v5, v3)
    if false {
        if (!true) {
            switch (v6 + (v6 * v6)) {
            case -2147483648:
                if (0 <= 2147483648) {
                    v13 := GenS3{F0: 42, F1: 1}
                    v5 = 1
                    _ = v13
                }
            case 0:
                v14 := GenS4{F0: 2147483647}
                _ = v14
            case 552:
                var v15 bool = true
                _ = v15
            default:
                if len(v2) > 0 {
                    v2[0] = (^(v6))
                }
            }
            v16 := GenS2{F0: v3, F1: 255}
            _ = v16
        } else {
            var v17 int = (v6 * -2147483648)
            v17 -= -2147483648
            _ = v17
        }
    } else {
        var v18 float64 = float64(v4)
        var v19 uint8 = 42
        _ = v18
        _ = v19
    }
    v2 = append(v2, (v6 - (-(v6))))
    v2 = append(v2, (v6 * 9223372036854775807))
    var v20 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
