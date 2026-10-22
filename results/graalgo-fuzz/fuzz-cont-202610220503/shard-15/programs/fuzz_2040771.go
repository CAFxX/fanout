package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int; F2 int }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64 }
func genFn0(p0 int16, p1 uint16, p2 int32) bool {
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int8 = int8(p1)
        p1 += (p1 >> ((64) & 63))
        v7 := GenS1{F0: 4294967295}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int32 = (p2 / ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (24 != 1)
}
func genFn1(p0 string, p1 bool) int16 {
    for fi9 := 0; fi9 < 15; fi9++ {
        p0 = p0
        v10 := []int{fi9, (fi9 / ((fi9) | 1))}
        p0 = p0
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := []int{}
        switch p0 {
        case "":
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 uint64 = 1
                v15 := GenS1{F0: 4294967295}
                _ = fi13
                _ = v14
                _ = v15
            }
        case "a":
            if false {
                var v16 float32 = 0.1
                v16 = v16
                _ = v16
            } else {
                p1 = ((p0 == p0) || p1)
                if (355 >= 1) {
                    switch ((p0 + (p0 + "foo bar")) + p0) {
                    case "hello":
                        switch (fi11 % ((-2147483648) | 1)) {
                        case 487:
                            p0 = ((p0 + p0) + p0)
                        default:
                            for fi17 := 0; fi17 < 4; fi17++ {
                                var v18 int = fi11
                                _ = fi17
                                _ = v18
                            }
                        }
                    }
                }
                if p1 {
                    fmt.Println(fi11, fi11)
                    var v19 int16 = 0
                    var v20 uint64 = 1
                    _ = v19
                    _ = v20
                } else {
                    switch -288 {
                    case 32768:
                        var v21 uint8 = 128
                        _ = v21
                    default:
                        _ = len(v12)
                    }
                    v22 := GenS1{F0: 1}
                    var v23 int8 = 0
                    _ = v22
                    _ = v23
                }
            }
        default:
            var v24 int16 = 395
            _ = v24
        }
        _ = fi11
        _ = v12
    }
    if ((((1e10 + (((0.0 + (float32(-2.25) / 1.5)) * float32((1e10 + 1e10))) - float32(0.0))) * (1.5 * 0.1)) + 3.14159) <= 0.1) {
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 int16 = 0
            var v27 uint16 = 2
            if true {
                switch -2147483648 {
                case -2147483648:
                    p0 = ""
                }
            }
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []string{"hello", "\n\t"}
    var v2 uint64 = 948
    var v3 uint16 = 506
    v4 := make(map[uint32]bool)
    var v28 uint = 2
    for fi29 := 0; fi29 < 5; fi29++ {
        _ = len(v4)
        _ = fi29
    }
    if (378 <= (^(v3))) {
        v30 := v4[426]
        fmt.Println(v30, v2, v2)
        _ = v30
    }
    if false {
        var v31 uint64 = (v2 - (v2 >> (((v28 >> (((v28 % (((^(v28))) | 1))) & 63))) & 63)))
        switch "hello" {
        case "a":
            var v32 uint64 = (v31 / ((uint64(v3)) | 1))
            _ = v32
        default:
            for fi33 := 0; fi33 < 13; fi33++ {
                var v34 uint64 = (v2 & v2)
                _ = fi33
                _ = v34
            }
        }
        v3 += (v3 ^ v3)
        _ = v31
    }
    v1 = append(v1, "a")
    if (1000 != -621) {
        v35 := GenS2{F0: (2 < v3)}
        _ = v35
    }
    if true {
        fmt.Println(v3, v2)
    } else {
        fmt.Println(v3)
    }
    v36 := GenS3{F0: 1, F1: -549, F2: -2147483648}
    var v37 int64 = 1
    if false {
        _ = len(v1)
        if false {
            if true {
                switch 1 {
                case 32767:
                    var v38 uint = (v28 % ((v28) | 1))
                    _ = v38
                case 131:
                    switch ((("a" + "\n\t") + "hello") + "\n\t") {
                    case "hello":
                        switch 0 {
                        case 32768:
                            v39 := GenS4{F0: (!true), F1: "foo bar"}
                            _ = v39
                        default:
                            v3 -= (v3 >> ((42) & 63))
                        }
                    case "\n\t":
                        v40 := GenS5{F0: 0, F1: 2}
                        _ = v40
                    case "foo bar":
                        var v41 int64 = (^(v37))
                        _ = v41
                    }
                default:
                    if false {
                        if ((v2 + v2) <= v2) {
                            for fi42 := 0; fi42 < 10; fi42++ {
                                switch 1 {
                                case 241:
                                    if true {
                                        v28 -= v28
                                        var v43 int8 = int8(fi42)
                                        if true {
                                            fmt.Println(v43, v37)
                                            for fi44 := 0; fi44 < 1; fi44++ {
                                                var v45 uint64 = v2
                                                switch -1 {
                                                case -2147483648:
                                                    if (true || true) {
                                                        v3 = 32767
                                                        v46 := GenS5{F0: -1, F1: 42}
                                                        if true {
                                                            if true {
                                                                var v47 uint = (v28 ^ (^(v28)))
                                                                v48 := v4[42]
                                                                switch ((("foo bar" + "foo bar") + "a") + ("foo bar" + "a")) {
                                                                case "hello":
                                                                    var v49 uint16 = (v3 ^ v3)
                                                                    _ = v49
                                                                case "foo bar":
                                                                    if (!(3.14159 != 0.0)) {
                                                                        _ = len(v4)
                                                                        switch (("a" + "") + "\n\t") {
                                                                        case "a":
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                var v50 bool = false
                                                                                _ = v50
                                                                            case "hello":
                                                                                v51 := GenS6{F0: (v3 << (((v47 >> ((255) & 63))) & 63)), F1: 32767}
                                                                                _ = v51
                                                                            default:
                                                                                if ("foo bar" == "hello") {
                                                                                    v2 = uint64(v28)
                                                                                    v43 -= (v43 ^ v43)
                                                                                    var v52 uint32 = 2147483648
                                                                                    _ = v52
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v53 := GenS7{F0: "hello"}
                                                                        v3 -= 32767
                                                                        _ = v53
                                                                    }
                                                                }
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                            var v54 uint64 = (v2 % ((v2) | 1))
                                                            fmt.Println(fi42, v45)
                                                            v54 = (^(v45))
                                                            var v55 float64 = (1.5 - (1e10 * (0.0 / float64((1e10 - 3.14159)))))
                                                            _ = len(v1)
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        if false {
                                                            if (float64(-2.25) == 0.0) {
                                                                var v56 uint32 = 4294967295
                                                                if (((v56 * (v56 >> ((v28) & 63))) == v56) && false) {
                                                                    v28 -= (v28 | (v28 / ((v28) | 1)))
                                                                }
                                                                _ = v56
                                                            } else {
                                                                if (!(42 < 183)) {
                                                                    var v57 uint = (^(v28))
                                                                    _ = v57
                                                                } else {
                                                                    v37 = -1
                                                                    v58 := GenS2{F0: (-1 >= -32768)}
                                                                    _ = v58
                                                                }
                                                            }
                                                        }
                                                        _ = v46
                                                    } else {
                                                        var v59 uint = v28
                                                        _ = v59
                                                    }
                                                default:
                                                    v28 = (v28 % (((v28 | 0)) | 1))
                                                }
                                                _ = fi44
                                                _ = v45
                                            }
                                        } else {
                                            v37 += v37
                                        }
                                        _ = v43
                                    }
                                case 42:
                                    v2 += (v2 ^ v2)
                                }
                                var v60 bool = (2147483647 >= 65535)
                                var v61 int32 = int32(v37)
                                _ = fi42
                                _ = v60
                                _ = v61
                            }
                            switch 1 {
                            case 32767:
                                for fi62 := 0; fi62 < 18; fi62++ {
                                    v28 = v28
                                    v28 = (v28 >> ((v28) & 63))
                                    _ = fi62
                                }
                            }
                        } else {
                            v37 += v37
                        }
                        v2 -= (^(v2))
                    }
                }
            }
        }
        for fi63 := 0; fi63 < 1; fi63++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v64 int = -1
            _ = fi63
            _ = v64
        }
    } else {
        var v65 int32 = -819
        if false {
            v28 -= (v28 | uint(v3))
            v66 := GenS8{F0: 1}
            var v67 uint = (v28 - v28)
            _ = v66
            _ = v67
        } else {
            var v68 int32 = v65
            _ = v68
        }
        var v69 int8 = 88
        _ = v65
        _ = v69
    }
    v3 = 799
    switch 128 {
    case -69:
        for fi70 := 0; fi70 < 9; fi70++ {
            if (((0.0 - 3.14159) + 3.14159) > -2.25) {
                var v71 int8 = 25
                fmt.Println(v2, v2, fi70)
                _ = v71
            }
            var v72 int8 = 127
            v3 = 2
            _ = fi70
            _ = v72
        }
    case 42:
        fmt.Println(v3, v2)
    case -2147483648:
        v73 := v4[0]
        _ = v73
    default:
        delete(v4, 2147483647)
    }
    var v74 uint32 = 977
    fmt.Println(v2)
    v75 := GenS8{F0: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v74
    _ = v75
}
