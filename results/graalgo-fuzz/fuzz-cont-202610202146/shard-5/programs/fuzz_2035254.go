package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0() uint8 {
    v6 := GenS1{F0: (("\n\t" + "hello") + ("hello" + "")), F1: 42, F2: -2147483648}
    v7 := GenS1{F0: ("\n\t" + "foo bar"), F1: 22, F2: 42}
    v8 := []int{127, 398}
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 string = ""
    var v2 float32 = 0.0
    v3 := make(map[uint32]bool)
    var v4 uint8 = 204
    var v5 bool = true
    if true {
        v3[2147483648] = v5
        fmt.Println(v1)
    }
    var v9 uint = 4294967295
    if v5 {
        v10 := []int{-2147483648, 985}
        for fi11 := 0; fi11 < 5; fi11++ {
            if v5 {
                v3[4294967295] = false
            } else {
                v10 = append(v10, fi11)
                v12 := GenS2{F0: 2147483647, F1: 1}
                _ = v12
            }
            _ = fi11
        }
        var v13 int64 = 651
        _ = v10
        _ = v13
    } else {
        var v14 int = 128
        _ = v14
    }
    v15 := GenS1{F0: "foo bar", F1: genFn0(), F2: 2147483647}
    v2 += (0.1 + (float32((-2.25 * (0.1 + 3.14159))) - v2))
    if v5 {
        v4 -= (v4 | v4)
    }
    var v16 uint = 0
    v1 = (((("" + "hello") + (((v1 + v1) + "a") + v1)) + (v1 + v1)) + v1)
    v16 -= 42
    if (101 != 42) {
        fmt.Println(v1, v1)
        v1 = ("a" + "hello")
        v17 := v3[42]
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 bool = (!(4 > 126))
        var v20 int16 = 42
        var v21 uint16 = 0
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    switch v1 {
    case "a":
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := GenS2{F0: 33, F1: -49}
            v16 -= 170
            v24 := GenS1{F0: "foo bar", F1: (v4 << ((v16) & 63)), F2: (-(fi22))}
            _ = fi22
            _ = v23
            _ = v24
        }
    case "foo bar":
        var v25 uint64 = 93
        _ = v25
    default:
        v4 += (v4 * 1)
    }
    var v26 float32 = 0.1
    v27 := GenS2{F0: 1, F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
