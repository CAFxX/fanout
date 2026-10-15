package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS3 struct { F0 string; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint16 }
func genFn0(p0 bool, p1 int, p2 uint) bool {
    var v4 bool = p0
    v5 := []int{(p1 >> (((^(p2))) & 63))}
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int32 = 2147483647
            v9 := GenS1{F0: -128}
            fmt.Println(fi6, fi6)
            _ = fi7
            _ = v8
            _ = v9
        }
        if v4 {
            var v10 int64 = -9223372036854775808
            _ = v10
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int64 = -1
            fmt.Println(v4)
            _ = fi11
            _ = v12
        }
        _ = fi6
    }
    switch "hello" {
    case "a":
        v13 := GenS1{F0: 1}
        _ = v13
    case "hello":
        var v14 uint = (p2 - (p2 ^ 4294967295))
        _ = v14
    default:
        if v4 {
            for fi15 := 0; fi15 < 4; fi15++ {
                fmt.Println(p2, v4)
                _ = fi15
            }
            if true {
                var v16 uint16 = 32767
                _ = v16
            }
            if (0.0 <= 0.0) {
                for fi17 := 0; fi17 < 8; fi17++ {
                    v18 := GenS2{F0: 42, F1: 0, F2: 223}
                    for fi19 := 0; fi19 < 18; fi19++ {
                        _ = len(v5)
                        p2 = p2
                        fmt.Println(fi17, fi17, fi17)
                        _ = fi19
                    }
                    p1 = (p1 * p1)
                    _ = fi17
                    _ = v18
                }
            } else {
                _ = len(v5)
                for fi20 := 0; fi20 < 10; fi20++ {
                    if v4 {
                        var v21 int32 = -2147483648
                        if len(v5) > 0 {
                            v5[0] = 127
                        }
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 16; fi22++ {
                            p1 = fi20
                            var v23 int32 = -1
                            var v24 int64 = -1
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if (("foo bar" + "") <= ("\n\t" + ("hello" + (("foo bar" + "hello") + "hello")))) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if (p0 || false) {
                                var v25 uint8 = 0
                                var v26 uint16 = 334
                                if len(v5) > 0 {
                                    v5[0] = p1
                                }
                                _ = v25
                                _ = v26
                            } else {
                                if true {
                                    v27 := GenS3{F0: ("a" + "foo bar"), F1: 1, F2: 255}
                                    var v28 float64 = 1e10
                                    p0 = p0
                                    _ = v27
                                    _ = v28
                                } else {
                                    switch (fi20 & p1) {
                                    case 4294967295:
                                        p0 = (v4 && (!v4))
                                    case 2147483647:
                                        p2 += p2
                                    case 1:
                                        for fi29 := 0; fi29 < 8; fi29++ {
                                            var v30 uint16 = 0
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = fi29
                                            _ = v30
                                        }
                                    default:
                                        p2 = p2
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = p1
                                    }
                                }
                            }
                        } else {
                            v31 := GenS2{F0: 1, F1: -9223372036854775808, F2: 255}
                            p2 += (p2 % ((255) | 1))
                            if v4 {
                                if v4 {
                                    switch fi20 {
                                    case -1:
                                        for fi32 := 0; fi32 < 10; fi32++ {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if (!(v4 || p0)) {
                                                switch 42 {
                                                case 4294967295:
                                                    var v33 bool = p0
                                                    _ = v33
                                                case -472:
                                                    p0 = p0
                                                }
                                                if (!true) {
                                                    var v34 float64 = 0.0
                                                    fmt.Println(p1, p2)
                                                    var v35 uint64 = 820
                                                    _ = v34
                                                    _ = v35
                                                }
                                            }
                                            var v36 uint16 = 32767
                                            _ = fi32
                                            _ = v36
                                        }
                                    case -892:
                                        var v37 uint32 = 0
                                        _ = v37
                                    case -9223372036854775808:
                                        fmt.Println(p0, p1)
                                    default:
                                        fmt.Println(v4, p0)
                                    }
                                    var v38 int8 = 9
                                    _ = v38
                                } else {
                                    var v39 int32 = 0
                                    v40 := GenS1{F0: 100}
                                    _ = v39
                                    _ = v40
                                }
                            }
                            _ = v31
                        }
                    } else {
                        var v41 int64 = -1
                        _ = v41
                    }
                    var v42 uint64 = 42
                    _ = fi20
                    _ = v42
                }
            }
        } else {
            if p0 {
                v43 := GenS1{F0: 42}
                _ = v43
            } else {
                if false {
                    for fi44 := 0; fi44 < 16; fi44++ {
                        var v45 string = ("hello" + "hello")
                        _ = fi44
                        _ = v45
                    }
                    v4 = (!(v4 || true))
                    switch p1 {
                    case 9223372036854775807:
                        p2 += p2
                    }
                }
                switch "hello" {
                case "":
                    var v46 uint64 = 481
                    _ = v46
                case "foo bar":
                    var v47 uint16 = 255
                    _ = v47
                }
                _ = len(v5)
            }
            var v48 int = (p1 % ((-2147483648) | 1))
            _ = v48
        }
    }
    p2 -= (p2 * (p2 + 243))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (992 < 307)
}

func main() {
    var v1 int16 = 42
    var v2 int16 = v1
    v3 := make(map[int]bool)
    for fi49 := 0; fi49 < 12; fi49++ {
        for fi50 := 0; fi50 < 7; fi50++ {
            switch (fi49 & (^(fi49))) {
            case -676:
                v1 -= (v2 - 5)
            }
            _ = fi50
        }
        for fi51 := 0; fi51 < 6; fi51++ {
            var v52 uint32 = 0
            if true {
                _ = len(v3)
            }
            if (863 < 65535) {
                fmt.Println(fi51)
                v3[(fi51 - fi51)] = false
            }
            v53 := []int{fi49, (fi51 ^ fi51)}
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = fi49
    }
    v1 = v2
    v2 = v1
    fmt.Println(v1, v1)
    if (1 >= 32767) {
        v1 = v2
    } else {
        fmt.Println(v1, v1)
        switch -9223372036854775808 {
        case 1:
            v54 := GenS1{F0: 100}
            _ = v54
        case -9223372036854775808:
            v55 := []int{}
            _ = v55
        case -1:
            fmt.Println(v2, v2)
        default:
            var v56 float64 = 3.14159
            _ = v56
        }
    }
    v2 -= v2
    v57 := GenS4{F0: 42, F1: 906}
    fmt.Println(v2, v2)
    var v58 uint64 = 951
    if (!(100 > -1)) {
        v59 := GenS4{F0: 2147483648, F1: 971}
        _ = v59
    } else {
        if (!false) {
            if (!false) {
                if false {
                    if ((v58 == v58) && (!false)) {
                        v1 += 0
                        if true {
                            fmt.Println(v58)
                        }
                        if (!true) {
                            v60 := []int{}
                            v3[226] = (1 < 65535)
                            _ = v60
                        }
                    }
                }
                v61 := GenS2{F0: 127, F1: 0, F2: 65535}
                _ = v61
            }
        } else {
            var v62 int32 = 61
            var v63 bool = (false && (127 != -128))
            _ = v62
            _ = v63
        }
        v1 += (v1 >> ((1) & 63))
        for fi64 := 0; fi64 < 17; fi64++ {
            delete(v3, fi64)
            _ = fi64
        }
    }
    switch ("a" + "a") {
    case "":
        if (!(true && (0 == 9223372036854775807))) {
            fmt.Println(v58, v1, v1)
            var v65 uint = 2
            _ = v65
        } else {
            v66 := v3[-1]
            v58 -= 1000
            switch "a" {
            case "\n\t":
                if (0 <= 2147483647) {
                    v66 = true
                    if (!false) {
                        v67 := []int{0, 965, 9223372036854775807, 42}
                        var v68 int32 = 0
                        if len(v67) > 0 {
                            v67[0] = 167
                        }
                        _ = v67
                        _ = v68
                    } else {
                        for fi69 := 0; fi69 < 4; fi69++ {
                            v2 = (^(v2))
                            _ = fi69
                        }
                        if (!v66) {
                            var v70 int8 = 42
                            _ = v70
                        } else {
                            if v66 {
                                for fi71 := 0; fi71 < 11; fi71++ {
                                    v1 = (v2 * v1)
                                    _ = fi71
                                }
                                v66 = false
                                fmt.Println(v58)
                            } else {
                                for fi72 := 0; fi72 < 11; fi72++ {
                                    for fi73 := 0; fi73 < 15; fi73++ {
                                        v58 += (v58 << ((0) & 63))
                                        if v66 {
                                            v66 = (-1 >= 9223372036854775807)
                                            fmt.Println(v2)
                                        }
                                        _ = fi73
                                    }
                                    _ = fi72
                                }
                                var v74 bool = v66
                                v3[32768] = v74
                                _ = v74
                            }
                        }
                        fmt.Println(v1, v66)
                    }
                } else {
                    v75 := GenS2{F0: 128, F1: -349, F2: 65535}
                    _ = v75
                }
            }
            _ = v66
        }
    }
    var v76 string = "foo bar"
    v58 += (v58 / ((1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v58:", v58)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v57
    _ = v58
    _ = v76
}
