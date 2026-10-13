package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 bool; F2 int16 }
type GenS5 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint16 = 0
    var v2 int32 = 1
    var v3 int32 = (v2 + (v2 & (v2 << ((18446744073709551615) & 63))))
    v4 := []int{-2147483648, 359, -237, -1}
    v5 := make(map[int]bool)
    if (1 == 1) {
        var v6 float32 = float32(-2.25)
        delete(v5, int(v1))
        _ = v6
    } else {
        var v7 uint8 = 242
        var v8 string = "foo bar"
        if (917 <= 4294967295) {
            for fi9 := 0; fi9 < 14; fi9++ {
                switch ("\n\t" + v8) {
                case "a":
                    var v10 uint8 = 127
                    _ = v10
                case "":
                    if (-812 == (fi9 << ((4294967295) & 63))) {
                        if false {
                            v5[(^(fi9))] = ((true || (false && true)) && (false && (true && ((^(v1)) == v1))))
                            var v11 uint32 = 2
                            _ = v11
                        } else {
                            for fi12 := 0; fi12 < 9; fi12++ {
                                var v13 int16 = 264
                                _ = fi12
                                _ = v13
                            }
                            _ = len(v5)
                        }
                    }
                }
                _ = fi9
            }
        } else {
            v14 := GenS1{F0: 1, F1: 547}
            v15 := GenS2{F0: false}
            var v16 uint32 = 42
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v7
        _ = v8
    }
    var v17 int64 = -1
    if (42 != -160) {
        var v18 int64 = v17
        v19 := GenS2{F0: true}
        var v20 int16 = 322
        _ = v18
        _ = v19
        _ = v20
    }
    if ((v3 == v2) && true) {
        if true {
            for fi21 := 0; fi21 < 14; fi21++ {
                v22 := GenS1{F0: -76, F1: v3}
                if (!false) {
                    fmt.Println(v1, fi21, v3)
                    fmt.Println(v3, v17)
                    if (302 == 4294967295) {
                        v2 += v2
                        switch ("a" + "a") {
                        case "\n\t":
                            for fi23 := 0; fi23 < 12; fi23++ {
                                if false {
                                    var v24 uint64 = 0
                                    _ = v24
                                } else {
                                    var v25 uint32 = 65535
                                    var v26 int32 = 2147483647
                                    switch (("" + "") + ("\n\t" + "hello")) {
                                    case "":
                                        var v27 float32 = 3.14159
                                        var v28 float32 = v27
                                        fmt.Println(v25, v27)
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = v25
                                    _ = v26
                                }
                                v2 = v3
                                var v29 bool = false
                                _ = fi23
                                _ = v29
                            }
                        case "foo bar":
                            var v30 uint64 = 42
                            _ = v30
                        default:
                            switch fi21 {
                            case 32768:
                                var v31 int = -419
                                _ = v31
                            case 42:
                                var v32 uint16 = (^(v1))
                                _ = v32
                            case -2147483648:
                                var v33 uint16 = (^(v1))
                                _ = v33
                            }
                        }
                    }
                }
                var v34 int32 = v2
                _ = fi21
                _ = v22
                _ = v34
            }
            var v35 uint16 = 1
            if (!false) {
                v3 = 1
            } else {
                var v36 float32 = 3.14159
                _ = v36
            }
            _ = v35
        } else {
            var v37 uint8 = 0
            fmt.Println(v37, v2)
            _ = v37
        }
    } else {
        var v38 float64 = 0.0
        var v39 uint16 = (^(v1))
        _ = v38
        _ = v39
    }
    if false {
        v40 := GenS2{F0: true}
        var v41 uint8 = 183
        switch 4294967295 {
        case -9223372036854775808:
            if true {
                if (!false) {
                    for fi42 := 0; fi42 < 15; fi42++ {
                        var v43 int8 = -128
                        for fi44 := 0; fi44 < 2; fi44++ {
                            v3 = 414
                            _ = fi44
                        }
                        var v45 uint8 = (v41 * 0)
                        _ = fi42
                        _ = v43
                        _ = v45
                    }
                    fmt.Println(v2)
                }
                if true {
                    var v46 float32 = 1.5
                    _ = v46
                }
            } else {
                v3 = (v3 * (-(v3)))
                if (true && (((-100 < 127) && (int(v17) != 32767)) || (true && false))) {
                    v47 := GenS2{F0: true}
                    _ = v47
                } else {
                    for fi48 := 0; fi48 < 6; fi48++ {
                        v49 := GenS1{F0: 127, F1: 0}
                        var v50 float32 = 1e10
                        var v51 uint8 = 31
                        _ = fi48
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                    v52 := GenS3{F0: "hello"}
                    _ = v52
                }
                _ = len(v4)
            }
        case -1:
            v3 = 2147483647
        case 4294967295:
            fmt.Println(v17, v41, v2)
        }
        _ = v40
        _ = v41
    }
    switch 32768 {
    case -1:
        if ("\n\t" == "\n\t") {
            for fi53 := 0; fi53 < 20; fi53++ {
                v54 := GenS2{F0: (v1 != v1)}
                switch 42 {
                case 1:
                    var v55 string = ("" + ("\n\t" + "\n\t"))
                    _ = v55
                case 4294967295:
                    for fi56 := 0; fi56 < 4; fi56++ {
                        v17 -= v17
                        _ = fi56
                    }
                default:
                    v57 := GenS4{F0: 255, F1: (((1e10 + 0.0) / ((-2.25 * (1e10 * 1e10)) * 1e10)) == (3.14159 - 1e10)), F2: 1000}
                    _ = v57
                }
                if (!false) {
                    v1 += (v1 | v1)
                } else {
                    fmt.Println(v1, v3)
                    var v58 int32 = v2
                    v1 += 590
                    _ = v58
                }
                _ = fi53
                _ = v54
            }
        } else {
            var v59 uint32 = 755
            switch "" {
            case "foo bar":
                if ((32767 != 32767) || false) {
                    fmt.Println(v2)
                    switch -2147483648 {
                    case 2147483647:
                        var v60 uint16 = v1
                        _ = v60
                    case 127:
                        var v61 int = 2147483647
                        _ = v61
                    case 128:
                        v59 = (v59 >> ((904) & 63))
                    }
                    v1 += v1
                } else {
                    fmt.Println(v2)
                }
            case "a":
                v62 := GenS5{F0: true, F1: 100}
                _ = v62
            case "":
                switch ("foo bar" + "hello") {
                case "hello":
                    fmt.Println(v59)
                default:
                    _ = len(v5)
                }
            default:
                var v63 uint32 = (^(v59))
                _ = v63
            }
            _ = v59
        }
    case 826:
        switch "a" {
        case "foo bar":
            v2 = (v3 & (^(v2)))
        default:
            v2 -= v2
        }
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        var v65 uint16 = (v1 % ((v1) | 1))
        fmt.Println(v65, fi64)
        _ = fi64
        _ = v65
    }
    _ = len(v4)
    for fi66 := 0; fi66 < 1; fi66++ {
        var v67 uint = 977
        _ = fi66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
