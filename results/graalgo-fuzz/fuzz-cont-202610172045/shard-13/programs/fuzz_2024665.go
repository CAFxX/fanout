package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int32 }
type GenS2 struct { F0 bool; F1 bool }

func main() {
    var v1 bool = (!false)
    var v2 uint16 = 174
    v3 := make(map[int]string)
    v4 := []int{-832}
    v5 := []int{}
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    v6 := GenS1{F0: 4294967295, F1: 945, F2: -238}
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int32 = 42
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(v1, v2)
    fmt.Println(v1, v1)
    v10 := GenS2{F0: v1, F1: (18446744073709551615 <= 1)}
    if ("a" > (("hello" + "a") + (("a" + "\n\t") + "hello"))) {
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS1{F0: 4294967295, F1: 1, F2: -2147483648}
            _ = fi11
            _ = v12
        }
        switch "a" {
        case "foo bar":
            if v1 {
                v13 := GenS2{F0: false, F1: false}
                _ = v13
            }
        case "":
            v1 = v1
        case "hello":
            fmt.Println(v1, v2)
        default:
            if len(v4) > 0 {
                v4[0] = 572
            }
        }
        fmt.Println(v1)
    } else {
        var v14 int64 = -9223372036854775808
        if ((-32 == -8) && v1) {
            for fi15 := 0; fi15 < 6; fi15++ {
                for fi16 := 0; fi16 < 17; fi16++ {
                    switch 128 {
                    case 2147483647:
                        v2 += (v2 / (((v2 ^ 65535)) | 1))
                    case -1:
                        switch fi16 {
                        case -672:
                            var v17 int16 = 0
                            _ = v17
                        case 32768:
                            v5 = append(v5, (fi15 % ((fi15) | 1)))
                        default:
                            fmt.Println(v1)
                        }
                    case 32767:
                        var v18 string = ("\n\t" + "a")
                        _ = v18
                    }
                    fmt.Println(fi15, fi15, v1)
                    _ = fi16
                }
                if len(v4) > 0 {
                    v4[0] = fi15
                }
                v1 = (1 >= 2147483647)
                _ = fi15
            }
        } else {
            v2 -= v2
            v2 -= v2
        }
        delete(v3, -2147483648)
        _ = v14
    }
    v1 = v1
    v19 := GenS1{F0: 739, F1: 4294967295, F2: 1}
    v4 = append(v4, -882)
    switch -368 {
    case 4294967295:
        for fi20 := 0; fi20 < 20; fi20++ {
            v1 = (32767 <= 776)
            _ = fi20
        }
    case 127:
        fmt.Println(v1, v2)
    case -9223372036854775808:
        v2 += (v2 + 32767)
    default:
        if false {
            fmt.Println(v2)
            var v21 uint8 = 0
            if (-1 == 444) {
                v22 := GenS1{F0: 42, F1: 1, F2: -1}
                v1 = (((!(1.5 < 0.1)) && (v1 || (4294967295 >= 215))) || (uint32(v2) < 2147483648))
                _ = v22
            }
            _ = v21
        } else {
            v23 := GenS2{F0: (float32(3.14159) <= 1.5), F1: false}
            fmt.Println(v1, v2)
            var v24 int8 = 0
            _ = v23
            _ = v24
        }
    }
    if v1 {
        fmt.Println(v1, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
