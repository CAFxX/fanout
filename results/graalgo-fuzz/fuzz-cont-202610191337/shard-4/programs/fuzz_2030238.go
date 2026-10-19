package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 string; F1 uint32; F2 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16; F1 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16; F1 uint32 }
func genFn0() uint {
    if (((!(165 > 127)) && true) && (42 == -632)) {
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 float64 = 0.1
            v7 := GenS1{F0: 42, F1: -29, F2: "\n\t"}
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        if false {
            var v8 int16 = 1000
            v9 := []int{128}
            _ = v8
            _ = v9
        } else {
            var v10 string = ""
            _ = v10
        }
        var v11 uint32 = 65535
        switch "foo bar" {
        case "":
            v12 := []int{-49, 4294967295, 127}
            _ = v12
        case "\n\t":
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := []int{}
                for fi15 := 0; fi15 < 10; fi15++ {
                    fmt.Println(fi15)
                    var v16 int64 = 42
                    switch (("hello" + "") + "hello") {
                    case "\n\t":
                        _ = len(v14)
                    default:
                        v11 = v11
                    }
                    _ = fi15
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
        }
        _ = v11
    }
    var v17 int32 = 0
    fmt.Println(v17)
    _ = v17
    return 42
}
func genFn1() int {
    var v18 uint8 = 1
    var v19 int32 = -2147483648
    _ = v18
    _ = v19
    return -664
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    v20 := GenS2{F0: "hello"}
    switch "hello" {
    case "hello":
        switch "\n\t" {
        case "hello":
            if false {
                var v21 uint32 = 2
                switch 549 {
                case 1:
                    v21 += v21
                case -527:
                    v1 += v1
                }
                fmt.Println(v21, v21)
                _ = v21
            }
        case "a":
            v22 := GenS3{F0: 868, F1: 1, F2: 0}
            _ = v22
        default:
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 int = (fi23 << ((65535) & 63))
                if true {
                    delete(v3, genFn1())
                    for fi25 := 0; fi25 < 20; fi25++ {
                        if false {
                            fmt.Println(fi25, fi23, fi25)
                            if true {
                                v24 = 4294967295
                                v2[1] = (0 >= 127)
                                var v26 uint = 4294967295
                                v27 := GenS2{F0: "foo bar"}
                                if (954 > 1) {
                                    v26 = v26
                                    var v28 uint32 = 1
                                    _ = v28
                                }
                                _ = v26
                                _ = v27
                            } else {
                                fmt.Println(v1, v1)
                            }
                            var v29 uint64 = 2
                            _ = v29
                        } else {
                            switch 231 {
                            case 128:
                                var v30 int32 = -502
                                _ = v30
                            case 254:
                                v31 := GenS4{F0: (("\n\t" + "a") + "\n\t"), F1: 128}
                                _ = v31
                            }
                            v24 += fi23
                            v1 = (^(v1))
                            v1 -= 42
                            fmt.Println(v1, fi25)
                        }
                        fmt.Println(fi23, v1, v1)
                        _ = fi25
                    }
                    if (2147483647 < 0) {
                        var v32 int8 = 42
                        v33 := GenS2{F0: "\n\t"}
                        _ = v32
                        _ = v33
                    } else {
                        var v34 bool = (105 == 375)
                        if (0 <= 1) {
                            if true {
                                for fi35 := 0; fi35 < 6; fi35++ {
                                    if false {
                                        var v36 float32 = 0.0
                                        if (!(true || true)) {
                                            v37 := GenS4{F0: (((("foo bar" + "hello") + "") + "foo bar") + ("foo bar" + "a")), F1: 0}
                                            v24 = genFn1()
                                            v38 := GenS5{F0: -79, F1: "hello"}
                                            _ = v37
                                            _ = v38
                                        } else {
                                            v39 := GenS4{F0: "", F1: 128}
                                            var v40 int32 = 0
                                            _ = v39
                                            _ = v40
                                        }
                                        v1 = (v1 ^ v1)
                                        if (v36 > 1.5) {
                                            switch ("foo bar" + "foo bar") {
                                            case "a":
                                                var v41 int = (fi23 * 128)
                                                _ = v41
                                            default:
                                                var v42 int64 = 183
                                                _ = v42
                                            }
                                        }
                                        switch (v24 * v24) {
                                        case -151:
                                            var v43 int = (fi23 % ((4294967295) | 1))
                                            _ = v43
                                        }
                                        _ = v36
                                    }
                                    v34 = v34
                                    _ = fi35
                                }
                                v44 := GenS5{F0: 127, F1: ("a" + ("foo bar" + ("\n\t" + "a")))}
                                _ = v44
                            } else {
                                v45 := GenS6{F0: ("hello" + (("hello" + "hello") + ("\n\t" + "\n\t"))), F1: 2147483648, F2: 255}
                                _ = v45
                            }
                        }
                        _ = v34
                    }
                }
                for fi46 := 0; fi46 < 12; fi46++ {
                    var v47 int = (fi46 >> ((65535) & 63))
                    var v48 uint16 = 2
                    v49 := v4[632]
                    var v50 int16 = v1
                    v51 := []int{}
                    _ = fi46
                    _ = v47
                    _ = v48
                    _ = v49
                    _ = v50
                    _ = v51
                }
                _ = fi23
                _ = v24
            }
        }
    }
    if (2147483648 != 2147483647) {
        if true {
            fmt.Println(v1)
            v52 := GenS7{F0: 1}
            _ = v52
        } else {
            var v53 int16 = (v1 ^ (^(v1)))
            fmt.Println(v1, v53)
            _ = v53
        }
        switch -9223372036854775808 {
        case 128:
            var v54 int = 1
            _ = v54
        case 32768:
            v55 := GenS6{F0: ("\n\t" + ("\n\t" + "\n\t")), F1: 2147483647, F2: 158}
            _ = v55
        }
    }
    v56 := v4[297]
    var v57 uint64 = 42
    fmt.Println(v57)
    switch v56 {
    case "hello":
        for fi58 := 0; fi58 < 20; fi58++ {
            if (((!true) && true) && true) {
                for fi59 := 0; fi59 < 19; fi59++ {
                    fmt.Println(v1, v1)
                    v1 = v1
                    _ = fi59
                }
                switch fi58 {
                case 42:
                    var v60 float32 = -2.25
                    _ = v60
                case 127:
                    for fi61 := 0; fi61 < 20; fi61++ {
                        fmt.Println(v56, fi58)
                        v56 = v56
                        _ = fi61
                    }
                default:
                    v62 := GenS8{F0: 127}
                    _ = v62
                }
                v63 := GenS2{F0: v56}
                _ = v63
            }
            _ = fi58
        }
    case "":
        v56 = "hello"
    default:
        fmt.Println(v57, v56)
    }
    fmt.Println(v56, v56)
    switch (v56 + "hello") {
    case "\n\t":
        fmt.Println(v1, v56)
    case "":
        for fi64 := 0; fi64 < 16; fi64++ {
            v56 = ((v56 + v56) + "\n\t")
            if true {
                v56 = v56
                v56 = ("\n\t" + (("foo bar" + v56) + "a"))
            } else {
                v57 -= 18446744073709551615
                fmt.Println(fi64, v1)
                switch ("\n\t" + v56) {
                case "":
                    for fi65 := 0; fi65 < 18; fi65++ {
                        if false {
                            if ((1e10 * float32((0.1 + 3.14159))) <= 3.14159) {
                                if (!(-1 == 706)) {
                                    var v66 int = genFn1()
                                    var v67 string = "a"
                                    v68 := []int{-1, 850, (v66 - fi64)}
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                                var v69 float32 = -2.25
                                var v70 uint32 = 1
                                if true {
                                    v71 := GenS9{F0: (v1 * (v1 / ((32767) | 1))), F1: 979}
                                    var v72 int8 = -112
                                    _ = v71
                                    _ = v72
                                }
                                _ = v69
                                _ = v70
                            } else {
                                fmt.Println(v1, fi65)
                                switch (^(fi65)) {
                                case 9223372036854775807:
                                    v56 = v56
                                case 42:
                                    var v73 int16 = -1
                                    _ = v73
                                case -154:
                                    var v74 uint64 = 42
                                    _ = v74
                                }
                                v1 -= v1
                                fmt.Println(fi64, v1)
                            }
                        } else {
                            var v75 bool = (65535 < 4294967295)
                            if (v75 && (v75 || v75)) {
                                if (9223372036854775807 > 539) {
                                    var v76 uint32 = 939
                                    _ = v76
                                } else {
                                    var v77 uint32 = 2147483648
                                    _ = v77
                                }
                            }
                            if v75 {
                                var v78 int = fi65
                                switch (v78 & genFn1()) {
                                case -9223372036854775808:
                                    fmt.Println(v78, v75)
                                case 1:
                                    v79 := GenS10{F0: 65535}
                                    _ = v79
                                case 32767:
                                    v78 += fi65
                                }
                                _ = v78
                            }
                            _ = v75
                        }
                        _ = fi65
                    }
                case "\n\t":
                    fmt.Println(v56, v57)
                case "hello":
                    v1 = 32767
                }
            }
            if false {
                v1 += -936
            }
            _ = fi64
        }
    case "a":
        if false {
            var v80 int64 = -1
            switch -819 {
            case 32767:
                fmt.Println(v1)
            case 185:
                var v81 uint64 = 18446744073709551615
                _ = v81
            }
            _ = v80
        }
    }
    v82 := []int{}
    v57 = (v57 - (v57 - (^(v57))))
    v4[130] = ((("foo bar" + "") + "") + "\n\t")
    if true {
        if true {
            v83 := GenS11{F0: 0, F1: 808}
            _ = v83
        } else {
            for fi84 := 0; fi84 < 8; fi84++ {
                delete(v3, genFn1())
                _ = fi84
            }
        }
    } else {
        for fi85 := 0; fi85 < 7; fi85++ {
            v1 = v1
            var v86 string = "\n\t"
            v82 = append(v82, fi85)
            _ = fi85
            _ = v86
        }
        var v87 uint = 979
        var v88 uint64 = (v57 % ((v57) | 1))
        _ = v87
        _ = v88
    }
    var v89 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v82:", len(v82), v82)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v82
    _ = v89
}
