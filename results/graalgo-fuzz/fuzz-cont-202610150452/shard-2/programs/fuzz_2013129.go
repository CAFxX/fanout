package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }

func main() {
    var v1 float32 = (3.14159 - 3.14159)
    var v2 int32 = -1
    var v3 uint8 = 138
    switch "\n\t" {
    case "a":
        v2 += (v2 & v2)
    }
    var v4 bool = (((504 > 4294967295) || false) || false)
    v4 = (127 == 127)
    var v5 int32 = v2
    if (0.0 >= 3.14159) {
        var v6 int64 = 9223372036854775807
        if (!v4) {
            v4 = (254 > v3)
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 int16 = 1000
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 int32 = (v2 >> ((2) & 63))
                    v1 -= v1
                    _ = fi9
                    _ = v10
                }
                _ = fi7
                _ = v8
            }
        } else {
            var v11 uint8 = (^(v3))
            _ = v11
        }
        fmt.Println(v3, v3)
        _ = v6
    }
    fmt.Println(v3, v4)
    v4 = (-18 < 0)
    v1 = (v1 / (((v1 + v1) * 0.1) - v1))
    switch 0 {
    case 4294967295:
        var v12 uint16 = uint16(v3)
        _ = v12
    default:
        v13 := GenS1{F0: 1, F1: 2147483647}
        _ = v13
    }
    switch -745 {
    case 9223372036854775807:
        var v14 float64 = -2.25
        _ = v14
    case -2147483648:
        var v15 uint32 = 0
        _ = v15
    case 421:
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := []int{fi16}
            switch "a" {
            case "a":
                var v18 uint = 18446744073709551615
                _ = v18
            case "hello":
                _ = len(v17)
            case "\n\t":
                v19 := GenS1{F0: int16(v5), F1: 101}
                _ = v19
            default:
                var v20 float64 = 1e10
                _ = v20
            }
            _ = fi16
            _ = v17
        }
    default:
        var v21 string = ("hello" + "\n\t")
        _ = v21
    }
    v22 := GenS1{F0: 42, F1: 190}
    for fi23 := 0; fi23 < 13; fi23++ {
        v3 += v3
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
