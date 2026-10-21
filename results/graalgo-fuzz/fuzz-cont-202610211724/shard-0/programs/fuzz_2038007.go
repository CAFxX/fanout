package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int32; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint8; F2 string }
func genFn0(p0 uint16, p1 int16) uint8 {
    p0 = p0
    p0 += 2
    for fi5 := 0; fi5 < 10; fi5++ {
        if false {
            v6 := []int{-9223372036854775808}
            if true {
                var v7 uint = 255
                v8 := GenS1{F0: fi5}
                v9 := GenS2{F0: fi5, F1: 388, F2: (fi5 | int(p0))}
                _ = v7
                _ = v8
                _ = v9
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            p0 = 255
            _ = v6
        }
        switch fi5 {
        case 1:
            fmt.Println(p1, fi5, p0)
        case -2147483648:
            if true {
                switch "\n\t" {
                case "":
                    var v10 int16 = p1
                    _ = v10
                case "hello":
                    switch "foo bar" {
                    case "foo bar":
                        var v11 int16 = -1
                        _ = v11
                    case "\n\t":
                        var v12 uint8 = 0
                        _ = v12
                    case "":
                        p0 -= (p0 ^ p0)
                    }
                }
                var v13 int32 = 42
                _ = v13
            } else {
                v14 := GenS1{F0: (-(fi5))}
                _ = v14
            }
        default:
            v15 := GenS2{F0: 128, F1: -2147483648, F2: (fi5 % ((fi5) | 1))}
            _ = v15
        }
        var v16 float64 = float64(1e10)
        _ = fi5
        _ = v16
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int32 = -422
    var v2 int16 = 32767
    var v3 int32 = (-(v1))
    var v4 int16 = (v2 ^ 0)
    switch 32768 {
    case 128:
        if ((!false) || (-128 > 42)) {
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 uint64 = 0
                _ = fi17
                _ = v18
            }
        }
    case 437:
        var v19 int32 = (v1 >> ((783) & 63))
        _ = v19
    case -9223372036854775808:
        var v20 uint = 18446744073709551615
        _ = v20
    }
    v21 := []int{2147483647}
    switch "a" {
    case "hello":
        var v22 int16 = v2
        _ = v22
    case "":
        if false {
            var v23 int16 = 710
            _ = v23
        }
    }
    v21 = append(v21, -1)
    if true {
        v3 -= (^(v1))
        fmt.Println(v3, v2)
        if (true && (!false)) {
            fmt.Println(v4)
        } else {
            _ = len(v21)
            v1 -= (-(v3))
        }
    } else {
        fmt.Println(v2, v3, v1)
        var v24 int32 = v3
        if len(v21) > 0 {
            v21[0] = -9223372036854775808
        }
        _ = v24
    }
    switch 32768 {
    case 0:
        var v25 uint64 = 0
        _ = v25
    case 1:
        switch "hello" {
        case "\n\t":
            for fi26 := 0; fi26 < 16; fi26++ {
                v2 -= (v2 - 1000)
                switch -318 {
                case 42:
                    v3 -= 0
                case 429:
                    v27 := GenS2{F0: (fi26 | fi26), F1: -675, F2: -2147483648}
                    _ = v27
                default:
                    if true {
                        fmt.Println(v3, v1)
                        v28 := GenS3{F0: ("\n\t" + "\n\t")}
                        _ = v28
                    }
                }
                if (271 == 1) {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        var v30 float32 = 0.0
                        _ = fi29
                        _ = v30
                    }
                    if false {
                        v1 = (v1 | (v3 * (v1 + (v3 + v1))))
                    } else {
                        var v31 uint8 = 128
                        _ = v31
                    }
                } else {
                    fmt.Println(fi26)
                }
                _ = fi26
            }
        }
    case 42:
        v4 += v2
    }
    v2 = v2
    var v32 float64 = 3.14159
    if (v3 < v1) {
        var v33 int64 = -1
        v1 -= -2147483648
        fmt.Println(v32)
        _ = v33
    } else {
        fmt.Println(v32, v2)
    }
    v34 := GenS4{F0: 128, F1: 1, F2: (("\n\t" + ("a" + ("a" + "foo bar"))) + ("hello" + (("" + "") + "\n\t")))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v34
    _ = v4
}
