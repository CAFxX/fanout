package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int64; F2 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 int; F2 int16 }
type GenS8 struct { F0 uint16; F1 int32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int; F1 uint8; F2 uint16 }
func genFn0() int64 {
    var v7 uint64 = 2
    var v8 int = 127
    v9 := []int{}
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 int16 = 1000
    var v2 uint64 = 18446744073709551615
    var v3 float64 = 3.14159
    v4 := make(map[string]uint32)
    var v5 int32 = -1
    var v6 int8 = 100
    v10 := GenS1{F0: false, F1: false, F2: -1}
    fmt.Println(v3, v6, v5)
    switch "\n\t" {
    case "\n\t":
        v2 += v2
    case "hello":
        v5 += v5
    case "":
        v11 := []int{}
        _ = v11
    }
    v2 = v2
    fmt.Println(v2, v3)
    for fi12 := 0; fi12 < 11; fi12++ {
        switch "foo bar" {
        case "hello":
            switch "a" {
            case "foo bar":
                var v13 uint = 42
                _ = v13
            default:
                switch fi12 {
                case -9223372036854775808:
                    v3 = float64(0.1)
                case -909:
                    v6 += -1
                default:
                    switch "" {
                    case "a":
                        v1 = 42
                    case "\n\t":
                        if true {
                            fmt.Println(v5, fi12)
                            var v14 int16 = 32767
                            _ = v14
                        } else {
                            if true {
                                v4["\n\t"] = 4294967295
                                var v15 int16 = (v1 % (((v1 & v1)) | 1))
                                _ = v15
                            } else {
                                v2 += (v2 >> ((0) & 63))
                                var v16 uint32 = 4294967295
                                _ = v16
                            }
                            v6 -= (v6 & v6)
                        }
                    case "hello":
                        v17 := GenS2{F0: -1}
                        _ = v17
                    }
                }
            }
        case "foo bar":
            if true {
                v18 := []int{int(v2), (fi12 * fi12)}
                if ((4294967295 < 42) && true) {
                    if ((fi12 == int(v6)) || true) {
                        v19 := GenS2{F0: -9223372036854775808}
                        var v20 uint16 = 257
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v5 -= int32(v2)
                            fmt.Println(fi21, fi21, v3)
                            switch fi21 {
                            case 32768:
                                if true {
                                    fmt.Println(v20)
                                }
                            default:
                                v22 := GenS2{F0: 677}
                                _ = v22
                            }
                            _ = fi21
                        }
                        _ = v19
                        _ = v20
                    }
                    switch "foo bar" {
                    case "":
                        v2 = 0
                    case "a":
                        switch ("foo bar" + (("foo bar" + ("foo bar" + (((("" + ("a" + ("foo bar" + ("foo bar" + "\n\t")))) + (("a" + ("a" + ((((("hello" + "a") + "\n\t") + "a") + ("\n\t" + "\n\t")) + "a"))) + "")) + "hello") + ("" + "\n\t")))) + "foo bar")) {
                        case "hello":
                            var v23 float32 = ((1.5 / (float32(v3) - (0.0 + -2.25))) - (0.1 - (0.0 / 1e10)))
                            _ = v23
                        }
                    case "\n\t":
                        switch "a" {
                        case "":
                            if false {
                                v24 := GenS1{F0: true, F1: (!false), F2: fi12}
                                _ = v24
                            } else {
                                _ = len(v4)
                            }
                        case "a":
                            for fi25 := 0; fi25 < 19; fi25++ {
                                v6 += -128
                                v26 := v4[""]
                                switch "a" {
                                case "a":
                                    fmt.Println(v6, v1)
                                case "\n\t":
                                    var v27 string = "hello"
                                    _ = v27
                                case "hello":
                                    switch "\n\t" {
                                    case "a":
                                        v28 := GenS1{F0: (!true), F1: false, F2: (fi25 + (^(fi12)))}
                                        _ = v28
                                    case "hello":
                                        delete(v4, ("a" + ""))
                                    case "\n\t":
                                        if (fi12 == -2147483648) {
                                            v29 := GenS1{F0: ((!((((v5 < (v5 - -2147483648)) || true) || (1 != 200)) || ((-1 <= 0) && true))) || false), F1: false, F2: fi25}
                                            if true {
                                                v26 += (v26 - 2147483647)
                                            } else {
                                                v26 = 1
                                            }
                                            switch fi12 {
                                            case -123:
                                                v30 := GenS2{F0: -349}
                                                _ = v30
                                            }
                                            _ = v29
                                        }
                                    }
                                default:
                                    delete(v4, "a")
                                }
                                if (v1 != (v1 % (((v1 + v1)) | 1))) {
                                    v31 := v4["foo bar"]
                                    v32 := GenS3{F0: true, F1: "hello"}
                                    _ = v31
                                    _ = v32
                                } else {
                                    var v33 int64 = genFn0()
                                    var v34 string = "hello"
                                    fmt.Println(v1)
                                    v26 -= v26
                                    _ = v33
                                    _ = v34
                                }
                                _ = fi25
                                _ = v26
                            }
                        case "hello":
                            v35 := GenS4{F0: 1}
                            _ = v35
                        }
                    default:
                        var v36 int = 601
                        _ = v36
                    }
                    var v37 uint8 = uint8(v2)
                    _ = v37
                } else {
                    var v38 int = (-(fi12))
                    switch (("" + (("foo bar" + "") + (("foo bar" + "\n\t") + "a"))) + "") {
                    case "":
                        if false {
                            v2 += v2
                            var v39 uint16 = 255
                            _ = v39
                        } else {
                            switch ((("a" + ("" + ("foo bar" + "\n\t"))) + (("\n\t" + "a") + ("\n\t" + ("a" + "a")))) + "foo bar") {
                            case "a":
                                if false {
                                    for fi40 := 0; fi40 < 7; fi40++ {
                                        if false {
                                            v41 := GenS3{F0: false, F1: "\n\t"}
                                            var v42 uint32 = 2147483648
                                            _ = v41
                                            _ = v42
                                        }
                                        _ = fi40
                                    }
                                } else {
                                    if (v3 > v3) {
                                        var v43 uint64 = 1
                                        for fi44 := 0; fi44 < 19; fi44++ {
                                            if true {
                                                v45 := GenS4{F0: 1}
                                                if (true || true) {
                                                    var v46 int32 = v5
                                                    _ = len(v18)
                                                    _ = v46
                                                } else {
                                                    v3 += v3
                                                }
                                                if true {
                                                    if (!true) {
                                                        var v47 uint = 0
                                                        v48 := GenS2{F0: 811}
                                                        var v49 uint8 = 42
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    v50 := GenS5{F0: true, F1: 238, F2: 0}
                                                    _ = v50
                                                }
                                                v51 := GenS6{F0: "foo bar"}
                                                _ = v45
                                                _ = v51
                                            } else {
                                                var v52 int32 = 42
                                                fmt.Println(v5, v38, v6)
                                                if (true && ((v38 > (fi44 >> ((42) & 63))) && false)) {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        fmt.Println(v38, v52, v52)
                                                    case "":
                                                        if true {
                                                            if len(v18) > 0 {
                                                                v18[0] = (-(v38))
                                                            }
                                                            v6 += (-(v6))
                                                            switch (v38 | -31) {
                                                            case 127:
                                                                switch ("\n\t" + "\n\t") {
                                                                case "":
                                                                    v3 = 0.0
                                                                case "a":
                                                                    v53 := v4["hello"]
                                                                    _ = v53
                                                                }
                                                            }
                                                        } else {
                                                            if (false || true) {
                                                                var v54 int = v38
                                                                fmt.Println(v54)
                                                                _ = v54
                                                            } else {
                                                                var v55 uint32 = 2147483647
                                                                if false {
                                                                    v38 = (fi12 * -623)
                                                                } else {
                                                                    v2 -= v43
                                                                }
                                                                _ = v55
                                                            }
                                                            if false {
                                                                v56 := GenS1{F0: (!false), F1: true, F2: (fi44 | 2147483647)}
                                                                var v57 int64 = -998
                                                                _ = v56
                                                                _ = v57
                                                            } else {
                                                                v58 := GenS3{F0: (1 > 240), F1: ((("\n\t" + "a") + "foo bar") + "foo bar")}
                                                                _ = len(v4)
                                                                if (true && false) {
                                                                    fmt.Println(v1, v6, v38)
                                                                } else {
                                                                    if (true && (752 >= 2)) {
                                                                        v59 := GenS7{F0: (!false), F1: 526, F2: v1}
                                                                        _ = v59
                                                                    } else {
                                                                        var v60 float64 = float64(1.5)
                                                                        if (0 > 27) {
                                                                            var v61 bool = true
                                                                            if v61 {
                                                                                delete(v4, "foo bar")
                                                                                fmt.Println(v6, v60, v38)
                                                                            } else {
                                                                                switch "" {
                                                                                case "":
                                                                                    var v62 int64 = -1
                                                                                    _ = v62
                                                                                }
                                                                                var v63 uint32 = 1
                                                                                var v64 uint32 = v63
                                                                                _ = v63
                                                                                _ = v64
                                                                            }
                                                                            v38 -= (fi12 * fi44)
                                                                            var v65 int32 = 2147483647
                                                                            if len(v18) >= 2 {
                                                                                v18 = v18[1:len(v18)]
                                                                            }
                                                                            _ = v61
                                                                            _ = v65
                                                                        }
                                                                        v1 += v1
                                                                        _ = v60
                                                                    }
                                                                    v66 := GenS8{F0: 0, F1: (v52 + v52)}
                                                                    _ = v66
                                                                }
                                                                _ = v58
                                                            }
                                                            v38 = fi44
                                                            if false {
                                                                v1 = v1
                                                            }
                                                        }
                                                    default:
                                                        var v67 int32 = v52
                                                        _ = v67
                                                    }
                                                    v52 -= int32(v2)
                                                    var v68 uint32 = 2147483648
                                                    var v69 string = ((("\n\t" + ("\n\t" + (("foo bar" + ("\n\t" + "\n\t")) + ("" + "")))) + ("a" + ("hello" + "hello"))) + "\n\t")
                                                    var v70 float32 = float32(v3)
                                                    _ = v68
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    v18 = append(v18, -1)
                                                    v71 := GenS5{F0: (1 != 473), F1: -1, F2: 2}
                                                    v72 := v4[(("" + (("" + "hello") + "")) + "foo bar")]
                                                    _ = v71
                                                    _ = v72
                                                }
                                                if false {
                                                    v73 := GenS3{F0: (-431 < v5), F1: "foo bar"}
                                                    v74 := GenS9{F0: -509}
                                                    _ = v73
                                                    _ = v74
                                                }
                                                _ = v52
                                            }
                                            _ = fi44
                                        }
                                        switch v38 {
                                        case 4294967295:
                                            for fi75 := 0; fi75 < 8; fi75++ {
                                                v18 = append(v18, (fi12 | 127))
                                                _ = fi75
                                            }
                                        default:
                                            if true {
                                                v2 -= v2
                                            } else {
                                                delete(v4, "foo bar")
                                                v6 += -1
                                                for fi76 := 0; fi76 < 11; fi76++ {
                                                    var v77 uint = 255
                                                    v6 = v6
                                                    fmt.Println(v38, v43, v5)
                                                    _ = fi76
                                                    _ = v77
                                                }
                                            }
                                        }
                                        _ = v43
                                    }
                                    v4[("hello" + "a")] = 4294967295
                                    switch v38 {
                                    case 1:
                                        if true {
                                            if len(v18) > 0 {
                                                v18[0] = v38
                                            }
                                        } else {
                                            var v78 uint64 = v2
                                            _ = len(v18)
                                            if len(v18) > 0 {
                                                v18[0] = (v38 % ((fi12) | 1))
                                            }
                                            _ = v78
                                        }
                                    case -2147483648:
                                        v4["a"] = 2147483647
                                    default:
                                        v2 = v2
                                    }
                                }
                            default:
                                for fi79 := 0; fi79 < 10; fi79++ {
                                    switch "" {
                                    case "a":
                                        v5 = (v5 >> ((4294967295) & 63))
                                    case "hello":
                                        fmt.Println(v2)
                                    }
                                    _ = fi79
                                }
                            }
                            v80 := GenS6{F0: "\n\t"}
                            _ = v80
                        }
                    }
                    _ = v38
                }
                if (fi12 > -9223372036854775808) {
                    var v81 uint = 4294967295
                    var v82 uint = 255
                    v81 += 2
                    _ = v81
                    _ = v82
                }
                _ = v18
            }
        case "a":
            var v83 uint16 = 1
            _ = v83
        default:
            for fi84 := 0; fi84 < 16; fi84++ {
                v85 := []int{fi84, (fi84 + fi84), (fi84 % (((fi12 + (^(fi12)))) | 1))}
                _ = fi84
                _ = v85
            }
        }
        _ = fi12
    }
    var v86 float64 = v3
    v86 = -2.25
    v87 := GenS10{F0: -879, F1: 158, F2: 1}
    if (true && true) {
        v88 := []int{-2147483648, 32767}
        _ = v88
    }
    if (!(42 <= 0)) {
        if (((137 < v2) && (!true)) && true) {
            switch 127 {
            case 128:
                fmt.Println(v6)
            default:
                delete(v4, "a")
            }
        }
    } else {
        fmt.Println(v2)
        var v89 uint64 = (^(v2))
        v6 = int8(v2)
        _ = v89
    }
    v86 += ((v86 + v3) / (1e10 / v3))
    v3 = v86
    var v90 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v86
    _ = v87
    _ = v90
}
