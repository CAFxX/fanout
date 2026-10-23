package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 string = "a"
    var v4 int8 = -128
    var v5 float64 = (0.1 * 0.1)
    fmt.Println(v4, v4, v4)
    switch -728 {
    case 2147483647:
        v3 = (v3 + (v3 + ("foo bar" + v3)))
    case 0:
        var v6 uint64 = 98
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            switch v3 {
            case "":
                fmt.Println(fi8, fi8)
                v9 := GenS1{F0: 32767}
                _ = v9
            case "hello":
                var v10 uint64 = 18446744073709551615
                _ = v10
            case "a":
                var v11 bool = true
                _ = v11
            default:
                var v12 uint32 = 0
                _ = v12
            }
            var v13 float32 = 1e10
            v14 := GenS1{F0: (fi7 % ((fi7) | 1))}
            _ = fi8
            _ = v13
            _ = v14
        }
        switch (v3 + v3) {
        case "\n\t":
            delete(v2, 2147483648)
        case "hello":
            var v15 int16 = 42
            _ = v15
        case "foo bar":
            v5 += (float64(float32(1.5)) / v5)
        default:
            if false {
                v5 = v5
                var v16 int8 = 42
                v3 = (v3 + v3)
                _ = v16
            }
        }
        _ = fi7
    }
    if false {
        if (true && false) {
            var v17 float64 = ((v5 + (v5 + v5)) + ((v5 - ((1e10 * v5) * 0.0)) - (v5 - 1.5)))
            _ = v17
        }
        v18 := []int{-1, 128}
        delete(v1, 127)
        _ = v18
    }
    v19 := GenS1{F0: -215}
    v5 = v5
    v20 := GenS1{F0: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
