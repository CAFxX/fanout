package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int64; F1 int32; F2 int32 }
func genFn0() uint16 {
    v7 := GenS1{F0: 100, F1: -467}
    var v8 float32 = 0.1
    var v9 bool = (!false)
    v9 = v9
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint8, p1 string) int {
    p0 = p0
    var v10 string = p1
    _ = p0
    _ = p1
    _ = v10
    return -1
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]bool)
    var v5 int16 = -220
    var v6 bool = true
    var v11 int64 = 827
    var v12 string = ("foo bar" + "foo bar")
    switch int(v3) {
    case -790:
        if true {
            v12 = ("" + ("foo bar" + "foo bar"))
            for fi13 := 0; fi13 < 7; fi13++ {
                v11 += (v11 * int64(v5))
                v14 := GenS2{F0: v11, F1: 525, F2: 87}
                for fi15 := 0; fi15 < 5; fi15++ {
                    fmt.Println(v5, v5, v11)
                    _ = fi15
                }
                _ = fi13
                _ = v14
            }
        }
    case 32768:
        v16 := GenS2{F0: v11, F1: 0, F2: int32(v11)}
        _ = v16
    case -580:
        v5 += v5
    }
    var v17 uint = 491
    if false {
        if v6 {
            v3 = v3
        }
        v12 = "hello"
    }
    if (42 > 966) {
        var v18 float32 = (1e10 + (float32(float64(-2.25)) + 1.5))
        v19 := v2[32767]
        v4[1] = false
        _ = v18
        _ = v19
    } else {
        v11 = v11
    }
    v20 := GenS1{F0: -93, F1: -1}
    var v21 float64 = 1e10
    v22 := GenS2{F0: 516, F1: 707, F2: 42}
    var v23 int16 = (v5 + -409)
    var v24 int32 = 0
    if (v17 != v17) {
        var v25 int32 = v24
        _ = v25
    }
    var v26 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
