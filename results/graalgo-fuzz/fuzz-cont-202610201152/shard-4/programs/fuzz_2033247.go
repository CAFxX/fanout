package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int32 }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 int16; F1 int64; F2 uint32 }
type GenS11 struct { F0 int }
func genFn0(p0 int) uint16 {
    fmt.Println(p0)
    var v7 uint32 = 2
    v8 := GenS1{F0: false, F1: 1, F2: 42}
    v9 := GenS2{F0: -11, F1: 4294967295, F2: (!false)}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 107
}

func main() {
    var v1 int32 = 87
    var v2 bool = true
    var v3 uint64 = 1
    var v4 float32 = 1e10
    var v5 bool = (true && v2)
    v6 := make(map[string]int)
    if (2 != 216) {
        fmt.Println(v4)
        v5 = ((1.5 / 1.0) == float32(((float64(1.5) * 0.0) * (float64(v4) - 0.0))))
    }
    if v2 {
        v10 := GenS2{F0: -1, F1: 347, F2: (18446744073709551615 > 65535)}
        v11 := v6[("foo bar" + "")]
        _ = v10
        _ = v11
    } else {
        switch "a" {
        case "\n\t":
            for fi12 := 0; fi12 < 3; fi12++ {
                if (!v5) {
                    var v13 float32 = v4
                    v1 = v1
                    v14 := GenS3{F0: 584}
                    _ = v13
                    _ = v14
                } else {
                    switch (^(fi12)) {
                    case 23:
                        if v5 {
                            for fi15 := 0; fi15 < 19; fi15++ {
                                v16 := GenS4{F0: 255}
                                _ = fi15
                                _ = v16
                            }
                            v17 := []int{fi12, -282, 555, (-(fi12))}
                            _ = v17
                        }
                    case 128:
                        var v18 uint8 = 127
                        _ = v18
                    case 761:
                        if v5 {
                            if true {
                                var v19 uint16 = genFn0(fi12)
                                fmt.Println(v5)
                                _ = v19
                            } else {
                                v20 := []int{9223372036854775807}
                                _ = v20
                            }
                        } else {
                            v21 := GenS5{F0: v2, F1: 90}
                            _ = v21
                        }
                    default:
                        if (216 <= 255) {
                            fmt.Println(fi12, v5, v5)
                        } else {
                            var v22 uint32 = 132
                            _ = v22
                        }
                    }
                    var v23 uint64 = 18446744073709551615
                    var v24 string = ""
                    _ = v23
                    _ = v24
                }
                if (!v5) {
                    if v5 {
                        for fi25 := 0; fi25 < 19; fi25++ {
                            v26 := GenS6{F0: (fi25 >> ((721) & 63))}
                            _ = fi25
                            _ = v26
                        }
                        switch -1 {
                        case 128:
                            switch "" {
                            case "":
                                switch ("a" + (("foo bar" + ("" + ("a" + "foo bar"))) + "a")) {
                                case "hello":
                                    if v2 {
                                        var v27 int = fi12
                                        if true {
                                            _ = len(v6)
                                        } else {
                                            v27 -= 127
                                            v28 := GenS7{F0: (v3 - (v3 >> ((255) & 63)))}
                                            v29 := GenS3{F0: 1}
                                            _ = v28
                                            _ = v29
                                        }
                                        for fi30 := 0; fi30 < 5; fi30++ {
                                            v31 := GenS8{F0: fi30}
                                            _ = fi30
                                            _ = v31
                                        }
                                        _ = v27
                                    }
                                case "a":
                                    switch 128 {
                                    case -78:
                                        for fi32 := 0; fi32 < 19; fi32++ {
                                            switch (^(fi12)) {
                                            case 189:
                                                v5 = false
                                            }
                                            v5 = false
                                            var v33 int32 = 1
                                            _ = fi32
                                            _ = v33
                                        }
                                    case 42:
                                        for fi34 := 0; fi34 < 15; fi34++ {
                                            v35 := []int{fi12, fi34, fi34, int(v1)}
                                            _ = fi34
                                            _ = v35
                                        }
                                    case 2147483647:
                                        v36 := GenS3{F0: 0}
                                        _ = v36
                                    }
                                case "foo bar":
                                    if (v2 || v2) {
                                        v4 = v4
                                    } else {
                                        v5 = v5
                                    }
                                }
                            }
                        case -643:
                            var v37 uint32 = 65535
                            _ = v37
                        }
                        var v38 int32 = v1
                        _ = v38
                    } else {
                        if true {
                            v39 := []int{fi12, fi12}
                            switch fi12 {
                            case 391:
                                v3 += 18446744073709551615
                            case 4294967295:
                                for fi40 := 0; fi40 < 2; fi40++ {
                                    delete(v6, "")
                                    _ = fi40
                                }
                            default:
                                v6[("a" + "a")] = (fi12 / ((fi12) | 1))
                            }
                            var v41 float64 = 3.14159
                            _ = v39
                            _ = v41
                        } else {
                            fmt.Println(fi12)
                        }
                        delete(v6, "")
                        if (!((("foo bar" + ("\n\t" + "\n\t")) + "") <= "")) {
                            if (v3 != (v3 - (v3 + v3))) {
                                v2 = v2
                                fmt.Println(fi12, v4)
                            }
                            for fi42 := 0; fi42 < 1; fi42++ {
                                if (v5 && (2147483647 > (v1 % (((v1 / (((v1 | 1)) | 1))) | 1)))) {
                                    v4 = (-2.25 - 0.1)
                                    switch "foo bar" {
                                    case "\n\t":
                                        v1 = (v1 ^ (v1 - 606))
                                    case "a":
                                        v1 = (-(v1))
                                    case "foo bar":
                                        v3 -= v3
                                    }
                                    v43 := GenS9{F0: 1, F1: 65535}
                                    _ = v43
                                } else {
                                    v44 := []int{4294967295}
                                    _ = v44
                                }
                                v45 := []int{(fi42 ^ fi42), (fi12 | (-(fi42)))}
                                _ = fi42
                                _ = v45
                            }
                            var v46 uint = 649
                            _ = v46
                        } else {
                            var v47 float32 = 0.1
                            fmt.Println(v1, v4)
                            v48 := GenS8{F0: (fi12 / ((32768) | 1))}
                            _ = v47
                            _ = v48
                        }
                    }
                    fmt.Println(v2, v1, v3)
                    for fi49 := 0; fi49 < 3; fi49++ {
                        var v50 uint32 = 0
                        _ = fi49
                        _ = v50
                    }
                } else {
                    var v51 bool = v2
                    switch "a" {
                    case "a":
                        switch (fi12 ^ fi12) {
                        case -2147483648:
                            fmt.Println(v51)
                        }
                    }
                    _ = v51
                }
                _ = fi12
            }
        default:
            fmt.Println(v4)
        }
    }
    v4 = v4
    switch 2147483647 {
    case 0:
        var v52 int8 = -100
        _ = v52
    case 493:
        if v2 {
            var v53 float64 = float64(float32(((float64(((float32(((-2.25 / (-2.25 * ((0.1 + (((float64(1.5) - 0.0) / 0.1) * (3.14159 * 1e10))) * 1e10))) + -2.25)) + ((v4 + float32(1e10)) - v4)) / 0.1)) * 3.14159) / 1.5)))
            v4 = (v4 + -2.25)
            v6[("foo bar" + "foo bar")] = 623
            _ = v53
        }
    case -2147483648:
        _ = len(v6)
    default:
        if v2 {
            if (((0.0 - ((0.1 - 3.14159) + -2.25)) != 3.14159) || (3.14159 == (0.1 / 1.0))) {
                if v5 {
                    var v54 uint64 = v3
                    _ = v54
                }
                if (true && v5) {
                    if v5 {
                        for fi55 := 0; fi55 < 9; fi55++ {
                            fmt.Println(v1, v2)
                            _ = fi55
                        }
                    }
                }
                switch 996 {
                case 32767:
                    v56 := []int{42, 42, 42}
                    _ = v56
                case -2147483648:
                    fmt.Println(v2, v1)
                case 137:
                    if (!v2) {
                        for fi57 := 0; fi57 < 11; fi57++ {
                            for fi58 := 0; fi58 < 4; fi58++ {
                                switch ((("foo bar" + ("\n\t" + "\n\t")) + ("hello" + "\n\t")) + ("a" + "hello")) {
                                case "hello":
                                    fmt.Println(fi57)
                                }
                                v59 := GenS6{F0: (fi58 >> ((1) & 63))}
                                _ = fi58
                                _ = v59
                            }
                            _ = fi57
                        }
                    }
                default:
                    for fi60 := 0; fi60 < 20; fi60++ {
                        v61 := GenS10{F0: int16(fi60), F1: 431, F2: 2}
                        _ = fi60
                        _ = v61
                    }
                }
            }
            v62 := v6["hello"]
            _ = v62
        }
    }
    v63 := v6[""]
    v64 := GenS8{F0: (v63 * 0)}
    switch ("a" + "") {
    case "a":
        v65 := GenS1{F0: true, F1: v1, F2: (v1 >> ((75) & 63))}
        _ = v65
    case "foo bar":
        v66 := GenS6{F0: v63}
        _ = v66
    case "hello":
        v67 := GenS1{F0: (127 >= 0), F1: 1, F2: v1}
        _ = v67
    }
    v68 := GenS11{F0: (v63 + 121)}
    for fi69 := 0; fi69 < 4; fi69++ {
        var v70 int64 = 42
        v70 = v70
        _ = fi69
        _ = v70
    }
    v4 = v4
    for fi71 := 0; fi71 < 19; fi71++ {
        for fi72 := 0; fi72 < 8; fi72++ {
            var v73 bool = (((("a" + "\n\t") + ("hello" + "foo bar")) == (("\n\t" + ("\n\t" + "")) + "")) || v5)
            var v74 int8 = 127
            _ = fi72
            _ = v73
            _ = v74
        }
        if ((("foo bar" + (("" + ("\n\t" + "foo bar")) + "\n\t")) != "foo bar") && true) {
            fmt.Println(v63)
        }
        fmt.Println(v3)
        _ = fi71
    }
    v2 = (!v5)
    var v75 int32 = 0
    v4 += 1.5
    v76 := GenS4{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v68
    _ = v75
    _ = v76
}
