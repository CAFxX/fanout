package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int8; F1 bool; F2 uint8 }
type GenS5 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 float64 = 1.5
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 float32 = ((0.0 - -2.25) * float32(float64((1e10 / 0.1))))
        var v6 uint = 0
        v7 := GenS1{F0: -32768, F1: (^(v6))}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    v1[-434] = -734
    if true {
        fmt.Println(v3)
        var v8 int32 = 823
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS1{F0: -32768, F1: 18446744073709551615}
            switch (fi9 * 42) {
            case -215:
                v11 := GenS2{F0: 0, F1: (false || ((42 > -9223372036854775808) && true)), F2: 1}
                _ = v11
            }
            var v12 int16 = 0
            _ = fi9
            _ = v10
            _ = v12
        }
        _ = v8
    } else {
        switch -1 {
        case -9223372036854775808:
            if (v3 == 0.1) {
                for fi13 := 0; fi13 < 2; fi13++ {
                    if ((fi13 >> ((65535) & 63)) >= fi13) {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            v3 += (v3 / v3)
                            _ = fi14
                        }
                    }
                    var v15 float32 = 0.1
                    _ = fi13
                    _ = v15
                }
                v2[4294967295] = "a"
                if false {
                    var v16 uint16 = 2
                    _ = v16
                }
            } else {
                v3 -= (1.5 + v3)
                fmt.Println(v3)
            }
        }
        v17 := GenS3{F0: false, F1: "foo bar"}
        v18 := GenS3{F0: false, F1: "\n\t"}
        _ = v17
        _ = v18
    }
    var v19 float64 = -2.25
    switch "" {
    case "foo bar":
        v20 := GenS1{F0: 32767, F1: 4294967295}
        _ = v20
    case "a":
        v21 := []int{-738, 211, -1, -1}
        _ = v21
    case "\n\t":
        for fi22 := 0; fi22 < 20; fi22++ {
            v3 += v19
            fmt.Println(fi22)
            _ = fi22
        }
    default:
        var v23 float32 = (((-2.25 - float32((float64(1.5) + -2.25))) / -2.25) * -2.25)
        _ = v23
    }
    var v24 int = 9223372036854775807
    if ((634 > 2) && false) {
        v25 := []int{v24, (v24 >> ((255) & 63)), v24, 2147483647}
        _ = v25
    } else {
        var v26 uint32 = 0
        v27 := GenS4{F0: 1, F1: false, F2: 13}
        v26 = (v26 + v26)
        _ = v26
        _ = v27
    }
    v24 = 4294967295
    var v28 int64 = 9223372036854775807
    if true {
        var v29 uint16 = 32767
        v3 += float64(3.14159)
        if true {
            if true {
                if false {
                    if (!(false || (153 == 396))) {
                        fmt.Println(v24)
                        _ = len(v1)
                    }
                    switch v24 {
                    case 128:
                        var v30 int32 = 0
                        _ = v30
                    }
                }
            } else {
                if ((false && true) || false) {
                    if true {
                        v29 = 0
                        var v31 uint8 = 127
                        fmt.Println(v19, v28)
                        _ = v31
                    }
                    v3 = v3
                    if false {
                        var v32 uint8 = 42
                        if (true && (v19 < v3)) {
                            switch "hello" {
                            case "":
                                v3 = v3
                            default:
                                fmt.Println(v3, v24, v28)
                            }
                            v28 -= (v28 % (((^(v28))) | 1))
                        }
                        v33 := v1[v24]
                        _ = v32
                        _ = v33
                    } else {
                        switch (("a" + "a") + ((("\n\t" + "") + "a") + "")) {
                        case "\n\t":
                            v24 = 1
                        case "foo bar":
                            for fi34 := 0; fi34 < 6; fi34++ {
                                v35 := v2[(^(fi34))]
                                _ = fi34
                                _ = v35
                            }
                        case "a":
                            v3 += v19
                        }
                        switch (v24 << ((18446744073709551615) & 63)) {
                        case -9223372036854775808:
                            for fi36 := 0; fi36 < 14; fi36++ {
                                var v37 uint64 = uint64(v29)
                                switch v24 {
                                case 1:
                                    fmt.Println(v3, v37)
                                case 32768:
                                    v29 -= v29
                                }
                                _ = fi36
                                _ = v37
                            }
                        }
                    }
                }
            }
        } else {
            if true {
                v38 := []int{(v24 ^ -9223372036854775808), (v24 * 32767), (v24 | v24)}
                v39 := GenS5{F0: 95, F1: 2147483648}
                _ = v38
                _ = v39
            } else {
                var v40 int = (-(v24))
                _ = v40
            }
            v41 := GenS4{F0: -32, F1: false, F2: 0}
            _ = v41
        }
        _ = v29
    }
    if false {
        var v42 uint32 = 0
        if (1000 < 847) {
            var v43 uint = 65535
            var v44 int8 = -32
            _ = v43
            _ = v44
        }
        var v45 float64 = (v3 / (v19 - v19))
        _ = v42
        _ = v45
    }
    v46 := []int{v24}
    var v47 uint16 = 32767
    for fi48 := 0; fi48 < 6; fi48++ {
        var v49 string = "hello"
        _ = fi48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v46
    _ = v47
}
