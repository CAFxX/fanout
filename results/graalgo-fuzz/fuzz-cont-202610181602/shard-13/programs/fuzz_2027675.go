package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint, p1 bool, p2 uint16) uint64 {
    p2 -= 255
    var v5 uint16 = (p2 << (((p0 * p0)) & 63))
    var v6 float64 = 1e10
    p0 = (p0 + 18446744073709551615)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 1
    v3 := make(map[int]string)
    var v4 float32 = (3.14159 - ((1e10 + float32((3.14159 - (-2.25 + 0.1)))) / (((1e10 / (-2.25 * 0.1)) * (1e10 - 0.1)) + 1.5)))
    v2 = v2
    if true {
        v4 = (((3.14159 * v4) + v4) + (v4 + v4))
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int64 = -437
            _ = fi7
            _ = v8
        }
    } else {
        v9 := []int{4294967295, 128, 42}
        v2 = v2
        _ = v9
    }
    v4 = (v4 / v4)
    if (386 <= 688) {
        switch -1 {
        case 0:
            fmt.Println(v4)
        case -1:
            var v10 uint64 = 665
            _ = v10
        case -9223372036854775808:
            v11 := GenS1{F0: 42}
            _ = v11
        }
    } else {
        var v12 uint16 = 615
        _ = v12
    }
    v13 := GenS1{F0: 88}
    if (3.14159 == (0.0 - v4)) {
        v2 = v2
        fmt.Println(v4, v4)
        for fi14 := 0; fi14 < 19; fi14++ {
            fmt.Println(v4, fi14)
            switch fi14 {
            case -571:
                if true {
                    v15 := GenS1{F0: 255}
                    v2 += v2
                    v2 = v2
                    _ = v15
                } else {
                    var v16 int32 = 2147483647
                    fmt.Println(v2)
                    fmt.Println(v4)
                    _ = v16
                }
            default:
                v17 := []int{(fi14 >> ((1) & 63))}
                _ = v17
            }
            v2 -= (^(v2))
            _ = fi14
        }
    }
    switch "foo bar" {
    case "foo bar":
        var v18 bool = true
        _ = v18
    default:
        _ = len(v3)
    }
    if (9223372036854775807 <= 127) {
        delete(v1, -908)
    }
    switch 354 {
    case 4294967295:
        v4 = float32(3.14159)
    }
    switch (("" + "") + "hello") {
    case "hello":
        var v19 int64 = -9223372036854775808
        _ = v19
    case "\n\t":
        var v20 int64 = -300
        _ = v20
    case "a":
        if (!(!true)) {
            fmt.Println(v2, v2)
        }
    default:
        var v21 int8 = 29
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
