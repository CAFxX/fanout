package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 uint64; F2 int64 }

func main() {
    var v1 uint = 461
    v2 := make(map[int]string)
    var v3 string = ("foo bar" + ((("hello" + "\n\t") + ("foo bar" + (("" + "hello") + "foo bar"))) + ""))
    var v4 uint = v1
    v5 := make(map[string]int)
    if (!(946 > 81)) {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: int16(fi6), F1: 0}
            _ = fi6
            _ = v7
        }
        v3 = ("foo bar" + ("" + ((v3 + v3) + v3)))
    } else {
        v3 = (v3 + v3)
    }
    var v8 int32 = 42
    var v9 float32 = (3.14159 / (3.14159 + -2.25))
    v10 := []int{-239, 4294967295, 128}
    var v11 uint64 = 0
    if (!((true || (true && (65535 > 0))) || false)) {
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := GenS2{F0: uint64(v1), F1: 1}
            if (false || false) {
                var v14 uint32 = 2147483647
                if false {
                    if true {
                        _ = len(v10)
                        v11 -= v11
                        v15 := GenS1{F0: 42, F1: 4294967295}
                        _ = v15
                    } else {
                        v10 = append(v10, fi12)
                    }
                }
                if (true && false) {
                    v1 += v1
                    switch "foo bar" {
                    case "":
                        if false {
                            for fi16 := 0; fi16 < 15; fi16++ {
                                var v17 uint16 = 65535
                                v14 = (v14 - v14)
                                switch ("\n\t" + "foo bar") {
                                case "hello":
                                    var v18 int8 = -128
                                    _ = v18
                                }
                                _ = fi16
                                _ = v17
                            }
                            v3 = v3
                            var v19 uint16 = 449
                            _ = v19
                        }
                    default:
                        if false {
                            v20 := v5[v3]
                            v21 := GenS3{F0: 892}
                            v22 := GenS1{F0: 258, F1: (^(v1))}
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    }
                    var v23 int16 = 0
                    _ = v23
                } else {
                    v11 = (^(v11))
                }
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        v9 += 1.5
        v9 = v9
    }
    if (!false) {
        fmt.Println(v11)
        v8 -= v8
        var v24 string = ""
        _ = v24
    }
    var v25 string = (((v3 + (v3 + "hello")) + "\n\t") + "hello")
    if (v3 != v3) {
        v26 := GenS3{F0: 2}
        switch "a" {
        case "hello":
            var v27 uint8 = 127
            _ = v27
        case "\n\t":
            v28 := v5[v25]
            _ = v28
        case "a":
            fmt.Println(v1, v25)
        default:
            v1 = v4
        }
        _ = v26
    }
    v29 := GenS2{F0: (v11 | uint64(v1)), F1: 22}
    var v30 uint32 = 750
    var v31 float32 = (v9 - 1e10)
    if false {
        if true {
            if false {
                v4 = (v4 & (v4 << (((v1 - (v4 - v4))) & 63)))
            } else {
                var v32 uint32 = v30
                fmt.Println(v11, v31, v9)
                _ = v32
            }
            v4 += v1
            v33 := GenS4{F0: (v8 ^ v8)}
            _ = v33
        } else {
            if true {
                v34 := GenS2{F0: 935, F1: 5}
                var v35 uint8 = uint8(v4)
                _ = v34
                _ = v35
            } else {
                var v36 float64 = 0.0
                v37 := GenS3{F0: 297}
                _ = v36
                _ = v37
            }
            fmt.Println(v30, v9, v25)
            v1 = v4
        }
        v38 := GenS4{F0: 621}
        _ = v38
    }
    v30 = (^(v30))
    v10 = append(v10, 9223372036854775807)
    fmt.Println(v3, v3)
    v39 := GenS5{F0: 2, F1: (^(v11)), F2: -576}
    v5[(((v3 + v3) + "") + v3)] = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
