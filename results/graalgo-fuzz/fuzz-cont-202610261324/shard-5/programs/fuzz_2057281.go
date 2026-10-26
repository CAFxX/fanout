package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0(p0 int8, p1 int8) int16 {
    fmt.Println(p1, p1)
    var v4 int16 = -32768
    _ = p0
    _ = p1
    _ = v4
    return v4
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 0
    var v3 int64 = 9223372036854775807
    v5 := GenS1{F0: 825, F1: 255}
    if (!(42 < 127)) {
        if false {
            v6 := []int{-635}
            v7 := GenS1{F0: v2, F1: -1}
            _ = v6
            _ = v7
        }
        var v8 uint8 = 127
        _ = v8
    }
    switch -902 {
    case -1:
        var v9 bool = (6 == 42)
        _ = v9
    case 1:
        if ((^(v2)) <= 954) {
            v10 := []int{-1, 32768, -56, -1}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := []int{(fi11 ^ fi11), (fi11 << ((42) & 63)), fi11}
                v2 -= v2
                _ = fi11
                _ = v12
            }
        }
    }
    v3 = 42
    fmt.Println(v3, v2)
    v13 := GenS1{F0: v2, F1: -764}
    var v14 uint8 = 1
    var v15 float64 = -2.25
    var v16 int64 = v3
    var v17 float32 = ((3.14159 * (0.1 + float32((0.0 / v15)))) * -2.25)
    v3 -= v3
    var v18 int8 = 127
    switch "hello" {
    case "\n\t":
        fmt.Println(v2)
    case "foo bar":
        switch 9223372036854775807 {
        case 1:
            switch ("a" + "hello") {
            case "\n\t":
                v3 = (v3 / ((-9223372036854775808) | 1))
            case "a":
                v14 -= 0
            case "foo bar":
                v16 = (v16 >> ((42) & 63))
            }
        }
    case "a":
        v3 = 9223372036854775807
    default:
        v16 -= v16
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v3 = v16
        _ = fi19
    }
    v20 := GenS1{F0: (^(v2)), F1: -397}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v5
}
