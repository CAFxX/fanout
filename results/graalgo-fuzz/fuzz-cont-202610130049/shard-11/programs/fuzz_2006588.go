package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 bool) bool {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: 42, F1: "a"}
        for fi6 := 0; fi6 < 20; fi6++ {
            if (127 > 41) {
                p0 = p0
                p0 = (0 > 2147483647)
                switch "\n\t" {
                case "a":
                    v7 := GenS2{F0: 1, F1: ((969 < 0) && true)}
                    _ = v7
                default:
                    v8 := []int{(fi4 ^ (fi4 / ((-1) | 1)))}
                    _ = v8
                }
                v9 := []int{fi6}
                p0 = p0
                _ = v9
            }
            p0 = p0
            var v10 float32 = (float32(3.14159) + float32(-2.25))
            _ = fi6
            _ = v10
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := []int{(fi11 & (fi4 / ((fi11) | 1))), fi4}
            _ = len(v12)
            v13 := GenS3{F0: 917}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi4
        _ = v5
    }
    if p0 {
        v14 := []int{1, 2147483647}
        _ = v14
    } else {
        switch -9223372036854775808 {
        case 32768:
            v15 := GenS2{F0: 969, F1: (-32768 != 1)}
            _ = v15
        case 653:
            var v16 int = 32768
            _ = v16
        }
        v17 := GenS1{F0: 968, F1: "hello"}
        p0 = (!p0)
        _ = v17
    }
    _ = p0
    return (0 != 0)
}
func genFn1(p0 int16, p1 uint8, p2 uint16) int8 {
    var v18 uint = 100
    var v19 int8 = 59
    var v20 uint64 = 42
    v21 := []int{-2147483648, 2147483647, 749}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    _ = v21
    return v19
}

func main() {
    v1 := []int{-117, -215, 0}
    v2 := make(map[string]bool)
    var v3 bool = ((1.5 - (((1.5 - ((1.5 - float64(1.5)) - 0.0)) * 1e10) * 0.0)) != ((1.5 / 1e10) + -2.25))
    var v22 int32 = -1
    v23 := GenS4{F0: genFn1(32767, 126, 792), F1: false}
    if len(v1) > 0 {
        v1[0] = -1
    }
    var v24 uint = 0
    var v25 uint = v24
    var v26 uint32 = 894
    v27 := GenS5{F0: 796, F1: 0}
    fmt.Println(v26, v3)
    _ = len(v2)
    v28 := GenS6{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
