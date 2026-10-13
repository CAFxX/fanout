package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint16, p1 int64, p2 int32) uint {
    p0 = (p0 << ((1) & 63))
    switch int(p0) {
    case -149:
        v6 := []int{32767, 4294967295, 1}
        _ = v6
    case -58:
        var v7 float32 = -2.25
        _ = v7
    default:
        var v8 int32 = int32(p1)
        _ = v8
    }
    if (false || (42 > 37)) {
        fmt.Println(p1)
        fmt.Println(p2)
        v9 := []int{}
        _ = v9
    }
    v10 := []int{128, -633, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 0
}
func genFn1(p0 int64, p1 bool) string {
    fmt.Println(p0)
    var v11 int32 = -1
    _ = p0
    _ = p1
    _ = v11
    return ("" + "\n\t")
}

func main() {
    var v1 bool = (!(128 < 128))
    var v2 uint32 = 42
    var v3 int = -2147483648
    v4 := make(map[uint32]uint32)
    var v5 int16 = 1
    v12 := []int{(v3 ^ (v3 / ((-9223372036854775808) | 1))), (v3 ^ 948), 4294967295}
    v13 := v4[(v2 / ((v2) | 1))]
    fmt.Println(v5, v1)
    if true {
        var v14 bool = (!false)
        _ = v14
    } else {
        if v1 {
            v5 = v5
        } else {
            v15 := GenS1{F0: v5}
            _ = v15
        }
    }
    v16 := GenS1{F0: 32767}
    if v1 {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            v5 += (v5 & v5)
            _ = fi17
        }
    }
    var v18 string = "a"
    var v19 float64 = 1.5
    v2 = (v2 + 2147483648)
    var v20 uint16 = 0
    v2 = 444
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
