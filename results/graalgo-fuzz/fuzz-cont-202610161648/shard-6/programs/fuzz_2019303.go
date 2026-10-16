package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int64; F1 bool; F2 int }
func genFn0(p0 int32, p1 uint16) int {
    v7 := []int{-502, 42, 32768}
    p1 = p1
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint32 = 0
        var v10 float64 = 1e10
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 333
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 872
    v3 := []string{"\n\t", ("\n\t" + "hello")}
    var v4 uint16 = 585
    var v5 uint64 = 0
    var v6 uint64 = v5
    var v11 uint16 = (^(v4))
    v5 += (^(v5))
    switch ("" + "hello") {
    case "":
        v4 += v11
    default:
        v11 -= (v11 | (^(v4)))
    }
    v12 := GenS1{F0: v4, F1: v11}
    if true {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint32 = 0
            fmt.Println(v5, v2, v11)
            v11 = v4
            _ = fi13
            _ = v14
        }
        v4 = v11
    }
    switch (("hello" + ("a" + "foo bar")) + "a") {
    case "hello":
        var v15 uint16 = (^(v4))
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        switch "foo bar" {
        case "":
            var v17 int = fi16
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS2{F0: 9223372036854775807, F1: false, F2: fi18}
                _ = len(v1)
                var v20 uint = 4294967295
                _ = fi18
                _ = v19
                _ = v20
            }
        case "foo bar":
            fmt.Println(v5, v4, v6)
        }
        _ = fi16
    }
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
