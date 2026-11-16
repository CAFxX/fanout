package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 int64; F1 int64; F2 uint8 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 uint64 = 0
    var v2 float32 = 1.5
    var v3 int16 = -32768
    v4 := make(map[uint32]bool)
    var v5 string = ("a" + ("a" + "\n\t"))
    var v6 int32 = 42
    switch "a" {
    case "hello":
        v2 = (v2 / float32(((0.1 * (0.1 - 3.14159)) / 1.0)))
    }
    v7 := GenS1{F0: 42}
    if false {
        v2 = -2.25
        if true {
            if (2 >= 65535) {
                if true {
                    if false {
                        var v8 uint64 = v1
                        v2 = v2
                        if true {
                            v5 = v5
                            if false {
                                delete(v4, 4294967295)
                                fmt.Println(v8, v8, v3)
                                if false {
                                    v9 := []int{-361, -9223372036854775808, -493}
                                    if (false || false) {
                                        for fi10 := 0; fi10 < 2; fi10++ {
                                            if ((((v2 + v2) / (((-2.25 / v2) * -2.25) / (v2 - v2))) <= 0.1) || true) {
                                                var v11 float32 = float32(-2.25)
                                                _ = v11
                                            }
                                            _ = fi10
                                        }
                                        var v12 int16 = v3
                                        v4[65535] = true
                                        _ = v12
                                    }
                                    v5 = "a"
                                    _ = v9
                                }
                            }
                            if (4294967295 != 42) {
                                if true {
                                    switch 0 {
                                    case 32767:
                                        fmt.Println(v8, v8)
                                    default:
                                        var v13 float64 = 1.5
                                        _ = v13
                                    }
                                }
                                switch v5 {
                                case "a":
                                    v4[2147483648] = (false && (661 != 65535))
                                case "\n\t":
                                    delete(v4, 4294967295)
                                case "hello":
                                    v14 := GenS1{F0: 255}
                                    _ = v14
                                default:
                                    var v15 float32 = 1e10
                                    _ = v15
                                }
                                var v16 int = 1
                                _ = v16
                            } else {
                                var v17 uint64 = 421
                                v18 := GenS2{F0: 42, F1: -83, F2: v3}
                                _ = v17
                                _ = v18
                            }
                        } else {
                            v19 := v4[0]
                            switch v5 {
                            case "a":
                                v20 := GenS1{F0: 255}
                                _ = v20
                            default:
                                switch "" {
                                case "a":
                                    var v21 uint8 = uint8(v8)
                                    _ = v21
                                case "foo bar":
                                    for fi22 := 0; fi22 < 9; fi22++ {
                                        var v23 int16 = (v3 + (v3 - 1000))
                                        _ = fi22
                                        _ = v23
                                    }
                                }
                            }
                            var v24 int = 4294967295
                            _ = v19
                            _ = v24
                        }
                        _ = v8
                    }
                    fmt.Println(v2)
                } else {
                    fmt.Println(v5)
                }
            } else {
                var v25 uint64 = (v1 | (^(v1)))
                v26 := GenS3{F0: -731, F1: 335, F2: 0}
                v27 := GenS3{F0: -9223372036854775808, F1: 9223372036854775807, F2: 74}
                _ = v25
                _ = v26
                _ = v27
            }
        } else {
            v6 += -2147483648
            v28 := []int{-173}
            for fi29 := 0; fi29 < 20; fi29++ {
                for fi30 := 0; fi30 < 3; fi30++ {
                    switch (fi29 / ((fi30) | 1)) {
                    case 42:
                        v3 = v3
                    }
                    delete(v4, 2147483647)
                    if true {
                        v6 = -2147483648
                        var v31 int16 = v3
                        _ = v31
                    } else {
                        v32 := GenS1{F0: 156}
                        fmt.Println(v6, v2)
                        if false {
                            if false {
                                v4[2147483647] = false
                                v33 := GenS4{F0: v6, F1: int(v6)}
                                var v34 bool = false
                                _ = v33
                                _ = v34
                            } else {
                                v35 := GenS5{F0: 2}
                                v5 = v5
                                switch "hello" {
                                case "":
                                    var v36 float32 = v2
                                    _ = v36
                                case "hello":
                                    v5 = "foo bar"
                                case "a":
                                    if ((v6 | (v6 % ((v6) | 1))) != v6) {
                                        fmt.Println(fi29, v3, v6)
                                    } else {
                                        v6 -= v6
                                        v5 = v5
                                    }
                                }
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                                _ = v35
                            }
                        } else {
                            if true {
                                var v37 bool = ((false && (v6 < v6)) || (1 > -127))
                                var v38 int64 = 42
                                var v39 uint8 = 201
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                        }
                        _ = v32
                    }
                    _ = fi30
                }
                _ = fi29
            }
            _ = v28
        }
        switch 1 {
        case -1:
            var v40 uint16 = 32767
            _ = v40
        case 127:
            fmt.Println(v6, v3)
        case 32767:
            v5 = v5
        default:
            v4[uint32(v1)] = (4294967295 != 1)
        }
    } else {
        fmt.Println(v3, v5)
        fmt.Println(v5)
    }
    v41 := []int{-655, 2147483647, -200, 0}
    switch 4294967295 {
    case -894:
        fmt.Println(v5, v1, v6)
    case -792:
        v2 = ((float32(-2.25) / v2) / (v2 / 0.1))
    case -287:
        fmt.Println(v2, v3, v2)
    default:
        var v42 uint16 = 42
        _ = v42
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        if len(v41) >= 2 {
            v41 = v41[1:len(v41)]
        }
        v44 := GenS5{F0: 85}
        var v45 uint64 = (^(v1))
        _ = fi43
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 14; fi46++ {
        v3 = v3
        _ = fi46
    }
    v4[1] = true
    var v47 float64 = 3.14159
    var v48 string = (v5 + v5)
    for fi49 := 0; fi49 < 14; fi49++ {
        v47 -= (v47 * 3.14159)
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
