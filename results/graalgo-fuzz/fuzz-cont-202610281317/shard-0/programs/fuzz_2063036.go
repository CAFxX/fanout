package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -65
    v3 := make(map[uint32]string)
    var v4 int8 = v2
    v2 = v4
    v5 := []int{128, 2147483647, 690, 9223372036854775807}
    fmt.Println(v4)
    fmt.Println(v4, v2)
    var v6 bool = (4294967295 == 42)
    fmt.Println(v2, v2)
    v2 = v4
    if v6 {
        v7 := GenS1{F0: 42, F1: -494}
        if (v2 >= v4) {
            v1[42] = (!false)
        } else {
            for fi8 := 0; fi8 < 17; fi8++ {
                _ = len(v5)
                _ = fi8
            }
        }
        v9 := GenS1{F0: 287, F1: 127}
        _ = v7
        _ = v9
    } else {
        var v10 int16 = 98
        _ = v10
    }
    v11 := GenS2{F0: ""}
    var v12 float32 = 1.5
    var v13 int = 1
    switch "foo bar" {
    case "a":
        v14 := GenS2{F0: "hello"}
        _ = v14
    default:
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 int16 = -988
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi15
            _ = v16
        }
    }
    fmt.Println(v4)
    fmt.Println(v13, v12)
    v2 -= (v2 | v2)
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
