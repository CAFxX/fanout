package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int; F2 int8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 int16; F1 uint8; F2 int32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 uint16; F1 int8; F2 int8 }
func genFn0(p0 bool) int {
    p0 = (-197 <= -985)
    switch -9223372036854775808 {
    case 215:
        if (p0 && false) {
            v7 := GenS1{F0: 32767}
            v8 := GenS1{F0: 4294967295}
            fmt.Println(p0)
            _ = v7
            _ = v8
        } else {
            v9 := []int{42, 32768}
            var v10 string = (("foo bar" + "\n\t") + ("a" + ""))
            _ = v10
            _ = v9
        }
    case -1:
        var v11 float64 = (float64((3.14159 + (float32(0.0) * ((3.14159 / -2.25) * (1e10 - ((1.5 / 1.5) * (((1.5 + 1.5) * 1.5) + float32((0.1 / -2.25))))))))) + 1.5)
        _ = v11
    case 128:
        var v12 uint64 = 4294967295
        _ = v12
    default:
        var v13 uint8 = 1
        _ = v13
    }
    p0 = (4294967295 <= 287)
    switch ("hello" + "hello") {
    case "foo bar":
        fmt.Println(p0)
    }
    _ = p0
    return 4294967295
}

func main() {
    var v1 float32 = float32(0.0)
    v2 := []uint32{42}
    var v3 bool = true
    v4 := make(map[string]int)
    v5 := []string{("\n\t" + (("hello" + "") + (((("a" + ("hello" + "")) + (((("" + ("" + "foo bar")) + "\n\t") + "hello") + "\n\t")) + (("foo bar" + "") + "foo bar")) + ((("a" + ("hello" + (((("a" + "foo bar") + "hello") + ("" + ("a" + "foo bar"))) + ""))) + ("" + ("a" + "\n\t"))) + ("" + "a")))))}
    v6 := []uint32{0, 81, 2}
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 float32 = (v1 * float32(3.14159))
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        if v3 {
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := GenS2{F0: v3, F1: 365, F2: -33}
                switch "\n\t" {
                case "foo bar":
                    if (127 != 0) {
                        v19 := v4["\n\t"]
                        _ = v19
                    }
                case "hello":
                    if (-9223372036854775808 <= 0) {
                        v20 := GenS3{F0: 0, F1: 0}
                        _ = v20
                    }
                default:
                    fmt.Println(fi17, fi16)
                    var v21 string = (("hello" + ("a" + "\n\t")) + ((("a" + "\n\t") + "hello") + ""))
                    v22 := GenS4{F0: 42}
                    _ = v21
                    _ = v22
                }
                v23 := GenS1{F0: fi16}
                if true {
                    var v24 float64 = 1e10
                    var v25 float64 = float64(1e10)
                    _ = len(v6)
                    if v3 {
                        var v26 string = ""
                        var v27 uint = 65535
                        _ = len(v4)
                        _ = v26
                        _ = v27
                    } else {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v3 = v3
                        } else {
                            _ = len(v2)
                        }
                        fmt.Println(v25, v24)
                        v24 = (v24 + float64(v1))
                    }
                    var v28 float64 = v25
                    v29 := v4[(("" + (("hello" + "") + "")) + ((("a" + ("hello" + "foo bar")) + ("a" + "hello")) + "foo bar"))]
                    v30 := GenS5{F0: -54, F1: 42}
                    v31 := GenS5{F0: 1, F1: -891}
                    var v32 bool = (127 > 241)
                    fmt.Println(v32, v29, v32)
                    var v33 int32 = 2147483647
                    v1 -= -2.25
                    var v34 int = 127
                    _ = v24
                    _ = v25
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                _ = fi17
                _ = v18
                _ = v23
            }
        } else {
            var v35 uint = 217
            for fi36 := 0; fi36 < 18; fi36++ {
                v1 = v1
                if v3 {
                    v6 = append(v6, 992)
                    v1 -= 3.14159
                } else {
                    v1 = v1
                    var v37 int16 = 515
                    v38 := GenS5{F0: -1, F1: 1}
                    _ = v37
                    _ = v38
                }
                _ = fi36
            }
            var v39 int64 = 831
            _ = v35
            _ = v39
        }
        if v3 {
            var v40 uint8 = 1
            v41 := GenS6{F0: 769, F1: v40, F2: 0}
            var v42 float64 = 3.14159
            _ = v40
            _ = v41
            _ = v42
        } else {
            v3 = v3
        }
        _ = fi16
    }
    v3 = false
    fmt.Println(v3, v1)
    _ = len(v2)
    if (2 <= 225) {
        if false {
            if v3 {
                var v43 int32 = -877
                var v44 float64 = 0.0
                for fi45 := 0; fi45 < 20; fi45++ {
                    v43 -= v43
                    for fi46 := 0; fi46 < 5; fi46++ {
                        v47 := GenS5{F0: v43, F1: -32768}
                        v43 = (v43 | v43)
                        if (-101 >= 0) {
                            fmt.Println(v43, v3)
                            if (!(42 == 4294967295)) {
                                delete(v4, ("a" + ("hello" + "a")))
                            } else {
                                if (v43 > -2147483648) {
                                    var v48 uint = 1
                                    _ = v48
                                }
                                v4[("foo bar" + "\n\t")] = -299
                                var v49 float32 = v1
                                v44 = v44
                                _ = v49
                            }
                            _ = len(v2)
                            var v50 uint64 = 2
                            if v3 {
                                switch genFn0((1000 < 0)) {
                                case -9223372036854775808:
                                    var v51 int16 = 1000
                                    _ = v51
                                case 2147483647:
                                    switch "a" {
                                    case "\n\t":
                                        var v52 uint64 = (v50 | v50)
                                        switch ("" + "foo bar") {
                                        case "":
                                            v1 += v1
                                        case "a":
                                            var v53 float32 = (1e10 / v1)
                                            _ = v53
                                        }
                                        _ = v52
                                    }
                                case -115:
                                    v43 = v43
                                default:
                                    v44 = v44
                                }
                            }
                            v54 := GenS3{F0: 0, F1: 2}
                            _ = v50
                            _ = v54
                        }
                        _ = fi46
                        _ = v47
                    }
                    var v55 uint = 183
                    _ = fi45
                    _ = v55
                }
                _ = v43
                _ = v44
            } else {
                switch "\n\t" {
                case "a":
                    v1 -= (v1 / (v1 + v1))
                case "":
                    v1 -= 0.1
                case "\n\t":
                    v1 = float32((0.1 * 3.14159))
                }
                if v3 {
                    if (v3 || (!v3)) {
                        if v3 {
                            if (1.5 > (3.14159 - (((((((3.14159 * float64(3.14159)) * 3.14159) * 0.0) + (0.1 / 1e10)) - (((((float64(((v1 * v1) - (v1 / (0.1 + ((float32((1e10 / 1.0)) * 1e10) * 3.14159))))) - float64(v1)) + ((0.0 - (1.5 * (1e10 + 1e10))) + 0.0)) / -2.25) / float64((-2.25 / v1))) / float64(v1))) * 0.1) + (0.1 + 1e10)))) {
                                fmt.Println(v3)
                                if true {
                                    v56 := GenS7{F0: -32768}
                                    _ = v56
                                }
                                v3 = (((0.0 + (1e10 + -2.25)) / float64(1.5)) > -2.25)
                            }
                            v57 := GenS4{F0: 65535}
                            v3 = v3
                            _ = v57
                        }
                        var v58 uint64 = 732
                        _ = v58
                    }
                    var v59 int64 = 959
                    _ = v59
                }
            }
            v1 -= v1
            v60 := GenS8{F0: (255 == 0), F1: 0}
            _ = v60
        }
        v61 := GenS9{F0: 32767, F1: -34, F2: 0}
        for fi62 := 0; fi62 < 12; fi62++ {
            var v63 int16 = -1
            _ = len(v4)
            _ = fi62
            _ = v63
        }
        _ = v61
    }
    v3 = true
    var v64 bool = false
    v64 = (v64 && false)
    for fi65 := 0; fi65 < 4; fi65++ {
        delete(v4, "foo bar")
        var v66 uint16 = 500
        _ = fi65
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
