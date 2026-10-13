package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int32 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool; F1 string; F2 int8 }
type GenS10 struct { F0 uint; F1 uint64; F2 int16 }
type GenS11 struct { F0 int32; F1 uint }
func genFn0() uint8 {
    if true {
        if true {
            var v6 float32 = float32(-2.25)
            _ = v6
        }
        v7 := []int{}
        _ = v7
    } else {
        var v8 uint16 = 1
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: 442, F1: fi9, F2: 2147483647}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS2{F0: -100, F1: "hello"}
            _ = fi11
            _ = v12
        }
        _ = v8
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 bool = true
        if (!v14) {
            var v15 string = ("hello" + "hello")
            _ = v15
        } else {
            v14 = v14
            v16 := GenS2{F0: 100, F1: ("foo bar" + "foo bar")}
            var v17 uint16 = 255
            _ = v16
            _ = v17
        }
        if v14 {
            var v18 uint16 = 32767
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS3{F0: 9223372036854775807, F1: v14}
                fmt.Println(v14, fi19)
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                switch fi13 {
                case 42:
                    v14 = ((0.0 / ((3.14159 / 1.0) / ((0.1 / -2.25) - 0.0))) > -2.25)
                case 4294967295:
                    if v14 {
                        v14 = (2147483648 > 852)
                        v22 := []int{(fi13 / ((fi13) | 1)), fi13, fi13, (fi13 << ((42) & 63))}
                        _ = v22
                    } else {
                        var v23 uint16 = (^(v18))
                        _ = v23
                    }
                case 1:
                    v24 := []int{}
                    _ = v24
                }
                if (v14 || (256 < -1)) {
                    v25 := GenS2{F0: -1, F1: (("foo bar" + "\n\t") + "hello")}
                    if false {
                        v26 := GenS2{F0: 127, F1: ("" + "\n\t")}
                        _ = v26
                    } else {
                        v27 := GenS4{F0: 0, F1: 914}
                        v28 := GenS5{F0: 0}
                        _ = v27
                        _ = v28
                    }
                    _ = v25
                }
                _ = fi21
            }
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    var v29 uint8 = 127
    var v30 bool = (333 > 723)
    _ = v29
    _ = v30
    return (^(v29))
}
func genFn1() int32 {
    var v31 int16 = 1
    fmt.Println(v31)
    fmt.Println(v31)
    v32 := GenS6{F0: 32767, F1: 78}
    _ = v31
    _ = v32
    return 0
}

func main() {
    var v1 int64 = 0
    var v2 int32 = 42
    var v3 uint8 = 18
    v4 := []string{}
    var v5 uint32 = 4294967295
    _ = len(v4)
    fmt.Println(v5)
    if len(v4) > 0 {
        v4[0] = ("hello" + "")
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    v33 := GenS7{F0: uint16(v3), F1: 351}
    if false {
        if true {
            var v34 int16 = 32767
            _ = v34
        }
        if true {
            switch "" {
            case "foo bar":
                v2 = 0
            case "hello":
                v5 -= (^(v5))
            default:
                v35 := GenS6{F0: 1000, F1: v3}
                _ = v35
            }
        } else {
            if (2147483647 == 1) {
                v3 -= v3
                v1 = 970
            }
            v36 := GenS8{F0: v1}
            _ = v36
        }
    } else {
        v2 = 42
        switch ("" + "a") {
        case "\n\t":
            if false {
                v3 -= (v3 | (v3 | v3))
            } else {
                if (42 >= -1) {
                    v2 += v2
                }
            }
        default:
            v3 -= (v3 >> ((354) & 63))
        }
    }
    var v37 int64 = v1
    var v38 float32 = (((0.1 * 1e10) - float32(1.5)) / 1e10)
    fmt.Println(v38, v2)
    fmt.Println(v1, v2)
    v4 = append(v4, "a")
    if ((v37 == 0) && ("foo bar" <= "hello")) {
        var v39 uint = 18446744073709551615
        for fi40 := 0; fi40 < 2; fi40++ {
            v3 = v3
            if (255 > 65535) {
                v37 += v1
                fmt.Println(v39, v3)
                fmt.Println(v38, v5)
            }
            var v41 float64 = (((1e10 - -2.25) * 0.0) / -2.25)
            _ = fi40
            _ = v41
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            for fi43 := 0; fi43 < 19; fi43++ {
                v37 = v37
                v4 = append(v4, (("hello" + ("" + "hello")) + "hello"))
                if false {
                    var v44 uint = 4294967295
                    _ = v44
                }
                if (-2.25 > 0.0) {
                    if false {
                        var v45 bool = true
                        _ = v45
                    }
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                }
                fmt.Println(fi43, v37)
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                var v46 int32 = -2147483648
                switch "hello" {
                case "\n\t":
                    fmt.Println(v46, v39, v3)
                case "hello":
                    var v47 string = "\n\t"
                    _ = v47
                case "a":
                    switch "hello" {
                    case "foo bar":
                        v46 -= genFn1()
                        v3 -= (v3 << ((v39) & 63))
                        var v48 string = ""
                        _ = v48
                    case "hello":
                        if (1.5 == float64(float32((1e10 + 1e10)))) {
                            v49 := GenS3{F0: v1, F1: (55 != 127)}
                            v37 = 0
                            _ = v49
                        }
                    default:
                        if (((v39 | 818) != (v39 + v39)) || false) {
                            fmt.Println(v3, v1)
                            v50 := GenS8{F0: v37}
                            _ = v50
                        } else {
                            v46 += int32(v3)
                            if true {
                                if true {
                                    switch ("" + "foo bar") {
                                    case "a":
                                        var v51 int8 = 108
                                        _ = v51
                                    case "hello":
                                        v52 := GenS9{F0: false, F1: ("\n\t" + ""), F2: 0}
                                        _ = v52
                                    case "":
                                        var v53 uint32 = v5
                                        _ = v53
                                    }
                                }
                                switch fi42 {
                                case 127:
                                    fmt.Println(v39)
                                case 143:
                                    var v54 uint16 = 42
                                    _ = v54
                                case -2147483648:
                                    fmt.Println(v38, v46)
                                default:
                                    var v55 float64 = (1e10 / 0.1)
                                    var v56 uint = (v39 - v39)
                                    if false {
                                        v3 = (v3 / (((v3 & (v3 ^ 0))) | 1))
                                        switch (fi43 * (fi42 << ((v56) & 63))) {
                                        case -2147483648:
                                            v46 = v2
                                        default:
                                            fmt.Println(v38, v39, v1)
                                        }
                                        switch fi42 {
                                        case 32768:
                                            if true {
                                                if len(v4) > 0 {
                                                    v4[0] = ""
                                                }
                                            }
                                        case -2147483648:
                                            if true {
                                                if (-2.25 == v55) {
                                                    var v57 uint = 42
                                                    v5 -= (v5 << ((v39) & 63))
                                                    _ = v57
                                                } else {
                                                    var v58 int8 = 42
                                                    _ = v58
                                                }
                                            }
                                        case 9223372036854775807:
                                            v3 += genFn0()
                                        }
                                    } else {
                                        v4 = append(v4, "foo bar")
                                        v59 := GenS10{F0: 0, F1: 23, F2: 638}
                                        _ = v59
                                    }
                                    _ = v55
                                    _ = v56
                                }
                            }
                            if (genFn0() < (v3 ^ 30)) {
                                v60 := GenS6{F0: 42, F1: 47}
                                var v61 uint64 = 692
                                var v62 uint8 = v3
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            v3 -= (v3 ^ 128)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v63 := GenS11{F0: genFn1(), F1: v39}
                            _ = v63
                        }
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                    }
                default:
                    var v64 uint32 = v5
                    _ = v64
                }
                v65 := GenS10{F0: (v39 / (((v39 * 775)) | 1)), F1: 0, F2: int16(v1)}
                _ = fi43
                _ = v46
                _ = v65
            }
            _ = fi42
        }
        _ = v39
    }
    if true {
        var v66 float32 = (0.1 - float32(float64((0.1 - 1.5))))
        _ = v66
    } else {
        var v67 float32 = v38
        if (!false) {
            v68 := GenS6{F0: 713, F1: 0}
            var v69 int8 = 42
            var v70 string = "a"
            _ = v68
            _ = v69
            _ = v70
        }
        var v71 float32 = v38
        _ = v67
        _ = v71
    }
    v4 = append(v4, "hello")
    fmt.Println(v5, v1, v3)
    v3 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
