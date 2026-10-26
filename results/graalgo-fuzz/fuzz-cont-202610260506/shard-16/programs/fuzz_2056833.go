package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }

func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    v3 := []uint32{2147483647, 250}
    v4 := make(map[int]string)
    var v5 int = -571
    v5 += v5
    _ = len(v4)
    if ("foo bar" != "hello") {
        v6 := v4[v5]
        _ = v6
    }
    v5 = v5
    v5 += v5
    fmt.Println(v5, v5)
    v5 = v5
    if len(v3) > 0 {
        v3[0] = 911
    }
    v1 -= (v1 | 0)
    v1 -= -9223372036854775808
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!false) {
        for fi7 := 0; fi7 < 15; fi7++ {
            if true {
                v8 := GenS1{F0: ("\n\t" + (("a" + "hello") + "a")), F1: 100}
                fmt.Println(v5, fi7)
                for fi9 := 0; fi9 < 9; fi9++ {
                    v10 := GenS2{F0: -1, F1: (-(v1)), F2: -9223372036854775808}
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                fmt.Println(v5)
                v4[fi7] = ""
                var v11 float32 = 1.5
                _ = v11
            }
            _ = fi7
        }
        var v12 float64 = 3.14159
        _ = v12
    } else {
        var v13 int8 = int8(v1)
        var v14 int32 = 42
        _ = v13
        _ = v14
    }
    var v15 uint16 = 1
    v5 += -2147483648
    var v16 int16 = 42
    v17 := GenS1{F0: (("\n\t" + "hello") + ""), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
