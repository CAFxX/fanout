package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
func genFn0(p0 int32) int16 {
    switch "\n\t" {
    case "\n\t":
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int64 = 9223372036854775807
            var v8 int32 = p0
            var v9 int = 9223372036854775807
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        p0 -= p0
        var v11 uint = 65535
        _ = fi10
        _ = v11
    }
    var v12 float64 = 1.5
    switch 1 {
    case 4294967295:
        switch 127 {
        case -1:
            fmt.Println(p0)
        }
    case 32767:
        var v13 int32 = (-(p0))
        _ = v13
    }
    _ = p0
    _ = v12
    return -32768
}
func genFn1() bool {
    var v14 float32 = -2.25
    v14 = (((v14 + (v14 + v14)) - v14) * 1e10)
    var v15 int32 = 217
    v14 = ((v14 - v14) / v14)
    _ = v14
    _ = v15
    return true
}

func main() {
    var v1 uint64 = 845
    var v2 int8 = 20
    var v3 uint = 93
    var v4 int8 = 100
    var v5 uint64 = v1
    switch "" {
    case "foo bar":
        var v16 int64 = 9223372036854775807
        _ = v16
    case "a":
        var v17 float64 = 3.14159
        _ = v17
    case "hello":
        v4 += -67
    }
    switch 1 {
    case 194:
        v18 := []int{-200, 32767, 263}
        _ = v18
    case 634:
        var v19 bool = false
        _ = v19
    case 42:
        switch "foo bar" {
        case "":
            var v20 string = ""
            _ = v20
        case "\n\t":
            v4 = (v4 & (v4 / (((v4 >> ((255) & 63))) | 1)))
        case "foo bar":
            var v21 string = ("foo bar" + "hello")
            _ = v21
        default:
            v22 := []int{127}
            _ = v22
        }
    }
    if true {
        var v23 uint16 = 0
        if (65535 >= (v23 ^ 42)) {
            fmt.Println(v4)
        } else {
            var v24 string = "a"
            v25 := GenS1{F0: 295, F1: v24, F2: (v23 / ((v23) | 1))}
            _ = v24
            _ = v25
        }
        v26 := GenS1{F0: -1, F1: (("\n\t" + "\n\t") + "hello"), F2: (^(v23))}
        _ = v23
        _ = v26
    } else {
        v27 := GenS1{F0: -1, F1: "\n\t", F2: 65535}
        _ = v27
    }
    var v28 string = "foo bar"
    if false {
        v3 += 1
    }
    if true {
        v3 = (v3 / ((v3) | 1))
        v2 = (v4 - v4)
    }
    v28 = ("" + v28)
    v29 := GenS2{F0: 32767, F1: true, F2: 755}
    var v30 int = -1
    var v31 uint32 = 2147483648
    v32 := []int{v30, -780}
    for fi33 := 0; fi33 < 2; fi33++ {
        v2 -= (v2 << (((v3 >> ((133) & 63))) & 63))
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
