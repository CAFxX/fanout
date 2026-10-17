package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 string; F2 string }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 bool }
type GenS9 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint16, p2 uint64) int64 {
    var v5 int8 = 1
    p0 = true
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := []int{}
        var v8 bool = p0
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return -582
}
func genFn1(p0 int64, p1 uint16) string {
    fmt.Println(p0)
    var v10 int32 = -508
    var v11 uint = 4294967295
    var v12 float64 = 1e10
    v10 += 1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483647
    var v3 float64 = 3.14159
    v4 := make(map[string]string)
    if (!false) {
        var v13 int64 = -1
        fmt.Println(v3)
        fmt.Println(v13, v2, v13)
        _ = v13
    }
    if true {
        var v14 float32 = (1e10 - -2.25)
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 uint = 42
            v17 := GenS1{F0: int(v16)}
            for fi18 := 0; fi18 < 6; fi18++ {
                switch "" {
                case "":
                    switch ("foo bar" + "") {
                    case "":
                        if ((((float64(1e10) - 0.1) + (v3 - v3)) / float64(v14)) < (-2.25 - 1.5)) {
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                switch "\n\t" {
                                case "a":
                                    v14 = v14
                                case "\n\t":
                                    if false {
                                        v19 := GenS2{F0: 199}
                                        _ = v19
                                    }
                                }
                                if true {
                                    var v20 uint16 = 460
                                    var v21 string = "\n\t"
                                    _ = v20
                                    _ = v21
                                } else {
                                    v14 -= ((v14 - ((-2.25 / 0.1) + (1e10 * v14))) * ((v14 / v14) - 1.5))
                                    var v22 string = "foo bar"
                                    _ = v22
                                }
                                if true {
                                    v23 := GenS3{F0: 2, F1: 137, F2: (^(v2))}
                                    fmt.Println(fi18, fi15)
                                    v14 -= 3.14159
                                    var v24 uint64 = 1
                                    _ = v23
                                    _ = v24
                                }
                            case "":
                                var v25 string = "foo bar"
                                var v26 int32 = -2147483648
                                v27 := GenS4{F0: 1}
                                _ = v25
                                _ = v26
                                _ = v27
                            case "hello":
                                switch "a" {
                                case "foo bar":
                                    v16 += v16
                                default:
                                    v2 -= (v2 + (^(v2)))
                                }
                            }
                            if true {
                                v14 = v14
                            }
                            if (!false) {
                                switch "foo bar" {
                                case "hello":
                                    v28 := GenS5{F0: 42, F1: ("hello" + ("\n\t" + ((((("\n\t" + "foo bar") + "foo bar") + "\n\t") + (("\n\t" + "foo bar") + "hello")) + ((("" + ("\n\t" + ("" + "foo bar"))) + "foo bar") + ("a" + "foo bar"))))), F2: "\n\t"}
                                    _ = v28
                                case "foo bar":
                                    fmt.Println(fi15, fi15)
                                }
                                var v29 int16 = 42
                                var v30 int = (fi15 | fi18)
                                _ = v29
                                _ = v30
                            }
                        } else {
                            fmt.Println(fi18, v16)
                            if false {
                                fmt.Println(v14, fi18)
                                var v31 string = "hello"
                                v31 = v31
                                _ = v31
                            }
                        }
                    case "foo bar":
                        if true {
                            var v32 uint16 = 35
                            _ = v32
                        }
                    case "hello":
                        delete(v4, (("" + (("hello" + "foo bar") + "\n\t")) + "a"))
                    }
                case "hello":
                    var v33 bool = (-310 < 42)
                    _ = v33
                case "a":
                    v16 -= v16
                default:
                    if (-2.25 != ((1.5 * ((float64(float32((v3 * (-2.25 / (-2.25 - float64(v14)))))) + (3.14159 - (1.5 - v3))) / ((v3 - (v3 * (1.5 - float64(v14)))) / (0.0 + v3)))) / (3.14159 - (((float64(0.0) - float64(v14)) / ((v3 / float64(-2.25)) + float64(v14))) - (-2.25 * (v3 / (0.0 / (0.0 * v3)))))))) {
                        if (0 == 938) {
                            switch 32768 {
                            case 32768:
                                if (!true) {
                                    v3 = (v3 * ((((1e10 * (v3 - v3)) + ((v3 + (((v3 / -2.25) * float64((v14 - float32(0.0)))) * 0.1)) + v3)) * float64(v14)) * v3))
                                    v34 := GenS6{F0: v16, F1: "hello"}
                                    v35 := GenS2{F0: 0}
                                    _ = v34
                                    _ = v35
                                } else {
                                    v36 := GenS1{F0: 521}
                                    v2 += 431
                                    v3 -= 3.14159
                                    v3 = float64((v14 + 3.14159))
                                    var v37 uint16 = 42
                                    var v38 uint16 = (v37 << ((v16) & 63))
                                    v37 = v37
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                            case 2147483647:
                                var v39 float64 = 1.5
                                _ = v39
                            case -1:
                                v16 += v16
                            }
                            fmt.Println(v16, v16, v2)
                        }
                        if (false || (((^(v2)) != v2) || true)) {
                            v40 := GenS7{F0: 32767}
                            _ = v40
                        }
                        fmt.Println(fi15, v16, fi15)
                        if true {
                            fmt.Println(fi15, v2)
                        }
                        var v41 uint = (^(v16))
                        _ = v41
                    } else {
                        var v42 uint = v16
                        if (-73 == 42) {
                            var v43 string = "a"
                            _ = v43
                        } else {
                            v3 = (1.5 + 3.14159)
                        }
                        _ = v42
                    }
                }
                v44 := v4[("a" + ("foo bar" + ""))]
                _ = fi18
                _ = v44
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    } else {
        if (!true) {
            switch 32768 {
            case 2147483647:
                v4["a"] = ("hello" + ("a" + "\n\t"))
            case -1:
                if (false || false) {
                    for fi45 := 0; fi45 < 7; fi45++ {
                        fmt.Println(v2, fi45)
                        _ = fi45
                    }
                    var v46 uint16 = 814
                    v46 -= v46
                    _ = v46
                } else {
                    fmt.Println(v2, v2)
                }
            }
            v47 := GenS3{F0: (^(v2)), F1: 714, F2: v2}
            v2 = v2
            _ = v47
        } else {
            v48 := []int{128, 42}
            _ = v48
        }
        fmt.Println(v3, v3)
        var v49 float64 = v3
        _ = v49
    }
    fmt.Println(v2)
    v3 = v3
    if false {
        if false {
            switch -875 {
            case -1:
                fmt.Println(v3)
            case 0:
                v3 = 0.1
            default:
                switch "foo bar" {
                case "a":
                    v50 := GenS5{F0: -815, F1: "hello", F2: "hello"}
                    _ = v50
                case "foo bar":
                    v3 -= v3
                }
            }
            switch "hello" {
            case "foo bar":
                delete(v4, ("foo bar" + ""))
            }
            if (-34 >= -9223372036854775808) {
                v51 := GenS8{F0: 685, F1: (true && ((false || true) || false))}
                fmt.Println(v3, v3)
                v3 = ((((float64(((3.14159 + 0.0) * float32((v3 / (0.1 - v3))))) - v3) * float64(-2.25)) / 0.1) * float64(-2.25))
                _ = v51
            } else {
                if false {
                    fmt.Println(v2)
                } else {
                    v52 := v4[(("foo bar" + ("foo bar" + (("foo bar" + "\n\t") + ("hello" + ("foo bar" + "foo bar"))))) + "foo bar")]
                    v53 := GenS4{F0: 2}
                    switch 9223372036854775807 {
                    case -2147483648:
                        v2 = (v2 | v2)
                    case 42:
                        for fi54 := 0; fi54 < 18; fi54++ {
                            if true {
                                v3 = (0.0 - 3.14159)
                                for fi55 := 0; fi55 < 5; fi55++ {
                                    v56 := GenS2{F0: 42}
                                    _ = fi55
                                    _ = v56
                                }
                                v3 += 0.0
                            }
                            switch (fi54 ^ (fi54 >> ((uint(v2)) & 63))) {
                            case -9223372036854775808:
                                v57 := []int{(fi54 - fi54), fi54}
                                _ = v57
                            default:
                                v52 = ((v52 + (v52 + ("a" + "a"))) + "hello")
                            }
                            _ = fi54
                        }
                    case -1:
                        if true {
                            v2 = (v2 >> ((42) & 63))
                            fmt.Println(v52, v52)
                        } else {
                            if true {
                                for fi58 := 0; fi58 < 19; fi58++ {
                                    fmt.Println(fi58, fi58)
                                    v2 = v2
                                    _ = fi58
                                }
                                v59 := GenS9{F0: 2}
                                v3 = v3
                                _ = v59
                            }
                        }
                    default:
                        fmt.Println(v3, v52, v3)
                    }
                    _ = v52
                    _ = v53
                }
            }
        } else {
            if ((128 == 127) && (-1 < 0)) {
                var v60 uint8 = 42
                v2 += v2
                v60 += (^(v60))
                _ = v60
            }
            v3 = 1.5
        }
        v61 := []int{0, 127, 128, 0}
        _ = v61
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        if (false || false) {
            v63 := []int{}
            _ = v63
        }
        for fi64 := 0; fi64 < 18; fi64++ {
            v65 := GenS1{F0: fi64}
            var v66 float32 = -2.25
            _ = fi64
            _ = v65
            _ = v66
        }
        _ = fi62
    }
    fmt.Println(v3, v3)
    var v67 float32 = (-2.25 - float32(v3))
    for fi68 := 0; fi68 < 11; fi68++ {
        fmt.Println(v67)
        _ = fi68
    }
    switch 230 {
    case 32768:
        var v69 uint64 = 4294967295
        _ = v69
    case 2147483647:
        var v70 float64 = v3
        _ = v70
    case 301:
        v3 += ((v3 / (0.1 / (float64((float32(v3) * (1e10 / 1e10))) / ((v3 / 1e10) * v3)))) - (v3 * 1.5))
    }
    if true {
        for fi71 := 0; fi71 < 5; fi71++ {
            v3 -= ((v3 / float64(((3.14159 + v67) + (-2.25 / v67)))) - (3.14159 * 1e10))
            _ = fi71
        }
        switch "hello" {
        case "hello":
            for fi72 := 0; fi72 < 8; fi72++ {
                fmt.Println(v3)
                _ = fi72
            }
        case "a":
            v73 := GenS9{F0: 32767}
            _ = v73
        case "\n\t":
            var v74 string = "hello"
            _ = v74
        }
    } else {
        if (3.14159 != (v3 / ((0.1 / ((v3 + -2.25) - (v3 - v3))) + (v3 * v3)))) {
            v75 := GenS1{F0: -1}
            _ = v75
        } else {
            if true {
                for fi76 := 0; fi76 < 19; fi76++ {
                    v77 := GenS6{F0: 0, F1: "foo bar"}
                    v3 = 3.14159
                    _ = fi76
                    _ = v77
                }
                v3 += v3
            } else {
                v3 -= v3
            }
            delete(v4, "\n\t")
            for fi78 := 0; fi78 < 4; fi78++ {
                v79 := []int{fi78, fi78, fi78}
                v67 = (v67 - (v67 - (0.0 * (float32(0.0) + ((v67 - (0.1 + (1.5 - 0.0))) + v67)))))
                _ = fi78
                _ = v79
            }
        }
        var v80 uint16 = 42
        _ = v80
    }
    v3 -= float64(((((1e10 + (v67 * float32(v3))) / v67) / (((1e10 / (float32((v3 - (-2.25 + (((v3 + v3) / 1.0) + 0.0)))) / 3.14159)) * (float32(v3) - 0.0)) + (1.5 - (v67 - 1e10)))) - v67))
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
}
