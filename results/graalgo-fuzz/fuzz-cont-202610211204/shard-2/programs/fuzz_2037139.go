package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 int; F1 string; F2 bool }
type GenS3 struct { F0 uint; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int; F2 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16; F1 int16; F2 string }
type GenS9 struct { F0 int32; F1 int8; F2 uint }
type GenS10 struct { F0 uint; F1 int8 }
func genFn0(p0 int16, p1 int8, p2 uint16) uint64 {
    v4 := []int{32768, -296, 218, 4294967295}
    v5 := GenS1{F0: (p0 & (-(p0))), F1: -719, F2: 32767}
    v4 = append(v4, -666)
    _ = len(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 86
}
func genFn1(p0 int8, p1 uint32, p2 string) int {
    v6 := GenS1{F0: 836, F1: 544, F2: 2}
    v7 := GenS1{F0: 32767, F1: 0, F2: 32767}
    p2 = (p2 + "")
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 691
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: -1, F1: 78, F2: 32767}
        v1 += 127
        _ = fi8
        _ = v9
    }
    var v10 uint = 65535
    var v11 uint16 = uint16(v10)
    if false {
        v11 = 32767
    } else {
        if false {
            _ = len(v2)
            switch ("a" + "hello") {
            case "foo bar":
                var v12 int8 = 42
                _ = v12
            default:
                if ("foo bar" <= "a") {
                    if ((("hello" != ("\n\t" + "hello")) && (2147483648 > 2)) && (v10 != v10)) {
                        for fi13 := 0; fi13 < 20; fi13++ {
                            delete(v3, 0)
                            _ = fi13
                        }
                        if false {
                            _ = len(v2)
                            fmt.Println(v1, v1)
                            for fi14 := 0; fi14 < 19; fi14++ {
                                v1 = 42
                                var v15 int8 = 0
                                if (((v1 << ((652) & 63)) >= 128) && false) {
                                    switch "hello" {
                                    case "":
                                        for fi16 := 0; fi16 < 18; fi16++ {
                                            if false {
                                                if true {
                                                    delete(v2, "a")
                                                    if (((v15 + v15) < v15) || (!true)) {
                                                        v17 := []int{int(v1), fi16}
                                                        v3[163] = (-(fi16))
                                                        _ = v17
                                                    }
                                                }
                                                if (v1 <= v1) {
                                                    fmt.Println(fi16)
                                                    var v18 float64 = 1e10
                                                    if true {
                                                        v18 -= ((v18 / ((((3.14159 - float64(0.0)) + ((float64((0.0 * float32(1.5))) - v18) + float64(0.1))) / (v18 + v18)) + ((float64(1.5) / (3.14159 / 3.14159)) / -2.25))) * 1e10)
                                                        v10 += 18446744073709551615
                                                        v19 := []int{}
                                                        v20 := GenS2{F0: fi16, F1: "\n\t", F2: true}
                                                        if true {
                                                            v21 := GenS1{F0: 32767, F1: 2147483647, F2: (v11 & 32767)}
                                                            if true {
                                                                var v22 uint16 = 2
                                                                v23 := GenS1{F0: 1000, F1: 237, F2: 1}
                                                                _ = v22
                                                                _ = v23
                                                            } else {
                                                                fmt.Println(v1, v10, v18)
                                                            }
                                                            _ = v21
                                                        } else {
                                                            v24 := GenS3{F0: v10, F1: (v1 ^ 255), F2: (-(v15))}
                                                            _ = v24
                                                        }
                                                        _ = v19
                                                        _ = v20
                                                    } else {
                                                        v25 := GenS4{F0: v10}
                                                        v26 := GenS4{F0: v10}
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                    v11 = 32767
                                                    v27 := GenS5{F0: v10}
                                                    _ = v18
                                                    _ = v27
                                                }
                                                var v28 float32 = (-2.25 / float32((0.1 - (1e10 / 3.14159))))
                                                _ = v28
                                            } else {
                                                if true {
                                                    v29 := GenS6{F0: v10, F1: fi16, F2: (!true)}
                                                    _ = v29
                                                }
                                                v30 := v3[348]
                                                _ = v30
                                            }
                                            v15 -= v15
                                            v1 = v1
                                            v31 := GenS7{F0: (fi16 * (fi14 / ((fi16) | 1)))}
                                            _ = fi16
                                            _ = v31
                                        }
                                    case "hello":
                                        v32 := []int{(fi14 & (fi14 + 32767)), (fi14 + (fi14 * genFn1(v15, 2147483648, "hello"))), 9223372036854775807, int(v11)}
                                        _ = v32
                                    default:
                                        for fi33 := 0; fi33 < 17; fi33++ {
                                            v34 := GenS8{F0: v11, F1: 1, F2: ""}
                                            _ = fi33
                                            _ = v34
                                        }
                                    }
                                    v15 = (v15 % ((0) | 1))
                                    if false {
                                        _ = len(v3)
                                        if (false && true) {
                                            var v35 int32 = 354
                                            v36 := GenS1{F0: 32767, F1: v35, F2: v11}
                                            v11 -= v11
                                            _ = v35
                                            _ = v36
                                        } else {
                                            var v37 uint = v10
                                            v38 := GenS9{F0: 42, F1: (v15 - (v15 / ((v15) | 1))), F2: (v10 & v37)}
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                } else {
                                    for fi39 := 0; fi39 < 8; fi39++ {
                                        v40 := GenS10{F0: 1, F1: (^(v15))}
                                        v41 := []int{(fi14 + (fi14 + (-(fi14)))), fi39, fi39, -431}
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                    v42 := GenS2{F0: (fi14 * fi14), F1: "\n\t", F2: true}
                                    if (!(v10 < v10)) {
                                        v1 = (v1 | 127)
                                    }
                                    _ = v42
                                }
                                _ = fi14
                                _ = v15
                            }
                        } else {
                            fmt.Println(v10)
                        }
                    } else {
                        v1 = v1
                    }
                    for fi43 := 0; fi43 < 9; fi43++ {
                        v10 = uint(v11)
                        for fi44 := 0; fi44 < 11; fi44++ {
                            if true {
                                if (true && (2 == 1)) {
                                    switch (fi44 | fi44) {
                                    case -738:
                                        delete(v3, 0)
                                    case 0:
                                        fmt.Println(v11)
                                    }
                                    fmt.Println(fi44, v10)
                                    v11 = v11
                                }
                                if true {
                                    v1 += (v1 + 1)
                                    fmt.Println(v1, v11, fi44)
                                    delete(v2, (("\n\t" + ("\n\t" + "foo bar")) + ((("\n\t" + "") + ("a" + "hello")) + "hello")))
                                    _ = len(v3)
                                }
                                v1 -= v1
                            } else {
                                var v45 uint8 = (v1 >> ((v10) & 63))
                                var v46 uint = v10
                                v47 := GenS5{F0: v46}
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                            var v48 uint = v10
                            var v49 int = fi43
                            _ = fi44
                            _ = v48
                            _ = v49
                        }
                        switch (fi43 >> ((v10) & 63)) {
                        case -9223372036854775808:
                            v10 = (v10 | 631)
                        }
                        _ = fi43
                    }
                    v50 := []int{genFn1(-1, 65535, "foo bar"), -2147483648, -9223372036854775808, 0}
                    _ = v50
                } else {
                    fmt.Println(v10)
                    fmt.Println(v10, v11)
                }
            }
        } else {
            var v51 int = 899
            v52 := []int{561, (v51 + (^(v51))), v51, (v51 - 806)}
            if len(v52) > 0 {
                v52[0] = v51
            }
            _ = v51
            _ = v52
        }
        v11 = 1
    }
    var v53 uint16 = 65535
    switch "hello" {
    case "foo bar":
        var v54 string = ((("hello" + "") + ("a" + "foo bar")) + "foo bar")
        _ = v54
    default:
        var v55 int8 = -123
        _ = v55
    }
    v2[("foo bar" + "foo bar")] = 0
    var v56 uint8 = (v1 % (((v1 ^ (^(v1)))) | 1))
    v56 += v1
    var v57 string = "foo bar"
    v58 := GenS1{F0: 0, F1: 2147483647, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v53
    _ = v56
    _ = v57
    _ = v58
}
