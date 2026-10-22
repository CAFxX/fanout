package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int32; F2 uint16 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint; F1 uint16; F2 int16 }
type GenS10 struct { F0 uint32; F1 bool }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS13 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 uint32 = 2147483647
    v4 := GenS1{F0: (^(v3))}
    fmt.Println(v3)
    v3 = (^(v3))
    if true {
        fmt.Println(v3)
        fmt.Println(v3)
        var v5 uint32 = v3
        _ = v5
    } else {
        for fi6 := 0; fi6 < 5; fi6++ {
            v3 = 4294967295
            if false {
                for fi7 := 0; fi7 < 1; fi7++ {
                    v8 := GenS2{F0: 255, F1: 497, F2: v3}
                    _ = fi7
                    _ = v8
                }
            } else {
                var v9 uint64 = 42
                if false {
                    v10 := GenS3{F0: (!(!true)), F1: 1000, F2: 0}
                    v9 = v9
                    var v11 int16 = int16(v9)
                    _ = v10
                    _ = v11
                }
                _ = v9
            }
            for fi12 := 0; fi12 < 17; fi12++ {
                switch 32768 {
                case 932:
                    var v13 float32 = 3.14159
                    _ = v13
                case 2147483647:
                    fmt.Println(fi12, fi6, fi6)
                case -992:
                    var v14 uint16 = 65535
                    _ = v14
                }
                _ = fi12
            }
            _ = fi6
        }
        v15 := GenS4{F0: ("foo bar" + "foo bar")}
        v16 := GenS1{F0: 4294967295}
        _ = v15
        _ = v16
    }
    var v17 float32 = -2.25
    if ((2 == 0) || false) {
        v3 += 42
        for fi18 := 0; fi18 < 10; fi18++ {
            switch (fi18 ^ fi18) {
            case 1:
                v19 := GenS5{F0: true, F1: 2147483647, F2: 832}
                _ = v19
            case -369:
                v3 = (v3 % (((v3 + v3)) | 1))
            default:
                if false {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v21 := GenS6{F0: false, F1: 1}
                        var v22 string = ("hello" + ("foo bar" + ("a" + "foo bar")))
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    if true {
                        if false {
                            var v23 uint32 = (v3 + v3)
                            _ = v23
                        } else {
                            var v24 bool = false
                            v25 := GenS3{F0: ((1 > 0) && (1 == 1)), F1: 814, F2: 42}
                            v26 := []int{139, 1, 906, (fi18 + 347)}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        if (1 >= 122) {
                            for fi27 := 0; fi27 < 17; fi27++ {
                                var v28 bool = (9223372036854775807 >= 0)
                                if (v3 > (^(v3))) {
                                    v29 := GenS7{F0: -21}
                                    _ = v29
                                }
                                _ = fi27
                                _ = v28
                            }
                        } else {
                            v30 := []int{fi18, fi18, (fi18 >> ((18446744073709551615) & 63)), (fi18 >> ((255) & 63))}
                            _ = v30
                        }
                        fmt.Println(fi18)
                    } else {
                        switch "a" {
                        case "foo bar":
                            v31 := []int{fi18, 266}
                            _ = v31
                        case "":
                            var v32 uint16 = 0
                            _ = v32
                        case "a":
                            v33 := []int{(fi18 / (((fi18 >> ((0) & 63))) | 1)), (fi18 | 2147483647), fi18, fi18}
                            _ = v33
                        default:
                            switch ((("hello" + "hello") + "foo bar") + "hello") {
                            case "a":
                                v17 = (float32(0.0) / ((v17 + v17) / v17))
                            case "":
                                v34 := []int{fi18, (^(fi18)), fi18}
                                _ = v34
                            case "hello":
                                switch (-(fi18)) {
                                case 9223372036854775807:
                                    v35 := GenS8{F0: ("a" + "foo bar")}
                                    _ = v35
                                case 0:
                                    var v36 bool = true
                                    _ = v36
                                }
                            default:
                                if true {
                                    if ((true && false) || false) {
                                        v37 := GenS9{F0: 1, F1: 2, F2: 271}
                                        for fi38 := 0; fi38 < 3; fi38++ {
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v39 float64 = 3.14159
                                                _ = v39
                                            case "hello":
                                                var v40 int32 = 675
                                                _ = v40
                                            default:
                                                fmt.Println(v17, fi18)
                                            }
                                            delete(v1, 401)
                                            _ = fi38
                                        }
                                        for fi41 := 0; fi41 < 6; fi41++ {
                                            switch "a" {
                                            case "a":
                                                if (1 < 2) {
                                                    var v42 float64 = 0.0
                                                    _ = v42
                                                } else {
                                                    switch (((("hello" + ("foo bar" + (("hello" + "foo bar") + "\n\t"))) + ("" + "a")) + ("" + "hello")) + "foo bar") {
                                                    case "hello":
                                                        fmt.Println(v3)
                                                    }
                                                }
                                            case "":
                                                v43 := GenS2{F0: 1, F1: 43, F2: v3}
                                                _ = v43
                                            }
                                            v44 := GenS2{F0: 0, F1: int32(fi18), F2: v3}
                                            var v45 float32 = v17
                                            _ = fi41
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = v37
                                    }
                                }
                            }
                        }
                        for fi46 := 0; fi46 < 19; fi46++ {
                            v17 = 1.5
                            v17 = 1e10
                            v3 = v3
                            _ = fi46
                        }
                    }
                } else {
                    for fi47 := 0; fi47 < 8; fi47++ {
                        switch -855 {
                        case -216:
                            switch (fi47 + (fi47 << ((992) & 63))) {
                            case -650:
                                v1[(^(v3))] = (v3 | (v3 & (v3 / ((v3) | 1))))
                                var v48 float32 = -2.25
                                _ = v48
                            case 409:
                                v3 = (^(v3))
                            case 2147483647:
                                v49 := GenS10{F0: v3, F1: (!true)}
                                _ = v49
                            }
                        case 42:
                            v50 := GenS4{F0: "\n\t"}
                            _ = v50
                        case 858:
                            if (false || false) {
                                var v51 uint8 = 42
                                _ = v51
                            }
                        default:
                            var v52 int8 = -87
                            _ = v52
                        }
                        _ = fi47
                    }
                    v53 := []int{fi18}
                    _ = v53
                }
            }
            v54 := GenS11{F0: 316}
            var v55 int16 = 0
            _ = fi18
            _ = v54
            _ = v55
        }
    }
    v17 = float32(1e10)
    v3 -= v3
    for fi56 := 0; fi56 < 3; fi56++ {
        for fi57 := 0; fi57 < 10; fi57++ {
            v58 := v1[42]
            v59 := GenS12{F0: 4294967295, F1: -1, F2: 42}
            if (!true) {
                v17 = float32(1e10)
                var v60 uint16 = 32767
                v61 := GenS9{F0: 1, F1: (v60 - v60), F2: 1}
                _ = v60
                _ = v61
            } else {
                v62 := GenS12{F0: 642, F1: 0, F2: 32767}
                _ = v62
            }
            _ = fi57
            _ = v58
            _ = v59
        }
        v63 := []int{128}
        _ = fi56
        _ = v63
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        var v64 int16 = -1
        _ = v64
    case "hello":
        v65 := GenS4{F0: (("foo bar" + ("foo bar" + "")) + "foo bar")}
        _ = v65
    }
    fmt.Println(v17, v17)
    v66 := GenS13{F0: (^(v3))}
    if (1 >= 4294967295) {
        var v67 uint16 = 42
        _ = v67
    }
    switch 4294967295 {
    case 1:
        v17 += 3.14159
    }
    v3 = (v3 ^ (^(v3)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v66
}
