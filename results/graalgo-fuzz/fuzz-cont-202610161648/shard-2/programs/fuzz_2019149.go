package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    v4 := []uint32{2, 2147483648}
    var v5 uint16 = 65535
    var v6 string = "\n\t"
    if (v6 < (v6 + ("" + v6))) {
        var v7 int32 = -450
        for fi8 := 0; fi8 < 13; fi8++ {
            v5 -= v5
            _ = fi8
        }
        _ = v7
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            switch v6 {
            case "foo bar":
                v5 = 295
            case "hello":
                v5 = v5
            }
            fmt.Println(v5, fi9)
            _ = fi10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := v3[fi9]
            if ((-349 == -977) || v12) {
                var v13 uint = 0
                v5 = (v5 | v5)
                _ = v13
            } else {
                var v14 int32 = 0
                _ = v14
            }
            var v15 uint8 = 2
            _ = fi11
            _ = v12
            _ = v15
        }
        v16 := GenS1{F0: false}
        _ = fi9
        _ = v16
    }
    switch 0 {
    case 426:
        v5 = (^(v5))
    }
    if (true && ((v5 << ((42) & 63)) < 1)) {
        v17 := GenS2{F0: 42, F1: -1}
        v4 = append(v4, 0)
        _ = v17
    }
    var v18 uint32 = 1
    v6 = (v6 + "")
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
