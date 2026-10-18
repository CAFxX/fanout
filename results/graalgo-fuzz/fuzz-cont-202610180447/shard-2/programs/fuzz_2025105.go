package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS10 struct { F0 uint16; F1 string }
type GenS11 struct { F0 uint32; F1 int64 }
type GenS12 struct { F0 bool; F1 uint16; F2 uint }
func genFn0(p0 uint16, p1 string, p2 int32) bool {
    v5 := GenS1{F0: "foo bar"}
    v6 := []int{42, 430, -1, 127}
    var v7 int8 = 127
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 uint64, p1 int16, p2 bool) int {
    if (p0 != 604) {
        if p2 {
            if false {
                p0 = p0
                p0 += p0
                var v8 int64 = 1
                _ = v8
            }
            v9 := GenS1{F0: "foo bar"}
            _ = v9
        }
    } else {
        v10 := GenS2{F0: 127, F1: 65535}
        if ("a" > "a") {
            fmt.Println(p2, p0)
            v11 := []int{0, 32768, 0, -1}
            _ = v11
        }
        var v12 int32 = 42
        _ = v10
        _ = v12
    }
    if ((("hello" + "\n\t") + "a") == "foo bar") {
        p2 = (!false)
        v13 := GenS3{F0: ("\n\t" + ("" + "foo bar"))}
        v14 := []int{127, 9223372036854775807, -2147483648}
        _ = v13
        _ = v14
    } else {
        v15 := []int{653, -1, 9223372036854775807}
        _ = v15
    }
    fmt.Println(p0)
    switch 127 {
    case 912:
        if true {
            for fi16 := 0; fi16 < 14; fi16++ {
                if (65535 < 65535) {
                    p1 = p1
                } else {
                    if p2 {
                        switch ((("" + "foo bar") + "") + "") {
                        case "foo bar":
                            v17 := []int{825}
                            _ = v17
                        case "hello":
                            switch (fi16 / ((-1) | 1)) {
                            case 221:
                                p0 = (p0 % ((4294967295) | 1))
                            case 934:
                                var v18 bool = p2
                                _ = v18
                            default:
                                var v19 int16 = 42
                                _ = v19
                            }
                        default:
                            p0 += p0
                        }
                        v20 := []int{}
                        switch fi16 {
                        case 128:
                            var v21 float32 = float32(0.0)
                            _ = v21
                        }
                        _ = v20
                    } else {
                        v22 := []int{fi16}
                        var v23 int16 = (p1 & (^(p1)))
                        _ = v22
                        _ = v23
                    }
                }
                var v24 int16 = p1
                if ((0.0 / 1e10) != (float32(0.1) * 1.5)) {
                    if p2 {
                        for fi25 := 0; fi25 < 3; fi25++ {
                            p0 = p0
                            var v26 int16 = p1
                            p1 = 32767
                            _ = fi25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 11; fi27++ {
                            v28 := GenS4{F0: 32767}
                            _ = fi27
                            _ = v28
                        }
                        var v29 uint32 = 22
                        _ = v29
                    }
                    if false {
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            v30 := GenS5{F0: p2, F1: 491}
                            _ = v30
                        case "hello":
                            p1 -= (p1 + (v24 >> ((630) & 63)))
                        case "foo bar":
                            v31 := []int{int(p1)}
                            _ = v31
                        default:
                            fmt.Println(p2, p1, v24)
                        }
                        var v32 float64 = ((-2.25 + 1.5) - (0.0 * 1e10))
                        _ = v32
                    }
                    p1 -= p1
                }
                _ = fi16
                _ = v24
            }
            v33 := GenS6{F0: 503, F1: 42}
            for fi34 := 0; fi34 < 16; fi34++ {
                p0 += 11
                var v35 bool = ((847 < 312) || (61 > 20))
                _ = fi34
                _ = v35
            }
            _ = v33
        }
    case 4294967295:
        for fi36 := 0; fi36 < 14; fi36++ {
            for fi37 := 0; fi37 < 12; fi37++ {
                switch "\n\t" {
                case "a":
                    fmt.Println(fi36, fi36, fi36)
                default:
                    if true {
                        v38 := GenS1{F0: ""}
                        _ = v38
                    }
                }
                var v39 float64 = (0.1 + (-2.25 / 1.0))
                v40 := GenS7{F0: 0, F1: (("\n\t" + "hello") + "\n\t")}
                _ = fi37
                _ = v39
                _ = v40
            }
            _ = fi36
        }
    default:
        fmt.Println(p2)
    }
    var v41 bool = ((128 < 1) && p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v41
    return 101
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    delete(v3, 255)
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 uint = 976
        v44 := GenS8{F0: 1, F1: 4294967295}
        _ = fi42
        _ = v43
        _ = v44
    }
    v45 := GenS7{F0: 9223372036854775807, F1: "hello"}
    var v46 uint = 255
    var v47 int16 = 1
    v47 -= 0
    v47 = v47
    v46 -= v46
    var v48 uint64 = 18446744073709551615
    switch 565 {
    case 2147483647:
        var v49 int = 127
        _ = v49
    default:
        if true {
            var v50 uint = v46
            switch 128 {
            case 0:
                if true {
                    v48 = (v48 / (((v48 - 2)) | 1))
                    switch "" {
                    case "\n\t":
                        v51 := GenS8{F0: 2147483647, F1: 2147483648}
                        _ = v51
                    case "":
                        switch "foo bar" {
                        case "foo bar":
                            if (32767 != 335) {
                                v47 = int16(v48)
                                v48 = v48
                                for fi52 := 0; fi52 < 5; fi52++ {
                                    if (!(!true)) {
                                        v53 := GenS3{F0: "foo bar"}
                                        if (!(!((!true) || false))) {
                                            if false {
                                                v54 := GenS9{F0: 18446744073709551615, F1: -2147483648, F2: 2}
                                                delete(v2, genFn1(v48, 514, (!(65535 >= 625))))
                                                _ = v54
                                            } else {
                                                switch fi52 {
                                                case -9223372036854775808:
                                                    v55 := []int{fi52, (fi52 & genFn1(0, v47, false))}
                                                    _ = v55
                                                case 42:
                                                    v56 := v3[65535]
                                                    _ = v56
                                                case 127:
                                                    v57 := v4[2147483648]
                                                    _ = v57
                                                }
                                            }
                                            var v58 int32 = -1
                                            fmt.Println(v47, v48)
                                            _ = v58
                                        }
                                        v59 := GenS10{F0: 0, F1: (("a" + "") + "")}
                                        _ = v53
                                        _ = v59
                                    } else {
                                        var v60 uint = v50
                                        var v61 bool = true
                                        _ = v60
                                        _ = v61
                                    }
                                    _ = fi52
                                }
                            } else {
                                for fi62 := 0; fi62 < 9; fi62++ {
                                    if (v47 < v47) {
                                        v63 := GenS7{F0: 0, F1: "foo bar"}
                                        _ = len(v2)
                                        v64 := []int{fi62}
                                        _ = v63
                                        _ = v64
                                    }
                                    fmt.Println(fi62)
                                    _ = fi62
                                }
                                fmt.Println(v50, v47, v46)
                            }
                        default:
                            v50 = v50
                        }
                    case "hello":
                        if true {
                            var v65 uint = (v46 | (v46 * (^(v50))))
                            if ("a" <= "\n\t") {
                                if (true && (true || false)) {
                                    v50 -= v46
                                    var v66 int32 = 1
                                    var v67 int64 = 1
                                    _ = v66
                                    _ = v67
                                } else {
                                    v68 := []int{genFn1((v48 >> ((v65) & 63)), (-(v47)), (float32(0.1) >= 1.5)), genFn1((v48 >> ((18446744073709551615) & 63)), v47, false)}
                                    _ = v68
                                }
                            } else {
                                fmt.Println(v46, v50, v48)
                            }
                            fmt.Println(v47)
                            _ = v65
                        }
                    }
                    if false {
                        switch 32768 {
                        case 127:
                            var v69 int16 = (v47 + (v47 + (v47 & 1)))
                            _ = v69
                        case -79:
                            var v70 int8 = 42
                            _ = v70
                        case 128:
                            delete(v2, 32767)
                        default:
                            v47 += 42
                        }
                        switch int(v46) {
                        case 199:
                            fmt.Println(v46, v46)
                        case 128:
                            switch 32767 {
                            case 365:
                                if false {
                                    v71 := v3[322]
                                    _ = v71
                                }
                            case 2147483647:
                                if (!(("foo bar" == "") || (true || true))) {
                                    v72 := GenS7{F0: int64(v47), F1: ("a" + "a")}
                                    v73 := []int{-92, genFn1((^(v48)), v47, false)}
                                    _ = v72
                                    _ = v73
                                }
                            case -1:
                                var v74 uint16 = 42
                                _ = v74
                            default:
                                switch "a" {
                                case "a":
                                    if (true || (-1 < 21)) {
                                        var v75 bool = true
                                        _ = v75
                                    } else {
                                        fmt.Println(v50, v47)
                                        var v76 int = 42
                                        _ = v76
                                    }
                                default:
                                    _ = len(v1)
                                }
                            }
                        case 42:
                            v48 = 362
                        }
                        v46 += (v46 % ((v50) | 1))
                    }
                } else {
                    v47 = v47
                    var v77 float32 = -2.25
                    _ = v77
                }
            case 660:
                v78 := GenS11{F0: 0, F1: -9223372036854775808}
                _ = v78
            case -2147483648:
                if true {
                    var v79 uint64 = v48
                    v80 := GenS8{F0: -475, F1: 892}
                    _ = v79
                    _ = v80
                }
            }
            var v81 uint8 = 145
            _ = v50
            _ = v81
        } else {
            v46 = 4294967295
            for fi82 := 0; fi82 < 9; fi82++ {
                v83 := GenS5{F0: (42 < -9223372036854775808), F1: 42}
                v84 := []int{(fi82 / ((fi82) | 1)), fi82, (fi82 | (^(fi82))), fi82}
                v85 := GenS12{F0: (float64(1.5) == (-2.25 * 1.5)), F1: 65535, F2: v46}
                _ = fi82
                _ = v83
                _ = v84
                _ = v85
            }
        }
    }
    fmt.Println(v47)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
