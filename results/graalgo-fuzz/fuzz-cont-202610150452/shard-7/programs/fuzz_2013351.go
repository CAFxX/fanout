package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16; F2 int64 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int32; F1 uint8; F2 int32 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8; F1 uint; F2 int16 }
type GenS10 struct { F0 int32 }
func genFn0(p0 int64) uint8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{(^(fi6)), -277, -9223372036854775808, -9223372036854775808}
        _ = fi6
        _ = v7
    }
    var v8 uint = 1
    _ = p0
    _ = v8
    return 127
}

func main() {
    var v1 int8 = 100
    var v2 int = 4294967295
    v3 := make(map[string]int)
    var v4 float32 = 1e10
    v5 := make(map[int]string)
    fmt.Println(v1)
    if (false && ((-32768 > 1000) || true)) {
        if false {
            var v9 uint = 42
            for fi10 := 0; fi10 < 18; fi10++ {
                if (2147483647 < 2147483648) {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        v2 = fi10
                        _ = fi11
                    }
                    var v12 int32 = 0
                    _ = v12
                } else {
                    v13 := []int{fi10, (fi10 << (((v9 * v9)) & 63))}
                    var v14 int32 = -1
                    _ = v13
                    _ = v14
                }
                _ = fi10
            }
            if (!(603 < 2)) {
                for fi15 := 0; fi15 < 6; fi15++ {
                    v9 += (v9 / ((2) | 1))
                    var v16 uint16 = 255
                    _ = fi15
                    _ = v16
                }
            } else {
                var v17 int64 = 0
                _ = v17
            }
            _ = v9
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            if (float64(float32((-2.25 / 1.0))) <= 3.14159) {
                var v19 float32 = v4
                if (111 > int32(fi18)) {
                    if false {
                        if (true && (false && false)) {
                            fmt.Println(v2, fi18, v2)
                        }
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v21 := GenS1{F0: "foo bar"}
                            v22 := GenS2{F0: 2, F1: -143, F2: -203}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    } else {
                        for fi23 := 0; fi23 < 20; fi23++ {
                            v1 -= (v1 - v1)
                            fmt.Println(v1)
                            if false {
                                var v24 uint = 644
                                if (!true) {
                                    switch "hello" {
                                    case "\n\t":
                                        v25 := []int{(fi18 >> ((v24) & 63)), fi23, fi23}
                                        _ = v25
                                    case "foo bar":
                                        v24 += v24
                                    case "":
                                        var v26 int16 = -392
                                        _ = v26
                                    default:
                                        if true {
                                            switch fi18 {
                                            case -9223372036854775808:
                                                var v27 int64 = -969
                                                _ = v27
                                            case 42:
                                                v28 := []int{fi23, -2147483648}
                                                _ = v28
                                            case -888:
                                                var v29 float32 = v19
                                                v30 := GenS3{F0: 127, F1: 0}
                                                if true {
                                                    if (!false) {
                                                        if (("foo bar" <= ("foo bar" + ((("" + "a") + "a") + ""))) && false) {
                                                            var v31 int64 = -1
                                                            _ = v31
                                                        } else {
                                                            v32 := GenS3{F0: (^(v1)), F1: 1}
                                                            _ = v32
                                                        }
                                                    } else {
                                                        v4 += 0.1
                                                        var v33 float64 = 3.14159
                                                        _ = v33
                                                    }
                                                    delete(v3, ("a" + ""))
                                                }
                                                var v34 string = ("a" + "\n\t")
                                                _ = v29
                                                _ = v30
                                                _ = v34
                                            }
                                            var v35 string = (("\n\t" + ((("a" + "") + "a") + ("hello" + ""))) + ("\n\t" + "a"))
                                            _ = v35
                                        }
                                    }
                                } else {
                                    if false {
                                        delete(v3, "a")
                                    }
                                    v19 = float32((0.1 - float64(v4)))
                                    v36 := GenS4{F0: 1, F1: genFn0(0), F2: 42}
                                    _ = v36
                                }
                                switch "hello" {
                                case "a":
                                    v37 := []int{(fi23 & v2), fi23, (-(fi18))}
                                    _ = v37
                                }
                                if (628 > 1) {
                                    switch ("a" + ("hello" + "")) {
                                    case "hello":
                                        _ = len(v5)
                                    case "":
                                        switch 128 {
                                        case 1:
                                            var v38 bool = (32767 <= 42)
                                            _ = v38
                                        case 42:
                                            switch fi23 {
                                            case -418:
                                                var v39 bool = true
                                                _ = v39
                                            case -504:
                                                v4 -= -2.25
                                            case 32768:
                                                if true {
                                                    var v40 bool = false
                                                    var v41 int64 = 0
                                                    _ = v40
                                                    _ = v41
                                                }
                                            default:
                                                var v42 uint8 = 1
                                                _ = v42
                                            }
                                        default:
                                            delete(v5, int(v1))
                                        }
                                    default:
                                        v43 := GenS5{F0: (844 >= 847), F1: 490}
                                        _ = v43
                                    }
                                    if (-2147483648 >= 0) {
                                        fmt.Println(v2)
                                    } else {
                                        if (((v2 / ((-2147483648) | 1)) >= (fi18 & (fi23 & 128))) && true) {
                                            v44 := []int{}
                                            var v45 float64 = ((1e10 / 3.14159) * -2.25)
                                            _ = v44
                                            _ = v45
                                        }
                                        if (255 != 255) {
                                            fmt.Println(v1, v1)
                                            fmt.Println(fi23, v1, fi23)
                                            fmt.Println(v24, fi18, fi23)
                                            var v46 int8 = v1
                                            v4 += v4
                                            _ = v46
                                        } else {
                                            v47 := GenS2{F0: 65535, F1: 1000, F2: int64(v1)}
                                            v48 := []int{2147483647, -1, fi23, v2}
                                            if (2147483647 >= 1) {
                                                _ = len(v48)
                                                v2 += fi18
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                }
                                switch "a" {
                                case "":
                                    v49 := v3[(("" + "a") + "hello")]
                                    _ = v49
                                }
                                _ = v24
                            }
                            _ = fi23
                        }
                        _ = len(v3)
                        v1 += (v1 - (v1 << ((2) & 63)))
                    }
                    v19 -= v4
                } else {
                    _ = len(v5)
                    if true {
                        v50 := GenS6{F0: 601, F1: (!true), F2: 18446744073709551615}
                        var v51 uint = 2
                        _ = v50
                        _ = v51
                    }
                    switch fi18 {
                    case 857:
                        for fi52 := 0; fi52 < 13; fi52++ {
                            if (3.14159 >= -2.25) {
                                if (32767 == 32767) {
                                    var v53 bool = false
                                    _ = v53
                                }
                            }
                            if true {
                                switch ("hello" + ("hello" + (("hello" + "foo bar") + ("hello" + (("" + "hello") + (("\n\t" + "a") + "a")))))) {
                                case "hello":
                                    var v54 int16 = 287
                                    _ = v54
                                case "foo bar":
                                    fmt.Println(fi18, v1)
                                }
                                v1 += v1
                                v55 := []int{-460}
                                _ = v55
                            } else {
                                fmt.Println(v19)
                            }
                            fmt.Println(v19, v1, v2)
                            var v56 uint = 1
                            _ = fi52
                            _ = v56
                        }
                    }
                }
                _ = v19
            } else {
                v2 = (-(fi18))
                fmt.Println(v2)
                var v57 uint8 = 42
                _ = v57
            }
            for fi58 := 0; fi58 < 20; fi58++ {
                v2 = (^(fi58))
                _ = fi58
            }
            _ = fi18
        }
    } else {
        if true {
            v4 = v4
            for fi59 := 0; fi59 < 15; fi59++ {
                delete(v3, "foo bar")
                for fi60 := 0; fi60 < 8; fi60++ {
                    if (!("foo bar" == "a")) {
                        fmt.Println(v4, fi60, v1)
                        _ = len(v5)
                    }
                    switch "\n\t" {
                    case "foo bar":
                        var v61 uint32 = 596
                        _ = v61
                    case "a":
                        v62 := GenS7{F0: 2, F1: -32768}
                        _ = v62
                    }
                    v2 = (v2 / (((fi60 ^ fi60)) | 1))
                    _ = fi60
                }
                if true {
                    fmt.Println(v1, v4, v2)
                    v4 -= v4
                } else {
                    fmt.Println(v1, v2)
                }
                _ = fi59
            }
        } else {
            fmt.Println(v1, v2, v4)
            if (true && (0 < 37)) {
                v63 := GenS8{F0: "\n\t"}
                if false {
                    v2 = v2
                    if false {
                        if false {
                            var v64 uint16 = 32767
                            fmt.Println(v1)
                            v1 -= 42
                            _ = v64
                        } else {
                            var v65 float32 = 1e10
                            _ = v65
                        }
                    }
                } else {
                    for fi66 := 0; fi66 < 5; fi66++ {
                        switch (fi66 + 42) {
                        case 32768:
                            fmt.Println(v4, fi66, v4)
                        case 127:
                            v67 := GenS5{F0: true, F1: 65535}
                            _ = v67
                        case 1:
                            switch "foo bar" {
                            case "\n\t":
                                v1 = v1
                            case "hello":
                                v68 := []int{(fi66 ^ fi66)}
                                _ = v68
                            case "foo bar":
                                if (9223372036854775807 == 9223372036854775807) {
                                    var v69 uint64 = 4294967295
                                    if (!false) {
                                        v70 := []int{}
                                        fmt.Println(v1)
                                        _ = v70
                                    } else {
                                        v71 := []int{v2, (fi66 / (((v2 + (v2 * fi66))) | 1)), (^(v2)), 9223372036854775807}
                                        v5[v2] = (("foo bar" + "hello") + ("hello" + "foo bar"))
                                        v2 = (v2 & fi66)
                                        _ = v71
                                    }
                                    _ = v69
                                }
                            default:
                                v72 := []int{(fi66 ^ 259), (fi66 >> ((1) & 63)), (v2 & (v2 - (-(fi66))))}
                                _ = v72
                            }
                        default:
                            var v73 uint64 = 42
                            _ = v73
                        }
                        for fi74 := 0; fi74 < 3; fi74++ {
                            var v75 uint8 = 127
                            var v76 bool = true
                            v75 -= (v75 + 0)
                            _ = fi74
                            _ = v75
                            _ = v76
                        }
                        _ = len(v3)
                        _ = fi66
                    }
                    v77 := GenS6{F0: 600, F1: true, F2: 4294967295}
                    _ = v77
                }
                v4 -= ((3.14159 * v4) - v4)
                _ = v63
            } else {
                v78 := GenS9{F0: 0, F1: 42, F2: -1}
                v4 = (v4 - (float32((0.1 + (1.5 * float64((0.0 + v4))))) + v4))
                _ = v78
            }
            switch v2 {
            case 42:
                v4 = (1e10 / v4)
            }
        }
        for fi79 := 0; fi79 < 7; fi79++ {
            v80 := GenS1{F0: ("foo bar" + "\n\t")}
            v81 := []int{}
            var v82 uint16 = 150
            _ = fi79
            _ = v80
            _ = v81
            _ = v82
        }
        var v83 int16 = int16(v2)
        _ = v83
    }
    var v84 uint = 500
    fmt.Println(v2, v2, v4)
    fmt.Println(v2, v2, v84)
    var v85 uint16 = 1
    delete(v5, v2)
    v84 -= v84
    for fi86 := 0; fi86 < 17; fi86++ {
        v84 = (v84 % ((18446744073709551615) | 1))
        if ((0 < 1) || (-746 >= 0)) {
            v84 -= (v84 & (v84 + 177))
            v84 = v84
        } else {
            v85 = (^(v85))
            v1 = v1
            v3[""] = (fi86 ^ (v2 ^ (v2 << ((42) & 63))))
        }
        for fi87 := 0; fi87 < 10; fi87++ {
            if (((uint64(v85) < uint64(v84)) && true) || false) {
                fmt.Println(fi86)
                if false {
                    var v88 uint32 = 64
                    _ = v88
                }
                v89 := []int{176}
                _ = v89
            }
            _ = fi87
        }
        _ = fi86
    }
    v4 += (0.1 / v4)
    v85 = 349
    var v90 uint64 = 1
    v91 := []int{v2}
    for fi92 := 0; fi92 < 16; fi92++ {
        switch (-(fi92)) {
        case 32767:
            var v93 uint = v84
            _ = v93
        default:
            switch (v2 - v2) {
            case 1:
                v94 := GenS10{F0: 2147483647}
                _ = v94
            default:
                v90 = 42
            }
        }
        v2 = fi92
        _ = fi92
    }
    v85 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v90:", v90)
    fmt.Println("slicelen v91:", len(v91), v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v85
    _ = v90
    _ = v91
}
