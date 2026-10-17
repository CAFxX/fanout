package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint) int32 {
    fmt.Println(p0)
    switch 757 {
    case 4294967295:
        var v7 uint32 = 2
        _ = v7
    case 753:
        if true {
            var v8 uint64 = 18446744073709551615
            var v9 bool = true
            var v10 float64 = ((1.5 + 1e10) - (0.1 * 1e10))
            _ = v10
            _ = v8
            _ = v9
        }
    default:
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(fi11, p0)
            var v12 uint32 = 2
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    p0 -= 397
    if false {
        v14 := []int{-9223372036854775808, -474, 2147483647}
        _ = v14
    }
    _ = p0
    return -2147483648
}

func main() {
    var v1 int16 = -199
    v2 := make(map[string]bool)
    var v3 uint = 0
    var v4 int32 = -911
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]string)
    v15 := []int{42, -385}
    switch 42 {
    case 431:
        if false {
            var v16 int16 = 799
            _ = v16
        }
    case 32768:
        var v17 int64 = -1
        _ = v17
    }
    var v18 uint16 = 876
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 uint32 = 0
        v21 := GenS1{F0: 42, F1: uint64(v3)}
        var v22 int16 = 32767
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: 776}
    var v24 uint64 = 2
    var v25 int = -231
    switch ("foo bar" + ("hello" + "foo bar")) {
    case "":
        v3 += (v3 + (v3 << ((v3) & 63)))
    case "a":
        var v26 int16 = (v1 / (((^(v1))) | 1))
        _ = v26
    case "\n\t":
        for fi27 := 0; fi27 < 15; fi27++ {
            fmt.Println(v24, fi27, v18)
            v1 -= v1
            var v28 int16 = -1
            _ = fi27
            _ = v28
        }
    }
    v4 += v4
    var v29 float32 = 0.0
    for fi30 := 0; fi30 < 17; fi30++ {
        v15 = append(v15, int(v18))
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 int32 = v4
            v18 -= (v18 * (v18 % ((v18) | 1)))
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    if false {
        var v33 float32 = -2.25
        _ = v33
    } else {
        _ = len(v15)
    }
    v29 = v29
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
