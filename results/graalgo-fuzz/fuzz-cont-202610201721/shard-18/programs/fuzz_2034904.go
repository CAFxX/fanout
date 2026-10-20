package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0() int {
    v4 := GenS1{F0: 2147483647, F1: 123}
    v5 := []int{4294967295, -1}
    _ = v4
    _ = v5
    return 32767
}
func genFn1(p0 int16) int16 {
    fmt.Println(p0)
    v6 := GenS1{F0: 42, F1: 255}
    v7 := []int{genFn0(), -9223372036854775808, 32767}
    for fi8 := 0; fi8 < 14; fi8++ {
        switch fi8 {
        case 32767:
            fmt.Println(p0, p0)
        case -412:
            switch (fi8 / (((fi8 ^ genFn0())) | 1)) {
            case 4294967295:
                var v9 int8 = 1
                _ = v9
            case -285:
                for fi10 := 0; fi10 < 3; fi10++ {
                    switch (^(fi10)) {
                    case 127:
                        switch fi10 {
                        case -2147483648:
                            var v11 uint8 = 42
                            _ = v11
                        case 127:
                            v12 := GenS2{F0: (65535 == 508), F1: 195}
                            _ = v12
                        }
                    default:
                        switch fi10 {
                        case 128:
                            v13 := GenS2{F0: (-128 < 1), F1: 255}
                            _ = v13
                        default:
                            p0 = 1
                        }
                    }
                    fmt.Println(fi8, fi10)
                    var v14 int = fi8
                    _ = fi10
                    _ = v14
                }
            case 996:
                if true {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := GenS2{F0: (508 <= 1), F1: 1}
                        _ = fi15
                        _ = v16
                    }
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        var v17 int8 = 105
                        _ = v17
                    case "":
                        switch (fi8 + fi8) {
                        case 852:
                            var v18 uint = 255
                            _ = v18
                        case 128:
                            if len(v7) > 0 {
                                v7[0] = genFn0()
                            }
                        default:
                            fmt.Println(p0, fi8)
                        }
                    }
                }
            }
        case 128:
            var v19 uint = 193
            _ = v19
        default:
            var v20 float32 = 3.14159
            _ = v20
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v21 := GenS1{F0: 624, F1: 65535}
        _ = fi8
        _ = v21
    }
    v7 = append(v7, 42)
    _ = p0
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483648, 2147483647, 2, 2}
    v3 := []uint32{386, 0, 65535}
    if (((((!false) && ((-2.25 - (-2.25 + 0.1)) < (0.1 + (0.0 + 1.5)))) || false) || (490 == 1)) && (1 == 118)) {
        switch 2147483647 {
        case 42:
            switch ("\n\t" + "a") {
            case "a":
                for fi22 := 0; fi22 < 20; fi22++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(fi22)
                    _ = fi22
                }
            }
        case 968:
            var v23 int16 = 0
            _ = v23
        }
        for fi24 := 0; fi24 < 11; fi24++ {
            _ = len(v2)
            _ = fi24
        }
    }
    var v25 int16 = -32768
    var v26 int = 1
    var v27 string = "hello"
    fmt.Println(v27, v27, v25)
    if (18446744073709551615 != 42) {
        var v28 int = (^(v26))
        var v29 string = v27
        switch genFn0() {
        case -2147483648:
            fmt.Println(v25)
        case -9223372036854775808:
            for fi30 := 0; fi30 < 8; fi30++ {
                v25 = v25
                _ = fi30
            }
        case 2147483647:
            for fi31 := 0; fi31 < 5; fi31++ {
                fmt.Println(v27, v25)
                _ = fi31
            }
        }
        _ = v28
        _ = v29
    }
    v25 -= (v25 | (^(v25)))
    v25 = (v25 * v25)
    v25 = (v25 + 42)
    var v32 int8 = 127
    if true {
        for fi33 := 0; fi33 < 11; fi33++ {
            v25 += v25
            _ = fi33
        }
        v32 += (v32 + 100)
    }
    switch (v27 + "") {
    case "\n\t":
        fmt.Println(v26)
    case "hello":
        v25 -= v25
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v27 = "foo bar"
        _ = fi34
    }
    _ = len(v1)
    if false {
        var v35 float64 = 1e10
        var v36 bool = (3.14159 == (1e10 / 1.5))
        v35 = float64(0.1)
        _ = v35
        _ = v36
    } else {
        v37 := GenS3{F0: 476, F1: 757}
        _ = v37
    }
    v25 = v25
    var v38 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v38
}
