package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint; F1 int16; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 bool) uint16 {
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{fi5, fi5}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 255
}
func genFn1(p0 string, p1 int) string {
    fmt.Println(p1, p0)
    v7 := GenS1{F0: (p1 | p1)}
    v8 := GenS1{F0: p1}
    if true {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 string = ((((("hello" + p0) + "a") + ("foo bar" + "foo bar")) + "") + (("\n\t" + p0) + "hello"))
            var v11 int64 = 9223372036854775807
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(p0)
    } else {
        p0 = p0
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 string = (((p0 + (p0 + p0)) + ((p0 + (("\n\t" + p0) + p0)) + "foo bar")) + ((p0 + p0) + ""))
            var v14 int64 = 0
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    if false {
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ("" + ((p0 + "hello") + ((((p0 + (p0 + p0)) + p0) + (p0 + p0)) + "a")))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 0
    v3 := []uint32{}
    var v15 float32 = float32((-2.25 + 0.0))
    if true {
        v2 += v2
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int16 = 32767
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: 0}
        var v19 string = "foo bar"
        _ = v18
        _ = v19
    }
    v2 = v2
    v20 := GenS1{F0: 42}
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 string = ("" + "")
        if (!false) {
            fmt.Println(v2, v2)
            v22 = "\n\t"
        }
        _ = fi21
        _ = v22
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            if true {
                var v23 bool = (true && (!(("" >= "") && (true && false))))
                v23 = (9223372036854775807 < -707)
                if (0 > -1) {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        var v25 int16 = 42
                        v23 = true
                        _ = fi24
                        _ = v25
                    }
                    _ = len(v3)
                    if true {
                        switch "foo bar" {
                        case "hello":
                            v15 -= v15
                        case "a":
                            var v26 uint32 = 4294967295
                            _ = v26
                        case "":
                            v1[127] = ""
                        }
                        v27 := GenS1{F0: 1}
                        v2 -= v2
                        _ = v27
                    }
                }
                _ = v23
            }
            v2 -= (v2 * (v2 & 2147483647))
        } else {
            for fi28 := 0; fi28 < 18; fi28++ {
                v15 += ((((v15 / (v15 - v15)) + (v15 * ((3.14159 - v15) - (v15 / v15)))) + v15) + float32((1.5 - -2.25)))
                v2 += v2
                _ = fi28
            }
            if (((("a" + "a") + "\n\t") + "\n\t") != "hello") {
                var v29 uint32 = 65535
                fmt.Println(v15, v15)
                if true {
                    fmt.Println(v2, v29)
                } else {
                    v15 = (float32(1e10) + ((float32(0.1) * ((3.14159 / 1e10) * -2.25)) - v15))
                    var v30 uint = 1
                    var v31 uint16 = 303
                    _ = v30
                    _ = v31
                }
                _ = v29
            }
        }
    case "foo bar":
        switch 127 {
        case 0:
            v32 := GenS1{F0: 32767}
            _ = v32
        }
    case "hello":
        fmt.Println(v2, v2)
    default:
        v2 = (v2 >> ((951) & 63))
    }
    var v33 int = 32767
    for fi34 := 0; fi34 < 10; fi34++ {
        switch "foo bar" {
        case "":
            for fi35 := 0; fi35 < 15; fi35++ {
                v36 := GenS1{F0: -9223372036854775808}
                if (182 == 95) {
                    var v37 int8 = -1
                    if true {
                        fmt.Println(fi35, v15)
                    }
                    if false {
                        var v38 uint16 = 32767
                        v39 := v1[1]
                        var v40 int32 = v2
                        _ = v38
                        _ = v39
                        _ = v40
                    } else {
                        v41 := GenS1{F0: 9223372036854775807}
                        _ = v41
                    }
                    fmt.Println(v15, v33)
                    _ = v37
                }
                _ = fi35
                _ = v36
            }
        }
        _ = fi34
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        var v43 string = "a"
        _ = fi42
        _ = v43
    }
    var v44 int16 = 0
    v2 += -1
    var v45 string = ("" + "a")
    switch 32768 {
    case 150:
        v46 := GenS2{F0: false, F1: v33}
        _ = v46
    case 538:
        for fi47 := 0; fi47 < 17; fi47++ {
            v48 := GenS2{F0: true, F1: 4294967295}
            _ = fi47
            _ = v48
        }
    }
    var v49 int = 42
    v1[644] = "hello"
    var v50 uint16 = 255
    if (712 >= 294) {
        if (-1 <= 1) {
            for fi51 := 0; fi51 < 6; fi51++ {
                for fi52 := 0; fi52 < 16; fi52++ {
                    var v53 uint = 78
                    if true {
                        var v54 uint8 = uint8(v53)
                        if (v54 < v54) {
                            var v55 uint16 = (^(v50))
                            _ = len(v1)
                            _ = v55
                        } else {
                            if false {
                                var v56 int8 = 1
                                var v57 int = fi51
                                var v58 int64 = -467
                                _ = v56
                                _ = v57
                                _ = v58
                            } else {
                                var v59 string = ""
                                var v60 int = (v49 ^ (v49 + -1))
                                _ = v59
                                _ = v60
                            }
                            v15 = v15
                            if ((153 > (v54 * v54)) && (!((true && (65535 <= 640)) && (73 <= 842)))) {
                                var v61 int16 = v44
                                if true {
                                    fmt.Println(v33, v50, v45)
                                    v45 = v45
                                }
                                v62 := GenS1{F0: v33}
                                _ = v61
                                _ = v62
                            } else {
                                var v63 uint = v53
                                if false {
                                    v64 := GenS3{F0: (v63 + 294), F1: v44, F2: 100}
                                    _ = len(v3)
                                    if (!(-9223372036854775808 != 628)) {
                                        if len(v3) > 0 {
                                            v3[0] = 2147483648
                                        }
                                        if true {
                                            if (!true) {
                                                v49 += 128
                                                _ = len(v3)
                                                fmt.Println(v33, v44, v44)
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    }
                                    v54 += (^(v54))
                                    var v65 uint8 = v54
                                    _ = v64
                                    _ = v65
                                } else {
                                    v53 += v53
                                    v45 = ((v45 + v45) + "")
                                }
                                v66 := GenS1{F0: (v33 + (v49 + -110))}
                                _ = v63
                                _ = v66
                            }
                            v67 := GenS4{F0: 127, F1: v50}
                            _ = v67
                        }
                        v3 = append(v3, 2147483647)
                        _ = v54
                    }
                    if false {
                        v44 = (v44 / (((v44 / ((v44) | 1))) | 1))
                    }
                    _ = fi52
                    _ = v53
                }
                _ = fi51
            }
            switch (v49 ^ (-(v33))) {
            case 42:
                var v68 uint16 = (^(v50))
                _ = v68
            default:
                v50 += v50
            }
        } else {
            fmt.Println(v15, v2)
            fmt.Println(v44)
        }
        switch v45 {
        case "\n\t":
            if (932 == 905) {
                v33 -= v33
                var v69 uint16 = 65535
                v70 := GenS5{F0: 2, F1: "foo bar"}
                _ = v69
                _ = v70
            } else {
                v15 -= v15
                var v71 int16 = 0
                _ = v71
            }
        case "hello":
            for fi72 := 0; fi72 < 14; fi72++ {
                v1[2147483648] = "foo bar"
                fmt.Println(v2, v45)
                _ = fi72
            }
        default:
            switch (((v45 + (v45 + v45)) + "\n\t") + "") {
            case "":
                var v73 string = v45
                _ = v73
            default:
                var v74 uint32 = 2147483647
                _ = v74
            }
        }
    }
    for fi75 := 0; fi75 < 2; fi75++ {
        switch (^(v33)) {
        case 2147483647:
            var v76 string = v45
            _ = v76
        case -9223372036854775808:
            v50 = 10
        case 0:
            v45 = v45
        }
        _ = fi75
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v44
    _ = v45
    _ = v49
    _ = v50
}
