package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int16 }
func genFn0(p0 string, p1 uint8, p2 int8) uint {
    v5 := []int{127, 128, -243}
    var v6 int = 127
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: (p2 >> ((4294967295) & 63)), F1: 389, F2: 1000}
        _ = fi7
        _ = v8
    }
    p1 = 151
    var v9 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v9
    return 255
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]int)
    var v3 uint8 = 47
    var v4 uint16 = 42
    var v10 uint64 = 18446744073709551615
    for fi11 := 0; fi11 < 12; fi11++ {
        v4 = v4
        v2[883] = fi11
        _ = fi11
    }
    if (42 > 519) {
        var v12 string = (("\n\t" + ("hello" + (((("a" + "hello") + "a") + ("\n\t" + "hello")) + ""))) + "")
        var v13 int = 2147483647
        v14 := []int{127, v13}
        _ = v12
        _ = v13
        _ = v14
    } else {
        v15 := []int{2147483647, 9223372036854775807, 483, -2147483648}
        _ = v15
    }
    var v16 int32 = -1
    v4 += uint16(v3)
    var v17 int64 = 0
    v18 := []int{-106, 177}
    v10 = uint64(v1)
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
