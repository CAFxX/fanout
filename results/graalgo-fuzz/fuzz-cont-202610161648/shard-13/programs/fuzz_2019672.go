package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 bool; F1 int; F2 uint32 }
func genFn0(p0 uint32, p1 int64, p2 int32) string {
    fmt.Println(p0, p2)
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            p2 -= (p2 * p2)
            _ = fi6
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    return ("hello" + "a")
}
func genFn1() int {
    if (42 > 2) {
        switch (((("a" + "hello") + "\n\t") + "a") + (("\n\t" + ("foo bar" + ("\n\t" + ("foo bar" + "\n\t")))) + "\n\t")) {
        case "foo bar":
            for fi7 := 0; fi7 < 15; fi7++ {
                fmt.Println(fi7)
                _ = fi7
            }
        case "":
            var v8 bool = (3.14159 >= (1.5 / (1.5 * 3.14159)))
            _ = v8
        }
        v9 := GenS1{F0: 754}
        v10 := []int{0, 4294967295, -759, 9223372036854775807}
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 4294967295}
    var v12 uint = 654
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 float64 = 3.14159
    var v2 float32 = 3.14159
    var v3 uint32 = 811
    var v4 bool = (!true)
    v4 = (-234 < 1)
    fmt.Println(v1, v2, v1)
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int = fi14
            if v4 {
                v16 := []int{(fi14 * (fi14 - fi14))}
                if true {
                    if (!true) {
                        if (v15 >= (-(v15))) {
                            v3 = v3
                            switch "foo bar" {
                            case "hello":
                                if ((-2.25 / 1.0) > v1) {
                                    if v4 {
                                        _ = len(v16)
                                        var v17 uint = 2
                                        v3 = (v3 << (((v17 - v17)) & 63))
                                        _ = v17
                                    }
                                }
                            default:
                                switch (v15 | (v15 & -1)) {
                                case 42:
                                    v16 = append(v16, v15)
                                case -95:
                                    switch fi13 {
                                    case 851:
                                        if len(v16) > 0 {
                                            v16[0] = genFn1()
                                        }
                                    case 128:
                                        if v4 {
                                            if v4 {
                                                switch (^(fi13)) {
                                                case -1:
                                                    var v18 uint64 = 2
                                                    _ = v18
                                                case 0:
                                                    var v19 int32 = 1
                                                    _ = v19
                                                case 648:
                                                    fmt.Println(v3, v1, fi13)
                                                default:
                                                    v3 -= 2147483647
                                                }
                                            } else {
                                                switch 2147483647 {
                                                case -1:
                                                    if false {
                                                        var v20 int = v15
                                                        var v21 bool = v4
                                                        _ = v20
                                                        _ = v21
                                                    } else {
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                        var v22 int16 = 0
                                                        var v23 int64 = -1
                                                        v3 = (v3 | v3)
                                                        v3 += v3
                                                        _ = v22
                                                        _ = v23
                                                    }
                                                default:
                                                    v24 := GenS1{F0: 65535}
                                                    _ = v24
                                                }
                                                var v25 float64 = (float64(v2) - v1)
                                                switch (fi13 % (((^(v15))) | 1)) {
                                                case 4294967295:
                                                    v16 = append(v16, -349)
                                                case 9223372036854775807:
                                                    switch ("\n\t" + "") {
                                                    case "a":
                                                        v26 := GenS1{F0: 4294967295}
                                                        _ = v26
                                                    case "":
                                                        v4 = true
                                                    }
                                                default:
                                                    if v4 {
                                                        v27 := GenS2{F0: v3}
                                                        _ = v27
                                                    } else {
                                                        v28 := GenS3{F0: false, F1: -899}
                                                        _ = len(v16)
                                                        fmt.Println(v25, v25)
                                                        _ = v28
                                                    }
                                                }
                                                _ = v25
                                            }
                                        }
                                        var v29 bool = (!v4)
                                        v3 += 1
                                        _ = v29
                                    case -593:
                                        v4 = (125 == 128)
                                    default:
                                        v15 = (fi13 * (fi13 / ((genFn1()) | 1)))
                                    }
                                case -512:
                                    if ((32767 > 32767) && (float32(v1) <= 1e10)) {
                                        fmt.Println(v4, fi14, v2)
                                        if len(v16) > 0 {
                                            v16[0] = v15
                                        }
                                        v30 := GenS3{F0: (2147483647 >= (fi13 | (fi13 << ((0) & 63)))), F1: -1}
                                        _ = v30
                                    } else {
                                        v2 = (v2 * -2.25)
                                    }
                                }
                            }
                        }
                        v2 = (float32(1.5) + v2)
                    } else {
                        var v31 uint64 = 1
                        v32 := GenS3{F0: false, F1: -1}
                        var v33 uint8 = 151
                        v34 := GenS2{F0: v3}
                        _ = v31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                } else {
                    if v4 {
                        v1 += (1e10 / (((v1 * v1) / (((float64((-2.25 - (v2 * v2))) / v1) / v1) - (0.1 * v1))) + v1))
                        fmt.Println(fi14, v15)
                        var v35 uint8 = 128
                        _ = v35
                    }
                    fmt.Println(v3)
                    var v36 int32 = 1
                    fmt.Println(v36, v4)
                    var v37 int = (fi14 + -937)
                    _ = v36
                    _ = v37
                }
                if ("a" <= ("foo bar" + "a")) {
                    var v38 int32 = -1
                    var v39 string = ("hello" + "hello")
                    v39 = "a"
                    _ = v38
                    _ = v39
                }
                _ = v16
            }
            var v40 int = (-(fi13))
            _ = fi14
            _ = v15
            _ = v40
        }
        v41 := []int{127, (fi13 << ((864) & 63)), (fi13 + genFn1())}
        var v42 uint16 = 255
        _ = fi13
        _ = v41
        _ = v42
    }
    if (0 > 1000) {
        var v43 int8 = 0
        _ = v43
    } else {
        if true {
            v44 := GenS3{F0: true, F1: 947}
            if (v2 <= 0.0) {
                var v45 int8 = -88
                v46 := GenS4{F0: 9223372036854775807, F1: (0 > 2147483647)}
                _ = v45
                _ = v46
            }
            _ = v44
        }
        v47 := []int{}
        v2 -= (((0.1 * 0.0) * (v2 / ((v2 / 1e10) + v2))) * 1e10)
        _ = v47
    }
    switch (("hello" + "\n\t") + "foo bar") {
    case "":
        v4 = v4
    case "a":
        for fi48 := 0; fi48 < 18; fi48++ {
            v49 := []int{}
            for fi50 := 0; fi50 < 19; fi50++ {
                if (127 >= 2) {
                    v51 := GenS5{F0: (!(!true)), F1: genFn1(), F2: (v3 >> ((18446744073709551615) & 63))}
                    var v52 int32 = 42
                    _ = v51
                    _ = v52
                }
                var v53 int16 = -1
                _ = fi50
                _ = v53
            }
            v49 = append(v49, 677)
            _ = fi48
            _ = v49
        }
    }
    if (18446744073709551615 == 42) {
        for fi54 := 0; fi54 < 6; fi54++ {
            for fi55 := 0; fi55 < 3; fi55++ {
                v56 := []int{int(v3), (fi54 & fi54), fi55}
                var v57 uint8 = uint8(v3)
                _ = fi55
                _ = v56
                _ = v57
            }
            _ = fi54
        }
        for fi58 := 0; fi58 < 4; fi58++ {
            var v59 uint16 = 2
            v3 += v3
            _ = fi58
            _ = v59
        }
    }
    var v60 uint = 0
    v61 := GenS4{F0: -1, F1: true}
    v62 := GenS3{F0: (v60 != (v60 & (^(v60)))), F1: 42}
    for fi63 := 0; fi63 < 19; fi63++ {
        v64 := GenS5{F0: false, F1: (fi63 | fi63), F2: v3}
        switch (fi63 ^ int(v60)) {
        case 84:
            var v65 int = -2147483648
            _ = v65
        case 115:
            var v66 int8 = 127
            _ = v66
        default:
            v67 := []int{32768}
            _ = v67
        }
        v68 := []int{(fi63 % ((int(v3)) | 1)), fi63}
        _ = fi63
        _ = v64
        _ = v68
    }
    var v69 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v61
    _ = v62
    _ = v69
}
