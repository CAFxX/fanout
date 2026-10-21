package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint8; F1 bool; F2 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32; F1 uint; F2 int16 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 bool }

func main() {
    var v1 uint16 = 65535
    var v2 uint32 = 264
    v3 := make(map[string]int)
    var v4 uint = uint(v2)
    switch 259 {
    case -585:
        v5 := v3["hello"]
        _ = v5
    case 0:
        v4 += v4
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS1{F0: fi6}
            switch (fi7 & (fi7 - fi6)) {
            case 128:
                v9 := GenS2{F0: 9223372036854775807}
                _ = v9
            case 490:
                fmt.Println(v1, v4)
            case 127:
                v10 := GenS3{F0: 1, F1: -1, F2: v1}
                _ = v10
            }
            var v11 int16 = 302
            _ = fi7
            _ = v11
            _ = v8
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            v1 += v1
            switch ("\n\t" + "") {
            case "a":
                v13 := GenS4{F0: (v4 - 687), F1: 2, F2: 32767}
                _ = v13
            }
            var v14 int = fi6
            _ = fi12
            _ = v14
        }
        _ = fi6
    }
    v15 := GenS2{F0: -349}
    if (!false) {
        _ = len(v3)
        v4 -= uint(v1)
        fmt.Println(v2, v2)
    } else {
        var v16 uint = 42
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            fmt.Println(fi17)
            v2 = 0
            if true {
                v19 := GenS5{F0: false, F1: -1}
                _ = v19
            }
            _ = fi18
        }
        var v20 bool = (v4 != (v4 ^ 99))
        for fi21 := 0; fi21 < 17; fi21++ {
            v1 = (v1 ^ v1)
            var v22 uint64 = 1
            v23 := []int{(fi21 ^ int(v1))}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi17
        _ = v20
    }
    switch "\n\t" {
    case "\n\t":
        switch "hello" {
        case "a":
            v24 := GenS5{F0: (true || true), F1: 42}
            _ = v24
        case "":
            switch 1 {
            case 127:
                v25 := GenS6{F0: 0, F1: true, F2: (0 < 42)}
                _ = v25
            case 128:
                if true {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        v27 := GenS7{F0: (v4 << ((v4) & 63))}
                        v4 = (v4 / ((v4) | 1))
                        _ = fi26
                        _ = v27
                    }
                    switch 9223372036854775807 {
                    case -156:
                        v3[(("a" + "") + (("a" + "foo bar") + ("foo bar" + "foo bar")))] = 32768
                    case 4294967295:
                        if (!((!(!(!(!(127 != 1))))) || false)) {
                            v28 := GenS8{F0: 2}
                            v29 := GenS9{F0: 128}
                            v30 := []int{32768, -2147483648}
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    case 281:
                        v4 = (^(v4))
                    default:
                        var v31 int16 = 1000
                        _ = v31
                    }
                    v32 := GenS6{F0: 42, F1: (3.14159 <= (0.0 + 1.5)), F2: false}
                    _ = v32
                }
            case 2147483647:
                v1 -= (^(v1))
            }
        default:
            var v33 int = int(v1)
            _ = v33
        }
    case "hello":
        v34 := GenS10{F0: (v2 >> ((214) & 63)), F1: v4, F2: 32767}
        _ = v34
    default:
        v35 := GenS6{F0: 199, F1: true, F2: ((v2 >= v2) && true)}
        _ = v35
    }
    v3[""] = -240
    var v36 bool = (1 <= -9223372036854775808)
    if (108 > 66) {
        if (-928 < 1) {
            v37 := v3[""]
            if (v1 <= v1) {
                v38 := []int{}
                v39 := GenS7{F0: v4}
                _ = v38
                _ = v39
            }
            switch v37 {
            case 32767:
                v37 = (v37 / (((v37 * v37)) | 1))
            default:
                if false {
                    v2 -= v2
                    if (44 < -90) {
                        if ((^(v4)) == (^(v4))) {
                            for fi40 := 0; fi40 < 6; fi40++ {
                                var v41 bool = false
                                _ = fi40
                                _ = v41
                            }
                            for fi42 := 0; fi42 < 19; fi42++ {
                                for fi43 := 0; fi43 < 16; fi43++ {
                                    if ((42 >= 42) && v36) {
                                        v44 := GenS11{F0: "foo bar"}
                                        if (-394 <= -2147483648) {
                                            v45 := []int{(^(fi42)), -857, 32768}
                                            if (!true) {
                                                switch (v37 >> (((^(v4))) & 63)) {
                                                case 2147483647:
                                                    v46 := GenS12{F0: v36}
                                                    _ = v46
                                                default:
                                                    v37 = (v37 << (((v4 % ((v4) | 1))) & 63))
                                                }
                                            }
                                            var v47 int = 32767
                                            _ = v45
                                            _ = v47
                                        }
                                        var v48 int8 = 0
                                        _ = v44
                                        _ = v48
                                    }
                                    v49 := []int{(-(fi42))}
                                    if len(v49) >= 2 {
                                        v49 = v49[1:len(v49)]
                                    }
                                    _ = fi43
                                    _ = v49
                                }
                                if true {
                                    for fi50 := 0; fi50 < 13; fi50++ {
                                        v51 := []int{(-(fi42)), fi50, (fi42 << ((18446744073709551615) & 63))}
                                        _ = fi50
                                        _ = v51
                                    }
                                } else {
                                    for fi52 := 0; fi52 < 8; fi52++ {
                                        var v53 uint32 = (v2 ^ v2)
                                        _ = fi52
                                        _ = v53
                                    }
                                    fmt.Println(v36, v37)
                                    fmt.Println(v1, v2, v36)
                                }
                                v54 := []int{(v37 ^ (fi42 | fi42))}
                                _ = fi42
                                _ = v54
                            }
                        }
                        var v55 uint16 = 42
                        _ = v55
                    } else {
                        switch "a" {
                        case "foo bar":
                            v56 := GenS1{F0: (v37 - v37)}
                            _ = v56
                        }
                    }
                }
            }
            _ = v37
        }
        v1 = (v1 * v1)
        v57 := v3[""]
        _ = v57
    } else {
        v1 -= v1
        v2 += (^(v2))
    }
    if false {
        if (v1 < uint16(v2)) {
            v36 = (302 < 18446744073709551615)
            v1 = 255
            switch 4294967295 {
            case 32767:
                for fi58 := 0; fi58 < 13; fi58++ {
                    for fi59 := 0; fi59 < 2; fi59++ {
                        fmt.Println(v36, fi58, v4)
                        _ = fi59
                    }
                    fmt.Println(v4)
                    _ = fi58
                }
            case 42:
                v3[""] = 127
            case 128:
                v60 := []int{0, -1}
                _ = v60
            }
        }
    } else {
        v36 = v36
        for fi61 := 0; fi61 < 8; fi61++ {
            v36 = (true || v36)
            var v62 int16 = -1
            _ = fi61
            _ = v62
        }
        var v63 int64 = 42
        _ = v63
    }
    for fi64 := 0; fi64 < 17; fi64++ {
        if v36 {
            var v65 int16 = 42
            fmt.Println(v36)
            _ = v65
        }
        _ = fi64
    }
    var v66 uint8 = 162
    fmt.Println(v36)
    v67 := GenS10{F0: v2, F1: (v4 >> (((v4 - v4)) & 63)), F2: -32768}
    v68 := GenS4{F0: 2, F1: (v66 - v66), F2: 2}
    v66 += (^(v66))
    if v36 {
        v66 = (v66 + (v66 ^ v66))
        v69 := []int{1, 530, 32767, -621}
        v36 = true
        _ = v69
    } else {
        for fi70 := 0; fi70 < 1; fi70++ {
            switch (^(fi70)) {
            case -863:
                v71 := GenS7{F0: (^(v4))}
                _ = v71
            case 4294967295:
                var v72 int16 = 32767
                _ = v72
            case 128:
                v73 := []int{}
                _ = v73
            default:
                switch (fi70 << ((v4) & 63)) {
                case -828:
                    v74 := []int{-260}
                    _ = v74
                case 4294967295:
                    var v75 int = fi70
                    _ = v75
                case -721:
                    var v76 int8 = -128
                    _ = v76
                default:
                    var v77 int64 = -817
                    _ = v77
                }
            }
            _ = fi70
        }
        v1 = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v66
    _ = v67
    _ = v68
}
