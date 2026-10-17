package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int32; F1 int8; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16; F1 uint16; F2 string }
type GenS9 struct { F0 bool; F1 int32; F2 uint32 }
type GenS10 struct { F0 bool; F1 uint32 }
type GenS11 struct { F0 int8; F1 int64; F2 uint32 }
type GenS12 struct { F0 int32; F1 string }
func genFn0() int64 {
    var v4 string = "\n\t"
    v5 := []int{449}
    var v6 int64 = 496
    for fi7 := 0; fi7 < 7; fi7++ {
        if (!(42 >= 42)) {
            v6 -= (v6 >> ((255) & 63))
            if len(v5) > 0 {
                v5[0] = (fi7 >> ((65535) & 63))
            }
            if (262 <= 4294967295) {
                v4 = v4
            } else {
                switch 127 {
                case 42:
                    fmt.Println(v4, v6, fi7)
                case 960:
                    v6 = -774
                default:
                    if true {
                        v8 := GenS1{F0: 32768, F1: 49, F2: ((fi7 & 419) > (fi7 >> ((82) & 63)))}
                        v9 := GenS1{F0: fi7, F1: 0, F2: true}
                        _ = v8
                        _ = v9
                    }
                }
            }
        }
        v10 := GenS2{F0: ("foo bar" + v4)}
        fmt.Println(fi7, fi7, v6)
        _ = fi7
        _ = v10
    }
    var v11 uint8 = 127
    _ = v11
    _ = v4
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 bool) uint {
    if (1.5 >= 1e10) {
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(p0, p0)
            var v13 float32 = ((3.14159 + 3.14159) / (((0.0 + (0.0 * 0.1)) + ((1.5 / 1.5) - (((1e10 - 1.5) / 1.0) / 1.5))) + (0.1 - 0.0)))
            _ = fi12
            _ = v13
        }
    }
    switch "a" {
    case "a":
        var v14 string = ("a" + (("a" + (((("hello" + "") + "a") + "") + "")) + "foo bar"))
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        if (!((-2.25 >= float64((1.5 / (1.5 / float32(3.14159))))) || ((42 < 458) || p0))) {
            v16 := []int{401, (fi15 & fi15), fi15}
            v17 := GenS2{F0: "foo bar"}
            for fi18 := 0; fi18 < 5; fi18++ {
                fmt.Println(fi18, fi15)
                var v19 int = fi18
                switch (^(fi15)) {
                case 179:
                    v16 = append(v16, fi18)
                }
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        } else {
            if p0 {
                switch "" {
                case "\n\t":
                    v20 := GenS3{F0: 2147483647, F1: 514, F2: 2}
                    _ = v20
                case "foo bar":
                    fmt.Println(p0, fi15)
                case "a":
                    var v21 bool = (!(-1 <= genFn0()))
                    _ = v21
                default:
                    p0 = p0
                }
            } else {
                v22 := []int{4294967295, -372, 4294967295, (-(fi15))}
                _ = v22
            }
            var v23 uint64 = 2
            _ = v23
        }
        p0 = (1000 <= 32767)
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 float32 = ((((0.0 + 1e10) - (0.0 - 3.14159)) - 0.1) * float32(1e10))
            v26 := GenS4{F0: 0, F1: 222}
            fmt.Println(p0)
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi15
    }
    fmt.Println(p0)
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 int64 = genFn0()
        v29 := GenS3{F0: 0, F1: 2, F2: 255}
        _ = fi27
        _ = v28
        _ = v29
    }
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 int32 = 152
    v2 := make(map[uint32]bool)
    var v3 int64 = -1
    v30 := GenS3{F0: 327, F1: 18446744073709551615, F2: 311}
    switch 0 {
    case 2147483647:
        var v31 int8 = -125
        _ = v31
    }
    var v32 uint32 = 4294967295
    switch ("a" + "") {
    case "":
        var v33 int16 = -32768
        _ = v33
    case "\n\t":
        var v34 uint64 = 4294967295
        _ = v34
    case "a":
        v35 := v2[(v32 + (^(v32)))]
        _ = v35
    default:
        v36 := GenS5{F0: -2147483648, F1: 1, F2: 4}
        _ = v36
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        for fi38 := 0; fi38 < 13; fi38++ {
            if false {
                switch "foo bar" {
                case "hello":
                    var v39 string = ("\n\t" + ("" + ("hello" + "\n\t")))
                    _ = v39
                case "\n\t":
                    v40 := GenS6{F0: 31, F1: 1}
                    _ = v40
                case "foo bar":
                    if true {
                        v41 := GenS2{F0: "\n\t"}
                        v42 := GenS7{F0: fi37}
                        _ = v41
                        _ = v42
                    }
                    v1 -= (v1 % ((-2147483648) | 1))
                    var v43 uint64 = 4294967295
                    _ = v43
                default:
                    var v44 float32 = -2.25
                    _ = v44
                }
            }
            switch "hello" {
            case "a":
                v3 = (v3 >> ((40) & 63))
            default:
                v45 := []int{(fi37 >> ((uint(v32)) & 63)), fi37}
                _ = v45
            }
            var v46 string = "\n\t"
            _ = fi38
            _ = v46
        }
        if true {
            v47 := []int{-2147483648}
            v3 += 42
            _ = v47
        } else {
            var v48 uint64 = 2
            switch ("\n\t" + (("\n\t" + "\n\t") + "")) {
            case "foo bar":
                if ("a" <= "a") {
                    var v49 uint16 = 255
                    for fi50 := 0; fi50 < 13; fi50++ {
                        switch -1 {
                        case -161:
                            fmt.Println(v49)
                        default:
                            if true {
                                var v51 int64 = v3
                                v2[(v32 ^ (v32 & v32))] = (v3 != int64(v49))
                                _ = v51
                            }
                        }
                        switch fi50 {
                        case 1:
                            fmt.Println(fi37, v32)
                        default:
                            if true {
                                var v52 float32 = 0.1
                                _ = v52
                            }
                        }
                        v49 -= (v49 & 304)
                        var v53 int16 = int16(v3)
                        var v54 int = (fi50 >> ((42) & 63))
                        _ = fi50
                        _ = v53
                        _ = v54
                    }
                    _ = v49
                } else {
                    v55 := GenS3{F0: v32, F1: 4294967295, F2: 42}
                    _ = v55
                }
            case "\n\t":
                v56 := []int{}
                _ = v56
            default:
                switch (fi37 - fi37) {
                case -9223372036854775808:
                    var v57 uint16 = 0
                    _ = v57
                case 128:
                    switch ("a" + ("\n\t" + ("a" + "\n\t"))) {
                    case "hello":
                        v1 += (v1 - (-(v1)))
                    case "foo bar":
                        switch (-(fi37)) {
                        case 42:
                            v3 -= v3
                        }
                    default:
                        var v58 int8 = 127
                        _ = v58
                    }
                case 9223372036854775807:
                    for fi59 := 0; fi59 < 16; fi59++ {
                        v60 := GenS5{F0: (v1 + v1), F1: 127, F2: 42}
                        _ = fi59
                        _ = v60
                    }
                }
            }
            _ = v48
        }
        v3 = genFn0()
        _ = fi37
    }
    switch ("a" + ("hello" + "a")) {
    case "":
        if ((1e10 >= float32(3.14159)) && true) {
            for fi61 := 0; fi61 < 1; fi61++ {
                v62 := []int{fi61, -1, fi61}
                v3 = 9223372036854775807
                _ = fi61
                _ = v62
            }
            if (!true) {
                for fi63 := 0; fi63 < 7; fi63++ {
                    if (v3 < 1) {
                        for fi64 := 0; fi64 < 16; fi64++ {
                            v65 := []int{fi63, fi64, fi63}
                            fmt.Println(v1, v32, fi64)
                            v3 = v3
                            _ = fi64
                            _ = v65
                        }
                        for fi66 := 0; fi66 < 2; fi66++ {
                            v67 := GenS8{F0: 0, F1: 255, F2: "foo bar"}
                            v32 = (v32 | v32)
                            switch "a" {
                            case "hello":
                                switch (("" + "") + (("\n\t" + "") + "foo bar")) {
                                case "":
                                    fmt.Println(v3, v1)
                                case "\n\t":
                                    v1 = 1
                                }
                            case "":
                                v68 := []int{}
                                _ = v68
                            case "a":
                                v69 := GenS9{F0: (0 >= 4294967295), F1: (v1 << ((478) & 63)), F2: 1}
                                _ = v69
                            }
                            _ = fi66
                            _ = v67
                        }
                        var v70 int64 = 42
                        _ = v70
                    } else {
                        v71 := GenS10{F0: true, F1: (v32 & v32)}
                        for fi72 := 0; fi72 < 2; fi72++ {
                            var v73 bool = true
                            _ = fi72
                            _ = v73
                        }
                        var v74 uint32 = (v32 - (v32 ^ (v32 / ((2147483648) | 1))))
                        _ = v71
                        _ = v74
                    }
                    switch ("" + ("" + (((("a" + "") + ("" + ("hello" + "a"))) + "") + "foo bar"))) {
                    case "a":
                        v3 -= (v3 & genFn0())
                    case "foo bar":
                        var v75 float32 = 1.5
                        _ = v75
                    case "hello":
                        var v76 int8 = 28
                        _ = v76
                    default:
                        var v77 int16 = -32768
                        _ = v77
                    }
                    _ = fi63
                }
                fmt.Println(v1)
            }
        }
    case "\n\t":
        fmt.Println(v32, v1)
    }
    if ((0.0 - (-2.25 - (float32(-2.25) / 1.5))) < 0.1) {
        var v78 uint32 = 1
        v79 := GenS3{F0: (v78 & v78), F1: 0, F2: 621}
        var v80 float32 = 1e10
        _ = v78
        _ = v79
        _ = v80
    } else {
        v32 += (v32 - (v32 << ((65535) & 63)))
        v81 := GenS2{F0: "hello"}
        _ = v81
    }
    for fi82 := 0; fi82 < 3; fi82++ {
        var v83 uint32 = (v32 % (((v32 / ((v32) | 1))) | 1))
        var v84 int32 = v1
        _ = fi82
        _ = v83
        _ = v84
    }
    for fi85 := 0; fi85 < 20; fi85++ {
        v86 := GenS4{F0: 127, F1: 0}
        _ = fi85
        _ = v86
    }
    for fi87 := 0; fi87 < 12; fi87++ {
        for fi88 := 0; fi88 < 19; fi88++ {
            v3 -= (v3 << ((255) & 63))
            _ = fi88
        }
        _ = fi87
    }
    if (233 > 4294967295) {
        if (false && true) {
            v3 += 0
            var v89 int32 = v1
            for fi90 := 0; fi90 < 7; fi90++ {
                v91 := GenS11{F0: 82, F1: (^(v3)), F2: (^(v32))}
                var v92 float32 = float32(-2.25)
                v93 := []int{}
                _ = fi90
                _ = v91
                _ = v92
                _ = v93
            }
            _ = v89
        }
        switch ("a" + "hello") {
        case "\n\t":
            var v94 uint8 = 255
            _ = v94
        case "foo bar":
            _ = len(v2)
        case "":
            switch ((("foo bar" + ("hello" + "\n\t")) + "a") + "\n\t") {
            case "foo bar":
                var v95 float32 = float32(float64(1e10))
                _ = v95
            default:
                fmt.Println(v1, v32, v3)
            }
        default:
            switch "" {
            case "\n\t":
                v32 += v32
            default:
                v96 := GenS12{F0: -2147483648, F1: "\n\t"}
                _ = v96
            }
        }
    }
    var v97 int16 = -549
    v32 += v32
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v97
}
