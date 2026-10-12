package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16) int32 {
    if ((100 < -128) && false) {
        switch ("a" + "") {
        case "hello":
            var v5 string = ""
            _ = v5
        case "a":
            fmt.Println(p0)
        case "foo bar":
            for fi6 := 0; fi6 < 16; fi6++ {
                p0 = (p0 - -1)
                _ = fi6
            }
        default:
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := []int{(fi7 >> ((2) & 63)), fi7, (-(fi7)), 127}
                _ = fi7
                _ = v8
            }
        }
        var v9 uint16 = 42
        _ = v9
    } else {
        v10 := []int{}
        _ = v10
    }
    p0 += -975
    _ = p0
    return -445
}
func genFn1(p0 int32, p1 uint8, p2 int16) int64 {
    var v11 uint32 = 389
    var v12 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{65535}
    var v3 uint16 = 296
    var v4 uint32 = 545
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += 859
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 uint16 = 670
        _ = fi13
        _ = v14
    }
    if true {
        var v15 float32 = -2.25
        var v16 uint16 = v3
        fmt.Println(v15, v3, v15)
        _ = v15
        _ = v16
    } else {
        fmt.Println(v3, v4)
        v3 = v3
        var v17 float64 = -2.25
        _ = v17
    }
    v18 := GenS1{F0: 39}
    var v19 int = 42
    delete(v1, v19)
    var v20 uint32 = v4
    if (!(-1 <= 370)) {
        fmt.Println(v3, v3)
        _ = len(v1)
        var v21 float32 = float32(1.5)
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
