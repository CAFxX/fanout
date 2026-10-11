package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint8; F2 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 uint, p1 int64) int32 {
    v6 := []int{-262}
    v6 = append(v6, 127)
    _ = p0
    _ = p1
    _ = v6
    return 244
}

func main() {
    v1 := []string{("hello" + ""), ("\n\t" + ("\n\t" + "\n\t")), "a"}
    v2 := []string{"\n\t", "foo bar", ("" + ("" + "a"))}
    v3 := make(map[uint32]string)
    var v4 float64 = (float64(float32(0.1)) - 1e10)
    var v5 uint = 255
    v4 = v4
    var v7 float32 = (float32(((1.5 / 1.5) + (-2.25 * (float64(0.0) + v4)))) - (0.1 * 1e10))
    for fi8 := 0; fi8 < 5; fi8++ {
        if true {
            var v9 uint32 = 2
            v4 = (v4 - (((((float64(v7) + 3.14159) - v4) / v4) + (v4 * -2.25)) + 1e10))
            var v10 float32 = v7
            _ = v10
            _ = v9
        } else {
            var v11 int = fi8
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            v7 = v7
            _ = fi12
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v14 uint32 = 1
            v14 += 0
            _ = fi13
            _ = v14
        }
        _ = fi8
    }
    var v15 uint8 = 157
    v5 += v5
    if true {
        v16 := GenS1{F0: 581, F1: 42, F2: (!((1 >= -1) || false))}
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int16 = 812
            switch ((("foo bar" + ("\n\t" + "hello")) + "hello") + ("\n\t" + ("hello" + "foo bar"))) {
            case "\n\t":
                v5 += v5
            case "a":
                var v19 string = "hello"
                _ = v19
            case "foo bar":
                v20 := GenS1{F0: 255, F1: 2, F2: false}
                _ = v20
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                _ = len(v3)
                switch "" {
                case "\n\t":
                    v5 += 838
                case "foo bar":
                    v4 = v4
                default:
                    v22 := GenS2{F0: 2, F1: fi21}
                    _ = v22
                }
                v23 := GenS3{F0: 683}
                if false {
                    var v24 int32 = 42
                    v7 += (v7 - v7)
                    var v25 int16 = (^(v18))
                    _ = v24
                    _ = v25
                } else {
                    v4 = (-2.25 * ((v4 + (v4 * v4)) * v4))
                    v26 := GenS4{F0: v15, F1: (("\n\t" + ("" + "a")) + "a")}
                    v5 = 255
                    _ = v26
                }
                v1 = append(v1, ("hello" + "a"))
                if (2 < 869) {
                    if true {
                        if false {
                            switch ("hello" + "hello") {
                            case "hello":
                                var v27 bool = (42 <= (v5 % ((0) | 1)))
                                _ = v27
                            case "foo bar":
                                var v28 float64 = (v4 + (v4 / v4))
                                _ = v28
                            }
                            v15 = (v15 / ((2) | 1))
                            if true {
                                fmt.Println(v18)
                            } else {
                                if false {
                                    v15 = (^(v15))
                                    switch fi21 {
                                    case 1:
                                        if true {
                                            var v29 string = ""
                                            _ = v29
                                        } else {
                                            var v30 string = "\n\t"
                                            v15 -= 127
                                            _ = v30
                                        }
                                    case 127:
                                        var v31 uint = (v5 * v5)
                                        _ = v31
                                    default:
                                        var v32 uint8 = 1
                                        _ = v32
                                    }
                                } else {
                                    v15 = v15
                                    switch (fi17 ^ fi17) {
                                    case 128:
                                        v4 = 3.14159
                                        v5 += v5
                                    case 32767:
                                        var v33 int16 = (v18 + v18)
                                        _ = v33
                                    case -2147483648:
                                        var v34 float64 = float64((v7 / 1.0))
                                        _ = v34
                                    }
                                    v35 := GenS3{F0: 0}
                                    _ = v35
                                }
                                v2 = append(v2, ("\n\t" + ""))
                                fmt.Println(v15, fi21, v15)
                            }
                        }
                        switch (fi21 << ((v5) & 63)) {
                        case -258:
                            v18 = v18
                        default:
                            fmt.Println(fi17)
                        }
                    }
                    switch ((("\n\t" + "foo bar") + (((("foo bar" + "") + "a") + (("a" + ("hello" + ("\n\t" + "hello"))) + "\n\t")) + "")) + ("" + "")) {
                    case "a":
                        if false {
                            var v36 int64 = 974
                            _ = v36
                        } else {
                            fmt.Println(v15, fi21)
                            var v37 float64 = 3.14159
                            var v38 int64 = 1
                            v39 := GenS3{F0: 258}
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    case "":
                        v40 := GenS2{F0: 4294967295, F1: fi21}
                        _ = v40
                    default:
                        fmt.Println(v15, v18, v4)
                    }
                    if (true && true) {
                        var v41 uint32 = 4294967295
                        _ = v41
                    } else {
                        fmt.Println(v5, v18)
                        var v42 uint16 = 32767
                        var v43 int = fi17
                        v4 = v4
                        _ = v42
                        _ = v43
                    }
                }
                _ = fi21
                _ = v23
            }
            _ = fi17
            _ = v18
        }
        v15 -= v15
        _ = v16
    } else {
        if false {
            v44 := v3[0]
            for fi45 := 0; fi45 < 8; fi45++ {
                for fi46 := 0; fi46 < 4; fi46++ {
                    if true {
                        switch "a" {
                        case "\n\t":
                            if true {
                                v4 -= -2.25
                            } else {
                                v47 := GenS5{F0: (397 > 42)}
                                _ = v47
                            }
                        }
                        switch ("a" + "foo bar") {
                        case "a":
                            if (!(!false)) {
                                fmt.Println(fi46)
                                v48 := GenS2{F0: 4294967295, F1: (fi46 / ((fi46) | 1))}
                                v49 := GenS3{F0: 2147483647}
                                _ = v48
                                _ = v49
                            }
                        case "":
                            switch v44 {
                            case "foo bar":
                                var v50 uint32 = 65535
                                _ = v50
                            case "a":
                                if (!false) {
                                    var v51 string = v44
                                    _ = v51
                                }
                            default:
                                var v52 int8 = -24
                                _ = v52
                            }
                        case "hello":
                            v53 := GenS6{F0: (42 > -673), F1: 255, F2: v44}
                            _ = v53
                        default:
                            delete(v3, 2147483648)
                        }
                        if true {
                            v54 := GenS7{F0: 18446744073709551615}
                            v55 := GenS4{F0: v15, F1: (v44 + (((v44 + "a") + ("hello" + "")) + "foo bar"))}
                            _ = v54
                            _ = v55
                        } else {
                            v7 = 1e10
                            var v56 string = "a"
                            var v57 float64 = ((v4 / v4) / (float64(v7) * -2.25))
                            _ = v56
                            _ = v57
                        }
                    } else {
                        if (-9223372036854775808 <= -381) {
                            fmt.Println(fi46, fi46, v4)
                        }
                        if false {
                            v3[65535] = ("foo bar" + ("" + v44))
                            switch "\n\t" {
                            case "a":
                                if (v15 <= (v15 / ((v15) | 1))) {
                                    var v58 float64 = 1e10
                                    _ = v58
                                }
                            case "\n\t":
                                var v59 int32 = -597
                                _ = v59
                            default:
                                switch v44 {
                                case "a":
                                    v2 = append(v2, ("" + "hello"))
                                }
                                v7 += (((v7 - v7) + v7) * (3.14159 * v7))
                                var v60 int32 = -256
                                _ = v60
                            }
                        }
                    }
                    var v61 uint32 = 1
                    v62 := GenS8{F0: v61}
                    _ = fi46
                    _ = v61
                    _ = v62
                }
                var v63 int8 = -128
                for fi64 := 0; fi64 < 11; fi64++ {
                    var v65 uint32 = 42
                    delete(v3, v65)
                    v3[(v65 << ((4294967295) & 63))] = "\n\t"
                    if true {
                        fmt.Println(v65)
                        if (!((!(382 != -1)) || false)) {
                            fmt.Println(v7)
                            v4 += v4
                            if false {
                                _ = len(v2)
                                v63 += v63
                            }
                        } else {
                            if false {
                                _ = len(v3)
                                v66 := v3[v65]
                                var v67 float64 = (v4 * 0.1)
                                var v68 uint8 = v15
                                _ = v66
                                _ = v67
                                _ = v68
                            } else {
                                delete(v3, v65)
                                var v69 int32 = -393
                                v69 = v69
                                v15 = (v15 & v15)
                                _ = v69
                            }
                            v65 += 0
                        }
                        if false {
                            v1 = append(v1, v44)
                        }
                    }
                    _ = fi64
                    _ = v65
                }
                _ = fi45
                _ = v63
            }
            _ = v44
        }
        v7 -= v7
    }
    v4 = (0.0 * (((3.14159 - (v4 / 1.0)) / 1e10) / (((float64((v7 - 0.1)) - v4) - v4) + (1e10 * v4))))
    var v70 string = "a"
    var v71 float32 = -2.25
    for fi72 := 0; fi72 < 6; fi72++ {
        for fi73 := 0; fi73 < 3; fi73++ {
            fmt.Println(v71, v4, fi72)
            v74 := GenS4{F0: 42, F1: (v70 + ("" + v70))}
            v5 += v5
            switch "\n\t" {
            case "a":
                fmt.Println(v5, v4, v15)
            case "hello":
                fmt.Println(v4, v71, fi73)
            }
            var v75 int = 32768
            _ = fi73
            _ = v74
            _ = v75
        }
        _ = fi72
    }
    v3[245] = (v70 + v70)
    var v76 int16 = 42
    var v77 uint8 = (v15 * v15)
    if ((v5 % (((v5 * 65535)) | 1)) == (^(v5))) {
        fmt.Println(v4, v15)
        var v78 string = (((((("hello" + (v70 + v70)) + "hello") + v70) + v70) + ((v70 + v70) + v70)) + "hello")
        for fi79 := 0; fi79 < 10; fi79++ {
            if false {
                v71 = 1.5
            } else {
                for fi80 := 0; fi80 < 8; fi80++ {
                    v81 := GenS1{F0: 2, F1: 4294967295, F2: true}
                    _ = fi80
                    _ = v81
                }
                for fi82 := 0; fi82 < 10; fi82++ {
                    v83 := v3[557]
                    _ = fi82
                    _ = v83
                }
                for fi84 := 0; fi84 < 16; fi84++ {
                    var v85 uint64 = 4294967295
                    switch v78 {
                    case "foo bar":
                        v4 = (1e10 / ((v4 - float64(((v7 / (1.5 - (v71 / 0.1))) / 1e10))) / v4))
                    default:
                        if (!(2147483647 < 268)) {
                            v70 = v78
                        } else {
                            var v86 uint = (v5 + (v5 & v5))
                            v87 := GenS3{F0: 1}
                            var v88 bool = (false || false)
                            _ = v86
                            _ = v87
                            _ = v88
                        }
                    }
                    _ = fi84
                    _ = v85
                }
            }
            _ = fi79
        }
        _ = v78
    } else {
        v71 = v71
        if false {
            v89 := GenS8{F0: 704}
            _ = v89
        } else {
            _ = len(v2)
        }
        var v90 uint = v5
        _ = v90
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v70
    _ = v71
    _ = v76
    _ = v77
}
