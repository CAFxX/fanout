package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 string }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 bool; F2 int32 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 uint64; F1 int8 }
func genFn0(p0 int64) bool {
    if (-483 != 1) {
        var v6 uint16 = 65535
        _ = v6
    } else {
        switch -2147483648 {
        case 238:
            var v7 string = "hello"
            _ = v7
        case 32768:
            p0 = p0
        default:
            v8 := GenS1{F0: -1, F1: "foo bar", F2: (("hello" + ("" + "a")) + ("foo bar" + ""))}
            _ = v8
        }
        v9 := GenS1{F0: 452, F1: "foo bar", F2: ""}
        v10 := GenS1{F0: 0, F1: (("hello" + "") + ""), F2: "foo bar"}
        _ = v10
        _ = v9
    }
    switch 746 {
    case 954:
        var v11 int64 = p0
        _ = v11
    case 100:
        var v12 uint8 = 0
        _ = v12
    }
    var v13 uint8 = 255
    _ = p0
    _ = v13
    return false
}
func genFn1(p0 uint8, p1 int) int {
    switch p1 {
    case -956:
        if (830 != uint(p0)) {
            switch (("hello" + ("hello" + ("a" + "a"))) + ("hello" + ("foo bar" + ("" + "")))) {
            case "foo bar":
                p0 = 255
            default:
                v14 := []int{31, p1, p1, p1}
                _ = v14
            }
            var v15 float32 = ((1.5 / 1.0) - (1e10 + 0.1))
            if ((("foo bar" + "foo bar") <= "") || false) {
                var v16 uint8 = (^(p0))
                switch "" {
                case "foo bar":
                    p0 -= p0
                }
                _ = v16
            }
            _ = v15
        }
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        switch p1 {
        case 2147483647:
            switch (fi17 * (p1 - (p1 % ((2147483647) | 1)))) {
            case -9223372036854775808:
                for fi18 := 0; fi18 < 10; fi18++ {
                    fmt.Println(fi18)
                    p0 -= 86
                    _ = fi18
                }
            case 14:
                var v19 uint8 = 143
                _ = v19
            case -801:
                p1 -= 42
            default:
                for fi20 := 0; fi20 < 14; fi20++ {
                    fmt.Println(fi17, p0, p1)
                    var v21 int16 = -32768
                    _ = fi20
                    _ = v21
                }
            }
        case 664:
            v22 := []int{}
            _ = v22
        }
        _ = fi17
    }
    v23 := []int{(p1 * (p1 - p1)), (p1 << ((4294967295) & 63)), (p1 / (((p1 + p1)) | 1)), (^(p1))}
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    _ = p0
    _ = p1
    _ = v23
    return (-(p1))
}

func main() {
    var v1 uint16 = 2
    var v2 uint8 = 2
    var v3 int64 = -1
    v4 := make(map[uint32]bool)
    v5 := []string{("hello" + (((("\n\t" + "") + (("foo bar" + "a") + (("" + "hello") + "\n\t"))) + ("\n\t" + ("foo bar" + ("hello" + "a")))) + "foo bar")), "hello"}
    for fi24 := 0; fi24 < 20; fi24++ {
        switch "foo bar" {
        case "hello":
            v3 += v3
        case "foo bar":
            switch fi24 {
            case 2147483647:
                v25 := GenS2{F0: v1, F1: (^(v2))}
                _ = v25
            case 306:
                for fi26 := 0; fi26 < 9; fi26++ {
                    if true {
                        var v27 uint8 = v2
                        _ = v27
                    }
                    if (1e10 == 1e10) {
                        var v28 uint16 = (v1 - 0)
                        switch "hello" {
                        case "a":
                            v29 := GenS1{F0: 811, F1: "a", F2: "\n\t"}
                            _ = v29
                        }
                        _ = v28
                    }
                    _ = fi26
                }
            case 841:
                if false {
                    var v30 uint16 = (v1 ^ (v1 << ((937) & 63)))
                    if (0.0 <= 1.5) {
                        for fi31 := 0; fi31 < 7; fi31++ {
                            v30 -= (v1 << ((802) & 63))
                            switch 127 {
                            case 4294967295:
                                switch 4294967295 {
                                case 779:
                                    var v32 int64 = (^(v3))
                                    _ = v32
                                case -2147483648:
                                    v1 += 2
                                case -149:
                                    var v33 bool = false
                                    _ = v33
                                }
                            }
                            _ = fi31
                        }
                        v1 = uint16(v2)
                        for fi34 := 0; fi34 < 11; fi34++ {
                            v35 := GenS3{F0: v3}
                            _ = fi34
                            _ = v35
                        }
                    } else {
                        for fi36 := 0; fi36 < 11; fi36++ {
                            v3 = v3
                            _ = fi36
                        }
                    }
                    _ = v30
                }
            }
        default:
            fmt.Println(v3)
        }
        _ = fi24
    }
    var v37 int8 = 0
    var v38 uint8 = v2
    v38 = v2
    var v39 int = 42
    v39 = (^(v39))
    v1 = 42
    var v40 uint8 = v38
    var v41 int8 = 42
    for fi42 := 0; fi42 < 19; fi42++ {
        fmt.Println(v2, v38)
        delete(v4, 816)
        _ = len(v4)
        _ = fi42
    }
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    if (false || (float64(0.1) >= (-2.25 - 3.14159))) {
        v43 := GenS4{F0: v1, F1: true, F2: int32(v39)}
        v4[4294967295] = ((((0.0 / float64((0.1 * 3.14159))) / 1.0) * 3.14159) > 0.0)
        v5 = append(v5, ("\n\t" + "\n\t"))
        _ = v43
    } else {
        fmt.Println(v41, v39, v37)
        if true {
            var v44 int16 = -32768
            v45 := GenS5{F0: 100, F1: 879}
            _ = len(v4)
            _ = v44
            _ = v45
        } else {
            v5 = append(v5, "hello")
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (v40 > v38) {
                v4[764] = false
            }
        }
    }
    fmt.Println(v41, v39, v3)
    if false {
        delete(v4, 753)
        v46 := GenS6{F0: 1, F1: 1}
        if (18446744073709551615 < 65535) {
            v47 := v4[65535]
            fmt.Println(v38, v1, v1)
            var v48 int64 = (v3 & -9223372036854775808)
            _ = v47
            _ = v48
        } else {
            v37 = (v41 >> ((18446744073709551615) & 63))
            v1 += v1
            fmt.Println(v40)
        }
        _ = v46
    } else {
        switch v39 {
        case -290:
            v49 := GenS7{F0: -1, F1: (^(v2))}
            _ = v49
        case 4294967295:
            for fi50 := 0; fi50 < 11; fi50++ {
                v51 := v4[226]
                _ = fi50
                _ = v51
            }
        default:
            v52 := GenS1{F0: 0, F1: "hello", F2: "hello"}
            _ = v52
        }
        switch v39 {
        case 128:
            var v53 uint16 = (v1 / ((497) | 1))
            _ = v53
        case 32768:
            var v54 uint = 18446744073709551615
            _ = v54
        case 0:
            v55 := GenS8{F0: 2, F1: v41}
            _ = v55
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
