package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint; F2 int }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int64; F1 uint; F2 uint16 }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 string; F1 string; F2 uint8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint8; F1 uint64; F2 int32 }

func main() {
    var v1 int16 = -481
    var v2 uint = 714
    var v3 float32 = -2.25
    for fi4 := 0; fi4 < 15; fi4++ {
        if (true && true) {
            v5 := GenS1{F0: -1, F1: (v1 >> ((v2) & 63)), F2: "hello"}
            v2 = (^(v2))
            _ = v5
        }
        var v6 int32 = 42
        _ = fi4
        _ = v6
    }
    fmt.Println(v2, v3)
    v2 = v2
    var v7 int8 = -1
    v8 := GenS2{F0: (v2 % ((444) | 1)), F1: 758}
    v9 := GenS2{F0: (v2 - v2), F1: 9223372036854775807}
    if (!true) {
        v2 += 0
    } else {
        v10 := GenS1{F0: -46, F1: (-(v1)), F2: "foo bar"}
        if (1 > 1) {
            v11 := []int{215, 42, 128}
            var v12 float32 = v3
            fmt.Println(v2, v7)
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 18; fi13++ {
                for fi14 := 0; fi14 < 7; fi14++ {
                    v15 := GenS2{F0: v2, F1: 0}
                    if (1e10 > float32(3.14159)) {
                        v16 := GenS2{F0: (^(v2)), F1: 0}
                        v17 := GenS3{F0: 437}
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(v7)
                        v18 := []int{(fi13 ^ -764)}
                        _ = v18
                    }
                    v19 := GenS1{F0: 0, F1: v1, F2: "hello"}
                    _ = fi14
                    _ = v15
                    _ = v19
                }
                v20 := GenS1{F0: -2147483648, F1: (v1 * (-(v1))), F2: "hello"}
                switch "hello" {
                case "a":
                    for fi21 := 0; fi21 < 19; fi21++ {
                        fmt.Println(v2)
                        _ = fi21
                    }
                default:
                    for fi22 := 0; fi22 < 5; fi22++ {
                        if true {
                            if true {
                                switch 496 {
                                case 32768:
                                    v23 := GenS4{F0: (v2 - (v2 << ((v2) & 63))), F1: (v2 + (v2 + v2)), F2: (fi13 >> ((v2) & 63))}
                                    _ = v23
                                case -2147483648:
                                    v24 := GenS4{F0: v2, F1: (v2 / ((65535) | 1)), F2: -578}
                                    _ = v24
                                case 2147483647:
                                    v25 := []int{(fi13 & (^(fi13))), (fi13 % ((fi22) | 1))}
                                    _ = v25
                                default:
                                    v26 := GenS5{F0: 245, F1: v2}
                                    _ = v26
                                }
                                v7 -= (v7 >> ((v2) & 63))
                                v7 = (v7 * v7)
                                var v27 uint64 = 908
                                var v28 string = "\n\t"
                                _ = v27
                                _ = v28
                            }
                            fmt.Println(v3)
                        } else {
                            if (42 > 0) {
                                v29 := []int{(fi22 | fi13), fi22, fi13, (fi13 / (((fi13 | (^(fi13)))) | 1))}
                                v29 = append(v29, fi22)
                                if len(v29) > 0 {
                                    v29[0] = -1
                                }
                                _ = v29
                            } else {
                                var v30 float64 = 1e10
                                v3 = 1.5
                                if true {
                                    var v31 uint64 = 56
                                    switch "hello" {
                                    case "a":
                                        v31 = v31
                                    case "foo bar":
                                        var v32 uint64 = v31
                                        _ = v32
                                    case "\n\t":
                                        v33 := GenS6{F0: -417, F1: (v2 * (v2 >> ((v2) & 63))), F2: 16}
                                        _ = v33
                                    }
                                    var v34 uint8 = 127
                                    _ = v31
                                    _ = v34
                                }
                                _ = v30
                            }
                            var v35 uint64 = 501
                            _ = v35
                        }
                        switch 2147483647 {
                        case 127:
                            v1 += -1
                        case 128:
                            switch ((("foo bar" + "hello") + "a") + "a") {
                            case "foo bar":
                                fmt.Println(v1)
                            case "\n\t":
                                switch (fi22 + int(v7)) {
                                case -393:
                                    var v36 string = "a"
                                    _ = v36
                                default:
                                    fmt.Println(v3)
                                }
                            case "a":
                                fmt.Println(v1, v3, v3)
                            }
                        case 42:
                            switch 128 {
                            case 128:
                                var v37 int64 = 1
                                _ = v37
                            default:
                                v3 += ((v3 / v3) / (float32((3.14159 * ((1.5 / -2.25) + 3.14159))) / 1e10))
                            }
                        default:
                            switch fi13 {
                            case 42:
                                v38 := GenS3{F0: 181}
                                _ = v38
                            case -1:
                                fmt.Println(v2)
                                if true {
                                    if (!true) {
                                        fmt.Println(v1, v1)
                                    } else {
                                        v2 = (^(v2))
                                    }
                                }
                                v2 = (v2 << (((v2 << ((255) & 63))) & 63))
                            case 1:
                                var v39 int64 = 197
                                _ = v39
                            default:
                                var v40 int8 = v7
                                _ = v40
                            }
                        }
                        var v41 float32 = v3
                        _ = fi22
                        _ = v41
                    }
                }
                _ = fi13
                _ = v20
            }
            v42 := GenS7{F0: (v7 << (((^(v2))) & 63)), F1: v1}
            v43 := GenS7{F0: 1, F1: (v1 & v1)}
            _ = v42
            _ = v43
        }
        var v44 uint64 = 181
        _ = v10
        _ = v44
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        v1 += (v1 - v1)
        switch fi45 {
        case 32767:
            fmt.Println(fi45, v3, fi45)
        default:
            if (42 != 2) {
                v46 := GenS6{F0: -9223372036854775808, F1: (v2 - 18446744073709551615), F2: 959}
                var v47 uint16 = 880
                _ = v46
                _ = v47
            }
        }
        _ = fi45
    }
    var v48 float64 = -2.25
    v49 := []int{9223372036854775807, 4294967295, -338}
    v7 += (v7 ^ (^(v7)))
    v49 = append(v49, 851)
    v48 += ((v48 / 1.0) - v48)
    var v50 uint64 = 759
    for fi51 := 0; fi51 < 20; fi51++ {
        if true {
            if len(v49) >= 2 {
                v49 = v49[1:len(v49)]
            }
            v50 = 4294967295
            for fi52 := 0; fi52 < 10; fi52++ {
                v48 += v48
                switch fi51 {
                case -70:
                    v49 = append(v49, 1)
                default:
                    fmt.Println(v3, fi52, fi51)
                }
                v53 := GenS8{F0: "foo bar", F1: ("hello" + ""), F2: 127}
                if ((v48 + 1.5) == (((0.0 + float64(float32(float64(1.5)))) / (v48 - 0.0)) * ((1e10 + v48) + v48))) {
                    v54 := GenS9{F0: (v1 << ((v2) & 63))}
                    var v55 int16 = v1
                    _ = v54
                    _ = v55
                }
                if ((!false) && (((v7 >> ((v2) & 63)) == (-(v7))) || true)) {
                    if true {
                        switch "hello" {
                        case "hello":
                            switch fi51 {
                            case 649:
                                v3 += 3.14159
                            }
                        }
                        if len(v49) >= 2 {
                            v49 = v49[1:len(v49)]
                        }
                    } else {
                        fmt.Println(fi51)
                        if false {
                            v2 = v2
                        } else {
                            v1 = v1
                            var v56 int8 = v7
                            _ = v56
                        }
                        switch "\n\t" {
                        case "hello":
                            _ = len(v49)
                        case "foo bar":
                            if ((!false) || false) {
                                if len(v49) > 0 {
                                    v49[0] = fi51
                                }
                            }
                        case "\n\t":
                            var v57 uint = (v2 << ((v2) & 63))
                            _ = v57
                        default:
                            if (float32(v48) >= v3) {
                                if false {
                                    if true {
                                        v2 = uint(v50)
                                        var v58 uint16 = 635
                                        if len(v49) >= 2 {
                                            v49 = v49[1:len(v49)]
                                        }
                                        _ = v58
                                    }
                                    v2 += (^(v2))
                                }
                                v7 = v7
                                var v59 int = int(v1)
                                _ = v59
                            } else {
                                if len(v49) >= 2 {
                                    v49 = v49[1:len(v49)]
                                }
                            }
                        }
                        v49 = append(v49, (fi52 ^ (^(fi51))))
                    }
                    var v60 int16 = v1
                    _ = v60
                }
                _ = fi52
                _ = v53
            }
        } else {
            v49 = append(v49, (fi51 << (((v2 >> ((v2) & 63))) & 63)))
        }
        var v61 int = -317
        v62 := GenS10{F0: 127, F1: v50, F2: -2147483648}
        _ = fi51
        _ = v61
        _ = v62
    }
    var v63 uint8 = 2
    var v64 uint16 = uint16(v63)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v50:", v50)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v50
    _ = v63
    _ = v64
    _ = v7
    _ = v8
    _ = v9
}
