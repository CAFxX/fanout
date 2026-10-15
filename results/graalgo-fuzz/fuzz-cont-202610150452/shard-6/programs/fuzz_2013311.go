package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int; F1 int64; F2 int64 }
type GenS6 struct { F0 uint16; F1 int16; F2 string }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 uint; F1 int64; F2 uint16 }
func genFn0(p0 int64) int64 {
    v4 := GenS1{F0: (p0 & p0), F1: "\n\t"}
    v5 := GenS2{F0: 2147483647, F1: 65535}
    _ = p0
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 uint8, p1 int16) int {
    switch "hello" {
    case "a":
        fmt.Println(p1)
    default:
        var v6 int64 = 0
        _ = v6
    }
    if (1 != -229) {
        switch -9223372036854775808 {
        case -951:
            switch "a" {
            case "hello":
                fmt.Println(p1)
            case "\n\t":
                for fi7 := 0; fi7 < 11; fi7++ {
                    for fi8 := 0; fi8 < 2; fi8++ {
                        if false {
                            v9 := []int{(fi7 | fi7), (^(fi8))}
                            if true {
                                var v10 int16 = p1
                                var v11 uint32 = 805
                                _ = v10
                                _ = v11
                            } else {
                                fmt.Println(p1)
                                p0 -= 128
                                p1 = (p1 + 42)
                            }
                            _ = v9
                        }
                        _ = fi8
                    }
                    p0 += p0
                    _ = fi7
                }
            }
        case -2147483648:
            fmt.Println(p1)
        default:
            if false {
                v12 := GenS1{F0: 182, F1: "\n\t"}
                for fi13 := 0; fi13 < 18; fi13++ {
                    if (fi13 > -2147483648) {
                        var v14 uint64 = 0
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 4; fi15++ {
                        if (false || true) {
                            v16 := []int{(^(fi15)), fi13, (fi13 >> ((255) & 63))}
                            switch (fi13 % ((fi15) | 1)) {
                            case 128:
                                if ("foo bar" < ("" + ("foo bar" + (("foo bar" + ("hello" + "")) + "foo bar")))) {
                                    var v17 uint16 = 32767
                                    var v18 int = 32768
                                    v19 := GenS3{F0: 2}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                            default:
                                if len(v16) > 0 {
                                    v16[0] = fi13
                                }
                            }
                            _ = v16
                        }
                        if false {
                            p1 = (-(p1))
                            v20 := GenS4{F0: "foo bar", F1: 100}
                            var v21 int32 = -2147483648
                            _ = v20
                            _ = v21
                        } else {
                            var v22 int8 = 127
                            p1 = p1
                            v23 := GenS5{F0: fi15, F1: 9223372036854775807, F2: -881}
                            _ = v22
                            _ = v23
                        }
                        _ = fi15
                    }
                    _ = fi13
                }
                _ = v12
            } else {
                p0 = 1
                p0 = 42
            }
        }
        var v24 uint32 = 257
        _ = v24
    }
    v25 := []int{0, 1, 1}
    var v26 int = 32768
    fmt.Println(v26, p1)
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return 2147483647
}

func main() {
    var v1 int8 = -1
    var v2 uint32 = 87
    var v3 uint32 = 2147483647
    v3 += v3
    if true {
        v27 := GenS2{F0: -1, F1: v2}
        var v28 uint32 = 2
        switch 32768 {
        case 32767:
            if true {
                if (-3 >= -9223372036854775808) {
                    v29 := []int{2147483647, 1, -728, 32768}
                    v2 = (^(v2))
                    v3 -= v3
                    _ = v29
                }
            } else {
                v30 := GenS6{F0: 0, F1: 42, F2: ("" + "a")}
                switch 113 {
                case -544:
                    var v31 float32 = 1e10
                    _ = v31
                case 2147483647:
                    v32 := []int{}
                    _ = v32
                case -2147483648:
                    switch 9223372036854775807 {
                    case 127:
                        for fi33 := 0; fi33 < 16; fi33++ {
                            var v34 int8 = 100
                            _ = fi33
                            _ = v34
                        }
                    default:
                        var v35 int16 = -1
                        _ = v35
                    }
                default:
                    v28 += v2
                }
                v3 = 0
                _ = v30
            }
        case 144:
            v36 := []int{genFn1(1, 32767), -9223372036854775808}
            _ = v36
        }
        _ = v27
        _ = v28
    }
    v3 -= 699
    v37 := GenS2{F0: 983, F1: 65535}
    fmt.Println(v3, v2, v3)
    v1 = (v1 << ((65535) & 63))
    if true {
        switch 555 {
        case 464:
            fmt.Println(v1, v3, v1)
        }
        var v38 int16 = -32768
        var v39 int16 = -1
        _ = v38
        _ = v39
    } else {
        v1 += v1
        if false {
            for fi40 := 0; fi40 < 6; fi40++ {
                fmt.Println(v1, v2, v1)
                v41 := GenS7{F0: 588, F1: -9223372036854775808}
                _ = fi40
                _ = v41
            }
        }
    }
    fmt.Println(v2)
    v42 := []int{0, 127, 9223372036854775807, -9223372036854775808}
    v3 -= 2147483648
    switch -280 {
    case -924:
        var v43 uint64 = 2
        _ = v43
    default:
        var v44 uint64 = 4294967295
        _ = v44
    }
    v45 := GenS7{F0: 0, F1: -1}
    if (!(genFn1(1, 32767) > 9223372036854775807)) {
        fmt.Println(v1)
        v46 := GenS8{F0: 634, F1: 42}
        v47 := GenS9{F0: 89, F1: -176, F2: 302}
        _ = v46
        _ = v47
    } else {
        for fi48 := 0; fi48 < 14; fi48++ {
            v3 += 382
            fmt.Println(fi48, v2)
            _ = fi48
        }
        var v49 uint16 = uint16(v3)
        _ = v49
    }
    if len(v42) > 0 {
        v42[0] = 42
    }
    if false {
        v50 := GenS2{F0: -285, F1: (v3 ^ v2)}
        for fi51 := 0; fi51 < 14; fi51++ {
            v42 = append(v42, fi51)
            for fi52 := 0; fi52 < 7; fi52++ {
                var v53 uint8 = 126
                v2 = (v2 >> ((421) & 63))
                v54 := GenS7{F0: v2, F1: 1}
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = fi51
        }
        _ = v50
    }
    var v55 bool = false
    for fi56 := 0; fi56 < 5; fi56++ {
        fmt.Println(v55)
        v57 := GenS1{F0: 42, F1: "a"}
        _ = fi56
        _ = v57
    }
    for fi58 := 0; fi58 < 18; fi58++ {
        for fi59 := 0; fi59 < 20; fi59++ {
            if len(v42) >= 2 {
                v42 = v42[1:len(v42)]
            }
            v60 := GenS2{F0: -498, F1: 42}
            _ = fi59
            _ = v60
        }
        for fi61 := 0; fi61 < 2; fi61++ {
            v1 -= (v1 + 1)
            if (-32768 <= 1) {
                if true {
                    if len(v42) >= 2 {
                        v42 = v42[1:len(v42)]
                    }
                } else {
                    v42 = append(v42, (fi58 | (-(fi61))))
                    v3 -= (v3 + (v2 * v3))
                }
                var v62 uint16 = 32767
                switch 4294967295 {
                case 32768:
                    var v63 string = (("foo bar" + ("hello" + ("foo bar" + "foo bar"))) + "foo bar")
                    _ = v63
                case 381:
                    var v64 int32 = 42
                    _ = v64
                case 1:
                    v65 := GenS2{F0: 2147483647, F1: v2}
                    _ = v65
                }
                if len(v42) >= 2 {
                    v42 = v42[1:len(v42)]
                }
                var v66 uint = 548
                var v67 uint16 = v62
                _ = v62
                _ = v66
                _ = v67
            } else {
                var v68 uint8 = 255
                v68 -= 0
                _ = v68
            }
            if (v55 || (449 < 650)) {
                if len(v42) > 0 {
                    v42[0] = fi61
                }
                if len(v42) > 0 {
                    v42[0] = genFn1(54, -881)
                }
            }
            _ = fi61
        }
        _ = fi58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v42
    _ = v45
    _ = v55
}
