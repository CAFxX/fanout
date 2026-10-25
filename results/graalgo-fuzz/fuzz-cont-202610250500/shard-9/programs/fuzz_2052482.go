package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint }
type GenS3 struct { F0 string; F1 uint; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint32, p1 uint) uint8 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int = fi4
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: -1}
    p0 += uint32(p1)
    var v7 int16 = 382
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 int16, p1 uint64) bool {
    if (!(false || (0 == 42))) {
        var v8 uint64 = (p1 >> ((255) & 63))
        _ = v8
    } else {
        p0 += 214
        v9 := GenS2{F0: 9223372036854775807, F1: 115, F2: 1}
        var v10 bool = false
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 100}
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS3{F0: (("\n\t" + "") + (("" + (("hello" + "\n\t") + "")) + "\n\t")), F1: 65535, F2: 2}
        v14 := []int{}
        _ = len(v14)
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int16 = p0
    for fi16 := 0; fi16 < 14; fi16++ {
        v15 += v15
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v15
    return true
}

func main() {
    var v1 uint64 = 0
    var v2 int8 = 127
    var v3 bool = false
    var v17 uint64 = (^(v1))
    var v18 int16 = 1
    var v19 float32 = float32(-2.25)
    v20 := []int{-2147483648, 0, -497, 140}
    switch "" {
    case "a":
        v21 := GenS4{F0: 599}
        _ = v21
    default:
        var v22 float32 = (-2.25 * 0.1)
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        var v24 int8 = -128
        var v25 int16 = (v18 - v18)
        _ = fi23
        _ = v24
        _ = v25
    }
    v1 += (v1 ^ (v17 & v1))
    var v26 int64 = 1
    v20 = append(v20, 256)
    v20 = append(v20, 586)
    if len(v20) > 0 {
        v20[0] = -497
    }
    switch -1 {
    case -204:
        fmt.Println(v17, v19, v19)
    case -295:
        var v27 uint64 = v1
        _ = v27
    case 542:
        fmt.Println(v17, v19, v1)
    }
    if ((v18 / (((v18 | v18)) | 1)) > 301) {
        v28 := GenS2{F0: v26, F1: 2, F2: 1}
        fmt.Println(v18)
        if ((v18 - (^(v18))) < (v18 * (v18 + v18))) {
            switch 128 {
            case -9223372036854775808:
                if len(v20) > 0 {
                    v20[0] = 0
                }
            case -556:
                v29 := GenS5{F0: 18446744073709551615, F1: 128}
                _ = v29
            case -98:
                for fi30 := 0; fi30 < 17; fi30++ {
                    _ = len(v20)
                    _ = fi30
                }
            default:
                var v31 uint16 = 1
                _ = v31
            }
            v2 = 127
            var v32 uint32 = 756
            _ = v32
        } else {
            v18 = -1
            v33 := GenS2{F0: (^(v26)), F1: 41, F2: 255}
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 int32 = 0
                v36 := GenS5{F0: v1, F1: 0}
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = v33
        }
        _ = v28
    }
    fmt.Println(v17, v26)
    fmt.Println(v2, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
}
