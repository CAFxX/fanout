package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int16 }
func genFn0(p0 uint, p1 int32) uint16 {
    v6 := []int{32767, 32767, -77}
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    _ = p0
    _ = p1
    _ = v6
    return 32767
}
func genFn1(p0 int64, p1 uint) uint32 {
    p1 = (^(p1))
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint16 = 42
        var v9 uint32 = 517
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 10; fi10++ {
            p0 = (p0 ^ p0)
            var v11 bool = (0 == 2)
            if ((p0 << ((p1) & 63)) > p0) {
                v12 := []int{(fi10 >> ((65535) & 63)), (fi10 ^ 512)}
                _ = v12
            } else {
                if true {
                    p0 += p0
                }
                for fi13 := 0; fi13 < 1; fi13++ {
                    p0 -= p0
                    v14 := []int{128, fi10, (fi10 << (((p1 / ((42) | 1))) & 63))}
                    _ = len(v14)
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi10
            _ = v11
        }
        if (32767 > -9223372036854775808) {
            var v15 int = 0
            _ = v15
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 bool = (!(!true))
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 int64 = 376
    var v5 uint32 = 877
    if v1 {
        if v1 {
            v5 -= 2
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 float64 = 1e10
                for fi18 := 0; fi18 < 5; fi18++ {
                    if (v1 || true) {
                        fmt.Println(fi18, v1)
                    }
                    if (!(!(v4 != v4))) {
                        v17 = ((float64(1e10) / 1.0) - v17)
                        var v19 string = "a"
                        _ = v19
                    }
                    v20 := []int{}
                    _ = fi18
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
        } else {
            v4 = v4
            var v21 int64 = v4
            for fi22 := 0; fi22 < 7; fi22++ {
                v23 := []int{fi22, fi22, fi22, (fi22 + (fi22 % ((fi22) | 1)))}
                switch fi22 {
                case 32768:
                    for fi24 := 0; fi24 < 18; fi24++ {
                        if (127 >= 1) {
                            v5 = (v5 / ((v5) | 1))
                            var v25 float32 = -2.25
                            var v26 int8 = -128
                            _ = v25
                            _ = v26
                        } else {
                            if v1 {
                                delete(v2, "a")
                            }
                        }
                        _ = fi24
                    }
                case 2147483647:
                    v27 := GenS1{F0: -38}
                    _ = v27
                default:
                    switch "\n\t" {
                    case "":
                        v2[(("" + "") + ("a" + ""))] = (-(fi22))
                    }
                }
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 8; fi29++ {
                v1 = v1
                var v30 uint = 694
                v31 := GenS1{F0: -114}
                v5 = (v5 + v5)
                _ = fi29
                _ = v30
                _ = v31
            }
            fmt.Println(v1)
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 uint64 = 1
                if v1 {
                    var v34 uint16 = uint16(v33)
                    _ = v34
                } else {
                    if (true && v1) {
                        v35 := GenS1{F0: -90}
                        v36 := []int{}
                        _ = v35
                        _ = v36
                    }
                }
                _ = fi32
                _ = v33
            }
            _ = fi28
        }
    }
    if (v1 && (v1 || false)) {
        for fi37 := 0; fi37 < 9; fi37++ {
            v1 = false
            v38 := GenS2{F0: 0, F1: int16(fi37)}
            if v1 {
                fmt.Println(fi37, v4, v1)
            } else {
                v39 := []int{(fi37 - fi37), (fi37 + fi37), fi37, (fi37 * fi37)}
                _ = v39
            }
            _ = fi37
            _ = v38
        }
        v4 = -9223372036854775808
        if false {
            var v40 uint16 = 1
            v40 -= v40
            var v41 uint64 = 4294967295
            _ = v40
            _ = v41
        } else {
            v4 -= (v4 & -1)
            v42 := GenS2{F0: 128, F1: 32767}
            _ = v42
        }
    }
    fmt.Println(v5, v4)
    if ((v4 - -493) == 790) {
        fmt.Println(v4)
    }
    v43 := GenS1{F0: -1}
    v44 := []int{-221, -1, -981, 32767}
    v45 := GenS1{F0: -1}
    if len(v44) > 0 {
        v44[0] = -1
    }
    switch -1 {
    case 9223372036854775807:
        if len(v44) >= 2 {
            v44 = v44[1:len(v44)]
        }
    case 0:
        var v46 uint16 = 42
        _ = v46
    case -2147483648:
        switch ("\n\t" + (("" + "\n\t") + (("\n\t" + "foo bar") + "\n\t"))) {
        case "foo bar":
            switch "" {
            case "hello":
                if (!(-44 != 0)) {
                    fmt.Println(v5, v5, v4)
                } else {
                    fmt.Println(v1, v5, v4)
                }
            case "a":
                for fi47 := 0; fi47 < 8; fi47++ {
                    v48 := GenS2{F0: 128, F1: 494}
                    if false {
                        v44 = append(v44, -52)
                        for fi49 := 0; fi49 < 13; fi49++ {
                            var v50 uint8 = 1
                            _ = fi49
                            _ = v50
                        }
                        for fi51 := 0; fi51 < 8; fi51++ {
                            if (1 == 1) {
                                var v52 uint16 = 508
                                v5 += v5
                                var v53 uint = 1
                                _ = v52
                                _ = v53
                            }
                            var v54 int = fi47
                            if v1 {
                                v55 := GenS2{F0: 128, F1: 1000}
                                v56 := GenS3{F0: 1}
                                _ = v55
                                _ = v56
                            }
                            _ = fi51
                            _ = v54
                        }
                    }
                    _ = fi47
                    _ = v48
                }
            case "":
                v5 -= (v5 % ((42) | 1))
            default:
                v57 := GenS4{F0: 1, F1: 0, F2: 639}
                _ = v57
            }
        case "a":
            v1 = (0 < 42)
        case "":
            if (((-32768 > 349) && v1) || true) {
                v5 -= 0
            }
        default:
            var v58 uint64 = 797
            _ = v58
        }
    default:
        fmt.Println(v1, v5, v5)
    }
    v4 += (v4 ^ v4)
    if (2147483648 != 0) {
        v5 -= genFn1(v4, 18446744073709551615)
        v4 -= 173
        v44 = append(v44, 750)
    } else {
        v44 = append(v44, -597)
        switch -32 {
        case 2147483647:
            if (-2.25 < ((1.5 * -2.25) + 0.1)) {
                v4 += v4
            } else {
                fmt.Println(v1)
            }
        default:
            v4 = (v4 * 0)
        }
    }
    v59 := GenS1{F0: 1}
    if (v1 || v1) {
        var v60 int16 = -32768
        fmt.Println(v4)
        _ = v60
    } else {
        fmt.Println(v1, v5)
    }
    if v1 {
        v5 -= (v5 ^ (v5 ^ 4294967295))
        v4 += (-(v4))
        var v61 bool = (819 == 2)
        _ = v61
    }
    v44 = append(v44, 32768)
    for fi62 := 0; fi62 < 7; fi62++ {
        delete(v3, fi62)
        v4 -= -9223372036854775808
        var v63 uint16 = 255
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v59
}
