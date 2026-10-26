package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
func genFn0() uint {
    var v4 float64 = ((0.1 / ((1.5 * -2.25) / 1e10)) + float64(3.14159))
    v5 := []int{32768, 0, -2147483648, 644}
    fmt.Println(v4)
    v6 := GenS1{F0: 984, F1: 65535, F2: 18446744073709551615}
    _ = v4
    _ = v5
    _ = v6
    return 104
}
func genFn1(p0 int32) int16 {
    v7 := []int{}
    p0 -= -544
    v8 := GenS1{F0: 193, F1: 2147483647, F2: 4294967295}
    if false {
        var v9 uint = genFn0()
        _ = v9
    }
    _ = p0
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 int = 487
    v2 := []int{(v1 % ((v1) | 1)), (v1 ^ v1), v1}
    var v3 int32 = -2147483648
    var v10 int16 = 42
    if (false || (!(-42 == 100))) {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        var v11 uint = 765
        _ = v11
    } else {
        switch ("foo bar" + "foo bar") {
        case "hello":
            fmt.Println(v3, v3, v10)
        case "":
            v3 += 0
        case "\n\t":
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: 42, F1: 65535, F2: 141}
                fmt.Println(v10, v1)
                _ = fi12
                _ = v13
            }
        default:
            if false {
                var v14 string = ((("foo bar" + "\n\t") + "hello") + "\n\t")
                _ = v14
            } else {
                switch "" {
                case "a":
                    fmt.Println(v1, v10, v3)
                }
            }
        }
    }
    v3 = (v3 | v3)
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 int32 = (v3 >> ((287) & 63))
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint64 = 1
        _ = fi17
        _ = v18
    }
    switch ("" + "\n\t") {
    case "foo bar":
        v19 := GenS2{F0: v1, F1: 1, F2: 65535}
        _ = v19
    }
    if true {
        var v20 int64 = 1
        if true {
            var v21 uint16 = 0
            _ = v21
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v20
    } else {
        fmt.Println(v10, v3, v10)
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v10 = v10
        if (1.5 == (-2.25 + 0.1)) {
            v10 += (v10 - v10)
            v1 = (v1 * -9223372036854775808)
            if true {
                var v23 float64 = 1e10
                _ = v23
            }
        }
        _ = fi22
    }
    v2 = append(v2, 128)
    _ = len(v2)
    var v24 uint64 = 2
    var v25 int32 = (v3 * v3)
    v26 := GenS1{F0: 1, F1: 116, F2: uint(v24)}
    var v27 float32 = 0.0
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 int16 = (v10 << ((255) & 63))
        if false {
            if (false || (42 < 1)) {
                v2 = append(v2, 321)
                var v30 int32 = 2147483647
                _ = v30
            }
            var v31 bool = true
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
