package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint8; F2 bool }
type GenS4 struct { F0 int32; F1 int32 }
func genFn0() uint8 {
    var v5 int32 = 0
    if true {
        v5 = v5
        var v6 int = 75
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint32 = 42
            fmt.Println(v6, fi7)
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    _ = v5
    return 168
}
func genFn1(p0 int64, p1 uint16, p2 uint16) int16 {
    v9 := GenS1{F0: (("" + ("foo bar" + ("\n\t" + ""))) + ("" + "hello")), F1: true}
    p1 = (p2 << ((23) & 63))
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS2{F0: 498, F1: 1}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 255
    var v3 uint8 = (v2 - v2)
    v4 := []uint32{2, 853, uint32(v3), 2147483648}
    fmt.Println(v3, v2)
    var v12 uint = 1
    v12 += v12
    switch (("" + ("\n\t" + "foo bar")) + (("hello" + "a") + "a")) {
    case "\n\t":
        if false {
            var v13 uint32 = 0
            _ = v13
        } else {
            if false {
                _ = len(v4)
                v14 := GenS3{F0: 1, F1: 1, F2: true}
                _ = v14
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 int64 = -1
                v16 += (-(v16))
                _ = fi15
                _ = v16
            }
        }
    case "":
        v17 := GenS4{F0: 769, F1: -1}
        _ = v17
    default:
        var v18 uint32 = 2
        _ = v18
    }
    switch -210 {
    case 4294967295:
        v12 = (v12 % ((4294967295) | 1))
    default:
        v19 := GenS2{F0: 14, F1: 4294967295}
        _ = v19
    }
    fmt.Println(v12, v2, v12)
    var v20 uint = v12
    v21 := GenS3{F0: -704, F1: v2, F2: (!(v20 != (v12 * 4294967295)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
