package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint }
func genFn0() string {
    var v5 int = -387
    var v6 int32 = 42
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1() int8 {
    switch 4294967295 {
    case 127:
        v7 := GenS1{F0: 385, F1: (-1 > 105)}
        _ = v7
    default:
        if (true && (!(true && (!true)))) {
            v8 := GenS2{F0: ("hello" + "")}
            v9 := []int{232}
            if true {
                var v10 uint16 = 788
                _ = v10
            }
            _ = v8
            _ = v9
        }
    }
    var v11 int64 = -773
    var v12 int8 = 127
    v11 -= (v11 - v11)
    _ = v11
    _ = v12
    return (v12 ^ v12)
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 65535
    var v3 uint64 = 1
    var v4 int = -8
    var v13 uint = 327
    if ("" < ("foo bar" + "hello")) {
        v3 = v3
        var v14 int64 = -71
        var v15 uint16 = 542
        _ = v14
        _ = v15
    } else {
        v16 := GenS3{F0: (false || ((^(v3)) == (^(v3)))), F1: (v13 & 117), F2: 0}
        _ = v16
    }
    v17 := GenS1{F0: 4294967295, F1: true}
    if false {
        for fi18 := 0; fi18 < 5; fi18++ {
            if true {
                var v19 uint8 = 128
                var v20 int64 = -1
                _ = len(v1)
                _ = v19
                _ = v20
            }
            if false {
                var v21 float32 = 3.14159
                _ = v21
            }
            _ = fi18
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch ("hello" + ("\n\t" + "hello")) {
            case "a":
                fmt.Println(v13, fi22, v4)
            case "hello":
                var v23 int8 = genFn1()
                _ = v23
            case "\n\t":
                var v24 float32 = 0.0
                _ = v24
            default:
                v2 -= v2
            }
            _ = fi22
        }
    }
    fmt.Println(v2)
    v13 = 65535
    var v25 uint = v13
    switch "hello" {
    case "a":
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 bool = true
            v25 = v13
            if len(v1) > 0 {
                v1[0] = (v2 - 2147483647)
            }
            _ = fi26
            _ = v27
        }
    }
    if true {
        var v28 float64 = 3.14159
        if true {
            for fi29 := 0; fi29 < 2; fi29++ {
                v2 -= (v2 | (v2 & (v2 | v2)))
                var v30 int16 = 1000
                v31 := GenS1{F0: (v3 + 248), F1: ((42 > v25) || false)}
                _ = fi29
                _ = v30
                _ = v31
            }
            if true {
                v1 = append(v1, v2)
            }
            v25 = (v25 + (v13 ^ 18446744073709551615))
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            v3 = (v3 >> ((255) & 63))
            _ = fi32
        }
        _ = v28
    }
    if len(v1) > 0 {
        v1[0] = uint32(v3)
    }
    fmt.Println(v2, v13, v25)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 32767 {
    case -2147483648:
        fmt.Println(v3)
    case 128:
        v1 = append(v1, v2)
    case -682:
        v2 -= 917
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        v34 := GenS4{F0: 638, F1: 255}
        v13 = v25
        _ = fi33
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
