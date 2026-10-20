package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 int16) int {
    var v6 float32 = 0.1
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float32 = (-2.25 + 1e10)
        v8 -= v8
        _ = fi7
        _ = v8
    }
    switch 127 {
    case 603:
        var v9 int8 = 127
        _ = v9
    case 9223372036854775807:
        var v10 int32 = -320
        _ = v10
    }
    _ = p0
    _ = v6
    return -703
}
func genFn1(p0 string, p1 uint8) uint16 {
    var v11 bool = false
    fmt.Println(v11)
    _ = p0
    _ = p1
    _ = v11
    return 65535
}

func main() {
    var v1 uint16 = 63
    var v2 uint64 = 0
    var v3 int16 = 0
    v4 := make(map[uint32]uint32)
    v5 := []string{"hello", "\n\t"}
    v12 := GenS1{F0: 2, F1: 42}
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS1{F0: 220, F1: 792}
        _ = fi13
        _ = v14
    }
    _ = len(v4)
    v5 = append(v5, "")
    var v15 uint = 65535
    for fi16 := 0; fi16 < 12; fi16++ {
        if (0.1 > 1.5) {
            v15 += (^(v15))
            v17 := GenS2{F0: false, F1: 736}
            _ = v17
        }
        v5 = append(v5, "foo bar")
        _ = fi16
    }
    v18 := GenS2{F0: true, F1: 9223372036854775807}
    var v19 uint32 = 42
    var v20 uint32 = (^(v19))
    v19 += 115
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
