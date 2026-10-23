package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint) uint64 {
    p0 = 42
    p1 = (^(p1))
    var v7 uint8 = (^(p0))
    _ = p0
    _ = p1
    _ = v7
    return 972
}
func genFn1(p0 int32, p1 uint32) int8 {
    var v8 float32 = 0.1
    var v9 int8 = 127
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{fi10, (fi10 - (-(fi10))), 32767}
        p1 -= (p1 * (p1 & (p1 ^ (^(p1)))))
        _ = fi10
        _ = v11
    }
    var v12 int8 = 1
    _ = p0
    _ = p1
    _ = v12
    _ = v8
    _ = v9
    return (v12 - v9)
}

func main() {
    var v1 int8 = -66
    var v2 int = 32768
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]uint32)
    var v5 int32 = -2147483648
    var v6 uint64 = 42
    var v13 uint32 = 2147483647
    fmt.Println(v6, v2)
    for fi14 := 0; fi14 < 6; fi14++ {
        if false {
            if (false && false) {
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := []int{2147483647, int(v5)}
                    v2 += (fi15 ^ (v2 + 1))
                    _ = fi15
                    _ = v16
                }
            } else {
                v2 -= 458
                v17 := GenS1{F0: (v5 + v5), F1: (v2 >> ((2) & 63)), F2: v3}
                _ = v17
            }
        } else {
            if false {
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := GenS1{F0: 2147483647, F1: v2, F2: int64(fi18)}
                    var v20 int32 = (v5 | v5)
                    fmt.Println(v3)
                    if (int(v13) <= 42) {
                        switch "foo bar" {
                        case "hello":
                            switch fi18 {
                            case 9223372036854775807:
                                v21 := GenS2{F0: ""}
                                _ = v21
                            case 345:
                                var v22 float64 = 1.5
                                _ = v22
                            case 1:
                                _ = len(v4)
                            default:
                                v23 := GenS3{F0: (^(v6)), F1: "a"}
                                _ = v23
                            }
                        case "foo bar":
                            _ = len(v4)
                        case "\n\t":
                            fmt.Println(v20)
                        default:
                            v24 := []int{127, 128, v2}
                            _ = v24
                        }
                        var v25 uint8 = 26
                        fmt.Println(v6, fi18)
                        _ = v25
                    } else {
                        v4[fi14] = (^(v13))
                        if (false || ((fi14 >> ((4294967295) & 63)) == (fi14 & (v2 & 487)))) {
                            v5 = v5
                            var v26 bool = (!true)
                            if true {
                                v27 := GenS4{F0: v13, F1: v13, F2: 0}
                                _ = v27
                            } else {
                                fmt.Println(v26)
                                var v28 float64 = 1.5
                                v1 += 100
                                _ = v28
                            }
                            _ = v26
                        }
                        switch "" {
                        case "foo bar":
                            v29 := GenS5{F0: 991}
                            _ = v29
                        case "hello":
                            if (fi14 > fi14) {
                                fmt.Println(v3, v2, v20)
                                var v30 int16 = -32768
                                _ = v30
                            } else {
                                v5 = v5
                                v5 += (v20 & int32(fi14))
                            }
                        case "a":
                            v6 = genFn0(139, 202)
                        default:
                            var v31 uint = 2
                            _ = v31
                        }
                        switch "a" {
                        case "\n\t":
                            v6 += v6
                        case "a":
                            v32 := v4[int(v3)]
                            _ = v32
                        }
                        v2 = (v2 << ((819) & 63))
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            v13 -= uint32(v6)
        }
        v2 = fi14
        _ = fi14
    }
    var v33 int64 = 9223372036854775807
    v33 += int64(v5)
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v2)
    default:
        if true {
            if true {
                if ((0.0 >= -2.25) || false) {
                    var v34 uint16 = 1
                    _ = v34
                } else {
                    switch (^(v2)) {
                    case 57:
                        switch v2 {
                        case 42:
                            if ((^(v13)) == (v13 % (((^(v13))) | 1))) {
                                var v35 int64 = v33
                                if true {
                                    var v36 string = "foo bar"
                                    _ = v36
                                } else {
                                    v13 = (v13 / (((^(v13))) | 1))
                                }
                                _ = v35
                            } else {
                                v33 += (v33 / ((0) | 1))
                            }
                        default:
                            var v37 float32 = (((1e10 + 0.1) * float32(0.1)) / 0.1)
                            _ = v37
                        }
                    default:
                        v38 := []int{}
                        _ = v38
                    }
                }
                v2 -= v2
                fmt.Println(v6)
            } else {
                var v39 float64 = 1.5
                var v40 bool = ((v6 >> ((65535) & 63)) > v6)
                _ = v39
                _ = v40
            }
            if (!(true && (112 <= v1))) {
                v13 = (v13 ^ v13)
                for fi41 := 0; fi41 < 12; fi41++ {
                    if false {
                        v42 := []int{42, (v2 | fi41)}
                        if (!(-2.25 <= ((0.0 * 3.14159) * float32(((0.1 / 1.0) / 0.1))))) {
                            var v43 uint64 = (^(v6))
                            _ = v43
                        }
                        _ = v42
                    } else {
                        var v44 bool = false
                        v45 := GenS4{F0: v13, F1: v13, F2: 255}
                        _ = v44
                        _ = v45
                    }
                    var v46 string = "\n\t"
                    _ = fi41
                    _ = v46
                }
                fmt.Println(v2, v6, v33)
            }
        } else {
            v4[-887] = 0
            if true {
                for fi47 := 0; fi47 < 19; fi47++ {
                    fmt.Println(v2)
                    _ = fi47
                }
            } else {
                v48 := GenS1{F0: (v5 / (((-(v5))) | 1)), F1: (v2 % (((v2 / ((v2) | 1))) | 1)), F2: 718}
                switch ("\n\t" + "a") {
                case "hello":
                    var v49 float32 = float32(-2.25)
                    _ = v49
                case "\n\t":
                    v50 := v4[v2]
                    _ = v50
                }
                for fi51 := 0; fi51 < 18; fi51++ {
                    for fi52 := 0; fi52 < 14; fi52++ {
                        var v53 int16 = 32767
                        v54 := GenS5{F0: 32767}
                        v55 := GenS5{F0: 2}
                        _ = fi52
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                    _ = fi51
                }
                _ = v48
            }
            v3 = 0
        }
    }
    if true {
        v2 = 62
        var v56 uint32 = 42
        v57 := GenS2{F0: "\n\t"}
        _ = v56
        _ = v57
    } else {
        var v58 bool = false
        _ = v58
    }
    if false {
        var v59 string = "a"
        v13 = (v13 >> ((1) & 63))
        _ = v59
    } else {
        v60 := GenS4{F0: v13, F1: (v13 / ((730) | 1)), F2: 65535}
        fmt.Println(v5)
        v61 := v4[(-(v2))]
        _ = v60
        _ = v61
    }
    v13 = (^(v13))
    fmt.Println(v6)
    v5 = (v5 << ((18446744073709551615) & 63))
    switch -409 {
    case -9223372036854775808:
        var v62 float64 = 1e10
        _ = v62
    default:
        var v63 uint = 65535
        _ = v63
    }
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
