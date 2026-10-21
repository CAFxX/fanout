package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 string }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    var v1 float32 = 0.0
    var v2 int32 = -981
    var v3 uint16 = 32767
    var v4 uint8 = 2
    var v5 uint16 = (v3 >> ((uint(v4)) & 63))
    v2 += (^(v2))
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 255, F1: v5, F2: ("" + "hello")}
        switch "hello" {
        case "\n\t":
            for fi8 := 0; fi8 < 16; fi8++ {
                v2 -= int32(v4)
                var v9 uint16 = v3
                _ = fi8
                _ = v9
            }
        default:
            if (0 >= 2) {
                v10 := []int{fi6, (fi6 ^ fi6)}
                _ = v10
            } else {
                var v11 bool = true
                _ = v11
            }
        }
        _ = fi6
        _ = v7
    }
    v12 := []int{-705, 4294967295, -2147483648}
    for fi13 := 0; fi13 < 17; fi13++ {
        if true {
            v1 = float32((((3.14159 / 0.1) / (float64((float32(0.1) * v1)) * 1e10)) - 0.0))
        } else {
            var v14 bool = (false && (0 == 2))
            var v15 uint64 = 1
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    if (true && false) {
        var v16 string = ("a" + "\n\t")
        _ = v16
    } else {
        var v17 uint64 = 1
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            for fi19 := 0; fi19 < 15; fi19++ {
                v3 -= v3
                _ = fi19
            }
            _ = fi18
        }
        _ = v17
    }
    v2 = v2
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if false {
        v3 = v5
    }
    if (74 <= 1) {
        switch "a" {
        case "foo bar":
            var v20 uint = 746
            _ = v20
        }
    } else {
        v2 -= (^(v2))
        _ = len(v12)
    }
    v12 = append(v12, 1)
    var v21 bool = (!(float64(v1) > 1e10))
    var v22 float32 = (v1 + 0.1)
    v23 := GenS2{F0: v4, F1: -430}
    for fi24 := 0; fi24 < 11; fi24++ {
        v1 += (((1.5 - v22) * (1.5 + v1)) - v1)
        for fi25 := 0; fi25 < 10; fi25++ {
            if ((-128 > 1) || (!false)) {
                v22 = v1
                switch (fi25 / ((-542) | 1)) {
                case -9223372036854775808:
                    var v26 bool = true
                    _ = v26
                case -689:
                    if (-15 <= -113) {
                        if ((v5 + (^(v3))) != (v5 - v5)) {
                            switch (fi24 + (fi25 << ((0) & 63))) {
                            case 9223372036854775807:
                                if v21 {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                } else {
                                    var v27 bool = (351 < 2)
                                    _ = v27
                                }
                            case 32767:
                                v12 = append(v12, int(v4))
                            case 32768:
                                v21 = true
                            }
                        } else {
                            v28 := GenS3{F0: v21, F1: "a"}
                            var v29 int8 = -128
                            _ = v28
                            _ = v29
                        }
                    } else {
                        fmt.Println(v5, v21)
                        v30 := GenS2{F0: (^(v4)), F1: (-(fi24))}
                        var v31 uint32 = 4294967295
                        v32 := GenS4{F0: (v4 - 87), F1: fi25}
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                case 2147483647:
                    v5 = v5
                }
            }
            _ = fi25
        }
        _ = fi24
    }
    if len(v12) > 0 {
        v12[0] = 631
    }
    v22 = float32(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
