package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0() int8 {
    var v7 int16 = 1000
    v8 := GenS1{F0: 42}
    if (9223372036854775807 > 927) {
        v7 -= v7
        v7 += (-(v7))
    }
    v7 -= -273
    _ = v7
    _ = v8
    return 81
}

func main() {
    v1 := []uint32{}
    var v2 uint = 65535
    var v3 bool = (0 <= 181)
    v4 := make(map[uint32]bool)
    var v5 int8 = -1
    v6 := []string{"", "", ("foo bar" + ((("hello" + ((("\n\t" + ("hello" + "")) + "foo bar") + ("foo bar" + "\n\t"))) + "a") + ""))}
    v2 += 65535
    var v9 uint8 = 64
    v4[0] = (!false)
    if v3 {
        var v10 bool = false
        _ = v10
    }
    var v11 string = ""
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int32 = 330
        var v14 int64 = 42
        v15 := GenS1{F0: (v2 - (v2 >> ((uint(v9)) & 63)))}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    if (uint64(v9) <= 42) {
        v5 = (^(v5))
    }
    v16 := GenS2{F0: v3, F1: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
