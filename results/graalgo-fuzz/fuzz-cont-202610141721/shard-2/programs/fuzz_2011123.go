package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int; F1 bool; F2 int32 }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{65535, 65535}
    var v3 int16 = 1
    var v4 int64 = -813
    var v5 uint8 = 0
    fmt.Println(v4, v3)
    v5 += (v5 * v5)
    v6 := GenS1{F0: 164, F1: 731}
    var v7 bool = ((!(v5 >= v5)) && false)
    v8 := GenS1{F0: 489, F1: 255}
    v7 = v7
    if true {
        v7 = (v7 && ((42 == 65535) || v7))
        fmt.Println(v4, v4)
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: 65535, F1: 18446744073709551615}
        if v7 {
            v7 = v7
            var v11 uint32 = 4294967295
            _ = v11
        } else {
            switch -964 {
            case -158:
                for fi12 := 0; fi12 < 7; fi12++ {
                    var v13 uint64 = 42
                    switch "\n\t" {
                    case "":
                        var v14 uint = 96
                        _ = v14
                    default:
                        v15 := v1[42]
                        _ = v15
                    }
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    _ = fi12
                    _ = v13
                }
            case 397:
                if v7 {
                    var v16 int64 = (v4 << ((210) & 63))
                    fmt.Println(v16, v4, v5)
                    switch "a" {
                    case "":
                        if true {
                            for fi17 := 0; fi17 < 15; fi17++ {
                                _ = len(v2)
                                v18 := GenS2{F0: v4, F1: v4}
                                v16 -= -1
                                _ = fi17
                                _ = v18
                            }
                            v4 -= -9223372036854775808
                        } else {
                            if false {
                                v19 := GenS3{F0: 4294967295, F1: -522}
                                _ = v19
                            }
                        }
                    }
                    _ = v16
                } else {
                    if false {
                        fmt.Println(v3)
                    } else {
                        v20 := GenS4{F0: 77, F1: (v3 & (v3 + (v3 ^ (v3 ^ (-(v3))))))}
                        var v21 float32 = 1.5
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(v7, fi9, v7)
                }
            }
            v22 := v1[42]
            _ = v22
        }
        v23 := GenS5{F0: -2147483648, F1: ((3.14159 > (3.14159 - float64(3.14159))) && true), F2: 1}
        _ = fi9
        _ = v10
        _ = v23
    }
    switch 0 {
    case 2147483647:
        v4 -= (^(v4))
    case 127:
        var v24 uint32 = 2147483647
        _ = v24
    default:
        v7 = (v3 <= v3)
    }
    var v25 uint64 = uint64(v5)
    var v26 int16 = (^(v3))
    if v7 {
        switch 1 {
        case 127:
            var v27 int = -9223372036854775808
            _ = v27
        case 128:
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 uint8 = (^(v5))
                _ = fi28
                _ = v29
            }
        case -536:
            if false {
                var v30 string = "\n\t"
                var v31 uint16 = 1
                for fi32 := 0; fi32 < 2; fi32++ {
                    _ = len(v1)
                    _ = fi32
                }
                _ = v30
                _ = v31
            }
        }
    }
    var v33 uint16 = 65535
    fmt.Println(v7, v7, v4)
    v26 -= -1
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 20; fi35++ {
            fmt.Println(fi35, v26)
            var v36 bool = v7
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 7; fi37++ {
            var v38 int32 = -2147483648
            _ = fi37
            _ = v38
        }
        _ = fi34
    }
    if (2147483647 < 1) {
        fmt.Println(v4, v33)
        v39 := GenS6{F0: (v3 % ((v3) | 1))}
        v33 -= v33
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
