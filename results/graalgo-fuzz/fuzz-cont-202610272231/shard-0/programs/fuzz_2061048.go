package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint; F2 int8 }
func genFn0() int {
    var v7 uint16 = 2
    for fi8 := 0; fi8 < 4; fi8++ {
        v7 += 895
        v9 := GenS1{F0: 0, F1: 627, F2: 5}
        v7 += (v7 / (((v7 / ((2) | 1))) | 1))
        _ = fi8
        _ = v9
    }
    v10 := []int{}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    switch "a" {
    case "hello":
        v7 = v7
    case "foo bar":
        if (0.1 > 0.1) {
            switch (("hello" + "\n\t") + "hello") {
            case "foo bar":
                switch -972 {
                case 2147483647:
                    v7 = v7
                case 73:
                    if (!(!(0.0 <= (1.5 + -2.25)))) {
                        if len(v10) > 0 {
                            v10[0] = -1
                        }
                    }
                }
            case "":
                for fi11 := 0; fi11 < 13; fi11++ {
                    v7 += v7
                    if (720 == 302) {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    } else {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            v7 -= 32767
                            _ = fi12
                        }
                    }
                    _ = fi11
                }
            default:
                if true {
                    v7 = v7
                    switch (("hello" + "") + (("a" + "foo bar") + "hello")) {
                    case "":
                        var v13 int16 = 1000
                        _ = v13
                    case "a":
                        v7 += v7
                    case "foo bar":
                        for fi14 := 0; fi14 < 7; fi14++ {
                            var v15 uint = 239
                            for fi16 := 0; fi16 < 17; fi16++ {
                                if len(v10) > 0 {
                                    v10[0] = fi14
                                }
                                _ = fi16
                            }
                            _ = fi14
                            _ = v15
                        }
                    }
                }
            }
            v17 := GenS2{F0: 1}
            _ = v17
        } else {
            var v18 uint8 = 46
            _ = v18
        }
    }
    _ = v10
    _ = v7
    return -95
}
func genFn1(p0 int16) int {
    var v19 int16 = p0
    v20 := GenS3{F0: true}
    _ = p0
    _ = v19
    _ = v20
    return -198
}

func main() {
    var v1 uint32 = 539
    var v2 float64 = (1e10 / 1.5)
    var v3 int32 = -988
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    var v6 int64 = -9223372036854775808
    var v21 uint32 = (^(v1))
    var v22 int8 = -63
    delete(v4, "a")
    fmt.Println(v3, v3)
    fmt.Println(v6)
    if true {
        switch -699 {
        case 121:
            var v23 int64 = v6
            _ = v23
        case 167:
            _ = len(v4)
        case 42:
            v24 := []int{127, -2147483648, -1}
            _ = v24
        default:
            fmt.Println(v22)
        }
        if true {
            v2 -= (3.14159 / v2)
            switch -2147483648 {
            case 2147483647:
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 int32 = (v3 << ((59) & 63))
                    v27 := GenS3{F0: (!(-2147483648 != (v3 ^ 801)))}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            case 9223372036854775807:
                for fi28 := 0; fi28 < 16; fi28++ {
                    v29 := GenS4{F0: 42}
                    v30 := []int{(fi28 ^ (fi28 % (((fi28 % ((0) | 1))) | 1)))}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            case 827:
                var v31 float64 = (0.1 - v2)
                _ = v31
            default:
                v4["a"] = genFn0()
            }
        } else {
            var v32 uint = 42
            _ = v32
        }
    }
    v33 := []int{}
    if false {
        for fi34 := 0; fi34 < 2; fi34++ {
            switch "" {
            case "hello":
                v35 := GenS5{F0: 42, F1: 294, F2: (v22 % ((v22) | 1))}
                _ = v35
            default:
                var v36 uint64 = 18446744073709551615
                _ = v36
            }
            switch (fi34 % (((fi34 + fi34)) | 1)) {
            case 4294967295:
                v2 += 1.5
            case -9:
                for fi37 := 0; fi37 < 6; fi37++ {
                    var v38 int64 = v6
                    v22 = -128
                    _ = fi37
                    _ = v38
                }
            case -13:
                var v39 int32 = v3
                _ = v39
            default:
                fmt.Println(v2, v3, v21)
            }
            _ = fi34
        }
        fmt.Println(v2)
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 string = ""
        v42 := GenS2{F0: 65535}
        _ = fi40
        _ = v41
        _ = v42
    }
    v1 -= 1
    v5[(v1 | 0)] = 1
    if len(v33) > 0 {
        v33[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
