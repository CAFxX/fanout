package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
func genFn0(p0 string, p1 bool, p2 int8) bool {
    switch p0 {
    case "":
        p1 = ((!(1 == 969)) || false)
    }
    var v6 uint64 = 1
    fmt.Println(v6, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ((42 >= 255) && ((((1.5 / 1.0) - (0.1 + ((((0.0 + 0.0) / (1e10 - 0.1)) - -2.25) + 3.14159))) - 0.0) < (0.1 * -2.25)))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{154}
    v3 := make(map[int]uint32)
    var v4 int = 0
    v5 := make(map[int]string)
    switch (v4 - -170) {
    case -2147483648:
        if (!true) {
            v7 := GenS1{F0: "foo bar", F1: 255, F2: 0}
            _ = v7
        } else {
            var v8 int16 = 32767
            if ((1e10 > 1.5) && (v4 >= 778)) {
                v8 = v8
            }
            v9 := v3[v4]
            _ = v8
            _ = v9
        }
    case 32767:
        v10 := v5[(v4 << ((255) & 63))]
        _ = v10
    }
    var v11 float32 = 1e10
    if len(v2) > 0 {
        v2[0] = (v4 & v4)
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        switch ("" + "") {
        case "":
            fmt.Println(v11, v11, fi12)
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = -1
            }
        case "a":
            v11 = v11
        default:
            fmt.Println(fi12)
        }
        _ = fi12
    }
    if len(v2) > 0 {
        v2[0] = (^(v4))
    }
    var v13 uint8 = 132
    var v14 int8 = -128
    v15 := GenS1{F0: "hello", F1: 2, F2: 26}
    var v16 uint = 255
    switch -1 {
    case 127:
        var v17 float64 = 0.1
        _ = v17
    case 12:
        var v18 uint64 = 42
        _ = v18
    }
    var v19 bool = (18446744073709551615 == 18446744073709551615)
    var v20 float32 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
