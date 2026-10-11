package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 uint64 }
func genFn0() uint32 {
    var v5 float32 = 0.0
    var v6 uint = 4294967295
    var v7 float32 = 1.5
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{127, -38, -2147483648}
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    var v8 int16 = -504
    var v9 int64 = -9223372036854775808
    if (v8 > (^(v8))) {
        _ = len(v1)
        var v10 uint64 = 2
        v9 += 9223372036854775807
        _ = v10
    }
    var v11 int32 = -2147483648
    fmt.Println(v8, v11)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        _ = len(v3)
        if false {
            var v13 float64 = (3.14159 + 1.5)
            v2 = append(v2, fi12)
            _ = v13
        } else {
            switch "a" {
            case "foo bar":
                var v14 int16 = -269
                _ = v14
            case "hello":
                v15 := GenS1{F0: 721}
                _ = v15
            }
            v8 += 1
            if true {
                if (!true) {
                    var v16 uint16 = 519
                    _ = v16
                }
            }
        }
        _ = fi12
    }
    var v17 bool = ((-(v8)) <= (v8 + (v8 * -32768)))
    fmt.Println(v9, v11)
    v18 := GenS2{F0: 42, F1: v9, F2: 1}
    switch 42 {
    case -9223372036854775808:
        v17 = (-2.25 > (1e10 - (3.14159 / 1.5)))
    default:
        delete(v4, 4294967295)
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        switch "hello" {
        case "\n\t":
            v11 -= -296
        default:
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 int8 = 39
                var v22 uint64 = 18446744073709551615
                var v23 uint32 = 148
                delete(v1, ("" + ""))
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        }
        v17 = true
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
