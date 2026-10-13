package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int8 }
func genFn0(p0 string, p1 uint32) uint64 {
    switch -9223372036854775808 {
    case -234:
        var v7 int = 2147483647
        _ = v7
    default:
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                p0 = "\n\t"
                _ = fi9
            }
            var v10 uint16 = 170
            _ = fi8
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        if false {
            v12 := GenS1{F0: "\n\t", F1: 785, F2: -128}
            _ = v12
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 int32, p1 uint64, p2 bool) uint8 {
    switch "hello" {
    case "foo bar":
        var v13 float32 = -2.25
        _ = v13
    case "\n\t":
        v14 := GenS1{F0: "\n\t", F1: 686, F2: 0}
        _ = v14
    case "a":
        var v15 string = ("\n\t" + "\n\t")
        _ = v15
    default:
        v16 := GenS1{F0: "\n\t", F1: 134, F2: 81}
        _ = v16
    }
    switch "\n\t" {
    case "\n\t":
        if true {
            var v17 float32 = ((0.0 + (0.0 - -2.25)) * 0.1)
            var v18 bool = false
            p1 += 42
            _ = v17
            _ = v18
        }
    case "foo bar":
        switch "foo bar" {
        case "hello":
            v19 := GenS1{F0: "hello", F1: 255, F2: -128}
            _ = v19
        case "":
            if (p2 || (!(29 == 1))) {
                p0 -= p0
                var v20 string = "a"
                _ = v20
            }
        case "\n\t":
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 int = (-(fi21))
                _ = fi21
                _ = v22
            }
        default:
            if p2 {
                p2 = (((970 <= 670) && ((p2 || (615 >= 0)) && (!(65535 != 474)))) || p2)
                var v23 uint8 = 127
                var v24 uint16 = 909
                _ = v23
                _ = v24
            }
        }
    default:
        p1 = 4294967295
    }
    var v25 string = ""
    v26 := GenS1{F0: "", F1: 155, F2: 100}
    for fi27 := 0; fi27 < 20; fi27++ {
        p1 += 2
        var v28 uint32 = 2147483647
        var v29 uint16 = 752
        _ = fi27
        _ = v28
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return 128
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{65535, 757, 2147483648}
    var v3 uint8 = 0
    var v4 uint32 = 1
    var v5 float32 = float32(0.1)
    var v6 int32 = 42
    fmt.Println(v4, v4, v3)
    var v30 string = "hello"
    var v31 float64 = (0.1 * 1.5)
    v3 = v3
    var v32 bool = ((!(uint32(v3) != (v4 << ((42) & 63)))) || (!(42 < 42)))
    switch "\n\t" {
    case "foo bar":
        if v32 {
            v6 = (v6 % ((v6) | 1))
            var v33 int16 = 1
            _ = v33
        } else {
            var v34 uint32 = (^(v4))
            _ = v34
        }
    case "":
        switch v30 {
        case "foo bar":
            v35 := GenS1{F0: v30, F1: 639, F2: 1}
            _ = v35
        case "\n\t":
            v5 = -2.25
        }
    }
    fmt.Println(v5, v3)
    v30 = (v30 + ("foo bar" + ((v30 + v30) + v30)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
