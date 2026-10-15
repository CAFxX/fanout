package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int8; F2 int }
type GenS5 struct { F0 uint8; F1 uint8; F2 string }

func main() {
    var v1 int = -9223372036854775808
    v2 := []int{(^(v1)), v1, v1, 1}
    var v3 string = ("" + ("foo bar" + "\n\t"))
    var v4 int64 = -464
    var v5 uint32 = 0
    switch v3 {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            for fi6 := 0; fi6 < 11; fi6++ {
                if true {
                    v3 = v3
                    if false {
                        v7 := GenS1{F0: "foo bar"}
                        fmt.Println(v1, v3, v5)
                        _ = v7
                    }
                    v3 = v3
                } else {
                    v3 = "foo bar"
                    v4 -= v4
                    switch 0 {
                    case 1:
                        fmt.Println(v4, fi6)
                    default:
                        v5 -= (v5 ^ v5)
                    }
                }
                if (-782 < v4) {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        v5 = 2147483648
                        switch fi8 {
                        case 996:
                            v9 := GenS1{F0: (v3 + (v3 + v3))}
                            _ = v9
                        case 1:
                            fmt.Println(v4, v3, v3)
                        case 4294967295:
                            var v10 uint8 = 252
                            _ = v10
                        default:
                            fmt.Println(v3)
                        }
                        if false {
                            v1 = fi6
                            v3 = (("hello" + v3) + v3)
                            v2 = append(v2, fi6)
                        }
                        _ = fi8
                    }
                    for fi11 := 0; fi11 < 16; fi11++ {
                        var v12 int16 = 396
                        v5 -= (v5 | (v5 << ((255) & 63)))
                        _ = fi11
                        _ = v12
                    }
                    var v13 uint16 = 42
                    _ = v13
                } else {
                    if len(v2) > 0 {
                        v2[0] = 127
                    }
                    var v14 string = v3
                    _ = v14
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 uint = 65535
                    _ = fi15
                    _ = v16
                }
                _ = fi6
            }
        }
    default:
        v1 = (v1 / ((v1) | 1))
    }
    var v17 string = (v3 + v3)
    v5 += v5
    fmt.Println(v5, v17)
    v17 = v3
    v18 := GenS1{F0: v17}
    v3 = v3
    var v19 int = v1
    for fi20 := 0; fi20 < 7; fi20++ {
        fmt.Println(v17, v19)
        _ = fi20
    }
    fmt.Println(v17)
    v21 := GenS2{F0: 418, F1: int(v4)}
    v17 = (v17 + "a")
    switch "a" {
    case "a":
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 int8 = 0
            var v24 int16 = 1
            _ = fi22
            _ = v23
            _ = v24
        }
    case "":
        for fi25 := 0; fi25 < 5; fi25++ {
            if (32767 >= 908) {
                for fi26 := 0; fi26 < 5; fi26++ {
                    var v27 int16 = 0
                    if (1 == -128) {
                        v28 := GenS3{F0: 0}
                        v27 = (v27 % ((int16(fi26)) | 1))
                        fmt.Println(v3)
                        _ = v28
                    }
                    _ = fi26
                    _ = v27
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                for fi29 := 0; fi29 < 6; fi29++ {
                    switch (-(fi29)) {
                    case 42:
                        if (!(1 < 715)) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v17 = (v3 + "\n\t")
                        }
                    case 4294967295:
                        var v30 uint = 475
                        _ = v30
                    case 9223372036854775807:
                        if false {
                            if false {
                                var v31 float32 = 3.14159
                                if false {
                                    v32 := GenS1{F0: (v17 + "foo bar")}
                                    var v33 int64 = v4
                                    _ = v32
                                    _ = v33
                                }
                                fmt.Println(v3, v1)
                                _ = v31
                            } else {
                                v34 := GenS4{F0: 1, F1: -128, F2: fi29}
                                v5 -= 23
                                v35 := GenS5{F0: 128, F1: 194, F2: v17}
                                _ = v34
                                _ = v35
                            }
                            if (true && (7 < 42)) {
                                v19 = v19
                                var v36 float64 = -2.25
                                _ = v36
                            }
                        } else {
                            fmt.Println(v1)
                        }
                    }
                    if (!((3.14159 + -2.25) >= float64(1.5))) {
                        v17 = "foo bar"
                        v2 = append(v2, (v1 - 146))
                    } else {
                        var v37 int64 = v4
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                        _ = v37
                    }
                    _ = fi29
                }
                fmt.Println(v17, v3, v3)
            }
            _ = fi25
        }
    }
    var v38 uint32 = v5
    if (!false) {
        var v39 int = v19
        if (false && false) {
            var v40 int32 = -1
            if true {
                if false {
                    for fi41 := 0; fi41 < 5; fi41++ {
                        fmt.Println(v5)
                        v3 = v3
                        var v42 uint16 = 227
                        _ = fi41
                        _ = v42
                    }
                }
                var v43 bool = true
                _ = v43
            }
            _ = v40
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = v39
    }
    var v44 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v5
}
