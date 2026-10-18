package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int16; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint32; F1 int8; F2 int8 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 int; F1 uint8; F2 bool }
func genFn0() bool {
    if false {
        v5 := GenS1{F0: 236, F1: -9223372036854775808}
        _ = v5
    }
    var v6 uint64 = 0
    if true {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint16 = 42
            _ = fi7
            _ = v8
        }
        var v9 uint = 18446744073709551615
        _ = v9
    } else {
        v6 += v6
    }
    _ = v6
    return (!false)
}

func main() {
    var v1 string = ("" + (("foo bar" + "") + "foo bar"))
    var v2 bool = true
    var v3 int16 = 42
    var v4 int64 = 9223372036854775807
    var v10 uint64 = 42
    v4 = 62
    v11 := []int{378, 4294967295, 2147483647, -736}
    for fi12 := 0; fi12 < 20; fi12++ {
        switch -709 {
        case 0:
            if len(v11) > 0 {
                v11[0] = fi12
            }
        case 32767:
            var v13 int16 = (-(v3))
            _ = v13
        }
        var v14 float64 = -2.25
        if v2 {
            if false {
                _ = len(v11)
                v3 = 69
                v10 -= v10
            }
        }
        _ = fi12
        _ = v14
    }
    var v15 int8 = 127
    fmt.Println(v4, v3)
    switch 0 {
    case 0:
        switch 42 {
        case 42:
            fmt.Println(v2, v1)
        case 128:
            switch (v1 + v1) {
            case "hello":
                fmt.Println(v15, v4)
            default:
                v4 += v4
            }
        }
    case 413:
        switch -2147483648 {
        case -856:
            var v16 uint64 = v10
            _ = v16
        case 848:
            v15 = (v15 + v15)
        case 190:
            v1 = ((v1 + ("\n\t" + (v1 + v1))) + "")
        }
    case 4294967295:
        var v17 int8 = v15
        _ = v17
    }
    fmt.Println(v1)
    v11 = append(v11, 422)
    for fi18 := 0; fi18 < 6; fi18++ {
        if (!(!(!(uint8(v10) >= 51)))) {
            switch fi18 {
            case 877:
                fmt.Println(fi18)
            default:
                for fi19 := 0; fi19 < 11; fi19++ {
                    if true {
                        switch fi18 {
                        case 32768:
                            var v20 int8 = 0
                            _ = v20
                        }
                    } else {
                        fmt.Println(fi19, v3)
                    }
                    _ = fi19
                }
            }
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 string = ("" + (((v1 + v1) + v1) + ("a" + "")))
        _ = fi21
        _ = v22
    }
    if v2 {
        var v23 bool = ((-356 >= v4) || v2)
        _ = v23
    } else {
        for fi24 := 0; fi24 < 12; fi24++ {
            v25 := GenS2{F0: 0, F1: v4, F2: 65535}
            if (303 < 810) {
                var v26 string = "\n\t"
                v1 = "\n\t"
                _ = v26
            }
            var v27 uint16 = 65535
            _ = fi24
            _ = v25
            _ = v27
        }
    }
    var v28 uint = 42
    v4 = (-(v4))
    v28 += v28
    if v2 {
        if (v2 && v2) {
            v29 := GenS3{F0: (v4 << ((2) & 63)), F1: (v4 ^ (v4 - 1))}
            _ = v29
        }
        v11 = append(v11, 127)
        if v2 {
            if v2 {
                if false {
                    v3 = v3
                }
                for fi30 := 0; fi30 < 7; fi30++ {
                    v31 := GenS3{F0: (v4 - v4), F1: 753}
                    _ = fi30
                    _ = v31
                }
            } else {
                if v2 {
                    var v32 uint32 = uint32(v28)
                    _ = v32
                } else {
                    fmt.Println(v1, v3)
                    v3 = v3
                    if true {
                        v1 = "foo bar"
                        if (v15 <= 1) {
                            var v33 int64 = (v4 - v4)
                            v34 := GenS4{F0: 2147483647, F1: (^(v15)), F2: (^(v15))}
                            for fi35 := 0; fi35 < 2; fi35++ {
                                if (!v2) {
                                    var v36 int64 = -1
                                    fmt.Println(v15, v1)
                                    v37 := GenS4{F0: 9, F1: 126, F2: -128}
                                    _ = v36
                                    _ = v37
                                } else {
                                    var v38 float64 = float64(1e10)
                                    if v2 {
                                        fmt.Println(v33, v28, v3)
                                        v39 := GenS5{F0: 0, F1: (v10 % ((155) | 1))}
                                        _ = v39
                                    } else {
                                        for fi40 := 0; fi40 < 14; fi40++ {
                                            if len(v11) > 0 {
                                                v11[0] = fi40
                                            }
                                            _ = fi40
                                        }
                                        for fi41 := 0; fi41 < 13; fi41++ {
                                            fmt.Println(v3, v38)
                                            _ = len(v11)
                                            fmt.Println(v2, v4)
                                            _ = fi41
                                        }
                                        v4 += (v4 + 9223372036854775807)
                                    }
                                    var v42 float32 = float32(v38)
                                    _ = v38
                                    _ = v42
                                }
                                var v43 int = fi35
                                _ = fi35
                                _ = v43
                            }
                            _ = v33
                            _ = v34
                        } else {
                            for fi44 := 0; fi44 < 4; fi44++ {
                                v2 = (!(v1 > v1))
                                _ = fi44
                            }
                            if v2 {
                                var v45 float32 = 3.14159
                                switch -486 {
                                case 819:
                                    for fi46 := 0; fi46 < 15; fi46++ {
                                        var v47 string = (v1 + (v1 + (v1 + "\n\t")))
                                        _ = fi46
                                        _ = v47
                                    }
                                case 1:
                                    var v48 int8 = v15
                                    _ = v48
                                case -709:
                                    switch ((v1 + v1) + (((v1 + ("foo bar" + "")) + v1) + "a")) {
                                    case "foo bar":
                                        v49 := GenS6{F0: 118, F1: 2147483648}
                                        _ = v49
                                    case "a":
                                        v50 := GenS1{F0: 1, F1: (-(v4))}
                                        _ = v50
                                    case "hello":
                                        if v2 {
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch 127 {
                                                case -858:
                                                    v10 = v10
                                                case 32768:
                                                    var v51 string = v1
                                                    _ = v51
                                                }
                                            case "a":
                                                var v52 uint8 = 128
                                                _ = v52
                                            case "":
                                                fmt.Println(v3, v1, v2)
                                            default:
                                                v1 = ("" + "\n\t")
                                            }
                                            if ((v28 % ((v28) | 1)) != 65535) {
                                                fmt.Println(v4, v45, v2)
                                                for fi53 := 0; fi53 < 17; fi53++ {
                                                    if v2 {
                                                        fmt.Println(v10, v3, v28)
                                                        switch v1 {
                                                        case "":
                                                            for fi54 := 0; fi54 < 4; fi54++ {
                                                                var v55 bool = false
                                                                _ = fi54
                                                                _ = v55
                                                            }
                                                        default:
                                                            _ = len(v11)
                                                        }
                                                        v28 -= v28
                                                    }
                                                    _ = len(v11)
                                                    v28 = (^(v28))
                                                    _ = fi53
                                                }
                                            }
                                        } else {
                                            fmt.Println(v15, v10, v28)
                                        }
                                    }
                                }
                                _ = v45
                            }
                        }
                        fmt.Println(v28)
                    } else {
                        v2 = (v3 != (^(v3)))
                        v10 = 42
                        switch 1 {
                        case 127:
                            switch 32767 {
                            case 9223372036854775807:
                                v15 = -1
                            case 42:
                                var v56 int32 = 42
                                _ = v56
                            case 127:
                                var v57 int8 = 1
                                _ = v57
                            }
                        case -827:
                            if true {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                fmt.Println(v3)
                            } else {
                                if v2 {
                                    for fi58 := 0; fi58 < 2; fi58++ {
                                        v1 = ((v1 + "\n\t") + (v1 + ((v1 + "") + "foo bar")))
                                        _ = fi58
                                    }
                                } else {
                                    var v59 int16 = (v3 >> ((v28) & 63))
                                    _ = len(v11)
                                    v4 = (v4 | v4)
                                    _ = v59
                                }
                                var v60 uint64 = 1
                                _ = v60
                            }
                        case 9223372036854775807:
                            v11 = append(v11, -2147483648)
                        default:
                            fmt.Println(v1)
                        }
                    }
                }
                v10 = (^(v10))
                v1 = ""
            }
            for fi61 := 0; fi61 < 3; fi61++ {
                var v62 float64 = 3.14159
                for fi63 := 0; fi63 < 3; fi63++ {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = fi63
                }
                v64 := GenS7{F0: fi61, F1: 127, F2: ((v1 + v1) >= (("" + (("hello" + (v1 + "a")) + v1)) + ("a" + v1)))}
                _ = fi61
                _ = v62
                _ = v64
            }
            var v65 float64 = (0.1 / (float64(float32(3.14159)) * float64(((0.1 * 0.1) - 3.14159))))
            _ = v65
        } else {
            var v66 float32 = -2.25
            for fi67 := 0; fi67 < 18; fi67++ {
                var v68 int16 = (v3 | 32767)
                if (v10 >= 42) {
                    v66 = ((v66 / (v66 * (v66 + 1e10))) * -2.25)
                }
                _ = fi67
                _ = v68
            }
            _ = v66
        }
    }
    var v69 string = "hello"
    if ((-2147483648 == -1) || false) {
        fmt.Println(v4, v10)
    } else {
        v3 = (v3 >> ((v28) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v69
}
