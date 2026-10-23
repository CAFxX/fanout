package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int, p2 string) int {
    var v6 int32 = 689
    fmt.Println(v6, p2)
    var v7 uint16 = 519
    p0 -= (p0 | (p0 & p0))
    if (1 < 369) {
        v8 := GenS1{F0: v7, F1: 9223372036854775807}
        p1 -= (p1 % ((int(v6)) | 1))
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := []int{(fi9 - fi9), (fi9 % (((fi9 - 641)) | 1))}
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}
func genFn1(p0 int64, p1 uint64) uint16 {
    p1 += (p1 - p1)
    p1 = 0
    var v11 string = (("hello" + (("foo bar" + "a") + ("a" + (("foo bar" + "hello") + ("foo bar" + ("hello" + "hello")))))) + ((("foo bar" + "\n\t") + "foo bar") + "hello"))
    _ = p0
    _ = p1
    _ = v11
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{128}
    var v3 int32 = 807
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    v12 := v1[42]
    v3 += 1
    v13 := GenS1{F0: genFn1(1, 2), F1: 42}
    if false {
        var v14 uint8 = 1
        v14 += (v14 % (((v14 | v14)) | 1))
        fmt.Println(v12, v12, v14)
        _ = v14
    } else {
        if true {
            v12 -= (v12 ^ (v12 >> ((42) & 63)))
        }
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 int32 = v3
        _ = fi15
        _ = v16
    }
    var v17 float32 = (3.14159 * 3.14159)
    _ = len(v5)
    v18 := GenS2{F0: (v12 ^ v12)}
    var v19 int8 = -128
    v2 = append(v2, -1)
    for fi20 := 0; fi20 < 10; fi20++ {
        switch ("a" + (("\n\t" + ((("hello" + ("hello" + "a")) + "a") + "")) + "")) {
        case "hello":
            v21 := GenS2{F0: (v12 + 1)}
            _ = v21
        default:
            fmt.Println(v3)
        }
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
