package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int32 }
func genFn0() int64 {
    if (128 <= 127) {
        switch ("a" + ("foo bar" + "hello")) {
        case "\n\t":
            for fi6 := 0; fi6 < 14; fi6++ {
                for fi7 := 0; fi7 < 19; fi7++ {
                    v8 := []int{-658, 42, (fi6 + (^(fi7))), 32767}
                    _ = fi7
                    _ = v8
                }
                var v9 string = ""
                v10 := GenS1{F0: 615, F1: 128}
                _ = fi6
                _ = v10
                _ = v9
            }
        }
        var v11 uint32 = 65535
        fmt.Println(v11)
        _ = v11
    }
    var v12 int16 = -1
    v12 += (v12 % (((v12 << ((42) & 63))) | 1))
    var v13 float64 = (0.1 * -2.25)
    v12 = 74
    _ = v12
    _ = v13
    return -1
}
func genFn1(p0 int32, p1 int, p2 uint8) int64 {
    var v14 float64 = ((-2.25 * 3.14159) + (0.1 / (1.5 - ((float64((1.5 + 1.5)) / 1e10) - -2.25))))
    var v15 float32 = 1e10
    switch (("a" + "hello") + "hello") {
    case "a":
        var v16 string = "\n\t"
        _ = v16
    case "hello":
        v17 := []int{p1, -2147483648}
        _ = v17
    default:
        v18 := []int{p1, p1}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 796
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 int8 = 73
    var v3 bool = false
    var v4 uint8 = 127
    v5 := make(map[int]int)
    v19 := GenS1{F0: 1, F1: v4}
    v20 := GenS1{F0: -1, F1: (^(v4))}
    var v21 int8 = v2
    v2 += -128
    var v22 int16 = 1
    if true {
        if v3 {
            var v23 uint = 42
            var v24 int16 = (v22 ^ (v22 << ((v23) & 63)))
            _ = v23
            _ = v24
        }
    } else {
        fmt.Println(v3, v22)
        if false {
            if v3 {
                v5[128] = -1
            }
        }
        var v25 int8 = (v21 >> ((279) & 63))
        _ = v25
    }
    var v26 bool = v3
    v27 := GenS2{F0: -787}
    if true {
        fmt.Println(v26, v26)
        v3 = (18446744073709551615 >= 0)
    }
    v21 = (v2 ^ v2)
    var v28 uint64 = 18446744073709551615
    switch "a" {
    case "\n\t":
        v1 -= float32((0.0 - 0.0))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
