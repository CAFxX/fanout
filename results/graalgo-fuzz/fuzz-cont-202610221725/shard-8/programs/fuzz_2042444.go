package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int32; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := []uint32{}
    var v2 string = "a"
    v3 := make(map[int]uint32)
    var v4 float64 = float64(1.5)
    var v5 int8 = -56
    v6 := GenS1{F0: 1}
    var v7 string = v2
    v1 = append(v1, 865)
    var v8 int8 = -1
    var v9 bool = (false || true)
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            fmt.Println(v9, v7)
            v8 = (-(v5))
            _ = len(v1)
            v8 = (^(v8))
            v12 := GenS1{F0: 32767}
            if true {
                var v13 int = (fi10 - fi10)
                _ = v13
            }
            v14 := GenS2{F0: 42}
            switch (v2 + (("hello" + "") + (v2 + (((v2 + v2) + "a") + ((v2 + v2) + (("a" + v7) + v2)))))) {
            case "\n\t":
                if ((3.14159 - (v4 + v4)) > float64(float32((1e10 + ((0.0 * (v4 * ((((0.1 + 1e10) + (1.5 * float64(1.5))) * (v4 + v4)) + v4))) * -2.25))))) {
                    v15 := GenS3{F0: 65535, F1: 4294967295}
                    v9 = ((false && (v5 > v5)) || (2 != 1))
                    v8 += v8
                    _ = v15
                } else {
                    var v16 uint = 311
                    var v17 uint32 = 5
                    _ = v16
                    _ = v17
                }
            case "":
                if (!v9) {
                    fmt.Println(v4, v5)
                    var v18 uint8 = 42
                    var v19 uint = 703
                    v20 := GenS3{F0: 1, F1: 42}
                    _ = v18
                    _ = v19
                    _ = v20
                }
            default:
                var v21 uint32 = 0
                _ = v21
            }
            var v22 uint64 = 599
            v9 = v9
            _ = fi11
            _ = v12
            _ = v14
            _ = v22
        }
        _ = fi10
    }
    switch 1 {
    case 32767:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -2147483648:
        v3[25] = 2
    }
    fmt.Println(v9, v7)
    if v9 {
        switch 9223372036854775807 {
        case -140:
            switch ("" + v7) {
            case "a":
                _ = len(v3)
            case "":
                if false {
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v24 := v3[fi23]
                        v5 -= (v5 ^ 0)
                        _ = fi23
                        _ = v24
                    }
                    v25 := GenS3{F0: 334, F1: 471}
                    _ = v25
                } else {
                    v5 = (-(v8))
                    if (v9 || true) {
                        switch "\n\t" {
                        case "\n\t":
                            var v26 int = -384
                            _ = v26
                        case "a":
                            fmt.Println(v9)
                        }
                    }
                    v2 = v7
                }
            case "\n\t":
                if (255 == 61) {
                    switch 127 {
                    case 4294967295:
                        switch 128 {
                        case -156:
                            var v27 int64 = 643
                            _ = v27
                        case 32767:
                            v7 = (((v7 + "hello") + ("\n\t" + (v2 + ("\n\t" + (("foo bar" + (v7 + ((v7 + v7) + "foo bar"))) + ("hello" + v2)))))) + ("\n\t" + ("hello" + v2)))
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case 1:
                        var v28 uint16 = 255
                        _ = v28
                    }
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                }
            }
        }
        if (v9 || v9) {
            for fi29 := 0; fi29 < 12; fi29++ {
                var v30 uint16 = 469
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if v9 {
                    v9 = v9
                }
                _ = fi29
                _ = v30
            }
            if true {
                v4 = (v4 * (-2.25 + 0.0))
            } else {
                if (69 < 745) {
                    v8 = v8
                    if ((-2.25 - (((v4 - v4) / 1.0) / v4)) == 1e10) {
                        v31 := v3[-2147483648]
                        fmt.Println(v5, v31, v7)
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 18; fi32++ {
                        for fi33 := 0; fi33 < 17; fi33++ {
                            if v9 {
                                if (1000 > 1000) {
                                    v8 -= (v8 % ((-128) | 1))
                                } else {
                                    v7 = (("foo bar" + (v7 + ("hello" + v7))) + v2)
                                }
                                fmt.Println(v4, fi33, v8)
                                if (331 == 42) {
                                    switch (((("" + v7) + ("hello" + ((v2 + ("hello" + v7)) + "hello"))) + (v7 + ("foo bar" + ("\n\t" + (v7 + ("hello" + "a")))))) + (v2 + "hello")) {
                                    case "":
                                        var v34 float32 = (3.14159 / 0.1)
                                        _ = v34
                                    case "hello":
                                        v5 -= (v8 << ((65535) & 63))
                                    case "\n\t":
                                        v5 = -53
                                    }
                                    v2 = v7
                                    var v35 bool = (42 < 968)
                                    _ = v35
                                }
                            }
                            var v36 int8 = (v8 * 42)
                            v3[fi32] = 564
                            fmt.Println(v2)
                            _ = fi33
                            _ = v36
                        }
                        _ = fi32
                    }
                } else {
                    if (!(1.5 == v4)) {
                        var v37 int64 = -685
                        _ = v37
                    } else {
                        v4 -= float64(0.1)
                        v38 := GenS1{F0: 23}
                        _ = v38
                    }
                    if len(v1) > 0 {
                        v1[0] = 883
                    }
                }
                var v39 int32 = 636
                for fi40 := 0; fi40 < 12; fi40++ {
                    v41 := GenS4{F0: false, F1: v9}
                    var v42 uint8 = 0
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = v39
            }
        } else {
            v8 = v5
            if (v9 || v9) {
                var v43 int32 = 1
                var v44 int = -1
                _ = v43
                _ = v44
            }
            v1 = append(v1, 1)
        }
        v45 := GenS5{F0: -2147483648, F1: v5, F2: 2}
        _ = v45
    } else {
        v5 += v5
        if (!true) {
            switch 32768 {
            case 32767:
                switch ("hello" + ("a" + "a")) {
                case "":
                    v46 := GenS6{F0: 2, F1: (^(v8)), F2: 659}
                    _ = v46
                default:
                    switch v7 {
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "\n\t":
                        if false {
                            fmt.Println(v2, v9, v2)
                            var v47 uint8 = 32
                            _ = v47
                        }
                    case "a":
                        for fi48 := 0; fi48 < 7; fi48++ {
                            if (-1 >= -322) {
                                var v49 bool = true
                                delete(v3, (fi48 ^ 9223372036854775807))
                                _ = v49
                            }
                            var v50 int64 = 1
                            var v51 int8 = 0
                            _ = fi48
                            _ = v50
                            _ = v51
                        }
                    default:
                        fmt.Println(v9, v5, v4)
                    }
                }
            case -9223372036854775808:
                for fi52 := 0; fi52 < 7; fi52++ {
                    v1 = append(v1, 2)
                    switch v7 {
                    case "":
                        var v53 int32 = 0
                        _ = v53
                    case "hello":
                        v54 := GenS7{F0: 2}
                        _ = v54
                    }
                    v4 = 1.5
                    _ = fi52
                }
            case 4294967295:
                v8 = (v5 | 42)
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v55 int = 9223372036854775807
            _ = v55
        }
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
