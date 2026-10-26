package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
func genFn0(p0 int32, p1 string, p2 uint) int {
    v7 := []int{32768}
    var v8 uint = (p2 + (p2 + p2))
    var v9 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -971
}

func main() {
    var v1 string = "a"
    var v2 uint8 = 128
    var v3 int8 = 1
    var v4 uint64 = 0
    var v5 string = "foo bar"
    var v6 float64 = (((-2.25 * 1e10) / 1e10) - float64(1.5))
    v10 := []int{42, -986, -590, 32768}
    if true {
        v11 := GenS1{F0: "", F1: 32767, F2: 421}
        var v12 float64 = v6
        fmt.Println(v4, v5)
        _ = v11
        _ = v12
    } else {
        v13 := GenS1{F0: v5, F1: -32768, F2: -2147483648}
        v14 := GenS1{F0: ("hello" + "\n\t"), F1: 1, F2: -2147483648}
        v5 = v1
        _ = v13
        _ = v14
    }
    var v15 string = (v5 + (v1 + v5))
    v5 = v1
    switch (("" + (v15 + "")) + v15) {
    case "a":
        v16 := GenS1{F0: v1, F1: 65, F2: -1}
        _ = v16
    case "":
        var v17 int16 = -32768
        _ = v17
    }
    v3 += (v3 & (v3 * v3))
    v3 = v3
    var v18 int32 = 2147483647
    for fi19 := 0; fi19 < 16; fi19++ {
        switch (((v1 + v5) + v1) + v5) {
        case "\n\t":
            var v20 uint64 = 1
            _ = v20
        case "hello":
            v3 = (v3 * v3)
        default:
            var v21 int64 = 9223372036854775807
            _ = v21
        }
        _ = fi19
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
