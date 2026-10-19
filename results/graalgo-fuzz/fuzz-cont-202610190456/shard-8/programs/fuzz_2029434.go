package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 string; F1 bool; F2 int64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint8; F2 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 2
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    var v6 int16 = 32767
    var v7 uint = v2
    switch ("a" + "hello") {
    case "foo bar":
        v8 := []int{32768}
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 12; fi9++ {
            switch ("foo bar" + ("foo bar" + "\n\t")) {
            case "\n\t":
                switch (fi9 ^ int(v2)) {
                case 9223372036854775807:
                    v7 -= (v7 & 18446744073709551615)
                default:
                    v10 := GenS1{F0: -906, F1: fi9, F2: 4294967295}
                    _ = v10
                }
            case "foo bar":
                var v11 float32 = 1e10
                _ = v11
            case "":
                var v12 int64 = 42
                _ = v12
            }
            _ = fi9
        }
    default:
        for fi13 := 0; fi13 < 1; fi13++ {
            v7 = (v2 - v7)
            _ = fi13
        }
    }
    switch "\n\t" {
    case "hello":
        switch "a" {
        case "\n\t":
            if (!(9223372036854775807 >= -1)) {
                for fi14 := 0; fi14 < 5; fi14++ {
                    if ((-2.25 - 1e10) > (((1e10 - 0.0) - float32(0.1)) + (3.14159 / float32(1e10)))) {
                        v15 := GenS2{F0: ((("" + ("foo bar" + (("foo bar" + (("foo bar" + "\n\t") + ("hello" + (("\n\t" + "foo bar") + ("hello" + ""))))) + "\n\t"))) + "") + ("\n\t" + ("hello" + "a"))), F1: (1 == 1), F2: 1}
                        for fi16 := 0; fi16 < 11; fi16++ {
                            v17 := []int{}
                            if ("\n\t" <= ("hello" + "foo bar")) {
                                if true {
                                    v2 -= 482
                                    var v18 uint64 = 2
                                    _ = v18
                                } else {
                                    v19 := GenS3{F0: 42, F1: (0 >= 329)}
                                    _ = v19
                                }
                                var v20 uint8 = 28
                                _ = v20
                            }
                            var v21 bool = (!false)
                            _ = fi16
                            _ = v17
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 19; fi22++ {
                            var v23 uint8 = 128
                            switch ("foo bar" + ("\n\t" + ("hello" + "foo bar"))) {
                            case "\n\t":
                                v24 := []int{(fi22 & (fi22 / ((fi22) | 1))), (-(fi22)), fi14}
                                _ = v24
                            }
                            if (true && (333 != 2147483647)) {
                                if (false || (true || true)) {
                                    switch "a" {
                                    case "hello":
                                        v25 := GenS4{F0: ("\n\t" + "a")}
                                        _ = v25
                                    case "\n\t":
                                        v5[fi22] = 42
                                        v7 = (^(v7))
                                    default:
                                        v6 += (v6 * (^(v6)))
                                    }
                                }
                                fmt.Println(v2)
                                if false {
                                    v23 -= (^(v23))
                                    var v26 int64 = 42
                                    _ = v26
                                } else {
                                    fmt.Println(fi22, v23)
                                    if (-10 < 105) {
                                        var v27 uint16 = 939
                                        delete(v4, 2)
                                        switch "\n\t" {
                                        case "hello":
                                            v6 -= v6
                                        default:
                                            v28 := GenS3{F0: 0, F1: (-2147483648 >= 1)}
                                            _ = v28
                                        }
                                        v29 := GenS2{F0: "", F1: false, F2: -179}
                                        v30 := GenS3{F0: 1, F1: ((((true || true) && true) && true) || false)}
                                        _ = v27
                                        _ = v29
                                        _ = v30
                                    } else {
                                        switch "hello" {
                                        case "a":
                                            switch 127 {
                                            case 725:
                                                v7 = (v2 * (^(v7)))
                                            default:
                                                v6 = (v6 << ((v7) & 63))
                                            }
                                        case "\n\t":
                                            fmt.Println(fi14)
                                        }
                                        if (9223372036854775807 < 549) {
                                            v7 = (^(v2))
                                        }
                                        v7 = (v2 / (((v2 - 2)) | 1))
                                        v31 := GenS2{F0: "", F1: false, F2: -1}
                                        _ = v31
                                    }
                                    if (225 < -564) {
                                        var v32 int16 = v6
                                        _ = v32
                                    } else {
                                        var v33 int32 = 1
                                        var v34 int16 = (v6 & v6)
                                        delete(v4, 980)
                                        var v35 int8 = 42
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                }
                                v6 -= 32767
                            }
                            _ = fi22
                            _ = v23
                        }
                        _ = v15
                    } else {
                        if false {
                            var v36 bool = false
                            _ = v36
                        }
                        v37 := GenS2{F0: ("" + "a"), F1: (true || (false || false)), F2: 0}
                        if false {
                            v6 = 38
                        }
                        _ = v37
                    }
                    _ = fi14
                }
            } else {
                var v38 uint64 = 229
                fmt.Println(v38, v2)
                _ = v38
            }
        case "foo bar":
            var v39 int = 32768
            _ = v39
        default:
            var v40 uint8 = 42
            _ = v40
        }
    default:
        for fi41 := 0; fi41 < 1; fi41++ {
            switch int(v6) {
            case 42:
                var v42 int = fi41
                _ = v42
            case 4294967295:
                v7 += 255
            }
            var v43 uint = v2
            for fi44 := 0; fi44 < 16; fi44++ {
                var v45 uint64 = 1
                var v46 int64 = -9223372036854775808
                var v47 float32 = (-2.25 + float32(0.0))
                _ = fi44
                _ = v45
                _ = v46
                _ = v47
            }
            _ = fi41
            _ = v43
        }
    }
    for fi48 := 0; fi48 < 8; fi48++ {
        var v49 int64 = -1
        switch fi48 {
        case 128:
            var v50 string = ""
            _ = v50
        case -1:
            v2 += v7
        case 1:
            var v51 uint32 = 2147483648
            _ = v51
        }
        _ = fi48
        _ = v49
    }
    var v52 int32 = 0
    if (!(1 > 1)) {
        for fi53 := 0; fi53 < 15; fi53++ {
            v54 := GenS5{F0: 0, F1: 820}
            _ = fi53
            _ = v54
        }
    }
    v55 := GenS6{F0: 255, F1: 42, F2: -1}
    v56 := v3["foo bar"]
    var v57 bool = ((!(false && false)) && true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v7
}
