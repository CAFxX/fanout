package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int; F2 string }

func main() {
    var v1 float64 = 0.0
    var v2 float32 = 0.0
    var v3 int32 = -1
    v4 := make(map[string]bool)
    v2 -= (float32((v1 * v1)) + v2)
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 uint8 = 128
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint16 = 32767
            _ = fi7
            _ = v8
        }
        switch ("" + "hello") {
        case "foo bar":
            delete(v4, (("" + "foo bar") + "\n\t"))
        case "\n\t":
            var v9 float32 = v2
            _ = v9
        case "":
            if (!false) {
                _ = len(v4)
            }
        default:
            fmt.Println(fi5)
        }
        _ = fi5
        _ = v6
    }
    if ("hello" <= "hello") {
        v2 = float32(-2.25)
    }
    if (4294967295 >= -37) {
        var v10 int16 = 561
        _ = v10
    } else {
        var v11 int32 = 42
        _ = v11
    }
    v4["hello"] = (((1e10 + -2.25) == float64((v2 / 3.14159))) || true)
    v12 := v4["a"]
    var v13 int32 = v3
    fmt.Println(v13, v1, v3)
    v14 := GenS1{F0: 42, F1: 0, F2: (127 == -1)}
    if (true && (1 != 0)) {
        switch 9223372036854775807 {
        case -1:
            if v12 {
                if ((0 == 18446744073709551615) || v12) {
                    v3 = v3
                    v15 := GenS1{F0: 42, F1: 18446744073709551615, F2: ((4294967295 != 269) || v12)}
                    v16 := v4["hello"]
                    _ = v15
                    _ = v16
                }
            } else {
                switch "a" {
                case "":
                    fmt.Println(v1, v3)
                default:
                    var v17 uint8 = 255
                    _ = v17
                }
            }
        case 9223372036854775807:
            v12 = (false || false)
        default:
            v18 := GenS2{F0: 665}
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            for fi20 := 0; fi20 < 18; fi20++ {
                v3 = v3
                _ = fi20
            }
            v2 += v2
            v2 += (v2 - float32((v1 / v1)))
            _ = fi19
        }
        v21 := GenS2{F0: 42}
        _ = v21
    } else {
        v13 = (v3 - (v13 - (-(v3))))
    }
    var v22 float64 = (v1 * ((v1 / v1) / -2.25))
    v23 := GenS3{F0: 128, F1: -2147483648, F2: "hello"}
    v2 = -2.25
    v24 := []int{9223372036854775807, 32768}
    if true {
        fmt.Println(v3, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
