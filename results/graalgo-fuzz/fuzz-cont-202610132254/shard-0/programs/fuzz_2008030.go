package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 int; F2 int64 }
func genFn0(p0 uint8, p1 int8, p2 uint16) uint32 {
    v6 := GenS1{F0: -337, F1: -32768}
    fmt.Println(p1, p0)
    var v7 bool = (626 > 365)
    var v8 string = ""
    if (932 < 2147483648) {
        switch 2147483647 {
        case 1:
            v9 := GenS2{F0: 847}
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int32) uint8 {
    fmt.Println(p0)
    var v10 uint32 = 102
    _ = p0
    _ = v10
    return 205
}

func main() {
    var v1 uint = 473
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    var v5 uint = 1
    for fi11 := 0; fi11 < 15; fi11++ {
        v5 = (v1 << ((217) & 63))
        _ = fi11
    }
    v1 += v1
    v12 := GenS1{F0: -1, F1: 42}
    var v13 uint64 = 1
    var v14 uint8 = 128
    v13 += v13
    _ = len(v2)
    v15 := []int{32767, -2147483648, 2147483647}
    v1 += 42
    v5 = 1
    fmt.Println(v5)
    switch ("" + "foo bar") {
    case "a":
        v14 = genFn1(1)
    default:
        v1 -= (^(v5))
    }
    v14 -= 127
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint16 = 1
        _ = fi16
        _ = v17
    }
    var v18 float64 = 1e10
    for fi19 := 0; fi19 < 8; fi19++ {
        v1 = (v1 << ((672) & 63))
        _ = fi19
    }
    v20 := GenS2{F0: v13}
    v21 := GenS3{F0: -961, F1: 1, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
