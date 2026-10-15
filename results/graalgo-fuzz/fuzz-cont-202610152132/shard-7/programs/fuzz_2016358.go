package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }
func genFn0(p0 int64, p1 uint32, p2 uint8) bool {
    var v6 int32 = 1
    if true {
        var v7 int = 0
        if false {
            var v8 float32 = (3.14159 / 3.14159)
            p2 = (p2 | p2)
            _ = v8
        }
        fmt.Println(p0)
        _ = v7
    }
    var v9 float64 = (0.0 - 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return (!(751 <= 32767))
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 string = ("a" + "hello")
    var v4 string = "\n\t"
    v5 := []uint32{2147483647, 1, 0, 0}
    if false {
        var v10 uint32 = 65535
        _ = v10
    }
    fmt.Println(v3)
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            _ = len(v2)
            var v13 int32 = 42
            v4 = v3
            switch v4 {
            case "hello":
                v14 := GenS1{F0: -1}
                var v15 float64 = 0.1
                fmt.Println(v3, v13)
                v3 = ((v4 + "\n\t") + v3)
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
                _ = v14
                _ = v15
            case "a":
                var v16 int64 = 1
                _ = v16
            }
            v13 = -655
            _ = fi12
            _ = v13
        }
        v3 = "a"
        _ = fi11
    }
    fmt.Println(v3)
    var v17 int = 4294967295
    v17 = (v17 >> ((2) & 63))
    fmt.Println(v4)
    if len(v5) > 0 {
        v5[0] = 2147483648
    }
    fmt.Println(v17, v17)
    var v18 bool = (33 >= 0)
    v19 := v2[271]
    switch v17 {
    case -2147483648:
        var v20 uint = 491
        _ = v20
    }
    if v18 {
        v17 -= (v17 % (((^(v17))) | 1))
        var v21 uint64 = 42
        _ = v21
    } else {
        var v22 float64 = (0.1 + -2.25)
        _ = v22
    }
    if true {
        v23 := GenS2{F0: 128, F1: 526, F2: (-2147483648 >= 0)}
        _ = v23
    }
    v19 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
