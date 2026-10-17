package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint16; F2 int64 }
type GenS5 struct { F0 int8; F1 bool; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int; F1 uint }
func genFn0(p0 uint16) int16 {
    var v7 int8 = 116
    v8 := GenS1{F0: (v7 >> ((4294967295) & 63))}
    v9 := GenS1{F0: v7}
    if (18446744073709551615 <= 365) {
        v10 := GenS2{F0: uint64(p0), F1: 2}
        _ = v10
    }
    v11 := []int{246}
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 1000
}
func genFn1() bool {
    var v12 uint16 = 1
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int32 = 2147483647
        for fi15 := 0; fi15 < 18; fi15++ {
            v12 -= v12
            _ = fi15
        }
        v16 := []int{}
        _ = fi13
        _ = v14
        _ = v16
    }
    if true {
        var v17 uint8 = 2
        for fi18 := 0; fi18 < 4; fi18++ {
            if ((42 > 4294967295) || (42 == 18446744073709551615)) {
                for fi19 := 0; fi19 < 9; fi19++ {
                    v17 += v17
                    v12 = (v12 / ((v12) | 1))
                    v17 = (v17 - (v17 / ((127) | 1)))
                    _ = fi19
                }
                v20 := GenS2{F0: 395, F1: v17}
                v21 := GenS1{F0: 23}
                _ = v20
                _ = v21
            } else {
                v12 += v12
                v22 := GenS2{F0: 2, F1: v17}
                _ = v22
            }
            if (((fi18 * (fi18 | fi18)) > fi18) && true) {
                v12 = 890
            } else {
                var v23 int = (fi18 * (fi18 & fi18))
                var v24 bool = true
                for fi25 := 0; fi25 < 15; fi25++ {
                    if v24 {
                        var v26 bool = (v12 > 42)
                        _ = v26
                    } else {
                        var v27 uint8 = v17
                        v23 = fi18
                        v17 = (^(v17))
                        _ = v27
                    }
                    _ = fi25
                }
                _ = v23
                _ = v24
            }
            v12 -= v12
            _ = fi18
        }
        v17 = v17
        _ = v17
    } else {
        var v28 bool = (-32768 < 42)
        v29 := GenS3{F0: 49}
        v30 := []int{}
        _ = v28
        _ = v29
        _ = v30
    }
    _ = v12
    return true
}

func main() {
    var v1 uint = 1
    var v2 uint = 65535
    var v3 float32 = 3.14159
    var v4 uint16 = 65535
    v5 := make(map[uint32]bool)
    var v6 int = 722
    for fi31 := 0; fi31 < 20; fi31++ {
        fmt.Println(v1, v1, v3)
        v1 -= (v1 & (v2 << ((72) & 63)))
        var v32 uint = v1
        _ = fi31
        _ = v32
    }
    v33 := GenS2{F0: 4294967295, F1: uint8(v4)}
    v34 := GenS3{F0: 2}
    v1 += v1
    v35 := v5[4294967295]
    v36 := GenS1{F0: 100}
    v37 := GenS1{F0: 92}
    var v38 int64 = 42
    v35 = v35
    var v39 uint64 = uint64(v4)
    if false {
        for fi40 := 0; fi40 < 18; fi40++ {
            fmt.Println(v1, v39)
            for fi41 := 0; fi41 < 19; fi41++ {
                if (float32(1.5) >= 0.1) {
                    var v42 int32 = 42
                    delete(v5, 2)
                    _ = v42
                } else {
                    delete(v5, 42)
                    var v43 uint8 = 127
                    delete(v5, 1)
                    _ = v43
                }
                _ = fi41
            }
            if true {
                v2 -= 538
            } else {
                v44 := GenS1{F0: 48}
                switch ((("" + "") + ("hello" + "a")) + ("a" + ((("hello" + "") + ((((((("a" + ((("hello" + (((((((("" + ("\n\t" + (("" + ("a" + "")) + "hello"))) + "") + "foo bar") + "a") + "") + "\n\t") + "") + "a")) + "foo bar") + ("foo bar" + "\n\t"))) + "foo bar") + "\n\t") + "\n\t") + "a") + "") + "\n\t")) + ""))) {
                case "a":
                    v1 = v2
                case "\n\t":
                    for fi45 := 0; fi45 < 8; fi45++ {
                        v38 += -1
                        _ = fi45
                    }
                case "":
                    if (false && v35) {
                        fmt.Println(v6)
                    }
                }
                if v35 {
                    v46 := GenS3{F0: 42}
                    var v47 int64 = 42
                    switch 9223372036854775807 {
                    case -491:
                        fmt.Println(v4, v1, v2)
                    default:
                        var v48 int8 = 103
                        _ = v48
                    }
                    _ = v46
                    _ = v47
                } else {
                    var v49 string = "foo bar"
                    for fi50 := 0; fi50 < 8; fi50++ {
                        v51 := GenS4{F0: -1, F1: v4, F2: (v38 >> ((65535) & 63))}
                        fmt.Println(v38, v6, fi50)
                        v52 := GenS5{F0: 127, F1: (-9223372036854775808 > 0), F2: 42}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    fmt.Println(v4)
                    _ = v49
                }
                _ = v44
            }
            _ = fi40
        }
    } else {
        fmt.Println(v39, v4)
        switch "foo bar" {
        case "a":
            v39 += v39
        case "foo bar":
            v53 := []int{(v6 / ((v6) | 1)), (v6 / (((v6 - v6)) | 1)), (-(v6))}
            _ = v53
        case "hello":
            switch int(v39) {
            case 0:
                v3 = (3.14159 - float32(0.0))
            case -9223372036854775808:
                v3 = (v3 / -2.25)
            case -913:
                var v54 uint8 = 75
                _ = v54
            }
        }
    }
    var v55 uint32 = 250
    v56 := GenS6{F0: 1, F1: uint8(v39), F2: (v55 | (v55 * (v55 / ((1) | 1))))}
    v39 += v39
    for fi57 := 0; fi57 < 7; fi57++ {
        if ("" != "\n\t") {
            v1 -= 65535
            switch int(v55) {
            case 0:
                var v58 int64 = (^(v38))
                _ = v58
            case 32767:
                v4 = (v4 | (v4 / (((v4 << ((v2) & 63))) | 1)))
            }
            fmt.Println(v35)
        } else {
            v59 := GenS7{F0: v38, F1: "\n\t"}
            var v60 uint8 = 128
            v1 = 1
            _ = v59
            _ = v60
        }
        var v61 int32 = 386
        var v62 string = "foo bar"
        _ = fi57
        _ = v61
        _ = v62
    }
    if ((v1 << ((18446744073709551615) & 63)) > uint(v4)) {
        for fi63 := 0; fi63 < 9; fi63++ {
            if v35 {
                switch "a" {
                case "":
                    var v64 int8 = 0
                    _ = v64
                }
                switch int(v39) {
                case 128:
                    v65 := GenS7{F0: (v38 | -9223372036854775808), F1: ""}
                    _ = v65
                case 4294967295:
                    v66 := GenS7{F0: -1, F1: "foo bar"}
                    _ = v66
                default:
                    var v67 int64 = v38
                    _ = v67
                }
            }
            switch "hello" {
            case "hello":
                v6 += (fi63 & 127)
            case "":
                fmt.Println(v4)
            }
            delete(v5, 296)
            _ = fi63
        }
        var v68 float32 = v3
        _ = v68
    } else {
        v69 := GenS8{F0: (^(v39))}
        if (v38 == (v38 ^ (v38 * (^(v38))))) {
            switch "foo bar" {
            case "hello":
                fmt.Println(v3)
            case "foo bar":
                for fi70 := 0; fi70 < 8; fi70++ {
                    v71 := GenS5{F0: 127, F1: true, F2: v55}
                    if v35 {
                        for fi72 := 0; fi72 < 10; fi72++ {
                            v73 := []int{(fi72 / (((v6 ^ 127)) | 1)), int(v55), -301}
                            switch "\n\t" {
                            case "foo bar":
                                v4 = v4
                            case "":
                                fmt.Println(v39, v4)
                            case "hello":
                                if (v39 > (v39 - (v39 | v39))) {
                                    v73 = append(v73, (^(fi72)))
                                }
                            default:
                                v2 -= v2
                            }
                            v74 := GenS7{F0: 9223372036854775807, F1: ("" + "\n\t")}
                            _ = fi72
                            _ = v73
                            _ = v74
                        }
                        v75 := GenS9{F0: v6, F1: v2}
                        var v76 int32 = 0
                        _ = v75
                        _ = v76
                    }
                    _ = fi70
                    _ = v71
                }
            case "a":
                v5[v55] = (42 == 0)
            }
            var v77 uint = 0
            v55 = 65535
            _ = v77
        }
        for fi78 := 0; fi78 < 6; fi78++ {
            v1 = 65535
            _ = fi78
        }
        _ = v69
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v6
}
