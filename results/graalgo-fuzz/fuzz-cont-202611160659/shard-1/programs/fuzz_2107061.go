package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
func genFn0(p0 uint16) int {
    var v5 int = 267
    v6 := GenS1{F0: (^(v5)), F1: 70}
    fmt.Println(v5)
    v5 += int(p0)
    _ = p0
    _ = v5
    _ = v6
    return (-(v5))
}
func genFn1(p0 int8, p1 int64) int32 {
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 float32 = 1.5
            _ = fi8
            _ = v9
        }
        var v10 uint32 = 42
        p0 += p0
        _ = fi7
        _ = v10
    }
    fmt.Println(p0)
    v11 := GenS1{F0: 0, F1: 128}
    p0 = (-(p0))
    _ = p0
    _ = p1
    _ = v11
    return 1
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []uint32{2}
    var v3 int64 = -9223372036854775808
    var v4 uint32 = uint32(v1)
    v3 += 9223372036854775807
    v1 += (v1 << ((v1) & 63))
    v3 -= (-(v3))
    if (false && (1000 == 0)) {
        switch "" {
        case "hello":
            fmt.Println(v3, v3, v3)
        case "\n\t":
            v3 = (v3 << ((v1) & 63))
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                v2 = append(v2, v4)
                var v13 uint64 = 114
                var v14 uint16 = 2
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    }
    var v15 int32 = genFn1(-1, -9223372036854775808)
    v2 = append(v2, v4)
    if true {
        var v16 int16 = 32767
        v16 += (^(v16))
        _ = v16
    }
    var v17 uint = 4294967295
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 uint = v1
        v20 := GenS1{F0: 32767, F1: 0}
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int32 = (v15 & 0)
            fmt.Println(v15, fi21, v15)
            _ = fi21
            _ = v22
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    if ((true && (42 <= 100)) && true) {
        switch "hello" {
        case "foo bar":
            v23 := GenS1{F0: 127, F1: 0}
            _ = v23
        }
        var v24 uint16 = 0
        _ = v24
    } else {
        if len(v2) > 0 {
            v2[0] = v4
        }
        v15 = 2147483647
        v4 -= 1
    }
    v25 := GenS1{F0: 1, F1: 255}
    var v26 bool = (v1 > uint(v4))
    for fi27 := 0; fi27 < 4; fi27++ {
        v15 = v15
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
