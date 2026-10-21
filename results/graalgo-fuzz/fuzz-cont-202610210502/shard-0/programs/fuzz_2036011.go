package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0(p0 uint8, p1 int8) int32 {
    if (!true) {
        p0 += p0
        var v5 int8 = (p1 ^ p1)
        _ = v5
    } else {
        if true {
            var v6 uint16 = 1
            _ = v6
        } else {
            v7 := GenS1{F0: (p1 >> ((772) & 63))}
            v8 := GenS1{F0: (p1 ^ p1)}
            _ = v7
            _ = v8
        }
    }
    v9 := []int{}
    p1 += p1
    p0 = (p0 ^ 2)
    _ = len(v9)
    _ = p0
    _ = p1
    _ = v9
    return -622
}
func genFn1(p0 uint8, p1 uint16, p2 bool) bool {
    switch "a" {
    case "a":
        var v10 int16 = int16(p1)
        _ = v10
    default:
        v11 := GenS1{F0: -128}
        _ = v11
    }
    v12 := []int{}
    var v13 int32 = 1
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 string = ("foo bar" + ("" + ("hello" + "a")))
        _ = fi14
        _ = v15
    }
    if (-698 >= 0) {
        p1 -= 65535
        p1 = 30
        if ((^(p0)) > (p0 % ((0) | 1))) {
            v16 := GenS1{F0: 0}
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return (18446744073709551615 == 4294967295)
}

func main() {
    var v1 int8 = 1
    var v2 uint64 = 897
    var v3 uint16 = 42
    var v4 float32 = 0.0
    v17 := []int{9223372036854775807}
    v2 = (^(v2))
    var v18 uint64 = (v2 % ((v2) | 1))
    var v19 bool = false
    switch 1 {
    case -9223372036854775808:
        var v20 uint8 = 0
        _ = v20
    case 0:
        v21 := GenS1{F0: (v1 & 0)}
        _ = v21
    default:
        var v22 int = 128
        _ = v22
    }
    var v23 int16 = -1
    for fi24 := 0; fi24 < 16; fi24++ {
        v23 += (v23 ^ (v23 ^ (v23 ^ 47)))
        for fi25 := 0; fi25 < 10; fi25++ {
            if (!((((1.5 - float64((v4 * 0.0))) + (1e10 * 0.1)) - 1.5) != (((((-2.25 / (-2.25 - float64(v4))) * float64(1.5)) + 0.1) / 0.1) * 0.0))) {
                _ = len(v17)
                var v26 uint = 2
                _ = v26
            }
            _ = fi25
        }
        _ = fi24
    }
    v3 = (v3 % ((255) | 1))
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 int32 = -2147483648
        for fi29 := 0; fi29 < 3; fi29++ {
            v17 = append(v17, fi27)
            _ = fi29
        }
        v30 := GenS2{F0: 624, F1: v19}
        _ = fi27
        _ = v28
        _ = v30
    }
    v19 = true
    v19 = v19
    var v31 float64 = 1.5
    fmt.Println(v23, v31)
    v17 = append(v17, 4294967295)
    switch "foo bar" {
    case "":
        var v32 uint8 = 115
        _ = v32
    case "\n\t":
        v23 = 1000
    }
    var v33 int8 = -1
    v23 = v23
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v33
    _ = v4
}
