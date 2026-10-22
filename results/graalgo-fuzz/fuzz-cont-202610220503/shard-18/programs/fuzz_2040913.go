package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int; F2 uint }

func main() {
    var v1 uint = 4294967295
    var v2 int8 = 1
    v3 := []string{"foo bar", "hello", (("" + "foo bar") + "\n\t")}
    v4 := make(map[string]int)
    var v5 float64 = 0.0
    v6 := GenS1{F0: 32767}
    if ((v5 * v5) >= 3.14159) {
        if true {
            switch "foo bar" {
            case "":
                v1 = (v1 * 295)
            }
        }
        var v7 uint32 = 42
        _ = v7
    } else {
        delete(v4, "hello")
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        switch (fi8 >> (((v1 ^ (v1 / (((v1 % ((65535) | 1))) | 1)))) & 63)) {
        case -48:
            fmt.Println(v1)
        case 127:
            v9 := GenS1{F0: -32768}
            _ = v9
        default:
            var v10 uint8 = 127
            _ = v10
        }
        delete(v4, "hello")
        _ = fi8
    }
    if (127 >= 84) {
        var v11 int = 2147483647
        if false {
            _ = len(v3)
            var v12 float32 = 1.5
            _ = v12
        }
        switch "" {
        case "a":
            v13 := GenS2{F0: -1, F1: v11, F2: 587}
            _ = v13
        case "\n\t":
            fmt.Println(v11)
        case "foo bar":
            v2 = (v2 + v2)
        default:
            if true {
                for fi14 := 0; fi14 < 12; fi14++ {
                    v11 -= (v11 - (fi14 + (v11 * 32768)))
                    v5 -= (v5 - -2.25)
                    _ = fi14
                }
                var v15 float32 = 1.5
                _ = v15
            }
        }
        _ = v11
    } else {
        var v16 uint8 = 2
        _ = v16
    }
    v1 = (v1 - (v1 << ((v1) & 63)))
    var v17 uint16 = 24
    var v18 int64 = 0
    fmt.Println(v17, v2)
    v2 = (v2 >> (((v1 ^ v1)) & 63))
    fmt.Println(v17, v17, v18)
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 float32 = (((0.1 * float32(1e10)) + 0.0) + 1.5)
            fmt.Println(v2, v2)
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v5 -= v5
    var v22 uint16 = v17
    v17 = v22
    for fi23 := 0; fi23 < 11; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 float32 = (1e10 - -2.25)
            _ = fi24
            _ = v25
        }
        fmt.Println(v2, v22)
        var v26 int8 = -16
        _ = fi23
        _ = v26
    }
    var v27 int64 = (v18 * v18)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
