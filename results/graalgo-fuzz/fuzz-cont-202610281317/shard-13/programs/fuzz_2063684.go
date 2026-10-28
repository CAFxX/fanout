package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 string, p1 uint8) uint32 {
    var v6 int64 = 1
    p0 = p0
    fmt.Println(p0, v6, p0)
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 1
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v5 string = (("foo bar" + "\n\t") + "\n\t")
    v5 = "hello"
    if true {
        fmt.Println(v5, v2)
        var v7 int = -2147483648
        _ = v7
    }
    var v8 uint = 902
    switch v5 {
    case "\n\t":
        switch 185 {
        case -1:
            v2 += v2
        case 128:
            v9 := GenS1{F0: 65535, F1: v2, F2: 32767}
            _ = v9
        case -941:
            var v10 uint = v8
            _ = v10
        }
    default:
        v11 := GenS2{F0: 128, F1: 42, F2: 255}
        _ = v11
    }
    v12 := []int{}
    if true {
        v12 = append(v12, 2147483647)
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS3{F0: -9223372036854775808}
            v15 := GenS2{F0: int(v8), F1: 1, F2: 2}
            v2 -= v2
            _ = fi13
            _ = v14
            _ = v15
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    } else {
        for fi16 := 0; fi16 < 10; fi16++ {
            v2 = (v2 * v2)
            v17 := GenS1{F0: 86, F1: v2, F2: 737}
            _ = fi16
            _ = v17
        }
        var v18 string = "\n\t"
        fmt.Println(v18, v5)
        _ = v18
    }
    var v19 string = "\n\t"
    var v20 uint32 = uint32(v8)
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint32 = v20
        v2 -= 1
        v23 := GenS3{F0: -1}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
