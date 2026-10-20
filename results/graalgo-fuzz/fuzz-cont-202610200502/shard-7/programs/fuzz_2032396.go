package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
func genFn0(p0 bool, p1 int) uint64 {
    switch "foo bar" {
    case "":
        v6 := []int{(p1 >> ((65535) & 63)), -9223372036854775808, (p1 << ((65535) & 63))}
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 2; fi7++ {
            p0 = p0
            p0 = false
            var v8 bool = (11 < 937)
            _ = fi7
            _ = v8
        }
    }
    var v9 int = -192
    v10 := []int{}
    fmt.Println(p1, p0, p0)
    p0 = p0
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int8 = 2
    var v2 int8 = (v1 - (v1 % (((v1 ^ v1)) | 1)))
    var v3 uint64 = 18446744073709551615
    var v4 int8 = v1
    v5 := make(map[int]int)
    var v11 float32 = 0.1
    _ = len(v5)
    if true {
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := []int{-187, (-(fi12)), fi12, fi12}
            _ = fi12
            _ = v13
        }
        fmt.Println(v2)
    }
    if (genFn0((32767 == 65535), 1) < (^(v3))) {
        v14 := GenS1{F0: 1, F1: uint16(v3)}
        _ = v14
    } else {
        var v15 uint32 = uint32(v3)
        _ = v15
    }
    var v16 float32 = 0.1
    v17 := []int{}
    v1 -= v4
    if len(v17) > 0 {
        v17[0] = 1
    }
    _ = len(v5)
    v2 -= (v4 >> ((1) & 63))
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
