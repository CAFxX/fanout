package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int32; F1 bool; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 1
    var v3 int = 0
    var v4 string = "\n\t"
    var v5 int16 = -971
    v6 := make(map[string]bool)
    var v7 float64 = ((1.5 / 0.1) + (1.5 * (0.1 - ((float64(-2.25) - 0.0) / (1e10 / 3.14159)))))
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 773}
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 2
    for fi11 := 0; fi11 < 5; fi11++ {
        switch v3 {
        case 32768:
            for fi12 := 0; fi12 < 9; fi12++ {
                if (65535 == 65535) {
                    var v13 float64 = -2.25
                    _ = v13
                } else {
                    v14 := v1[(v4 + v4)]
                    switch ("foo bar" + "") {
                    case "":
                        var v15 int32 = 42
                        _ = v15
                    }
                    _ = v14
                }
                if (v4 != v4) {
                    v4 = ("\n\t" + ((v4 + "") + "foo bar"))
                }
                if (127 <= 127) {
                    if (65535 < 23) {
                        v16 := GenS1{F0: 745}
                        v2 -= (-(v2))
                        _ = v16
                    } else {
                        switch (("" + (v4 + ("\n\t" + v4))) + ("" + ("a" + ""))) {
                        case "":
                            v2 = v2
                        case "\n\t":
                            var v17 int8 = 1
                            _ = v17
                        default:
                            var v18 uint32 = 65535
                            _ = v18
                        }
                        v7 = float64(((-2.25 - (-2.25 * -2.25)) * 1e10))
                    }
                    _ = len(v1)
                }
                if false {
                    _ = len(v6)
                }
                v7 = -2.25
                _ = fi12
            }
        case -9223372036854775808:
            v19 := GenS1{F0: 379}
            _ = v19
        default:
            if (!((0 >= 1) || false)) {
                fmt.Println(fi11, v10, v7)
            }
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            if true {
                fmt.Println(v5, fi20, v10)
            } else {
                var v21 uint16 = 32767
                _ = v21
            }
            var v22 uint = uint(v10)
            v22 = v22
            _ = fi20
            _ = v22
        }
        _ = fi11
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            if false {
                var v25 bool = true
                var v26 uint16 = 1
                if (!(42 <= 0)) {
                    v25 = false
                    if true {
                        var v27 string = v4
                        _ = v27
                    }
                    var v28 int = (fi23 + -1)
                    v6[("hello" + (("a" + "") + ("" + (v4 + ""))))] = (v5 > v5)
                    _ = v28
                }
                v29 := []int{32768, (v3 | v3)}
                v6[("hello" + "foo bar")] = (!(v10 <= v10))
                _ = v25
                _ = v26
                _ = v29
            } else {
                var v30 uint32 = 312
                var v31 int16 = (v5 ^ v5)
                _ = v30
                _ = v31
            }
            var v32 int64 = -719
            if (v5 < -451) {
                if (!true) {
                    var v33 float64 = float64(3.14159)
                    var v34 bool = true
                    v35 := GenS1{F0: 2}
                    _ = v33
                    _ = v34
                    _ = v35
                } else {
                    v36 := GenS2{F0: (^(v10)), F1: 100, F2: 127}
                    if ((("hello" + "a") + (v4 + ("\n\t" + ((("\n\t" + v4) + (v4 + v4)) + "\n\t")))) <= v4) {
                        v37 := []int{fi24, (^(fi24)), (fi24 - fi23), fi23}
                        if false {
                            v7 -= v7
                            if true {
                                v37 = append(v37, fi23)
                                v7 = (float64(0.0) / v7)
                                switch fi24 {
                                case 127:
                                    var v38 int = 2147483647
                                    _ = v38
                                case 458:
                                    fmt.Println(fi24, v4)
                                case -9223372036854775808:
                                    if len(v37) >= 2 {
                                        v37 = v37[1:len(v37)]
                                    }
                                }
                            } else {
                                v37 = append(v37, (fi24 & fi23))
                            }
                            if false {
                                var v39 int32 = 0
                                v32 = 481
                                _ = v39
                            }
                            v3 = 1
                            var v40 int32 = (^(v2))
                            _ = v40
                        } else {
                            if false {
                                if (true && false) {
                                    v37 = append(v37, (^(fi23)))
                                } else {
                                    v5 = v5
                                }
                                if false {
                                    v3 += v3
                                    if false {
                                        var v41 float64 = v7
                                        _ = v41
                                    } else {
                                        v42 := GenS3{F0: 1, F1: v5}
                                        v2 = v2
                                        _ = v42
                                    }
                                }
                            } else {
                                var v43 int32 = (^(v2))
                                v44 := GenS4{F0: v43, F1: "hello"}
                                var v45 int8 = -37
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            v46 := GenS1{F0: 32767}
                            _ = v46
                        }
                        fmt.Println(fi24, fi23, v2)
                        _ = v37
                    } else {
                        v2 -= 42
                    }
                    _ = v36
                }
                var v47 bool = false
                v47 = true
                fmt.Println(v3, v4)
                v48 := []int{fi23, fi23}
                _ = v47
                _ = v48
            }
            _ = fi24
            _ = v32
        }
        if (42 < 58) {
            switch (v4 + ((v4 + "hello") + "\n\t")) {
            case "":
                if (501 != 624) {
                    switch v4 {
                    case "a":
                        fmt.Println(v10, v3, v10)
                    case "":
                        if false {
                            switch v3 {
                            case 4294967295:
                                v49 := GenS3{F0: 62, F1: v5}
                                _ = v49
                            }
                            fmt.Println(v7, v2, v2)
                        } else {
                            switch fi23 {
                            case 0:
                                switch "hello" {
                                case "a":
                                    v10 = 0
                                }
                            case 394:
                                v10 = (v10 % ((v10) | 1))
                            }
                        }
                    case "hello":
                        v50 := GenS3{F0: 1, F1: (^(v5))}
                        _ = v50
                    }
                    for fi51 := 0; fi51 < 8; fi51++ {
                        v4 = v4
                        v3 += fi23
                        var v52 bool = (int8(v2) == 1)
                        switch (v3 >> ((1) & 63)) {
                        case -222:
                            if v52 {
                                var v53 uint32 = 2147483647
                                _ = v53
                            } else {
                                v3 -= 1
                            }
                        case 4294967295:
                            v54 := GenS3{F0: 42, F1: v5}
                            var v55 int32 = (-(v2))
                            _ = v54
                            _ = v55
                        case -1:
                            fmt.Println(v7, v3)
                        }
                        _ = fi51
                        _ = v52
                    }
                    v3 = v3
                }
            default:
                var v56 bool = true
                _ = v56
            }
            switch (v4 + ("a" + "\n\t")) {
            case "a":
                v10 = (^(v10))
            case "":
                if false {
                    switch v4 {
                    case "\n\t":
                        if false {
                            var v57 float32 = ((-2.25 - 1.5) * 1.5)
                            _ = v57
                        }
                    default:
                        var v58 int = -306
                        _ = v58
                    }
                } else {
                    var v59 string = v4
                    var v60 string = ""
                    if (!false) {
                        fmt.Println(v4)
                        if (true || true) {
                            v60 = (v4 + (v59 + (v4 + v60)))
                        }
                        v61 := []int{(-(v3)), (v3 * v3), -2147483648}
                        _ = v61
                    }
                    _ = v59
                    _ = v60
                }
            case "\n\t":
                v3 = -235
            default:
                switch (v4 + ("a" + v4)) {
                case "":
                    v10 = 793
                case "foo bar":
                    switch v3 {
                    case 1:
                        for fi62 := 0; fi62 < 5; fi62++ {
                            if false {
                                v3 += (fi62 << ((65535) & 63))
                                fmt.Println(v7, v3, v3)
                                if (42 <= fi62) {
                                    var v63 bool = (!(65535 < 1))
                                    var v64 uint8 = 42
                                    _ = v63
                                    _ = v64
                                }
                                if (v2 > v2) {
                                    v5 = v5
                                    v65 := []int{-908}
                                    _ = v65
                                }
                            } else {
                                v3 -= v3
                            }
                            v10 -= v10
                            delete(v1, "a")
                            v66 := []int{fi23, (v3 | fi62), 1, fi62}
                            if (!((true && (!((v3 << ((255) & 63)) > (fi23 >> ((1) & 63))))) && (42 <= 1))) {
                                if (v5 > (v5 >> ((18446744073709551615) & 63))) {
                                    v7 = v7
                                    _ = len(v66)
                                }
                            } else {
                                if true {
                                    v10 += 416
                                    switch v4 {
                                    case "foo bar":
                                        var v67 float32 = 1e10
                                        _ = v67
                                    case "\n\t":
                                        var v68 uint8 = 1
                                        _ = v68
                                    }
                                    fmt.Println(v2)
                                }
                            }
                            _ = fi62
                            _ = v66
                        }
                    case 127:
                        var v69 uint32 = 277
                        _ = v69
                    }
                case "hello":
                    if (42 > 4294967295) {
                        v70 := GenS1{F0: 1}
                        v7 -= v7
                        _ = v70
                    } else {
                        v4 = ("a" + ((v4 + v4) + "a"))
                        fmt.Println(fi23)
                    }
                default:
                    if true {
                        v10 -= v10
                    } else {
                        v4 = ("a" + (((v4 + "\n\t") + v4) + v4))
                        v71 := GenS1{F0: 255}
                        v4 = "hello"
                        _ = v71
                    }
                }
            }
        } else {
            for fi72 := 0; fi72 < 2; fi72++ {
                v73 := GenS5{F0: (v2 % ((v2) | 1)), F1: false, F2: -9223372036854775808}
                var v74 uint64 = 633
                _ = fi72
                _ = v73
                _ = v74
            }
            switch v4 {
            case "a":
                var v75 bool = false
                _ = v75
            case "\n\t":
                v5 = (v5 - (^(v5)))
            case "foo bar":
                v4 = ""
            }
            var v76 uint = 255
            _ = v76
        }
        _ = fi23
    }
    var v77 uint32 = 1
    var v78 bool = true
    if v78 {
        fmt.Println(v4, v77)
        if (v4 >= (v4 + v4)) {
            for fi79 := 0; fi79 < 17; fi79++ {
                var v80 int16 = v5
                switch "\n\t" {
                case "a":
                    for fi81 := 0; fi81 < 18; fi81++ {
                        var v82 uint8 = 138
                        fmt.Println(v77, v77)
                        v83 := GenS4{F0: (v2 * -1), F1: "a"}
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                }
                v84 := GenS5{F0: (v2 << ((255) & 63)), F1: (v78 && (0.1 < v7)), F2: int64(v2)}
                _ = fi79
                _ = v80
                _ = v84
            }
        }
    }
    v85 := GenS4{F0: v2, F1: v4}
    fmt.Println(v3, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v77
    _ = v78
    _ = v85
}
