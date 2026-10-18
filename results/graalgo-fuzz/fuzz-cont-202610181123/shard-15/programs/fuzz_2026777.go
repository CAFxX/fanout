package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint; F1 int; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int64 }
func genFn0() uint32 {
    var v4 int = -482
    if false {
        var v5 int = (v4 + v4)
        _ = v5
    } else {
        v6 := GenS1{F0: "foo bar", F1: 127, F2: 65535}
        switch "foo bar" {
        case "hello":
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint64 = 713
                _ = fi7
                _ = v8
            }
        case "foo bar":
            var v9 uint16 = 42
            _ = v9
        case "a":
            fmt.Println(v4)
        }
        _ = v6
    }
    fmt.Println(v4)
    if false {
        for fi10 := 0; fi10 < 5; fi10++ {
            if (("\n\t" + ((("\n\t" + "foo bar") + "hello") + "")) == ("" + "a")) {
                var v11 uint64 = 1
                var v12 int64 = 381
                var v13 int16 = 1000
                _ = v11
                _ = v12
                _ = v13
            } else {
                var v14 float64 = (float64(0.1) - 1e10)
                fmt.Println(v14, v4)
                _ = v14
            }
            var v15 uint32 = 2147483648
            _ = fi10
            _ = v15
        }
        v4 += 0
    } else {
        v16 := GenS2{F0: 2, F1: 100}
        _ = v16
    }
    _ = v4
    return 2147483647
}
func genFn1(p0 int64) int16 {
    v17 := []int{32768, 9223372036854775807, -250}
    var v18 float64 = -2.25
    _ = p0
    _ = v17
    _ = v18
    return 1000
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 892
    var v3 int = 4294967295
    var v19 int8 = int8(v2)
    v3 -= 127
    switch (v3 ^ (v3 + (v3 >> ((1) & 63)))) {
    case 32768:
        for fi20 := 0; fi20 < 3; fi20++ {
            switch "hello" {
            case "hello":
                for fi21 := 0; fi21 < 17; fi21++ {
                    if true {
                        if false {
                            var v22 uint = 42
                            var v23 float32 = 0.0
                            if false {
                                if (float64(v23) == 1.5) {
                                    var v24 int16 = 42
                                    _ = v24
                                }
                            } else {
                                v19 += v19
                                v3 = 32767
                            }
                            _ = v22
                            _ = v23
                        } else {
                            v25 := v1["foo bar"]
                            switch -9223372036854775808 {
                            case 0:
                                var v26 uint64 = 667
                                _ = v26
                            case -336:
                                v27 := []int{(fi21 * (v3 * (v3 << ((4294967295) & 63)))), (fi21 >> ((0) & 63)), fi20, (-(fi21))}
                                _ = v27
                            case 127:
                                fmt.Println(fi21, v19, fi20)
                            }
                            v28 := []int{127, v3}
                            _ = v25
                            _ = v28
                        }
                        delete(v1, (("a" + "\n\t") + "a"))
                    }
                    var v29 float32 = 0.1
                    _ = fi21
                    _ = v29
                }
            default:
                for fi30 := 0; fi30 < 7; fi30++ {
                    switch fi30 {
                    case -9223372036854775808:
                        v2 -= (v2 * (v2 * v2))
                    default:
                        switch "\n\t" {
                        case "hello":
                            switch (-(fi20)) {
                            case -1:
                                switch (fi30 >> ((42) & 63)) {
                                case -1:
                                    var v31 int = (^(v3))
                                    _ = v31
                                case 32768:
                                    v3 += 306
                                default:
                                    delete(v1, "a")
                                }
                            }
                        case "\n\t":
                            var v32 uint64 = (v2 << ((4294967295) & 63))
                            _ = v32
                        case "a":
                            fmt.Println(v3, v19)
                        default:
                            var v33 int32 = -889
                            _ = v33
                        }
                    }
                    switch -351 {
                    case -950:
                        var v34 int8 = v19
                        _ = v34
                    case 1:
                        v35 := []int{}
                        _ = v35
                    case 128:
                        fmt.Println(v2, v3, v2)
                    }
                    _ = fi30
                }
            }
            v2 -= 0
            for fi36 := 0; fi36 < 19; fi36++ {
                switch fi20 {
                case 448:
                    var v37 uint8 = 2
                    _ = v37
                case 9223372036854775807:
                    var v38 float32 = 1.5
                    _ = v38
                case 32768:
                    var v39 uint8 = 127
                    _ = v39
                default:
                    v40 := GenS3{F0: 166, F1: (v3 | -498), F2: 141}
                    _ = v40
                }
                if false {
                    if true {
                        var v41 bool = true
                        if ((((1e10 + (-2.25 + (3.14159 + ((0.1 - 0.1) / (1.5 - 0.0))))) != (0.1 * (0.0 - (float32((float64(-2.25) - 0.1)) + 0.1)))) && false) && v41) {
                            delete(v1, ("a" + "a"))
                            v42 := GenS1{F0: ("\n\t" + "a"), F1: 0, F2: 703}
                            _ = v42
                        }
                        v19 -= v19
                        _ = v41
                    }
                } else {
                    v19 += (v19 | (v19 & (v19 + 42)))
                }
                _ = fi36
            }
            _ = fi20
        }
    }
    if (32767 != 32767) {
        for fi43 := 0; fi43 < 6; fi43++ {
            var v44 float32 = -2.25
            _ = fi43
            _ = v44
        }
        for fi45 := 0; fi45 < 13; fi45++ {
            v19 -= (v19 + (v19 ^ 100))
            if false {
                v2 = (v2 << ((4294967295) & 63))
                switch "hello" {
                case "":
                    v19 = 1
                case "hello":
                    for fi46 := 0; fi46 < 13; fi46++ {
                        v1[("foo bar" + "hello")] = "foo bar"
                        if false {
                            v47 := []int{fi46, (fi45 % ((-2147483648) | 1))}
                            _ = v47
                        }
                        if (42 != 127) {
                            v19 = 127
                            v48 := GenS3{F0: 18446744073709551615, F1: v3, F2: 128}
                            _ = v48
                        }
                        _ = fi46
                    }
                case "a":
                    v49 := GenS2{F0: 42, F1: v19}
                    _ = v49
                }
            } else {
                var v50 uint16 = 1
                fmt.Println(v3, v2, v50)
                v51 := GenS1{F0: ("" + ""), F1: 0, F2: 919}
                _ = v50
                _ = v51
            }
            _ = fi45
        }
    }
    var v52 uint8 = 2
    if ((!(("\n\t" + "hello") > "foo bar")) || (42 <= -49)) {
        fmt.Println(v19, v3)
        fmt.Println(v3, v52)
        v2 = (^(v2))
    }
    v53 := GenS1{F0: "foo bar", F1: v52, F2: 1}
    switch (-(v3)) {
    case 829:
        v54 := GenS4{F0: (v19 * -1), F1: 255}
        _ = v54
    default:
        if true {
            switch "hello" {
            case "":
                v3 = 1
            default:
                fmt.Println(v2, v2)
            }
            for fi55 := 0; fi55 < 11; fi55++ {
                var v56 bool = true
                v3 = 0
                var v57 int = -9223372036854775808
                _ = fi55
                _ = v56
                _ = v57
            }
        } else {
            v58 := []int{32768, v3}
            _ = v58
        }
    }
    v3 = (v3 ^ int(v52))
    for fi59 := 0; fi59 < 18; fi59++ {
        var v60 float64 = 0.0
        var v61 uint8 = (v52 + uint8(v2))
        _ = fi59
        _ = v60
        _ = v61
    }
    switch 93 {
    case 2147483647:
        switch ("\n\t" + "a") {
        case "hello":
            v62 := GenS1{F0: ((("" + "\n\t") + "foo bar") + "\n\t"), F1: (v52 | 129), F2: 1}
            _ = v62
        }
    }
    v63 := GenS5{F0: v52, F1: (-(v19))}
    v64 := GenS5{F0: (v52 ^ (^(v52))), F1: -128}
    if true {
        v3 -= (v3 & 2147483647)
        var v65 uint64 = 607
        _ = v65
    } else {
        var v66 int16 = 32767
        delete(v1, (("a" + "") + "\n\t"))
        v2 -= (v2 << ((1) & 63))
        _ = v66
    }
    if (-185 >= -331) {
        var v67 uint8 = (v52 * (v52 % ((1) | 1)))
        for fi68 := 0; fi68 < 2; fi68++ {
            var v69 uint = 2
            var v70 int8 = (^(v19))
            var v71 bool = (0.1 != (((1e10 - (1.5 * -2.25)) / 1.0) * 0.1))
            _ = fi68
            _ = v69
            _ = v70
            _ = v71
        }
        fmt.Println(v67, v19, v3)
        _ = v67
    } else {
        fmt.Println(v2)
    }
    var v72 int16 = -32768
    v73 := GenS6{F0: 42}
    switch 230 {
    case 32768:
        fmt.Println(v72, v3)
    case 1:
        var v74 int64 = 1
        _ = v74
    case -1:
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v52
    _ = v53
    _ = v63
    _ = v64
    _ = v72
    _ = v73
}
