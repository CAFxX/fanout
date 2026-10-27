package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int; F2 uint }
func genFn0(p0 uint8, p1 int16) int8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint16 = 1
        if (0 <= p1) {
            p0 = p0
            p1 -= (p1 / (((^(p1))) | 1))
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                fmt.Println(p0, p0)
                p0 -= (p0 << ((890) & 63))
                v9 := []int{32767}
                _ = fi8
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    if true {
        switch "a" {
        case "a":
            for fi10 := 0; fi10 < 3; fi10++ {
                switch -9223372036854775808 {
                case -448:
                    if true {
                        v11 := []int{fi10, 784, (fi10 & 32768)}
                        _ = v11
                    }
                case -944:
                    var v12 int64 = 1
                    _ = v12
                }
                var v13 int8 = 100
                var v14 bool = ((!false) || ((!true) && (("" + "\n\t") <= "\n\t")))
                _ = fi10
                _ = v13
                _ = v14
            }
        case "foo bar":
            v15 := GenS1{F0: 898, F1: -2147483648, F2: 9223372036854775807}
            _ = v15
        }
    }
    p0 -= p0
    _ = p0
    _ = p1
    return 127
}
func genFn1() uint8 {
    var v16 uint64 = 0
    var v17 int = 1
    if (((428 < 24) && false) && ((true && false) && true)) {
        v16 = 489
        v18 := GenS1{F0: (v16 & v16), F1: -2147483648, F2: (v17 % ((42) | 1))}
        _ = v18
    } else {
        switch "\n\t" {
        case "a":
            var v19 int64 = -66
            _ = v19
        case "hello":
            v20 := GenS2{F0: -1, F1: 1}
            _ = v20
        case "":
            v21 := GenS1{F0: v16, F1: 2147483647, F2: 757}
            _ = v21
        default:
            if (42 == -1) {
                v22 := []int{v17, -1}
                if (9223372036854775807 <= 9223372036854775807) {
                    var v23 uint = 18446744073709551615
                    v17 -= (-(v17))
                    _ = v23
                }
                switch "\n\t" {
                case "foo bar":
                    _ = len(v22)
                case "\n\t":
                    switch ("a" + "a") {
                    case "a":
                        var v24 uint64 = v16
                        _ = v24
                    case "\n\t":
                        fmt.Println(v16, v17)
                    }
                default:
                    for fi25 := 0; fi25 < 19; fi25++ {
                        v26 := GenS3{F0: 0, F1: ((^(fi25)) >= (fi25 & 9223372036854775807)), F2: 255}
                        for fi27 := 0; fi27 < 12; fi27++ {
                            fmt.Println(fi25)
                            var v28 float32 = -2.25
                            _ = fi27
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                    }
                }
                _ = v22
            } else {
                v29 := []int{(^(v17)), (v17 >> ((18446744073709551615) & 63)), v17}
                v30 := GenS4{F0: -1}
                _ = v29
                _ = v30
            }
        }
    }
    v31 := GenS2{F0: -32768, F1: int32(v17)}
    _ = v16
    _ = v17
    _ = v31
    return 127
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{9223372036854775807, 538, 4294967295, 2147483647}
    var v3 int8 = 1
    var v4 int32 = 2147483647
    var v5 uint64 = 42
    for fi32 := 0; fi32 < 12; fi32++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (true || ((0.1 + 1.5) > 0.0)) {
            v33 := GenS5{F0: 0, F1: (fi32 << ((704) & 63)), F2: uint(v5)}
            _ = v33
        } else {
            for fi34 := 0; fi34 < 7; fi34++ {
                v3 = (v3 + v3)
                fmt.Println(v3)
                v35 := v1["\n\t"]
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 3; fi36++ {
                if (!true) {
                    v5 -= (v5 | (^(v5)))
                } else {
                    v4 = (v4 & 42)
                }
                if (18446744073709551615 != 4294967295) {
                    var v37 int8 = (v3 & 12)
                    var v38 uint16 = 32767
                    _ = v37
                    _ = v38
                } else {
                    v3 = genFn0(2, -279)
                    var v39 int32 = v4
                    var v40 string = "\n\t"
                    _ = v39
                    _ = v40
                }
                v4 = v4
                switch ("" + "a") {
                case "hello":
                    var v41 uint64 = v5
                    _ = v41
                case "":
                    var v42 float32 = 3.14159
                    _ = v42
                default:
                    v4 -= v4
                }
                _ = fi36
            }
        }
        _ = fi32
    }
    if len(v2) > 0 {
        v2[0] = int(v5)
    }
    fmt.Println(v5, v5)
    fmt.Println(v3, v5, v3)
    v5 -= v5
    v43 := GenS4{F0: -9223372036854775808}
    var v44 bool = false
    v45 := GenS1{F0: 42, F1: v4, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
