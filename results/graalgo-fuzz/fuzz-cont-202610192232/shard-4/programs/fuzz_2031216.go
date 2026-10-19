package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{2147483647, -493}
    var v3 uint8 = 0
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 11; fi5++ {
        v3 += (v3 - v3)
        _ = fi5
    }
    v3 += v3
    v3 = v3
    v3 = v3
    v6 := GenS1{F0: -648, F1: 2147483647, F2: 2}
    switch 9223372036854775807 {
    case 899:
        switch -9223372036854775808 {
        case -2147483648:
            var v7 uint = 65535
            _ = v7
        case 898:
            fmt.Println(v3)
        default:
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS1{F0: fi8, F1: (fi8 | 708), F2: 0}
                v3 = v3
                _ = fi8
                _ = v9
            }
        }
    case 127:
        v3 += 2
    }
    fmt.Println(v3)
    switch 689 {
    case 42:
        v3 += (v3 * v3)
    case -565:
        var v10 int16 = -32768
        _ = v10
    case 32768:
        v11 := GenS1{F0: 42, F1: 4294967295, F2: 32767}
        _ = v11
    default:
        switch 127 {
        case 42:
            v3 -= 255
        case -836:
            v12 := v4[2147483648]
            _ = v12
        case -344:
            v13 := GenS1{F0: 143, F1: 4294967295, F2: 0}
            _ = v13
        default:
            switch -719 {
            case 128:
                v3 += (v3 ^ v3)
            case -9223372036854775808:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
    }
    fmt.Println(v3)
    if ((false && false) || true) {
        var v14 uint32 = 2147483647
        _ = len(v4)
        for fi15 := 0; fi15 < 4; fi15++ {
            v4[v14] = true
            switch 127 {
            case 316:
                v16 := GenS1{F0: -1, F1: 1, F2: 32767}
                _ = v16
            case 42:
                if true {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        _ = len(v1)
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if false {
                                v14 = v14
                                if false {
                                    v14 -= (^(v14))
                                }
                                var v18 uint16 = 2
                                _ = v18
                            } else {
                                v3 -= v3
                            }
                        } else {
                            var v19 string = (("a" + "hello") + ("foo bar" + ("a" + "")))
                            var v20 int8 = -40
                            _ = len(v1)
                            _ = v19
                            _ = v20
                        }
                        _ = fi17
                    }
                    var v21 int16 = 32767
                    _ = v21
                } else {
                    v3 = (v3 | (v3 % ((v3) | 1)))
                }
            default:
                var v22 uint64 = 449
                _ = v22
            }
            _ = fi15
        }
        _ = v14
    }
    var v23 float32 = 0.0
    if (0 > 1) {
        for fi24 := 0; fi24 < 17; fi24++ {
            v3 -= v3
            if len(v2) > 0 {
                v2[0] = 128
            }
            var v25 uint32 = 980
            _ = fi24
            _ = v25
        }
        var v26 float64 = (3.14159 + (1.5 - float64(v23)))
        fmt.Println(v3, v26)
        _ = v26
    } else {
        var v27 uint = 577
        _ = v27
    }
    switch 32767 {
    case 1:
        fmt.Println(v23, v23)
    case 32767:
        if false {
            for fi28 := 0; fi28 < 15; fi28++ {
                fmt.Println(v3, fi28)
                v1[493] = 9223372036854775807
                var v29 uint = 18446744073709551615
                _ = fi28
                _ = v29
            }
        } else {
            var v30 int = -2147483648
            var v31 int64 = 42
            v23 = v23
            _ = v30
            _ = v31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v6
}
