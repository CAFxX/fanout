package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int16; F2 int16 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 int; F2 int64 }

func main() {
    var v1 uint64 = 42
    v2 := []string{"", "", "a"}
    v3 := make(map[int]uint32)
    var v4 int = -2147483648
    v5 := []uint32{}
    v6 := make(map[uint32]int)
    fmt.Println(v1, v4)
    delete(v6, 2)
    fmt.Println(v4, v1)
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            if (!false) {
                fmt.Println(fi7, fi7)
                var v9 bool = true
                switch -2147483648 {
                case 876:
                    if v9 {
                        var v10 string = "hello"
                        v1 += v1
                        fmt.Println(v10, v1)
                        _ = v10
                    } else {
                        v11 := GenS1{F0: -128, F1: 128, F2: 4294967295}
                        _ = v11
                    }
                default:
                    var v12 uint32 = 65535
                    _ = v12
                }
                if ((v1 < v1) && ("" >= "a")) {
                    if ((!((^(v1)) >= v1)) || true) {
                        v13 := GenS1{F0: 100, F1: (v4 & (fi8 / (((-(v4))) | 1))), F2: 42}
                        v9 = true
                        var v14 uint64 = 0
                        if (127 > 128) {
                            v15 := GenS1{F0: 127, F1: (v4 * (v4 * v4)), F2: 65535}
                            v16 := GenS1{F0: 127, F1: (v4 + fi8), F2: 430}
                            _ = v15
                            _ = v16
                        }
                        fmt.Println(v9, v9, v4)
                        _ = v13
                        _ = v14
                    }
                    if (!("a" <= ("hello" + ("\n\t" + "a")))) {
                        fmt.Println(v4, v9, fi7)
                        switch ((("" + "foo bar") + "\n\t") + "hello") {
                        case "":
                            v9 = v9
                        default:
                            v17 := GenS2{F0: (v1 ^ 18446744073709551615), F1: 4294967295}
                            _ = v17
                        }
                        fmt.Println(fi7, fi7, v9)
                        v18 := GenS3{F0: 0, F1: 1000, F2: 1000}
                        _ = v18
                    }
                    var v19 uint16 = 1
                    _ = v19
                } else {
                    if false {
                        if v9 {
                            var v20 int16 = 1000
                            _ = v20
                        }
                        var v21 uint16 = 1
                        fmt.Println(v21, fi8, v4)
                        v9 = ((false || true) && false)
                        v22 := GenS4{F0: v21, F1: false}
                        _ = v21
                        _ = v22
                    }
                    v1 = (v1 | (v1 >> ((4294967295) & 63)))
                    switch "hello" {
                    case "foo bar":
                        v23 := GenS4{F0: 2, F1: true}
                        _ = v23
                    case "":
                        v24 := GenS3{F0: 979, F1: 42, F2: 532}
                        _ = v24
                    }
                }
                v2 = append(v2, "")
                var v25 int32 = 994
                if v9 {
                    var v26 uint32 = 42
                    switch int(v26) {
                    case -9223372036854775808:
                        if v9 {
                            v27 := GenS5{F0: 65535}
                            _ = v27
                        } else {
                            v25 = (v25 - (v25 & 2147483647))
                            _ = len(v2)
                            switch fi7 {
                            case 127:
                                var v28 float64 = 0.0
                                _ = v28
                            default:
                                var v29 int16 = 958
                                _ = v29
                            }
                        }
                    case 127:
                        switch ("hello" + ("\n\t" + "hello")) {
                        case "hello":
                            v30 := GenS2{F0: uint64(v26), F1: 182}
                            _ = v30
                        case "\n\t":
                            switch (fi7 | (^(fi7))) {
                            case -253:
                                var v31 uint32 = v26
                                _ = v31
                            }
                        }
                        v9 = ((v9 || (65535 != 255)) && ((((!(1e10 <= 0.1)) && v9) && ((fi7 * v4) != (-(v4)))) || (-300 == -9223372036854775808)))
                        v32 := GenS3{F0: v26, F1: -868, F2: 1}
                        fmt.Println(v9, v25, v4)
                        _ = v32
                    }
                    var v33 uint8 = 199
                    _ = v26
                    _ = v33
                }
                _ = v25
                _ = v9
            }
            _ = fi8
        }
        var v34 int64 = 1
        fmt.Println(v34, v1)
        _ = fi7
        _ = v34
    }
    if (98 > 255) {
        var v35 float64 = 3.14159
        _ = v35
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v36 := GenS6{F0: v1, F1: (v4 ^ 42), F2: -596}
    if (!true) {
        for fi37 := 0; fi37 < 3; fi37++ {
            var v38 int8 = 0
            _ = fi37
            _ = v38
        }
        var v39 uint32 = 979
        for fi40 := 0; fi40 < 19; fi40++ {
            v6[v39] = (v4 << ((739) & 63))
            _ = fi40
        }
        _ = v39
    }
    v41 := v3[(-(v4))]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
