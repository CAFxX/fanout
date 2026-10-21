package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 string; F1 uint64; F2 int64 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS7 struct { F0 string; F1 uint8 }
func genFn0() uint64 {
    v7 := []int{}
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint8 = 127
            v11 := GenS1{F0: 175, F1: "hello", F2: v10}
            if (0.1 < float64((-2.25 / (3.14159 / float32((0.1 * (1.5 - 0.1))))))) {
                v10 = (v10 % ((0) | 1))
                if ((15 != 32768) && false) {
                    var v12 int32 = 1
                    _ = v12
                } else {
                    var v13 int32 = 642
                    _ = v13
                }
                v10 = 0
            } else {
                switch fi9 {
                case 127:
                    var v14 uint8 = (v10 ^ (^(v10)))
                    _ = v14
                case 971:
                    var v15 bool = ((false || true) || (169 > (^(v10))))
                    var v16 uint = 255
                    v17 := GenS1{F0: 988, F1: "hello", F2: 0}
                    v18 := GenS1{F0: -1, F1: "hello", F2: v10}
                    if (1 != 0) {
                        _ = len(v7)
                        if len(v7) > 0 {
                            v7[0] = fi9
                        }
                        var v19 int8 = -1
                        var v20 int32 = -2147483648
                        switch "a" {
                        case "hello":
                            var v21 float32 = (((1e10 - 1e10) / 1e10) / -2.25)
                            _ = v21
                        }
                        _ = v19
                        _ = v20
                    } else {
                        if false {
                            var v22 uint16 = 255
                            _ = v22
                        }
                        v10 = (^(v10))
                    }
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                case 128:
                    v10 = (^(v10))
                default:
                    var v23 uint = 18446744073709551615
                    _ = v23
                }
                switch (fi8 | fi8) {
                case 32767:
                    if true {
                        v7 = append(v7, -645)
                        switch (fi9 >> ((614) & 63)) {
                        case 0:
                            v24 := GenS1{F0: 1, F1: ("hello" + ("hello" + "foo bar")), F2: (v10 << ((778) & 63))}
                            _ = v24
                        }
                        if (!true) {
                            v10 += (^(v10))
                            v25 := GenS1{F0: 807, F1: ("foo bar" + "hello"), F2: v10}
                            v10 = v10
                            _ = v25
                        } else {
                            v10 = 2
                            v26 := GenS1{F0: 9223372036854775807, F1: (("hello" + "a") + "foo bar"), F2: (v10 * v10)}
                            var v27 uint64 = 103
                            _ = v26
                            _ = v27
                        }
                    }
                case 42:
                    v28 := GenS2{F0: 32767, F1: "\n\t"}
                    _ = v28
                }
                v10 = 127
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v29 float64 = 1.5
            fmt.Println(fi8, fi9, v10)
            _ = fi9
            _ = v10
            _ = v11
            _ = v29
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    var v30 uint64 = 42
    if true {
        v30 = (v30 >> ((4294967295) & 63))
    } else {
        v31 := GenS3{F0: "\n\t", F1: (v30 << ((0) & 63)), F2: 1}
        _ = v31
    }
    _ = v30
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 42
    var v3 bool = ((128 >= 1) || true)
    var v4 float64 = 3.14159
    var v5 int8 = -128
    var v6 int8 = v5
    switch "a" {
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            v6 = v5
        }
    case "a":
        var v32 int8 = v5
        _ = v32
    }
    v33 := GenS4{F0: v2, F1: genFn0(), F2: 127}
    if v3 {
        var v34 float32 = (float32(float64(1e10)) + -2.25)
        v5 += v6
        _ = v34
    }
    v35 := []int{-2147483648}
    fmt.Println(v6)
    v36 := GenS5{F0: 151}
    if (v2 == v2) {
        v37 := GenS5{F0: 127}
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 int16 = 32767
            v6 -= v6
            _ = fi38
            _ = v39
        }
        v40 := GenS5{F0: 127}
        _ = v37
        _ = v40
    } else {
        fmt.Println(v4, v4, v2)
    }
    v41 := GenS1{F0: (v2 * v2), F1: ("a" + "a"), F2: 255}
    v42 := GenS6{F0: 0, F1: 0, F2: 65535}
    v43 := v1["\n\t"]
    for fi44 := 0; fi44 < 13; fi44++ {
        for fi45 := 0; fi45 < 12; fi45++ {
            var v46 uint8 = 201
            _ = fi45
            _ = v46
        }
        var v47 int64 = v2
        var v48 uint = 4294967295
        _ = fi44
        _ = v47
        _ = v48
    }
    if (-2.25 <= 3.14159) {
        switch ("hello" + "\n\t") {
        case "":
            for fi49 := 0; fi49 < 15; fi49++ {
                v35 = append(v35, 844)
                _ = fi49
            }
        default:
            v2 = (v2 ^ v2)
        }
        var v50 uint32 = 42
        _ = v50
    }
    v51 := GenS7{F0: "hello", F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v51
    _ = v6
}
