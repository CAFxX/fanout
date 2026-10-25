package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int }
type GenS2 struct { F0 string; F1 uint; F2 uint }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint16 = 898
    v2 := make(map[int]int)
    v3 := []uint32{0, 65535, 2147483648, 4294967295}
    var v4 string = "foo bar"
    var v5 bool = (((0.1 + (0.1 - (-2.25 * 1.5))) != -2.25) || (!true))
    if len(v3) > 0 {
        v3[0] = 1
    }
    v5 = (!((v5 && (v5 && v5)) || (-1 < 1)))
    v6 := GenS1{F0: 287, F1: 729, F2: 369}
    var v7 uint16 = v1
    for fi8 := 0; fi8 < 13; fi8++ {
        switch ("" + v4) {
        case "":
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 bool = v5
                if v10 {
                    switch fi8 {
                    case -499:
                        var v11 bool = v5
                        v12 := v2[fi9]
                        _ = v11
                        _ = v12
                    case 269:
                        var v13 int64 = 42
                        _ = v13
                    case 0:
                        var v14 int16 = -1
                        _ = v14
                    default:
                        switch -82 {
                        case 1:
                            if v5 {
                                v10 = true
                                v10 = (v10 && (v4 != ("a" + "foo bar")))
                            } else {
                                fmt.Println(v4, v7)
                                var v15 int16 = 368
                                _ = v15
                            }
                        case 128:
                            v3 = append(v3, 42)
                        case 9223372036854775807:
                            v4 = ((v4 + "") + v4)
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                } else {
                    switch (v4 + (v4 + "a")) {
                    case "hello":
                        var v16 bool = (fi9 < fi9)
                        _ = v16
                    case "\n\t":
                        v17 := GenS1{F0: 0, F1: 764, F2: (fi9 + (fi9 & fi8))}
                        _ = v17
                    case "a":
                        v4 = ("hello" + (v4 + v4))
                    }
                    v7 += v1
                }
                var v18 uint64 = 1
                _ = fi9
                _ = v10
                _ = v18
            }
        case "a":
            var v19 int64 = -763
            _ = v19
        }
        _ = len(v2)
        _ = fi8
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS2{F0: "", F1: 844, F2: 4294967295}
        v22 := GenS3{F0: 0, F1: v4}
        _ = fi20
        _ = v21
        _ = v22
    }
    switch 42 {
    case -9223372036854775808:
        v23 := GenS4{F0: (v7 * (v1 & 66))}
        _ = v23
    case 467:
        v24 := GenS2{F0: v4, F1: 1, F2: 1}
        _ = v24
    }
    switch 619 {
    case 127:
        var v25 bool = v5
        _ = v25
    }
    switch v4 {
    case "foo bar":
        var v26 uint = 708
        _ = v26
    }
    if (v4 > "\n\t") {
        if (!(v5 || v5)) {
            var v27 int8 = 1
            _ = v27
        } else {
            var v28 int16 = 42
            v7 = v1
            _ = v28
        }
    }
    var v29 uint64 = 979
    v2[-9223372036854775808] = -9223372036854775808
    v30 := GenS2{F0: v4, F1: 65535, F2: 65535}
    if (1 == 819) {
        var v31 uint64 = (^(v29))
        var v32 int64 = 1
        v1 -= 42
        _ = v31
        _ = v32
    }
    v4 = ""
    var v33 int16 = -936
    v4 = (v4 + v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
