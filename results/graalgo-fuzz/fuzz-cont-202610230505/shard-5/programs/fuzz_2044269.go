package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 uint32, p1 uint32) int {
    switch "\n\t" {
    case "hello":
        p1 = 2
    case "foo bar":
        v4 := GenS1{F0: p1, F1: 1}
        _ = v4
    case "\n\t":
        var v5 int = -2147483648
        _ = v5
    default:
        var v6 int32 = -2147483648
        _ = v6
    }
    v7 := []int{0}
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 int8) int8 {
    v8 := []int{510, 127}
    p0 = p0
    v8 = append(v8, 128)
    var v9 int8 = 127
    var v10 float64 = 1e10
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int8 = 127
    v3 := []uint32{18, (v1 | (v1 << ((4294967295) & 63))), v1}
    var v11 int32 = 1
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int16 = 1
        _ = fi12
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 3; fi14++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi14
        }
        var v15 string = ""
        var v16 bool = (2 > 65535)
        _ = v15
        _ = v16
    } else {
        if (!(0 != 9223372036854775807)) {
            var v17 uint = 4294967295
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := GenS1{F0: (v1 ^ v1), F1: 4294967295}
                _ = fi18
                _ = v19
            }
            v2 -= genFn1((-(v2)))
            _ = v17
        }
        fmt.Println(v2)
    }
    v20 := GenS2{F0: 65535, F1: (v1 & v1), F2: 560}
    fmt.Println(v2)
    for fi21 := 0; fi21 < 4; fi21++ {
        if (148 < v1) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                v11 += (v11 * 644)
                v1 = 1
                v23 := GenS1{F0: (v1 & 2147483648), F1: 42}
                _ = fi22
                _ = v23
            }
        } else {
            v24 := GenS2{F0: 42, F1: v1, F2: 18446744073709551615}
            if false {
                v11 = -1
            }
            _ = v24
        }
        _ = fi21
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        for fi26 := 0; fi26 < 11; fi26++ {
            _ = len(v3)
            v1 += v1
            _ = fi26
        }
        v27 := GenS1{F0: (v1 / ((v1) | 1)), F1: 42}
        v1 += 2147483647
        _ = fi25
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        v1 = v1
        fmt.Println(fi28, v11, v11)
        v29 := GenS3{F0: (!true), F1: ((("a" + ("\n\t" + "")) + (((((("a" + "") + "a") + ("" + "a")) + "a") + ("\n\t" + "hello")) + ("foo bar" + ""))) + "")}
        _ = fi28
        _ = v29
    }
    fmt.Println(v2, v11)
    var v30 float32 = 0.0
    v31 := GenS3{F0: true, F1: ("hello" + "hello")}
    fmt.Println(v2)
    var v32 float32 = v30
    if len(v3) > 0 {
        v3[0] = v1
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 uint = 904
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v32 += 0.1
        _ = len(v3)
        v30 += v30
        _ = fi35
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
