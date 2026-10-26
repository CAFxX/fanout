package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 bool; F1 uint64; F2 uint }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int16; F1 string; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 7
    var v3 bool = (((1e10 - -2.25) + -2.25) < 1e10)
    v4 := make(map[string]uint32)
    var v5 uint = 65535
    v3 = (1e10 != float64((float32(1e10) / 1.0)))
    var v6 float64 = 3.14159
    v7 := GenS1{F0: 32767, F1: v5}
    var v8 int32 = -1
    switch -281 {
    case -9223372036854775808:
        var v9 int = 42
        _ = v9
    case -2147483648:
        switch -171 {
        case 4294967295:
            for fi10 := 0; fi10 < 4; fi10++ {
                switch fi10 {
                case 111:
                    v11 := GenS2{F0: v3, F1: uint64(v5), F2: 18446744073709551615}
                    _ = v11
                case 0:
                    fmt.Println(v3, v5, fi10)
                default:
                    v12 := GenS3{F0: int(v5), F1: 692}
                    _ = v12
                }
                _ = fi10
            }
        case 32767:
            v8 = v8
        case 127:
            var v13 int = 2147483647
            _ = v13
        default:
            var v14 int = -276
            _ = v14
        }
    default:
        v15 := GenS4{F0: 53, F1: 569}
        _ = v15
    }
    if ((!(-1 < 0)) || (32767 != 42)) {
        if true {
            v16 := GenS5{F0: -32768, F1: "\n\t", F2: 255}
            v17 := GenS6{F0: v8}
            _ = v16
            _ = v17
        } else {
            if v3 {
                if (2 != 32767) {
                    if v3 {
                        var v18 bool = v3
                        v18 = (434 < v5)
                        v19 := []int{-194}
                        _ = v18
                        _ = v19
                    } else {
                        v2 = (^(v2))
                    }
                    for fi20 := 0; fi20 < 6; fi20++ {
                        v2 = v2
                        v21 := GenS3{F0: fi20, F1: 255}
                        for fi22 := 0; fi22 < 15; fi22++ {
                            switch "hello" {
                            case "a":
                                if v3 {
                                    var v23 uint = 451
                                    var v24 uint32 = 4294967295
                                    var v25 uint16 = 662
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 string = "foo bar"
                                    var v27 uint64 = 2
                                    fmt.Println(fi20)
                                    _ = v26
                                    _ = v27
                                }
                            case "\n\t":
                                var v28 uint = 255
                                _ = v28
                            case "hello":
                                v29 := GenS7{F0: true}
                                _ = v29
                            }
                            v30 := []int{(^(fi20)), (-(fi22)), (fi22 | (fi22 % ((fi22) | 1))), -2147483648}
                            fmt.Println(v2)
                            _ = fi22
                            _ = v30
                        }
                        _ = fi20
                        _ = v21
                    }
                }
                if v3 {
                    v3 = (!false)
                    var v31 bool = true
                    v6 -= v6
                    _ = v31
                } else {
                    v3 = true
                }
                var v32 int = -9223372036854775808
                _ = v32
            } else {
                if true {
                    fmt.Println(v8, v3, v6)
                } else {
                    var v33 int64 = 0
                    var v34 int8 = 127
                    _ = v33
                    _ = v34
                }
                var v35 uint16 = 42
                _ = v35
            }
        }
    }
    if (0 <= 42) {
        v2 = (v2 << ((v5) & 63))
    } else {
        v36 := GenS1{F0: 32768, F1: (v5 << ((v5) & 63))}
        fmt.Println(v6, v5)
        fmt.Println(v8)
        _ = v36
    }
    switch "\n\t" {
    case "a":
        v8 += v8
    }
    if (0 >= 535) {
        v2 += (^(v2))
    } else {
        v37 := []int{761}
        _ = v37
    }
    v8 = v8
    v2 = v2
    delete(v4, "")
    for fi38 := 0; fi38 < 18; fi38++ {
        if (v8 > v8) {
            v8 = (v8 << (((v5 + 4294967295)) & 63))
            v5 = v5
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
