package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 int32; F1 uint16; F2 int64 }
type GenS9 struct { F0 uint16 }

func main() {
    var v1 float32 = (1.5 / 1.0)
    var v2 int = 482
    v3 := make(map[int]int)
    if ((("\n\t" + "\n\t") >= (("" + "") + "a")) || ("a" != "foo bar")) {
        var v4 int8 = 1
        _ = v4
    } else {
        v2 = (-(v2))
        v5 := GenS1{F0: 42, F1: 1, F2: 51}
        if true {
            if false {
                v1 += v1
            } else {
                if true {
                    v6 := []int{-594, v2, (-(v2)), 2147483647}
                    delete(v3, v2)
                    _ = v6
                } else {
                    v7 := []int{v2, v2}
                    _ = v7
                }
                for fi8 := 0; fi8 < 3; fi8++ {
                    if true {
                        v1 += v1
                        var v9 uint = 65535
                        _ = v9
                    }
                    v10 := []int{(v2 % ((3) | 1)), (v2 << ((0) & 63))}
                    _ = len(v3)
                    _ = fi8
                    _ = v10
                }
            }
            v11 := GenS2{F0: 4294967295, F1: 0, F2: 61}
            v1 = (((0.0 * float32(1e10)) - v1) - float32(3.14159))
            _ = v11
        } else {
            for fi12 := 0; fi12 < 4; fi12++ {
                v1 += float32(1.5)
                if true {
                    var v13 uint16 = 213
                    v13 = 675
                    fmt.Println(v1, v2)
                    _ = v13
                } else {
                    var v14 uint16 = 827
                    _ = v14
                }
                _ = fi12
            }
            var v15 float64 = 3.14159
            var v16 int64 = 0
            _ = v15
            _ = v16
        }
        _ = v5
    }
    fmt.Println(v2, v1)
    if true {
        var v17 uint8 = 2
        for fi18 := 0; fi18 < 2; fi18++ {
            _ = len(v3)
            _ = fi18
        }
        v1 = (v1 + v1)
        _ = v17
    } else {
        var v19 uint64 = 704
        _ = v19
    }
    v20 := GenS3{F0: 42}
    v21 := []int{v2, 811, v2}
    v22 := GenS1{F0: 42, F1: 603, F2: 169}
    for fi23 := 0; fi23 < 13; fi23++ {
        v2 = 127
        for fi24 := 0; fi24 < 19; fi24++ {
            switch (fi24 + 32767) {
            case 336:
                var v25 int8 = -1
                _ = v25
            case -9223372036854775808:
                fmt.Println(v2)
            case 4294967295:
                v1 = v1
            }
            switch fi23 {
            case 71:
                if false {
                    _ = len(v21)
                    var v26 uint8 = 255
                    _ = v26
                } else {
                    fmt.Println(fi24, fi24)
                }
            case 0:
                fmt.Println(v2, v2)
            case 2147483647:
                v27 := GenS4{F0: 32767, F1: 76, F2: -425}
                if (!((1 < 1000) && false)) {
                    var v28 int8 = 127
                    switch v2 {
                    case 128:
                        fmt.Println(fi24)
                        if (123 >= v28) {
                            v29 := GenS4{F0: 727, F1: 33, F2: 1}
                            v30 := v3[(fi23 << ((255) & 63))]
                            v31 := GenS5{F0: 653}
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    case -492:
                        if true {
                            switch "foo bar" {
                            case "foo bar":
                                delete(v3, (fi24 * 508))
                            case "hello":
                                var v32 uint32 = 2147483648
                                _ = v32
                            case "":
                                v28 = v28
                            }
                            switch "hello" {
                            case "hello":
                                var v33 int32 = 750
                                _ = v33
                            case "foo bar":
                                v28 += v28
                            case "a":
                                v28 = (-(v28))
                            }
                            delete(v3, (fi24 ^ -601))
                            if true {
                                if true {
                                    var v34 int32 = -994
                                    var v35 uint16 = 42
                                    var v36 bool = ("a" != "a")
                                    v36 = ((!(1000 != 42)) && true)
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                } else {
                                    if (1.5 <= -2.25) {
                                        v37 := GenS1{F0: 1000, F1: 2, F2: 2}
                                        v38 := GenS2{F0: 0, F1: 1000, F2: 127}
                                        _ = v37
                                        _ = v38
                                    } else {
                                        if ((true || false) || (-787 < 9223372036854775807)) {
                                            switch fi23 {
                                            case 9223372036854775807:
                                                v39 := GenS1{F0: -616, F1: 182, F2: 1}
                                                _ = v39
                                            }
                                            v40 := GenS3{F0: 18446744073709551615}
                                            _ = v40
                                        } else {
                                            v2 -= (v2 << ((65535) & 63))
                                        }
                                    }
                                    if (962 == 1) {
                                        v41 := GenS6{F0: "a", F1: 391}
                                        var v42 string = ("\n\t" + "foo bar")
                                        _ = v41
                                        _ = v42
                                    } else {
                                        var v43 int = (v2 | (fi23 * -2147483648))
                                        _ = v43
                                    }
                                }
                            }
                            fmt.Println(v28, fi24)
                        }
                    case 449:
                        fmt.Println(v2, v2, v28)
                    default:
                        if false {
                            v44 := GenS2{F0: 65535, F1: 32767, F2: 0}
                            var v45 float64 = -2.25
                            _ = v44
                            _ = v45
                        }
                    }
                    _ = v28
                } else {
                    if true {
                        if (false && (("foo bar" + ("foo bar" + ("foo bar" + ""))) < "a")) {
                            v2 -= 550
                        } else {
                            _ = len(v21)
                            var v46 uint32 = 65535
                            _ = v46
                        }
                        var v47 int16 = 736
                        var v48 int16 = v47
                        if (((18446744073709551615 <= 288) || ((1 != 2147483648) && (1 == 1))) || true) {
                            v49 := GenS7{F0: 426, F1: 1}
                            _ = v49
                        }
                        switch "foo bar" {
                        case "a":
                            switch 127 {
                            case -1:
                                _ = len(v3)
                            case 9223372036854775807:
                                v50 := GenS8{F0: 42, F1: 482, F2: -1}
                                _ = v50
                            case 128:
                                v2 -= fi24
                            default:
                                var v51 string = ""
                                _ = v51
                            }
                        case "foo bar":
                            fmt.Println(fi24)
                        case "hello":
                            fmt.Println(fi24, fi23, v48)
                        }
                        fmt.Println(fi24)
                        _ = v47
                        _ = v48
                    }
                }
                var v52 uint8 = 1
                var v53 int64 = 85
                _ = v27
                _ = v52
                _ = v53
            }
            v1 += (((v1 + (v1 + v1)) / v1) - -2.25)
            _ = fi24
        }
        _ = fi23
    }
    v21 = append(v21, v2)
    fmt.Println(v2)
    v54 := GenS9{F0: 255}
    for fi55 := 0; fi55 < 11; fi55++ {
        fmt.Println(fi55, fi55)
        if true {
            var v56 string = ""
            _ = v56
        } else {
            v2 += fi55
            v57 := GenS2{F0: 2, F1: -51, F2: 255}
            for fi58 := 0; fi58 < 5; fi58++ {
                switch ("" + "") {
                case "a":
                    v1 += v1
                default:
                    var v59 float32 = (1.5 * 1.5)
                    _ = v59
                }
                _ = fi58
            }
            _ = v57
        }
        _ = fi55
    }
    var v60 bool = false
    if ((!true) && v60) {
        var v61 float32 = (float32(1.5) * v1)
        _ = v61
    }
    _ = len(v21)
    switch "foo bar" {
    case "\n\t":
        _ = len(v3)
    case "foo bar":
        v1 -= 1e10
    default:
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
    }
    for fi62 := 0; fi62 < 19; fi62++ {
        _ = len(v21)
        delete(v3, 431)
        v63 := GenS9{F0: 65535}
        _ = fi62
        _ = v63
    }
    if len(v21) > 0 {
        v21[0] = (v2 & (v2 << ((735) & 63)))
    }
    v2 -= 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v54
    _ = v60
}
