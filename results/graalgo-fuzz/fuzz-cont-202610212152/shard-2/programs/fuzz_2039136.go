package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{65535}
    var v3 uint64 = 4294967295
    var v4 string = ""
    var v5 float64 = 0.0
    if ((float64((0.1 + 1.5)) / ((1e10 + v5) - (-2.25 - v5))) <= v5) {
        switch 127 {
        case 536:
            var v6 uint = 65535
            _ = v6
        case -2147483648:
            v3 = 18446744073709551615
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                v2 = append(v2, 42)
                fmt.Println(v4, v4)
                fmt.Println(v3)
                _ = fi8
            }
            _ = fi7
        }
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(v1)
        _ = fi9
    }
    if true {
        fmt.Println(v4, v3)
    }
    fmt.Println(v1, v4)
    var v10 uint = 255
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 int8 = -61
    v10 += (v10 / ((4294967295) | 1))
    switch (v4 + v4) {
    case "\n\t":
        var v12 bool = (1 > -2147483648)
        _ = v12
    case "":
        for fi13 := 0; fi13 < 10; fi13++ {
            v4 = v4
            _ = fi13
        }
    case "foo bar":
        var v14 int64 = -9223372036854775808
        _ = v14
    default:
        if (42 <= 311) {
            var v15 uint16 = 65535
            var v16 uint = v10
            _ = v15
            _ = v16
        }
    }
    v3 = (v3 + v3)
    if true {
        var v17 int64 = 893
        for fi18 := 0; fi18 < 17; fi18++ {
            v11 += 1
            _ = fi18
        }
        _ = v17
    } else {
        fmt.Println(v10)
    }
    v19 := GenS1{F0: -1, F1: 1, F2: 2147483647}
    var v20 int64 = -241
    switch v4 {
    case "a":
        var v21 int = 229
        _ = v21
    }
    v4 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
