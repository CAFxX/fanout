package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint; F1 uint32; F2 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 string }
type GenS10 struct { F0 int32; F1 int; F2 uint16 }
type GenS11 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS12 struct { F0 int8; F1 int64 }
func genFn0(p0 uint32, p1 int16) int8 {
    if false {
        p0 = 2147483647
    } else {
        var v5 uint16 = 971
        v6 := GenS1{F0: 42}
        v7 := GenS2{F0: 155, F1: ("" + ("hello" + "a"))}
        _ = v5
        _ = v6
        _ = v7
    }
    switch "foo bar" {
    case "hello":
        v8 := []int{411, 9223372036854775807}
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 16; fi9++ {
            for fi10 := 0; fi10 < 14; fi10++ {
                v11 := GenS1{F0: p0}
                var v12 uint = 65535
                p1 += (p1 + (p1 * p1))
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    case "a":
        fmt.Println(p0, p1)
    }
    v13 := []int{32767, 42}
    var v14 uint64 = 0
    fmt.Println(p0, v14)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return -117
}
func genFn1() int8 {
    var v15 uint16 = 2
    if true {
        var v16 int16 = 42
        switch "hello" {
        case "hello":
            v17 := GenS1{F0: 2}
            _ = v17
        case "foo bar":
            for fi18 := 0; fi18 < 16; fi18++ {
                v19 := GenS3{F0: 785, F1: (v16 | 32767)}
                var v20 uint32 = 4294967295
                _ = fi18
                _ = v19
                _ = v20
            }
        case "\n\t":
            v15 -= v15
        default:
            v15 += (^(v15))
        }
        v16 += v16
        _ = v16
    } else {
        v21 := GenS4{F0: "", F1: 42, F2: -9223372036854775808}
        var v22 int32 = -97
        _ = v21
        _ = v22
    }
    var v23 string = ("\n\t" + "\n\t")
    v23 = (("a" + (v23 + v23)) + v23)
    v24 := []int{-9223372036854775808, 4294967295, 246, 0}
    _ = v15
    _ = v23
    _ = v24
    return -1
}

func main() {
    var v1 uint = 42
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    for fi25 := 0; fi25 < 6; fi25++ {
        if true {
            if (-9223372036854775808 == 42) {
                if true {
                    if (!true) {
                        if false {
                            fmt.Println(fi25)
                            v26 := GenS2{F0: 121, F1: ("a" + "a")}
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 2; fi27++ {
                                fmt.Println(fi25, fi25)
                                v28 := []int{128, (^(fi25)), fi25, fi27}
                                _ = fi27
                                _ = v28
                            }
                            v1 = 2
                        }
                        if ((1 <= -32768) && (fi25 > (^(fi25)))) {
                            v1 -= v1
                            switch (-(fi25)) {
                            case 4294967295:
                                for fi29 := 0; fi29 < 19; fi29++ {
                                    fmt.Println(v1)
                                    _ = fi29
                                }
                            default:
                                v1 -= (v1 << (((v1 >> ((4294967295) & 63))) & 63))
                            }
                        } else {
                            v1 += (^(v1))
                            v1 = 65535
                        }
                        fmt.Println(v1)
                    }
                    fmt.Println(fi25, fi25)
                } else {
                    fmt.Println(v1, v1)
                    v30 := GenS5{F0: 255, F1: 2147483648, F2: 32767}
                    if true {
                        v1 = (v1 - (v1 - v1))
                    } else {
                        v31 := GenS3{F0: 2147483647, F1: 0}
                        _ = v31
                    }
                    _ = v30
                }
                switch "a" {
                case "a":
                    if (1 == -128) {
                        v32 := GenS1{F0: 4294967295}
                        var v33 int32 = 1
                        _ = v32
                        _ = v33
                    }
                case "\n\t":
                    v34 := GenS6{F0: 42}
                    _ = v34
                }
            } else {
                v35 := GenS2{F0: -1, F1: "a"}
                _ = v35
            }
        } else {
            fmt.Println(v1, fi25)
            for fi36 := 0; fi36 < 4; fi36++ {
                v1 -= 65535
                switch (fi36 >> ((615) & 63)) {
                case 265:
                    if true {
                        var v37 int8 = -128
                        v37 = v37
                        if false {
                            switch fi36 {
                            case -1:
                                var v38 uint32 = 177
                                _ = v38
                            case 0:
                                var v39 int64 = 9223372036854775807
                                _ = v39
                            case -225:
                                fmt.Println(v1)
                            }
                            v37 += v37
                        } else {
                            if (false && true) {
                                fmt.Println(v1)
                            } else {
                                v4[2] = "foo bar"
                            }
                            fmt.Println(fi25, v1)
                            var v40 int32 = -1
                            switch "foo bar" {
                            case "":
                                v41 := GenS7{F0: (!((v40 >> ((v1) & 63)) >= v40)), F1: 2147483648, F2: 255}
                                _ = v41
                            }
                            _ = v40
                        }
                        _ = v37
                    } else {
                        if (true || true) {
                            var v42 string = ""
                            var v43 int64 = 457
                            v43 += 0
                            switch (fi25 << ((v1) & 63)) {
                            case 42:
                                v44 := GenS1{F0: 42}
                                _ = v44
                            }
                            _ = v42
                            _ = v43
                        } else {
                            v1 = v1
                        }
                        v1 = v1
                        switch (("hello" + "hello") + "\n\t") {
                        case "hello":
                            if false {
                                v45 := GenS5{F0: (^(v1)), F1: 2147483647, F2: 42}
                                fmt.Println(fi36)
                                if (191 >= 720) {
                                    v1 = v1
                                    v1 = 65535
                                    v1 = v1
                                } else {
                                    if true {
                                        v1 = 65535
                                        v1 = v1
                                        v1 = (v1 & 255)
                                    } else {
                                        var v46 int64 = -1
                                        _ = v46
                                    }
                                    var v47 float32 = 3.14159
                                    switch "hello" {
                                    case "hello":
                                        v47 += 0.1
                                    case "foo bar":
                                        v1 -= (v1 ^ v1)
                                    case "a":
                                        var v48 int = (fi36 & 2147483647)
                                        _ = v48
                                    }
                                    v49 := []int{-883, (fi25 >> ((v1) & 63))}
                                    _ = v47
                                    _ = v49
                                }
                                v50 := v4[0]
                                v51 := GenS2{F0: int64(fi25), F1: "hello"}
                                v52 := []int{fi25}
                                _ = v45
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                            v53 := GenS3{F0: 63, F1: -956}
                            _ = v53
                        default:
                            v54 := GenS3{F0: 65535, F1: 0}
                            _ = v54
                        }
                    }
                case -456:
                    v55 := []int{fi25, (^(fi36)), (-(fi36)), fi25}
                    _ = v55
                }
                v56 := GenS6{F0: -268}
                _ = fi36
                _ = v56
            }
            fmt.Println(v1, v1)
        }
        switch 127 {
        case -691:
            v1 = (v1 - (v1 - 2))
        case 32767:
            var v57 uint = v1
            _ = v57
        default:
            v58 := GenS8{F0: "foo bar"}
            _ = v58
        }
        _ = fi25
    }
    v59 := []int{32768, -2147483648, -594, 0}
    v59 = append(v59, -158)
    var v60 bool = false
    fmt.Println(v1, v1)
    var v61 uint16 = 392
    var v62 uint = v1
    for fi63 := 0; fi63 < 10; fi63++ {
        if (0 < 42) {
            v64 := GenS9{F0: 1, F1: "hello"}
            fmt.Println(v1, v62)
            switch (fi63 + fi63) {
            case 42:
                switch (fi63 + fi63) {
                case 0:
                    v60 = v60
                }
            case 4294967295:
                fmt.Println(v62, v1, v60)
            default:
                if v60 {
                    v59 = append(v59, fi63)
                }
            }
            _ = v64
        }
        v60 = ("a" != "hello")
        _ = fi63
    }
    var v65 int8 = 100
    v62 -= (v62 % ((v1) | 1))
    for fi66 := 0; fi66 < 19; fi66++ {
        v1 = 938
        for fi67 := 0; fi67 < 3; fi67++ {
            var v68 string = ("hello" + "hello")
            v65 = -1
            switch "hello" {
            case "\n\t":
                if v60 {
                    v69 := GenS9{F0: 947, F1: "\n\t"}
                    _ = v69
                } else {
                    var v70 int = fi67
                    var v71 bool = (!v60)
                    _ = len(v59)
                    if (v65 != -1) {
                        v72 := GenS1{F0: 64}
                        _ = v72
                    }
                    fmt.Println(v71)
                    _ = v70
                    _ = v71
                }
            case "a":
                switch "" {
                case "hello":
                    if v60 {
                        var v73 uint = (^(v1))
                        _ = v73
                    } else {
                        var v74 int = fi67
                        _ = v74
                    }
                case "":
                    v60 = (int(v1) <= 32768)
                default:
                    if false {
                        v75 := GenS8{F0: ("hello" + "")}
                        fmt.Println(v60)
                        _ = v75
                    }
                }
                fmt.Println(v60, v1)
                fmt.Println(v1, v61)
            }
            _ = fi67
            _ = v68
        }
        v76 := GenS10{F0: 42, F1: (^(fi66)), F2: v61}
        _ = fi66
        _ = v76
    }
    if (v60 || (411 < 0)) {
        v77 := GenS11{F0: v65, F1: v61, F2: 65535}
        if v60 {
            v78 := GenS8{F0: ""}
            v65 -= (v65 + genFn0(882, 0))
            var v79 int = -346
            _ = v78
            _ = v79
        } else {
            v80 := GenS12{F0: v65, F1: -538}
            v2[((("" + "foo bar") + ("a" + "\n\t")) + ("foo bar" + (("\n\t" + "foo bar") + "")))] = 65535
            _ = v80
        }
        v1 -= v1
        _ = v77
    }
    for fi81 := 0; fi81 < 6; fi81++ {
        for fi82 := 0; fi82 < 19; fi82++ {
            _ = len(v2)
            _ = fi82
        }
        v83 := GenS10{F0: 0, F1: fi81, F2: v61}
        _ = fi81
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v65
}
