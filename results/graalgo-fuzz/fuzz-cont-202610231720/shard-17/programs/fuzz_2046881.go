package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 uint64, p1 int, p2 int16) uint64 {
    var v5 uint = 65535
    fmt.Println(p1)
    var v6 uint = (v5 * v5)
    if (-9223372036854775808 != -819) {
        var v7 int8 = int8(p2)
        var v8 int8 = 42
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ""
    var v3 int64 = -511
    var v4 uint64 = 4294967295
    var v9 int16 = 32767
    v10 := GenS1{F0: 42, F1: v2}
    var v11 int16 = v9
    v4 = genFn0(v4, -1, v9)
    if false {
        for fi12 := 0; fi12 < 20; fi12++ {
            fmt.Println(v4, fi12, v4)
            _ = fi12
        }
        fmt.Println(v3, v4, v2)
    }
    v11 -= (v11 << ((0) & 63))
    if true {
        v2 = "\n\t"
    } else {
        v4 += 1
        v4 = v4
        v13 := GenS1{F0: 551, F1: v2}
        _ = v13
    }
    v14 := GenS1{F0: 32767, F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
