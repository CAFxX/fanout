package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    var v3 int64 = 900
    v4 := []int{}
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(fi5, fi5)
        v3 = (v3 / (((v3 >> ((962) & 63))) | 1))
        _ = fi5
    }
    var v6 string = "foo bar"
    var v7 string = "\n\t"
    if (2147483647 == 114) {
        switch v7 {
        case "hello":
            var v8 float64 = ((3.14159 + ((float64(1e10) / 1.0) / 1.0)) * float64(((0.0 / 1.5) - 0.1)))
            _ = v8
        case "\n\t":
            _ = len(v2)
        case "foo bar":
            v9 := GenS1{F0: 4294967295, F1: -1}
            _ = v9
        default:
            var v10 int16 = 240
            _ = v10
        }
    } else {
        v7 = ((v7 + v7) + v6)
        v11 := GenS1{F0: 42, F1: 2147483647}
        _ = v11
    }
    v12 := GenS1{F0: 65535, F1: -1}
    v13 := GenS1{F0: 18446744073709551615, F1: 978}
    var v14 uint = 1
    switch 42 {
    case 0:
        if (!true) {
            v3 += (v3 - (v3 & int64(v14)))
            fmt.Println(v3, v7)
            v6 = "foo bar"
        } else {
            for fi15 := 0; fi15 < 19; fi15++ {
                fmt.Println(v6, fi15, v7)
                _ = fi15
            }
        }
    }
    if (!false) {
        var v16 int8 = -1
        _ = v16
    }
    var v17 int = 817
    v7 = v6
    var v18 int = (v17 ^ (v17 & v17))
    if false {
        v18 = (v18 / (((v17 >> ((581) & 63))) | 1))
        v14 = v14
        v19 := GenS2{F0: v6}
        _ = v19
    } else {
        var v20 string = (v6 + (v6 + ("hello" + v7)))
        v20 = v20
        var v21 uint64 = 695
        _ = v20
        _ = v21
    }
    fmt.Println(v6)
    var v22 string = (v7 + ((((("foo bar" + v6) + v6) + v7) + v6) + v6))
    var v23 int32 = 470
    for fi24 := 0; fi24 < 4; fi24++ {
        if (true && true) {
            v25 := GenS1{F0: v14, F1: v23}
            _ = len(v4)
            for fi26 := 0; fi26 < 15; fi26++ {
                v17 = v17
                _ = fi26
            }
            _ = v25
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v17 = v17
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
