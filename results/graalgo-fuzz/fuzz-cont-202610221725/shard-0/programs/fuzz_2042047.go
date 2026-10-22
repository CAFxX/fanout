package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint16, p1 int64, p2 int8) int32 {
    var v6 int = -824
    v7 := GenS1{F0: 128, F1: false}
    switch "\n\t" {
    case "foo bar":
        v8 := GenS2{F0: p2}
        _ = v8
    }
    fmt.Println(v6)
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            p0 = (p0 | (p0 / ((p0) | 1)))
            var v11 int16 = -32768
            switch v6 {
            case -9223372036854775808:
                switch "a" {
                case "\n\t":
                    v12 := GenS2{F0: int8(v6)}
                    _ = v12
                case "a":
                    fmt.Println(p1, p1)
                case "hello":
                    fmt.Println(v11, fi9, p1)
                }
            }
            _ = fi10
            _ = v11
        }
        if (("a" + "\n\t") == "foo bar") {
            v13 := []int{(v6 | fi9), v6}
            p2 += p2
            _ = v13
        }
        p1 = (p1 ^ p1)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int32, p1 uint) uint16 {
    fmt.Println(p1)
    var v14 int8 = 1
    fmt.Println(p1, p1, p1)
    for fi15 := 0; fi15 < 6; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := []int{fi15, (^(fi16)), fi15, (^(fi15))}
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    switch "" {
    case "":
        v18 := []int{-591, 2147483647, -569}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v14
    return 982
}

func main() {
    v1 := []int{-796, -352, -379}
    v2 := make(map[uint32]int)
    var v3 float64 = 1e10
    var v4 int = 2147483647
    var v5 uint32 = 42
    _ = len(v1)
    v19 := GenS2{F0: -43}
    var v20 uint32 = 552
    v5 -= 0
    fmt.Println(v5, v20, v3)
    _ = len(v2)
    switch ((("foo bar" + ("a" + "\n\t")) + "hello") + "foo bar") {
    case "\n\t":
        var v21 int64 = -566
        _ = v21
    default:
        v22 := GenS2{F0: 0}
        _ = v22
    }
    var v23 uint = 2
    v24 := GenS1{F0: 48, F1: (v4 < -9223372036854775808)}
    if (false && true) {
        var v25 string = "hello"
        var v26 int8 = -39
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        fmt.Println(v5)
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
