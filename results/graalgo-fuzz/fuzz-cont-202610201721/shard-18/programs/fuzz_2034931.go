package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 string; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int16; F2 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 uint }
func genFn0(p0 uint) uint8 {
    for fi5 := 0; fi5 < 10; fi5++ {
        switch "a" {
        case "":
            if true {
                p0 += 42
            } else {
                var v6 int16 = -749
                v6 = (v6 << ((p0) & 63))
                _ = v6
            }
        case "\n\t":
            p0 += p0
        default:
            if false {
                if (false && false) {
                    var v7 uint = p0
                    _ = v7
                }
                v8 := GenS1{F0: 2147483647}
                switch fi5 {
                case 32767:
                    var v9 int = fi5
                    _ = v9
                default:
                    var v10 float32 = (-2.25 - 3.14159)
                    _ = v10
                }
                _ = v8
            }
        }
        _ = fi5
    }
    p0 = (p0 % ((42) | 1))
    for fi11 := 0; fi11 < 20; fi11++ {
        p0 += (p0 + 42)
        _ = fi11
    }
    var v12 int16 = -638
    fmt.Println(p0, v12)
    _ = p0
    _ = v12
    return 2
}
func genFn1(p0 uint64, p1 uint16, p2 uint8) uint64 {
    p2 -= p2
    var v13 float32 = 0.0
    v14 := []int{32768}
    v13 = (1e10 / 1.5)
    var v15 int32 = 890
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return p0
}

func main() {
    var v1 int32 = 313
    v2 := []string{"hello", "\n\t", "hello"}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    if false {
        for fi16 := 0; fi16 < 13; fi16++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi16
        }
        if (127 >= 42) {
            v1 = 1
            v17 := GenS1{F0: 623}
            if (true || (!(true && (65535 == 975)))) {
                var v18 bool = (931 <= 18446744073709551615)
                _ = v18
            }
            _ = v17
        } else {
            v19 := GenS1{F0: 2}
            v20 := GenS1{F0: 65535}
            _ = v19
            _ = v20
        }
    }
    fmt.Println(v1)
    for fi21 := 0; fi21 < 14; fi21++ {
        switch (fi21 - (^(fi21))) {
        case 127:
            fmt.Println(v1, v1)
        default:
            v22 := GenS1{F0: 42}
            _ = v22
        }
        if true {
            var v23 uint64 = genFn1(4294967295, 406, 51)
            if (!false) {
                v24 := GenS2{F0: false, F1: "a", F2: (v23 ^ (v23 | genFn1(v23, 314, 255)))}
                v25 := GenS3{F0: -1, F1: 985}
                v23 = (v23 % (((v23 << ((4294967295) & 63))) | 1))
                _ = v24
                _ = v25
            } else {
                var v26 int16 = -358
                _ = v26
            }
            var v27 bool = ((v23 >> ((0) & 63)) != v23)
            _ = v23
            _ = v27
        }
        _ = fi21
    }
    if false {
        v1 = v1
        v2 = append(v2, ("hello" + "foo bar"))
    }
    switch ((("" + (("foo bar" + "a") + (("foo bar" + "\n\t") + (("\n\t" + "foo bar") + "a")))) + "") + "foo bar") {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            var v28 uint64 = 42
            _ = v28
        case "\n\t":
            var v29 uint32 = 134
            _ = v29
        case "hello":
            if false {
                v30 := GenS4{F0: (("hello" + "\n\t") == ("\n\t" + "a"))}
                var v31 int32 = (v1 % (((v1 >> ((18446744073709551615) & 63))) | 1))
                for fi32 := 0; fi32 < 19; fi32++ {
                    for fi33 := 0; fi33 < 4; fi33++ {
                        delete(v3, 2)
                        v34 := GenS3{F0: 1, F1: 42}
                        _ = len(v2)
                        _ = fi33
                        _ = v34
                    }
                    _ = fi32
                }
                _ = v30
                _ = v31
            }
        default:
            var v35 int16 = -600
            _ = v35
        }
    }
    switch 0 {
    case 128:
        if true {
            switch "hello" {
            case "":
                var v36 float32 = 1.5
                _ = v36
            case "foo bar":
                v1 = 1
            case "a":
                for fi37 := 0; fi37 < 19; fi37++ {
                    v1 = (v1 / ((v1) | 1))
                    var v38 int32 = 816
                    v39 := GenS5{F0: 255, F1: 528, F2: "foo bar"}
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            var v40 uint = 323
            v40 = v40
            _ = v40
        }
    case 720:
        v1 = (v1 - (v1 % ((-1) | 1)))
    default:
        if true {
            v1 = (v1 | (^(v1)))
            v1 += v1
        } else {
            switch -2147483648 {
            case -632:
                fmt.Println(v1)
            case -9223372036854775808:
                var v41 uint8 = 127
                _ = v41
            case -39:
                v42 := GenS3{F0: -112, F1: 2147483648}
                _ = v42
            }
            fmt.Println(v1)
        }
    }
    _ = len(v4)
    v1 = 2147483647
    v43 := GenS6{F0: -59}
    var v44 float64 = -2.25
    if false {
        var v45 int8 = 0
        _ = v45
    }
    switch 2147483647 {
    case 4294967295:
        v46 := GenS5{F0: 0, F1: -53, F2: "foo bar"}
        _ = v46
    case 127:
        for fi47 := 0; fi47 < 1; fi47++ {
            if true {
                for fi48 := 0; fi48 < 6; fi48++ {
                    fmt.Println(fi47, fi48, fi47)
                    _ = fi48
                }
            } else {
                v44 += (float64(0.0) * (v44 + v44))
                if false {
                    v49 := v4[0]
                    var v50 uint = 4294967295
                    _ = v49
                    _ = v50
                } else {
                    var v51 uint8 = 42
                    switch (fi47 - fi47) {
                    case -979:
                        v51 = (v51 % ((genFn0(424)) | 1))
                    default:
                        switch (fi47 + 4294967295) {
                        case 2147483647:
                            for fi52 := 0; fi52 < 13; fi52++ {
                                if true {
                                    v44 = v44
                                    if false {
                                        fmt.Println(fi52)
                                    }
                                    v44 = 0.1
                                }
                                fmt.Println(fi52)
                                fmt.Println(v51)
                                _ = fi52
                            }
                        case 860:
                            switch fi47 {
                            case -1:
                                if (1000 != -32768) {
                                    switch "hello" {
                                    case "":
                                        if true {
                                            for fi53 := 0; fi53 < 17; fi53++ {
                                                fmt.Println(fi53, v44, v1)
                                                if true {
                                                    v54 := GenS2{F0: (!false), F1: ("hello" + (("foo bar" + "\n\t") + "a")), F2: 2}
                                                    _ = v54
                                                }
                                                if (385 != 339) {
                                                    if true {
                                                        switch fi47 {
                                                        case -1:
                                                            fmt.Println(v1, v1, fi47)
                                                        default:
                                                            v55 := GenS7{F0: 0, F1: 1}
                                                            _ = v55
                                                        }
                                                        v56 := GenS2{F0: ((fi53 * fi53) != 2147483647), F1: ((("a" + ("foo bar" + (("hello" + "") + (("hello" + ("foo bar" + "\n\t")) + (("hello" + "foo bar") + (("foo bar" + (("hello" + ("a" + "hello")) + "\n\t")) + "hello")))))) + (("foo bar" + "hello") + "\n\t")) + ("" + "")), F2: 162}
                                                        v1 -= (v1 & (v1 ^ 456))
                                                        _ = v56
                                                    }
                                                }
                                                _ = fi53
                                            }
                                        }
                                    case "a":
                                        var v57 uint8 = (v51 << ((237) & 63))
                                        _ = v57
                                    }
                                    v1 += v1
                                    v51 -= (^(v51))
                                } else {
                                    if (!(!true)) {
                                        if (true || false) {
                                            fmt.Println(v1, v44)
                                        }
                                        v44 = float64(((0.0 - 1.5) / 1.5))
                                    }
                                    for fi58 := 0; fi58 < 4; fi58++ {
                                        delete(v3, 4294967295)
                                        _ = fi58
                                    }
                                    v59 := GenS4{F0: (!(true || true))}
                                    _ = v59
                                }
                            case 32767:
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                            case 2147483647:
                                for fi60 := 0; fi60 < 16; fi60++ {
                                    var v61 float32 = (1.5 - (float32(v44) + ((0.1 / 1.5) + (1.5 + 1e10))))
                                    v1 = v1
                                    switch "\n\t" {
                                    case "":
                                        v51 -= v51
                                        v61 = v61
                                    default:
                                        var v62 int64 = 1
                                        _ = v62
                                    }
                                    _ = fi60
                                    _ = v61
                                }
                            }
                        case -1:
                            v63 := v4[2147483647]
                            _ = v63
                        default:
                            switch fi47 {
                            case 558:
                                switch "hello" {
                                case "a":
                                    switch fi47 {
                                    case 32768:
                                        for fi64 := 0; fi64 < 9; fi64++ {
                                            fmt.Println(v51, fi47, v44)
                                            _ = fi64
                                        }
                                    case 720:
                                        v1 -= (v1 | 2147483647)
                                    }
                                case "":
                                    v51 -= v51
                                default:
                                    fmt.Println(v51, v44)
                                }
                            default:
                                delete(v4, 0)
                            }
                        }
                    }
                    _ = v51
                }
            }
            var v65 int64 = 9223372036854775807
            _ = fi47
            _ = v65
        }
    case -2147483648:
        for fi66 := 0; fi66 < 20; fi66++ {
            v67 := GenS8{F0: 2}
            for fi68 := 0; fi68 < 4; fi68++ {
                v44 -= 0.0
                if true {
                    var v69 uint8 = 12
                    _ = v69
                } else {
                    var v70 int32 = -1
                    _ = v70
                }
                _ = fi68
            }
            if false {
                delete(v4, 2147483648)
            } else {
                for fi71 := 0; fi71 < 5; fi71++ {
                    var v72 uint16 = 42
                    _ = fi71
                    _ = v72
                }
                if false {
                    if true {
                        v44 += (v44 - float64(float32(v44)))
                        for fi73 := 0; fi73 < 6; fi73++ {
                            v3[534] = (!true)
                            _ = len(v2)
                            _ = fi73
                        }
                    }
                }
            }
            _ = fi66
            _ = v67
        }
    }
    var v74 float64 = v44
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v74
}
