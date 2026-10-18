package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []string{(("" + "foo bar") + "\n\t"), ""}
    var v2 int = 1
    var v3 bool = ((100 <= 0) || (-9223372036854775808 > -78))
    var v4 uint8 = 0
    v5 := make(map[uint32]uint32)
    var v6 int32 = 0
    v2 = (v2 % ((32768) | 1))
    v1 = append(v1, "a")
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = 1e10
        v9 := GenS1{F0: 757}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        if v3 {
            var v11 string = "\n\t"
            _ = len(v5)
            var v12 int8 = 102
            _ = v11
            _ = v12
        } else {
            if true {
                v2 -= v2
            }
        }
        v13 := GenS1{F0: 1}
        var v14 int64 = -9223372036854775808
        _ = fi10
        _ = v13
        _ = v14
    }
    v4 += (^(v4))
    if (4294967295 == 2147483648) {
        var v15 int64 = 9223372036854775807
        v16 := GenS1{F0: 65535}
        _ = v15
        _ = v16
    } else {
        fmt.Println(v3, v6, v3)
    }
    switch "a" {
    case "\n\t":
        var v17 uint8 = (v4 + v4)
        _ = v17
    default:
        if v3 {
            v18 := GenS1{F0: 2}
            _ = v18
        } else {
            fmt.Println(v6, v2)
            if (!(v3 && ((!v3) || v3))) {
                fmt.Println(v2, v6, v3)
                v6 = (v6 + (v6 + int32(v4)))
            }
        }
    }
    var v19 float32 = (((1.5 + ((0.0 + -2.25) / 0.1)) + 0.0) / 3.14159)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
