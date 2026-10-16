package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0(p0 string) string {
    switch 483 {
    case 798:
        var v7 float32 = float32(0.1)
        _ = v7
    case -977:
        if ((-2147483648 > 748) && ((1e10 + (0.1 - 0.1)) > 0.0)) {
            p0 = ""
        }
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int8 = 127
        v9 += v9
        _ = fi8
        _ = v9
    }
    p0 = p0
    for fi10 := 0; fi10 < 12; fi10++ {
        if ((-9223372036854775808 <= -9223372036854775808) && false) {
            var v11 float64 = 0.1
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 string = p0
            v13 = p0
            var v14 float32 = 0.1
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    p0 = p0
    _ = p0
    return p0
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 2
    v3 := make(map[uint32]uint32)
    var v4 bool = false
    var v5 int32 = 463
    var v6 int = -311
    fmt.Println(v4, v1)
    fmt.Println(v1, v5)
    v2 -= v2
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int32 = 957
            if (!((("\n\t" + "foo bar") + ((("foo bar" + "a") + (("\n\t" + "foo bar") + ("foo bar" + "a"))) + ("hello" + "hello"))) <= ((("a" + "a") + "\n\t") + "\n\t"))) {
                v18 := GenS1{F0: 139, F1: (v6 / ((fi15) | 1)), F2: 348}
                var v19 float32 = 1.5
                if v4 {
                    var v20 uint16 = 73
                    fmt.Println(fi16)
                    _ = len(v3)
                    _ = v20
                }
                _ = v18
                _ = v19
            } else {
                v21 := GenS2{F0: 128, F1: 128}
                v4 = true
                v22 := GenS1{F0: 73, F1: (v6 << ((2) & 63)), F2: 42}
                _ = v21
                _ = v22
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v23 := []int{}
    fmt.Println(v6, v1)
    var v24 uint32 = v2
    var v25 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
