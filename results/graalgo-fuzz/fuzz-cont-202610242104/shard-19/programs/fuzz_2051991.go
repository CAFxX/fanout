package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint16 }
func genFn0(p0 uint32, p1 uint, p2 int) int8 {
    if (((0.1 - 1.5) == float32((0.0 + float64((0.1 - float32(((0.1 - 0.1) * 0.1))))))) && (4294967295 >= 4294967295)) {
        v6 := []int{(p2 * int(p0)), p2, 128, 0}
        v6 = append(v6, p2)
        _ = v6
    } else {
        fmt.Println(p1, p0)
    }
    if false {
        if true {
            switch "" {
            case "":
                var v7 int8 = -125
                _ = v7
            }
            p0 = p0
        }
    }
    v8 := GenS1{F0: 574, F1: 0, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    v1 := []uint32{2147483647, 16, 2147483648, 1}
    v2 := []int{-2147483648, -527, 32767}
    v3 := make(map[uint32]int)
    var v4 uint8 = 255
    v5 := make(map[string]bool)
    fmt.Println(v4)
    v4 -= 2
    var v9 uint16 = 42
    switch 2147483647 {
    case 42:
        if true {
            v1 = append(v1, 445)
        }
    case 32767:
        var v10 uint16 = (^(v9))
        _ = v10
    case 644:
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 8; fi12++ {
                v9 -= (v9 << ((65535) & 63))
                _ = fi12
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi11
        }
    }
    v13 := GenS2{F0: 4294967295, F1: 2, F2: 255}
    fmt.Println(v9, v9)
    if false {
        _ = len(v5)
        switch 926 {
        case 2147483647:
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint64 = 0
                _ = fi14
                _ = v15
            }
        }
        _ = len(v3)
    }
    v3[4294967295] = 2147483647
    var v16 string = "hello"
    var v17 int16 = -690
    var v18 uint = uint(v9)
    v18 -= 623
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
