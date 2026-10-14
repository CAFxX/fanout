package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int; F1 int64; F2 int }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int16; F1 int8; F2 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{32768, 0, 174}
    var v3 int = 0
    v4 := make(map[string]bool)
    var v5 bool = (42 <= 32767)
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := v1[1]
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS1{F0: "", F1: v5}
            _ = len(v1)
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 uint64 = 254
    switch "" {
    case "\n\t":
        var v11 float32 = float32(0.1)
        _ = v11
    }
    if v5 {
        v12 := GenS2{F0: -1, F1: 0, F2: 4294967295}
        v5 = (!(!(114 >= 18446744073709551615)))
        _ = v12
    }
    v10 += (v10 << ((50) & 63))
    switch ("a" + "\n\t") {
    case "\n\t":
        var v13 float64 = 3.14159
        _ = v13
    case "foo bar":
        v14 := GenS2{F0: 127, F1: 172, F2: 836}
        _ = v14
    }
    switch v3 {
    case -1:
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 uint = 127
            v16 += 2
            v10 += (v10 + v10)
            _ = fi15
            _ = v16
        }
    case 483:
        fmt.Println(v3)
    }
    var v17 float64 = (((0.0 + (1e10 / 0.1)) - (((0.1 * (3.14159 / 3.14159)) + ((float64(1e10) / 1.0) - (3.14159 * -2.25))) - 3.14159)) + 0.0)
    if (v17 == v17) {
        v18 := GenS3{F0: v3, F1: -840, F2: v3}
        fmt.Println(v10)
        for fi19 := 0; fi19 < 14; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                v10 += 2
                var v21 uint64 = (v10 - 1)
                if true {
                    if v5 {
                        switch "a" {
                        case "foo bar":
                            v22 := GenS1{F0: "hello", F1: (-813 != 9223372036854775807)}
                            _ = v22
                        default:
                            var v23 uint64 = (v21 << ((862) & 63))
                            _ = v23
                        }
                        var v24 int = -887
                        _ = v24
                    }
                    v25 := GenS2{F0: 0, F1: 255, F2: 2147483647}
                    _ = v25
                }
                _ = fi20
                _ = v21
            }
            switch "" {
            case "hello":
                var v26 uint = 0
                _ = v26
            case "\n\t":
                v10 += (v10 | (v10 << ((42) & 63)))
            case "foo bar":
                v17 = (float64(1e10) / v17)
            }
            var v27 float64 = float64((-2.25 / float32(v17)))
            _ = fi19
            _ = v27
        }
        _ = v18
    } else {
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 string = "a"
            v10 += 1
            v30 := GenS4{F0: v10, F1: 1}
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    var v31 uint8 = 42
    delete(v4, "a")
    fmt.Println(v31, v5)
    switch (v3 >> ((232) & 63)) {
    case 554:
        v5 = (42 > 61)
    case -9223372036854775808:
        for fi32 := 0; fi32 < 14; fi32++ {
            switch 128 {
            case 42:
                fmt.Println(v10, v10)
            case 127:
                if (v10 > 0) {
                    v33 := GenS3{F0: (fi32 % (((v3 ^ -9223372036854775808)) | 1)), F1: -9223372036854775808, F2: (fi32 << ((399) & 63))}
                    for fi34 := 0; fi34 < 5; fi34++ {
                        _ = len(v1)
                        v35 := GenS3{F0: 587, F1: 0, F2: fi34}
                        fmt.Println(v17)
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                } else {
                    var v36 string = ""
                    fmt.Println(v10)
                    _ = v36
                }
            case 224:
                switch "" {
                case "\n\t":
                    switch ("\n\t" + "") {
                    case "foo bar":
                        v37 := GenS5{F0: 65535, F1: 0}
                        _ = v37
                    default:
                        var v38 uint64 = (^(v10))
                        _ = v38
                    }
                default:
                    fmt.Println(v31)
                }
            default:
                v17 -= (float64((((0.0 / 1.0) * ((1.5 / ((0.1 / 3.14159) - (1e10 * (-2.25 + float32((float64((1.5 + 0.1)) - ((v17 + v17) - (3.14159 * 1e10)))))))) * 0.0)) / 1.5)) - (v17 - v17))
            }
            _ = fi32
        }
    case -2147483648:
        for fi39 := 0; fi39 < 16; fi39++ {
            switch "hello" {
            case "\n\t":
                var v40 bool = v5
                _ = v40
            default:
                fmt.Println(v10, v17, fi39)
            }
            switch fi39 {
            case 127:
                fmt.Println(v5, v31, fi39)
            }
            switch "" {
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "hello":
                var v41 int = (v3 * v3)
                _ = v41
            case "\n\t":
                fmt.Println(fi39)
            }
            _ = fi39
        }
    default:
        fmt.Println(v17)
    }
    var v42 int16 = 42
    switch v3 {
    case -9223372036854775808:
        for fi43 := 0; fi43 < 15; fi43++ {
            if false {
                v17 += v17
                v44 := GenS6{F0: v5, F1: uint16(v31)}
                _ = v44
            } else {
                var v45 uint16 = 542
                for fi46 := 0; fi46 < 13; fi46++ {
                    if (-2.25 != (float32((1.5 * v17)) - (1e10 * 1e10))) {
                        var v47 float32 = ((0.1 + 0.1) - (0.1 - 1.5))
                        _ = v47
                    } else {
                        var v48 uint32 = 708
                        switch fi46 {
                        case -301:
                            var v49 int64 = 42
                            _ = v49
                        case -966:
                            if v5 {
                                v2 = append(v2, (fi46 + 292))
                                var v50 int64 = -1
                                delete(v1, fi46)
                                _ = v50
                            } else {
                                fmt.Println(v31, v31)
                                if (-1 != 1) {
                                    v51 := GenS6{F0: (4294967295 == 2), F1: 2}
                                    v42 = 686
                                    v52 := GenS7{F0: (v42 & v42), F1: 51, F2: (v48 & uint32(v45))}
                                    _ = v51
                                    _ = v52
                                } else {
                                    fmt.Println(v3, v17)
                                    v10 += v10
                                    v53 := GenS8{F0: 2147483647}
                                    _ = v53
                                }
                                if v5 {
                                    switch (fi46 & (v3 & (fi43 ^ int(v42)))) {
                                    case 2147483647:
                                        var v54 int32 = 633
                                        _ = v54
                                    case -2147483648:
                                        v31 = v31
                                    case 128:
                                        fmt.Println(v42, v10)
                                    default:
                                        switch "\n\t" {
                                        case "\n\t":
                                            if (127 != 100) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v10 = (v10 / ((v10) | 1))
                                                case "":
                                                    var v55 string = (("\n\t" + ("a" + "hello")) + (("\n\t" + (("\n\t" + (("\n\t" + "\n\t") + "hello")) + "\n\t")) + ("" + "hello")))
                                                    _ = v55
                                                case "a":
                                                    var v56 int64 = -561
                                                    _ = v56
                                                }
                                                v31 += (v31 >> ((82) & 63))
                                            }
                                        case "a":
                                            switch 2147483647 {
                                            case -9223372036854775808:
                                                var v57 uint = 2
                                                _ = v57
                                            case 0:
                                                v58 := GenS9{F0: 2}
                                                _ = v58
                                            default:
                                                v59 := GenS8{F0: 42}
                                                fmt.Println(fi43)
                                                v10 -= 884
                                                _ = v59
                                            }
                                        default:
                                            var v60 string = ""
                                            _ = v60
                                        }
                                    }
                                    fmt.Println(v5)
                                }
                            }
                        case -2147483648:
                            fmt.Println(v42, v17)
                        default:
                            var v61 int8 = 1
                            _ = v61
                        }
                        v62 := GenS5{F0: (v48 << ((255) & 63)), F1: 222}
                        _ = v48
                        _ = v62
                    }
                    var v63 uint32 = uint32(v31)
                    _ = fi46
                    _ = v63
                }
                v31 += 10
                _ = v45
            }
            v5 = (!v5)
            for fi64 := 0; fi64 < 18; fi64++ {
                var v65 uint = 239
                switch "\n\t" {
                case "":
                    v66 := v1[v3]
                    _ = v66
                default:
                    v31 = v31
                }
                _ = fi64
                _ = v65
            }
            _ = fi43
        }
    case -2147483648:
        var v67 uint8 = (v31 & (v31 + v31))
        _ = v67
    case 128:
        if ((v42 - (v42 * (v42 + v42))) < v42) {
            fmt.Println(v42)
            v68 := v1[(v3 + (^(v3)))]
            v17 -= (v17 / v17)
            _ = v68
        } else {
            v3 = (-(v3))
            v69 := GenS10{F0: 2, F1: v10, F2: 65535}
            _ = v69
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v5
}
