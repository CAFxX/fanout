package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 399
    var v3 uint8 = 250
    v4 := make(map[string]string)
    var v5 int8 = 42
    var v6 int32 = 1
    v6 = v6
    if false {
        for fi7 := 0; fi7 < 5; fi7++ {
            v6 += v6
            _ = fi7
        }
        v8 := []int{127, 32768}
        v9 := GenS1{F0: 13, F1: uint16(v2)}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := []int{fi10, fi10, (^(fi10))}
            var v12 int64 = 42
            v5 += 0
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint32 = 682
        if false {
            v14 := []int{-1, 1}
            _ = v14
        }
        _ = v13
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS2{F0: 0, F1: 32767}
        v3 += v3
        switch "" {
        case "foo bar":
            v5 += (^(v5))
        case "\n\t":
            v2 += v2
        }
        _ = fi15
        _ = v16
    }
    v17 := []int{9223372036854775807, -197}
    if (false && true) {
        v18 := GenS2{F0: -9223372036854775808, F1: 668}
        v1[128] = 44
        _ = v18
    }
    var v19 float64 = float64((3.14159 / ((3.14159 + 1.5) + 0.0)))
    v20 := GenS3{F0: (v5 + v5)}
    if (2147483647 > 0) {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        fmt.Println(v19)
        v2 = (v2 % ((v2) | 1))
    }
    v6 = (v6 * 42)
    v3 -= 147
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v2, v6)
        v3 -= (^(v3))
        _ = fi21
    }
    if false {
        _ = len(v17)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
