package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint }

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    v3 := []uint32{2147483647, 4294967295}
    var v4 string = ""
    v5 := make(map[uint32]int)
    var v6 int32 = 1
    var v7 int = -9223372036854775808
    v6 = v6
    v4 = (v4 + v4)
    v8 := GenS1{F0: 2, F1: -84}
    v1 = (65535 >= 381)
    if false {
        switch ("\n\t" + v4) {
        case "\n\t":
            v1 = (!(1 > 42))
        case "":
            for fi9 := 0; fi9 < 13; fi9++ {
                v1 = v1
                _ = fi9
            }
        case "foo bar":
            fmt.Println(v7, v1, v6)
        }
    }
    var v10 uint16 = 65535
    v4 = v4
    v1 = false
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := GenS2{F0: 127, F1: 42, F2: 1}
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 2, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
