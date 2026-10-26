package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint64 }
func genFn0() int64 {
    var v6 string = ("a" + "")
    fmt.Println(v6)
    v7 := GenS1{F0: 1}
    v8 := []int{2147483647}
    fmt.Println(v6)
    _ = v6
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 uint8, p1 string, p2 int32) uint8 {
    switch 42 {
    case 42:
        fmt.Println(p1, p1)
    case 2147483647:
        v9 := []int{-395}
        _ = v9
    }
    var v10 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (^(p0))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 2
    v3 := make(map[uint32]bool)
    var v4 uint64 = 738
    var v5 float32 = 0.0
    switch "" {
    case "\n\t":
        v11 := GenS1{F0: 0}
        _ = v11
    default:
        var v12 uint16 = 42
        _ = v12
    }
    v13 := GenS2{F0: 42, F1: uint64(v2)}
    var v14 uint8 = genFn1((^(v2)), "\n\t", -2147483648)
    var v15 float32 = v5
    v5 += v15
    if true {
        switch ("a" + ("hello" + "")) {
        case "":
            v16 := GenS2{F0: 4294967295, F1: v4}
            _ = v16
        case "a":
            delete(v3, 2147483648)
        default:
            v14 = 73
        }
    }
    if (4294967295 >= 1) {
        switch -876 {
        case 32767:
            v17 := []int{9223372036854775807, 2147483647, 4294967295, 1}
            _ = v17
        case -2147483648:
            v15 = v5
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 int16 = 0
            _ = fi18
            _ = v19
        }
    }
    v5 += v5
    var v20 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
