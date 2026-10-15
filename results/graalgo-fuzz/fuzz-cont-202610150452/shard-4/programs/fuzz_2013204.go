package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint }
func genFn0(p0 uint64, p1 int16, p2 uint64) uint32 {
    if (1 <= 0) {
        v7 := []int{32767, 128, 370, 2147483647}
        v7 = append(v7, 32767)
        _ = v7
    }
    var v8 string = "a"
    p1 += (^(p1))
    if (42 <= 2147483647) {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: 2147483647, F1: 355, F2: 32767}
            p1 = -495
            _ = fi9
            _ = v10
        }
        var v11 uint = 762
        p1 = (p1 & (p1 & (p1 - 32767)))
        _ = v11
    } else {
        v12 := GenS2{F0: p1}
        _ = v12
    }
    if (0.1 == 3.14159) {
        var v13 int32 = -675
        _ = v13
    } else {
        if (2 != 255) {
            if (0 == 65535) {
                switch 127 {
                case 978:
                    fmt.Println(p1, p0, p2)
                case 570:
                    p1 += (p1 % (((p1 & p1)) | 1))
                case 0:
                    v14 := GenS3{F0: (p1 + (p1 / ((42) | 1))), F1: (false && (1 != 42)), F2: 4294967295}
                    _ = v14
                default:
                    var v15 uint32 = 523
                    _ = v15
                }
            }
            if true {
                v16 := GenS4{F0: int32(p2), F1: 0}
                fmt.Println(p1)
                fmt.Println(p2, v8)
                _ = v16
            } else {
                v17 := GenS4{F0: 766, F1: 65535}
                v8 = v8
                _ = v17
            }
            v18 := []int{}
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 65535
}

func main() {
    v1 := []int{-679, 513, 128, 128}
    var v2 uint64 = 18446744073709551615
    var v3 float32 = -2.25
    v4 := []string{"foo bar", "hello"}
    v5 := []uint32{42, 2147483647}
    var v6 int16 = 530
    v3 = 0.1
    v6 = 42
    if true {
        v2 = v2
        var v19 int32 = 1
        _ = v19
    }
    v3 = (-2.25 + v3)
    var v20 float64 = (float64(0.0) / float64((0.0 - float32(float64(0.1)))))
    var v21 bool = false
    v22 := GenS4{F0: 580, F1: 2147483648}
    v21 = v21
    var v23 uint16 = 255
    var v24 uint32 = 65535
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS5{F0: v2, F1: 709, F2: 42}
        var v27 uint8 = 0
        _ = fi25
        _ = v26
        _ = v27
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    v23 += 2
    var v28 int64 = 9223372036854775807
    var v29 int8 = 1
    var v30 uint16 = (^(v23))
    fmt.Println(v2, v6)
    for fi31 := 0; fi31 < 7; fi31++ {
        v23 = 32767
        var v32 int8 = 1
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
