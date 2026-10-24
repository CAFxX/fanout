package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint32; F1 int64 }
func genFn0() string {
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint = 4294967295
        var v9 float64 = 1e10
        var v10 uint32 = 417
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS2{F0: 2147483647, F1: 128}
    switch 32767 {
    case 4294967295:
        var v12 int32 = 42
        _ = v12
    case 482:
        var v13 uint64 = 42
        _ = v13
    default:
        v14 := []int{-716, 32767, 127}
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        if true {
            v16 := GenS2{F0: -422, F1: 42}
            _ = v16
        }
        _ = fi15
    }
    _ = v11
    _ = v6
    return "a"
}
func genFn1(p0 int64, p1 uint, p2 int32) uint8 {
    v17 := GenS1{F0: 213}
    p1 = p1
    var v18 bool = (true || false)
    p0 -= 9223372036854775807
    p0 = (p0 % ((0) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 2
}

func main() {
    var v1 int32 = -569
    var v2 uint = 65535
    v3 := make(map[string]string)
    var v4 float32 = 3.14159
    v5 := []string{"\n\t", ("\n\t" + ("" + "a")), "hello", "\n\t"}
    v19 := GenS3{F0: 2}
    v20 := GenS1{F0: 255}
    var v21 int64 = 42
    v4 = v4
    switch 933 {
    case -1:
        v22 := GenS4{F0: "a", F1: 996}
        _ = v22
    case -543:
        v2 += (v2 % (((v2 << (((v2 << (((v2 * 65535)) & 63))) & 63))) | 1))
    case -67:
        v21 = v21
    }
    var v23 int64 = (v21 >> ((v2) & 63))
    var v24 int32 = v1
    fmt.Println(v24)
    fmt.Println(v1, v2, v21)
    var v25 bool = false
    for fi26 := 0; fi26 < 19; fi26++ {
        if (!v25) {
            switch "a" {
            case "hello":
                v27 := GenS1{F0: 127}
                _ = v27
            case "\n\t":
                v28 := GenS5{F0: 42}
                _ = v28
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                switch 9223372036854775807 {
                case 128:
                    var v30 int32 = (v1 * (v24 * v24))
                    _ = v30
                case -727:
                    fmt.Println(fi26, v2, v2)
                }
                _ = fi29
            }
            var v31 uint16 = 255
            _ = v31
        } else {
            if true {
                _ = len(v3)
                if false {
                    switch "" {
                    case "":
                        if (!(-223 <= 42)) {
                            v21 += (v23 << ((v2) & 63))
                            v25 = (v25 || true)
                            var v32 int32 = 0
                            _ = v32
                        } else {
                            v3[""] = "a"
                        }
                    }
                } else {
                    if (377 >= 42) {
                        switch ("a" + "hello") {
                        case "":
                            v33 := GenS2{F0: v24, F1: (fi26 - (^(fi26)))}
                            _ = v33
                        case "foo bar":
                            v5 = append(v5, ("hello" + (("a" + ("\n\t" + ("\n\t" + ("hello" + "")))) + ("hello" + (("hello" + "") + "")))))
                        }
                        var v34 int64 = (v21 % ((v21) | 1))
                        _ = v34
                    } else {
                        if v25 {
                            var v35 string = "foo bar"
                            for fi36 := 0; fi36 < 19; fi36++ {
                                var v37 float64 = float64((v4 + ((v4 / 1.0) / ((float32(0.1) * 0.0) - v4))))
                                if v25 {
                                    var v38 uint64 = 42
                                    v37 = ((float64(3.14159) * v37) + (3.14159 * (1.5 / -2.25)))
                                    v25 = false
                                    _ = v38
                                }
                                _ = fi36
                                _ = v37
                            }
                            _ = v35
                        } else {
                            v39 := GenS1{F0: 2}
                            switch fi26 {
                            case 9223372036854775807:
                                var v40 int8 = -1
                                _ = v40
                            default:
                                var v41 int64 = (v23 | v23)
                                _ = v41
                            }
                            var v42 uint32 = 65535
                            _ = v39
                            _ = v42
                        }
                    }
                    if false {
                        var v43 uint16 = 1
                        _ = v43
                    } else {
                        switch (fi26 | (fi26 ^ (^(fi26)))) {
                        case 2147483647:
                            v21 -= v23
                        case 32767:
                            for fi44 := 0; fi44 < 7; fi44++ {
                                var v45 int64 = (v21 / ((42) | 1))
                                fmt.Println(v25)
                                var v46 string = (("hello" + "hello") + "a")
                                _ = fi44
                                _ = v45
                                _ = v46
                            }
                        case -470:
                            v4 += v4
                        default:
                            delete(v3, "\n\t")
                        }
                        var v47 bool = false
                        _ = v47
                    }
                    if ((v23 * v21) > v23) {
                        for fi48 := 0; fi48 < 6; fi48++ {
                            var v49 int32 = (-(v1))
                            switch (fi26 % ((fi26) | 1)) {
                            case -9223372036854775808:
                                v23 = -9223372036854775808
                            case -731:
                                if v25 {
                                    v24 = 688
                                    if false {
                                        delete(v3, ("" + "foo bar"))
                                    }
                                    delete(v3, "")
                                    var v50 uint = v2
                                    switch ("hello" + "a") {
                                    case "\n\t":
                                        var v51 uint16 = 0
                                        v52 := GenS5{F0: 0}
                                        if (v1 <= (v49 & (v24 % (((v1 - v24)) | 1)))) {
                                            var v53 uint64 = 1
                                            if ((0 <= 230) && v25) {
                                                if true {
                                                    var v54 uint32 = 435
                                                    _ = len(v5)
                                                    if (uint64(v54) < v53) {
                                                        var v55 uint8 = 128
                                                        _ = v55
                                                    } else {
                                                        v56 := v3[("foo bar" + "hello")]
                                                        fmt.Println(v23, fi26, v25)
                                                        _ = v56
                                                    }
                                                    v57 := GenS3{F0: (v51 / ((v51) | 1))}
                                                    _ = v54
                                                    _ = v57
                                                } else {
                                                    var v58 int16 = 1000
                                                    _ = v58
                                                }
                                                var v59 uint32 = 411
                                                _ = v59
                                            } else {
                                                fmt.Println(v1, fi26)
                                                v60 := GenS4{F0: ((((("a" + "hello") + (("a" + ("hello" + "a")) + "\n\t")) + "\n\t") + "") + (("" + "\n\t") + ("" + "foo bar"))), F1: 2147483648}
                                                var v61 uint32 = 1
                                                _ = v60
                                                _ = v61
                                            }
                                            var v62 string = "\n\t"
                                            _ = v53
                                            _ = v62
                                        }
                                        _ = v51
                                        _ = v52
                                    case "a":
                                        v23 -= 0
                                    case "hello":
                                        v2 -= v50
                                    default:
                                        if true {
                                            fmt.Println(v50, v4)
                                            v63 := GenS6{F0: fi48, F1: 255}
                                            switch "a" {
                                            case "":
                                                var v64 float64 = (1.5 / -2.25)
                                                _ = v64
                                            case "foo bar":
                                                var v65 uint8 = genFn1(9223372036854775807, (v2 >> ((v2) & 63)), (v49 << ((1) & 63)))
                                                _ = v65
                                            case "a":
                                                v66 := GenS3{F0: 890}
                                                _ = v66
                                            default:
                                                fmt.Println(v50)
                                            }
                                            _ = v63
                                        }
                                    }
                                    _ = v50
                                } else {
                                    switch fi48 {
                                    case 32767:
                                        var v67 int64 = int64(fi48)
                                        var v68 uint16 = 2
                                        var v69 bool = (0.1 < -2.25)
                                        v5 = append(v5, ("\n\t" + "a"))
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 uint32 = 639
                                    switch "" {
                                    case "\n\t":
                                        v49 -= (v1 * -355)
                                    case "hello":
                                        _ = len(v3)
                                    }
                                    _ = v70
                                }
                            default:
                                if (!(!true)) {
                                    v71 := GenS7{F0: (v23 & (v23 % (((v23 - (v21 % ((835) | 1)))) | 1))), F1: int8(v23)}
                                    v1 += (-(v49))
                                    _ = v71
                                } else {
                                    var v72 string = "\n\t"
                                    _ = v72
                                }
                            }
                            var v73 uint64 = 2
                            _ = fi48
                            _ = v49
                            _ = v73
                        }
                        fmt.Println(v4, v23)
                    } else {
                        var v74 int16 = -669
                        var v75 string = (("foo bar" + "\n\t") + "foo bar")
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v74
                        _ = v75
                    }
                }
            } else {
                v76 := GenS5{F0: 159}
                _ = v76
            }
        }
        _ = fi26
    }
    v77 := GenS6{F0: 128, F1: 42}
    if v25 {
        v2 = 1
        var v78 int64 = (v23 + v23)
        _ = v78
    }
    delete(v3, ("hello" + ("foo bar" + ("a" + "\n\t"))))
    for fi79 := 0; fi79 < 13; fi79++ {
        v80 := GenS4{F0: "", F1: 2147483648}
        fmt.Println(v24, v23, v25)
        v81 := GenS8{F0: 1, F1: (v21 + -9223372036854775808)}
        _ = fi79
        _ = v80
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v77
}
