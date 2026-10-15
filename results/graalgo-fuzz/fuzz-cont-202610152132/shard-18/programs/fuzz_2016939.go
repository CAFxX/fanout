package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 bool }
func genFn0(p0 int32, p1 int8, p2 int32) uint16 {
    var v5 int = -17
    for fi6 := 0; fi6 < 15; fi6++ {
        v5 += fi6
        switch (("hello" + "a") + "foo bar") {
        case "":
            v7 := GenS1{F0: 234}
            _ = v7
        case "a":
            for fi8 := 0; fi8 < 13; fi8++ {
                fmt.Println(fi6, v5)
                _ = fi8
            }
        case "foo bar":
            v9 := []int{fi6, (v5 ^ 42), (v5 * fi6)}
            _ = v9
        default:
            v5 = (-(fi6))
        }
        v10 := GenS2{F0: 192, F1: p1}
        _ = fi6
        _ = v10
    }
    p1 -= p1
    p1 = 1
    p2 = (p0 | (p0 * (p0 - -718)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}
func genFn1() uint8 {
    v11 := GenS1{F0: 18446744073709551615}
    v12 := GenS3{F0: -543}
    var v13 uint16 = 912
    v13 += v13
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int32 = -930
        _ = fi14
        _ = v15
    }
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := []int{}
    var v2 int8 = -32
    var v3 uint64 = 414
    var v4 int64 = 42
    for fi16 := 0; fi16 < 7; fi16++ {
        v4 -= v4
        var v17 uint32 = 447
        switch "foo bar" {
        case "\n\t":
            v17 = (^(v17))
        case "hello":
            v3 += (v3 - 703)
        case "a":
            v3 = v3
        default:
            v18 := GenS3{F0: fi16}
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    v3 -= v3
    switch 32768 {
    case 128:
        if false {
            for fi19 := 0; fi19 < 3; fi19++ {
                for fi20 := 0; fi20 < 20; fi20++ {
                    v3 -= 4294967295
                    if len(v1) > 0 {
                        v1[0] = fi20
                    }
                    _ = fi20
                }
                var v21 uint32 = 2147483647
                _ = fi19
                _ = v21
            }
            v22 := GenS4{F0: "hello", F1: 508}
            v4 -= (^(v4))
            _ = v22
        } else {
            v4 += 42
            v23 := GenS1{F0: v3}
            _ = v23
        }
    case 4294967295:
        fmt.Println(v2)
    case 2147483647:
        for fi24 := 0; fi24 < 6; fi24++ {
            switch (-(fi24)) {
            case -4:
                fmt.Println(v2, v3, v2)
            }
            v4 = (v4 - v4)
            switch ("foo bar" + "") {
            case "a":
                for fi25 := 0; fi25 < 4; fi25++ {
                    if false {
                        v1 = append(v1, fi25)
                        v26 := GenS5{F0: (65535 < 2)}
                        _ = v26
                    }
                    v4 -= v4
                    _ = fi25
                }
            case "\n\t":
                fmt.Println(v3)
            default:
                switch "" {
                case "a":
                    fmt.Println(v2, v2)
                default:
                    var v27 int64 = v4
                    _ = v27
                }
            }
            _ = fi24
        }
    default:
        var v28 uint = 4294967295
        _ = v28
    }
    var v29 int32 = 1
    v1 = append(v1, 4294967295)
    for fi30 := 0; fi30 < 5; fi30++ {
        fmt.Println(fi30)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi30
    }
    fmt.Println(v2)
    if ((3.14159 / (1e10 + 1e10)) > 1e10) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 = v2
    var v31 uint32 = 0
    v4 = (v4 / ((v4) | 1))
    if (true || false) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v29 = v29
        if (2 < 0) {
            var v32 float64 = ((0.1 - ((3.14159 + (-2.25 * (1e10 * 1.5))) / 1.0)) / 1.5)
            _ = v32
        }
    }
    v2 += v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v33 int8 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v33
    _ = v4
}
