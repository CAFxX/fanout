package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint64; F1 uint; F2 string }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int16; F1 uint; F2 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 582
    v3 := make(map[string]bool)
    v4 := GenS1{F0: 898, F1: 42, F2: -1}
    v2 = (v2 & (^(v2)))
    v2 = 4294967295
    fmt.Println(v2)
    v5 := GenS2{F0: false, F1: 9223372036854775807, F2: 559}
    fmt.Println(v2)
    v2 = (v2 >> ((18446744073709551615) & 63))
    switch 32768 {
    case -1:
        delete(v3, "hello")
    }
    switch -770 {
    case 888:
        var v6 int8 = 34
        _ = v6
    default:
        switch 32767 {
        case 128:
            delete(v1, 9223372036854775807)
        case 127:
            v7 := GenS3{F0: 42, F1: 100}
            _ = v7
        case 1:
            v8 := []int{}
            _ = v8
        default:
            v2 += v2
        }
    }
    var v9 float32 = -2.25
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint64 = (v2 / ((v2) | 1))
        _ = fi10
        _ = v11
    }
    if true {
        var v12 int8 = -128
        for fi13 := 0; fi13 < 1; fi13++ {
            if false {
                v14 := GenS2{F0: false, F1: -412, F2: 91}
                _ = v14
            }
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := GenS3{F0: 573, F1: -125}
                if false {
                    v17 := []int{(fi13 << ((839) & 63))}
                    _ = v17
                }
                var v18 int64 = 9223372036854775807
                _ = fi15
                _ = v16
                _ = v18
            }
            _ = fi13
        }
        var v19 float32 = v9
        _ = v12
        _ = v19
    } else {
        if false {
            v20 := []int{883}
            if (-85 < 127) {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                v2 = v2
            } else {
                var v21 float64 = 1.5
                if false {
                    v22 := GenS1{F0: 835, F1: 1, F2: 183}
                    fmt.Println(v9, v2, v21)
                    _ = len(v20)
                    _ = v22
                } else {
                    v23 := GenS4{F0: (v2 + v2), F1: 911, F2: ("foo bar" + (("hello" + "\n\t") + "foo bar"))}
                    fmt.Println(v21, v21, v21)
                    _ = v23
                }
                if true {
                    var v24 int32 = 646
                    var v25 uint16 = 0
                    _ = v24
                    _ = v25
                } else {
                    v3[("foo bar" + "hello")] = (0 != 0)
                    switch "a" {
                    case "":
                        if true {
                            for fi26 := 0; fi26 < 20; fi26++ {
                                v21 = 1.5
                                _ = fi26
                            }
                            switch 4294967295 {
                            case 2147483647:
                                var v27 uint32 = 1
                                _ = v27
                            case 32767:
                                v20 = append(v20, 32768)
                            default:
                                if true {
                                    if (v21 < (v21 / 1.0)) {
                                        if (9223372036854775807 > 625) {
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v21, v21)
                                            default:
                                                v2 += (v2 / ((v2) | 1))
                                            }
                                            v28 := GenS1{F0: 127, F1: 0, F2: -9223372036854775808}
                                            _ = v28
                                        } else {
                                            if (-9223372036854775808 >= -9223372036854775808) {
                                                v9 = v9
                                            } else {
                                                if (413 < -1) {
                                                    var v29 float64 = 1e10
                                                    if (("\n\t" + "\n\t") < "hello") {
                                                        var v30 float32 = v9
                                                        fmt.Println(v21)
                                                        _ = v30
                                                    } else {
                                                        fmt.Println(v9, v2, v21)
                                                    }
                                                    switch 978 {
                                                    case -1:
                                                        var v31 int8 = -128
                                                        _ = v31
                                                    case 0:
                                                        fmt.Println(v21, v2, v21)
                                                    }
                                                    _ = v29
                                                } else {
                                                    _ = len(v1)
                                                    if (2147483647 < 2147483647) {
                                                        if false {
                                                            v32 := GenS4{F0: (v2 ^ v2), F1: 42, F2: ""}
                                                            _ = v32
                                                        }
                                                        var v33 uint64 = v2
                                                        var v34 int32 = -1
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        v20 = append(v20, 32767)
                                                    }
                                                    if false {
                                                        for fi35 := 0; fi35 < 12; fi35++ {
                                                            v36 := GenS5{F0: 189, F1: (v2 | v2)}
                                                            delete(v3, "\n\t")
                                                            _ = fi35
                                                            _ = v36
                                                        }
                                                    } else {
                                                        v2 += 2
                                                        fmt.Println(v9, v21, v9)
                                                    }
                                                }
                                                if (2147483647 != 599) {
                                                    _ = len(v20)
                                                    v21 -= v21
                                                    if false {
                                                        for fi37 := 0; fi37 < 13; fi37++ {
                                                            var v38 float64 = float64(1e10)
                                                            var v39 uint = 4294967295
                                                            _ = fi37
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        v40 := GenS3{F0: 42, F1: 100}
                                                        _ = v40
                                                    }
                                                } else {
                                                    _ = len(v1)
                                                    var v41 int64 = 9223372036854775807
                                                    var v42 uint32 = 1
                                                    _ = v41
                                                    _ = v42
                                                }
                                            }
                                        }
                                        if true {
                                            fmt.Println(v21, v21)
                                        }
                                        var v43 int = 32768
                                        _ = v43
                                    } else {
                                        if (127 < 818) {
                                            v44 := GenS6{F0: (("a" + ("a" + "foo bar")) + ((("\n\t" + "hello") + "\n\t") + "hello")), F1: true}
                                            v45 := GenS7{F0: 223, F1: 9223372036854775807}
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                    var v46 int64 = 1
                                    v47 := GenS3{F0: 0, F1: 61}
                                    _ = v46
                                    _ = v47
                                } else {
                                    var v48 uint16 = 32767
                                    v49 := GenS8{F0: -1, F1: 858, F2: -1}
                                    if (!true) {
                                        var v50 int = -1
                                        v48 = v48
                                        v51 := GenS6{F0: "a", F1: false}
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = v48
                                    _ = v49
                                }
                            }
                            var v52 uint32 = 551
                            _ = v52
                        }
                    case "a":
                        var v53 float32 = v9
                        _ = v53
                    default:
                        v20 = append(v20, 32768)
                    }
                }
                _ = v21
            }
            switch "foo bar" {
            case "hello":
                var v54 int8 = 42
                _ = v54
            case "\n\t":
                if (!(!(false && (0.0 <= 1e10)))) {
                    _ = len(v3)
                    var v55 float32 = -2.25
                    v1[-2147483648] = 0
                    _ = v55
                }
            }
            _ = v20
        } else {
            v56 := GenS8{F0: -555, F1: 0, F2: 9223372036854775807}
            fmt.Println(v9)
            var v57 int8 = -128
            _ = v56
            _ = v57
        }
        v2 = 18446744073709551615
        var v58 uint = 692
        _ = v58
    }
    fmt.Println(v9, v9)
    var v59 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v59:", v59)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v9
}
