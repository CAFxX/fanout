package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 int32; F2 int64 }
type GenS6 struct { F0 bool; F1 uint32; F2 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 int64; F2 uint8 }
func genFn0(p0 uint64) uint {
    var v7 string = "hello"
    var v8 bool = ((!false) || true)
    p0 = p0
    _ = p0
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 uint8 = 1
    v2 := []uint32{2, 739}
    var v3 int8 = 39
    var v4 uint8 = v1
    v5 := make(map[uint32]string)
    var v6 int16 = -545
    var v9 uint8 = v1
    if false {
        v1 -= v1
    }
    var v10 uint8 = 128
    v11 := GenS1{F0: 0, F1: 65535, F2: (^(v3))}
    v12 := GenS2{F0: ((v1 * v10) > v9)}
    var v13 int32 = 91
    _ = len(v5)
    if false {
        switch 127 {
        case -2147483648:
            v14 := GenS2{F0: false}
            _ = v14
        case 128:
            for fi15 := 0; fi15 < 17; fi15++ {
                if false {
                    _ = len(v2)
                    for fi16 := 0; fi16 < 2; fi16++ {
                        var v17 uint16 = 32767
                        _ = fi16
                        _ = v17
                    }
                }
                var v18 uint8 = v4
                _ = fi15
                _ = v18
            }
        case 32767:
            if (32767 >= v6) {
                v9 = (^(v4))
            }
        default:
            for fi19 := 0; fi19 < 6; fi19++ {
                for fi20 := 0; fi20 < 6; fi20++ {
                    if ((32767 != 889) || true) {
                        var v21 uint8 = v4
                        switch "\n\t" {
                        case "foo bar":
                            switch (fi19 * fi19) {
                            case -1:
                                v22 := GenS1{F0: 2147483647, F1: 2, F2: (v3 - -128)}
                                _ = v22
                            case 0:
                                switch "a" {
                                case "":
                                    var v23 uint32 = 2
                                    _ = v23
                                }
                            case 32767:
                                var v24 float64 = (1.5 + (1e10 - -2.25))
                                _ = v24
                            }
                        case "":
                            var v25 float64 = (1.5 * -2.25)
                            _ = v25
                        default:
                            v26 := GenS1{F0: 42, F1: 2, F2: -1}
                            _ = v26
                        }
                        _ = v21
                    } else {
                        fmt.Println(v13)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v27 int64 = int64(v6)
                        v3 = (v3 & v3)
                        var v28 uint16 = 38
                        _ = v27
                        _ = v28
                    }
                    _ = fi20
                }
                v4 -= 2
                _ = fi19
            }
        }
        v29 := GenS3{F0: 955, F1: v13}
        v10 -= (v1 + 107)
        _ = v29
    } else {
        if ((-(v3)) >= v3) {
            var v30 int = 416
            v31 := GenS4{F0: v6}
            switch (v30 | v30) {
            case 896:
                v5[277] = "foo bar"
            case 1:
                switch "\n\t" {
                case "":
                    switch "\n\t" {
                    case "foo bar":
                        v2 = append(v2, 962)
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
            case 2147483647:
                v3 += 1
            default:
                fmt.Println(v4, v3, v6)
            }
            _ = v30
            _ = v31
        }
        v2 = append(v2, 2147483648)
        switch 42 {
        case 1:
            if (0 <= 996) {
                v9 = (v4 << ((genFn0(2)) & 63))
                fmt.Println(v13, v3, v4)
            }
        case -121:
            switch "a" {
            case "hello":
                v3 -= -1
            case "foo bar":
                fmt.Println(v9)
            }
        case -9223372036854775808:
            v9 -= v10
        default:
            for fi32 := 0; fi32 < 2; fi32++ {
                for fi33 := 0; fi33 < 3; fi33++ {
                    v4 = v9
                    switch int(v3) {
                    case 1:
                        fmt.Println(v9, v3)
                    case 4294967295:
                        v6 += (v6 & v6)
                    case 2147483647:
                        var v34 float32 = (3.14159 - (0.0 - 3.14159))
                        _ = v34
                    default:
                        var v35 int = fi33
                        switch (fi32 & (^(fi32))) {
                        case 2147483647:
                            fmt.Println(fi32, fi32, v35)
                        case 1:
                            if len(v2) > 0 {
                                v2[0] = 138
                            }
                        default:
                            v1 = v9
                        }
                        v2 = append(v2, 65535)
                        _ = v35
                    }
                    var v36 int16 = v6
                    _ = fi33
                    _ = v36
                }
                var v37 int64 = 42
                _ = fi32
                _ = v37
            }
        }
    }
    if ((!((42 <= genFn0(1)) || true)) || true) {
        v13 = v13
        var v38 int64 = 965
        _ = v38
    } else {
        if false {
            for fi39 := 0; fi39 < 15; fi39++ {
                var v40 uint32 = 1
                v6 -= (v6 | (v6 * 562))
                _ = fi39
                _ = v40
            }
            v2 = append(v2, 971)
        }
        var v41 int = -2147483648
        if false {
            var v42 int16 = v6
            _ = v42
        } else {
            fmt.Println(v4, v3, v10)
            if (float64(0.0) <= 1.5) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        _ = v41
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        switch "hello" {
        case "a":
            fmt.Println(v1, v1)
        case "":
            fmt.Println(v6, fi43)
        case "foo bar":
            for fi44 := 0; fi44 < 3; fi44++ {
                var v45 int32 = (v13 * v13)
                if false {
                    v6 = 42
                }
                fmt.Println(v1)
                v10 = v10
                _ = fi44
                _ = v45
            }
        default:
            switch (fi43 ^ 2147483647) {
            case 1:
                fmt.Println(v9)
            }
        }
        for fi46 := 0; fi46 < 3; fi46++ {
            if false {
                v47 := GenS5{F0: uint(v4), F1: v13, F2: 9223372036854775807}
                v2 = append(v2, 1)
                var v48 float32 = (1.5 + (3.14159 / 1.0))
                var v49 uint = 0
                _ = v47
                _ = v48
                _ = v49
            }
            switch -120 {
            case -510:
                var v50 string = "a"
                _ = v50
            case -121:
                switch 128 {
                case -2147483648:
                    var v51 float32 = float32(float64(0.0))
                    _ = v51
                case 32768:
                    if (!false) {
                        if (true && true) {
                            v52 := GenS5{F0: genFn0(18446744073709551615), F1: (v13 >> ((4294967295) & 63)), F2: 9223372036854775807}
                            v3 = v3
                            _ = v52
                        }
                    }
                case 247:
                    v4 += v9
                default:
                    switch (((("a" + "") + "foo bar") + "") + "") {
                    case "hello":
                        v5[4294967295] = ""
                    }
                }
            case -531:
                delete(v5, 1)
            }
            _ = fi46
        }
        _ = fi43
    }
    _ = len(v2)
    v53 := GenS6{F0: false, F1: 4294967295, F2: 9223372036854775807}
    v5[36] = ("a" + ("foo bar" + "foo bar"))
    for fi54 := 0; fi54 < 10; fi54++ {
        v55 := GenS7{F0: 435}
        var v56 int8 = (^(v3))
        switch "a" {
        case "\n\t":
            var v57 uint = 0
            _ = v57
        default:
            v58 := GenS8{F0: 0, F1: 1, F2: v9}
            _ = v58
        }
        _ = fi54
        _ = v55
        _ = v56
    }
    for fi59 := 0; fi59 < 13; fi59++ {
        if false {
            if true {
                if (!(v13 < v13)) {
                    for fi60 := 0; fi60 < 12; fi60++ {
                        v4 = v1
                        _ = fi60
                    }
                    v61 := GenS1{F0: 4294967295, F1: 32767, F2: (^(v3))}
                    var v62 string = ("a" + ("\n\t" + "foo bar"))
                    _ = v61
                    _ = v62
                }
                if (1 <= 97) {
                    switch "\n\t" {
                    case "a":
                        fmt.Println(fi59, v13, v4)
                    case "\n\t":
                        v63 := GenS7{F0: 2147483647}
                        _ = v63
                    default:
                        var v64 int64 = 1
                        _ = v64
                    }
                }
            }
            for fi65 := 0; fi65 < 10; fi65++ {
                switch fi59 {
                case -9223372036854775808:
                    _ = len(v2)
                case 4294967295:
                    fmt.Println(fi59)
                }
                var v66 int64 = -2
                var v67 int16 = v6
                var v68 int32 = -216
                _ = fi65
                _ = v66
                _ = v67
                _ = v68
            }
        }
        _ = fi59
    }
    v6 += 794
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v9
}
