package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint }
type GenS2 struct { F0 uint }
func genFn0(p0 uint64) uint {
    var v4 int = 1
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint8 = 85
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 uint16 = 306
    v3 += 32767
    fmt.Println(v3)
    var v7 int16 = -759
    fmt.Println(v7, v7)
    v8 := GenS1{F0: -99, F1: (v7 | (v7 / ((v7) | 1))), F2: 652}
    fmt.Println(v3, v3)
    v3 += (^(v3))
    if (!false) {
        v3 = 1
    }
    var v9 uint8 = 82
    v10 := GenS2{F0: 18446744073709551615}
    v3 -= v3
    var v11 int = -1
    var v12 bool = (!false)
    _ = len(v2)
    var v13 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
