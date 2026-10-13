package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int16; F1 int32; F2 int64 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 string; F1 uint32; F2 string }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 540
    v3 := make(map[uint32]bool)
    v2 = 236
    v4 := []int{-242}
    v5 := GenS1{F0: 9223372036854775807, F1: -98}
    v6 := GenS2{F0: 100, F1: 1}
    fmt.Println(v2)
    v2 = v2
    v7 := v3[2]
    if ((4294967295 > 4294967295) || false) {
        v8 := GenS3{F0: 32767, F1: -1, F2: 0}
        switch -289 {
        case -2147483648:
            switch -9223372036854775808 {
            case 366:
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 int8 = -128
                    v11 := GenS4{F0: 82, F1: -870}
                    if (2 >= 350) {
                        if (!v7) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v12 := GenS5{F0: 1}
                            _ = v12
                        }
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        default:
            var v13 int8 = 127
            _ = v13
        }
        _ = v8
    }
    if ((1.5 / 0.1) <= (((0.1 * (-2.25 - 1.5)) - 3.14159) - 0.0)) {
        if (v7 && v7) {
            v4 = append(v4, 582)
            delete(v1, ("hello" + "hello"))
            _ = len(v1)
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                if (true && v7) {
                    fmt.Println(fi14, v2, v2)
                    v2 = (^(v2))
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v16 := v3[981]
                        switch "\n\t" {
                        case "":
                            v2 -= (v2 % (((v2 << ((732) & 63))) | 1))
                        }
                        _ = fi15
                        _ = v16
                    }
                } else {
                    v2 = (v2 >> ((2) & 63))
                    var v17 bool = false
                    v18 := GenS6{F0: v7, F1: ""}
                    _ = v17
                    _ = v18
                }
                _ = fi14
            }
            switch -1 {
            case 32767:
                var v19 float32 = (((3.14159 - 0.0) + float32(0.1)) / 1.0)
                _ = v19
            case 315:
                fmt.Println(v2, v2)
            }
            v7 = v7
        }
    } else {
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 bool = true
            v22 := GenS4{F0: 1, F1: -1}
            if v21 {
                if ((^(v2)) < 65535) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch -791 {
                    case 42:
                        if false {
                            v3[448] = ((0.1 >= (1e10 / -2.25)) || v7)
                        }
                    case 32768:
                        for fi23 := 0; fi23 < 12; fi23++ {
                            v24 := GenS7{F0: v2, F1: 1, F2: v2}
                            v25 := GenS8{F0: 2, F1: (("\n\t" + "a") + ("" + ""))}
                            switch fi20 {
                            case 9223372036854775807:
                                v26 := GenS9{F0: 42}
                                _ = v26
                            default:
                                fmt.Println(fi20, v7)
                            }
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    default:
                        for fi27 := 0; fi27 < 5; fi27++ {
                            var v28 int = (fi20 << ((1) & 63))
                            v21 = v21
                            v29 := GenS10{F0: "\n\t", F1: 0, F2: "hello"}
                            v30 := GenS11{F0: v21}
                            if (0 > 583) {
                                if v7 {
                                    v21 = (!v21)
                                } else {
                                    v28 += (-(fi27))
                                }
                                v31 := GenS9{F0: 898}
                                switch (("hello" + ("foo bar" + "\n\t")) + "hello") {
                                case "a":
                                    switch "hello" {
                                    case "\n\t":
                                        var v32 uint32 = 281
                                        _ = v32
                                    case "foo bar":
                                        v33 := GenS1{F0: 0, F1: 0}
                                        _ = v33
                                    default:
                                        fmt.Println(v21, v28)
                                    }
                                }
                                _ = v31
                            } else {
                                if ((^(v2)) == v2) {
                                    v34 := GenS12{F0: 0}
                                    var v35 int16 = 1000
                                    _ = v34
                                    _ = v35
                                } else {
                                    var v36 int8 = 100
                                    switch "a" {
                                    case "hello":
                                        var v37 int8 = v36
                                        _ = v37
                                    case "a":
                                        var v38 float64 = 3.14159
                                        _ = v38
                                    case "\n\t":
                                        v28 = (v28 | v28)
                                    default:
                                        v39 := GenS12{F0: 799}
                                        _ = v39
                                    }
                                    fmt.Println(v21, v7, fi20)
                                    v40 := GenS4{F0: v2, F1: 42}
                                    _ = v36
                                    _ = v40
                                }
                            }
                            var v41 int8 = 100
                            v3[65535] = v7
                            _ = fi27
                            _ = v28
                            _ = v29
                            _ = v30
                            _ = v41
                        }
                    }
                }
                fmt.Println(v7, v7, v2)
            }
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v7 = v7
    for fi42 := 0; fi42 < 4; fi42++ {
        if (!(!(("" + "") < "foo bar"))) {
            v2 += 1
            v43 := v1[(("foo bar" + "foo bar") + ("foo bar" + "a"))]
            delete(v1, ("" + ""))
            _ = v43
        } else {
            switch "" {
            case "foo bar":
                delete(v3, 1)
            case "\n\t":
                var v44 string = "hello"
                _ = v44
            }
        }
        v7 = v7
        _ = fi42
    }
    fmt.Println(v7, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
