package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0(p0 bool) uint64 {
    p0 = ((("a" + (("hello" + "\n\t") + ("\n\t" + "foo bar"))) + "foo bar") == "foo bar")
    v6 := []int{612, 1, 128, 9223372036854775807}
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 0
}
func genFn1(p0 int8, p1 uint8, p2 int32) uint8 {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: "a", F1: (-(p2))}
            fmt.Println(p0)
            v10 := GenS1{F0: (("foo bar" + "hello") + ("" + "")), F1: p2}
            p0 = p0
            fmt.Println(fi7, p2)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v11 := GenS2{F0: 687, F1: 42, F2: 4294967295}
    p1 -= p1
    var v12 int32 = p2
    v13 := []int{-468}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 138
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 42
    var v3 uint = 4294967295
    var v4 uint8 = 119
    var v5 uint16 = 2
    fmt.Println(v2, v2, v5)
    v3 = v3
    v14 := GenS2{F0: (v5 + v5), F1: 786, F2: v3}
    var v15 int32 = -2147483648
    v16 := GenS2{F0: v5, F1: 713, F2: (^(v3))}
    v17 := GenS3{F0: (!false), F1: 911}
    fmt.Println(v4)
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int16 = -32768
        v2 -= (fi18 + (fi18 >> ((1) & 63)))
        _ = fi18
        _ = v19
    }
    v15 = v15
    v20 := GenS3{F0: true, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
