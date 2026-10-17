package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int32 }
func genFn0() uint {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int8 = 107
        v9 := []int{(-(fi7)), fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        v10 := []int{-541, 128, 127, 874}
        _ = v10
    }
    var v11 uint = 688
    _ = v11
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 0
    var v3 int16 = 32767
    var v4 float32 = 0.0
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 685
    switch v2 {
    case 4294967295:
        v4 = v4
    case 127:
        v12 := []int{(v2 / ((454) | 1)), (v2 + (v2 << ((648) & 63))), (v2 + 32768)}
        _ = v12
    default:
        var v13 uint32 = (v6 + v6)
        _ = v13
    }
    v3 -= (v3 - v3)
    var v14 uint16 = 1
    v15 := GenS1{F0: 42, F1: 0}
    var v16 float32 = float32(((0.1 * -2.25) * float64((-2.25 + 1.5))))
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int32 = -1
        var v19 int64 = -9223372036854775808
        _ = fi17
        _ = v18
        _ = v19
    }
    switch int(v14) {
    case 32767:
        for fi20 := 0; fi20 < 19; fi20++ {
            v6 += 4294967295
            _ = fi20
        }
    case 1:
        delete(v1, (v2 << ((18446744073709551615) & 63)))
    case 824:
        if (v14 >= v14) {
            v21 := GenS2{F0: (v2 << ((242) & 63)), F1: (v14 | (v14 * (v14 ^ v14)))}
            var v22 string = "hello"
            v23 := GenS3{F0: 0}
            _ = v21
            _ = v22
            _ = v23
        }
    default:
        var v24 uint64 = uint64(v14)
        _ = v24
    }
    v4 -= float32(1e10)
    v6 -= v6
    if (v4 != 0.1) {
        v14 = (^(v14))
        v25 := GenS4{F0: 18446744073709551615, F1: (v6 << ((561) & 63)), F2: -1}
        _ = v25
    } else {
        v3 = -294
    }
    v3 -= v3
    switch "hello" {
    case "foo bar":
        v14 = 0
    case "":
        delete(v5, 0)
    }
    var v26 float32 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
