package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0(p0 int32, p1 uint32) uint64 {
    v6 := []int{253, 1, -8, 9223372036854775807}
    if false {
        var v7 int32 = (p0 << ((65535) & 63))
        _ = v7
    } else {
        var v8 float32 = ((float32((0.0 + (3.14159 * 0.1))) - ((((-2.25 / (1e10 / 1.0)) / 0.1) / 1.5) - 1e10)) / float32((float64(1.5) + 1.5)))
        v9 := GenS1{F0: 4294967295, F1: 0, F2: 158}
        var v10 float32 = (1.5 - -2.25)
        _ = v10
        _ = v8
        _ = v9
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v6
    return 17
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]string)
    var v3 float32 = (-2.25 * 3.14159)
    var v4 uint32 = 198
    var v5 bool = true
    if (!v5) {
        var v11 uint = 384
        _ = v11
    } else {
        v1 = uint(v4)
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := []int{}
        var v14 int64 = 163
        v5 = v5
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        fmt.Println(v3, v1, v5)
        var v16 uint = (^(v1))
        var v17 int32 = 37
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 18446744073709551615, F1: -1, F2: 1}
    var v19 string = ("a" + "foo bar")
    v20 := GenS1{F0: 184, F1: 0, F2: 255}
    switch (v19 + "a") {
    case "":
        var v21 float64 = (1.5 * (-2.25 / 1.0))
        _ = v21
    case "foo bar":
        v1 = (v1 >> ((1) & 63))
    case "a":
        for fi22 := 0; fi22 < 10; fi22++ {
            v5 = (4294967295 < (v4 ^ 123))
            if false {
                delete(v2, (^(v4)))
                var v23 int16 = 504
                var v24 float64 = -2.25
                _ = v23
                _ = v24
            }
            var v25 string = v19
            _ = fi22
            _ = v25
        }
    default:
        fmt.Println(v4)
    }
    var v26 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
