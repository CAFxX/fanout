package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int16 }
func genFn0() uint32 {
    var v7 uint16 = 195
    v8 := []int{}
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 42, F1: 100, F2: 2147483647}
        var v11 float64 = float64(0.0)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 bool = false
    _ = v12
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    var v1 int = 814
    var v2 string = "\n\t"
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    v5 := []string{((("" + v2) + "foo bar") + v2), (v2 + v2), (v2 + "foo bar"), "hello"}
    var v6 int8 = -1
    var v13 float64 = 3.14159
    fmt.Println(v13)
    var v14 uint64 = 4294967295
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    v15 := GenS1{F0: 149, F1: 100, F2: 59}
    for fi16 := 0; fi16 < 20; fi16++ {
        v1 += 4294967295
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            switch v2 {
            case "foo bar":
                v13 = v13
            case "":
                v5 = append(v5, (((v2 + v2) + ((v2 + "\n\t") + v2)) + "a"))
            case "\n\t":
                var v18 uint32 = 2147483648
                _ = v18
            default:
                var v19 uint32 = genFn0()
                _ = v19
            }
            _ = fi17
        }
        _ = fi16
    }
    fmt.Println(v2, v1, v14)
    if true {
        switch v2 {
        case "\n\t":
            _ = len(v5)
        case "":
            var v20 int8 = (v6 + 1)
            _ = v20
        case "hello":
            var v21 uint64 = (v14 | v14)
            _ = v21
        }
        v22 := GenS1{F0: 65535, F1: (v6 >> ((772) & 63)), F2: 42}
        fmt.Println(v6)
        _ = v22
    }
    if false {
        v4[945] = (v1 | (v1 / ((v1) | 1)))
        var v23 float32 = (((0.0 + 3.14159) - float32(v13)) - (((-2.25 / (3.14159 / 0.1)) * (-2.25 * 0.0)) + -2.25))
        if (32767 != 14) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 int64 = -9223372036854775808
                _ = fi24
                _ = v25
            }
        }
        _ = v23
    } else {
        var v26 int16 = 42
        _ = v26
    }
    switch v1 {
    case -820:
        fmt.Println(v14, v6, v13)
    }
    if true {
        v2 = (v2 + (v2 + ((v2 + v2) + v2)))
    } else {
        v27 := GenS1{F0: 408, F1: (v6 - (v6 ^ -128)), F2: 923}
        for fi28 := 0; fi28 < 1; fi28++ {
            v29 := GenS2{F0: (-(v6))}
            switch v2 {
            case "a":
                v2 = v2
            default:
                if true {
                    switch "" {
                    case "hello":
                        var v30 bool = false
                        _ = v30
                    case "\n\t":
                        var v31 uint8 = 2
                        _ = v31
                    case "foo bar":
                        var v32 float32 = 3.14159
                        _ = v32
                    default:
                        v2 = v2
                    }
                    if true {
                        v33 := GenS1{F0: 0, F1: (v6 % ((-60) | 1)), F2: -1}
                        v34 := GenS1{F0: 42, F1: 1, F2: 0}
                        for fi35 := 0; fi35 < 15; fi35++ {
                            if true {
                                switch v2 {
                                case "\n\t":
                                    var v36 uint = 0
                                    _ = v36
                                case "":
                                    fmt.Println(v6, fi35, v2)
                                }
                                v37 := GenS2{F0: v6}
                                _ = v37
                            } else {
                                switch v2 {
                                case "\n\t":
                                    delete(v4, 2)
                                case "foo bar":
                                    var v38 uint16 = 0
                                    _ = v38
                                default:
                                    if ((((-2.25 - (((float32(v13) * 1e10) + 0.1) * ((0.1 / 3.14159) * 0.0))) * 1e10) * (float32((v13 * v13)) / 1.5)) <= (float32((v13 + 0.0)) / 3.14159)) {
                                        v39 := GenS3{F0: fi28}
                                        fmt.Println(v1)
                                        var v40 uint32 = 42
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v3[v2] = true
                                        var v41 uint = 0
                                        _ = v41
                                    }
                                    v6 = (-(v6))
                                    v42 := GenS4{F0: 0}
                                    _ = v42
                                }
                                var v43 uint8 = 0
                                var v44 uint32 = 2
                                v1 += (v1 % ((int(v44)) | 1))
                                _ = v43
                                _ = v44
                            }
                            v1 += (v1 >> ((4294967295) & 63))
                            var v45 int8 = 100
                            if (!false) {
                                switch v2 {
                                case "":
                                    switch ((("a" + v2) + "a") + v2) {
                                    case "\n\t":
                                        switch "a" {
                                        case "hello":
                                            var v46 int32 = 42
                                            _ = v46
                                        default:
                                            var v47 bool = true
                                            _ = v47
                                        }
                                    case "hello":
                                        if false {
                                            v2 = v2
                                        }
                                    case "a":
                                        v48 := GenS5{F0: -2147483648}
                                        _ = v48
                                    default:
                                        if false {
                                            v49 := GenS6{F0: (-(fi35)), F1: "\n\t"}
                                            fmt.Println(fi35, v6, fi28)
                                            var v50 float64 = (float64((1e10 / float32((0.0 - v13)))) + (v13 + 1.5))
                                            _ = v49
                                            _ = v50
                                        } else {
                                            var v51 uint32 = 1
                                            _ = v51
                                        }
                                    }
                                case "a":
                                    v52 := GenS5{F0: 0}
                                    _ = v52
                                case "foo bar":
                                    v14 += 642
                                default:
                                    v13 = 1e10
                                }
                                v53 := GenS7{F0: 0}
                                var v54 uint64 = 42
                                _ = v53
                                _ = v54
                            }
                            _ = fi35
                            _ = v45
                        }
                        _ = v33
                        _ = v34
                    }
                }
            }
            _ = len(v5)
            _ = fi28
            _ = v29
        }
        v2 = "a"
        _ = v27
    }
    var v55 int = v1
    v5 = append(v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
