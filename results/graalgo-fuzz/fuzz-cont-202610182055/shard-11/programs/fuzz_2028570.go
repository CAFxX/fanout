package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint32 }
func genFn0() int16 {
    v7 := GenS1{F0: 1, F1: -477}
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 bool = (!true)
            fmt.Println(fi9, fi8)
            fmt.Println(fi9)
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807}
        var v12 bool = true
        _ = fi8
        _ = v11
        _ = v12
    }
    var v13 int64 = 1
    var v14 int32 = -1
    v14 -= -1
    _ = v13
    _ = v14
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 42
    var v2 uint8 = 128
    var v3 uint = uint(v1)
    v4 := []uint32{}
    var v5 string = ("hello" + "foo bar")
    v6 := make(map[string]int)
    v2 -= v2
    var v15 bool = (!true)
    v16 := GenS2{F0: 55, F1: -545}
    v3 = v3
    if (!((("foo bar" + "foo bar") + "foo bar") >= ((v5 + "a") + (("foo bar" + v5) + v5)))) {
        v17 := GenS3{F0: (v1 - v1)}
        fmt.Println(v3, v1)
        _ = v17
    } else {
        for fi18 := 0; fi18 < 13; fi18++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v19 float64 = -2.25
            _ = fi18
            _ = v19
        }
    }
    if v15 {
        v2 = v2
    }
    if v15 {
        fmt.Println(v3, v2)
    }
    v20 := GenS1{F0: (v3 >> (((v3 * 42)) & 63)), F1: 0}
    switch (v5 + v5) {
    case "\n\t":
        v15 = (42 < (v1 * (v1 - 2147483648)))
    case "a":
        fmt.Println(v3, v2)
    default:
        if v15 {
            fmt.Println(v2, v5)
            v2 += (v2 + 77)
            var v21 uint32 = 4294967295
            _ = v21
        }
    }
    fmt.Println(v5, v3)
    v15 = true
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
