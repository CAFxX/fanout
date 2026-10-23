package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 int16; F2 string }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 int32 }
type GenS11 struct { F0 int; F1 uint32 }
func genFn0(p0 int) uint {
    p0 = p0
    var v6 int8 = 127
    var v7 string = ("\n\t" + "a")
    _ = p0
    _ = v6
    _ = v7
    return 161
}
func genFn1(p0 uint) uint32 {
    v8 := []int{-1, -9223372036854775808, 128, -353}
    if (550 != 0) {
        v9 := GenS1{F0: -293, F1: "hello"}
        for fi10 := 0; fi10 < 12; fi10++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            p0 -= p0
            _ = fi10
        }
        _ = v9
    }
    p0 += 56
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: fi11, F1: ((("a" + ("hello" + (((((("foo bar" + "foo bar") + "") + "\n\t") + "hello") + "hello") + ((((("foo bar" + "\n\t") + ("hello" + ((("\n\t" + "\n\t") + "\n\t") + "hello"))) + "foo bar") + "hello") + "foo bar")))) + "hello") + "foo bar")}
        v13 := GenS1{F0: fi11, F1: (("foo bar" + "\n\t") + "hello")}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = len(v8)
    _ = p0
    _ = v8
    return 1
}

func main() {
    var v1 uint32 = 646
    var v2 int32 = -1
    v3 := make(map[string]int)
    var v4 int = -650
    v5 := make(map[uint32]string)
    if false {
        v14 := []int{v4, (v4 + 4294967295)}
        var v15 int64 = 1
        _ = v14
        _ = v15
    } else {
        v5[(v1 + v1)] = ("a" + "a")
    }
    var v16 uint8 = 127
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS2{F0: -1, F1: 255, F2: v16}
        _ = fi17
        _ = v18
    }
    var v19 float64 = ((0.1 - -2.25) - (0.0 + 3.14159))
    v4 += v4
    fmt.Println(v16, v19)
    var v20 int16 = 32767
    var v21 uint = 510
    v22 := GenS3{F0: 2, F1: (v16 | 127)}
    v23 := GenS4{F0: -497}
    v24 := GenS5{F0: 1}
    v25 := GenS4{F0: 9223372036854775807}
    if false {
        v21 = (v21 * genFn0(int(v2)))
        var v26 int64 = 9223372036854775807
        switch 4294967295 {
        case 338:
            v21 = v21
        case 81:
            for fi27 := 0; fi27 < 13; fi27++ {
                switch "\n\t" {
                case "\n\t":
                    for fi28 := 0; fi28 < 16; fi28++ {
                        switch "hello" {
                        case "a":
                            fmt.Println(fi28, v26)
                        }
                        switch (fi28 / ((int(v21)) | 1)) {
                        case -998:
                            if (!true) {
                                v5[v1] = "foo bar"
                                v1 = 2147483648
                                switch (v4 & v4) {
                                case 1:
                                    var v29 float64 = (v19 / v19)
                                    _ = v29
                                default:
                                    switch (v4 ^ int(v26)) {
                                    case 0:
                                        var v30 bool = false
                                        _ = v30
                                    case 42:
                                        var v31 int8 = 42
                                        _ = v31
                                    default:
                                        var v32 float64 = (1e10 + float64((1e10 / 3.14159)))
                                        _ = v32
                                    }
                                }
                            } else {
                                if (!(!true)) {
                                    var v33 uint64 = uint64(v16)
                                    v19 += v19
                                    _ = v33
                                } else {
                                    if true {
                                        _ = len(v3)
                                    } else {
                                        switch ("foo bar" + "") {
                                        case "":
                                            if false {
                                                fmt.Println(v16)
                                                if true {
                                                    if (true || false) {
                                                        v5[(v1 >> ((uint(v1)) & 63))] = (("foo bar" + (("\n\t" + "a") + ("a" + ("" + (("a" + "hello") + ("foo bar" + "foo bar")))))) + "foo bar")
                                                    }
                                                    if true {
                                                        if true {
                                                            v34 := GenS5{F0: (v20 + -1)}
                                                            _ = v34
                                                        } else {
                                                            v26 -= (v26 / ((v26) | 1))
                                                        }
                                                    } else {
                                                        var v35 string = "\n\t"
                                                        _ = v35
                                                    }
                                                } else {
                                                    v2 = int32(v26)
                                                    var v36 string = "\n\t"
                                                    _ = v36
                                                }
                                            }
                                        case "hello":
                                            v19 = 1.5
                                        case "a":
                                            if (-1 != v2) {
                                                v37 := v3["\n\t"]
                                                v21 -= (v21 / ((v21) | 1))
                                                _ = v37
                                            } else {
                                                delete(v5, 2147483648)
                                                v38 := GenS4{F0: -3}
                                                switch 42 {
                                                case -1:
                                                    var v39 uint32 = 1
                                                    _ = v39
                                                case -363:
                                                    var v40 float32 = 1.5
                                                    _ = v40
                                                case 1:
                                                    v41 := GenS2{F0: 70, F1: 32767, F2: 218}
                                                    _ = v41
                                                default:
                                                    var v42 float64 = (v19 * v19)
                                                    _ = v42
                                                }
                                                fmt.Println(v20, v20, v2)
                                                var v43 uint8 = (v16 / ((v16) | 1))
                                                var v44 uint8 = (^(v16))
                                                _ = v38
                                                _ = v43
                                                _ = v44
                                            }
                                        }
                                        if false {
                                            var v45 int16 = v20
                                            v2 = (v2 + v2)
                                            var v46 uint = v21
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                }
                                v4 = 32768
                            }
                        case -927:
                            v2 = v2
                        case -449:
                            var v47 uint32 = 672
                            _ = v47
                        }
                        _ = fi28
                    }
                case "":
                    switch "a" {
                    case "\n\t":
                        v48 := GenS1{F0: fi27, F1: ("hello" + "")}
                        _ = v48
                    case "":
                        switch (fi27 >> ((v21) & 63)) {
                        case 29:
                            v19 = v19
                        }
                    default:
                        var v49 uint = 4294967295
                        _ = v49
                    }
                default:
                    v50 := GenS1{F0: v4, F1: "a"}
                    _ = v50
                }
                var v51 int64 = (v26 - (v26 % ((int64(v21)) | 1)))
                _ = fi27
                _ = v51
            }
        case 4294967295:
            _ = len(v3)
        }
        _ = v26
    }
    v52 := []int{int(v21), 127, v4, (v4 | int(v20))}
    switch "" {
    case "hello":
        v1 += v1
    case "foo bar":
        var v53 uint16 = 596
        _ = v53
    }
    for fi54 := 0; fi54 < 17; fi54++ {
        if (true && true) {
            var v55 bool = ((v2 & -1) <= v2)
            switch (fi54 / ((int(v2)) | 1)) {
            case 42:
                if v55 {
                    v56 := GenS1{F0: (-(fi54)), F1: ("foo bar" + ("\n\t" + "hello"))}
                    _ = v56
                } else {
                    v57 := GenS6{F0: uint16(v21), F1: v20, F2: ("hello" + "foo bar")}
                    for fi58 := 0; fi58 < 3; fi58++ {
                        v16 = (v16 - 255)
                        var v59 int32 = (v2 << ((v21) & 63))
                        switch (v4 & fi58) {
                        case 128:
                            if ((-2147483648 != (v59 % ((int32(v20)) | 1))) && v55) {
                                if (true || v55) {
                                    switch fi58 {
                                    case 1:
                                        var v60 int = (fi58 - (v4 + fi54))
                                        _ = v60
                                    default:
                                        v19 = v19
                                    }
                                    switch (fi54 >> (((v21 % ((471) | 1))) & 63)) {
                                    case 32768:
                                        if (v1 < (v1 - (v1 % ((v1) | 1)))) {
                                            if true {
                                                switch int(v16) {
                                                case 1:
                                                    fmt.Println(fi58, v2)
                                                case 9223372036854775807:
                                                    fmt.Println(v1)
                                                }
                                                fmt.Println(v20, v19, v1)
                                            }
                                            switch int(v1) {
                                            case 128:
                                                var v61 float32 = (3.14159 * float32(v19))
                                                _ = v61
                                            default:
                                                if v55 {
                                                    var v62 uint16 = 192
                                                    if v55 {
                                                        var v63 bool = true
                                                        if false {
                                                            v64 := GenS7{F0: ((("foo bar" + "foo bar") + (("" + ((("foo bar" + "\n\t") + "a") + "")) + (((("foo bar" + "hello") + ((("\n\t" + ((("" + ("foo bar" + ("foo bar" + ""))) + ("" + "hello")) + "\n\t")) + "hello") + "\n\t")) + "foo bar") + "hello"))) + ""), F1: -2147483648}
                                                            if len(v52) > 0 {
                                                                v52[0] = 127
                                                            }
                                                            v65 := GenS3{F0: 821, F1: v16}
                                                            v66 := GenS8{F0: 0, F1: (v20 % ((-32768) | 1))}
                                                            v19 += (float64(0.1) * (v19 - -2.25))
                                                            var v67 float32 = (1e10 - 3.14159)
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                            _ = v67
                                                        } else {
                                                            v68 := GenS9{F0: fi58}
                                                            var v69 int16 = -32768
                                                            v69 += v69
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                        _ = v63
                                                    } else {
                                                        var v70 float64 = (v19 + 1e10)
                                                        v52 = append(v52, 244)
                                                        _ = v70
                                                    }
                                                    _ = v62
                                                }
                                            }
                                            var v71 uint64 = 308
                                            _ = v71
                                        } else {
                                            v72 := GenS2{F0: 1, F1: 42, F2: v16}
                                            _ = v72
                                        }
                                    case 128:
                                        v4 = v4
                                    case 4294967295:
                                        v52 = append(v52, (fi54 % (((fi54 >> ((2) & 63))) | 1)))
                                    default:
                                        v73 := v3["hello"]
                                        _ = v73
                                    }
                                    switch ("\n\t" + "") {
                                    case "a":
                                        var v74 uint16 = 0
                                        _ = v74
                                    case "\n\t":
                                        v20 += (-(v20))
                                        v19 = v19
                                        v59 = v2
                                    case "hello":
                                        var v75 int64 = 1
                                        var v76 uint16 = 1
                                        _ = v75
                                        _ = v76
                                    }
                                    fmt.Println(v2, fi58, v55)
                                } else {
                                    if len(v52) > 0 {
                                        v52[0] = fi54
                                    }
                                    v19 += 3.14159
                                }
                            }
                        case -635:
                            v16 += v16
                        case 38:
                            v77 := GenS8{F0: 2147483647, F1: (-(v20))}
                            _ = v77
                        }
                        _ = fi58
                        _ = v59
                    }
                    _ = v57
                }
            case -9223372036854775808:
                var v78 uint = (v21 * (v21 + 0))
                _ = v78
            default:
                var v79 float64 = (v19 - 0.1)
                _ = v79
            }
            _ = v55
        }
        v80 := GenS10{F0: (v4 | (fi54 + fi54)), F1: (v2 | (v2 + (v2 >> ((v21) & 63))))}
        fmt.Println(v2)
        _ = fi54
        _ = v80
    }
    v81 := GenS10{F0: (v4 / ((v4) | 1)), F1: (v2 | (v2 ^ (v2 | v2)))}
    v82 := GenS11{F0: v4, F1: (v1 * genFn1((v21 - (v21 | 2))))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v52:", len(v52), v52)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v81
    _ = v82
}
