package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 string, p1 int64) int32 {
    p0 = p0
    p0 = (p0 + "a")
    v5 := GenS1{F0: p1, F1: -108}
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    var v1 bool = (947 <= 65535)
    var v2 int8 = 1
    var v3 int = 128
    var v4 uint = 49
    fmt.Println(v2, v3)
    var v6 int8 = -1
    v7 := []int{v3, v3}
    v7 = append(v7, 42)
    for fi8 := 0; fi8 < 18; fi8++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    var v9 float32 = (1e10 / (1.5 / 1.0))
    var v10 int32 = 2147483647
    if false {
        v6 += (v6 + (-(v2)))
    } else {
        if true {
            var v11 uint8 = 127
            switch "foo bar" {
            case "":
                if v1 {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 int64 = 627
                        _ = fi12
                        _ = v13
                    }
                    if v1 {
                        switch 128 {
                        case -9223372036854775808:
                            v14 := GenS1{F0: -9223372036854775808, F1: v10}
                            _ = v14
                        case 240:
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        default:
                            v15 := GenS1{F0: 1, F1: (-(v10))}
                            _ = v15
                        }
                        v11 -= v11
                        for fi16 := 0; fi16 < 20; fi16++ {
                            v17 := GenS2{F0: v1}
                            if len(v7) > 0 {
                                v7[0] = fi16
                            }
                            if v1 {
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    var v19 int8 = v6
                                    v20 := GenS1{F0: -109, F1: (v10 & v10)}
                                    if ((("foo bar" + "\n\t") + "foo bar") != ("hello" + "foo bar")) {
                                        var v21 uint32 = 156
                                        if v1 {
                                            v22 := GenS3{F0: fi16, F1: 42, F2: (v19 / (((v19 + (v6 >> (((v4 << ((uint(v11)) & 63))) & 63)))) | 1))}
                                            switch v3 {
                                            case -708:
                                                var v23 float64 = (0.1 * (float64(1.5) - float64(v9)))
                                                _ = v23
                                            default:
                                                v21 = 1
                                            }
                                            if ((v10 << ((v4) & 63)) != (v10 - (v10 & v10))) {
                                                v24 := GenS1{F0: -9223372036854775808, F1: v10}
                                                _ = v24
                                            } else {
                                                fmt.Println(v21, v4, v4)
                                                v25 := GenS4{F0: v4}
                                                v26 := GenS2{F0: false}
                                                _ = v25
                                                _ = v26
                                            }
                                            v6 = v6
                                            _ = v22
                                        } else {
                                            v4 += (v4 ^ 43)
                                        }
                                        _ = v21
                                    }
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                            }
                            _ = fi16
                            _ = v17
                        }
                    }
                    switch "hello" {
                    case "":
                        fmt.Println(v9, v11)
                    case "a":
                        for fi27 := 0; fi27 < 2; fi27++ {
                            _ = len(v7)
                            _ = fi27
                        }
                    default:
                        var v28 int8 = v2
                        _ = v28
                    }
                }
            }
            _ = v11
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            if false {
                fmt.Println(v3, fi29, v2)
                if true {
                    v1 = (!(159 >= 156))
                }
                v9 = 1e10
            }
            _ = fi29
        }
    }
    _ = len(v7)
    if false {
        switch v3 {
        case -261:
            var v30 uint64 = 664
            _ = v30
        case -692:
            fmt.Println(v2, v2, v2)
        }
    }
    fmt.Println(v6)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v31 int = v3
    for fi32 := 0; fi32 < 5; fi32++ {
        for fi33 := 0; fi33 < 17; fi33++ {
            var v34 uint64 = 1
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v35 uint16 = 2
    v9 += float32((((0.0 + (-2.25 * (float64((((float32(0.1) * (v9 + v9)) * 0.0) - ((3.14159 * v9) - float32(-2.25)))) - 3.14159))) * 3.14159) - 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
