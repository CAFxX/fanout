package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 string }
func genFn0(p0 int32) int32 {
    for fi4 := 0; fi4 < 3; fi4++ {
        fmt.Println(p0)
        v5 := GenS1{F0: -23, F1: 258, F2: "hello"}
        var v6 uint8 = 128
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{}
    _ = p0
    _ = v7
    return (p0 << ((2) & 63))
}

func main() {
    var v1 int16 = -1
    v2 := []int{32767, 639, 32767}
    var v3 int16 = -1
    v3 = 1000
    switch "" {
    case "foo bar":
        v3 = v3
    case "\n\t":
        for fi8 := 0; fi8 < 19; fi8++ {
            v2 = append(v2, 0)
            fmt.Println(v3, fi8)
            _ = fi8
        }
    }
    if (!true) {
        for fi9 := 0; fi9 < 11; fi9++ {
            v1 += (v1 / (((v3 | (-(v3)))) | 1))
            v1 += -782
            _ = fi9
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            if true {
                if true {
                    v1 += v1
                    v3 += (v3 << ((2) & 63))
                    var v10 string = "hello"
                    _ = v10
                }
                v3 = 42
            }
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "hello":
            _ = len(v2)
        }
        var v11 int32 = 1
        _ = v11
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(fi12, fi12)
            fmt.Println(fi13, v3, fi12)
            _ = fi13
        }
        _ = fi12
    }
    var v14 float32 = (float32(3.14159) / 3.14159)
    v3 = v1
    var v15 float32 = v14
    _ = len(v2)
    fmt.Println(v15, v3, v3)
    var v16 int = 2147483647
    v17 := GenS1{F0: -128, F1: 42, F2: ""}
    for fi18 := 0; fi18 < 2; fi18++ {
        if len(v2) > 0 {
            v2[0] = (v16 / ((v16) | 1))
        }
        switch -2147483648 {
        case -345:
            fmt.Println(v1, v1)
        case 2147483647:
            if false {
                v15 -= v15
            } else {
                fmt.Println(v14, v15)
            }
        default:
            var v19 uint64 = 18446744073709551615
            _ = v19
        }
        _ = fi18
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
