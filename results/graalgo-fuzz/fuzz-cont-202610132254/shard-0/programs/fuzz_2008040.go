package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 string; F2 int8 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 uint64; F2 int64 }

func main() {
    var v1 uint16 = 65535
    var v2 int64 = -1
    var v3 float32 = (float32(-2.25) * (0.1 + 0.0))
    v4 := make(map[uint32]uint32)
    var v5 uint16 = (^(v1))
    if ("" <= (("foo bar" + "foo bar") + (((("" + "") + "\n\t") + "") + "a"))) {
        v5 -= v1
    } else {
        if (123 == 243) {
            var v6 uint16 = 692
            if (475 < 0) {
                v7 := []int{-9223372036854775808, 9223372036854775807, 128}
                if true {
                    v8 := GenS1{F0: 249}
                    if true {
                        v9 := GenS2{F0: (v6 | v1), F1: 127, F2: 128}
                        var v10 uint = 2
                        _ = v10
                        _ = v9
                    } else {
                        var v11 string = "hello"
                        v11 = "hello"
                        _ = v11
                    }
                    _ = v8
                }
                fmt.Println(v6)
                _ = v7
            } else {
                v2 += (v2 >> ((369) & 63))
                var v12 bool = true
                fmt.Println(v12, v1, v6)
                _ = v12
            }
            _ = v6
        } else {
            if (255 <= 591) {
                var v13 float64 = 0.0
                switch 4294967295 {
                case -1:
                    v14 := GenS3{F0: 122}
                    _ = v14
                case 32768:
                    if false {
                        v13 += (v13 / v13)
                        for fi15 := 0; fi15 < 10; fi15++ {
                            for fi16 := 0; fi16 < 9; fi16++ {
                                v1 = (v1 ^ 1)
                                if (!(42 != 0)) {
                                    v17 := GenS4{F0: 2, F1: "hello", F2: -1}
                                    _ = v17
                                }
                                _ = fi16
                            }
                            fmt.Println(v3, v1, fi15)
                            var v18 uint8 = 2
                            _ = fi15
                            _ = v18
                        }
                    }
                }
                _ = v13
            }
            fmt.Println(v2)
            v19 := v4[2]
            _ = v19
        }
    }
    if true {
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v20 int8 = 127
            _ = v20
        default:
            if (2 < 2) {
                var v21 int16 = 1000
                var v22 int16 = 0
                delete(v4, 4294967295)
                _ = v21
                _ = v22
            }
        }
        v23 := GenS3{F0: 1}
        v24 := v4[4294967295]
        _ = v23
        _ = v24
    } else {
        fmt.Println(v1, v3, v5)
        _ = len(v4)
    }
    switch 32767 {
    case 9223372036854775807:
        var v25 uint8 = 1
        _ = v25
    case 127:
        fmt.Println(v2, v3)
    case -226:
        fmt.Println(v2, v5, v1)
    default:
        v26 := GenS4{F0: 4294967295, F1: "hello", F2: 42}
        _ = v26
    }
    v5 += 42
    var v27 int8 = -85
    var v28 int32 = 1
    if true {
        v2 += -1
        v4[2147483648] = 978
    } else {
        if (-9223372036854775808 < -9223372036854775808) {
            for fi29 := 0; fi29 < 12; fi29++ {
                v30 := GenS3{F0: 2}
                _ = fi29
                _ = v30
            }
        } else {
            v5 = (v5 >> ((998) & 63))
        }
        for fi31 := 0; fi31 < 2; fi31++ {
            switch 0 {
            case 648:
                var v32 int64 = v2
                _ = v32
            case 9223372036854775807:
                v5 = v5
            }
            var v33 string = "\n\t"
            _ = fi31
            _ = v33
        }
        for fi34 := 0; fi34 < 5; fi34++ {
            var v35 float64 = 3.14159
            _ = fi34
            _ = v35
        }
    }
    v1 = 598
    for fi36 := 0; fi36 < 6; fi36++ {
        v37 := GenS5{F0: 18446744073709551615, F1: 42}
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        if false {
            var v39 uint16 = v1
            v1 = (v1 << ((18446744073709551615) & 63))
            fmt.Println(v2)
            _ = v39
        }
        _ = fi38
    }
    if true {
        switch -9223372036854775808 {
        case 127:
            switch 9223372036854775807 {
            case 481:
                if (false && false) {
                    switch -591 {
                    case 32768:
                        if ("" >= "\n\t") {
                            for fi40 := 0; fi40 < 1; fi40++ {
                                fmt.Println(fi40, v5, v27)
                                v41 := GenS2{F0: 0, F1: fi40, F2: int(v28)}
                                _ = fi40
                                _ = v41
                            }
                        }
                    case 0:
                        v27 += int8(v5)
                    case 2147483647:
                        fmt.Println(v2, v3, v28)
                    }
                    var v42 int8 = (v27 << ((18446744073709551615) & 63))
                    _ = v42
                }
            case -373:
                var v43 int64 = v2
                _ = v43
            case 32767:
                v27 = v27
            }
        case 128:
            var v44 uint32 = 65535
            _ = v44
        case 4294967295:
            var v45 uint64 = 530
            _ = v45
        default:
            v28 = v28
        }
        fmt.Println(v3, v27, v27)
    }
    if true {
        if true {
            var v46 int8 = (v27 * 127)
            _ = v46
        }
        for fi47 := 0; fi47 < 9; fi47++ {
            v48 := []int{}
            if (false || true) {
                v28 += (v28 * (v28 * v28))
                var v49 int32 = 2147483647
                _ = v49
            }
            var v50 uint8 = 85
            _ = fi47
            _ = v48
            _ = v50
        }
        v51 := GenS6{F0: int16(v5)}
        _ = v51
    } else {
        if (false || (!(-873 < -32768))) {
            for fi52 := 0; fi52 < 11; fi52++ {
                for fi53 := 0; fi53 < 12; fi53++ {
                    v28 = 493
                    var v54 string = "\n\t"
                    v55 := GenS7{F0: (fi53 * (fi53 / (((-(fi52))) | 1))), F1: 1, F2: v2}
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                _ = fi52
            }
        }
        for fi56 := 0; fi56 < 15; fi56++ {
            if true {
                var v57 uint64 = 2
                v1 -= v5
                _ = v57
            } else {
                for fi58 := 0; fi58 < 3; fi58++ {
                    v59 := GenS1{F0: 27}
                    v60 := GenS3{F0: 1}
                    v61 := v4[42]
                    if true {
                        switch (fi58 - fi58) {
                        case -664:
                            fmt.Println(v61, fi58)
                        }
                        v61 = (^(v61))
                        v62 := v4[506]
                        _ = v62
                    } else {
                        var v63 float32 = -2.25
                        _ = v63
                    }
                    _ = fi58
                    _ = v59
                    _ = v60
                    _ = v61
                }
                var v64 float64 = 1.5
                fmt.Println(v28, v2)
                _ = v64
            }
            _ = fi56
        }
        var v65 bool = (2 > 1)
        _ = v65
    }
    var v66 int8 = v27
    v27 += (v66 ^ v27)
    var v67 uint32 = 67
    var v68 uint64 = 2
    var v69 bool = false
    var v70 uint32 = v67
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
