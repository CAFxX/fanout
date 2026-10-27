package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int64; F1 int }
func genFn0(p0 int16, p1 int) uint8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(p1, fi8, fi8)
            var v9 int64 = 705
            var v10 int64 = -865
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 bool = (false && true)
        _ = fi7
        _ = v11
    }
    fmt.Println(p1, p1)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int16) uint8 {
    v12 := GenS1{F0: -1}
    v13 := []int{42}
    _ = p0
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := []uint32{42, 162}
    v2 := make(map[int]uint32)
    var v3 uint16 = 55
    v4 := []int{-9223372036854775808, 4294967295}
    v5 := []int{-9223372036854775808, 0, 963, 2147483647}
    v6 := []string{(("foo bar" + ("" + "a")) + (("" + (("a" + "\n\t") + "foo bar")) + ("a" + "a"))), ("hello" + "foo bar"), "foo bar"}
    var v14 bool = true
    v3 += (v3 << ((uint(v3)) & 63))
    var v15 int8 = 1
    if true {
        var v16 int32 = 0
        _ = v16
    } else {
        switch 32768 {
        case -796:
            if v14 {
                if ((1 != 42) && v14) {
                    v15 -= 1
                    v14 = false
                } else {
                    if v14 {
                        var v17 bool = v14
                        var v18 string = "hello"
                        _ = v17
                        _ = v18
                    } else {
                        v3 = 42
                        switch -146 {
                        case -2147483648:
                            for fi19 := 0; fi19 < 5; fi19++ {
                                switch 2147483647 {
                                case 864:
                                    var v20 uint32 = 2147483647
                                    _ = v20
                                case 42:
                                    var v21 int16 = 1
                                    _ = v21
                                default:
                                    for fi22 := 0; fi22 < 17; fi22++ {
                                        v15 = -1
                                        v15 = v15
                                        _ = fi22
                                    }
                                }
                                v23 := GenS2{F0: 2}
                                if (2 <= 2) {
                                    for fi24 := 0; fi24 < 5; fi24++ {
                                        var v25 uint64 = 18446744073709551615
                                        var v26 int = 42
                                        v27 := v2[4294967295]
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                } else {
                                    for fi28 := 0; fi28 < 9; fi28++ {
                                        if v14 {
                                            v14 = v14
                                            var v29 uint64 = 4294967295
                                            _ = v29
                                        }
                                        _ = len(v2)
                                        _ = fi28
                                    }
                                    v5 = append(v5, (fi19 % ((fi19) | 1)))
                                    v2[9223372036854775807] = 0
                                }
                                _ = fi19
                                _ = v23
                            }
                        case 1:
                            fmt.Println(v14, v15, v3)
                        }
                        v15 = v15
                    }
                    if (!v14) {
                        v30 := GenS3{F0: 42, F1: 2, F2: 65535}
                        _ = v30
                    }
                }
                v14 = v14
                v2[int(v3)] = 1
            }
        case 32767:
            delete(v2, 0)
        }
        var v31 string = "a"
        v31 = (v31 + v31)
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        for fi33 := 0; fi33 < 6; fi33++ {
            v2[(^(fi32))] = 1
            _ = fi33
        }
        v3 = (v3 + 65535)
        switch 0 {
        case 253:
            if v14 {
                var v34 int = (fi32 * fi32)
                switch "" {
                case "hello":
                    v35 := GenS2{F0: 0}
                    _ = v35
                default:
                    if v14 {
                        if (!(-1 >= 0)) {
                            if v14 {
                                v3 += (v3 + (v3 / ((2) | 1)))
                                if ((!(!(v14 || v14))) || (false && (!true))) {
                                    var v36 string = "foo bar"
                                    _ = v36
                                } else {
                                    for fi37 := 0; fi37 < 13; fi37++ {
                                        if v14 {
                                            var v38 uint8 = 0
                                            var v39 uint = 2
                                            _ = v38
                                            _ = v39
                                        }
                                        var v40 float32 = 3.14159
                                        _ = fi37
                                        _ = v40
                                    }
                                    v15 += 127
                                }
                                var v41 int8 = (-(v15))
                                _ = v41
                            }
                        } else {
                            for fi42 := 0; fi42 < 1; fi42++ {
                                v43 := GenS4{F0: 9223372036854775807, F1: (v34 << ((4294967295) & 63))}
                                _ = fi42
                                _ = v43
                            }
                        }
                        if v14 {
                            fmt.Println(v3)
                            if v14 {
                                fmt.Println(v14, fi32, v14)
                                var v44 uint8 = genFn1(42)
                                _ = v44
                            }
                        }
                    } else {
                        var v45 int64 = 42
                        _ = v45
                    }
                }
                _ = v34
            } else {
                for fi46 := 0; fi46 < 11; fi46++ {
                    if true {
                        delete(v2, (^(fi32)))
                    }
                    _ = fi46
                }
                if (v14 || (0.0 < (1e10 - 3.14159))) {
                    var v47 bool = v14
                    _ = v47
                } else {
                    v3 = 65535
                }
                fmt.Println(fi32)
            }
        case -950:
            v15 = 127
        case 32768:
            if (128 > 2) {
                var v48 int32 = -14
                v3 -= v3
                for fi49 := 0; fi49 < 8; fi49++ {
                    var v50 uint8 = 115
                    _ = fi49
                    _ = v50
                }
                _ = v48
            }
        }
        _ = fi32
    }
    for fi51 := 0; fi51 < 3; fi51++ {
        v3 = v3
        v52 := GenS1{F0: -1}
        if (255 <= 18446744073709551615) {
            switch "foo bar" {
            case "foo bar":
                if false {
                    switch "hello" {
                    case "a":
                        if ((fi51 * -345) < (fi51 >> ((2) & 63))) {
                            var v53 uint = 0
                            for fi54 := 0; fi54 < 1; fi54++ {
                                v3 += (v3 ^ 604)
                                v3 += 42
                                _ = fi54
                            }
                            _ = v53
                        } else {
                            v15 += 42
                            fmt.Println(fi51)
                        }
                    default:
                        v3 += (v3 - v3)
                    }
                    v15 = (^(v15))
                    v55 := GenS1{F0: 9223372036854775807}
                    _ = v55
                }
            case "hello":
                _ = len(v2)
            case "\n\t":
                v14 = (false && v14)
            }
            var v56 uint32 = 2147483648
            fmt.Println(fi51, fi51, v14)
            _ = v56
        }
        _ = fi51
        _ = v52
    }
    for fi57 := 0; fi57 < 3; fi57++ {
        if false {
            if (!v14) {
                v14 = false
            } else {
                for fi58 := 0; fi58 < 3; fi58++ {
                    var v59 bool = false
                    _ = len(v6)
                    _ = fi58
                    _ = v59
                }
            }
            var v60 uint = 4294967295
            v60 = (v60 & v60)
            _ = v60
        }
        _ = fi57
    }
    v61 := GenS1{F0: -918}
    v15 = 0
    _ = len(v6)
    switch 964 {
    case 2147483647:
        if (-895 == 42) {
            var v62 uint8 = 42
            var v63 uint16 = 0
            _ = v62
            _ = v63
        }
    case 0:
        fmt.Println(v3, v3, v15)
    default:
        if (v14 && v14) {
            fmt.Println(v14, v3)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
