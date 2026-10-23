package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint8 = 101
    v2 := make(map[uint32]bool)
    var v3 uint64 = 852
    v4 := make(map[int]bool)
    var v5 int = -9223372036854775808
    var v6 uint = 1
    v1 = (v1 + (^(v1)))
    var v7 int8 = -83
    switch (("hello" + "hello") + "\n\t") {
    case "hello":
        var v8 uint32 = 2147483647
        _ = v8
    default:
        var v9 uint32 = 42
        _ = v9
    }
    v3 -= (v3 / (((v3 ^ 0)) | 1))
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := v4[(fi10 ^ (-(v5)))]
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int8 = (v7 >> ((v6) & 63))
            var v14 int8 = 38
            _ = len(v4)
            var v15 int64 = -9223372036854775808
            var v16 int16 = 0
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
            _ = v16
        }
        v5 += -1
        _ = fi10
        _ = v11
    }
    fmt.Println(v5, v6)
    if false {
        var v17 uint32 = 2147483648
        var v18 int16 = -206
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: int8(v3)}
    var v20 int = (v5 << (((^(v6))) & 63))
    v5 = int(v1)
    var v21 uint = v6
    for fi22 := 0; fi22 < 10; fi22++ {
        v21 += (^(v21))
        _ = fi22
    }
    v6 -= (v21 ^ (v21 + 255))
    var v23 string = "foo bar"
    switch (v20 * 32768) {
    case 42:
        var v24 float32 = (1e10 - 0.1)
        _ = v24
    case -9223372036854775808:
        v25 := []int{(^(v20))}
        _ = v25
    case 32768:
        v3 = 2
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
