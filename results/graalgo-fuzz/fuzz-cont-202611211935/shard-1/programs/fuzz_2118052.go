package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0() int8 {
    var v6 uint16 = 42
    if (1 >= 640) {
        v7 := GenS1{F0: 780, F1: 128}
        fmt.Println(v6)
        var v8 float64 = (1e10 * 0.1)
        _ = v7
        _ = v8
    }
    _ = v6
    return -85
}

func main() {
    var v1 uint16 = 576
    var v2 float32 = -2.25
    var v3 int64 = -9223372036854775808
    var v4 int = -620
    var v5 float64 = 1.5
    var v9 int16 = int16(v4)
    var v10 bool = false
    v10 = (!v10)
    if v10 {
        var v11 uint = 255
        if (!v10) {
            v12 := []int{(v4 * v4), (v4 >> (((^(v11))) & 63))}
            _ = v12
        }
        var v13 uint16 = (^(v1))
        _ = v11
        _ = v13
    } else {
        var v14 int = 0
        fmt.Println(v10, v1)
        _ = v14
    }
    if v10 {
        v2 -= (v2 * (v2 + v2))
        var v15 float64 = 3.14159
        v1 -= (v1 | v1)
        _ = v15
    } else {
        v10 = v10
        v10 = ((v10 || ((v10 && true) || (v10 && v10))) && v10)
        v16 := []int{v4, v4}
        _ = v16
    }
    v10 = false
    for fi17 := 0; fi17 < 20; fi17++ {
        v4 = (fi17 ^ v4)
        switch ("foo bar" + ("foo bar" + "hello")) {
        case "hello":
            fmt.Println(v2, v2, v5)
        case "\n\t":
            var v18 float32 = v2
            _ = v18
        }
        _ = fi17
    }
    if v10 {
        v19 := []int{(^(v4))}
        fmt.Println(v9, v2)
        if len(v19) > 0 {
            v19[0] = int(v9)
        }
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        fmt.Println(v3)
        for fi21 := 0; fi21 < 3; fi21++ {
            if (false && v10) {
                if false {
                    v4 += (v4 | (v4 & fi21))
                }
            }
            v4 = (fi21 - (v4 << ((0) & 63)))
            switch 0 {
            case 32768:
                v1 = (^(v1))
            }
            fmt.Println(v5, v3, v2)
            v22 := []int{fi20, -2147483648, fi20, (fi21 & 127)}
            _ = len(v22)
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 uint8 = 1
            _ = fi23
            _ = v24
        }
        _ = fi20
    }
    v25 := GenS1{F0: 686, F1: (v4 & v4)}
    var v26 int8 = int8(v3)
    if v10 {
        v27 := GenS1{F0: 1, F1: 32767}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 uint8 = 0
            var v30 float32 = (((v2 * v2) / (1.5 / ((float32(v5) - 0.1) / ((v2 + v2) - v2)))) - 1e10)
            _ = fi28
            _ = v29
            _ = v30
        }
        v4 += v4
    }
    var v31 bool = (65535 > 2)
    v32 := []int{v4, (v4 + 32768), (v4 << ((0) & 63))}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
