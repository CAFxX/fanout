package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 int }
type GenS4 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string; F1 int; F2 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 42
    var v3 bool = true
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 2, F1: 55, F2: 0}
        var v9 int16 = -154
        v6 = v6
        _ = fi7
        _ = v8
        _ = v9
    }
    v2 = v2
    v2 += v2
    var v10 int16 = 42
    v11 := []int{}
    if v3 {
        v10 = (v10 + v10)
        v2 = 100
        switch 701 {
        case 9223372036854775807:
            if (true || ((-(v2)) <= int8(v6))) {
                if (1 > -197) {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        for fi13 := 0; fi13 < 6; fi13++ {
                            fmt.Println(v6, v2, v6)
                            var v14 int16 = 222
                            v3 = ((v6 << ((347) & 63)) == v6)
                            _ = fi13
                            _ = v14
                        }
                        var v15 float32 = (1e10 * float32(((((((3.14159 * (0.0 + 1e10)) - (0.0 - -2.25)) + ((-2.25 + (1e10 - 0.1)) / 0.1)) * 0.0) * float64((0.0 + 3.14159))) * ((3.14159 / 3.14159) * ((1e10 * 0.0) - float64(3.14159))))))
                        _ = fi12
                        _ = v15
                    }
                    var v16 uint64 = 748
                    var v17 bool = (float32(3.14159) < (((0.1 - 0.1) - 3.14159) - -2.25))
                    _ = v16
                    _ = v17
                }
                var v18 uint32 = 4294967295
                switch -707 {
                case -592:
                    switch "foo bar" {
                    case "":
                        var v19 int8 = (^(v2))
                        _ = v19
                    case "a":
                        v18 = (v18 << ((326) & 63))
                    case "\n\t":
                        v2 = (v2 % ((-128) | 1))
                    default:
                        fmt.Println(v18, v2, v10)
                    }
                case 259:
                    switch -1 {
                    case 32767:
                        if true {
                            v10 = v10
                            var v20 uint = 0
                            for fi21 := 0; fi21 < 11; fi21++ {
                                for fi22 := 0; fi22 < 7; fi22++ {
                                    switch 32767 {
                                    case -510:
                                        v23 := GenS1{F0: v20, F1: 1, F2: int8(v20)}
                                        _ = v23
                                    case -9223372036854775808:
                                        v24 := GenS2{F0: 2}
                                        _ = v24
                                    }
                                    fmt.Println(v20, v3, v20)
                                    v2 = int8(fi22)
                                    _ = fi22
                                }
                                var v25 int64 = (v6 & (v6 % ((v6) | 1)))
                                _ = fi21
                                _ = v25
                            }
                            _ = v20
                        } else {
                            v26 := GenS1{F0: 674, F1: 42, F2: (v2 % ((v2) | 1))}
                            _ = v26
                        }
                    case 2147483647:
                        v18 -= v18
                    default:
                        switch 743 {
                        case -2147483648:
                            var v27 float64 = ((((0.0 / 1.5) + 1.5) / 1.0) + 3.14159)
                            _ = v27
                        case 42:
                            var v28 int16 = 189
                            _ = v28
                        }
                    }
                case 1:
                    v29 := GenS3{F0: 742, F1: 912, F2: -9223372036854775808}
                    _ = v29
                }
                _ = v18
            } else {
                var v30 int64 = -375
                for fi31 := 0; fi31 < 9; fi31++ {
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 int = 8
                        _ = fi32
                        _ = v33
                    }
                    var v34 uint32 = 396
                    _ = fi31
                    _ = v34
                }
                fmt.Println(v3, v6)
                _ = v30
            }
        case 1:
            for fi35 := 0; fi35 < 13; fi35++ {
                switch (-(fi35)) {
                case 127:
                    var v36 int16 = (v10 | (-(v10)))
                    _ = v36
                case 32768:
                    v37 := GenS4{F0: 4294967295, F1: 834, F2: 30}
                    _ = v37
                default:
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                }
                var v38 int64 = int64(v10)
                _ = fi35
                _ = v38
            }
        case 32767:
            for fi39 := 0; fi39 < 16; fi39++ {
                v11 = append(v11, 32767)
                var v40 uint64 = 4294967295
                _ = fi39
                _ = v40
            }
        default:
            for fi41 := 0; fi41 < 7; fi41++ {
                v42 := GenS2{F0: 2}
                _ = fi41
                _ = v42
            }
        }
    }
    var v43 int16 = (v10 << ((887) & 63))
    fmt.Println(v3, v3, v3)
    if (((v2 != v2) || (true || true)) || (!v3)) {
        v43 -= v43
        v44 := GenS5{F0: v10, F1: 0}
        if v3 {
            if false {
                v11 = append(v11, 127)
                var v45 int8 = v2
                v11 = append(v11, -1)
                _ = v45
            } else {
                v10 -= -522
                v46 := GenS6{F0: -73, F1: v10}
                v47 := GenS7{F0: (v3 && (false || true))}
                _ = v46
                _ = v47
            }
            var v48 uint8 = 68
            for fi49 := 0; fi49 < 5; fi49++ {
                if false {
                    var v50 string = "hello"
                    _ = v50
                } else {
                    v10 -= (v43 >> ((2) & 63))
                    for fi51 := 0; fi51 < 11; fi51++ {
                        v52 := GenS8{F0: 882, F1: 2}
                        if v3 {
                            v53 := GenS9{F0: 0}
                            var v54 int32 = 0
                            fmt.Println(fi51)
                            _ = v53
                            _ = v54
                        } else {
                            fmt.Println(v2, v3, v3)
                        }
                        if v3 {
                            fmt.Println(v10, v2, fi49)
                            delete(v4, (^(fi51)))
                        } else {
                            var v55 float32 = 1.5
                            if v3 {
                                v56 := GenS7{F0: v3}
                                _ = v56
                            }
                            switch "hello" {
                            case "":
                                v3 = v3
                            default:
                                _ = len(v11)
                            }
                            _ = v55
                        }
                        _ = fi51
                        _ = v52
                    }
                }
                _ = fi49
            }
            _ = v48
        }
        _ = v44
    }
    v57 := GenS10{F0: "hello", F1: 1, F2: 18446744073709551615}
    for fi58 := 0; fi58 < 2; fi58++ {
        v11 = append(v11, fi58)
        fmt.Println(v3)
        var v59 string = ("\n\t" + "hello")
        _ = fi58
        _ = v59
    }
    for fi60 := 0; fi60 < 5; fi60++ {
        v61 := GenS7{F0: (!v3)}
        v62 := GenS5{F0: v43, F1: 255}
        fmt.Println(fi60, v2)
        _ = fi60
        _ = v61
        _ = v62
    }
    for fi63 := 0; fi63 < 17; fi63++ {
        if true {
            v3 = ((!(1 > 255)) && v3)
            v3 = (!false)
            v6 -= (v6 % ((v6) | 1))
        }
        _ = fi63
    }
    v10 += 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v57
    _ = v6
}
