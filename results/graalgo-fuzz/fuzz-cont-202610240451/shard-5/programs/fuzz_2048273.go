package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int; F2 string }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 string; F2 bool }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS10 struct { F0 int64; F1 int16 }
type GenS11 struct { F0 int16 }
func genFn0(p0 int8, p1 uint8) int16 {
    fmt.Println(p0, p0)
    v4 := GenS1{F0: 42, F1: 2147483647}
    p0 += (p0 | (p0 - 100))
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 uint8) uint16 {
    var v5 string = ("a" + ("\n\t" + ((("\n\t" + ((("hello" + ("hello" + ("foo bar" + "a"))) + "hello") + "foo bar")) + ("\n\t" + "\n\t")) + "\n\t")))
    fmt.Println(v5, p0)
    v6 := GenS2{F0: 697, F1: 2147483647, F2: 42}
    _ = p0
    _ = v5
    _ = v6
    return 2
}

func main() {
    v1 := []int{4294967295, 4294967295}
    var v2 uint16 = 65535
    v3 := make(map[int]bool)
    _ = len(v1)
    switch "a" {
    case "\n\t":
        var v7 int8 = 1
        _ = v7
    case "hello":
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
    default:
        v8 := GenS3{F0: 699, F1: 0, F2: "hello"}
        _ = v8
    }
    v9 := GenS4{F0: true, F1: genFn0(127, 73)}
    v2 -= v2
    var v10 float64 = float64(0.1)
    if true {
        if len(v1) > 0 {
            v1[0] = -1
        }
        v2 = (v2 + v2)
        if true {
            var v11 bool = true
            _ = v11
        } else {
            var v12 int64 = 0
            for fi13 := 0; fi13 < 3; fi13++ {
                v10 -= v10
                v10 = v10
                _ = fi13
            }
            _ = v12
        }
    }
    if (false || true) {
        v14 := GenS5{F0: (v2 <= v2), F1: 4294967295}
        var v15 uint64 = 0
        var v16 int = 760
        _ = v14
        _ = v15
        _ = v16
    } else {
        var v17 uint8 = 128
        if false {
            v2 -= v2
        } else {
            v18 := GenS6{F0: 2}
            switch -9223372036854775808 {
            case 392:
                var v19 int16 = -32768
                _ = v19
            case 1:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v20 int64 = 679
                _ = v20
            }
            _ = v18
        }
        _ = v17
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        if false {
            var v22 bool = (("foo bar" + "a") < ("foo bar" + "hello"))
            if (127 != -1) {
                for fi23 := 0; fi23 < 13; fi23++ {
                    var v24 uint32 = 4294967295
                    fmt.Println(v24, fi23)
                    if v22 {
                        if (v22 || v22) {
                            delete(v3, (fi23 / ((fi23) | 1)))
                            v2 = (v2 ^ (v2 % ((genFn1(42)) | 1)))
                            var v25 int32 = -573
                            _ = v25
                        } else {
                            v24 -= 2147483648
                            if v22 {
                                switch (fi21 % (((fi23 + (fi23 << ((42) & 63)))) | 1)) {
                                case 127:
                                    if ((!(316 >= 2)) && v22) {
                                        v2 = (v2 % ((0) | 1))
                                    }
                                }
                            } else {
                                var v26 int = (fi23 >> ((65535) & 63))
                                switch "a" {
                                case "a":
                                    fmt.Println(fi23)
                                case "":
                                    var v27 int16 = -578
                                    _ = v27
                                case "hello":
                                    var v28 float32 = float32((float64((-2.25 / 1.0)) - float64((0.1 + float32(float64(0.0))))))
                                    _ = v28
                                default:
                                    var v29 uint = 1
                                    _ = v29
                                }
                                _ = v26
                            }
                        }
                        switch 0 {
                        case -704:
                            fmt.Println(v10)
                        case -2147483648:
                            v2 = 42
                            var v30 int64 = 0
                            _ = v30
                        case 529:
                            if false {
                                _ = len(v1)
                                var v31 uint = 2
                                if ("a" != ("a" + ("hello" + ""))) {
                                    v24 -= 1
                                    var v32 bool = v22
                                    var v33 uint64 = 0
                                    _ = v32
                                    _ = v33
                                } else {
                                    v34 := GenS2{F0: 42, F1: 384, F2: (v24 & (^(v24)))}
                                    var v35 int = -864
                                    v36 := GenS5{F0: v22, F1: v24}
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                _ = v31
                            }
                        }
                        if v22 {
                            if (2147483647 == 437) {
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                v3[(-(fi23))] = ((("" + ("foo bar" + "foo bar")) + ((("a" + "") + "foo bar") + ("foo bar" + "a"))) < "\n\t")
                            } else {
                                if false {
                                    if ((!v22) || v22) {
                                        v24 = 727
                                    }
                                    var v37 uint = 995
                                    var v38 int8 = 42
                                    var v39 uint64 = 2
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                } else {
                                    if v22 {
                                        switch (fi21 + fi23) {
                                        case -156:
                                            if v22 {
                                                if v22 {
                                                    fmt.Println(v2)
                                                }
                                                switch ("\n\t" + "hello") {
                                                case "":
                                                    delete(v3, fi21)
                                                }
                                            }
                                        }
                                        var v40 float64 = v10
                                        v41 := GenS7{F0: (^(v2)), F1: ((((("foo bar" + "") + ("hello" + "foo bar")) + "") + "foo bar") + (("" + "") + ("a" + "hello"))), F2: (127 > 1)}
                                        var v42 int32 = int32(fi23)
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                    var v43 int = (fi21 / ((128) | 1))
                                    _ = v43
                                }
                            }
                            switch ((("hello" + "\n\t") + "foo bar") + "hello") {
                            case "foo bar":
                                fmt.Println(fi23, fi21)
                            case "hello":
                                v2 = v2
                            case "a":
                                if true {
                                    var v44 uint32 = v24
                                    fmt.Println(v2, fi23)
                                    v45 := GenS8{F0: 919, F1: 255}
                                    _ = v44
                                    _ = v45
                                }
                            }
                            fmt.Println(v10, fi21)
                        } else {
                            var v46 float32 = 1e10
                            _ = v46
                        }
                    }
                    v3[757] = (v22 || v22)
                    _ = fi23
                    _ = v24
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                delete(v3, (fi21 / (((-(fi21))) | 1)))
            } else {
                var v47 float32 = (3.14159 * -2.25)
                switch "a" {
                case "hello":
                    switch fi21 {
                    case 38:
                        for fi48 := 0; fi48 < 4; fi48++ {
                            var v49 int16 = 857
                            if true {
                                delete(v3, (fi48 >> ((65535) & 63)))
                                v22 = v22
                                v49 += v49
                            } else {
                                v50 := GenS8{F0: (fi21 * (fi21 - (fi21 + fi48))), F1: 0}
                                if (true || (2147483647 < 46)) {
                                    switch fi21 {
                                    case -2147483648:
                                        if len(v1) > 0 {
                                            v1[0] = 2147483647
                                        }
                                    case 0:
                                        switch (("" + ("\n\t" + ("hello" + (("foo bar" + "a") + "foo bar")))) + ("foo bar" + "foo bar")) {
                                        case "":
                                            if (v22 && ((-(v49)) < genFn0(-125, 1))) {
                                                v22 = ("foo bar" >= "foo bar")
                                                var v51 int = fi21
                                                v22 = (1000 == (v49 << ((1) & 63)))
                                                v51 -= fi48
                                                switch fi48 {
                                                case -1:
                                                    switch -1 {
                                                    case 128:
                                                        switch ("a" + "") {
                                                        case "hello":
                                                            switch ((((((("foo bar" + "\n\t") + "") + ("\n\t" + ("a" + "hello"))) + "hello") + "") + "foo bar") + "hello") {
                                                            case "foo bar":
                                                                var v52 string = ""
                                                                _ = v52
                                                            default:
                                                                delete(v3, (fi21 ^ 4294967295))
                                                            }
                                                        }
                                                    default:
                                                        if true {
                                                            if (!v22) {
                                                                fmt.Println(v2, v47, v47)
                                                            } else {
                                                                if v22 {
                                                                    var v53 uint64 = 2
                                                                    v47 += 1.5
                                                                    _ = v53
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        fmt.Println(v10, v2)
                                                                    }
                                                                    v54 := GenS5{F0: (v47 == ((v47 + v47) * (v47 / v47))), F1: 42}
                                                                    v55 := GenS1{F0: 435, F1: 1}
                                                                    v56 := GenS8{F0: v51, F1: 1}
                                                                    v47 = float32(float64(float32(v10)))
                                                                    _ = v54
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    if false {
                                                                        switch (v51 >> ((396) & 63)) {
                                                                        case 127:
                                                                            v57 := GenS9{F0: 1, F1: (v2 * v2), F2: -9223372036854775808}
                                                                            _ = v57
                                                                        case -468:
                                                                            delete(v3, (fi48 | (fi21 | (fi48 + (fi48 >> ((4294967295) & 63))))))
                                                                        case 282:
                                                                            var v58 float32 = v47
                                                                            _ = v58
                                                                        default:
                                                                            v22 = true
                                                                        }
                                                                        fmt.Println(v2, v10)
                                                                        v1 = append(v1, fi48)
                                                                    }
                                                                default:
                                                                    if (-128 > 1) {
                                                                        if (!v22) {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                        var v59 uint8 = 127
                                                                        var v60 int8 = -128
                                                                        v61 := GenS7{F0: (v2 % (((v2 / ((32767) | 1))) | 1)), F1: "\n\t", F2: false}
                                                                        var v62 uint16 = v2
                                                                        v51 = (fi48 & v51)
                                                                        _ = v59
                                                                        _ = v60
                                                                        _ = v61
                                                                        _ = v62
                                                                    } else {
                                                                        v47 -= (-2.25 * float32(v10))
                                                                    }
                                                                }
                                                                if v22 {
                                                                    var v63 int = 127
                                                                    _ = v63
                                                                } else {
                                                                    var v64 int16 = (v49 ^ genFn0(42, 0))
                                                                    v2 += v2
                                                                    v65 := GenS6{F0: 4294967295}
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                            }
                                                        } else {
                                                            switch 1 {
                                                            case -2147483648:
                                                                v49 = -32768
                                                            default:
                                                                v66 := GenS4{F0: false, F1: v49}
                                                                _ = v66
                                                            }
                                                            var v67 int64 = int64(v51)
                                                            v67 += v67
                                                            _ = v67
                                                        }
                                                    }
                                                case -178:
                                                    v10 -= (v10 / ((-2.25 * (v10 / (v10 * 0.0))) * ((v10 * v10) * 1e10)))
                                                }
                                                delete(v3, v51)
                                                var v68 string = "a"
                                                if true {
                                                    v2 += v2
                                                    switch (v51 ^ (fi21 | fi48)) {
                                                    case 32767:
                                                        v2 += 32767
                                                    case 128:
                                                        v69 := GenS10{F0: 1, F1: -1}
                                                        _ = v69
                                                    }
                                                }
                                                _ = v51
                                                _ = v68
                                            }
                                        case "foo bar":
                                            fmt.Println(fi21)
                                        }
                                    default:
                                        if (15 == 52) {
                                            v2 = 0
                                            var v70 uint16 = 190
                                            _ = v70
                                        }
                                    }
                                }
                                var v71 uint16 = (v2 >> ((42) & 63))
                                _ = v50
                                _ = v71
                            }
                            v47 = (1.5 / v47)
                            fmt.Println(v47, v10, v22)
                            _ = fi48
                            _ = v49
                        }
                    case 4294967295:
                        if len(v1) > 0 {
                            v1[0] = (fi21 | 127)
                        }
                    default:
                        var v72 uint16 = v2
                        _ = v72
                    }
                case "a":
                    for fi73 := 0; fi73 < 5; fi73++ {
                        fmt.Println(v2, v47)
                        _ = fi73
                    }
                case "foo bar":
                    v10 -= (((v10 * v10) - v10) - v10)
                }
                _ = v47
            }
            v74 := GenS11{F0: 904}
            _ = v22
            _ = v74
        }
        _ = fi21
    }
    v3[32768] = (false || true)
    var v75 uint = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v75:", v75)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v75
    _ = v9
}
