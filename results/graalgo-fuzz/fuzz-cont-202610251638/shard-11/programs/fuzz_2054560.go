package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 string }
func genFn0(p0 int32, p1 uint) string {
    switch -817 {
    case 552:
        var v5 float64 = 1.5
        _ = v5
    case 33:
        var v6 uint16 = 2
        _ = v6
    case 643:
        var v7 string = "hello"
        _ = v7
    default:
        p1 = 65535
    }
    switch ("hello" + "a") {
    case "hello":
        v8 := GenS1{F0: 65535}
        _ = v8
    default:
        v9 := []int{}
        _ = v9
    }
    p1 = 0
    var v10 int64 = -1
    v11 := []int{0, 127, 128}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return ("" + ((("\n\t" + "foo bar") + "\n\t") + ""))
}
func genFn1() uint16 {
    var v12 int32 = 0
    switch "hello" {
    case "foo bar":
        if true {
            v13 := []int{1, 791}
            for fi14 := 0; fi14 < 9; fi14++ {
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := GenS1{F0: 255}
                    fmt.Println(fi14, fi15, fi15)
                    _ = fi15
                    _ = v16
                }
                v12 += v12
                _ = fi14
            }
            if true {
                switch "hello" {
                case "foo bar":
                    _ = len(v13)
                case "a":
                    v17 := GenS2{F0: true, F1: 1, F2: 1}
                    _ = v17
                }
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(fi18, v12)
                    v12 = (v12 % ((v12) | 1))
                    _ = fi18
                }
            } else {
                var v19 int = 128
                _ = v19
            }
            _ = v13
        }
    default:
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 uint16 = 32767
            switch "hello" {
            case "":
                if false {
                    var v22 int32 = -2147483648
                    _ = v22
                } else {
                    fmt.Println(v12)
                    var v23 uint16 = v21
                    if (138 >= 42) {
                        switch (-(fi20)) {
                        case 42:
                            v24 := []int{fi20, fi20}
                            _ = v24
                        case 0:
                            v25 := GenS3{F0: 4294967295, F1: 0}
                            _ = v25
                        case -2147483648:
                            for fi26 := 0; fi26 < 6; fi26++ {
                                v27 := GenS1{F0: (v23 + v21)}
                                v28 := GenS3{F0: 0, F1: 65535}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        default:
                            if false {
                                fmt.Println(v12)
                                var v29 float64 = -2.25
                                _ = v29
                            } else {
                                v30 := []int{(-(fi20)), fi20}
                                v23 = 1
                                for fi31 := 0; fi31 < 6; fi31++ {
                                    v30 = append(v30, (fi31 >> ((18446744073709551615) & 63)))
                                    var v32 float32 = (1e10 + (1.5 - (0.0 / 1.0)))
                                    v12 -= v12
                                    switch ("hello" + ("" + "")) {
                                    case "":
                                        var v33 int8 = 1
                                        _ = v33
                                    case "hello":
                                        var v34 int32 = (v12 % ((v12) | 1))
                                        _ = v34
                                    default:
                                        var v35 bool = false
                                        var v36 int64 = -377
                                        _ = v35
                                        _ = v36
                                    }
                                    fmt.Println(fi20)
                                    v21 = (v21 & v23)
                                    v37 := GenS3{F0: 4294967295, F1: 1}
                                    _ = fi31
                                    _ = v32
                                    _ = v37
                                }
                                _ = v30
                            }
                        }
                    }
                    _ = v23
                }
            case "hello":
                var v38 int16 = 1000
                _ = v38
            default:
                v39 := GenS1{F0: 201}
                _ = v39
            }
            _ = fi20
            _ = v21
        }
    }
    var v40 float32 = float32((1.5 + float64((0.1 + (1.5 + 1.5)))))
    _ = v12
    _ = v40
    return 42
}

func main() {
    var v1 int16 = -422
    var v2 bool = (1 < 417)
    var v3 bool = true
    v4 := make(map[string]uint32)
    for fi41 := 0; fi41 < 3; fi41++ {
        if (((v1 ^ (v1 % ((v1) | 1))) >= (-(v1))) || v3) {
            var v42 int = fi41
            v3 = v2
            _ = v42
        } else {
            switch ((("hello" + "hello") + ("hello" + "")) + "\n\t") {
            case "hello":
                switch "a" {
                case "hello":
                    var v43 int16 = 0
                    _ = v43
                default:
                    v44 := GenS2{F0: false, F1: 255, F2: 42}
                    _ = v44
                }
            }
        }
        if ("hello" <= "") {
            fmt.Println(v2, v2)
        } else {
            var v45 int64 = 1
            fmt.Println(v3, v3)
            _ = v45
        }
        switch (fi41 / (((fi41 | (fi41 >> ((18446744073709551615) & 63)))) | 1)) {
        case -9223372036854775808:
            v46 := GenS1{F0: 35}
            _ = v46
        default:
            var v47 uint8 = 76
            _ = v47
        }
        _ = fi41
    }
    var v48 float64 = (3.14159 + 0.1)
    for fi49 := 0; fi49 < 4; fi49++ {
        if (!true) {
            v50 := []int{}
            if v3 {
                var v51 int32 = 2147483647
                v48 = (float64(1.5) - (v48 + v48))
                _ = v51
            } else {
                var v52 int64 = 42
                if (v2 && false) {
                    if len(v50) > 0 {
                        v50[0] = fi49
                    }
                }
                _ = v52
            }
            if v2 {
                v53 := v4[("hello" + "\n\t")]
                v54 := GenS4{F0: 18446744073709551615}
                v1 -= (v1 % (((^(v1))) | 1))
                _ = v53
                _ = v54
            }
            _ = v50
        } else {
            if false {
                v4["\n\t"] = 854
                if true {
                    v1 = int16(fi49)
                    var v55 uint64 = 1
                    _ = v55
                }
                var v56 int8 = 127
                _ = v56
            }
            switch (("foo bar" + ("hello" + "\n\t")) + ("hello" + "\n\t")) {
            case "a":
                if v2 {
                    for fi57 := 0; fi57 < 13; fi57++ {
                        delete(v4, "hello")
                        switch "a" {
                        case "hello":
                            v58 := []int{-204}
                            _ = v58
                        case "":
                            v59 := GenS5{F0: 1, F1: "foo bar"}
                            _ = v59
                        default:
                            v60 := GenS6{F0: 343, F1: 127}
                            _ = v60
                        }
                        fmt.Println(v48)
                        _ = fi57
                    }
                }
            default:
                v48 -= 1.5
            }
            var v61 bool = v3
            _ = v61
        }
        fmt.Println(fi49)
        v62 := []int{0, fi49, fi49, (-(fi49))}
        _ = fi49
        _ = v62
    }
    if v2 {
        for fi63 := 0; fi63 < 18; fi63++ {
            if false {
                switch fi63 {
                case -1:
                    var v64 int64 = 669
                    _ = v64
                default:
                    v3 = true
                }
                v1 += 1000
            } else {
                if v2 {
                    v1 = (v1 / (((v1 % ((v1) | 1))) | 1))
                    v2 = ((true && false) || v3)
                    v2 = v2
                }
                v3 = (42 >= 42)
            }
            fmt.Println(v3, v2)
            _ = fi63
        }
        var v65 uint16 = 32767
        _ = v65
    } else {
        for fi66 := 0; fi66 < 16; fi66++ {
            v67 := []int{(^(fi66)), fi66}
            var v68 int = 886
            v69 := v4[(("" + ("hello" + ("a" + ""))) + "")]
            _ = fi66
            _ = v67
            _ = v68
            _ = v69
        }
    }
    var v70 uint64 = 4294967295
    v48 = v48
    var v71 int = 1
    var v72 uint32 = 2147483647
    switch 336 {
    case 0:
        for fi73 := 0; fi73 < 4; fi73++ {
            if v2 {
                v2 = (v1 <= v1)
                var v74 uint32 = (v72 >> ((214) & 63))
                _ = v74
            }
            _ = fi73
        }
    default:
        for fi75 := 0; fi75 < 14; fi75++ {
            for fi76 := 0; fi76 < 19; fi76++ {
                var v77 float32 = 0.1
                switch fi76 {
                case -241:
                    v78 := GenS2{F0: v2, F1: 194, F2: 42}
                    _ = v78
                case 1:
                    if ((((v77 - v77) * float32((v48 * ((1.5 - float64(((v77 - (v77 / v77)) / float32(v48)))) - float64(-2.25))))) / 1e10) == (v77 / (v77 + 0.1))) {
                        var v79 float32 = v77
                        v77 = (0.1 + v79)
                        var v80 int32 = 726
                        v81 := v4["hello"]
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            v77 = v79
                        case "a":
                            var v82 int64 = 1
                            _ = v82
                        }
                        if false {
                            var v83 int32 = 1
                            delete(v4, ("" + ("\n\t" + (("a" + "hello") + ("foo bar" + "a")))))
                            _ = v83
                        }
                        fmt.Println(v81)
                        v72 += v81
                        _ = v79
                        _ = v80
                        _ = v81
                    } else {
                        v71 = (fi76 << ((255) & 63))
                        v84 := GenS7{F0: ("\n\t" + "")}
                        var v85 int16 = v1
                        _ = v84
                        _ = v85
                    }
                case -9223372036854775808:
                    var v86 int64 = 427
                    _ = v86
                }
                fmt.Println(v1)
                _ = fi76
                _ = v77
            }
            v87 := GenS1{F0: 2}
            v72 -= v72
            _ = fi75
            _ = v87
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v70
    _ = v71
    _ = v72
}
