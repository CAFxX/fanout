package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 bool) uint64 {
    v6 := []int{1, -431, -1}
    if false {
        var v7 bool = (p0 && p0)
        var v8 int16 = 0
        v9 := GenS1{F0: 2147483647, F1: 42, F2: 2}
        _ = v7
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 510, F1: 0, F2: 0}
        _ = len(v6)
        p0 = (0.1 != 3.14159)
        _ = v10
    }
    p0 = false
    var v11 float64 = (3.14159 + ((0.0 / ((0.1 / 1.0) + (3.14159 * float64(0.1)))) + (0.0 * 1.5)))
    _ = p0
    _ = v11
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint16 = 572
    var v2 int = -75
    v3 := make(map[string]string)
    var v4 int16 = 767
    var v5 uint = 4294967295
    var v12 int32 = 2147483647
    if ((float64(3.14159) <= 0.0) && true) {
        var v13 uint16 = v1
        _ = v13
    } else {
        for fi14 := 0; fi14 < 5; fi14++ {
            switch "a" {
            case "foo bar":
                for fi15 := 0; fi15 < 7; fi15++ {
                    var v16 uint8 = 187
                    fmt.Println(v5, fi14)
                    _ = fi15
                    _ = v16
                }
            case "\n\t":
                var v17 uint16 = (^(v1))
                _ = v17
            }
            fmt.Println(fi14, v5)
            _ = fi14
        }
    }
    v18 := GenS2{F0: 2}
    v19 := GenS2{F0: 17}
    fmt.Println(v5, v2, v1)
    for fi20 := 0; fi20 < 12; fi20++ {
        v12 = -400
        _ = fi20
    }
    var v21 float64 = 3.14159
    switch (((("" + ("hello" + "")) + "\n\t") + "foo bar") + "hello") {
    case "\n\t":
        var v22 int64 = -1
        _ = v22
    case "a":
        v23 := []int{}
        _ = v23
    }
    if true {
        var v24 int8 = -128
        switch "\n\t" {
        case "a":
            _ = len(v3)
        case "\n\t":
            v5 = 1
        case "":
            v2 -= 32767
        default:
            var v25 string = "hello"
            _ = v25
        }
        v26 := v3["\n\t"]
        _ = v24
        _ = v26
    }
    if true {
        for fi27 := 0; fi27 < 15; fi27++ {
            if true {
                fmt.Println(v4)
            }
            _ = fi27
        }
    } else {
        var v28 int16 = (v4 >> ((v5) & 63))
        v29 := []int{(v2 * (v2 % ((42) | 1))), 42, v2, 126}
        switch "hello" {
        case "\n\t":
            var v30 uint64 = 18446744073709551615
            _ = v30
        case "hello":
            if (false && false) {
                var v31 int32 = v12
                _ = v31
            }
        default:
            fmt.Println(v1, v28, v12)
        }
        _ = v28
        _ = v29
    }
    var v32 float64 = ((-2.25 - v21) / ((float64(3.14159) - v21) / ((0.1 * v21) - (v21 + v21))))
    v4 += (v4 ^ v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
