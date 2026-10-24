package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 bool, p1 int64, p2 uint) uint {
    if (2 >= 128) {
        p0 = p0
        var v5 float64 = 1.5
        p2 = (p2 % (((p2 >> (((p2 + (p2 % (((^(p2))) | 1)))) & 63))) | 1))
        _ = v5
    }
    var v6 float64 = 1e10
    p0 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1() int8 {
    var v7 float64 = 0.0
    fmt.Println(v7)
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 42
    var v2 uint16 = 2
    var v3 uint8 = 42
    v4 := make(map[int]uint32)
    var v8 int64 = 1
    switch "foo bar" {
    case "":
        v8 -= (v8 / ((124) | 1))
    case "foo bar":
        v1 -= 2
    }
    v9 := []int{}
    var v10 uint16 = (v2 * (^(v1)))
    if false {
        v8 -= v8
    } else {
        if true {
            var v11 uint32 = 2147483648
            _ = v11
        }
        v2 -= 863
        if true {
            delete(v4, 405)
            for fi12 := 0; fi12 < 2; fi12++ {
                v8 -= 1
                v3 = (v3 | v3)
                _ = fi12
            }
        }
    }
    fmt.Println(v2, v3)
    v13 := GenS1{F0: (18446744073709551615 == 255)}
    v1 -= 255
    if (3.14159 < 0.0) {
        v14 := GenS1{F0: false}
        v9 = append(v9, 1)
        _ = v14
    }
    var v15 float32 = (1e10 / 0.1)
    if (!(0 >= 732)) {
        var v16 uint8 = v3
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        var v17 uint32 = 0
        _ = v16
        _ = v17
    } else {
        fmt.Println(v3, v3)
        fmt.Println(v1, v8)
        fmt.Println(v8, v1, v1)
    }
    var v18 int16 = 532
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
