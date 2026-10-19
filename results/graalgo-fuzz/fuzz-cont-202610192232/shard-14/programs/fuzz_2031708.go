package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 uint32; F2 bool }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]string)
    var v3 uint32 = 743
    var v4 int = -9223372036854775808
    v5 := []int{42}
    var v6 uint64 = 2
    switch v4 {
    case -9223372036854775808:
        if false {
            if true {
                v1 = (^(v1))
            } else {
                for fi7 := 0; fi7 < 17; fi7++ {
                    if ((v6 + v6) <= v6) {
                        v3 = (^(v3))
                        for fi8 := 0; fi8 < 12; fi8++ {
                            var v9 uint32 = v3
                            var v10 bool = false
                            v11 := GenS1{F0: v1}
                            _ = fi8
                            _ = v10
                            _ = v11
                            _ = v9
                        }
                        v4 -= fi7
                    } else {
                        var v12 int64 = 9223372036854775807
                        _ = v12
                    }
                    v5 = append(v5, -111)
                    fmt.Println(v1)
                    _ = fi7
                }
                if true {
                    var v13 int8 = -40
                    _ = v13
                } else {
                    v14 := GenS2{F0: v4, F1: 42, F2: v6}
                    v1 = (v1 * v1)
                    _ = v14
                }
                for fi15 := 0; fi15 < 17; fi15++ {
                    if true {
                        for fi16 := 0; fi16 < 9; fi16++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi16
                        }
                        v6 = (v6 % ((1) | 1))
                        v5 = append(v5, -113)
                    } else {
                        if (-1 == -697) {
                            for fi17 := 0; fi17 < 6; fi17++ {
                                if (fi15 == -2147483648) {
                                    var v18 uint = 100
                                    v19 := GenS3{F0: 42, F1: (((v3 ^ (v3 << ((v18) & 63))) > 0) || (false && false))}
                                    switch fi15 {
                                    case 9223372036854775807:
                                        v20 := GenS2{F0: fi17, F1: 63, F2: (v6 % ((v6) | 1))}
                                        _ = v20
                                    }
                                    _ = v18
                                    _ = v19
                                } else {
                                    if (false && (!((42 > 42) && false))) {
                                        fmt.Println(fi15, v3)
                                    }
                                    v1 = 1
                                }
                                _ = fi17
                            }
                        }
                    }
                    v5 = append(v5, fi15)
                    _ = fi15
                }
            }
            fmt.Println(v6, v3, v6)
        } else {
            v3 = (v3 % (((v3 * (^(v3)))) | 1))
            v1 += v1
            _ = len(v2)
        }
    }
    v6 = 4294967295
    fmt.Println(v4, v6, v3)
    if true {
        v21 := GenS4{F0: -167}
        _ = v21
    } else {
        var v22 int8 = 42
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 uint32 = 164
            var v25 uint16 = (v1 * uint16(v24))
            _ = fi23
            _ = v24
            _ = v25
        }
        for fi26 := 0; fi26 < 10; fi26++ {
            fmt.Println(v22, v3)
            delete(v2, 32767)
            _ = fi26
        }
        _ = v22
    }
    fmt.Println(v3, v3, v4)
    delete(v2, (v4 - -1))
    v27 := GenS3{F0: 42, F1: (364 != (v3 << ((493) & 63)))}
    var v28 float64 = ((((((float64(((1e10 / (-2.25 + 1.5)) / (3.14159 * 1e10))) / (0.1 + float64((((3.14159 * (1.5 - 1e10)) * 3.14159) / (-2.25 + 3.14159))))) / float64(((0.1 - (-2.25 - 1e10)) - 1.5))) * 1.5) / 1.0) * 1e10) + 0.1)
    fmt.Println(v4)
    _ = len(v5)
    var v29 float32 = -2.25
    v30 := GenS5{F0: 65535, F1: (v3 << ((4294967295) & 63)), F2: (818 == 755)}
    var v31 int8 = -1
    switch "hello" {
    case "hello":
        v32 := v2[(v4 & (v4 + v4))]
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
