package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int64 }

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]bool)
    var v3 int32 = 0
    fmt.Println(v1, v3)
    switch 32767 {
    case -2147483648:
        fmt.Println(v1, v3)
    case 333:
        v1 += (v1 / (((v1 ^ (v1 ^ (v1 * 127)))) | 1))
    case -239:
        if false {
            switch "foo bar" {
            case "foo bar":
                if (1.5 < 0.0) {
                    fmt.Println(v3)
                }
            }
            v4 := []int{42}
            _ = v4
        } else {
            v5 := GenS1{F0: 2}
            _ = v5
        }
    default:
        v6 := GenS2{F0: 0, F1: 647}
        _ = v6
    }
    var v7 int32 = 0
    if ("" > "foo bar") {
        v8 := []int{}
        switch 128 {
        case 2147483647:
            v3 = (v7 + v3)
        }
        _ = v8
    } else {
        v9 := GenS2{F0: 0, F1: 1}
        fmt.Println(v3, v7, v3)
        _ = v9
    }
    var v10 uint64 = 343
    var v11 uint32 = 42
    v12 := GenS1{F0: 247}
    for fi13 := 0; fi13 < 10; fi13++ {
        v1 = v1
        var v14 uint32 = uint32(v1)
        switch (fi13 | 32768) {
        case 32767:
            v15 := GenS1{F0: 42}
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v16 float64 = -2.25
    v17 := []int{1, 32768, 4294967295}
    v10 -= (^(v10))
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(v3, v7, v1)
        v16 -= (v16 * (v16 - float64(((-2.25 + 3.14159) + 1e10))))
        _ = fi18
    }
    if ((!false) && (-501 == 42)) {
        v19 := GenS1{F0: 255}
        switch ("hello" + "\n\t") {
        case "foo bar":
            v3 = 0
        case "\n\t":
            v20 := GenS2{F0: 42, F1: 9223372036854775807}
            _ = v20
        case "hello":
            v1 -= (v1 | (v1 | v1))
        }
        _ = v19
    }
    var v21 uint32 = 4294967295
    if ((52 != 42) || false) {
        var v22 uint8 = 102
        if true {
            var v23 bool = false
            _ = v23
        }
        var v24 int64 = 1
        _ = v22
        _ = v24
    } else {
        v3 = (-(v3))
        v21 += v11
        if true {
            var v25 int8 = -81
            _ = v25
        }
    }
    v3 -= v3
    v26 := GenS1{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v7
}
