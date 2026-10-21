package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint; F1 int; F2 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 int8; F2 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 int32; F1 string; F2 int }
type GenS10 struct { F0 uint64; F1 uint64 }
type GenS11 struct { F0 bool }
func genFn0(p0 int16) uint32 {
    v7 := []int{9223372036854775807, 32767, 398, 822}
    v8 := GenS1{F0: 127, F1: "foo bar", F2: true}
    _ = p0
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 2
    v3 := make(map[uint32]bool)
    v4 := []string{"hello", "hello"}
    v5 := []string{("\n\t" + "\n\t"), ("foo bar" + ("hello" + "hello")), "a"}
    v6 := []string{"", (("a" + "foo bar") + "hello")}
    if (!false) {
        v9 := GenS2{F0: 32767}
        v10 := GenS3{F0: true, F1: 1}
        switch 229 {
        case 9223372036854775807:
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 uint16 = (v2 << ((0) & 63))
                for fi13 := 0; fi13 < 4; fi13++ {
                    if (false || false) {
                        v1 = (^(v1))
                        v2 = 32767
                        v14 := GenS2{F0: fi11}
                        _ = v14
                    } else {
                        if false {
                            v15 := GenS4{F0: 2, F1: (-(fi11)), F2: 128}
                            if false {
                                var v16 float64 = 3.14159
                                _ = v16
                            }
                            _ = v15
                        } else {
                            if (!((fi11 & 782) < fi13)) {
                                var v17 int8 = 91
                                _ = v17
                            }
                            if false {
                                var v18 uint64 = v1
                                v12 += v2
                                _ = v18
                            } else {
                                v12 += 65535
                                if (!(0 >= -1)) {
                                    switch ("foo bar" + "foo bar") {
                                    case "hello":
                                        fmt.Println(v12)
                                    default:
                                        v12 = 65535
                                    }
                                    v19 := GenS3{F0: false, F1: -79}
                                    _ = v19
                                } else {
                                    var v20 uint = 255
                                    v5 = append(v5, ("hello" + (("\n\t" + "hello") + "hello")))
                                    _ = v20
                                }
                            }
                            v2 = v12
                        }
                        switch fi13 {
                        case 127:
                            if false {
                                v1 += (v1 - v1)
                            }
                        case -993:
                            fmt.Println(fi11)
                        case -1:
                            var v21 uint = 4294967295
                            _ = v21
                        }
                        switch ("\n\t" + "hello") {
                        case "a":
                            v2 += (^(v12))
                        case "hello":
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                        case "":
                            var v22 float32 = 0.0
                            _ = v22
                        default:
                            var v23 bool = true
                            _ = v23
                        }
                        var v24 int16 = -828
                        v12 += (v2 & (v2 * v12))
                        fmt.Println(fi13, v24, v24)
                        _ = v24
                    }
                    if true {
                        v25 := GenS5{F0: (v1 / ((v1) | 1))}
                        v26 := GenS4{F0: 4294967295, F1: 32768, F2: 1}
                        _ = v25
                        _ = v26
                    } else {
                        if ("a" > ("foo bar" + ("" + ""))) {
                            if false {
                                v27 := GenS6{F0: 1, F1: -1, F2: 42}
                                fmt.Println(v2, v2)
                                fmt.Println(v12)
                                fmt.Println(v12, v1, fi13)
                                _ = v27
                            } else {
                                var v28 float64 = ((0.0 - 0.1) - (0.1 + 1.5))
                                _ = v28
                            }
                        }
                        if (false || (0 >= 19)) {
                            v29 := GenS7{F0: (v1 >> ((4294967295) & 63))}
                            _ = v29
                        }
                        var v30 float32 = 1e10
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        switch ("" + (("hello" + "") + "a")) {
                        case "hello":
                            var v31 uint16 = v12
                            _ = v31
                        case "":
                            switch "" {
                            case "a":
                                v30 = v30
                            case "\n\t":
                                v32 := GenS1{F0: 178, F1: "\n\t", F2: true}
                                _ = v32
                            case "":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                v33 := GenS1{F0: 255, F1: "hello", F2: false}
                                _ = v33
                            }
                            v1 = (v1 - 927)
                        case "a":
                            var v34 uint8 = 162
                            _ = v34
                        default:
                            switch "a" {
                            case "\n\t":
                                if ((fi13 / (((fi13 / ((fi11) | 1))) | 1)) > 729) {
                                    if (v30 != 3.14159) {
                                        var v35 float64 = 1e10
                                        v30 = v30
                                        _ = v35
                                    } else {
                                        if true {
                                            var v36 int = fi13
                                            fmt.Println(fi11, v36)
                                            v37 := GenS8{F0: 184, F1: 32767}
                                            _ = v36
                                            _ = v37
                                        } else {
                                            _ = len(v3)
                                        }
                                    }
                                    v38 := GenS9{F0: 2147483647, F1: "foo bar", F2: fi11}
                                    var v39 uint16 = v12
                                    _ = v38
                                    _ = v39
                                } else {
                                    fmt.Println(v2, v2, v12)
                                    v4 = append(v4, ((("hello" + "a") + "") + "foo bar"))
                                    v30 = v30
                                }
                            case "hello":
                                v30 -= 1.5
                            }
                        }
                        _ = v30
                    }
                    _ = fi13
                }
                var v40 float32 = float32(0.1)
                _ = fi11
                _ = v12
                _ = v40
            }
        case 2147483647:
            if true {
                v2 = v2
            } else {
                for fi41 := 0; fi41 < 13; fi41++ {
                    v1 = (v1 >> ((845) & 63))
                    for fi42 := 0; fi42 < 1; fi42++ {
                        v2 -= v2
                        _ = fi42
                    }
                    _ = fi41
                }
            }
        case -601:
            var v43 int = -9223372036854775808
            _ = v43
        }
        _ = v10
        _ = v9
    }
    var v44 float64 = 3.14159
    _ = len(v4)
    switch "foo bar" {
    case "":
        var v45 int32 = -1
        _ = v45
    case "foo bar":
        v46 := GenS10{F0: v1, F1: 1}
        _ = v46
    case "hello":
        var v47 uint64 = v1
        _ = v47
    }
    fmt.Println(v1, v2, v44)
    for fi48 := 0; fi48 < 2; fi48++ {
        v44 += v44
        if false {
            var v49 uint32 = 38
            v1 += (v1 & 1)
            _ = v49
        } else {
            for fi50 := 0; fi50 < 17; fi50++ {
                v44 += float64(0.1)
                if false {
                    if ((783 < genFn0(-583)) || false) {
                        switch "" {
                        case "\n\t":
                            var v51 int = (-(fi50))
                            _ = v51
                        }
                    }
                    v2 = v2
                    v52 := GenS11{F0: (1 < -32768)}
                    _ = v52
                } else {
                    var v53 int32 = -963
                    delete(v3, genFn0(-699))
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v53
                }
                _ = fi50
            }
            var v54 string = ("hello" + "a")
            _ = v54
        }
        _ = fi48
    }
    var v55 int16 = 42
    for fi56 := 0; fi56 < 4; fi56++ {
        for fi57 := 0; fi57 < 1; fi57++ {
            v2 += v2
            _ = fi57
        }
        v1 -= uint64(v2)
        switch fi56 {
        case -916:
            v44 = float64(-2.25)
        case -1:
            for fi58 := 0; fi58 < 11; fi58++ {
                v1 += 4294967295
                v1 = 0
                if (!(9223372036854775807 == int64(v55))) {
                    v2 = (^(v2))
                } else {
                    v1 += v1
                    v2 -= v2
                    v55 += 126
                }
                _ = fi58
            }
        case -2147483648:
            v3[genFn0(v55)] = true
        }
        _ = fi56
    }
    delete(v3, 42)
    var v59 uint32 = 65535
    v44 += v44
    if false {
        if (996 >= 65535) {
            v1 = (v1 * v1)
            fmt.Println(v59, v55, v59)
        }
    }
    delete(v3, (v59 & 2147483647))
    fmt.Println(v59)
    var v60 int64 = 1
    if ((((("" + "\n\t") + (("foo bar" + ("foo bar" + "\n\t")) + (("\n\t" + ("\n\t" + ("foo bar" + "hello"))) + "a"))) + "\n\t") + "a") != "hello") {
        v2 -= 0
        if (128 > 127) {
            _ = len(v6)
            for fi61 := 0; fi61 < 5; fi61++ {
                if false {
                    v60 = v60
                    v62 := GenS2{F0: fi61}
                    _ = len(v3)
                    _ = v62
                }
                switch fi61 {
                case 2147483647:
                    v55 = v55
                case 42:
                    if ((v1 - v1) == v1) {
                        for fi63 := 0; fi63 < 20; fi63++ {
                            v55 = v55
                            var v64 int = 127
                            _ = fi63
                            _ = v64
                        }
                    } else {
                        var v65 int8 = -1
                        _ = v65
                    }
                default:
                    v44 = (v44 - 0.1)
                }
                var v66 uint8 = 42
                _ = fi61
                _ = v66
            }
            v2 += v2
        }
        var v67 int32 = 477
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v55
    _ = v59
    _ = v6
    _ = v60
}
