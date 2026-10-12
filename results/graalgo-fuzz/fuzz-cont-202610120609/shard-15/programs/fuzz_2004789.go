package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0() int32 {
    switch 542 {
    case 128:
        v7 := GenS1{F0: 63, F1: 344, F2: -1}
        _ = v7
    case -751:
        v8 := []int{271, 38, 127, 559}
        _ = v8
    case 4294967295:
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 float32 = 3.14159
            _ = fi9
            _ = v10
        }
    }
    var v11 uint = 828
    var v12 uint32 = 1
    fmt.Println(v12)
    _ = v11
    _ = v12
    return 2147483647
}
func genFn1() int64 {
    var v13 uint64 = 71
    var v14 uint64 = (v13 & 4294967295)
    v13 -= (v13 * (v13 & v14))
    _ = v13
    _ = v14
    return 39
}

func main() {
    var v1 uint = 65535
    var v2 int8 = 0
    var v3 string = "hello"
    var v4 uint8 = 0
    var v5 int8 = (-(v2))
    v6 := make(map[uint32]string)
    v5 = v5
    switch v3 {
    case "":
        switch "\n\t" {
        case "a":
            v3 = v3
        }
    case "a":
        v5 += (^(v5))
    case "hello":
        switch 9223372036854775807 {
        case -766:
            var v15 int32 = 27
            _ = v15
        }
    }
    if (3.14159 > 3.14159) {
        var v16 float32 = -2.25
        _ = v16
    }
    fmt.Println(v1, v5, v5)
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int8 = v5
            v1 = 572
            _ = len(v6)
            _ = fi18
            _ = v19
        }
        var v20 uint64 = 2
        _ = fi17
        _ = v20
    }
    v6[1] = v3
    v21 := v6[317]
    fmt.Println(v5, v1)
    for fi22 := 0; fi22 < 2; fi22++ {
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 int = 762
            v25 := []int{}
            v2 = (v2 >> (((v1 * 327)) & 63))
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 bool = true
        v28 := v6[2]
        var v29 int32 = genFn0()
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    v4 += (v4 | (^(v4)))
    v30 := GenS2{F0: 42, F1: 2147483647}
    if (false && false) {
        var v31 int32 = 42
        fmt.Println(v21)
        _ = v31
    } else {
        var v32 string = v21
        v4 -= (v4 & (v4 / (((v4 + (v4 ^ v4))) | 1)))
        fmt.Println(v32)
        _ = v32
    }
    var v33 int32 = 2147483647
    v34 := GenS2{F0: 42, F1: 1}
    var v35 int16 = 1000
    switch -9223372036854775808 {
    case -1:
        v36 := GenS2{F0: 42, F1: 0}
        _ = v36
    case 4294967295:
        var v37 int64 = int64(v2)
        _ = v37
    default:
        v5 -= (v5 % (((v2 / ((int8(v33)) | 1))) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
