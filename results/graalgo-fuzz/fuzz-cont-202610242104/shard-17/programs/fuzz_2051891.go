package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 uint8; F1 bool; F2 bool }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS8 struct { F0 int64; F1 uint32; F2 uint }
type GenS9 struct { F0 uint8; F1 uint16 }
func genFn0() int {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 string = ""
        v7 = ""
        v7 = v7
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 2147483648
    _ = v8
    return -1
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 42
    v3 := make(map[uint32]string)
    var v4 int64 = 590
    v5 := []string{"a", "foo bar", "foo bar", (("hello" + "a") + (("foo bar" + "") + "\n\t"))}
    v4 = 0
    _ = len(v5)
    if true {
        v9 := GenS1{F0: (v2 ^ 65535), F1: 42, F2: -47}
        _ = v9
    }
    v10 := GenS1{F0: (v2 - (v2 >> (((v2 % ((1) | 1))) & 63))), F1: 557, F2: 1}
    if (1 != 127) {
        switch genFn0() {
        case -9223372036854775808:
            fmt.Println(v2, v4)
        default:
            var v11 bool = false
            _ = v11
        }
        switch 741 {
        case -9223372036854775808:
            v12 := GenS2{F0: 127, F1: (((1.5 - (3.14159 - 3.14159)) / 1.5) >= (1e10 * (1.5 - 3.14159))), F2: (true || true)}
            _ = v12
        case 1:
            v5 = append(v5, "\n\t")
        default:
            v2 = (^(v2))
        }
    }
    switch "\n\t" {
    case "foo bar":
        var v13 uint64 = 1
        _ = v13
    case "hello":
        var v14 int16 = -1
        _ = v14
    case "\n\t":
        fmt.Println(v4, v4)
    default:
        v15 := GenS2{F0: 2, F1: ((!(false || (127 != 0))) && false), F2: (!(1e10 < -2.25))}
        _ = v15
    }
    v16 := GenS3{F0: v2, F1: 127}
    if true {
        v1["a"] = genFn0()
        v17 := GenS4{F0: 2}
        _ = v17
    } else {
        v2 -= v2
        if (true && false) {
            for fi18 := 0; fi18 < 6; fi18++ {
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 float32 = 1e10
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            if (false || true) {
                var v21 int64 = (v4 ^ (-(v4)))
                v22 := v1["a"]
                _ = v21
                _ = v22
            }
        }
    }
    fmt.Println(v2, v4)
    v23 := GenS4{F0: 251}
    if (2 != 1) {
        v24 := GenS5{F0: ("a" + ""), F1: v4}
        switch "\n\t" {
        case "a":
            switch (("hello" + "\n\t") + "") {
            case "":
                var v25 float64 = 0.0
                _ = v25
            }
        case "foo bar":
            v26 := GenS5{F0: (("hello" + "hello") + "\n\t"), F1: 9223372036854775807}
            _ = v26
        default:
            switch "hello" {
            case "\n\t":
                for fi27 := 0; fi27 < 14; fi27++ {
                    v3[4294967295] = "foo bar"
                    _ = fi27
                }
            case "":
                v2 = (^(v2))
            case "foo bar":
                v4 = (v4 / ((9223372036854775807) | 1))
            default:
                var v28 bool = false
                _ = v28
            }
        }
        _ = v24
    } else {
        v29 := GenS6{F0: 514}
        switch genFn0() {
        case 32767:
            fmt.Println(v4, v2)
        case 9223372036854775807:
            v4 += (v4 >> (((v2 | 1)) & 63))
        case 498:
            v4 += -844
        default:
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 string = "foo bar"
                switch (fi30 >> ((1) & 63)) {
                case 128:
                    v32 := GenS7{F0: 65535, F1: (1 != 246), F2: 4294967295}
                    _ = v32
                case 0:
                    switch (fi30 % (((fi30 / ((fi30) | 1))) | 1)) {
                    case 0:
                        v31 = v31
                    default:
                        switch v31 {
                        case "":
                            v2 = v2
                        case "a":
                            if ((false || (-2.25 <= ((1.5 / (float64((0.1 - ((3.14159 * (1e10 + 0.0)) + 0.1))) - float64(float32(float64(0.0))))) * 1.5))) || false) {
                                switch (-(fi30)) {
                                case -1:
                                    v4 += v4
                                case -699:
                                    switch (v31 + (("\n\t" + ("a" + "\n\t")) + ("a" + v31))) {
                                    case "":
                                        if true {
                                            v4 += v4
                                            var v33 float64 = 0.1
                                            for fi34 := 0; fi34 < 15; fi34++ {
                                                var v35 bool = (42 <= 1)
                                                _ = len(v5)
                                                v33 = -2.25
                                                _ = fi34
                                                _ = v35
                                            }
                                            _ = v33
                                        }
                                    case "a":
                                        if len(v5) > 0 {
                                            v5[0] = v31
                                        }
                                    }
                                default:
                                    v36 := GenS2{F0: 127, F1: (!false), F2: (-2147483648 == 2147483647)}
                                    _ = v36
                                }
                            } else {
                                for fi37 := 0; fi37 < 7; fi37++ {
                                    v4 = (v4 % ((-999) | 1))
                                    if true {
                                        if false {
                                            var v38 uint64 = 4294967295
                                            if (true || (true && false)) {
                                                fmt.Println(fi37, v2)
                                                var v39 int32 = 0
                                                v40 := GenS8{F0: (v4 * v4), F1: 4294967295, F2: v2}
                                                _ = v39
                                                _ = v40
                                            }
                                            v41 := GenS9{F0: 0, F1: 563}
                                            _ = v38
                                            _ = v41
                                        } else {
                                            fmt.Println(v2, fi37, v31)
                                            if (false || false) {
                                                var v42 uint = (^(v2))
                                                var v43 bool = (!false)
                                                fmt.Println(v2, fi30)
                                                if (0 != 1000) {
                                                    v31 = (((v31 + v31) + "hello") + ((v31 + ("\n\t" + (("" + (("hello" + v31) + "a")) + v31))) + "\n\t"))
                                                    fmt.Println(v43)
                                                    fmt.Println(v31)
                                                } else {
                                                    var v44 uint8 = 1
                                                    switch (fi37 * (fi30 / (((fi37 | int(v42))) | 1))) {
                                                    case -292:
                                                        var v45 int8 = -96
                                                        _ = v45
                                                    }
                                                    _ = v44
                                                }
                                                _ = v42
                                                _ = v43
                                            } else {
                                                v2 -= (v2 + (v2 >> (((v2 | (v2 >> ((v2) & 63)))) & 63)))
                                                var v46 uint8 = 61
                                                _ = v46
                                            }
                                            switch ((v31 + (v31 + v31)) + v31) {
                                            case "a":
                                                var v47 uint8 = 205
                                                _ = v47
                                            case "\n\t":
                                                var v48 int16 = -824
                                                _ = v48
                                            case "hello":
                                                var v49 int8 = -1
                                                _ = v49
                                            }
                                        }
                                        _ = len(v3)
                                        v50 := GenS6{F0: genFn0()}
                                        _ = v50
                                    } else {
                                        if (!(v4 < (v4 * -704))) {
                                            _ = len(v5)
                                            if (v31 != v31) {
                                                fmt.Println(v2, v2)
                                            } else {
                                                if (false && true) {
                                                    fmt.Println(v31, fi37)
                                                    v4 = (v4 ^ (v4 | -9223372036854775808))
                                                }
                                                var v51 uint16 = 0
                                                var v52 int = 677
                                                fmt.Println(v52, v2)
                                                if false {
                                                    _ = len(v3)
                                                } else {
                                                    var v53 int = (fi37 ^ genFn0())
                                                    _ = v53
                                                }
                                                _ = v51
                                                _ = v52
                                            }
                                            var v54 int16 = 432
                                            _ = v54
                                        }
                                        if ((!true) || (0 == -32768)) {
                                            switch ((v31 + v31) + ("" + v31)) {
                                            case "\n\t":
                                                var v55 int16 = -1
                                                _ = v55
                                            case "foo bar":
                                                v2 = (^(v2))
                                            case "hello":
                                                v56 := GenS3{F0: 65535, F1: 128}
                                                var v57 int16 = -32768
                                                var v58 uint64 = 881
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            default:
                                                fmt.Println(fi30, v31)
                                            }
                                            delete(v3, 587)
                                        } else {
                                            switch -2147483648 {
                                            case -9223372036854775808:
                                                var v59 uint32 = 2
                                                _ = v59
                                            case 0:
                                                _ = len(v1)
                                            default:
                                                v31 = v31
                                            }
                                            var v60 int = genFn0()
                                            v31 = v31
                                            _ = v60
                                        }
                                    }
                                    v31 = v31
                                    _ = fi37
                                }
                                fmt.Println(fi30, v2, v31)
                                switch ((v31 + (v31 + "hello")) + "\n\t") {
                                case "a":
                                    fmt.Println(v31)
                                case "foo bar":
                                    v2 += 0
                                case "\n\t":
                                    var v61 uint8 = 127
                                    _ = v61
                                }
                            }
                        default:
                            switch genFn0() {
                            case 32768:
                                fmt.Println(fi30, v4, fi30)
                            case 290:
                                v31 = "a"
                            case 9223372036854775807:
                                var v62 uint8 = 197
                                _ = v62
                            }
                        }
                    }
                }
                v63 := GenS9{F0: 0, F1: 255}
                _ = fi30
                _ = v31
                _ = v63
            }
        }
        switch 42 {
        case 9223372036854775807:
            v4 += v4
        case 32767:
            fmt.Println(v4, v2)
        default:
            _ = len(v5)
        }
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
