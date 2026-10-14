package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -1
    v3 := []string{}
    v4 := make(map[int]int)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2)
    var v5 float32 = 0.0
    var v6 uint16 = 1
    fmt.Println(v2, v2, v5)
    switch (((("hello" + "a") + "") + "foo bar") + "a") {
    case "\n\t":
        for fi7 := 0; fi7 < 19; fi7++ {
            _ = len(v3)
            v3 = append(v3, "")
            _ = fi7
        }
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    switch "foo bar" {
    case "\n\t":
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 42}
            switch "foo bar" {
            case "\n\t":
                v6 -= (^(v6))
            default:
                var v10 bool = true
                _ = v10
            }
            fmt.Println(v6)
            _ = fi8
            _ = v9
        }
    case "hello":
        if false {
            var v11 uint32 = 2147483647
            v11 = 2147483647
            v11 = v11
            _ = v11
        }
    default:
        switch ("" + "\n\t") {
        case "a":
            v5 -= 1e10
        default:
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS2{F0: fi12, F1: 2147483647}
                if true {
                    v14 := GenS3{F0: 65535}
                    _ = v14
                }
                v5 = 1e10
                _ = fi12
                _ = v13
            }
        }
    }
    fmt.Println(v5, v5)
    var v15 uint = 18446744073709551615
    fmt.Println(v6, v5, v6)
    fmt.Println(v15, v15, v6)
    for fi16 := 0; fi16 < 9; fi16++ {
        v3 = append(v3, (("" + ((("a" + "a") + "hello") + "foo bar")) + "\n\t"))
        v2 = (v2 - v2)
        _ = fi16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int16 = 32767
        switch "\n\t" {
        case "foo bar":
            var v19 int16 = v18
            _ = v19
        case "\n\t":
            var v20 float32 = v5
            _ = v20
        case "hello":
            v15 -= (v15 << ((v15) & 63))
        }
        v5 = (0.1 * ((v5 + 1e10) * v5))
        _ = fi17
        _ = v18
    }
    v21 := GenS2{F0: 128, F1: -866}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
