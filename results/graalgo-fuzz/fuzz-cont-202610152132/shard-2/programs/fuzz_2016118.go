package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 bool; F2 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 int; F2 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int64; F1 bool }
func genFn0() uint32 {
    v4 := []int{}
    var v5 int16 = -32768
    if true {
        var v6 int8 = 100
        fmt.Println(v6)
        if false {
            v7 := GenS1{F0: 43, F1: 2, F2: 4294967295}
            v5 -= (v5 + (v5 >> ((165) & 63)))
            _ = v7
        } else {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint8 = 139
                fmt.Println(v9)
                _ = len(v4)
                _ = fi8
                _ = v9
            }
            _ = len(v4)
        }
        _ = v6
    } else {
        v5 -= 0
        v5 = v5
        v5 -= v5
    }
    v5 = 533
    switch ("" + ((("" + "hello") + "a") + "hello")) {
    case "foo bar":
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(fi10)
            var v11 bool = (-128 == 127)
            _ = fi10
            _ = v11
        }
    case "\n\t":
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int64 = 110
            _ = fi12
            _ = v13
        }
    case "a":
        fmt.Println(v5)
    default:
        var v14 uint16 = 65535
        _ = v14
    }
    _ = v4
    _ = v5
    return 65535
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 1
    var v3 float32 = 3.14159
    var v15 int16 = -32768
    v2 = (^(v2))
    v3 += v3
    v16 := GenS2{F0: 752, F1: 2}
    fmt.Println(v3, v3)
    var v17 uint8 = (v2 | (v2 << ((42) & 63)))
    v18 := []int{219, -2147483648, -321, -602}
    if true {
        switch 9223372036854775807 {
        case -188:
            switch 0 {
            case 9223372036854775807:
                var v19 string = ("foo bar" + "a")
                _ = v19
            }
        case -261:
            if false {
                var v20 int32 = -422
                v1[-576] = ""
                _ = v20
            }
        case 912:
            var v21 uint8 = (v17 << ((18446744073709551615) & 63))
            _ = v21
        default:
            switch "\n\t" {
            case "a":
                fmt.Println(v15, v15)
            case "hello":
                var v22 int8 = 42
                _ = v22
            }
        }
        if true {
            var v23 uint = 4294967295
            v24 := v1[int(v15)]
            _ = v23
            _ = v24
        }
        if (!true) {
            var v25 string = ""
            if false {
                v25 = v25
                if (false && true) {
                    fmt.Println(v15)
                }
            } else {
                var v26 uint8 = (v2 >> ((347) & 63))
                var v27 float64 = ((-2.25 + (0.0 + 0.0)) * (0.1 + 3.14159))
                switch (v25 + "") {
                case "":
                    v1[127] = "a"
                case "hello":
                    v28 := GenS2{F0: 0, F1: 255}
                    _ = v28
                }
                _ = v26
                _ = v27
            }
            _ = v25
        }
    } else {
        for fi29 := 0; fi29 < 18; fi29++ {
            if (!false) {
                switch ("" + (((("foo bar" + ("hello" + "a")) + "") + ("foo bar" + "\n\t")) + ("" + "foo bar"))) {
                case "a":
                    if (1e10 == 3.14159) {
                        if true {
                            if false {
                                if ((("hello" + "hello") == "foo bar") && false) {
                                    delete(v1, fi29)
                                    switch fi29 {
                                    case -593:
                                        _ = len(v18)
                                    }
                                    var v30 float32 = v3
                                    _ = v30
                                } else {
                                    _ = len(v18)
                                    for fi31 := 0; fi31 < 2; fi31++ {
                                        _ = len(v1)
                                        v32 := GenS2{F0: 65535, F1: (v2 - (v2 / (((v17 + 105)) | 1)))}
                                        var v33 int64 = 275
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                    }
                                }
                                v34 := GenS2{F0: uint32(v2), F1: 0}
                                _ = v34
                            }
                            var v35 uint8 = (v2 * v2)
                            v36 := GenS3{F0: 0, F1: "\n\t"}
                            _ = v35
                            _ = v36
                        }
                    } else {
                        if true {
                            v17 = (v2 | 2)
                        }
                        v37 := GenS3{F0: 886, F1: ("a" + (("hello" + "\n\t") + "\n\t"))}
                        _ = v37
                    }
                case "\n\t":
                    v2 = 255
                default:
                    var v38 uint64 = 1
                    _ = v38
                }
            }
            for fi39 := 0; fi39 < 20; fi39++ {
                var v40 float32 = 0.0
                _ = fi39
                _ = v40
            }
            _ = fi29
        }
        switch 500 {
        case 32767:
            v41 := GenS1{F0: v2, F1: 477, F2: 2}
            _ = v41
        case 1:
            var v42 uint64 = 2
            _ = v42
        case 677:
            for fi43 := 0; fi43 < 1; fi43++ {
                v44 := GenS2{F0: 4294967295, F1: v2}
                v45 := GenS1{F0: 127, F1: uint64(v17), F2: 2147483648}
                _ = fi43
                _ = v44
                _ = v45
            }
        }
        switch "\n\t" {
        case "foo bar":
            v46 := GenS3{F0: 18446744073709551615, F1: ("foo bar" + "\n\t")}
            _ = v46
        case "a":
            v47 := GenS1{F0: (v17 << ((2) & 63)), F1: 42, F2: genFn0()}
            _ = v47
        case "hello":
            v15 -= v15
        default:
            v48 := GenS1{F0: v2, F1: 42, F2: genFn0()}
            _ = v48
        }
    }
    switch "foo bar" {
    case "hello":
        v3 -= float32(-2.25)
    case "":
        if (-2147483648 <= 42) {
            for fi49 := 0; fi49 < 4; fi49++ {
                if (genFn0() > 0) {
                    fmt.Println(v3, v15, fi49)
                    v50 := GenS4{F0: (65535 <= 312)}
                    _ = v50
                }
                if (912 != 65535) {
                    v3 -= (v3 - -2.25)
                    if true {
                        var v51 uint64 = 15
                        var v52 bool = false
                        _ = v51
                        _ = v52
                    }
                } else {
                    switch (("foo bar" + ("hello" + "hello")) + "a") {
                    case "\n\t":
                        switch fi49 {
                        case 128:
                            fmt.Println(v15, v17, v17)
                        default:
                            var v53 float64 = ((((((0.0 + -2.25) - 0.1) + (0.1 - (float64(float32(-2.25)) + float64(v3)))) + 0.0) / (-2.25 + float64(v3))) + 1e10)
                            _ = v53
                        }
                    case "hello":
                        var v54 int8 = 42
                        _ = v54
                    case "foo bar":
                        v55 := GenS5{F0: v2, F1: false, F2: (fi49 * 32767)}
                        _ = v55
                    default:
                        switch (fi49 * 32768) {
                        case -2147483648:
                            switch 127 {
                            case 32767:
                                v56 := GenS6{F0: 837}
                                _ = v56
                            case 42:
                                var v57 uint = 1
                                _ = v57
                            default:
                                switch (("foo bar" + "hello") + "") {
                                case "foo bar":
                                    v58 := GenS7{F0: (v15 | (v15 * -935)), F1: fi49, F2: 100}
                                    _ = v58
                                }
                            }
                        case 127:
                            v59 := GenS8{F0: -521}
                            _ = v59
                        default:
                            if false {
                                v3 = v3
                                var v60 int32 = -2147483648
                                switch "a" {
                                case "":
                                    var v61 int8 = 0
                                    _ = v61
                                }
                                _ = v60
                            }
                        }
                    }
                    v15 = v15
                }
                _ = fi49
            }
            v3 -= 3.14159
        }
    case "a":
        var v62 uint = 255
        _ = v62
    default:
        var v63 string = "a"
        _ = v63
    }
    _ = len(v18)
    for fi64 := 0; fi64 < 3; fi64++ {
        var v65 string = ""
        _ = fi64
        _ = v65
    }
    var v66 uint16 = 2
    switch "" {
    case "a":
        if len(v18) > 0 {
            v18[0] = int(v66)
        }
    case "hello":
        v67 := GenS4{F0: true}
        _ = v67
    case "":
        v17 += 128
    default:
        v68 := GenS9{F0: 9223372036854775807, F1: (!false)}
        _ = v68
    }
    v3 += ((0.0 + v3) / v3)
    v69 := GenS8{F0: 0}
    var v70 int16 = 96
    var v71 int16 = 576
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v66:", v66)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v66
    _ = v69
    _ = v70
    _ = v71
}
