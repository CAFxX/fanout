package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "\n\t"
    var v3 uint8 = 128
    v4 := []uint32{}
    var v5 uint32 = 4294967295
    v6 := GenS1{F0: -1}
    v7 := GenS1{F0: -1}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := v1[249]
        v3 = (v3 * 36)
        var v10 float64 = ((0.1 - (3.14159 + 0.1)) * float64(0.0))
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int16 = 0
    fmt.Println(v11, v3, v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (-9223372036854775808 <= 1) {
        v12 := GenS2{F0: 82, F1: true, F2: 0}
        v11 += (v11 - (-(v11)))
        if true {
            var v13 int = 128
            _ = v13
        }
        _ = v12
    } else {
        delete(v1, -1)
        var v14 float64 = 1e10
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS2{F0: 578, F1: false, F2: -1}
        for fi17 := 0; fi17 < 4; fi17++ {
            v5 = v5
            v4 = append(v4, 245)
            var v18 int16 = (v11 / (((v11 * v11)) | 1))
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            _ = len(v1)
            _ = fi19
        }
        _ = fi15
        _ = v16
    }
    if false {
        v20 := GenS1{F0: 0}
        _ = v20
    } else {
        var v21 bool = (!false)
        for fi22 := 0; fi22 < 17; fi22++ {
            v21 = (2 <= 0)
            _ = fi22
        }
        fmt.Println(v3, v2)
        _ = v21
    }
    var v23 uint8 = v3
    v3 = v3
    if false {
        v11 += v11
        v24 := v1[32767]
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := GenS3{F0: -124, F1: -167}
            switch "" {
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = v24
                }
            default:
                var v27 string = v2
                _ = v27
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    var v28 int32 = -743
    v29 := GenS4{F0: 0}
    v11 -= (-(v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
