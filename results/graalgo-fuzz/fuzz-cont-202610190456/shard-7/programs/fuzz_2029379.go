package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 bool; F2 uint32 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint; F1 uint8; F2 int16 }
type GenS10 struct { F0 int32; F1 int8 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    var v5 uint32 = 2147483648
    switch "hello" {
    case "":
        for fi6 := 0; fi6 < 13; fi6++ {
            delete(v2, (v5 & 0))
            v7 := []int{(^(fi6))}
            _ = fi6
            _ = v7
        }
    case "\n\t":
        var v8 uint32 = 0
        _ = v8
    case "hello":
        var v9 int = 128
        _ = v9
    }
    switch "a" {
    case "hello":
        _ = len(v4)
    case "":
        v1 = v1
    case "foo bar":
        if false {
            v10 := []int{-776, 995}
            if (2 <= 42) {
                v1 += (^(v1))
            }
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 string = "hello"
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v5 -= (^(v5))
        v14 := []int{int(v1), 0, (fi13 >> ((255) & 63))}
        switch fi13 {
        case 32768:
            var v15 uint32 = 2147483648
            _ = v15
        case 32767:
            if (-1 > (fi13 - fi13)) {
                if (true && false) {
                    if (65535 < 482) {
                        if true {
                            v5 = 1
                        } else {
                            for fi16 := 0; fi16 < 13; fi16++ {
                                v17 := GenS1{F0: 42, F1: fi13, F2: true}
                                _ = fi16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 1; fi18++ {
                                if true {
                                    var v19 float64 = 3.14159
                                    v20 := GenS2{F0: fi18, F1: 65535}
                                    if true {
                                        var v21 float32 = 0.1
                                        fmt.Println(v1, v1, fi18)
                                        switch (("hello" + "foo bar") + "foo bar") {
                                        case "a":
                                            switch fi18 {
                                            case 295:
                                                v22 := GenS1{F0: 897, F1: 32767, F2: false}
                                                _ = v22
                                            case 128:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v23 uint16 = 2
                                                    _ = v23
                                                case "\n\t":
                                                    if (-253 <= (-(v1))) {
                                                        var v24 uint8 = 1
                                                        _ = v24
                                                    } else {
                                                        v5 += 42
                                                        v25 := GenS2{F0: fi13, F1: 0}
                                                        v1 -= v1
                                                        _ = v25
                                                    }
                                                default:
                                                    v21 = 1.5
                                                }
                                            case -2147483648:
                                                v26 := GenS3{F0: (v5 + (v5 | 65535)), F1: 42, F2: 2147483647}
                                                _ = v26
                                            }
                                            fmt.Println(v5)
                                        case "hello":
                                            v5 = (^(v5))
                                        }
                                        v21 += (v21 + ((float32(1e10) - v21) * (float32((float64(1e10) + 1.5)) - ((1.5 - v21) + v21))))
                                        v27 := GenS4{F0: 0, F1: 255, F2: 56}
                                        v21 = 0.0
                                        v28 := GenS3{F0: (v5 % ((v5) | 1)), F1: 4294967295, F2: (v5 | (^(v5)))}
                                        _ = v21
                                        _ = v27
                                        _ = v28
                                    } else {
                                        v29 := GenS5{F0: ("foo bar" + ("a" + ""))}
                                        v5 += 1
                                        var v30 float64 = (v19 + v19)
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 uint8 = 0
                                    _ = v19
                                    _ = v20
                                    _ = v31
                                }
                                _ = fi18
                            }
                            var v32 uint64 = 980
                            _ = v32
                        }
                        v33 := GenS6{F0: -1, F1: (0 <= v1), F2: v5}
                        _ = v33
                    }
                }
                if (!(-9223372036854775808 != -9223372036854775808)) {
                    switch fi13 {
                    case 777:
                        if (true || (false && false)) {
                            v1 -= (-(v1))
                        }
                    }
                    v14 = append(v14, (fi13 - -1))
                } else {
                    v34 := GenS1{F0: 42, F1: 128, F2: (1 >= 2)}
                    _ = v34
                }
                var v35 int64 = 50
                _ = v35
            }
        default:
            v36 := GenS3{F0: (v5 ^ v5), F1: 969, F2: 600}
            _ = v36
        }
        _ = fi13
        _ = v14
    }
    v5 = v5
    v5 = (v5 ^ v5)
    v5 += v5
    var v37 int = -756
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 uint32 = v5
        delete(v4, fi38)
        _ = fi38
        _ = v39
    }
    if true {
        for fi40 := 0; fi40 < 1; fi40++ {
            if (false || true) {
                if true {
                    if true {
                        v41 := GenS7{F0: -1, F1: 31}
                        _ = v41
                    } else {
                        var v42 uint64 = 42
                        v43 := GenS1{F0: 42, F1: v37, F2: true}
                        _ = v42
                        _ = v43
                    }
                    if false {
                        switch 32767 {
                        case -993:
                            var v44 int16 = (v1 | v1)
                            _ = v44
                        }
                        if true {
                            fmt.Println(v37)
                            fmt.Println(v1)
                        }
                    }
                    v45 := GenS1{F0: 2147483647, F1: v37, F2: (165 == (fi40 ^ v37))}
                    _ = v45
                } else {
                    v46 := GenS8{F0: 0}
                    switch ("\n\t" + "a") {
                    case "hello":
                        v5 = v5
                    case "a":
                        switch v37 {
                        case 4294967295:
                            for fi47 := 0; fi47 < 6; fi47++ {
                                var v48 uint8 = 227
                                if (true || true) {
                                    if false {
                                        fmt.Println(v37)
                                        var v49 uint = 18446744073709551615
                                        _ = v49
                                    } else {
                                        fmt.Println(fi40, fi47, v5)
                                        if (177 != (v48 % ((2) | 1))) {
                                            switch (("a" + ("hello" + "\n\t")) + ((("" + ("" + ((((("foo bar" + "a") + "") + "hello") + "hello") + "foo bar"))) + "hello") + ((("" + ("hello" + (("\n\t" + "") + "a"))) + "foo bar") + "a"))) {
                                            case "\n\t":
                                                if false {
                                                    var v50 int16 = (v1 | v1)
                                                    _ = v50
                                                }
                                            default:
                                                v2[(v5 ^ 4294967295)] = false
                                            }
                                            v5 = (^(v5))
                                        }
                                        if false {
                                            var v51 uint64 = 42
                                            fmt.Println(fi47, fi47)
                                            _ = v51
                                        } else {
                                            var v52 uint32 = v5
                                            _ = v52
                                        }
                                    }
                                    v5 -= 4294967295
                                }
                                v37 = 508
                                _ = fi47
                                _ = v48
                            }
                        }
                    case "\n\t":
                        for fi53 := 0; fi53 < 13; fi53++ {
                            v54 := GenS9{F0: 18446744073709551615, F1: 198, F2: 1}
                            v5 -= 945
                            v55 := GenS6{F0: -1, F1: true, F2: v5}
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    }
                    _ = v46
                }
                switch 4294967295 {
                case 1:
                    if (!false) {
                        fmt.Println(v37)
                        v1 = v1
                    } else {
                        fmt.Println(v1, v5, fi40)
                        if true {
                            switch (^(v37)) {
                            case -9223372036854775808:
                                var v56 int = (v37 % ((9223372036854775807) | 1))
                                _ = v56
                            case 42:
                                var v57 uint32 = 475
                                _ = v57
                            }
                            v58 := GenS10{F0: 549, F1: 42}
                            for fi59 := 0; fi59 < 19; fi59++ {
                                var v60 int64 = 1
                                _ = fi59
                                _ = v60
                            }
                            _ = v58
                        }
                    }
                case 42:
                    v1 = 0
                }
                v1 = 1
            } else {
                var v61 float64 = -2.25
                var v62 int = 2147483647
                _ = v61
                _ = v62
            }
            _ = fi40
        }
        var v63 int16 = v1
        v1 = 137
        _ = v63
    } else {
        v4[(v37 / ((int(v5)) | 1))] = -9223372036854775808
    }
    v64 := []int{v37}
    v5 = 1
    v65 := GenS3{F0: (v5 << ((4294967295) & 63)), F1: 0, F2: 2147483648}
    v1 -= 81
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v64:", len(v64), v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v64
    _ = v65
}
