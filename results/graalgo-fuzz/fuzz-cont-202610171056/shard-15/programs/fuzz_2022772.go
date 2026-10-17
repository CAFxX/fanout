package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 int32; F1 int; F2 uint }
type GenS5 struct { F0 bool; F1 int; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64 }
func genFn0() int64 {
    v5 := GenS1{F0: -1, F1: 32767, F2: -2147483648}
    if ((255 > 140) && true) {
        var v6 uint16 = 1
        var v7 float64 = 0.0
        v8 := GenS1{F0: 0, F1: v6, F2: 2147483647}
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 int64 = 42
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 255, F2: -267}
    var v11 int8 = 100
    _ = v10
    _ = v11
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 int32 = -548
    v2 := make(map[uint32]int)
    var v3 uint64 = 0
    var v4 bool = false
    v12 := GenS1{F0: 42, F1: 65535, F2: v1}
    v13 := GenS2{F0: 544}
    var v14 int64 = -219
    for fi15 := 0; fi15 < 4; fi15++ {
        if (!v4) {
            v4 = (v4 || true)
        } else {
            v3 += 1
            switch fi15 {
            case -9223372036854775808:
                v4 = v4
            default:
                switch (-(fi15)) {
                case 127:
                    v3 = v3
                case -1:
                    fmt.Println(v1, v3, v3)
                case 128:
                    v16 := GenS3{F0: fi15, F1: "foo bar", F2: 42}
                    _ = v16
                default:
                    fmt.Println(v1)
                }
            }
            v4 = true
        }
        var v17 bool = v4
        v17 = (785 <= 1)
        _ = fi15
        _ = v17
    }
    fmt.Println(v3, v3)
    v1 = v1
    switch "foo bar" {
    case "foo bar":
        var v18 uint8 = 128
        _ = v18
    case "\n\t":
        for fi19 := 0; fi19 < 17; fi19++ {
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint16 = 32767
                delete(v2, 274)
                v22 := GenS4{F0: -2147483648, F1: fi20, F2: 763}
                _ = fi20
                _ = v21
                _ = v22
            }
            if (3.14159 == ((3.14159 + 1.5) * 0.1)) {
                if v4 {
                    switch "" {
                    case "":
                        v1 = (v1 - v1)
                    default:
                        v1 -= v1
                    }
                } else {
                    if (740 <= 18446744073709551615) {
                        var v23 int8 = -26
                        v4 = ((v4 && (false || false)) || (!false))
                        var v24 int32 = (v1 ^ 42)
                        _ = v23
                        _ = v24
                    }
                }
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        v25 := GenS3{F0: (fi19 - fi19), F1: "foo bar", F2: -1}
                        _ = v25
                    case "":
                        if true {
                            switch ("foo bar" + "\n\t") {
                            case "a":
                                v4 = v4
                            }
                            for fi26 := 0; fi26 < 11; fi26++ {
                                fmt.Println(v3, v3)
                                if v4 {
                                    var v27 uint32 = 452
                                    var v28 bool = v4
                                    v27 = (v27 + 549)
                                    fmt.Println(v3, v1)
                                    v29 := []int{-228}
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                } else {
                                    v1 += -1
                                    var v30 uint = 4294967295
                                    fmt.Println(v4, v4, v14)
                                    _ = v30
                                }
                                v31 := GenS4{F0: v1, F1: (fi26 + (fi26 << ((0) & 63))), F2: 1}
                                _ = fi26
                                _ = v31
                            }
                            var v32 uint16 = 65535
                            _ = v32
                        }
                    default:
                        v14 -= -1
                    }
                    if v4 {
                        v4 = v4
                        var v33 uint32 = 42
                        v34 := GenS5{F0: v4, F1: (fi19 % ((fi19) | 1)), F2: 255}
                        _ = v33
                        _ = v34
                    } else {
                        if v4 {
                            fmt.Println(fi19, v3, fi19)
                        } else {
                            var v35 bool = v4
                            _ = v35
                        }
                        var v36 int64 = 42
                        _ = v36
                    }
                    if v4 {
                        for fi37 := 0; fi37 < 5; fi37++ {
                            switch "foo bar" {
                            case "a":
                                v14 -= v14
                            }
                            switch "hello" {
                            case "hello":
                                var v38 float32 = float32(0.0)
                                _ = v38
                            case "foo bar":
                                if (genFn0() == genFn0()) {
                                    v14 = (v14 | (v14 << ((364) & 63)))
                                } else {
                                    switch 4294967295 {
                                    case 684:
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (v14 <= (v14 / ((v14) | 1))) {
                                                var v39 bool = (genFn0() > v14)
                                                _ = v39
                                            }
                                        case "hello":
                                            var v40 int32 = (v1 >> ((2) & 63))
                                            _ = v40
                                        default:
                                            v41 := GenS1{F0: (v14 ^ (v14 ^ genFn0())), F1: 65535, F2: v1}
                                            _ = v41
                                        }
                                    default:
                                        v14 += 42
                                    }
                                }
                            default:
                                v42 := GenS6{F0: (v1 % ((v1) | 1))}
                                _ = v42
                            }
                            _ = fi37
                        }
                        v14 += v14
                    }
                } else {
                    for fi43 := 0; fi43 < 13; fi43++ {
                        var v44 uint16 = 2
                        var v45 uint8 = 2
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 10; fi46++ {
                        fmt.Println(fi46, v3, v3)
                        var v47 uint16 = 2
                        delete(v2, 2)
                        _ = fi46
                        _ = v47
                    }
                }
            } else {
                delete(v2, 2147483647)
            }
            _ = fi19
        }
    case "hello":
        v48 := []int{}
        _ = v48
    }
    v49 := GenS7{F0: 52}
    var v50 bool = (255 != 65535)
    if v50 {
        fmt.Println(v1)
        fmt.Println(v14)
    }
    var v51 int16 = 42
    v52 := []int{4294967295}
    for fi53 := 0; fi53 < 15; fi53++ {
        _ = len(v52)
        _ = fi53
    }
    var v54 uint32 = 2147483648
    v50 = false
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v54
}
