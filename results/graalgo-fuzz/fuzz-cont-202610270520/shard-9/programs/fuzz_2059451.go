package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint }
type GenS6 struct { F0 int; F1 bool; F2 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 uint8 = 185
    v4 := make(map[uint32]bool)
    var v5 uint8 = 2
    if true {
        switch (("a" + ((("\n\t" + "hello") + "") + "a")) + ("foo bar" + "hello")) {
        case "a":
            v5 += 1
        case "foo bar":
            var v6 float64 = 1.5
            _ = v6
        case "hello":
            v5 = v3
        default:
            v7 := []int{-241}
            _ = v7
        }
        switch -9223372036854775808 {
        case -756:
            var v8 uint16 = 932
            _ = v8
        case 882:
            v9 := GenS1{F0: 1, F1: 404}
            _ = v9
        case 617:
            var v10 int64 = 1
            _ = v10
        }
    }
    v3 -= (v5 / ((v5) | 1))
    v11 := GenS2{F0: 127, F1: 421, F2: 65535}
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS3{F0: 0, F1: -1}
        v5 = 128
        v14 := GenS1{F0: (^(fi12)), F1: 1}
        _ = fi12
        _ = v13
        _ = v14
    }
    switch "foo bar" {
    case "":
        if (-25 == 1) {
            var v15 int64 = 9223372036854775807
            _ = v15
        }
    case "foo bar":
        if (("foo bar" + ((("hello" + (("\n\t" + ("\n\t" + "hello")) + "")) + "") + "a")) == "hello") {
            for fi16 := 0; fi16 < 15; fi16++ {
                _ = len(v4)
                _ = fi16
            }
            v3 = 128
        } else {
            if (-1 != int32(v5)) {
                v3 = (v5 ^ v3)
            }
            if false {
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 float32 = (0.0 - 1e10)
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 uint = 2
                        if false {
                            v21 := GenS3{F0: uint64(v3), F1: 2147483647}
                            _ = v21
                        }
                        fmt.Println(v20, fi17)
                        _ = fi19
                        _ = v20
                    }
                    if true {
                        fmt.Println(v3)
                    } else {
                        switch -400 {
                        case 9223372036854775807:
                            if (4294967295 >= 103) {
                                delete(v1, (((("foo bar" + ("" + "")) + "a") + "") + "\n\t"))
                            }
                        case 1:
                            v22 := []int{fi17, fi17, -812, -337}
                            _ = v22
                        case -2147483648:
                            switch (((("foo bar" + "foo bar") + "foo bar") + "hello") + "a") {
                            case "foo bar":
                                fmt.Println(v3, v18, v18)
                            case "hello":
                                v23 := GenS3{F0: 850, F1: -1}
                                _ = v23
                            case "":
                                fmt.Println(v5, v3)
                            default:
                                var v24 uint8 = (v3 % ((v3) | 1))
                                _ = v24
                            }
                        }
                    }
                    _ = fi17
                    _ = v18
                }
            }
            var v25 bool = false
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        if (0 != 42) {
            switch fi26 {
            case 322:
                var v27 int8 = 127
                _ = v27
            }
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS1{F0: fi26, F1: 4294967295}
                _ = len(v2)
                fmt.Println(fi26, fi28, v5)
                var v30 uint32 = uint32(v5)
                v31 := []int{}
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            v3 -= v5
        } else {
            if false {
                _ = len(v1)
            } else {
                var v32 uint = 255
                _ = len(v1)
                _ = v32
            }
            var v33 int16 = 42
            if ((!true) && (1 >= 1)) {
                v33 += (v33 + (-(v33)))
                for fi34 := 0; fi34 < 11; fi34++ {
                    v35 := GenS1{F0: (-(fi26)), F1: 18446744073709551615}
                    _ = fi34
                    _ = v35
                }
            } else {
                switch fi26 {
                case 953:
                    var v36 uint = 42
                    _ = v36
                default:
                    var v37 float32 = (1e10 / 1.5)
                    _ = v37
                }
            }
            _ = v33
        }
        if ((true && (!false)) && true) {
            v5 -= 128
        } else {
            switch fi26 {
            case -9223372036854775808:
                v3 = v5
            case -2147483648:
                v3 -= (v3 | 11)
            default:
                fmt.Println(v5, v3)
            }
            v1["hello"] = 955
        }
        _ = fi26
    }
    if true {
        if (42 > 1) {
            for fi38 := 0; fi38 < 8; fi38++ {
                if true {
                    var v39 float32 = (-2.25 + (3.14159 - ((0.0 + -2.25) + 1e10)))
                    for fi40 := 0; fi40 < 13; fi40++ {
                        v39 = (1e10 / 1.5)
                        v41 := GenS4{F0: 2147483647, F1: 42}
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                }
                v42 := GenS5{F0: (-9223372036854775808 > 735), F1: 717, F2: 271}
                for fi43 := 0; fi43 < 7; fi43++ {
                    v3 -= v3
                    if false {
                        v44 := []int{(fi38 + 32767)}
                        if (false && (42 != 1)) {
                            if (false && true) {
                                delete(v4, 784)
                            } else {
                                if len(v44) >= 2 {
                                    v44 = v44[1:len(v44)]
                                }
                            }
                            fmt.Println(v3)
                        } else {
                            v45 := GenS4{F0: 2147483647, F1: v5}
                            var v46 string = "\n\t"
                            v3 = v5
                            _ = v45
                            _ = v46
                        }
                        v3 = v5
                        _ = v44
                    } else {
                        v47 := GenS6{F0: 4294967295, F1: true, F2: 4294967295}
                        v3 = (v3 << ((18446744073709551615) & 63))
                        if (fi38 < -1) {
                            if ((-(fi43)) <= -2147483648) {
                                v5 = v3
                                var v48 string = "a"
                                fmt.Println(v48, v48)
                                _ = v48
                            }
                            var v49 int = fi43
                            fmt.Println(v3)
                            _ = v49
                        }
                        _ = v47
                    }
                    var v50 uint64 = uint64(v5)
                    _ = fi43
                    _ = v50
                }
                _ = fi38
                _ = v42
            }
        }
        v3 = (v5 * (v3 >> ((18446744073709551615) & 63)))
    } else {
        switch -1 {
        case 1:
            v2["\n\t"] = 0
        default:
            switch -9223372036854775808 {
            case 127:
                fmt.Println(v3)
            case 42:
                var v51 bool = (790 <= 1)
                _ = v51
            case 32768:
                var v52 int8 = 0
                _ = v52
            }
        }
        switch ("hello" + "foo bar") {
        case "\n\t":
            for fi53 := 0; fi53 < 15; fi53++ {
                v54 := GenS6{F0: 2147483647, F1: false, F2: 65535}
                _ = fi53
                _ = v54
            }
        default:
            var v55 int16 = -757
            _ = v55
        }
        if (!(1 != 292)) {
            v5 += v5
            var v56 int16 = 1
            _ = v56
        }
    }
    var v57 int8 = -128
    if true {
        v3 = 128
        v58 := GenS4{F0: 0, F1: (v3 * (v5 * (v3 ^ v5)))}
        _ = v58
    }
    if false {
        if false {
            v59 := GenS5{F0: true, F1: 42, F2: 124}
            _ = len(v2)
            _ = v59
        } else {
            if (-1 == -1) {
                var v60 bool = true
                v61 := GenS2{F0: 2, F1: 4294967295, F2: 42}
                fmt.Println(v60)
                _ = v60
                _ = v61
            } else {
                for fi62 := 0; fi62 < 6; fi62++ {
                    v63 := GenS7{F0: ""}
                    _ = fi62
                    _ = v63
                }
                fmt.Println(v3, v57, v57)
                switch 4294967295 {
                case 0:
                    v64 := GenS8{F0: 42}
                    _ = v64
                default:
                    switch 42 {
                    case 32767:
                        switch -9223372036854775808 {
                        case 32767:
                            for fi65 := 0; fi65 < 17; fi65++ {
                                v66 := v1[("" + "\n\t")]
                                _ = fi65
                                _ = v66
                            }
                        case 128:
                            v67 := []int{42, 32768, 4294967295, 32767}
                            _ = v67
                        }
                    }
                }
            }
            v5 = (^(v5))
        }
    } else {
        v68 := GenS2{F0: (v3 & v3), F1: 4294967295, F2: 0}
        for fi69 := 0; fi69 < 14; fi69++ {
            if false {
                var v70 bool = (!(1.5 > float64(-2.25)))
                switch fi69 {
                case 128:
                    var v71 int32 = 0
                    _ = v71
                }
                delete(v2, ("foo bar" + "\n\t"))
                _ = v70
            } else {
                v5 = (v3 % ((150) | 1))
            }
            _ = fi69
        }
        _ = v68
    }
    if (3.14159 != -2.25) {
        var v72 int16 = -17
        v3 = 2
        v3 -= 128
        _ = v72
    }
    var v73 int8 = (v57 & 0)
    v73 = (v57 - (v73 ^ v57))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v73
}
