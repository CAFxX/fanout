package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }
type GenS2 struct { F0 string }

func main() {
    v1 := []string{}
    v2 := []string{}
    var v3 uint8 = 185
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    _ = len(v5)
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(fi6, v3)
        fmt.Println(fi6, fi6)
        fmt.Println(fi6, v3)
        _ = fi6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 2, F1: 0, F2: 65535}
        v9 := GenS1{F0: 166, F1: 750, F2: 255}
        var v10 uint8 = (^(v3))
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 bool = ((!(65535 <= 883)) && true)
    var v12 uint64 = uint64(v3)
    for fi13 := 0; fi13 < 7; fi13++ {
        if v11 {
            var v14 uint8 = (v3 << ((310) & 63))
            var v15 float32 = float32(1.5)
            _ = v14
            _ = v15
        } else {
            switch "" {
            case "":
                var v16 float64 = 0.0
                _ = v16
            case "foo bar":
                v17 := GenS1{F0: 2147483648, F1: 884, F2: 42}
                _ = v17
            default:
                if v11 {
                    var v18 uint16 = 65535
                    _ = v18
                }
            }
            var v19 uint32 = 627
            if v11 {
                var v20 uint16 = 255
                _ = v20
            }
            _ = v19
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS1{F0: uint32(v3), F1: 0, F2: 1}
            var v23 bool = ((fi13 ^ fi21) > (fi21 * 0))
            v1 = append(v1, "a")
            _ = fi21
            _ = v22
            _ = v23
        }
        v12 = v12
        _ = fi13
    }
    v12 = v12
    fmt.Println(v11)
    v24 := GenS1{F0: 1, F1: 4294967295, F2: 42}
    var v25 int8 = 0
    v26 := GenS2{F0: "foo bar"}
    fmt.Println(v3)
    v11 = v11
    var v27 float64 = -2.25
    v12 += (v12 % ((v12) | 1))
    if true {
        v28 := GenS1{F0: 2147483647, F1: 777, F2: uint(v3)}
        _ = v28
    } else {
        fmt.Println(v25, v27, v11)
    }
    v29 := GenS1{F0: 2147483648, F1: 42, F2: 357}
    v12 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
