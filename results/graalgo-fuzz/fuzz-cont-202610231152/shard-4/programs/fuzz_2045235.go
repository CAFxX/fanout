package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int; F2 int8 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 int64; F1 uint8 }
func genFn0() int64 {
    var v5 int16 = 1
    var v6 uint = 307
    v5 = v5
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint64 = 444
    v2 := make(map[uint32]uint32)
    v3 := []string{(("a" + ("foo bar" + "a")) + ("\n\t" + (("hello" + "foo bar") + ("hello" + "hello")))), "a", "foo bar"}
    var v4 uint8 = 42
    var v7 int16 = 42
    var v8 int8 = -1
    fmt.Println(v8)
    switch "a" {
    case "foo bar":
        switch int(v4) {
        case 127:
            fmt.Println(v8)
        case 128:
            v4 -= v4
        }
    default:
        for fi9 := 0; fi9 < 5; fi9++ {
            switch (-(fi9)) {
            case 128:
                var v10 float32 = 1.5
                _ = v10
            }
            _ = len(v3)
            _ = fi9
        }
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int = 462
        if (!(-1 <= 0)) {
            var v13 int64 = 1
            for fi14 := 0; fi14 < 17; fi14++ {
                _ = len(v2)
                _ = fi14
            }
            if (("foo bar" + "") >= "foo bar") {
                var v15 int32 = -171
                switch "\n\t" {
                case "\n\t":
                    if false {
                        var v16 float32 = 3.14159
                        for fi17 := 0; fi17 < 7; fi17++ {
                            var v18 float64 = (1.5 + -2.25)
                            v19 := GenS1{F0: v8, F1: (v4 - (^(v4)))}
                            var v20 uint = 938
                            v21 := GenS1{F0: 14, F1: 5}
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        if false {
                            v22 := GenS1{F0: 0, F1: v4}
                            fmt.Println(v13, v12, v1)
                            for fi23 := 0; fi23 < 17; fi23++ {
                                v7 = v7
                                if false {
                                    v4 -= v4
                                } else {
                                    v24 := GenS2{F0: int16(v12), F1: 556}
                                    fmt.Println(fi23, v16)
                                    _ = v24
                                }
                                var v25 uint = 0
                                _ = fi23
                                _ = v25
                            }
                            _ = v22
                        }
                        _ = v16
                    } else {
                        for fi26 := 0; fi26 < 17; fi26++ {
                            switch fi26 {
                            case 80:
                                switch "hello" {
                                case "\n\t":
                                    switch 32767 {
                                    case 0:
                                        v27 := GenS3{F0: ("a" == ""), F1: (42 >= (v1 / (((v1 << ((610) & 63))) | 1)))}
                                        _ = v27
                                    case 127:
                                        v28 := v2[0]
                                        _ = v28
                                    }
                                case "hello":
                                    if (255 <= (v4 | v4)) {
                                        fmt.Println(fi26)
                                        v29 := GenS4{F0: 1}
                                        _ = v29
                                    } else {
                                        var v30 uint16 = 741
                                        v30 = (^(v30))
                                        _ = v30
                                    }
                                case "":
                                    v13 = (v13 << ((0) & 63))
                                }
                            case 543:
                                if true {
                                    v12 -= fi26
                                    var v31 uint64 = 973
                                    _ = v31
                                }
                                v32 := GenS4{F0: 2}
                                _ = v32
                            }
                            switch (v12 & 4294967295) {
                            case 118:
                                var v33 uint64 = 1
                                _ = v33
                            case 4294967295:
                                fmt.Println(v1, v13, v12)
                            case -1:
                                var v34 float32 = float32((1.5 / ((float64(1.5) - (1.5 - 1e10)) / (1.5 / (float64((1e10 * 1.5)) * 0.1)))))
                                _ = v34
                            default:
                                v35 := GenS1{F0: v8, F1: 79}
                                _ = v35
                            }
                            _ = fi26
                        }
                    }
                case "":
                    v13 -= 1
                default:
                    if true {
                        switch -661 {
                        case 9223372036854775807:
                            v13 -= (-(v13))
                        default:
                            var v36 uint16 = 552
                            _ = v36
                        }
                        delete(v2, 0)
                    }
                }
                _ = v15
            } else {
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 int16 = 42
                    if ((!(v8 == (v8 / ((v8) | 1)))) && false) {
                        var v39 int8 = (v8 % ((v8) | 1))
                        _ = v39
                    } else {
                        switch -1 {
                        case -9223372036854775808:
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                        case -2147483648:
                            if (false && true) {
                                v40 := GenS2{F0: (-(v7)), F1: 0}
                                var v41 float32 = 0.0
                                v42 := GenS1{F0: v8, F1: (v4 * (v4 * 85))}
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                        default:
                            v43 := GenS5{F0: 9223372036854775807, F1: fi37, F2: (-(v8))}
                            _ = v43
                        }
                        v44 := GenS6{F0: (^(v8)), F1: 158}
                        v45 := GenS7{F0: "", F1: 0}
                        _ = v44
                        _ = v45
                    }
                    _ = fi37
                    _ = v38
                }
                v2[42] = 0
            }
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    v46 := GenS7{F0: ("a" + "foo bar"), F1: uint32(v1)}
    for fi47 := 0; fi47 < 9; fi47++ {
        fmt.Println(v8, v1)
        fmt.Println(v8, v4)
        _ = len(v2)
        _ = fi47
    }
    var v48 int = 128
    v7 -= v7
    v49 := GenS1{F0: v8, F1: v4}
    var v50 float32 = 1.5
    for fi51 := 0; fi51 < 17; fi51++ {
        if false {
            v52 := GenS6{F0: v8, F1: 42}
            var v53 string = ("foo bar" + ("a" + "a"))
            var v54 uint64 = 1
            _ = v52
            _ = v53
            _ = v54
        }
        delete(v2, 0)
        fmt.Println(fi51, v50, v4)
        _ = fi51
    }
    for fi55 := 0; fi55 < 17; fi55++ {
        switch ((("a" + "") + (((("" + (("foo bar" + "hello") + "")) + "foo bar") + "a") + "foo bar")) + "") {
        case "\n\t":
            v4 = (v4 << ((42) & 63))
        default:
            v7 += 1000
        }
        var v56 string = "hello"
        var v57 uint8 = (v4 & v4)
        _ = fi55
        _ = v56
        _ = v57
    }
    if false {
        v58 := GenS1{F0: v8, F1: v4}
        switch (-(v48)) {
        case 730:
            v7 += -1
        case 32767:
            v59 := GenS8{F0: 0, F1: v4}
            _ = v59
        }
        _ = v58
    }
    v4 = 184
    var v60 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v48
    _ = v49
    _ = v50
    _ = v60
    _ = v7
    _ = v8
}
