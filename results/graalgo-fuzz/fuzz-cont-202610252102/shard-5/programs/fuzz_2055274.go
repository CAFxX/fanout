package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 uint; F1 bool }
func genFn0() int {
    var v6 bool = true
    fmt.Println(v6)
    switch ("foo bar" + "foo bar") {
    case "a":
        v6 = false
    case "\n\t":
        var v7 uint32 = 42
        _ = v7
    case "":
        var v8 uint8 = 1
        _ = v8
    }
    _ = v6
    return -915
}
func genFn1(p0 uint8, p1 int32, p2 string) int8 {
    var v9 uint8 = 225
    v9 -= 220
    if true {
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: (p0 | v9)}
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS2{F0: 32767}
                if ((p1 | -2147483648) > (p1 % ((p1) | 1))) {
                    v14 := []int{fi12, 485}
                    var v15 bool = (-911 >= -32768)
                    _ = len(v14)
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: p0}
                if (1 < 1) {
                    var v17 string = (p2 + ("" + p2))
                    v18 := GenS1{F0: (v9 << ((255) & 63))}
                    v19 := []int{(fi12 ^ (fi12 - fi12))}
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v20 := []int{int(p0), (fi12 >> ((18446744073709551615) & 63)), (fi10 >> ((144) & 63))}
                _ = fi12
                _ = v13
                _ = v16
                _ = v20
            }
            if (!true) {
                var v21 float64 = 1.5
                fmt.Println(p2, p0)
                if true {
                    fmt.Println(p2, p0, v21)
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v23 := []int{-234, 32767, 2147483647, 11}
                        _ = fi22
                        _ = v23
                    }
                    switch (p2 + (p2 + "\n\t")) {
                    case "":
                        if false {
                            switch fi10 {
                            case -1:
                                for fi24 := 0; fi24 < 19; fi24++ {
                                    var v25 float64 = 0.1
                                    v26 := []int{(fi24 & (fi24 >> ((846) & 63))), fi24, (-(fi10))}
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                            case 128:
                                var v27 float32 = 1e10
                                _ = v27
                            case 32767:
                                v28 := GenS1{F0: p0}
                                _ = v28
                            }
                            var v29 uint = 0
                            switch (p2 + "") {
                            case "a":
                                fmt.Println(p1)
                            default:
                                var v30 int16 = -1
                                _ = v30
                            }
                            _ = v29
                        }
                    case "a":
                        var v31 uint16 = 255
                        _ = v31
                    }
                } else {
                    switch (p2 + "hello") {
                    case "foo bar":
                        var v32 float64 = (0.0 - 3.14159)
                        _ = v32
                    case "a":
                        v33 := []int{(fi10 >> ((2) & 63)), fi10, (^(fi10)), (^(fi10))}
                        _ = v33
                    }
                    var v34 float32 = float32((v21 * (v21 - 0.1)))
                    var v35 uint8 = v9
                    _ = v34
                    _ = v35
                }
                _ = v21
            } else {
                fmt.Println(v9, fi10)
                for fi36 := 0; fi36 < 6; fi36++ {
                    v37 := GenS3{F0: p2}
                    _ = fi36
                    _ = v37
                }
            }
            _ = fi10
            _ = v11
        }
        var v38 float32 = 3.14159
        _ = v38
    }
    switch p2 {
    case "hello":
        if (false || true) {
            for fi39 := 0; fi39 < 19; fi39++ {
                for fi40 := 0; fi40 < 12; fi40++ {
                    var v41 int16 = 0
                    v42 := GenS3{F0: "\n\t"}
                    v43 := []int{4294967295, 283}
                    _ = fi40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                p1 = p1
                _ = fi39
            }
            v44 := GenS1{F0: (p0 / ((255) | 1))}
            _ = v44
        } else {
            for fi45 := 0; fi45 < 2; fi45++ {
                fmt.Println(p2, p0)
                _ = fi45
            }
        }
    case "foo bar":
        var v46 float64 = 0.1
        _ = v46
    case "a":
        p0 += 196
    }
    var v47 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v47
    _ = v9
    return -87
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 1e10
    v3 := make(map[uint32]int)
    var v4 float32 = float32((1.5 / ((float64(v2) - (1.5 - (0.1 / (1e10 * 0.1)))) - (-2.25 / float64(v2)))))
    v5 := make(map[uint32]bool)
    fmt.Println(v4)
    if (!false) {
        for fi48 := 0; fi48 < 3; fi48++ {
            for fi49 := 0; fi49 < 5; fi49++ {
                v4 += -2.25
                _ = fi49
            }
            v50 := v5[2]
            _ = fi48
            _ = v50
        }
        if ((!true) || ((1e10 * 1.5) <= 1.5)) {
            v4 = float32((0.0 - 0.0))
            var v51 float32 = v2
            _ = v51
        } else {
            var v52 float32 = -2.25
            var v53 uint16 = 42
            v54 := GenS1{F0: 128}
            _ = v52
            _ = v53
            _ = v54
        }
        v55 := GenS1{F0: 189}
        _ = v55
    } else {
        if (3.14159 != (float32(0.1) * float32(-2.25))) {
            for fi56 := 0; fi56 < 12; fi56++ {
                for fi57 := 0; fi57 < 13; fi57++ {
                    if true {
                        _ = len(v3)
                        v4 += (((v2 / ((1.5 + 0.1) - v4)) - v4) - 1.5)
                    } else {
                        var v58 float32 = (float32(3.14159) + v4)
                        if false {
                            var v59 uint32 = 65535
                            var v60 int16 = -253
                            _ = v59
                            _ = v60
                        }
                        v2 = 3.14159
                        _ = v58
                    }
                    _ = fi57
                }
                v61 := GenS4{F0: 42}
                switch "" {
                case "hello":
                    if false {
                        v62 := v3[65535]
                        switch ("\n\t" + "foo bar") {
                        case "foo bar":
                            if true {
                                v63 := GenS4{F0: 18446744073709551615}
                                _ = v63
                            } else {
                                v5[65535] = (-211 >= -1)
                            }
                        case "a":
                            v64 := GenS5{F0: "hello", F1: -32768}
                            _ = v64
                        case "hello":
                            v65 := GenS2{F0: 105}
                            _ = v65
                        }
                        if true {
                            v62 += v62
                        } else {
                            fmt.Println(v2, fi56, fi56)
                        }
                        _ = v62
                    } else {
                        var v66 float32 = ((v4 * v4) * v2)
                        v2 = v66
                        if ((42 != 0) && false) {
                            var v67 int16 = -516
                            _ = v67
                        } else {
                            v66 += (v4 + float32(1e10))
                            v68 := GenS3{F0: ""}
                            _ = v68
                        }
                        _ = v66
                    }
                case "foo bar":
                    v69 := GenS2{F0: 42}
                    _ = v69
                case "":
                    var v70 int16 = -1
                    _ = v70
                default:
                    fmt.Println(v4, fi56)
                }
                _ = fi56
                _ = v61
            }
            var v71 uint = 382
            _ = v71
        }
        if true {
            fmt.Println(v4)
            for fi72 := 0; fi72 < 5; fi72++ {
                fmt.Println(v2)
                if (255 <= 255) {
                    var v73 int = fi72
                    switch (^(v73)) {
                    case -9223372036854775808:
                        v74 := []int{fi72, fi72}
                        _ = v74
                    }
                    var v75 uint = 4294967295
                    _ = v73
                    _ = v75
                } else {
                    for fi76 := 0; fi76 < 7; fi76++ {
                        v77 := v3[42]
                        _ = fi76
                        _ = v77
                    }
                }
                _ = fi72
            }
            v2 -= v2
        } else {
            for fi78 := 0; fi78 < 5; fi78++ {
                fmt.Println(v2, v4)
                _ = fi78
            }
            v79 := GenS3{F0: "hello"}
            _ = v79
        }
    }
    if false {
        var v80 uint32 = 4294967295
        switch 366 {
        case -907:
            v5[65535] = true
        case 32767:
            if false {
                var v81 int16 = 1000
                v82 := GenS2{F0: 42}
                delete(v3, (v80 * v80))
                _ = v81
                _ = v82
            }
        default:
            v4 -= (float32(((1.5 + (float64(float32(1e10)) / 1e10)) * (0.1 - 3.14159))) / 1.5)
        }
        v2 = v4
        _ = v80
    }
    var v83 bool = (!true)
    delete(v5, 0)
    switch "a" {
    case "a":
        var v84 int32 = -35
        _ = v84
    case "":
        delete(v1, 0)
    default:
        fmt.Println(v83)
    }
    if (!v83) {
        var v85 uint = 973
        _ = v85
    } else {
        if true {
            v4 = ((v2 * float32((((((3.14159 + 0.0) / float64(v2)) - 0.0) - float64((v2 * 0.1))) - ((float64(v4) / 0.1) + float64((v2 - ((((1.5 + 1e10) + (1.5 + v4)) * (1.5 + (3.14159 + 0.0))) + v4))))))) + ((v4 * (v2 * v4)) * (0.1 + 3.14159)))
            v86 := GenS5{F0: ("hello" + "\n\t"), F1: 0}
            var v87 uint8 = 28
            _ = v86
            _ = v87
        }
    }
    switch -215 {
    case 4294967295:
        v88 := GenS6{F0: 66, F1: 32767}
        _ = v88
    case 271:
        v83 = ((false || ((v83 && true) || false)) || (2 < 4294967295))
    case 1:
        var v89 bool = false
        _ = v89
    default:
        if (529 <= 4294967295) {
            for fi90 := 0; fi90 < 12; fi90++ {
                if v83 {
                    var v91 int8 = 42
                    fmt.Println(v83, v2)
                    _ = v91
                } else {
                    v92 := []int{(^(fi90)), 959}
                    for fi93 := 0; fi93 < 1; fi93++ {
                        var v94 int64 = 0
                        _ = fi93
                        _ = v94
                    }
                    v95 := GenS1{F0: 255}
                    _ = v92
                    _ = v95
                }
                _ = fi90
            }
        } else {
            for fi96 := 0; fi96 < 17; fi96++ {
                switch (("hello" + "\n\t") + (("\n\t" + "\n\t") + "hello")) {
                case "":
                    if (v83 && (409 <= 1)) {
                        v83 = ((-350 <= 0) || (true && v83))
                        var v97 uint = 255
                        _ = v97
                    }
                }
                _ = fi96
            }
        }
    }
    fmt.Println(v2)
    var v98 int16 = 1000
    var v99 uint64 = 42
    var v100 uint = 42
    v101 := GenS7{F0: (v100 * (v100 | (v100 - v100))), F1: ((0.0 / 1.5) != ((1e10 * (float64(((v2 + v4) * 3.14159)) - (1e10 * float64(1e10)))) * (((0.0 - float64(0.1)) - ((float64(v4) + (0.0 * 3.14159)) / 1e10)) - 0.0)))}
    fmt.Println(v83, v4, v100)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v98
    _ = v99
}
