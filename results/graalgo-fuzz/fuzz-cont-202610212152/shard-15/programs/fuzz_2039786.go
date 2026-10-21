package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0() uint {
    var v6 uint64 = 553
    v6 += (v6 >> ((255) & 63))
    var v7 uint64 = (^(v6))
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 int, p1 uint16, p2 uint32) uint8 {
    p1 -= (p1 & (p1 ^ (p1 << ((42) & 63))))
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 127
}

func main() {
    var v1 float64 = 1e10
    var v2 uint32 = 2147483648
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]bool)
    var v9 uint64 = 285
    var v10 uint32 = (v2 + 4294967295)
    if true {
        v11 := GenS1{F0: (v2 << ((0) & 63)), F1: -128}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 2; fi12++ {
            v9 = 2
            if false {
                var v13 uint64 = v9
                v10 = v10
                _ = v13
            }
            switch ("\n\t" + ("" + (((((("" + "hello") + "foo bar") + "hello") + ("foo bar" + "\n\t")) + "\n\t") + "a"))) {
            case "hello":
                v9 = 748
            }
            _ = fi12
        }
        v4[v10] = v10
        var v14 int8 = 127
        _ = v14
    }
    var v15 bool = false
    var v16 uint16 = 2
    v17 := GenS1{F0: v2, F1: 1}
    _ = len(v5)
    switch -669 {
    case -9223372036854775808:
        v4[v10] = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
