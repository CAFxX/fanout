package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 uint16; F1 string; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint16; F1 int8; F2 int }
type GenS8 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint32; F1 int16; F2 int16 }

func main() {
    v1 := []uint32{42, 42}
    v2 := make(map[uint32]int)
    var v3 string = "\n\t"
    v4 := make(map[int]bool)
    v5 := GenS1{F0: (v3 + ("foo bar" + v3)), F1: 1}
    var v6 string = "a"
    var v7 int32 = -867
    v8 := GenS1{F0: v3, F1: 0}
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            switch v3 {
            case "a":
                if false {
                    fmt.Println(fi10)
                } else {
                    v11 := GenS1{F0: (v6 + v3), F1: int64(v7)}
                    v3 = "hello"
                    v7 += (^(v7))
                    _ = v11
                }
            }
            v12 := v4[fi10]
            _ = fi10
            _ = v12
        }
        _ = fi9
    }
    v3 = v6
    if (false && false) {
        switch v3 {
        case "\n\t":
            var v13 float32 = 1e10
            _ = v13
        case "hello":
            if true {
                fmt.Println(v7, v3, v6)
                v14 := GenS1{F0: v6, F1: -1}
                _ = v14
            } else {
                var v15 uint = 255
                _ = v15
            }
        }
        switch -823 {
        case 919:
            v16 := GenS1{F0: "hello", F1: -706}
            _ = v16
        default:
            _ = len(v1)
        }
        fmt.Println(v6, v7)
    }
    if (65535 > 255) {
        delete(v4, 701)
    } else {
        var v17 float32 = ((1.5 - (-2.25 + 3.14159)) * 3.14159)
        _ = len(v4)
        _ = v17
    }
    v18 := GenS2{F0: 873, F1: 784}
    v3 = ((v3 + ((v6 + "\n\t") + "hello")) + (v6 + ("" + v3)))
    if len(v1) > 0 {
        v1[0] = 1
    }
    if false {
        v3 = v3
        v19 := GenS1{F0: v3, F1: 1}
        _ = v19
    }
    if (("\n\t" + ("a" + v6)) > v3) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v6)
        if (false && (false && (-98 != 4))) {
            switch v6 {
            case "hello":
                switch 0 {
                case 42:
                    if (-1 <= 100) {
                        fmt.Println(v6)
                        v20 := GenS2{F0: 11, F1: 0}
                        _ = v20
                    } else {
                        switch 32768 {
                        case 2147483647:
                            v6 = v3
                        case -233:
                            v6 = v3
                        default:
                            v21 := GenS2{F0: 65535, F1: 261}
                            _ = v21
                        }
                        var v22 int = -2147483648
                        v3 = v3
                        _ = v22
                    }
                default:
                    _ = len(v1)
                }
            case "a":
                switch (("hello" + v3) + v6) {
                case "hello":
                    switch v3 {
                    case "\n\t":
                        if true {
                            for fi23 := 0; fi23 < 14; fi23++ {
                                v24 := v4[(^(fi23))]
                                for fi25 := 0; fi25 < 7; fi25++ {
                                    var v26 int64 = 0
                                    _ = fi25
                                    _ = v26
                                }
                                _ = fi23
                                _ = v24
                            }
                            var v27 int32 = (v7 + v7)
                            switch v6 {
                            case "hello":
                                var v28 uint64 = 673
                                _ = v28
                            case "":
                                v27 += v27
                            default:
                                v27 -= v7
                            }
                            _ = v27
                        }
                    case "a":
                        var v29 bool = false
                        _ = v29
                    }
                case "foo bar":
                    for fi30 := 0; fi30 < 12; fi30++ {
                        var v31 int16 = 0
                        v32 := GenS3{F0: (v7 % ((1) | 1))}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
            default:
                v3 = v6
            }
            if false {
                fmt.Println(v7, v3)
                fmt.Println(v3, v6)
            } else {
                switch "foo bar" {
                case "foo bar":
                    v33 := GenS4{F0: -59, F1: 65535}
                    _ = v33
                case "\n\t":
                    var v34 uint32 = 724
                    _ = v34
                case "a":
                    switch 610 {
                    case -2147483648:
                        fmt.Println(v3)
                    case -340:
                        if true {
                            v6 = (v6 + ((v6 + v6) + v6))
                        }
                    default:
                        var v35 uint16 = 2
                        _ = v35
                    }
                }
                fmt.Println(v7, v7)
            }
            var v36 uint = 4294967295
            _ = v36
        } else {
            for fi37 := 0; fi37 < 12; fi37++ {
                switch fi37 {
                case 128:
                    if (32767 != 1) {
                        switch 328 {
                        case 322:
                            switch v3 {
                            case "hello":
                                var v38 uint8 = 2
                                _ = v38
                            case "a":
                                for fi39 := 0; fi39 < 16; fi39++ {
                                    v40 := GenS4{F0: 127, F1: 173}
                                    v7 += (v7 % ((-855) | 1))
                                    if true {
                                        fmt.Println(fi37, v3, fi39)
                                        switch ((v3 + v6) + ("" + "\n\t")) {
                                        case "":
                                            var v41 bool = (1 != 2)
                                            _ = v41
                                        case "hello":
                                            var v42 uint8 = 127
                                            _ = v42
                                        case "\n\t":
                                            v6 = (v6 + ("hello" + v3))
                                        }
                                        v3 = "\n\t"
                                    }
                                    _ = fi39
                                    _ = v40
                                }
                            default:
                                switch (fi37 | fi37) {
                                case -776:
                                    fmt.Println(v3, fi37)
                                case 925:
                                    var v43 uint64 = 18446744073709551615
                                    _ = v43
                                case 127:
                                    v6 = v6
                                default:
                                    for fi44 := 0; fi44 < 15; fi44++ {
                                        if false {
                                            var v45 float64 = 3.14159
                                            if ((v3 + "\n\t") < ("\n\t" + (v6 + (("hello" + (v6 + v3)) + v3)))) {
                                                var v46 uint = 65535
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v46
                                            } else {
                                                v3 = v6
                                                v3 = v3
                                            }
                                            v47 := GenS5{F0: 255, F1: ((("" + (v3 + "foo bar")) + v6) + v3), F2: 42}
                                            _ = v45
                                            _ = v47
                                        } else {
                                            v48 := GenS1{F0: v6, F1: 821}
                                            v49 := GenS4{F0: -118, F1: 0}
                                            if (((-(fi44)) >= (fi44 / ((128) | 1))) && (!(false || ("hello" != "")))) {
                                                switch fi44 {
                                                case 820:
                                                    v50 := GenS5{F0: 228, F1: "\n\t", F2: 65535}
                                                    delete(v2, 751)
                                                    _ = v50
                                                case 128:
                                                    var v51 int8 = 100
                                                    _ = v51
                                                case -9223372036854775808:
                                                    v6 = (v3 + v6)
                                                default:
                                                    var v52 int16 = -83
                                                    _ = v52
                                                }
                                                if true {
                                                    if (v6 != "foo bar") {
                                                        v1 = append(v1, 196)
                                                    } else {
                                                        if true {
                                                            v53 := GenS6{F0: 4294967295, F1: 0, F2: 65535}
                                                            _ = len(v1)
                                                            _ = v53
                                                        } else {
                                                            switch v3 {
                                                            case "":
                                                                fmt.Println(v6, fi37)
                                                            }
                                                            fmt.Println(v7, v7)
                                                            var v54 uint = 607
                                                            _ = v54
                                                        }
                                                        var v55 uint = 70
                                                        _ = v55
                                                    }
                                                }
                                                var v56 float32 = float32((3.14159 - 0.0))
                                                fmt.Println(v7, v6)
                                                var v57 int32 = v7
                                                _ = v56
                                                _ = v57
                                            } else {
                                                v6 = "\n\t"
                                                v58 := GenS7{F0: 255, F1: 0, F2: int(v7)}
                                                _ = v58
                                            }
                                            v59 := GenS8{F0: 2, F1: -32768, F2: 927}
                                            _ = v48
                                            _ = v49
                                            _ = v59
                                        }
                                        _ = fi44
                                    }
                                }
                            }
                        case -784:
                            var v60 uint16 = 179
                            _ = v60
                        case 299:
                            var v61 int8 = -67
                            _ = v61
                        }
                        switch fi37 {
                        case 42:
                            for fi62 := 0; fi62 < 17; fi62++ {
                                fmt.Println(fi37, fi37)
                                v7 = (v7 * (v7 % ((v7) | 1)))
                                _ = fi62
                            }
                        case 127:
                            var v63 float64 = 0.1
                            _ = v63
                        }
                        v7 = v7
                    } else {
                        switch ((("" + (v6 + "hello")) + "a") + v6) {
                        case "":
                            v3 = ("\n\t" + "hello")
                        case "a":
                            fmt.Println(v6)
                        case "hello":
                            var v64 int64 = 1
                            _ = v64
                        default:
                            for fi65 := 0; fi65 < 4; fi65++ {
                                v4[fi65] = false
                                _ = fi65
                            }
                        }
                        v66 := GenS9{F0: 255}
                        _ = v66
                    }
                case -9223372036854775808:
                    switch "\n\t" {
                    case "\n\t":
                        v67 := GenS9{F0: 32767}
                        _ = v67
                    case "foo bar":
                        v3 = v3
                    case "":
                        delete(v4, fi37)
                    }
                case -687:
                    v7 -= v7
                default:
                    v3 = v3
                }
                v68 := GenS10{F0: 65535, F1: -32768, F2: 172}
                _ = fi37
                _ = v68
            }
        }
    }
    fmt.Println(v3, v7)
    var v69 bool = false
    var v70 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v7
    _ = v70
    _ = v8
}
