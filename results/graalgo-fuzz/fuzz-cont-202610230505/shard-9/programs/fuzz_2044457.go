package main

import (
    "fmt"
)

func genFn0(p0 int64, p1 int8, p2 uint16) int {
    var v5 int = 32767
    var v6 int64 = (p0 << ((822) & 63))
    v7 := []int{(v5 << ((42) & 63))}
    var v8 int8 = (-(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]int)
    var v3 uint8 = 127
    var v4 uint32 = 4294967295
    v4 += v4
    var v9 int8 = 1
    switch "foo bar" {
    case "a":
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint16 = 2
            switch (^(fi10)) {
            case -9223372036854775808:
                switch "hello" {
                case "\n\t":
                    if ((1.5 * (0.0 / (3.14159 + ((1e10 + -2.25) / float32(((((v1 * float64(((1.5 - 3.14159) - float32((v1 - (v1 / 1.0)))))) + -2.25) + v1) - float64(0.0))))))) != -2.25) {
                        var v12 bool = (float64((0.0 - 3.14159)) < v1)
                        _ = v12
                    }
                case "a":
                    var v13 uint16 = (v11 >> ((42) & 63))
                    _ = v13
                case "":
                    v14 := []int{genFn0(690, v9, v11)}
                    _ = v14
                default:
                    fmt.Println(v3, fi10, v3)
                }
            case -518:
                var v15 int32 = 1
                _ = v15
            }
            _ = fi10
            _ = v11
        }
    case "\n\t":
        v3 = v3
    case "":
        var v16 int64 = 0
        _ = v16
    default:
        var v17 int32 = -473
        _ = v17
    }
    v1 -= v1
    v9 -= 42
    v4 -= v4
    v1 = 3.14159
    fmt.Println(v9, v3)
    v4 -= uint32(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
