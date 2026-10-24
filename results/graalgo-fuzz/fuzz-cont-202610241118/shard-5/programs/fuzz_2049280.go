package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int32; F1 int32; F2 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 bool; F1 int64; F2 uint64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64 }
func genFn0() uint32 {
    var v6 uint16 = 1
    v7 := GenS1{F0: 939, F1: v6, F2: (v6 + 489)}
    v8 := GenS1{F0: 208, F1: 342, F2: (^(v6))}
    if false {
        var v9 bool = (260 < 65535)
        fmt.Println(v9)
        v10 := GenS2{F0: 1, F1: v9}
        _ = v10
        _ = v9
    } else {
        v6 += 2
    }
    var v11 float64 = 0.0
    _ = v11
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1() int {
    var v12 uint16 = 255
    v12 -= v12
    v12 = (v12 ^ v12)
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int64 = -9223372036854775808
        var v15 float64 = (0.1 * float64(0.1))
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v12
    return 127
}

func main() {
    v1 := []int{-2147483648, 128}
    v2 := []uint32{4294967295, 4294967295}
    var v3 int64 = 42
    var v4 int8 = -128
    var v5 int16 = 42
    v3 = (v3 << ((1) & 63))
    var v16 uint8 = 2
    _ = len(v2)
    var v17 uint32 = 42
    var v18 float64 = 1e10
    var v19 float64 = v18
    if true {
        if (false && (2 >= (v17 - genFn0()))) {
            fmt.Println(v17, v18, v4)
            _ = len(v1)
        } else {
            v18 -= v18
            fmt.Println(v5)
            if (!(!(v4 >= -1))) {
                for fi20 := 0; fi20 < 4; fi20++ {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        if false {
                            v18 = ((float64((1.5 / 1.0)) / 1.5) / v18)
                        } else {
                            v22 := GenS1{F0: -1, F1: 1, F2: 1}
                            _ = v22
                        }
                        var v23 float32 = 0.0
                        if true {
                            v17 -= 2147483648
                            var v24 float32 = v23
                            _ = v24
                        }
                        _ = fi21
                        _ = v23
                    }
                    v4 = v4
                    _ = fi20
                }
                switch "a" {
                case "hello":
                    v1 = append(v1, -2147483648)
                case "foo bar":
                    if true {
                        v25 := GenS3{F0: v16, F1: 740}
                        v26 := GenS4{F0: 384, F1: 1}
                        _ = v25
                        _ = v26
                    }
                }
            }
        }
        v27 := GenS5{F0: "\n\t", F1: (v4 & (v4 | int8(v17)))}
        for fi28 := 0; fi28 < 20; fi28++ {
            if false {
                for fi29 := 0; fi29 < 17; fi29++ {
                    v30 := GenS6{F0: 766, F1: 846, F2: (fi29 >> ((255) & 63))}
                    var v31 float32 = (float32((0.1 / float64(3.14159))) + 0.1)
                    v19 = (3.14159 - v19)
                    switch (fi29 >> ((65535) & 63)) {
                    case 2147483647:
                        if true {
                            var v32 float64 = v19
                            if false {
                                _ = len(v2)
                                var v33 uint16 = 42
                                v34 := GenS7{F0: 237}
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(v31, v3, v19)
                            _ = v32
                        } else {
                            v35 := GenS1{F0: (v5 ^ (^(v5))), F1: 569, F2: 42}
                            switch "a" {
                            case "hello":
                                v36 := GenS8{F0: v17, F1: v5}
                                _ = v36
                            case "a":
                                v3 += (v3 | 0)
                            default:
                                switch -493 {
                                case -342:
                                    v37 := GenS9{F0: true, F1: v3, F2: 259}
                                    _ = v37
                                }
                            }
                            _ = v35
                        }
                    case 42:
                        switch "foo bar" {
                        case "foo bar":
                            var v38 uint = 65535
                            _ = v38
                        case "a":
                            v39 := GenS1{F0: 1000, F1: 409, F2: 4}
                            _ = v39
                        default:
                            var v40 uint8 = 1
                            _ = v40
                        }
                    }
                    switch (-(fi28)) {
                    case 9223372036854775807:
                        v41 := GenS10{F0: v16}
                        _ = v41
                    case 345:
                        v3 = v3
                    case 741:
                        v18 += v19
                    default:
                        var v42 uint = 2
                        _ = v42
                    }
                    fmt.Println(v5)
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                for fi43 := 0; fi43 < 11; fi43++ {
                    v44 := GenS11{F0: 0}
                    v3 += 42
                    _ = fi43
                    _ = v44
                }
                v5 = 32767
            } else {
                for fi45 := 0; fi45 < 17; fi45++ {
                    v5 -= (v5 | v5)
                    _ = fi45
                }
            }
            if false {
                for fi46 := 0; fi46 < 1; fi46++ {
                    var v47 int8 = 127
                    var v48 bool = false
                    var v49 int = 42
                    _ = fi46
                    _ = v47
                    _ = v48
                    _ = v49
                }
            }
            var v50 string = "\n\t"
            _ = fi28
            _ = v50
        }
        _ = v27
    }
    var v51 float64 = float64(((3.14159 - 0.1) - 1e10))
    switch -1 {
    case 32767:
        v3 -= v3
    case 2147483647:
        for fi52 := 0; fi52 < 1; fi52++ {
            for fi53 := 0; fi53 < 7; fi53++ {
                if (false || (116 < 32767)) {
                    var v54 int64 = 1
                    switch "" {
                    case "a":
                        _ = len(v2)
                    case "hello":
                        switch "foo bar" {
                        case "\n\t":
                            v51 += 1e10
                        case "":
                            var v55 bool = (false && true)
                            _ = v55
                        case "hello":
                            var v56 int64 = (v3 >> ((42) & 63))
                            _ = v56
                        }
                    default:
                        v18 = -2.25
                    }
                    fmt.Println(v3, fi52)
                    _ = v54
                } else {
                    v51 -= 1.5
                }
                fmt.Println(v5)
                var v57 float64 = (v18 - v19)
                _ = fi53
                _ = v57
            }
            _ = len(v2)
            v58 := GenS7{F0: 261}
            _ = fi52
            _ = v58
        }
    case -26:
        var v59 float64 = (v19 + float64(((((float32(v19) * 1.5) / -2.25) + 3.14159) + 3.14159)))
        _ = v59
    default:
        v51 = (0.0 + 3.14159)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
