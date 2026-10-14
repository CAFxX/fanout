package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 string }
type GenS6 struct { F0 uint16; F1 int64 }
func genFn0(p0 string, p1 int64, p2 bool) uint32 {
    p1 = (-(p1))
    if p2 {
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 string = p0
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := GenS1{F0: -5, F1: -32768}
                if false {
                    v6 = "a"
                }
                v6 = "hello"
                _ = fi7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
        p0 = (p0 + p0)
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: 42, F1: 42}
            _ = fi9
            _ = v10
        }
    } else {
        var v11 uint64 = 1
        fmt.Println(p2)
        _ = v11
    }
    var v12 uint64 = 4294967295
    var v13 uint64 = (v12 % (((^(v12))) | 1))
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            fmt.Println(fi15, fi15, fi14)
            fmt.Println(p0, p2, p1)
            p2 = p2
            switch p0 {
            case "a":
                var v16 uint16 = 0
                _ = v16
            case "hello":
                p0 = "foo bar"
            default:
                var v17 int32 = 411
                _ = v17
            }
            if false {
                var v18 int8 = 100
                if ("a" == (p0 + p0)) {
                    fmt.Println(p0, p2)
                    switch p0 {
                    case "":
                        fmt.Println(p1)
                    }
                    var v19 float64 = 3.14159
                    if ((((-2.25 / 1.5) * (-2.25 * 3.14159)) - float32(1.5)) != ((1.5 / -2.25) - ((0.0 * 0.1) * (3.14159 * 0.1)))) {
                        if p2 {
                            switch (("a" + p0) + "a") {
                            case "hello":
                                var v20 float64 = (v19 - (v19 + ((v19 * -2.25) + v19)))
                                _ = v20
                            case "\n\t":
                                switch p0 {
                                case "foo bar":
                                    var v21 int8 = v18
                                    _ = v21
                                }
                            case "a":
                                var v22 bool = true
                                _ = v22
                            }
                            if (2 <= 984) {
                                v18 -= v18
                                p1 = p1
                            }
                            switch fi15 {
                            case 32768:
                                v23 := GenS1{F0: -128, F1: 0}
                                p1 -= p1
                                var v24 int64 = 9223372036854775807
                                _ = v23
                                _ = v24
                            default:
                                var v25 float64 = v19
                                _ = v25
                            }
                            if (v13 >= (v13 << ((65535) & 63))) {
                                if ((v18 <= v18) || ((float32(float64((3.14159 * (3.14159 - 1.5)))) - 0.1) <= (3.14159 / (float32(v19) - 1e10)))) {
                                    v26 := []int{(fi15 ^ int(v18)), 32768}
                                    p2 = false
                                    _ = v26
                                }
                            } else {
                                var v27 bool = (v18 <= 0)
                                fmt.Println(fi15, v27)
                                _ = v27
                            }
                        }
                        fmt.Println(v18, p0, v13)
                        v28 := []int{1, -9223372036854775808}
                        _ = v28
                    }
                    var v29 int32 = 1
                    _ = v19
                    _ = v29
                }
                v30 := []int{fi15, (fi14 / ((4294967295) | 1))}
                fmt.Println(p1)
                p0 = "a"
                var v31 string = p0
                v32 := GenS2{F0: (p0 + "a"), F1: p1}
                _ = v18
                _ = v30
                _ = v31
                _ = v32
            }
            v33 := []int{fi14, fi15}
            fmt.Println(v12, p1)
            switch fi15 {
            case 722:
                v34 := GenS2{F0: (p0 + ((("a" + ("" + p0)) + ("foo bar" + "a")) + p0)), F1: 0}
                _ = v34
            case 9223372036854775807:
                var v35 uint32 = 479
                _ = v35
            default:
                var v36 float64 = 0.0
                _ = v36
            }
            _ = fi15
            _ = v33
        }
        var v37 int64 = 970
        p2 = p2
        _ = fi14
        _ = v37
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 2
}
func genFn1(p0 int16, p1 string) int {
    v38 := []int{127, 32768}
    var v39 uint = 65535
    v40 := GenS3{F0: "", F1: 365}
    var v41 int32 = -16
    if (366 >= 2) {
        v38 = append(v38, 891)
    } else {
        var v42 float64 = 1.5
        _ = v42
    }
    _ = p0
    _ = p1
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    return 127
}

func main() {
    v1 := []string{}
    v2 := []string{}
    var v3 string = "hello"
    var v4 int8 = -29
    var v43 uint8 = 0
    for fi44 := 0; fi44 < 4; fi44++ {
        for fi45 := 0; fi45 < 13; fi45++ {
            v3 = (((v3 + "") + v3) + ("\n\t" + ((("a" + v3) + v3) + (("foo bar" + v3) + "foo bar"))))
            var v46 uint8 = v43
            _ = fi45
            _ = v46
        }
        var v47 uint16 = uint16(v43)
        _ = fi44
        _ = v47
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v48 bool = (1 > 201)
    if false {
        v49 := GenS1{F0: v4, F1: -1}
        for fi50 := 0; fi50 < 14; fi50++ {
            switch 967 {
            case 394:
                fmt.Println(v43, v4, v43)
            default:
                v51 := GenS1{F0: v4, F1: 1}
                _ = v51
            }
            var v52 uint = 39
            _ = fi50
            _ = v52
        }
        _ = v49
    }
    v53 := GenS2{F0: v3, F1: 1}
    switch "" {
    case "":
        fmt.Println(v3, v3, v48)
    case "foo bar":
        for fi54 := 0; fi54 < 5; fi54++ {
            if v48 {
                v55 := GenS4{F0: -128}
                v56 := GenS5{F0: v43, F1: 42, F2: (v3 + (v3 + "foo bar"))}
                _ = v55
                _ = v56
            }
            _ = fi54
        }
    case "a":
        var v57 uint32 = 225
        _ = v57
    }
    for fi58 := 0; fi58 < 17; fi58++ {
        v48 = ((-49 >= 42) && v48)
        _ = len(v1)
        _ = fi58
    }
    switch "foo bar" {
    case "":
        v3 = "\n\t"
    case "foo bar":
        if v48 {
            var v59 uint8 = 168
            if (24 <= v59) {
                if v48 {
                    v4 = v4
                    for fi60 := 0; fi60 < 18; fi60++ {
                        for fi61 := 0; fi61 < 8; fi61++ {
                            fmt.Println(v48, v43, v4)
                            var v62 int8 = v4
                            var v63 uint8 = v59
                            _ = fi61
                            _ = v62
                            _ = v63
                        }
                        _ = fi60
                    }
                }
            } else {
                if v48 {
                    fmt.Println(v48, v59)
                    fmt.Println(v59, v3)
                    fmt.Println(v43, v3)
                } else {
                    for fi64 := 0; fi64 < 5; fi64++ {
                        fmt.Println(v3, v4)
                        _ = fi64
                    }
                    fmt.Println(v48, v48, v43)
                    switch ("hello" + ("" + v3)) {
                    case "foo bar":
                        for fi65 := 0; fi65 < 7; fi65++ {
                            fmt.Println(v59, fi65)
                            v4 -= -52
                            fmt.Println(v4)
                            _ = fi65
                        }
                    default:
                        for fi66 := 0; fi66 < 18; fi66++ {
                            for fi67 := 0; fi67 < 13; fi67++ {
                                fmt.Println(fi66, v48, fi67)
                                v3 = "a"
                                _ = fi67
                            }
                            if false {
                                v43 = (v43 >> ((1) & 63))
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi66
                        }
                    }
                }
                switch 128 {
                case -227:
                    v68 := GenS6{F0: 65535, F1: 1}
                    _ = v68
                case 4294967295:
                    v3 = (v3 + v3)
                case 809:
                    for fi69 := 0; fi69 < 11; fi69++ {
                        for fi70 := 0; fi70 < 5; fi70++ {
                            switch ((v3 + v3) + ("\n\t" + v3)) {
                            case "hello":
                                var v71 uint8 = 1
                                _ = v71
                            default:
                                var v72 string = ("foo bar" + ((((v3 + ("\n\t" + v3)) + ((v3 + ("a" + ("a" + "\n\t"))) + "foo bar")) + v3) + v3))
                                _ = v72
                            }
                            v59 -= (v59 ^ (^(v59)))
                            if len(v1) > 0 {
                                v1[0] = v3
                            }
                            _ = fi70
                        }
                        _ = fi69
                    }
                }
            }
            var v73 float64 = 3.14159
            _ = v59
            _ = v73
        } else {
            v48 = (false || true)
        }
    }
    v43 -= v43
    var v74 float64 = 0.0
    for fi75 := 0; fi75 < 3; fi75++ {
        v43 += 0
        _ = fi75
    }
    v74 = v74
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v48
    _ = v53
    _ = v74
}
