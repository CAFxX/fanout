package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    var v1 float32 = 1.5
    v2 := []string{("" + "")}
    v3 := []uint32{654, 2, 4294967295, 415}
    var v4 int8 = -44
    v5 := []int{9223372036854775807, -9223372036854775808, 128}
    v1 = (v1 + -2.25)
    if false {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: "foo bar", F1: fi6, F2: 58}
            if len(v3) > 0 {
                v3[0] = 0
            }
            _ = fi6
            _ = v7
        }
        v1 += float32(3.14159)
        var v8 float64 = 0.0
        _ = v8
    } else {
        v4 -= -92
    }
    v9 := GenS1{F0: "\n\t", F1: -9223372036854775808, F2: 208}
    _ = len(v2)
    switch 128 {
    case 563:
        if false {
            if ((true && ((v4 | v4) < 0)) || true) {
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 int64 = 487
                    v11 = (^(v11))
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: "\n\t", F1: -534, F2: 255}
                _ = v12
            }
            var v13 uint16 = 1
            fmt.Println(v13, v1, v1)
            _ = v13
        }
    case 4294967295:
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint64 = 403
            var v16 float32 = 3.14159
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 int32 = 13
                v1 += (float32(-2.25) * v1)
                _ = fi17
                _ = v18
            }
            _ = fi14
            _ = v15
            _ = v16
        }
    case 445:
        _ = len(v2)
    default:
        var v19 string = "hello"
        _ = v19
    }
    v4 -= -91
    fmt.Println(v4, v4)
    v20 := GenS1{F0: "a", F1: 168, F2: 0}
    if (false && ((v4 / (((v4 + v4)) | 1)) == 42)) {
        if len(v5) > 0 {
            v5[0] = 1
        }
        if (42 < -32768) {
            fmt.Println(v1, v4)
        } else {
            _ = len(v2)
            switch 32768 {
            case 1:
                v1 -= v1
            case 887:
                v21 := GenS2{F0: (v4 | v4), F1: 4294967295}
                _ = v21
            case 148:
                v4 = (v4 / ((v4) | 1))
            default:
                switch 42 {
                case -1:
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v23 := GenS2{F0: 27, F1: 706}
                        fmt.Println(fi22)
                        v24 := GenS2{F0: v4, F1: 4294967295}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                case 128:
                    fmt.Println(v1)
                }
            }
        }
        v25 := GenS3{F0: 0, F1: (^(v4))}
        _ = v25
    } else {
        var v26 float64 = ((0.1 + float64(3.14159)) + -2.25)
        v1 += v1
        _ = v26
    }
    var v27 uint32 = 65535
    v2 = append(v2, ("foo bar" + "hello"))
    var v28 int8 = v4
    if (9223372036854775807 != 9223372036854775807) {
        fmt.Println(v27, v27)
        v27 = v27
        v4 = (v28 + v28)
    }
    v29 := GenS4{F0: 0, F1: 4294967295}
    if (!false) {
        v1 -= -2.25
    } else {
        v5 = append(v5, 4294967295)
        for fi30 := 0; fi30 < 12; fi30++ {
            v1 = (float32(1.5) / ((1e10 - v1) / float32(0.1)))
            _ = fi30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
