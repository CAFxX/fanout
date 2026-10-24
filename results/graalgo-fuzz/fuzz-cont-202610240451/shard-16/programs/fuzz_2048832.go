package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int8; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint64; F1 int16; F2 int16 }

func main() {
    var v1 int32 = -1
    var v2 int8 = -128
    v3 := make(map[uint32]int)
    fmt.Println(v1, v1)
    var v4 uint32 = 217
    for fi5 := 0; fi5 < 1; fi5++ {
        v3[v4] = fi5
        v6 := GenS1{F0: 0}
        v2 = 42
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: 2, F1: "hello"}
    v2 -= (-(v2))
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint16 = 823
        if ((true || (v2 >= (v2 / (((v2 ^ -111)) | 1)))) && (false || (v1 == -172))) {
            var v10 uint32 = v4
            v10 -= v10
            var v11 int16 = 909
            _ = v10
            _ = v11
        } else {
            v3[v4] = fi8
        }
        _ = fi8
        _ = v9
    }
    var v12 uint32 = (v4 % (((v4 / ((v4) | 1))) | 1))
    var v13 uint64 = uint64(v12)
    fmt.Println(v13, v2, v1)
    v3[294] = 32768
    if false {
        if false {
            for fi14 := 0; fi14 < 6; fi14++ {
                switch fi14 {
                case 9223372036854775807:
                    v15 := GenS3{F0: (v2 & int8(v1)), F1: (v12 << ((14) & 63)), F2: -221}
                    _ = v15
                case 127:
                    fmt.Println(v1)
                default:
                    for fi16 := 0; fi16 < 11; fi16++ {
                        if false {
                            delete(v3, v4)
                        }
                        _ = fi16
                    }
                }
                if true {
                    var v17 int64 = -9223372036854775808
                    switch "" {
                    case "a":
                        if false {
                            var v18 float32 = 1.5
                            var v19 int16 = 585
                            _ = v18
                            _ = v19
                        }
                    case "hello":
                        for fi20 := 0; fi20 < 19; fi20++ {
                            var v21 int64 = 407
                            switch (("" + ("hello" + "foo bar")) + "a") {
                            case "":
                                fmt.Println(v1)
                            case "foo bar":
                                v13 -= 1
                                v2 = (v2 % (((v2 * v2)) | 1))
                            case "hello":
                                v22 := GenS2{F0: 18446744073709551615, F1: "foo bar"}
                                _ = v22
                            }
                            _ = fi20
                            _ = v21
                        }
                    case "":
                        v3[2147483648] = fi14
                    default:
                        v12 = (v12 * v12)
                    }
                    var v23 uint8 = 128
                    _ = v17
                    _ = v23
                } else {
                    v24 := GenS4{F0: v13, F1: -1, F2: 375}
                    _ = v24
                }
                _ = fi14
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                v26 := GenS4{F0: 42, F1: 0, F2: 685}
                _ = fi25
                _ = v26
            }
        } else {
            var v27 int = 2147483647
            v3[v12] = (v27 * (v27 & -815))
            var v28 uint64 = v13
            _ = v27
            _ = v28
        }
        var v29 int64 = 42
        _ = v29
    }
    v4 += (v4 ^ (^(v4)))
    var v30 float64 = 3.14159
    v31 := []int{int(v1), 32767, 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v7
}
