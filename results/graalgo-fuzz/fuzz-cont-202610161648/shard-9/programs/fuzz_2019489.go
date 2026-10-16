package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint32 }

func main() {
    var v1 bool = true
    var v2 bool = true
    var v3 uint8 = 0
    var v4 uint32 = 2147483647
    var v5 uint16 = 48
    switch 2147483647 {
    case 32767:
        var v6 bool = (!v1)
        _ = v6
    case 842:
        switch 4294967295 {
        case 32767:
            v7 := GenS1{F0: 107, F1: 42, F2: (^(v3))}
            _ = v7
        }
    }
    v8 := GenS2{F0: (("" + "a") + "\n\t"), F1: -1}
    v1 = (-19 >= 101)
    if false {
        if true {
            var v9 uint64 = 198
            v10 := []int{-131, 2147483647, 0, -9223372036854775808}
            var v11 bool = (2147483647 < 108)
            _ = v10
            _ = v11
            _ = v9
        } else {
            v4 += 1
            var v12 uint8 = v3
            _ = v12
        }
        v3 = (v3 >> ((0) & 63))
    } else {
        var v13 uint16 = (v5 - v5)
        for fi14 := 0; fi14 < 12; fi14++ {
            switch "foo bar" {
            case "a":
                if true {
                    var v15 int16 = 456
                    _ = v15
                }
            default:
                if v2 {
                    var v16 uint = 65535
                    if ((v3 * v3) < (v3 * 128)) {
                        switch (fi14 * (fi14 | (fi14 ^ (fi14 << (((^(v16))) & 63))))) {
                        case 32768:
                            v3 = 63
                        }
                        var v17 uint8 = v3
                        _ = v17
                    } else {
                        switch fi14 {
                        case 0:
                            fmt.Println(v4)
                        case 2147483647:
                            for fi18 := 0; fi18 < 6; fi18++ {
                                v16 -= v16
                                _ = fi18
                            }
                        case -553:
                            if (32767 >= -519) {
                                for fi19 := 0; fi19 < 13; fi19++ {
                                    v5 -= 359
                                    v4 = uint32(v16)
                                    _ = fi19
                                }
                                for fi20 := 0; fi20 < 3; fi20++ {
                                    switch fi20 {
                                    case -819:
                                        var v21 int16 = -32768
                                        _ = v21
                                    case -2147483648:
                                        v22 := GenS3{F0: -128}
                                        _ = v22
                                    case -9223372036854775808:
                                        var v23 float32 = -2.25
                                        _ = v23
                                    default:
                                        v24 := []int{}
                                        _ = v24
                                    }
                                    if (32767 == -32768) {
                                        v25 := []int{fi14, 4294967295, fi20}
                                        _ = v25
                                    } else {
                                        v26 := []int{}
                                        _ = v26
                                    }
                                    _ = fi20
                                }
                            }
                        default:
                            for fi27 := 0; fi27 < 20; fi27++ {
                                fmt.Println(v4, fi14)
                                v16 -= (v16 % ((v16) | 1))
                                _ = fi27
                            }
                        }
                        v4 += 920
                    }
                    var v28 int = (fi14 - -2147483648)
                    _ = v16
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 9; fi29++ {
                        v30 := GenS2{F0: "\n\t", F1: 0}
                        if v1 {
                            if (18446744073709551615 <= 4294967295) {
                                v31 := []int{fi14}
                                v32 := GenS2{F0: "\n\t", F1: -1}
                                _ = v31
                                _ = v32
                            }
                            v33 := []int{}
                            _ = v33
                        } else {
                            switch fi29 {
                            case 108:
                                v34 := GenS2{F0: "hello", F1: 1}
                                _ = v34
                            case 0:
                                if (!v1) {
                                    fmt.Println(fi14, fi29)
                                    fmt.Println(fi29, v13)
                                    if (!(!false)) {
                                        v35 := GenS4{F0: (false && (183 < (^(v5)))), F1: v3}
                                        _ = v35
                                    }
                                    v2 = (v4 <= 784)
                                }
                            case 9223372036854775807:
                                fmt.Println(v4, fi14)
                            }
                            var v36 uint16 = v13
                            _ = v36
                        }
                        if v1 {
                            fmt.Println(v4, v1, v2)
                            fmt.Println(v2, v4, v1)
                            v2 = (0 != 17)
                            if ((fi14 & fi14) == 32768) {
                                v37 := GenS5{F0: 1, F1: 563}
                                _ = v37
                            } else {
                                fmt.Println(v1, fi14, v1)
                                fmt.Println(v1)
                            }
                            var v38 uint8 = 128
                            fmt.Println(v38, v3, v2)
                            _ = v38
                        }
                        _ = fi29
                        _ = v30
                    }
                }
            }
            _ = fi14
        }
        _ = v13
    }
    v39 := []int{42, 726, -614, 209}
    v2 = (4294967295 <= 2)
    for fi40 := 0; fi40 < 18; fi40++ {
        if v1 {
            v41 := GenS6{F0: -212, F1: v4, F2: v4}
            var v42 float64 = 3.14159
            _ = v41
            _ = v42
        }
        _ = fi40
    }
    v2 = v1
    v1 = true
    v1 = (9223372036854775807 > -561)
    v1 = ("" != (("hello" + "") + "a"))
    if (!v2) {
        v43 := GenS3{F0: -128}
        if len(v39) > 0 {
            v39[0] = -1
        }
        _ = v43
    }
    if v1 {
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
        v3 -= (v3 / ((v3) | 1))
    }
    v5 = (v5 ^ v5)
    if len(v39) > 0 {
        v39[0] = -871
    }
    var v44 int8 = 42
    var v45 int16 = -189
    if (!v1) {
        fmt.Println(v5, v4, v4)
    } else {
        v4 -= (v4 * (v4 ^ uint32(v5)))
        switch "" {
        case "":
            _ = len(v39)
        case "\n\t":
            switch ("foo bar" + "a") {
            case "\n\t":
                _ = len(v39)
            case "foo bar":
                fmt.Println(v4, v44)
            case "hello":
                if len(v39) > 0 {
                    v39[0] = -1
                }
            default:
                fmt.Println(v45, v5)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v8
}
