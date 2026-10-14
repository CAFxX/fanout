package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
type GenS4 struct { F0 int32; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint8; F2 int }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 int64; F1 uint16; F2 bool }
func genFn0(p0 uint, p1 uint8, p2 int64) uint64 {
    var v7 uint8 = p1
    v8 := GenS1{F0: -1}
    p0 += uint(v7)
    v9 := []int{128, -999, 778}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int = 42
    var v2 string = "a"
    var v3 uint8 = 1
    var v4 float64 = 1.5
    var v5 string = "hello"
    var v6 uint32 = 127
    if true {
        v4 += v4
    } else {
        var v10 uint = 255
        var v11 float32 = 0.1
        var v12 float64 = v4
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 string = (v2 + "a")
    switch v13 {
    case "foo bar":
        if (42 <= 184) {
            for fi14 := 0; fi14 < 2; fi14++ {
                if (v6 >= 2147483648) {
                    var v15 int16 = 42
                    fmt.Println(v13, v3, v6)
                    _ = v15
                }
                _ = fi14
            }
        } else {
            v16 := []int{}
            _ = v16
        }
    default:
        if ((!(-258 == (v1 - v1))) || ((v6 + 2) <= (v6 * (v6 + 2)))) {
            v17 := GenS2{F0: false, F1: (v6 - v6), F2: 9223372036854775807}
            v4 -= 0.1
            v2 = "foo bar"
            _ = v17
        } else {
            v18 := []int{-135, (v1 ^ (-(v1))), 32767}
            v19 := GenS3{F0: ((v3 % ((v3) | 1)) != (v3 | (^(v3)))), F1: v6, F2: "hello"}
            _ = v18
            _ = v19
        }
    }
    if true {
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS4{F0: -2147483648, F1: 32767, F2: 2147483647}
            _ = fi20
            _ = v21
        }
        v22 := []int{v1, (v1 ^ v1), v1, v1}
        v22 = append(v22, -656)
        _ = v22
    }
    v23 := []int{v1, v1}
    var v24 uint32 = v6
    if true {
        if true {
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 bool = true
                if true {
                    var v27 float32 = 0.0
                    v28 := GenS5{F0: 255, F1: (v3 << ((458) & 63)), F2: 1}
                    _ = len(v23)
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v2 = v2
                        _ = fi29
                    }
                    var v30 uint8 = v3
                    _ = v30
                }
                var v31 uint16 = 140
                _ = fi25
                _ = v26
                _ = v31
            }
        } else {
            v23 = append(v23, (^(v1)))
            var v32 float32 = ((0.0 - 0.0) - 0.1)
            _ = v32
        }
        var v33 uint16 = 255
        _ = v33
    }
    v23 = append(v23, v1)
    for fi34 := 0; fi34 < 6; fi34++ {
        switch (fi34 & 2147483647) {
        case 4294967295:
            v2 = "foo bar"
        }
        switch (v1 % ((128) | 1)) {
        case 42:
            v13 = (v2 + (v13 + ""))
        case 1:
            v23 = append(v23, (-(v1)))
        default:
            var v35 uint = 18446744073709551615
            _ = v35
        }
        var v36 float32 = (3.14159 * -2.25)
        _ = fi34
        _ = v36
    }
    switch 0 {
    case 2147483647:
        var v37 uint64 = 0
        _ = v37
    case 0:
        v1 = v1
    }
    var v38 int32 = 2147483647
    for fi39 := 0; fi39 < 19; fi39++ {
        _ = len(v23)
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        var v40 string = (v13 + ((v5 + "hello") + ""))
        _ = fi39
        _ = v40
    }
    if (0.1 >= -2.25) {
        v5 = "a"
        v5 = v2
        var v41 float64 = ((v4 * float64((3.14159 + 0.0))) / (1.5 * (v4 * v4)))
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        v43 := GenS2{F0: false, F1: uint32(v3), F2: 922}
        for fi44 := 0; fi44 < 8; fi44++ {
            switch v2 {
            case "":
                v45 := GenS3{F0: (-1 != -1), F1: (v24 & 158), F2: (v13 + v2)}
                _ = v45
            case "a":
                fmt.Println(fi42, v24)
            default:
                v46 := GenS5{F0: 65535, F1: (v3 * (v3 - 0)), F2: fi44}
                _ = v46
            }
            _ = fi44
        }
        v23 = append(v23, v1)
        _ = fi42
        _ = v43
    }
    var v47 uint16 = 65535
    for fi48 := 0; fi48 < 10; fi48++ {
        if (v13 < "") {
            if false {
                if false {
                    v49 := GenS4{F0: (v38 ^ 244), F1: -300, F2: 127}
                    _ = v49
                } else {
                    var v50 int = fi48
                    if false {
                        v5 = (v2 + (("foo bar" + "hello") + v5))
                        var v51 uint64 = genFn0(431, (v3 - 255), 747)
                        _ = len(v23)
                        _ = v51
                    } else {
                        v5 = ((v5 + "\n\t") + ("hello" + ((v2 + v2) + (((((v5 + (v13 + v2)) + (v5 + ((("" + v5) + "a") + ("" + "\n\t")))) + ("" + v2)) + "a") + "foo bar"))))
                    }
                    if false {
                        for fi52 := 0; fi52 < 14; fi52++ {
                            var v53 bool = false
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            if true {
                                v54 := GenS6{F0: -128, F1: v1}
                                var v55 int8 = 0
                                var v56 bool = false
                                _ = v54
                                _ = v55
                                _ = v56
                            }
                            var v57 uint64 = 18446744073709551615
                            _ = fi52
                            _ = v53
                            _ = v57
                        }
                        var v58 float32 = -2.25
                        _ = v58
                    } else {
                        v59 := GenS7{F0: -781, F1: 255, F2: false}
                        _ = v59
                    }
                    _ = v50
                }
                for fi60 := 0; fi60 < 7; fi60++ {
                    v61 := GenS2{F0: false, F1: 4294967295, F2: 9223372036854775807}
                    var v62 uint32 = v24
                    _ = fi60
                    _ = v61
                    _ = v62
                }
            }
        } else {
            var v63 uint32 = (v24 - 2147483648)
            _ = v63
        }
        fmt.Println(v47, v5)
        _ = len(v23)
        _ = fi48
    }
    v23 = append(v23, (v1 + -2147483648))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
