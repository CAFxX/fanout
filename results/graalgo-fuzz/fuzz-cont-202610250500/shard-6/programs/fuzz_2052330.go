package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 float64 = 0.1
    v4 := []string{}
    var v5 int64 = 9223372036854775807
    if true {
        switch "" {
        case "foo bar":
            var v6 int = 2147483647
            _ = v6
        case "":
            v7 := GenS1{F0: 289, F1: ("a" + "\n\t")}
            _ = v7
        default:
            for fi8 := 0; fi8 < 11; fi8++ {
                v9 := GenS2{F0: 1000, F1: 0}
                var v10 bool = ((((0.1 / 3.14159) / ((0.1 * 1.5) - float32((v3 / (v3 / float64((1e10 + 1.5))))))) + 0.0) <= 1e10)
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    }
    if false {
        var v11 float64 = 1e10
        v4 = append(v4, "")
        var v12 int8 = -1
        _ = v11
        _ = v12
    } else {
        var v13 uint = 787
        v5 = v5
        _ = v13
    }
    var v14 int16 = 1000
    fmt.Println(v5, v5)
    var v15 bool = (32767 == 786)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint64 = 766
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 2
    var v19 uint8 = v18
    var v20 float64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
