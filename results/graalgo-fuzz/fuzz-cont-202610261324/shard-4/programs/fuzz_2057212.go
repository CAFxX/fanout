package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
type GenS2 struct { F0 int16; F1 int; F2 int32 }

func main() {
    v1 := []string{"", "\n\t", ("\n\t" + (("\n\t" + "\n\t") + "foo bar")), "a"}
    var v2 string = (((("a" + "a") + (("foo bar" + ((("\n\t" + "hello") + "a") + "a")) + "")) + ("" + "a")) + "a")
    var v3 float32 = (1e10 + 3.14159)
    var v4 uint64 = 796
    var v5 uint = 255
    v6 := GenS1{F0: uint64(v5), F1: -1, F2: false}
    v5 -= (v5 ^ uint(v4))
    v4 = (v4 + v4)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int8 = -28
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v2, v5)
    v4 -= (v4 & (v4 + 4294967295))
    var v9 uint16 = 2
    v9 -= (v9 / ((42) | 1))
    if true {
        if len(v1) > 0 {
            v1[0] = v2
        }
        var v10 int32 = -946
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int = (fi11 * fi11)
            var v13 int = (-(fi11))
            var v14 uint64 = (v4 / ((v4) | 1))
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = v10
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v15 := GenS2{F0: 42, F1: -2147483648, F2: 0}
        v1 = append(v1, v2)
        _ = v15
    }
    var v16 string = ""
    if (3.14159 != -2.25) {
        fmt.Println(v5)
        var v17 uint32 = 65535
        var v18 uint = uint(v17)
        _ = v17
        _ = v18
    } else {
        var v19 int8 = 127
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
