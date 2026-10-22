package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int64 }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int; F2 string }
type GenS6 struct { F0 int8; F1 int16; F2 uint16 }

func main() {
    var v1 string = ("foo bar" + ("foo bar" + ("hello" + "a")))
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    v4 := []string{}
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    v5 := GenS1{F0: 42, F1: 32767, F2: 0}
    v3[("a" + ((v1 + ((v1 + "a") + v1)) + v1))] = 944
    switch 378 {
    case 32768:
        for fi6 := 0; fi6 < 16; fi6++ {
            v4 = append(v4, v1)
            fmt.Println(v1)
            if false {
                v7 := GenS2{F0: 32767, F1: (((3.14159 * (0.0 * ((((-2.25 * 0.1) * 1e10) - 0.0) + 3.14159))) / -2.25) >= float32((-2.25 + -2.25))), F2: 2147483647}
                var v8 uint32 = 4294967295
                v9 := GenS3{F0: (fi6 | (fi6 ^ (fi6 ^ fi6))), F1: (v8 / ((2147483648) | 1))}
                _ = v7
                _ = v8
                _ = v9
            } else {
                var v10 int8 = -60
                v11 := GenS4{F0: ("" + ("foo bar" + ((v1 + v1) + v1)))}
                v2[2147483648] = fi6
                _ = v10
                _ = v11
            }
            _ = fi6
        }
    case -110:
        v1 = (v1 + (((("a" + (v1 + (("hello" + ("" + (v1 + ("foo bar" + v1)))) + v1))) + (((("" + "\n\t") + "a") + v1) + (("" + v1) + v1))) + v1) + (v1 + "hello")))
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS5{F0: 42, F1: (fi12 << ((81) & 63)), F2: v1}
        v14 := GenS5{F0: 255, F1: fi12, F2: ((("" + v1) + ("\n\t" + ("" + "a"))) + v1)}
        for fi15 := 0; fi15 < 20; fi15++ {
            if false {
                v1 = ("a" + v1)
                var v16 int = (fi12 % ((32767) | 1))
                var v17 float32 = ((((0.0 * -2.25) - ((((((-2.25 + (0.0 + 1e10)) - 1e10) - -2.25) / 3.14159) * ((0.1 / 3.14159) / ((0.0 + float32((3.14159 - 1.5))) + 0.0))) + 1.5)) - 1.5) * (-2.25 / 1.5))
                v16 += fi15
                _ = v16
                _ = v17
            } else {
                var v18 int64 = 1
                if false {
                    delete(v2, 2147483647)
                    if (!true) {
                        v1 = v1
                        v18 += 1
                        v1 = (v1 + v1)
                        v1 = v1
                        var v19 bool = true
                        _ = v19
                    }
                }
                v18 += 1
                _ = v18
            }
            _ = fi15
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v20 int16 = 647
    v20 = (v20 * -120)
    if (false || (v1 >= "a")) {
        fmt.Println(v20, v1)
    } else {
        v21 := GenS6{F0: 127, F1: -1, F2: 169}
        _ = v21
    }
    if len(v4) > 0 {
        v4[0] = v1
    }
    switch -1 {
    case 2147483647:
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS3{F0: fi22, F1: 505}
            switch v1 {
            case "foo bar":
                var v24 uint32 = 2147483647
                _ = v24
            case "hello":
                var v25 string = v1
                _ = v25
            case "":
                v4 = append(v4, v1)
            }
            _ = fi22
            _ = v23
        }
    default:
        delete(v2, 1)
    }
    if true {
        if (32767 > v20) {
            var v26 uint8 = 0
            _ = len(v4)
            _ = v26
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
