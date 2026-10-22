package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint16; F1 uint64; F2 int }
func genFn0(p0 int32, p1 int8, p2 int16) uint16 {
    p0 = p0
    var v7 int32 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 uint8, p1 int32, p2 int8) string {
    v8 := GenS1{F0: p2, F1: 12, F2: 21}
    v9 := []int{32767, 404, 32767, -9223372036854775808}
    var v10 float32 = (3.14159 / (0.1 + 0.1))
    v10 = v10
    p1 = (p1 % (((p1 * -643)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]string)
    var v3 uint32 = 2147483648
    var v4 string = "a"
    v5 := make(map[string]bool)
    v6 := []int{-24}
    var v11 int8 = 100
    v12 := GenS2{F0: 246, F1: 1}
    v11 -= 1
    v6 = append(v6, -139)
    v13 := GenS2{F0: 1, F1: 4294967295}
    var v14 float32 = float32(1e10)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if false {
        v15 := GenS3{F0: false, F1: v3, F2: 1}
        if ((1e10 == 3.14159) || true) {
            v16 := GenS4{F0: 1, F1: "\n\t"}
            _ = v16
        }
        _ = v15
    }
    v17 := GenS5{F0: 32767, F1: 548, F2: 32768}
    switch 127 {
    case 127:
        for fi18 := 0; fi18 < 5; fi18++ {
            v14 = ((v14 - ((-2.25 / (float32(1.5) - 0.0)) + float32((0.0 * 1e10)))) / v14)
            _ = fi18
        }
    default:
        _ = len(v2)
    }
    v1 += 0.0
    var v19 int64 = -779
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
