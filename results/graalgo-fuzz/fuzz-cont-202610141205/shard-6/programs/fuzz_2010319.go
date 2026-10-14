package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 string }

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]bool)
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]string)
    v6 := []int{}
    v7 := GenS1{F0: 18446744073709551615, F1: 32768}
    if false {
        v1 = 0
        for fi8 := 0; fi8 < 17; fi8++ {
            v2[(fi8 | fi8)] = (264 > 738)
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
    } else {
        v1 = v1
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v3 = append(v3, 161)
        _ = fi10
    }
    v11 := GenS2{F0: 0}
    v1 = (v1 % ((v1) | 1))
    var v12 uint8 = 128
    if (1 == 2) {
        var v13 uint32 = 902
        if (!false) {
            v2[2147483647] = (100 != 38)
            switch int(v13) {
            case 4294967295:
                for fi14 := 0; fi14 < 20; fi14++ {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        if false {
                            if false {
                                var v16 uint32 = (v13 << ((42) & 63))
                                if true {
                                    if true {
                                        var v17 float32 = -2.25
                                        _ = v17
                                    }
                                }
                                v12 = v12
                                _ = v16
                            } else {
                                fmt.Println(fi15, fi15)
                            }
                            v13 += 860
                        } else {
                            var v18 float64 = (1e10 + 0.1)
                            v19 := GenS2{F0: v1}
                            v20 := GenS3{F0: uint16(v13), F1: ("\n\t" + "")}
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        switch "foo bar" {
                        case "hello":
                            var v21 int32 = -1
                            _ = v21
                        }
                        _ = fi15
                    }
                    var v22 int8 = 42
                    var v23 int16 = -168
                    _ = fi14
                    _ = v22
                    _ = v23
                }
            case -999:
                fmt.Println(v13, v13, v1)
            case -981:
                var v24 uint16 = 723
                _ = v24
            }
        }
        v1 = (v1 * 255)
        _ = v13
    } else {
        _ = len(v2)
    }
    v1 = (v1 << ((4294967295) & 63))
    fmt.Println(v1, v1)
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
