package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint64 }

func main() {
    var v1 int = 128
    var v2 uint8 = 2
    v3 := []int{253, (v1 & (v1 | v1))}
    var v4 uint = 18446744073709551615
    var v5 int64 = -924
    if true {
        if (-2.25 == 0.0) {
            if ((v1 | 1) != v1) {
                v6 := GenS1{F0: -9223372036854775808, F1: 568}
                _ = v6
            }
        }
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS2{F0: 32767, F1: (v2 + (v2 >> (((v4 ^ v4)) & 63)))}
                _ = fi8
                _ = v9
            }
            _ = len(v3)
            if false {
                fmt.Println(v4)
                fmt.Println(fi7, v5)
                for fi10 := 0; fi10 < 17; fi10++ {
                    v5 += (^(v5))
                    var v11 int32 = 1
                    if false {
                        v3 = append(v3, (fi10 & -9223372036854775808))
                        v2 = (^(v2))
                        var v12 uint64 = 671
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi7
        }
        v3 = append(v3, v1)
    } else {
        if len(v3) > 0 {
            v3[0] = (v1 * int(v4))
        }
        v3 = append(v3, v1)
        var v13 uint64 = 0
        _ = v13
    }
    v5 += (^(v5))
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(fi15, fi15, v2)
            switch "hello" {
            case "foo bar":
                v3 = append(v3, fi15)
            default:
                v16 := GenS3{F0: 1, F1: v5, F2: uint64(v4)}
                _ = v16
            }
            _ = fi15
        }
        var v17 uint16 = 65535
        _ = fi14
        _ = v17
    }
    _ = len(v3)
    if (1 <= -632) {
        var v18 string = "a"
        _ = v18
    } else {
        var v19 int8 = -128
        if false {
            for fi20 := 0; fi20 < 11; fi20++ {
                for fi21 := 0; fi21 < 16; fi21++ {
                    v1 = 0
                    v5 -= (^(v5))
                    if len(v3) > 0 {
                        v3[0] = fi21
                    }
                    _ = fi21
                }
                fmt.Println(fi20, v2, v5)
                _ = fi20
            }
            v3 = append(v3, (v1 << ((4294967295) & 63)))
        } else {
            var v22 string = ""
            _ = v22
        }
        _ = v19
    }
    switch 350 {
    case 1:
        var v23 uint32 = 160
        _ = v23
    case 4294967295:
        if true {
            v5 -= (v5 - 9223372036854775807)
            var v24 int8 = -1
            _ = v24
        }
    case 32768:
        if (9223372036854775807 == 0) {
            v5 += 42
        } else {
            v25 := GenS3{F0: 462, F1: v5, F2: 4294967295}
            switch (v1 << (((v4 % ((v4) | 1))) & 63)) {
            case 32767:
                for fi26 := 0; fi26 < 14; fi26++ {
                    v4 = v4
                    _ = fi26
                }
            default:
                var v27 uint32 = 1
                _ = v27
            }
            _ = len(v3)
            _ = v25
        }
    }
    switch "a" {
    case "a":
        var v28 int32 = 42
        _ = v28
    case "\n\t":
        v29 := GenS2{F0: -979, F1: 255}
        _ = v29
    }
    if false {
        v4 -= (^(v4))
        var v30 uint16 = 65535
        _ = v30
    }
    _ = len(v3)
    var v31 uint = (v4 | 0)
    var v32 int = v1
    v33 := GenS3{F0: 923, F1: (v5 / ((v5) | 1)), F2: 688}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
