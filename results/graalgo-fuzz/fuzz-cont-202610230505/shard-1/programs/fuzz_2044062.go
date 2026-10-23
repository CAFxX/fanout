package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 int32, p1 int16, p2 int8) int64 {
    var v5 float32 = -2.25
    for fi6 := 0; fi6 < 8; fi6++ {
        if true {
            var v7 float64 = 1e10
            for fi8 := 0; fi8 < 8; fi8++ {
                if true {
                    v5 -= v5
                    p2 = (p2 / ((p2) | 1))
                    fmt.Println(fi6)
                    p2 = p2
                    if true {
                        p1 = (p1 % (((p1 ^ p1)) | 1))
                    } else {
                        switch fi8 {
                        case 32767:
                            switch -2147483648 {
                            case 32768:
                                v9 := GenS1{F0: (p2 << ((2) & 63)), F1: "\n\t", F2: 2}
                                _ = v9
                            case -617:
                                switch fi6 {
                                case 54:
                                    p2 += (p2 | (p2 % ((p2) | 1)))
                                }
                            }
                        case -9223372036854775808:
                            v10 := []int{}
                            _ = v10
                        }
                    }
                }
                fmt.Println(p2, p1)
                _ = fi8
            }
            if (1.5 <= (((-2.25 - ((v5 / float32((v7 * 3.14159))) - v5)) * (v5 - float32(v7))) * v5)) {
                var v11 int16 = p1
                p1 = 346
                fmt.Println(p0, v5)
                _ = v11
            }
            _ = v7
        }
        var v12 uint64 = 2
        v13 := []int{127, -2147483648, fi6}
        _ = fi6
        _ = v12
        _ = v13
    }
    var v14 int = 939
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v5
    return 42
}

func main() {
    var v1 int64 = -1
    var v2 int = 128
    var v3 int64 = 0
    v4 := []string{}
    v3 -= genFn0(843, 733, -1)
    v3 = v1
    for fi15 := 0; fi15 < 18; fi15++ {
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v3 = genFn0(642, -792, 127)
            v16 := GenS2{F0: 255, F1: 65535}
            _ = v16
        }
        _ = fi15
    }
    fmt.Println(v2)
    if true {
        v1 -= v3
        var v17 uint64 = 2
        _ = v17
    } else {
        for fi18 := 0; fi18 < 17; fi18++ {
            v2 = v2
            if false {
                var v19 bool = (18446744073709551615 >= 674)
                _ = v19
            }
            _ = fi18
        }
        fmt.Println(v3)
        v1 = 748
    }
    var v20 uint8 = 42
    var v21 uint8 = v20
    v3 -= (v3 - 0)
    v22 := GenS3{F0: v1, F1: 2147483647}
    var v23 string = (("" + "hello") + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
