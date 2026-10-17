package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int16; F1 string; F2 bool }
type GenS6 struct { F0 int8; F1 string; F2 uint32 }
func genFn0() int8 {
    var v6 uint32 = 1
    if (4294967295 < 32767) {
        switch "\n\t" {
        case "foo bar":
            var v7 int16 = 42
            _ = v7
        }
        v6 -= v6
    } else {
        v6 += (v6 * (v6 & (v6 & v6)))
        switch "foo bar" {
        case "":
            var v8 string = "a"
            _ = v8
        case "foo bar":
            if ((true || false) || true) {
                v6 -= v6
                v9 := GenS1{F0: 2, F1: 42, F2: 127}
                for fi10 := 0; fi10 < 16; fi10++ {
                    v11 := GenS1{F0: 42, F1: 18446744073709551615, F2: fi10}
                    v12 := []int{}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v9
            } else {
                switch 127 {
                case -9223372036854775808:
                    if (0 <= 2147483647) {
                        switch (("" + (("hello" + "foo bar") + "foo bar")) + "") {
                        case "a":
                            var v13 bool = true
                            _ = v13
                        }
                        var v14 int32 = 881
                        _ = v14
                    }
                }
            }
        case "hello":
            v6 -= (v6 >> ((888) & 63))
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 int32 = 2147483647
            _ = fi15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v6 = (v6 % ((42) | 1))
        _ = fi17
    }
    _ = v6
    return 100
}
func genFn1(p0 uint16, p1 uint8) uint32 {
    var v18 int16 = 42
    v18 = v18
    var v19 bool = ((100 > 2) || true)
    if (-843 >= -711) {
        v20 := []int{-475, 76}
        var v21 int32 = 0
        _ = v20
        _ = v21
    } else {
        if ((v18 & (v18 % (((v18 << ((18446744073709551615) & 63))) | 1))) < 882) {
            var v22 uint8 = (p1 / ((p1) | 1))
            switch ("" + ("hello" + ((("" + "") + "foo bar") + "foo bar"))) {
            case "hello":
                for fi23 := 0; fi23 < 2; fi23++ {
                    if (fi23 <= 4294967295) {
                        var v24 int = -850
                        v24 += (-(fi23))
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 1; fi25++ {
                            var v26 int32 = 0
                            v27 := []int{fi25}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        fmt.Println(v22, p1)
                        v28 := []int{(fi23 * (fi23 << ((255) & 63))), (fi23 | (fi23 | -280)), (-(fi23)), fi23}
                        _ = v28
                    }
                    if (((3.14159 * (0.1 * ((1e10 + 1e10) / float32(((3.14159 * 1e10) * (float64(0.1) + 1.5)))))) * 3.14159) <= 0.1) {
                        v19 = (!v19)
                    } else {
                        v29 := GenS1{F0: 853, F1: 1, F2: 0}
                        _ = v29
                    }
                    fmt.Println(v18, p1)
                    _ = fi23
                }
            case "foo bar":
                for fi30 := 0; fi30 < 9; fi30++ {
                    fmt.Println(v19, v22, v18)
                    if ((0.0 + float32((1e10 + (-2.25 + 3.14159)))) > (3.14159 * (1.5 / (0.1 + 1.5)))) {
                        for fi31 := 0; fi31 < 12; fi31++ {
                            p1 = (p1 ^ v22)
                            _ = fi31
                        }
                        v22 = (p1 ^ (^(v22)))
                    } else {
                        p1 = (p1 ^ p1)
                        if v19 {
                            v18 -= (^(v18))
                        } else {
                            switch (fi30 & fi30) {
                            case 32768:
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    switch 1 {
                                    case -100:
                                        var v33 int16 = 32767
                                        _ = v33
                                    case 127:
                                        switch "hello" {
                                        case "hello":
                                            p0 -= p0
                                        case "\n\t":
                                            fmt.Println(v19, p1)
                                        }
                                    case 4294967295:
                                        v34 := []int{fi30, -9223372036854775808, -614, fi32}
                                        _ = v34
                                    default:
                                        v35 := []int{fi30}
                                        _ = v35
                                    }
                                    var v36 int64 = int64(v18)
                                    _ = fi32
                                    _ = v36
                                }
                            case 1:
                                var v37 uint8 = 1
                                _ = v37
                            }
                            p1 = 127
                            v38 := GenS1{F0: 0, F1: 18446744073709551615, F2: (fi30 + fi30)}
                            _ = v38
                        }
                        p0 -= (p0 & (^(p0)))
                    }
                    _ = fi30
                }
            case "":
                fmt.Println(v19, p0, v22)
            }
            _ = v22
        }
        v19 = false
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 4294967295
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 100
    var v3 int8 = v2
    var v4 int16 = 0
    var v5 float32 = ((1e10 + (((((1e10 / 0.1) / 1e10) * (1e10 - float32((1e10 - 1.5)))) * 3.14159) + 0.0)) + ((0.1 / 1e10) * -2.25))
    if (false || ((((float64(float32(1e10)) + 0.0) * 0.1) + (float64(1e10) + (3.14159 + (0.1 + (-2.25 * (1e10 / (1.5 + float64((v5 / ((v5 * 3.14159) * v5)))))))))) >= ((-2.25 * 0.0) * 1.5))) {
        switch "\n\t" {
        case "hello":
            v5 -= v5
        case "":
            switch "hello" {
            case "hello":
                for fi39 := 0; fi39 < 18; fi39++ {
                    fmt.Println(v3, v5)
                    _ = fi39
                }
            case "foo bar":
                if (-85 >= 0) {
                    for fi40 := 0; fi40 < 3; fi40++ {
                        v5 -= 0.0
                        _ = fi40
                    }
                } else {
                    v2 += v3
                    for fi41 := 0; fi41 < 1; fi41++ {
                        for fi42 := 0; fi42 < 16; fi42++ {
                            var v43 float32 = v5
                            _ = fi42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 1; fi44++ {
                            v45 := GenS1{F0: 2, F1: 42, F2: fi41}
                            _ = fi44
                            _ = v45
                        }
                        _ = fi41
                    }
                }
            }
        case "\n\t":
            for fi46 := 0; fi46 < 6; fi46++ {
                var v47 float64 = float64(((((1e10 + 1e10) / v5) + v5) - v5))
                _ = fi46
                _ = v47
            }
        }
        for fi48 := 0; fi48 < 14; fi48++ {
            var v49 int = (^(fi48))
            _ = len(v1)
            var v50 uint = 1
            _ = fi48
            _ = v49
            _ = v50
        }
        for fi51 := 0; fi51 < 1; fi51++ {
            v52 := []int{(-(fi51)), fi51, fi51}
            _ = fi51
            _ = v52
        }
    } else {
        v53 := []int{}
        v54 := GenS2{F0: 32767}
        if len(v53) >= 2 {
            v53 = v53[1:len(v53)]
        }
        _ = v53
        _ = v54
    }
    var v55 int = 42
    for fi56 := 0; fi56 < 14; fi56++ {
        if false {
            v5 -= (v5 / ((float32((0.0 + (-2.25 / 0.1))) / ((0.1 / v5) * v5)) * (v5 + v5)))
        }
        v57 := v1[(("hello" + "") + ("\n\t" + (("hello" + "foo bar") + "foo bar")))]
        v57 = (v57 * fi56)
        _ = fi56
        _ = v57
    }
    var v58 int32 = 42
    switch 677 {
    case 1:
        var v59 int = 9223372036854775807
        _ = v59
    case 32767:
        v4 -= v4
    case 2147483647:
        v55 = (v55 << ((722) & 63))
    }
    v58 = -207
    if true {
        fmt.Println(v58, v2)
    } else {
        if ((0.0 - 0.0) < 0.0) {
            for fi60 := 0; fi60 < 10; fi60++ {
                v61 := GenS3{F0: 2}
                for fi62 := 0; fi62 < 8; fi62++ {
                    fmt.Println(v58)
                    _ = fi62
                }
                var v63 uint32 = 623
                _ = fi60
                _ = v61
                _ = v63
            }
            v64 := GenS4{F0: 924, F1: false}
            var v65 string = "\n\t"
            _ = v64
            _ = v65
        }
        var v66 string = "a"
        _ = v66
    }
    for fi67 := 0; fi67 < 18; fi67++ {
        v55 = fi67
        switch (^(v55)) {
        case 128:
            var v68 uint64 = 362
            _ = v68
        case 1:
            for fi69 := 0; fi69 < 5; fi69++ {
                v55 += (^(fi69))
                v5 = 1.5
                _ = fi69
            }
        case -430:
            v70 := GenS4{F0: 42, F1: (!false)}
            _ = v70
        default:
            if (((42 <= 42) && ((1e10 <= (1e10 / 1e10)) || true)) || true) {
                switch ("hello" + ("hello" + "a")) {
                case "":
                    v71 := GenS4{F0: 42, F1: true}
                    _ = v71
                case "a":
                    v72 := GenS5{F0: 1, F1: (("foo bar" + ("" + "")) + "a"), F2: true}
                    _ = v72
                default:
                    var v73 int32 = 0
                    _ = v73
                }
                v3 = genFn0()
            } else {
                for fi74 := 0; fi74 < 15; fi74++ {
                    v3 = -64
                    v55 = (fi67 ^ v55)
                    fmt.Println(v58, v2, v58)
                    _ = fi74
                }
            }
        }
        _ = fi67
    }
    v75 := GenS6{F0: (v3 | v2), F1: "foo bar", F2: 4294967295}
    switch v55 {
    case 0:
        fmt.Println(v3, v2)
    default:
        var v76 float64 = float64(0.0)
        _ = v76
    }
    if ("foo bar" >= "a") {
        fmt.Println(v58)
        for fi77 := 0; fi77 < 20; fi77++ {
            v3 = 42
            switch (fi77 / ((fi77) | 1)) {
            case 32767:
                var v78 bool = false
                _ = v78
            case 42:
                for fi79 := 0; fi79 < 2; fi79++ {
                    var v80 string = "hello"
                    _ = len(v1)
                    _ = fi79
                    _ = v80
                }
            }
            var v81 bool = true
            _ = fi77
            _ = v81
        }
    } else {
        switch v55 {
        case 443:
            switch (("" + "") + "foo bar") {
            case "":
                v82 := GenS5{F0: v4, F1: "\n\t", F2: ((2 >= 4294967295) && false)}
                _ = v82
            case "a":
                v83 := []int{42, -952, v55, (-(v55))}
                _ = v83
            }
        case -2147483648:
            for fi84 := 0; fi84 < 15; fi84++ {
                if ((float64(((v5 - v5) * float32(-2.25))) >= -2.25) && true) {
                    var v85 uint32 = genFn1(125, 0)
                    switch ("foo bar" + ("" + ("foo bar" + ("a" + "hello")))) {
                    case "":
                        for fi86 := 0; fi86 < 7; fi86++ {
                            var v87 uint64 = 732
                            var v88 uint32 = v85
                            v87 -= v87
                            _ = fi86
                            _ = v87
                            _ = v88
                        }
                    case "hello":
                        var v89 uint16 = uint16(v85)
                        _ = v89
                    case "a":
                        v58 = v58
                    default:
                        var v90 string = "foo bar"
                        _ = v90
                    }
                    _ = v85
                }
                _ = fi84
            }
        case 4294967295:
            _ = len(v1)
        }
    }
    v55 += (v55 - (v55 * (v55 << ((255) & 63))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v58
    _ = v75
}
