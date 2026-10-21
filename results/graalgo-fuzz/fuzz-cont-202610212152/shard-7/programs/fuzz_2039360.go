package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint8; F1 int64; F2 uint32 }
func genFn0(p0 string, p1 string) int {
    switch 0 {
    case -399:
        if true {
            v6 := []int{268, 0, 4294967295, 128}
            for fi7 := 0; fi7 < 14; fi7++ {
                switch p0 {
                case "":
                    for fi8 := 0; fi8 < 11; fi8++ {
                        if false {
                            v9 := GenS1{F0: 852, F1: 42, F2: 127}
                            var v10 int32 = 42
                            var v11 uint8 = 42
                            _ = v10
                            _ = v11
                            _ = v9
                        }
                        var v12 uint16 = 32767
                        p0 = p0
                        _ = fi8
                        _ = v12
                    }
                case "foo bar":
                    var v13 uint32 = 42
                    _ = v13
                case "hello":
                    for fi14 := 0; fi14 < 13; fi14++ {
                        switch (-(fi7)) {
                        case 2147483647:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            p1 = p1
                        }
                        _ = fi14
                    }
                }
                switch (fi7 % ((-1) | 1)) {
                case 2147483647:
                    switch fi7 {
                    case -9223372036854775808:
                        fmt.Println(p0)
                    }
                case -643:
                    fmt.Println(fi7)
                case 32768:
                    if ((0 <= 2) && (((65535 > 82) || (false || (75 <= -115))) || (!(false && (255 == 2))))) {
                        if true {
                            for fi15 := 0; fi15 < 10; fi15++ {
                                var v16 int16 = -1
                                v16 = (v16 & v16)
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if len(v6) > 0 {
                                    v6[0] = (fi7 * (fi7 << ((380) & 63)))
                                }
                                _ = fi15
                                _ = v16
                            }
                            var v17 int16 = -816
                            _ = v17
                        }
                    }
                default:
                    var v18 int32 = -2147483648
                    _ = v18
                }
                _ = fi7
            }
            _ = len(v6)
            _ = v6
        } else {
            p0 = p0
        }
    case 1:
        var v19 string = p1
        _ = v19
    case 9223372036854775807:
        var v20 uint = 4294967295
        _ = v20
    default:
        v21 := GenS2{F0: 0}
        _ = v21
    }
    p1 = (p1 + (p1 + (p1 + p0)))
    if (((true && false) && true) && (127 == -57)) {
        v22 := GenS1{F0: 18446744073709551615, F1: 752, F2: 0}
        _ = v22
    }
    var v23 int16 = 1
    _ = p0
    _ = p1
    _ = v23
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 2147483648
    var v3 int16 = 0
    v4 := make(map[string]uint32)
    var v5 uint32 = v2
    var v24 bool = (3.14159 != (1.5 - 1.5))
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := []int{}
        _ = fi25
        _ = v26
    }
    var v27 int8 = 76
    var v28 int64 = -9223372036854775808
    fmt.Println(v27, v27, v27)
    fmt.Println(v5, v5)
    switch "\n\t" {
    case "":
        v29 := v1[4294967295]
        _ = v29
    case "hello":
        v30 := []int{-578, 128, 4294967295, 128}
        _ = v30
    }
    v28 += -9223372036854775808
    v31 := GenS2{F0: v28}
    switch "hello" {
    case "a":
        switch -778 {
        case 0:
            var v32 int8 = (v27 & v27)
            _ = v32
        case 4294967295:
            v5 -= v2
        default:
            var v33 uint32 = v2
            _ = v33
        }
    case "":
        for fi34 := 0; fi34 < 18; fi34++ {
            switch -1 {
            case 128:
                var v35 string = "foo bar"
                _ = v35
            case 127:
                var v36 uint = 122
                _ = v36
            case 1:
                v37 := GenS1{F0: 621, F1: v5, F2: int8(v2)}
                _ = v37
            default:
                v38 := GenS1{F0: 18446744073709551615, F1: v5, F2: -1}
                _ = v38
            }
            fmt.Println(fi34, v24)
            for fi39 := 0; fi39 < 10; fi39++ {
                v40 := GenS3{F0: 8, F1: 143, F2: 121}
                v41 := GenS3{F0: 0, F1: 18446744073709551615, F2: 42}
                v42 := GenS3{F0: 255, F1: 65535, F2: 42}
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
            _ = fi34
        }
    default:
        v43 := GenS4{F0: v28, F1: 42}
        _ = v43
    }
    v44 := GenS2{F0: (v28 >> ((688) & 63))}
    fmt.Println(v3, v2)
    if false {
        var v45 float64 = (0.1 - 0.1)
        var v46 uint8 = 183
        _ = v45
        _ = v46
    }
    v47 := GenS5{F0: 232, F1: (^(v28)), F2: (^(v2))}
    v48 := v4["a"]
    v2 += 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v44
    _ = v47
    _ = v48
    _ = v5
}
