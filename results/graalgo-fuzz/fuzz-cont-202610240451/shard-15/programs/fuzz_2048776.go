package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0(p0 bool, p1 uint64) int16 {
    v7 := []int{2147483647, 2147483647}
    p0 = (38 >= 32768)
    _ = p0
    _ = p1
    _ = v7
    return 225
}

func main() {
    v1 := []string{}
    v2 := []string{"\n\t", "", "hello", "hello"}
    v3 := []uint32{164, 4294967295, 2147483647, 4294967295}
    var v4 uint16 = 42
    var v5 int64 = 1
    v6 := []uint32{}
    switch ("a" + "") {
    case "hello":
        var v8 bool = false
        _ = v8
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "":
        switch ("foo bar" + ("\n\t" + "a")) {
        case "":
            fmt.Println(v5)
        case "a":
            v4 += (v4 % ((42) | 1))
        default:
            var v9 float64 = (0.0 / -2.25)
            _ = v9
        }
    default:
        v5 = v5
    }
    v10 := GenS1{F0: 4294967295, F1: -2147483648}
    var v11 uint = 971
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint = 255
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(v5, v11, v11)
        v4 = v4
        _ = fi14
    }
    v5 = 9223372036854775807
    var v15 int16 = 1000
    v16 := GenS2{F0: 203, F1: -1}
    v11 = v11
    if true {
        if true {
            v2 = append(v2, "")
        }
    }
    _ = len(v6)
    var v17 bool = true
    fmt.Println(v4, v17)
    for fi18 := 0; fi18 < 15; fi18++ {
        v15 -= (-(v15))
        _ = fi18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
