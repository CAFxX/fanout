package main

import (
    "fmt"
)

func genFn0(p0 int16, p1 uint8, p2 uint8) int32 {
    switch "foo bar" {
    case "":
        for fi4 := 0; fi4 < 17; fi4++ {
            p1 += 255
            var v5 uint16 = 430
            _ = fi4
            _ = v5
        }
    case "a":
        var v6 float64 = 1e10
        _ = v6
    case "hello":
        fmt.Println(p2)
    }
    var v7 int8 = -128
    v8 := []int{4294967295}
    fmt.Println(p0, v7)
    p1 += 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int = -8
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v9 float64 = 0.1
    var v10 int32 = genFn0(42, 1, 42)
    fmt.Println(v10, v9)
    v9 = (v9 * (float64(0.1) * (v9 - (v9 / (1e10 - ((1.5 - ((v9 + v9) / (0.1 - (v9 / (v9 - ((1.5 + (float64(3.14159) - v9)) + 1.5)))))) - ((float64(float32(v9)) / 3.14159) * v9)))))))
    for fi11 := 0; fi11 < 15; fi11++ {
        switch -2147483648 {
        case -788:
            v1 = 128
        default:
            v12 := []int{v1, v1, fi11, fi11}
            _ = v12
        }
        v9 += v9
        _ = fi11
    }
    v9 = v9
    v10 += v10
    v13 := []int{(v1 >> ((4294967295) & 63)), (v1 ^ v1), (v1 << ((548) & 63))}
    v10 = v10
    switch "a" {
    case "a":
        v9 += (v9 - (0.0 + v9))
    default:
        var v14 uint32 = 65535
        _ = v14
    }
    fmt.Println(v9, v9)
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint64 = 2
        v16 = v16
        if false {
            var v17 float64 = (v9 / v9)
            fmt.Println(v1, v10, v17)
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
