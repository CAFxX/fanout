package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 858
    var v3 float64 = 1e10
    v4 := []int{}
    v5 := []string{}
    if (!true) {
        v6 := GenS1{F0: true, F1: 924, F2: (v2 + 0)}
        v7 := GenS1{F0: true, F1: 0, F2: 774}
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: (v3 >= v3), F1: 1, F2: (v2 / ((v2) | 1))}
    if ((!(true && ((18446744073709551615 > 2) && (65535 != 0)))) && false) {
        switch "" {
        case "a":
            switch -9223372036854775808 {
            case 32768:
                var v9 int = 4294967295
                _ = v9
            case 128:
                v10 := GenS1{F0: false, F1: 4294967295, F2: 2147483648}
                _ = v10
            case -2147483648:
                var v11 int = 42
                _ = v11
            }
        case "hello":
            var v12 uint = 42
            _ = v12
        case "\n\t":
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v1 = ("foo bar" + v1)
            } else {
                for fi13 := 0; fi13 < 2; fi13++ {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        if len(v5) > 0 {
                            v5[0] = v1
                        }
                        var v15 float64 = 0.0
                        _ = len(v4)
                        _ = fi14
                        _ = v15
                    }
                    _ = fi13
                }
                var v16 int64 = 9223372036854775807
                fmt.Println(v16, v16, v2)
                _ = v16
            }
        default:
            if false {
                switch ("a" + "a") {
                case "a":
                    v17 := GenS1{F0: (!true), F1: uint64(v2), F2: 283}
                    _ = v17
                case "\n\t":
                    switch "hello" {
                    case "hello":
                        var v18 int64 = 0
                        _ = v18
                    }
                }
                v1 = ("hello" + (("\n\t" + ("" + v1)) + ("hello" + v1)))
            }
        }
        fmt.Println(v3)
    }
    var v19 int64 = 1
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 uint8 = 182
        v19 -= v19
        v2 += v2
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: (4294967295 <= 18446744073709551615), F1: 156, F2: (v2 | 2)}
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v3, v3)
    case "\n\t":
        switch 32767 {
        case 32767:
            v3 = 1e10
        case -669:
            if (2 >= 42) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (true || (-187 < 221)) {
                    v1 = v1
                    if (1 >= 972) {
                        v23 := GenS1{F0: true, F1: 0, F2: (^(v2))}
                        v24 := GenS2{F0: 176, F1: v2, F2: v19}
                        if true {
                            v25 := GenS3{F0: -1, F1: 128}
                            _ = v25
                        } else {
                            _ = len(v5)
                        }
                        _ = v23
                        _ = v24
                    } else {
                        v26 := GenS2{F0: 42, F1: 42, F2: 1}
                        _ = v26
                    }
                    switch "foo bar" {
                    case "a":
                        v27 := GenS1{F0: true, F1: 42, F2: (v2 / ((v2) | 1))}
                        _ = v27
                    }
                }
            }
        case 1:
            if true {
                v19 += (^(v19))
            } else {
                var v28 uint8 = 255
                _ = v28
            }
        }
    default:
        v2 = 2147483647
    }
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = 127
    }
    if len(v4) > 0 {
        v4[0] = -743
    }
    fmt.Println(v3)
    v29 := GenS4{F0: 785}
    v19 = v19
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 int32 = 0
        v32 := GenS2{F0: 2, F1: 65535, F2: v19}
        for fi33 := 0; fi33 < 2; fi33++ {
            v2 -= (v2 * v2)
            v31 += v31
            var v34 uint32 = v2
            _ = fi33
            _ = v34
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
