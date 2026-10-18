package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 uint64; F2 int }
type GenS8 struct { F0 int8; F1 int16; F2 int8 }
type GenS9 struct { F0 int64; F1 int16; F2 uint64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint16; F1 int8 }
func genFn0(p0 int, p1 uint8) int {
    p0 = (p0 + (p0 ^ 1))
    var v7 float32 = (0.0 - 1.5)
    var v8 uint32 = 65535
    switch "foo bar" {
    case "":
        v9 := []int{(^(p0)), p0, -9223372036854775808}
        _ = v9
    default:
        p0 += p0
    }
    var v10 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 uint, p1 bool, p2 uint) uint64 {
    p2 = 264
    var v11 float32 = float32(1e10)
    var v12 bool = p1
    var v13 uint = 63
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 890
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 int8 = -36
    v4 := []string{"hello", "\n\t"}
    var v5 uint32 = 4294967295
    v6 := make(map[int]int)
    v5 = v5
    if (-32768 == 1) {
        v4 = append(v4, "hello")
        for fi14 := 0; fi14 < 3; fi14++ {
            if false {
                var v15 uint32 = (v5 / (((v5 & 1)) | 1))
                v3 -= (^(v3))
                var v16 float32 = (-2.25 + (3.14159 * 0.1))
                _ = v15
                _ = v16
            } else {
                v5 = 42
            }
            switch fi14 {
            case 0:
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 bool = (!false)
                    if (42 != 411) {
                        v18 = (((-1 != 42) && false) && v18)
                        fmt.Println(v5)
                    }
                    switch fi14 {
                    case -9223372036854775808:
                        v5 = v5
                    default:
                        if v18 {
                            switch (fi17 + -9223372036854775808) {
                            case 9223372036854775807:
                                v19 := GenS1{F0: 0, F1: "a"}
                                _ = v19
                            case 32768:
                                v20 := GenS1{F0: 1, F1: (((((("\n\t" + (("foo bar" + "") + "")) + "") + ("a" + "")) + "a") + "foo bar") + (("" + "\n\t") + "\n\t"))}
                                _ = v20
                            }
                            v21 := GenS1{F0: 32767, F1: ("a" + ("foo bar" + ("hello" + "hello")))}
                            _ = v21
                        } else {
                            v22 := v2[fi14]
                            var v23 uint32 = 1
                            var v24 int16 = -32768
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                    _ = fi17
                    _ = v18
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi14
        }
        var v25 int32 = -1
        _ = v25
    } else {
        fmt.Println(v5, v5)
        if (53 > 255) {
            _ = len(v2)
            switch "" {
            case "\n\t":
                v26 := GenS2{F0: 0, F1: int64(v5)}
                _ = v26
            case "a":
                _ = len(v4)
            }
        } else {
            var v27 uint8 = 36
            fmt.Println(v3)
            _ = len(v4)
            _ = v27
        }
        if (true || true) {
            fmt.Println(v3, v5)
            for fi28 := 0; fi28 < 2; fi28++ {
                if (!(0 <= 1000)) {
                    switch ("foo bar" + ((("hello" + (("foo bar" + "\n\t") + ("" + "foo bar"))) + "\n\t") + ("a" + "a"))) {
                    case "\n\t":
                        v29 := GenS3{F0: (fi28 >> ((4294967295) & 63))}
                        _ = v29
                    case "a":
                        v30 := GenS4{F0: 42, F1: 0, F2: 9223372036854775807}
                        _ = v30
                    default:
                        v31 := GenS1{F0: 334, F1: (("\n\t" + "a") + "")}
                        _ = v31
                    }
                    if true {
                        if (false && false) {
                            v32 := GenS1{F0: 255, F1: ("\n\t" + (("a" + ("\n\t" + ("a" + "hello"))) + "\n\t"))}
                            var v33 string = ("\n\t" + "")
                            switch "\n\t" {
                            case "a":
                                var v34 uint = 42
                                _ = v34
                            case "\n\t":
                                fmt.Println(v3, v5)
                            case "foo bar":
                                if (v3 < (v3 & v3)) {
                                    v3 = v3
                                    var v35 int = genFn0(fi28, 127)
                                    _ = v35
                                }
                            default:
                                var v36 int = genFn0(fi28, 42)
                                _ = v36
                            }
                            _ = v32
                            _ = v33
                        } else {
                            var v37 int8 = (v3 >> ((4294967295) & 63))
                            v37 += (^(v3))
                            fmt.Println(v37, v37)
                            _ = v37
                        }
                        v38 := GenS5{F0: 4294967295, F1: 255}
                        _ = v38
                    }
                    var v39 int8 = (v3 * v3)
                    _ = v39
                }
                var v40 float32 = (1.5 * (1.5 / float32(0.1)))
                _ = fi28
                _ = v40
            }
            v5 = v5
        } else {
            v41 := GenS6{F0: 165}
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        if (false && (false || (213 == -36))) {
            if false {
                v5 -= v5
                for fi43 := 0; fi43 < 15; fi43++ {
                    var v44 float64 = 1.5
                    _ = fi43
                    _ = v44
                }
                switch fi42 {
                case 722:
                    var v45 uint32 = v5
                    _ = v45
                }
            } else {
                var v46 int32 = 0
                _ = v46
            }
        } else {
            var v47 uint = 2
            fmt.Println(v5, fi42)
            _ = v47
        }
        if (uint8(v5) > 255) {
            if (42 <= 0) {
                var v48 uint64 = genFn1(973, (0 == 1), 42)
                _ = len(v6)
                _ = v48
            } else {
                for fi49 := 0; fi49 < 10; fi49++ {
                    v50 := GenS4{F0: 0, F1: 222, F2: 0}
                    v51 := GenS7{F0: "foo bar", F1: 2, F2: 42}
                    var v52 uint32 = v5
                    var v53 float64 = ((0.1 - 3.14159) - 0.0)
                    fmt.Println(v5, v52)
                    _ = fi49
                    _ = v50
                    _ = v51
                    _ = v52
                    _ = v53
                }
                v3 = v3
            }
            _ = len(v2)
        } else {
            v54 := GenS8{F0: v3, F1: 1, F2: (v3 * 1)}
            _ = len(v6)
            _ = v54
        }
        switch "\n\t" {
        case "a":
            for fi55 := 0; fi55 < 4; fi55++ {
                fmt.Println(v5)
                switch "hello" {
                case "a":
                    switch genFn0((^(fi42)), 68) {
                    case 32768:
                        fmt.Println(fi42, fi55, fi42)
                    }
                case "hello":
                    v56 := GenS9{F0: 729, F1: -1, F2: genFn1(18446744073709551615, true, 1)}
                    _ = v56
                case "":
                    _ = len(v2)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi55
            }
        case "foo bar":
            var v57 int16 = 756
            _ = v57
        case "\n\t":
            for fi58 := 0; fi58 < 14; fi58++ {
                _ = len(v4)
                _ = fi58
            }
        default:
            var v59 int64 = -1
            _ = v59
        }
        _ = fi42
    }
    if true {
        fmt.Println(v3, v5)
        var v60 int32 = 42
        _ = v60
    } else {
        var v61 float32 = 1e10
        v61 += 1.5
        _ = v61
    }
    v5 = 2
    var v62 int32 = -2147483648
    var v63 int64 = 1
    switch (("hello" + ("\n\t" + "foo bar")) + ("foo bar" + ("" + ("" + ("foo bar" + "hello"))))) {
    case "hello":
        v6[-9223372036854775808] = genFn0(2147483647, 2)
    }
    for fi64 := 0; fi64 < 14; fi64++ {
        v65 := v2[(fi64 + 302)]
        _ = fi64
        _ = v65
    }
    if false {
        if (!false) {
            v66 := GenS10{F0: 4294967295}
            _ = v66
        }
        v3 += (v3 & v3)
    }
    var v67 string = ""
    if (103 >= 678) {
        var v68 int16 = 1000
        v69 := GenS1{F0: 884, F1: v67}
        _ = v68
        _ = v69
    }
    v70 := GenS7{F0: "", F1: 1, F2: 128}
    for fi71 := 0; fi71 < 1; fi71++ {
        for fi72 := 0; fi72 < 1; fi72++ {
            if false {
                switch "hello" {
                case "\n\t":
                    switch (fi71 * int(v62)) {
                    case -493:
                        v5 += v5
                    case -1:
                        fmt.Println(fi71)
                        v62 = v62
                    default:
                        var v73 int32 = (v62 % (((v62 % ((int32(fi71)) | 1))) | 1))
                        _ = v73
                    }
                }
            } else {
                v5 = (v5 & (v5 * 969))
                if true {
                    delete(v6, (fi71 ^ 0))
                } else {
                    var v74 float64 = (-2.25 - -2.25)
                    if true {
                        v5 = (v5 >> ((2) & 63))
                    }
                    _ = v74
                }
                v67 = "hello"
            }
            _ = fi72
        }
        v75 := GenS1{F0: 169, F1: "foo bar"}
        _ = fi71
        _ = v75
    }
    if false {
        v4 = append(v4, (("\n\t" + v67) + v67))
    }
    v76 := GenS11{F0: 2, F1: v3}
    if true {
        if (v67 <= ((("a" + v67) + v67) + "foo bar")) {
            var v77 int64 = (v63 << ((2) & 63))
            var v78 bool = (v62 < v62)
            _ = v77
            _ = v78
        } else {
            if true {
                v6[128] = 4294967295
            } else {
                var v79 int32 = (-(v62))
                _ = v79
            }
        }
        var v80 uint64 = 2
        _ = len(v4)
        _ = v80
    } else {
        var v81 float32 = 1.5
        v62 = v62
        _ = v81
    }
    v4 = append(v4, "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v67
    _ = v70
    _ = v76
}
