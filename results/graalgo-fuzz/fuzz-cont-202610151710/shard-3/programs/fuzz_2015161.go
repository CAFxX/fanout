package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int8 = 103
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483647
    var v4 float64 = 0.1
    v4 = 1e10
    fmt.Println(v3, v3, v1)
    for fi5 := 0; fi5 < 4; fi5++ {
        v3 += (v3 | (v3 + (v3 & (v3 ^ v3))))
        _ = fi5
    }
    v6 := []int{-35, -602, 32767}
    v1 = 42
    v1 = 100
    var v7 int = 4294967295
    var v8 int = v7
    for fi9 := 0; fi9 < 19; fi9++ {
        if false {
            for fi10 := 0; fi10 < 12; fi10++ {
                v1 = (v1 >> ((2) & 63))
                var v11 uint8 = 127
                switch ((("hello" + "hello") + ("\n\t" + ("a" + ("a" + "hello")))) + (("\n\t" + "foo bar") + "foo bar")) {
                case "\n\t":
                    var v12 int = (^(fi9))
                    _ = v12
                case "":
                    v1 = v1
                case "foo bar":
                    if len(v6) > 0 {
                        v6[0] = 0
                    }
                }
                _ = fi10
                _ = v11
            }
        }
        switch ("" + ("\n\t" + (("a" + "foo bar") + "\n\t"))) {
        case "":
            fmt.Println(fi9, v4)
        }
        _ = fi9
    }
    v3 = (^(v3))
    switch (("a" + "a") + "a") {
    case "a":
        var v13 float64 = 3.14159
        _ = v13
    }
    v14 := GenS1{F0: v7, F1: 2, F2: -285}
    v3 = (v3 ^ v3)
    v15 := GenS2{F0: 42}
    if (v1 == v1) {
        fmt.Println(v3, v8)
    }
    switch ("a" + "hello") {
    case "\n\t":
        v16 := GenS2{F0: 0}
        _ = v16
    default:
        fmt.Println(v7, v4)
    }
    v7 -= (v8 - (v8 << ((345) & 63)))
    v17 := GenS1{F0: v8, F1: 2, F2: 806}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
