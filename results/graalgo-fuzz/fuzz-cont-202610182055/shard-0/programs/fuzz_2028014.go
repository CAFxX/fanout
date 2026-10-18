package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint; F1 uint8; F2 string }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 int32; F1 uint16; F2 int }
type GenS9 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 float32 = (0.1 - (0.1 + 3.14159))
    var v2 uint32 = 4294967295
    var v3 uint = 255
    var v4 float64 = (3.14159 + (-2.25 * float64(((v1 * (3.14159 * v1)) + v1))))
    var v5 string = ("a" + ((("\n\t" + ("\n\t" + "foo bar")) + ("foo bar" + "hello")) + "foo bar"))
    var v6 uint8 = 2
    var v7 uint8 = (v6 / ((0) | 1))
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v1, v7)
        v9 := GenS1{F0: 42, F1: 65535, F2: -500}
        v3 += v3
        _ = fi8
        _ = v9
    }
    var v10 int = -731
    var v11 uint64 = 1
    var v12 string = (((v5 + "") + ("foo bar" + v5)) + v5)
    var v13 string = ""
    switch v5 {
    case "hello":
        v14 := GenS2{F0: -128}
        _ = v14
    case "foo bar":
        switch v5 {
        case "foo bar":
            var v15 string = (v13 + v12)
            _ = v15
        case "hello":
            for fi16 := 0; fi16 < 13; fi16++ {
                v10 += 42
                _ = fi16
            }
        case "\n\t":
            var v17 int = (^(v10))
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint32 = (v2 & 42)
        var v20 string = "\n\t"
        v12 = v13
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        if false {
            switch 1 {
            case 4294967295:
                fmt.Println(v5, v7, v2)
            case -9223372036854775808:
                v22 := []int{}
                _ = v22
            case 42:
                v23 := []int{(-(fi21))}
                _ = v23
            }
            v24 := GenS1{F0: 42, F1: (v3 / ((v3) | 1)), F2: 74}
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(fi25)
                var v26 int8 = 42
                if (v13 < (v5 + (v13 + ""))) {
                    fmt.Println(v1)
                    v5 = v13
                    if (-824 >= 0) {
                        switch (-(v10)) {
                        case -1:
                            v27 := GenS2{F0: (^(v26))}
                            _ = v27
                        case -537:
                            v7 = (v6 / (((v7 / ((uint8(v3)) | 1))) | 1))
                        default:
                            switch v5 {
                            case "":
                                v28 := GenS3{F0: -1}
                                _ = v28
                            default:
                                switch ((v12 + v12) + ((v12 + v12) + v13)) {
                                case "":
                                    v26 = (v26 - int8(fi25))
                                    switch -852 {
                                    case 2147483647:
                                        v5 = ((v12 + v13) + v5)
                                    case 127:
                                        v10 = 32767
                                    case 32767:
                                        if (v2 == (^(v2))) {
                                            v6 = (v7 >> ((702) & 63))
                                            if false {
                                                var v29 uint32 = v2
                                                _ = v29
                                            } else {
                                                switch ((v5 + (v12 + "foo bar")) + v13) {
                                                case "":
                                                    var v30 uint8 = v7
                                                    _ = v30
                                                case "\n\t":
                                                    if (false || (-328 <= -5)) {
                                                        v31 := []int{}
                                                        v32 := GenS4{F0: 2147483647}
                                                        v1 = v1
                                                        var v33 int16 = 1000
                                                        var v34 int32 = 0
                                                        _ = v31
                                                        _ = v32
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        var v35 uint32 = uint32(v3)
                                                        var v36 int8 = (v26 / (((-(v26))) | 1))
                                                        fmt.Println(v26, v36, v6)
                                                        v36 -= v36
                                                        if false {
                                                            v37 := GenS5{F0: 1, F1: (v11 + (^(v11)))}
                                                            var v38 uint64 = 2
                                                            var v39 bool = false
                                                            _ = v37
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                }
                                            }
                                            v40 := GenS4{F0: 42}
                                            _ = v40
                                        }
                                    default:
                                        switch (v10 << ((242) & 63)) {
                                        case 127:
                                            v41 := GenS6{F0: v3, F1: (v7 << ((v3) & 63)), F2: v12}
                                            _ = v41
                                        case 9223372036854775807:
                                            fmt.Println(fi25)
                                        default:
                                            fmt.Println(v4)
                                        }
                                    }
                                    switch (v10 - 127) {
                                    case 127:
                                        fmt.Println(v5, v2, v2)
                                    case -829:
                                        switch "hello" {
                                        case "foo bar":
                                            if (!(((0.1 + (1e10 + (((0.0 / v1) * v1) * v1))) != v1) || ((v10 | 624) != v10))) {
                                                var v42 int8 = 5
                                                var v43 string = v13
                                                switch "a" {
                                                case "foo bar":
                                                    if true {
                                                        v6 = (v7 - 116)
                                                    } else {
                                                        switch int(v7) {
                                                        case 637:
                                                            var v44 bool = false
                                                            _ = v44
                                                        case -2147483648:
                                                            v26 = v26
                                                        case -314:
                                                            v45 := []int{(v10 << ((uint(v2)) & 63)), fi21, (fi25 / ((fi25) | 1)), (v10 * fi21)}
                                                            _ = v45
                                                        default:
                                                            v26 -= v42
                                                        }
                                                    }
                                                case "hello":
                                                    switch (v12 + "a") {
                                                    case "foo bar":
                                                        if false {
                                                            v46 := []int{-1, int(v11), 0}
                                                            v1 = ((float32((float64(v1) - v4)) + (v1 - v1)) - (v1 / v1))
                                                            _ = v46
                                                        }
                                                    case "hello":
                                                        var v47 uint32 = 0
                                                        _ = v47
                                                    default:
                                                        if (9223372036854775807 < -9223372036854775808) {
                                                            v48 := GenS7{F0: 127, F1: ((v3 * v3) <= v3)}
                                                            fmt.Println(v5)
                                                            v12 = ("a" + (v5 + "hello"))
                                                            var v49 uint16 = 255
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v3)
                                                }
                                                _ = v42
                                                _ = v43
                                            }
                                        case "\n\t":
                                            var v50 uint8 = (v6 >> ((165) & 63))
                                            _ = v50
                                        default:
                                            if (0 == -814) {
                                                fmt.Println(v4, v26, v4)
                                            }
                                        }
                                    case 0:
                                        switch -598 {
                                        case 42:
                                            v51 := []int{-517, 32768, int(v2), fi25}
                                            switch ("hello" + (((("hello" + v13) + "") + v12) + (("foo bar" + ("a" + v13)) + "\n\t"))) {
                                            case "\n\t":
                                                v52 := GenS8{F0: 2147483647, F1: 1, F2: 944}
                                                _ = v52
                                            case "hello":
                                                var v53 int8 = (v26 % (((v26 - (v26 % (((v26 * 92)) | 1)))) | 1))
                                                _ = v53
                                            case "a":
                                                fmt.Println(v5, v10, v26)
                                            }
                                            fmt.Println(v11, v13, v3)
                                            _ = v51
                                        case 218:
                                            v54 := []int{127, (fi21 >> (((^(v3))) & 63)), fi21}
                                            _ = v54
                                        }
                                    }
                                case "\n\t":
                                    var v55 int = -579
                                    _ = v55
                                }
                            }
                        }
                    } else {
                        v2 = (v2 << (((v3 ^ v3)) & 63))
                        var v56 uint16 = 987
                        if (true || true) {
                            v57 := GenS2{F0: v26}
                            _ = v57
                        }
                        v58 := GenS7{F0: 94, F1: (!(false && (v3 == uint(v6))))}
                        v12 = (v5 + ((v13 + "") + ((v5 + v5) + v13)))
                        _ = v56
                        _ = v58
                    }
                }
                _ = fi25
                _ = v26
            }
            _ = v24
        }
        var v59 float32 = v1
        v60 := GenS9{F0: 42, F1: 32767}
        _ = fi21
        _ = v59
        _ = v60
    }
    if false {
        var v61 uint64 = (v11 * v11)
        var v62 float64 = float64(3.14159)
        _ = v61
        _ = v62
    } else {
        for fi63 := 0; fi63 < 7; fi63++ {
            if ((((v1 * v1) + v1) < 1e10) || ((-1 < (fi63 * (-(v10)))) && false)) {
                v3 += v3
                v10 = v10
            } else {
                var v64 string = ""
                _ = v64
            }
            v12 = v12
            _ = fi63
        }
        v65 := []int{(v10 & (v10 >> ((v3) & 63))), (^(v10))}
        _ = v65
    }
    v66 := []int{}
    if len(v66) > 0 {
        v66[0] = v10
    }
    var v67 uint16 = 255
    _ = len(v66)
    v10 = (v10 * (-(v10)))
    var v68 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v7
}
