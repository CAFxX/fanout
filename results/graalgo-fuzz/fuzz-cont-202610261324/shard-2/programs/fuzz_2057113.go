package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0(p0 int16) int16 {
    var v7 uint32 = 2147483648
    v7 = 4294967295
    v8 := GenS1{F0: 881, F1: p0}
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := []int{}
        switch 4294967295 {
        case -9223372036854775808:
            var v11 int = fi9
            _ = v11
        case 1:
            v12 := GenS1{F0: 811, F1: p0}
            _ = v12
        case 32768:
            _ = len(v10)
        default:
            v7 += v7
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return (p0 | p0)
}

func main() {
    var v1 uint = 4294967295
    v2 := []string{("foo bar" + ("hello" + "a")), "", "a"}
    var v3 float64 = (1e10 * 1e10)
    var v4 float64 = float64(1e10)
    var v5 bool = true
    v6 := []int{32768, -2147483648, 4294967295}
    var v13 int16 = -1
    var v14 int32 = -1
    fmt.Println(v4, v3)
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println(v4, v5)
    _ = len(v6)
    switch ("a" + ((((("" + "a") + (("foo bar" + "a") + (("foo bar" + "") + "hello"))) + "") + "a") + "")) {
    case "":
        v6 = append(v6, 0)
    case "\n\t":
        v14 = -2147483648
    }
    fmt.Println(v4)
    v14 += (v14 / ((v14) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
