package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int8, p1 uint16) int64 {
    v6 := GenS1{F0: p0}
    var v7 int64 = -9223372036854775808
    var v8 int64 = v7
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 uint16, p1 uint, p2 uint64) int16 {
    var v9 uint16 = p0
    fmt.Println(p2)
    v10 := GenS2{F0: (^(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 788
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (0.1 / (1.5 + -2.25))
    var v3 int32 = 364
    v4 := make(map[string]uint32)
    var v5 string = ""
    fmt.Println(v5, v5)
    var v11 bool = true
    v12 := []int{32767, 128}
    var v13 int32 = (^(v3))
    v14 := GenS2{F0: 18446744073709551615}
    v15 := GenS3{F0: 1}
    for fi16 := 0; fi16 < 10; fi16++ {
        v12 = append(v12, (fi16 % (((fi16 % (((^(fi16))) | 1))) | 1)))
        v5 = ((v5 + v5) + v5)
        _ = fi16
    }
    var v17 uint8 = 188
    for fi18 := 0; fi18 < 9; fi18++ {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi18
    }
    v2 = (-2.25 * v2)
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v17)
        v3 = (v13 ^ v13)
        _ = fi19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        switch fi20 {
        case 495:
            var v21 uint64 = 4294967295
            _ = v21
        default:
            delete(v4, v5)
        }
        _ = len(v12)
        v11 = false
        _ = fi20
    }
    v11 = v11
    v22 := GenS1{F0: 100}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
