package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 int16; F2 string }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
type GenS4 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS5 struct { F0 int16; F1 int8; F2 int }
type GenS6 struct { F0 bool; F1 int64 }
func genFn0() uint32 {
    if true {
        if (42 < 123) {
            var v6 uint = 441
            _ = v6
        } else {
            switch ("foo bar" + (("" + "") + "a")) {
            case "foo bar":
                if false {
                    for fi7 := 0; fi7 < 11; fi7++ {
                        for fi8 := 0; fi8 < 10; fi8++ {
                            if true {
                                if true {
                                    if false {
                                        if (!(!false)) {
                                            fmt.Println(fi7)
                                            var v9 bool = true
                                            var v10 uint16 = 2
                                            _ = v10
                                            _ = v9
                                        } else {
                                            var v11 bool = false
                                            switch "a" {
                                            case "hello":
                                                v12 := []int{974, (fi8 % ((fi7) | 1))}
                                                _ = v12
                                            case "foo bar":
                                                if (!false) {
                                                    v13 := GenS1{F0: 2, F1: -288}
                                                    var v14 bool = true
                                                    var v15 int16 = 1
                                                    _ = v13
                                                    _ = v14
                                                    _ = v15
                                                }
                                            default:
                                                v11 = (1e10 >= 3.14159)
                                            }
                                            fmt.Println(fi7, fi7, v11)
                                            v16 := []int{(fi7 | -966), fi8}
                                            _ = len(v16)
                                            _ = v11
                                            _ = v16
                                        }
                                        if (682 > -224) {
                                            v17 := GenS1{F0: 865, F1: 42}
                                            _ = v17
                                        } else {
                                            var v18 float64 = (((float64(float32(0.1)) / -2.25) - ((3.14159 + (3.14159 + (3.14159 + (((3.14159 - (1e10 / 1e10)) * 1e10) * -2.25)))) - 1e10)) - (3.14159 - (-2.25 / 3.14159)))
                                            if (!true) {
                                                switch (fi8 | fi8) {
                                                case 42:
                                                    switch "a" {
                                                    case "":
                                                        switch "\n\t" {
                                                        case "a":
                                                            v19 := GenS1{F0: 255, F1: 1}
                                                            _ = v19
                                                        case "\n\t":
                                                            v18 -= (0.0 / 0.1)
                                                        default:
                                                            var v20 int16 = -1
                                                            _ = v20
                                                        }
                                                    case "\n\t":
                                                        v21 := GenS2{F0: 9223372036854775807, F1: -1, F2: ("a" + ((("" + "a") + "") + "foo bar"))}
                                                        _ = v21
                                                    case "hello":
                                                        v22 := []int{}
                                                        _ = v22
                                                    }
                                                }
                                                fmt.Println(v18)
                                            } else {
                                                switch "foo bar" {
                                                case "":
                                                    v18 += (v18 * 3.14159)
                                                default:
                                                    var v23 int = fi8
                                                    _ = v23
                                                }
                                                v18 -= 1e10
                                                v24 := []int{}
                                                _ = v24
                                            }
                                            fmt.Println(fi8, fi8)
                                            _ = v18
                                        }
                                        var v25 uint64 = 4294967295
                                        _ = v25
                                    }
                                } else {
                                    switch ("hello" + "\n\t") {
                                    case "\n\t":
                                        fmt.Println(fi7, fi7)
                                        switch (fi8 + 2147483647) {
                                        case 127:
                                            fmt.Println(fi7)
                                        case -10:
                                            v26 := []int{}
                                            _ = v26
                                        }
                                        v27 := GenS2{F0: 1, F1: -507, F2: "foo bar"}
                                        _ = v27
                                    }
                                }
                            } else {
                                if false {
                                    var v28 uint64 = 4294967295
                                    v29 := GenS3{F0: (v28 >> ((1) & 63)), F1: 1, F2: "\n\t"}
                                    fmt.Println(fi8, fi8)
                                    _ = v28
                                    _ = v29
                                }
                            }
                            _ = fi8
                        }
                        _ = fi7
                    }
                    var v30 uint = 42
                    fmt.Println(v30)
                    _ = v30
                }
            }
            if (273 > 128) {
                if false {
                    v31 := []int{972}
                    for fi32 := 0; fi32 < 7; fi32++ {
                        if len(v31) > 0 {
                            v31[0] = fi32
                        }
                        _ = fi32
                    }
                    for fi33 := 0; fi33 < 11; fi33++ {
                        var v34 int32 = 42
                        v35 := GenS3{F0: 0, F1: 65535, F2: ("" + "foo bar")}
                        fmt.Println(fi33, fi33)
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    _ = v31
                }
                if (32767 >= 0) {
                    var v36 int32 = 593
                    v36 -= v36
                    _ = v36
                } else {
                    v37 := GenS4{F0: 442, F1: -53, F2: 2}
                    for fi38 := 0; fi38 < 3; fi38++ {
                        v39 := []int{(fi38 | fi38)}
                        _ = fi38
                        _ = v39
                    }
                    _ = v37
                }
                var v40 uint16 = 42
                _ = v40
            } else {
                v41 := GenS4{F0: 4294967295, F1: 1, F2: 2}
                var v42 uint16 = 0
                v43 := []int{}
                _ = v41
                _ = v42
                _ = v43
            }
        }
    }
    var v44 int8 = 107
    v45 := []int{-465, 127, 7, 32768}
    var v46 int8 = v44
    _ = v44
    _ = v45
    _ = v46
    return 2
}

func main() {
    var v1 uint16 = 1
    var v2 int16 = -32768
    var v3 float64 = ((((3.14159 * 1.5) + 1.5) - 0.1) + (((0.0 - 3.14159) / 3.14159) - (0.0 * ((1e10 + -2.25) / 0.1))))
    var v4 int32 = 2147483647
    v5 := make(map[string]int)
    var v47 uint8 = 255
    v3 = v3
    for fi48 := 0; fi48 < 10; fi48++ {
        var v49 float32 = (1e10 * 1e10)
        var v50 float32 = ((v49 / (float32(0.1) * 3.14159)) * float32(v3))
        v49 += v49
        _ = fi48
        _ = v49
        _ = v50
    }
    var v51 int16 = (v2 | (^(v2)))
    v52 := GenS3{F0: 1, F1: genFn0(), F2: "\n\t"}
    v3 += (1.5 / v3)
    v53 := GenS5{F0: 0, F1: 1, F2: 2147483647}
    var v54 uint8 = (^(v47))
    for fi55 := 0; fi55 < 17; fi55++ {
        switch "hello" {
        case "\n\t":
            v54 += 127
        case "a":
            var v56 uint32 = 65535
            _ = v56
        default:
            var v57 string = "\n\t"
            _ = v57
        }
        if true {
            fmt.Println(v4, v3, v4)
            switch ((("\n\t" + "") + ("" + (("hello" + ("foo bar" + "foo bar")) + "foo bar"))) + "hello") {
            case "a":
                v58 := GenS6{F0: true, F1: 868}
                _ = v58
            case "\n\t":
                var v59 int64 = 814
                _ = v59
            case "foo bar":
                fmt.Println(v4, v4)
            default:
                switch (fi55 * 87) {
                case 127:
                    for fi60 := 0; fi60 < 14; fi60++ {
                        var v61 uint32 = 451
                        _ = fi60
                        _ = v61
                    }
                }
            }
            fmt.Println(v3, v1, v2)
        }
        var v62 int = int(v54)
        _ = fi55
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
