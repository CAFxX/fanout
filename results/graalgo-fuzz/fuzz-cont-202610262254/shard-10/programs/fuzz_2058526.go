package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int; F2 uint64 }
type GenS4 struct { F0 int64; F1 int64; F2 uint }
type GenS5 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0(p0 uint32, p1 uint64) uint16 {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 int8 = 67
        p1 -= 0
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 18446744073709551615, F1: 42, F2: v5}
            var v8 string = ("hello" + ("\n\t" + "hello"))
            v9 := GenS2{F0: (v5 % ((v5) | 1))}
            var v10 uint = 65535
            v11 := []int{127, int(v10)}
            _ = fi6
            _ = v10
            _ = v11
            _ = v7
            _ = v8
            _ = v9
        }
        _ = fi4
        _ = v5
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint64, p1 int16, p2 uint16) uint16 {
    p0 = (p0 % ((0) | 1))
    p2 += (^(p2))
    p0 = (p0 - (p0 ^ p0))
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{("foo bar" + ""), "hello", ("" + "foo bar")}
    v3 := make(map[int]uint32)
    var v12 int64 = 452
    fmt.Println(v12)
    v13 := GenS3{F0: (v12 >> ((2) & 63)), F1: -2147483648, F2: 42}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v14 := GenS3{F0: 42, F1: 128, F2: 2}
    v15 := GenS2{F0: 127}
    var v16 uint64 = 0
    var v17 uint = 782
    for fi18 := 0; fi18 < 4; fi18++ {
        if (!(false || (0.0 > 1e10))) {
            v19 := GenS2{F0: 42}
            _ = v19
        } else {
            if ("foo bar" > "\n\t") {
                if true {
                    v3[(fi18 ^ 2147483647)] = 423
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    _ = len(v1)
                }
            }
            if false {
                for fi20 := 0; fi20 < 6; fi20++ {
                    v16 = (v16 - v16)
                    v12 = (v12 * v12)
                    _ = fi20
                }
                switch "" {
                case "\n\t":
                    if false {
                        switch "\n\t" {
                        case "foo bar":
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                        case "\n\t":
                            var v21 int32 = 2147483647
                            _ = v21
                        default:
                            var v22 uint = (v17 / ((v17) | 1))
                            _ = v22
                        }
                        var v23 int32 = -764
                        fmt.Println(v23, v17)
                        _ = v23
                    } else {
                        _ = len(v1)
                    }
                }
                var v24 string = "foo bar"
                _ = v24
            } else {
                var v25 uint64 = v16
                _ = v25
            }
        }
        switch "" {
        case "foo bar":
            v26 := v3[(-(fi18))]
            _ = v26
        case "\n\t":
            v17 = 65535
        }
        _ = fi18
    }
    var v27 uint16 = 0
    var v28 bool = ((^(v16)) != (^(v16)))
    var v29 string = ((("" + "hello") + "\n\t") + "foo bar")
    for fi30 := 0; fi30 < 3; fi30++ {
        _ = len(v2)
        _ = fi30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 int64 = v12
        if (!v28) {
            switch ("a" + ("hello" + ("hello" + v29))) {
            case "\n\t":
                v16 = v16
            case "hello":
                var v33 bool = (!false)
                _ = v33
            }
        } else {
            for fi34 := 0; fi34 < 10; fi34++ {
                if (uint32(v17) <= 458) {
                    v29 = v29
                    switch (v29 + v29) {
                    case "hello":
                        _ = len(v3)
                    case "":
                        v1[v29] = v29
                    case "foo bar":
                        switch "" {
                        case "foo bar":
                            _ = len(v2)
                        case "hello":
                            var v35 int16 = -1
                            _ = v35
                        case "a":
                            var v36 string = v29
                            _ = v36
                        }
                    }
                    var v37 int32 = -1
                    _ = v37
                } else {
                    var v38 string = v29
                    v27 = (^(v27))
                    _ = v38
                }
                v17 = v17
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v29 = v29
                } else {
                    switch v29 {
                    case "a":
                        _ = len(v2)
                    case "hello":
                        v27 -= v27
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v39 float64 = 3.14159
                    var v40 uint16 = (v27 << ((v17) & 63))
                    _ = v39
                    _ = v40
                }
                _ = fi34
            }
            var v41 uint32 = 47
            _ = v41
        }
        _ = fi31
        _ = v32
    }
    switch 2147483647 {
    case 2147483647:
        v12 -= 42
    case 418:
        delete(v1, v29)
    case 9223372036854775807:
        if (false || ((v12 >> ((v17) & 63)) >= (^(v12)))) {
            var v42 float64 = (3.14159 + 1.5)
            if v28 {
                v12 = v12
                for fi43 := 0; fi43 < 4; fi43++ {
                    v44 := GenS1{F0: 255, F1: 1, F2: -128}
                    var v45 int32 = 42
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                fmt.Println(v12)
            } else {
                var v46 uint64 = v16
                _ = v46
            }
            v27 = (v27 * genFn1(v16, -1, v27))
            _ = v42
        }
    default:
        fmt.Println(v28)
    }
    if false {
        v28 = (v12 < 0)
        if v28 {
            for fi47 := 0; fi47 < 11; fi47++ {
                v16 = (v16 ^ (v16 * 0))
                v48 := GenS4{F0: 0, F1: (v12 ^ -9223372036854775808), F2: (v17 - 255)}
                for fi49 := 0; fi49 < 4; fi49++ {
                    v28 = v28
                    _ = fi49
                }
                _ = fi47
                _ = v48
            }
        } else {
            switch 128 {
            case 113:
                v12 += (v12 << ((v17) & 63))
            case 1:
                switch 4294967295 {
                case 42:
                    for fi50 := 0; fi50 < 16; fi50++ {
                        var v51 int = fi50
                        v16 = v16
                        _ = fi50
                        _ = v51
                    }
                case 4294967295:
                    v17 += 4294967295
                case -9223372036854775808:
                    if v28 {
                        switch "hello" {
                        case "hello":
                            switch 32767 {
                            case -725:
                                for fi52 := 0; fi52 < 18; fi52++ {
                                    v27 = genFn1((v16 ^ (v16 - 203)), -278, 2)
                                    switch (fi52 / (((fi52 / ((fi52) | 1))) | 1)) {
                                    case 9223372036854775807:
                                        fmt.Println(v27)
                                    }
                                    _ = fi52
                                }
                            case -2147483648:
                                var v53 uint = 65535
                                _ = v53
                            case 128:
                                v54 := GenS3{F0: v12, F1: -2147483648, F2: uint64(v17)}
                                _ = v54
                            default:
                                var v55 uint = 569
                                _ = v55
                            }
                        default:
                            for fi56 := 0; fi56 < 12; fi56++ {
                                v29 = "foo bar"
                                if (false || (!(genFn0(2147483647, 1) != v27))) {
                                    var v57 float32 = float32(0.1)
                                    _ = v57
                                }
                                _ = fi56
                            }
                        }
                        fmt.Println(v12, v28, v17)
                    } else {
                        switch 102 {
                        case 32767:
                            v58 := GenS5{F0: 128, F1: 42, F2: (v16 << (((v17 - v17)) & 63))}
                            _ = v58
                        default:
                            var v59 uint8 = 1
                            _ = v59
                        }
                    }
                default:
                    v60 := v3[32767]
                    _ = v60
                }
            default:
                v27 += v27
            }
            var v61 int32 = int32(v27)
            _ = v61
        }
    } else {
        var v62 int32 = 441
        if ((^(v16)) < 18446744073709551615) {
            v17 = (v17 | v17)
        }
        _ = v62
    }
    fmt.Println(v28)
    for fi63 := 0; fi63 < 19; fi63++ {
        if ((v27 * (v27 & v27)) <= genFn1(v16, 300, genFn0(2, 42))) {
            var v64 string = "a"
            _ = v64
        }
        _ = fi63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
