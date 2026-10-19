package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int; F2 uint }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int8; F1 uint }
func genFn0(p0 uint8, p1 int, p2 int32) int8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        p2 = (-(p2))
        var v8 uint32 = 2147483647
        v9 := GenS1{F0: 1, F1: (v8 >> ((0) & 63)), F2: (p2 * (p2 << ((2) & 63)))}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{}
    if true {
        p2 = (p2 * 0)
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            if true {
                var v12 float32 = 3.14159
                _ = v12
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                if false {
                    var v14 int16 = -808
                    _ = v14
                } else {
                    switch ("\n\t" + ("hello" + "hello")) {
                    case "a":
                        if len(v10) > 0 {
                            v10[0] = (p1 ^ (fi11 * (p1 - fi11)))
                        }
                    case "":
                        var v15 string = "\n\t"
                        _ = v15
                    case "\n\t":
                        var v16 float64 = 0.0
                        if true {
                            var v17 int8 = -128
                            fmt.Println(p0, v17, v16)
                            _ = v17
                        }
                        fmt.Println(fi13)
                        switch ("hello" + "hello") {
                        case "":
                            p1 = -26
                        case "hello":
                            v18 := GenS2{F0: 32767, F1: -37, F2: 18446744073709551615}
                            _ = v18
                        }
                        v10 = append(v10, (fi13 % (((fi13 & p1)) | 1)))
                        _ = v16
                    default:
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    }
                }
                var v19 int8 = 42
                fmt.Println(fi13)
                p1 -= fi13
                if (1 != -9223372036854775808) {
                    p1 -= -472
                    p1 -= fi13
                    v19 = (v19 & v19)
                    switch p1 {
                    case -601:
                        v20 := GenS1{F0: (v19 + (v19 + 127)), F1: 442, F2: p2}
                        _ = v20
                    }
                    if len(v10) > 0 {
                        v10[0] = (p1 % (((fi11 % ((p1) | 1))) | 1))
                    }
                    v19 += v19
                    var v21 float32 = 1.5
                    _ = v21
                } else {
                    switch (fi11 - 128) {
                    case 32768:
                        var v22 int64 = 419
                        _ = v22
                    case -326:
                        if (1000 < 42) {
                            var v23 int16 = -899
                            _ = v23
                        }
                    }
                    v24 := GenS1{F0: (v19 % ((-32) | 1)), F1: 1, F2: (p2 << ((4294967295) & 63))}
                    _ = v24
                }
                var v25 uint = 727
                _ = fi13
                _ = v19
                _ = v25
            }
            _ = fi11
        }
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}
func genFn1(p0 string) bool {
    switch -661 {
    case 0:
        var v26 bool = (0 <= 0)
        _ = v26
    case -409:
        v27 := GenS2{F0: 811, F1: -556, F2: 0}
        _ = v27
    case -481:
        fmt.Println(p0)
    default:
        v28 := []int{32768, 514, 606, 42}
        _ = v28
    }
    p0 = ""
    _ = p0
    return false
}

func main() {
    var v1 uint = 198
    v2 := []uint32{}
    v3 := make(map[int]int)
    v4 := []int{32768, 42}
    var v5 float32 = 0.0
    v6 := make(map[int]uint32)
    for fi29 := 0; fi29 < 3; fi29++ {
        v30 := GenS1{F0: -1, F1: 354, F2: -1}
        v5 = float32(1.5)
        _ = fi29
        _ = v30
    }
    var v31 int16 = 650
    v32 := GenS3{F0: false, F1: -1}
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        switch "\n\t" {
        case "a":
            if true {
                switch (("foo bar" + "hello") + "a") {
                case "a":
                    fmt.Println(v5)
                case "hello":
                    var v34 int16 = v31
                    _ = v34
                default:
                    if len(v4) > 0 {
                        v4[0] = fi33
                    }
                }
                v35 := GenS4{F0: 127, F1: v1}
                v1 = (v1 | (^(v1)))
                _ = v35
            } else {
                fmt.Println(fi33)
            }
        case "foo bar":
            if ((v31 > v31) || false) {
                for fi36 := 0; fi36 < 16; fi36++ {
                    var v37 uint32 = 2147483648
                    _ = fi36
                    _ = v37
                }
            } else {
                var v38 float64 = -2.25
                v5 = (v5 + ((1.5 / 0.1) * (1.5 - v5)))
                _ = v38
            }
        default:
            for fi39 := 0; fi39 < 14; fi39++ {
                var v40 int64 = 19
                _ = fi39
                _ = v40
            }
        }
        _ = fi33
    }
    switch 1 {
    case 9223372036854775807:
        if false {
            var v41 uint64 = 19
            v1 -= (v1 & 4294967295)
            v1 = uint(v41)
            _ = v41
        }
    case 0:
        v42 := GenS4{F0: 100, F1: 65535}
        _ = v42
    }
    v5 = 0.0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
