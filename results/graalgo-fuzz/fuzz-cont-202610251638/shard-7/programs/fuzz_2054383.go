package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 int32; F1 uint; F2 string }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint; F2 uint }
type GenS5 struct { F0 int16; F1 bool; F2 string }
type GenS6 struct { F0 int; F1 int; F2 int32 }
type GenS7 struct { F0 int8; F1 bool; F2 bool }
type GenS8 struct { F0 bool; F1 uint32; F2 int16 }
type GenS9 struct { F0 string; F1 uint64 }
func genFn0(p0 int64) uint32 {
    v5 := []int{764, 0, 0}
    switch "a" {
    case "":
        switch ("hello" + (("hello" + (("" + "hello") + "\n\t")) + "hello")) {
        case "\n\t":
            for fi6 := 0; fi6 < 20; fi6++ {
                switch "\n\t" {
                case "":
                    if false {
                        v7 := GenS1{F0: p0, F1: -1, F2: 0}
                        fmt.Println(fi6)
                        _ = v7
                    } else {
                        for fi8 := 0; fi8 < 5; fi8++ {
                            fmt.Println(fi8)
                            var v9 uint16 = 1
                            _ = fi8
                            _ = v9
                        }
                        for fi10 := 0; fi10 < 7; fi10++ {
                            p0 -= -9223372036854775808
                            v5 = append(v5, (fi6 << ((4294967295) & 63)))
                            _ = fi10
                        }
                    }
                default:
                    var v11 uint16 = 0
                    _ = v11
                }
                for fi12 := 0; fi12 < 15; fi12++ {
                    v13 := GenS2{F0: -783, F1: 2, F2: ("" + "")}
                    _ = fi12
                    _ = v13
                }
                if len(v5) > 0 {
                    v5[0] = fi6
                }
                _ = fi6
            }
        case "":
            var v14 bool = true
            _ = v14
        case "foo bar":
            if (100 > 1) {
                var v15 int32 = -1
                var v16 uint16 = 32767
                v17 := GenS3{F0: 18446744073709551615, F1: 1}
                _ = v15
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 10; fi18++ {
                    switch "\n\t" {
                    case "a":
                        v19 := GenS3{F0: 589, F1: 18446744073709551615}
                        _ = v19
                    case "hello":
                        _ = len(v5)
                    default:
                        for fi20 := 0; fi20 < 14; fi20++ {
                            fmt.Println(fi18)
                            var v21 uint64 = 2
                            if true {
                                switch fi20 {
                                case 766:
                                    var v22 uint8 = 1
                                    _ = v22
                                case 32767:
                                    v21 -= 227
                                }
                                p0 -= 42
                                switch 42 {
                                case 0:
                                    v23 := GenS4{F0: true, F1: 45, F2: 255}
                                    _ = v23
                                case 1:
                                    var v24 int64 = 1
                                    switch ("" + "a") {
                                    case "hello":
                                        var v25 float64 = 1e10
                                        _ = v25
                                    case "foo bar":
                                        if (-636 >= -725) {
                                            v26 := GenS1{F0: (-(p0)), F1: 1, F2: 100}
                                            p0 += (p0 | 0)
                                            _ = v26
                                        }
                                    }
                                    fmt.Println(fi18, fi18, p0)
                                    _ = v24
                                default:
                                    v27 := GenS4{F0: (!true), F1: 255, F2: 2}
                                    _ = v27
                                }
                            } else {
                                fmt.Println(v21, fi20, p0)
                            }
                            v5 = append(v5, fi18)
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = fi18
                }
                if (-52 == -2147483648) {
                    if len(v5) > 0 {
                        v5[0] = -9223372036854775808
                    }
                }
            }
        }
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int16 = 466
        switch (fi28 << ((65535) & 63)) {
        case -916:
            v29 = 32767
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 uint8 = 1
            _ = fi30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    _ = p0
    _ = v5
    return 2
}
func genFn1(p0 int64, p1 int16) uint16 {
    if (true || ((!true) || (true || false))) {
        var v32 uint64 = 2
        v32 -= (v32 | (v32 ^ (v32 + v32)))
        switch "hello" {
        case "foo bar":
            fmt.Println(v32, v32)
        case "a":
            var v33 uint16 = 42
            _ = v33
        case "":
            p1 += (p1 + p1)
        default:
            v34 := []int{9223372036854775807, 4294967295, -1, -280}
            _ = v34
        }
        _ = v32
    }
    v35 := []int{32767, -2147483648, -9223372036854775808, -483}
    p1 += 1
    fmt.Println(p0, p0)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v35
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 32767
    v3 := make(map[string]int)
    var v4 float64 = (-2.25 / 1.0)
    if true {
        var v36 int32 = 1
        _ = v36
    }
    v37 := GenS1{F0: 9223372036854775807, F1: -310, F2: 4}
    for fi38 := 0; fi38 < 14; fi38++ {
        for fi39 := 0; fi39 < 3; fi39++ {
            if (3.14159 == 0.1) {
                v2 = 42
                v40 := []int{}
                var v41 int16 = 32767
                _ = v40
                _ = v41
            } else {
                var v42 int16 = 1
                v43 := GenS5{F0: 1, F1: (!((!(0 > v42)) && (0 > 65535))), F2: "hello"}
                _ = v42
                _ = v43
            }
            var v44 string = "hello"
            if (!true) {
                v2 -= (fi38 + (v2 >> ((0) & 63)))
                if false {
                    v45 := GenS6{F0: 9223372036854775807, F1: fi38, F2: 673}
                    _ = v45
                } else {
                    v3[(v44 + v44)] = (v2 << ((2) & 63))
                    if false {
                        fmt.Println(v44, v44)
                        switch 4294967295 {
                        case 127:
                            switch "" {
                            case "foo bar":
                                if (128 < 127) {
                                    var v46 int64 = 0
                                    v47 := GenS6{F0: (fi39 % (((v2 ^ (v2 % ((-9223372036854775808) | 1)))) | 1)), F1: (fi38 | fi39), F2: -1}
                                    var v48 int16 = -1
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                                switch v2 {
                                case 1:
                                    v49 := GenS7{F0: 1, F1: true, F2: false}
                                    _ = v49
                                }
                            case "\n\t":
                                v1[v44] = 1
                            default:
                                var v50 uint = 4294967295
                                _ = v50
                            }
                        case 823:
                            v2 = fi38
                        }
                        v2 += (fi38 & v2)
                    }
                    v2 -= 0
                }
                v51 := v1[v44]
                var v52 uint32 = (v51 | v51)
                _ = v51
                _ = v52
            }
            v2 -= fi38
            _ = fi39
            _ = v44
        }
        if (-2.25 == 0.0) {
            v2 -= (^(v2))
        }
        _ = fi38
    }
    var v53 uint16 = 255
    for fi54 := 0; fi54 < 1; fi54++ {
        v53 -= 255
        v53 = genFn1(-9223372036854775808, 0)
        _ = fi54
    }
    for fi55 := 0; fi55 < 6; fi55++ {
        v56 := GenS8{F0: true, F1: 42, F2: 654}
        v2 = (fi55 * (fi55 ^ (v2 - (v2 + (v2 & fi55)))))
        fmt.Println(v53, v4, v2)
        _ = fi55
        _ = v56
    }
    var v57 uint = 721
    if true {
        delete(v3, "")
    }
    if false {
        v58 := GenS9{F0: "a", F1: 42}
        v59 := []int{-2147483648}
        _ = v58
        _ = v59
    } else {
        for fi60 := 0; fi60 < 20; fi60++ {
            var v61 float32 = (0.1 / 0.1)
            _ = len(v3)
            var v62 int8 = -89
            _ = fi60
            _ = v61
            _ = v62
        }
    }
    var v63 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v53
    _ = v57
    _ = v63
}
