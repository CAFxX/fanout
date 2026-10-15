package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 uint }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint; F1 uint64 }

func main() {
    var v1 int32 = -1
    var v2 uint64 = 18446744073709551615
    var v3 int8 = 1
    var v4 int8 = v3
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int16 = 39
        v8 := GenS1{F0: (v7 ^ (v7 | v7))}
        _ = fi6
        _ = v7
        _ = v8
    }
    v4 -= v4
    if false {
        v9 := GenS1{F0: -782}
        _ = v9
    } else {
        v10 := []int{32767, -1}
        v11 := GenS1{F0: 264}
        v12 := GenS1{F0: 949}
        _ = v10
        _ = v11
        _ = v12
    }
    if ((v4 + v4) >= v4) {
        var v13 float64 = 0.1
        v2 -= v2
        _ = v13
    }
    switch (("\n\t" + "a") + "") {
    case "\n\t":
        if true {
            _ = len(v5)
        }
    case "hello":
        v14 := GenS1{F0: -1}
        _ = v14
    case "":
        fmt.Println(v2, v4, v1)
    default:
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint = 18446744073709551615
            v17 := GenS2{F0: 1}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    if (!false) {
        for fi18 := 0; fi18 < 8; fi18++ {
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS3{F0: false, F1: -671, F2: 2}
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        v21 := GenS1{F0: 289}
        _ = v21
    }
    var v22 uint64 = v2
    for fi23 := 0; fi23 < 7; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 int8 = 1
            _ = fi24
            _ = v25
        }
        if false {
            for fi26 := 0; fi26 < 20; fi26++ {
                v3 = v3
                if (42 > 42) {
                    v27 := []int{(fi23 + (-(fi23))), fi26}
                    v27 = append(v27, -9223372036854775808)
                    _ = v27
                } else {
                    if (("a" > "\n\t") && true) {
                        v28 := []int{(fi26 % ((fi26) | 1)), fi23, (fi26 - fi26), (fi26 - fi26)}
                        v1 -= int32(v2)
                        _ = v28
                    }
                }
                _ = fi26
            }
            fmt.Println(fi23)
        } else {
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 int16 = 42
                if (!false) {
                    v31 := GenS3{F0: false, F1: fi23, F2: 42}
                    v1 = v1
                    _ = v31
                }
                _ = fi29
                _ = v30
            }
            for fi32 := 0; fi32 < 6; fi32++ {
                switch "foo bar" {
                case "\n\t":
                    v4 -= -84
                case "":
                    v33 := GenS4{F0: 0, F1: v4}
                    _ = v33
                default:
                    if false {
                        if false {
                            v4 -= (v4 ^ (v4 % ((v4) | 1)))
                        }
                    }
                }
                if (330 != 65535) {
                    if false {
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                var v34 bool = (fi23 != 128)
                                fmt.Println(fi32, fi23)
                                v35 := GenS5{F0: 1, F1: (("hello" + "") + ("a" + ("\n\t" + "foo bar")))}
                                _ = v34
                                _ = v35
                            }
                        }
                        fmt.Println(v22)
                    } else {
                        if (true && (!(4294967295 >= 2147483647))) {
                            var v36 uint = 2
                            _ = v36
                        }
                    }
                    v37 := GenS5{F0: 0, F1: ("" + (("hello" + "a") + "hello"))}
                    switch "" {
                    case "a":
                        if true {
                            switch -9223372036854775808 {
                            case 9223372036854775807:
                                var v38 float64 = ((float64(3.14159) * (1e10 + 1.5)) / 3.14159)
                                _ = v38
                            default:
                                if false {
                                    var v39 float64 = 1.5
                                    v40 := GenS6{F0: 255, F1: v22}
                                    _ = v39
                                    _ = v40
                                }
                            }
                            v41 := []int{fi23, (fi32 * fi23)}
                            _ = v41
                        } else {
                            v42 := GenS2{F0: 215}
                            _ = v42
                        }
                    case "":
                        switch fi23 {
                        case 466:
                            var v43 int32 = int32(fi32)
                            _ = v43
                        }
                    }
                    _ = v37
                } else {
                    fmt.Println(fi32)
                }
                _ = fi32
            }
        }
        _ = fi23
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        for fi45 := 0; fi45 < 12; fi45++ {
            var v46 uint32 = 1
            if (128 > 128) {
                var v47 int16 = 32767
                if (false || false) {
                    v2 = 1
                    fmt.Println(fi44, v4, v1)
                } else {
                    v22 = v2
                }
                var v48 uint64 = 0
                var v49 uint16 = 2
                var v50 int8 = (-(v4))
                _ = v47
                _ = v48
                _ = v49
                _ = v50
            }
            _ = fi45
            _ = v46
        }
        var v51 string = "a"
        _ = fi44
        _ = v51
    }
    v52 := []int{32768, -9223372036854775808, 32767, int(v3)}
    var v53 int8 = v3
    var v54 uint = uint(v2)
    v22 = (v22 - v2)
    for fi55 := 0; fi55 < 13; fi55++ {
        var v56 int64 = 447
        switch fi55 {
        case 4294967295:
            for fi57 := 0; fi57 < 1; fi57++ {
                var v58 int16 = -1
                var v59 uint8 = 74
                _ = len(v52)
                _ = len(v5)
                var v60 int8 = (-(v53))
                v4 -= v53
                _ = fi57
                _ = v58
                _ = v59
                _ = v60
            }
        }
        _ = fi55
        _ = v56
    }
    _ = len(v52)
    var v61 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v61
}
