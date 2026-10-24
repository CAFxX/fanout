package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 bool }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
type GenS3 struct { F0 int16; F1 string }

func main() {
    var v1 float32 = (float32(3.14159) / 1.0)
    v2 := make(map[string]string)
    v3 := []uint32{2147483647, 2147483648, 162}
    var v4 float64 = 1e10
    var v5 int = 32768
    var v6 uint8 = 128
    v4 += float64(v1)
    var v7 bool = (0 > 0)
    var v8 int = v5
    v9 := GenS1{F0: 83, F1: "foo bar", F2: v7}
    v6 -= 255
    fmt.Println(v5, v6, v1)
    for fi10 := 0; fi10 < 14; fi10++ {
        switch "a" {
        case "":
            for fi11 := 0; fi11 < 14; fi11++ {
                v7 = (513 == 241)
                fmt.Println(v1, v5, v1)
                _ = fi11
            }
        case "foo bar":
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS2{F0: ((v7 && true) || v7), F1: 873, F2: v7}
                switch "" {
                case "\n\t":
                    v6 = v6
                case "":
                    var v14 float32 = v1
                    fmt.Println(v7, fi12, v1)
                    var v15 uint32 = 42
                    _ = v14
                    _ = v15
                default:
                    v16 := GenS2{F0: v7, F1: 0, F2: v7}
                    _ = v16
                }
                var v17 float32 = v1
                _ = fi12
                _ = v13
                _ = v17
            }
        default:
            v5 += v8
        }
        v8 -= v8
        _ = len(v2)
        _ = fi10
    }
    var v18 uint32 = 4294967295
    v5 += v8
    v1 -= (3.14159 * float32(v4))
    if false {
        v19 := GenS2{F0: false, F1: 255, F2: ("hello" != "hello")}
        _ = v19
    } else {
        v6 -= (v6 >> ((345) & 63))
    }
    v6 += v6
    if len(v3) > 0 {
        v3[0] = 0
    }
    if true {
        switch "\n\t" {
        case "\n\t":
            var v20 float64 = v4
            _ = v20
        case "hello":
            for fi21 := 0; fi21 < 12; fi21++ {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 uint32 = (v18 | 2147483648)
                    var v24 uint8 = v6
                    if len(v3) > 0 {
                        v3[0] = v18
                    }
                    var v25 string = "\n\t"
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi21
            }
        default:
            for fi26 := 0; fi26 < 7; fi26++ {
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 uint32 = 128
                    v8 = (fi26 ^ v8)
                    _ = fi27
                    _ = v28
                }
                v29 := GenS3{F0: 1, F1: "\n\t"}
                _ = fi26
                _ = v29
            }
        }
        var v30 bool = (2 >= (v6 | (v6 - v6)))
        var v31 int16 = 868
        _ = v30
        _ = v31
    }
    fmt.Println(v6, v4)
    var v32 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
