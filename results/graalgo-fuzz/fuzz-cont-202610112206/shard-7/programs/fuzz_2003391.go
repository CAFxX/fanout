package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 bool }
func genFn0(p0 uint64, p1 uint64, p2 uint32) int8 {
    v5 := GenS1{F0: -2147483648}
    v6 := GenS1{F0: 32768}
    p0 = 2
    v7 := []int{330, 191, 4294967295, -749}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -127
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 int = v1
    for fi8 := 0; fi8 < 15; fi8++ {
        if (0.0 <= 0.1) {
            v9 := GenS1{F0: (fi8 | v1)}
            for fi10 := 0; fi10 < 17; fi10++ {
                v4 += -679
                _ = fi10
            }
            var v11 float32 = (1.5 + (3.14159 * 1.5))
            _ = v11
            _ = v9
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            v1 -= (v4 & v1)
            var v13 uint8 = 42
            var v14 uint16 = 680
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 string = "hello"
        _ = fi8
        _ = v15
    }
    var v16 float64 = ((0.1 * ((3.14159 + (0.1 / 3.14159)) * ((3.14159 * 1e10) * 3.14159))) + ((0.0 * (((3.14159 * 0.0) - 1e10) - ((((0.0 * float64(((3.14159 * 0.0) * -2.25))) - 1.5) / ((3.14159 - 1e10) + 3.14159)) + 3.14159))) / 1e10))
    v16 -= v16
    v17 := []int{v1, v4, 9223372036854775807, (v1 / ((1) | 1))}
    v18 := GenS2{F0: "\n\t"}
    if (0 <= -850) {
        _ = len(v17)
    }
    fmt.Println(v16, v1, v1)
    v19 := v2[2147483648]
    switch ("foo bar" + (("\n\t" + ("a" + "foo bar")) + "")) {
    case "":
        var v20 float64 = v16
        _ = v20
    default:
        v21 := v3[("\n\t" + "")]
        _ = v21
    }
    v19 -= v19
    var v22 uint64 = 42
    v23 := GenS3{F0: 0, F1: (((186 < 4294967295) && false) || (4294967295 < v22))}
    var v24 string = "foo bar"
    switch v24 {
    case "hello":
        delete(v3, v24)
    default:
        for fi25 := 0; fi25 < 13; fi25++ {
            v16 += float64(0.0)
            _ = fi25
        }
    }
    var v26 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
