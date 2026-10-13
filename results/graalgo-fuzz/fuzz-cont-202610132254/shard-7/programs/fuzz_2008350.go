package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint32 }
func genFn0(p0 uint32, p1 uint) string {
    if false {
        fmt.Println(p1, p0)
        p0 -= uint32(p1)
    }
    var v6 uint = 4294967295
    p0 = 0
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{(-(fi7)), (-(fi7)), fi7, 0}
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int8 = 1
            _ = fi9
            _ = v10
        }
        var v11 float32 = 1e10
        _ = fi7
        _ = v11
        _ = v8
    }
    if (!true) {
        v6 = (v6 * 65535)
        v12 := GenS1{F0: 622, F1: 4294967295, F2: (^(p0))}
        _ = v12
    } else {
        var v13 uint16 = 2
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    return "hello"
}
func genFn1(p0 int) int8 {
    v14 := []int{(^(p0)), (p0 + p0)}
    fmt.Println(p0)
    _ = p0
    _ = v14
    return -1
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{2147483647}
    var v3 uint64 = 2
    v4 := []string{}
    var v5 uint8 = 128
    if true {
        var v15 int = 32767
        fmt.Println(v15)
        _ = v15
    }
    fmt.Println(v3, v5)
    v2 = append(v2, 793)
    var v16 int8 = 105
    var v17 string = ("" + ("a" + ""))
    var v18 int8 = 1
    v19 := v1[((v17 + v17) + v17)]
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 string = "a"
        var v22 int = (fi20 / ((654) | 1))
        if true {
            v1["hello"] = (("a" + v21) + "a")
            if (255 < 229) {
                var v23 float64 = 1e10
                v22 += (v22 / ((fi20) | 1))
                _ = v23
            }
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    switch v17 {
    case "foo bar":
        var v24 int32 = 0
        _ = v24
    case "\n\t":
        fmt.Println(v18)
    case "a":
        v25 := GenS1{F0: uint16(v3), F1: 869, F2: 2147483648}
        _ = v25
    default:
        v18 += (-(v18))
    }
    v26 := GenS1{F0: 255, F1: 2, F2: 0}
    var v27 int = 32768
    var v28 int = (v27 + v27)
    if false {
        switch v17 {
        case "foo bar":
            v2 = append(v2, 65535)
        case "a":
            v2 = append(v2, 2)
        case "\n\t":
            v1["a"] = v19
        default:
            v28 = 0
        }
        v29 := GenS1{F0: 2, F1: uint32(v3), F2: 0}
        _ = v29
    }
    fmt.Println(v27)
    var v30 string = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
