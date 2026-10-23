package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int8; F2 int }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int8; F1 string; F2 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 string; F1 int64 }

func main() {
    var v1 uint64 = 0
    var v2 uint64 = 80
    var v3 int16 = 1
    var v4 uint8 = 184
    v5 := []int{-2147483648, 127, -1}
    v6 := make(map[uint32]uint32)
    v1 += (v1 | 0)
    var v7 int32 = -1
    switch "\n\t" {
    case "\n\t":
        if false {
            var v8 int64 = 1
            _ = v8
        } else {
            var v9 uint16 = 430
            _ = v9
        }
    case "hello":
        var v10 int16 = (v3 & v3)
        _ = v10
    }
    var v11 uint = 255
    v11 = (v11 - 373)
    if (true || false) {
        v12 := GenS1{F0: v3}
        _ = v12
    } else {
        if (1.5 < (1.5 / -2.25)) {
            v13 := GenS2{F0: v11, F1: 1, F2: -2147483648}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                if len(v5) > 0 {
                    v5[0] = (fi14 | -2147483648)
                }
                v11 = v11
                _ = fi14
            }
        }
        fmt.Println(v2, v11)
        if true {
            fmt.Println(v2)
            for fi15 := 0; fi15 < 3; fi15++ {
                switch "hello" {
                case "hello":
                    if true {
                        v16 := GenS2{F0: (v11 * (v11 << ((18446744073709551615) & 63))), F1: -55, F2: fi15}
                        _ = v16
                    } else {
                        if true {
                            v17 := GenS2{F0: 0, F1: 16, F2: (^(fi15))}
                            if false {
                                delete(v6, 2147483647)
                                if len(v5) > 0 {
                                    v5[0] = 282
                                }
                                v18 := GenS3{F0: v11, F1: 42}
                                _ = v18
                            }
                            for fi19 := 0; fi19 < 20; fi19++ {
                                if ("" != "hello") {
                                    v4 += v4
                                    fmt.Println(fi19, v2, v1)
                                    v2 -= (v1 << ((1) & 63))
                                } else {
                                    fmt.Println(fi15, v1)
                                    v4 = v4
                                    var v20 uint16 = 2
                                    if ((487 >= 508) || true) {
                                        var v21 uint = 4294967295
                                        _ = v21
                                    } else {
                                        v22 := GenS4{F0: v3, F1: v20}
                                        v23 := GenS4{F0: 1000, F1: v20}
                                        _ = v22
                                        _ = v23
                                    }
                                    switch 61 {
                                    case 924:
                                        v24 := GenS5{F0: -1, F1: "", F2: -2147483648}
                                        _ = v24
                                    }
                                    delete(v6, 2147483648)
                                    var v25 uint32 = 811
                                    _ = v20
                                    _ = v25
                                }
                                _ = fi19
                            }
                            _ = v17
                        } else {
                            fmt.Println(fi15)
                        }
                        switch "a" {
                        case "hello":
                            fmt.Println(v7)
                        }
                        switch int(v4) {
                        case -2147483648:
                            if (255 == 785) {
                                for fi26 := 0; fi26 < 4; fi26++ {
                                    var v27 uint8 = 0
                                    _ = fi26
                                    _ = v27
                                }
                                for fi28 := 0; fi28 < 2; fi28++ {
                                    v29 := GenS6{F0: (v2 << (((v11 >> ((uint(v4)) & 63))) & 63))}
                                    var v30 uint8 = v4
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                switch (fi15 - fi15) {
                                case -2147483648:
                                    v31 := v6[462]
                                    _ = v31
                                case -1:
                                    fmt.Println(fi15, v7, v4)
                                case 1:
                                    v3 -= (v3 & (v3 >> ((256) & 63)))
                                }
                            }
                        default:
                            fmt.Println(v4, v3)
                        }
                    }
                case "":
                    if false {
                        var v32 int64 = 42
                        _ = v32
                    } else {
                        switch "hello" {
                        case "\n\t":
                            switch "" {
                            case "foo bar":
                                v11 -= (v11 % (((v11 * (v11 + (v11 & (v11 + v11))))) | 1))
                            case "":
                                v1 -= (v1 % ((v1) | 1))
                            case "hello":
                                var v33 float64 = float64(1e10)
                                _ = v33
                            }
                        case "":
                            delete(v6, 2147483648)
                        }
                    }
                default:
                    switch "\n\t" {
                    case "a":
                        v4 = 155
                    default:
                        v34 := GenS7{F0: -262, F1: v7}
                        _ = v34
                    }
                }
                var v35 int = 127
                switch "foo bar" {
                case "":
                    for fi36 := 0; fi36 < 14; fi36++ {
                        v2 = (^(v2))
                        _ = fi36
                    }
                case "hello":
                    for fi37 := 0; fi37 < 17; fi37++ {
                        fmt.Println(fi15, v35)
                        v4 = v4
                        v38 := GenS8{F0: -15, F1: "hello"}
                        _ = fi37
                        _ = v38
                    }
                }
                _ = fi15
                _ = v35
            }
            fmt.Println(v2)
        }
    }
    v7 = v7
    v11 -= (v11 % ((366) | 1))
    for fi39 := 0; fi39 < 7; fi39++ {
        switch ("\n\t" + "a") {
        case "hello":
            fmt.Println(v4, v7)
        case "\n\t":
            v3 = (v3 & v3)
        default:
            v3 += (v3 & v3)
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            v7 = v7
            _ = fi40
        }
        _ = fi39
    }
    if true {
        v1 -= (v1 ^ (v2 + 1))
        v41 := GenS9{F0: "foo bar", F1: -1}
        _ = v41
    }
    v1 -= 18446744073709551615
    if false {
        v11 += (v11 / (((^(v11))) | 1))
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
