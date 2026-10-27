package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 bool; F1 int8 }
func genFn0(p0 uint16) uint64 {
    switch 769 {
    case 2147483647:
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{(fi7 % (((fi7 % ((int(p0)) | 1))) | 1)), fi7, 128, -481}
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(fi9, fi9)
        v10 := GenS1{F0: (fi9 % (((fi9 >> ((0) & 63))) | 1))}
        var v11 float32 = 0.0
        _ = fi9
        _ = v10
        _ = v11
    }
    p0 = (p0 ^ (^(p0)))
    _ = p0
    return 301
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 0
    v3 := make(map[string]int)
    var v4 float64 = (1.5 - float64(3.14159))
    v5 := []int{-252, 350}
    var v6 uint8 = 42
    if (32767 > 2) {
        v2 += (^(v2))
    }
    var v12 uint64 = 842
    if (!((false && true) && (!true))) {
        if (42 == 9223372036854775807) {
            var v13 bool = (4294967295 < (^(v12)))
            var v14 uint16 = 1
            _ = v13
            _ = v14
        }
    }
    var v15 int32 = -1
    if false {
        if true {
            if ((0.1 < 3.14159) || (!true)) {
                switch ("" + ("a" + "foo bar")) {
                case "hello":
                    v6 += (v6 >> ((255) & 63))
                case "\n\t":
                    var v16 int8 = -128
                    _ = v16
                }
                v17 := GenS1{F0: 9223372036854775807}
                v18 := GenS1{F0: 42}
                _ = v17
                _ = v18
            }
            switch ("foo bar" + "\n\t") {
            case "hello":
                v19 := GenS1{F0: 42}
                _ = v19
            case "":
                var v20 int8 = -61
                _ = v20
            default:
                switch "a" {
                case "foo bar":
                    switch -415 {
                    case 2147483647:
                        fmt.Println(v2, v2)
                    case -2147483648:
                        switch ("hello" + ("hello" + ("hello" + ("\n\t" + "")))) {
                        case "foo bar":
                            v4 -= (v4 / v4)
                        case "hello":
                            v12 = v12
                        case "\n\t":
                            for fi21 := 0; fi21 < 13; fi21++ {
                                fmt.Println(fi21)
                                v22 := GenS2{F0: -9223372036854775808, F1: "a"}
                                v23 := GenS3{F0: 297}
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        default:
                            v12 += genFn0(0)
                        }
                    case 127:
                        fmt.Println(v2, v12, v4)
                    default:
                        if (0 >= 2) {
                            v24 := GenS4{F0: "foo bar", F1: (true && (!true))}
                            v25 := GenS4{F0: "hello", F1: false}
                            v2 = (v2 / ((v2) | 1))
                            _ = v24
                            _ = v25
                        }
                    }
                case "hello":
                    var v26 int16 = v2
                    _ = v26
                case "":
                    if false {
                        if (false && (false && (uint(v12) > uint(v6)))) {
                            if (!true) {
                                v15 += (v15 % ((435) | 1))
                                var v27 bool = true
                                _ = v27
                            } else {
                                var v28 int16 = v2
                                v29 := v3[("a" + (("\n\t" + "\n\t") + ("" + ("" + "a"))))]
                                _ = v28
                                _ = v29
                            }
                            switch "" {
                            case "\n\t":
                                var v30 uint8 = v6
                                _ = v30
                            }
                        } else {
                            v2 = v2
                            var v31 int = -126
                            _ = v31
                        }
                        var v32 uint = 2
                        fmt.Println(v15, v32, v12)
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 20; fi33++ {
                            var v34 int8 = 0
                            var v35 uint32 = 65535
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        var v36 float32 = ((3.14159 * 3.14159) + ((float32(v4) + (3.14159 - 0.1)) - 0.0))
                        fmt.Println(v4, v2, v2)
                        _ = v36
                    }
                default:
                    if (32767 > 127) {
                        if len(v5) > 0 {
                            v5[0] = -796
                        }
                        var v37 bool = false
                        for fi38 := 0; fi38 < 13; fi38++ {
                            var v39 int16 = v2
                            var v40 uint64 = 4294967295
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        _ = v37
                    } else {
                        v41 := v1[574]
                        if true {
                            var v42 bool = false
                            _ = v42
                        }
                        _ = v41
                    }
                }
            }
        } else {
            var v43 int32 = v15
            var v44 float32 = 0.1
            _ = v43
            _ = v44
        }
        v45 := GenS5{F0: uint(v12), F1: 255}
        var v46 int = -9223372036854775808
        _ = v45
        _ = v46
    }
    var v47 uint = 4294967295
    if (!((^(v15)) <= int32(v6))) {
        delete(v3, (("hello" + "\n\t") + "a"))
    } else {
        v4 = -2.25
        for fi48 := 0; fi48 < 6; fi48++ {
            v47 = v47
            fmt.Println(fi48)
            _ = fi48
        }
    }
    v15 = (v15 - v15)
    var v49 uint8 = v6
    v50 := GenS6{F0: ((v47 - (v47 + uint(v12))) <= uint(v6)), F1: 15}
    var v51 int32 = v15
    fmt.Println(v51, v49, v49)
    var v52 uint16 = 2
    v15 = (-(v15))
    var v53 int8 = int8(v47)
    var v54 int = 36
    v51 = v51
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
