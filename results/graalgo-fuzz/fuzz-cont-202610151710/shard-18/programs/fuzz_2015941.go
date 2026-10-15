package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int64 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int64 = 1
    var v3 bool = false
    v4 := []int{9223372036854775807, 519, -1, 546}
    var v5 float32 = 1e10
    var v6 int16 = 0
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v3)
        _ = fi7
    }
    var v8 uint64 = 457
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS1{F0: (v6 & (v6 % ((v6) | 1))), F1: -927, F2: (v2 / (((v2 | v2)) | 1))}
        v11 := GenS2{F0: 65535}
        var v12 float32 = (v5 * v5)
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 2}
    v3 = v3
    var v14 float64 = (0.0 + ((1.5 - 0.1) / -2.25))
    var v15 int = 1
    var v16 int64 = v2
    v2 += (v16 - int64(v6))
    if len(v4) > 0 {
        v4[0] = v15
    }
    if v3 {
        var v17 uint64 = v8
        switch (v15 & -804) {
        case 32767:
            v18 := GenS3{F0: int32(v15)}
            _ = v18
        case -9223372036854775808:
            v19 := GenS4{F0: 4294967295}
            _ = v19
        case -1:
            var v20 int16 = -661
            _ = v20
        }
        fmt.Println(v16, v2)
        _ = v17
    }
    v2 -= 42
    if (float64(float32(float64((1.5 - (0.1 - (v5 * 0.1)))))) != float64(v5)) {
        if true {
            if v3 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v21 := GenS5{F0: (v8 & (v8 % ((v8) | 1))), F1: v16}
                _ = v21
            } else {
                fmt.Println(v15, v16, v1)
            }
        } else {
            v4 = append(v4, v15)
            for fi22 := 0; fi22 < 1; fi22++ {
                v4 = append(v4, (v15 & int(v1)))
                for fi23 := 0; fi23 < 18; fi23++ {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(fi22, v15, fi22)
                    case "":
                        var v24 uint64 = 574
                        _ = v24
                    default:
                        var v25 uint64 = v8
                        _ = v25
                    }
                    var v26 uint16 = 2
                    _ = fi23
                    _ = v26
                }
                _ = fi22
            }
        }
        fmt.Println(v15, v2)
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
