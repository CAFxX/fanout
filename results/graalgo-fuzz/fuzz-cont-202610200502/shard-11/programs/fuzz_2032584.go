package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint; F1 int32; F2 string }
type GenS5 struct { F0 int32; F1 string; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 uint }
func genFn0(p0 uint32, p1 int, p2 int) int32 {
    p1 += (p1 >> ((2) & 63))
    var v5 uint16 = 1
    fmt.Println(p2)
    switch ("hello" + "") {
    case "hello":
        if true {
            if true {
                for fi6 := 0; fi6 < 17; fi6++ {
                    switch p1 {
                    case -2147483648:
                        p1 -= fi6
                    case 4294967295:
                        v5 = 2
                    case -9223372036854775808:
                        var v7 int32 = 844
                        _ = v7
                    }
                    p1 = (p2 | fi6)
                    _ = fi6
                }
                fmt.Println(p2, v5)
            }
            if false {
                fmt.Println(p1, p1, v5)
            } else {
                for fi8 := 0; fi8 < 2; fi8++ {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        fmt.Println(fi9)
                        v10 := []int{p1, p2}
                        var v11 int8 = -59
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    _ = fi8
                }
            }
            v12 := []int{p1, p1, (p1 >> ((42) & 63))}
            _ = v12
        }
    case "":
        p0 -= p0
    case "foo bar":
        v13 := GenS1{F0: 981, F1: 1}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 uint8 = 42
    var v14 bool = (42 == 0)
    fmt.Println(v4)
    v4 = (v4 ^ v4)
    v15 := GenS2{F0: 1, F1: 0}
    var v16 int64 = 42
    switch "a" {
    case "foo bar":
        v17 := GenS2{F0: 255, F1: 0}
        _ = v17
    case "a":
        v18 := GenS3{F0: 545, F1: 18446744073709551615}
        _ = v18
    default:
        v19 := GenS4{F0: 4294967295, F1: genFn0(1, 4294967295, -1), F2: "foo bar"}
        _ = v19
    }
    var v20 float64 = float64(0.0)
    switch "hello" {
    case "\n\t":
        v21 := GenS3{F0: 32767, F1: uint(v4)}
        _ = v21
    case "a":
        var v22 string = ("\n\t" + "foo bar")
        _ = v22
    }
    switch "hello" {
    case "hello":
        for fi23 := 0; fi23 < 4; fi23++ {
            v20 = -2.25
            _ = fi23
        }
    case "a":
        switch 683 {
        case -348:
            if (128 < -1) {
                fmt.Println(v4, v14)
                v20 = 0.0
                fmt.Println(v14)
            }
        default:
            v24 := GenS5{F0: 2147483647, F1: "\n\t", F2: -1}
            _ = v24
        }
    default:
        v25 := GenS3{F0: 32767, F1: 65535}
        _ = v25
    }
    var v26 string = ("hello" + "\n\t")
    v20 = v20
    switch (v26 + v26) {
    case "":
        var v27 int16 = 32767
        _ = v27
    case "hello":
        v14 = v14
    case "\n\t":
        var v28 string = "\n\t"
        _ = v28
    }
    var v29 uint64 = 406
    var v30 uint = 165
    if (!false) {
        if (-2147483648 > -204) {
            v1[2] = 1
            if v14 {
                v31 := GenS6{F0: 297}
                for fi32 := 0; fi32 < 16; fi32++ {
                    var v33 string = "hello"
                    v34 := GenS6{F0: 2147483647}
                    if (32767 >= (fi32 / ((1) | 1))) {
                        if (-83 != -96) {
                            var v35 int32 = genFn0(883, fi32, -1)
                            _ = v35
                        } else {
                            if (false || true) {
                                for fi36 := 0; fi36 < 5; fi36++ {
                                    v37 := []int{fi32, fi32}
                                    _ = len(v3)
                                    _ = fi36
                                    _ = v37
                                }
                            }
                            var v38 string = v33
                            _ = v38
                        }
                    } else {
                        switch "hello" {
                        case "foo bar":
                            delete(v2, 2147483647)
                        case "":
                            var v39 uint16 = 42
                            _ = v39
                        default:
                            switch v26 {
                            case "a":
                                v20 -= (float64(0.0) * (v20 / v20))
                            case "hello":
                                v40 := GenS5{F0: -272, F1: "hello", F2: 127}
                                _ = v40
                            case "foo bar":
                                v41 := GenS7{F0: 1000, F1: 387}
                                _ = v41
                            }
                        }
                        for fi42 := 0; fi42 < 10; fi42++ {
                            v43 := GenS5{F0: 945, F1: "a", F2: -128}
                            if v14 {
                                switch fi32 {
                                case -2147483648:
                                    v2[uint32(v4)] = v14
                                case 324:
                                    var v44 int64 = (v16 / ((v16) | 1))
                                    _ = v44
                                default:
                                    if v14 {
                                        v45 := GenS3{F0: 32767, F1: (^(v30))}
                                        v29 += v29
                                        delete(v3, fi32)
                                        v46 := GenS4{F0: v30, F1: genFn0(4294967295, (-(fi32)), int(v30)), F2: v26}
                                        switch -9223372036854775808 {
                                        case 1:
                                            v47 := GenS8{F0: v30}
                                            _ = v47
                                        case -9223372036854775808:
                                            var v48 bool = (724 != 2)
                                            _ = v48
                                        }
                                        _ = v45
                                        _ = v46
                                    } else {
                                        v30 = (v30 ^ (v30 & (v30 << ((900) & 63))))
                                    }
                                }
                                if v14 {
                                    var v49 uint16 = 42
                                    _ = v49
                                }
                                switch (v33 + "") {
                                case "foo bar":
                                    v20 = v20
                                case "hello":
                                    v50 := []int{(fi42 ^ 2147483647)}
                                    var v51 int16 = int16(fi42)
                                    v52 := GenS7{F0: (v51 + 42), F1: (v30 / (((v30 - (v30 & (v30 << ((v30) & 63))))) | 1))}
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                case "\n\t":
                                    if v14 {
                                        var v53 float64 = v20
                                        v54 := []int{fi42, (fi32 | 32767), fi32}
                                        _ = v53
                                        _ = v54
                                    }
                                }
                            }
                            v20 = 0.0
                            _ = fi42
                            _ = v43
                        }
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                var v55 int8 = 14
                _ = v31
                _ = v55
            } else {
                var v56 float64 = ((v20 / v20) - v20)
                _ = v56
            }
        }
        for fi57 := 0; fi57 < 10; fi57++ {
            v20 = float64(1.5)
            _ = fi57
        }
        var v58 int8 = -128
        _ = v58
    }
    v59 := GenS8{F0: v30}
    switch "hello" {
    case "":
        if (-2147483648 != 2147483647) {
            var v60 uint64 = 199
            _ = v60
        } else {
            var v61 int8 = 127
            switch ("" + "\n\t") {
            case "a":
                v20 = -2.25
            }
            switch -78 {
            case -720:
                v20 -= v20
            default:
                fmt.Println(v29, v16)
            }
            _ = v61
        }
    case "hello":
        var v62 int8 = -96
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v59
}
