package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32; F1 int; F2 int8 }
func genFn0(p0 int16) uint16 {
    p0 = p0
    var v7 uint64 = 42
    v8 := GenS1{F0: 0, F1: 4294967295, F2: v7}
    if true {
        fmt.Println(v7, p0)
        p0 -= (p0 + -1)
    }
    p0 += 32767
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 255
    var v3 int = 128
    v4 := make(map[uint32]uint32)
    var v5 int32 = 2147483647
    v6 := []string{"a", "hello"}
    var v9 uint32 = 955
    var v10 uint64 = 42
    var v11 bool = (!(true && false))
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 0, F1: (^(v9)), F2: (v10 / (((v10 % ((150) | 1))) | 1))}
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 2147483647, F1: (^(v9)), F2: (^(v10))}
    if true {
        var v15 int = v3
        switch 9223372036854775807 {
        case 2147483647:
            for fi16 := 0; fi16 < 13; fi16++ {
                for fi17 := 0; fi17 < 6; fi17++ {
                    fmt.Println(v3, v9)
                    v3 = (fi17 & -2147483648)
                    if true {
                        switch "hello" {
                        case "hello":
                            v18 := GenS1{F0: v5, F1: v9, F2: v10}
                            _ = v18
                        default:
                            fmt.Println(v5, v9)
                        }
                        if len(v6) > 0 {
                            v6[0] = ("\n\t" + "a")
                        }
                        var v19 int16 = 0
                        _ = v19
                    }
                    _ = fi17
                }
                _ = fi16
            }
        case 42:
            var v20 int32 = int32(v10)
            _ = v20
        case 1:
            var v21 uint32 = v9
            _ = v21
        default:
            var v22 uint16 = genFn0(-570)
            _ = v22
        }
        _ = v15
    }
    switch v3 {
    case 4294967295:
        if len(v6) > 0 {
            v6[0] = ("hello" + "hello")
        }
    default:
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 13; fi24++ {
                if true {
                    var v25 uint64 = (v10 & v10)
                    _ = v25
                }
                _ = fi24
            }
            switch 42 {
            case 9223372036854775807:
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 uint = 76
                    switch v3 {
                    case 125:
                        fmt.Println(v11)
                    case 2147483647:
                        var v28 int64 = 0
                        _ = v28
                    case 127:
                        switch "" {
                        case "hello":
                            v2 = (v2 << ((v27) & 63))
                        case "foo bar":
                            v11 = (v11 && v11)
                            var v29 int32 = -610
                            switch fi26 {
                            case 4294967295:
                                v30 := GenS1{F0: (v5 | v29), F1: (v9 ^ v9), F2: v10}
                                _ = v30
                            case 9223372036854775807:
                                if v11 {
                                    fmt.Println(fi23, v11)
                                    if v11 {
                                        _ = len(v6)
                                        if (!false) {
                                            var v31 int = v3
                                            _ = v31
                                        }
                                        fmt.Println(v9)
                                    }
                                    v11 = true
                                    v11 = false
                                }
                            default:
                                v29 = 42
                            }
                            _ = v29
                        default:
                            v5 += v5
                        }
                    }
                    v3 = (fi23 * (fi26 / ((-9223372036854775808) | 1)))
                    _ = fi26
                    _ = v27
                }
            }
            _ = fi23
        }
    }
    delete(v4, v9)
    var v32 int32 = (v5 - (v5 * v5))
    v9 = (v9 << (((v2 * (v2 % (((v2 % (((v2 << (((v2 & (v2 + 192))) & 63))) | 1))) | 1)))) & 63))
    var v33 bool = v11
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 float64 = 1.5
        _ = fi34
        _ = v35
    }
    var v36 int8 = int8(v10)
    v37 := GenS2{F0: (v5 & v32), F1: (^(v3)), F2: (v36 % ((v36) | 1))}
    switch v3 {
    case 4294967295:
        for fi38 := 0; fi38 < 18; fi38++ {
            var v39 uint = (v2 >> ((65535) & 63))
            _ = fi38
            _ = v39
        }
    }
    if (104 < 170) {
        for fi40 := 0; fi40 < 8; fi40++ {
            v33 = (4294967295 < (v2 - (v2 + (v2 - 42))))
            switch v3 {
            case 32767:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi40
        }
    }
    var v41 int32 = int32(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v9
}
