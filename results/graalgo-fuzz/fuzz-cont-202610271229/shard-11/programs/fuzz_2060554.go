package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 int16 }
type GenS4 struct { F0 int32; F1 uint64; F2 int16 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int16; F2 int8 }
type GenS7 struct { F0 uint32; F1 int32 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := []int{fi4}
        _ = fi4
        _ = v5
    }
    if true {
        switch 42 {
        case 9223372036854775807:
            v6 := []int{-1}
            _ = v6
        case 128:
            var v7 int16 = 0
            _ = v7
        case -1:
            switch "a" {
            case "\n\t":
                var v8 int16 = 734
                _ = v8
            case "a":
                v9 := []int{42, 42, 32768, 32767}
                _ = v9
            default:
                var v10 float32 = 1.5
                _ = v10
            }
        }
        switch 0 {
        case 9223372036854775807:
            var v11 uint64 = 1
            _ = v11
        }
        if (127 != -1) {
            v12 := []int{}
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 float32 = (0.0 * (0.0 + (0.1 - 3.14159)))
                v14 -= (float32(0.0) + v14)
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: true, F1: "hello"}
            _ = v12
            _ = v15
        }
    } else {
        var v16 float64 = (0.1 + (((0.1 + 1e10) * 3.14159) * 1.5))
        if (-9223372036854775808 > -1) {
            v16 -= (3.14159 * v16)
            fmt.Println(v16)
            switch 584 {
            case 0:
                if true {
                    v16 = 1.5
                    if true {
                        var v17 int = -1
                        var v18 int8 = 127
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(v16)
                        if (true && true) {
                            if ((((v16 * 3.14159) - float64((((0.1 - 1.5) / (0.1 * float32(float64(float32((float64(0.0) + v16)))))) + 1.5))) + 1e10) <= 0.1) {
                                var v19 int8 = 127
                                _ = v19
                            } else {
                                fmt.Println(v16)
                            }
                        } else {
                            switch 32767 {
                            case 127:
                                v16 = (0.0 - v16)
                            case 538:
                                switch "" {
                                case "hello":
                                    v20 := []int{}
                                    _ = v20
                                case "":
                                    var v21 uint32 = 593
                                    _ = v21
                                }
                            }
                            var v22 float32 = (1e10 + 3.14159)
                            _ = v22
                        }
                        if true {
                            switch -708 {
                            case -105:
                                v23 := GenS1{F0: true, F1: "\n\t"}
                                _ = v23
                            case 127:
                                var v24 uint16 = 2
                                _ = v24
                            case 32767:
                                v16 = v16
                            default:
                                v16 = 3.14159
                            }
                        } else {
                            v16 += 3.14159
                        }
                    }
                    if false {
                        if (true || false) {
                            var v25 int16 = 0
                            _ = v25
                        }
                        var v26 int8 = 31
                        v27 := GenS1{F0: true, F1: ("" + (("\n\t" + ("\n\t" + "a")) + ("a" + "\n\t")))}
                        _ = v26
                        _ = v27
                    }
                }
            case 93:
                v28 := GenS1{F0: true, F1: (("\n\t" + "foo bar") + "")}
                _ = v28
            case -718:
                switch 0 {
                case 1:
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v30 := []int{fi29}
                        v16 = v16
                        _ = fi29
                        _ = v30
                    }
                case 42:
                    for fi31 := 0; fi31 < 19; fi31++ {
                        v16 += float64((0.1 * (0.1 / -2.25)))
                        v32 := GenS1{F0: true, F1: (("hello" + "") + "")}
                        _ = fi31
                        _ = v32
                    }
                case 4294967295:
                    for fi33 := 0; fi33 < 6; fi33++ {
                        var v34 bool = true
                        _ = fi33
                        _ = v34
                    }
                default:
                    if false {
                        if (-87 >= 0) {
                            var v35 int32 = -2147483648
                            fmt.Println(v16, v16)
                            _ = v35
                        } else {
                            v36 := []int{42, -9223372036854775808, 957}
                            v37 := GenS2{F0: 42, F1: 0, F2: 2147483648}
                            var v38 int = -2147483648
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    }
                }
            default:
                if true {
                    v16 = (float64(0.0) / float64((1.5 - (0.0 * 3.14159))))
                    var v39 int = 263
                    _ = v39
                }
            }
        }
        _ = v16
    }
    var v40 float32 = (0.0 - ((float32(0.1) - 0.0) - 1.5))
    if (false || false) {
        if false {
            for fi41 := 0; fi41 < 1; fi41++ {
                switch ("foo bar" + "a") {
                case "":
                    for fi42 := 0; fi42 < 10; fi42++ {
                        v40 += v40
                        _ = fi42
                    }
                }
                switch "\n\t" {
                case "foo bar":
                    var v43 uint32 = 0
                    _ = v43
                }
                _ = fi41
            }
        }
    }
    fmt.Println(v40)
    _ = v40
    return -32768
}

func main() {
    var v1 uint16 = 2
    var v2 int64 = int64(v1)
    v3 := []uint32{4294967295, 18, 4294967295}
    switch -91 {
    case -194:
        var v44 uint16 = v1
        _ = v44
    case 402:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v1, v2)
    if (2147483648 != 4294967295) {
        var v45 string = ("hello" + "a")
        _ = v45
    } else {
        if (!(752 > 4294967295)) {
            v1 += v1
            for fi46 := 0; fi46 < 10; fi46++ {
                if ("foo bar" >= "hello") {
                    v2 = (v2 % ((v2) | 1))
                    switch "hello" {
                    case "":
                        var v47 uint = 255
                        _ = v47
                    }
                    v2 = v2
                } else {
                    v48 := GenS3{F0: -2147483648, F1: (fi46 & fi46), F2: 1}
                    for fi49 := 0; fi49 < 16; fi49++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                        _ = fi49
                    }
                    _ = v48
                }
                switch "\n\t" {
                case "\n\t":
                    if (-2.25 <= 3.14159) {
                        v2 = (-(v2))
                    }
                default:
                    var v50 int8 = int8(v2)
                    _ = v50
                }
                fmt.Println(fi46, fi46)
                _ = fi46
            }
        } else {
            var v51 int32 = -2147483648
            switch 0 {
            case -335:
                var v52 string = ""
                _ = v52
            case 9223372036854775807:
                v1 -= v1
            case -882:
                for fi53 := 0; fi53 < 4; fi53++ {
                    fmt.Println(v1, v51, v51)
                    _ = fi53
                }
            default:
                v54 := GenS4{F0: v51, F1: 685, F2: 655}
                _ = v54
            }
            v2 = 389
            _ = v51
        }
    }
    v55 := GenS3{F0: 0, F1: 128, F2: 1}
    _ = len(v3)
    v56 := GenS5{F0: 42, F1: v1}
    v1 = v1
    v3 = append(v3, 0)
    v57 := GenS6{F0: 9, F1: genFn0(), F2: 1}
    switch 42 {
    case -786:
        v2 = (v2 - v2)
    case 32767:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 978:
        v58 := GenS2{F0: 2147483647, F1: 42, F2: 42}
        _ = v58
    default:
        var v59 int32 = -2147483648
        _ = v59
    }
    v1 += v1
    for fi60 := 0; fi60 < 1; fi60++ {
        var v61 int8 = 0
        _ = len(v3)
        _ = fi60
        _ = v61
    }
    switch -2147483648 {
    case 1:
        if false {
            if len(v3) > 0 {
                v3[0] = uint32(v1)
            }
            switch "foo bar" {
            case "foo bar":
                v1 += 32767
            }
        }
    case 4294967295:
        var v62 int8 = 127
        _ = v62
    case 127:
        var v63 bool = false
        _ = v63
    default:
        var v64 string = (("foo bar" + ("" + "")) + "foo bar")
        _ = v64
    }
    v65 := GenS7{F0: 0, F1: int32(v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
    _ = v56
    _ = v57
    _ = v65
}
