package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 bool; F1 uint; F2 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32; F1 uint; F2 int }
type GenS8 struct { F0 int16; F1 string; F2 uint8 }
func genFn0(p0 int16) uint64 {
    v5 := GenS1{F0: 0}
    p0 = (p0 % ((p0) | 1))
    v6 := []int{32768}
    _ = p0
    _ = v5
    _ = v6
    return 21
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]string)
    var v3 int8 = -73
    var v4 int32 = 1
    var v7 uint16 = 2
    v1 = (v1 * v1)
    v4 = (^(v4))
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS2{F0: (v4 << ((255) & 63)), F1: (v3 * 42)}
        var v10 int = int(v7)
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint8 = 1
    if false {
        v4 = 1
        fmt.Println(v1, v3)
    }
    if true {
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS3{F0: 1}
            fmt.Println(v1, fi12)
            _ = fi12
            _ = v13
        }
        switch 9223372036854775807 {
        case 127:
            switch "hello" {
            case "foo bar":
                switch "hello" {
                case "hello":
                    if true {
                        if ("foo bar" < "foo bar") {
                            var v14 uint16 = (v7 + (v7 + (^(v7))))
                            for fi15 := 0; fi15 < 13; fi15++ {
                                var v16 uint32 = (v1 % (((^(v1))) | 1))
                                switch "foo bar" {
                                case "a":
                                    v17 := v2[(v16 - v16)]
                                    _ = v17
                                default:
                                    var v18 int16 = -1
                                    _ = v18
                                }
                                delete(v2, (v1 % ((108) | 1)))
                                _ = fi15
                                _ = v16
                            }
                            v19 := v2[774]
                            _ = v14
                            _ = v19
                        }
                        if (v4 < (v4 - (v4 % ((v4) | 1)))) {
                            var v20 uint16 = (v7 + (v7 << ((42) & 63)))
                            fmt.Println(v7, v11, v3)
                            _ = v20
                        } else {
                            if false {
                                var v21 uint16 = v7
                                v3 -= v3
                                _ = v21
                            } else {
                                v1 = 191
                                v22 := GenS4{F0: 42, F1: 195}
                                if (461 < -769) {
                                    v23 := GenS3{F0: int64(v11)}
                                    _ = v23
                                }
                                _ = v22
                            }
                        }
                    } else {
                        v3 += -1
                    }
                case "foo bar":
                    for fi24 := 0; fi24 < 17; fi24++ {
                        var v25 int = (fi24 - (fi24 * (-(fi24))))
                        v1 -= 849
                        if (false || false) {
                            for fi26 := 0; fi26 < 10; fi26++ {
                                v4 = (v4 - (v4 - v4))
                                fmt.Println(v7)
                                v7 += v7
                                _ = fi26
                            }
                            v27 := GenS3{F0: 684}
                            if ((3.14159 < ((0.1 - (0.1 + (0.0 - ((-2.25 - float32(float64(float32(0.1)))) / ((1e10 - 0.1) - 1.5))))) / 0.1)) || (0 == 255)) {
                                var v28 bool = false
                                if (-457 == (v4 ^ 653)) {
                                    if false {
                                        if ((v3 ^ v3) >= (v3 % ((int8(v25)) | 1))) {
                                            v3 += (v3 << ((uint(v1)) & 63))
                                        }
                                        for fi29 := 0; fi29 < 1; fi29++ {
                                            if v28 {
                                                v25 = 0
                                                var v30 float64 = 1.5
                                                _ = v30
                                            }
                                            _ = fi29
                                        }
                                        for fi31 := 0; fi31 < 13; fi31++ {
                                            if v28 {
                                                if true {
                                                    v32 := GenS5{F0: ((!v28) || ((^(v11)) < (^(v11)))), F1: 65535, F2: fi24}
                                                    v4 += (v4 % (((v4 << ((850) & 63))) | 1))
                                                    v1 = (^(v1))
                                                    _ = v32
                                                }
                                                var v33 uint64 = 641
                                                var v34 int8 = v3
                                                v35 := []int{32768, (fi31 << ((963) & 63))}
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            } else {
                                                v36 := GenS6{F0: 42}
                                                v37 := GenS1{F0: 18446744073709551615}
                                                _ = v36
                                                _ = v37
                                            }
                                            fmt.Println(fi24)
                                            _ = fi31
                                        }
                                    } else {
                                        var v38 uint = 255
                                        switch 4294967295 {
                                        case 435:
                                            v38 -= (v38 & (v38 >> (((v38 / ((v38) | 1))) & 63)))
                                        case 32768:
                                            switch -2147483648 {
                                            case 32767:
                                                var v39 int8 = (v3 & 42)
                                                _ = v39
                                            }
                                        case 77:
                                            fmt.Println(v28, v4, v38)
                                        default:
                                            if v28 {
                                                v40 := v2[(v1 ^ (v1 ^ (v1 & v1)))]
                                                v25 = -9223372036854775808
                                                v38 -= 65535
                                                _ = v40
                                            }
                                        }
                                        _ = v38
                                    }
                                } else {
                                    switch "hello" {
                                    case "":
                                        for fi41 := 0; fi41 < 14; fi41++ {
                                            var v42 bool = (!(!((32767 > 1) && true)))
                                            _ = fi41
                                            _ = v42
                                        }
                                    case "foo bar":
                                        v43 := []int{}
                                        _ = v43
                                    default:
                                        switch ("\n\t" + "foo bar") {
                                        case "a":
                                            var v44 uint = 0
                                            _ = v44
                                        case "\n\t":
                                            _ = len(v2)
                                        }
                                    }
                                    for fi45 := 0; fi45 < 10; fi45++ {
                                        if v28 {
                                            var v46 uint = 517
                                            v47 := GenS5{F0: (v28 || v28), F1: 0, F2: (fi45 << ((uint(v7)) & 63))}
                                            v4 -= (v4 - (-(v4)))
                                            _ = v46
                                            _ = v47
                                        }
                                        _ = fi45
                                    }
                                }
                                var v48 int32 = 0
                                _ = v28
                                _ = v48
                            }
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                    }
                }
            case "hello":
                var v49 uint32 = 534
                _ = v49
            }
        case -122:
            var v50 int = 128
            _ = v50
        case -9223372036854775808:
            var v51 int16 = 1000
            _ = v51
        default:
            for fi52 := 0; fi52 < 5; fi52++ {
                v1 = (v1 ^ (^(v1)))
                for fi53 := 0; fi53 < 14; fi53++ {
                    if true {
                        var v54 uint16 = v7
                        _ = v54
                    }
                    fmt.Println(v1, fi52, fi52)
                    if true {
                        switch 128 {
                        case 443:
                            v55 := []int{fi53, 0, (fi53 - fi53), fi52}
                            _ = v55
                        default:
                            var v56 uint = 2
                            _ = v56
                        }
                        var v57 int8 = 42
                        v58 := []int{fi53}
                        v59 := GenS4{F0: 2, F1: 42}
                        _ = v57
                        _ = v58
                        _ = v59
                    } else {
                        var v60 uint64 = 793
                        _ = v60
                    }
                    _ = fi53
                }
                _ = fi52
            }
        }
        v4 += v4
    }
    v61 := GenS7{F0: 42, F1: 580, F2: 48}
    v7 -= 255
    if false {
        switch 128 {
        case 900:
            v62 := []int{}
            _ = v62
        default:
            for fi63 := 0; fi63 < 15; fi63++ {
                if (genFn0(1) < 466) {
                    v7 = v7
                }
                _ = fi63
            }
        }
        v7 -= v7
    } else {
        fmt.Println(v11)
        for fi64 := 0; fi64 < 8; fi64++ {
            for fi65 := 0; fi65 < 13; fi65++ {
                var v66 bool = ((false && false) && (false || (false && true)))
                var v67 uint32 = 653
                fmt.Println(v4)
                _ = fi65
                _ = v66
                _ = v67
            }
            if true {
                var v68 uint = 320
                v11 -= 42
                _ = v68
            }
            _ = fi64
        }
        v69 := GenS8{F0: -32768, F1: ("\n\t" + ("foo bar" + ("a" + "\n\t"))), F2: 188}
        _ = v69
    }
    v4 = (v4 / ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v7
}
