package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 int, p1 string, p2 uint32) bool {
    var v4 int8 = -110
    v5 := GenS1{F0: 96, F1: 9223372036854775807, F2: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (32767 > -443)
}

func main() {
    v1 := []int{332, -831, -9223372036854775808, -759}
    var v2 int16 = 1
    v3 := []string{}
    var v6 int16 = v2
    v6 -= 42
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 6; fi8++ {
        if false {
            for fi9 := 0; fi9 < 6; fi9++ {
                fmt.Println(v7, fi8)
                v10 := GenS1{F0: 100, F1: v7, F2: -2147483648}
                _ = fi9
                _ = v10
            }
            v6 -= v6
        } else {
            var v11 float64 = (-2.25 / (0.1 - (float64(3.14159) * 0.1)))
            var v12 int8 = 80
            for fi13 := 0; fi13 < 2; fi13++ {
                v11 = -2.25
                switch "hello" {
                case "":
                    var v14 float64 = (v11 - v11)
                    _ = v14
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "a":
                    switch 2147483647 {
                    case 9223372036854775807:
                        fmt.Println(v12, fi8, v7)
                    case 154:
                        switch (fi13 ^ fi8) {
                        case -9223372036854775808:
                            switch (fi8 & fi13) {
                            case 9223372036854775807:
                                v15 := GenS1{F0: (v12 - int8(v2)), F1: -1, F2: 1}
                                _ = v15
                            case -492:
                                v12 = (v12 - v12)
                            }
                        }
                    case 2147483647:
                        var v16 uint64 = 0
                        _ = v16
                    }
                default:
                    if false {
                        v17 := GenS1{F0: (v12 << ((65535) & 63)), F1: (v7 / ((v7) | 1)), F2: 1}
                        if true {
                            v7 = (v7 * (^(v7)))
                        }
                        _ = v17
                    } else {
                        if true {
                            fmt.Println(v11)
                            switch (("a" + "") + ("a" + "")) {
                            case "":
                                var v18 int16 = -32768
                                _ = v18
                            default:
                                var v19 uint = 233
                                _ = v19
                            }
                            var v20 int16 = 32767
                            _ = v20
                        }
                    }
                }
                switch fi8 {
                case -136:
                    var v21 int8 = (v12 * int8(v7))
                    _ = v21
                case 1:
                    v11 -= 0.1
                case -29:
                    switch (fi13 % ((fi13) | 1)) {
                    case 2147483647:
                        v2 -= (^(v6))
                    case 753:
                        var v22 string = ("foo bar" + "")
                        _ = v22
                    }
                }
                _ = fi13
            }
            _ = v11
            _ = v12
        }
        if false {
            v7 -= 0
        } else {
            var v23 string = ""
            var v24 int8 = 42
            _ = v23
            _ = v24
        }
        switch "\n\t" {
        case "foo bar":
            if false {
                var v25 int16 = int16(fi8)
                _ = v25
            }
        case "hello":
            v2 = (v6 >> ((18446744073709551615) & 63))
        default:
            v26 := GenS1{F0: -128, F1: (v7 / (((-(v7))) | 1)), F2: 2147483647}
            _ = v26
        }
        _ = fi8
    }
    v27 := GenS1{F0: -1, F1: v7, F2: -1}
    v28 := GenS2{F0: 127, F1: 2, F2: 42}
    _ = len(v3)
    v2 -= 42
    v29 := GenS3{F0: 1, F1: 127, F2: v7}
    if false {
        v6 = v2
        v1 = append(v1, 9223372036854775807)
    } else {
        v1 = append(v1, 504)
    }
    var v30 int8 = -101
    var v31 uint8 = 201
    v32 := GenS4{F0: "foo bar", F1: 4294967295}
    var v33 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v6
    _ = v7
}
