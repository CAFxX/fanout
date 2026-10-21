package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int16; F1 int; F2 uint32 }
func genFn0() string {
    var v4 bool = ((1e10 - (3.14159 + 0.0)) <= 1.5)
    var v5 int16 = 218
    var v6 int = -675
    _ = v4
    _ = v5
    _ = v6
    return "hello"
}
func genFn1(p0 uint8, p1 string) bool {
    if true {
        if (!true) {
            v7 := []int{-9223372036854775808}
            _ = v7
        } else {
            var v8 float64 = -2.25
            v9 := []int{9223372036854775807, 127}
            _ = v8
            _ = v9
        }
        v10 := []int{9223372036854775807}
        _ = v10
    }
    fmt.Println(p0)
    v11 := GenS1{F0: (("a" + (p1 + "a")) + "foo bar"), F1: p0, F2: int16(p0)}
    _ = p0
    _ = p1
    _ = v11
    return (703 > 9223372036854775807)
}

func main() {
    var v1 string = ("" + "hello")
    v2 := make(map[string]int)
    var v3 int32 = 42
    v12 := GenS1{F0: v1, F1: 127, F2: -32768}
    v1 = v1
    v13 := GenS2{F0: ("" + ("hello" + (v1 + ""))), F1: 546}
    if false {
        var v14 uint = 0
        _ = v14
    }
    var v15 int16 = 942
    switch 665 {
    case -460:
        if (!(9223372036854775807 > 691)) {
            if (0 < 1) {
                fmt.Println(v1)
                v15 -= (v15 ^ (v15 | (^(v15))))
            }
        }
    case 32768:
        fmt.Println(v3, v3, v3)
    case 835:
        v16 := GenS1{F0: "foo bar", F1: 42, F2: (v15 + v15)}
        _ = v16
    default:
        fmt.Println(v15, v1, v3)
    }
    fmt.Println(v15, v3)
    if true {
        for fi17 := 0; fi17 < 14; fi17++ {
            v15 -= (v15 * v15)
            v1 = ((("" + (v1 + "")) + (v1 + "a")) + ("" + v1))
            var v18 uint32 = 1
            _ = fi17
            _ = v18
        }
        v19 := GenS3{F0: (v15 ^ (v15 | v15)), F1: 2147483647, F2: 2147483648}
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
}
