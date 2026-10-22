package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 string; F2 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 float32 = (1.5 / (0.0 - float32(1e10)))
    var v4 float64 = 1.5
    var v5 uint8 = 217
    switch 32767 {
    case 534:
        var v6 uint64 = 42
        _ = v6
    case 32767:
        if (0 < -140) {
            v7 := GenS1{F0: 1, F1: 4294967295}
            v4 -= ((v4 - v4) + (float64(1e10) * v4))
            _ = v7
        } else {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS2{F0: (!(4294967295 >= 0)), F1: 127}
                _ = fi8
                _ = v9
            }
            var v10 uint32 = 2147483647
            _ = v10
        }
    case 9223372036854775807:
        if true {
            v4 = ((v4 / -2.25) - (v4 + (-2.25 / 0.1)))
            if (42 < -916) {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int8 = 0
                    _ = fi11
                    _ = v12
                }
            } else {
                fmt.Println(v3, v5, v5)
            }
        } else {
            v3 = v3
            fmt.Println(v5, v5, v4)
            fmt.Println(v5)
        }
    }
    var v13 int64 = int64(v5)
    v13 = 9223372036854775807
    var v14 string = "hello"
    if true {
        v14 = "a"
        switch 32768 {
        case -1:
            for fi15 := 0; fi15 < 12; fi15++ {
                for fi16 := 0; fi16 < 12; fi16++ {
                    v17 := GenS3{F0: v13, F1: 255, F2: 2}
                    fmt.Println(v3, fi15)
                    var v18 uint16 = 255
                    v19 := GenS4{F0: 0}
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                if false {
                    v13 -= -1
                    fmt.Println(v14)
                    if (!false) {
                        fmt.Println(v3)
                        var v20 uint8 = 1
                        _ = v20
                    } else {
                        var v21 string = "a"
                        _ = v21
                    }
                } else {
                    switch "hello" {
                    case "a":
                        fmt.Println(v4, fi15, v5)
                    }
                    for fi22 := 0; fi22 < 19; fi22++ {
                        var v23 int16 = 42
                        if (!((-128 > 42) || false)) {
                            v23 -= v23
                            v23 -= (^(v23))
                            fmt.Println(fi22, v14)
                        } else {
                            _ = len(v2)
                            if (255 == 383) {
                                v24 := GenS5{F0: (((-2.25 < v3) && (4294967295 == 18446744073709551615)) && true), F1: ("hello" + v14), F2: (fi15 ^ fi15)}
                                var v25 uint32 = 65535
                                var v26 uint = 255
                                if true {
                                    fmt.Println(v26)
                                    delete(v1, (^(v25)))
                                } else {
                                    v14 = ""
                                    v27 := GenS4{F0: 100}
                                    _ = v27
                                }
                                v28 := []int{fi22, (fi22 - 128), fi22}
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v28
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                }
                _ = fi15
            }
        case 42:
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 int16 = 32767
                _ = fi29
                _ = v30
            }
        case -666:
            switch ("" + v14) {
            case "foo bar":
                for fi31 := 0; fi31 < 8; fi31++ {
                    for fi32 := 0; fi32 < 18; fi32++ {
                        var v33 uint64 = 592
                        _ = fi32
                        _ = v33
                    }
                    _ = fi31
                }
            case "a":
                if false {
                    v13 = (-(v13))
                    var v34 float32 = (1e10 * (0.0 + (1e10 * v3)))
                    v35 := GenS2{F0: true, F1: int8(v5)}
                    _ = v34
                    _ = v35
                } else {
                    fmt.Println(v13, v13, v14)
                }
            case "":
                fmt.Println(v5)
            }
        }
        for fi36 := 0; fi36 < 14; fi36++ {
            if true {
                var v37 float32 = v3
                if (!false) {
                    v2[(fi36 >> ((0) & 63))] = 1
                }
                v37 += v37
                _ = v37
            } else {
                if true {
                    var v38 int8 = 100
                    switch v14 {
                    case "\n\t":
                        var v39 uint = 163
                        _ = v39
                    case "":
                        switch -2147483648 {
                        case 375:
                            v14 = v14
                        case 9223372036854775807:
                            v4 = 3.14159
                        case 32768:
                            var v40 uint16 = 0
                            _ = v40
                        default:
                            v3 -= v3
                        }
                    case "a":
                        v5 = 1
                    default:
                        var v41 int8 = (-(v38))
                        _ = v41
                    }
                    _ = v38
                }
            }
            fmt.Println(v13, fi36)
            _ = fi36
        }
    } else {
        v42 := GenS6{F0: -1}
        _ = v42
    }
    v43 := GenS7{F0: -1}
    var v44 uint32 = 19
    v13 -= -9223372036854775808
    var v45 uint = 1
    v14 = ((v14 + v14) + "")
    var v46 int16 = 1
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
