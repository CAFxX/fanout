package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 int32 = 941
    v4 := []uint32{}
    if false {
        switch 1 {
        case 32767:
            for fi5 := 0; fi5 < 9; fi5++ {
                for fi6 := 0; fi6 < 17; fi6++ {
                    v3 = -1
                    v3 = v3
                    _ = fi6
                }
                _ = fi5
            }
        case 0:
            var v7 uint64 = 42
            _ = v7
        }
        v3 = (^(v3))
    } else {
        if false {
            if false {
                for fi8 := 0; fi8 < 9; fi8++ {
                    v3 = v3
                    _ = fi8
                }
                var v9 uint32 = 4294967295
                v3 += (v3 - v3)
                _ = v9
            } else {
                v3 += v3
            }
            if (float32((0.0 / float64((1.5 + (1.5 - float32(((1.5 - 0.0) + (0.1 / 0.1)))))))) == 3.14159) {
                v3 = (-(v3))
                var v10 bool = (!((false && true) || ("a" >= (("" + "a") + ("\n\t" + "\n\t")))))
                v4 = append(v4, 42)
                _ = v10
            } else {
                switch ("hello" + "hello") {
                case "":
                    v3 = v3
                case "a":
                    if len(v4) > 0 {
                        v4[0] = 2
                    }
                case "foo bar":
                    _ = len(v4)
                default:
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                }
                for fi11 := 0; fi11 < 5; fi11++ {
                    v3 += int32(fi11)
                    _ = fi11
                }
                var v12 uint8 = 1
                _ = v12
            }
        }
        fmt.Println(v3)
    }
    _ = len(v1)
    v3 += (v3 - -2147483648)
    var v13 int32 = (v3 + (v3 - v3))
    delete(v1, -84)
    v3 -= v3
    var v14 float64 = -2.25
    if (-995 > 42) {
        fmt.Println(v3, v14, v13)
    } else {
        v15 := GenS1{F0: 18446744073709551615, F1: 255}
        for fi16 := 0; fi16 < 17; fi16++ {
            if false {
                for fi17 := 0; fi17 < 8; fi17++ {
                    if true {
                        v18 := GenS2{F0: 513}
                        _ = v18
                    } else {
                        var v19 uint16 = 0
                        switch fi16 {
                        case -2147483648:
                            v4 = append(v4, 276)
                        case 703:
                            fmt.Println(v3)
                        case -1:
                            if false {
                                v4 = append(v4, 42)
                            } else {
                                var v20 int8 = int8(v3)
                                fmt.Println(fi17, v19)
                                var v21 uint32 = 2
                                _ = v20
                                _ = v21
                            }
                        }
                        fmt.Println(fi16, v14)
                        _ = v19
                    }
                    if len(v4) > 0 {
                        v4[0] = 182
                    }
                    _ = len(v4)
                    _ = fi17
                }
            }
            _ = fi16
        }
        _ = v15
    }
    var v22 int64 = -1
    var v23 int8 = 100
    v24 := GenS3{F0: 18446744073709551615, F1: 736, F2: 2}
    var v25 int8 = v23
    for fi26 := 0; fi26 < 16; fi26++ {
        v4 = append(v4, 838)
        _ = fi26
    }
    _ = len(v2)
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 11; fi28++ {
            if len(v4) > 0 {
                v4[0] = 470
            }
            v22 -= (v22 ^ (v22 / (((v22 << ((2) & 63))) | 1)))
            fmt.Println(v3, v23)
            _ = fi28
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            switch -9223372036854775808 {
            case 2147483647:
                v23 = v25
            case 933:
                var v30 int8 = (v25 & 44)
                _ = v30
            case 127:
                v25 = v25
            }
            _ = fi29
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
