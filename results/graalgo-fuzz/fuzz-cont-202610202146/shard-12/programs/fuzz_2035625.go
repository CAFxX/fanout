package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8 }
func genFn0(p0 bool, p1 int) int32 {
    v5 := GenS1{F0: 1, F1: 2}
    p1 = p1
    p1 -= 1
    var v6 int8 = -128
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 string = (("a" + ("hello" + "foo bar")) + (("hello" + "foo bar") + "\n\t"))
    var v2 int64 = 42
    var v3 uint32 = 2
    var v4 int32 = 42
    if ((32767 >= 32768) && (((false || false) && (255 >= 94)) && (1 <= -32768))) {
        v7 := []int{42, -2147483648}
        _ = v7
    } else {
        var v8 uint64 = 2
        v2 += (v2 | v2)
        v8 -= v8
        _ = v8
    }
    v9 := GenS2{F0: v3, F1: 42, F2: v3}
    if (((("foo bar" + (v1 + "")) + "\n\t") + "hello") == (v1 + "")) {
        var v10 int16 = 42
        _ = v10
    }
    if false {
        switch 42 {
        case 32767:
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := []int{fi11, fi11, fi11, fi11}
                v4 = v4
                _ = fi11
                _ = v12
            }
        case 32768:
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := []int{}
                if (v3 < v3) {
                    var v15 string = ""
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 int8 = 42
                        fmt.Println(v2, fi16)
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                }
                v2 += (v2 & (v2 >> ((0) & 63)))
                _ = fi13
                _ = v14
            }
        case 845:
            var v18 uint32 = v3
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            if true {
                v4 = 1
                var v20 uint = 1
                if true {
                    switch (fi19 & (fi19 / (((^(fi19))) | 1))) {
                    case -68:
                        var v21 uint = v20
                        _ = v21
                    case -9223372036854775808:
                        var v22 uint16 = 479
                        _ = v22
                    default:
                        var v23 uint32 = v3
                        _ = v23
                    }
                    if (604 >= 42) {
                        if true {
                            for fi24 := 0; fi24 < 15; fi24++ {
                                v25 := GenS3{F0: 893, F1: -1, F2: (v3 << (((v20 / (((v20 >> ((v20) & 63))) | 1))) & 63))}
                                v1 = v1
                                _ = fi24
                                _ = v25
                            }
                        } else {
                            v1 = "foo bar"
                            for fi26 := 0; fi26 < 18; fi26++ {
                                v4 -= (^(v4))
                                if true {
                                    v27 := []int{(^(fi19)), fi26, 286}
                                    _ = v27
                                }
                                var v28 int8 = 0
                                var v29 int16 = 42
                                _ = fi26
                                _ = v28
                                _ = v29
                            }
                        }
                        var v30 int8 = 127
                        v31 := GenS1{F0: 1, F1: 0}
                        _ = v30
                        _ = v31
                    }
                    if ((v4 & -2147483648) == (v4 | genFn0(((v4 <= -2147483648) && (!true)), (^(fi19))))) {
                        if false {
                            var v32 float32 = (3.14159 * 1e10)
                            v33 := GenS3{F0: 480, F1: 100, F2: v3}
                            _ = v32
                            _ = v33
                        }
                    }
                } else {
                    v34 := GenS3{F0: 392, F1: 127, F2: v3}
                    _ = v34
                }
                _ = v20
            }
            if true {
                switch (fi19 / (((fi19 + fi19)) | 1)) {
                case 128:
                    v3 = 2147483647
                }
            }
            _ = fi19
        }
        v2 = -219
    } else {
        v35 := []int{128, -961, -1, -2147483648}
        _ = v35
    }
    var v36 uint64 = 421
    if (!(!(4294967295 == -1))) {
        switch -1 {
        case 42:
            v2 -= v2
        case -2147483648:
            v37 := []int{4294967295}
            _ = v37
        default:
            if true {
                switch v1 {
                case "hello":
                    v4 = int32(v36)
                case "":
                    for fi38 := 0; fi38 < 11; fi38++ {
                        var v39 uint8 = 127
                        _ = fi38
                        _ = v39
                    }
                }
                var v40 int16 = 42
                for fi41 := 0; fi41 < 8; fi41++ {
                    v42 := GenS4{F0: (v4 / ((genFn0((((0.1 + (0.1 - (1e10 + 0.1))) - (float32(1e10) + 3.14159)) <= (1e10 - (((-2.25 + 1.5) * 0.0) - ((1e10 / 0.1) - 0.0)))), -343)) | 1))}
                    v1 = "foo bar"
                    _ = fi41
                    _ = v42
                }
                _ = v40
            }
        }
    }
    v36 -= v36
    for fi43 := 0; fi43 < 2; fi43++ {
        switch ("" + ((v1 + (v1 + (("" + v1) + ""))) + v1)) {
        case "a":
            for fi44 := 0; fi44 < 1; fi44++ {
                var v45 int16 = -1
                if (!(2147483647 != (-(v4)))) {
                    v3 -= (v3 - (v3 * v3))
                } else {
                    if true {
                        switch 128 {
                        case 4294967295:
                            var v46 string = (v1 + "")
                            _ = v46
                        }
                    } else {
                        var v47 uint8 = 2
                        if true {
                            v47 = (^(v47))
                            v48 := GenS4{F0: (^(v4))}
                            fmt.Println(v36, v4)
                            _ = v48
                        } else {
                            var v49 uint8 = (^(v47))
                            _ = v49
                        }
                        fmt.Println(v4, v36)
                        _ = v47
                    }
                    v36 = 4294967295
                }
                v4 -= (v4 << ((4294967295) & 63))
                _ = fi44
                _ = v45
            }
        case "hello":
            for fi50 := 0; fi50 < 19; fi50++ {
                fmt.Println(v2, v1, fi50)
                v4 -= (v4 + genFn0((1.5 != (1.5 / 1.0)), 42))
                _ = fi50
            }
        default:
            v51 := GenS5{F0: (v4 & (^(v4))), F1: v1}
            _ = v51
        }
        if false {
            for fi52 := 0; fi52 < 9; fi52++ {
                var v53 uint16 = 978
                var v54 int16 = 32767
                _ = fi52
                _ = v53
                _ = v54
            }
            var v55 uint32 = v3
            if ((!(v55 != v55)) || (!false)) {
                var v56 int16 = 1
                v36 -= (^(v36))
                v56 += 42
                _ = v56
            } else {
                v57 := []int{fi43, (fi43 & (fi43 ^ (fi43 + (fi43 * 715))))}
                _ = v57
            }
            _ = v55
        }
        _ = fi43
    }
    if false {
        switch v1 {
        case "a":
            switch 561 {
            case 708:
                v58 := GenS6{F0: 32767}
                _ = v58
            case -9223372036854775808:
                for fi59 := 0; fi59 < 8; fi59++ {
                    if (1 == -9223372036854775808) {
                        var v60 uint16 = 32767
                        _ = v60
                    } else {
                        var v61 int = fi59
                        v62 := GenS2{F0: 2147483648, F1: (v2 + -498), F2: 65535}
                        _ = v61
                        _ = v62
                    }
                    _ = fi59
                }
            default:
                v1 = v1
            }
        case "hello":
            if (127 <= 127) {
                var v63 uint32 = (v3 + v3)
                _ = v63
            }
        case "":
            v64 := []int{}
            _ = v64
        }
    }
    if (!(126 < 788)) {
        var v65 int8 = 0
        _ = v65
    } else {
        fmt.Println(v4, v1)
        var v66 bool = false
        for fi67 := 0; fi67 < 10; fi67++ {
            var v68 int8 = 127
            v69 := GenS2{F0: (v3 << ((873) & 63)), F1: (v2 ^ 1), F2: 4294967295}
            _ = fi67
            _ = v68
            _ = v69
        }
        _ = v66
    }
    v4 = v4
    for fi70 := 0; fi70 < 13; fi70++ {
        switch 910 {
        case -2147483648:
            for fi71 := 0; fi71 < 20; fi71++ {
                var v72 int8 = 42
                v73 := GenS7{F0: -128}
                _ = fi71
                _ = v72
                _ = v73
            }
        case 881:
            fmt.Println(v4)
        }
        _ = fi70
    }
    switch v1 {
    case "hello":
        v36 += v36
    case "\n\t":
        fmt.Println(v4, v1)
    default:
        v36 -= (v36 | (v36 ^ (v36 - 18446744073709551615)))
    }
    var v74 int16 = int16(v3)
    v75 := []int{128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v74:", v74)
    fmt.Println("slicelen v75:", len(v75), v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v74
    _ = v75
    _ = v9
}
