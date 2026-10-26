package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }
type GenS3 struct { F0 bool; F1 int64; F2 string }
type GenS4 struct { F0 uint16; F1 int16; F2 int32 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint16; F1 int; F2 int64 }
type GenS7 struct { F0 uint; F1 uint64; F2 uint }
type GenS8 struct { F0 uint16; F1 int32; F2 int8 }
func genFn0() string {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint = 661
        v6 := []int{(^(fi4)), fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    if false {
        var v7 uint8 = 42
        v7 += 127
        v8 := []int{}
        _ = v7
        _ = v8
    }
    v9 := []int{1, 127, 32768}
    var v10 uint8 = 0
    v10 = (v10 | 42)
    _ = v10
    _ = v9
    return ""
}
func genFn1() int16 {
    for fi11 := 0; fi11 < 11; fi11++ {
        if ((-2.25 <= 1e10) || (127 != 0)) {
            var v12 string = "foo bar"
            v13 := GenS1{F0: 424}
            var v14 float32 = (0.0 - float32(1e10))
            _ = v12
            _ = v13
            _ = v14
        } else {
            var v15 int16 = -959
            switch ("\n\t" + "a") {
            case "":
                var v16 uint64 = 689
                _ = v16
            case "a":
                var v17 int8 = 1
                _ = v17
            }
            _ = v15
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        if (0 != 255) {
            fmt.Println(fi18)
            var v19 bool = true
            for fi20 := 0; fi20 < 5; fi20++ {
                if (!(287 == 222)) {
                    v19 = v19
                    switch ((("" + "foo bar") + ("a" + ("" + ""))) + "hello") {
                    case "\n\t":
                        if v19 {
                            if (18446744073709551615 > 18446744073709551615) {
                                v21 := GenS1{F0: 150}
                                _ = v21
                            }
                            var v22 uint16 = 0
                            _ = v22
                        } else {
                            v23 := GenS1{F0: 2}
                            _ = v23
                        }
                    case "hello":
                        v19 = ((0.1 + ((1.5 * (0.0 - float32((0.0 / 1.5)))) + 1.5)) == -2.25)
                        var v24 bool = (42 != 163)
                        var v25 int64 = -9223372036854775808
                        _ = v24
                        _ = v25
                    case "a":
                        v26 := GenS2{F0: 32767, F1: v19, F2: 18446744073709551615}
                        _ = v26
                    default:
                        fmt.Println(fi18)
                    }
                    switch (-(fi18)) {
                    case 821:
                        v27 := GenS3{F0: (1e10 < 0.1), F1: 42, F2: "hello"}
                        _ = v27
                    }
                } else {
                    v28 := []int{}
                    if v19 {
                        fmt.Println(fi18, fi18)
                        if (((1.5 > float32((0.1 + -2.25))) || (-1 == -128)) || (!v19)) {
                            var v29 float32 = 1.5
                            if len(v28) > 0 {
                                v28[0] = 128
                            }
                            _ = v29
                        }
                    } else {
                        if false {
                            v19 = (!false)
                            fmt.Println(v19)
                        } else {
                            v30 := GenS3{F0: (!((76 >= 1) || (("hello" + (("hello" + "a") + "a")) >= "\n\t"))), F1: 0, F2: "foo bar"}
                            if true {
                                if ((fi20 ^ fi18) > -201) {
                                    v19 = (2147483647 != 65535)
                                    v19 = false
                                }
                                fmt.Println(fi18)
                                v31 := GenS2{F0: 342, F1: ((fi18 + fi20) >= fi18), F2: 18446744073709551615}
                                _ = v31
                            }
                            v28 = append(v28, fi20)
                            _ = v30
                        }
                    }
                    v19 = v19
                    _ = v28
                }
                _ = fi20
            }
            _ = v19
        } else {
            var v32 float64 = (float64(1e10) * float64(0.0))
            v32 -= ((v32 / (v32 - -2.25)) + (v32 * v32))
            var v33 int32 = 1
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 16; fi34++ {
            v35 := []int{fi34}
            fmt.Println(fi34, fi34)
            _ = fi34
            _ = v35
        }
        _ = fi18
    }
    v36 := GenS2{F0: 308, F1: (!(2147483647 != -2147483648)), F2: 606}
    for fi37 := 0; fi37 < 11; fi37++ {
        var v38 float64 = 1e10
        for fi39 := 0; fi39 < 17; fi39++ {
            if (true && true) {
                v40 := GenS3{F0: ((-576 == 2147483647) || (1000 != -1)), F1: 0, F2: ""}
                var v41 uint32 = 42
                var v42 int = (fi37 & (fi37 % ((fi39) | 1)))
                _ = v40
                _ = v41
                _ = v42
            }
            _ = fi39
        }
        var v43 uint16 = 647
        _ = fi37
        _ = v38
        _ = v43
    }
    var v44 int32 = -2147483648
    _ = v36
    _ = v44
    return -32768
}

func main() {
    var v1 float32 = (3.14159 * (1e10 * (-2.25 * 3.14159)))
    var v2 int8 = 127
    v3 := make(map[string]bool)
    var v45 int = 700
    var v46 uint64 = 427
    if (v2 > 87) {
        v47 := []int{}
        _ = v47
    }
    switch "foo bar" {
    case "hello":
        for fi48 := 0; fi48 < 13; fi48++ {
            var v49 int16 = 569
            _ = fi48
            _ = v49
        }
    case "\n\t":
        v1 -= 0.1
    case "":
        v50 := GenS4{F0: 255, F1: 42, F2: 795}
        _ = v50
    default:
        fmt.Println(v45, v45, v1)
    }
    v51 := GenS2{F0: 2, F1: ((^(v2)) != (v2 ^ (-(v2)))), F2: 65535}
    v45 += v45
    delete(v3, "")
    v46 -= 311
    v52 := GenS5{F0: 754, F1: (!(0.0 <= 1.5))}
    switch "\n\t" {
    case "":
        var v53 uint8 = 0
        _ = v53
    case "\n\t":
        switch ("foo bar" + ("a" + "a")) {
        case "foo bar":
            v46 = 2
        case "a":
            for fi54 := 0; fi54 < 14; fi54++ {
                v46 = (v46 >> ((243) & 63))
                _ = fi54
            }
        }
    default:
        for fi55 := 0; fi55 < 17; fi55++ {
            if true {
                v56 := []int{(v45 | (fi55 + (v45 << ((49) & 63)))), (v45 | (v45 - (v45 + fi55)))}
                switch (fi55 << ((65535) & 63)) {
                case 128:
                    var v57 uint16 = 32767
                    _ = v57
                case -2147483648:
                    var v58 int = 9223372036854775807
                    _ = v58
                case 42:
                    v2 += 100
                default:
                    var v59 int64 = -1
                    _ = v59
                }
                for fi60 := 0; fi60 < 7; fi60++ {
                    switch "hello" {
                    case "":
                        var v61 int16 = 673
                        switch fi55 {
                        case 128:
                            v61 -= v61
                        }
                        v62 := GenS1{F0: 158}
                        v61 = (^(v61))
                        v45 = (fi60 + (v45 | int(v61)))
                        if true {
                            v3["a"] = false
                            v63 := GenS6{F0: 0, F1: fi55, F2: -920}
                            fmt.Println(v61, v46)
                            if true {
                                var v64 int16 = genFn1()
                                var v65 int = 4294967295
                                _ = v64
                                _ = v65
                            } else {
                                var v66 uint32 = 578
                                if false {
                                    v45 = (fi60 * 258)
                                }
                                fmt.Println(v66, v61)
                                _ = v66
                            }
                            _ = v63
                        } else {
                            v1 -= 0.1
                            var v67 uint = 42
                            if false {
                                v2 = (v2 / ((127) | 1))
                                v1 = (((0.0 - (v1 * v1)) / 3.14159) / float32((1e10 / float64(v1))))
                                v46 = 0
                                var v68 uint8 = 247
                                _ = v68
                            } else {
                                var v69 int16 = (v61 % ((-32768) | 1))
                                _ = v69
                            }
                            _ = v67
                        }
                        switch (fi60 * (v45 - fi60)) {
                        case -547:
                            switch (v45 & fi60) {
                            case 4294967295:
                                var v70 int = (^(fi55))
                                _ = v70
                            default:
                                v45 = (^(v45))
                            }
                        case 128:
                            if true {
                                v46 += 0
                                fmt.Println(v2)
                            } else {
                                var v71 float64 = 0.0
                                _ = v71
                            }
                        case 0:
                            var v72 bool = (!(true || (2147483648 == uint32(v46))))
                            _ = v72
                        }
                        _ = v61
                        _ = v62
                    case "\n\t":
                        if true {
                            var v73 float32 = 0.0
                            _ = len(v56)
                            _ = v73
                        } else {
                            var v74 uint16 = 1
                            v1 = v1
                            _ = v74
                        }
                    default:
                        var v75 uint32 = 2
                        _ = v75
                    }
                    v2 = 100
                    _ = fi60
                }
                _ = v56
            } else {
                for fi76 := 0; fi76 < 18; fi76++ {
                    if true {
                        v77 := []int{291, (v45 & -1)}
                        fmt.Println(fi55)
                        v45 = (v45 % (((v45 & 32767)) | 1))
                        if ((!(2147483648 < 2)) || false) {
                            v78 := GenS4{F0: 65535, F1: -562, F2: -422}
                            var v79 string = ((("" + ("hello" + "a")) + "\n\t") + ("hello" + "foo bar"))
                            fmt.Println(fi55)
                            _ = v78
                            _ = v79
                        } else {
                            v2 -= (v2 << ((4294967295) & 63))
                            var v80 string = ("foo bar" + "foo bar")
                            _ = v80
                        }
                        if false {
                            var v81 int16 = -533
                            v82 := GenS7{F0: 4294967295, F1: v46, F2: 222}
                            if true {
                                v1 += (v1 + 0.0)
                            } else {
                                var v83 uint16 = 255
                                _ = v83
                            }
                            _ = v81
                            _ = v82
                        }
                        v84 := GenS8{F0: 255, F1: 1, F2: (v2 % (((v2 + (v2 | 127))) | 1))}
                        fmt.Println(v1)
                        _ = v77
                        _ = v84
                    }
                    _ = fi76
                }
            }
            switch "foo bar" {
            case "a":
                v85 := GenS8{F0: 611, F1: 0, F2: v2}
                _ = v85
            case "":
                var v86 int = (fi55 * fi55)
                _ = v86
            case "\n\t":
                for fi87 := 0; fi87 < 6; fi87++ {
                    v45 -= fi55
                    if (1.5 != v1) {
                        var v88 uint64 = v46
                        v45 += fi87
                        fmt.Println(v46, fi55, v88)
                        _ = v88
                    }
                    _ = fi87
                }
            }
            fmt.Println(v46)
            _ = fi55
        }
    }
    v89 := []int{}
    v45 = (v45 * 42)
    var v90 int8 = v2
    fmt.Println(v2, v45, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v89:", len(v89), v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v51
    _ = v52
    _ = v89
    _ = v90
}
