package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint16) bool {
    v4 := GenS1{F0: -128}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v4
    return false
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    v3 := []int{32768}
    var v5 int8 = 127
    v3 = append(v3, -9223372036854775808)
    v6 := GenS1{F0: (v5 - v5)}
    switch "\n\t" {
    case "a":
        var v7 int32 = -1
        _ = v7
    case "":
        fmt.Println(v5)
    case "foo bar":
        fmt.Println(v5)
    default:
        if false {
            fmt.Println(v5)
        } else {
            v8 := GenS2{F0: (v5 * v5)}
            _ = v8
        }
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    v9 := GenS3{F0: 1000}
    var v10 uint32 = 376
    var v11 uint8 = 255
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int64 = 9223372036854775807
        v10 = (v10 ^ 2147483648)
        _ = fi12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
