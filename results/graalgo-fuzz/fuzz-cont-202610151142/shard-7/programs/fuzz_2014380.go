package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 int32, p1 uint8) string {
    p0 = 1
    p0 -= p0
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]int)
    var v3 int16 = -1
    v4 := []uint32{1}
    v3 += (v3 + (v3 | (v3 % ((0) | 1))))
    v3 = v3
    for fi5 := 0; fi5 < 2; fi5++ {
        switch (v1 ^ (v1 ^ (fi5 / ((v1) | 1)))) {
        case -1:
            v1 -= v1
        case -2147483648:
            var v6 int8 = -128
            _ = v6
        default:
            v3 = (-(v3))
        }
        _ = len(v4)
        var v7 string = "a"
        _ = fi5
        _ = v7
    }
    var v8 uint = 65535
    v8 -= v8
    var v9 uint64 = 2
    if ((v3 ^ v3) == 203) {
        var v10 uint8 = 0
        var v11 uint16 = 387
        _ = v10
        _ = v11
    }
    switch v1 {
    case -9223372036854775808:
        fmt.Println(v1, v8)
    case 42:
        switch "foo bar" {
        case "a":
            var v12 int16 = (^(v3))
            _ = v12
        default:
            for fi13 := 0; fi13 < 10; fi13++ {
                for fi14 := 0; fi14 < 20; fi14++ {
                    v9 += (v9 * 4294967295)
                    var v15 uint32 = 2147483648
                    fmt.Println(v15, v15, v9)
                    _ = fi14
                    _ = v15
                }
                var v16 uint8 = 42
                _ = fi13
                _ = v16
            }
        }
    }
    v17 := GenS1{F0: 2, F1: v8, F2: 42}
    var v18 string = ("\n\t" + "a")
    v19 := GenS1{F0: 255, F1: v8, F2: 2147483647}
    v20 := GenS1{F0: 0, F1: v8, F2: -2147483648}
    var v21 int64 = 0
    switch v1 {
    case 32768:
        var v22 uint16 = 42
        _ = v22
    case 615:
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
