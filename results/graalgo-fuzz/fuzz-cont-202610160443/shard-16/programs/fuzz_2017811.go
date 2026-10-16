package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint; F1 int32; F2 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 int; F2 uint16 }
func genFn0(p0 uint, p1 int) bool {
    v7 := GenS1{F0: 1000, F1: (true && (false && ((^(p0)) == (p0 << ((p0) & 63))))), F2: 2147483647}
    var v8 float64 = 0.1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (0.1 != 1e10)
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 int32 = 2147483647
    var v5 uint = 0
    var v6 int64 = 9223372036854775807
    var v9 uint = 0
    switch "" {
    case "foo bar":
        delete(v3, 0)
    default:
        switch ("hello" + "hello") {
        case "foo bar":
            v10 := []int{4294967295, 4294967295, -9223372036854775808, 32767}
            _ = v10
        case "":
            if (42 > 2) {
                v11 := []int{578, -1, 42, 127}
                _ = v11
            }
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if true {
            var v13 float64 = 3.14159
            var v14 int32 = (v4 * v4)
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := GenS2{F0: fi15, F1: 0}
                _ = fi15
                _ = v16
            }
            v17 := GenS2{F0: (-(fi12)), F1: 42}
            var v18 uint8 = 2
            _ = v17
            _ = v18
        }
        switch "" {
        case "foo bar":
            if false {
                for fi19 := 0; fi19 < 4; fi19++ {
                    v20 := GenS3{F0: 1}
                    switch "\n\t" {
                    case "":
                        var v21 float64 = (((1.5 * float64(3.14159)) + (0.0 - (((float64(((1.5 / (3.14159 / 0.1)) + -2.25)) - ((1e10 - 1.5) / (float64(-2.25) + float64(3.14159)))) + 1.5) - 3.14159))) - float64(3.14159))
                        _ = v21
                    default:
                        var v22 int64 = (v6 % (((v6 ^ (v6 >> (((^(v9))) & 63)))) | 1))
                        _ = v22
                    }
                    var v23 uint64 = 4294967295
                    v1 = 80
                    switch "" {
                    case "foo bar":
                        if false {
                            if (1 >= 54) {
                                var v24 uint64 = v23
                                var v25 int8 = 127
                                _ = v24
                                _ = v25
                            }
                        } else {
                            if ((v1 - 255) <= (v1 << ((v9) & 63))) {
                                v26 := []int{(fi19 ^ (fi12 / ((fi12) | 1))), fi12, 704, int(v1)}
                                _ = v26
                            }
                        }
                    case "hello":
                        var v27 bool = false
                        _ = v27
                    default:
                        switch 32768 {
                        case -506:
                            if false {
                                if (v6 < v6) {
                                    v6 = (v6 + (v6 >> ((283) & 63)))
                                    if (50 == 19) {
                                        if (false && (!((float32(1.5) - 0.1) <= 0.1))) {
                                            v28 := GenS4{F0: ((true && true) && (false && ((!true) && (!("\n\t" <= "a"))))), F1: 32767, F2: 127}
                                            switch "\n\t" {
                                            case "a":
                                                switch "\n\t" {
                                                case "":
                                                    if false {
                                                        v29 := GenS1{F0: 42, F1: true, F2: 0}
                                                        if (-745 <= (-(v4))) {
                                                            switch "" {
                                                            case "\n\t":
                                                                v30 := GenS5{F0: v9, F1: (-(v4)), F2: ("" + "a")}
                                                                _ = v30
                                                            default:
                                                                v23 += v23
                                                            }
                                                            v6 = int64(v23)
                                                        }
                                                        _ = v29
                                                    } else {
                                                        var v31 string = ("foo bar" + (("" + (("foo bar" + (("a" + "foo bar") + ("" + ("\n\t" + "a")))) + (("a" + "") + "foo bar"))) + (("foo bar" + "hello") + "")))
                                                        var v32 uint = (v9 + v9)
                                                        if ((v31 + ((v31 + ((((v31 + v31) + ((v31 + "\n\t") + v31)) + (v31 + v31)) + ("a" + ((v31 + v31) + v31)))) + "a")) < v31) {
                                                            v33 := GenS6{F0: v1}
                                                            v5 = v32
                                                            v6 += (v6 ^ v6)
                                                            _ = v33
                                                        } else {
                                                            v34 := GenS1{F0: 260, F1: (((((v1 ^ v1) != (^(v1))) && true) && ((128 < (fi12 | fi19)) && (!false))) && true), F2: 65535}
                                                            if true {
                                                                if true {
                                                                    var v35 bool = false
                                                                    fmt.Println(fi12)
                                                                    if (("foo bar" + v31) >= v31) {
                                                                        v1 = v1
                                                                        var v36 float64 = 1e10
                                                                        _ = v36
                                                                    } else {
                                                                        var v37 uint32 = 2
                                                                        delete(v2, (-(fi19)))
                                                                        _ = v37
                                                                    }
                                                                    _ = v35
                                                                } else {
                                                                    var v38 uint64 = (v23 - v23)
                                                                    v39 := []int{fi12, fi12}
                                                                    _ = v38
                                                                    _ = v39
                                                                }
                                                            }
                                                            _ = v34
                                                        }
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                default:
                                                    var v40 uint64 = v23
                                                    _ = v40
                                                }
                                            default:
                                                if ((^(v23)) == (v23 << ((uint(v23)) & 63))) {
                                                    var v41 string = ("" + "a")
                                                    v42 := GenS7{F0: (fi19 + (-(fi19)))}
                                                    _ = v41
                                                    _ = v42
                                                }
                                            }
                                            _ = v28
                                        } else {
                                            fmt.Println(fi19)
                                            delete(v2, -1)
                                        }
                                        if true {
                                            if ((false || true) || (100 != 0)) {
                                                v9 -= 952
                                                switch "" {
                                                case "foo bar":
                                                    var v43 uint = v9
                                                    _ = v43
                                                case "hello":
                                                    v44 := []int{int(v4)}
                                                    _ = v44
                                                }
                                                var v45 string = ("a" + (("a" + "a") + "foo bar"))
                                                _ = v45
                                            }
                                        } else {
                                            var v46 int = (fi12 >> ((v9) & 63))
                                            switch (v46 / (((^(fi19))) | 1)) {
                                            case 0:
                                                if (((v4 ^ 1) > (v4 / ((v4) | 1))) && (!true)) {
                                                    v5 -= (v9 ^ 4294967295)
                                                    v46 = -9223372036854775808
                                                }
                                                delete(v3, 0)
                                                v47 := GenS8{F0: 229}
                                                _ = v47
                                            case -9223372036854775808:
                                                switch v46 {
                                                case 1:
                                                    v48 := GenS9{F0: (v6 - (v6 - v6)), F1: (v46 / (((v46 + (fi19 | v46))) | 1)), F2: 65535}
                                                    _ = v48
                                                case 42:
                                                    switch (^(v46)) {
                                                    case 32767:
                                                        delete(v2, 1)
                                                        v6 += 1
                                                        var v49 uint8 = 42
                                                        delete(v2, (fi12 | 1))
                                                        var v50 int8 = 100
                                                        _ = v49
                                                        _ = v50
                                                    case -9223372036854775808:
                                                        var v51 uint8 = (v1 % ((v1) | 1))
                                                        _ = v51
                                                    }
                                                default:
                                                    if ((-(v6)) != (v6 | (^(v6)))) {
                                                        var v52 int8 = 100
                                                        _ = v52
                                                    }
                                                }
                                            }
                                            _ = v46
                                        }
                                        v5 = v5
                                        v1 = (v1 * v1)
                                    } else {
                                        if (!false) {
                                            if true {
                                                var v53 uint64 = 42
                                                v54 := GenS6{F0: (v1 % ((v1) | 1))}
                                                if true {
                                                    var v55 int16 = 1
                                                    if ((("" + ("hello" + "hello")) + "hello") == "") {
                                                        var v56 float64 = (((((0.1 * 0.1) * float64(1.5)) * float64(3.14159)) - ((0.1 + 3.14159) * (1.5 / 1.0))) - float64(3.14159))
                                                        _ = v56
                                                    }
                                                    v57 := GenS1{F0: (-(v55)), F1: (true && false), F2: 4294967295}
                                                    _ = v55
                                                    _ = v57
                                                }
                                                _ = v53
                                                _ = v54
                                            } else {
                                                switch (("foo bar" + ("hello" + ("" + ("" + "a")))) + ("a" + "hello")) {
                                                case "a":
                                                    v6 = 762
                                                case "hello":
                                                    v1 = 42
                                                case "":
                                                    var v58 float64 = 0.0
                                                    _ = v58
                                                default:
                                                    v9 = v9
                                                }
                                            }
                                            if (-32768 == -32768) {
                                                v6 = v6
                                                v59 := []int{fi19, 32767, (fi12 & (fi12 << ((v9) & 63)))}
                                                var v60 uint16 = 60
                                                _ = v59
                                                _ = v60
                                            }
                                            fmt.Println(v4, fi12, fi19)
                                        }
                                        if ((false && (885 == 0)) && false) {
                                            v61 := []int{int(v1), fi12}
                                            v62 := v2[fi19]
                                            _ = v61
                                            _ = v62
                                        } else {
                                            var v63 string = ("\n\t" + ("foo bar" + "foo bar"))
                                            var v64 int32 = 943
                                            var v65 int64 = v6
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                    }
                                }
                            } else {
                                switch (fi12 % (((fi19 / ((int(v9)) | 1))) | 1)) {
                                case -9223372036854775808:
                                    var v66 uint8 = v1
                                    _ = v66
                                }
                            }
                        case -464:
                            var v67 string = (("hello" + "hello") + ((("a" + "hello") + "\n\t") + ""))
                            _ = v67
                        case 127:
                            var v68 uint = 255
                            _ = v68
                        }
                    }
                    switch (-(fi12)) {
                    case 42:
                        switch "a" {
                        case "hello":
                            v5 = (v9 ^ v5)
                        case "a":
                            if false {
                                fmt.Println(v9, v1, fi12)
                                if false {
                                    var v69 int16 = 42
                                    var v70 int16 = int16(v6)
                                    _ = v69
                                    _ = v70
                                }
                                var v71 int32 = v4
                                v72 := []int{fi12, (-(fi19)), fi12}
                                _ = v71
                                _ = v72
                            }
                        case "":
                            v5 -= (v5 / ((222) | 1))
                        default:
                            fmt.Println(v1, v23)
                            var v73 uint32 = 2147483648
                            _ = v73
                        }
                    default:
                        if true {
                            if true {
                                v4 = (-(v4))
                                fmt.Println(v9, v4)
                                v3[42] = (-44 >= 42)
                            }
                            v74 := []int{-1, (^(fi19))}
                            _ = v74
                        } else {
                            v5 += v9
                        }
                        switch "a" {
                        case "hello":
                            v23 -= 774
                        }
                    }
                    var v75 int = (fi12 << (((v5 >> ((v9) & 63))) & 63))
                    _ = fi19
                    _ = v20
                    _ = v23
                    _ = v75
                }
            }
        }
        _ = fi12
    }
    v3[476] = true
    v5 += 4294967295
    delete(v2, -371)
    var v76 int8 = 1
    delete(v3, 2)
    for fi77 := 0; fi77 < 9; fi77++ {
        fmt.Println(fi77, v6)
        _ = fi77
    }
    switch 32767 {
    case 128:
        v4 += int32(v1)
    }
    v4 = (v4 % ((2147483647) | 1))
    v4 += 1
    for fi78 := 0; fi78 < 15; fi78++ {
        v5 = 0
        _ = fi78
    }
    for fi79 := 0; fi79 < 11; fi79++ {
        v1 -= (v1 * 42)
        var v80 uint64 = 2
        fmt.Println(v76)
        _ = fi79
        _ = v80
    }
    var v81 float64 = 0.1
    var v82 int = -1
    var v83 bool = true
    if false {
        v9 -= 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v81
    _ = v82
    _ = v83
    _ = v9
}
