package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 int32 }
type GenS4 struct { F0 int16; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int32 }
func genFn0(p0 uint, p1 uint8, p2 int) int64 {
    fmt.Println(p0, p1, p2)
    v6 := GenS1{F0: false}
    var v7 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (v7 * v7)
}
func genFn1(p0 uint16, p1 int8, p2 int16) int16 {
    v8 := []int{691, 0, 11, 1}
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS2{F0: 915, F1: 208}
        p2 = p2
        _ = fi9
        _ = v10
    }
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p2
}

func main() {
    var v1 int16 = -579
    var v2 uint8 = 2
    var v3 float64 = (0.0 - (((0.0 / 0.1) / (0.0 - (float64((-2.25 + (float32(3.14159) / 1.0))) - 0.0))) - -2.25))
    v4 := make(map[string]bool)
    v5 := []int{}
    var v11 int16 = 32767
    if ((true || (9223372036854775807 > -127)) || true) {
        _ = len(v4)
        v12 := GenS3{F0: 0, F1: ("hello" + "\n\t"), F2: -892}
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v11 = genFn1(255, 54, v11)
        _ = fi13
    }
    var v14 uint32 = 2147483648
    for fi15 := 0; fi15 < 14; fi15++ {
        if (152 != genFn0(42, (v2 / ((v2) | 1)), (^(fi15)))) {
            var v16 float32 = 1e10
            var v17 bool = (true && ((-1 < -9223372036854775808) || (false && false)))
            fmt.Println(v1, v3, v1)
            _ = v16
            _ = v17
        } else {
            v1 = v1
            switch (-(fi15)) {
            case 42:
                var v18 int32 = 42
                _ = v18
            case 4294967295:
                fmt.Println(v14, v1, fi15)
            case -390:
                if true {
                    v19 := GenS3{F0: v11, F1: "\n\t", F2: -994}
                    _ = v19
                }
            }
        }
        var v20 uint = 4294967295
        _ = fi15
        _ = v20
    }
    _ = len(v5)
    v21 := GenS4{F0: -1, F1: (v1 % ((v1) | 1)), F2: 2}
    v2 += (^(v2))
    v22 := GenS5{F0: 2, F1: uint64(v2), F2: -742}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
