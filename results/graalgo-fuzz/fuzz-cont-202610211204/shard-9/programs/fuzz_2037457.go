package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int16; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int32, p1 uint16) uint {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := []int{32768, (fi4 / ((fi4) | 1)), (fi4 / (((^(fi4))) | 1)), fi4}
        _ = fi4
        _ = v5
    }
    switch int(p1) {
    case 0:
        if false {
            p0 = (^(p0))
            v6 := GenS1{F0: -458, F1: 378}
            var v7 float64 = 0.0
            _ = v6
            _ = v7
        } else {
            v8 := []int{}
            var v9 int16 = -1
            var v10 uint32 = 935
            _ = v10
            _ = v8
            _ = v9
        }
    case 42:
        v11 := GenS2{F0: 1, F1: 969}
        _ = v11
    default:
        v12 := GenS3{F0: ((42 >= uint8(p1)) || (18446744073709551615 >= 66))}
        _ = v12
    }
    v13 := GenS3{F0: (1 != 0)}
    v14 := []int{}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 18446744073709551615
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]int)
    v3 := []string{"", "a"}
    fmt.Println(v1)
    if false {
        var v15 float64 = (v1 * (v1 + 1.5))
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := GenS4{F0: 2147483647, F1: -32768, F2: 42}
                _ = fi17
                _ = v18
            }
            v1 = v15
            delete(v2, 2)
            _ = fi16
        }
        _ = len(v2)
        _ = v15
    } else {
        v19 := GenS2{F0: -9223372036854775808, F1: 2147483647}
        _ = v19
    }
    var v20 int32 = 62
    switch 4294967295 {
    case -617:
        v21 := GenS5{F0: 42}
        _ = v21
    case 9223372036854775807:
        var v22 int16 = 0
        _ = v22
    case -2147483648:
        var v23 int64 = -9223372036854775808
        _ = v23
    default:
        v1 = 0.0
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 float32 = float32((float64(-2.25) - 0.1))
        var v26 int64 = 243
        if (!true) {
            v27 := GenS6{F0: 42, F1: 668}
            v1 = float64((v25 * v25))
            var v28 int16 = 0
            _ = v27
            _ = v28
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    v3 = append(v3, ("hello" + "hello"))
    for fi29 := 0; fi29 < 4; fi29++ {
        v1 = v1
        _ = fi29
    }
    v1 = 3.14159
    var v30 float32 = 1e10
    v1 += 3.14159
    if (65535 >= 255) {
        v31 := GenS3{F0: true}
        v32 := v2[0]
        _ = v31
        _ = v32
    } else {
        fmt.Println(v1, v20, v20)
        switch -9223372036854775808 {
        case -481:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -570:
            v2[0] = 4294967295
        default:
            var v33 string = "\n\t"
            _ = v33
        }
        v20 -= -1
    }
    v2[65535] = 2147483647
    v20 = (v20 >> ((genFn0(v20, 42)) & 63))
    fmt.Println(v30)
    if false {
        var v34 uint16 = 0
        var v35 int8 = 1
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
}
