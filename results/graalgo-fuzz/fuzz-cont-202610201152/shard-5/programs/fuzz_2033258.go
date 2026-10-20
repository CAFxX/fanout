package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []int{}
    v2 := []string{"\n\t", "\n\t", ("foo bar" + "hello")}
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v5 bool = false
    var v6 int32 = 270
    v3[-9223372036854775808] = 1
    switch 127 {
    case 875:
        var v7 uint16 = 65535
        _ = v7
    case 128:
        switch -519 {
        case 32768:
            switch "a" {
            case "hello":
                v8 := GenS1{F0: (!v5), F1: 1, F2: false}
                _ = v8
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = 1
                }
            default:
                v9 := GenS1{F0: true, F1: 881, F2: true}
                _ = v9
            }
        default:
            var v10 bool = true
            _ = v10
        }
    }
    var v11 uint64 = 1
    fmt.Println(v11, v5, v5)
    fmt.Println(v11, v6)
    v12 := GenS2{F0: 0}
    if (!(!v5)) {
        switch 158 {
        case 32767:
            var v13 string = "\n\t"
            _ = v13
        default:
            var v14 float64 = -2.25
            _ = v14
        }
    } else {
        var v15 uint16 = 398
        v16 := GenS3{F0: 0}
        _ = v15
        _ = v16
    }
    v11 -= (v11 + (v11 | 18446744073709551615))
    fmt.Println(v6, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
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
