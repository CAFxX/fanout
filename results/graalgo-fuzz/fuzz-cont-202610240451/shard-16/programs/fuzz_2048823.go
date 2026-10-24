package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint16 }
func genFn0() int32 {
    switch 127 {
    case 2147483647:
        var v7 string = "foo bar"
        _ = v7
    case -9223372036854775808:
        var v8 uint = 241
        _ = v8
    case 1:
        v9 := []int{4294967295, -557, 0}
        _ = v9
    }
    v10 := GenS1{F0: -128, F1: 0, F2: 1}
    if false {
        var v11 uint16 = 372
        v12 := GenS1{F0: 1, F1: 128, F2: (v11 - (v11 ^ 2))}
        v11 = (v11 >> ((0) & 63))
        _ = v11
        _ = v12
    } else {
        v13 := GenS1{F0: -1, F1: 80, F2: 0}
        _ = v13
    }
    v14 := []int{128, 42}
    if ((0 == 9223372036854775807) || (true || ((2 >= 1) && false))) {
        if true {
            if (-1 == 1000) {
                _ = len(v14)
                var v15 uint32 = 2
                _ = v15
            } else {
                if len(v14) > 0 {
                    v14[0] = -2147483648
                }
                var v16 uint16 = 954
                _ = v16
            }
            v17 := GenS1{F0: -128, F1: 127, F2: 33}
            _ = v17
        }
        v18 := GenS1{F0: 100, F1: 127, F2: 65535}
        var v19 float64 = 1.5
        _ = v18
        _ = v19
    }
    _ = v10
    _ = v14
    return 0
}
func genFn1() uint32 {
    var v20 uint32 = 2
    v21 := []int{-378, -1}
    _ = v20
    _ = v21
    return v20
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]bool)
    var v3 int8 = 127
    var v4 uint = 65535
    var v5 int32 = -1
    var v6 uint8 = 2
    v3 = v3
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint64 = v1
        _ = fi22
        _ = v23
    }
    var v24 bool = false
    var v25 int64 = -9223372036854775808
    switch "hello" {
    case "hello":
        v1 = (v1 - 775)
    case "\n\t":
        v5 -= (^(v5))
    default:
        fmt.Println(v25, v25, v24)
    }
    fmt.Println(v24, v6)
    var v26 int16 = -913
    if false {
        v5 += int32(v25)
    }
    fmt.Println(v24, v25)
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 string = "foo bar"
        _ = fi27
        _ = v28
    }
    var v29 int = -340
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
