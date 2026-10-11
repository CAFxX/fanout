package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int; F1 int8; F2 int16 }
type GenS3 struct { F0 int32; F1 string; F2 uint8 }
type GenS4 struct { F0 int32; F1 int; F2 uint64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint8, p1 int32, p2 uint) int {
    p1 += (p1 % ((p1) | 1))
    v6 := GenS1{F0: (p2 >> ((529) & 63)), F1: (p1 != p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -2147483648
}

func main() {
    v1 := []string{"\n\t", "hello", ("hello" + ""), "\n\t"}
    v2 := []string{"a", "hello", "", ""}
    var v3 int = -660
    v4 := make(map[uint32]string)
    var v5 float64 = 1.5
    var v7 uint = 65535
    v7 += (v7 / (((v7 >> (((v7 & v7)) & 63))) | 1))
    v5 = ((1e10 / 1.0) * v5)
    v8 := v4[1]
    fmt.Println(v8, v3, v8)
    switch "\n\t" {
    case "":
        v9 := GenS1{F0: (^(v7)), F1: false}
        _ = v9
    case "hello":
        fmt.Println(v5)
    }
    var v10 bool = false
    v8 = v8
    v4[42] = v8
    v11 := GenS2{F0: (v3 & v3), F1: -22, F2: -89}
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := v4[2147483648]
        if true {
            delete(v4, 0)
            v14 := GenS2{F0: -1, F1: 0, F2: 1000}
            _ = v14
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            switch "foo bar" {
            case "foo bar":
                if ((1000 > -32768) || (127 == 1)) {
                    switch v13 {
                    case "hello":
                        var v16 int8 = int8(v3)
                        _ = v16
                    case "":
                        v5 = -2.25
                    }
                }
            default:
                if v10 {
                    var v17 string = v13
                    if v10 {
                        var v18 uint32 = 2147483647
                        if ((fi12 >> (((v7 - 472)) & 63)) > 0) {
                            v13 = v13
                            _ = len(v1)
                        } else {
                            v13 = (v17 + "foo bar")
                            fmt.Println(v3, v3, fi15)
                            fmt.Println(fi12, fi15)
                        }
                        var v19 float32 = 0.0
                        switch fi15 {
                        case 127:
                            v20 := v4[v18]
                            _ = v20
                        case -323:
                            switch v3 {
                            case 4294967295:
                                v21 := v4[v18]
                                _ = v21
                            }
                        case 2147483647:
                            if (!(((v8 + ((v13 + (v17 + v17)) + v17)) + v13) == v13)) {
                                v3 += 32767
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = v17
                                }
                                v22 := GenS1{F0: 255, F1: v10}
                                switch (v13 + (v17 + "")) {
                                case "\n\t":
                                    fmt.Println(v13, fi15, v7)
                                }
                                v23 := GenS2{F0: 127, F1: -128, F2: -757}
                                _ = v22
                                _ = v23
                            }
                        default:
                            var v24 uint32 = v18
                            if v10 {
                                v19 = (-2.25 * float32((v5 + 1e10)))
                                switch genFn0(1, 0, v7) {
                                case 1:
                                    fmt.Println(v5, v19)
                                default:
                                    var v25 int64 = -9223372036854775808
                                    if false {
                                        if v10 {
                                            var v26 float64 = v5
                                            v27 := GenS3{F0: 853, F1: v8, F2: 2}
                                            v7 -= (v7 * (v7 * (v7 << (((v7 % (((^(v7))) | 1))) & 63))))
                                            _ = v26
                                            _ = v27
                                        } else {
                                            var v28 float32 = (float32(v5) * (3.14159 + v19))
                                            _ = v28
                                        }
                                        v29 := GenS4{F0: 42, F1: genFn0(2, -610, v7), F2: 4294967295}
                                        var v30 int32 = int32(v24)
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = v25
                                }
                            }
                            _ = v24
                        }
                        _ = v18
                        _ = v19
                    }
                    _ = v17
                }
                switch fi12 {
                case -9223372036854775808:
                    if (100 == 42) {
                        v3 = fi15
                    } else {
                        var v31 uint = v7
                        _ = v31
                    }
                case 942:
                    fmt.Println(fi15)
                case 0:
                    switch v13 {
                    case "a":
                        if (1 < 1) {
                            if true {
                                v13 = v13
                            } else {
                                switch "a" {
                                case "":
                                    v5 -= ((1e10 + 0.1) * v5)
                                }
                            }
                            v32 := v4[2]
                            _ = v32
                        } else {
                            v33 := GenS3{F0: int32(v3), F1: ("hello" + (((v8 + "foo bar") + ("" + ((("a" + (v13 + (v8 + ("hello" + v13)))) + v8) + "\n\t"))) + ((v8 + v13) + (v13 + (v13 + v13))))), F2: 127}
                            _ = len(v4)
                            if v10 {
                                v34 := GenS2{F0: -1, F1: 127, F2: 42}
                                var v35 bool = v10
                                var v36 int16 = 1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            var v37 bool = (127 <= 120)
                            v1 = append(v1, (((v8 + "a") + "foo bar") + "foo bar"))
                            _ = v33
                            _ = v37
                        }
                    case "foo bar":
                        switch (fi12 - 2147483647) {
                        case -60:
                            switch "foo bar" {
                            case "a":
                                fmt.Println(fi15, v5, v3)
                                switch v3 {
                                case 865:
                                    v13 = "hello"
                                default:
                                    if (v10 && (1 > 1)) {
                                        fmt.Println(fi15, v7)
                                    }
                                }
                            case "":
                                if ((v10 && (!false)) || v10) {
                                    switch "\n\t" {
                                    case "":
                                        switch v3 {
                                        case -9223372036854775808:
                                            var v38 int32 = 1
                                            _ = v38
                                        }
                                    }
                                    v39 := GenS3{F0: 1, F1: (v13 + (v13 + (("" + "foo bar") + v8))), F2: 255}
                                    _ = v39
                                }
                            default:
                                var v40 uint64 = 472
                                _ = v40
                            }
                        case -1:
                            var v41 uint32 = 65535
                            _ = v41
                        case 0:
                            if v10 {
                                var v42 string = "\n\t"
                                v5 += (float64(-2.25) / v5)
                                v43 := GenS4{F0: 42, F1: v3, F2: 4294967295}
                                var v44 uint = 18446744073709551615
                                _ = v42
                                _ = v43
                                _ = v44
                            } else {
                                if v10 {
                                    v13 = ("\n\t" + v8)
                                    if v10 {
                                        var v45 int64 = -1
                                        var v46 uint8 = 221
                                        v47 := GenS2{F0: 858, F1: 42, F2: 1000}
                                        var v48 int8 = 42
                                        switch (v8 + "foo bar") {
                                        case "hello":
                                            v49 := GenS1{F0: (v7 % ((v7) | 1)), F1: v10}
                                            _ = v49
                                        case "":
                                            v45 = v45
                                        case "\n\t":
                                            if false {
                                                var v50 uint = 252
                                                v51 := GenS5{F0: v45}
                                                _ = v50
                                                _ = v51
                                            } else {
                                                var v52 int64 = v45
                                                _ = v52
                                            }
                                            switch ("a" + "\n\t") {
                                            case "a":
                                                if (true || (v10 && v10)) {
                                                    v8 = ("" + v8)
                                                }
                                            case "\n\t":
                                                v46 -= v46
                                            }
                                            var v53 uint = v7
                                            _ = v53
                                        }
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                            }
                        default:
                            if v10 {
                                if v10 {
                                    v10 = (v10 && v10)
                                    var v54 uint = (v7 | v7)
                                    _ = v54
                                }
                                var v55 string = "\n\t"
                                _ = v55
                            } else {
                                v13 = v8
                            }
                        }
                    case "\n\t":
                        var v56 bool = (543 == 32767)
                        _ = v56
                    }
                default:
                    switch fi15 {
                    case 1:
                        v8 = (((("\n\t" + ((v13 + v13) + "\n\t")) + v13) + v13) + v8)
                    }
                }
            }
            switch "a" {
            case "hello":
                switch v13 {
                case "hello":
                    fmt.Println(v13, v13, v5)
                case "\n\t":
                    v8 = ""
                }
            case "a":
                var v57 uint32 = 375
                _ = v57
            case "":
                var v58 float64 = (((((float64((1.5 / ((float32(v5) + -2.25) - 1.5))) * v5) / (float64(1e10) * 1e10)) * 3.14159) / (v5 - v5)) + ((float64((1e10 * 3.14159)) * v5) + (1.5 + (3.14159 / ((float64(3.14159) - v5) / (float64((0.1 - -2.25)) + ((v5 + v5) + v5)))))))
                var v59 int32 = -1
                switch v13 {
                case "foo bar":
                    v3 = v3
                case "":
                    var v60 uint32 = 65535
                    _ = v60
                default:
                    var v61 int8 = -48
                    _ = v61
                }
                _ = v58
                _ = v59
            }
            if (v10 || v10) {
                v62 := GenS2{F0: fi12, F1: -15, F2: 32767}
                _ = v62
            }
            _ = fi15
        }
        _ = fi12
        _ = v13
    }
    if (128 < 255) {
        if ((-1 >= v3) || (863 < -2147483648)) {
            if len(v2) > 0 {
                v2[0] = ((v8 + v8) + (v8 + "a"))
            }
            for fi63 := 0; fi63 < 1; fi63++ {
                v3 = v3
                v5 += 3.14159
                _ = fi63
            }
            for fi64 := 0; fi64 < 7; fi64++ {
                for fi65 := 0; fi65 < 11; fi65++ {
                    fmt.Println(v7, v7, v10)
                    v66 := GenS3{F0: int32(fi64), F1: v8, F2: 2}
                    _ = fi65
                    _ = v66
                }
                v67 := GenS2{F0: genFn0(127, -317, 1), F1: 42, F2: 1}
                _ = fi64
                _ = v67
            }
        } else {
            v8 = "\n\t"
            switch (((v8 + v8) + "foo bar") + "") {
            case "a":
                v8 = "foo bar"
            case "hello":
                v8 = (((v8 + ((v8 + "foo bar") + "")) + ("\n\t" + ((("" + v8) + ("" + "foo bar")) + (v8 + "")))) + v8)
            }
        }
        v10 = false
    }
    var v68 uint32 = 423
    var v69 uint = 694
    v10 = ((65535 == 961) || true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
    _ = v7
    _ = v8
}
