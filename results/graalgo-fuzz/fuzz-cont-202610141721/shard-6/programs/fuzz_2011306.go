package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint32 }
func genFn0() uint {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int16 = 0
        var v7 bool = true
        for fi8 := 0; fi8 < 16; fi8++ {
            fmt.Println(v7, fi5, v6)
            switch fi5 {
            case 168:
                var v9 int8 = -1
                if ((2147483647 == 42) && v7) {
                    v9 = 100
                    var v10 string = "hello"
                    v9 += v9
                    v10 = "foo bar"
                    fmt.Println(v7, v7)
                    _ = v10
                }
                _ = v9
            }
            _ = fi8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v11 := []int{-415, 0}
    switch "" {
    case "\n\t":
        if true {
            if (false && false) {
                if (true && ((65535 >= 18446744073709551615) || ("foo bar" >= ("a" + "a")))) {
                    if false {
                        for fi12 := 0; fi12 < 9; fi12++ {
                            var v13 int8 = 42
                            _ = fi12
                            _ = v13
                        }
                        var v14 uint8 = 215
                        _ = v14
                    }
                    var v15 float32 = (1.5 + -2.25)
                    _ = v15
                }
            } else {
                switch 128 {
                case -572:
                    var v16 uint32 = 2147483647
                    _ = v16
                case 0:
                    var v17 int64 = 0
                    _ = v17
                default:
                    var v18 uint64 = 147
                    _ = v18
                }
            }
        }
    case "a":
        v11 = append(v11, 32767)
    case "hello":
        v11 = append(v11, 4294967295)
    }
    if true {
        var v19 int16 = 0
        _ = v19
    }
    var v20 uint8 = 128
    _ = v11
    _ = v20
    return 42
}
func genFn1() int16 {
    var v21 uint64 = 18446744073709551615
    var v22 bool = false
    v23 := GenS1{F0: 449, F1: 2147483648, F2: (v21 << ((934) & 63))}
    _ = v21
    _ = v22
    _ = v23
    return -32768
}

func main() {
    var v1 uint16 = 351
    v2 := []uint32{42}
    var v3 uint16 = (v1 & v1)
    var v4 uint16 = (v1 ^ 556)
    v24 := GenS1{F0: -2147483648, F1: 0, F2: 42}
    var v25 uint32 = 4294967295
    switch "hello" {
    case "foo bar":
        fmt.Println(v25, v25, v25)
    case "a":
        if (-128 == 100) {
            switch 32767 {
            case 42:
                fmt.Println(v1)
            case -1:
                v25 += (v25 >> ((255) & 63))
            case 430:
                if true {
                    switch "" {
                    case "":
                        for fi26 := 0; fi26 < 8; fi26++ {
                            _ = len(v2)
                            v27 := GenS1{F0: 420, F1: v25, F2: 18446744073709551615}
                            switch (fi26 - 718) {
                            case 840:
                                var v28 int32 = 1
                                _ = v28
                            case 2147483647:
                                v29 := GenS2{F0: -1, F1: 2}
                                _ = v29
                            }
                            _ = fi26
                            _ = v27
                        }
                    default:
                        v1 = (v3 | (v1 ^ v1))
                    }
                    v4 = 65535
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            if true {
                var v30 int8 = 43
                var v31 uint = genFn0()
                _ = v30
                _ = v31
            }
        }
    default:
        _ = len(v2)
    }
    switch "\n\t" {
    case "":
        for fi32 := 0; fi32 < 1; fi32++ {
            _ = len(v2)
            _ = fi32
        }
    case "\n\t":
        v4 = v3
    case "hello":
        switch 42 {
        case 42:
            v33 := GenS3{F0: 2147483648}
            _ = v33
        default:
            for fi34 := 0; fi34 < 17; fi34++ {
                v3 = (v3 | 1)
                var v35 int16 = 0
                if true {
                    v2 = append(v2, (v25 - 4294967295))
                    if ((false || (!true)) && false) {
                        v4 += (v4 * (v4 & v4))
                        var v36 uint64 = 282
                        _ = len(v2)
                        _ = v36
                    } else {
                        if false {
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v3, fi34)
                            default:
                                var v37 uint32 = (v25 & 1)
                                _ = v37
                            }
                        } else {
                            v1 += (^(v3))
                            for fi38 := 0; fi38 < 1; fi38++ {
                                v4 = (v4 / ((65535) | 1))
                                var v39 int = (fi34 | fi34)
                                _ = fi38
                                _ = v39
                            }
                        }
                        var v40 int16 = genFn1()
                        _ = v40
                    }
                    if true {
                        var v41 int8 = 1
                        _ = v41
                    }
                }
                _ = fi34
                _ = v35
            }
        }
    }
    v1 = (v1 + v1)
    var v42 uint64 = 2
    v43 := GenS3{F0: (v25 - (v25 & v25))}
    var v44 int64 = -348
    v2 = append(v2, 65535)
    switch -9223372036854775808 {
    case -950:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 650:
        v45 := GenS1{F0: 42, F1: (v25 << ((176) & 63)), F2: 982}
        _ = v45
    default:
        fmt.Println(v3, v3)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v46 uint = 42
    var v47 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v46
    _ = v47
}
