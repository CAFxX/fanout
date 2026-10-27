package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int; F2 string }
type GenS5 struct { F0 int8; F1 int16; F2 uint32 }
func genFn0() uint {
    if false {
        var v5 uint32 = 832
        fmt.Println(v5)
        _ = v5
    }
    v6 := []int{32768, 1, 4294967295, 128}
    if (!(42 <= -925)) {
        var v7 float32 = 0.0
        var v8 float32 = 0.1
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        _ = len(v6)
        if false {
            if len(v6) > 0 {
                v6[0] = (fi9 * -702)
            }
            var v10 uint64 = 0
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := GenS1{F0: (("\n\t" + ("" + "\n\t")) + "hello"), F1: 0, F2: 0}
                var v13 int8 = 13
                var v14 int = 127
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = v10
        }
        _ = fi9
    }
    _ = v6
    return 255
}
func genFn1() int16 {
    switch "hello" {
    case "hello":
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 float64 = (3.14159 * 0.1)
            _ = fi15
            _ = v16
        }
    }
    var v17 bool = (102 <= -2)
    _ = v17
    return -572
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -32768
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 0
    var v18 int8 = -125
    v19 := GenS2{F0: 1, F1: (v2 * -32768), F2: v18}
    switch -9223372036854775808 {
    case 127:
        for fi20 := 0; fi20 < 16; fi20++ {
            switch "hello" {
            case "\n\t":
                v2 += v2
            case "foo bar":
                v18 -= (v18 | (v18 << ((65535) & 63)))
            default:
                v4 += v4
            }
            var v21 float32 = (1.5 + 3.14159)
            _ = fi20
            _ = v21
        }
    case -57:
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v18)
        default:
            if ((!true) || false) {
                for fi22 := 0; fi22 < 16; fi22++ {
                    if (0.0 == -2.25) {
                        v23 := GenS1{F0: "hello", F1: -692, F2: (v4 % (((v4 - 0)) | 1))}
                        _ = v23
                    }
                    _ = fi22
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        v26 := GenS3{F0: (v4 / ((42) | 1))}
                        if (42 == 0) {
                            var v27 uint8 = 128
                            _ = v27
                        }
                        var v28 uint16 = v4
                        _ = fi25
                        _ = v26
                        _ = v28
                    }
                    _ = fi24
                }
                v29 := GenS4{F0: genFn1(), F1: 127, F2: "hello"}
                _ = v29
            }
        }
    case -244:
        var v30 uint16 = 65535
        _ = v30
    default:
        v2 = (-(v2))
    }
    v31 := v3[483]
    fmt.Println(v31)
    var v32 int32 = 2147483647
    v4 = 562
    for fi33 := 0; fi33 < 2; fi33++ {
        v34 := []int{(fi33 / ((fi33) | 1))}
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 uint64 = 18446744073709551615
            if true {
                v37 := GenS2{F0: -1, F1: (-(v2)), F2: (v18 + -128)}
                var v38 bool = true
                _ = v37
                _ = v38
            } else {
                if len(v34) > 0 {
                    v34[0] = 32768
                }
                v39 := GenS1{F0: (("\n\t" + ("foo bar" + ("a" + (("\n\t" + "\n\t") + (("" + "foo bar") + "\n\t"))))) + ""), F1: v32, F2: v4}
                _ = v39
            }
            _ = fi35
            _ = v36
        }
        if (((-(v2)) <= (v2 & v2)) && true) {
            for fi40 := 0; fi40 < 14; fi40++ {
                v41 := GenS1{F0: "foo bar", F1: (^(v32)), F2: v4}
                v42 := GenS5{F0: (v18 << ((525) & 63)), F1: -32768, F2: (^(v31))}
                _ = fi40
                _ = v41
                _ = v42
            }
            fmt.Println(fi33, v18, v32)
        }
        _ = fi33
        _ = v34
    }
    for fi43 := 0; fi43 < 18; fi43++ {
        var v44 int = fi43
        _ = fi43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
