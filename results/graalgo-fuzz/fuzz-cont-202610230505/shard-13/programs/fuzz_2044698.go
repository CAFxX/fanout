package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int8; F2 int }
type GenS4 struct { F0 uint16; F1 int16; F2 int }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int; F1 uint; F2 int8 }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint64 }
func genFn0(p0 int64, p1 uint64) uint8 {
    switch 2147483647 {
    case -2147483648:
        p0 = p0
    case 2147483647:
        fmt.Println(p0)
    default:
        v6 := []int{127, -1}
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 240
}
func genFn1(p0 int16, p1 int32, p2 int) int16 {
    v7 := GenS1{F0: 0, F1: 42}
    p0 -= p0
    v8 := GenS1{F0: 4294967295, F1: 554}
    fmt.Println(p0)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (^(p0))
}

func main() {
    v1 := []uint32{457}
    v2 := []string{("\n\t" + "")}
    var v3 float64 = 1.5
    var v4 uint64 = 42
    v5 := make(map[uint32]int)
    v9 := GenS2{F0: 65535}
    v3 -= (((0.1 - v3) + v3) + -2.25)
    v1 = append(v1, 911)
    fmt.Println(v3)
    var v10 uint32 = 2147483648
    if (464 == 42) {
        var v11 int64 = -1
        fmt.Println(v10, v10, v4)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v2 = append(v2, "\n\t")
                var v14 string = "a"
                _ = fi13
                _ = v14
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS1{F0: v4, F1: 289}
                v3 = (v3 + 1e10)
                _ = fi15
                _ = v16
            }
            _ = fi12
        }
    }
    v10 = v10
    v10 += (^(v10))
    var v17 float64 = v3
    v18 := GenS3{F0: -56, F1: 1, F2: 42}
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS1{F0: 370, F1: 42}
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: 111, F1: int8(v10), F2: -770}
    var v22 string = ""
    v3 += float64(0.1)
    var v23 uint16 = 0
    if (("" == "a") && true) {
        switch 42 {
        case -9223372036854775808:
            _ = len(v5)
        case 439:
            var v24 uint32 = 2147483647
            _ = v24
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            if false {
                if (false || (((-513 == 42) && true) && ((fi25 ^ fi25) != fi25))) {
                    fmt.Println(v10, v4, fi25)
                    for fi26 := 0; fi26 < 13; fi26++ {
                        var v27 int64 = -133
                        _ = fi26
                        _ = v27
                    }
                }
                var v28 float32 = 0.0
                _ = v28
            } else {
                var v29 float32 = float32(v17)
                _ = v29
            }
            v30 := GenS4{F0: 39, F1: int16(v10), F2: -2147483648}
            _ = fi25
            _ = v30
        }
        var v31 bool = (255 == genFn0(-637, (v4 | v4)))
        _ = v31
    } else {
        v3 = float64(-2.25)
        if (!(0 != 326)) {
            for fi32 := 0; fi32 < 19; fi32++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi32
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                if true {
                    if false {
                        v34 := GenS5{F0: "", F1: uint16(v10)}
                        fmt.Println(v23)
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 9; fi35++ {
                            if true {
                                v36 := GenS1{F0: (v4 + (v4 * (v4 >> ((0) & 63)))), F1: 42}
                                if true {
                                    var v37 uint64 = 18446744073709551615
                                    v37 = v4
                                    _ = v37
                                }
                                v38 := GenS6{F0: (fi33 & fi35), F1: 872, F2: 1}
                                _ = v36
                                _ = v38
                            }
                            var v39 uint32 = 65535
                            _ = fi35
                            _ = v39
                        }
                        if (true || (0 <= -257)) {
                            switch v22 {
                            case "\n\t":
                                switch 127 {
                                case -1:
                                    var v40 bool = (1000 >= 0)
                                    _ = v40
                                case -2147483648:
                                    switch v22 {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "a":
                                            v41 := GenS7{F0: (v23 + v23), F1: (^(v4)), F2: (v4 / (((v4 | v4)) | 1))}
                                            _ = v41
                                        case "hello":
                                            v3 = (0.0 * v17)
                                        }
                                    case "":
                                        v22 = (v22 + v22)
                                    case "\n\t":
                                        v42 := GenS4{F0: (^(v23)), F1: 42, F2: fi33}
                                        _ = v42
                                    }
                                case 0:
                                    fmt.Println(v10)
                                default:
                                    v23 += 255
                                }
                            case "foo bar":
                                v23 = v23
                            case "hello":
                                fmt.Println(v17)
                            }
                        } else {
                            var v43 string = ("foo bar" + v22)
                            fmt.Println(v43, v23)
                            for fi44 := 0; fi44 < 7; fi44++ {
                                v45 := v5[(v10 & v10)]
                                _ = fi44
                                _ = v45
                            }
                            _ = v43
                        }
                        var v46 uint32 = (^(v10))
                        _ = v46
                    }
                    if (!(-44 <= int32(v10))) {
                        if (false || true) {
                            v4 -= v4
                            switch (fi33 % ((fi33) | 1)) {
                            case 32768:
                                v4 = v4
                            case 1:
                                v4 -= v4
                            case 428:
                                for fi47 := 0; fi47 < 4; fi47++ {
                                    var v48 string = v22
                                    _ = fi47
                                    _ = v48
                                }
                            }
                            fmt.Println(v22, v10, v10)
                        } else {
                            fmt.Println(fi33, fi33, fi33)
                            if (v22 < ((v22 + v22) + v22)) {
                                for fi49 := 0; fi49 < 10; fi49++ {
                                    var v50 int64 = -9223372036854775808
                                    _ = fi49
                                    _ = v50
                                }
                            }
                            switch (v22 + v22) {
                            case "hello":
                                v10 = v10
                            case "a":
                                var v51 uint32 = (v10 - 2)
                                _ = v51
                            case "":
                                fmt.Println(v22)
                            default:
                                if (v23 == (v23 << ((uint(v4)) & 63))) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v10 += v10
                                }
                            }
                        }
                    }
                } else {
                    if true {
                        var v52 int64 = -732
                        _ = v52
                    } else {
                        v22 = ""
                        v22 = "foo bar"
                    }
                    v17 += v3
                }
                fmt.Println(v22, v22, v4)
                _ = fi33
            }
            v17 -= (v3 / v17)
        }
        for fi53 := 0; fi53 < 3; fi53++ {
            v3 = (1e10 - float64((-2.25 / 3.14159)))
            v17 += 0.0
            _ = fi53
        }
    }
    v3 = v17
    var v54 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v9
}
