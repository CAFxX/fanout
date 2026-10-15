package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }

func main() {
    var v1 uint64 = 2
    var v2 string = "foo bar"
    var v3 int32 = -2147483648
    v4 := []int{}
    var v5 int = 1
    var v6 uint16 = 602
    var v7 uint = 2
    v5 = v5
    switch 32767 {
    case 371:
        var v8 uint8 = 42
        _ = v8
    case 1:
        fmt.Println(v3, v2, v6)
    case 42:
        if true {
            v9 := GenS1{F0: -32768, F1: (v7 * (v7 << ((v7) & 63))), F2: v1}
            for fi10 := 0; fi10 < 12; fi10++ {
                fmt.Println(fi10, v1, fi10)
                v1 += v1
                _ = fi10
            }
            _ = v9
        }
    }
    if false {
        if len(v4) > 0 {
            v4[0] = (v5 + (v5 ^ v5))
        }
        var v11 uint8 = 127
        _ = v11
    }
    v7 = 18446744073709551615
    switch v5 {
    case -2147483648:
        fmt.Println(v6, v2)
    case 0:
        v12 := GenS1{F0: 1000, F1: (v7 + 615), F2: v1}
        _ = v12
    }
    v13 := GenS2{F0: 2, F1: 1}
    v14 := GenS1{F0: 496, F1: v7, F2: (v1 & (^(v1)))}
    var v15 int16 = 32767
    switch v5 {
    case -235:
        var v16 int64 = 42
        _ = v16
    }
    switch v5 {
    case 9223372036854775807:
        var v17 bool = true
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v1 = v1
        var v19 string = (v2 + (((v2 + "foo bar") + (v2 + v2)) + v2))
        for fi20 := 0; fi20 < 12; fi20++ {
            v3 -= v3
            var v21 int8 = 0
            fmt.Println(v2, v3)
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v2 = ""
    var v22 uint8 = 0
    v15 = -763
    v6 = (v6 * v6)
    switch "" {
    case "foo bar":
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint = 2
            _ = fi23
            _ = v24
        }
    case "":
        v22 = (v22 ^ v22)
    }
    var v25 int32 = (v3 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
