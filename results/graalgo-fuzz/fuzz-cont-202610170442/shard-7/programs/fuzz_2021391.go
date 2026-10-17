package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 int; F1 int16; F2 bool }
type GenS9 struct { F0 uint64; F1 int64; F2 uint32 }
func genFn0(p0 int8, p1 string, p2 int8) uint32 {
    v5 := GenS1{F0: -32768, F1: 42, F2: -1}
    p1 = p1
    switch p1 {
    case "a":
        p2 = p0
    case "\n\t":
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := []int{127, 1, 275}
            v8 := GenS2{F0: 2147483647}
            switch (fi6 & 32768) {
            case 752:
                if true {
                    fmt.Println(p0, p2)
                    fmt.Println(p1)
                } else {
                    var v9 uint16 = 306
                    var v10 int = 0
                    _ = v10
                    _ = v9
                }
            case -9223372036854775808:
                v11 := GenS1{F0: 32767, F1: 677, F2: 0}
                _ = v11
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    case "foo bar":
        p1 = ""
    default:
        var v12 int16 = -32768
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 378
}
func genFn1() int8 {
    var v13 int32 = 42
    for fi14 := 0; fi14 < 10; fi14++ {
        if (true && (v13 > v13)) {
            v15 := []int{fi14, -9223372036854775808, fi14, (fi14 - fi14)}
            _ = v15
        } else {
            v16 := []int{fi14, 688, fi14, (fi14 & (fi14 ^ (fi14 * fi14)))}
            var v17 uint16 = 42
            _ = v16
            _ = v17
        }
        switch 128 {
        case -873:
            if false {
                var v18 int = fi14
                var v19 int = 32768
                v19 -= 32767
                _ = v18
                _ = v19
            }
        case 32767:
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 bool = true
                switch (fi20 << ((419) & 63)) {
                case 1:
                    var v22 bool = v21
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
        case 2147483647:
            var v23 uint16 = 255
            _ = v23
        default:
            var v24 uint = 0
            _ = v24
        }
        _ = fi14
    }
    switch 2147483647 {
    case 0:
        v13 += v13
    case 128:
        v13 = (v13 % (((v13 ^ 42)) | 1))
    }
    _ = v13
    return 100
}

func main() {
    var v1 bool = false
    var v2 uint8 = 255
    var v3 uint16 = 2
    var v4 float32 = 3.14159
    fmt.Println(v3, v1, v4)
    v3 += 65535
    var v25 bool = (42 <= 32768)
    v1 = (1000 <= 0)
    v1 = ((42 > 930) || (!true))
    v26 := []int{42, 1}
    v27 := GenS2{F0: 443}
    var v28 uint = 0
    var v29 float32 = 0.1
    v30 := GenS2{F0: 0}
    v26 = append(v26, -232)
    if (!false) {
        var v31 uint32 = uint32(v2)
        var v32 float32 = (v29 / v4)
        v3 = (^(v3))
        _ = v31
        _ = v32
    } else {
        v33 := GenS1{F0: 42, F1: -614, F2: 1}
        _ = v33
    }
    _ = len(v26)
    if true {
        v34 := GenS2{F0: 209}
        if v25 {
            v3 -= (^(v3))
            for fi35 := 0; fi35 < 7; fi35++ {
                fmt.Println(v1, v2)
                v36 := GenS1{F0: 1, F1: -717, F2: 42}
                switch "hello" {
                case "foo bar":
                    v26 = append(v26, (fi35 % ((-375) | 1)))
                }
                _ = fi35
                _ = v36
            }
            if v1 {
                if (!(0 > 1)) {
                    var v37 uint64 = 42
                    _ = v37
                } else {
                    var v38 uint = 663
                    for fi39 := 0; fi39 < 11; fi39++ {
                        var v40 uint32 = 374
                        fmt.Println(v40, v28, v4)
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                }
            } else {
                switch "\n\t" {
                case "":
                    if (!(!(869 < v28))) {
                        for fi41 := 0; fi41 < 9; fi41++ {
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                            _ = fi41
                        }
                        var v42 uint = v28
                        _ = v42
                    } else {
                        var v43 int = 257
                        _ = v43
                    }
                case "hello":
                    for fi44 := 0; fi44 < 14; fi44++ {
                        var v45 int8 = 0
                        _ = fi44
                        _ = v45
                    }
                case "foo bar":
                    for fi46 := 0; fi46 < 7; fi46++ {
                        for fi47 := 0; fi47 < 8; fi47++ {
                            v48 := GenS3{F0: (^(v28)), F1: "a"}
                            if v25 {
                                v49 := GenS4{F0: 18446744073709551615}
                                switch "a" {
                                case "a":
                                    _ = len(v26)
                                }
                                _ = v49
                            }
                            var v50 int32 = 42
                            _ = fi47
                            _ = v48
                            _ = v50
                        }
                        _ = fi46
                    }
                }
                if (!v25) {
                    v29 += 1.5
                }
            }
        }
        var v51 bool = v25
        _ = v34
        _ = v51
    }
    if v1 {
        if (3.14159 != 3.14159) {
            v25 = (3.14159 == (0.0 * -2.25))
            v52 := GenS5{F0: (("\n\t" + "") + (("a" + "hello") + ("foo bar" + ("hello" + ("a" + "a")))))}
            if ((((((v29 + float32(1e10)) - 3.14159) - (v4 * v4)) * (v29 + (v4 / (1.5 / ((v29 - (0.1 + (v4 / 1.5))) * v4))))) * 3.14159) != 0.1) {
                v53 := GenS6{F0: genFn0(genFn1(), ("\n\t" + "a"), 100), F1: (^(v28))}
                v28 = 0
                _ = v53
            } else {
                for fi54 := 0; fi54 < 11; fi54++ {
                    if false {
                        fmt.Println(v2, fi54)
                        for fi55 := 0; fi55 < 20; fi55++ {
                            v1 = (9223372036854775807 > 1)
                            fmt.Println(fi55, v4, v29)
                            fmt.Println(v3, v29)
                            _ = fi55
                        }
                    }
                    v1 = (fi54 > (fi54 + (fi54 & 677)))
                    if (!false) {
                        v56 := GenS7{F0: int8(v2), F1: -1}
                        var v57 uint8 = v2
                        v58 := GenS8{F0: (-(fi54)), F1: 42, F2: false}
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    _ = fi54
                }
                v26 = append(v26, 128)
            }
            _ = v52
        } else {
            switch 127 {
            case 9223372036854775807:
                switch 2147483647 {
                case 0:
                    switch 32767 {
                    case -9223372036854775808:
                        _ = len(v26)
                    default:
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                    }
                default:
                    v29 = (1e10 * 1e10)
                }
            case 128:
                v59 := GenS7{F0: 100, F1: 42}
                _ = v59
            default:
                var v60 int16 = -32768
                _ = v60
            }
            if ((v2 << (((^(v28))) & 63)) == (v2 << (((^(v28))) & 63))) {
                v3 -= v3
            } else {
                v28 = v28
                var v61 int32 = -1
                v62 := GenS4{F0: 371}
                _ = v61
                _ = v62
            }
            var v63 int64 = 42
            _ = v63
        }
        var v64 int = 128
        if (genFn1() < 10) {
            var v65 int32 = 199
            v26 = append(v26, (v64 % ((v64) | 1)))
            _ = v65
        }
        _ = v64
    }
    if len(v26) > 0 {
        v26[0] = 32767
    }
    v66 := GenS9{F0: 0, F1: -1, F2: 1}
    v1 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v66
}
