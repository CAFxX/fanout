package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 int }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 bool; F1 uint64; F2 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS10 struct { F0 int32; F1 uint64; F2 int16 }
func genFn0(p0 uint8, p1 int8) bool {
    v6 := GenS1{F0: -1, F1: 32767}
    p1 -= 100
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 string = "foo bar"
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: -971, F1: 609}
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return false
}
func genFn1() uint {
    v10 := GenS1{F0: 21, F1: 42}
    var v11 bool = false
    _ = v10
    _ = v11
    return 641
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]int)
    var v3 float64 = 3.14159
    var v4 uint8 = 1
    var v5 int64 = 694
    var v12 int64 = (^(v5))
    if ((v5 + v12) != (v5 << ((18446744073709551615) & 63))) {
        var v13 int32 = 42
        _ = v13
    }
    if false {
        if true {
            var v14 uint = 772
            v15 := GenS2{F0: 32768, F1: v4, F2: 255}
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 8; fi16++ {
                for fi17 := 0; fi17 < 13; fi17++ {
                    if (!(42 >= v4)) {
                        v18 := GenS2{F0: fi16, F1: 0, F2: 255}
                        _ = v18
                    }
                    if (false || true) {
                        if (2 > 255) {
                            v1 -= 1
                            if false {
                                if (false && true) {
                                    v19 := []int{(fi16 ^ (fi16 & 528))}
                                    var v20 float64 = 0.0
                                    var v21 float32 = 1e10
                                    if (false || true) {
                                        var v22 int16 = 42
                                        if (!false) {
                                            v23 := GenS3{F0: 2, F1: (^(v22)), F2: v1}
                                            v21 = (v21 + 0.0)
                                            v5 -= (v5 * v5)
                                            _ = v23
                                        } else {
                                            fmt.Println(fi17)
                                            v21 = ((float32((v3 + (0.0 + v3))) / v21) * v21)
                                        }
                                        _ = v22
                                    }
                                    fmt.Println(v21, fi17)
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                            } else {
                                if false {
                                    if (true || (("foo bar" + "a") >= "foo bar")) {
                                        v3 = (v3 + v3)
                                    }
                                    _ = len(v2)
                                } else {
                                    fmt.Println(v5, v1)
                                }
                                v24 := GenS2{F0: fi16, F1: 170, F2: 2}
                                v25 := GenS4{F0: int64(fi16), F1: (v1 << ((344) & 63)), F2: (fi16 << ((genFn1()) & 63))}
                                _ = v24
                                _ = v25
                            }
                        } else {
                            var v26 uint = genFn1()
                            v27 := GenS5{F0: v12, F1: 1000}
                            var v28 int32 = -877
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        var v29 int8 = 127
                        var v30 int8 = (v29 / ((120) | 1))
                        switch "\n\t" {
                        case "a":
                            v29 += v30
                        case "":
                            v31 := []int{128}
                            _ = v31
                        case "hello":
                            v32 := GenS6{F0: ((^(fi16)) < fi16), F1: 18446744073709551615, F2: 1}
                            _ = v32
                        default:
                            switch (("a" + "foo bar") + "") {
                            case "\n\t":
                                fmt.Println(v12, v12)
                            default:
                                v30 = (v29 << ((4294967295) & 63))
                            }
                        }
                        _ = v29
                        _ = v30
                    } else {
                        if (248 >= 516) {
                            if ((2 > 766) || (true || (((("" + ("hello" + "\n\t")) + "") + "") >= ("foo bar" + ("foo bar" + "a"))))) {
                                if true {
                                    switch 653 {
                                    case 1:
                                        var v33 uint32 = 2
                                        _ = v33
                                    case 2147483647:
                                        var v34 uint32 = 2147483647
                                        _ = v34
                                    case 32767:
                                        _ = len(v2)
                                    default:
                                        v35 := GenS2{F0: 748, F1: 168, F2: 2}
                                        _ = v35
                                    }
                                    if ((v12 - 9223372036854775807) >= 42) {
                                        if true {
                                            var v36 uint8 = v4
                                            if ((!false) && (1 >= 2147483648)) {
                                                if ((!true) || false) {
                                                    v37 := v2["a"]
                                                    if false {
                                                        fmt.Println(v12, fi17, fi16)
                                                        v38 := []int{fi16}
                                                        var v39 int64 = -1
                                                        var v40 uint8 = (v4 % ((42) | 1))
                                                        v5 = v5
                                                        v37 -= 128
                                                        _ = v38
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                    if true {
                                                        v1 -= 4294967295
                                                        switch (fi16 - (fi16 & fi17)) {
                                                        case 0:
                                                            _ = len(v2)
                                                        }
                                                        fmt.Println(v36, v4)
                                                    }
                                                    var v41 string = (("hello" + (("" + "foo bar") + "hello")) + ("\n\t" + ("\n\t" + "a")))
                                                    _ = v37
                                                    _ = v41
                                                } else {
                                                    v3 = (1.5 / 1.0)
                                                    fmt.Println(v12)
                                                    v42 := v2["hello"]
                                                    _ = v42
                                                }
                                            }
                                            _ = v36
                                        } else {
                                            v43 := GenS7{F0: 1}
                                            delete(v2, ("hello" + "foo bar"))
                                            var v44 uint64 = (v1 / (((v1 - 18446744073709551615)) | 1))
                                            v4 -= v4
                                            var v45 int16 = 0
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        }
                                        v46 := []int{(-(fi17))}
                                        fmt.Println(fi17, fi16)
                                        _ = v46
                                    } else {
                                        var v47 bool = false
                                        _ = v47
                                    }
                                    var v48 float32 = 3.14159
                                    _ = v48
                                }
                                switch (fi16 | fi16) {
                                case 1:
                                    delete(v2, "foo bar")
                                case 0:
                                    if false {
                                        var v49 int16 = 1000
                                        fmt.Println(v12)
                                        _ = v49
                                    } else {
                                        v3 -= 0.1
                                        _ = len(v2)
                                    }
                                case 4294967295:
                                    v4 += 193
                                default:
                                    var v50 uint = 4294967295
                                    _ = v50
                                }
                            } else {
                                var v51 float32 = 0.0
                                v52 := GenS6{F0: (255 < 1), F1: (v1 * (v1 - 42)), F2: 0}
                                var v53 int8 = 1
                                v2[""] = fi17
                                var v54 float32 = v51
                                _ = v51
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                            var v55 string = (("foo bar" + "a") + "a")
                            fmt.Println(v1, fi16)
                            v5 = (v5 | (^(v12)))
                            _ = v55
                        } else {
                            v56 := GenS6{F0: true, F1: (v1 | (^(v1))), F2: -1}
                            _ = v56
                        }
                        v57 := []int{fi16, 4294967295}
                        _ = v57
                    }
                    fmt.Println(v4)
                    _ = fi17
                }
                _ = fi16
            }
            if (2 == 258) {
                v58 := []int{9223372036854775807, -819, 42}
                _ = v58
            }
        }
        for fi59 := 0; fi59 < 15; fi59++ {
            switch "foo bar" {
            case "\n\t":
                switch fi59 {
                case 706:
                    switch fi59 {
                    case -2147483648:
                        v60 := []int{fi59, 509}
                        _ = v60
                    case 4294967295:
                        var v61 int64 = 42
                        _ = v61
                    default:
                        for fi62 := 0; fi62 < 3; fi62++ {
                            var v63 int16 = 1000
                            v1 = 4294967295
                            _ = fi62
                            _ = v63
                        }
                    }
                case 128:
                    switch ("\n\t" + ("hello" + "")) {
                    case "":
                        v1 = (v1 ^ 744)
                    case "hello":
                        if (true || true) {
                            v64 := GenS7{F0: int64(v4)}
                            v3 += v3
                            _ = v64
                        } else {
                            var v65 int8 = 42
                            _ = v65
                        }
                    case "\n\t":
                        v66 := []int{fi59}
                        _ = v66
                    default:
                        if true {
                            v3 += 3.14159
                        } else {
                            var v67 float32 = (1.5 * (-2.25 - (-2.25 * 3.14159)))
                            _ = v67
                        }
                    }
                }
            case "hello":
                if false {
                    v5 -= v5
                    var v68 uint8 = (v4 | v4)
                    var v69 int16 = 0
                    _ = v68
                    _ = v69
                }
            }
            var v70 uint32 = 4294967295
            var v71 string = "a"
            _ = fi59
            _ = v70
            _ = v71
        }
        var v72 int8 = int8(v12)
        _ = v72
    }
    var v73 int = 4294967295
    var v74 int16 = 0
    for fi75 := 0; fi75 < 6; fi75++ {
        v76 := GenS6{F0: (!(949 >= 4294967295)), F1: 18446744073709551615, F2: -684}
        var v77 uint32 = 42
        v78 := []int{fi75, -166, (fi75 << ((255) & 63)), v73}
        _ = fi75
        _ = v76
        _ = v77
        _ = v78
    }
    v74 = (v74 | 1)
    v79 := GenS3{F0: (v4 + v4), F1: v74, F2: v1}
    v5 += (v5 - v5)
    fmt.Println(v12, v12)
    switch 128 {
    case 458:
        var v80 uint64 = v1
        _ = v80
    case -1:
        if (23 < 0) {
            delete(v2, "a")
            fmt.Println(v4)
        }
    case 127:
        switch (v73 % ((v73) | 1)) {
        case -1:
            v81 := GenS1{F0: 32767, F1: 1}
            _ = v81
        default:
            fmt.Println(v73, v4, v1)
        }
    default:
        v3 += 1.5
    }
    fmt.Println(v1, v3, v12)
    v82 := GenS8{F0: 0, F1: 32767}
    for fi83 := 0; fi83 < 11; fi83++ {
        v84 := GenS9{F0: v4, F1: 0, F2: v1}
        _ = fi83
        _ = v84
    }
    if (!(true && (-2.25 > 0.1))) {
        v3 += (-2.25 - (v3 - (v3 - 1e10)))
        switch ("hello" + "hello") {
        case "foo bar":
            for fi85 := 0; fi85 < 16; fi85++ {
                var v86 float64 = v3
                _ = fi85
                _ = v86
            }
        case "hello":
            var v87 int16 = (-(v74))
            _ = v87
        default:
            var v88 uint64 = (v1 << ((4294967295) & 63))
            _ = v88
        }
    } else {
        v4 += 173
    }
    v89 := GenS10{F0: int32(v74), F1: (^(v1)), F2: -133}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v73
    _ = v74
    _ = v79
    _ = v82
    _ = v89
}
