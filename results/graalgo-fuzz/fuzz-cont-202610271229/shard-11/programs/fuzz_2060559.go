package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint32, p1 int64, p2 string) uint {
    var v5 int = 2147483647
    if false {
        fmt.Println(p2, p0)
        fmt.Println(p1, p1, v5)
        var v6 int = v5
        _ = v6
    } else {
        p0 = p0
    }
    v7 := GenS1{F0: 127, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 951
}
func genFn1(p0 int16) int {
    fmt.Println(p0)
    p0 = (p0 + -52)
    p0 += 32767
    switch ("hello" + "") {
    case "a":
        var v8 int32 = -126
        _ = v8
    case "hello":
        v9 := []int{-1, 9223372036854775807, -2147483648}
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int8 = -2
    var v3 float32 = (float32((0.1 * ((3.14159 - (((0.0 / (float64(1e10) * 1e10)) * 0.0) * ((((1e10 - (0.1 - (((float64((0.0 / -2.25)) * (0.1 - 0.0)) - (3.14159 + 3.14159)) / 1.5))) / 1.5) + (-2.25 / (-2.25 / ((float64(0.1) + 1.5) - (0.0 + (-2.25 * (-2.25 + 3.14159))))))) - 3.14159))) - ((1e10 / -2.25) - (3.14159 * float64(-2.25)))))) / (((3.14159 - float32(float64((float32(0.1) - 1e10)))) * 1.5) - (-2.25 - (1e10 / ((0.0 - 1.5) + 0.1)))))
    var v4 uint16 = 255
    fmt.Println(v3)
    var v10 float32 = -2.25
    var v11 int32 = 470
    v12 := GenS1{F0: 0, F1: (v2 / (((-(v2))) | 1))}
    if true {
        v1 -= (v1 ^ v1)
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        if false {
            fmt.Println(v1, v2, v1)
            switch fi13 {
            case -146:
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v4, fi13)
                }
            case 9223372036854775807:
                v14 := []int{}
                _ = v14
            case 32767:
                v15 := GenS2{F0: v2, F1: 2}
                _ = v15
            default:
                var v16 int = (fi13 | (fi13 - genFn1(-1)))
                _ = v16
            }
        }
        _ = fi13
    }
    v17 := []int{0, -47, 2147483647, -59}
    if false {
        if true {
            switch 909 {
            case -11:
                for fi18 := 0; fi18 < 1; fi18++ {
                    v1 += v1
                    fmt.Println(fi18)
                    _ = fi18
                }
            case 9223372036854775807:
                var v19 int = genFn1(553)
                _ = v19
            case 32767:
                var v20 int8 = v2
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v11, v1)
        _ = fi21
    }
    v22 := GenS2{F0: (^(v2)), F1: 0}
    if len(v17) > 0 {
        v17[0] = 84
    }
    v4 = v4
    fmt.Println(v10, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
