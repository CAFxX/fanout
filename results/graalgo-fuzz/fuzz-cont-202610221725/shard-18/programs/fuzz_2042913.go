package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int8; F2 int }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32 }
func genFn0() int64 {
    v6 := GenS1{F0: 372, F1: 0, F2: 0}
    var v7 int16 = 0
    v8 := GenS2{F0: 344, F1: -1, F2: 1}
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 uint32 = 124
    var v2 int64 = 1
    var v3 bool = true
    var v4 float32 = 0.1
    v5 := []uint32{}
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint8 = uint8(v1)
        v2 = (v2 >> ((18446744073709551615) & 63))
        var v11 int8 = 42
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v4, v3)
    var v12 uint64 = 872
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v13 := GenS3{F0: (v2 >> ((65535) & 63)), F1: 0}
    v14 := GenS4{F0: false, F1: 32767, F2: v12}
    for fi15 := 0; fi15 < 16; fi15++ {
        if true {
            switch "hello" {
            case "hello":
                if v3 {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v3 = true
                        _ = fi16
                    }
                }
            case "":
                switch (fi15 << ((0) & 63)) {
                case -326:
                    v17 := GenS5{F0: ((v1 << ((1) & 63)) < 241)}
                    _ = v17
                default:
                    var v18 uint64 = v12
                    _ = v18
                }
            default:
                v19 := GenS5{F0: (!(((fi15 / ((9223372036854775807) | 1)) <= (^(fi15))) || v3))}
                _ = v19
            }
            v2 -= -1
            if true {
                if v3 {
                    if (42 <= 42) {
                        v20 := GenS3{F0: v2, F1: -715}
                        switch fi15 {
                        case -330:
                            if (!(v12 == (v12 / ((2) | 1)))) {
                                v12 = v12
                                v12 -= (v12 * v12)
                                switch (fi15 << ((301) & 63)) {
                                case 896:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case -1:
                                    _ = len(v5)
                                default:
                                    v21 := GenS6{F0: -2147483648, F1: ""}
                                    _ = v21
                                }
                            } else {
                                v4 = v4
                            }
                        case -2147483648:
                            fmt.Println(v4, v1, v4)
                        case -1:
                            for fi22 := 0; fi22 < 12; fi22++ {
                                fmt.Println(fi15, v2)
                                if v3 {
                                    _ = len(v5)
                                    fmt.Println(v4, v3)
                                }
                                var v23 int64 = v2
                                _ = fi22
                                _ = v23
                            }
                        }
                        fmt.Println(v4, v1)
                        _ = v20
                    }
                    v24 := GenS1{F0: -32768, F1: -1, F2: (^(v1))}
                    if (1 < 0) {
                        v2 = (v2 << ((302) & 63))
                        v2 += v2
                    } else {
                        v25 := GenS7{F0: 0}
                        var v26 int = fi15
                        _ = v25
                        _ = v26
                    }
                    _ = v24
                }
                v27 := GenS6{F0: -330, F1: ""}
                v2 += (v2 % ((184) | 1))
                _ = v27
            }
        }
        fmt.Println(v12)
        v4 = (1e10 * float32(3.14159))
        _ = fi15
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v2 -= v2
        _ = fi28
    }
    v3 = false
    _ = len(v5)
    v3 = ((true && v3) && v3)
    v2 -= v2
    var v29 string = ("hello" + ((("hello" + "a") + "a") + "a"))
    v3 = (((2 >= 1) && (v3 || v3)) && v3)
    v30 := GenS8{F0: 2147483648}
    v2 += v2
    v31 := GenS8{F0: (^(v1))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
