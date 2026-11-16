package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0(p0 int64) bool {
    if true {
        var v5 uint = 42
        v5 += v5
        _ = v5
    }
    if (!((1 >= 0) || true)) {
        p0 = 0
        fmt.Println(p0)
    }
    fmt.Println(p0)
    _ = p0
    return (42 <= 32768)
}
func genFn1(p0 uint, p1 int16) int64 {
    v6 := GenS1{F0: 42}
    if ((1 != 1) || (-2147483648 <= 4294967295)) {
        switch "foo bar" {
        case "":
            if (true || true) {
                v7 := GenS1{F0: 42}
                _ = v7
            }
        case "\n\t":
            var v8 string = ("" + ("\n\t" + ""))
            _ = v8
        }
        v9 := []int{915, 1}
        _ = v9
    } else {
        if false {
            var v10 float64 = (float64(0.0) / float64(-2.25))
            if true {
                var v11 int = 931
                p0 += p0
                _ = v11
            }
            _ = v10
        }
    }
    var v12 int = -1
    var v13 int64 = 60
    fmt.Println(v13, p1, v13)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v6
    return (^(v13))
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]int)
    var v4 uint = 255
    fmt.Println(v1)
    v4 -= v4
    v1 -= 559
    if true {
        v1 += (v4 << (((v1 | v1)) & 63))
        if false {
            fmt.Println(v4, v1)
            v2 = (v2 << (((^(v4))) & 63))
            var v14 uint64 = 2
            _ = v14
        }
    } else {
        if true {
            v2 -= (v2 << (((v1 >> ((v1) & 63))) & 63))
            fmt.Println(v1, v2)
            v1 = (^(v1))
        } else {
            v1 += (^(v4))
            v15 := []int{404, 32768, -286}
            _ = v15
        }
        var v16 float32 = 1e10
        _ = v16
    }
    v4 = 18446744073709551615
    v4 = (v1 / (((v4 & 0)) | 1))
    v2 = -1
    v17 := []int{1}
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 string = (("\n\t" + "hello") + "\n\t")
        if false {
            if (-151 < 1000) {
                for fi20 := 0; fi20 < 17; fi20++ {
                    v21 := GenS2{F0: 2147483647}
                    _ = fi20
                    _ = v21
                }
                v1 -= 2
                var v22 float32 = float32(((0.1 + float64((-2.25 - (0.1 + -2.25)))) / 1e10))
                _ = v22
            }
            v4 = (^(v4))
            if ("\n\t" < "hello") {
                for fi23 := 0; fi23 < 7; fi23++ {
                    v24 := v3["foo bar"]
                    if (false && (float32(3.14159) > 3.14159)) {
                        var v25 int16 = 42
                        if ((!false) || (((float32((1.5 / 1.0)) + 3.14159) + 1e10) == 3.14159)) {
                            if (-128 >= 42) {
                                v2 += (v2 + genFn1(v4, (v25 ^ (v25 * (v25 - 533)))))
                            } else {
                                switch (v19 + (v19 + "\n\t")) {
                                case "foo bar":
                                    v4 -= v1
                                case "a":
                                    fmt.Println(v25)
                                default:
                                    v4 -= (v4 / (((^(v1))) | 1))
                                }
                                v2 += 1
                                v26 := GenS1{F0: 2}
                                _ = v26
                            }
                        } else {
                            v27 := GenS3{F0: 0, F1: 2, F2: 42}
                            _ = v27
                        }
                        if ((v19 + v19) >= "hello") {
                            var v28 uint16 = 42
                            _ = v28
                        } else {
                            v4 = (v1 & v1)
                            v19 = "foo bar"
                        }
                        _ = v25
                    }
                    v1 += v4
                    if true {
                        v17 = append(v17, (fi23 | fi18))
                        var v29 int64 = 306
                        _ = v29
                    }
                    v4 += (v4 | v4)
                    delete(v3, "\n\t")
                    v1 = v1
                    var v30 uint64 = 748
                    v30 = (^(v30))
                    switch 2147483647 {
                    case 6:
                        var v31 int = (fi23 << (((v4 * (v1 << ((0) & 63)))) & 63))
                        _ = v31
                    }
                    _ = fi23
                    _ = v24
                    _ = v30
                }
            }
        } else {
            var v32 float64 = (1e10 + 0.1)
            if ((v2 / ((-9223372036854775808) | 1)) > v2) {
                var v33 float32 = -2.25
                v17 = append(v17, 718)
                fmt.Println(v2)
                _ = v33
            } else {
                switch (v19 + "foo bar") {
                case "a":
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                }
                var v34 int = fi18
                if (1 <= 0) {
                    v35 := GenS3{F0: 1, F1: 2, F2: 0}
                    _ = v35
                } else {
                    v36 := GenS3{F0: 2, F1: 0, F2: 0}
                    fmt.Println(v4, v1)
                    var v37 uint = v4
                    _ = v36
                    _ = v37
                }
                _ = v34
            }
            switch (v19 + v19) {
            case "\n\t":
                if (v2 > (-(v2))) {
                    for fi38 := 0; fi38 < 7; fi38++ {
                        var v39 uint8 = 108
                        fmt.Println(fi18)
                        _ = fi38
                        _ = v39
                    }
                    if false {
                        for fi40 := 0; fi40 < 6; fi40++ {
                            v32 = 3.14159
                            fmt.Println(v4, fi40)
                            _ = fi40
                        }
                    } else {
                        for fi41 := 0; fi41 < 8; fi41++ {
                            switch ("hello" + ((v19 + ("foo bar" + ("hello" + "a"))) + v19)) {
                            case "a":
                                var v42 uint32 = 394
                                _ = v42
                            case "\n\t":
                                switch 2147483647 {
                                case 722:
                                    switch (fi41 % (((fi18 & (fi18 & fi41))) | 1)) {
                                    case 593:
                                        v43 := GenS4{F0: 919}
                                        _ = v43
                                    case 278:
                                        delete(v3, "")
                                    case 32768:
                                        switch (fi41 >> (((v4 % ((v1) | 1))) & 63)) {
                                        case 4294967295:
                                            v1 -= (v4 | v1)
                                        case -723:
                                            switch "\n\t" {
                                            case "\n\t":
                                                delete(v3, "")
                                                var v44 int8 = 127
                                                fmt.Println(fi41)
                                                _ = v44
                                            case "a":
                                                v19 = "\n\t"
                                            }
                                        }
                                    }
                                }
                            case "foo bar":
                                v32 -= 0.1
                            }
                            if false {
                                v45 := GenS2{F0: -662}
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                switch v19 {
                                case "\n\t":
                                    fmt.Println(v4)
                                case "a":
                                    if true {
                                        _ = len(v17)
                                        fmt.Println(fi18, v19)
                                    }
                                default:
                                    var v46 int = -9223372036854775808
                                    _ = v46
                                }
                                _ = v45
                            }
                            var v47 uint32 = 2147483648
                            _ = fi41
                            _ = v47
                        }
                        if true {
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v2)
                            case "foo bar":
                                if false {
                                    if (1.5 <= v32) {
                                        var v48 uint16 = 1
                                        _ = v48
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        v49 := GenS5{F0: 1}
                                        _ = v49
                                    case "a":
                                        switch fi18 {
                                        case 42:
                                            var v50 uint32 = 860
                                            _ = v50
                                        case -639:
                                            v2 = v2
                                        default:
                                            delete(v3, v19)
                                        }
                                    }
                                } else {
                                    if ((2147483647 > -554) && true) {
                                        if true {
                                            var v51 int8 = 1
                                            for fi52 := 0; fi52 < 8; fi52++ {
                                                _ = len(v3)
                                                _ = fi52
                                            }
                                            _ = v51
                                        }
                                        fmt.Println(v32, v4)
                                    }
                                }
                            }
                        }
                    }
                }
            case "":
                delete(v3, v19)
            }
            _ = v32
        }
        _ = fi18
        _ = v19
    }
    if (2 != 1) {
        if (1000 != 2) {
            fmt.Println(v2, v1, v4)
            if (!true) {
                switch 13 {
                case 496:
                    v4 -= v4
                case 128:
                    delete(v3, "\n\t")
                default:
                    v2 = (v2 * 9223372036854775807)
                }
            }
            v53 := GenS6{F0: 824}
            _ = v53
        }
        fmt.Println(v4, v2, v1)
    }
    fmt.Println(v4, v1, v2)
    v2 = (v2 | -744)
    for fi54 := 0; fi54 < 20; fi54++ {
        _ = len(v17)
        _ = fi54
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    if true {
        v2 -= 9223372036854775807
    } else {
        var v55 float32 = float32(float64(3.14159))
        _ = len(v17)
        var v56 int8 = 77
        _ = v55
        _ = v56
    }
    fmt.Println(v2, v2, v2)
    v57 := GenS7{F0: 2, F1: 1, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v57
}
