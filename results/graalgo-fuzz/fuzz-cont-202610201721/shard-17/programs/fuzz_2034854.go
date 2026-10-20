package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint32) bool {
    p0 += 2147483647
    p0 -= (p0 * (^(p0)))
    v7 := []int{42}
    _ = p0
    _ = v7
    return (((float64(3.14159) + 0.0) * 1e10) < 1e10)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 4294967295
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    var v5 int16 = 1000
    var v6 uint64 = 18446744073709551615
    var v8 uint = (v2 % ((uint(v6)) | 1))
    v8 -= (v8 * 426)
    var v9 uint8 = 82
    switch "a" {
    case "":
        switch "hello" {
        case "foo bar":
            var v10 bool = true
            _ = v10
        }
    case "\n\t":
        var v11 float64 = ((-2.25 + -2.25) + float64(1e10))
        _ = v11
    case "foo bar":
        _ = len(v1)
    default:
        v6 = (v6 % (((^(v6))) | 1))
    }
    v6 += v6
    var v12 uint8 = (v9 >> ((v8) & 63))
    if (!true) {
        v13 := GenS1{F0: 209}
        if (!(false || ((false || (((1.5 * float32(3.14159)) * float32(0.1)) <= (-2.25 - (0.0 + 3.14159)))) || false))) {
            if len(v3) > 0 {
                v3[0] = 65535
            }
        }
        _ = v13
    }
    _ = len(v4)
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
