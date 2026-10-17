package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64; F1 int64; F2 bool }
type GenS4 struct { F0 string }
func genFn0() uint {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint8 = 255
        var v9 uint32 = 752
        var v10 int8 = 127
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if false {
        v11 := GenS1{F0: -122}
        var v12 int32 = 2147483647
        _ = v11
        _ = v12
    } else {
        var v13 float32 = 3.14159
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint16 = 42
        switch (fi14 ^ 4294967295) {
        case -947:
            for fi16 := 0; fi16 < 7; fi16++ {
                if (18446744073709551615 >= 839) {
                    v17 := []int{(fi14 * (-(fi14))), fi16}
                    fmt.Println(v15)
                    _ = v17
                } else {
                    v15 = v15
                }
                v15 = (^(v15))
                _ = fi16
            }
        case 416:
            switch 127 {
            case -1:
                var v18 uint8 = 1
                _ = v18
            case -551:
                v19 := []int{(fi14 - (fi14 & -235)), fi14, (fi14 << ((255) & 63))}
                _ = v19
            case 811:
                for fi20 := 0; fi20 < 16; fi20++ {
                    fmt.Println(fi14, fi20, v15)
                    if ((-32768 != 1000) || false) {
                        var v21 float32 = (3.14159 / (-2.25 + 1.5))
                        _ = v21
                    }
                    switch (-(fi14)) {
                    case 2147483647:
                        if false {
                            v22 := GenS2{F0: 986, F1: 42}
                            v15 = (v15 + (v15 + v15))
                            if (true || (48 != 1)) {
                                var v23 float64 = (0.1 * (3.14159 - 0.0))
                                v23 += v23
                                _ = v23
                            }
                            _ = v22
                        }
                    default:
                        if false {
                            v15 -= 1
                            v24 := []int{(fi14 + fi14)}
                            switch fi20 {
                            case 127:
                                var v25 int16 = 1000
                                _ = v25
                            case -1:
                                switch "\n\t" {
                                case "hello":
                                    var v26 uint64 = 42
                                    _ = v26
                                default:
                                    if false {
                                        v15 = (^(v15))
                                        var v27 uint = 571
                                        switch (fi14 ^ fi20) {
                                        case 535:
                                            _ = len(v24)
                                        case 271:
                                            fmt.Println(v15, fi20, v27)
                                        default:
                                            fmt.Println(fi14)
                                        }
                                        var v28 float32 = -2.25
                                        _ = v27
                                        _ = v28
                                    } else {
                                        v15 -= (v15 | (^(v15)))
                                        fmt.Println(fi20, fi14, v15)
                                        var v29 int8 = 115
                                        _ = v29
                                    }
                                    var v30 int = 4294967295
                                    switch "a" {
                                    case "hello":
                                        if len(v24) > 0 {
                                            v24[0] = -650
                                        }
                                    case "\n\t":
                                        var v31 bool = true
                                        _ = v31
                                    default:
                                        v30 = -1
                                    }
                                    var v32 uint16 = (^(v15))
                                    fmt.Println(v30, v32)
                                    _ = v30
                                    _ = v32
                                }
                            case 42:
                                v33 := GenS1{F0: 60}
                                _ = v33
                            default:
                                v15 = v15
                            }
                            if (1e10 != 1e10) {
                                v34 := GenS2{F0: 134, F1: -1}
                                _ = v34
                            }
                            _ = v24
                        }
                    }
                    _ = fi20
                }
            default:
                var v35 int = (fi14 << ((65535) & 63))
                _ = v35
            }
        }
        v36 := []int{}
        _ = fi14
        _ = v15
        _ = v36
    }
    return 37
}

func main() {
    var v1 uint16 = 2
    v2 := []string{"hello"}
    v3 := make(map[string]uint32)
    var v4 int8 = -3
    v5 := make(map[int]string)
    v6 := make(map[int]int)
    var v37 int16 = 1
    var v38 uint16 = v1
    var v39 int64 = -1
    fmt.Println(v37)
    fmt.Println(v4)
    for fi40 := 0; fi40 < 12; fi40++ {
        v41 := GenS1{F0: v4}
        _ = fi40
        _ = v41
    }
    switch "a" {
    case "foo bar":
        _ = len(v2)
    case "":
        for fi42 := 0; fi42 < 13; fi42++ {
            for fi43 := 0; fi43 < 6; fi43++ {
                v44 := GenS1{F0: v4}
                switch 32767 {
                case 128:
                    switch (("hello" + "") + "") {
                    case "foo bar":
                        fmt.Println(v37, v4)
                        var v45 uint16 = 952
                        _ = v45
                    case "":
                        fmt.Println(v39, v1, v39)
                    default:
                        var v46 float32 = (3.14159 / 1e10)
                        _ = v46
                    }
                case -9223372036854775808:
                    v37 = (v37 - (v37 ^ (v37 + 1000)))
                default:
                    var v47 int64 = (^(v39))
                    _ = v47
                }
                if len(v2) > 0 {
                    v2[0] = ("hello" + ("foo bar" + "\n\t"))
                }
                v48 := GenS2{F0: (v39 >> ((0) & 63)), F1: (v39 << ((255) & 63))}
                _ = fi43
                _ = v44
                _ = v48
            }
            if true {
                _ = len(v2)
                fmt.Println(v4)
            } else {
                for fi49 := 0; fi49 < 18; fi49++ {
                    v50 := GenS1{F0: v4}
                    fmt.Println(v38, v39, v1)
                    v51 := GenS3{F0: 1, F1: v39, F2: false}
                    _ = fi49
                    _ = v50
                    _ = v51
                }
            }
            _ = fi42
        }
    case "hello":
        if false {
            fmt.Println(v39, v39)
            fmt.Println(v39)
            if true {
                var v52 int32 = 42
                var v53 float32 = 3.14159
                _ = v52
                _ = v53
            }
        }
    }
    v1 -= (v38 >> ((65535) & 63))
    var v54 int64 = -849
    var v55 uint = 42
    var v56 uint64 = uint64(v55)
    for fi57 := 0; fi57 < 8; fi57++ {
        if (0 != 0) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi57
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v58 float32 = 1e10
    v59 := GenS4{F0: ""}
    if true {
        var v60 string = ("a" + (("" + "\n\t") + "\n\t"))
        _ = v60
    } else {
        v54 = 718
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v58
    _ = v59
    _ = v6
}
