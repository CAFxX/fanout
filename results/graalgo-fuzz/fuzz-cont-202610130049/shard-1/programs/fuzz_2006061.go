package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 1e10
    var v3 int8 = 42
    var v4 uint = 65535
    var v5 int16 = -1
    var v6 uint32 = 1
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 12; fi8++ {
        v4 = 65535
        _ = fi8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        if false {
            v1[v7] = true
        }
        switch v7 {
        case "\n\t":
            if (0.0 != (0.0 / -2.25)) {
                if false {
                    if false {
                        v10 := GenS1{F0: 1, F1: (v4 ^ 780), F2: v4}
                        v11 := GenS2{F0: 1, F1: v4}
                        _ = v10
                        _ = v11
                    } else {
                        if true {
                            var v12 int8 = (v3 | v3)
                            v13 := GenS1{F0: 409, F1: 646, F2: (^(v4))}
                            _ = v12
                            _ = v13
                        } else {
                            if (!true) {
                                var v14 string = "hello"
                                var v15 uint32 = (v6 & v6)
                                _ = v14
                                _ = v15
                            }
                            switch "\n\t" {
                            case "foo bar":
                                v16 := GenS1{F0: 32767, F1: v4, F2: (v4 << (((v4 * v4)) & 63))}
                                _ = v16
                            case "\n\t":
                                switch fi9 {
                                case 4294967295:
                                    var v17 uint64 = 628
                                    _ = v17
                                case 32768:
                                    v4 = (v4 / ((v4) | 1))
                                }
                            case "a":
                                delete(v1, v7)
                            }
                            var v18 int64 = 0
                            _ = v18
                        }
                        if (!(true || false)) {
                            for fi19 := 0; fi19 < 15; fi19++ {
                                v4 += (v4 + v4)
                                v20 := []int{int(v5)}
                                var v21 int64 = 772
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            v4 -= (v4 % ((v4) | 1))
                            v3 = v3
                        }
                        var v22 uint8 = 127
                        _ = v22
                    }
                    fmt.Println(v6, v2, v3)
                } else {
                    for fi23 := 0; fi23 < 19; fi23++ {
                        var v24 int64 = -9223372036854775808
                        _ = fi23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 12; fi25++ {
                        v26 := []int{fi9}
                        switch 32767 {
                        case -9223372036854775808:
                            var v27 int8 = (v3 ^ -128)
                            _ = v27
                        case 32768:
                            fmt.Println(fi25)
                        }
                        _ = fi25
                        _ = v26
                    }
                    if true {
                        switch (v7 + v7) {
                        case "hello":
                            v6 -= (v6 & (v6 << ((v4) & 63)))
                        default:
                            v28 := v1[v7]
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 4; fi29++ {
                            v30 := []int{0, fi29, (^(fi9)), fi9}
                            v2 = (((v2 / 1e10) * v2) * v2)
                            _ = fi29
                            _ = v30
                        }
                        if ((401 != (v6 << ((v4) & 63))) && false) {
                            for fi31 := 0; fi31 < 4; fi31++ {
                                delete(v1, v7)
                                var v32 bool = false
                                fmt.Println(v7, fi31, v5)
                                _ = fi31
                                _ = v32
                            }
                        } else {
                            v5 = 97
                            if true {
                                v3 = (v3 | (v3 / ((v3) | 1)))
                                v33 := []int{}
                                _ = v33
                            }
                            var v34 bool = (v4 < (^(v4)))
                            _ = v34
                        }
                    }
                }
            }
        case "a":
            v2 = 1e10
        case "hello":
            fmt.Println(v2, v3)
        }
        _ = fi9
    }
    if (255 <= 1) {
        if ("a" == v7) {
            var v35 string = ("hello" + v7)
            switch (v7 + v7) {
            case "foo bar":
                if (0 != 32767) {
                    var v36 string = "\n\t"
                    _ = v36
                } else {
                    _ = len(v1)
                }
            case "":
                for fi37 := 0; fi37 < 1; fi37++ {
                    v7 = "\n\t"
                    _ = fi37
                }
            }
            var v38 bool = (false || ((false || false) && true))
            _ = v35
            _ = v38
        }
        fmt.Println(v4, v2)
        _ = len(v1)
    }
    _ = len(v1)
    for fi39 := 0; fi39 < 15; fi39++ {
        v40 := GenS3{F0: (v5 & v5)}
        _ = fi39
        _ = v40
    }
    if true {
        fmt.Println(v7)
    }
    v41 := GenS2{F0: -826, F1: v4}
    var v42 uint8 = 230
    v43 := GenS4{F0: 435, F1: 662, F2: uint16(v42)}
    fmt.Println(v2, v7, v4)
    v44 := []int{int(v4), 4294967295, 320}
    v45 := GenS5{F0: "foo bar"}
    switch -632 {
    case 127:
        switch v7 {
        case "foo bar":
            v46 := GenS6{F0: (v6 & (^(v6))), F1: int(v42)}
            _ = v46
        default:
            v42 -= uint8(v4)
        }
    case 305:
        v2 = -2.25
    default:
        v6 += v6
    }
    v5 = (^(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
