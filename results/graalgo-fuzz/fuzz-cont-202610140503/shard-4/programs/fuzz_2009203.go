package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
func genFn0(p0 uint16) uint64 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 string = "hello"
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        p0 = 140
        _ = fi9
    }
    _ = p0
    return 42
}
func genFn1() int64 {
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int8 = 1
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    var v13 int = 4294967295
    if ((!((-128 <= 109) || ((9223372036854775807 < 42) && (0.1 >= 1.5)))) || false) {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 uint = 0
            v13 = 127
            var v16 uint = v15
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 bool = true
        v13 = (v13 << ((485) & 63))
        _ = v17
    } else {
        for fi18 := 0; fi18 < 11; fi18++ {
            v13 += v13
            _ = fi18
        }
        v13 = 2147483647
    }
    var v19 uint8 = 128
    _ = v13
    _ = v19
    return 0
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 128
    v3 := []int{2147483647, 127, -192}
    v4 := make(map[int]int)
    var v5 uint16 = 32767
    v6 := []string{""}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 += v2
    var v20 uint32 = 0
    v21 := GenS1{F0: 32767, F1: uint8(v5)}
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := v4[(fi22 >> (((v1 + 0)) & 63))]
        _ = fi22
        _ = v23
    }
    switch 128 {
    case -736:
        v2 = 42
    }
    v4[4294967295] = 4294967295
    v20 += v20
    v5 = v5
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(v20, v20, v2)
        _ = fi24
    }
    v2 = 127
    v25 := GenS1{F0: -200, F1: (v2 ^ (v2 - (v2 / ((v2) | 1))))}
    var v26 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
