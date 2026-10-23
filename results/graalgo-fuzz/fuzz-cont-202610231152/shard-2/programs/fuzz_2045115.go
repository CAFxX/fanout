package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 int8 }
func genFn0() int64 {
    var v5 int64 = 9223372036854775807
    v6 := []int{9223372036854775807, -17, -143}
    _ = v5
    _ = v6
    return 397
}
func genFn1(p0 uint32) int16 {
    v7 := []int{-9223372036854775808, 312, 0}
    p0 -= p0
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 uint64 = 1
    _ = len(v7)
    _ = p0
    _ = v7
    _ = v8
    return 421
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 354
    var v3 bool = false
    var v4 uint16 = 42
    v9 := GenS1{F0: 112}
    for fi10 := 0; fi10 < 13; fi10++ {
        v2 = v2
        if (true && (!v3)) {
            var v11 int = (fi10 * -240)
            v12 := GenS1{F0: 127}
            _ = v11
            _ = v12
        } else {
            v2 -= (^(v2))
        }
        _ = fi10
    }
    if ("foo bar" >= "a") {
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := []int{(fi13 - (fi13 * fi13)), fi13, 127, fi13}
            var v15 int32 = int32(v4)
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 float64 = 0.1
    v17 := []int{}
    var v18 uint = v2
    fmt.Println(v4, v3)
    var v19 int = -9223372036854775808
    fmt.Println(v18)
    v18 -= 255
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 uint32 = 1
        v18 += (^(v2))
        var v22 uint8 = 255
        _ = fi20
        _ = v21
        _ = v22
    }
    v16 = (((v16 / 3.14159) * float64(0.1)) / (-2.25 - 1.5))
    if false {
        switch v19 {
        case 42:
            if (!false) {
                var v23 float32 = (((1.5 * 1.5) / 1.0) - (-2.25 - ((1.5 - float32(v16)) / 1.5)))
                var v24 bool = v3
                v23 = v23
                _ = v23
                _ = v24
            }
        default:
            v18 -= v2
        }
        v2 -= (v18 % ((v18) | 1))
    } else {
        var v25 string = (("hello" + (("" + "a") + "")) + ((("" + "\n\t") + (("hello" + "") + (("" + "a") + ("" + ("foo bar" + ("hello" + "a")))))) + "hello"))
        if (-532 >= 521) {
            if ((32767 == v19) && v3) {
                fmt.Println(v3, v4)
            }
            switch ("foo bar" + (v25 + v25)) {
            case "a":
                v1[203] = 42
            case "foo bar":
                v26 := GenS2{F0: 1, F1: -128, F2: 0}
                _ = v26
            default:
                v18 = (v18 - (v18 << (((v2 & 42)) & 63)))
            }
        } else {
            v19 = v19
            v1[1] = 66
        }
        _ = v25
    }
    if (v3 && (!v3)) {
        v27 := GenS1{F0: 0}
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := GenS1{F0: 255}
        v19 += (fi28 ^ fi28)
        v30 := GenS1{F0: 42}
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 float64 = v16
    var v32 int = v19
    v31 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v9
}
