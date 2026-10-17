package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int8 }

func main() {
    var v1 float32 = (3.14159 - 3.14159)
    var v2 int = -237
    v3 := []uint32{0}
    for fi4 := 0; fi4 < 4; fi4++ {
        v1 = 0.1
        v5 := GenS1{F0: 42, F1: 4294967295}
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 237
    fmt.Println(v2)
    fmt.Println(v1)
    var v7 int32 = -2147483648
    fmt.Println(v7, v7)
    var v8 int = v2
    v1 = v1
    if false {
        fmt.Println(v6, v6)
    }
    var v9 uint = 2
    if (18446744073709551615 < 963) {
        switch (v8 << ((v9) & 63)) {
        case 32767:
            var v10 int8 = 42
            _ = v10
        case 32768:
            if false {
                if ((v9 - (^(v9))) == 0) {
                    fmt.Println(v2)
                    if true {
                        for fi11 := 0; fi11 < 15; fi11++ {
                            for fi12 := 0; fi12 < 18; fi12++ {
                                var v13 uint64 = 133
                                v14 := GenS2{F0: v13, F1: 969, F2: -128}
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            if ((v6 & v6) == v6) {
                                v9 = (v9 - v9)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                v7 += v7
                                if true {
                                    if (((v6 ^ (v6 & 865)) != v6) || (1 >= 1000)) {
                                        v6 = 720
                                        v3 = append(v3, v6)
                                        v2 = (v2 & v8)
                                    } else {
                                        v7 = (v7 & (v7 << ((815) & 63)))
                                    }
                                    _ = len(v3)
                                    var v15 float32 = v1
                                    _ = v15
                                }
                                var v16 int = (v2 & v2)
                                _ = v16
                            }
                            _ = fi11
                        }
                        fmt.Println(v8, v1)
                    }
                }
            }
        case -1:
            if len(v3) > 0 {
                v3[0] = v6
            }
        default:
            if true {
                switch ("" + ("" + (("foo bar" + "a") + "hello"))) {
                case "a":
                    for fi17 := 0; fi17 < 8; fi17++ {
                        if false {
                            for fi18 := 0; fi18 < 18; fi18++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v19 uint8 = 2
                                _ = fi18
                                _ = v19
                            }
                        }
                        v6 = v6
                        switch -372 {
                        case 1:
                            v6 += (^(v6))
                        case -195:
                            v2 += fi17
                        }
                        _ = fi17
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = (^(v6))
                    }
                }
            }
        }
        if false {
            for fi20 := 0; fi20 < 11; fi20++ {
                v2 = v8
                v6 = (^(v6))
                _ = fi20
            }
            _ = len(v3)
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v7 -= (v7 | (v7 % (((-(v7))) | 1)))
        for fi22 := 0; fi22 < 8; fi22++ {
            if true {
                v2 -= 4294967295
                v2 -= (fi21 | v8)
            }
            _ = fi22
        }
        _ = fi21
    }
    if true {
        switch (v2 + (v2 << (((v9 | (v9 | 4294967295))) & 63))) {
        case 4294967295:
            var v23 int = (v8 + -2147483648)
            _ = v23
        case 1:
            fmt.Println(v9, v7, v1)
        }
    } else {
        fmt.Println(v7)
    }
    var v24 uint8 = 92
    v25 := GenS2{F0: 42, F1: uint16(v9), F2: 42}
    v26 := GenS1{F0: 159, F1: (^(v6))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
