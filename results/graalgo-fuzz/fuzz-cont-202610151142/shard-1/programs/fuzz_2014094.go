package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }
type GenS2 struct { F0 string; F1 uint8 }
func genFn0(p0 bool, p1 bool) int8 {
    v5 := []int{127, -2147483648, 9223372036854775807, 913}
    if (!true) {
        fmt.Println(p0, p1)
        switch "foo bar" {
        case "":
            for fi6 := 0; fi6 < 16; fi6++ {
                if ((1e10 + 0.1) == 1.5) {
                    fmt.Println(p1)
                }
                v5 = append(v5, -2147483648)
                p0 = true
                _ = fi6
            }
        case "\n\t":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "a":
            var v7 int = -1
            _ = v7
        }
    }
    v5 = append(v5, 128)
    var v8 uint8 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 127
}

func main() {
    v1 := []uint32{353, 506}
    var v2 uint32 = 364
    v3 := make(map[uint32]bool)
    var v4 float32 = ((-2.25 / 3.14159) + float32(0.1))
    fmt.Println(v2)
    switch int(v2) {
    case 127:
        v1 = append(v1, v2)
    case 0:
        fmt.Println(v4, v4)
    case 4294967295:
        var v9 uint8 = 1
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint16 = 547
        fmt.Println(v4)
        _ = fi10
        _ = v11
    }
    var v12 float32 = v4
    _ = len(v3)
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 = 1.5
        _ = fi13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v2 = 644
        _ = fi14
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = 738
    fmt.Println(v12, v2)
    v15 := GenS1{F0: -1, F1: false, F2: 42}
    fmt.Println(v4, v12, v12)
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 uint64 = 385
        _ = fi16
        _ = v17
    }
    var v18 uint = 123
    var v19 float32 = ((v4 * v4) / (v4 - v4))
    fmt.Println(v19)
    v20 := GenS2{F0: "foo bar", F1: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
