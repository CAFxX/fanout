package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint8 = 33
    v2 := []string{"hello", "\n\t"}
    v3 := make(map[string]bool)
    var v4 uint16 = 273
    v5 := make(map[string]bool)
    var v6 bool = ((!(!false)) && true)
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 16}
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(v6, v6)
            v10 := GenS2{F0: 735, F1: 2147483647}
            if v6 {
                var v11 int16 = 32767
                v5[("" + "hello")] = v6
                var v12 int16 = (v11 / ((v11) | 1))
                _ = v11
                _ = v12
            } else {
                var v13 float32 = 1e10
                v1 = (v1 ^ v1)
                var v14 uint = 761
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = len(v2)
    var v15 float64 = 0.1
    v6 = (42 <= 90)
    switch "a" {
    case "":
        var v16 uint16 = v4
        _ = v16
    case "foo bar":
        v3["\n\t"] = false
    case "\n\t":
        v4 += 1
    default:
        v2 = append(v2, ("\n\t" + "a"))
    }
    v1 -= 158
    v2 = append(v2, "\n\t")
    v1 -= (v1 + 42)
    var v17 int8 = 1
    v18 := GenS3{F0: true}
    v17 += (v17 << ((255) & 63))
    switch 128 {
    case 4294967295:
        fmt.Println(v1, v15)
    case -1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -2147483648:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi19 := 0; fi19 < 17; fi19++ {
            for fi20 := 0; fi20 < 6; fi20++ {
                v21 := GenS2{F0: -740, F1: 2147483647}
                var v22 int8 = v17
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
