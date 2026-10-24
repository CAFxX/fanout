package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0(p0 int8) string {
    p0 = p0
    fmt.Println(p0)
    var v4 int64 = -1
    if ("" == ("foo bar" + "")) {
        v4 = -406
        v5 := []int{}
        _ = v5
    } else {
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = v4
    return ""
}

func main() {
    v1 := []uint32{419}
    var v2 int = 0
    var v3 int64 = 9223372036854775807
    if (18446744073709551615 == 4294967295) {
        v6 := GenS1{F0: 750, F1: 4294967295, F2: 1}
        v7 := GenS2{F0: ((false || (!(!(!false)))) && (v2 != -667))}
        _ = v6
        _ = v7
    }
    fmt.Println(v2, v2)
    var v8 float64 = float64(0.1)
    for fi9 := 0; fi9 < 12; fi9++ {
        if true {
            fmt.Println(v8, fi9)
            _ = len(v1)
        }
        v10 := GenS3{F0: (fi9 << ((4294967295) & 63)), F1: 0}
        if len(v1) > 0 {
            v1[0] = 2
        }
        _ = fi9
        _ = v10
    }
    v11 := GenS4{F0: v3, F1: 118}
    var v12 string = "a"
    v13 := GenS1{F0: -32768, F1: 18446744073709551615, F2: 2}
    switch v12 {
    case "a":
        if (!(false && false)) {
            v3 += v3
            v2 += (v2 & 76)
            v2 = (v2 % ((128) | 1))
        } else {
            fmt.Println(v2)
            v12 = v12
        }
    case "\n\t":
        v8 -= 3.14159
    case "foo bar":
        switch v2 {
        case -829:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int16 = 482
                v16 := GenS1{F0: v15, F1: 2, F2: 46}
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            switch ("a" + "foo bar") {
            case "\n\t":
                _ = len(v1)
            case "a":
                var v17 uint64 = 0
                _ = v17
            case "":
                switch -1 {
                case 128:
                    for fi18 := 0; fi18 < 16; fi18++ {
                        if (853 > 967) {
                            fmt.Println(v2)
                        }
                        _ = fi18
                    }
                case 2147483647:
                    if len(v1) > 0 {
                        v1[0] = 194
                    }
                }
            default:
                var v19 uint8 = 128
                _ = v19
            }
        }
    }
    v1 = append(v1, 42)
    switch (v12 + ("hello" + (v12 + ("hello" + v12)))) {
    case "":
        switch v12 {
        case "":
            for fi20 := 0; fi20 < 18; fi20++ {
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                _ = fi20
            }
        default:
            switch v12 {
            case "":
                if true {
                    switch ((v12 + (v12 + "\n\t")) + v12) {
                    case "":
                        switch "a" {
                        case "":
                            for fi21 := 0; fi21 < 17; fi21++ {
                                v12 = v12
                                _ = fi21
                            }
                        case "\n\t":
                            if (4294967295 != 2) {
                                var v22 int32 = 1
                                _ = v22
                            }
                        default:
                            if true {
                                v12 = v12
                            }
                        }
                    case "\n\t":
                        v3 -= 9223372036854775807
                    case "a":
                        v23 := GenS4{F0: v3, F1: -1}
                        _ = v23
                    default:
                        v24 := GenS1{F0: -654, F1: 1, F2: 42}
                        _ = v24
                    }
                    var v25 float32 = 1.5
                    _ = v25
                }
            case "hello":
                v2 += -201
            default:
                var v26 uint8 = 155
                _ = v26
            }
        }
    }
    v2 -= (-(v2))
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
}
