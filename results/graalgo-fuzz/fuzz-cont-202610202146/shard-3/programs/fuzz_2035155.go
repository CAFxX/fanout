package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int8; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 int }

func main() {
    v1 := []int{-1, 32767, 32767}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 float64 = float64(((1e10 / ((0.1 + 1e10) / ((((-2.25 * ((-2.25 * float32(float64(-2.25))) - -2.25)) - 1.5) / 0.1) / (float32(float64(1e10)) - (1e10 + float32(((float64((1.5 * -2.25)) - float64(1e10)) * 1.5))))))) / 3.14159))
    var v5 uint16 = 948
    var v6 int8 = -1
    if false {
        delete(v3, 42)
    }
    v7 := GenS1{F0: -1}
    switch -525 {
    case 127:
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 float32 = 1e10
            fmt.Println(v6)
            _ = fi8
            _ = v9
        }
    }
    var v10 uint8 = 216
    for fi11 := 0; fi11 < 6; fi11++ {
        switch "\n\t" {
        case "\n\t":
            var v12 int16 = -68
            _ = v12
        case "":
            v13 := GenS1{F0: fi11}
            _ = v13
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v6 = v6
        v15 := GenS2{F0: fi14, F1: (v6 * 97), F2: 32767}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        if true {
            var v17 int16 = -1
            _ = v17
        } else {
            if true {
                var v18 int32 = 1
                _ = v18
            } else {
                fmt.Println(v10)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 int16 = 42
                    _ = fi19
                    _ = v20
                }
            }
            var v21 float32 = 1e10
            _ = v21
        }
        var v22 uint8 = (v10 | v10)
        _ = fi16
        _ = v22
    }
    v2[609] = (16 > 2)
    _ = len(v1)
    var v23 bool = true
    v24 := GenS2{F0: 42, F1: (v6 | v6), F2: 42}
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 uint64 = 1
        _ = fi25
        _ = v26
    }
    switch 421 {
    case 42:
        if v23 {
            switch ("foo bar" + "hello") {
            case "a":
                v27 := GenS3{F0: 42}
                _ = v27
            case "":
                delete(v3, -9223372036854775808)
            default:
                v4 -= 1e10
            }
            v23 = v23
            v23 = v23
        }
    case 1:
        v28 := GenS4{F0: -1}
        _ = v28
    case -2147483648:
        for fi29 := 0; fi29 < 18; fi29++ {
            _ = len(v2)
            v1 = append(v1, 32768)
            _ = fi29
        }
    default:
        switch 42 {
        case -377:
            v30 := GenS5{F0: (^(v5)), F1: 233}
            _ = v30
        case 833:
            v10 = (v10 - v10)
        case -1:
            fmt.Println(v6, v10, v10)
        default:
            fmt.Println(v10, v6, v23)
        }
    }
    fmt.Println(v23, v4, v23)
    v31 := v2[258]
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 int16 = 1000
        _ = fi32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
