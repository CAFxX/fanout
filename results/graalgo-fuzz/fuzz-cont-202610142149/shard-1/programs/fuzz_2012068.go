package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16 }
func genFn0() int64 {
    v7 := []int{32767, 502}
    switch "foo bar" {
    case "hello":
        var v8 uint32 = 68
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 3; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                if (2147483647 == 2147483648) {
                    fmt.Println(fi9, fi10)
                    var v11 uint = 571
                    _ = v11
                }
                var v12 bool = (fi10 > fi9)
                _ = fi10
                _ = v12
            }
            var v13 bool = false
            _ = fi9
            _ = v13
        }
    default:
        var v14 bool = false
        _ = v14
    }
    var v15 int64 = 42
    _ = v15
    _ = v7
    return -67
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 bool = false
    var v4 int16 = -863
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]uint32)
    if (!v3) {
        var v16 uint64 = 2
        var v17 uint64 = v16
        _ = v16
        _ = v17
    }
    delete(v6, -2147483648)
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 float32 = (0.0 / 0.1)
        var v20 int8 = 42
        _ = fi18
        _ = v19
        _ = v20
    }
    v4 = (v4 >> ((65535) & 63))
    v21 := GenS1{F0: -1, F1: 0}
    _ = len(v1)
    var v22 uint8 = 124
    var v23 int64 = -9223372036854775808
    switch 4294967295 {
    case 369:
        var v24 uint = 0
        _ = v24
    default:
        if ((764 < 179) || (109 <= 787)) {
            var v25 uint8 = v22
            var v26 int = 2147483647
            _ = v25
            _ = v26
        } else {
            v27 := GenS1{F0: 2147483647, F1: 1}
            var v28 float64 = (0.1 + 0.1)
            _ = v27
            _ = v28
        }
    }
    v23 += (-(v23))
    fmt.Println(v3, v4, v4)
    v29 := GenS2{F0: (-(v4))}
    v3 = false
    var v30 int32 = -2147483648
    if ((v22 & (v22 & v22)) < (v22 * (^(v22)))) {
        v31 := GenS1{F0: v30, F1: 42}
        _ = v31
    } else {
        for fi32 := 0; fi32 < 9; fi32++ {
            fmt.Println(v22, v3)
            fmt.Println(v23, v23, v23)
            for fi33 := 0; fi33 < 2; fi33++ {
                fmt.Println(v4)
                _ = fi33
            }
            _ = fi32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
