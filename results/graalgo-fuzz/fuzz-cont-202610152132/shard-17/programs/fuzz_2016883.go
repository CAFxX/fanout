package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 bool }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 bool; F1 uint16; F2 int32 }
type GenS5 struct { F0 int64; F1 int; F2 bool }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint16) int64 {
    switch "a" {
    case "foo bar":
        for fi4 := 0; fi4 < 20; fi4++ {
            var v5 int = (fi4 & (fi4 >> ((1) & 63)))
            p1 = 42
            v6 := GenS1{F0: 1, F1: true}
            _ = fi4
            _ = v5
            _ = v6
        }
    case "hello":
        v7 := []int{188}
        _ = v7
    case "":
        if true {
            v8 := []int{42}
            for fi9 := 0; fi9 < 12; fi9++ {
                v8 = append(v8, 42)
                switch (fi9 & fi9) {
                case -2147483648:
                    var v10 int8 = -57
                    _ = v10
                case 9223372036854775807:
                    var v11 uint = 0
                    _ = v11
                default:
                    _ = len(v8)
                }
                _ = fi9
            }
            _ = v8
        }
    }
    switch 314 {
    case -231:
        v12 := []int{1}
        _ = v12
    case 1:
        if (!(!true)) {
            p0 += p0
        }
    case 0:
        p0 = 65535
    default:
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1() int8 {
    var v13 uint32 = 65535
    var v14 uint16 = 2
    v15 := []int{128}
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(fi16, v13, v13)
        v14 += v14
        var v17 uint64 = 771
        _ = fi16
        _ = v17
    }
    switch ("foo bar" + ("foo bar" + ("\n\t" + "foo bar"))) {
    case "\n\t":
        if ((^(v13)) != v13) {
            var v18 uint16 = (v14 + 0)
            for fi19 := 0; fi19 < 15; fi19++ {
                v14 = v18
                fmt.Println(v18, v14, v13)
                var v20 float32 = float32(0.0)
                _ = fi19
                _ = v20
            }
            v13 -= 65535
            _ = v18
        } else {
            _ = len(v15)
            if false {
                if (((false || ((("" > ("\n\t" + "foo bar")) && true) || (540 != 0))) && (-342 == 42)) && false) {
                    var v21 int32 = 1
                    if len(v15) > 0 {
                        v15[0] = 2147483647
                    }
                    _ = v21
                } else {
                    var v22 int64 = 0
                    _ = v22
                }
            }
        }
    case "":
        switch "a" {
        case "foo bar":
            var v23 float64 = ((1e10 - 0.1) + 0.1)
            _ = v23
        }
    case "hello":
        v24 := GenS1{F0: int32(v13), F1: (-260 == 128)}
        _ = v24
    default:
        var v25 string = "\n\t"
        _ = v25
    }
    _ = v13
    _ = v14
    _ = v15
    return -36
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]string)
    var v3 int32 = -41
    delete(v2, "\n\t")
    switch 323 {
    case 563:
        v26 := GenS1{F0: -1, F1: (0 > 1)}
        _ = v26
    case -1:
        var v27 string = "a"
        _ = v27
    default:
        v3 = (v3 | 2147483647)
    }
    if false {
        switch "hello" {
        case "":
            v28 := []int{386, 127, -1}
            _ = v28
        case "\n\t":
            v29 := []int{42, 32767}
            _ = v29
        case "hello":
            var v30 uint8 = (v1 / ((v1) | 1))
            _ = v30
        }
        v31 := GenS2{F0: -225, F1: (v1 - (v1 ^ 1)), F2: (!true)}
        v3 = -831
        _ = v31
    } else {
        var v32 int16 = 1
        switch 2147483647 {
        case -9223372036854775808:
            if (true && (!true)) {
                if true {
                    v2["a"] = "\n\t"
                } else {
                    v33 := GenS3{F0: v32, F1: "a"}
                    _ = v33
                }
                var v34 int = 0
                if (!((18446744073709551615 >= 434) && false)) {
                    v34 = (v34 ^ 4294967295)
                    fmt.Println(v1, v3, v1)
                    switch "a" {
                    case "foo bar":
                        v35 := GenS4{F0: (((-(v3)) >= v3) && false), F1: 32767, F2: v3}
                        _ = v35
                    case "hello":
                        v32 = (v32 | (v32 >> ((uint(v1)) & 63)))
                    default:
                        v36 := GenS5{F0: 42, F1: (v34 * (v34 ^ v34)), F2: false}
                        _ = v36
                    }
                } else {
                    fmt.Println(v3, v32)
                    var v37 string = "a"
                    _ = v37
                }
                _ = v34
            }
        case 0:
            v38 := v2["foo bar"]
            _ = v38
        }
        var v39 float64 = 3.14159
        _ = v32
        _ = v39
    }
    v1 -= 128
    var v40 int16 = 0
    for fi41 := 0; fi41 < 2; fi41++ {
        var v42 int8 = -13
        _ = fi41
        _ = v42
    }
    fmt.Println(v3, v40, v1)
    if (!false) {
        switch 127 {
        case 0:
            if (!(623 == 246)) {
                v1 -= 2
                for fi43 := 0; fi43 < 20; fi43++ {
                    var v44 uint16 = 295
                    _ = fi43
                    _ = v44
                }
                fmt.Println(v40)
            } else {
                delete(v2, ("hello" + "hello"))
                var v45 float32 = -2.25
                _ = v45
            }
        case 32768:
            fmt.Println(v3)
        }
        for fi46 := 0; fi46 < 8; fi46++ {
            switch -438 {
            case -1:
                fmt.Println(fi46, v3, v1)
            case 2147483647:
                if false {
                    var v47 uint8 = (^(v1))
                    var v48 bool = (uint64(v47) < 400)
                    _ = v47
                    _ = v48
                } else {
                    v3 = (v3 << ((65535) & 63))
                    if false {
                        switch "" {
                        case "":
                            var v49 uint8 = v1
                            _ = v49
                        case "a":
                            if (316 < 0) {
                                for fi50 := 0; fi50 < 18; fi50++ {
                                    if (false && (genFn0(65535, 1) > 42)) {
                                        var v51 int = (fi46 ^ fi46)
                                        var v52 int8 = -56
                                        _ = v51
                                        _ = v52
                                    } else {
                                        _ = len(v2)
                                        var v53 uint16 = 1
                                        v3 = v3
                                        _ = v53
                                    }
                                    _ = fi50
                                }
                                if false {
                                    switch (fi46 * 4294967295) {
                                    case 1:
                                        delete(v2, "hello")
                                    case 9223372036854775807:
                                        fmt.Println(v3, v40)
                                    case 307:
                                        v54 := []int{42, (fi46 >> ((42) & 63))}
                                        _ = v54
                                    }
                                    v1 = v1
                                    for fi55 := 0; fi55 < 12; fi55++ {
                                        if (false && false) {
                                            var v56 bool = true
                                            _ = v56
                                        } else {
                                            fmt.Println(fi46, v3, v1)
                                            v57 := GenS1{F0: v3, F1: (true && true)}
                                            v58 := v2[("foo bar" + (("\n\t" + "") + "\n\t"))]
                                            _ = v57
                                            _ = v58
                                        }
                                        var v59 bool = (18446744073709551615 == 597)
                                        _ = fi55
                                        _ = v59
                                    }
                                }
                                for fi60 := 0; fi60 < 19; fi60++ {
                                    v40 = (^(v40))
                                    v61 := GenS5{F0: 878, F1: fi46, F2: false}
                                    _ = fi60
                                    _ = v61
                                }
                            } else {
                                fmt.Println(v40)
                            }
                        case "hello":
                            if (((fi46 | fi46) < (^(fi46))) && (true && ((("a" + "\n\t") + ("a" + "hello")) < "\n\t"))) {
                                v1 = v1
                            } else {
                                v62 := GenS6{F0: 86, F1: 907}
                                _ = v62
                            }
                        default:
                            var v63 uint16 = 255
                            _ = v63
                        }
                        if (!false) {
                            delete(v2, "hello")
                            v40 = (v40 + (v40 + v40))
                            v1 = v1
                        } else {
                            var v64 int64 = 1
                            _ = v64
                        }
                    } else {
                        v65 := GenS7{F0: 255, F1: 2}
                        _ = v65
                    }
                    v1 = 128
                }
            case 32768:
                for fi66 := 0; fi66 < 18; fi66++ {
                    var v67 int32 = (-(v3))
                    _ = fi66
                    _ = v67
                }
            default:
                v68 := []int{(fi46 & (fi46 << ((1) & 63))), fi46, fi46, 42}
                _ = v68
            }
            fmt.Println(fi46)
            v69 := []int{}
            _ = fi46
            _ = v69
        }
    }
    for fi70 := 0; fi70 < 10; fi70++ {
        fmt.Println(v3, fi70)
        v71 := GenS5{F0: 198, F1: -2147483648, F2: true}
        _ = fi70
        _ = v71
    }
    v1 = 96
    if true {
        var v72 int16 = (v40 % ((v40) | 1))
        _ = v72
    }
    if (!true) {
        v1 -= 255
        switch -706 {
        case 128:
            var v73 int32 = (^(v3))
            _ = v73
        case -2147483648:
            var v74 int32 = (v3 / ((v3) | 1))
            _ = v74
        case 42:
            var v75 uint = 4294967295
            _ = v75
        }
    }
    v76 := GenS8{F0: 172, F1: 255}
    if false {
        if false {
            v77 := []int{127}
            v1 += v1
            _ = v77
        } else {
            v2["foo bar"] = "hello"
            for fi78 := 0; fi78 < 11; fi78++ {
                var v79 float32 = 1.5
                _ = fi78
                _ = v79
            }
            for fi80 := 0; fi80 < 19; fi80++ {
                v81 := GenS5{F0: -1, F1: -3, F2: (false || true)}
                v1 -= v1
                var v82 uint64 = 42
                _ = fi80
                _ = v81
                _ = v82
            }
        }
        v83 := GenS1{F0: v3, F1: (!true)}
        v84 := GenS9{F0: genFn1(), F1: 42}
        _ = v83
        _ = v84
    }
    var v85 uint8 = v1
    v86 := []int{691, -1}
    for fi87 := 0; fi87 < 12; fi87++ {
        v88 := GenS10{F0: 1}
        var v89 int8 = 73
        if true {
            if (false && true) {
                v40 = 42
                if (((float32(3.14159) / 1.0) != 1e10) || (false || true)) {
                    fmt.Println(v85, v3, v85)
                }
            }
            v1 += v1
        }
        _ = fi87
        _ = v88
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v85:", v85)
    fmt.Println("slicelen v86:", len(v86), v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v76
    _ = v85
    _ = v86
}
