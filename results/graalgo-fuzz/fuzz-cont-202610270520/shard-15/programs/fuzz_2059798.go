package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }

func main() {
    var v1 bool = true
    var v2 int64 = 877
    var v3 uint = 424
    v4 := []uint32{997, 4294967295, 1}
    var v5 uint16 = 255
    fmt.Println(v3, v2)
    if (!(247 > 65535)) {
        var v6 int = -9223372036854775808
        var v7 float32 = 0.1
        v7 += v7
        _ = v6
        _ = v7
    }
    fmt.Println(v3, v1, v5)
    fmt.Println(v5)
    for fi8 := 0; fi8 < 17; fi8++ {
        if (42 != -114) {
            v2 -= (v2 + v2)
            v9 := GenS1{F0: 4294967295, F1: "foo bar"}
            _ = v9
        } else {
            var v10 uint = (v3 - v3)
            if len(v4) > 0 {
                v4[0] = 65535
            }
            v11 := GenS1{F0: 2147483648, F1: ""}
            _ = v10
            _ = v11
        }
        v1 = v1
        _ = fi8
    }
    v2 = (-(v2))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (true || v1) {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int32 = 0
            var v14 bool = (int32(v5) != 1)
            var v15 int64 = (v2 & (^(v2)))
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
