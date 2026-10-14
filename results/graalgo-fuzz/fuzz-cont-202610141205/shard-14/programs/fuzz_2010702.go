package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int }

func main() {
    var v1 int32 = -747
    var v2 int32 = 1
    var v3 float64 = 1.5
    var v4 bool = (!true)
    var v5 uint8 = 112
    v2 = -843
    v6 := []int{}
    if (!false) {
        v7 := GenS1{F0: (v1 * (v1 % (((-(v2))) | 1))), F1: 4294967295, F2: 2}
        _ = v7
    } else {
        if (v5 == v5) {
            for fi8 := 0; fi8 < 12; fi8++ {
                if v4 {
                    v4 = ("\n\t" > "")
                    if (65535 != 155) {
                        if (65535 == 65535) {
                            v4 = v4
                            v3 -= v3
                        }
                    } else {
                        if true {
                            v2 += -2147483648
                        }
                    }
                } else {
                    var v9 int = fi8
                    _ = v9
                }
                v1 = (v1 % ((v2) | 1))
                switch "" {
                case "hello":
                    v10 := GenS2{F0: 111}
                    _ = v10
                case "\n\t":
                    if len(v6) > 0 {
                        v6[0] = fi8
                    }
                }
                _ = fi8
            }
            if len(v6) > 0 {
                v6[0] = 128
            }
        }
    }
    v11 := GenS2{F0: 331}
    v12 := GenS2{F0: 1}
    v13 := GenS1{F0: (v1 << ((2) & 63)), F1: -1, F2: v5}
    if (-397 < 0) {
        switch 4294967295 {
        case 1:
            var v14 string = "foo bar"
            _ = v14
        case -404:
            v1 = v1
        case 2147483647:
            for fi15 := 0; fi15 < 20; fi15++ {
                v3 = 0.1
                _ = fi15
            }
        default:
            v4 = (float64(-2.25) >= v3)
        }
    } else {
        v2 = (v2 / ((-2147483648) | 1))
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 float64 = (v3 * 1e10)
            _ = fi16
            _ = v17
        }
    }
    if v4 {
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int8 = -53
            _ = fi18
            _ = v19
        }
    }
    if true {
        for fi20 := 0; fi20 < 1; fi20++ {
            fmt.Println(v2)
            var v21 string = "\n\t"
            var v22 float32 = float32(0.0)
            _ = fi20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := GenS3{F0: (-(v2)), F1: fi23}
            _ = fi23
            _ = v24
        }
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
