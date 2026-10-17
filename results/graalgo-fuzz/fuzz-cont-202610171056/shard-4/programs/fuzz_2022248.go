package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint16; F1 string; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 string; F2 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int64; F1 uint8; F2 bool }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 uint; F1 uint }
func genFn0(p0 int) int {
    v4 := []int{p0, 498, p0, (p0 ^ -2147483648)}
    _ = len(v4)
    switch (("foo bar" + ("\n\t" + "a")) + "") {
    case "foo bar":
        var v5 float32 = ((float32(1.5) / -2.25) + 3.14159)
        _ = v5
    case "\n\t":
        p0 += (p0 + (p0 ^ 32767))
    case "hello":
        fmt.Println(p0)
    default:
        var v6 int16 = -1
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 bool = (1 < 133)
        var v9 int64 = -391
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v4
    return (p0 | (p0 + 128))
}
func genFn1(p0 int16) int8 {
    v10 := []int{2147483647, 34}
    v10 = append(v10, 127)
    for fi11 := 0; fi11 < 3; fi11++ {
        fmt.Println(p0, p0)
        _ = fi11
    }
    _ = p0
    _ = v10
    return -1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 float64 = (1e10 / 3.14159)
    switch "" {
    case "a":
        delete(v1, 32768)
    case "foo bar":
        for fi12 := 0; fi12 < 15; fi12++ {
            v3 = 0.1
            if (true || true) {
                v3 = (v3 * 0.0)
                v13 := GenS1{F0: 0}
                for fi14 := 0; fi14 < 2; fi14++ {
                    switch ("foo bar" + "\n\t") {
                    case "":
                        v15 := v2[994]
                        _ = v15
                    case "hello":
                        var v16 int8 = 100
                        _ = v16
                    case "\n\t":
                        v3 -= v3
                    }
                    fmt.Println(fi14, v3, fi14)
                    _ = fi14
                }
                _ = v13
            }
            v17 := v2[42]
            _ = fi12
            _ = v17
        }
    case "\n\t":
        switch ("foo bar" + "a") {
        case "hello":
            switch (("a" + "a") + "") {
            case "foo bar":
                if true {
                    v3 = 1.5
                    v18 := GenS1{F0: 255}
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 uint32 = 927
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                }
            case "a":
                for fi21 := 0; fi21 < 15; fi21++ {
                    switch "\n\t" {
                    case "foo bar":
                        if (false && (float64(0.0) <= -2.25)) {
                            v3 += v3
                        }
                    default:
                        var v22 bool = ((0 <= 201) && (!(2 > 993)))
                        _ = v22
                    }
                    _ = fi21
                }
            case "":
                switch "hello" {
                case "foo bar":
                    v3 -= 3.14159
                }
            }
        case "foo bar":
            v3 -= (v3 / -2.25)
        default:
            var v23 int64 = 690
            _ = v23
        }
    }
    v3 = (((v3 * (-2.25 / 0.1)) * float64((((3.14159 - (0.0 + (((float32((v3 * v3)) + 3.14159) / 3.14159) + 1e10))) - -2.25) / (0.1 - ((-2.25 + 1e10) - float32(v3)))))) - (((v3 / ((v3 * 1e10) - (v3 + 1.5))) - (float64((float32((1.5 * v3)) + 1e10)) / (v3 / 1.0))) / v3))
    var v24 uint = 118
    var v25 int16 = 0
    fmt.Println(v25)
    v25 += -1
    var v26 uint64 = 42
    v26 = v26
    switch int(v24) {
    case 527:
        for fi27 := 0; fi27 < 12; fi27++ {
            fmt.Println(fi27, v3, v24)
            _ = fi27
        }
    case 127:
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 string = ("\n\t" + ("a" + "\n\t"))
            for fi30 := 0; fi30 < 17; fi30++ {
                v31 := GenS2{F0: 42, F1: 1}
                v24 = v24
                fmt.Println(v26)
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 19; fi32++ {
                v1[genFn0(802)] = false
                v33 := GenS3{F0: (fi28 * fi28), F1: (fi28 / ((128) | 1))}
                v26 += v26
                _ = fi32
                _ = v33
            }
            _ = fi28
            _ = v29
        }
    }
    delete(v2, 2147483647)
    v26 += 524
    if false {
        if (156 <= 224) {
            var v34 uint32 = 2147483648
            v35 := GenS4{F0: 162, F1: "a", F2: (v24 & 42)}
            _ = v34
            _ = v35
        }
        var v36 string = ""
        _ = v36
    } else {
        for fi37 := 0; fi37 < 15; fi37++ {
            fmt.Println(v25, v3, v25)
            switch ("a" + ("foo bar" + ((("foo bar" + ("foo bar" + "\n\t")) + "hello") + ((("\n\t" + ("" + "hello")) + "foo bar") + "a")))) {
            case "\n\t":
                v38 := GenS5{F0: fi37}
                _ = v38
            case "hello":
                v3 += v3
            case "a":
                for fi39 := 0; fi39 < 16; fi39++ {
                    if (!true) {
                        if false {
                            var v40 float32 = (1e10 - (0.1 * float32(3.14159)))
                            v3 = float64(((v40 + 1e10) * float32(0.0)))
                            _ = v40
                        } else {
                            v26 += 2
                        }
                        if (float32(3.14159) < 1.5) {
                            v25 += v25
                        }
                        v24 = v24
                    }
                    v41 := GenS6{F0: "a", F1: "\n\t", F2: 4294967295}
                    var v42 int64 = -1
                    _ = fi39
                    _ = v41
                    _ = v42
                }
            }
            var v43 float32 = -2.25
            _ = fi37
            _ = v43
        }
        var v44 int = 9223372036854775807
        if (1e10 <= float32(0.0)) {
            v45 := GenS1{F0: 0}
            v46 := GenS2{F0: int64(v24), F1: (v24 + v24)}
            var v47 int16 = (v25 + v25)
            _ = v45
            _ = v46
            _ = v47
        } else {
            var v48 float64 = v3
            if false {
                if true {
                    v49 := GenS7{F0: 1}
                    _ = v49
                } else {
                    var v50 float64 = v3
                    v24 += v24
                    _ = v50
                }
                var v51 bool = (2 == 1)
                _ = v51
            } else {
                if false {
                    v52 := v2[42]
                    switch ((("foo bar" + "\n\t") + ("" + "hello")) + (("hello" + "foo bar") + "hello")) {
                    case "a":
                        var v53 uint8 = 129
                        _ = v53
                    case "foo bar":
                        if true {
                            v52 -= (v52 << (((^(v24))) & 63))
                            fmt.Println(v25)
                            for fi54 := 0; fi54 < 19; fi54++ {
                                v25 -= 32767
                                var v55 uint8 = 84
                                _ = fi54
                                _ = v55
                            }
                        }
                    }
                    v56 := GenS5{F0: (v44 * -191)}
                    _ = v52
                    _ = v56
                } else {
                    _ = len(v1)
                }
                if true {
                    for fi57 := 0; fi57 < 4; fi57++ {
                        fmt.Println(fi57)
                        _ = fi57
                    }
                    v58 := GenS8{F0: 1, F1: 1, F2: false}
                    v59 := GenS9{F0: 1, F1: 595}
                    _ = v58
                    _ = v59
                } else {
                    if false {
                        for fi60 := 0; fi60 < 12; fi60++ {
                            v25 -= (v25 >> ((v24) & 63))
                            var v61 float32 = ((0.1 + 3.14159) / 1e10)
                            _ = fi60
                            _ = v61
                        }
                    }
                }
            }
            for fi62 := 0; fi62 < 13; fi62++ {
                for fi63 := 0; fi63 < 12; fi63++ {
                    var v64 uint32 = 410
                    _ = fi63
                    _ = v64
                }
                _ = fi62
            }
            _ = v48
        }
        _ = v44
    }
    fmt.Println(v25, v25)
    if false {
        if ((42 != (^(v26))) && (-523 >= -176)) {
            _ = len(v1)
            switch "hello" {
            case "\n\t":
                var v65 uint64 = 2
                _ = v65
            case "hello":
                if true {
                    var v66 uint64 = (^(v26))
                    var v67 int16 = 32767
                    _ = v66
                    _ = v67
                }
            }
            var v68 int8 = 100
            _ = v68
        } else {
            if (v24 != (v24 + v24)) {
                var v69 int64 = -1
                v70 := v2[4294967295]
                if (!false) {
                    var v71 int32 = 2147483647
                    var v72 int16 = 529
                    if (false && (65535 <= 65535)) {
                        _ = len(v1)
                        for fi73 := 0; fi73 < 10; fi73++ {
                            var v74 bool = true
                            if (v74 && ((0.1 <= v3) && (v74 && (v74 && v74)))) {
                                fmt.Println(fi73, v74, v71)
                                if true {
                                    v25 = -187
                                    fmt.Println(v3)
                                    var v75 int8 = 14
                                    _ = v75
                                }
                            } else {
                                for fi76 := 0; fi76 < 6; fi76++ {
                                    v77 := GenS3{F0: (fi76 * -2147483648), F1: (fi73 & genFn0((fi73 & (-(fi76)))))}
                                    v78 := GenS10{F0: v24, F1: 2}
                                    _ = fi76
                                    _ = v77
                                    _ = v78
                                }
                            }
                            _ = fi73
                            _ = v74
                        }
                        v3 = float64(3.14159)
                    }
                    _ = v71
                    _ = v72
                } else {
                    v26 += v26
                    v79 := GenS7{F0: v69}
                    fmt.Println(v24, v25, v69)
                    _ = v79
                }
                _ = v69
                _ = v70
            } else {
                v24 += v24
                if (((-2.25 * (3.14159 / 0.1)) != 1.5) && ((v26 + 309) == (v26 << ((v24) & 63)))) {
                    var v80 uint8 = 42
                    var v81 float32 = (1.5 / (0.1 - float32(v3)))
                    _ = v80
                    _ = v81
                }
            }
            v3 += float64(3.14159)
            if (true && (false && true)) {
                v82 := GenS9{F0: 42, F1: 230}
                if (!(-2147483648 == 2147483647)) {
                    v83 := GenS4{F0: uint16(v24), F1: ("\n\t" + "foo bar"), F2: (v24 * 662)}
                    v24 = v24
                    _ = v83
                }
                _ = v82
            } else {
                fmt.Println(v26, v26)
                if (-1 > 683) {
                    var v84 float32 = 1e10
                    if true {
                        var v85 int8 = 1
                        var v86 int64 = 42
                        v84 -= (0.0 - 3.14159)
                        _ = v85
                        _ = v86
                    }
                    var v87 uint8 = 42
                    _ = v84
                    _ = v87
                }
            }
        }
    }
    var v88 float32 = 1.5
    v89 := []int{}
    fmt.Println(v26, v3, v3)
    fmt.Println(v25, v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v88:", v88)
    fmt.Println("slicelen v89:", len(v89), v89)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v88
    _ = v89
}
