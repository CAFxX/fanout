package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 string }
func genFn0() int16 {
    var v5 int = 908
    v6 := GenS1{F0: 0, F1: 18446744073709551615}
    var v7 uint16 = 848
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 42
    var v2 uint8 = 128
    var v3 uint8 = (v2 * v2)
    var v4 float64 = 0.0
    if (2147483648 < 65535) {
        v8 := []int{32768, 32767, 121}
        fmt.Println(v1, v4, v2)
        _ = v8
    }
    switch -643 {
    case -525:
        v2 -= (v2 + 2)
    }
    v9 := GenS2{F0: "\n\t"}
    v10 := GenS2{F0: ("foo bar" + ("hello" + ""))}
    for fi11 := 0; fi11 < 12; fi11++ {
        v2 += 71
        _ = fi11
    }
    switch 42 {
    case -9223372036854775808:
        fmt.Println(v1, v1)
    case 4294967295:
        switch "foo bar" {
        case "\n\t":
            var v12 int16 = 1
            _ = v12
        }
    }
    fmt.Println(v3, v1, v1)
    var v13 int = 272
    var v14 int16 = 0
    if true {
        v15 := []int{v13, int(v2)}
        _ = v15
    }
    v2 = v2
    v16 := []int{int(v1), (v13 / ((int(v3)) | 1))}
    for fi17 := 0; fi17 < 15; fi17++ {
        if len(v16) > 0 {
            v16[0] = v13
        }
        if (uint32(v3) != 2147483648) {
            v18 := GenS2{F0: "\n\t"}
            _ = v18
        }
        _ = fi17
    }
    if (-128 != 42) {
        var v19 float32 = 3.14159
        if (v14 <= (v14 ^ (v14 * (-(v14))))) {
            fmt.Println(v1, v13)
            v2 += v2
            fmt.Println(v1, v4, v4)
        }
        v4 = (float64(0.0) + float64(float32(float64(v19))))
        _ = v19
    }
    var v20 uint = 1
    switch ("foo bar" + "a") {
    case "a":
        v1 = 1
    case "\n\t":
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    default:
        v4 -= (1e10 * float64(1e10))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
