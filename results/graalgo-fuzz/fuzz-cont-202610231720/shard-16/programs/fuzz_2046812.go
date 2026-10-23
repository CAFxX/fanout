package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
func genFn0(p0 uint) int16 {
    v5 := GenS1{F0: "a", F1: p0}
    p0 -= (p0 & p0)
    var v6 int16 = 723
    _ = p0
    _ = v5
    _ = v6
    return (v6 / (((v6 | (v6 / ((v6) | 1)))) | 1))
}
func genFn1(p0 uint16, p1 uint64) uint {
    v7 := []int{31}
    p0 += p0
    p0 = p0
    fmt.Println(p1, p1)
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 322
}

func main() {
    var v1 string = ""
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    for fi9 := 0; fi9 < 1; fi9++ {
        v1 = "foo bar"
        _ = fi9
    }
    switch -2147483648 {
    case 768:
        var v10 int8 = -1
        _ = v10
    }
    v1 = v1
    v1 = v1
    v1 = v1
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    switch 32767 {
    case 831:
        switch v1 {
        case "a":
            v1 = (v1 + v1)
        case "hello":
            v1 = "foo bar"
        case "foo bar":
            var v12 string = v1
            _ = v12
        default:
            fmt.Println(v1)
        }
    }
    var v13 uint64 = 4294967295
    if (18446744073709551615 > 42) {
        v14 := []int{-9223372036854775808}
        _ = v14
    }
    var v15 string = (v1 + v1)
    var v16 int = 1
    v15 = ("" + "foo bar")
    delete(v3, v16)
    v15 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
