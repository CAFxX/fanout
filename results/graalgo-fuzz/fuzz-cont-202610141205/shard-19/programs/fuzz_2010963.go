package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 int16 }
func genFn0(p0 bool, p1 int32) uint32 {
    for fi5 := 0; fi5 < 8; fi5++ {
        switch fi5 {
        case -9223372036854775808:
            var v6 uint16 = 300
            _ = v6
        case 127:
            p0 = p0
        case 2147483647:
            v7 := []int{(fi5 % ((32767) | 1)), (fi5 >> ((70) & 63)), fi5}
            _ = v7
        }
        _ = fi5
    }
    switch -600 {
    case 32768:
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 int = (-(fi8))
                if true {
                    if p0 {
                        var v11 uint64 = 42
                        _ = v11
                    }
                    if true {
                        v12 := GenS1{F0: -709, F1: 694, F2: 891}
                        _ = v12
                    } else {
                        var v13 uint32 = 42
                        var v14 float32 = 0.1
                        if ((1e10 >= float32(3.14159)) || p0) {
                            v10 = v10
                            fmt.Println(fi9)
                            if p0 {
                                v15 := []int{(fi8 << ((931) & 63)), (v10 / ((fi8) | 1)), fi9}
                                var v16 int16 = 1000
                                var v17 int32 = p1
                                _ = v15
                                _ = v16
                                _ = v17
                            } else {
                                var v18 string = "foo bar"
                                _ = v18
                            }
                        }
                        _ = v13
                        _ = v14
                    }
                    v19 := GenS2{F0: 65535, F1: 65535}
                    _ = v19
                }
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case 9223372036854775807:
        if (((^(p1)) < (p1 - (p1 * 1))) || p0) {
            var v20 string = "\n\t"
            _ = v20
        } else {
            v21 := GenS1{F0: 1, F1: 0, F2: 0}
            v22 := []int{32768, -2147483648}
            _ = v21
            _ = v22
        }
    case -924:
        switch "" {
        case "a":
            fmt.Println(p0, p1)
        case "":
            p0 = true
        case "\n\t":
            p1 += (p1 + p1)
        default:
            v23 := GenS1{F0: 32767, F1: 2, F2: 1}
            _ = v23
        }
    }
    var v24 int32 = 1
    _ = p0
    _ = p1
    _ = v24
    return 2147483648
}

func main() {
    var v1 float64 = (((0.0 - ((0.1 + 3.14159) / 1.0)) * (1.5 / 1.0)) + float64((((-2.25 - float32(0.0)) - 0.0) - float32(1e10))))
    var v2 uint8 = 125
    v3 := []string{("foo bar" + ("hello" + ("" + "\n\t"))), ("" + "foo bar")}
    v4 := []int{}
    var v25 int32 = 0
    switch ("\n\t" + ("foo bar" + (("" + ("foo bar" + "hello")) + "hello"))) {
    case "a":
        switch -418 {
        case 32767:
            var v26 bool = false
            _ = v26
        case 128:
            var v27 float32 = 1e10
            _ = v27
        }
    case "hello":
        v25 += -90
    }
    v25 = v25
    v25 = v25
    v4 = append(v4, 42)
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 float64 = (3.14159 + 1e10)
        _ = fi28
        _ = v29
    }
    fmt.Println(v25, v1)
    v30 := GenS2{F0: 2, F1: 304}
    var v31 int64 = -584
    v1 -= v1
    var v32 bool = (!(!false))
    fmt.Println(v25, v1, v32)
    if len(v4) > 0 {
        v4[0] = 793
    }
    v33 := GenS3{F0: 321, F1: 100, F2: 380}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
