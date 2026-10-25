package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 uint32; F1 int8; F2 uint }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 int32; F1 int8; F2 int64 }
func genFn0(p0 uint8, p1 bool, p2 uint) uint8 {
    var v5 int64 = -1
    fmt.Println(p2, v5, p2)
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            p2 -= (p2 ^ p2)
            v8 := GenS1{F0: "\n\t", F1: 723}
            _ = fi7
            _ = v8
        }
        v5 = 0
        _ = fi6
    }
    if p1 {
        var v9 int64 = (v5 * -783)
        _ = v9
    } else {
        switch "\n\t" {
        case "":
            switch "" {
            case "\n\t":
                v10 := GenS1{F0: ("a" + "foo bar"), F1: 42}
                _ = v10
            case "":
                if true {
                    fmt.Println(p2, p1, p1)
                }
            }
        case "a":
            v11 := GenS1{F0: "a", F1: 0}
            _ = v11
        case "hello":
            var v12 int16 = 311
            _ = v12
        }
        v13 := []int{-477, 813, 526}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p0 << ((p2) & 63))
}
func genFn1(p0 int32) int16 {
    var v14 uint16 = 32767
    v15 := GenS2{F0: 976}
    fmt.Println(v14, v14)
    _ = p0
    _ = v14
    _ = v15
    return 514
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 0
    var v3 uint32 = 2
    v4 := make(map[uint32]uint32)
    fmt.Println(v2, v3)
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := v4[v3]
        _ = fi16
        _ = v17
    }
    v18 := GenS3{F0: 65535}
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS4{F0: ("a" + ""), F1: 986}
        v21 := GenS3{F0: 65535}
        v22 := GenS5{F0: 255}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 uint16 = (^(v2))
    if true {
        v24 := GenS6{F0: -2147483648, F1: 1}
        _ = v24
    } else {
        var v25 uint64 = 4294967295
        var v26 int16 = -304
        _ = v25
        _ = v26
    }
    v27 := []int{-544, 42, 128}
    switch -2147483648 {
    case 42:
        if true {
            if true {
                var v28 string = "hello"
                _ = v28
            } else {
                _ = len(v27)
            }
            switch 0 {
            case -9223372036854775808:
                if (!(-455 >= -2147483648)) {
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v30 := GenS3{F0: 255}
                        _ = fi29
                        _ = v30
                    }
                    v23 += (^(v23))
                    if true {
                        if (455 != 1) {
                            if (127 >= -128) {
                                v3 = 65535
                            }
                            fmt.Println(v23, v23)
                        } else {
                            switch -1 {
                            case -293:
                                switch int(v2) {
                                case 182:
                                    v31 := GenS7{F0: v3, F1: 42, F2: 1}
                                    _ = v31
                                }
                            case 546:
                                v3 = v3
                            case -1:
                                var v32 int8 = 127
                                _ = v32
                            default:
                                v4[v3] = (v3 >> ((18446744073709551615) & 63))
                            }
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            fmt.Println(v2, v3)
                        }
                        v33 := GenS8{F0: "hello", F1: true}
                        _ = v33
                    }
                } else {
                    v2 = 42
                    var v34 bool = (-991 >= -9223372036854775808)
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                    _ = v34
                }
            case 32767:
                for fi35 := 0; fi35 < 12; fi35++ {
                    fmt.Println(v2, v23, v23)
                    v3 = v3
                    _ = fi35
                }
            case 230:
                v3 = v3
            }
        } else {
            v36 := GenS5{F0: 605}
            v37 := GenS4{F0: ("" + ((("foo bar" + "a") + "") + ((("foo bar" + "") + "foo bar") + ("\n\t" + "a")))), F1: 0}
            _ = v36
            _ = v37
        }
    case 9223372036854775807:
        fmt.Println(v23, v23, v23)
    default:
        if true {
            delete(v1, (^(v3)))
            switch -899 {
            case 128:
                for fi38 := 0; fi38 < 13; fi38++ {
                    var v39 int8 = -128
                    v39 += (v39 / ((v39) | 1))
                    delete(v1, (v3 | v3))
                    _ = fi38
                    _ = v39
                }
            case 0:
                if false {
                    v40 := GenS5{F0: 255}
                    _ = v40
                }
            case -2147483648:
                v23 = (v2 + (v2 << ((2) & 63)))
            }
            v23 += 432
        }
    }
    v41 := GenS9{F0: 249, F1: -99, F2: 42}
    switch "a" {
    case "foo bar":
        switch 9223372036854775807 {
        case 441:
            var v42 int16 = 572
            _ = v42
        }
    case "":
        _ = len(v1)
    case "a":
        var v43 int32 = 486
        _ = v43
    }
    v23 = 975
    fmt.Println(v23, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v41
}
