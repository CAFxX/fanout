package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int; F1 uint; F2 uint }
type GenS4 struct { F0 int8; F1 uint64 }

func main() {
    var v1 uint32 = 1
    var v2 uint64 = 1
    var v3 uint32 = 0
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 111, F1: false}
        fmt.Println(v3)
        v7 := GenS2{F0: 42, F1: -125}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{4294967295}
    v1 = (v3 >> ((854) & 63))
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int = (fi9 << ((65535) & 63))
        v2 = (v2 + 731)
        fmt.Println(v3, v4, v3)
        _ = fi9
        _ = v10
    }
    v2 += (v2 >> ((255) & 63))
    v11 := GenS3{F0: -1, F1: 65535, F2: uint(v3)}
    if false {
        fmt.Println(v2)
        v4 = (v4 + v4)
        switch 127 {
        case 937:
            if (0 >= 100) {
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(v3, v1, v2)
                    switch "hello" {
                    case "hello":
                        for fi13 := 0; fi13 < 20; fi13++ {
                            var v14 int64 = 42
                            _ = fi13
                            _ = v14
                        }
                    case "":
                        for fi15 := 0; fi15 < 17; fi15++ {
                            var v16 int8 = 0
                            if true {
                                fmt.Println(v16, v2)
                            } else {
                                var v17 int = (-(fi15))
                                v18 := GenS1{F0: 255, F1: (!false)}
                                if (uint(v1) <= 0) {
                                    switch (v4 + (v4 + v4)) {
                                    case "":
                                        switch 128 {
                                        case 127:
                                            fmt.Println(v3, v4, v16)
                                        case 32767:
                                            v19 := GenS3{F0: 42, F1: 65535, F2: 0}
                                            _ = v19
                                        case -634:
                                            fmt.Println(fi15, v1, v4)
                                        }
                                    default:
                                        fmt.Println(fi12, fi12, v17)
                                        v20 := GenS1{F0: 42, F1: true}
                                        _ = v20
                                    }
                                }
                                _ = v17
                                _ = v18
                            }
                            var v21 uint16 = 170
                            _ = fi15
                            _ = v16
                            _ = v21
                        }
                    default:
                        var v22 int8 = -1
                        _ = v22
                    }
                    _ = fi12
                }
                v2 = (v2 + v2)
            }
        case 448:
            var v23 uint8 = 0
            _ = v23
        case 127:
            if (false || (0 != uint16(v3))) {
                for fi24 := 0; fi24 < 11; fi24++ {
                    fmt.Println(fi24)
                    v3 += 2
                    _ = fi24
                }
            } else {
                fmt.Println(v3)
                v2 += v2
                var v25 int32 = 42
                _ = v25
            }
        default:
            var v26 uint64 = (v2 ^ v2)
            _ = v26
        }
    }
    var v27 uint8 = 23
    if len(v8) > 0 {
        v8[0] = 42
    }
    v8 = append(v8, -9223372036854775808)
    switch 32767 {
    case 0:
        v1 -= v3
    case -1:
        fmt.Println(v27, v27, v4)
    case -9223372036854775808:
        if len(v8) > 0 {
            v8[0] = 22
        }
    default:
        for fi28 := 0; fi28 < 15; fi28++ {
            fmt.Println(fi28, fi28)
            _ = fi28
        }
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        var v30 float64 = 1e10
        if false {
            v31 := GenS4{F0: -1, F1: (v2 / ((v2) | 1))}
            var v32 uint64 = v2
            _ = v31
            _ = v32
        } else {
            fmt.Println(v1, v3)
            v33 := GenS2{F0: 42, F1: int8(fi29)}
            var v34 uint64 = (v2 ^ (v2 ^ v2))
            _ = v33
            _ = v34
        }
        _ = fi29
        _ = v30
    }
    v27 = v27
    for fi35 := 0; fi35 < 5; fi35++ {
        var v36 string = "foo bar"
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v8
}
