package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint }

func main() {
    var v1 bool = (1 > 128)
    var v2 uint64 = 1
    var v3 uint = 515
    if v1 {
        if (1.5 != 1.5) {
            v1 = v1
            switch "hello" {
            case "a":
                v4 := []int{32767, -24, -9223372036854775808, 9223372036854775807}
                _ = v4
            default:
                var v5 float64 = 1e10
                _ = v5
            }
        } else {
            var v6 int64 = 1
            _ = v6
        }
        v1 = (65535 >= 2147483648)
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 int = (-(fi7))
                v1 = false
                var v10 int32 = -951
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    } else {
        v11 := []int{4294967295}
        var v12 string = "\n\t"
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            v3 = (v3 << ((v3) & 63))
            _ = fi14
        }
        _ = fi13
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            v3 = (v3 ^ v3)
            _ = fi16
        }
        _ = fi15
    }
    if (1 >= 1) {
        v3 = 617
        var v17 int64 = -1
        v3 += (v3 % ((v3) | 1))
        _ = v17
    } else {
        var v18 int = 42
        _ = v18
    }
    var v19 int64 = -1
    v1 = (128 <= int(v19))
    if false {
        fmt.Println(v1)
        switch "\n\t" {
        case "hello":
            fmt.Println(v3, v3, v19)
        }
    }
    v20 := GenS1{F0: 0, F1: -994, F2: v2}
    if (false || v1) {
        var v21 uint = (v3 - v3)
        _ = v21
    } else {
        if v1 {
            if (32767 >= 478) {
                var v22 int64 = (v19 - (v19 & 1))
                for fi23 := 0; fi23 < 20; fi23++ {
                    fmt.Println(v1, fi23)
                    _ = fi23
                }
                var v24 uint = (v3 - (^(v3)))
                _ = v22
                _ = v24
            } else {
                fmt.Println(v1, v2)
                v25 := GenS1{F0: -836, F1: int32(v19), F2: (v2 / ((943) | 1))}
                var v26 string = "hello"
                _ = v25
                _ = v26
            }
            v27 := []int{-2147483648, -1}
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            _ = v27
        } else {
            v19 = 898
        }
    }
    if v1 {
        v28 := GenS2{F0: v3, F1: v19}
        v1 = (float32((1e10 * 1e10)) != -2.25)
        if (252 != (v19 >> ((18446744073709551615) & 63))) {
            if v1 {
                if ((1e10 * float32((3.14159 + 1.5))) == (0.1 * (float32((0.1 - 0.1)) - (float32((0.1 - 1e10)) + (3.14159 * 1e10))))) {
                    fmt.Println(v1, v1)
                    if false {
                        var v29 int8 = 1
                        v1 = true
                        v30 := GenS2{F0: (v3 << (((v3 + uint(v2))) & 63)), F1: (^(v19))}
                        _ = v29
                        _ = v30
                    }
                } else {
                    if true {
                        for fi31 := 0; fi31 < 15; fi31++ {
                            for fi32 := 0; fi32 < 14; fi32++ {
                                v33 := []int{}
                                _ = fi32
                                _ = v33
                            }
                            switch "foo bar" {
                            case "":
                                v1 = (false || v1)
                            case "a":
                                var v34 int16 = -1
                                _ = v34
                            case "hello":
                                var v35 int32 = -631
                                _ = v35
                            }
                            _ = fi31
                        }
                    }
                    v3 = (v3 >> (((v3 & v3)) & 63))
                    var v36 int8 = 100
                    _ = v36
                }
                for fi37 := 0; fi37 < 17; fi37++ {
                    for fi38 := 0; fi38 < 20; fi38++ {
                        switch -2147483648 {
                        case -9223372036854775808:
                            if (true && v1) {
                                v1 = ((-127 >= 127) || v1)
                                var v39 int = 128
                                var v40 uint16 = 280
                                _ = v39
                                _ = v40
                            } else {
                                fmt.Println(v1, v2, v3)
                                v41 := GenS2{F0: v3, F1: (^(v19))}
                                _ = v41
                            }
                        case 1:
                            var v42 uint64 = (^(v2))
                            _ = v42
                        case -1:
                            v1 = ((!(!(-2147483648 == 0))) || (42 == -1))
                        default:
                            if ((3.14159 + 0.0) >= float64((-2.25 + 0.0))) {
                                v43 := GenS1{F0: 1, F1: -2147483648, F2: (v2 % ((v2) | 1))}
                                v3 = (v3 / ((4294967295) | 1))
                                _ = v43
                            } else {
                                switch 42 {
                                case 9223372036854775807:
                                    var v44 float64 = 3.14159
                                    _ = v44
                                case 128:
                                    v45 := GenS2{F0: 885, F1: (-(v19))}
                                    _ = v45
                                case 1:
                                    v3 -= v3
                                }
                                v46 := []int{(^(fi37)), -1}
                                _ = v46
                            }
                        }
                        _ = fi38
                    }
                    v47 := GenS1{F0: -393, F1: 42, F2: v2}
                    v1 = ((199 <= 1) || true)
                    _ = fi37
                    _ = v47
                }
                v48 := []int{32767, 9223372036854775807, 128}
                _ = v48
            }
        } else {
            v49 := GenS3{F0: v3}
            _ = v49
        }
        _ = v28
    }
    if v1 {
        v50 := []int{428, 288}
        var v51 float32 = 1e10
        _ = v50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
