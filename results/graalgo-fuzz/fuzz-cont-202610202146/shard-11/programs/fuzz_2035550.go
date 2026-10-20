package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 bool }
type GenS2 struct { F0 uint; F1 int; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS4 struct { F0 int8; F1 string; F2 uint32 }
type GenS5 struct { F0 int32; F1 int16 }
func genFn0(p0 int64, p1 int8) int {
    var v5 int16 = 902
    v6 := []int{2147483647, -428, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    v1 := []int{127}
    v2 := []uint32{42, 799}
    v3 := make(map[int]bool)
    var v4 int16 = 0
    if true {
        v7 := GenS1{F0: 2147483647, F1: 4294967295, F2: true}
        v8 := GenS2{F0: 255, F1: 42, F2: 1}
        fmt.Println(v4)
        _ = v7
        _ = v8
    } else {
        v4 += (v4 % ((v4) | 1))
        var v9 bool = false
        _ = v9
    }
    v4 = -1
    var v10 int64 = -1
    for fi11 := 0; fi11 < 18; fi11++ {
        v10 = (v10 / (((v10 + v10)) | 1))
        for fi12 := 0; fi12 < 19; fi12++ {
            if false {
                _ = len(v3)
                if (((float32(1.5) - 3.14159) - 3.14159) == ((0.1 - 0.1) / 1.0)) {
                    if ((1 < 0) && (1 > 255)) {
                        v13 := GenS2{F0: 65535, F1: genFn0(v10, 42), F2: 42}
                        v14 := GenS3{F0: 1, F1: 42, F2: (v10 ^ v10)}
                        var v15 uint16 = 880
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        switch (fi12 / (((fi11 + (fi12 & -2147483648))) | 1)) {
                        case 42:
                            v4 -= (v4 >> ((65535) & 63))
                        default:
                            v4 -= (v4 >> ((0) & 63))
                        }
                        _ = len(v1)
                    }
                    if (!("" >= ("hello" + (("hello" + ("\n\t" + "foo bar")) + "a")))) {
                        fmt.Println(v10, fi11)
                    }
                    switch -174 {
                    case 9223372036854775807:
                        var v16 int32 = -97
                        _ = v16
                    default:
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                    }
                }
            }
            _ = fi12
        }
        _ = fi11
    }
    v10 += v10
    v17 := GenS4{F0: -128, F1: "\n\t", F2: 4294967295}
    v10 = 42
    if (!false) {
        var v18 int64 = (v10 ^ -1)
        if len(v1) > 0 {
            v1[0] = 594
        }
        _ = v18
    }
    var v19 int = -752
    var v20 uint16 = 1
    var v21 int16 = v4
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS5{F0: -1, F1: v21}
        var v24 float64 = 1.5
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
