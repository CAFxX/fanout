package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0(p0 int64) bool {
    p0 = p0
    p0 -= (p0 / ((-397) | 1))
    _ = p0
    return false
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint16 = 42
    v3 := make(map[uint32]bool)
    var v4 bool = true
    v5 := make(map[string]string)
    if ((v4 && (int8(v2) < -128)) || v4) {
        v6 := GenS1{F0: (v4 || v4)}
        if v4 {
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS2{F0: 27}
                _ = fi7
                _ = v8
            }
            fmt.Println(v4, v2)
            fmt.Println(v1)
        }
        v4 = v4
        _ = v6
    }
    if (!v4) {
        for fi9 := 0; fi9 < 10; fi9++ {
            v4 = true
            _ = fi9
        }
        v10 := v5["\n\t"]
        _ = v10
    }
    v1 -= (v1 ^ (-(v1)))
    var v11 int32 = 494
    var v12 uint8 = 0
    v5["foo bar"] = ("" + "hello")
    if false {
        var v13 uint32 = 2
        v14 := GenS1{F0: ("" <= ("foo bar" + "a"))}
        fmt.Println(v2, v1, v1)
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS2{F0: 485}
            var v17 string = "\n\t"
            v18 := GenS3{F0: 1, F1: -835}
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    }
    v2 = v2
    if v4 {
        v4 = v4
    }
    if v4 {
        v4 = ((-(v11)) != -1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
