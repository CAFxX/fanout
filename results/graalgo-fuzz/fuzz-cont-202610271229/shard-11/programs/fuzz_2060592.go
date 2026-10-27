package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint }
type GenS3 struct { F0 int; F1 int16; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint8; F1 string; F2 int32 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint8 }
func genFn0(p0 uint32, p1 bool, p2 uint64) uint16 {
    fmt.Println(p0, p1, p1)
    fmt.Println(p2, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return 775
}
func genFn1(p0 int64, p1 uint64) int16 {
    p0 = (p0 / ((p0) | 1))
    v5 := []int{-955}
    p1 = 1
    fmt.Println(p0)
    p0 += -9223372036854775808
    _ = p0
    _ = p1
    _ = v5
    return -25
}

func main() {
    v1 := []uint32{2147483647, 2, 1, 80}
    v2 := make(map[int]uint32)
    var v3 int = -471
    var v4 int8 = 1
    v4 = (v4 ^ 3)
    switch (v3 - (v3 & (-(v3)))) {
    case -1:
        if (false && true) {
            fmt.Println(v4, v3)
            fmt.Println(v4)
            switch "a" {
            case "\n\t":
                fmt.Println(v4)
            case "a":
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
            case "":
                var v6 uint64 = 76
                _ = v6
            }
        }
    case 2147483647:
        fmt.Println(v3, v4)
    default:
        switch ("a" + "hello") {
        case "a":
            switch v3 {
            case 4294967295:
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
            }
        case "\n\t":
            v4 += (v4 >> ((255) & 63))
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    var v7 uint32 = 1
    v4 -= v4
    v4 += v4
    v4 = v4
    v8 := GenS1{F0: 171, F1: (v7 / ((v7) | 1)), F2: 128}
    v7 -= 4294967295
    switch (v3 % ((v3) | 1)) {
    case 9223372036854775807:
        if ((^(v7)) > v7) {
            _ = len(v1)
            var v9 float32 = -2.25
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 int = (fi11 & (fi10 + fi11))
                    var v13 int32 = 225
                    if true {
                        _ = len(v1)
                        if (!(true && true)) {
                            var v14 uint32 = 1
                            switch ("a" + "hello") {
                            case "foo bar":
                                v9 += 0.1
                            case "hello":
                                v15 := GenS1{F0: int32(fi10), F1: v14, F2: 597}
                                _ = v15
                            case "":
                                if false {
                                    v14 = v14
                                    v16 := GenS1{F0: (v13 | v13), F1: 2, F2: 1}
                                    var v17 int64 = 42
                                    _ = v16
                                    _ = v17
                                }
                            }
                            _ = v14
                        }
                        v18 := GenS1{F0: (v13 & v13), F1: (v7 | (v7 & 439)), F2: 2}
                        _ = v18
                    } else {
                        v9 += float32((1.5 - 1.5))
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
            _ = v9
        } else {
            for fi19 := 0; fi19 < 4; fi19++ {
                if (-1 > 1) {
                    delete(v2, 9223372036854775807)
                    var v20 uint64 = 707
                    v20 = 687
                    _ = v20
                } else {
                    if ((genFn0(v7, (0 != 2), 42) <= 148) || true) {
                        var v21 uint32 = 65535
                        switch "hello" {
                        case "a":
                            v7 += 419
                        case "foo bar":
                            for fi22 := 0; fi22 < 5; fi22++ {
                                var v23 int8 = (v4 % (((v4 / ((v4) | 1))) | 1))
                                v21 = 159
                                _ = fi22
                                _ = v23
                            }
                        case "":
                            v21 -= 2147483648
                        default:
                            var v24 float64 = 1e10
                            _ = v24
                        }
                        var v25 uint32 = (v7 ^ v21)
                        _ = v21
                        _ = v25
                    } else {
                        var v26 uint32 = 65535
                        for fi27 := 0; fi27 < 16; fi27++ {
                            if ((!(true && false)) || true) {
                                v28 := GenS2{F0: (v26 - 81), F1: 477, F2: 18446744073709551615}
                                _ = v28
                            } else {
                                fmt.Println(v4, v26)
                                v4 += v4
                            }
                            v7 += (^(v7))
                            _ = fi27
                        }
                        v29 := GenS3{F0: (fi19 * 127), F1: genFn1(-1, 4294967295), F2: v4}
                        _ = v26
                        _ = v29
                    }
                }
                v3 += -783
                fmt.Println(v7, v4, v3)
                _ = fi19
            }
            var v30 uint64 = 0
            _ = v30
        }
    }
    v4 = -114
    switch (v3 * v3) {
    case 365:
        v4 -= (^(v4))
    default:
        if (!((!(!true)) || ((4294967295 <= 885) && ("hello" <= "a")))) {
            if true {
                var v31 float64 = 3.14159
                v31 += 1e10
                _ = v31
            } else {
                v4 = v4
                switch "a" {
                case "":
                    switch ("foo bar" + ("a" + ("a" + ""))) {
                    case "foo bar":
                        for fi32 := 0; fi32 < 5; fi32++ {
                            switch "" {
                            case "foo bar":
                                var v33 int16 = 42
                                _ = v33
                            }
                            for fi34 := 0; fi34 < 1; fi34++ {
                                var v35 int32 = 0
                                switch fi34 {
                                case -1:
                                    v7 -= v7
                                    v36 := v2[(fi34 * (fi34 ^ (fi32 % ((-839) | 1))))]
                                    v37 := GenS4{F0: ""}
                                    var v38 bool = false
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                case 0:
                                    fmt.Println(v4)
                                case -219:
                                    var v39 float32 = ((1.5 + 3.14159) / 1.0)
                                    _ = v39
                                }
                                _ = fi34
                                _ = v35
                            }
                            _ = fi32
                        }
                    default:
                        for fi40 := 0; fi40 < 4; fi40++ {
                            fmt.Println(v7, v4)
                            var v41 float64 = 3.14159
                            var v42 bool = (!("hello" < "hello"))
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    }
                case "hello":
                    v43 := GenS1{F0: 890, F1: (v7 * 560), F2: 255}
                    _ = v43
                case "\n\t":
                    v7 = 2147483648
                default:
                    var v44 float64 = float64(0.1)
                    _ = v44
                }
                v45 := GenS5{F0: 2, F1: v3}
                _ = v45
            }
            if (740 != 2) {
                var v46 uint = 457
                v3 += 625
                _ = v46
            } else {
                v3 = -2147483648
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if (0 > 399) {
                v47 := GenS2{F0: v7, F1: 657, F2: 255}
                _ = v47
            } else {
                var v48 int16 = 0
                _ = v48
            }
            if true {
                var v49 uint = 255
                if true {
                    fmt.Println(v3, v7)
                    v49 = v49
                }
                _ = v49
            } else {
                v50 := GenS4{F0: ("a" + "\n\t")}
                fmt.Println(v4)
                v7 = 42
                _ = v50
            }
            v51 := GenS6{F0: 1, F1: "hello", F2: -1}
            _ = v51
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi52 := 0; fi52 < 7; fi52++ {
        v4 = v4
        _ = fi52
    }
    var v53 int64 = 1
    var v54 int32 = 0
    for fi55 := 0; fi55 < 12; fi55++ {
        fmt.Println(v4, v4, v53)
        var v56 int = 4294967295
        _ = fi55
        _ = v56
    }
    if (3.14159 >= 0.0) {
        v53 = v53
        fmt.Println(v7, v4)
    } else {
        if true {
            _ = len(v1)
            v53 = v53
        }
        var v57 bool = true
        for fi58 := 0; fi58 < 17; fi58++ {
            v59 := GenS7{F0: 194, F1: 4294967295, F2: 42}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v7 -= v7
            _ = fi58
            _ = v59
        }
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v7
    _ = v8
}
