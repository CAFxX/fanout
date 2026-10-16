package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32, p1 string) int {
    p0 -= (p0 - p0)
    v5 := []int{-744, 932, 1, -2147483648}
    var v6 uint8 = 158
    v7 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -798
}
func genFn1(p0 int16, p1 int64, p2 uint32) uint8 {
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: 18446744073709551615}
        p2 += p2
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: (p2 >> ((42) & 63))}
    p2 -= 546
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = float32(1.5)
    var v3 uint32 = 42
    v4 := make(map[int]bool)
    _ = len(v1)
    delete(v1, 2147483647)
    fmt.Println(v2)
    if (!(3.14159 <= -2.25)) {
        v11 := GenS1{F0: 42}
        v2 = (float32(1.5) / v2)
        v3 = v3
        _ = v11
    }
    if false {
        delete(v4, -1)
        v3 += v3
    } else {
        if false {
            var v12 int = -407
            _ = v12
        } else {
            var v13 string = "foo bar"
            _ = v13
        }
    }
    fmt.Println(v2)
    switch ("foo bar" + "hello") {
    case "hello":
        _ = len(v4)
    }
    v2 -= v2
    v3 += v3
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(v2, v2)
        var v15 int8 = 100
        _ = fi14
        _ = v15
    }
    var v16 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
