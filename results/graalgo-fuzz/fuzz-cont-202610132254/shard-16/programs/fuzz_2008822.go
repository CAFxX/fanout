package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int) uint32 {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 950
}
func genFn1(p0 int8, p1 int64, p2 int32) uint16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        p2 += (p2 * p2)
        p0 = p0
        _ = fi7
    }
    switch "\n\t" {
    case "":
        v8 := []int{9223372036854775807}
        _ = v8
    case "a":
        var v9 uint32 = 0
        _ = v9
    default:
        var v10 uint64 = 369
        _ = v10
    }
    p0 = p0
    p1 += (p1 | (p1 << ((255) & 63)))
    _ = p0
    _ = p1
    _ = p2
    return 857
}

func main() {
    var v1 bool = (!false)
    var v2 int = 708
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    v5 := []int{}
    v6 := make(map[uint32]bool)
    var v11 int64 = -1
    v2 = (v2 << ((4294967295) & 63))
    var v12 string = ("hello" + "hello")
    if v1 {
        v13 := GenS1{F0: 453}
        _ = v13
    }
    v5 = append(v5, (^(v2)))
    if v1 {
        v11 = v11
        v14 := GenS1{F0: 0}
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(v12, v12)
            _ = fi15
        }
        _ = v14
    } else {
        v12 = ((v12 + ((v12 + (v12 + "")) + v12)) + v12)
        var v16 int = 1
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS2{F0: 1}
        if (v1 && false) {
            v11 -= v11
            v19 := GenS1{F0: 4294967295}
            _ = len(v3)
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    var v20 uint32 = 422
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 uint8 = 0
        _ = fi21
        _ = v22
    }
    _ = len(v5)
    var v23 uint64 = 711
    var v24 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
