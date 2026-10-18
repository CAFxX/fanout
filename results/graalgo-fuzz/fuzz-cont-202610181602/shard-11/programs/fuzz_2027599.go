package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 int16, p1 int8, p2 int32) bool {
    var v6 string = ("" + ("a" + ("\n\t" + ("foo bar" + (("foo bar" + "\n\t") + "")))))
    fmt.Println(v6, v6, p1)
    var v7 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 int = 9223372036854775807
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    if ((127 <= 25) && (!true)) {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int64 = 9223372036854775807
            v10 := GenS1{F0: fi8, F1: "hello"}
            _ = fi8
            _ = v10
            _ = v9
        }
        v3 = -146
    } else {
        var v11 float64 = 1e10
        v3 -= (v3 | v3)
        v3 = v3
        _ = v11
    }
    fmt.Println(v3)
    v3 = (-(v3))
    v3 -= v3
    var v12 float64 = -2.25
    v13 := GenS2{F0: (-(v3)), F1: 86}
    v3 += (-(v3))
    v12 -= v12
    v3 = v3
    for fi14 := 0; fi14 < 8; fi14++ {
        if (!(!false)) {
            v12 += v12
            fmt.Println(fi14, fi14, v12)
            v3 = 32768
        }
        _ = fi14
    }
    v15 := GenS2{F0: 0, F1: 127}
    v16 := GenS3{F0: 9223372036854775807, F1: -1}
    if true {
        var v17 bool = true
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
