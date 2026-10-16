package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 bool = false
    v4 := []int{}
    _ = len(v4)
    v5 := GenS1{F0: 128, F1: 709}
    fmt.Println(v3)
    v3 = (!(637 <= 42))
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    v3 = false
    var v6 uint16 = 0
    if v3 {
        v3 = (42 > 1)
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint8 = 2
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int32 = -729
            v12 := GenS2{F0: "\n\t"}
            _ = fi10
            _ = v11
            _ = v12
        }
        if v3 {
            v3 = (v3 || v3)
            v3 = (9223372036854775807 <= -9223372036854775808)
        }
        _ = fi7
    }
    _ = len(v1)
    fmt.Println(v3, v6)
    if v3 {
        v6 = v6
        fmt.Println(v3)
    } else {
        v4 = append(v4, 2147483647)
        if v3 {
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 float64 = 1.5
                v15 := GenS3{F0: -32768}
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            v2["foo bar"] = 2147483647
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 uint32 = 65535
                var v18 uint64 = 18446744073709551615
                _ = fi16
                _ = v17
                _ = v18
            }
        }
    }
    v3 = false
    switch -854 {
    case 0:
        v19 := GenS3{F0: -193}
        _ = v19
    default:
        switch 128 {
        case 1:
            var v20 uint64 = 18446744073709551615
            _ = v20
        default:
            for fi21 := 0; fi21 < 11; fi21++ {
                fmt.Println(v6, v6, v6)
                switch ("hello" + "foo bar") {
                case "hello":
                    fmt.Println(fi21, v3, fi21)
                case "foo bar":
                    var v22 float64 = 0.0
                    _ = v22
                }
                _ = fi21
            }
        }
    }
    var v23 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
