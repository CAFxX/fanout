package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint8 }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        if (!false) {
            switch fi6 {
            case 981:
                var v7 int64 = 1
                _ = v7
            case 42:
                v8 := []int{(fi6 ^ fi6), (^(fi6))}
                _ = v8
            }
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 uint16 = 255
                _ = fi9
                _ = v10
            }
            var v11 float32 = float32(0.0)
            _ = v11
        }
        v12 := []int{(fi6 + fi6), (^(fi6))}
        _ = fi6
        _ = v12
    }
    var v13 int32 = -836
    fmt.Println(v13)
    _ = v13
    return 255
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]uint32)
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]bool)
    var v5 uint8 = 1
    v14 := []int{0, 0}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if (v5 < (v5 & (v5 + v5))) {
        var v15 uint64 = 42
        _ = v15
    }
    var v16 int8 = 100
    var v17 int32 = 2147483647
    v1 += 4294967295
    v18 := GenS1{F0: (v1 - v1), F1: (^(v5))}
    v19 := GenS2{F0: 65535, F1: v17, F2: 0}
    fmt.Println(v17, v5)
    fmt.Println(v17, v17, v17)
    v16 += v16
    for fi20 := 0; fi20 < 2; fi20++ {
        switch 128 {
        case 42:
            v17 += (v17 / ((int32(v5)) | 1))
        case -977:
            v14 = append(v14, (-(fi20)))
        case -2147483648:
            v21 := GenS3{F0: uint8(v1)}
            _ = v21
        }
        _ = fi20
    }
    var v22 bool = false
    if v22 {
        var v23 uint8 = (v5 * (v5 & 0))
        if len(v14) > 0 {
            v14[0] = 2147483647
        }
        v24 := GenS1{F0: (^(v1)), F1: v5}
        _ = v23
        _ = v24
    } else {
        v25 := GenS4{F0: 0, F1: (v1 << ((4294967295) & 63)), F2: v5}
        _ = v25
    }
    var v26 uint64 = 2
    v22 = false
    var v27 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
