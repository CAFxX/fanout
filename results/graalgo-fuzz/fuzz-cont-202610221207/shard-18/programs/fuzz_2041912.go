package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 int32, p1 uint32, p2 uint8) uint64 {
    var v4 uint8 = p2
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 string = ("foo bar" + "hello")
        v7 := GenS1{F0: (v4 ^ p2), F1: v4}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 212
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 1
    var v3 int = 128
    delete(v1, "\n\t")
    var v8 uint32 = (v2 >> ((65535) & 63))
    fmt.Println(v8)
    switch (v3 / ((v3) | 1)) {
    case 127:
        fmt.Println(v8)
    case -1:
        var v9 int = v3
        _ = v9
    default:
        var v10 float32 = 1.5
        _ = v10
    }
    v11 := GenS2{F0: -9223372036854775808}
    v8 -= 2147483648
    for fi12 := 0; fi12 < 5; fi12++ {
        if true {
            fmt.Println(v2)
        } else {
            v1[""] = false
            if (196 != 42) {
                var v13 float64 = ((-2.25 / 1.0) - (-2.25 + (0.1 * -2.25)))
                fmt.Println(v2, v3)
                _ = v13
            }
        }
        var v14 float32 = 1e10
        _ = fi12
        _ = v14
    }
    v8 = v8
    v15 := GenS3{F0: 100, F1: 255}
    v16 := []int{}
    fmt.Println(v8, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
