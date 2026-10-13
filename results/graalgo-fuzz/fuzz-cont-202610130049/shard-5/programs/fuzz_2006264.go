package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 bool }
func genFn0(p0 int64) uint16 {
    switch "foo bar" {
    case "foo bar":
        switch (("foo bar" + "foo bar") + "") {
        case "hello":
            p0 += (-(p0))
        case "foo bar":
            v4 := GenS1{F0: (^(p0)), F1: 350, F2: 0}
            _ = v4
        default:
            p0 = (-(p0))
        }
    }
    switch "foo bar" {
    case "\n\t":
        for fi5 := 0; fi5 < 1; fi5++ {
            v6 := GenS2{F0: 0, F1: 4294967295}
            if true {
                if true {
                    var v7 uint16 = 230
                    fmt.Println(p0, fi5, p0)
                    _ = v7
                }
                fmt.Println(fi5)
            } else {
                p0 += (-(p0))
                if true {
                    var v8 uint8 = 1
                    p0 = p0
                    var v9 int64 = (^(p0))
                    _ = v8
                    _ = v9
                }
                v10 := GenS3{F0: (p0 >> ((385) & 63)), F1: ("foo bar" + "foo bar")}
                _ = v10
            }
            v11 := GenS4{F0: 1000, F1: fi5}
            _ = fi5
            _ = v11
            _ = v6
        }
    case "foo bar":
        fmt.Println(p0)
    case "a":
        v12 := []int{-1, -2147483648, 858, -2147483648}
        _ = v12
    default:
        var v13 float64 = ((1.5 - float64((((1.5 + ((0.1 / 3.14159) - 1.5)) / (0.1 - 3.14159)) - -2.25))) + float64(0.1))
        _ = v13
    }
    if false {
        v14 := []int{133, 127}
        _ = v14
    }
    p0 = (p0 % (((p0 % ((p0) | 1))) | 1))
    p0 -= (p0 % ((-976) | 1))
    _ = p0
    return 1
}
func genFn1(p0 int, p1 uint8) uint16 {
    p1 -= (p1 * p1)
    for fi15 := 0; fi15 < 3; fi15++ {
        switch "a" {
        case "a":
            if (0 >= 4294967295) {
                v16 := []int{fi15, fi15, (fi15 - (fi15 + (p0 % (((p0 << ((18446744073709551615) & 63))) | 1)))), -854}
                var v17 float32 = float32(float64((1.5 - 3.14159)))
                v16 = append(v16, (fi15 | (fi15 | p0)))
                _ = v16
                _ = v17
            } else {
                if false {
                    fmt.Println(p0)
                    p0 = (^(fi15))
                    v18 := []int{(-(fi15))}
                    _ = v18
                } else {
                    p0 = 0
                }
            }
        case "\n\t":
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 bool = ((!false) && (!false))
                _ = fi19
                _ = v20
            }
        case "hello":
            switch (p0 ^ (fi15 ^ (fi15 - p0))) {
            case -1:
                fmt.Println(fi15, fi15, fi15)
            case 9223372036854775807:
                var v21 uint32 = 0
                _ = v21
            default:
                if ((false && true) && (uint32(p1) > 640)) {
                    v22 := GenS5{F0: false}
                    fmt.Println(p0, p1)
                    _ = v22
                }
            }
        }
        p0 -= (-(fi15))
        p1 += (p1 << ((1) & 63))
        _ = fi15
    }
    var v23 uint32 = 2147483647
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 uint16 = 894
        var v26 uint32 = 795
        v27 := []int{fi24, p0}
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v23
    return 32767
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]int)
    var v3 float64 = 1e10
    var v28 string = "hello"
    v29 := []int{0, 32768}
    var v30 uint8 = 42
    if false {
        var v31 int64 = 42
        _ = v31
    } else {
        v28 = (("\n\t" + v28) + v28)
        _ = len(v29)
    }
    var v32 uint = 18446744073709551615
    fmt.Println(v3, v1)
    v33 := GenS5{F0: true}
    v3 += -2.25
    if true {
        v3 = (v3 * v3)
        fmt.Println(v3)
    }
    v34 := GenS1{F0: 9223372036854775807, F1: (v1 << ((v32) & 63)), F2: 639}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
}
