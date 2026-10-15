package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int32; F2 uint }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int32; F1 uint16; F2 string }

func main() {
    var v1 int16 = 1
    var v2 float64 = -2.25
    var v3 uint32 = 332
    var v4 float64 = v2
    v4 -= 3.14159
    v5 := []int{}
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: (v1 * (v1 & v1))}
        switch "foo bar" {
        case "hello":
            var v8 int8 = -128
            _ = v8
        default:
            if true {
                if (!false) {
                    v2 = v4
                    for fi9 := 0; fi9 < 15; fi9++ {
                        var v10 int16 = (v1 * v1)
                        _ = fi9
                        _ = v10
                    }
                }
                if (-1 == 0) {
                    switch ("hello" + "hello") {
                    case "foo bar":
                        if true {
                            v11 := GenS1{F0: v1}
                            v4 += v2
                            _ = v11
                        } else {
                            fmt.Println(v1)
                            v1 += v1
                            var v12 int = fi6
                            _ = v12
                        }
                    case "a":
                        var v13 int32 = 0
                        _ = v13
                    }
                    v3 = (^(v3))
                    if len(v5) > 0 {
                        v5[0] = fi6
                    }
                }
            }
        }
        var v14 string = "\n\t"
        _ = fi6
        _ = v14
        _ = v7
    }
    fmt.Println(v2, v1, v4)
    if false {
        v4 += 1.5
        v5 = append(v5, 32768)
        for fi15 := 0; fi15 < 10; fi15++ {
            v2 -= 1.5
            _ = fi15
        }
    } else {
        v2 = v4
        switch "\n\t" {
        case "foo bar":
            v3 = v3
        }
        _ = len(v5)
    }
    v3 += v3
    switch "a" {
    case "foo bar":
        v3 -= (v3 % ((1) | 1))
    }
    if len(v5) > 0 {
        v5[0] = 42
    }
    switch 42 {
    case 2147483647:
        switch 42 {
        case -132:
            if false {
                v16 := GenS2{F0: -686}
                _ = v16
            } else {
                switch "foo bar" {
                case "foo bar":
                    v17 := GenS3{F0: 18446744073709551615, F1: 0, F2: 930}
                    _ = v17
                case "":
                    v18 := GenS4{F0: 77, F1: 4294967295}
                    _ = v18
                case "a":
                    v19 := GenS5{F0: -456, F1: 255, F2: (((("a" + "foo bar") + "") + "hello") + "")}
                    _ = v19
                }
                v1 += (v1 % ((v1) | 1))
            }
        case -391:
            v20 := GenS4{F0: 235, F1: 16}
            _ = v20
        }
    case -1:
        fmt.Println(v4, v3)
    case 1:
        fmt.Println(v4, v4)
    }
    v4 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
