package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 string; F2 int64 }
type GenS9 struct { F0 int32; F1 string }
func genFn0(p0 int16, p1 string) string {
    var v5 int16 = (p0 ^ 1)
    var v6 int8 = -116
    var v7 float32 = 0.0
    if true {
        var v8 int32 = -2147483648
        _ = v8
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            v7 = (v7 - v7)
            _ = fi9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 string = p1
            if true {
                var v12 bool = true
                switch fi10 {
                case -1:
                    if (144 > 128) {
                        v11 = "a"
                        v13 := GenS1{F0: -1, F1: 42, F2: 10}
                        var v14 int32 = 429
                        _ = v13
                        _ = v14
                    }
                case 32768:
                    v15 := []int{(-(fi10)), fi10, (fi10 << ((42) & 63))}
                    _ = v15
                case 1:
                    v16 := GenS1{F0: 42, F1: -240, F2: 2}
                    _ = v16
                default:
                    v6 = (^(v6))
                }
                _ = v12
            } else {
                for fi17 := 0; fi17 < 13; fi17++ {
                    v18 := GenS1{F0: 42, F1: 11, F2: 1}
                    _ = fi17
                    _ = v18
                }
                v6 = 42
            }
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 float64 = (0.0 - 1.5)
    var v2 bool = (42 > 397)
    var v3 int8 = -1
    var v4 int32 = 2147483647
    v2 = (469 > 42)
    if v2 {
        if true {
            if ((v4 >= (v4 | 42)) || (255 > 182)) {
                if ((-2.25 / (((1.5 * float32(-2.25)) / 0.1) * (1e10 - float32(v1)))) == -2.25) {
                    v19 := []int{-529}
                    for fi20 := 0; fi20 < 11; fi20++ {
                        if len(v19) > 0 {
                            v19[0] = fi20
                        }
                        if false {
                            v21 := GenS2{F0: 446}
                            for fi22 := 0; fi22 < 19; fi22++ {
                                v23 := GenS2{F0: 750}
                                if v2 {
                                    v4 = 279
                                }
                                _ = len(v19)
                                v24 := GenS2{F0: 18446744073709551615}
                                v3 -= v3
                                v25 := GenS3{F0: 0, F1: "a"}
                                _ = fi22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            var v26 uint8 = 128
                            _ = v21
                            _ = v26
                        } else {
                            if len(v19) > 0 {
                                v19[0] = fi20
                            }
                            v4 = v4
                            v1 += 1.5
                        }
                        v2 = v2
                        _ = fi20
                    }
                    var v27 int32 = v4
                    _ = v19
                    _ = v27
                }
            }
            v28 := GenS4{F0: 0, F1: "foo bar"}
            _ = v28
        } else {
            switch ("" + ("a" + ("hello" + "foo bar"))) {
            case "\n\t":
                v29 := GenS3{F0: 1, F1: "hello"}
                _ = v29
            case "":
                v30 := []int{874, 128}
                _ = v30
            case "foo bar":
                if ("foo bar" > "foo bar") {
                    v31 := GenS3{F0: 1, F1: "\n\t"}
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v33 := GenS3{F0: 2, F1: "\n\t"}
                        _ = fi32
                        _ = v33
                    }
                    _ = v31
                }
            }
            v34 := []int{}
            _ = v34
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            for fi36 := 0; fi36 < 10; fi36++ {
                if v2 {
                    v37 := GenS1{F0: v4, F1: 1, F2: 573}
                    v4 -= v4
                    _ = v37
                }
                if (1 == 2) {
                    v38 := GenS4{F0: 54, F1: ""}
                    _ = v38
                }
                var v39 uint8 = 1
                _ = fi36
                _ = v39
            }
            _ = fi35
        }
    } else {
        v40 := GenS4{F0: 18446744073709551615, F1: ("\n\t" + "foo bar")}
        for fi41 := 0; fi41 < 16; fi41++ {
            if (v4 > 2147483647) {
                for fi42 := 0; fi42 < 20; fi42++ {
                    v43 := []int{(fi41 & (fi41 / ((0) | 1))), fi42}
                    if true {
                        v1 = ((3.14159 + (3.14159 * v1)) + 1.5)
                        v4 = v4
                        v2 = true
                    } else {
                        var v44 int8 = 127
                        _ = v44
                    }
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                    _ = fi42
                    _ = v43
                }
                v4 = (v4 & -22)
                v45 := GenS3{F0: 65535, F1: ("hello" + "\n\t")}
                _ = v45
            }
            var v46 uint16 = 42
            _ = fi41
            _ = v46
        }
        if v2 {
            switch ("hello" + "") {
            case "":
                v47 := GenS5{F0: 82}
                _ = v47
            case "a":
                v1 -= (0.1 - v1)
            default:
                var v48 uint16 = 2
                _ = v48
            }
            if false {
                v49 := GenS5{F0: 2147483647}
                v50 := []int{127, 127, int(v4)}
                _ = v49
                _ = v50
            }
            if (("foo bar" != "") || (((0.1 - (v1 + v1)) / 1e10) >= v1)) {
                v51 := GenS5{F0: 2}
                if false {
                    if (0 == 157) {
                        var v52 bool = (v2 && v2)
                        v53 := []int{2147483647, -1, 356}
                        _ = v52
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 8; fi54++ {
                        v55 := GenS6{F0: 127}
                        v56 := GenS7{F0: 321}
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                }
                v1 = float64(0.0)
                _ = v51
            }
        } else {
            v4 = v4
            if (216 < int(v4)) {
                v57 := GenS8{F0: 42, F1: "", F2: 9223372036854775807}
                v58 := []int{}
                _ = len(v58)
                _ = v57
                _ = v58
            }
        }
        _ = v40
    }
    v4 -= int32(v3)
    v1 = 3.14159
    v2 = (32 != 194)
    if ((141 >= 27) && v2) {
        if v2 {
            var v59 uint16 = 30
            var v60 int64 = 1
            _ = v59
            _ = v60
        } else {
            var v61 int16 = 1
            var v62 int32 = 2147483647
            _ = v61
            _ = v62
        }
        v4 += (v4 << ((352) & 63))
        v1 += v1
    }
    switch (((("hello" + (("" + "hello") + "a")) + "hello") + "") + ("foo bar" + "\n\t")) {
    case "hello":
        for fi63 := 0; fi63 < 7; fi63++ {
            for fi64 := 0; fi64 < 10; fi64++ {
                var v65 float64 = -2.25
                _ = fi64
                _ = v65
            }
            v3 = v3
            for fi66 := 0; fi66 < 10; fi66++ {
                v3 -= v3
                var v67 string = ""
                _ = fi66
                _ = v67
            }
            _ = fi63
        }
    case "":
        v1 = (v1 - float64(0.1))
    }
    v3 = v3
    if v2 {
        v68 := []int{626, -2147483648, 734, -9223372036854775808}
        _ = v68
    } else {
        v69 := []int{965, -801, -331, 1}
        if ("\n\t" < "foo bar") {
            var v70 uint64 = 109
            var v71 int = 0
            _ = v70
            _ = v71
        } else {
            var v72 int16 = 42
            var v73 uint64 = 1
            v69 = append(v69, 32767)
            _ = v72
            _ = v73
        }
        _ = v69
    }
    for fi74 := 0; fi74 < 8; fi74++ {
        if true {
            var v75 uint16 = 149
            switch "" {
            case "":
                v76 := GenS9{F0: v4, F1: ""}
                _ = v76
            case "a":
                switch (("" + "a") + "foo bar") {
                case "\n\t":
                    var v77 uint16 = (v75 & (v75 << ((0) & 63)))
                    _ = v77
                case "a":
                    var v78 uint32 = 65535
                    _ = v78
                }
            case "hello":
                v75 -= 32767
            }
            v75 -= 205
            _ = v75
        } else {
            if v2 {
                fmt.Println(v4)
                var v79 int8 = 1
                _ = v79
            }
        }
        v80 := GenS7{F0: 754}
        var v81 bool = (42 > 0)
        _ = fi74
        _ = v80
        _ = v81
    }
    fmt.Println(v3)
    v82 := []int{32767, 504}
    var v83 float32 = -2.25
    for fi84 := 0; fi84 < 20; fi84++ {
        var v85 uint16 = 0
        _ = fi84
        _ = v85
    }
    v2 = (!(3.14159 > v83))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v82:", len(v82), v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v82
    _ = v83
}
