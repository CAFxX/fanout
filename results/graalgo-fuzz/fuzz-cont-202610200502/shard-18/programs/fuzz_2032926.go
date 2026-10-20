package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint; F2 int8 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint) uint8 {
    if false {
        p0 = (^(p0))
        for fi5 := 0; fi5 < 3; fi5++ {
            p0 += (^(p0))
            p0 = (^(p0))
            _ = fi5
        }
        v6 := []int{-2}
        _ = v6
    } else {
        fmt.Println(p1)
    }
    p1 += 1
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 int64, p1 uint, p2 int32) uint8 {
    if false {
        var v7 uint64 = 0
        _ = v7
    } else {
        var v8 int64 = p0
        p1 -= p1
        _ = v8
    }
    var v9 int = -2147483648
    fmt.Println(p1, p0)
    switch "hello" {
    case "a":
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int8 = 127
            _ = fi10
            _ = v11
        }
    }
    var v12 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return genFn0(1, p1)
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]bool)
    v3 := []string{""}
    v4 := []string{"", "a", ("" + "a")}
    v13 := GenS1{F0: 1}
    if false {
        v2["foo bar"] = (787 == 0)
    } else {
        var v14 uint16 = 101
        var v15 uint64 = 1
        v16 := v2[((("" + "hello") + "foo bar") + ("a" + ("a" + "\n\t")))]
        _ = v14
        _ = v15
        _ = v16
    }
    switch (((("foo bar" + (("hello" + "") + "foo bar")) + "a") + "") + "hello") {
    case "a":
        var v17 int8 = -17
        _ = v17
    case "":
        var v18 string = "\n\t"
        _ = v18
    }
    var v19 string = (("" + "\n\t") + "foo bar")
    var v20 float32 = ((v1 + float32(-2.25)) + (0.0 - 3.14159))
    _ = len(v4)
    var v21 int8 = 1
    delete(v2, (("a" + ("hello" + ("\n\t" + "a"))) + ("" + ("" + "hello"))))
    if (false && (42 != 255)) {
        v1 = v1
    }
    var v22 int64 = -9223372036854775808
    v23 := v2[v19]
    fmt.Println(v23, v20, v23)
    if (!(1e10 != (v20 / v20))) {
        v21 = -56
        if v23 {
            switch "a" {
            case "foo bar":
                v22 = (v22 >> ((18446744073709551615) & 63))
            case "a":
                switch 0 {
                case -1:
                    if ((-32768 > 1000) && true) {
                        fmt.Println(v21, v20, v20)
                        if v23 {
                            fmt.Println(v20, v20, v23)
                        }
                        fmt.Println(v21, v20, v19)
                    } else {
                        v21 = (^(v21))
                        for fi24 := 0; fi24 < 15; fi24++ {
                            for fi25 := 0; fi25 < 14; fi25++ {
                                v23 = v23
                                v26 := GenS2{F0: 65535, F1: (v22 / ((-836) | 1))}
                                _ = fi25
                                _ = v26
                            }
                            _ = fi24
                        }
                        if (v19 > ((v19 + v19) + v19)) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            v20 += 3.14159
                            if (v20 < v1) {
                                var v27 string = ""
                                var v28 uint = 65535
                                if v23 {
                                    fmt.Println(v1)
                                    if false {
                                        if ((((v27 + ((v27 + "") + v19)) + "") + (v27 + "a")) == v27) {
                                            if true {
                                                v27 = "a"
                                                v27 = v19
                                                if len(v3) > 0 {
                                                    v3[0] = v27
                                                }
                                            } else {
                                                v3 = append(v3, v19)
                                                var v29 int16 = 560
                                                _ = v29
                                            }
                                            v1 = 1.5
                                            v30 := v2[v19]
                                            _ = v30
                                        } else {
                                            fmt.Println(v20, v22, v21)
                                            switch ("hello" + "") {
                                            case "foo bar":
                                                var v31 uint = v28
                                                _ = v31
                                            case "hello":
                                                fmt.Println(v19)
                                            case "\n\t":
                                                var v32 int32 = 42
                                                _ = v32
                                            default:
                                                v33 := GenS3{F0: 1, F1: ("hello" + "a")}
                                                _ = v33
                                            }
                                        }
                                        switch 9223372036854775807 {
                                        case -354:
                                            switch v19 {
                                            case "hello":
                                                v1 = 1.5
                                            case "":
                                                v27 = v27
                                            case "foo bar":
                                                var v34 int = 872
                                                _ = v34
                                            default:
                                                v35 := GenS4{F0: -2147483648}
                                                _ = v35
                                            }
                                        case -9223372036854775808:
                                            v36 := GenS5{F0: 42, F1: v28, F2: -128}
                                            _ = v36
                                        }
                                        var v37 int = 4294967295
                                        _ = v37
                                    } else {
                                        v38 := GenS3{F0: 1, F1: ""}
                                        fmt.Println(v21, v23, v20)
                                        v22 -= v22
                                        _ = v38
                                    }
                                    switch (v27 + v19) {
                                    case "a":
                                        var v39 int16 = -928
                                        _ = v39
                                    }
                                } else {
                                    var v40 int16 = 0
                                    _ = v40
                                }
                                _ = v27
                                _ = v28
                            }
                        }
                    }
                case 2147483647:
                    v23 = v23
                default:
                    var v41 string = v19
                    _ = v41
                }
            case "\n\t":
                var v42 int = -201
                _ = v42
            }
        } else {
            v20 += v20
            _ = len(v2)
        }
        for fi43 := 0; fi43 < 13; fi43++ {
            switch (fi43 >> ((643) & 63)) {
            case -9223372036854775808:
                fmt.Println(v22, v22)
            case -202:
                delete(v2, v19)
            default:
                var v44 float32 = v1
                _ = v44
            }
            var v45 int8 = v21
            _ = fi43
            _ = v45
        }
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    v46 := GenS6{F0: 1, F1: 2}
    if v23 {
        v47 := GenS2{F0: 838, F1: -1}
        if v23 {
            v23 = v23
            for fi48 := 0; fi48 < 17; fi48++ {
                if false {
                    v19 = (((v19 + v19) + (("hello" + v19) + v19)) + "a")
                }
                _ = fi48
            }
        }
        if (-2147483648 == 32767) {
            switch 128 {
            case -2147483648:
                if true {
                    _ = len(v2)
                    var v49 int8 = -128
                    if v23 {
                        if ((!v23) || false) {
                            v23 = (2 >= 128)
                            var v50 int16 = -32768
                            var v51 int = 1
                            _ = v50
                            _ = v51
                        } else {
                            switch v19 {
                            case "\n\t":
                                var v52 uint32 = 2147483648
                                _ = v52
                            case "hello":
                                if v23 {
                                    fmt.Println(v1)
                                    v53 := v2[v19]
                                    _ = v53
                                }
                            case "a":
                                fmt.Println(v23)
                            default:
                                if v23 {
                                    if (float32(((0.0 + (3.14159 / -2.25)) / (1e10 - ((3.14159 - 3.14159) - 1.5)))) >= v1) {
                                        for fi54 := 0; fi54 < 1; fi54++ {
                                            if (2 > 86) {
                                                if true {
                                                    v1 = v20
                                                }
                                                v19 = v19
                                            } else {
                                                v21 -= (v21 / ((-118) | 1))
                                            }
                                            v55 := GenS7{F0: 1000}
                                            _ = fi54
                                            _ = v55
                                        }
                                        v56 := GenS4{F0: -1}
                                        _ = v56
                                    } else {
                                        if ((float32(((float64(v20) / 1e10) - 0.0)) - ((v20 + -2.25) + ((v20 + v20) - (v20 - v1)))) != -2.25) {
                                            var v57 float32 = 1e10
                                            var v58 float64 = 0.1
                                            _ = v57
                                            _ = v58
                                        }
                                    }
                                    var v59 uint8 = 78
                                    v23 = (v23 || (-672 <= 0))
                                    _ = v59
                                } else {
                                    fmt.Println(v49)
                                    if false {
                                        v1 = v1
                                        for fi60 := 0; fi60 < 7; fi60++ {
                                            var v61 uint16 = 2
                                            _ = fi60
                                            _ = v61
                                        }
                                    } else {
                                        v49 += v21
                                        var v62 int64 = (v22 + -602)
                                        _ = v62
                                    }
                                    if (!v23) {
                                        var v63 bool = true
                                        _ = v63
                                    } else {
                                        v64 := GenS2{F0: 861, F1: -982}
                                        switch "a" {
                                        case "foo bar":
                                            _ = len(v2)
                                        default:
                                            if ((v22 - (v22 % (((v22 + v22)) | 1))) == (v22 | (v22 / (((v22 & 1)) | 1)))) {
                                                for fi65 := 0; fi65 < 3; fi65++ {
                                                    var v66 int64 = 477
                                                    _ = fi65
                                                    _ = v66
                                                }
                                                var v67 int = 32767
                                                _ = v67
                                            } else {
                                                v1 += 0.0
                                                v19 = "\n\t"
                                                v3 = append(v3, v19)
                                            }
                                        }
                                        _ = v64
                                    }
                                }
                            }
                            v20 = float32(1.5)
                        }
                        if false {
                            var v68 uint32 = 78
                            _ = v68
                        } else {
                            var v69 uint32 = 2
                            _ = v69
                        }
                    }
                    _ = v49
                } else {
                    for fi70 := 0; fi70 < 8; fi70++ {
                        var v71 float64 = (float64(v20) + 3.14159)
                        _ = fi70
                        _ = v71
                    }
                    var v72 int16 = 717
                    _ = v72
                }
            case -846:
                var v73 uint64 = 42
                _ = v73
            case 32767:
                v1 -= (((v20 * v20) / v20) - v20)
            }
            v4 = append(v4, v19)
            var v74 uint32 = 103
            _ = v74
        }
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v46
}
