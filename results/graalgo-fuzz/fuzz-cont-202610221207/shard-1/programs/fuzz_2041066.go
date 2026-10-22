package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint; F1 int32 }
func genFn0(p0 string, p1 int16, p2 bool) uint32 {
    var v6 bool = false
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 string = (("a" + p0) + p0)
        switch v8 {
        case "foo bar":
            for fi9 := 0; fi9 < 6; fi9++ {
                switch "hello" {
                case "\n\t":
                    fmt.Println(p0, p2, p2)
                    if (1 >= 32767) {
                        if false {
                            fmt.Println(p0)
                        } else {
                            var v10 int16 = (p1 << ((255) & 63))
                            switch "a" {
                            case "hello":
                                if true {
                                    fmt.Println(p2, v8, p1)
                                    var v11 float32 = (0.1 / (0.1 / (0.1 - 3.14159)))
                                    var v12 uint = 1
                                    v8 = p0
                                    p0 = "\n\t"
                                    _ = v11
                                    _ = v12
                                }
                            default:
                                p1 = (p1 ^ 42)
                            }
                            _ = v10
                        }
                    } else {
                        p0 = (p0 + "\n\t")
                        v13 := GenS1{F0: 4294967295, F1: 0, F2: ((255 == 832) || (v6 || (-1 <= 559)))}
                        v14 := GenS2{F0: (-(p1))}
                        var v15 uint16 = 65535
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    if v6 {
                        v16 := []int{fi7, fi7, (fi9 + fi7), (fi7 + -2147483648)}
                        _ = v16
                    } else {
                        switch v8 {
                        case "":
                            v8 = p0
                        case "\n\t":
                            var v17 uint64 = 42
                            _ = v17
                        case "a":
                            p0 = v8
                        }
                        if v6 {
                            switch -1 {
                            case 127:
                                if v6 {
                                    v18 := []int{}
                                    _ = v18
                                } else {
                                    fmt.Println(v8, p1, p2)
                                    p1 = p1
                                    v6 = ((3.14159 == 3.14159) || p2)
                                }
                            default:
                                p2 = ((^(fi7)) == 127)
                            }
                        } else {
                            var v19 int64 = -1
                            _ = v19
                        }
                        if (p2 || (!p2)) {
                            p0 = v8
                            v20 := GenS3{F0: 123, F1: 2}
                            v21 := GenS4{F0: -676}
                            _ = v20
                            _ = v21
                        } else {
                            var v22 bool = false
                            _ = v22
                        }
                    }
                    var v23 int16 = -688
                    var v24 uint32 = 2147483647
                    v25 := []int{}
                    _ = v23
                    _ = v24
                    _ = v25
                case "hello":
                    switch (fi9 + fi7) {
                    case 128:
                        switch (fi9 >> ((2) & 63)) {
                        case 0:
                            fmt.Println(p1, v8, fi9)
                        case -473:
                            var v26 bool = p2
                            _ = v26
                        }
                    case 4294967295:
                        var v27 float32 = 3.14159
                        _ = v27
                    default:
                        var v28 uint16 = 815
                        _ = v28
                    }
                case "foo bar":
                    if p2 {
                        var v29 uint8 = 127
                        _ = v29
                    }
                }
                _ = fi9
            }
        case "hello":
            v30 := GenS3{F0: -128, F1: 78}
            _ = v30
        }
        _ = fi7
        _ = v8
    }
    p0 = p0
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1000
}

func main() {
    var v1 int32 = 42
    var v2 uint8 = 19
    var v3 uint32 = 65535
    var v4 int64 = -1
    var v5 float32 = 1.5
    v3 = v3
    v31 := GenS3{F0: 52, F1: uint(v2)}
    fmt.Println(v5, v2, v5)
    for fi32 := 0; fi32 < 4; fi32++ {
        var v33 uint8 = 1
        v34 := []int{}
        v1 -= (v1 >> ((0) & 63))
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := []int{2147483647, -511}
    v36 := GenS1{F0: 1, F1: int64(v3), F2: ((v5 + 1.5) <= (((v5 + (float32(0.0) - v5)) - (v5 / (float32(((1.5 + 3.14159) + 3.14159)) - (float32(-2.25) - v5)))) - 1e10))}
    v37 := GenS1{F0: 1, F1: v4, F2: (false && (true || true))}
    switch 42 {
    case 127:
        for fi38 := 0; fi38 < 5; fi38++ {
            var v39 float32 = ((((v5 + v5) - 0.0) * v5) - 1e10)
            v39 = (((0.0 + float32(3.14159)) - 0.1) / 1.0)
            if false {
                var v40 int = fi38
                _ = v40
            } else {
                for fi41 := 0; fi41 < 16; fi41++ {
                    v42 := GenS3{F0: 42, F1: 42}
                    _ = fi41
                    _ = v42
                }
            }
            _ = fi38
            _ = v39
        }
    case 128:
        for fi43 := 0; fi43 < 19; fi43++ {
            var v44 bool = false
            _ = fi43
            _ = v44
        }
    }
    fmt.Println(v4, v2, v1)
    fmt.Println(v5)
    v45 := GenS3{F0: 127, F1: 255}
    for fi46 := 0; fi46 < 12; fi46++ {
        var v47 int32 = v1
        _ = fi46
        _ = v47
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        if true {
            for fi49 := 0; fi49 < 9; fi49++ {
                if true {
                    switch (("\n\t" + ("\n\t" + (((("\n\t" + "foo bar") + ((("foo bar" + "foo bar") + ((("foo bar" + "foo bar") + "\n\t") + "hello")) + "")) + (("a" + "\n\t") + ("a" + ("" + "a")))) + "\n\t"))) + (("a" + ("" + "foo bar")) + "a")) {
                    case "\n\t":
                        if false {
                            v50 := GenS5{F0: v2, F1: ((("" + "") + "\n\t") + ("foo bar" + ""))}
                            if ((1.5 / ((float64(v5) / 3.14159) / float64(1.5))) == 0.1) {
                                fmt.Println(fi49)
                            } else {
                                _ = len(v35)
                                var v51 int = (fi48 * fi49)
                                var v52 float64 = (0.0 + (0.0 + float64(v5)))
                                var v53 uint = 4294967295
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            _ = v50
                        } else {
                            v35 = append(v35, fi49)
                            v5 += 0.0
                        }
                    default:
                        var v54 uint = 975
                        _ = v54
                    }
                    var v55 int8 = -1
                    _ = v55
                }
                v56 := GenS6{F0: (v2 * 128)}
                var v57 int32 = 42
                if true {
                    v5 = v5
                    v1 += v57
                    var v58 int16 = -51
                    _ = v58
                }
                v3 -= (v3 / (((v3 ^ v3)) | 1))
                var v59 bool = ((0.0 / (float64(0.1) + float64(v5))) > 0.0)
                _ = fi49
                _ = v56
                _ = v57
                _ = v59
            }
            if true {
                if true {
                    _ = len(v35)
                    v60 := GenS7{F0: 0}
                    _ = v60
                }
                fmt.Println(v4, v2, v2)
            } else {
                v4 += int64(v3)
                if len(v35) >= 2 {
                    v35 = v35[1:len(v35)]
                }
            }
            for fi61 := 0; fi61 < 16; fi61++ {
                var v62 uint8 = (v2 % ((v2) | 1))
                v63 := GenS8{F0: 42}
                v64 := GenS9{F0: 65535, F1: (v1 % ((v1) | 1))}
                _ = fi61
                _ = v62
                _ = v63
                _ = v64
            }
        }
        for fi65 := 0; fi65 < 17; fi65++ {
            fmt.Println(v3, v4)
            fmt.Println(v3, v1)
            _ = fi65
        }
        var v66 bool = (42 != 18446744073709551615)
        _ = fi48
        _ = v66
    }
    var v67 uint64 = 333
    for fi68 := 0; fi68 < 12; fi68++ {
        fmt.Println(v3)
        fmt.Println(v3, v5)
        _ = fi68
    }
    for fi69 := 0; fi69 < 19; fi69++ {
        var v70 uint16 = 1
        fmt.Println(v67)
        switch 248 {
        case 0:
            switch ("foo bar" + "foo bar") {
            case "":
                v71 := GenS4{F0: int64(v67)}
                _ = v71
            case "\n\t":
                v5 -= v5
            case "foo bar":
                var v72 float32 = 1e10
                _ = v72
            default:
                if true {
                    for fi73 := 0; fi73 < 17; fi73++ {
                        v5 = (v5 - v5)
                        var v74 int16 = int16(v70)
                        _ = fi73
                        _ = v74
                    }
                    for fi75 := 0; fi75 < 11; fi75++ {
                        v76 := GenS4{F0: (v4 << ((65535) & 63))}
                        if false {
                            switch "hello" {
                            case "":
                                if len(v35) >= 2 {
                                    v35 = v35[1:len(v35)]
                                }
                            case "a":
                                _ = len(v35)
                            default:
                                switch fi75 {
                                case -2147483648:
                                    v1 -= (v1 / (((v1 | v1)) | 1))
                                case 32768:
                                    v70 = 65535
                                case 405:
                                    fmt.Println(v5, v5, v3)
                                }
                            }
                            v5 += (float32(1.5) + (v5 * v5))
                            v3 -= (v3 >> ((1) & 63))
                        }
                        fmt.Println(v5, v4, v4)
                        _ = fi75
                        _ = v76
                    }
                    var v77 int = fi69
                    _ = v77
                }
            }
        case 32767:
            v70 -= v70
        case 1:
            fmt.Println(v4, v67)
        }
        _ = fi69
        _ = v70
    }
    fmt.Println(v1, v67, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v45
    _ = v5
    _ = v67
}
