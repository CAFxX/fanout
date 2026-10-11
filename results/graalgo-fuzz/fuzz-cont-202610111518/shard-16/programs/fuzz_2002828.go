package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int32; F2 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 string; F1 int16; F2 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16; F1 int16; F2 string }
type GenS10 struct { F0 uint16; F1 uint; F2 string }
func genFn0(p0 int, p1 uint, p2 int16) bool {
    switch "\n\t" {
    case "a":
        switch ("a" + "") {
        case "":
            for fi4 := 0; fi4 < 17; fi4++ {
                var v5 uint = p1
                var v6 bool = false
                switch "foo bar" {
                case "hello":
                    if true {
                        v7 := GenS1{F0: 128}
                        _ = v7
                    }
                case "":
                    v8 := GenS2{F0: 42}
                    _ = v8
                default:
                    v5 = (^(p1))
                }
                _ = fi4
                _ = v5
                _ = v6
            }
        default:
            fmt.Println(p1, p0)
        }
    case "foo bar":
        if true {
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS3{F0: 1, F1: -462, F2: false}
                v11 := GenS2{F0: -294}
                v12 := GenS1{F0: 2}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            p2 = (p2 | -316)
        }
    default:
        fmt.Println(p1)
    }
    var v13 float64 = (0.1 - -2.25)
    var v14 uint = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return ((v13 / v13) != v13)
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -1
    var v3 int16 = 0
    if (((42 != 42) && (127 != -882)) && true) {
        if (!(1e10 == 1e10)) {
            var v15 int32 = 42
            v16 := GenS4{F0: 694}
            v17 := GenS5{F0: int64(v2)}
            _ = v15
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 int8 = v2
                var v20 uint8 = 1
                v20 -= v20
                _ = fi18
                _ = v19
                _ = v20
            }
            v3 = v3
            if true {
                v1[int(v2)] = ((("hello" + "foo bar") + ("foo bar" + "")) + ("" + "a"))
                fmt.Println(v2, v3)
                fmt.Println(v3, v2)
            } else {
                v2 = v2
            }
        }
        switch (("foo bar" + "") + ("a" + (("" + "\n\t") + ("a" + "foo bar")))) {
        case "\n\t":
            v21 := GenS6{F0: v2, F1: 4294967295}
            _ = v21
        default:
            switch 493 {
            case -2147483648:
                v2 = (v2 >> ((42) & 63))
            case 32767:
                switch "a" {
                case "":
                    if true {
                        var v22 int = 4294967295
                        v23 := []int{v22, (v22 ^ v22)}
                        if true {
                            _ = len(v1)
                            for fi24 := 0; fi24 < 18; fi24++ {
                                var v25 string = ""
                                _ = fi24
                                _ = v25
                            }
                            var v26 string = "foo bar"
                            _ = v26
                        }
                        _ = v22
                        _ = v23
                    } else {
                        delete(v1, 128)
                        v2 = v2
                        var v27 int16 = v3
                        _ = v27
                    }
                }
            case 0:
                var v28 uint32 = 998
                _ = v28
            default:
                v2 = v2
            }
        }
        v29 := GenS7{F0: "foo bar", F1: -338, F2: 128}
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        for fi31 := 0; fi31 < 16; fi31++ {
            fmt.Println(fi31)
            v32 := GenS4{F0: 2147483648}
            _ = fi31
            _ = v32
        }
        v33 := v1[2147483647]
        var v34 int16 = -255
        _ = fi30
        _ = v33
        _ = v34
    }
    if (!(0 < 255)) {
        for fi35 := 0; fi35 < 3; fi35++ {
            var v36 int8 = v2
            v36 = (v36 ^ (v36 * v36))
            if false {
                v37 := GenS1{F0: 127}
                delete(v1, -9223372036854775808)
                var v38 uint8 = 127
                _ = v37
                _ = v38
            }
            _ = fi35
            _ = v36
        }
    } else {
        if ((4294967295 >= 754) || true) {
            var v39 float64 = (0.0 * 1.5)
            var v40 float32 = 3.14159
            v41 := GenS7{F0: "a", F1: v3, F2: -600}
            _ = v39
            _ = v40
            _ = v41
        } else {
            fmt.Println(v3, v2)
        }
    }
    fmt.Println(v3)
    fmt.Println(v2)
    for fi42 := 0; fi42 < 7; fi42++ {
        if false {
            v43 := GenS8{F0: 1}
            _ = v43
        } else {
            v44 := GenS9{F0: v3, F1: (^(v3)), F2: "\n\t"}
            v2 = -1
            _ = v44
        }
        if false {
            for fi45 := 0; fi45 < 8; fi45++ {
                var v46 bool = true
                v47 := []int{}
                _ = fi45
                _ = v46
                _ = v47
            }
            v48 := GenS6{F0: 106, F1: 4294967295}
            if (false && false) {
                var v49 float64 = -2.25
                var v50 uint = 42
                var v51 uint = (v50 & v50)
                _ = v49
                _ = v50
                _ = v51
            }
            _ = v48
        }
        fmt.Println(fi42, fi42)
        _ = fi42
    }
    for fi52 := 0; fi52 < 6; fi52++ {
        var v53 uint8 = 0
        v2 = v2
        _ = fi52
        _ = v53
    }
    switch -962 {
    case 582:
        var v54 float32 = 3.14159
        _ = v54
    case 32767:
        var v55 string = ("foo bar" + "a")
        _ = v55
    }
    switch 2147483647 {
    case 32768:
        var v56 int32 = -729
        _ = v56
    default:
        v57 := GenS10{F0: 0, F1: 0, F2: (("a" + ("a" + "\n\t")) + ("a" + "a"))}
        _ = v57
    }
    for fi58 := 0; fi58 < 9; fi58++ {
        for fi59 := 0; fi59 < 17; fi59++ {
            v60 := []int{128, fi58, fi58}
            _ = fi59
            _ = v60
        }
        switch 42 {
        case -873:
            v61 := GenS2{F0: -2147483648}
            _ = v61
        case 32767:
            var v62 uint8 = 0
            _ = v62
        case 756:
            var v63 int64 = 0
            _ = v63
        default:
            v2 -= -91
        }
        fmt.Println(v3)
        _ = fi58
    }
    if false {
        v3 += v3
        if (143 < 2) {
            var v64 float32 = float32(-2.25)
            var v65 bool = (9223372036854775807 > 1)
            var v66 int64 = 42
            _ = v64
            _ = v65
            _ = v66
        } else {
            fmt.Println(v3)
            var v67 int32 = -2147483648
            if false {
                v1[127] = ""
            }
            _ = v67
        }
    }
    v1[42] = ("" + "hello")
    _ = len(v1)
    var v68 int = -9223372036854775808
    v68 = (v68 << ((127) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v68
}
