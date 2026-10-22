package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 0
    v3 := make(map[string]int)
    v4 := []uint32{2147483648, 487, 42}
    var v5 uint = 454
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: (v2 >> (((v5 - (v5 ^ (^(v5))))) & 63)), F1: 0, F2: 0}
            var v9 uint64 = 474
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 uint8 = 1
    v11 := GenS1{F0: (v2 + int16(v10)), F1: 2147483647, F2: -1}
    switch -2147483648 {
    case -538:
        if true {
            switch "a" {
            case "a":
                fmt.Println(v10, v10)
            case "":
                v10 = (v10 & v10)
            case "\n\t":
                delete(v3, ("hello" + (("hello" + "a") + ("a" + "a"))))
            }
            v5 = v5
        }
    case -699:
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: (^(v2)), F1: -1, F2: 100}
            _ = fi12
            _ = v13
        }
    case -9223372036854775808:
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(v2)
            var v15 uint8 = v10
            for fi16 := 0; fi16 < 3; fi16++ {
                v15 += v10
                var v17 uint32 = 4294967295
                var v18 uint = 103
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi14
            _ = v15
        }
    }
    if (128 < -194) {
        if true {
            if true {
                v2 = 32767
                switch "a" {
                case "hello":
                    v10 -= (^(v10))
                default:
                    var v19 uint8 = 2
                    _ = v19
                }
                if len(v4) > 0 {
                    v4[0] = 2
                }
            }
            fmt.Println(v2)
        }
    }
    v20 := GenS1{F0: -32768, F1: 2147483647, F2: 0}
    switch "\n\t" {
    case "":
        v21 := GenS1{F0: (-(v2)), F1: -2147483648, F2: 1}
        _ = v21
    case "foo bar":
        var v22 float64 = (float64((((0.1 - 1e10) / -2.25) / -2.25)) / 1.0)
        _ = v22
    case "\n\t":
        var v23 float32 = 0.0
        _ = v23
    default:
        v24 := GenS2{F0: 42, F1: 255, F2: v5}
        _ = v24
    }
    var v25 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
