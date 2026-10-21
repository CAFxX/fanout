package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 uint; F2 bool }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint16, p1 int32) uint {
    var v5 uint32 = 238
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int16 = 598
        p0 = p0
        v8 := GenS1{F0: true, F1: fi6, F2: 203}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 bool = false
    v10 := []int{9223372036854775807, 32767}
    p1 = (p1 - -831)
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v9
    return 243
}
func genFn1(p0 string) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return (65535 != 255)
}

func main() {
    v1 := []uint32{4294967295}
    var v2 int64 = 42
    var v3 uint = 797
    var v4 int16 = -1
    v3 = genFn0(678, 0)
    if (!(100 < 76)) {
        if true {
            var v11 bool = false
            _ = v11
        }
        var v12 int32 = 42
        fmt.Println(v3, v4)
        _ = v12
    } else {
        var v13 bool = true
        switch -927 {
        case -353:
            switch 128 {
            case 1:
                for fi14 := 0; fi14 < 3; fi14++ {
                    switch fi14 {
                    case 127:
                        for fi15 := 0; fi15 < 20; fi15++ {
                            var v16 uint16 = 2
                            _ = fi15
                            _ = v16
                        }
                    }
                    _ = fi14
                }
            case 32767:
                if (("hello" + "hello") >= ("" + "foo bar")) {
                    if ((v13 && v13) || (v13 && v13)) {
                        if true {
                            fmt.Println(v2)
                            v17 := GenS2{F0: 32768}
                            _ = v17
                        }
                    } else {
                        var v18 bool = false
                        v19 := GenS2{F0: 0}
                        _ = v18
                        _ = v19
                    }
                    var v20 int32 = 0
                    v1 = append(v1, 1)
                    _ = v20
                } else {
                    var v21 float32 = (-2.25 * 1.5)
                    _ = v21
                }
            case 2147483647:
                if false {
                    v2 += -9223372036854775808
                    var v22 int64 = (-(v2))
                    _ = v22
                }
            default:
                v2 -= -9223372036854775808
            }
        case 4294967295:
            if len(v1) > 0 {
                v1[0] = 42
            }
        case 975:
            v4 += (v4 & v4)
        }
        switch 0 {
        case 32767:
            var v23 int32 = 229
            _ = v23
        case 127:
            switch "\n\t" {
            case "hello":
                switch 1 {
                case 128:
                    v4 += v4
                case 42:
                    for fi24 := 0; fi24 < 9; fi24++ {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            var v26 int8 = -1
                            _ = fi25
                            _ = v26
                        }
                        v2 -= (v2 - 1)
                        _ = fi24
                    }
                }
            case "a":
                _ = len(v1)
            default:
                v27 := GenS3{F0: 255, F1: 2147483647}
                _ = v27
            }
        case -667:
            if false {
                fmt.Println(v13, v3)
                for fi28 := 0; fi28 < 7; fi28++ {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v30 := GenS4{F0: 0}
                        _ = fi29
                        _ = v30
                    }
                    switch "\n\t" {
                    case "hello":
                        switch fi28 {
                        case 32767:
                            v2 -= 9223372036854775807
                        case 808:
                            switch "\n\t" {
                            case "\n\t":
                                for fi31 := 0; fi31 < 7; fi31++ {
                                    fmt.Println(v4, v3)
                                    _ = fi31
                                }
                            case "a":
                                var v32 string = ""
                                _ = v32
                            }
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case "\n\t":
                        var v33 uint = genFn0(1, -1)
                        _ = v33
                    case "":
                        fmt.Println(v3, v13)
                    default:
                        v2 = v2
                    }
                    v2 += v2
                    _ = fi28
                }
            } else {
                if v13 {
                    var v34 uint16 = 988
                    var v35 uint8 = 127
                    _ = v34
                    _ = v35
                }
            }
        }
        _ = v13
    }
    v36 := GenS5{F0: 1000, F1: v3, F2: (65535 >= 303)}
    v37 := GenS6{F0: v3, F1: 4294967295}
    switch ((("foo bar" + "foo bar") + "") + "") {
    case "foo bar":
        var v38 uint64 = 218
        _ = v38
    default:
        v1 = append(v1, 4294967295)
    }
    if false {
        _ = len(v1)
    } else {
        var v39 uint8 = 1
        for fi40 := 0; fi40 < 4; fi40++ {
            var v41 string = "hello"
            var v42 int8 = 73
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = v39
    }
    var v43 uint8 = 212
    v43 -= 255
    for fi44 := 0; fi44 < 15; fi44++ {
        switch "a" {
        case "":
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        default:
            v43 += (v43 - v43)
        }
        if true {
            var v45 int32 = -1
            fmt.Println(fi44, v4, v45)
            v2 = v2
            _ = v45
        } else {
            var v46 bool = (!(32767 >= v4))
            v4 += v4
            switch "\n\t" {
            case "a":
                var v47 uint64 = 42
                _ = v47
            case "foo bar":
                v48 := GenS7{F0: v43, F1: int16(fi44)}
                _ = v48
            case "":
                if (v46 || (v46 || true)) {
                    var v49 bool = (42 == 42)
                    switch "hello" {
                    case "":
                        v49 = (v43 > v43)
                    case "foo bar":
                        if v49 {
                            if true {
                                if v49 {
                                    v2 = (v2 >> ((v3) & 63))
                                }
                                v1 = append(v1, 42)
                            }
                            for fi50 := 0; fi50 < 13; fi50++ {
                                v3 = (v3 << ((genFn0(32767, -1)) & 63))
                                var v51 uint = (v3 >> ((v3) & 63))
                                var v52 int32 = 2147483647
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            fmt.Println(fi44, fi44)
                        }
                    }
                    v1 = append(v1, 42)
                    _ = v49
                }
            }
            _ = v46
        }
        var v53 uint64 = 0
        _ = fi44
        _ = v53
    }
    for fi54 := 0; fi54 < 14; fi54++ {
        for fi55 := 0; fi55 < 9; fi55++ {
            var v56 int64 = v2
            _ = fi55
            _ = v56
        }
        _ = fi54
    }
    if ("" <= ("foo bar" + ("\n\t" + ("a" + ("foo bar" + "foo bar"))))) {
        var v57 uint64 = 499
        var v58 int32 = -1
        _ = v57
        _ = v58
    }
    for fi59 := 0; fi59 < 6; fi59++ {
        v1 = append(v1, 511)
        switch "a" {
        case "\n\t":
            switch "hello" {
            case "hello":
                if len(v1) > 0 {
                    v1[0] = 1
                }
            }
        case "foo bar":
            fmt.Println(v2)
        }
        _ = fi59
    }
    for fi60 := 0; fi60 < 4; fi60++ {
        if false {
            v43 = (v43 | 2)
            if (!(!(2 < 404))) {
                fmt.Println(fi60)
            } else {
                v43 = 2
                v2 = v2
            }
            fmt.Println(v4, v3, v43)
        } else {
            fmt.Println(v3, fi60, v3)
        }
        _ = fi60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v43
}
