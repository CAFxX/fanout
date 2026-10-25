package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() uint {
    if true {
        var v4 uint64 = 2
        fmt.Println(v4)
        _ = v4
    } else {
        var v5 int64 = -9223372036854775808
        var v6 uint8 = 2
        switch "hello" {
        case "":
            fmt.Println(v5, v6)
        case "foo bar":
            if (1000 >= -32768) {
                v6 = v6
            }
        }
        _ = v5
        _ = v6
    }
    var v7 int16 = -83
    _ = v7
    return 42
}
func genFn1(p0 uint64, p1 int8, p2 uint32) int16 {
    p1 -= p1
    p2 -= p2
    for fi8 := 0; fi8 < 5; fi8++ {
        p0 = p0
        var v9 uint8 = 222
        v10 := []int{fi8, fi8, 0}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 bool = true
    var v12 int8 = (p1 >> ((93) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (!(42 <= 2))
    var v3 uint16 = 2
    v13 := GenS1{F0: 1}
    var v14 float32 = 3.14159
    switch "a" {
    case "\n\t":
        switch -1 {
        case 128:
            v3 = v3
        case -602:
            v2 = (!v2)
        default:
            var v15 uint32 = 1
            _ = v15
        }
    case "hello":
        switch "hello" {
        case "":
            _ = len(v1)
        case "foo bar":
            for fi16 := 0; fi16 < 8; fi16++ {
                v17 := []int{}
                v18 := v1[9223372036854775807]
                switch (fi16 % ((127) | 1)) {
                case 71:
                    if len(v17) > 0 {
                        v17[0] = fi16
                    }
                }
                _ = fi16
                _ = v17
                _ = v18
            }
        case "a":
            var v19 float64 = (1.5 * 3.14159)
            _ = v19
        default:
            fmt.Println(v3, v3)
        }
    }
    _ = len(v1)
    var v20 int64 = -1
    v21 := GenS1{F0: 9223372036854775807}
    var v22 uint64 = 4294967295
    var v23 float32 = (0.0 / (0.1 * (v14 * (float32(0.0) / v14))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
