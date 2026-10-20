package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 int32; F2 uint }

func main() {
    var v1 uint32 = 320
    v2 := make(map[string]int)
    var v3 uint8 = 167
    var v4 uint32 = (^(v1))
    var v5 int64 = -1
    var v6 int64 = (-(v5))
    v7 := GenS1{F0: -74, F1: 369, F2: 1}
    var v8 int64 = v5
    fmt.Println(v8, v3)
    if (65535 < 412) {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := []int{fi9, fi9}
            switch "foo bar" {
            case "hello":
                v11 := GenS1{F0: 100, F1: 4294967295, F2: 1}
                _ = v11
            case "":
                switch "" {
                case "hello":
                    v12 := v2["a"]
                    _ = v12
                }
            case "a":
                fmt.Println(v3, v1, fi9)
            default:
                v13 := GenS2{F0: 32767, F1: 42, F2: 164}
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        v6 = v6
        for fi14 := 0; fi14 < 2; fi14++ {
            fmt.Println(v3, v1)
            if (42 <= 0) {
                v4 = 2147483648
                if false {
                    var v15 int32 = 2147483647
                    _ = v15
                }
                for fi16 := 0; fi16 < 19; fi16++ {
                    if (2 != 961) {
                        if (774 > 2) {
                            switch "foo bar" {
                            case "":
                                if ((585 > 2) || (((0.1 / 0.1) <= 0.1) || false)) {
                                    var v17 int8 = 42
                                    _ = v17
                                }
                            case "hello":
                                v1 = (v1 << ((4294967295) & 63))
                            default:
                                v3 -= (v3 / ((v3) | 1))
                            }
                            fmt.Println(v5, v1)
                            fmt.Println(fi16, v5, v1)
                            if (!true) {
                                var v18 uint = 42
                                _ = v18
                            } else {
                                v19 := GenS2{F0: 2, F1: 1, F2: 127}
                                fmt.Println(fi14)
                                var v20 int16 = -841
                                _ = v19
                                _ = v20
                            }
                        } else {
                            fmt.Println(v4)
                        }
                        switch (fi14 % ((int(v3)) | 1)) {
                        case -2147483648:
                            v21 := GenS1{F0: 6, F1: uint64(v3), F2: 230}
                            _ = v21
                        default:
                            if false {
                                if true {
                                    var v22 int8 = int8(v4)
                                    v8 -= v5
                                    v22 = v22
                                    _ = v22
                                }
                                if true {
                                    var v23 int64 = (v8 << ((0) & 63))
                                    if (0.0 <= float64(1.5)) {
                                        v24 := GenS3{F0: 1, F1: (v1 << ((649) & 63)), F2: 57}
                                        v3 -= v3
                                        _ = v24
                                    }
                                    fmt.Println(v1, v8)
                                    _ = v23
                                } else {
                                    fmt.Println(v8, v3)
                                }
                            }
                        }
                        var v25 float64 = 1e10
                        _ = v25
                    } else {
                        switch fi14 {
                        case -2147483648:
                            v26 := GenS4{F0: 2, F1: (v3 + v3)}
                            _ = v26
                        case 42:
                            v27 := GenS2{F0: 948, F1: 42, F2: v3}
                            _ = v27
                        }
                        var v28 int32 = 1
                        _ = v28
                    }
                    _ = fi16
                }
            }
            v1 -= v4
            _ = fi14
        }
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        v30 := GenS5{F0: 0, F1: 42}
        _ = fi29
        _ = v30
    }
    var v31 string = "foo bar"
    v32 := []int{}
    var v33 float64 = -2.25
    var v34 uint32 = v1
    fmt.Println(v34, v34, v4)
    for fi35 := 0; fi35 < 6; fi35++ {
        if false {
            delete(v2, "\n\t")
            switch "a" {
            case "hello":
                var v36 uint64 = 420
                _ = v36
            case "\n\t":
                for fi37 := 0; fi37 < 14; fi37++ {
                    switch v31 {
                    case "a":
                        fmt.Println(v6, v6)
                    case "\n\t":
                        if false {
                            if (v3 != (v3 * v3)) {
                                var v38 string = "hello"
                                _ = v38
                            }
                            if false {
                                var v39 float32 = 0.0
                                _ = v39
                            } else {
                                _ = len(v32)
                                if true {
                                    v40 := GenS6{F0: 63, F1: ((v3 >= v3) && false)}
                                    _ = v40
                                } else {
                                    v41 := v2[""]
                                    _ = v41
                                }
                            }
                        }
                    case "hello":
                        v3 += (v3 & (v3 - 255))
                    }
                    var v42 bool = true
                    switch int(v3) {
                    case 9223372036854775807:
                        fmt.Println(v3, v31, v3)
                    }
                    _ = fi37
                    _ = v42
                }
            }
            var v43 uint8 = (v3 - v3)
            _ = v43
        }
        for fi44 := 0; fi44 < 14; fi44++ {
            if false {
                if true {
                    switch ((v31 + v31) + v31) {
                    case "\n\t":
                        v31 = ((v31 + (v31 + (v31 + (v31 + v31)))) + v31)
                        if (true || ((255 < 18446744073709551615) || (127 >= -1))) {
                            fmt.Println(v33)
                            v45 := GenS3{F0: 0, F1: (v4 | v34), F2: 42}
                            _ = v45
                        }
                        if len(v32) >= 2 {
                            v32 = v32[1:len(v32)]
                        }
                        if (false && true) {
                            v31 = v31
                            v6 -= (-(v8))
                            if (!false) {
                                if (((fi35 - fi35) > fi35) || false) {
                                    fmt.Println(v34)
                                } else {
                                    _ = len(v2)
                                }
                                if len(v32) > 0 {
                                    v32[0] = 9223372036854775807
                                }
                                v46 := GenS4{F0: 1, F1: (v3 << ((0) & 63))}
                                v47 := GenS7{F0: 0}
                                _ = v46
                                _ = v47
                            }
                        }
                    }
                }
                if false {
                    if false {
                        var v48 int8 = 1
                        _ = v48
                    }
                    v49 := GenS5{F0: 2, F1: 42}
                    if ((v8 & (v6 % (((-(v5))) | 1))) != v8) {
                        fmt.Println(v8)
                        v50 := GenS8{F0: fi35}
                        v4 -= 2147483647
                        var v51 uint64 = 1
                        _ = v50
                        _ = v51
                    }
                    _ = v49
                } else {
                    v52 := GenS2{F0: 2, F1: 973, F2: v3}
                    fmt.Println(v4)
                    var v53 float64 = (v33 - 0.1)
                    var v54 int8 = int8(fi44)
                    _ = v52
                    _ = v53
                    _ = v54
                }
                var v55 int64 = 9223372036854775807
                _ = v55
            } else {
                v56 := GenS9{F0: 878, F1: 42, F2: 18446744073709551615}
                var v57 uint64 = 0
                switch v31 {
                case "hello":
                    var v58 int64 = v5
                    _ = v58
                default:
                    switch fi44 {
                    case 42:
                        v32 = append(v32, fi44)
                    case -9223372036854775808:
                        var v59 uint32 = v1
                        _ = v59
                    case 32768:
                        var v60 int64 = v8
                        _ = v60
                    default:
                        v3 -= 2
                    }
                }
                _ = v56
                _ = v57
            }
            switch fi35 {
            case 127:
                switch v31 {
                case "":
                    if false {
                        var v61 int8 = 127
                        _ = v61
                    } else {
                        v33 = (v33 * 3.14159)
                        v3 += v3
                        fmt.Println(v6, v1, v4)
                    }
                case "\n\t":
                    var v62 float32 = -2.25
                    _ = v62
                }
            }
            _ = len(v2)
            switch ("foo bar" + (("hello" + v31) + v31)) {
            case "":
                v34 = (v34 - v4)
            case "hello":
                _ = len(v2)
            case "\n\t":
                v63 := GenS9{F0: 255, F1: int32(v8), F2: 0}
                _ = v63
            }
            fmt.Println(v34)
            _ = fi44
        }
        for fi64 := 0; fi64 < 10; fi64++ {
            fmt.Println(v8)
            if false {
                v31 = (((v31 + ("hello" + "foo bar")) + v31) + "\n\t")
                if true {
                    fmt.Println(v5, v4, fi35)
                    v8 += v6
                }
                fmt.Println(v1, v8, v34)
            }
            _ = fi64
        }
        _ = fi35
    }
    if len(v32) > 0 {
        v32[0] = 2147483647
    }
    var v65 int = 128
    v8 += (v6 * 42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v7
    _ = v8
}
