package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int; F1 uint64; F2 string }
type GenS7 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 string; F2 int64 }

func main() {
    v1 := []string{"foo bar", ("hello" + "a")}
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    var v5 uint8 = 1
    var v6 int32 = 1
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v7 uint32 = 65535
        _ = v7
    }
    v5 = (v5 | v5)
    switch ("foo bar" + "a") {
    case "foo bar":
        v8 := GenS1{F0: "foo bar", F1: v5}
        _ = v8
    case "":
        var v9 bool = true
        _ = v9
    case "hello":
        v6 += (^(v6))
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if ((^(fi10)) == fi10) {
            for fi11 := 0; fi11 < 17; fi11++ {
                v6 = v6
                fmt.Println(fi10, fi11)
                switch fi11 {
                case -9223372036854775808:
                    if false {
                        if (42 < (fi10 - fi10)) {
                            v6 = v6
                            if (("a" + "") == ("hello" + "")) {
                                if true {
                                    v12 := v4[("a" + "foo bar")]
                                    v13 := GenS1{F0: "foo bar", F1: (^(v5))}
                                    v12 = v12
                                    var v14 float32 = 1.5
                                    _ = v12
                                    _ = v13
                                    _ = v14
                                } else {
                                    v5 = v5
                                    _ = len(v2)
                                    var v15 uint32 = 65535
                                    v16 := GenS1{F0: "hello", F1: v5}
                                    _ = v15
                                    _ = v16
                                }
                                switch (fi11 & fi10) {
                                case -2147483648:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case 197:
                                    if false {
                                        v17 := GenS2{F0: 65535}
                                        _ = v17
                                    } else {
                                        switch ("a" + ((("a" + "") + "hello") + "\n\t")) {
                                        case "\n\t":
                                            if (false && false) {
                                                var v18 uint64 = 42
                                                _ = v18
                                            }
                                        case "":
                                            delete(v2, 1)
                                        case "foo bar":
                                            v6 = (v6 >> ((2) & 63))
                                        default:
                                            switch fi11 {
                                            case 2147483647:
                                                var v19 int64 = 0
                                                _ = v19
                                            case 341:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                        v20 := GenS2{F0: 698}
                                        var v21 int32 = 2147483647
                                        _ = v20
                                        _ = v21
                                    }
                                case 0:
                                    if (false && false) {
                                        v2[636] = (uint(v5) >= 102)
                                    } else {
                                        fmt.Println(fi11)
                                        v5 = v5
                                    }
                                }
                                var v22 uint64 = 729
                                _ = v22
                            } else {
                                switch (-(fi11)) {
                                case 32767:
                                    var v23 int8 = 127
                                    _ = v23
                                default:
                                    v6 = (v6 + 1)
                                }
                                v4["hello"] = "a"
                                v5 -= (v5 << ((42) & 63))
                            }
                            var v24 int8 = int8(v5)
                            _ = v24
                        } else {
                            if false {
                                var v25 float32 = ((float32(-2.25) + 3.14159) * 1.5)
                                v26 := GenS1{F0: "\n\t", F1: v5}
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS2{F0: 770}
                                switch ("hello" + "\n\t") {
                                case "hello":
                                    v6 += v6
                                    if (!false) {
                                        v6 = (v6 % ((v6) | 1))
                                        var v28 int16 = -238
                                        _ = v28
                                    } else {
                                        if ((!true) && true) {
                                            v5 += v5
                                            if (598 <= 71) {
                                                v29 := GenS1{F0: (("" + "foo bar") + "\n\t"), F1: (v5 >> ((110) & 63))}
                                                _ = v29
                                            } else {
                                                v6 = v6
                                                delete(v2, 65535)
                                                if (!true) {
                                                    var v30 uint32 = 846
                                                    v31 := GenS3{F0: 1, F1: 322, F2: 557}
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    v6 = v6
                                                    v1 = append(v1, (("" + "hello") + "hello"))
                                                }
                                            }
                                        } else {
                                            v1 = append(v1, "\n\t")
                                            v32 := GenS3{F0: 2, F1: 865, F2: 4294967295}
                                            _ = v32
                                        }
                                        if false {
                                            var v33 int64 = -936
                                            v1 = append(v1, "foo bar")
                                            _ = v33
                                        }
                                    }
                                    delete(v2, 65535)
                                }
                                v5 += 128
                                _ = v27
                            }
                            switch fi10 {
                            case 127:
                                var v34 uint = 743
                                _ = v34
                            default:
                                switch "hello" {
                                case "\n\t":
                                    if true {
                                        v35 := GenS4{F0: 0, F1: -1}
                                        if false {
                                            v4["foo bar"] = "foo bar"
                                        } else {
                                            v36 := GenS2{F0: 1}
                                            _ = v36
                                        }
                                        _ = len(v2)
                                        _ = v35
                                    } else {
                                        v37 := GenS5{F0: -1, F1: (("hello" + ("hello" + "hello")) + ("" + ("a" + (("a" + "a") + ""))))}
                                        switch (("hello" + "foo bar") + "a") {
                                        case "hello":
                                            v38 := v3[27]
                                            _ = v38
                                        case "\n\t":
                                            var v39 uint = 2
                                            _ = v39
                                        default:
                                            _ = len(v1)
                                        }
                                        _ = v37
                                    }
                                    var v40 int64 = 424
                                    var v41 uint16 = 32767
                                    if (!(float64((3.14159 + -2.25)) >= 1e10)) {
                                        var v42 int16 = 1
                                        _ = v42
                                    } else {
                                        if (!true) {
                                            v40 = -1
                                            fmt.Println(fi10, v5)
                                            switch 4294967295 {
                                            case 42:
                                                if false {
                                                    switch ("a" + "a") {
                                                    case "hello":
                                                        if (0 >= -846) {
                                                            v41 += 2
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(v6, v40, fi10)
                                                    case "":
                                                        var v43 uint32 = 1
                                                        _ = v43
                                                    default:
                                                        v44 := GenS4{F0: -85, F1: 0}
                                                        _ = v44
                                                    }
                                                    if ((v5 ^ v5) > v5) {
                                                        var v45 string = "a"
                                                        _ = v45
                                                    } else {
                                                        v5 = v5
                                                    }
                                                    var v46 bool = true
                                                    _ = v46
                                                }
                                            case 4294967295:
                                                var v47 string = "\n\t"
                                                _ = v47
                                            case -9223372036854775808:
                                                var v48 string = (("" + "a") + "hello")
                                                _ = v48
                                            }
                                        } else {
                                            fmt.Println(fi11, v6)
                                            v5 -= v5
                                            delete(v3, 65535)
                                        }
                                    }
                                    _ = v40
                                    _ = v41
                                default:
                                    switch "foo bar" {
                                    case "":
                                        var v49 float64 = ((1e10 * (3.14159 + (-2.25 * 1e10))) - -2.25)
                                        _ = v49
                                    default:
                                        var v50 int8 = -1
                                        _ = v50
                                    }
                                    if (467 > 2) {
                                        var v51 float64 = 0.1
                                        var v52 uint64 = 1
                                        _ = len(v4)
                                        var v53 int = fi10
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    } else {
                                        v5 = v5
                                        v54 := GenS4{F0: 1, F1: -9223372036854775808}
                                        v55 := GenS6{F0: (fi10 - (fi10 - fi10)), F1: 0, F2: ("hello" + "a")}
                                        _ = v54
                                        _ = v55
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                }
                            }
                            v56 := GenS4{F0: 28, F1: 192}
                            _ = v56
                        }
                    } else {
                        v6 += -392
                    }
                case 128:
                    v6 -= v6
                }
                _ = fi11
            }
        } else {
            if (v6 > 42) {
                v5 = 1
            } else {
                switch fi10 {
                case 127:
                    for fi57 := 0; fi57 < 16; fi57++ {
                        var v58 uint32 = 4294967295
                        if (!true) {
                            v59 := GenS5{F0: 0, F1: "hello"}
                            _ = v59
                        } else {
                            var v60 uint16 = 255
                            v61 := GenS3{F0: 1, F1: -1, F2: 42}
                            var v62 uint8 = v5
                            switch "foo bar" {
                            case "\n\t":
                                _ = len(v1)
                            case "a":
                                if false {
                                    if true {
                                        v6 = v6
                                    } else {
                                        v62 = v5
                                    }
                                    fmt.Println(v62, fi57, v6)
                                } else {
                                    v6 = v6
                                    if false {
                                        var v63 uint16 = (v60 + 318)
                                        v58 = (^(v58))
                                        if false {
                                            var v64 float32 = 0.1
                                            v64 = (0.1 - (0.1 + v64))
                                            var v65 uint8 = 218
                                            _ = v64
                                            _ = v65
                                        }
                                        _ = v63
                                    } else {
                                        v6 = (-(v6))
                                        var v66 int16 = int16(v58)
                                        if ((v6 << ((318) & 63)) < 1) {
                                            v62 -= v5
                                            v60 = 65535
                                            var v67 uint = 786
                                            _ = v67
                                        } else {
                                            v68 := GenS7{F0: 0, F1: true, F2: 197}
                                            _ = v68
                                        }
                                        _ = v66
                                    }
                                }
                            case "foo bar":
                                if true {
                                    fmt.Println(v60, v6)
                                }
                            }
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                        _ = fi57
                        _ = v58
                    }
                default:
                    fmt.Println(v5, v5)
                }
                v69 := GenS7{F0: 0, F1: false, F2: 1}
                var v70 uint16 = 32767
                _ = v69
                _ = v70
            }
            v71 := GenS8{F0: (v5 / ((v5) | 1))}
            _ = v71
        }
        v6 -= (v6 ^ -1)
        _ = fi10
    }
    delete(v2, 2)
    var v72 uint64 = 2
    v73 := GenS9{F0: 9223372036854775807, F1: "\n\t", F2: 9223372036854775807}
    v72 = (v72 * v72)
    v72 = (v72 + 688)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v73
}
