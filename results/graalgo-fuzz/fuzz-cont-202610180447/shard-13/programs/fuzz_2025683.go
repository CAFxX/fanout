package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []string{"hello", "foo bar", "a"}
    var v2 uint16 = 1
    v3 := make(map[int]string)
    var v4 float64 = (-2.25 * 0.1)
    var v5 int64 = int64(v2)
    fmt.Println(v2, v2, v5)
    fmt.Println(v4)
    delete(v3, -2147483648)
    v6 := GenS1{F0: 128}
    v5 -= v5
    v4 += v4
    v7 := GenS1{F0: 97}
    fmt.Println(v5, v2, v5)
    fmt.Println(v2, v5)
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            fmt.Println(v5, v2)
            if true {
                fmt.Println(v4, fi8)
            }
            fmt.Println(fi8)
            _ = fi9
        }
        _ = fi8
    }
    if false {
        var v10 uint16 = 2
        _ = v10
    }
    v5 = v5
    var v11 int = 32768
    v11 += (v11 / (((v11 % ((4294967295) | 1))) | 1))
    var v12 uint32 = 0
    switch "a" {
    case "a":
        v13 := GenS1{F0: 42}
        _ = v13
    case "foo bar":
        var v14 uint = 18446744073709551615
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 7; fi15++ {
            switch "hello" {
            case "a":
                _ = len(v3)
            case "foo bar":
                fmt.Println(v5, fi15, v11)
            default:
                v4 = -2.25
            }
            v5 -= v5
            _ = fi15
        }
        fmt.Println(v11, v5, v5)
        v11 -= (v11 + v11)
    }
    var v16 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
