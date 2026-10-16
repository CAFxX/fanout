package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint64; F1 int32; F2 int }
type GenS4 struct { F0 uint64; F1 string; F2 int8 }
func genFn0() int32 {
    v5 := []int{1}
    if false {
        if true {
            var v6 int8 = 100
            _ = v6
        } else {
            var v7 bool = false
            _ = v7
        }
    } else {
        v8 := GenS1{F0: 194, F1: -942}
        switch "\n\t" {
        case "a":
            var v9 bool = false
            _ = v9
        default:
            var v10 uint16 = 65535
            _ = v10
        }
        if len(v5) > 0 {
            v5[0] = 32767
        }
        _ = v8
    }
    var v11 int = -377
    _ = v11
    _ = v5
    return 2147483647
}
func genFn1() string {
    v12 := GenS2{F0: -32768, F1: 65535}
    var v13 uint = 0
    _ = v12
    _ = v13
    return ("a" + "foo bar")
}

func main() {
    var v1 int64 = -515
    var v2 int32 = -1
    v3 := []uint32{2147483648, 2, 2147483648}
    v4 := []string{"hello", "", "foo bar"}
    switch 1 {
    case 2147483647:
        _ = len(v4)
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v2 -= genFn0()
        v15 := GenS1{F0: 0, F1: (^(v1))}
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 255, F1: (v1 % ((v1) | 1))}
    v1 -= v1
    fmt.Println(v1, v1)
    v3 = append(v3, 0)
    var v17 uint64 = 42
    if (0 <= 255) {
        for fi18 := 0; fi18 < 14; fi18++ {
            if false {
                v2 = v2
                v17 += (v17 + (v17 >> ((0) & 63)))
                v19 := GenS1{F0: 254, F1: 1}
                _ = v19
            }
            for fi20 := 0; fi20 < 17; fi20++ {
                switch "a" {
                case "foo bar":
                    v21 := GenS3{F0: (v17 * v17), F1: v2, F2: fi18}
                    _ = v21
                case "hello":
                    if false {
                        var v22 int = (fi20 ^ 898)
                        var v23 string = "hello"
                        var v24 float32 = 3.14159
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                }
                v2 += (v2 | (^(v2)))
                var v25 float64 = 1.5
                _ = fi20
                _ = v25
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi18
        }
        var v26 uint8 = 226
        for fi27 := 0; fi27 < 4; fi27++ {
            fmt.Println(v1, fi27)
            var v28 uint8 = 128
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    if true {
        if (0.1 >= 1.5) {
            fmt.Println(v17, v17, v2)
            if true {
                v1 += -295
            } else {
                if true {
                    _ = len(v3)
                    if false {
                        fmt.Println(v17, v2)
                        _ = len(v3)
                        for fi29 := 0; fi29 < 6; fi29++ {
                            for fi30 := 0; fi30 < 19; fi30++ {
                                switch (fi30 << ((42) & 63)) {
                                case -752:
                                    if false {
                                        var v31 int16 = -611
                                        _ = v31
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = ("" + ((("hello" + (("hello" + "") + "hello")) + ("" + "a")) + ("\n\t" + "\n\t")))
                                    }
                                case 128:
                                    v17 -= 838
                                case 1:
                                    v3 = append(v3, 2147483647)
                                default:
                                    v32 := GenS4{F0: 1, F1: (("hello" + "a") + "\n\t"), F2: 42}
                                    _ = v32
                                }
                                _ = fi30
                            }
                            switch 2147483647 {
                            case -2147483648:
                                fmt.Println(fi29, v1, v17)
                            case -172:
                                if true {
                                    v2 -= (v2 % ((1) | 1))
                                    for fi33 := 0; fi33 < 8; fi33++ {
                                        var v34 int = int(v2)
                                        if (0.0 != 1e10) {
                                            _ = len(v4)
                                            var v35 uint16 = 32767
                                            _ = v35
                                        }
                                        _ = fi33
                                        _ = v34
                                    }
                                }
                            case 42:
                                fmt.Println(fi29)
                            default:
                                v17 = v17
                            }
                            _ = fi29
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                        for fi36 := 0; fi36 < 7; fi36++ {
                            v17 += (v17 ^ (v17 >> ((255) & 63)))
                            v1 = (v1 / ((v1) | 1))
                            if true {
                                fmt.Println(v1)
                            }
                            _ = fi36
                        }
                    }
                    if ("foo bar" >= "foo bar") {
                        fmt.Println(v17, v1, v2)
                        v37 := GenS2{F0: 270, F1: 65535}
                        switch "\n\t" {
                        case "":
                            if (!(((true || (v17 < 42)) || (((0.0 - 1.5) >= 1e10) && (834 <= 65535))) || false)) {
                                v17 -= (^(v17))
                            } else {
                                for fi38 := 0; fi38 < 7; fi38++ {
                                    v39 := GenS4{F0: 0, F1: "hello", F2: 42}
                                    _ = fi38
                                    _ = v39
                                }
                                if true {
                                    var v40 uint64 = 42
                                    _ = v40
                                }
                            }
                        default:
                            var v41 uint = 4294967295
                            _ = v41
                        }
                        _ = v37
                    } else {
                        var v42 uint = uint(v17)
                        for fi43 := 0; fi43 < 14; fi43++ {
                            switch (fi43 ^ fi43) {
                            case 2147483647:
                                v17 -= 0
                            }
                            _ = fi43
                        }
                        _ = v42
                    }
                } else {
                    for fi44 := 0; fi44 < 1; fi44++ {
                        var v45 int64 = 42
                        fmt.Println(v17, v17, v1)
                        _ = fi44
                        _ = v45
                    }
                    var v46 int8 = 127
                    if false {
                        fmt.Println(v1, v2, v17)
                        if false {
                            v47 := GenS4{F0: (v17 << ((871) & 63)), F1: ("hello" + ("foo bar" + ("" + (("hello" + "a") + "a")))), F2: v46}
                            _ = v47
                        }
                    }
                    _ = v46
                }
                for fi48 := 0; fi48 < 16; fi48++ {
                    for fi49 := 0; fi49 < 17; fi49++ {
                        var v50 int16 = 0
                        v51 := GenS4{F0: (^(v17)), F1: ("a" + ""), F2: 22}
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    fmt.Println(v1)
                    _ = fi48
                }
                v3 = append(v3, 0)
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    if true {
        v52 := GenS1{F0: 42, F1: 1}
        switch "hello" {
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "a":
            if len(v3) > 0 {
                v3[0] = uint32(v17)
            }
        case "":
            v53 := GenS3{F0: v17, F1: 101, F2: 42}
            _ = v53
        }
        _ = v52
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        v55 := GenS1{F0: 191, F1: v1}
        var v56 uint8 = 0
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi54
        _ = v55
        _ = v56
    }
    fmt.Println(v17, v2)
    for fi57 := 0; fi57 < 9; fi57++ {
        fmt.Println(v17)
        _ = fi57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
