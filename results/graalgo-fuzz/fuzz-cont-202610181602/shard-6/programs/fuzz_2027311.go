package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint; F1 uint64; F2 int32 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint16) int8 {
    p0 -= p0
    var v5 int8 = 80
    var v6 int16 = 1000
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 string, p1 uint32, p2 int32) bool {
    var v7 int16 = -71
    var v8 uint64 = 940
    if ((^(p2)) == (p2 & p2)) {
        v8 = v8
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 string = "\n\t"
            v7 = v7
            v8 = 4294967295
            _ = fi9
            _ = v10
        }
        p1 += p1
    }
    v11 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    _ = v8
    return (80 >= 42)
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]uint32)
    var v3 int16 = -1
    v4 := make(map[string]string)
    v12 := v4["hello"]
    v13 := GenS2{F0: 605, F1: 685, F2: v3}
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint = 727
        _ = fi14
        _ = v15
    }
    switch (((("foo bar" + v12) + v12) + v12) + "") {
    case "\n\t":
        if (!(9223372036854775807 < 77)) {
            v1 = float64((1e10 / 3.14159))
            v16 := []int{42}
            _ = v16
        }
    case "a":
        for fi17 := 0; fi17 < 14; fi17++ {
            if (0 > 42) {
                var v18 int = (fi17 / (((^(fi17))) | 1))
                _ = v18
            }
            var v19 int = fi17
            _ = fi17
            _ = v19
        }
    case "foo bar":
        var v20 float32 = -2.25
        _ = v20
    default:
        v21 := GenS3{F0: 18446744073709551615, F1: 18446744073709551615, F2: 987}
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 bool = true
        _ = fi22
        _ = v23
    }
    var v24 uint32 = 2147483647
    for fi25 := 0; fi25 < 6; fi25++ {
        switch v12 {
        case "hello":
            for fi26 := 0; fi26 < 17; fi26++ {
                v1 += float64(0.1)
                var v27 uint64 = 4294967295
                _ = fi26
                _ = v27
            }
        case "foo bar":
            var v28 uint64 = 1
            _ = v28
        }
        v12 = "foo bar"
        fmt.Println(v12, v24, fi25)
        _ = fi25
    }
    switch "\n\t" {
    case "":
        if (false || false) {
            v29 := GenS3{F0: 0, F1: 408, F2: 197}
            _ = v29
        } else {
            v12 = "\n\t"
            var v30 float64 = (v1 - 3.14159)
            v3 = -32768
            _ = v30
        }
    case "a":
        var v31 float32 = 3.14159
        _ = v31
    case "hello":
        v24 = (v24 & (^(v24)))
    default:
        if true {
            v1 += ((v1 + -2.25) + v1)
            fmt.Println(v12, v24, v24)
            var v32 float64 = v1
            _ = v32
        }
    }
    var v33 uint16 = 2
    var v34 int64 = 1
    fmt.Println(v3, v34)
    v35 := GenS2{F0: uint32(v33), F1: 2, F2: v3}
    v1 = (float64(0.1) - v1)
    v2[(^(v24))] = (v24 ^ (^(v24)))
    v36 := GenS4{F0: false}
    v12 = ((((("foo bar" + v12) + ("\n\t" + "hello")) + ("a" + v12)) + ((("\n\t" + ("foo bar" + ((v12 + (("" + (v12 + ("hello" + v12))) + "")) + "hello"))) + v12) + (v12 + "foo bar"))) + v12)
    v37 := []int{-2147483648, -265}
    v1 -= float64(((((0.0 + 3.14159) - float32(0.0)) / (1e10 * -2.25)) - 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
