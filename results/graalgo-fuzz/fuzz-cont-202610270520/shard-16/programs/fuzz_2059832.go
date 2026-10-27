package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int16 }
type GenS2 struct { F0 int32; F1 string; F2 bool }

func main() {
    var v1 uint32 = 42
    var v2 int32 = -870
    var v3 uint32 = (v1 >> ((110) & 63))
    var v4 uint64 = 2
    var v5 int8 = 117
    var v6 uint32 = (v3 / ((1) | 1))
    v7 := []int{-9223372036854775808, 188, int(v5), 1}
    var v8 int16 = -32768
    if false {
        switch 533 {
        case 4294967295:
            for fi9 := 0; fi9 < 1; fi9++ {
                v1 = v1
                if (false || ("a" <= (("a" + "a") + ""))) {
                    switch "" {
                    case "hello":
                        if len(v7) > 0 {
                            v7[0] = (fi9 ^ 9223372036854775807)
                        }
                    default:
                        switch "a" {
                        case "\n\t":
                            if (!true) {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                switch 32767 {
                                case 42:
                                    if len(v7) > 0 {
                                        v7[0] = 9223372036854775807
                                    }
                                default:
                                    v4 += (v4 & 42)
                                }
                                for fi10 := 0; fi10 < 19; fi10++ {
                                    v7 = append(v7, fi9)
                                    _ = fi10
                                }
                            } else {
                                v11 := GenS1{F0: 1, F1: true, F2: (v8 & v8)}
                                _ = v11
                            }
                        }
                    }
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    fmt.Println(v2, v8)
                } else {
                    v5 += (v5 / (((v5 - 127)) | 1))
                }
                _ = fi9
            }
        case -1:
            var v12 uint = 2
            _ = v12
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 5; fi14++ {
                if false {
                    v3 += (v6 / ((42) | 1))
                    v15 := GenS1{F0: 2, F1: false, F2: -1}
                    _ = v15
                }
                _ = len(v7)
                _ = fi14
            }
            v6 = v3
            for fi16 := 0; fi16 < 1; fi16++ {
                v7 = append(v7, (fi16 | 4294967295))
                if (1 != 65535) {
                    v7 = append(v7, 128)
                }
                _ = fi16
            }
            _ = fi13
        }
    }
    v17 := GenS1{F0: uint8(v4), F1: (42 == uint(v6)), F2: (v8 & v8)}
    if true {
        fmt.Println(v1, v4, v5)
        if (-683 == 42) {
            for fi18 := 0; fi18 < 7; fi18++ {
                v1 = (v3 % (((^(v6))) | 1))
                _ = fi18
            }
            v19 := GenS2{F0: (v2 - v2), F1: "", F2: (!false)}
            _ = v19
        }
    } else {
        v20 := GenS2{F0: v2, F1: "hello", F2: true}
        _ = v20
    }
    v7 = append(v7, -2147483648)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v5, v3)
    var v21 uint64 = v4
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
