package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 bool }
type GenS3 struct { F0 string; F1 uint32; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint32; F2 int }
func genFn0() uint16 {
    v7 := GenS1{F0: 284}
    var v8 uint64 = 4294967295
    fmt.Println(v8)
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 bool, p1 uint16, p2 int32) uint64 {
    if p0 {
        var v9 float32 = (3.14159 * 0.1)
        for fi10 := 0; fi10 < 12; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := GenS2{F0: -18, F1: (v9 < (float32((0.1 * 3.14159)) + v9)), F2: (520 == 4294967295)}
                if p0 {
                    if p0 {
                        if (("hello" != ("foo bar" + "hello")) && (!p0)) {
                            v13 := []int{}
                            v14 := GenS1{F0: 195}
                            _ = v13
                            _ = v14
                        }
                        p1 = p1
                        switch (("foo bar" + ((("a" + "foo bar") + "hello") + "")) + "") {
                        case "":
                            var v15 uint8 = 159
                            _ = v15
                        case "\n\t":
                            v16 := GenS3{F0: "foo bar", F1: 4294967295, F2: (!(!false))}
                            _ = v16
                        case "foo bar":
                            p2 -= 0
                            var v17 uint = 255
                            _ = v17
                        }
                    }
                    v9 += (3.14159 * 0.1)
                } else {
                    v18 := GenS2{F0: 127, F1: (!false), F2: p0}
                    if (1.5 <= 1.5) {
                        p1 -= p1
                    } else {
                        fmt.Println(fi10, p1, p2)
                        v19 := []int{(fi11 / ((fi11) | 1)), (fi11 & (fi11 % (((fi10 | fi10)) | 1)))}
                        if ((fi10 ^ fi11) == fi11) {
                            fmt.Println(p0)
                        }
                        var v20 uint64 = 2
                        _ = v19
                        _ = v20
                    }
                    _ = v18
                }
                _ = fi11
                _ = v12
            }
            if p0 {
                fmt.Println(p1, v9, p1)
            }
            _ = fi10
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 7; fi22++ {
                v23 := []int{128, (fi21 - (fi22 % ((fi22) | 1))), fi22, 42}
                var v24 uint8 = 255
                switch "a" {
                case "a":
                    fmt.Println(p0)
                case "foo bar":
                    v25 := GenS2{F0: 1, F1: (v24 != (^(v24))), F2: p0}
                    _ = v25
                case "hello":
                    var v26 int8 = -1
                    _ = v26
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            if ((p0 && p0) && p0) {
                if p0 {
                    fmt.Println(p1)
                    v27 := []int{-9223372036854775808, fi21}
                    if p0 {
                        p2 += (^(p2))
                        var v28 float32 = (((float32(0.1) - ((1e10 - v9) * (0.1 + -2.25))) - (v9 * v9)) - ((v9 * float32((1.5 / 1e10))) + 3.14159))
                        v28 -= ((v9 * 1e10) - float32((0.1 * float64(-2.25))))
                        _ = v28
                    }
                    _ = v27
                }
                fmt.Println(fi21, p1, p1)
                var v29 uint8 = 173
                _ = v29
            } else {
                for fi30 := 0; fi30 < 14; fi30++ {
                    v31 := []int{(fi30 & fi30), 4294967295, (fi21 - fi21)}
                    switch "" {
                    case "foo bar":
                        p0 = false
                    default:
                        switch "hello" {
                        case "a":
                            p0 = true
                        default:
                            if (!(3.14159 == 0.0)) {
                                var v32 uint32 = 386
                                var v33 int64 = 1
                                p1 = 32767
                                var v34 float32 = (v9 + 0.0)
                                var v35 uint8 = 128
                                _ = v32
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        }
                    }
                    if p0 {
                        switch "a" {
                        case "foo bar":
                            var v36 uint = 0
                            _ = v36
                        case "a":
                            v9 = (v9 / v9)
                        case "hello":
                            if len(v31) > 0 {
                                v31[0] = (fi21 * (fi21 >> ((586) & 63)))
                            }
                        }
                        v37 := GenS1{F0: 1}
                        if (true || p0) {
                            fmt.Println(fi21, v9, fi30)
                            p0 = true
                            var v38 string = ("hello" + "hello")
                            _ = v38
                        } else {
                            v9 = v9
                            var v39 uint8 = 127
                            _ = v39
                        }
                        var v40 uint64 = 1
                        if p0 {
                            var v41 uint = 115
                            v31 = append(v31, 42)
                            _ = v41
                        }
                        _ = v37
                        _ = v40
                    } else {
                        p1 = (p1 << ((1) & 63))
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                    }
                    var v42 int16 = -858
                    var v43 int = (fi30 - (fi21 >> ((2) & 63)))
                    p2 = (p2 + p2)
                    _ = fi30
                    _ = v31
                    _ = v42
                    _ = v43
                }
                v9 = 1e10
                if (!(uint8(p1) != 1)) {
                    var v44 uint8 = 42
                    var v45 int32 = p2
                    var v46 int16 = int16(fi21)
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 9; fi47++ {
                        v48 := GenS4{F0: (!(!((fi21 ^ 128) >= int(p1))))}
                        var v49 uint64 = 0
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                }
            }
            _ = fi21
        }
        _ = v9
    }
    var v50 bool = (!true)
    switch "hello" {
    case "a":
        var v51 int16 = -32768
        _ = v51
    case "foo bar":
        var v52 uint64 = 262
        _ = v52
    }
    v53 := []int{-553, 127, 42, -316}
    _ = p0
    _ = p1
    _ = p2
    _ = v50
    _ = v53
    return 445
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int16 = 32767
    var v3 int32 = -1
    v4 := make(map[uint32]int)
    var v5 uint64 = 412
    var v6 uint32 = v1
    v54 := GenS4{F0: true}
    v55 := []int{-1, int(v3)}
    v3 = v3
    v56 := GenS2{F0: 100, F1: (55 != 0), F2: false}
    fmt.Println(v1, v2)
    v57 := GenS2{F0: int8(v2), F1: ((65535 != (v1 << ((65535) & 63))) || (136 >= 255)), F2: false}
    var v58 float64 = 1e10
    for fi59 := 0; fi59 < 10; fi59++ {
        var v60 bool = true
        switch fi59 {
        case 2147483647:
            switch fi59 {
            case 128:
                var v61 bool = v60
                _ = v61
            case 672:
                v62 := GenS3{F0: ("hello" + ("a" + "a")), F1: (v1 >> ((2) & 63)), F2: ((9223372036854775807 == 1) || v60)}
                _ = v62
            default:
                v4[2147483647] = fi59
            }
        case 9223372036854775807:
            v55 = append(v55, fi59)
        default:
            if (127 >= -1) {
                if true {
                    _ = len(v55)
                    var v63 uint = 4294967295
                    v64 := GenS3{F0: (("hello" + "hello") + "foo bar"), F1: (v6 * (v6 / ((v6) | 1))), F2: false}
                    _ = v63
                    _ = v64
                } else {
                    fmt.Println(v3)
                }
                var v65 int64 = 644
                _ = v65
            }
        }
        _ = fi59
        _ = v60
    }
    var v66 int64 = -490
    if true {
        v67 := GenS3{F0: ("foo bar" + "\n\t"), F1: v6, F2: false}
        v68 := GenS5{F0: -1, F1: v1, F2: 1}
        _ = v67
        _ = v68
    }
    var v69 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
    _ = v66
    _ = v69
}
