package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 int64; F2 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 int32 }

func main() {
    var v1 int32 = 0
    var v2 uint8 = 42
    var v3 uint32 = 1
    var v4 string = ("\n\t" + "")
    var v5 int64 = -443
    v1 = 111
    v2 -= (v2 + (v2 * v2))
    fmt.Println(v2, v4, v2)
    v6 := GenS1{F0: 0, F1: 1}
    v5 = v5
    var v7 int32 = v1
    for fi8 := 0; fi8 < 15; fi8++ {
        v3 += (v3 << ((255) & 63))
        var v9 int32 = -1
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v4, v7)
            var v11 bool = (true || false)
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v12 float64 = 1e10
    v13 := []int{-325}
    v14 := GenS1{F0: 538, F1: 1}
    switch 1 {
    case -2147483648:
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := GenS2{F0: (v2 >> ((2) & 63)), F1: uint16(v2)}
            _ = fi15
            _ = v16
        }
    case 1:
        var v17 uint = 980
        _ = v17
    default:
        if len(v13) > 0 {
            v13[0] = 486
        }
    }
    var v18 uint64 = 2
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            switch "\n\t" {
            case "":
                v21 := GenS3{F0: 42}
                _ = v21
            case "a":
                var v22 uint16 = 65535
                fmt.Println(v7, fi19, v7)
                fmt.Println(v5)
                _ = v22
            case "\n\t":
                v2 -= v2
            default:
                v23 := GenS3{F0: 674}
                _ = v23
            }
            fmt.Println(v1, fi20, v3)
            var v24 int32 = (-(v1))
            _ = fi20
            _ = v24
        }
        if (4294967295 < 4294967295) {
            for fi25 := 0; fi25 < 7; fi25++ {
                v18 -= (v18 ^ v18)
                fmt.Println(v1)
                fmt.Println(v5, v7, v7)
                _ = fi25
            }
        } else {
            if len(v13) > 0 {
                v13[0] = fi19
            }
        }
        if true {
            v26 := GenS2{F0: 1, F1: 891}
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 string = (v4 + ((v4 + "hello") + v4))
                if false {
                    v29 := GenS2{F0: v2, F1: 42}
                    v4 = "foo bar"
                    _ = v29
                }
                if true {
                    if ((v18 - 0) < v18) {
                        if ((v2 <= (v2 % (((v2 & 99)) | 1))) || (v12 < v12)) {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            var v30 float32 = float32(((v12 / 3.14159) / v12))
                            _ = v30
                        }
                        if true {
                            switch "\n\t" {
                            case "a":
                                if ((!true) && false) {
                                    if (!(v28 < "hello")) {
                                        var v31 float64 = float64(0.0)
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        if len(v13) > 0 {
                                            v13[0] = 4294967295
                                        }
                                        var v32 uint32 = (v3 | (v3 / ((v3) | 1)))
                                        v33 := GenS4{F0: v5}
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v5 = (v5 - v5)
                                    }
                                    fmt.Println(v12, v3, fi19)
                                }
                            }
                            v3 = (v3 ^ v3)
                            if (float64(1e10) == v12) {
                                switch v28 {
                                case "hello":
                                    if true {
                                        switch (fi27 / (((fi27 & fi19)) | 1)) {
                                        case -360:
                                            v34 := GenS5{F0: 42}
                                            _ = v34
                                        default:
                                            var v35 int64 = 42
                                            switch fi27 {
                                            case -1:
                                                var v36 int16 = 42
                                                _ = v36
                                            case 2147483647:
                                                v28 = "foo bar"
                                            case -2147483648:
                                                v28 = ""
                                            default:
                                                var v37 float32 = 1.5
                                                _ = v37
                                            }
                                            _ = v35
                                        }
                                    } else {
                                        v38 := GenS2{F0: v2, F1: 43}
                                        _ = v38
                                    }
                                }
                                if true {
                                    var v39 int16 = -32768
                                    if (v4 < v28) {
                                        v28 = ((v28 + ("a" + v4)) + v28)
                                        v40 := GenS4{F0: (v5 / ((v5) | 1))}
                                        _ = v40
                                    } else {
                                        fmt.Println(v12)
                                    }
                                    _ = v39
                                } else {
                                    fmt.Println(v7, v1, v28)
                                    v13 = append(v13, fi19)
                                    v41 := GenS6{F0: -102}
                                    var v42 int = fi27
                                    fmt.Println(v12, v5, v18)
                                    switch (v42 << ((255) & 63)) {
                                    case -599:
                                        v13 = append(v13, (v42 - 4294967295))
                                    case 1:
                                        v28 = "foo bar"
                                    }
                                    if (!(v4 < v4)) {
                                        if len(v13) > 0 {
                                            v13[0] = (fi27 + fi27)
                                        }
                                        v13 = append(v13, (v42 / ((-2147483648) | 1)))
                                        switch v28 {
                                        case "a":
                                            v5 = (v5 << ((4294967295) & 63))
                                        case "foo bar":
                                            v43 := GenS1{F0: 1, F1: (v18 - 96)}
                                            _ = v43
                                        }
                                    } else {
                                        v4 = v28
                                        var v44 int32 = (-(v7))
                                        _ = v44
                                    }
                                    v7 = 845
                                    _ = v41
                                    _ = v42
                                }
                                fmt.Println(v4, v1)
                            } else {
                                _ = len(v13)
                                var v45 uint = 2
                                v45 -= (v45 / ((v45) | 1))
                                _ = v45
                            }
                        } else {
                            v7 = -2147483648
                            fmt.Println(v2)
                        }
                    } else {
                        if true {
                            v46 := GenS2{F0: v2, F1: 255}
                            if len(v13) > 0 {
                                v13[0] = (^(fi27))
                            }
                            switch fi27 {
                            case 32768:
                                v47 := GenS7{F0: v5, F1: (v5 | (v5 | (v5 + 1))), F2: 1}
                                _ = v47
                            case -1:
                                if false {
                                    v13 = append(v13, fi19)
                                    _ = len(v13)
                                } else {
                                    v48 := GenS8{F0: 212}
                                    _ = v48
                                }
                            case -783:
                                v18 -= (^(v18))
                            default:
                                if true {
                                    switch "" {
                                    case "hello":
                                        v49 := GenS3{F0: 255}
                                        _ = v49
                                    case "":
                                        switch (("a" + ("foo bar" + "hello")) + v28) {
                                        case "":
                                            v3 = 65535
                                        case "\n\t":
                                            var v50 uint16 = 65535
                                            _ = v50
                                        case "hello":
                                            v51 := GenS7{F0: v5, F1: v5, F2: 0}
                                            _ = v51
                                        }
                                    }
                                    v3 = v3
                                } else {
                                    v13 = append(v13, (fi27 & fi19))
                                    if false {
                                        var v52 uint16 = 32767
                                        v28 = (v28 + "a")
                                        _ = v52
                                    } else {
                                        var v53 int16 = -32768
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        v54 := GenS9{F0: "", F1: (v7 | v7)}
                                        if len(v13) > 0 {
                                            v13[0] = -2147483648
                                        }
                                        _ = v53
                                        _ = v54
                                    }
                                }
                            }
                            _ = v46
                        }
                        switch v4 {
                        case "hello":
                            var v55 int16 = 364
                            _ = v55
                        case "foo bar":
                            var v56 uint16 = 65535
                            _ = v56
                        }
                        if true {
                            v57 := GenS3{F0: 205}
                            _ = v57
                        }
                    }
                }
                _ = fi27
                _ = v28
            }
            _ = v26
        } else {
            var v58 int = (fi19 - 32767)
            var v59 int16 = -1
            fmt.Println(v59, v1, fi19)
            _ = v58
            _ = v59
        }
        _ = fi19
    }
    if ((float64((0.0 - ((-2.25 + 1e10) / 3.14159))) - (float64(((0.0 / (0.0 + 1e10)) * 0.0)) / (3.14159 + 0.0))) != (((0.0 * v12) / 1.5) / (v12 / 1.0))) {
        v60 := GenS2{F0: 3, F1: 1000}
        _ = v60
    } else {
        v61 := GenS9{F0: v4, F1: (v1 << ((4294967295) & 63))}
        var v62 int8 = 0
        switch "a" {
        case "":
            var v63 string = "a"
            _ = v63
        case "hello":
            var v64 float64 = v12
            _ = v64
        case "\n\t":
            v12 = 3.14159
        default:
            var v65 uint16 = 993
            _ = v65
        }
        _ = v61
        _ = v62
    }
    fmt.Println(v3, v18, v3)
    switch ("hello" + v4) {
    case "a":
        fmt.Println(v1, v3, v5)
    case "":
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    case "hello":
        if len(v13) > 0 {
            v13[0] = 960
        }
    }
    v7 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
