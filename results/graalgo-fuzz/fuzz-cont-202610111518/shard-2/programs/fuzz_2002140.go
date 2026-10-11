package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
func genFn0(p0 string, p1 uint8, p2 int8) uint16 {
    for fi7 := 0; fi7 < 6; fi7++ {
        if true {
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS1{F0: 2}
                _ = fi8
                _ = v9
            }
            p1 = (^(p1))
            for fi10 := 0; fi10 < 5; fi10++ {
                p0 = ""
                p2 -= (p2 + p2)
                _ = fi10
            }
        } else {
            var v11 int16 = 0
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 uint16 = 0
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 float32 = 0.1
                v16 := GenS1{F0: 1}
                if (!false) {
                    var v17 float64 = -2.25
                    _ = v17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v11
        }
        p2 += p2
        _ = fi7
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    return 910
}
func genFn1(p0 uint64) uint32 {
    for fi18 := 0; fi18 < 10; fi18++ {
        switch "a" {
        case "a":
            if false {
                v19 := []int{879, (fi18 >> ((0) & 63))}
                if false {
                    if ((4294967295 >= 42) || (65535 < 2)) {
                        if true {
                            for fi20 := 0; fi20 < 1; fi20++ {
                                p0 = (p0 * (p0 - (p0 ^ (p0 << ((511) & 63)))))
                                p0 += 0
                                v21 := GenS1{F0: p0}
                                _ = fi20
                                _ = v21
                            }
                            if len(v19) > 0 {
                                v19[0] = fi18
                            }
                        } else {
                            fmt.Println(fi18, p0)
                            p0 -= p0
                            if (true && (false || (992 > 753))) {
                                v22 := GenS1{F0: (^(p0))}
                                _ = v22
                            } else {
                                p0 = 42
                            }
                        }
                        var v23 float32 = 1e10
                        var v24 int16 = -208
                        _ = v23
                        _ = v24
                    } else {
                        p0 -= 2
                        v19 = append(v19, fi18)
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v26 := GenS1{F0: 945}
                            var v27 int = (^(fi25))
                            switch -52 {
                            case 32767:
                                var v28 int16 = -538
                                _ = v28
                            default:
                                p0 = (p0 * p0)
                            }
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                    var v29 float32 = 3.14159
                    p0 -= p0
                    _ = v29
                } else {
                    if (!(false || false)) {
                        fmt.Println(p0, fi18)
                        var v30 int16 = 1
                        v31 := GenS1{F0: 0}
                        _ = v30
                        _ = v31
                    }
                }
                _ = v19
            }
        case "hello":
            for fi32 := 0; fi32 < 1; fi32++ {
                v33 := GenS2{F0: 1, F1: 356, F2: 0}
                p0 -= p0
                var v34 uint8 = 0
                _ = fi32
                _ = v33
                _ = v34
            }
        case "foo bar":
            if false {
                p0 += p0
                var v35 float32 = ((1.5 - 1.5) / (3.14159 / 1.0))
                _ = v35
            }
        default:
            p0 = p0
        }
        _ = fi18
    }
    v36 := GenS2{F0: (p0 % ((18446744073709551615) | 1)), F1: 869, F2: 128}
    v37 := []int{-2147483648, 42, 42, int(p0)}
    _ = p0
    _ = v36
    _ = v37
    return 470
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 127
    var v3 uint = 910
    v4 := []uint32{1, 65535, 247, 20}
    v5 := []int{-1, -9223372036854775808}
    var v6 int8 = 0
    v38 := GenS3{F0: (("" + "\n\t") + ("a" + "")), F1: 1, F2: 42}
    v3 = v3
    v39 := GenS3{F0: ("foo bar" + "foo bar"), F1: 32767, F2: -32768}
    _ = len(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v6, v3, v6)
    if (-2147483648 >= 137) {
        v2 += (v2 - v2)
        if (!true) {
            for fi40 := 0; fi40 < 7; fi40++ {
                if true {
                    fmt.Println(fi40, v2, v6)
                    var v41 int64 = 42
                    _ = v41
                }
                v1[65535] = 32768
                _ = fi40
            }
            v42 := v1[2147483648]
            v3 -= 744
            _ = v42
        } else {
            v6 -= 0
            for fi43 := 0; fi43 < 8; fi43++ {
                var v44 int8 = 66
                switch "foo bar" {
                case "foo bar":
                    switch 32767 {
                    case -362:
                        v45 := GenS3{F0: "\n\t", F1: 42, F2: 0}
                        _ = v45
                    case -2147483648:
                        var v46 int8 = int8(v2)
                        _ = v46
                    case 127:
                        for fi47 := 0; fi47 < 5; fi47++ {
                            v44 = (v6 ^ (v6 & (v6 - v6)))
                            _ = fi47
                        }
                    }
                default:
                    fmt.Println(v44, v3, fi43)
                }
                for fi48 := 0; fi48 < 15; fi48++ {
                    fmt.Println(v6, v2, v6)
                    var v49 uint32 = 489
                    v49 = v49
                    _ = fi48
                    _ = v49
                }
                _ = fi43
                _ = v44
            }
        }
    }
    var v50 uint8 = (v2 * (v2 - 128))
    var v51 int32 = 0
    v6 += int8(v51)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
