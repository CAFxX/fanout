package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{42, 4294967295, 42, 2}
    v3 := make(map[int]int)
    var v4 float32 = float32(1e10)
    var v5 int = 627
    if (1 >= 42) {
        var v6 int = v5
        _ = v6
    }
    v7 := GenS1{F0: -728}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch (v5 + (v5 | (v5 >> ((42) & 63)))) {
    case 9223372036854775807:
        if (((0.1 / (1.5 / v4)) * ((v4 * v4) + 0.0)) == 0.1) {
            v5 = v5
            if len(v2) > 0 {
                v2[0] = 682
            }
        } else {
            _ = len(v2)
        }
    case 724:
        if true {
            switch 4294967295 {
            case 0:
                if (false && (((true && (-1 <= 42)) && (1 <= -1)) || (true && (!((374 > -32768) && (-783 > 42)))))) {
                    v5 = (v5 >> ((1) & 63))
                    v8 := GenS1{F0: int32(v5)}
                    var v9 int16 = -32768
                    _ = v8
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 18; fi10++ {
                        v4 = (0.1 / 3.14159)
                        v4 -= v4
                        _ = fi10
                    }
                    v4 = ((1e10 / 1.0) / -2.25)
                    if ((0.0 + v4) <= 0.0) {
                        v11 := GenS1{F0: 2147483647}
                        for fi12 := 0; fi12 < 5; fi12++ {
                            var v13 int16 = 42
                            v13 -= v13
                            v4 -= float32(1e10)
                            _ = fi12
                            _ = v13
                        }
                        v5 += v5
                        _ = v11
                    } else {
                        v4 = (float32((0.0 / 1.0)) + -2.25)
                    }
                }
            case 1:
                var v14 uint8 = 1
                _ = v14
            case 2147483647:
                var v15 int64 = 42
                _ = v15
            default:
                if (!(false && (false && false))) {
                    fmt.Println(v5, v5)
                    v16 := GenS1{F0: -2147483648}
                    v5 -= (v5 ^ v5)
                    _ = v16
                } else {
                    fmt.Println(v4)
                    v4 += float32(-2.25)
                }
            }
            fmt.Println(v5, v5)
        } else {
            v5 += 9223372036854775807
        }
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v4, fi17)
        _ = fi17
    }
    var v18 uint32 = 1
    switch (^(v5)) {
    case 127:
        switch -1 {
        case 1:
            switch ("hello" + "") {
            case "":
                var v19 uint32 = 791
                _ = v19
            case "\n\t":
                for fi20 := 0; fi20 < 18; fi20++ {
                    v21 := GenS2{F0: 814, F1: -574, F2: 1}
                    switch (v5 % (((-(fi20))) | 1)) {
                    case 0:
                        for fi22 := 0; fi22 < 17; fi22++ {
                            fmt.Println(fi22, v5, fi20)
                            v23 := GenS1{F0: -815}
                            switch "hello" {
                            case "":
                                if (!false) {
                                    v4 = ((1.5 * 1.5) + 0.1)
                                    if (true && true) {
                                        var v24 int32 = 0
                                        fmt.Println(v18)
                                        _ = v24
                                    } else {
                                        v18 = v18
                                        v25 := GenS1{F0: -466}
                                        _ = v25
                                    }
                                }
                            case "foo bar":
                                var v26 int32 = 94
                                v4 = v4
                                v3[128] = (fi22 & (fi22 - fi20))
                                switch -2147483648 {
                                case 217:
                                    var v27 int16 = 32767
                                    _ = v27
                                case -9223372036854775808:
                                    v28 := GenS2{F0: 1, F1: v26, F2: 32767}
                                    _ = v28
                                default:
                                    if (798 < 68) {
                                        v29 := GenS1{F0: 619}
                                        if (1 < 963) {
                                            var v30 uint = 7
                                            switch v5 {
                                            case 127:
                                                fmt.Println(fi20, v30)
                                            }
                                            _ = v30
                                        }
                                        if false {
                                            fmt.Println(v4)
                                            var v31 string = "foo bar"
                                            _ = v31
                                        } else {
                                            v32 := GenS2{F0: 42, F1: (v26 << ((1) & 63)), F2: -857}
                                            v26 += 0
                                            if (1 < 0) {
                                                var v33 int64 = -9223372036854775808
                                                var v34 int32 = v26
                                                _ = v33
                                                _ = v34
                                            }
                                            v35 := GenS1{F0: (v26 - (v26 & v26))}
                                            v5 = 9223372036854775807
                                            v4 = (float32((0.1 / 3.14159)) * v4)
                                            _ = v32
                                            _ = v35
                                        }
                                        v4 -= v4
                                        _ = v29
                                    } else {
                                        var v36 int8 = -19
                                        fmt.Println(fi22, fi20)
                                        _ = v36
                                    }
                                    fmt.Println(v18)
                                    fmt.Println(fi22, v26, fi22)
                                    var v37 int64 = 0
                                    if false {
                                        var v38 int32 = v26
                                        v39 := GenS3{F0: 0}
                                        v37 -= v37
                                        _ = v38
                                        _ = v39
                                    }
                                    if true {
                                        v26 = 42
                                        fmt.Println(v18, v26)
                                    } else {
                                        _ = len(v1)
                                        v18 += v18
                                        fmt.Println(v4, v5, fi20)
                                        v2 = append(v2, 0)
                                        v2 = append(v2, 65535)
                                        var v40 uint32 = (^(v18))
                                        _ = v40
                                    }
                                    v41 := v3[(fi20 - fi20)]
                                    if true {
                                        v26 = v26
                                    }
                                    var v42 float32 = (1e10 / ((v4 - v4) * 1.5))
                                    _ = v37
                                    _ = v41
                                    _ = v42
                                }
                                delete(v3, (fi20 / (((fi20 << ((42) & 63))) | 1)))
                                _ = v26
                            }
                            _ = fi22
                            _ = v23
                        }
                    default:
                        var v43 int64 = 1
                        _ = v43
                    }
                    _ = fi20
                    _ = v21
                }
            case "foo bar":
                switch "a" {
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        v44 := GenS3{F0: 1}
                        _ = v44
                    default:
                        fmt.Println(v18, v4, v18)
                    }
                case "hello":
                    var v45 uint32 = (v18 - 2)
                    _ = v45
                }
            default:
                if (!("a" == ("foo bar" + ("hello" + "")))) {
                    for fi46 := 0; fi46 < 11; fi46++ {
                        delete(v1, v18)
                        var v47 uint = 4294967295
                        _ = fi46
                        _ = v47
                    }
                    var v48 bool = (18446744073709551615 <= 42)
                    v5 += (v5 + -31)
                    _ = v48
                }
            }
        case 9223372036854775807:
            v4 = float32(((float64(0.1) + -2.25) / (0.0 - 1.5)))
        case 128:
            var v49 int64 = -309
            _ = v49
        }
    case 32768:
        if false {
            v50 := GenS2{F0: 0, F1: 0, F2: -1}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v2) > 0 {
                v2[0] = v18
            }
            _ = v50
        } else {
            var v51 int64 = 515
            _ = v51
        }
    default:
        var v52 float64 = 1.5
        _ = v52
    }
    var v53 uint8 = 0
    switch "hello" {
    case "foo bar":
        switch v5 {
        case -829:
            fmt.Println(v18)
        }
    default:
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v54 := GenS3{F0: 323}
            _ = v54
        }
    }
    v5 = v5
    v55 := GenS1{F0: -1}
    v56 := GenS2{F0: 2, F1: -2147483648, F2: -1}
    var v57 int16 = -651
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v55
    _ = v56
    _ = v57
    _ = v7
}
