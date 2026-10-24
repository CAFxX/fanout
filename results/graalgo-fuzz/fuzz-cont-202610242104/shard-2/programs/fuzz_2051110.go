package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    var v3 int = 453
    v4 := make(map[string]uint32)
    var v5 uint32 = 42
    v2 = append(v2, (v5 - v5))
    switch "" {
    case "foo bar":
        delete(v4, "a")
    }
    v5 -= v5
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int32 = 128
        for fi8 := 0; fi8 < 15; fi8++ {
            v3 = (v3 | fi8)
            v9 := v4["\n\t"]
            v10 := GenS1{F0: 18446744073709551615, F1: "hello", F2: -77}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int16 = -32768
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v5 = (^(v5))
        _ = fi13
    }
    var v14 uint32 = v5
    var v15 int8 = -1
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
