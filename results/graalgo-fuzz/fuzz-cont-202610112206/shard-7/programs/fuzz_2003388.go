package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int16; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int32; F2 int }

func main() {
    v1 := []string{(("hello" + ("a" + "\n\t")) + ((("hello" + ("hello" + "")) + "hello") + "a")), "foo bar", "hello"}
    v2 := make(map[string]int)
    var v3 int64 = -1
    v4 := make(map[int]uint32)
    v5 := []int{42, 42, -693}
    var v6 string = ""
    delete(v2, "a")
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int = (fi7 >> ((897) & 63))
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: v3}
            v3 = (v3 % (((v3 % ((-1) | 1))) | 1))
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v11 uint32 = 42
    v12 := GenS2{F0: 1, F1: 977, F2: 4294967295}
    if (!false) {
        fmt.Println(v11, v3, v3)
        v11 = (^(v11))
        v6 = (v6 + v6)
    } else {
        if true {
            var v13 int16 = -1
            fmt.Println(v13, v13)
            _ = v13
        }
        fmt.Println(v3, v11, v3)
        switch "\n\t" {
        case "a":
            v6 = ("" + v6)
        case "foo bar":
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 int = 42
                _ = fi14
                _ = v15
            }
        }
    }
    var v16 int = 127
    v17 := GenS3{F0: 4294967295}
    if true {
        for fi18 := 0; fi18 < 15; fi18++ {
            if false {
                v6 = (("foo bar" + "foo bar") + "")
                var v19 uint = 226
                switch "foo bar" {
                case "foo bar":
                    for fi20 := 0; fi20 < 11; fi20++ {
                        if true {
                            var v21 int8 = 127
                            v11 = v11
                            _ = v21
                        }
                        _ = fi20
                    }
                case "a":
                    switch 32768 {
                    case -1:
                        for fi22 := 0; fi22 < 5; fi22++ {
                            switch v16 {
                            case -460:
                                v16 = (fi22 / ((42) | 1))
                                if true {
                                    v23 := GenS3{F0: fi18}
                                    fmt.Println(fi18, v19, v19)
                                    v16 += v16
                                    var v24 int32 = 2147483647
                                    _ = v23
                                    _ = v24
                                } else {
                                    v19 = (v19 + (v19 / ((42) | 1)))
                                    fmt.Println(v16, v6)
                                    if true {
                                        if ((1 < 100) || true) {
                                            var v25 int8 = 42
                                            v16 += (v16 / ((fi22) | 1))
                                            _ = v25
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            delete(v2, ("hello" + (v6 + v6)))
                                        case "":
                                            _ = len(v2)
                                        }
                                    } else {
                                        if (!(v19 <= 4294967295)) {
                                            var v26 uint32 = v11
                                            _ = v26
                                        }
                                        var v27 float32 = (0.0 + (-2.25 + 0.0))
                                        _ = v27
                                    }
                                }
                            case 275:
                                v1 = append(v1, "hello")
                            default:
                                if (false && false) {
                                    if (uint64(v19) < 42) {
                                        var v28 int = (fi18 / (((^(fi22))) | 1))
                                        v28 += 4294967295
                                        v28 += (fi22 ^ fi18)
                                        _ = v28
                                    }
                                    var v29 float32 = float32((0.0 + -2.25))
                                    _ = v29
                                }
                            }
                            if false {
                                switch ((v6 + "a") + (("\n\t" + ("foo bar" + (v6 + v6))) + "\n\t")) {
                                case "":
                                    v30 := GenS1{F0: v3}
                                    _ = v30
                                case "\n\t":
                                    v16 -= (v16 >> ((v19) & 63))
                                case "a":
                                    if len(v5) > 0 {
                                        v5[0] = (-(fi18))
                                    }
                                default:
                                    v5 = append(v5, (-(fi18)))
                                }
                                fmt.Println(v11, v16, v6)
                            } else {
                                v6 = v6
                                fmt.Println(fi22, fi18)
                                v11 = 2147483647
                            }
                            v16 = fi18
                            _ = fi22
                        }
                    case -2147483648:
                        var v31 int8 = -42
                        _ = v31
                    case -175:
                        fmt.Println(v19)
                    }
                case "hello":
                    v16 = fi18
                }
                _ = v19
            } else {
                v32 := GenS2{F0: -128, F1: -43, F2: (fi18 ^ v16)}
                _ = v32
            }
            var v33 bool = false
            _ = fi18
            _ = v33
        }
        v6 = v6
        switch v16 {
        case -56:
            var v34 bool = true
            _ = v34
        }
    }
    var v35 float64 = (((float64(((3.14159 * 1.5) / float32((float64((-2.25 * 1.5)) - float64(-2.25))))) + -2.25) * (-2.25 + (1e10 * 3.14159))) / float64(1.5))
    if false {
        v11 -= v11
    }
    var v36 uint8 = 128
    var v37 int16 = 32767
    for fi38 := 0; fi38 < 17; fi38++ {
        v39 := GenS3{F0: (-(fi38))}
        if ((false && (1 != (v16 % (((fi38 + fi38)) | 1)))) || false) {
            v40 := v2[((v6 + "\n\t") + v6)]
            v11 += v11
            _ = v40
        } else {
            var v41 int32 = -2147483648
            v4[(v16 - (v16 ^ (fi38 % (((^(v16))) | 1))))] = (v11 | 609)
            _ = v41
        }
        if false {
            fmt.Println(v35, v6, v11)
            for fi42 := 0; fi42 < 13; fi42++ {
                if (true && (false || (3.14159 > ((1.5 * v35) * 1e10)))) {
                    v16 -= fi42
                    v3 = v3
                    var v43 int32 = 0
                    _ = v43
                }
                _ = fi42
            }
            v11 = (v11 ^ (v11 ^ v11))
        } else {
            var v44 uint = 0
            if ((v37 >> (((^(v44))) & 63)) < v37) {
                v45 := GenS4{F0: 0, F1: v36, F2: 0}
                _ = v45
            } else {
                v11 += (^(v11))
                for fi46 := 0; fi46 < 3; fi46++ {
                    v35 = float64(-2.25)
                    _ = fi46
                }
                v47 := GenS5{F0: 2, F1: -1, F2: (fi38 | 1)}
                _ = v47
            }
            _ = v44
        }
        _ = fi38
        _ = v39
    }
    v3 = 9223372036854775807
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 32768 {
    case 2147483647:
        fmt.Println(v6, v36)
    case 32768:
        var v48 int = v16
        _ = v48
    }
    v35 = float64(1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
