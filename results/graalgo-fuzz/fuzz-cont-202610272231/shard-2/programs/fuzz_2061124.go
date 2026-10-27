package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
func genFn0(p0 uint32) uint64 {
    for fi4 := 0; fi4 < 13; fi4++ {
        if true {
            for fi5 := 0; fi5 < 20; fi5++ {
                p0 = p0
                _ = fi5
            }
            v6 := []int{542, (^(fi4))}
            _ = v6
        }
        _ = fi4
    }
    v7 := []int{32768, 9223372036854775807, 1, 32768}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 string) int32 {
    p0 = p0
    v8 := GenS1{F0: p0, F1: 1000, F2: 2}
    var v9 float32 = -2.25
    switch "\n\t" {
    case "foo bar":
        if (0 != 127) {
            v9 = v9
            var v10 int = 32767
            _ = v10
        }
    case "\n\t":
        switch p0 {
        case "a":
            var v11 uint64 = 956
            _ = v11
        case "hello":
            v12 := GenS1{F0: p0, F1: 42, F2: 1}
            _ = v12
        }
    case "hello":
        var v13 uint16 = 1
        _ = v13
    default:
        v9 += ((v9 - ((v9 * 0.1) / v9)) * v9)
    }
    v9 += v9
    _ = p0
    _ = v8
    _ = v9
    return -507
}

func main() {
    var v1 uint8 = 151
    var v2 int64 = 1
    var v3 float64 = 0.1
    v14 := GenS1{F0: "", F1: int16(v1), F2: genFn0(65535)}
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int8 = -1
        v2 += (v2 >> ((uint(v1)) & 63))
        _ = fi15
        _ = v16
    }
    v1 = v1
    var v17 uint = 18446744073709551615
    if (0.0 > (3.14159 + (float32((v3 - 0.1)) - -2.25))) {
        v18 := GenS2{F0: v2, F1: 70, F2: v17}
        var v19 int32 = genFn1("hello")
        var v20 int32 = genFn1(("hello" + "hello"))
        _ = v18
        _ = v19
        _ = v20
    }
    if true {
        if false {
            v1 = v1
        } else {
            switch "foo bar" {
            case "\n\t":
                v1 -= 127
            case "hello":
                var v21 int32 = 1
                _ = v21
            }
            v22 := GenS1{F0: ("hello" + ""), F1: 1, F2: 2}
            _ = v22
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int64 = v2
            _ = fi23
            _ = v24
        }
    }
    v25 := []int{-2147483648}
    v17 = v17
    var v26 float64 = ((v3 + (v3 + float64(0.0))) + ((v3 / (float64(1.5) / (((((3.14159 - 0.0) - 0.1) - 1.5) / ((v3 + v3) + v3)) - (0.1 + 3.14159)))) * 3.14159))
    for fi27 := 0; fi27 < 9; fi27++ {
        switch "\n\t" {
        case "a":
            v2 -= (v2 << ((v17) & 63))
        }
        for fi28 := 0; fi28 < 20; fi28++ {
            v2 += (v2 % ((v2) | 1))
            _ = fi28
        }
        _ = fi27
    }
    if len(v25) > 0 {
        v25[0] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
