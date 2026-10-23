package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int64; F2 int64 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS6 struct { F0 uint8; F1 int16; F2 uint }
func genFn0(p0 int8, p1 uint8, p2 int64) int {
    var v4 bool = (true || false)
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int = fi5
        v7 := []int{v6}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        switch (fi8 + (fi8 >> ((18446744073709551615) & 63))) {
        case 32768:
            v9 := GenS1{F0: p1}
            _ = v9
        case -1:
            v10 := GenS1{F0: p1}
            _ = v10
        default:
            var v11 int16 = -32768
            _ = v11
        }
        _ = fi8
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -380
}
func genFn1(p0 uint8, p1 int8, p2 uint8) bool {
    for fi12 := 0; fi12 < 16; fi12++ {
        p2 -= p2
        p1 = (p1 % ((0) | 1))
        _ = fi12
    }
    var v13 int64 = 486
    if true {
        v13 -= (^(v13))
        if true {
            v14 := []int{32768, 458}
            for fi15 := 0; fi15 < 18; fi15++ {
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := GenS2{F0: 42}
                    var v18 int64 = 42
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                p0 -= (p2 + (^(p2)))
                _ = fi15
            }
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (1 > -32768)
}

func main() {
    var v1 string = ("hello" + "\n\t")
    var v2 uint64 = 1
    var v3 uint64 = 1
    var v19 uint32 = 0
    var v20 uint32 = (v19 << ((318) & 63))
    if (float32(((((1.5 * 0.0) + float64((0.0 + 1e10))) * float64(3.14159)) + 3.14159)) < (((0.1 + (-2.25 * (0.1 / (1.5 * (0.1 * ((3.14159 + 1e10) - 1e10)))))) + float32(1.5)) + 1.5)) {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 bool = true
            fmt.Println(v19)
            var v23 int16 = -1
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 float64 = 1.5
        _ = v24
    }
    v1 = v1
    if (!(18446744073709551615 > 502)) {
        var v25 string = "\n\t"
        _ = v25
    }
    v1 = (v1 + v1)
    var v26 uint64 = 42
    v1 = (v1 + (v1 + ""))
    var v27 int64 = -1
    for fi28 := 0; fi28 < 10; fi28++ {
        if false {
            var v29 uint64 = 4294967295
            switch "" {
            case "hello":
                var v30 int = 99
                _ = v30
            case "foo bar":
                v27 = (v27 * 1)
            }
            _ = v29
        } else {
            for fi31 := 0; fi31 < 9; fi31++ {
                v32 := GenS2{F0: (v3 & v3)}
                if (fi28 < 127) {
                    var v33 int8 = 0
                    if true {
                        var v34 int32 = -529
                        var v35 uint32 = v20
                        var v36 uint = 699
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    var v37 int64 = 9223372036854775807
                    var v38 int16 = 32767
                    var v39 string = "a"
                    _ = v33
                    _ = v37
                    _ = v38
                    _ = v39
                }
                v40 := GenS2{F0: v26}
                _ = fi31
                _ = v32
                _ = v40
            }
            v3 = 0
        }
        v41 := GenS2{F0: v2}
        _ = fi28
        _ = v41
    }
    var v42 float64 = 0.1
    var v43 uint32 = (v19 >> ((0) & 63))
    v20 += (v20 >> ((4294967295) & 63))
    if (0 < 255) {
        v44 := GenS1{F0: 236}
        if (!true) {
            v45 := GenS3{F0: v27, F1: -1, F2: v27}
            _ = v45
        }
        _ = v44
    }
    switch ((v1 + ((((v1 + v1) + v1) + v1) + "a")) + v1) {
    case "foo bar":
        if (65535 <= 976) {
            v26 = v2
            switch 127 {
            case 128:
                fmt.Println(v2)
            case 355:
                fmt.Println(v26, v27)
            }
        } else {
            v1 = "hello"
            v2 = v3
            v46 := []int{}
            _ = v46
        }
    case "\n\t":
        var v47 uint = 918
        _ = v47
    case "a":
        for fi48 := 0; fi48 < 12; fi48++ {
            fmt.Println(v20, v1, v2)
            if true {
                for fi49 := 0; fi49 < 15; fi49++ {
                    fmt.Println(v1, v26, v19)
                    _ = fi49
                }
            } else {
                if true {
                    switch fi48 {
                    case -190:
                        for fi50 := 0; fi50 < 19; fi50++ {
                            v43 = (v20 - (v20 / ((v20) | 1)))
                            v51 := []int{fi48, (^(fi48)), fi50, 128}
                            _ = fi50
                            _ = v51
                        }
                    case 127:
                        var v52 string = "hello"
                        _ = v52
                    }
                }
                switch "" {
                case "foo bar":
                    var v53 int = (fi48 % ((-2147483648) | 1))
                    _ = v53
                }
                v54 := []int{fi48, genFn0(0, 212, 1), fi48}
                _ = v54
            }
            _ = fi48
        }
    default:
        if ((!true) || (!(!(v27 < v27)))) {
            var v55 string = "\n\t"
            if (((v3 - (v26 + 1)) <= (^(v2))) || (v1 != "\n\t")) {
                if (v3 != 2) {
                    v56 := []int{9223372036854775807, 42, -9223372036854775808}
                    v26 += (v3 * 18446744073709551615)
                    v3 -= (^(v2))
                    _ = v56
                } else {
                    v57 := []int{127, -2147483648}
                    fmt.Println(v42, v2, v19)
                    if (!(1 < 255)) {
                        v58 := GenS4{F0: true, F1: 0, F2: (v43 + 42)}
                        if len(v57) > 0 {
                            v57[0] = genFn0(100, 128, 9223372036854775807)
                        }
                        switch v1 {
                        case "":
                            if (1000 < -32768) {
                                v2 = 18446744073709551615
                            } else {
                                if (((v1 + "\n\t") + "\n\t") != (("foo bar" + "hello") + (v55 + ("foo bar" + "foo bar")))) {
                                    fmt.Println(v27, v20)
                                } else {
                                    var v59 float64 = v42
                                    _ = v59
                                }
                                v57 = append(v57, genFn0(100, 255, (v27 & 915)))
                                if (((float32(-2.25) - 0.1) >= 0.0) || true) {
                                    var v60 int8 = 1
                                    fmt.Println(v19, v60)
                                    v20 = 311
                                    _ = v60
                                } else {
                                    v61 := GenS4{F0: (!(v55 > ("foo bar" + v1))), F1: 42, F2: v43}
                                    for fi62 := 0; fi62 < 5; fi62++ {
                                        for fi63 := 0; fi63 < 12; fi63++ {
                                            fmt.Println(v43, v43, v20)
                                            var v64 uint8 = 128
                                            v64 = v64
                                            _ = fi63
                                            _ = v64
                                        }
                                        switch (fi62 & (fi62 | fi62)) {
                                        case 128:
                                            switch fi62 {
                                            case 695:
                                                v27 = v27
                                            case -545:
                                                fmt.Println(v42, v43, v27)
                                            }
                                        case -600:
                                            v3 -= (v2 & (^(v3)))
                                        }
                                        _ = fi62
                                    }
                                    var v65 int32 = 541
                                    _ = v61
                                    _ = v65
                                }
                            }
                        case "foo bar":
                            v66 := GenS5{F0: -128, F1: 65535, F2: 128}
                            _ = v66
                        }
                        _ = v58
                    }
                    _ = v57
                }
                var v67 string = "foo bar"
                _ = v67
            }
            if true {
                if (((-(v27)) == 436) || false) {
                    fmt.Println(v42)
                    v19 = (v43 / ((42) | 1))
                }
                var v68 uint16 = 314
                if (!false) {
                    var v69 int8 = 42
                    v70 := []int{-2147483648, -1, 128, 32768}
                    var v71 int64 = 0
                    _ = v69
                    _ = v70
                    _ = v71
                } else {
                    v72 := []int{-9223372036854775808, 423}
                    for fi73 := 0; fi73 < 7; fi73++ {
                        v74 := GenS3{F0: 9223372036854775807, F1: v27, F2: (v27 << ((65535) & 63))}
                        if len(v72) > 0 {
                            v72[0] = -213
                        }
                        for fi75 := 0; fi75 < 17; fi75++ {
                            v27 = v27
                            _ = fi75
                        }
                        _ = fi73
                        _ = v74
                    }
                    _ = v72
                }
                _ = v68
            } else {
                v76 := []int{128, 861, 32768}
                if len(v76) >= 2 {
                    v76 = v76[1:len(v76)]
                }
                var v77 uint32 = (v20 | (v20 + 4294967295))
                _ = v76
                _ = v77
            }
            _ = v55
        }
    }
    v78 := []int{32768}
    v79 := GenS2{F0: (v2 / ((v3) | 1))}
    for fi80 := 0; fi80 < 4; fi80++ {
        if true {
            fmt.Println(v19)
        } else {
            v3 = 0
            for fi81 := 0; fi81 < 8; fi81++ {
                switch fi80 {
                case -502:
                    v82 := GenS6{F0: 176, F1: 1, F2: 1}
                    var v83 uint64 = v2
                    _ = v82
                    _ = v83
                }
                _ = fi81
            }
            v1 = "\n\t"
        }
        for fi84 := 0; fi84 < 7; fi84++ {
            var v85 uint64 = (v26 - uint64(v43))
            v3 = v3
            _ = fi84
            _ = v85
        }
        _ = fi80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v78:", len(v78), v78)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v42
    _ = v43
    _ = v78
    _ = v79
}
