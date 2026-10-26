package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 int16; F2 int }

func main() {
    var v1 bool = true
    var v2 int16 = 42
    var v3 uint16 = 65535
    var v4 uint8 = 182
    v5 := make(map[uint32]int)
    if (!true) {
        fmt.Println(v2, v1)
        fmt.Println(v1, v1, v4)
        v6 := GenS1{F0: 315}
        _ = v6
    } else {
        v2 = 1000
    }
    v3 += 32767
    for fi7 := 0; fi7 < 4; fi7++ {
        switch (("\n\t" + "hello") + "\n\t") {
        case "a":
            var v8 int8 = -1
            _ = v8
        case "foo bar":
            v1 = (-9223372036854775808 < -1)
        default:
            if (1 >= 42) {
                var v9 string = ("a" + "\n\t")
                fmt.Println(v2)
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 int32 = 1
                    v5[691] = fi10
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                if v1 {
                    switch (-(fi7)) {
                    case -920:
                        if false {
                            v2 = (-(v2))
                        } else {
                            v3 = (v3 | (v3 / ((2) | 1)))
                            switch (fi7 % ((128) | 1)) {
                            case 1:
                                var v12 float64 = (1.5 - 1.5)
                                _ = v12
                            case 2147483647:
                                var v13 uint8 = (v4 - (v4 >> ((1) & 63)))
                                _ = v13
                            case 42:
                                var v14 float64 = (-2.25 + 1e10)
                                _ = v14
                            }
                        }
                    case 32767:
                        for fi15 := 0; fi15 < 12; fi15++ {
                            v2 += -778
                            var v16 int8 = 12
                            v3 = v3
                            var v17 uint16 = 978
                            var v18 uint32 = 862
                            v19 := GenS1{F0: 42}
                            var v20 uint = 63
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    case 1:
                        var v21 int64 = -476
                        _ = v21
                    }
                    if (v1 && (v4 <= (v4 + v4))) {
                        v1 = (!((!(409 <= -309)) && (!v1)))
                        var v22 int8 = 127
                        _ = v22
                    } else {
                        v2 = 42
                    }
                }
                var v23 bool = ((!(("hello" > "hello") || v1)) && (!(42 < 2)))
                if (v1 && (!(0 <= 9223372036854775807))) {
                    fmt.Println(v2, fi7)
                    if (!false) {
                        v3 = (v3 << ((18446744073709551615) & 63))
                        var v24 float64 = -2.25
                        _ = v24
                    } else {
                        v23 = (v23 && ((3.14159 - -2.25) < 1.5))
                    }
                }
                _ = v23
            }
        }
        v4 = (^(v4))
        _ = fi7
    }
    var v25 bool = v1
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int64 = 92
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        delete(v5, uint32(v3))
        _ = fi28
    }
    v25 = false
    for fi29 := 0; fi29 < 7; fi29++ {
        v4 -= (v4 << ((0) & 63))
        _ = fi29
    }
    var v30 int16 = v2
    if ("a" <= "\n\t") {
        for fi31 := 0; fi31 < 2; fi31++ {
            v4 = v4
            fmt.Println(v30)
            if v1 {
                v2 += (v2 + v2)
                var v32 uint8 = v4
                switch "hello" {
                case "foo bar":
                    fmt.Println(v32)
                case "a":
                    v33 := GenS2{F0: 1}
                    _ = v33
                default:
                    switch "\n\t" {
                    case "":
                        var v34 uint8 = (v4 ^ v4)
                        _ = v34
                    case "foo bar":
                        v32 = (v4 - v4)
                    }
                }
                _ = v32
            }
            _ = fi31
        }
        switch ("" + ("foo bar" + "foo bar")) {
        case "\n\t":
            var v35 int = 127
            _ = v35
        case "hello":
            if (-128 <= -73) {
                v36 := GenS3{F0: (v3 <= v3), F1: -2147483648, F2: 65535}
                v4 -= 138
                _ = v36
            } else {
                switch -9223372036854775808 {
                case 42:
                    v30 = v30
                case -1:
                    for fi37 := 0; fi37 < 18; fi37++ {
                        var v38 int = fi37
                        v38 = (fi37 / (((v38 + -2147483648)) | 1))
                        v39 := []int{2147483647, fi37, 127, (v38 & (v38 << ((65535) & 63)))}
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                case -600:
                    var v40 float32 = 1.5
                    _ = v40
                }
                v41 := GenS4{F0: (v30 * v30), F1: 798}
                fmt.Println(v25, v4, v2)
                _ = v41
            }
        }
        var v42 int16 = -896
        _ = v42
    }
    v43 := GenS5{F0: 1}
    v3 -= (v3 * (v3 + uint16(v4)))
    _ = len(v5)
    v44 := GenS1{F0: 18446744073709551615}
    if (v25 && (71 != -128)) {
        for fi45 := 0; fi45 < 8; fi45++ {
            var v46 uint64 = 959
            if false {
                _ = len(v5)
            }
            _ = fi45
            _ = v46
        }
        switch "\n\t" {
        case "a":
            if true {
                if v25 {
                    v1 = v25
                    delete(v5, 4294967295)
                    v2 = 405
                }
            } else {
                v1 = false
                delete(v5, 839)
                if (2 == 2147483648) {
                    var v47 int32 = 130
                    _ = v47
                }
            }
        }
    } else {
        v4 += v4
    }
    if ((v2 * (v2 >> ((512) & 63))) <= (v2 - v30)) {
        if false {
            for fi48 := 0; fi48 < 16; fi48++ {
                if (!v1) {
                    if v25 {
                        v49 := GenS3{F0: (42 > v30), F1: (-(fi48)), F2: 2}
                        var v50 string = "hello"
                        v2 = v30
                        _ = v49
                        _ = v50
                    }
                }
                for fi51 := 0; fi51 < 3; fi51++ {
                    v52 := v5[601]
                    v1 = (!(v30 < v30))
                    _ = fi51
                    _ = v52
                }
                _ = fi48
            }
            if (127 < 0) {
                for fi53 := 0; fi53 < 14; fi53++ {
                    if v1 {
                        switch (fi53 * fi53) {
                        case 2147483647:
                            var v54 int = fi53
                            _ = v54
                        case -51:
                            switch fi53 {
                            case -1:
                                switch -9223372036854775808 {
                                case -591:
                                    if v1 {
                                        v55 := []int{fi53}
                                        v56 := GenS3{F0: (42 != 0), F1: (fi53 + -2147483648), F2: 1}
                                        _ = v55
                                        _ = v56
                                    } else {
                                        v57 := GenS6{F0: 1, F1: 1, F2: fi53}
                                        _ = v57
                                    }
                                }
                            case 9223372036854775807:
                                v30 += 42
                            case 0:
                                v25 = v1
                            }
                        case -2147483648:
                            fmt.Println(v4)
                        default:
                            v4 = (v4 + 37)
                        }
                        switch ("a" + "hello") {
                        case "foo bar":
                            v58 := v5[65535]
                            _ = v58
                        default:
                            for fi59 := 0; fi59 < 15; fi59++ {
                                var v60 uint = 0
                                _ = fi59
                                _ = v60
                            }
                        }
                        fmt.Println(v1, v2)
                    } else {
                        switch (fi53 << ((944) & 63)) {
                        case 4294967295:
                            var v61 uint16 = v3
                            _ = v61
                        case 0:
                            var v62 uint32 = 2147483648
                            _ = v62
                        }
                        _ = len(v5)
                        var v63 int8 = 0
                        _ = v63
                    }
                    _ = fi53
                }
                for fi64 := 0; fi64 < 10; fi64++ {
                    v4 = (v4 >> ((18446744073709551615) & 63))
                    switch fi64 {
                    case 32767:
                        fmt.Println(v25)
                    case 128:
                        for fi65 := 0; fi65 < 19; fi65++ {
                            var v66 float32 = float32((-2.25 * ((3.14159 / 0.1) - (1e10 / 0.1))))
                            v67 := GenS1{F0: 42}
                            var v68 int64 = 0
                            v69 := v5[2147483648]
                            var v70 int32 = 278
                            var v71 int = (fi64 << ((2) & 63))
                            _ = fi65
                            _ = v66
                            _ = v67
                            _ = v68
                            _ = v69
                            _ = v70
                            _ = v71
                        }
                    default:
                        var v72 uint64 = 276
                        _ = v72
                    }
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        _ = len(v5)
                    case "foo bar":
                        v25 = false
                    default:
                        v73 := []int{(fi64 - 0), fi64, fi64, (-(fi64))}
                        _ = v73
                    }
                    _ = fi64
                }
            } else {
                for fi74 := 0; fi74 < 5; fi74++ {
                    v3 = uint16(v4)
                    _ = fi74
                }
            }
        } else {
            var v75 int16 = 663
            _ = len(v5)
            var v76 uint16 = (v3 | (^(v3)))
            _ = v75
            _ = v76
        }
        v1 = true
        var v77 int64 = 0
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
