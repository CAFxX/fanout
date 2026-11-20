package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64 }
func genFn0() bool {
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := []int{42, (fi6 ^ fi6)}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if false {
            var v8 int16 = -32768
            var v9 float32 = ((0.1 - 0.1) * (1.5 - (0.1 - 0.1)))
            v9 -= (v9 / (float32(3.14159) / 1e10))
            _ = v8
            _ = v9
        } else {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v10 uint32 = 2147483647
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    var v11 uint = 0
    if true {
        v12 := []int{42}
        _ = v12
    } else {
        v13 := GenS1{F0: 595}
        v11 = v11
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(v11, v11)
            _ = fi14
        }
        _ = v13
    }
    v15 := []int{353, -1, -1, 128}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = v11
    _ = v15
    return (!true)
}

func main() {
    var v1 float64 = 1.5
    var v2 int = 4294967295
    var v3 int64 = 9223372036854775807
    var v4 string = ((((("hello" + "") + "") + "a") + ("" + "foo bar")) + "a")
    var v5 string = "\n\t"
    v16 := []int{-2147483648, v2}
    fmt.Println(v3, v5, v3)
    v1 = float64(1.5)
    var v17 string = ((v5 + v5) + (((((v5 + v5) + v4) + v4) + ((v5 + "") + "\n\t")) + ((("a" + ("a" + "a")) + v4) + v4)))
    var v18 uint32 = 2
    var v19 int16 = 686
    fmt.Println(v3, v18, v19)
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    _ = len(v16)
    var v20 uint64 = uint64(v18)
    v16 = append(v16, -380)
    for fi21 := 0; fi21 < 7; fi21++ {
        v16 = append(v16, int(v20))
        var v22 uint8 = 95
        var v23 uint8 = v22
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: ("\n\t" + "hello"), F1: 42}
    if true {
        if true {
            v3 = v3
            v4 = ((v5 + v17) + v4)
        } else {
            v17 = "a"
        }
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    } else {
        var v25 bool = (-2.25 != 1.5)
        if v25 {
            if v25 {
                switch v17 {
                case "a":
                    var v26 bool = (true && (!(493 > (v20 / ((v20) | 1)))))
                    _ = v26
                case "hello":
                    if true {
                        switch 2147483647 {
                        case 1:
                            fmt.Println(v18, v4, v20)
                        }
                        var v27 uint32 = 2147483648
                        for fi28 := 0; fi28 < 18; fi28++ {
                            var v29 int = (fi28 * v2)
                            fmt.Println(v1, v25)
                            if v25 {
                                var v30 bool = v25
                                fmt.Println(v19, v5, v17)
                                fmt.Println(v29, v2)
                                _ = v30
                            } else {
                                v20 = (v20 >> ((4294967295) & 63))
                                var v31 int64 = v3
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    fmt.Println(v2, v25, v19)
                                    fmt.Println(v4, v4, v29)
                                    _ = fi32
                                }
                                _ = v31
                            }
                            _ = fi28
                            _ = v29
                        }
                        _ = v27
                    }
                case "\n\t":
                    for fi33 := 0; fi33 < 8; fi33++ {
                        switch "foo bar" {
                        case "hello":
                            if (v25 && v25) {
                                v34 := GenS3{F0: 1, F1: 143, F2: 42}
                                switch 1 {
                                case 564:
                                    if len(v16) > 0 {
                                        v16[0] = (v2 + fi33)
                                    }
                                case 9223372036854775807:
                                    switch (v2 << ((18446744073709551615) & 63)) {
                                    case 32767:
                                        v35 := GenS2{F0: (("a" + v4) + v4), F1: 691}
                                        _ = v35
                                    }
                                case 829:
                                    for fi36 := 0; fi36 < 10; fi36++ {
                                        v19 += (^(v19))
                                        _ = fi36
                                    }
                                }
                                v5 = v5
                                _ = v34
                            }
                        case "a":
                            if v25 {
                                fmt.Println(v18, v20)
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                for fi37 := 0; fi37 < 1; fi37++ {
                                    if len(v16) > 0 {
                                        v16[0] = fi33
                                    }
                                    var v38 int64 = (v3 ^ 1)
                                    _ = fi37
                                    _ = v38
                                }
                            } else {
                                v1 = float64(1.5)
                            }
                        case "":
                            for fi39 := 0; fi39 < 17; fi39++ {
                                if v25 {
                                    var v40 uint32 = 865
                                    if v25 {
                                        if len(v16) > 0 {
                                            v16[0] = fi39
                                        }
                                        _ = len(v16)
                                    }
                                    _ = v40
                                } else {
                                    if v25 {
                                        var v41 uint8 = 20
                                        _ = v41
                                    }
                                    switch (v2 >> ((2) & 63)) {
                                    case -2147483648:
                                        var v42 uint = 286
                                        v20 -= 18446744073709551615
                                        _ = v42
                                    case 127:
                                        var v43 float64 = 1.5
                                        _ = v43
                                    default:
                                        if len(v16) > 0 {
                                            v16[0] = -173
                                        }
                                    }
                                }
                                _ = fi39
                            }
                        default:
                            v3 = (v3 & (v3 - (v3 * 0)))
                        }
                        v20 += v20
                        v17 = "foo bar"
                        _ = fi33
                    }
                default:
                    if (v3 != 9223372036854775807) {
                        v16 = append(v16, int(v20))
                        if true {
                            if (!v25) {
                                var v44 int = v2
                                _ = v44
                            } else {
                                if v25 {
                                    fmt.Println(v3, v17, v2)
                                }
                                fmt.Println(v1, v20)
                            }
                        }
                        v45 := GenS4{F0: 689, F1: v18}
                        _ = v45
                    } else {
                        var v46 uint = 210
                        _ = v46
                    }
                }
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
            } else {
                fmt.Println(v25, v17, v20)
            }
            switch v2 {
            case -2147483648:
                _ = len(v16)
            case 128:
                for fi47 := 0; fi47 < 1; fi47++ {
                    fmt.Println(v18, v19, fi47)
                    _ = fi47
                }
            default:
                for fi48 := 0; fi48 < 4; fi48++ {
                    if false {
                        for fi49 := 0; fi49 < 9; fi49++ {
                            v25 = v25
                            _ = fi49
                        }
                    }
                    _ = fi48
                }
            }
        } else {
            var v50 bool = (!(127 <= 127))
            switch v17 {
            case "hello":
                var v51 uint64 = (v20 >> ((0) & 63))
                _ = v51
            case "foo bar":
                v52 := GenS5{F0: true}
                _ = v52
            }
            _ = v50
        }
        for fi53 := 0; fi53 < 19; fi53++ {
            if v25 {
                v4 = ("foo bar" + v17)
                if ((((1e10 * (-2.25 * 0.1)) * 1e10) - -2.25) <= 1.5) {
                    fmt.Println(v18)
                } else {
                    var v54 int32 = 1
                    var v55 uint = 1
                    _ = v54
                    _ = v55
                }
            }
            v56 := GenS6{F0: v3}
            for fi57 := 0; fi57 < 15; fi57++ {
                v58 := GenS3{F0: 1, F1: 128, F2: 42}
                v2 = -1
                if (!(v5 > (v17 + ("\n\t" + v4)))) {
                    switch (fi53 << ((18446744073709551615) & 63)) {
                    case -691:
                        v59 := GenS5{F0: (!(v20 < 244))}
                        _ = v59
                    case 4294967295:
                        v2 += fi57
                    }
                    if v25 {
                        var v60 int16 = v19
                        v20 -= (^(v20))
                        var v61 uint8 = 42
                        v62 := GenS1{F0: 551}
                        v61 += v61
                        _ = v60
                        _ = v61
                        _ = v62
                    } else {
                        if v25 {
                            v16 = append(v16, (fi53 * 67))
                            v63 := GenS4{F0: v18, F1: (v18 % ((v18) | 1))}
                            var v64 uint = 4294967295
                            _ = v63
                            _ = v64
                        }
                    }
                }
                _ = len(v16)
                v1 = v1
                _ = fi57
                _ = v58
            }
            _ = fi53
            _ = v56
        }
        _ = v25
    }
    if len(v16) > 0 {
        v16[0] = 184
    }
    v18 -= (v18 * v18)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
