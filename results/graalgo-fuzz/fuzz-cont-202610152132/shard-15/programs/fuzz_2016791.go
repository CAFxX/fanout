package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
type GenS3 struct { F0 int; F1 uint64; F2 int16 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 int; F1 uint16 }
type GenS10 struct { F0 int64; F1 uint; F2 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int32; F1 uint8 }
func genFn0(p0 int32, p1 int) uint {
    p0 += (p0 - p0)
    v4 := []int{(p1 % ((605) | 1))}
    p1 = p1
    _ = p0
    _ = p1
    _ = v4
    return 4294967295
}
func genFn1() int16 {
    v5 := GenS1{F0: ""}
    var v6 int64 = 9223372036854775807
    var v7 bool = (0 != -32768)
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 float32 = (1e10 * 1.5)
    v2 := []int{9223372036854775807, 0, -491, 42}
    v3 := make(map[string]int)
    delete(v3, "foo bar")
    if (!true) {
        v8 := GenS2{F0: genFn1(), F1: -2147483648, F2: -436}
        v9 := GenS3{F0: -9223372036854775808, F1: 18446744073709551615, F2: -799}
        for fi10 := 0; fi10 < 4; fi10++ {
            v1 = 0.0
            v11 := GenS4{F0: -82, F1: 2147483647}
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    } else {
        if true {
            for fi12 := 0; fi12 < 10; fi12++ {
                for fi13 := 0; fi13 < 18; fi13++ {
                    if false {
                        v14 := v3["\n\t"]
                        v3[""] = fi13
                        _ = v14
                    } else {
                        switch ("a" + "hello") {
                        case "a":
                            fmt.Println(v1, fi13, fi12)
                        }
                        var v15 uint8 = 2
                        if false {
                            var v16 uint32 = uint32(v15)
                            var v17 uint = 2
                            if len(v2) > 0 {
                                v2[0] = fi13
                            }
                            switch "a" {
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = fi13
                                }
                            case "":
                                v18 := GenS5{F0: (v15 ^ uint8(v17))}
                                _ = v18
                            default:
                                if false {
                                    var v19 int64 = -1
                                    var v20 uint64 = 42
                                    switch -2147483648 {
                                    case -9223372036854775808:
                                        if true {
                                            _ = len(v3)
                                            var v21 uint16 = 0
                                            fmt.Println(v20, v21, fi12)
                                            _ = v21
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                var v22 int32 = -2147483648
                                                _ = v22
                                            }
                                        }
                                    case -2147483648:
                                        if true {
                                            if ((genFn1() > 42) && (true || true)) {
                                                var v23 int16 = 0
                                                var v24 int16 = (v23 >> ((18446744073709551615) & 63))
                                                v3["\n\t"] = fi13
                                                _ = v23
                                                _ = v24
                                            } else {
                                                var v25 uint64 = v20
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if (!false) {
                                                    v26 := GenS6{F0: -100}
                                                    if true {
                                                        var v27 int = (fi12 & int(v19))
                                                        if (v1 == float32(1.5)) {
                                                            v15 -= 127
                                                            v25 = 18446744073709551615
                                                        }
                                                        switch (-(fi12)) {
                                                        case 9223372036854775807:
                                                            switch 128 {
                                                            case 127:
                                                                v17 -= v17
                                                            }
                                                        default:
                                                            v28 := GenS5{F0: 42}
                                                            _ = v28
                                                        }
                                                        var v29 uint8 = v15
                                                        _ = v27
                                                        _ = v29
                                                    }
                                                    _ = v26
                                                } else {
                                                    if true {
                                                        var v30 bool = false
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v31 uint8 = v15
                                                            _ = v31
                                                        case "":
                                                            var v32 string = ("a" + ("hello" + "\n\t"))
                                                            _ = v32
                                                        default:
                                                            v3["hello"] = (-(fi13))
                                                        }
                                                        _ = v30
                                                    } else {
                                                        v33 := GenS5{F0: v15}
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if true {
                                                                var v34 int32 = 2147483647
                                                                _ = v34
                                                            } else {
                                                                v1 -= v1
                                                                if true {
                                                                    var v35 string = (("hello" + "") + "foo bar")
                                                                    _ = v35
                                                                }
                                                            }
                                                        }
                                                        _ = v33
                                                    }
                                                    if (!(!true)) {
                                                        v36 := GenS3{F0: 32768, F1: 1, F2: 1}
                                                        _ = v36
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = fi12
                                                        }
                                                        var v37 uint = v17
                                                        var v38 uint32 = (v16 % ((v16) | 1))
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    v16 = uint32(v15)
                                                }
                                                _ = v25
                                            }
                                            var v39 int32 = 624
                                            _ = v39
                                        } else {
                                            fmt.Println(fi12, v15)
                                        }
                                    case -191:
                                        var v40 uint8 = (v15 | (v15 | (v15 * v15)))
                                        _ = v40
                                    default:
                                        v41 := v3[""]
                                        _ = v41
                                    }
                                    if (!(!(!true))) {
                                        if (!(25 == int16(v16))) {
                                            if (-128 >= -1) {
                                                v42 := v3["a"]
                                                switch (fi12 % ((v42) | 1)) {
                                                case -2147483648:
                                                    if true {
                                                        if true {
                                                            _ = len(v3)
                                                            v42 -= fi13
                                                            fmt.Println(fi13)
                                                            fmt.Println(v16, v1, v19)
                                                        } else {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                if (1e10 == v1) {
                                                                    var v43 int64 = (^(v19))
                                                                    var v44 uint32 = (v16 * (v16 >> (((v17 - uint(v15))) & 63)))
                                                                    fmt.Println(v42, fi12, v20)
                                                                    _ = v43
                                                                    _ = v44
                                                                } else {
                                                                    fmt.Println(v42)
                                                                    if (float32(1.5) < v1) {
                                                                        v42 += fi12
                                                                        v1 += 0.0
                                                                        v45 := v3[("" + "")]
                                                                        if true {
                                                                            v20 = (v20 | 18446744073709551615)
                                                                        } else {
                                                                            v46 := GenS3{F0: -9223372036854775808, F1: v20, F2: 42}
                                                                            if (v20 != 474) {
                                                                                v2 = append(v2, (fi12 / ((fi12) | 1)))
                                                                                var v47 float64 = 1e10
                                                                                _ = v47
                                                                            }
                                                                            _ = v46
                                                                        }
                                                                        _ = v45
                                                                    }
                                                                    v15 -= v15
                                                                }
                                                            case "":
                                                                if (209 < 295) {
                                                                    v17 -= (v17 >> ((uint(v20)) & 63))
                                                                }
                                                            case "hello":
                                                                var v48 uint64 = (v20 & 42)
                                                                _ = v48
                                                            }
                                                            v20 = (v20 | (v20 / ((v20) | 1)))
                                                        }
                                                        v17 += uint(v15)
                                                        _ = len(v2)
                                                    } else {
                                                        switch (v42 | (v42 | v42)) {
                                                        case 601:
                                                            var v49 float32 = v1
                                                            _ = v49
                                                        case -401:
                                                            var v50 int = int(v19)
                                                            _ = v50
                                                        case 127:
                                                            var v51 uint32 = v16
                                                            _ = v51
                                                        default:
                                                            if true {
                                                                if (!((true || (0 == (v16 & v16))) || false)) {
                                                                    v1 = (float32(((float64(v1) - float64((v1 + (v1 * (0.0 / float32((1e10 * 1.5))))))) + 1e10)) * 0.0)
                                                                    v15 -= 2
                                                                    fmt.Println(fi13)
                                                                } else {
                                                                    v19 -= -156
                                                                }
                                                            }
                                                        }
                                                        _ = len(v3)
                                                    }
                                                default:
                                                    v15 += 135
                                                }
                                                _ = v42
                                            } else {
                                                var v52 float64 = -2.25
                                                switch fi13 {
                                                case -2147483648:
                                                    var v53 string = ("a" + "a")
                                                    _ = v53
                                                case -968:
                                                    switch 2147483647 {
                                                    case 1:
                                                        fmt.Println(v52, v20)
                                                    case 343:
                                                        v16 += 888
                                                    case -1:
                                                        fmt.Println(v1, fi13)
                                                        v54 := GenS7{F0: 3, F1: (-(v19))}
                                                        _ = v54
                                                    }
                                                }
                                                _ = v52
                                            }
                                        }
                                        if false {
                                            v55 := GenS8{F0: int8(v20), F1: 0}
                                            _ = v55
                                        } else {
                                            var v56 int = (fi13 << (((v17 / ((genFn0(-561, (fi12 / ((32768) | 1)))) | 1))) & 63))
                                            if (9223372036854775807 < 128) {
                                                v20 -= v20
                                                v15 += 127
                                            }
                                            v57 := GenS9{F0: v56, F1: 2}
                                            _ = v56
                                            _ = v57
                                        }
                                        if false {
                                            if false {
                                                if false {
                                                    delete(v3, "foo bar")
                                                    if ((v1 + v1) >= float32(float64(v1))) {
                                                        switch "hello" {
                                                        case "a":
                                                            v58 := GenS10{F0: v19, F1: v17, F2: (v20 << (((^(v17))) & 63))}
                                                            _ = v58
                                                        default:
                                                            v3["foo bar"] = (fi13 * (-(fi13)))
                                                        }
                                                        var v59 string = ((("a" + ("" + "foo bar")) + ("hello" + "hello")) + "")
                                                        var v60 bool = (v59 < v59)
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                    v61 := GenS11{F0: v17}
                                                    _ = v61
                                                }
                                                switch fi13 {
                                                case 32767:
                                                    var v62 uint8 = (^(v15))
                                                    _ = v62
                                                }
                                                if false {
                                                    v16 += (^(v16))
                                                } else {
                                                    v1 -= (float32((0.1 / 1e10)) * (((float32(3.14159) - (0.1 / float32((float64(float32((0.1 / (((1.5 / 3.14159) - (((1.5 * 3.14159) / 1.0) + ((3.14159 + 0.0) / (((1e10 / 0.1) * 0.0) - float64((float32(3.14159) - (1e10 / float32(-2.25)))))))) - ((0.0 / float64(1.5)) - (3.14159 - ((1e10 - -2.25) / (1.5 + 1.5)))))))) + 3.14159)))) / (1e10 * v1)) - float32(-2.25)))
                                                    delete(v3, "foo bar")
                                                }
                                            }
                                            var v63 float32 = v1
                                            _ = v63
                                        }
                                    }
                                    var v64 float32 = 0.1
                                    v16 -= (v16 - v16)
                                    _ = v19
                                    _ = v20
                                    _ = v64
                                } else {
                                    fmt.Println(fi12)
                                    v65 := v3[("foo bar" + (("a" + ("hello" + "foo bar")) + (("a" + "foo bar") + ("foo bar" + "hello"))))]
                                    _ = v65
                                }
                            }
                            if true {
                                v66 := GenS11{F0: genFn0(-335, fi13)}
                                fmt.Println(v16, v15)
                                fmt.Println(fi13)
                                _ = v66
                            } else {
                                switch fi13 {
                                case 32768:
                                    v67 := GenS12{F0: 2147483647, F1: (^(v15))}
                                    _ = v67
                                default:
                                    v15 -= (v15 + 0)
                                }
                            }
                            _ = v16
                            _ = v17
                        }
                        _ = v15
                    }
                    v1 = v1
                    v1 = v1
                    _ = fi13
                }
                _ = fi12
            }
        }
    }
    if false {
        switch 335 {
        case -1:
            var v68 uint16 = 255
            _ = v68
        case 9223372036854775807:
            v69 := GenS8{F0: 114, F1: 544}
            _ = v69
        }
        var v70 bool = true
        _ = v70
    }
    _ = len(v2)
    for fi71 := 0; fi71 < 11; fi71++ {
        v1 += ((v1 / v1) * v1)
        _ = fi71
    }
    v1 += v1
    v1 = ((v1 + v1) + float32(float64(-2.25)))
    v72 := GenS4{F0: 127, F1: 42}
    v73 := v3["a"]
    var v74 uint = 65535
    if true {
        if false {
            var v75 bool = true
            var v76 float32 = (v1 / v1)
            _ = v75
            _ = v76
        } else {
            fmt.Println(v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v72
    _ = v73
    _ = v74
}
