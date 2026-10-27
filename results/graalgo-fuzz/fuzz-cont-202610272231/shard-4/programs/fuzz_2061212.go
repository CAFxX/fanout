package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint; F2 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 int8 }

func main() {
    var v1 int8 = 100
    var v2 int8 = v1
    v3 := make(map[string]string)
    v4 := []uint32{0, 0}
    var v5 int16 = 1
    _ = len(v4)
    switch "foo bar" {
    case "hello":
        v6 := GenS1{F0: 127, F1: ((-(v2)) > v2), F2: 0}
        _ = v6
    default:
        var v7 int = 32768
        _ = v7
    }
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "a" {
    case "\n\t":
        var v8 uint8 = 128
        _ = v8
    case "":
        if true {
            var v9 int64 = 1
            if (v1 == (v2 ^ (v1 + (v1 + v2)))) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 float64 = 0.1
                    v9 = 0
                    v9 += (v9 ^ (v9 * -9223372036854775808))
                    _ = fi10
                    _ = v11
                }
                fmt.Println(v1, v2, v1)
            }
            if ((42 > 1) || true) {
                v3["foo bar"] = ("\n\t" + "a")
                if true {
                    v9 += (v9 | (-(v9)))
                    v12 := GenS2{F0: "foo bar"}
                    var v13 uint = 443
                    _ = v12
                    _ = v13
                }
            } else {
                switch 32767 {
                case 42:
                    v3["foo bar"] = "a"
                default:
                    v3["hello"] = "a"
                }
                var v14 int16 = v5
                fmt.Println(v1, v1)
                _ = v14
            }
            _ = v9
        } else {
            v5 = (v5 + v5)
        }
    case "hello":
        for fi15 := 0; fi15 < 9; fi15++ {
            v2 += (v2 * (v2 * 0))
            v16 := v3[""]
            _ = fi15
            _ = v16
        }
    }
    var v17 string = ("" + "hello")
    var v18 int8 = (v1 | v1)
    var v19 float64 = (1e10 * (float64((1.5 - 0.1)) / (1e10 * (float64(0.1) / (-2.25 / 1e10)))))
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int32 = -517
        v17 = v17
        for fi22 := 0; fi22 < 3; fi22++ {
            v5 += (v5 - v5)
            var v23 int32 = v21
            switch (fi22 | (fi20 ^ (fi20 | (fi22 & -2147483648)))) {
            case 32768:
                v24 := GenS2{F0: "a"}
                _ = v24
            }
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    if len(v4) > 0 {
        v4[0] = 145
    }
    if false {
        fmt.Println(v5)
        if true {
            if true {
                var v25 string = v17
                _ = v25
            }
        }
    } else {
        var v26 int8 = -119
        for fi27 := 0; fi27 < 18; fi27++ {
            fmt.Println(v17)
            switch ((v17 + v17) + ("" + "hello")) {
            case "":
                v26 = (-(v1))
            case "hello":
                if false {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        var v29 int64 = -751
                        if (!true) {
                            var v30 uint = 18446744073709551615
                            _ = v30
                        } else {
                            var v31 int16 = v5
                            fmt.Println(v2, v2)
                            var v32 int8 = (v26 + v1)
                            _ = v31
                            _ = v32
                        }
                        switch (("a" + (v17 + v17)) + v17) {
                        case "hello":
                            v33 := GenS2{F0: "hello"}
                            _ = v33
                        case "\n\t":
                            var v34 int64 = 233
                            _ = v34
                        case "foo bar":
                            v18 = v2
                        }
                        _ = fi28
                        _ = v29
                    }
                    if true {
                        v1 -= v18
                        v35 := GenS3{F0: (v5 % ((v5) | 1))}
                        _ = len(v4)
                        _ = v35
                    } else {
                        v2 += (-(v2))
                        switch 407 {
                        case 26:
                            var v36 int8 = 100
                            _ = v36
                        default:
                            if ("\n\t" <= v17) {
                                for fi37 := 0; fi37 < 9; fi37++ {
                                    v18 = (v18 ^ v18)
                                    _ = fi37
                                }
                            }
                        }
                        fmt.Println(fi27, v5)
                    }
                    for fi38 := 0; fi38 < 19; fi38++ {
                        var v39 int64 = 154
                        if len(v4) > 0 {
                            v4[0] = 836
                        }
                        _ = fi38
                        _ = v39
                    }
                } else {
                    v40 := GenS4{F0: 32767, F1: 709, F2: v17}
                    fmt.Println(v19, v26)
                    _ = v40
                }
            default:
                fmt.Println(v5, v18, v17)
            }
            _ = fi27
        }
        for fi41 := 0; fi41 < 7; fi41++ {
            fmt.Println(v17)
            delete(v3, "hello")
            v42 := GenS2{F0: "a"}
            _ = fi41
            _ = v42
        }
        _ = v26
    }
    if false {
        switch 1 {
        case 1:
            var v43 uint16 = 1
            _ = v43
        case 9223372036854775807:
            switch "\n\t" {
            case "foo bar":
                v44 := GenS5{F0: 0}
                _ = v44
            case "hello":
                _ = len(v4)
            case "":
                for fi45 := 0; fi45 < 10; fi45++ {
                    if false {
                        var v46 int = -2147483648
                        v47 := v3[("\n\t" + ("hello" + v17))]
                        _ = v46
                        _ = v47
                    } else {
                        v1 += v1
                    }
                    var v48 float64 = float64(((1e10 * 0.0) * (((-2.25 - (float32(float64((3.14159 * 1e10))) + 1e10)) - (3.14159 / 1.0)) - (1e10 * 0.0))))
                    v49 := GenS5{F0: -2147483648}
                    _ = fi45
                    _ = v48
                    _ = v49
                }
            default:
                fmt.Println(v18, v17)
            }
        case 32767:
            switch ((v17 + v17) + (("" + v17) + (("foo bar" + v17) + (("" + v17) + v17)))) {
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = 42
                }
            default:
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            }
        default:
            _ = len(v3)
        }
        var v50 uint32 = 0
        _ = v50
    } else {
        var v51 int8 = -1
        v4 = append(v4, 2147483647)
        v52 := v3[v17]
        _ = v51
        _ = v52
    }
    switch 127 {
    case -636:
        var v53 uint64 = 1
        _ = v53
    case -558:
        v54 := GenS3{F0: (v5 >> ((255) & 63))}
        _ = v54
    default:
        v19 = (((v19 / v19) / 1.0) + float64(((((3.14159 * 1.5) / 0.1) - 1.5) / (0.0 + 1e10))))
    }
    if true {
        v55 := GenS2{F0: v17}
        var v56 uint8 = 255
        v57 := GenS3{F0: (v5 | (-(v5)))}
        _ = v55
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 17; fi58++ {
        switch (fi58 - fi58) {
        case 654:
            var v59 uint64 = 960
            _ = v59
        case 2147483647:
            for fi60 := 0; fi60 < 6; fi60++ {
                if (42 != 669) {
                    v17 = "a"
                } else {
                    var v61 string = ""
                    switch (v17 + v17) {
                    case "a":
                        v5 = v5
                    case "foo bar":
                        delete(v3, "hello")
                    }
                    var v62 float64 = v19
                    _ = v61
                    _ = v62
                }
                switch v17 {
                case "foo bar":
                    fmt.Println(v5, fi58)
                case "a":
                    if true {
                        v17 = v17
                    } else {
                        fmt.Println(v19)
                        var v63 int = -972
                        fmt.Println(v18)
                        v64 := GenS6{F0: int64(fi60), F1: 45}
                        _ = len(v3)
                        _ = v63
                        _ = v64
                    }
                    if false {
                        var v65 int16 = v5
                        switch v17 {
                        case "":
                            v19 = (float64(0.0) + 1e10)
                        case "hello":
                            fmt.Println(v5, fi60, v5)
                        }
                        _ = v65
                    } else {
                        v66 := GenS6{F0: 9223372036854775807, F1: v2}
                        if len(v4) > 0 {
                            v4[0] = 1
                        }
                        v67 := GenS1{F0: 0, F1: (42 < 0), F2: -1}
                        v19 += v19
                        _ = v66
                        _ = v67
                    }
                    v17 = v17
                case "\n\t":
                    var v68 uint8 = 42
                    _ = v68
                default:
                    fmt.Println(v19, v18, v17)
                }
                _ = fi60
            }
        default:
            var v69 int64 = 42
            _ = v69
        }
        _ = len(v3)
        fmt.Println(v17, v18)
        _ = fi58
    }
    for fi70 := 0; fi70 < 9; fi70++ {
        var v71 uint64 = 2
        if false {
            v17 = (((v17 + v17) + "hello") + v17)
        }
        if true {
            var v72 float32 = 0.0
            v1 = -1
            var v73 bool = true
            _ = v72
            _ = v73
        }
        _ = fi70
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
