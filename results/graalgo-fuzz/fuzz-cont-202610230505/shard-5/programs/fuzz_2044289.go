package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int8, p1 uint16) uint {
    p0 = p0
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            if false {
                fmt.Println(p1, fi7)
                var v9 int32 = 42
                var v10 uint16 = (p1 >> ((387) & 63))
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: 2147483648, F1: 1, F2: 65535}
            switch (^(fi7)) {
            case 0:
                p1 += p1
            default:
                if true {
                    p0 -= (p0 ^ p0)
                } else {
                    v12 := GenS2{F0: 88, F1: false, F2: 2147483647}
                    var v13 uint = 255
                    _ = v12
                    _ = v13
                }
            }
            if false {
                v14 := GenS3{F0: "foo bar"}
                _ = v14
            } else {
                p0 += p0
                v15 := GenS1{F0: 42, F1: 1, F2: 2147483648}
                if true {
                    var v16 uint64 = 915
                    _ = v16
                } else {
                    v17 := []int{(fi7 >> ((327) & 63)), fi8, fi7, fi7}
                    p0 = p0
                    if (true || true) {
                        var v18 float32 = ((-2.25 * -2.25) * float32(1e10))
                        _ = v18
                    } else {
                        if true {
                            p1 += (^(p1))
                            switch (fi8 * (fi7 ^ fi8)) {
                            case 146:
                                v19 := GenS1{F0: 2, F1: 2, F2: 2147483647}
                                _ = v19
                            default:
                                p0 -= p0
                            }
                        } else {
                            if true {
                                var v20 uint = 2
                                var v21 int8 = (p0 | (p0 << ((105) & 63)))
                                _ = v20
                                _ = v21
                            }
                        }
                    }
                    _ = v17
                }
                _ = v15
            }
            var v22 uint = 359
            _ = fi8
            _ = v11
            _ = v22
        }
        var v23 uint = 0
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 uint16 = p1
            var v26 int8 = p0
            switch ("hello" + "\n\t") {
            case "\n\t":
                var v27 int16 = 32767
                _ = v27
            case "":
                var v28 uint8 = 2
                _ = v28
            case "hello":
                v29 := GenS3{F0: ("\n\t" + "hello")}
                _ = v29
            default:
                switch "foo bar" {
                case "":
                    var v30 float32 = 1e10
                    _ = v30
                case "foo bar":
                    v23 = v23
                    v31 := []int{fi24, fi24}
                    _ = v31
                case "hello":
                    v32 := GenS3{F0: (("foo bar" + ("a" + "foo bar")) + "hello")}
                    _ = v32
                default:
                    var v33 bool = false
                    _ = v33
                }
            }
            v23 = v23
            var v34 float32 = -2.25
            var v35 uint32 = uint32(v23)
            _ = fi24
            _ = v25
            _ = v26
            _ = v34
            _ = v35
        }
        _ = fi7
        _ = v23
    }
    fmt.Println(p1, p0)
    fmt.Println(p1, p0)
    p0 += (p0 >> ((uint(p1)) & 63))
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1() int {
    var v36 bool = (-1 >= 1)
    if (!(!v36)) {
        if (1 >= 1) {
            if (!v36) {
                fmt.Println(v36)
            }
        } else {
            v37 := GenS4{F0: "\n\t"}
            _ = v37
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            v36 = false
            _ = fi38
        }
    }
    v36 = ((127 < 42) || (42 >= 255))
    var v39 uint8 = 64
    v40 := GenS2{F0: -1, F1: (692 == 993), F2: 2147483648}
    _ = v36
    _ = v39
    _ = v40
    return -284
}

func main() {
    var v1 float32 = float32(1e10)
    var v2 int = 32768
    var v3 int64 = -877
    var v4 uint = 0
    v5 := make(map[uint32]bool)
    v6 := []int{(v2 % ((108) | 1))}
    switch v2 {
    case 9223372036854775807:
        for fi41 := 0; fi41 < 15; fi41++ {
            var v42 string = "a"
            _ = fi41
            _ = v42
        }
    default:
        fmt.Println(v3)
    }
    v2 -= v2
    v43 := v5[1]
    v44 := GenS3{F0: "a"}
    delete(v5, 42)
    if (!v43) {
        v43 = (!(-423 > 42))
        var v45 bool = v43
        _ = v45
    }
    v2 = (v2 % ((v2) | 1))
    if v43 {
        if (!v43) {
            switch (((("foo bar" + ((("a" + (("hello" + "foo bar") + ("foo bar" + "hello"))) + "hello") + "foo bar")) + "hello") + "foo bar") + "hello") {
            case "foo bar":
                switch (((("hello" + "\n\t") + "hello") + "") + "\n\t") {
                case "foo bar":
                    v46 := GenS2{F0: 2, F1: (!false), F2: 4294967295}
                    _ = v46
                case "\n\t":
                    v43 = (v43 || (!(genFn0(14, 974) < (^(v4)))))
                case "hello":
                    v3 = v3
                default:
                    for fi47 := 0; fi47 < 1; fi47++ {
                        if v43 {
                            v43 = ((v4 & (^(v4))) >= 136)
                            var v48 int16 = 394
                            _ = v48
                        }
                        switch (fi47 * fi47) {
                        case -574:
                            v49 := GenS5{F0: (v3 % ((int64(v2)) | 1))}
                            _ = v49
                        default:
                            var v50 uint = v4
                            _ = v50
                        }
                        var v51 bool = (122 <= 0)
                        _ = fi47
                        _ = v51
                    }
                }
            case "hello":
                if v43 {
                    var v52 uint16 = 2
                    var v53 int16 = -1
                    _ = v52
                    _ = v53
                }
            }
            for fi54 := 0; fi54 < 6; fi54++ {
                switch "hello" {
                case "foo bar":
                    if (-32768 == -32768) {
                        if (1000 < 42) {
                            if len(v6) > 0 {
                                v6[0] = (v2 ^ (^(fi54)))
                            }
                            switch (v2 | (v2 - (v2 >> ((v4) & 63)))) {
                            case 42:
                                v1 += v1
                            case -951:
                                if false {
                                    var v55 int = -2147483648
                                    var v56 int32 = 42
                                    _ = v55
                                    _ = v56
                                }
                            }
                        } else {
                            v43 = false
                            v3 = (-(v3))
                        }
                    } else {
                        var v57 bool = (((true && true) && v43) && (v43 || v43))
                        _ = v57
                    }
                case "":
                    fmt.Println(v43, v3, v3)
                case "a":
                    if (65535 > 65535) {
                        for fi58 := 0; fi58 < 8; fi58++ {
                            v59 := GenS3{F0: "hello"}
                            _ = len(v5)
                            if v43 {
                                v3 += (v3 * v3)
                                if len(v6) > 0 {
                                    v6[0] = v2
                                }
                            }
                            _ = fi58
                            _ = v59
                        }
                    } else {
                        switch genFn1() {
                        case -2147483648:
                            if true {
                                v60 := GenS6{F0: (v3 - (v3 & (v3 - (v3 << ((919) & 63))))), F1: 2}
                                if (42 != 289) {
                                    v2 -= v2
                                    v61 := GenS2{F0: -49, F1: ((42 <= 2) || (1.5 >= -2.25)), F2: 2147483648}
                                    var v62 uint8 = 0
                                    _ = v61
                                    _ = v62
                                } else {
                                    v43 = (410 > (v3 + (v3 - v3)))
                                    var v63 bool = (genFn0(105, 65535) == 65535)
                                    v4 = 2
                                    _ = v63
                                }
                                _ = v60
                            }
                        case -614:
                            switch fi54 {
                            case -1:
                                var v64 uint16 = 32767
                                _ = v64
                            case 296:
                                v2 += (fi54 << ((v4) & 63))
                            case 790:
                                if (v4 == v4) {
                                    if true {
                                        if ((float64(1e10) != 0.0) && v43) {
                                            if (-253 != 1) {
                                                fmt.Println(v43, v1)
                                                v65 := GenS4{F0: ("foo bar" + "hello")}
                                                _ = v65
                                            } else {
                                                for fi66 := 0; fi66 < 17; fi66++ {
                                                    _ = len(v5)
                                                    delete(v5, 65535)
                                                    var v67 int64 = 0
                                                    if v43 {
                                                        switch "\n\t" {
                                                        case "a":
                                                            v67 += (v3 & -62)
                                                            v1 = 3.14159
                                                            v1 = v1
                                                        case "\n\t":
                                                            v2 = fi54
                                                        case "hello":
                                                            var v68 int = 127
                                                            _ = v68
                                                        }
                                                    }
                                                    _ = fi66
                                                    _ = v67
                                                }
                                                for fi69 := 0; fi69 < 18; fi69++ {
                                                    var v70 uint32 = 42
                                                    if v43 {
                                                        var v71 uint32 = 339
                                                        if (v71 <= (v71 * 1)) {
                                                            v72 := GenS5{F0: (v3 / ((9223372036854775807) | 1))}
                                                            v71 -= (v70 % (((v70 ^ (v70 / (((v70 - v71)) | 1)))) | 1))
                                                            _ = v72
                                                        } else {
                                                            v73 := GenS7{F0: -99, F1: 96}
                                                            _ = v73
                                                        }
                                                        v74 := GenS8{F0: 65535, F1: uint8(v71)}
                                                        _ = v71
                                                        _ = v74
                                                    } else {
                                                        v70 += (v70 | v70)
                                                    }
                                                    v43 = true
                                                    var v75 bool = v43
                                                    _ = fi69
                                                    _ = v70
                                                    _ = v75
                                                }
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                        }
                                        v4 = (^(v4))
                                    }
                                }
                            }
                        case -9223372036854775808:
                            var v76 int64 = v3
                            _ = v76
                        default:
                            var v77 int16 = 42
                            _ = v77
                        }
                        for fi78 := 0; fi78 < 13; fi78++ {
                            v6 = append(v6, -1)
                            var v79 uint16 = 42
                            _ = fi78
                            _ = v79
                        }
                    }
                }
                _ = fi54
            }
        } else {
            v5[4294967295] = (v1 < float32(1e10))
            v4 -= v4
        }
    }
    var v80 float32 = (1e10 - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v80
}
