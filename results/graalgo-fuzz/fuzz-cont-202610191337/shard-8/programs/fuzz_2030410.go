package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0(p0 bool) int8 {
    if p0 {
        v6 := GenS1{F0: -143, F1: 7}
        var v7 int32 = -739
        _ = v6
        _ = v7
    }
    var v8 uint32 = 65535
    var v9 float32 = (float32(0.0) * (0.0 + 3.14159))
    var v10 bool = p0
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return -128
}
func genFn1(p0 uint64) uint32 {
    v11 := GenS2{F0: 42, F1: 1}
    switch "hello" {
    case "foo bar":
        var v12 uint32 = 0
        _ = v12
    default:
        var v13 float32 = 1.5
        _ = v13
    }
    _ = p0
    _ = v11
    return 42
}

func main() {
    var v1 bool = false
    var v2 uint32 = 0
    v3 := make(map[uint32]string)
    var v4 int32 = -1
    var v5 uint = 18446744073709551615
    var v14 int8 = -1
    v4 += v4
    for fi15 := 0; fi15 < 18; fi15++ {
        if (240 != 1) {
            var v16 uint = (v5 ^ (v5 ^ 1))
            _ = v16
        } else {
            v17 := []int{}
            _ = v17
        }
        _ = fi15
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := []int{127, fi18, fi18}
        _ = fi18
        _ = v19
    }
    var v20 uint = 338
    var v21 uint = (v5 - 133)
    v22 := GenS1{F0: v4, F1: 884}
    fmt.Println(v14, v1, v20)
    v1 = true
    fmt.Println(v4, v1, v4)
    if false {
        for fi23 := 0; fi23 < 20; fi23++ {
            if (((float32((-2.25 + float64(-2.25))) * float32(0.0)) >= 0.1) && (v1 || (true || v1))) {
                v14 += (v14 / (((v14 | v14)) | 1))
                v4 = (v4 * v4)
            } else {
                v24 := []int{fi23}
                var v25 float32 = 3.14159
                _ = v24
                _ = v25
            }
            _ = fi23
        }
    }
    fmt.Println(v5, v14, v21)
    var v26 float32 = 1.5
    v5 += v20
    v27 := GenS1{F0: 2147483647, F1: 296}
    v28 := GenS1{F0: (v4 - (v4 >> (((v20 >> ((v21) & 63))) & 63))), F1: -270}
    if true {
        var v29 int8 = (v14 + 92)
        var v30 int = int(v4)
        _ = v29
        _ = v30
    } else {
        v31 := GenS3{F0: 127, F1: uint16(v2)}
        for fi32 := 0; fi32 < 8; fi32++ {
            fmt.Println(v5, v2, v26)
            v20 = (v20 << ((2) & 63))
            _ = fi32
        }
        v33 := GenS4{F0: 127, F1: -1, F2: genFn0((1e10 < -2.25))}
        _ = v31
        _ = v33
    }
    v34 := GenS1{F0: v4, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
