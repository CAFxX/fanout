package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint8; F1 int64; F2 int }
type GenS5 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 uint64) int32 {
    var v4 int64 = -9223372036854775808
    var v5 int8 = 87
    v4 = v4
    fmt.Println(v4, v5)
    if true {
        if (1 <= 42) {
            p0 = p0
            v6 := []int{}
            p0 -= (p0 / (((p0 / (((p0 | (p0 | (p0 * 128)))) | 1))) | 1))
            _ = v6
        }
        v7 := GenS1{F0: v4, F1: v4, F2: p1}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS2{F0: 255, F1: 0}
            for fi10 := 0; fi10 < 2; fi10++ {
                if false {
                    fmt.Println(p0, fi8, fi8)
                    var v11 int32 = 42
                    _ = v11
                }
                _ = fi10
            }
            switch "" {
            case "hello":
                for fi12 := 0; fi12 < 9; fi12++ {
                    v5 += v5
                    if false {
                        var v13 int32 = 2147483647
                        v14 := []int{(fi12 << ((1) & 63))}
                        _ = v13
                        _ = v14
                    }
                    _ = fi12
                }
            case "a":
                v15 := []int{int(v5), fi8}
                _ = v15
            case "\n\t":
                if true {
                    for fi16 := 0; fi16 < 13; fi16++ {
                        fmt.Println(fi16, fi8)
                        _ = fi16
                    }
                } else {
                    v17 := []int{p0, p0}
                    p0 -= 32768
                    _ = v17
                }
            default:
                v18 := GenS3{F0: 65535, F1: -1}
                _ = v18
            }
            _ = fi8
            _ = v9
        }
        if false {
            for fi19 := 0; fi19 < 8; fi19++ {
                fmt.Println(p0, fi19, v5)
                v20 := []int{32767, (^(fi19))}
                switch (-(p0)) {
                case 128:
                    if (!(1000 > 761)) {
                        var v21 bool = ("foo bar" == "a")
                        _ = v21
                    }
                case 606:
                    for fi22 := 0; fi22 < 17; fi22++ {
                        v23 := GenS1{F0: (v4 & 1), F1: v4, F2: 1}
                        _ = len(v20)
                        p1 -= p1
                        _ = fi22
                        _ = v23
                    }
                case 0:
                    v20 = append(v20, (p0 / (((fi19 >> ((899) & 63))) | 1)))
                }
                _ = fi19
                _ = v20
            }
            v4 = (-(v4))
            v24 := []int{(p0 / (((p0 * 42)) | 1)), (p0 + int(p1))}
            _ = v24
        } else {
            if true {
                fmt.Println(v5)
            }
            p1 -= p1
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 int8 = v5
                v27 := GenS1{F0: v4, F1: v4, F2: p1}
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        if true {
            v5 = 100
            if (-69 <= 1) {
                v5 = int8(v4)
                p0 = (^(p0))
                var v28 int = p0
                _ = v28
            }
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}

func main() {
    v1 := []int{42, 42, 0}
    var v2 string = ("foo bar" + ((("foo bar" + "hello") + (("foo bar" + ("foo bar" + "foo bar")) + ((("foo bar" + "hello") + ("\n\t" + "\n\t")) + "hello"))) + "hello"))
    var v3 int = -1
    var v29 float32 = 1.5
    v30 := GenS4{F0: 125, F1: 849, F2: v3}
    for fi31 := 0; fi31 < 4; fi31++ {
        v32 := GenS5{F0: 2, F1: 1, F2: 1}
        for fi33 := 0; fi33 < 17; fi33++ {
            switch fi33 {
            case 42:
                var v34 int = (^(fi33))
                _ = v34
            }
            v35 := GenS3{F0: 32767, F1: -1}
            _ = fi33
            _ = v35
        }
        fmt.Println(fi31)
        _ = fi31
        _ = v32
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    v36 := GenS2{F0: 178, F1: 1}
    if true {
        v29 = v29
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = len(v1)
    } else {
        if (128 != 128) {
            if len(v1) > 0 {
                v1[0] = (v3 & v3)
            }
        }
        for fi37 := 0; fi37 < 16; fi37++ {
            if len(v1) > 0 {
                v1[0] = (v3 << ((880) & 63))
            }
            var v38 uint8 = 127
            if (!(18446744073709551615 < 42)) {
                var v39 string = v2
                _ = v39
            } else {
                v1 = append(v1, v3)
                if true {
                    v29 -= -2.25
                    var v40 uint32 = 2
                    _ = v40
                }
            }
            _ = fi37
            _ = v38
        }
    }
    v41 := GenS2{F0: 125, F1: 18446744073709551615}
    if len(v1) > 0 {
        v1[0] = (^(v3))
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 uint64 = 1
        if false {
            if ((!false) || false) {
                for fi44 := 0; fi44 < 18; fi44++ {
                    switch v2 {
                    case "foo bar":
                        v45 := GenS6{F0: -128}
                        _ = v45
                    default:
                        v46 := GenS6{F0: 42}
                        _ = v46
                    }
                    _ = fi44
                }
                fmt.Println(v2, v43, v29)
            }
            for fi47 := 0; fi47 < 1; fi47++ {
                v48 := GenS3{F0: 32767, F1: 42}
                if false {
                    _ = len(v1)
                    v29 -= (v29 / v29)
                }
                if false {
                    var v49 int64 = 321
                    v3 += v3
                    v50 := GenS1{F0: -1, F1: v49, F2: (v43 / (((v43 >> ((3) & 63))) | 1))}
                    _ = v49
                    _ = v50
                } else {
                    v3 = (fi42 % ((-572) | 1))
                }
                _ = fi47
                _ = v48
            }
        }
        _ = len(v1)
        _ = fi42
        _ = v43
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v51 int64 = 0
        _ = v51
    case -21:
        switch v3 {
        case 1:
            switch ("\n\t" + v2) {
            case "a":
                var v52 int64 = 42
                _ = v52
            case "":
                if true {
                    if true {
                        v53 := GenS1{F0: -385, F1: -1, F2: 18446744073709551615}
                        _ = v53
                    } else {
                        if true {
                            if true {
                                var v54 int64 = 894
                                _ = v54
                            }
                        } else {
                            var v55 uint32 = 65535
                            _ = v55
                        }
                    }
                    _ = len(v1)
                    var v56 int8 = 68
                    _ = v56
                }
            case "foo bar":
                if (genFn0(v3, 18446744073709551615) == -802) {
                    switch v3 {
                    case 127:
                        v3 = (v3 / ((v3) | 1))
                    case 9223372036854775807:
                        for fi57 := 0; fi57 < 6; fi57++ {
                            for fi58 := 0; fi58 < 7; fi58++ {
                                var v59 uint = 2
                                if false {
                                    if len(v1) > 0 {
                                        v1[0] = (-(fi57))
                                    }
                                    var v60 bool = (true || ((!true) || (67 < 869)))
                                    _ = len(v1)
                                    _ = v60
                                } else {
                                    fmt.Println(fi57)
                                    v29 -= ((v29 / v29) + (v29 - (0.1 - 1.5)))
                                    v2 = "a"
                                }
                                fmt.Println(fi57)
                                _ = fi58
                                _ = v59
                            }
                            _ = fi57
                        }
                    default:
                        v3 = -512
                    }
                }
            default:
                var v61 int16 = -32768
                _ = v61
            }
        case -73:
            fmt.Println(v29)
        default:
            for fi62 := 0; fi62 < 3; fi62++ {
                if false {
                    var v63 uint = 42
                    _ = v63
                } else {
                    for fi64 := 0; fi64 < 1; fi64++ {
                        if ((147 < 1) || true) {
                            v65 := GenS7{F0: 901, F1: 1}
                            switch fi62 {
                            case 1:
                                var v66 int16 = 32767
                                _ = v66
                            case -486:
                                v1 = append(v1, (fi64 / (((fi62 & fi62)) | 1)))
                                v67 := GenS7{F0: -248, F1: -1}
                                v29 = 1e10
                                _ = v67
                            }
                            switch 32768 {
                            case 244:
                                v3 -= fi62
                            }
                            _ = v65
                        }
                        _ = fi64
                    }
                    v68 := GenS8{F0: -934, F1: 0}
                    _ = v68
                }
                _ = fi62
            }
        }
    }
    switch ("" + v2) {
    case "foo bar":
        for fi69 := 0; fi69 < 20; fi69++ {
            v3 = (v3 << ((233) & 63))
            _ = fi69
        }
    case "hello":
        var v70 int16 = 42
        _ = v70
    }
    v1 = append(v1, 0)
    var v71 int64 = -984
    v3 = v3
    v29 = v29
    v3 -= -2147483648
    switch v2 {
    case "hello":
        var v72 uint64 = 206
        _ = v72
    case "\n\t":
        v29 = v29
    default:
        var v73 int8 = -128
        _ = v73
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v41
    _ = v71
}
