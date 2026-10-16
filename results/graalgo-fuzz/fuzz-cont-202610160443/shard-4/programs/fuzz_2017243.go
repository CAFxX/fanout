package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
func genFn0(p0 int16, p1 bool, p2 int64) int {
    p2 -= 412
    if (18446744073709551615 <= 65535) {
        fmt.Println(p0)
        for fi5 := 0; fi5 < 16; fi5++ {
            v6 := []int{fi5, 128}
            p2 = 42
            var v7 uint32 = 728
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        v8 := []int{32768}
        v9 := GenS1{F0: -9223372036854775808, F1: (!p1)}
        v10 := GenS1{F0: 756, F1: (!(0 != 1))}
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 0
}

func main() {
    v1 := []string{}
    var v2 float32 = 0.1
    var v3 uint = 1
    var v4 float64 = ((0.1 - 0.1) / 1e10)
    var v12 string = "a"
    switch v12 {
    case "a":
        var v13 float64 = v4
        _ = v13
    case "":
        var v14 int16 = 0
        _ = v14
    default:
        for fi15 := 0; fi15 < 13; fi15++ {
            switch (fi15 % ((fi15) | 1)) {
            case 0:
                v16 := GenS1{F0: 1, F1: false}
                _ = v16
            }
            var v17 bool = false
            var v18 float32 = 0.1
            _ = fi15
            _ = v17
            _ = v18
        }
    }
    var v19 int8 = 42
    v12 = ""
    if false {
        var v20 int = 4294967295
        _ = v20
    }
    switch "a" {
    case "\n\t":
        var v21 int8 = (-(v19))
        _ = v21
    case "hello":
        if true {
            v4 -= v4
        } else {
            v2 += (v2 / 1e10)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v22 int32 = -2147483648
            _ = v22
        }
    case "a":
        v12 = "\n\t"
    default:
        fmt.Println(v4, v12, v12)
    }
    v23 := GenS1{F0: genFn0(1000, (v19 <= (v19 | (v19 / (((v19 * v19)) | 1)))), 42), F1: true}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v24 int = 9223372036854775807
    fmt.Println(v12, v4, v3)
    v4 += v4
    v24 = v24
    v19 -= 1
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v25 int8 = v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
