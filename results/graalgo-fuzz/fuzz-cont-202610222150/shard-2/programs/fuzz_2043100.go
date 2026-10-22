package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 string; F1 int; F2 int64 }
type GenS3 struct { F0 uint8; F1 int32; F2 int64 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint32; F1 int32; F2 int32 }
func genFn0() uint16 {
    var v4 int64 = 1
    if (3.14159 < (1e10 - ((0.1 + 1e10) / 3.14159))) {
        v5 := GenS1{F0: -1, F1: v4, F2: 2}
        var v6 bool = true
        v7 := []int{}
        _ = v5
        _ = v6
        _ = v7
    }
    switch "" {
    case "":
        var v8 int16 = 686
        _ = v8
    default:
        fmt.Println(v4)
    }
    fmt.Println(v4)
    if (true && (42 > 142)) {
        for fi9 := 0; fi9 < 1; fi9++ {
            v4 += v4
            switch "foo bar" {
            case "hello":
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := GenS1{F0: -42, F1: (v4 / ((v4) | 1)), F2: 1}
                    _ = fi10
                    _ = v11
                }
            case "foo bar":
                v12 := []int{18, (^(fi9)), 127}
                _ = v12
            default:
                var v13 uint = 4294967295
                _ = v13
            }
            v4 -= v4
            _ = fi9
        }
        switch "" {
        case "a":
            for fi14 := 0; fi14 < 19; fi14++ {
                v4 += int64(fi14)
                for fi15 := 0; fi15 < 12; fi15++ {
                    switch ("hello" + "") {
                    case "a":
                        var v16 int64 = -9223372036854775808
                        _ = v16
                    case "":
                        v4 -= (v4 + (^(v4)))
                        var v17 string = "\n\t"
                        v4 = v4
                        _ = v17
                    case "hello":
                        v4 -= v4
                    }
                    if (32767 <= fi15) {
                        v18 := []int{(fi15 >> ((0) & 63)), -1, 0}
                        var v19 uint = 42
                        if true {
                            var v20 uint = 65535
                            _ = v20
                        }
                        _ = v18
                        _ = v19
                    } else {
                        v21 := GenS1{F0: -128, F1: v4, F2: 2147483647}
                        _ = v21
                    }
                    var v22 uint8 = 128
                    _ = fi15
                    _ = v22
                }
                _ = fi14
            }
        case "foo bar":
            for fi23 := 0; fi23 < 10; fi23++ {
                for fi24 := 0; fi24 < 6; fi24++ {
                    v4 = v4
                    v4 = v4
                    v4 = (v4 ^ v4)
                    _ = fi24
                }
                var v25 int32 = -2147483648
                _ = fi23
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 uint32 = 2
            v4 = 689
            _ = fi26
            _ = v27
        }
    }
    _ = v4
    return 778
}
func genFn1(p0 int64, p1 int) uint64 {
    p0 -= p0
    v28 := GenS2{F0: "a", F1: (p1 * (p1 / (((p1 % ((p1) | 1))) | 1))), F2: -244}
    v29 := GenS3{F0: 42, F1: 2147483647, F2: p0}
    v30 := GenS3{F0: 0, F1: 573, F2: -1}
    _ = p0
    _ = p1
    _ = v28
    _ = v29
    _ = v30
    return 4294967295
}

func main() {
    var v1 uint64 = 323
    var v2 uint16 = 89
    var v3 int16 = 865
    if false {
        var v31 uint = 843
        _ = v31
    } else {
        var v32 bool = (!(76 <= 100))
        v33 := []int{-9223372036854775808, 9223372036854775807}
        _ = v32
        _ = v33
    }
    if true {
        switch "hello" {
        case "\n\t":
            for fi34 := 0; fi34 < 15; fi34++ {
                fmt.Println(v1)
                var v35 int32 = 0
                _ = fi34
                _ = v35
            }
        case "hello":
            v3 = v3
        }
    } else {
        v2 = (^(v2))
        var v36 int32 = -1
        v2 -= v2
        _ = v36
    }
    fmt.Println(v3, v2)
    var v37 uint16 = v2
    if (true || true) {
        if false {
            switch 2147483647 {
            case -1:
                fmt.Println(v1, v2)
            case 1:
                if false {
                    v2 = (v2 & (v2 << ((807) & 63)))
                    v38 := []int{655}
                    _ = v38
                } else {
                    var v39 uint16 = 889
                    fmt.Println(v2)
                    _ = v39
                }
            default:
                if (true || ((true && false) || false)) {
                    var v40 int32 = 1
                    v1 -= genFn1(-455, 42)
                    _ = v40
                }
            }
            v41 := GenS4{F0: (-128 > -128), F1: 744}
            fmt.Println(v3, v1, v37)
            _ = v41
        }
        v1 += (v1 - (v1 << ((18446744073709551615) & 63)))
    }
    v42 := GenS2{F0: "a", F1: 2147483647, F2: 1}
    v43 := GenS5{F0: 2147483647, F1: -675, F2: 0}
    v1 -= 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v42
    _ = v43
}
