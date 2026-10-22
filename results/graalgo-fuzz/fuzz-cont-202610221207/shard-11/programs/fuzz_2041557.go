package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int }
func genFn0(p0 int8) int64 {
    var v7 int64 = 869
    var v8 uint16 = 624
    var v9 uint = 2
    if false {
        fmt.Println(p0, v8, v7)
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (v7 / ((v7) | 1))
}

func main() {
    var v1 uint16 = 0
    var v2 bool = (!(((1.5 + 0.0) + 3.14159) == -2.25))
    v3 := make(map[int]int)
    v4 := []int{}
    var v5 int8 = 42
    v6 := make(map[int]int)
    switch "foo bar" {
    case "foo bar":
        var v10 string = "a"
        _ = v10
    case "":
        if v2 {
            for fi11 := 0; fi11 < 8; fi11++ {
                if v2 {
                    v1 = 42
                    v2 = false
                    var v12 int8 = (v5 * v5)
                    _ = v12
                } else {
                    if v2 {
                        var v13 uint64 = 18446744073709551615
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v15 := GenS1{F0: (v1 >> ((18446744073709551615) & 63))}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi14
                        _ = v15
                    }
                }
                if len(v4) > 0 {
                    v4[0] = -1
                }
                _ = fi11
            }
            var v16 float64 = -2.25
            _ = v16
        }
    default:
        delete(v6, -623)
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        switch "\n\t" {
        case "hello":
            for fi18 := 0; fi18 < 4; fi18++ {
                v19 := GenS1{F0: (v1 - (v1 - (v1 << ((65535) & 63))))}
                _ = fi18
                _ = v19
            }
        }
        _ = fi17
    }
    v20 := GenS2{F0: -103, F1: 65535, F2: 4294967295}
    fmt.Println(v2, v1)
    v21 := GenS1{F0: (^(v1))}
    v22 := GenS3{F0: v2, F1: v2}
    var v23 bool = (0 < 0)
    var v24 int = 4294967295
    v25 := GenS4{F0: v24}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
