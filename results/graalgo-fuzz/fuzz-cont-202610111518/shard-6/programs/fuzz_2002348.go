package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int32; F1 int; F2 bool }
type GenS8 struct { F0 int16; F1 int }
func genFn0() uint8 {
    v4 := []int{-189, 127}
    var v5 uint32 = 171
    if ("\n\t" <= ("" + "foo bar")) {
        v5 = v5
        _ = len(v4)
    }
    v5 = 42
    _ = v4
    _ = v5
    return 1
}
func genFn1() uint64 {
    switch 526 {
    case 127:
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 bool = (false && true)
            var v8 int32 = 461
            _ = fi6
            _ = v7
            _ = v8
        }
    case 1:
        switch "" {
        case "a":
            v9 := GenS1{F0: 32767}
            _ = v9
        case "foo bar":
            var v10 bool = false
            _ = v10
        default:
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int8 = -99
                v13 := GenS2{F0: (fi11 >= (fi11 / (((fi11 - fi11)) | 1))), F1: 42}
                switch fi11 {
                case -9223372036854775808:
                    switch "hello" {
                    case "hello":
                        v14 := GenS3{F0: 0, F1: genFn0()}
                        _ = v14
                    case "":
                        fmt.Println(fi11)
                    }
                case 32767:
                    var v15 int16 = 150
                    _ = v15
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    case 590:
        var v16 int = 0
        _ = v16
    }
    v17 := []int{-2147483648, 4294967295, 127}
    var v18 int16 = 731
    if (382 > 480) {
        switch 2147483647 {
        case 881:
            v17 = append(v17, 127)
        default:
            v19 := GenS1{F0: -945}
            _ = v19
        }
    } else {
        fmt.Println(v18)
        for fi20 := 0; fi20 < 15; fi20++ {
            if true {
                switch "\n\t" {
                case "hello":
                    for fi21 := 0; fi21 < 10; fi21++ {
                        if (!false) {
                            v18 += (v18 | v18)
                            switch "\n\t" {
                            case "\n\t":
                                v18 -= v18
                            default:
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            }
                            if false {
                                v18 += 42
                                switch "\n\t" {
                                case "":
                                    var v22 bool = false
                                    _ = v22
                                default:
                                    fmt.Println(fi20, v18, v18)
                                }
                                var v23 uint64 = 575
                                v17 = append(v17, (fi21 * (^(fi21))))
                                _ = v23
                            }
                            v24 := GenS3{F0: 32767, F1: genFn0()}
                            _ = v24
                        } else {
                            var v25 float32 = 1e10
                            v17 = append(v17, (fi21 ^ fi20))
                            _ = v25
                        }
                        _ = fi21
                    }
                }
                var v26 uint = 255
                _ = v26
            } else {
                fmt.Println(v18, v18)
                v18 -= v18
                v18 = (v18 % (((v18 % (((v18 >> ((65535) & 63))) | 1))) | 1))
            }
            for fi27 := 0; fi27 < 12; fi27++ {
                switch "foo bar" {
                case "foo bar":
                    var v28 bool = (1.5 >= -2.25)
                    _ = v28
                }
                _ = fi27
            }
            var v29 uint8 = 22
            _ = fi20
            _ = v29
        }
        var v30 int8 = -1
        _ = v30
    }
    _ = v17
    _ = v18
    return 42
}

func main() {
    v1 := []int{375, -569}
    v2 := []uint32{2, 395, 1000}
    v3 := make(map[string]uint32)
    var v31 int16 = 32767
    if (1 > 0) {
        v31 = v31
        v31 = v31
        v31 = (v31 + 1)
    } else {
        v3[(("foo bar" + "") + "\n\t")] = 50
        _ = len(v3)
        var v32 float64 = (0.0 - 1e10)
        _ = v32
    }
    v33 := GenS4{F0: 18446744073709551615}
    fmt.Println(v31)
    var v34 int = 9223372036854775807
    for fi35 := 0; fi35 < 14; fi35++ {
        for fi36 := 0; fi36 < 10; fi36++ {
            v3["a"] = 1
            var v37 int64 = 42
            v34 = fi36
            delete(v3, ("" + "foo bar"))
            var v38 float64 = 0.0
            v39 := GenS3{F0: 0, F1: 128}
            fmt.Println(fi35)
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        v31 += 32767
        _ = fi35
    }
    var v40 uint8 = 0
    switch "a" {
    case "foo bar":
        v41 := GenS5{F0: -2147483648, F1: 2, F2: (v40 * (v40 | (v40 & 38)))}
        _ = v41
    }
    var v42 bool = (true || false)
    var v43 int = (v34 & v34)
    for fi44 := 0; fi44 < 18; fi44++ {
        for fi45 := 0; fi45 < 17; fi45++ {
            var v46 int8 = 1
            var v47 uint8 = 0
            if (!v42) {
                v40 = 42
                if false {
                    v43 = 42
                    if v42 {
                        _ = len(v1)
                        if (9223372036854775807 == 42) {
                            if (255 < 1) {
                                fmt.Println(fi44, v34, v43)
                                v34 -= (fi45 % (((fi45 / ((fi44) | 1))) | 1))
                            }
                            var v48 float64 = 0.0
                            var v49 uint8 = (v40 * v40)
                            _ = v48
                            _ = v49
                        } else {
                            v43 += 42
                            delete(v3, (("foo bar" + (("a" + "") + (("\n\t" + ("foo bar" + "\n\t")) + "hello"))) + ("\n\t" + ((("hello" + "") + "hello") + ("a" + "")))))
                            var v50 int = int(v31)
                            _ = v50
                        }
                        if true {
                            var v51 int = fi45
                            v52 := GenS6{F0: -1, F1: v31}
                            _ = v51
                            _ = v52
                        } else {
                            v31 -= v31
                        }
                    } else {
                        var v53 int32 = 39
                        switch (^(fi44)) {
                        case 1:
                            v40 = (^(v47))
                        case -144:
                            v54 := GenS3{F0: 285, F1: (^(v40))}
                            _ = v54
                        }
                        v46 -= v46
                        _ = len(v3)
                        _ = v53
                    }
                }
            }
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi44
    }
    if v42 {
        v34 = -497
        if v42 {
            var v55 int64 = 42
            v56 := GenS7{F0: 1, F1: (-(v43)), F2: (1.5 > (1e10 * float32((1e10 / float64(1e10)))))}
            var v57 uint16 = 115
            _ = v55
            _ = v56
            _ = v57
        }
        var v58 int64 = 465
        _ = v58
    } else {
        fmt.Println(v34, v34)
        switch "\n\t" {
        case "":
            v42 = v42
        case "a":
            v43 -= v43
        case "\n\t":
            fmt.Println(v42, v31)
        default:
            switch "hello" {
            case "a":
                for fi59 := 0; fi59 < 16; fi59++ {
                    if (1 <= -105) {
                        fmt.Println(v31, v42)
                    }
                    var v60 int8 = -128
                    v31 = 0
                    _ = fi59
                    _ = v60
                }
            default:
                v43 += v34
            }
        }
    }
    v61 := v3["a"]
    fmt.Println(v42, v40, v34)
    switch ("\n\t" + "foo bar") {
    case "":
        v43 = (v34 % (((^(v34))) | 1))
    case "foo bar":
        switch "foo bar" {
        case "a":
            v62 := GenS8{F0: (v31 << ((621) & 63)), F1: -343}
            _ = v62
        case "":
            for fi63 := 0; fi63 < 14; fi63++ {
                fmt.Println(v43, v40)
                _ = fi63
            }
        }
    case "a":
        var v64 bool = (((0.1 / 1.0) != 0.0) && (((float64(3.14159) - (1e10 - 0.1)) - (3.14159 * 0.1)) >= 1.5))
        _ = v64
    default:
        v34 -= v34
    }
    var v65 uint16 = 0
    v66 := GenS1{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v61:", v61)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v40
    _ = v42
    _ = v43
    _ = v61
    _ = v65
    _ = v66
}
