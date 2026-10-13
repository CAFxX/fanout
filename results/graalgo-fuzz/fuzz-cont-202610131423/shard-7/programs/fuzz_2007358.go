package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 int16 = 0
    v2 := make(map[string]string)
    var v3 bool = false
    v4 := []string{"", ("hello" + ("\n\t" + ("\n\t" + ""))), (("foo bar" + (("foo bar" + "a") + "foo bar")) + "foo bar")}
    var v5 int64 = 662
    var v6 int16 = v1
    var v7 int8 = 100
    v5 = -698
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint32 = 260
        _ = fi8
        _ = v9
    }
    v7 -= (v7 * (v7 << ((763) & 63)))
    fmt.Println(v7, v3, v7)
    v7 = (v7 * -110)
    if (v3 || (float32(((0.0 * float64(0.0)) / float64((1.5 / 0.1)))) >= ((0.1 * 3.14159) * (0.1 - 0.0)))) {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint = 4294967295
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v3)
    var v12 uint16 = 128
    var v13 float64 = float64((3.14159 / ((1.5 * 0.1) / ((0.0 - (float32(0.0) * ((((0.1 / 1.0) * (((1.5 - 0.0) * ((0.0 - 0.0) / 3.14159)) + 1.5)) + (((3.14159 + (1e10 / 1e10)) - 1e10) + -2.25)) + 3.14159))) - (0.0 - (0.1 - 3.14159))))))
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            if true {
                v13 = ((0.0 / (((float64((0.1 * (0.0 + (3.14159 + 3.14159)))) / v13) * 1e10) + (v13 - v13))) * 0.0)
                v12 = 631
            }
            v16 := GenS1{F0: 255, F1: (^(v6)), F2: 65535}
            if (-630 > fi14) {
                v17 := GenS1{F0: v12, F1: 0, F2: 0}
                v12 -= (^(v12))
                fmt.Println(v5)
                _ = v17
            }
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            var v18 uint64 = 18446744073709551615
            v18 = (^(v18))
            _ = fi15
            _ = v16
            _ = v18
        }
        var v19 float64 = v13
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := GenS2{F0: "", F1: -2147483648}
            v2[""] = "foo bar"
            v22 := GenS2{F0: "\n\t", F1: 0}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi14
        _ = v19
    }
    if (0.0 >= 1e10) {
        fmt.Println(v5, v5)
        fmt.Println(v7, v6, v1)
    } else {
        var v23 uint32 = 2147483648
        v6 += -223
        v24 := GenS1{F0: (v12 ^ 32767), F1: (v1 + (^(v6))), F2: (^(v23))}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        switch "" {
        case "\n\t":
            switch 32767 {
            case 9223372036854775807:
                v12 = 65535
            }
        default:
            v12 = (v12 & (v12 | v12))
        }
        _ = fi25
    }
    if true {
        var v26 int16 = 42
        v13 += v13
        var v27 int = -2147483648
        _ = v26
        _ = v27
    } else {
        delete(v2, "hello")
    }
    var v28 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
