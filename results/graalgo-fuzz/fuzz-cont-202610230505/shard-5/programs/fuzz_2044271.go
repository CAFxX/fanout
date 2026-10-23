package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 1
    v3 := make(map[int]bool)
    var v4 int16 = -639
    switch "a" {
    case "hello":
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 int = fi5
            _ = fi5
            _ = v6
        }
    }
    var v7 uint8 = 225
    if (!(0 == 9223372036854775807)) {
        if false {
            if (1 == 9223372036854775807) {
                v2 -= (v2 & (-(v2)))
                var v8 int8 = 42
                _ = v8
            }
            v9 := GenS1{F0: "hello"}
            _ = v9
        } else {
            switch 9223372036854775807 {
            case -179:
                var v10 uint64 = 529
                _ = v10
            }
        }
        v7 = (v7 & (v7 | v7))
    }
    v2 += (v2 - v2)
    v11 := []int{795, -9223372036854775808, 127, 32768}
    var v12 uint64 = 4294967295
    if false {
        var v13 string = "a"
        v14 := GenS1{F0: "a"}
        _ = v13
        _ = v14
    }
    v11 = append(v11, -266)
    v7 = (^(v7))
    var v15 string = "hello"
    v16 := GenS2{F0: v15, F1: uint64(v7)}
    v17 := GenS1{F0: v15}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
