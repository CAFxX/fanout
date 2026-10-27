package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int32; F1 int; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint8 = 128
    var v2 int64 = -1
    var v3 uint8 = 255
    var v4 string = ("\n\t" + "")
    var v5 int16 = -408
    v5 = v5
    v3 -= (^(v1))
    var v6 int8 = 42
    v1 += 184
    switch 361 {
    case -9223372036854775808:
        switch "\n\t" {
        case "hello":
            switch 9223372036854775807 {
            case -2147483648:
                v7 := []int{32767, 4294967295, 301}
                _ = v7
            case 1:
                if ((42 > 870) || true) {
                    fmt.Println(v1)
                    for fi8 := 0; fi8 < 8; fi8++ {
                        if (-2147483648 == 1) {
                            var v9 float64 = 0.0
                            for fi10 := 0; fi10 < 7; fi10++ {
                                switch (v4 + (v4 + ((v4 + v4) + "foo bar"))) {
                                case "foo bar":
                                    v1 -= v3
                                    v11 := []int{fi8, (fi8 << ((uint(v3)) & 63))}
                                    _ = v11
                                case "":
                                    var v12 string = (v4 + v4)
                                    _ = v12
                                case "\n\t":
                                    v6 += v6
                                    fmt.Println(fi8, v2, v1)
                                default:
                                    var v13 int64 = (v2 & -419)
                                    _ = v13
                                }
                                var v14 uint32 = 2147483647
                                v5 = -32768
                                v4 = v4
                                _ = fi10
                                _ = v14
                            }
                            v15 := GenS1{F0: 2, F1: 4294967295}
                            _ = v15
                            _ = v9
                        }
                        var v16 uint8 = 255
                        for fi17 := 0; fi17 < 8; fi17++ {
                            fmt.Println(v6, fi17)
                            fmt.Println(v1, v4, fi8)
                            _ = fi17
                        }
                        _ = fi8
                        _ = v16
                    }
                }
            default:
                v18 := []int{}
                _ = v18
            }
        case "foo bar":
            if true {
                v3 -= v3
            } else {
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 bool = (0.1 != 0.0)
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 11; fi21++ {
                    var v22 float64 = 0.1
                    v23 := []int{}
                    var v24 float32 = 1e10
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                v25 := []int{-444}
                _ = v25
            }
        default:
            fmt.Println(v2, v4, v5)
        }
    case 127:
        v6 = (v6 * v6)
    }
    v6 = (v6 / (((v6 - v6)) | 1))
    v26 := GenS2{F0: 2147483647, F1: 128, F2: false}
    v27 := []int{-819, 2147483647, 9223372036854775807, int(v3)}
    v28 := GenS3{F0: v5}
    if false {
        var v29 string = v4
        var v30 uint64 = 1
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        v32 := GenS4{F0: v2}
        v33 := GenS3{F0: 738}
        _ = fi31
        _ = v32
        _ = v33
    }
    if (570 >= 116) {
        var v34 uint = 65535
        if (((v2 * (v2 >> ((68) & 63))) == (v2 & (v2 % (((-(v2))) | 1)))) || (42 > -2147483648)) {
            v6 = int8(v34)
        } else {
            switch -513 {
            case -9223372036854775808:
                _ = len(v27)
            case 334:
                if len(v27) > 0 {
                    v27[0] = int(v2)
                }
            case 127:
                fmt.Println(v34, v6)
            default:
                var v35 string = "\n\t"
                _ = v35
            }
            var v36 float32 = 0.0
            fmt.Println(v5)
            _ = v36
        }
        for fi37 := 0; fi37 < 8; fi37++ {
            fmt.Println(v2, v34, v3)
            _ = fi37
        }
        _ = v34
    }
    v27 = append(v27, -9223372036854775808)
    var v38 float32 = 0.0
    v39 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
