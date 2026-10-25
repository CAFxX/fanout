package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint }
func genFn0(p0 uint) uint {
    var v6 float32 = 1e10
    v7 := GenS1{F0: 127, F1: "\n\t", F2: 56}
    _ = p0
    _ = v6
    _ = v7
    return 283
}

func main() {
    var v1 bool = false
    v2 := []int{}
    v3 := make(map[uint32]bool)
    var v4 uint16 = 42
    var v5 bool = false
    delete(v3, 632)
    fmt.Println(v4, v5, v5)
    for fi8 := 0; fi8 < 5; fi8++ {
        if (111 >= -806) {
            fmt.Println(v1, v4, v1)
            delete(v3, 1)
        }
        _ = fi8
    }
    switch "" {
    case "hello":
        fmt.Println(v5, v5, v4)
    case "a":
        var v9 int32 = -718
        _ = v9
    case "foo bar":
        fmt.Println(v4)
    default:
        var v10 uint32 = 2147483647
        _ = v10
    }
    fmt.Println(v4, v1, v5)
    fmt.Println(v1, v5)
    var v11 uint = 18446744073709551615
    v11 -= uint(v4)
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: 100, F1: ("" + "\n\t"), F2: 13}
        switch fi12 {
        case -202:
            fmt.Println(v11, v11)
        case 729:
            for fi14 := 0; fi14 < 15; fi14++ {
                switch ("" + "hello") {
                case "hello":
                    v5 = v1
                case "foo bar":
                    var v15 int32 = 0
                    _ = v15
                }
                switch ("hello" + ("foo bar" + "hello")) {
                case "":
                    v1 = v1
                    var v16 float64 = ((0.0 - 1e10) / 1.5)
                    switch "\n\t" {
                    case "foo bar":
                        var v17 int32 = 852
                        _ = v17
                    case "":
                        if false {
                            v5 = (!v1)
                            var v18 int8 = -1
                            _ = len(v2)
                            _ = v18
                        }
                    }
                    _ = v16
                }
                _ = fi14
            }
        default:
            if (!(!v5)) {
                v19 := GenS1{F0: -61, F1: "hello", F2: 0}
                v5 = v5
                v20 := GenS2{F0: 492, F1: v11}
                _ = v19
                _ = v20
            }
        }
        _ = len(v3)
        _ = fi12
        _ = v13
    }
    var v21 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
