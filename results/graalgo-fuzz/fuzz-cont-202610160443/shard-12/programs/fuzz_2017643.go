package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int16 }

func main() {
    var v1 float32 = ((0.0 + 1.5) * float32(0.1))
    var v2 uint64 = 1
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 bool = false
    var v6 bool = (2 > 2)
    var v7 int16 = -32768
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint64 = (v2 / ((v2) | 1))
        fmt.Println(v2, v2)
        _ = fi8
        _ = v9
    }
    switch ("hello" + (("foo bar" + "\n\t") + "a")) {
    case "\n\t":
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := []int{128, (fi10 % ((2147483647) | 1))}
            _ = fi10
            _ = v11
        }
    default:
        v5 = false
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint16 = 226
        fmt.Println(v5, v2)
        v5 = false
        _ = fi12
        _ = v13
    }
    fmt.Println(v7)
    var v14 bool = v6
    v15 := GenS1{F0: "\n\t", F1: (v2 ^ (v2 + v2)), F2: v7}
    if false {
        v7 = v7
        fmt.Println(v5)
    }
    fmt.Println(v6, v1, v7)
    v16 := GenS2{F0: v2, F1: v2, F2: (v7 + v7)}
    var v17 int32 = -1
    for fi18 := 0; fi18 < 1; fi18++ {
        fmt.Println(v5, v6, fi18)
        _ = fi18
    }
    fmt.Println(v6)
    v5 = v5
    if (4294967295 == 315) {
        fmt.Println(v1, v7)
        var v19 uint8 = 255
        _ = v19
    } else {
        var v20 uint8 = 127
        if ("a" != ("" + "a")) {
            v20 = (v20 - (^(v20)))
            v1 -= v1
        } else {
            var v21 uint64 = 18446744073709551615
            var v22 string = "hello"
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 uint16 = 2
                v25 := []int{631, 9223372036854775807, fi23}
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v21
            _ = v22
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
