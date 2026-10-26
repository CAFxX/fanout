package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int; F1 uint8; F2 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int; F1 uint32 }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]bool)
    var v3 int8 = (-(v1))
    var v4 int64 = 1
    var v5 uint8 = 54
    v4 -= v4
    delete(v2, 4294967295)
    v6 := v2[428]
    switch "foo bar" {
    case "a":
        for fi7 := 0; fi7 < 10; fi7++ {
            fmt.Println(v6, v4)
            if (("" + ("hello" + "foo bar")) == "a") {
                fmt.Println(v3)
                for fi8 := 0; fi8 < 1; fi8++ {
                    v3 = (v3 << ((65535) & 63))
                    var v9 int = fi7
                    fmt.Println(v3, v1, v9)
                    _ = fi8
                    _ = v9
                }
                var v10 float32 = -2.25
                _ = v10
            } else {
                var v11 bool = v6
                var v12 int64 = 1
                switch "foo bar" {
                case "\n\t":
                    var v13 uint16 = 255
                    _ = v13
                case "hello":
                    v12 += v4
                case "":
                    if v11 {
                        switch "" {
                        case "foo bar":
                            v14 := []int{(-(fi7)), (fi7 >> ((969) & 63))}
                            _ = v14
                        case "\n\t":
                            var v15 float32 = ((((3.14159 / 1.0) + 0.0) * ((((1.5 - 3.14159) * -2.25) / -2.25) / 1.0)) / 3.14159)
                            _ = v15
                        case "":
                            if (2147483647 == 4294967295) {
                                if v11 {
                                    for fi16 := 0; fi16 < 11; fi16++ {
                                        v12 -= (v12 & (v12 % ((v12) | 1)))
                                        v11 = (472 <= v12)
                                        v17 := []int{}
                                        _ = fi16
                                        _ = v17
                                    }
                                    v18 := []int{2147483647, fi7, int(v3)}
                                    for fi19 := 0; fi19 < 8; fi19++ {
                                        fmt.Println(v4, fi7, v3)
                                        fmt.Println(v3, v1)
                                        _ = fi19
                                    }
                                    _ = v18
                                } else {
                                    v4 -= v4
                                    var v20 float32 = 1.5
                                    var v21 uint32 = 0
                                    _ = v20
                                    _ = v21
                                }
                            } else {
                                if (32767 >= 255) {
                                    if v6 {
                                        v5 = 2
                                    }
                                    if (fi7 < fi7) {
                                        var v22 uint64 = 4294967295
                                        v23 := GenS1{F0: -9223372036854775808, F1: 723}
                                        _ = v22
                                        _ = v23
                                    }
                                    var v24 uint32 = 1
                                    _ = v24
                                }
                                v1 = (v3 + (v1 + (^(v3))))
                            }
                        default:
                            if (!(!(1.5 <= (0.0 * 0.0)))) {
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    if v6 {
                                        v6 = v11
                                        var v26 int64 = -351
                                        v27 := v2[524]
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = fi25
                                }
                            }
                        }
                        if v6 {
                            for fi28 := 0; fi28 < 6; fi28++ {
                                v2[1] = (42 >= 0)
                                _ = fi28
                            }
                            var v29 uint64 = 1
                            var v30 uint = 18446744073709551615
                            _ = v29
                            _ = v30
                        }
                        v31 := GenS2{F0: 18446744073709551615}
                        _ = v31
                    } else {
                        v32 := GenS2{F0: 2}
                        var v33 int32 = 0
                        switch ("foo bar" + "hello") {
                        case "a":
                            var v34 uint8 = (v5 >> ((184) & 63))
                            _ = v34
                        default:
                            v11 = false
                        }
                        _ = v32
                        _ = v33
                    }
                default:
                    if v11 {
                        v35 := GenS3{F0: 2}
                        _ = v35
                    }
                }
                _ = v11
                _ = v12
            }
            v36 := []int{(-(fi7))}
            _ = fi7
            _ = v36
        }
    case "\n\t":
        var v37 float64 = (-2.25 + ((1.5 / (3.14159 * 0.1)) * 1e10))
        _ = v37
    default:
        for fi38 := 0; fi38 < 16; fi38++ {
            v1 = v3
            _ = fi38
        }
    }
    if v6 {
        var v39 uint64 = 2
        _ = v39
    }
    var v40 float64 = 1e10
    if (!((!v6) || v6)) {
        fmt.Println(v1)
    } else {
        v40 -= 1.5
    }
    var v41 int8 = (v3 / ((1) | 1))
    if v6 {
        if v6 {
            for fi42 := 0; fi42 < 19; fi42++ {
                v43 := GenS3{F0: 65535}
                var v44 uint32 = 42
                _ = fi42
                _ = v43
                _ = v44
            }
            var v45 uint16 = 1
            _ = v45
        }
        fmt.Println(v40)
    } else {
        for fi46 := 0; fi46 < 14; fi46++ {
            var v47 int16 = 32767
            v48 := []int{int(v4), fi46}
            v49 := GenS3{F0: 0}
            _ = fi46
            _ = v47
            _ = v48
            _ = v49
        }
        v41 += (v3 << ((163) & 63))
    }
    switch "foo bar" {
    case "hello":
        v40 -= v40
    case "a":
        var v50 int16 = 1000
        _ = v50
    }
    fmt.Println(v1)
    v51 := GenS2{F0: 1}
    v52 := v2[2147483647]
    if (((((("\n\t" + "a") + ("" + ("foo bar" + ""))) + ("\n\t" + ((("\n\t" + "hello") + "") + "\n\t"))) + "foo bar") + "") >= "a") {
        switch "hello" {
        case "\n\t":
            if (524 <= -1) {
                v53 := GenS1{F0: (v4 >> ((65) & 63)), F1: -519}
                var v54 bool = (-1 >= 42)
                _ = v53
                _ = v54
            } else {
                switch "a" {
                case "a":
                    v55 := GenS4{F0: -612, F1: 645, F2: (v5 - v5)}
                    _ = v55
                case "\n\t":
                    var v56 float64 = (v40 / 1.5)
                    _ = v56
                case "":
                    _ = len(v2)
                default:
                    var v57 int = 1
                    _ = v57
                }
                for fi58 := 0; fi58 < 20; fi58++ {
                    switch 1 {
                    case 2147483647:
                        v41 = v41
                    case 9223372036854775807:
                        v59 := GenS5{F0: (fi58 | (-(fi58))), F1: v5, F2: ""}
                        _ = v59
                    case -373:
                        if (v3 <= -128) {
                            v60 := GenS6{F0: 1}
                            fmt.Println(fi58)
                            _ = v60
                        }
                    }
                    _ = fi58
                }
            }
        default:
            v6 = (0 > 18446744073709551615)
        }
        v3 = (v3 | v1)
    }
    if (("\n\t" + ("foo bar" + "hello")) > "a") {
        delete(v2, 16)
        switch ("hello" + "foo bar") {
        case "foo bar":
            switch 2147483647 {
            case 9223372036854775807:
                switch ((("foo bar" + "\n\t") + "") + "a") {
                case "hello":
                    delete(v2, 2)
                case "foo bar":
                    for fi61 := 0; fi61 < 1; fi61++ {
                        v62 := GenS4{F0: -2147483648, F1: 885, F2: (v5 >> ((42) & 63))}
                        v63 := GenS7{F0: 1}
                        v64 := GenS8{F0: fi61, F1: v52}
                        _ = fi61
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                }
            case -1:
                v5 -= v5
            case 1:
                v65 := GenS7{F0: 1}
                _ = v65
            default:
                v6 = ((!v52) || (1 <= 1))
            }
        case "\n\t":
            v66 := GenS9{F0: -2147483648, F1: 65535}
            _ = v66
        }
        v1 = (v1 - 127)
    }
    var v67 float64 = -2.25
    var v68 int16 = 0
    for fi69 := 0; fi69 < 5; fi69++ {
        switch 22 {
        case -414:
            v70 := GenS5{F0: fi69, F1: (v5 | 1), F2: "hello"}
            _ = v70
        case -9223372036854775808:
            if v52 {
                v40 -= (v40 / (v67 + v67))
                fmt.Println(fi69)
            } else {
                fmt.Println(v4, v52)
                if true {
                    fmt.Println(v52)
                    var v71 uint = 2
                    _ = v71
                }
            }
        }
        _ = fi69
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v67
    _ = v68
}
