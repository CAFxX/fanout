package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 int32) int16 {
    p1 += (p1 | (p1 << ((615) & 63)))
    switch -655 {
    case 127:
        p0 = "a"
    default:
        var v5 float64 = 0.0
        _ = v5
    }
    if false {
        var v6 bool = true
        _ = v6
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint8 = 2
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int16 = -1
            var v11 float64 = 0.0
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            if false {
                fmt.Println(p0, p0)
            }
            v12 := GenS1{F0: 158, F1: 998, F2: 1}
            v13 := GenS1{F0: 2147483648, F1: 117, F2: 65535}
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: 462, F1: 1, F2: 131}
            p1 = 0
            v15 := []int{74}
            _ = v14
            _ = v15
        }
    }
    p1 += (^(p1))
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint = 42
    var v2 uint64 = 2
    v3 := make(map[uint32]string)
    var v4 uint = (v1 >> ((v1) & 63))
    fmt.Println(v2)
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int16 = 32767
        _ = fi16
        _ = v17
    }
    var v18 uint = v1
    fmt.Println(v4)
    fmt.Println(v2, v4, v1)
    v19 := GenS1{F0: uint32(v1), F1: 470, F2: 42}
    v20 := []int{4294967295}
    v21 := GenS2{F0: 2}
    v1 -= v18
    fmt.Println(v4)
    v18 -= (v4 % (((v4 | 930)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
