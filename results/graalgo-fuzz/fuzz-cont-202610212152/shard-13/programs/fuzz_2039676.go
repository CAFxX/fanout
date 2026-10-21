package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint16; F2 string }
type GenS5 struct { F0 int16; F1 uint32; F2 int64 }
type GenS6 struct { F0 int32; F1 uint32; F2 string }
type GenS7 struct { F0 uint16; F1 uint; F2 int16 }
type GenS8 struct { F0 int32; F1 string; F2 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 int8; F2 uint32 }
type GenS11 struct { F0 uint32; F1 int32; F2 bool }
type GenS12 struct { F0 int8 }
func genFn0() uint8 {
    if ((!true) && (1 >= 1)) {
        var v7 int8 = -1
        v7 += 65
        v8 := []int{}
        _ = v7
        _ = v8
    } else {
        var v9 uint8 = 213
        v10 := []int{-1}
        if true {
            v9 = v9
        } else {
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 int16 = 410
                v9 = (v9 >> ((226) & 63))
                _ = fi11
                _ = v12
            }
            fmt.Println(v9)
            fmt.Println(v9)
        }
        _ = v10
        _ = v9
    }
    var v13 int8 = 42
    var v14 float64 = ((0.1 - 0.0) + 0.1)
    _ = v13
    _ = v14
    return 42
}
func genFn1(p0 int32, p1 bool, p2 uint) uint32 {
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int64 = 888
        switch ("\n\t" + (((("foo bar" + ("hello" + ("\n\t" + ("hello" + "a")))) + "foo bar") + "\n\t") + "hello")) {
        case "":
            if (fi15 <= fi15) {
                v16 -= 42
                p0 = p0
            } else {
                switch fi15 {
                case 42:
                    p2 += 255
                }
            }
        }
        _ = fi15
        _ = v16
    }
    p0 += 2147483647
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int16 = 1
        var v19 uint = 255
        p2 -= (p2 % (((p2 + 96)) | 1))
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    v1 := []int{}
    var v2 uint64 = 0
    var v3 int8 = -39
    v4 := make(map[string]bool)
    var v5 bool = true
    var v6 uint64 = 42
    fmt.Println(v5, v5)
    if ((v5 || v5) && (int64(v3) <= 42)) {
        v20 := GenS1{F0: ("" + "foo bar"), F1: 1, F2: 1}
        switch "foo bar" {
        case "a":
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS2{F0: 42, F1: (-(v3))}
                if ((((("\n\t" + "foo bar") + "\n\t") + "") + "") == "a") {
                    var v23 uint64 = (v6 ^ 0)
                    if (v6 == v6) {
                        switch (-(fi21)) {
                        case 128:
                            switch "foo bar" {
                            case "\n\t":
                                if v5 {
                                    var v24 uint16 = 255
                                    v2 = uint64(v24)
                                    _ = v24
                                } else {
                                    fmt.Println(v2)
                                    v5 = false
                                    v25 := GenS2{F0: 2, F1: (v3 >> ((18446744073709551615) & 63))}
                                    _ = v25
                                }
                            case "":
                                if v5 {
                                    var v26 int = fi21
                                    v5 = false
                                    _ = v26
                                }
                            }
                        case 790:
                            v2 -= 1
                        }
                        v2 -= (v23 >> ((42) & 63))
                        switch "" {
                        case "hello":
                            switch int(v2) {
                            case 1:
                                if v5 {
                                    var v27 uint32 = genFn1(-779, v5, 1)
                                    v1 = append(v1, (fi21 % (((fi21 & (fi21 - (fi21 << ((255) & 63))))) | 1)))
                                    _ = v27
                                } else {
                                    v23 -= (v2 - v23)
                                }
                            }
                        case "a":
                            switch (fi21 * fi21) {
                            case 128:
                                v28 := GenS2{F0: 65535, F1: v3}
                                _ = v28
                            default:
                                if v5 {
                                    var v29 uint = 255
                                    _ = v29
                                } else {
                                    if v5 {
                                        delete(v4, "")
                                        fmt.Println(fi21, v23, v5)
                                    } else {
                                        v30 := GenS3{F0: 68, F1: 2}
                                        if len(v1) > 0 {
                                            v1[0] = (fi21 ^ fi21)
                                        }
                                        for fi31 := 0; fi31 < 10; fi31++ {
                                            v32 := GenS1{F0: "foo bar", F1: 0, F2: 1}
                                            fmt.Println(fi21, fi31, v3)
                                            _ = fi31
                                            _ = v32
                                        }
                                        _ = v30
                                    }
                                    v3 = 1
                                    delete(v4, "hello")
                                }
                            }
                        case "foo bar":
                            fmt.Println(v23, fi21)
                        }
                    }
                    if (v5 || true) {
                        _ = len(v1)
                        if v5 {
                            if false {
                                var v33 uint = 255
                                _ = v33
                            } else {
                                if false {
                                    var v34 float64 = 1.5
                                    if true {
                                        v3 += 42
                                    } else {
                                        _ = len(v4)
                                        v2 += (v2 * 4294967295)
                                        switch fi21 {
                                        case 127:
                                            var v35 int8 = v3
                                            _ = v35
                                        }
                                    }
                                    _ = v34
                                } else {
                                    v3 = 1
                                    if (98 == (v23 & (v6 % (((^(v23))) | 1)))) {
                                        if ((v6 & v2) == (^(v2))) {
                                            if true {
                                                if v5 {
                                                    fmt.Println(v2, v3, v2)
                                                    var v36 uint32 = genFn1(2147483647, (fi21 > fi21), 2)
                                                    switch "" {
                                                    case "hello":
                                                        for fi37 := 0; fi37 < 11; fi37++ {
                                                            v6 += v23
                                                            _ = len(v4)
                                                            v38 := v4["a"]
                                                            _ = fi37
                                                            _ = v38
                                                        }
                                                    case "\n\t":
                                                        v2 = v2
                                                    }
                                                    _ = v36
                                                }
                                                switch "" {
                                                case "\n\t":
                                                    if v5 {
                                                        v39 := GenS4{F0: v5, F1: 65535, F2: "a"}
                                                        if false {
                                                            v40 := GenS4{F0: (("foo bar" < "a") || (!v5)), F1: 1, F2: "hello"}
                                                            _ = len(v1)
                                                            v41 := GenS3{F0: 0, F1: 65535}
                                                            _ = v40
                                                            _ = v41
                                                        } else {
                                                            v4["foo bar"] = ((-(fi21)) > fi21)
                                                        }
                                                        var v42 int = fi21
                                                        _ = v39
                                                        _ = v42
                                                    }
                                                }
                                                for fi43 := 0; fi43 < 8; fi43++ {
                                                    var v44 uint8 = 127
                                                    if false {
                                                        if v5 {
                                                            switch fi21 {
                                                            case 127:
                                                                v6 = 0
                                                            case 42:
                                                                delete(v4, "\n\t")
                                                                if v5 {
                                                                    var v45 int16 = 32767
                                                                    _ = v45
                                                                } else {
                                                                    v5 = (v2 < v2)
                                                                }
                                                            case 9223372036854775807:
                                                                if v5 {
                                                                    v46 := GenS4{F0: true, F1: 42, F2: ("a" + "\n\t")}
                                                                    _ = v46
                                                                } else {
                                                                    delete(v4, "a")
                                                                    fmt.Println(v44)
                                                                    _ = len(v1)
                                                                }
                                                            default:
                                                                _ = len(v1)
                                                            }
                                                            v6 = v6
                                                            v44 = (v44 << ((958) & 63))
                                                        }
                                                        v3 = (v3 / ((42) | 1))
                                                    } else {
                                                        v5 = v5
                                                        var v47 int = fi21
                                                        _ = v47
                                                    }
                                                    var v48 uint8 = genFn0()
                                                    _ = fi43
                                                    _ = v44
                                                    _ = v48
                                                }
                                            } else {
                                                var v49 int16 = 1
                                                v2 = v6
                                                _ = v49
                                            }
                                            v6 = v2
                                        }
                                        delete(v4, "a")
                                        v23 -= 18446744073709551615
                                    }
                                    if (v5 || false) {
                                        v50 := GenS4{F0: (!(1e10 < ((((0.1 * 0.1) / 3.14159) + 3.14159) * -2.25))), F1: 1, F2: ""}
                                        _ = v50
                                    } else {
                                        for fi51 := 0; fi51 < 20; fi51++ {
                                            v23 += (v6 << ((2) & 63))
                                            _ = fi51
                                        }
                                        if (v3 < (v3 % ((1) | 1))) {
                                            switch "hello" {
                                            case "a":
                                                var v52 uint = 18446744073709551615
                                                _ = v52
                                            case "\n\t":
                                                if v5 {
                                                    var v53 uint = 215
                                                    var v54 int = fi21
                                                    _ = v53
                                                    _ = v54
                                                }
                                            case "hello":
                                                var v55 string = ((("hello" + "hello") + "foo bar") + ("" + ("a" + ("foo bar" + "foo bar"))))
                                                _ = v55
                                            }
                                        }
                                        var v56 int = (fi21 + (fi21 - -9223372036854775808))
                                        _ = v56
                                    }
                                }
                                if v5 {
                                    if v5 {
                                        v6 -= 390
                                    }
                                } else {
                                    var v57 int8 = v3
                                    v58 := GenS5{F0: 1, F1: 855, F2: 42}
                                    _ = v57
                                    _ = v58
                                }
                            }
                        }
                        var v59 uint8 = 127
                        _ = v59
                    } else {
                        v5 = true
                        if true {
                            v60 := GenS3{F0: 16, F1: 356}
                            v61 := GenS6{F0: 42, F1: 2, F2: "foo bar"}
                            v2 = v23
                            _ = v60
                            _ = v61
                        } else {
                            var v62 int = fi21
                            _ = v62
                        }
                    }
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
        default:
            switch 128 {
            case 127:
                v6 += (v6 >> ((255) & 63))
            }
        }
        v63 := GenS7{F0: 42, F1: 255, F2: -32768}
        _ = v20
        _ = v63
    }
    if true {
        for fi64 := 0; fi64 < 8; fi64++ {
            if false {
                v3 -= (v3 % ((-128) | 1))
                fmt.Println(v5, v3, v3)
                var v65 float32 = 0.1
                _ = v65
            } else {
                for fi66 := 0; fi66 < 1; fi66++ {
                    if (993 != 42) {
                        fmt.Println(v6, v3, v6)
                    }
                    v67 := GenS8{F0: 428, F1: "", F2: "foo bar"}
                    fmt.Println(v6, v3, fi66)
                    _ = fi66
                    _ = v67
                }
            }
            switch "foo bar" {
            case "":
                var v68 int8 = (v3 % (((v3 * (v3 / (((v3 - (v3 / ((42) | 1)))) | 1)))) | 1))
                _ = v68
            case "\n\t":
                if (false && (v2 == 0)) {
                    v2 += v6
                    v5 = (65535 >= uint16(v2))
                } else {
                    fmt.Println(v3)
                    switch "" {
                    case "hello":
                        var v69 float32 = (3.14159 * 1.5)
                        _ = v69
                    case "a":
                        v70 := v4["foo bar"]
                        _ = v70
                    case "foo bar":
                        v5 = true
                    }
                }
            default:
                v71 := GenS6{F0: 42, F1: 2, F2: "foo bar"}
                _ = v71
            }
            if true {
                v72 := GenS9{F0: (fi64 - fi64)}
                var v73 int32 = -45
                _ = v72
                _ = v73
            }
            _ = fi64
        }
    } else {
        if (323 > 1) {
            switch "a" {
            case "hello":
                v6 = v2
            default:
                var v74 bool = true
                _ = v74
            }
            v5 = v5
            var v75 uint8 = 128
            _ = v75
        } else {
            var v76 uint32 = 2147483647
            _ = v76
        }
        v77 := GenS10{F0: -2147483648, F1: (v3 % (((v3 & (v3 >> ((255) & 63)))) | 1)), F2: 65535}
        if v5 {
            if false {
                for fi78 := 0; fi78 < 16; fi78++ {
                    v79 := GenS11{F0: 1, F1: 1, F2: (!v5)}
                    v80 := GenS1{F0: ("" + ("foo bar" + ("hello" + "\n\t"))), F1: 1, F2: int32(v6)}
                    _ = fi78
                    _ = v79
                    _ = v80
                }
                var v81 uint32 = 42
                _ = v81
            }
            var v82 int64 = -110
            _ = v82
        }
        _ = v77
    }
    v3 += v3
    for fi83 := 0; fi83 < 16; fi83++ {
        fmt.Println(v2)
        if (v5 && v5) {
            v5 = v5
            switch ("a" + "a") {
            case "a":
                var v84 string = "a"
                _ = v84
            case "foo bar":
                var v85 uint64 = 651
                _ = v85
            }
        }
        _ = fi83
    }
    fmt.Println(v6, v2)
    fmt.Println(v6, v5)
    var v86 uint32 = 1
    var v87 bool = (false || true)
    if v87 {
        v6 = (v2 ^ v2)
    } else {
        if true {
            switch 4294967295 {
            case 128:
                v88 := GenS6{F0: 42, F1: 2147483647, F2: "foo bar"}
                _ = v88
            default:
                var v89 float32 = float32(3.14159)
                _ = v89
            }
            v90 := GenS12{F0: (v3 >> ((2) & 63))}
            var v91 int16 = 1
            _ = v90
            _ = v91
        }
    }
    var v92 uint = 42
    var v93 bool = (!(!(!(v3 < (v3 / (((v3 + v3)) | 1))))))
    v4[("\n\t" + ("" + "\n\t"))] = (!(255 < 175))
    var v94 int8 = v3
    if false {
        fmt.Println(v92, v6, v92)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch 2147483647 {
    case 32768:
        v95 := GenS10{F0: 769, F1: v3, F2: v86}
        _ = v95
    }
    if false {
        v96 := GenS1{F0: ("foo bar" + ("foo bar" + (("a" + "foo bar") + "foo bar"))), F1: 2147483647, F2: -1}
        v3 += 127
        v4[(((("" + "\n\t") + ("a" + "a")) + "a") + "hello")] = (429 <= -522)
        _ = v96
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v86
    _ = v87
    _ = v92
    _ = v93
    _ = v94
}
