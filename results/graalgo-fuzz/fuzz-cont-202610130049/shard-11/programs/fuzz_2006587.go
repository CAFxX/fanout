package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint64; F1 bool; F2 string }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 int8 }
func genFn0(p0 int8, p1 bool) int16 {
    fmt.Println(p0, p0)
    p1 = (42 <= p0)
    switch "hello" {
    case "a":
        p0 = 0
    case "foo bar":
        fmt.Println(p1, p1)
    case "":
        switch "foo bar" {
        case "a":
            v4 := GenS1{F0: 127, F1: 100}
            _ = v4
        case "foo bar":
            v5 := GenS1{F0: 0, F1: -1}
            _ = v5
        default:
            if p1 {
                for fi6 := 0; fi6 < 2; fi6++ {
                    if (0.0 < 0.1) {
                        if false {
                            var v7 uint8 = 1
                            _ = v7
                        }
                        var v8 int = (fi6 | (fi6 >> ((65535) & 63)))
                        switch (fi6 << ((900) & 63)) {
                        case 1:
                            v9 := GenS2{F0: 911}
                            _ = v9
                        case 2147483647:
                            for fi10 := 0; fi10 < 11; fi10++ {
                                v11 := GenS3{F0: 0, F1: 0}
                                var v12 int = 1
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                        case -1:
                            var v13 float64 = float64(0.0)
                            _ = v13
                        }
                        _ = v8
                    } else {
                        if (false || p1) {
                            var v14 int = (fi6 % ((274) | 1))
                            fmt.Println(p0, v14)
                            for fi15 := 0; fi15 < 10; fi15++ {
                                var v16 bool = p1
                                var v17 uint64 = 4294967295
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            _ = v14
                        } else {
                            var v18 string = "\n\t"
                            for fi19 := 0; fi19 < 4; fi19++ {
                                v20 := GenS1{F0: 2, F1: 0}
                                _ = fi19
                                _ = v20
                            }
                            _ = v18
                        }
                    }
                    _ = fi6
                }
                p1 = (float64((0.0 * ((-2.25 / 1.0) / 1.5))) >= ((((1e10 + (1.5 / 3.14159)) / 3.14159) - float64(float32((float64(1e10) + 0.0)))) + (0.1 + float64(1.5))))
                p0 = (-(p0))
            } else {
                v21 := GenS3{F0: 32767, F1: 1}
                _ = v21
            }
        }
    default:
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS4{F0: 1, F1: true, F2: ("\n\t" + "a")}
            v24 := []int{}
            switch 902 {
            case -582:
                p1 = ((32767 > 32767) && p1)
            case 657:
                var v25 int = fi22
                _ = v25
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    switch ("" + "") {
    case "foo bar":
        v26 := GenS2{F0: 1}
        _ = v26
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    if (255 <= 990) {
        var v27 int16 = 42
        v27 += -1
        _ = v27
    }
    var v28 int16 = genFn0(-128, false)
    var v29 uint32 = 2147483648
    delete(v3, "\n\t")
    v29 = (v29 << ((18446744073709551615) & 63))
    if true {
        v30 := GenS1{F0: 2, F1: -128}
        _ = v30
    } else {
        delete(v2, 684)
        v28 = (v28 * -32768)
        v29 -= v29
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        delete(v3, ((("\n\t" + (("hello" + "\n\t") + "foo bar")) + ("a" + (("" + "foo bar") + ((("hello" + "\n\t") + "") + "foo bar")))) + "foo bar"))
        _ = fi31
    }
    var v32 float64 = 1.5
    fmt.Println(v29, v28)
    for fi33 := 0; fi33 < 20; fi33++ {
        v34 := GenS3{F0: (v28 | (v28 >> ((495) & 63))), F1: 255}
        _ = fi33
        _ = v34
    }
    var v35 uint32 = (v29 >> ((42) & 63))
    switch 32767 {
    case 128:
        if true {
            v36 := GenS3{F0: (v28 << ((0) & 63)), F1: 0}
            switch 32768 {
            case 740:
                var v37 uint32 = (v29 << ((2) & 63))
                _ = v37
            case 42:
                v38 := []int{-2147483648, 742}
                _ = v38
            }
            _ = v36
        } else {
            v32 = v32
            var v39 float64 = v32
            _ = v39
        }
    case 42:
        v40 := v2[42]
        _ = v40
    case -2147483648:
        if false {
            v41 := GenS4{F0: 18446744073709551615, F1: true, F2: "foo bar"}
            _ = v41
        }
    }
    switch (("foo bar" + ("" + "hello")) + "") {
    case "hello":
        delete(v1, v29)
    case "a":
        if false {
            if (1 == 968) {
                v42 := GenS1{F0: 255, F1: 79}
                var v43 uint32 = 485
                _ = v42
                _ = v43
            }
            v44 := []int{32767, 32767, -2147483648}
            _ = len(v1)
            _ = v44
        } else {
            switch "" {
            case "":
                v1[v29] = (v35 * v35)
            case "foo bar":
                if false {
                    fmt.Println(v32, v28, v29)
                    var v45 int32 = int32(v28)
                    fmt.Println(v28)
                    _ = v45
                } else {
                    switch ("" + "a") {
                    case "":
                        v28 = (v28 + genFn0(0, ((false || false) && (813 < 4294967295))))
                    case "hello":
                        switch -65 {
                        case 344:
                            v46 := GenS4{F0: 891, F1: (1e10 == (v32 - v32)), F2: (("\n\t" + "") + "\n\t")}
                            _ = v46
                        case -2147483648:
                            switch "" {
                            case "a":
                                switch "a" {
                                case "a":
                                    v47 := GenS2{F0: 105}
                                    _ = v47
                                }
                            default:
                                if (false && ((v29 / ((v29) | 1)) != (v29 | 42))) {
                                    v48 := GenS4{F0: 911, F1: false, F2: "foo bar"}
                                    _ = v48
                                } else {
                                    fmt.Println(v28, v35)
                                    for fi49 := 0; fi49 < 18; fi49++ {
                                        if true {
                                            fmt.Println(v35, v28, v28)
                                            if true {
                                                switch fi49 {
                                                case 9223372036854775807:
                                                    var v50 float32 = 1e10
                                                    _ = v50
                                                default:
                                                    fmt.Println(v29, v29)
                                                }
                                                v51 := GenS1{F0: 128, F1: 88}
                                                var v52 uint8 = 255
                                                _ = v51
                                                _ = v52
                                            }
                                            for fi53 := 0; fi53 < 5; fi53++ {
                                                var v54 uint16 = 32767
                                                _ = fi53
                                                _ = v54
                                            }
                                        }
                                        var v55 int32 = 2147483647
                                        for fi56 := 0; fi56 < 2; fi56++ {
                                            var v57 uint32 = (v35 | v35)
                                            _ = fi56
                                            _ = v57
                                        }
                                        _ = fi49
                                        _ = v55
                                    }
                                }
                            }
                        default:
                            var v58 int = 9223372036854775807
                            _ = v58
                        }
                    case "foo bar":
                        switch 2147483647 {
                        case 32767:
                            v28 = genFn0(42, (false && (669 > 32767)))
                        default:
                            switch 32768 {
                            case -1:
                                if (true || (!((v28 >> ((65535) & 63)) > (v28 - (v28 << ((150) & 63)))))) {
                                    if (35 < 0) {
                                        fmt.Println(v32, v28)
                                    } else {
                                        var v59 bool = false
                                        _ = v59
                                    }
                                    if false {
                                        fmt.Println(v29, v29)
                                        for fi60 := 0; fi60 < 13; fi60++ {
                                            v61 := []int{252}
                                            for fi62 := 0; fi62 < 18; fi62++ {
                                                v35 = (v29 / (((^(v29))) | 1))
                                                _ = fi62
                                            }
                                            var v63 uint8 = 2
                                            _ = fi60
                                            _ = v61
                                            _ = v63
                                        }
                                    } else {
                                        v64 := GenS5{F0: 773, F1: 32767}
                                        var v65 bool = false
                                        var v66 int16 = v28
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                    if (((1.5 + 1e10) >= 0.0) || true) {
                                        if false {
                                            v29 = 2
                                            var v67 int32 = 0
                                            _ = v67
                                        }
                                        for fi68 := 0; fi68 < 15; fi68++ {
                                            v28 = (v28 | -746)
                                            fmt.Println(v35)
                                            var v69 uint64 = 1
                                            _ = fi68
                                            _ = v69
                                        }
                                    }
                                } else {
                                    v29 = v29
                                    v70 := GenS6{F0: (1 < 42), F1: -2147483648}
                                    _ = v70
                                }
                            case 32768:
                                var v71 float64 = v32
                                _ = v71
                            case 1:
                                for fi72 := 0; fi72 < 3; fi72++ {
                                    var v73 uint8 = 255
                                    _ = fi72
                                    _ = v73
                                }
                            }
                        }
                    }
                    switch "foo bar" {
                    case "":
                        v74 := GenS2{F0: 65535}
                        _ = v74
                    case "foo bar":
                        var v75 int32 = 0
                        _ = v75
                    default:
                        var v76 int = -720
                        _ = v76
                    }
                    fmt.Println(v35, v28)
                }
            }
            if true {
                v3["hello"] = -333
                v77 := []int{1, 0, 127, -1}
                v29 += (v29 & (v29 / (((v35 | v35)) | 1)))
                _ = v77
            }
            var v78 float32 = 1.5
            _ = v78
        }
    case "\n\t":
        var v79 int32 = int32(v29)
        _ = v79
    default:
        for fi80 := 0; fi80 < 15; fi80++ {
            if false {
                if false {
                    for fi81 := 0; fi81 < 9; fi81++ {
                        var v82 bool = true
                        switch fi81 {
                        case -249:
                            switch fi80 {
                            case 4294967295:
                                v82 = ((!false) && false)
                            }
                        case 2147483647:
                            v83 := v2[fi81]
                            _ = v83
                        }
                        _ = fi81
                        _ = v82
                    }
                    v84 := GenS5{F0: -1, F1: 1}
                    _ = v84
                }
                fmt.Println(fi80, v32, v28)
                var v85 uint16 = 42
                _ = v85
            } else {
                fmt.Println(v29, fi80)
                fmt.Println(v32, fi80, v28)
            }
            v32 = ((1.5 * v32) * v32)
            _ = fi80
        }
    }
    v29 += 1
    v86 := []int{289, 69}
    v32 = 3.14159
    v87 := GenS7{F0: -128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v86:", len(v86), v86)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v35
    _ = v86
    _ = v87
}
