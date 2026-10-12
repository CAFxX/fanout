package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int64; F1 int32; F2 int8 }
func genFn0(p0 bool, p1 int16) uint64 {
    switch 1 {
    case 2147483647:
        var v4 string = "hello"
        _ = v4
    case -1:
        if p0 {
            v5 := []int{}
            _ = v5
        }
    case 32768:
        v6 := GenS1{F0: -143}
        _ = v6
    }
    if (386 <= 42) {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int64 = 0
            v9 := []int{fi7, (^(fi7)), (^(fi7))}
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        var v10 uint32 = 140
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                if (!(p0 && (42 > 0))) {
                    v13 := GenS1{F0: -2147483648}
                    switch (fi11 - 0) {
                    case -9223372036854775808:
                        var v14 uint64 = 2
                        _ = v14
                    default:
                        var v15 float32 = 1e10
                        _ = v15
                    }
                    var v16 string = "a"
                    _ = v13
                    _ = v16
                } else {
                    var v17 uint32 = (^(v10))
                    _ = v17
                }
                _ = fi12
            }
            v18 := GenS1{F0: 906}
            _ = fi11
            _ = v18
        }
        switch "foo bar" {
        case "foo bar":
            if p0 {
                switch "\n\t" {
                case "a":
                    if ((-(p1)) > p1) {
                        switch "foo bar" {
                        case "":
                            p0 = (!false)
                        case "a":
                            v19 := GenS2{F0: -1, F1: 42, F2: uint8(v10)}
                            _ = v19
                        }
                        var v20 uint64 = 768
                        _ = v20
                    }
                case "foo bar":
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v22 := []int{(fi21 >> ((uint(v10)) & 63)), (fi21 / ((42) | 1)), (fi21 / ((-403) | 1))}
                        switch "a" {
                        case "a":
                            _ = len(v22)
                        case "foo bar":
                            if p0 {
                                v10 = v10
                            }
                        default:
                            p1 += (p1 << ((4294967295) & 63))
                        }
                        _ = fi21
                        _ = v22
                    }
                default:
                    var v23 string = "a"
                    _ = v23
                }
                p0 = p0
                var v24 string = "foo bar"
                _ = v24
            }
        case "a":
            v25 := []int{127, -1}
            _ = v25
        case "\n\t":
            var v26 int32 = -492
            _ = v26
        default:
            var v27 int = 1
            _ = v27
        }
        _ = v10
    }
    var v28 uint32 = 180
    _ = p0
    _ = p1
    _ = v28
    return 0
}
func genFn1(p0 bool, p1 bool) int16 {
    var v29 uint16 = 0
    var v30 uint = 227
    v31 := GenS2{F0: -9223372036854775808, F1: (^(v30)), F2: 244}
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    _ = v31
    return 1000
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int8 = -115
    var v3 float32 = 0.1
    var v32 uint64 = 973
    v2 -= (v2 * (-(v2)))
    switch "" {
    case "a":
        v33 := GenS1{F0: int32(v32)}
        _ = v33
    default:
        if (42 != 9223372036854775807) {
            switch 128 {
            case 127:
                if true {
                    switch int(v2) {
                    case 695:
                        v34 := GenS1{F0: 2147483647}
                        _ = v34
                    case 1:
                        fmt.Println(v1, v1, v32)
                    case 42:
                        v35 := []int{0}
                        _ = v35
                    }
                }
            }
            var v36 int64 = 42
            _ = v36
        } else {
            if false {
                v1 += (v1 << ((42) & 63))
            } else {
                v3 += 1.5
            }
            for fi37 := 0; fi37 < 11; fi37++ {
                v38 := GenS3{F0: 141, F1: (42 == -32768)}
                _ = fi37
                _ = v38
            }
            if true {
                v39 := GenS1{F0: 466}
                _ = v39
            }
        }
    }
    for fi40 := 0; fi40 < 10; fi40++ {
        v41 := []int{-2147483648}
        switch (fi40 / ((fi40) | 1)) {
        case -2147483648:
            fmt.Println(v1, fi40, v32)
        case 437:
            if true {
                if len(v41) >= 2 {
                    v41 = v41[1:len(v41)]
                }
                v2 += (v2 >> ((42) & 63))
                _ = len(v41)
            } else {
                for fi42 := 0; fi42 < 9; fi42++ {
                    if true {
                        var v43 int64 = -1
                        _ = v43
                    }
                    if (false || true) {
                        if ((fi40 < fi42) && false) {
                            switch -2147483648 {
                            case 32768:
                                var v44 int64 = 1
                                _ = v44
                            }
                            if len(v41) >= 2 {
                                v41 = v41[1:len(v41)]
                            }
                            if len(v41) > 0 {
                                v41[0] = (fi42 << ((255) & 63))
                            }
                        }
                        var v45 int8 = (v2 + v2)
                        v1 = v1
                        _ = v45
                    }
                    v46 := GenS1{F0: 2147483647}
                    _ = fi42
                    _ = v46
                }
                v32 = (v32 ^ 18446744073709551615)
                v32 = (v32 | v32)
            }
        case 918:
            v47 := GenS1{F0: -2147483648}
            _ = v47
        }
        _ = fi40
        _ = v41
    }
    v48 := GenS4{F0: ("\n\t" + ""), F1: -64}
    v49 := []int{4294967295}
    switch -9223372036854775808 {
    case 0:
        var v50 uint64 = v32
        _ = v50
    case -9223372036854775808:
        var v51 uint = 18446744073709551615
        _ = v51
    }
    fmt.Println(v2)
    v52 := GenS2{F0: -537, F1: 255, F2: 52}
    fmt.Println(v1)
    v3 -= (float32(1e10) - (0.1 * 0.1))
    for fi53 := 0; fi53 < 18; fi53++ {
        switch "hello" {
        case "hello":
            for fi54 := 0; fi54 < 18; fi54++ {
                var v55 uint = 65535
                _ = fi54
                _ = v55
            }
        case "\n\t":
            v56 := GenS3{F0: 191, F1: true}
            _ = v56
        case "a":
            if ((false || (false && false)) && false) {
                v2 = (v2 * v2)
                if (0 != 255) {
                    for fi57 := 0; fi57 < 19; fi57++ {
                        v58 := GenS5{F0: -775, F1: 522, F2: int8(v32)}
                        v32 += (v32 * (v32 % ((321) | 1)))
                        _ = fi57
                        _ = v58
                    }
                } else {
                    fmt.Println(fi53)
                    if false {
                        switch "hello" {
                        case "a":
                            if (2 != 42) {
                                if len(v49) >= 2 {
                                    v49 = v49[1:len(v49)]
                                }
                                v49 = append(v49, (fi53 % (((fi53 % ((32768) | 1))) | 1)))
                                v59 := GenS3{F0: 114, F1: (false || ((v32 ^ (^(v32))) < (^(v32))))}
                                _ = v59
                            }
                        case "foo bar":
                            if len(v49) > 0 {
                                v49[0] = fi53
                            }
                        default:
                            v49 = append(v49, 9223372036854775807)
                        }
                        v2 = 100
                        _ = len(v49)
                    }
                    fmt.Println(v3, v3, v2)
                }
                var v60 uint16 = 831
                _ = v60
            }
        }
        if false {
            fmt.Println(v2)
        } else {
            v32 = (v32 & (v32 * v32))
        }
        for fi61 := 0; fi61 < 4; fi61++ {
            var v62 bool = (false || (-2147483648 >= 0))
            _ = fi61
            _ = v62
        }
        _ = fi53
    }
    v1 = (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v49:", len(v49), v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v48
    _ = v49
    _ = v52
}
