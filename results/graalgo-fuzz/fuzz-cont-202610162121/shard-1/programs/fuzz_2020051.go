package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int32; F2 int64 }
func genFn0(p0 int8) uint32 {
    var v4 int32 = -13
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int = fi5
        v7 := GenS1{F0: 781, F1: 908, F2: 2}
        v8 := []int{fi5, (^(v6)), (fi5 + (^(v6)))}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 bool, p1 uint64, p2 uint64) int32 {
    if p0 {
        var v9 bool = false
        v9 = p0
        var v10 int16 = 790
        _ = v10
        _ = v9
    } else {
        v11 := []int{}
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        fmt.Println(p0, p1)
        _ = v11
    }
    v12 := GenS2{F0: (!(-2.25 == (((1.5 / float64(float32((3.14159 - 0.1)))) * 1e10) + (1.5 - -2.25)))), F1: 4294967295, F2: -75}
    p0 = false
    p2 += (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -2147483648
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -680
    var v3 uint32 = 2147483648
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS3{F0: 32768, F1: 128}
        _ = fi13
        _ = v14
    }
    v15 := GenS4{F0: -91}
    for fi16 := 0; fi16 < 15; fi16++ {
        v2 -= genFn1(((618 != uint64(v3)) || true), 867, 250)
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS4{F0: 100}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v19 := GenS5{F0: "\n\t", F1: 485, F2: 9223372036854775807}
    v2 = genFn1((32767 > 0), 842, 691)
    fmt.Println(v2, v2)
    var v20 uint8 = 109
    v20 += 5
    fmt.Println(v2)
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 int = (^(fi21))
        switch (v22 | fi21) {
        case 128:
            v23 := v1[(v22 * fi21)]
            _ = v23
        default:
            v20 = (^(v20))
        }
        var v24 uint8 = (v20 & 2)
        _ = fi21
        _ = v22
        _ = v24
    }
    v2 = (v2 & 884)
    v25 := GenS1{F0: (v3 % (((v3 & v3)) | 1)), F1: 0, F2: 18446744073709551615}
    v2 += (v2 ^ 416)
    fmt.Println(v2, v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
}
