package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int32 = -749
    var v2 int32 = v1
    var v3 uint8 = 14
    var v4 string = ("\n\t" + "")
    fmt.Println(v4, v1, v2)
    v5 := []int{}
    var v6 int8 = -1
    var v7 float64 = float64(-2.25)
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(fi9, v2)
            if (2 < 2) {
                v5 = append(v5, (fi9 ^ -429))
                v10 := GenS1{F0: 65535, F1: v6, F2: 2}
                _ = v10
            } else {
                v11 := GenS2{F0: 1000}
                _ = v11
            }
            if (42 < 42) {
                v3 += (v3 ^ (v3 % (((v3 % ((v3) | 1))) | 1)))
                v5 = append(v5, fi9)
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(v4, fi8)
                fmt.Println(v4, fi8)
            }
            var v12 string = v4
            var v13 int32 = (v2 | (^(v2)))
            _ = fi9
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    if false {
        v3 += v3
    }
    _ = len(v5)
    v4 = ("" + ((v4 + "") + ""))
    v5 = append(v5, -2147483648)
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    switch int(v6) {
    case 2147483647:
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int32 = 0
            _ = fi14
            _ = v15
        }
    default:
        switch (v4 + "a") {
        case "a":
            switch 871 {
            case 128:
                var v16 int8 = (-(v6))
                _ = v16
            }
        case "hello":
            fmt.Println(v6)
        default:
            v17 := GenS2{F0: 699}
            _ = v17
        }
    }
    v2 = v1
    _ = len(v5)
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (-1 <= 47) {
                fmt.Println(v1)
                v6 -= (v6 / (((v6 | 127)) | 1))
                var v20 uint64 = 801
                _ = v20
            } else {
                switch (fi18 | fi18) {
                case 127:
                    v3 = v3
                case 4294967295:
                    var v21 bool = false
                    _ = v21
                case 225:
                    v2 = (v1 + 2147483647)
                }
            }
            if (18446744073709551615 != 2) {
                if (false && false) {
                    v4 = v4
                    fmt.Println(v7)
                    v2 = (^(v2))
                }
            } else {
                v2 = -1
            }
            _ = fi19
        }
        _ = fi18
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v3 -= (v3 % ((v3) | 1))
        _ = fi22
    }
    v4 = "\n\t"
    for fi23 := 0; fi23 < 9; fi23++ {
        if len(v5) > 0 {
            v5[0] = fi23
        }
        v24 := GenS1{F0: 222, F1: v6, F2: 1}
        fmt.Println(v7)
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 0, F1: (v6 >> ((18446744073709551615) & 63)), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
