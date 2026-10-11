package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int16; F1 bool }
func genFn0() string {
    var v5 uint8 = 205
    var v6 string = ""
    v6 = (v6 + "hello")
    fmt.Println(v5)
    fmt.Println(v6, v5)
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 int32) bool {
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7, p0)
        _ = fi7
    }
    if ((!(!(4294967295 >= 65535))) && true) {
        if (!(1 > 296)) {
            var v8 int16 = 32767
            var v9 int8 = -13
            _ = v8
            _ = v9
        }
    }
    if true {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int16 = -32768
            var v12 string = ("" + "")
            v13 := GenS1{F0: (-(v11)), F1: 2}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        p0 += p0
    } else {
        switch "hello" {
        case "hello":
            fmt.Println(p0)
        case "\n\t":
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int64 = 276
                v16 := []int{}
                _ = fi14
                _ = v15
                _ = v16
            }
        case "foo bar":
            p0 = p0
        default:
            v17 := GenS2{F0: 0, F1: 127, F2: 42}
            _ = v17
        }
        var v18 int = 4294967295
        var v19 bool = (true && false)
        _ = v18
        _ = v19
    }
    v20 := []int{9223372036854775807, 127, 1}
    _ = p0
    _ = v20
    return true
}

func main() {
    var v1 uint8 = 255
    var v2 string = ("\n\t" + "foo bar")
    var v3 uint64 = 2
    var v4 int = 0
    v2 = v2
    if false {
        var v21 int64 = 0
        _ = v21
    } else {
        v22 := []int{(^(v4)), 1, 2147483647}
        _ = v22
    }
    v4 = v4
    if ((^(v1)) >= v1) {
        var v23 uint16 = 834
        _ = v23
    }
    var v24 uint16 = 867
    for fi25 := 0; fi25 < 18; fi25++ {
        v1 -= v1
        if (!(false && (v3 < 42))) {
            var v26 float32 = 0.0
            _ = v26
        } else {
            v27 := []int{fi25, fi25, fi25, (fi25 & fi25)}
            _ = v27
        }
        var v28 uint32 = 2147483647
        _ = fi25
        _ = v28
    }
    if false {
        v29 := []int{(v4 + -2147483648), v4}
        for fi30 := 0; fi30 < 11; fi30++ {
            v31 := GenS3{F0: "foo bar"}
            fmt.Println(v24, v24)
            _ = fi30
            _ = v31
        }
        _ = v29
    } else {
        if true {
            switch "hello" {
            case "\n\t":
                var v32 int64 = 9223372036854775807
                _ = v32
            default:
                var v33 int64 = 0
                _ = v33
            }
            var v34 string = ("a" + ("foo bar" + v2))
            v34 = (v2 + "hello")
            _ = v34
        } else {
            switch (v4 * v4) {
            case -961:
                if true {
                    if (((v24 % (((^(v24))) | 1)) > v24) && false) {
                        var v35 string = "a"
                        _ = v35
                    }
                    var v36 uint64 = (v3 | (v3 % ((489) | 1)))
                    v36 += v3
                    _ = v36
                }
            case -354:
                v1 = (v1 << ((767) & 63))
            default:
                if (0 >= 255) {
                    v37 := []int{v4, (v4 | v4), (v4 / (((^(v4))) | 1)), (-(v4))}
                    var v38 int8 = -128
                    _ = v37
                    _ = v38
                } else {
                    var v39 string = ("hello" + ((v2 + v2) + v2))
                    if false {
                        v40 := []int{}
                        switch 32767 {
                        case 128:
                            v2 = v39
                        case -1:
                            fmt.Println(v4, v1)
                        case 32768:
                            fmt.Println(v1, v39, v24)
                        default:
                            v1 = 1
                        }
                        v3 += 1
                        _ = v40
                    } else {
                        fmt.Println(v2)
                        v41 := []int{v4, (v4 + v4), 591}
                        _ = v41
                    }
                    _ = v39
                }
            }
            if true {
                for fi42 := 0; fi42 < 18; fi42++ {
                    var v43 int32 = 77
                    var v44 uint = 42
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            }
            if false {
                var v45 uint32 = 42
                var v46 uint32 = v45
                _ = v45
                _ = v46
            }
        }
        v47 := GenS2{F0: 389, F1: 42, F2: v24}
        if false {
            v24 -= 255
            switch -2147483648 {
            case 0:
                fmt.Println(v4, v2)
            case -1:
                switch "foo bar" {
                case "foo bar":
                    v48 := GenS1{F0: 1, F1: (v3 & 18446744073709551615)}
                    _ = v48
                case "hello":
                    v49 := []int{(v4 ^ (-(v4))), (v4 ^ -2147483648)}
                    _ = v49
                case "\n\t":
                    for fi50 := 0; fi50 < 11; fi50++ {
                        v51 := GenS1{F0: 1, F1: (v3 ^ v3)}
                        v3 = (v3 & (v3 + (v3 + 0)))
                        v52 := []int{(^(fi50))}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                }
            }
        } else {
            if (-1 > 32767) {
                var v53 uint8 = v1
                if false {
                    v3 += v3
                    switch v4 {
                    case 128:
                        v4 = v4
                    case 2147483647:
                        for fi54 := 0; fi54 < 13; fi54++ {
                            v3 -= (v3 / ((1) | 1))
                            fmt.Println(v4, v24, v53)
                            var v55 uint32 = 2
                            _ = fi54
                            _ = v55
                        }
                    }
                    switch v2 {
                    case "hello":
                        var v56 int = -650
                        _ = v56
                    case "":
                        v57 := GenS2{F0: 2147483648, F1: -128, F2: 0}
                        _ = v57
                    case "foo bar":
                        fmt.Println(v24, v1)
                    default:
                        switch v4 {
                        case 32767:
                            v4 = v4
                        }
                    }
                }
                for fi58 := 0; fi58 < 11; fi58++ {
                    fmt.Println(v53)
                    _ = fi58
                }
                _ = v53
            } else {
                for fi59 := 0; fi59 < 3; fi59++ {
                    v60 := GenS4{F0: ((((v2 + "") + v2) + (v2 + (v2 + v2))) + (v2 + "foo bar")), F1: 2}
                    _ = fi59
                    _ = v60
                }
                var v61 int8 = 42
                v1 = v1
                _ = v61
            }
            if true {
                var v62 float32 = float32(0.0)
                _ = v62
            }
        }
        _ = v47
    }
    v63 := []int{v4}
    v64 := GenS5{F0: 42, F1: true}
    for fi65 := 0; fi65 < 3; fi65++ {
        if (!(-9223372036854775808 <= 9223372036854775807)) {
            if (0.1 == -2.25) {
                v63 = append(v63, int(v3))
            } else {
                v2 = "hello"
            }
            v3 = (^(v3))
        } else {
            var v66 uint8 = (v1 + v1)
            var v67 int16 = 1
            _ = v66
            _ = v67
        }
        var v68 uint64 = (v3 >> ((2) & 63))
        _ = fi65
        _ = v68
    }
    v2 = ("hello" + v2)
    var v69 int16 = -32768
    var v70 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v63
    _ = v64
    _ = v69
    _ = v70
}
