package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int, p1 int64, p2 uint32) uint16 {
    if true {
        v7 := GenS1{F0: 42, F1: (p1 + (p1 - p1)), F2: 255}
        _ = v7
    } else {
        p2 -= p2
    }
    p0 += p0
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: 1000, F1: (p1 * p1), F2: 255}
        var v10 float32 = (1.5 * (1.5 - (-2.25 - float32(((1.5 - 1e10) + 1.5)))))
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1() uint {
    var v11 uint = 4294967295
    v11 -= 563
    v12 := GenS2{F0: true}
    _ = v11
    _ = v12
    return (v11 - (v11 * v11))
}

func main() {
    var v1 int64 = 42
    var v2 uint8 = 127
    v3 := []string{"foo bar", "hello"}
    v4 := make(map[int]bool)
    var v5 int8 = 1
    var v6 int64 = -1
    v3 = append(v3, "\n\t")
    for fi13 := 0; fi13 < 12; fi13++ {
        _ = len(v4)
        v14 := GenS2{F0: false}
        _ = fi13
        _ = v14
    }
    var v15 int = -1
    v15 += (v15 ^ (v15 << ((4294967295) & 63)))
    for fi16 := 0; fi16 < 10; fi16++ {
        v5 = v5
        v2 = v2
        var v17 int64 = (v6 << ((65535) & 63))
        _ = fi16
        _ = v17
    }
    var v18 uint64 = 204
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 uint32 = 755
        var v21 int8 = -51
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: -1, F1: (v6 * (^(v1))), F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
