package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0(p0 bool, p1 int32, p2 int) uint16 {
    var v7 uint64 = 1
    p0 = (-128 < -1)
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: -32, F1: "a"}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = -594
    var v3 uint64 = 1
    v4 := make(map[string]string)
    var v5 uint8 = 127
    var v6 int64 = 1
    v10 := []int{-42, 2147483647}
    fmt.Println(v2)
    for fi11 := 0; fi11 < 10; fi11++ {
        if (0.1 > (((0.0 + 3.14159) * -2.25) + 0.0)) {
            fmt.Println(v2, v6)
        } else {
            fmt.Println(v6, fi11)
            for fi12 := 0; fi12 < 13; fi12++ {
                if (true && true) {
                    var v13 int32 = 0
                    fmt.Println(fi12)
                    var v14 int16 = -1
                    _ = v13
                    _ = v14
                }
                _ = fi12
            }
        }
        v1 += -1
        _ = fi11
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint64 = 4294967295
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 7; fi18++ {
            v3 = 532
            v6 = v2
            var v19 uint8 = 107
            v1 = v2
            _ = fi18
            _ = v19
        }
        _ = fi15
    }
    fmt.Println(v6, v3, v3)
    v3 = (v3 & (^(v3)))
    v4[""] = "foo bar"
    if ((v3 & (v3 * (v3 - (v3 | (v3 / ((v3) | 1)))))) == (v3 & v3)) {
        if false {
            var v20 uint32 = 2147483647
            var v21 int8 = 0
            if len(v10) > 0 {
                v10[0] = -2147483648
            }
            _ = v20
            _ = v21
        }
        fmt.Println(v5, v1)
        v22 := GenS2{F0: 65535, F1: v3}
        _ = v22
    } else {
        var v23 string = ""
        switch 4294967295 {
        case 127:
            _ = len(v4)
        case 4294967295:
            var v24 int16 = -1
            _ = v24
        case -300:
            var v25 uint16 = uint16(v5)
            _ = v25
        }
        var v26 float64 = -2.25
        _ = v23
        _ = v26
    }
    v27 := v4[("a" + "foo bar")]
    v3 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
