package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 string; F2 int64 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 int64; F1 int32; F2 uint16 }
type GenS12 struct { F0 uint8 }
func genFn0() uint32 {
    switch "a" {
    case "\n\t":
        var v4 int64 = 9223372036854775807
        _ = v4
    case "foo bar":
        var v5 float32 = -2.25
        _ = v5
    default:
        v6 := []int{-118, 1, 32767}
        _ = v6
    }
    v7 := GenS1{F0: 18446744073709551615, F1: -1, F2: 42}
    _ = v7
    return 65535
}
func genFn1(p0 int16) string {
    v8 := GenS1{F0: 65535, F1: 4294967295, F2: (^(p0))}
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := []int{(^(fi9))}
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 2147483647
    v12 := GenS1{F0: 2, F1: 127, F2: p0}
    v13 := GenS2{F0: "foo bar"}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    _ = v8
    return "foo bar"
}

func main() {
    var v1 string = "hello"
    v2 := []uint32{2147483648, 65535, 42}
    var v3 uint64 = 42
    v14 := GenS3{F0: (v3 % (((v3 ^ v3)) | 1))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v3)
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint = 369
        _ = fi15
        _ = v16
    }
    v17 := GenS4{F0: -101, F1: -2147483648}
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 bool = (0.1 != 1e10)
        switch "hello" {
        case "a":
            if true {
                if v19 {
                    for fi20 := 0; fi20 < 11; fi20++ {
                        var v21 int32 = 0
                        _ = fi20
                        _ = v21
                    }
                    if true {
                        if false {
                            if (1e10 <= 1.5) {
                                switch (fi18 | 1) {
                                case 32768:
                                    v3 += 1
                                case 1:
                                    fmt.Println(v1, fi18, v3)
                                case 127:
                                    var v22 uint32 = 2147483648
                                    _ = v22
                                }
                                switch v1 {
                                case "a":
                                    v19 = (v19 && true)
                                }
                            }
                            for fi23 := 0; fi23 < 3; fi23++ {
                                if (true && v19) {
                                    var v24 uint32 = 105
                                    if (v24 > (^(v24))) {
                                        if true {
                                            if v19 {
                                                if (0 != 2) {
                                                    var v25 uint = 42
                                                    _ = v25
                                                } else {
                                                    fmt.Println(fi18, v24, v19)
                                                }
                                            }
                                            fmt.Println(v3)
                                            fmt.Println(v1, fi23, v1)
                                        } else {
                                            v24 += (^(v24))
                                            switch 32767 {
                                            case -1:
                                                switch v1 {
                                                case "":
                                                    if v19 {
                                                        if (80 == 0) {
                                                            if (!(!v19)) {
                                                                switch 2147483647 {
                                                                case 0:
                                                                    v24 += genFn0()
                                                                    v19 = (!true)
                                                                case 4294967295:
                                                                    v26 := GenS5{F0: -1}
                                                                    _ = v26
                                                                case 768:
                                                                    v19 = false
                                                                default:
                                                                    var v27 uint64 = (v3 & 0)
                                                                    _ = len(v2)
                                                                    _ = v27
                                                                }
                                                                var v28 int16 = 1
                                                                if ("a" >= "hello") {
                                                                    v3 = (^(v3))
                                                                    v1 = v1
                                                                } else {
                                                                    fmt.Println(fi23, fi23, v3)
                                                                    var v29 uint8 = 1
                                                                    fmt.Println(v24, v28, fi23)
                                                                    _ = v29
                                                                }
                                                                _ = v28
                                                            } else {
                                                                if v19 {
                                                                    fmt.Println(fi18, fi18, v19)
                                                                    var v30 uint = 4294967295
                                                                    _ = v30
                                                                } else {
                                                                    fmt.Println(v3, v3, fi23)
                                                                }
                                                                var v31 int8 = -128
                                                                _ = v31
                                                            }
                                                            v19 = (!v19)
                                                        } else {
                                                            fmt.Println(v1, v24, v1)
                                                        }
                                                    } else {
                                                        v32 := GenS6{F0: 42, F1: v1, F2: 1}
                                                        if ((v3 / (((v3 % (((^(v3))) | 1))) | 1)) > v3) {
                                                            v33 := GenS3{F0: (v3 & 245)}
                                                            _ = v33
                                                        }
                                                        _ = v32
                                                    }
                                                default:
                                                    fmt.Println(fi18, v19, v1)
                                                }
                                            default:
                                                fmt.Println(v24, v1)
                                            }
                                        }
                                        v19 = (!false)
                                    } else {
                                        var v34 uint64 = v3
                                        _ = v34
                                    }
                                    if (!((v3 + (v3 >> ((76) & 63))) < (v3 << ((4294967295) & 63)))) {
                                        if v19 {
                                            switch v1 {
                                            case "":
                                                if ((-2.25 - (-2.25 + -2.25)) != 1.5) {
                                                    _ = len(v2)
                                                    v24 = (v24 ^ (v24 & (v24 % (((v24 % ((genFn0()) | 1))) | 1))))
                                                    v3 = (v3 | 725)
                                                } else {
                                                    v3 = v3
                                                    var v35 int64 = 265
                                                    _ = v35
                                                }
                                            case "hello":
                                                _ = len(v2)
                                            case "a":
                                                switch fi23 {
                                                case 32768:
                                                    if (!v19) {
                                                        v3 = (v3 & (^(v3)))
                                                        fmt.Println(fi23, fi18, v3)
                                                        v24 = (v24 * 2147483647)
                                                        if len(v2) > 0 {
                                                            v2[0] = v24
                                                        }
                                                    } else {
                                                        if false {
                                                            v3 += (v3 & v3)
                                                            fmt.Println(v3, v24, v24)
                                                            var v36 uint32 = v24
                                                            var v37 int8 = -1
                                                            v38 := GenS7{F0: 885, F1: 0}
                                                            if false {
                                                                v36 += (v24 ^ v24)
                                                            } else {
                                                                var v39 uint8 = 128
                                                                _ = v39
                                                            }
                                                            _ = v36
                                                            _ = v37
                                                            _ = v38
                                                        } else {
                                                            v40 := GenS2{F0: "foo bar"}
                                                            _ = v40
                                                        }
                                                    }
                                                case 128:
                                                    v41 := GenS7{F0: v3, F1: 32767}
                                                    _ = v41
                                                }
                                            }
                                        }
                                    }
                                    v42 := GenS8{F0: 2, F1: 32767, F2: uint16(v24)}
                                    v43 := GenS6{F0: 9223372036854775807, F1: ((("foo bar" + ((v1 + "a") + ("" + "\n\t"))) + (v1 + "\n\t")) + v1), F2: 9223372036854775807}
                                    _ = v24
                                    _ = v42
                                    _ = v43
                                }
                                _ = fi23
                            }
                        } else {
                            v44 := GenS9{F0: v19}
                            for fi45 := 0; fi45 < 19; fi45++ {
                                v46 := GenS2{F0: "hello"}
                                v19 = true
                                v47 := GenS5{F0: 413}
                                if (v19 || ((245 <= -682) && v19)) {
                                    if v19 {
                                        var v48 int8 = 38
                                        switch (("a" + v1) + (v1 + v1)) {
                                        case "":
                                            if (!true) {
                                                v19 = false
                                                if len(v2) > 0 {
                                                    v2[0] = 0
                                                }
                                                var v49 int64 = int64(v3)
                                                _ = v49
                                            } else {
                                                var v50 uint32 = genFn0()
                                                v51 := GenS3{F0: (^(v3))}
                                                v52 := GenS5{F0: -1}
                                                var v53 uint16 = 902
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                                _ = v53
                                            }
                                        case "foo bar":
                                            v54 := GenS5{F0: 42}
                                            _ = v54
                                        }
                                        v3 = 125
                                        _ = v48
                                    } else {
                                        v55 := GenS2{F0: v1}
                                        var v56 uint32 = 2
                                        _ = v55
                                        _ = v56
                                    }
                                } else {
                                    v19 = (!v19)
                                    switch (("foo bar" + v1) + "\n\t") {
                                    case "":
                                        v1 = "hello"
                                    case "\n\t":
                                        v3 = 42
                                    case "foo bar":
                                        var v57 int8 = 127
                                        _ = v57
                                    default:
                                        _ = len(v2)
                                        var v58 uint8 = 127
                                        var v59 string = "foo bar"
                                        var v60 float32 = 1e10
                                        var v61 int32 = 922
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    }
                                }
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                            _ = v44
                        }
                    } else {
                        _ = len(v2)
                        var v62 float64 = 1.5
                        _ = v62
                    }
                } else {
                    if v19 {
                        fmt.Println(fi18, v1)
                        var v63 int = (fi18 >> ((4294967295) & 63))
                        _ = v63
                    } else {
                        var v64 uint64 = 4294967295
                        v65 := GenS10{F0: 2, F1: 18446744073709551615}
                        _ = v64
                        _ = v65
                    }
                    var v66 int8 = int8(v3)
                    _ = v66
                }
            } else {
                v1 = ("hello" + "\n\t")
                var v67 uint8 = 2
                v2 = append(v2, uint32(v3))
                _ = v67
            }
        default:
            v68 := GenS11{F0: -1, F1: 0, F2: 2}
            _ = v68
        }
        _ = fi18
        _ = v19
    }
    v1 = v1
    v69 := GenS2{F0: "\n\t"}
    v70 := GenS12{F0: 255}
    var v71 int16 = 32767
    v3 -= 1
    v72 := GenS9{F0: false}
    v3 += v3
    v1 = (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
