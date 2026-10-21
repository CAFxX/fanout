package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint32, p1 string) uint8 {
    if (128 >= 1) {
        if (836 > 0) {
            if (int8(p0) == 100) {
                v7 := []int{549, 0, 32768}
                var v8 int32 = 2147483647
                var v9 float64 = float64(float32((3.14159 * (0.0 + 3.14159))))
                _ = v7
                _ = v8
                _ = v9
            }
            if (true || (true || (p1 != "a"))) {
                v10 := []int{0, 32768, -808, -9223372036854775808}
                if false {
                    var v11 int8 = -1
                    p1 = (p1 + p1)
                    _ = v11
                }
                _ = v10
            }
        }
    } else {
        var v12 uint32 = 115
        var v13 uint16 = 0
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 0, F1: 65535}
    _ = p0
    _ = p1
    _ = v14
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 2
    var v3 float32 = (float32((((float64(-2.25) - 3.14159) + ((0.1 * (1.5 / 0.1)) * 0.1)) / 3.14159)) - 1.5)
    v4 := []uint32{310, 0, 2147483648, 42}
    var v5 int = -9223372036854775808
    var v6 uint16 = 255
    v5 = (v5 >> ((65535) & 63))
    switch "foo bar" {
    case "":
        v2 += (^(v2))
    case "\n\t":
        switch (v5 ^ v5) {
        case 499:
            for fi15 := 0; fi15 < 4; fi15++ {
                fmt.Println(fi15, v6, v5)
                _ = fi15
            }
        case 4294967295:
            _ = len(v4)
        case 139:
            v2 += v2
        }
    case "hello":
        fmt.Println(v3, v2)
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        v16 := GenS1{F0: 57, F1: 2}
        _ = v16
    }
    if false {
        delete(v1, 1)
        v1[2147483647] = v5
        v3 = (v3 * (3.14159 - v3))
    } else {
        v17 := GenS2{F0: 2}
        v5 += v5
        _ = v17
    }
    var v18 int8 = -72
    var v19 int = -1
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    fmt.Println(v19)
    for fi20 := 0; fi20 < 15; fi20++ {
        v3 = (v3 - (v3 / v3))
        v6 -= (v6 + v6)
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
