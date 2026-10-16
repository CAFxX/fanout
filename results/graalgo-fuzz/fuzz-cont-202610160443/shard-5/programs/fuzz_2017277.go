package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 uint; F1 int8; F2 int16 }
type GenS8 struct { F0 int16 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float32 = 0.0
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint8 = 127
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 42}
        _ = fi6
        _ = v10
        _ = v7
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            switch (fi11 << ((374) & 63)) {
            case 4294967295:
                if true {
                    v13 := []int{fi11}
                    var v14 int = fi11
                    _ = v13
                    _ = v14
                }
            case -344:
                v15 := GenS2{F0: 2}
                _ = v15
            default:
                if (1 < 1) {
                    var v16 uint64 = 4294967295
                    _ = v16
                } else {
                    if false {
                        var v17 int8 = -128
                        var v18 int8 = (v17 << ((625) & 63))
                        _ = v17
                        _ = v18
                    }
                }
            }
            _ = fi12
        }
        var v19 float32 = 0.1
        _ = fi11
        _ = v19
    }
    var v20 int16 = -417
    if (!false) {
        v20 = (v20 >> ((255) & 63))
    }
    if (true && (("foo bar" + "a") != "hello")) {
        v21 := GenS1{F0: 1}
        v22 := GenS2{F0: 1}
        _ = v21
        _ = v22
    }
    _ = v20
    return 63
}

func main() {
    var v1 uint8 = 1
    var v2 uint32 = 708
    var v3 int32 = 2147483647
    v4 := []int{-591}
    var v5 uint64 = 0
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS3{F0: 255, F1: 17, F2: -1}
        if true {
            fmt.Println(v1, v3, fi23)
        }
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 string = ""
        _ = len(v4)
        for fi27 := 0; fi27 < 8; fi27++ {
            if ((v5 | 2) > 18446744073709551615) {
                v2 += v2
            } else {
                var v28 int16 = -32768
                _ = v28
            }
            var v29 uint64 = (v5 >> ((255) & 63))
            v30 := GenS4{F0: ("foo bar" + (v26 + v26))}
            switch (fi27 * 42) {
            case -2147483648:
                var v31 bool = true
                _ = v31
            default:
                v29 = 2
                v32 := GenS5{F0: v26}
                if (!(255 <= 42)) {
                    var v33 int = (fi27 ^ (fi27 / ((fi27) | 1)))
                    _ = v33
                }
                _ = v32
            }
            _ = fi27
            _ = v29
            _ = v30
        }
        _ = fi25
        _ = v26
    }
    v4 = append(v4, -9223372036854775808)
    var v34 int16 = -383
    var v35 int64 = -9223372036854775808
    v34 += 32767
    switch "a" {
    case "foo bar":
        fmt.Println(v34, v3)
    case "\n\t":
        fmt.Println(v1)
    default:
        for fi36 := 0; fi36 < 8; fi36++ {
            if false {
                fmt.Println(v5, v34)
                v1 = uint8(v5)
                var v37 bool = true
                _ = v37
            }
            v38 := GenS6{F0: (v35 >> ((0) & 63)), F1: genFn0()}
            _ = fi36
            _ = v38
        }
    }
    v1 -= v1
    if len(v4) > 0 {
        v4[0] = 42
    }
    switch "hello" {
    case "hello":
        for fi39 := 0; fi39 < 13; fi39++ {
            if true {
                switch 0 {
                case 32767:
                    v40 := GenS3{F0: 1, F1: 2, F2: (v34 << ((255) & 63))}
                    _ = v40
                case -133:
                    v41 := GenS6{F0: 1, F1: 1}
                    _ = v41
                case 818:
                    var v42 uint = 0
                    _ = v42
                }
                if false {
                    v5 = (v5 | 894)
                    v4 = append(v4, 1)
                }
            }
            _ = fi39
        }
    case "a":
        if true {
            var v43 uint = 137
            var v44 int64 = (v35 % ((v35) | 1))
            _ = v43
            _ = v44
        } else {
            switch ("\n\t" + "") {
            case "a":
                _ = len(v4)
            case "":
                for fi45 := 0; fi45 < 1; fi45++ {
                    for fi46 := 0; fi46 < 17; fi46++ {
                        switch 127 {
                        case 32767:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 42:
                            var v47 int8 = 0
                            _ = v47
                        default:
                            v48 := GenS7{F0: 0, F1: 0, F2: (v34 % ((1000) | 1))}
                            _ = v48
                        }
                        if (-1 > -113) {
                            v3 = (v3 + (v3 % ((v3) | 1)))
                            switch fi46 {
                            case 32767:
                                v5 -= (v5 * 802)
                            case -1:
                                v1 = 2
                            }
                            _ = len(v4)
                        } else {
                            v3 -= v3
                        }
                        v49 := GenS3{F0: 1, F1: 127, F2: -339}
                        _ = fi46
                        _ = v49
                    }
                    var v50 uint8 = (v1 << ((65535) & 63))
                    _ = fi45
                    _ = v50
                }
            case "\n\t":
                var v51 int16 = v34
                _ = v51
            }
            if false {
                for fi52 := 0; fi52 < 16; fi52++ {
                    if (false || false) {
                        if (v2 >= v2) {
                            v53 := GenS2{F0: 2}
                            _ = v53
                        } else {
                            v54 := GenS8{F0: (v34 * v34)}
                            _ = v54
                        }
                        v34 += -467
                    }
                    var v55 uint64 = (v5 + uint64(v1))
                    for fi56 := 0; fi56 < 5; fi56++ {
                        v3 -= (v3 >> ((4294967295) & 63))
                        _ = fi56
                    }
                    _ = fi52
                    _ = v55
                }
                switch -9223372036854775808 {
                case 32767:
                    switch "\n\t" {
                    case "\n\t":
                        switch 2147483647 {
                        case 972:
                            v3 = 42
                        default:
                            for fi57 := 0; fi57 < 9; fi57++ {
                                for fi58 := 0; fi58 < 11; fi58++ {
                                    var v59 uint32 = (v2 & (^(v2)))
                                    _ = fi58
                                    _ = v59
                                }
                                _ = fi57
                            }
                        }
                    case "":
                        v60 := GenS6{F0: (v35 + -727), F1: 1}
                        _ = v60
                    }
                case 712:
                    v34 += (-(v34))
                case -1:
                    if len(v4) > 0 {
                        v4[0] = -778
                    }
                }
                switch 0 {
                case 128:
                    var v61 int8 = 0
                    _ = v61
                case 9223372036854775807:
                    fmt.Println(v3)
                case 4294967295:
                    var v62 int8 = int8(v35)
                    _ = v62
                }
            } else {
                fmt.Println(v35, v34, v1)
                v3 = v3
            }
        }
    }
    if true {
        v4 = append(v4, 2147483647)
        switch "foo bar" {
        case "foo bar":
            for fi63 := 0; fi63 < 18; fi63++ {
                if (true || ((!true) && true)) {
                    v4 = append(v4, fi63)
                } else {
                    _ = len(v4)
                }
                var v64 int32 = v3
                _ = fi63
                _ = v64
            }
        case "\n\t":
            for fi65 := 0; fi65 < 16; fi65++ {
                v34 += v34
                var v66 float64 = 0.1
                _ = fi65
                _ = v66
            }
        }
    } else {
        _ = len(v4)
    }
    for fi67 := 0; fi67 < 5; fi67++ {
        var v68 int64 = v35
        v5 = 365
        _ = fi67
        _ = v68
    }
    var v69 int8 = -1
    var v70 int = 0
    var v71 float64 = ((1e10 + 0.0) + (float64(1.5) * 1e10))
    v4 = append(v4, -2147483648)
    var v72 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
