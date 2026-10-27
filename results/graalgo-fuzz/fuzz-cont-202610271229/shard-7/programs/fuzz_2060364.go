package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 bool; F2 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint64; F1 int32; F2 uint }
func genFn0() uint {
    v6 := GenS1{F0: 42}
    var v7 int16 = 0
    for fi8 := 0; fi8 < 3; fi8++ {
        if (0 != 65535) {
            switch 127 {
            case 847:
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 int16 = -62
                    v10 -= (v7 & -1)
                    _ = fi9
                    _ = v10
                }
            default:
                switch (^(fi8)) {
                case 42:
                    for fi11 := 0; fi11 < 8; fi11++ {
                        fmt.Println(fi11, v7, fi8)
                        v7 = (v7 / ((v7) | 1))
                        _ = fi11
                    }
                case -2147483648:
                    fmt.Println(fi8, v7)
                case 9223372036854775807:
                    var v12 uint64 = 0
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        if (true && (32767 == v7)) {
                            v7 += v7
                            switch ("hello" + "foo bar") {
                            case "":
                                fmt.Println(v7, fi8, v7)
                            default:
                                var v14 int16 = 1
                                _ = v14
                            }
                            if (!true) {
                                v7 = v7
                            } else {
                                if true {
                                    switch ("a" + "hello") {
                                    case "hello":
                                        v15 := GenS1{F0: 159}
                                        _ = v15
                                    case "\n\t":
                                        if ((2147483647 != -2147483648) && (1 <= 255)) {
                                            var v16 uint8 = 2
                                            fmt.Println(v7, v7)
                                            if ((-2.25 + (0.0 - 3.14159)) >= (float64((1e10 - float32((-2.25 / (3.14159 * (1e10 / 1.5)))))) + -2.25)) {
                                                v16 = v16
                                                v7 += v7
                                                v17 := GenS2{F0: 18446744073709551615}
                                                _ = v17
                                            } else {
                                                v18 := GenS3{F0: (^(v16))}
                                                _ = v18
                                            }
                                            _ = v16
                                        }
                                    }
                                    v19 := []int{fi8, fi8, fi13, (-(fi13))}
                                    if (-49 < 0) {
                                        var v20 float32 = 1.5
                                        var v21 uint32 = 1
                                        _ = v20
                                        _ = v21
                                    } else {
                                        v22 := GenS2{F0: 4294967295}
                                        fmt.Println(fi13)
                                        v7 = (-(v7))
                                        _ = v22
                                    }
                                    _ = v19
                                }
                            }
                        } else {
                            var v23 bool = true
                            v24 := []int{878, 32768, 4294967295}
                            var v25 float32 = 0.1
                            v26 := GenS1{F0: 255}
                            _ = v23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v27 := []int{-808}
                        _ = fi13
                        _ = v27
                    }
                }
            }
            v7 = v7
            v7 += v7
        }
        _ = fi8
    }
    v7 -= v7
    v7 += 1
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int, p1 int, p2 string) int32 {
    p1 += p1
    if (!(42 > 1)) {
        for fi28 := 0; fi28 < 2; fi28++ {
            v29 := GenS1{F0: 32767}
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 float32 = (-2.25 * float32(float64(float32(((-2.25 + ((((1.5 * 1.5) + 0.0) * float64(1.5)) + (3.14159 + float64((0.0 / float32(3.14159)))))) - float64(((1.5 - 1e10) * ((0.1 * float32(1.5)) + ((1e10 / ((3.14159 - (1.5 * (0.1 + (float32(float64(-2.25)) + 1.5)))) * 1e10)) / 1e10)))))))))
                var v32 uint16 = 33
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi28
            _ = v29
        }
        for fi33 := 0; fi33 < 17; fi33++ {
            for fi34 := 0; fi34 < 8; fi34++ {
                var v35 bool = true
                v36 := []int{}
                fmt.Println(p2)
                _ = len(v36)
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = fi33
        }
        fmt.Println(p2, p1, p0)
    } else {
        var v37 int = 935
        if true {
            if false {
                p2 = p2
                for fi38 := 0; fi38 < 4; fi38++ {
                    p2 = "foo bar"
                    _ = fi38
                }
                v39 := GenS3{F0: 118}
                _ = v39
            } else {
                fmt.Println(p0, p0, p2)
            }
            if false {
                var v40 uint64 = 521
                if (2147483647 != 1) {
                    v41 := GenS4{F0: "a", F1: (genFn0() != 42), F2: -76}
                    _ = v41
                }
                _ = v40
            }
            fmt.Println(v37)
        }
        if false {
            p1 += p0
            v37 -= 740
        }
        _ = v37
    }
    _ = p0
    _ = p1
    _ = p2
    return 711
}

func main() {
    var v1 uint32 = 42
    var v2 string = "\n\t"
    var v3 bool = false
    var v4 string = "hello"
    var v5 uint16 = 1
    switch 0 {
    case -316:
        for fi42 := 0; fi42 < 2; fi42++ {
            var v43 float64 = ((1.5 + 3.14159) / -2.25)
            v44 := GenS2{F0: 0}
            _ = fi42
            _ = v43
            _ = v44
        }
    case 4294967295:
        if v3 {
            for fi45 := 0; fi45 < 9; fi45++ {
                v46 := GenS5{F0: v1}
                v2 = "\n\t"
                _ = fi45
                _ = v46
            }
            v4 = "foo bar"
        } else {
            v47 := []int{42, 9223372036854775807, 42, 266}
            _ = v47
        }
    case 32768:
        v2 = "a"
    }
    v5 += v5
    v2 = v4
    if (!(!true)) {
        v48 := []int{}
        _ = len(v48)
        switch -636 {
        case 128:
            fmt.Println(v3, v2)
        default:
            for fi49 := 0; fi49 < 9; fi49++ {
                v48 = append(v48, fi49)
                for fi50 := 0; fi50 < 20; fi50++ {
                    v48 = append(v48, fi50)
                    v51 := GenS6{F0: 100}
                    v2 = (((v2 + v2) + v2) + v2)
                    v52 := GenS6{F0: 127}
                    v53 := GenS7{F0: 26}
                    _ = fi50
                    _ = v51
                    _ = v52
                    _ = v53
                }
                if v3 {
                    v54 := GenS8{F0: 1, F1: 2147483647, F2: genFn0()}
                    if v3 {
                        var v55 int = -2147483648
                        for fi56 := 0; fi56 < 3; fi56++ {
                            _ = len(v48)
                            fmt.Println(v2)
                            _ = fi56
                        }
                        _ = v55
                    } else {
                        for fi57 := 0; fi57 < 13; fi57++ {
                            var v58 uint16 = 42
                            var v59 int16 = -354
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                    }
                    _ = v54
                }
                _ = fi49
            }
        }
        _ = v48
    } else {
        var v60 int64 = 0
        _ = v60
    }
    if ((false || false) && false) {
        v61 := []int{127, 61, 9223372036854775807, -347}
        _ = v61
    }
    if true {
        var v62 uint64 = 18446744073709551615
        _ = v62
    }
    fmt.Println(v1)
    for fi63 := 0; fi63 < 7; fi63++ {
        v64 := GenS5{F0: v1}
        fmt.Println(v2)
        v65 := GenS8{F0: 2, F1: -2147483648, F2: 18446744073709551615}
        _ = fi63
        _ = v64
        _ = v65
    }
    v66 := GenS4{F0: v4, F1: ((v5 >= (v5 ^ 65535)) || (-128 >= 1)), F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
}
