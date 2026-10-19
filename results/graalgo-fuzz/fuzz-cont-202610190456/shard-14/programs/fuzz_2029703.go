package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 int16; F2 int16 }
type GenS6 struct { F0 int32; F1 uint16; F2 string }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int8, p2 int32) uint16 {
    var v6 float64 = 3.14159
    fmt.Println(v6, v6)
    switch -1 {
    case 127:
        fmt.Println(p0, p2, p2)
    case -1:
        v7 := GenS1{F0: 2147483647, F1: 36, F2: -1}
        _ = v7
    case 32768:
        if true {
            var v8 int = 2147483647
            for fi9 := 0; fi9 < 19; fi9++ {
                for fi10 := 0; fi10 < 16; fi10++ {
                    if (p1 == int8(p2)) {
                        var v11 int32 = (p2 / ((int32(p0)) | 1))
                        v12 := GenS2{F0: (^(p1))}
                        _ = v11
                        _ = v12
                    }
                    var v13 uint16 = 42
                    v6 = (((v6 * v6) - (3.14159 * v6)) * v6)
                    v14 := []int{-465}
                    fmt.Println(p0, v8)
                    _ = fi10
                    _ = v13
                    _ = v14
                }
                _ = fi9
            }
            var v15 int = (v8 ^ v8)
            _ = v15
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 int8 = -1
    var v3 uint16 = 255
    var v4 bool = (0 > 0)
    var v5 int32 = 145
    v2 = v2
    v16 := GenS2{F0: v2}
    if true {
        v4 = (!v4)
        switch "a" {
        case "":
            for fi17 := 0; fi17 < 4; fi17++ {
                v1 = append(v1, 1)
                switch (("foo bar" + ("\n\t" + "foo bar")) + (("foo bar" + "foo bar") + (("" + ("\n\t" + ("\n\t" + "hello"))) + ("a" + ("a" + "hello"))))) {
                case "hello":
                    switch ("foo bar" + ("\n\t" + "foo bar")) {
                    case "foo bar":
                        v4 = (119 == 2)
                    case "hello":
                        v18 := GenS3{F0: fi17, F1: v4}
                        _ = v18
                    case "":
                        switch fi17 {
                        case 32767:
                            var v19 int32 = v5
                            _ = v19
                        case -9223372036854775808:
                            var v20 bool = v4
                            _ = v20
                        }
                    }
                case "foo bar":
                    var v21 uint8 = 2
                    _ = v21
                default:
                    if (false && (!true)) {
                        v22 := GenS4{F0: 255}
                        var v23 int32 = -2147483648
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v22
                        _ = v23
                    } else {
                        v24 := GenS1{F0: 42, F1: 170, F2: (fi17 % ((fi17) | 1))}
                        if (v4 || true) {
                            var v25 bool = true
                            if v25 {
                                if true {
                                    if v25 {
                                        for fi26 := 0; fi26 < 6; fi26++ {
                                            _ = len(v1)
                                            v27 := GenS4{F0: 655}
                                            var v28 int64 = -140
                                            _ = fi26
                                            _ = v27
                                            _ = v28
                                        }
                                        v29 := GenS3{F0: 42, F1: ((0.1 / 1.0) < (float64(1.5) + (1e10 - (-2.25 / 1.5))))}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v29
                                    }
                                    fmt.Println(v25, fi17)
                                    var v30 uint64 = 4294967295
                                    _ = v30
                                }
                            } else {
                                fmt.Println(v4, v3, fi17)
                                if false {
                                    for fi31 := 0; fi31 < 9; fi31++ {
                                        v32 := GenS5{F0: (v5 << ((18446744073709551615) & 63)), F1: 0, F2: -894}
                                        _ = fi31
                                        _ = v32
                                    }
                                    for fi33 := 0; fi33 < 18; fi33++ {
                                        if v25 {
                                            if v25 {
                                                var v34 float64 = (1.5 / 1.0)
                                                v1 = append(v1, 0)
                                                _ = v34
                                            } else {
                                                var v35 int32 = (v5 >> ((255) & 63))
                                                v5 = v5
                                                _ = v35
                                            }
                                            switch fi17 {
                                            case 4294967295:
                                                switch "\n\t" {
                                                case "hello":
                                                    fmt.Println(v2, v3)
                                                case "a":
                                                    var v36 bool = v25
                                                    _ = v36
                                                default:
                                                    v4 = true
                                                }
                                            }
                                            v2 = (v2 % (((v2 * (-(v2)))) | 1))
                                        } else {
                                            var v37 uint16 = (v3 << ((255) & 63))
                                            v1 = append(v1, 42)
                                            if (0 == 2) {
                                                var v38 string = "a"
                                                _ = v38
                                            } else {
                                                if v25 {
                                                    var v39 int32 = (-(v5))
                                                    _ = v39
                                                }
                                            }
                                            var v40 uint64 = 758
                                            v25 = v25
                                            _ = v37
                                            _ = v40
                                        }
                                        _ = fi33
                                    }
                                    v25 = (!v4)
                                }
                            }
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                            _ = v25
                        } else {
                            for fi41 := 0; fi41 < 12; fi41++ {
                                fmt.Println(v4, fi41, fi41)
                                fmt.Println(v5)
                                v42 := GenS6{F0: v5, F1: 606, F2: ("foo bar" + ("" + ("" + "\n\t")))}
                                switch fi41 {
                                case 32768:
                                    v2 += (v2 % ((127) | 1))
                                case 128:
                                    v43 := GenS7{F0: 648, F1: (^(v3))}
                                    _ = v43
                                }
                                _ = fi41
                                _ = v42
                            }
                            var v44 int8 = v2
                            _ = v44
                        }
                        var v45 float32 = 1.5
                        _ = v24
                        _ = v45
                    }
                }
                _ = fi17
            }
        }
        v1 = append(v1, 2147483647)
    } else {
        var v46 float32 = 1.5
        var v47 int8 = (v2 << ((2) & 63))
        v5 = 2147483647
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 1; fi48++ {
        v49 := GenS6{F0: v5, F1: 42, F2: "hello"}
        fmt.Println(v5, fi48)
        _ = fi48
        _ = v49
    }
    v50 := GenS8{F0: genFn0(238, v2, int32(v3))}
    v2 -= v2
    var v51 uint = 4294967295
    v52 := GenS4{F0: v51}
    var v53 int32 = v5
    fmt.Println(v2, v5, v5)
    switch (("\n\t" + (("\n\t" + "a") + ((("hello" + "\n\t") + "") + "a"))) + "hello") {
    case "a":
        switch -1 {
        case 32768:
            v1 = append(v1, 0)
        case -1:
            v4 = (!(1 == 100))
        case 1:
            v2 -= (v2 - v2)
        }
    default:
        fmt.Println(v4, v4, v5)
    }
    var v54 uint = (^(v51))
    if true {
        _ = len(v1)
        v55 := GenS7{F0: 65535, F1: genFn0(194, v2, (v5 / ((718) | 1)))}
        v54 = (v51 - v54)
        _ = v55
    } else {
        v3 += 32767
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    fmt.Println(v53, v54, v54)
    var v56 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v56
}
