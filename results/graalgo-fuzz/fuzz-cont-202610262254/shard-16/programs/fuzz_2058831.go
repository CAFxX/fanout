package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint64; F1 bool }

func main() {
    var v1 float32 = (1.5 * float32(0.0))
    var v2 int32 = 950
    var v3 uint8 = 97
    var v4 int = 127
    var v5 int8 = 42
    var v6 uint32 = 579
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 127, F1: -921}
        v3 = (^(v3))
        _ = fi7
        _ = v8
    }
    v9 := []int{(v4 - (v4 ^ v4)), v4}
    v6 = 518
    switch -461 {
    case -1:
        if false {
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 int32 = v2
                var v12 int = (-(fi10))
                for fi13 := 0; fi13 < 7; fi13++ {
                    fmt.Println(v5, v1)
                    if true {
                        _ = len(v9)
                        var v14 uint16 = 1
                        _ = v14
                    } else {
                        v5 -= (^(v5))
                        var v15 int32 = (v2 ^ int32(v6))
                        _ = v15
                    }
                    if len(v9) > 0 {
                        v9[0] = (-(v12))
                    }
                    var v16 uint8 = 127
                    switch "a" {
                    case "hello":
                        if false {
                            if (769 == -1) {
                                v17 := GenS2{F0: -365, F1: uint16(v3)}
                                switch fi13 {
                                case 0:
                                    var v18 string = ""
                                    _ = v18
                                }
                                _ = v17
                            } else {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                fmt.Println(v11, v4, v6)
                            }
                            switch "foo bar" {
                            case "":
                                v19 := GenS1{F0: -38, F1: fi10}
                                _ = v19
                            case "\n\t":
                                v11 = (^(v11))
                            case "foo bar":
                                v9 = append(v9, -9223372036854775808)
                            }
                            v5 = (^(v5))
                        }
                        switch (^(fi10)) {
                        case 42:
                            if true {
                                _ = len(v9)
                            }
                        case -1:
                            var v20 int32 = (v2 >> ((0) & 63))
                            _ = v20
                        default:
                            fmt.Println(fi10, v11)
                        }
                        fmt.Println(v11, v16, fi13)
                    case "a":
                        if (128 < 14) {
                            var v21 int8 = (v5 >> ((2) & 63))
                            var v22 float32 = v1
                            _ = v21
                            _ = v22
                        } else {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            var v23 int32 = v2
                            v24 := GenS3{F0: 359, F1: false}
                            _ = v23
                            _ = v24
                        }
                    default:
                        v1 = v1
                    }
                    _ = fi13
                    _ = v16
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            v25 := GenS2{F0: 32767, F1: 0}
            _ = v25
        }
    default:
        var v26 float64 = 0.0
        _ = v26
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    if ((v6 == 65535) || (552 < 4294967295)) {
        if (v5 == (v5 * -121)) {
            v2 = v2
            v6 -= (v6 * uint32(v3))
        }
    }
    v27 := GenS1{F0: (-(v5)), F1: 32768}
    switch (^(v4)) {
    case -2147483648:
        for fi28 := 0; fi28 < 8; fi28++ {
            if true {
                v4 -= v4
            }
            _ = fi28
        }
    case 1:
        v29 := GenS1{F0: (v5 - (-(v5))), F1: v4}
        _ = v29
    default:
        v9 = append(v9, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
