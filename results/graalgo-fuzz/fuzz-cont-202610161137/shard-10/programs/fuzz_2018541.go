package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 bool; F1 bool; F2 uint16 }

func main() {
    var v1 float64 = 1e10
    v2 := []string{""}
    var v3 uint = 18446744073709551615
    var v4 uint8 = 42
    for fi5 := 0; fi5 < 18; fi5++ {
        if ((1e10 + 1e10) < 1.5) {
            _ = len(v2)
        } else {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v1, v3)
            default:
                fmt.Println(fi5)
            }
            var v6 int16 = 1
            _ = v6
        }
        var v7 uint = v3
        _ = fi5
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: (v3 >> ((4294967295) & 63))}
    fmt.Println(v1, v3, v4)
    v3 -= (v3 >> (((v3 & 18446744073709551615)) & 63))
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            v4 = v4
            var v11 string = ""
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: 1, F1: (v3 * 2)}
        if false {
            var v13 uint32 = 4294967295
            v14 := GenS1{F0: v13, F1: (v3 << ((v3) & 63))}
            v13 -= (v13 - 2147483648)
            _ = v13
            _ = v14
        }
        _ = fi9
        _ = v12
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(v1, fi15, v3)
        var v16 uint16 = 428
        _ = fi15
        _ = v16
    }
    if true {
        v2 = append(v2, "foo bar")
        v17 := GenS2{F0: (!true), F1: (-256 < 182), F2: 643}
        _ = v17
    }
    v18 := GenS2{F0: false, F1: true, F2: 1}
    var v19 bool = true
    v4 += uint8(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += (v3 / (((v3 | (v3 + (v3 & (v3 / (((v3 >> ((v3) & 63))) | 1)))))) | 1))
    var v20 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
