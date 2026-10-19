package main

import (
    "fmt"
)

func genFn0(p0 int32) int {
    fmt.Println(p0)
    var v4 uint16 = 255
    var v5 uint8 = 1
    _ = p0
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    _ = len(v2)
    if (345 < 2147483648) {
        var v6 uint8 = 128
        var v7 string = "hello"
        _ = v6
        _ = v7
    }
    v8 := []int{9223372036854775807, 2147483647}
    var v9 int16 = 980
    v1 = (0 >= 15)
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 float32 = float32(1.5)
        v3[4294967295] = 274
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 int32 = 49
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int64 = 561
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    var v16 uint = 42
    for fi17 := 0; fi17 < 13; fi17++ {
        v8 = append(v8, genFn0(-2147483648))
        fmt.Println(v16, fi17)
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
