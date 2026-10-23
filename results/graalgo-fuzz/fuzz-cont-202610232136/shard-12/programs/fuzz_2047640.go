package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int16; F2 int32 }
type GenS5 struct { F0 int32; F1 int8; F2 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 0
    var v3 uint16 = 386
    v4 := []int{271, -2147483648, -2147483648, -609}
    var v5 int8 = 42
    fmt.Println(v2, v3, v3)
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 float32 = (1e10 * 0.0)
            var v8 float64 = -2.25
            _ = fi6
            _ = v7
            _ = v8
        }
        v5 += v5
    }
    v2 = v2
    var v9 uint8 = 255
    var v10 bool = (0 >= 2147483647)
    var v11 uint64 = 42
    if (42 >= 820) {
        v3 += v3
        var v12 int8 = 42
        fmt.Println(v9, v11, v10)
        _ = v12
    } else {
        v13 := GenS1{F0: 0, F1: (v3 - v3), F2: (!false)}
        v4 = append(v4, -2147483648)
        _ = v13
    }
    var v14 uint32 = v2
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint32 = v2
        var v17 int32 = 790
        v11 = (^(v11))
        _ = fi15
        _ = v16
        _ = v17
    }
    if v10 {
        fmt.Println(v5, v5, v5)
        v2 = (v2 | v2)
        var v18 int8 = v5
        _ = v18
    } else {
        if v10 {
            v14 = (v14 - v14)
        }
        v19 := GenS1{F0: -146, F1: (v3 >> ((2) & 63)), F2: v10}
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(v5, v3)
            if (!v10) {
                if (163 != (v14 << ((4294967295) & 63))) {
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v1[(-(fi20))] = 910
                        var v22 uint16 = v3
                        _ = fi21
                        _ = v22
                    }
                    v14 += 65535
                    _ = len(v1)
                } else {
                    v23 := v1[(-(fi20))]
                    v24 := GenS2{F0: fi20}
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 int = (-(fi20))
                    _ = fi25
                    _ = v26
                }
            } else {
                delete(v1, fi20)
            }
            _ = fi20
        }
        _ = v19
    }
    var v27 int32 = 1
    for fi28 := 0; fi28 < 3; fi28++ {
        fmt.Println(v3, v3, v5)
        v11 = (v11 ^ 0)
        _ = fi28
    }
    if false {
        if ((-2.25 * (1.5 - -2.25)) > -2.25) {
            for fi29 := 0; fi29 < 14; fi29++ {
                for fi30 := 0; fi30 < 16; fi30++ {
                    v10 = true
                    fmt.Println(v9)
                    if v10 {
                        var v31 int = (fi30 | fi29)
                        switch "a" {
                        case "a":
                            if (((v9 ^ 167) != v9) || v10) {
                                if v10 {
                                    fmt.Println(fi30, fi29, v2)
                                    fmt.Println(fi29)
                                    fmt.Println(v10, fi30, v14)
                                } else {
                                    v32 := GenS1{F0: v27, F1: v3, F2: v10}
                                    var v33 int = (v31 & fi29)
                                    _ = v32
                                    _ = v33
                                }
                            } else {
                                if ((!false) && (1.5 > (3.14159 + (3.14159 / float64(0.1))))) {
                                    var v34 float32 = 0.1
                                    _ = v34
                                } else {
                                    v14 = (v14 % ((2147483647) | 1))
                                }
                                if v10 {
                                    v27 = v27
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        v11 -= 0
                                    case "":
                                        v35 := GenS3{F0: 18446744073709551615, F1: 6, F2: (v3 - v3)}
                                        if (!v10) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            fmt.Println(fi29, v2)
                                        } else {
                                            v1[(v31 % (((fi30 * 1)) | 1))] = v31
                                            v36 := GenS2{F0: (v31 << ((4294967295) & 63))}
                                            var v37 float32 = ((((0.1 / (-2.25 - (3.14159 + (-2.25 * 3.14159)))) * 1e10) * (1e10 / 0.1)) * ((3.14159 + ((-2.25 * 0.0) + -2.25)) * 1.5))
                                            _ = v36
                                            _ = v37
                                        }
                                        if v10 {
                                            var v38 uint = 2
                                            _ = v38
                                        } else {
                                            v27 = (v27 - (v27 / (((^(v27))) | 1)))
                                        }
                                        _ = v35
                                    }
                                    fmt.Println(v3)
                                }
                                if v10 {
                                    v3 = v3
                                    v9 -= 1
                                    var v39 uint = 969
                                    _ = v39
                                }
                                var v40 uint16 = uint16(v2)
                                var v41 bool = true
                                _ = v40
                                _ = v41
                            }
                        case "foo bar":
                            v42 := GenS4{F0: (v14 | uint32(v3)), F1: -335, F2: (v27 - (-(v27)))}
                            _ = v42
                        case "hello":
                            if (42 <= 122) {
                                _ = len(v4)
                                v43 := GenS3{F0: 18446744073709551615, F1: 381, F2: v3}
                                _ = v43
                            }
                        }
                        var v44 int32 = v27
                        _ = v31
                        _ = v44
                    } else {
                        v14 = v14
                        v14 += v14
                        var v45 string = ("\n\t" + "")
                        var v46 int64 = 1
                        delete(v1, (fi29 * fi30))
                        if false {
                            v47 := v1[(fi30 + (fi29 | (-(fi29))))]
                            fmt.Println(v9)
                            var v48 bool = ((!((v46 * 42) <= (^(v46)))) && false)
                            _ = v47
                            _ = v48
                        } else {
                            v45 = v45
                            v2 = (v2 ^ v2)
                            v9 -= v9
                            if v10 {
                                var v49 float64 = -2.25
                                _ = v49
                            }
                            v1[(fi29 | (fi29 * (fi30 * (fi30 % (((^(fi29))) | 1)))))] = fi30
                            switch 127 {
                            case 4294967295:
                                switch (v45 + v45) {
                                case "hello":
                                    switch ((v45 + (v45 + (v45 + v45))) + ("foo bar" + v45)) {
                                    case "foo bar":
                                        if v10 {
                                            var v50 uint32 = v2
                                            var v51 uint = 65535
                                            delete(v1, fi29)
                                            v50 += v14
                                            _ = v50
                                            _ = v51
                                        }
                                    case "hello":
                                        var v52 string = v45
                                        _ = v52
                                    case "a":
                                        _ = len(v4)
                                    }
                                case "a":
                                    switch fi30 {
                                    case 0:
                                        v14 -= (v14 + v2)
                                    case 128:
                                        switch ((v45 + v45) + v45) {
                                        case "\n\t":
                                            v27 += v27
                                        case "foo bar":
                                            var v53 bool = (v46 <= v46)
                                            _ = v53
                                        }
                                        var v54 int = (-(fi29))
                                        fmt.Println(fi29, v46, v14)
                                        _ = v54
                                    case 4294967295:
                                        if (4294967295 >= 18446744073709551615) {
                                            var v55 uint16 = (v3 & v3)
                                            if (v9 > (^(v9))) {
                                                if (v11 <= v11) {
                                                    var v56 float64 = ((0.1 / 1.0) - (1e10 * -2.25))
                                                    _ = len(v4)
                                                    if true {
                                                        v56 = float64(1e10)
                                                        switch (("a" + v45) + (v45 + "")) {
                                                        case "foo bar":
                                                            v4 = append(v4, (-(fi30)))
                                                        case "a":
                                                            delete(v1, (fi29 ^ fi29))
                                                        case "hello":
                                                            var v57 int16 = int16(v27)
                                                            _ = v57
                                                        }
                                                        var v58 int = fi30
                                                        _ = v58
                                                    }
                                                    _ = v56
                                                } else {
                                                    var v59 uint64 = 0
                                                    _ = v59
                                                }
                                                v2 += v14
                                                v60 := GenS4{F0: 697, F1: 1, F2: (-(v27))}
                                                _ = v60
                                            }
                                            _ = v55
                                        }
                                    }
                                case "\n\t":
                                    v14 += (v14 & 155)
                                }
                            case 2147483647:
                                v61 := GenS5{F0: v27, F1: 76, F2: -1}
                                _ = v61
                            default:
                                v62 := GenS6{F0: v27}
                                _ = v62
                            }
                            switch v45 {
                            case "foo bar":
                                v63 := GenS3{F0: 84, F1: 0, F2: (v3 - (v3 >> ((1) & 63)))}
                                _ = v63
                            case "hello":
                                v46 -= (v46 ^ 42)
                            default:
                                fmt.Println(fi29, v10)
                            }
                        }
                        fmt.Println(v45, v11, v45)
                        _ = v45
                        _ = v46
                    }
                    switch (^(fi30)) {
                    case 9223372036854775807:
                        v64 := GenS7{F0: v11, F1: 184, F2: v9}
                        _ = v64
                    case -970:
                        if v10 {
                            var v65 int64 = -1
                            var v66 uint8 = v9
                            _ = v65
                            _ = v66
                        }
                    }
                    fmt.Println(v2, v3)
                    _ = fi30
                }
                _ = fi29
            }
        }
        v9 += v9
        fmt.Println(v27, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
