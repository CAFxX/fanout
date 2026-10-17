package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 bool; F1 bool; F2 uint }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint64; F1 int8; F2 bool }
type GenS7 struct { F0 int64; F1 uint; F2 uint8 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS9 struct { F0 uint; F1 int8 }
type GenS10 struct { F0 bool; F1 uint32 }
type GenS11 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 0
    v3 := make(map[string]bool)
    var v4 int = -1
    v5 := make(map[string]bool)
    var v6 float32 = (0.0 + -2.25)
    v7 := GenS1{F0: true, F1: v2, F2: 149}
    fmt.Println(v2)
    v6 -= ((float32(float64(((1.5 + (0.1 * v6)) * v6))) * v6) * v6)
    fmt.Println(v2, v6, v6)
    v2 -= 1
    var v8 string = "a"
    var v9 uint16 = 2
    switch v8 {
    case "foo bar":
        for fi10 := 0; fi10 < 15; fi10++ {
            v8 = "foo bar"
            switch (fi10 / ((v4) | 1)) {
            case 32767:
                var v11 string = "\n\t"
                _ = v11
            }
            fmt.Println(v9, v4, fi10)
            _ = fi10
        }
    case "\n\t":
        v12 := []int{(-(v4)), v4, v4, (v4 & v4)}
        _ = v12
    case "":
        v13 := GenS1{F0: (false || false), F1: v2, F2: 353}
        _ = v13
    default:
        switch v4 {
        case 9223372036854775807:
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 5; fi15++ {
                    fmt.Println(v8, v8, v2)
                    _ = fi15
                }
                var v16 int16 = 0
                _ = fi14
                _ = v16
            }
        case 41:
            var v17 uint32 = 2
            _ = v17
        case 32768:
            v18 := GenS2{F0: 1}
            _ = v18
        default:
            var v19 int8 = -128
            _ = v19
        }
    }
    var v20 bool = true
    for fi21 := 0; fi21 < 3; fi21++ {
        switch fi21 {
        case -319:
            v22 := GenS2{F0: 411}
            _ = v22
        }
        v23 := GenS1{F0: false, F1: (v2 - (v2 & 0)), F2: 931}
        _ = fi21
        _ = v23
    }
    if (!v20) {
        fmt.Println(v4)
        switch v8 {
        case "a":
            v2 -= (v2 - v2)
        default:
            var v24 int64 = v2
            _ = v24
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 bool = (!v20)
            v27 := []int{9223372036854775807, (v4 - (^(v4))), (^(v4)), v4}
            v28 := GenS3{F0: -233, F1: v8}
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
    } else {
        switch 2147483647 {
        case -2147483648:
            if v20 {
                var v29 uint16 = 408
                v30 := []int{(v4 << ((883) & 63)), (v4 - 0), v4}
                v30 = append(v30, 610)
                _ = v29
                _ = v30
            }
        default:
            v20 = ((v8 + v8) > v8)
        }
        var v31 uint16 = (v9 & v9)
        switch (("a" + v8) + "") {
        case "a":
            for fi32 := 0; fi32 < 5; fi32++ {
                v31 += (v9 << ((782) & 63))
                _ = fi32
            }
        case "foo bar":
            var v33 int = (v4 % ((32767) | 1))
            _ = v33
        default:
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 int64 = v2
                var v36 uint32 = 51
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        _ = v31
    }
    v2 -= v2
    v9 -= v9
    if (42 <= 1) {
        var v37 int16 = 692
        v20 = v20
        _ = v37
    } else {
        for fi38 := 0; fi38 < 15; fi38++ {
            var v39 int64 = v2
            v40 := GenS3{F0: -292, F1: (("\n\t" + "foo bar") + "foo bar")}
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    v41 := GenS4{F0: v20, F1: false, F2: 1}
    fmt.Println(v20, v4, v6)
    if v20 {
        if (207 == 0) {
            switch "a" {
            case "hello":
                for fi42 := 0; fi42 < 1; fi42++ {
                    switch 4294967295 {
                    case 473:
                        v43 := GenS2{F0: 68}
                        _ = v43
                    case -1:
                        var v44 float32 = float32(3.14159)
                        _ = v44
                    default:
                        delete(v3, v8)
                    }
                    v45 := v1["hello"]
                    if (2 < 1) {
                        var v46 string = "hello"
                        switch v8 {
                        case "":
                            v8 = ("hello" + v45)
                        case "hello":
                            v47 := GenS4{F0: v20, F1: (216 > -698), F2: 757}
                            _ = v47
                        case "foo bar":
                            v48 := GenS4{F0: false, F1: v20, F2: 42}
                            _ = v48
                        default:
                            v2 += (v2 & (v2 | (v2 - (v2 + (v2 << ((255) & 63))))))
                        }
                        _ = len(v1)
                        _ = v46
                    }
                    _ = fi42
                    _ = v45
                }
            case "foo bar":
                switch ((v8 + v8) + v8) {
                case "a":
                    if true {
                        v2 += v2
                    } else {
                        v49 := GenS5{F0: 65535, F1: v4}
                        if (42 < 4294967295) {
                            if (true || (515 >= 1)) {
                                for fi50 := 0; fi50 < 6; fi50++ {
                                    switch "foo bar" {
                                    case "hello":
                                        for fi51 := 0; fi51 < 17; fi51++ {
                                            fmt.Println(v8, v2)
                                            _ = fi51
                                        }
                                    default:
                                        if (((!(!v20)) || (v20 || ((v8 + "a") < (v8 + "\n\t")))) || v20) {
                                            v8 = v8
                                            v6 = 0.0
                                            v52 := []int{0}
                                            _ = v52
                                        } else {
                                            v8 = v8
                                            if (false && (!(!((^(v9)) > v9)))) {
                                                fmt.Println(v9)
                                                var v53 float32 = ((0.0 * v6) / v6)
                                                v53 -= 3.14159
                                                _ = v53
                                            }
                                            v8 = "\n\t"
                                        }
                                    }
                                    for fi54 := 0; fi54 < 14; fi54++ {
                                        switch v8 {
                                        case "foo bar":
                                            v2 = (v2 | (v2 & 0))
                                            v55 := GenS3{F0: 0, F1: ((v8 + (v8 + v8)) + (((("\n\t" + v8) + v8) + "hello") + "foo bar"))}
                                            _ = v55
                                        case "a":
                                            v56 := []int{fi50}
                                            _ = v56
                                        case "":
                                            if (v8 <= v8) {
                                                if false {
                                                    v57 := []int{(fi50 & fi50)}
                                                    v58 := GenS6{F0: 669, F1: 0, F2: (((^(v2)) <= -1) && true)}
                                                    v59 := GenS1{F0: (false && v20), F1: 42, F2: uint64(v9)}
                                                    v4 += -2147483648
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                }
                                                v60 := GenS7{F0: (v2 | v2), F1: 514, F2: 1}
                                                _ = v60
                                            } else {
                                                v61 := []int{}
                                                var v62 string = ""
                                                _ = v61
                                                _ = v62
                                            }
                                        default:
                                            var v63 int64 = (v2 >> ((255) & 63))
                                            _ = v63
                                        }
                                        switch (fi50 * 4294967295) {
                                        case 514:
                                            v8 = v8
                                        case 43:
                                            var v64 uint = 25
                                            _ = v64
                                        case 787:
                                            if true {
                                                v2 = v2
                                                v6 -= v6
                                            } else {
                                                _ = len(v1)
                                                if false {
                                                    var v65 int16 = -579
                                                    v65 += (v65 / ((v65) | 1))
                                                    var v66 int = (v4 % (((fi50 & fi50)) | 1))
                                                    v67 := v3[(((v8 + v8) + "") + v8)]
                                                    v68 := GenS8{F0: 79, F1: (v9 % (((v9 ^ (v9 - 0))) | 1)), F2: 100}
                                                    _ = v65
                                                    _ = v66
                                                    _ = v67
                                                    _ = v68
                                                }
                                                v69 := GenS9{F0: 1, F1: -54}
                                                _ = v69
                                            }
                                        }
                                        _ = fi54
                                    }
                                    _ = fi50
                                }
                                var v70 float64 = 3.14159
                                v70 = 1e10
                                _ = v70
                            }
                        }
                        _ = v49
                    }
                case "":
                    v20 = false
                default:
                    var v71 uint16 = v9
                    _ = v71
                }
            case "\n\t":
                if v20 {
                    fmt.Println(v8, v6)
                    fmt.Println(v9)
                    if v20 {
                        if (686 < 1) {
                            v4 = (v4 >> ((42) & 63))
                        } else {
                            v72 := GenS9{F0: 1, F1: -1}
                            var v73 bool = ((!false) && true)
                            switch "\n\t" {
                            case "hello":
                                v74 := GenS10{F0: v73, F1: 2147483648}
                                _ = v74
                            default:
                                var v75 bool = false
                                _ = v75
                            }
                            _ = v72
                            _ = v73
                        }
                    }
                }
            }
        } else {
            var v76 int32 = -2147483648
            var v77 float32 = (v6 / (((v6 - 0.0) / v6) * v6))
            v20 = v20
            _ = v76
            _ = v77
        }
    } else {
        switch "hello" {
        case "a":
            var v78 float64 = 1e10
            _ = v78
        }
    }
    if (!(!((0.0 + (1.5 * ((((((v6 / float32(float64((v6 * v6)))) - (v6 + float32(1.5))) * ((v6 - ((v6 / v6) * (v6 + float32((1e10 * float64((-2.25 * ((((v6 * (1e10 - v6)) / float32(1.5)) - v6) / 1.0)))))))) / v6)) + 3.14159) * -2.25) + -2.25))) == v6))) {
        for fi79 := 0; fi79 < 19; fi79++ {
            if v20 {
                fmt.Println(v8, v20, v6)
                v80 := GenS11{F0: false, F1: true}
                _ = v80
            } else {
                if ((v4 >> ((18446744073709551615) & 63)) != -766) {
                    v9 = 42
                    var v81 int8 = 78
                    var v82 string = v8
                    _ = v81
                    _ = v82
                }
            }
            _ = fi79
        }
        v83 := GenS9{F0: 65535, F1: -115}
        fmt.Println(v9, v6)
        _ = v83
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
