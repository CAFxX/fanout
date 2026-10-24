package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 string; F2 int32 }
type GenS5 struct { F0 int8; F1 uint32; F2 string }
type GenS6 struct { F0 string; F1 int; F2 uint8 }
type GenS7 struct { F0 int32; F1 uint8; F2 uint }
type GenS8 struct { F0 uint8; F1 uint64; F2 bool }
type GenS9 struct { F0 uint64; F1 int; F2 int64 }
type GenS10 struct { F0 int16; F1 uint; F2 uint8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint64 }
func genFn0(p0 bool, p1 int32) int {
    switch ("hello" + ((("hello" + ("hello" + "hello")) + "\n\t") + "hello")) {
    case "":
        for fi4 := 0; fi4 < 16; fi4++ {
            if (42 >= -107) {
                v5 := []int{}
                p1 -= (p1 & (^(p1)))
                _ = v5
            } else {
                v6 := GenS1{F0: 2, F1: 4294967295}
                _ = v6
            }
            _ = fi4
        }
    case "hello":
        var v7 int16 = 1
        _ = v7
    case "foo bar":
        v8 := GenS1{F0: 42, F1: 4294967295}
        _ = v8
    default:
        fmt.Println(p0, p0)
    }
    v9 := []int{}
    p0 = p0
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            switch 127 {
            case 4294967295:
                p1 += p1
            case 9223372036854775807:
                p1 += p1
            default:
                switch "" {
                case "":
                    v12 := GenS1{F0: 2, F1: 49}
                    _ = v12
                default:
                    fmt.Println(fi10, fi10, p0)
                }
                var v13 uint32 = 1
                var v14 float32 = (1e10 / 3.14159)
                if false {
                    p0 = false
                    var v15 bool = p0
                    if (!false) {
                        _ = len(v9)
                        var v16 string = ("hello" + "hello")
                        _ = v16
                    }
                    _ = v15
                }
                switch "\n\t" {
                case "":
                    fmt.Println(v13)
                default:
                    fmt.Println(p0, fi11)
                }
                switch fi10 {
                case 4294967295:
                    if len(v9) > 0 {
                        v9[0] = (fi10 / ((fi10) | 1))
                    }
                case 983:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                default:
                    if len(v9) > 0 {
                        v9[0] = fi11
                    }
                }
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        if false {
            var v17 uint32 = 840
            for fi18 := 0; fi18 < 8; fi18++ {
                p0 = p0
                fmt.Println(p1)
                v17 -= (v17 >> ((42) & 63))
                _ = fi18
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 int64 = 1
                _ = fi19
                _ = v20
            }
            _ = v17
        } else {
            var v21 float64 = (0.1 / 1.0)
            p0 = (12 == 73)
            p0 = true
            _ = v21
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return 128
}
func genFn1(p0 int) uint {
    switch (("hello" + (("\n\t" + ("\n\t" + "\n\t")) + "a")) + "a") {
    case "hello":
        switch (("a" + "foo bar") + ("hello" + "a")) {
        case "a":
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS2{F0: 554}
                _ = fi22
                _ = v23
            }
        case "":
            if true {
                fmt.Println(p0)
                v24 := []int{}
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 uint16 = 42
                    p0 += genFn0(true, 2147483647)
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
        default:
            var v27 uint = 255
            _ = v27
        }
    }
    if (!(-9223372036854775808 > 1)) {
        if (42 >= 3) {
            v28 := GenS3{F0: 65535}
            v29 := []int{}
            _ = v28
            _ = v29
        }
        v30 := []int{}
        p0 = p0
        _ = v30
    } else {
        switch p0 {
        case 128:
            p0 = p0
        }
    }
    var v31 uint32 = 2147483647
    if true {
        if false {
            var v32 int32 = 840
            fmt.Println(v32, p0, v32)
            v33 := []int{(p0 - (p0 % (((p0 * (p0 % ((genFn0(((v31 >> ((862) & 63)) == (^(v31))), v32)) | 1)))) | 1))), genFn0(true, (v32 - (v32 << ((4294967295) & 63)))), (p0 % (((p0 % ((128) | 1))) | 1)), -585}
            _ = v32
            _ = v33
        }
    } else {
        var v34 int = genFn0(((3.14159 / 3.14159) >= (-2.25 + (3.14159 + 1e10))), -2147483648)
        _ = v34
    }
    _ = p0
    _ = v31
    return 255
}

func main() {
    v1 := []string{"foo bar", "a"}
    v2 := make(map[int]string)
    var v3 uint32 = 2147483648
    if (((v3 * v3) < 2147483647) && true) {
        for fi35 := 0; fi35 < 3; fi35++ {
            v3 -= (v3 & (v3 & v3))
            var v36 int64 = -1
            v37 := GenS4{F0: 18446744073709551615, F1: "hello", F2: -2147483648}
            _ = fi35
            _ = v36
            _ = v37
        }
        v38 := v2[-293]
        _ = v38
    }
    switch 32767 {
    case -9223372036854775808:
        v3 -= (v3 - 59)
    case 4294967295:
        var v39 string = "\n\t"
        _ = v39
    case -606:
        _ = len(v1)
    default:
        v1 = append(v1, "")
    }
    var v40 int = 32768
    for fi41 := 0; fi41 < 11; fi41++ {
        var v42 uint8 = 70
        _ = fi41
        _ = v42
    }
    fmt.Println(v3)
    var v43 string = "foo bar"
    fmt.Println(v43, v43, v43)
    if true {
        v40 -= (v40 >> ((1) & 63))
        if true {
            if true {
                var v44 int8 = 0
                _ = v44
            }
            var v45 string = (v43 + (v43 + "hello"))
            for fi46 := 0; fi46 < 8; fi46++ {
                if true {
                    fmt.Println(v3)
                    v45 = v43
                    var v47 int16 = 32767
                    _ = v47
                } else {
                    fmt.Println(v40)
                    if (false || (fi46 > 841)) {
                        v48 := GenS2{F0: (fi46 - v40)}
                        _ = v48
                    }
                    var v49 bool = false
                    _ = v49
                }
                if (127 > 42) {
                    switch ((v43 + v43) + v45) {
                    case "foo bar":
                        for fi50 := 0; fi50 < 2; fi50++ {
                            v45 = v45
                            v51 := GenS1{F0: 42, F1: v3}
                            _ = fi50
                            _ = v51
                        }
                    default:
                        v52 := GenS5{F0: 0, F1: (v3 / ((v3) | 1)), F2: v43}
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 7; fi53++ {
                        var v54 float32 = float32(3.14159)
                        if ((-2.25 * (-2.25 * 0.0)) >= (0.0 + 1.5)) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                            if false {
                                switch "foo bar" {
                                case "":
                                    v55 := GenS6{F0: "a", F1: fi46, F2: 0}
                                    _ = v55
                                case "a":
                                    if (((!false) && (1 != 42)) || (42 != 1)) {
                                        v3 = (v3 ^ v3)
                                        var v56 uint32 = v3
                                        _ = v56
                                    } else {
                                        v57 := GenS7{F0: -2147483648, F1: 1, F2: 622}
                                        v58 := GenS8{F0: 2, F1: 2, F2: (!true)}
                                        _ = v57
                                        _ = v58
                                    }
                                    v45 = (v45 + "foo bar")
                                    var v59 float64 = (1e10 + -2.25)
                                    _ = v59
                                case "\n\t":
                                    if true {
                                        v60 := GenS9{F0: 2, F1: -2147483648, F2: -810}
                                        _ = v60
                                    } else {
                                        if ((1.5 >= float32(((0.0 * (1e10 + (3.14159 + -2.25))) * 0.0))) && (35 != 4294967295)) {
                                            v61 := GenS10{F0: 358, F1: 4294967295, F2: 52}
                                            v62 := GenS10{F0: -1, F1: 2, F2: 2}
                                            _ = v61
                                            _ = v62
                                        } else {
                                            if (!true) {
                                                var v63 bool = false
                                                v63 = (!v63)
                                                if (!(726 <= 728)) {
                                                    fmt.Println(v54)
                                                    fmt.Println(v40, fi53, v45)
                                                    _ = len(v1)
                                                } else {
                                                    v64 := GenS6{F0: (("\n\t" + v45) + "foo bar"), F1: (-(v40)), F2: 233}
                                                    _ = v64
                                                }
                                                _ = v63
                                            }
                                        }
                                    }
                                default:
                                    fmt.Println(v54, v43)
                                }
                                fmt.Println(v45)
                                v45 = v45
                            } else {
                                if true {
                                    fmt.Println(fi53, v54)
                                    var v65 uint64 = 4294967295
                                    if true {
                                        var v66 uint64 = 1
                                        _ = v66
                                    } else {
                                        var v67 int16 = 32767
                                        fmt.Println(v54)
                                        if false {
                                            switch "a" {
                                            case "\n\t":
                                                switch genFn0(false, 0) {
                                                case 978:
                                                    fmt.Println(v45)
                                                case 0:
                                                    var v68 float32 = v54
                                                    _ = v68
                                                }
                                            case "foo bar":
                                                v2[v40] = ("foo bar" + ((((v45 + v45) + "foo bar") + "a") + v43))
                                            case "hello":
                                                if false {
                                                    v65 += (v65 << ((genFn1((-(fi53)))) & 63))
                                                    v69 := GenS11{F0: 4294967295}
                                                    v70 := GenS6{F0: v45, F1: genFn0(false, 495), F2: 231}
                                                    if (1.5 > v54) {
                                                        if false {
                                                            var v71 bool = (!true)
                                                            v3 -= uint32(v65)
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    v67 = v67
                                                    if len(v1) > 0 {
                                                        v1[0] = "foo bar"
                                                    }
                                                }
                                                v72 := GenS8{F0: 229, F1: 923, F2: (!(!true))}
                                                if false {
                                                    v45 = v45
                                                    v65 = (v65 + v65)
                                                } else {
                                                    v73 := GenS12{F0: 0}
                                                    _ = v73
                                                }
                                                _ = v72
                                            }
                                            var v74 uint16 = 2
                                            _ = v74
                                        }
                                        _ = v67
                                    }
                                    _ = v65
                                } else {
                                    var v75 string = "foo bar"
                                    if len(v1) > 0 {
                                        v1[0] = v43
                                    }
                                    v75 = v43
                                    _ = v75
                                }
                                var v76 uint16 = 42
                                _ = v76
                            }
                        } else {
                            if ((v3 > (v3 ^ v3)) || false) {
                                v3 -= (v3 << ((725) & 63))
                                var v77 uint64 = 0
                                if (2 == 32767) {
                                    switch fi46 {
                                    case -841:
                                        _ = len(v1)
                                    case 127:
                                        var v78 bool = (-32768 > 0)
                                        _ = v78
                                    case 42:
                                        if false {
                                            if (956 != genFn1(699)) {
                                                switch ((v43 + v43) + "") {
                                                case "":
                                                    if (!false) {
                                                        _ = len(v1)
                                                        fmt.Println(v43, fi53, v54)
                                                        var v79 int32 = 1
                                                        _ = v79
                                                    }
                                                case "hello":
                                                    v43 = (v43 + (((v45 + (((v43 + "foo bar") + v45) + (v43 + v43))) + "") + (("" + ((v45 + (v45 + "hello")) + (v43 + "a"))) + v45)))
                                                }
                                                fmt.Println(v3, v54)
                                            } else {
                                                v43 = v45
                                            }
                                            v80 := GenS13{F0: 1}
                                            _ = v80
                                        }
                                    }
                                } else {
                                    var v81 int8 = -1
                                    delete(v2, (fi46 + v40))
                                    _ = v81
                                }
                                _ = v77
                            }
                        }
                        _ = fi53
                        _ = v54
                    }
                } else {
                    for fi82 := 0; fi82 < 4; fi82++ {
                        var v83 int64 = 9223372036854775807
                        v40 = (-(fi82))
                        v83 += v83
                        fmt.Println(v3, v43, v40)
                        v2[127] = ("hello" + v45)
                        fmt.Println(fi82, fi82, fi82)
                        switch v45 {
                        case "hello":
                            v3 += 65535
                        case "\n\t":
                            v84 := GenS7{F0: -319, F1: 255, F2: 898}
                            _ = v84
                        case "foo bar":
                            v85 := GenS5{F0: 127, F1: (v3 / (((v3 / (((v3 / ((2147483647) | 1))) | 1))) | 1)), F2: (("foo bar" + v45) + v45)}
                            _ = v85
                        }
                        v43 = v43
                        v3 = (v3 * v3)
                        v83 = -1
                        _ = fi82
                        _ = v83
                    }
                    for fi86 := 0; fi86 < 8; fi86++ {
                        v87 := GenS14{F0: 18446744073709551615}
                        _ = fi86
                        _ = v87
                    }
                }
                fmt.Println(v3)
                _ = fi46
            }
            _ = v45
        } else {
            fmt.Println(v40)
            var v88 uint32 = v3
            _ = v88
        }
    } else {
        fmt.Println(v40, v3, v40)
    }
    var v89 int64 = 0
    var v90 uint64 = 18446744073709551615
    var v91 int32 = 0
    v89 = v89
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v43
    _ = v89
    _ = v90
    _ = v91
}
