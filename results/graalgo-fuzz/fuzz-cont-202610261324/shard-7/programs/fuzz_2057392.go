package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int8; F1 uint64; F2 int16 }
type GenS7 struct { F0 uint16; F1 int; F2 uint32 }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 uint16; F1 int32; F2 string }
func genFn0(p0 int64, p1 uint) int64 {
    var v7 uint64 = 1
    if (0 > -6) {
        v8 := GenS1{F0: 9223372036854775807, F1: uint16(p1)}
        var v9 float64 = (0.0 / 1.0)
        fmt.Println(p0, p1)
        _ = v8
        _ = v9
    } else {
        switch 127 {
        case 0:
            var v10 uint8 = 255
            _ = v10
        case -9223372036854775808:
            v11 := []int{-9223372036854775808, int(p1), 454}
            _ = v11
        case -274:
            if (32767 != -32768) {
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 uint = (p1 >> (((p1 << ((255) & 63))) & 63))
                    v14 := []int{fi12}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                p1 += (p1 / ((862) | 1))
            } else {
                for fi15 := 0; fi15 < 2; fi15++ {
                    fmt.Println(v7, fi15)
                    _ = fi15
                }
                if (("foo bar" + "a") <= (("" + (("a" + ("foo bar" + "\n\t")) + "hello")) + (("foo bar" + ("" + ((((("" + "") + "a") + "hello") + "\n\t") + "\n\t"))) + "hello"))) {
                    var v16 uint64 = (v7 << ((p1) & 63))
                    if (!(!(!(-2.25 >= 1e10)))) {
                        v17 := GenS1{F0: -917, F1: 167}
                        if false {
                            if (0 < (v7 % ((1) | 1))) {
                                v18 := GenS2{F0: -2147483648, F1: 0, F2: (p0 + (p0 >> ((p1) & 63)))}
                                _ = v18
                            }
                            if ((false && true) || (1 == -1)) {
                                p0 -= p0
                            } else {
                                for fi19 := 0; fi19 < 11; fi19++ {
                                    v20 := GenS3{F0: (("\n\t" + "") + "a"), F1: 2}
                                    var v21 float64 = 0.0
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                                var v22 uint8 = 127
                                _ = v22
                            }
                            fmt.Println(v7)
                        }
                        v23 := GenS3{F0: ("foo bar" + ("" + "hello")), F1: 32767}
                        _ = v17
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v16 = v16
                            _ = fi24
                        }
                        if ((p1 - p1) < 2) {
                            switch "hello" {
                            case "a":
                                v25 := GenS2{F0: 2147483647, F1: 42, F2: 0}
                                _ = v25
                            case "hello":
                                var v26 bool = (!true)
                                _ = v26
                            }
                            p1 += (p1 + 65535)
                            p1 -= (p1 | 1)
                        } else {
                            fmt.Println(v7, v7)
                            v27 := GenS4{F0: "foo bar", F1: 0}
                            _ = v27
                        }
                    }
                    _ = v16
                }
            }
        }
        if true {
            for fi28 := 0; fi28 < 3; fi28++ {
                v29 := []int{}
                _ = fi28
                _ = v29
            }
        }
        for fi30 := 0; fi30 < 7; fi30++ {
            for fi31 := 0; fi31 < 13; fi31++ {
                if true {
                    v7 = 60
                    var v32 float32 = 3.14159
                    _ = v32
                }
                _ = fi31
            }
            for fi33 := 0; fi33 < 1; fi33++ {
                var v34 int8 = -128
                if true {
                    v35 := []int{fi30}
                    if len(v35) > 0 {
                        v35[0] = (fi30 - (fi30 / ((fi30) | 1)))
                    }
                    _ = v35
                } else {
                    v36 := GenS5{F0: 0, F1: 255}
                    _ = v36
                }
                _ = fi33
                _ = v34
            }
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 uint = 18446744073709551615
                _ = fi37
                _ = v38
            }
            _ = fi30
        }
    }
    p1 = 2
    p1 -= 4294967295
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 int8, p2 int) uint {
    var v39 int32 = -2147483648
    for fi40 := 0; fi40 < 18; fi40++ {
        if true {
            if ((false && (0 < 0)) || (1 >= 648)) {
                fmt.Println(v39)
            }
        } else {
            v39 = 1
        }
        v41 := []int{750, (fi40 + (^(fi40))), (fi40 / (((-(fi40))) | 1)), (-(fi40))}
        fmt.Println(fi40, p0)
        _ = fi40
        _ = v41
    }
    v42 := GenS6{F0: p1, F1: 2, F2: 216}
    v43 := []int{(p2 * p2), p2, (p2 - p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    _ = v42
    _ = v43
    return (p0 | p0)
}

func main() {
    var v1 int16 = 1
    var v2 string = "\n\t"
    v3 := make(map[string]bool)
    var v4 int16 = (v1 & (v1 >> ((1) & 63)))
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]uint32)
    fmt.Println(v4, v1, v1)
    var v44 bool = (100 >= 100)
    switch v2 {
    case "hello":
        if (-31 == -33) {
            fmt.Println(v44, v44, v4)
            switch 4294967295 {
            case 4294967295:
                v4 -= (v4 & v4)
            case 128:
                var v45 float64 = float64((0.1 + -2.25))
                _ = v45
            case 9223372036854775807:
                fmt.Println(v44, v4, v44)
            }
        } else {
            var v46 int8 = 0
            _ = v46
        }
    case "a":
        for fi47 := 0; fi47 < 16; fi47++ {
            var v48 int16 = v4
            v49 := GenS7{F0: 42, F1: 2147483647, F2: 65535}
            _ = fi47
            _ = v48
            _ = v49
        }
    }
    var v50 uint64 = 0
    v51 := GenS8{F0: int32(v4), F1: v44}
    v50 -= v50
    v52 := []int{-2147483648, 42, -9223372036854775808}
    v53 := GenS9{F0: 931, F1: 4294967295}
    v54 := GenS7{F0: 2, F1: 0, F2: 4294967295}
    v44 = ("foo bar" != ("foo bar" + v2))
    var v55 int64 = -901
    if ((1.5 + (0.1 / 1.0)) != 3.14159) {
        v56 := GenS10{F0: 255, F1: 0, F2: v2}
        if true {
            if len(v52) >= 2 {
                v52 = v52[1:len(v52)]
            }
        }
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
