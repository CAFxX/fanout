package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int32; F1 uint16; F2 int32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint; F1 int; F2 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int16; F1 int32; F2 uint }
type GenS9 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint32) int {
    var v4 uint = 121
    var v5 uint32 = p1
    v6 := GenS1{F0: "a"}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int16 = -32768
        v9 := []int{fi7, int(v8)}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -518
}

func main() {
    v1 := []uint32{500, 898, 532}
    var v2 uint16 = 255
    var v3 int16 = -1
    v10 := GenS2{F0: 42, F1: 32768}
    if false {
        v2 = (v2 * 32767)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v11 int64 = 0
        _ = v11
    } else {
        if (("hello" + ("hello" + "foo bar")) != "a") {
            if true {
                switch "foo bar" {
                case "\n\t":
                    var v12 uint64 = 0
                    _ = v12
                case "foo bar":
                    var v13 int = 9223372036854775807
                    _ = v13
                case "a":
                    switch 0 {
                    case -1:
                        switch 0 {
                        case 0:
                            for fi14 := 0; fi14 < 14; fi14++ {
                                var v15 int16 = (v3 ^ (v3 * v3))
                                v1 = append(v1, 4294967295)
                                var v16 string = ("foo bar" + "a")
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                        case -2147483648:
                            var v17 uint8 = 128
                            _ = v17
                        }
                    case -9223372036854775808:
                        v18 := GenS3{F0: -1, F1: v2, F2: 1}
                        _ = v18
                    case 128:
                        v19 := GenS4{F0: (!false), F1: 4294967295}
                        _ = v19
                    default:
                        v20 := GenS5{F0: 4294967295, F1: 32767, F2: ((1.5 * (1e10 / float64((1.5 + 3.14159)))) > 1.5)}
                        _ = v20
                    }
                default:
                    v21 := GenS6{F0: (^(v2))}
                    _ = v21
                }
                _ = len(v1)
                var v22 int8 = 1
                _ = v22
            }
            v23 := GenS6{F0: 2}
            var v24 uint16 = v2
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            v3 = 320
            fmt.Println(fi25, v2)
            for fi26 := 0; fi26 < 10; fi26++ {
                if (!true) {
                    if ("hello" >= "\n\t") {
                        v27 := GenS5{F0: 0, F1: (fi25 >> ((4294967295) & 63)), F2: (true || false)}
                        var v28 uint16 = 0
                        fmt.Println(v2)
                        _ = v27
                        _ = v28
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                    }
                } else {
                    if ((v2 % (((v2 - 442)) | 1)) < 1) {
                        v29 := GenS7{F0: true, F1: 42, F2: (v2 + v2)}
                        if false {
                            var v30 uint8 = 255
                            v3 = (v3 ^ (-(v3)))
                            v30 -= 251
                            if len(v1) > 0 {
                                v1[0] = 801
                            }
                            if false {
                                _ = len(v1)
                                var v31 int16 = v3
                                _ = v31
                            } else {
                                var v32 uint64 = 450
                                _ = v32
                            }
                            _ = v30
                        } else {
                            fmt.Println(v2, v2, v2)
                        }
                        v2 = v2
                        fmt.Println(v3)
                        _ = v29
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = len(v1)
                _ = fi26
            }
            _ = fi25
        }
        if (!(-212 > 128)) {
            if (-32768 >= 310) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v33 := GenS8{F0: 42, F1: 87, F2: 496}
                _ = v33
            }
            for fi34 := 0; fi34 < 16; fi34++ {
                v35 := GenS4{F0: (fi34 >= (-(fi34))), F1: 2147483648}
                _ = fi34
                _ = v35
            }
            v2 += 975
        } else {
            v2 = v2
            if false {
                var v36 float64 = -2.25
                _ = v36
            }
        }
    }
    var v37 uint64 = 0
    switch genFn0(18446744073709551615, 1) {
    case 9223372036854775807:
        _ = len(v1)
    case 937:
        var v38 int32 = 0
        _ = v38
    case -9223372036854775808:
        if true {
            if (-802 > -310) {
                v2 += (v2 << ((18446744073709551615) & 63))
            }
        }
    default:
        v37 -= (v37 / (((v37 / ((v37) | 1))) | 1))
    }
    var v39 int = -1
    fmt.Println(v39, v3)
    v40 := GenS9{F0: 1}
    var v41 uint64 = v37
    v42 := GenS1{F0: "\n\t"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
