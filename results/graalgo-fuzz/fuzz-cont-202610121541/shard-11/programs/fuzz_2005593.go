package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 bool; F1 int64; F2 string }
type GenS4 struct { F0 string; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint8; F1 bool; F2 int }

func main() {
    var v1 uint16 = 2
    var v2 int32 = -168
    v3 := make(map[uint32]uint32)
    var v4 uint16 = (v1 ^ (v1 | (^(v1))))
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    v7 := GenS1{F0: 42, F1: -32768, F2: "foo bar"}
    fmt.Println(v1)
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := v5[42]
        _ = fi8
        _ = v9
    }
    v10 := []int{}
    if false {
        switch "\n\t" {
        case "foo bar":
            var v11 uint32 = 0
            _ = v11
        case "\n\t":
            fmt.Println(v4)
        default:
            if true {
                v4 += v1
            } else {
                fmt.Println(v1, v2)
            }
        }
        var v12 int8 = -128
        _ = v12
    }
    var v13 uint8 = 1
    _ = len(v3)
    fmt.Println(v13, v13)
    v13 += 2
    switch -2147483648 {
    case 0:
        switch 128 {
        case 1:
            for fi14 := 0; fi14 < 14; fi14++ {
                if (!true) {
                    v15 := GenS1{F0: (fi14 % ((-425) | 1)), F1: 992, F2: "\n\t"}
                    if false {
                        var v16 uint8 = v13
                        if (true || true) {
                            switch fi14 {
                            case 32767:
                                fmt.Println(v13, v1, v4)
                            case 984:
                                var v17 float64 = 1e10
                                _ = v17
                            default:
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                            switch fi14 {
                            case -2147483648:
                                v18 := GenS1{F0: fi14, F1: -235, F2: ("a" + "hello")}
                                _ = v18
                            case 32767:
                                var v19 float32 = 3.14159
                                _ = v19
                            case 4294967295:
                                v4 += v1
                            default:
                                v20 := GenS2{F0: 704, F1: fi14}
                                _ = v20
                            }
                        }
                        _ = v16
                    } else {
                        var v21 uint64 = 0
                        v22 := GenS1{F0: (fi14 / ((fi14) | 1)), F1: 1000, F2: "hello"}
                        _ = v21
                        _ = v22
                    }
                    delete(v5, 369)
                    _ = v15
                }
                _ = fi14
            }
        default:
            var v23 float32 = (float32(-2.25) + 3.14159)
            _ = v23
        }
    case -944:
        for fi24 := 0; fi24 < 17; fi24++ {
            v25 := GenS3{F0: false, F1: 42, F2: "a"}
            for fi26 := 0; fi26 < 18; fi26++ {
                v10 = append(v10, (-(fi26)))
                if true {
                    v27 := GenS2{F0: int32(fi26), F1: fi26}
                    switch "\n\t" {
                    case "foo bar":
                        if ((false || (!(0 > 0))) || (!(uint16(v13) < 32767))) {
                            if (true && true) {
                                v4 += v4
                            }
                            _ = len(v10)
                        } else {
                            if (false || true) {
                                var v28 float32 = 1e10
                                fmt.Println(v28, fi24)
                                var v29 float64 = (1e10 - -2.25)
                                _ = v28
                                _ = v29
                            }
                        }
                    default:
                        var v30 int16 = -32768
                        _ = v30
                    }
                    switch (("foo bar" + "foo bar") + "foo bar") {
                    case "a":
                        switch fi26 {
                        case -1:
                            switch "\n\t" {
                            case "a":
                                var v31 int16 = -304
                                _ = v31
                            case "\n\t":
                                if (((true || true) || true) || false) {
                                    fmt.Println(fi26, fi24, fi26)
                                } else {
                                    fmt.Println(v4, fi24, v1)
                                    v10 = append(v10, (-(fi26)))
                                }
                            default:
                                if true {
                                    v32 := GenS2{F0: v2, F1: (-(fi24))}
                                    v33 := GenS2{F0: (-(v2)), F1: (fi24 >> ((4294967295) & 63))}
                                    _ = v32
                                    _ = v33
                                } else {
                                    v1 -= (v4 | v1)
                                    v10 = append(v10, (fi24 | fi26))
                                    var v34 int16 = -32768
                                    _ = v34
                                }
                            }
                            var v35 int8 = 1
                            _ = v35
                        case 491:
                            v4 += v4
                        case -289:
                            v4 += (v4 << ((747) & 63))
                        }
                    case "\n\t":
                        v36 := GenS4{F0: "foo bar", F1: 2147483648, F2: v2}
                        _ = v36
                    case "":
                        fmt.Println(v13, v2, v2)
                    }
                    _ = v27
                }
                _ = len(v5)
                _ = fi26
            }
            var v37 int8 = -123
            _ = fi24
            _ = v25
            _ = v37
        }
    case 773:
        for fi38 := 0; fi38 < 6; fi38++ {
            for fi39 := 0; fi39 < 15; fi39++ {
                v2 -= v2
                _ = fi39
            }
            _ = fi38
        }
    default:
        fmt.Println(v1, v2, v2)
    }
    var v40 uint = 2
    for fi41 := 0; fi41 < 20; fi41++ {
        fmt.Println(fi41)
        _ = fi41
    }
    if false {
        v2 -= (^(v2))
        if true {
            fmt.Println(v13, v4, v2)
            v40 += 65535
        } else {
            var v42 string = "a"
            var v43 float32 = -2.25
            _ = v42
            _ = v43
        }
        switch "" {
        case "a":
            v40 = 2
        case "\n\t":
            v2 += v2
        default:
            for fi44 := 0; fi44 < 19; fi44++ {
                var v45 uint16 = (v1 | (v4 >> (((v40 >> ((v40) & 63))) & 63)))
                v6[42] = ""
                _ = fi44
                _ = v45
            }
        }
    } else {
        switch "" {
        case "a":
            v46 := GenS1{F0: 42, F1: 42, F2: "\n\t"}
            _ = v46
        }
        fmt.Println(v13, v2)
    }
    switch (("a" + ("hello" + ("" + ("" + "\n\t")))) + "\n\t") {
    case "\n\t":
        for fi47 := 0; fi47 < 14; fi47++ {
            v40 += 276
            v1 = 595
            if (false && true) {
                fmt.Println(fi47)
            } else {
                var v48 int64 = 963
                switch fi47 {
                case -1:
                    if (false || (!true)) {
                        var v49 float64 = -2.25
                        v49 -= (3.14159 + v49)
                        for fi50 := 0; fi50 < 1; fi50++ {
                            _ = len(v3)
                            fmt.Println(fi50, v2, fi50)
                            v51 := GenS5{F0: v13, F1: ((-(v48)) <= -1), F2: fi47}
                            var v52 int16 = 724
                            var v53 bool = true
                            _ = fi50
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                        _ = v49
                    }
                }
                _ = v48
            }
            _ = fi47
        }
    default:
        v54 := GenS3{F0: false, F1: 42, F2: "\n\t"}
        _ = v54
    }
    var v55 uint = (v40 & (^(v40)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v55
    _ = v6
    _ = v7
}
