package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int8 }

func main() {
    var v1 int32 = 0
    var v2 int32 = -2147483648
    v3 := make(map[int]string)
    var v4 uint32 = 0
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]uint32)
    var v7 uint64 = 1
    if (!false) {
        var v8 uint8 = 222
        v9 := []int{32768, 734, -1}
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 65535}
        fmt.Println(v2)
        for fi11 := 0; fi11 < 10; fi11++ {
            if false {
                if (v4 <= (v4 >> ((293) & 63))) {
                    v12 := []int{(fi11 & fi11), fi11}
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v2 += 1
                        _ = fi13
                    }
                    _ = v12
                } else {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        var v15 int64 = 42
                        if ((true || false) || false) {
                            v16 := GenS1{F0: 65535}
                            _ = v16
                        }
                        _ = fi14
                        _ = v15
                    }
                    if false {
                        v7 += (v7 >> ((4294967295) & 63))
                    }
                }
                v17 := GenS2{F0: 469, F1: (-(fi11))}
                _ = v17
            } else {
                v4 += (^(v4))
                v2 = v1
                v1 = v2
            }
            _ = fi11
        }
        _ = v10
    }
    fmt.Println(v2, v7, v2)
    fmt.Println(v4)
    for fi18 := 0; fi18 < 14; fi18++ {
        v4 = (^(v4))
        for fi19 := 0; fi19 < 5; fi19++ {
            switch "foo bar" {
            case "":
                var v20 float64 = 3.14159
                _ = v20
            case "hello":
                var v21 int16 = 0
                _ = v21
            default:
                if true {
                    v22 := GenS1{F0: 65535}
                    _ = v22
                } else {
                    v23 := []int{fi18, 128, (fi19 | 1), 4294967295}
                    v5[v4] = 42
                    _ = v23
                }
            }
            _ = fi19
        }
        switch -1 {
        case 620:
            v1 = (v2 / ((1) | 1))
        case 411:
            v24 := GenS1{F0: 735}
            _ = v24
        case 601:
            for fi25 := 0; fi25 < 10; fi25++ {
                if (!true) {
                    v7 += v7
                    delete(v5, (^(v4)))
                    if true {
                        v26 := GenS2{F0: fi25, F1: fi25}
                        v4 = (^(v4))
                        _ = v26
                    } else {
                        v4 -= (v4 - v4)
                        v27 := []int{}
                        switch (fi25 >> ((2) & 63)) {
                        case -9223372036854775808:
                            v2 -= 1
                        case 1:
                            v4 -= (v4 >> ((736) & 63))
                        case -716:
                            v28 := GenS3{F0: -1}
                            _ = v28
                        default:
                            fmt.Println(v7)
                        }
                        _ = v27
                    }
                    v29 := []int{fi18, fi25}
                    _ = v29
                } else {
                    var v30 uint = 18446744073709551615
                    _ = v30
                }
                _ = fi25
            }
        default:
            _ = len(v6)
        }
        _ = fi18
    }
    fmt.Println(v7, v2)
    switch -2147483648 {
    case 32767:
        v31 := GenS4{F0: -977, F1: 0}
        _ = v31
    case 32768:
        v32 := []int{32767, 1}
        _ = v32
    }
    v2 = (v2 / (((v2 << ((4294967295) & 63))) | 1))
    fmt.Println(v4)
    if false {
        if true {
            if false {
                switch "hello" {
                case "hello":
                    if false {
                        var v33 uint32 = v4
                        _ = v33
                    }
                }
                if false {
                    v4 = (v4 / (((v4 % ((0) | 1))) | 1))
                }
            }
            for fi34 := 0; fi34 < 19; fi34++ {
                fmt.Println(fi34, v7)
                v1 -= 0
                switch "\n\t" {
                case "":
                    for fi35 := 0; fi35 < 4; fi35++ {
                        v1 = (v2 ^ (v2 | -2147483648))
                        if (v4 < (v4 & (v4 & v4))) {
                            fmt.Println(v4, v2, fi35)
                        } else {
                            v36 := GenS1{F0: 2}
                            v37 := []int{fi35, fi35}
                            _ = v36
                            _ = v37
                        }
                        _ = fi35
                    }
                }
                _ = fi34
            }
        }
        var v38 int64 = 773
        v39 := []int{}
        _ = v38
        _ = v39
    } else {
        if true {
            var v40 string = "a"
            var v41 uint16 = 32767
            _ = v40
            _ = v41
        }
    }
    var v42 uint64 = (v7 % ((127) | 1))
    if (-72 != -95) {
        v42 = v7
    } else {
        switch -9223372036854775808 {
        case 128:
            for fi43 := 0; fi43 < 18; fi43++ {
                v3[fi43] = ("hello" + ((("hello" + "") + "a") + "foo bar"))
                v42 -= (^(v7))
                var v44 int64 = -1
                _ = fi43
                _ = v44
            }
        case 32768:
            fmt.Println(v4)
        }
        for fi45 := 0; fi45 < 20; fi45++ {
            var v46 uint32 = (v4 - (v4 - 792))
            v4 -= 65535
            _ = fi45
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        v48 := GenS2{F0: fi47, F1: (fi47 >> ((760) & 63))}
        switch (("" + "hello") + "") {
        case "\n\t":
            _ = len(v3)
        case "hello":
            _ = len(v5)
        case "foo bar":
            var v49 uint32 = 179
            _ = v49
        default:
            var v50 uint8 = 128
            _ = v50
        }
        _ = fi47
        _ = v48
    }
    fmt.Println(v4, v7)
    for fi51 := 0; fi51 < 3; fi51++ {
        v52 := GenS2{F0: (fi51 | 9223372036854775807), F1: (^(fi51))}
        for fi53 := 0; fi53 < 5; fi53++ {
            var v54 uint32 = v4
            if (true && (!true)) {
                switch ("" + "\n\t") {
                case "hello":
                    switch ("a" + "hello") {
                    case "":
                        v42 = v7
                        var v55 int32 = 0
                        if true {
                            var v56 int32 = (v2 | -982)
                            _ = v56
                        } else {
                            var v57 int64 = 9223372036854775807
                            if (!false) {
                                fmt.Println(fi53, v1)
                            } else {
                                switch (("foo bar" + "foo bar") + ("a" + "a")) {
                                case "foo bar":
                                    v42 = v7
                                case "":
                                    v2 += v2
                                }
                                if true {
                                    var v58 uint = 0
                                    v59 := GenS1{F0: 769}
                                    _ = v58
                                    _ = v59
                                } else {
                                    if ((v57 <= (v57 % (((-(v57))) | 1))) || false) {
                                        if (true && (4294967295 != fi53)) {
                                            v57 += (v57 & v57)
                                            switch (fi51 ^ (fi51 / (((fi51 % (((fi51 * fi51)) | 1))) | 1))) {
                                            case 42:
                                                fmt.Println(v55)
                                            case 1:
                                                switch (fi53 >> ((438) & 63)) {
                                                case 367:
                                                    v5[(^(v54))] = 404
                                                case -2147483648:
                                                    v6["foo bar"] = v54
                                                }
                                            case 127:
                                                v60 := []int{654, (-(fi53)), fi51, 2147483647}
                                                _ = v60
                                            default:
                                                v54 = 1
                                            }
                                            v4 += (v54 >> ((1) & 63))
                                        }
                                    }
                                }
                                var v61 float64 = 0.0
                                var v62 int = 32768
                                v4 += v54
                                _ = v61
                                _ = v62
                            }
                            var v63 int64 = v57
                            var v64 int8 = -1
                            v65 := GenS1{F0: 32767}
                            _ = v57
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        _ = v55
                    }
                case "a":
                    var v66 int64 = 1
                    _ = v66
                }
            }
            _ = fi53
            _ = v54
        }
        _ = len(v6)
        _ = fi51
        _ = v52
    }
    switch "\n\t" {
    case "a":
        if (((float32((float64(0.0) + -2.25)) / (3.14159 * 1.5)) < float32(0.1)) && (true && false)) {
            v2 = v1
        }
    case "hello":
        switch "" {
        case "":
            if (false || ((4294967295 <= 255) || (127 <= 127))) {
                v67 := GenS3{F0: -16}
                v68 := []int{-273, 482, 42, 9223372036854775807}
                var v69 string = ""
                _ = v67
                _ = v68
                _ = v69
            }
        default:
            v1 = -976
        }
    case "foo bar":
        v2 = (v1 - (v1 | -1))
    default:
        v2 += -1
    }
    var v70 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v70
}
