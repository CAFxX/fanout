package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    v3 := []string{"\n\t"}
    v4 := []int{32767, -9223372036854775808}
    var v5 int8 = 87
    var v6 uint8 = 187
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 bool = ((0.1 >= 1e10) || false)
        if (v6 < v6) {
            for fi9 := 0; fi9 < 9; fi9++ {
                v5 += v5
                switch (fi7 + (^(fi7))) {
                case 2147483647:
                    var v10 int32 = 0
                    _ = v10
                case -468:
                    var v11 string = ("foo bar" + ("\n\t" + (("foo bar" + "hello") + "a")))
                    _ = v11
                }
                v8 = v8
                _ = fi9
            }
            if v8 {
                switch fi7 {
                case 9223372036854775807:
                    v12 := GenS1{F0: 456, F1: false}
                    _ = v12
                case 32767:
                    v13 := GenS2{F0: 0, F1: 65535}
                    _ = v13
                case 127:
                    var v14 uint = 255
                    _ = v14
                }
                if v8 {
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v1[fi15] = fi15
                        if (!(96 == uint(v6))) {
                            v16 := GenS2{F0: -927, F1: 804}
                            var v17 float32 = 0.0
                            _ = v16
                            _ = v17
                        } else {
                            var v18 int64 = 78
                            v5 = v5
                            v5 += 127
                            if ((true || v8) || true) {
                                if (2 <= 4294967295) {
                                    var v19 float64 = (0.0 - (0.1 - 1e10))
                                    var v20 int = (fi15 * -620)
                                    _ = v19
                                    _ = v20
                                }
                                fmt.Println(v8)
                                v8 = v8
                                v8 = ((700 <= 255) && (int32(fi7) != 42))
                            } else {
                                var v21 int16 = 42
                                var v22 int64 = v18
                                fmt.Println(v21, v21)
                                v8 = (3.14159 < 0.1)
                                if (!(v21 > (v21 & 1))) {
                                    switch fi15 {
                                    case 128:
                                        v4 = append(v4, fi7)
                                    default:
                                        v5 = (^(v5))
                                    }
                                    var v23 int64 = -9223372036854775808
                                    _ = len(v3)
                                    _ = v23
                                } else {
                                    fmt.Println(fi15)
                                    v22 -= (v18 & -1)
                                    if (fi7 > fi15) {
                                        var v24 float32 = (1e10 - float32((-2.25 - -2.25)))
                                        _ = v24
                                    } else {
                                        v4 = append(v4, (fi15 - 4294967295))
                                        switch fi15 {
                                        case -2147483648:
                                            if ("hello" == "hello") {
                                                v8 = v8
                                                fmt.Println(fi15, fi7, v6)
                                                switch "hello" {
                                                case "hello":
                                                    fmt.Println(v21, v6)
                                                    v21 += 1
                                                    v25 := v1[fi7]
                                                    _ = v25
                                                case "a":
                                                    v26 := GenS1{F0: 42, F1: v8}
                                                    _ = v26
                                                }
                                            } else {
                                                fmt.Println(v8, v8)
                                            }
                                        case 1:
                                            switch "hello" {
                                            case "":
                                                _ = len(v4)
                                            case "a":
                                                if (v8 && v8) {
                                                    v18 += 42
                                                    var v27 uint8 = 233
                                                    var v28 uint32 = 370
                                                    _ = v27
                                                    _ = v28
                                                }
                                            case "\n\t":
                                                var v29 uint32 = 4294967295
                                                _ = v29
                                            default:
                                                v30 := GenS1{F0: 0, F1: v8}
                                                var v31 uint64 = 0
                                                var v32 float32 = 1.5
                                                var v33 uint64 = v31
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            }
                                        case -834:
                                            v34 := GenS3{F0: v5}
                                            _ = v34
                                        default:
                                            v35 := GenS2{F0: -73, F1: 65535}
                                            _ = v35
                                        }
                                    }
                                }
                                _ = v21
                                _ = v22
                            }
                            v36 := GenS4{F0: -121}
                            _ = v18
                            _ = v36
                        }
                        var v37 int64 = 42
                        v4 = append(v4, (^(fi7)))
                        if (!("foo bar" <= "hello")) {
                            var v38 bool = true
                            _ = v38
                        }
                        _ = fi15
                        _ = v37
                    }
                }
            } else {
                fmt.Println(v6, v8)
                _ = len(v2)
                v39 := GenS3{F0: v5}
                _ = v39
            }
            v6 = v6
        }
        _ = fi7
        _ = v8
    }
    if true {
        for fi40 := 0; fi40 < 14; fi40++ {
            if false {
                v41 := GenS5{F0: 73, F1: (v5 >> ((2) & 63)), F2: v6}
                for fi42 := 0; fi42 < 11; fi42++ {
                    var v43 string = "foo bar"
                    _ = fi42
                    _ = v43
                }
                _ = v41
            } else {
                for fi44 := 0; fi44 < 16; fi44++ {
                    delete(v2, 2147483648)
                    var v45 string = ""
                    _ = fi44
                    _ = v45
                }
            }
            _ = fi40
        }
        v6 = 1
        v46 := GenS4{F0: 32767}
        _ = v46
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v6, v6)
    case "a":
        v47 := GenS3{F0: (-(v5))}
        _ = v47
    default:
        v5 = (v5 / (((v5 << ((469) & 63))) | 1))
    }
    v6 -= v6
    if len(v3) > 0 {
        v3[0] = ""
    }
    if ((3.14159 + (float32((0.1 - 3.14159)) - -2.25)) <= ((float32(3.14159) * -2.25) * 0.0)) {
        var v48 int32 = 24
        v49 := GenS4{F0: 1}
        _ = v48
        _ = v49
    }
    var v50 int8 = (v5 % ((v5) | 1))
    v5 -= (v5 | v50)
    for fi51 := 0; fi51 < 9; fi51++ {
        fmt.Println(fi51, fi51, v6)
        if false {
            var v52 int = fi51
            if true {
                v50 = -29
                var v53 uint = 472
                _ = v53
            }
            _ = v52
        }
        var v54 uint32 = 2147483647
        _ = fi51
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
