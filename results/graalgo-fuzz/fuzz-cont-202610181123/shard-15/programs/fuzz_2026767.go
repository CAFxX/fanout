package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint64; F1 int16; F2 int }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 int8; F2 int }
type GenS9 struct { F0 string; F1 int32 }
func genFn0(p0 uint64, p1 uint32) int8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        p0 = 4294967295
        _ = fi5
    }
    v6 := []int{364}
    for fi7 := 0; fi7 < 18; fi7++ {
        p0 = uint64(p1)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return 38
}

func main() {
    var v1 uint = 65535
    var v2 uint64 = 18446744073709551615
    var v3 uint64 = 1
    v4 := []uint32{}
    v1 -= v1
    switch (("hello" + "") + "hello") {
    case "\n\t":
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint16 = 2
            var v10 string = "foo bar"
            _ = fi8
            _ = v10
            _ = v9
        }
    case "hello":
        fmt.Println(v1, v2, v2)
    default:
        fmt.Println(v1, v2, v2)
    }
    fmt.Println(v2)
    fmt.Println(v1, v3)
    if (int(v3) >= -666) {
        if len(v4) > 0 {
            v4[0] = 0
        }
        v2 -= 1
        fmt.Println(v3, v2)
    } else {
        if (true && (!(1 < 269))) {
            var v11 int16 = 42
            v12 := GenS1{F0: 1, F1: (("foo bar" + "") + ("" + "a"))}
            _ = v11
            _ = v12
        } else {
            var v13 uint32 = 1
            for fi14 := 0; fi14 < 9; fi14++ {
                v3 = v3
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := GenS2{F0: v1, F1: 28}
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 10; fi17++ {
                    v3 = v2
                    v18 := GenS3{F0: 0, F1: 100}
                    if false {
                        switch "a" {
                        case "a":
                            _ = len(v4)
                        }
                        if (false && false) {
                            v13 += uint32(v3)
                        }
                    } else {
                        if true {
                            var v19 float32 = 3.14159
                            if false {
                                switch "hello" {
                                case "a":
                                    var v20 uint8 = 128
                                    _ = v20
                                case "\n\t":
                                    switch fi14 {
                                    case 42:
                                        fmt.Println(v2)
                                    case -596:
                                        v21 := GenS2{F0: v1, F1: 127}
                                        _ = v21
                                    default:
                                        switch "hello" {
                                        case "foo bar":
                                            v22 := GenS3{F0: (v2 + v3), F1: -32}
                                            _ = v22
                                        default:
                                            v23 := GenS2{F0: 591, F1: 0}
                                            _ = v23
                                        }
                                    }
                                case "foo bar":
                                    v4 = append(v4, 2147483647)
                                    var v24 float64 = 1e10
                                    v25 := GenS3{F0: (v3 * (^(v3))), F1: genFn0(498, 255)}
                                    v24 -= v24
                                    _ = v24
                                    _ = v25
                                }
                                var v26 int = 421
                                _ = v26
                            }
                            _ = v19
                        } else {
                            v27 := GenS4{F0: (^(v2)), F1: 784, F2: fi14}
                            v28 := GenS5{F0: 127, F1: 255}
                            _ = v27
                            _ = v28
                        }
                        v13 = (v13 >> ((631) & 63))
                    }
                    var v29 int16 = -947
                    _ = fi17
                    _ = v18
                    _ = v29
                }
                _ = fi14
            }
            v13 -= (v13 & v13)
            _ = v13
        }
        fmt.Println(v2, v3)
        var v30 int32 = 2147483647
        _ = v30
    }
    v31 := GenS6{F0: false, F1: 2}
    v3 += (v2 | (v3 & (v2 >> (((v1 - v1)) & 63))))
    var v32 uint32 = 140
    v33 := GenS2{F0: (^(v1)), F1: 127}
    v34 := GenS6{F0: false, F1: 65535}
    v35 := GenS4{F0: 0, F1: 1, F2: 9223372036854775807}
    v3 = v3
    if false {
        switch "a" {
        case "foo bar":
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 uint64 = v3
                _ = fi36
                _ = v37
            }
        case "hello":
            var v38 string = ("hello" + "foo bar")
            _ = v38
        case "\n\t":
            var v39 float32 = ((float32((-2.25 / (1e10 * (1.5 + (((float64((((0.0 + -2.25) / 3.14159) - 3.14159)) / 3.14159) - float64(float32(0.1))) * 1e10))))) * float32((0.0 + 1.5))) * 1e10)
            _ = v39
        }
    } else {
        v3 = (v3 - 4294967295)
        v3 = (v3 / (((v3 >> (((v1 >> ((v1) & 63))) & 63))) | 1))
        v40 := GenS1{F0: 32767, F1: (("foo bar" + "foo bar") + "foo bar")}
        _ = v40
    }
    v41 := GenS5{F0: 42, F1: 1}
    v42 := GenS7{F0: 42}
    if (!false) {
        for fi43 := 0; fi43 < 19; fi43++ {
            switch "hello" {
            case "hello":
                if false {
                    var v44 int = (-(fi43))
                    _ = v44
                }
            case "a":
                var v45 float64 = (0.0 * 0.0)
                _ = v45
            }
            for fi46 := 0; fi46 < 18; fi46++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi46
            }
            v1 = (v1 & 1)
            _ = fi43
        }
        for fi47 := 0; fi47 < 17; fi47++ {
            if (((42 >= 32) && (!false)) || (1.5 <= -2.25)) {
                fmt.Println(v1)
            } else {
                v48 := GenS8{F0: int16(fi47), F1: 42, F2: (fi47 * (^(fi47)))}
                v4 = append(v4, v32)
                fmt.Println(v3)
                _ = v48
            }
            if (!false) {
                var v49 uint32 = 2147483648
                if (!false) {
                    v3 = 1
                    for fi50 := 0; fi50 < 1; fi50++ {
                        var v51 bool = true
                        var v52 int = fi47
                        if v51 {
                            v53 := GenS9{F0: "\n\t", F1: 42}
                            v4 = append(v4, v49)
                            var v54 uint32 = v49
                            _ = v53
                            _ = v54
                        } else {
                            fmt.Println(v3, v52)
                            var v55 bool = (-1 != -74)
                            _ = v55
                        }
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                } else {
                    fmt.Println(v32, v49)
                    switch ("foo bar" + "hello") {
                    case "hello":
                        var v56 int32 = -1
                        _ = v56
                    case "a":
                        var v57 uint = (v1 - v1)
                        _ = v57
                    }
                }
                switch "foo bar" {
                case "foo bar":
                    switch 2147483647 {
                    case -997:
                        _ = len(v4)
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            var v58 int = (fi47 ^ fi47)
                            _ = v58
                        }
                    }
                case "\n\t":
                    v2 = 735
                default:
                    var v59 uint8 = 171
                    _ = v59
                }
                _ = v49
            } else {
                switch 4294967295 {
                case 1:
                    v1 = (v1 % ((205) | 1))
                }
            }
            if ((2147483647 < 748) && ("foo bar" <= "hello")) {
                var v60 float64 = float64(1.5)
                for fi61 := 0; fi61 < 6; fi61++ {
                    var v62 int8 = 1
                    v63 := GenS3{F0: v3, F1: 1}
                    if (213 != 146) {
                        if len(v4) > 0 {
                            v4[0] = v32
                        }
                    } else {
                        switch (fi47 / ((42) | 1)) {
                        case 2147483647:
                            v1 -= v1
                        case 127:
                            v1 += (v1 % ((v1) | 1))
                        case -761:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            _ = len(v4)
                            fmt.Println(v32, v32)
                            var v64 int64 = 942
                            _ = v64
                        }
                    }
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                _ = v60
            }
            _ = fi47
        }
        var v65 float64 = 0.0
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v42
}
