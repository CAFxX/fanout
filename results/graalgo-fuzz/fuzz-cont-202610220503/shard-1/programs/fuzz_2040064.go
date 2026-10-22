package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (0.1 * 0.0)
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 uint16 = 2
    if false {
        var v6 string = ("foo bar" + "hello")
        var v7 int32 = 715
        var v8 int64 = -1
        _ = v6
        _ = v7
        _ = v8
    }
    switch -9223372036854775808 {
    case 42:
        var v9 bool = (32767 <= -189)
        _ = v9
    case -2147483648:
        switch "hello" {
        case "\n\t":
            v2 = (v2 - v2)
        case "hello":
            v2 += (((((v2 + v2) / v2) / v2) / float32(-2.25)) * v2)
        case "a":
            v10 := GenS1{F0: ""}
            _ = v10
        }
    }
    var v11 int64 = 579
    if (false || false) {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 float64 = (1.5 + 0.1)
            _ = fi12
            _ = v13
        }
        var v14 int8 = 100
        v15 := GenS1{F0: ""}
        _ = v14
        _ = v15
    }
    v11 = v11
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS1{F0: ("\n\t" + "hello")}
        var v18 float64 = (1.5 + float64(v2))
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := []int{9223372036854775807}
    v19 = append(v19, -642)
    v1[317] = 789
    v2 += (v2 - v2)
    v20 := GenS2{F0: (v5 ^ v5), F1: 4294967295}
    v5 = (v5 >> ((242) & 63))
    v2 = (0.0 - (v2 / float32((0.0 + 1e10))))
    v21 := GenS1{F0: (("hello" + ("\n\t" + "foo bar")) + "hello")}
    v2 = ((v2 - 3.14159) - (float32((1.5 + (((-2.25 - -2.25) * 1e10) * 3.14159))) * 1.5))
    fmt.Println(v5)
    fmt.Println(v11)
    v2 += -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
