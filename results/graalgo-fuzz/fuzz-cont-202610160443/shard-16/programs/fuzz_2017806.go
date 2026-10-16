package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 string }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0() string {
    var v4 uint64 = 2
    if false {
        if (-1 > 132) {
            v4 += 4294967295
            v4 -= 16
            fmt.Println(v4)
        }
    }
    _ = v4
    return ""
}
func genFn1(p0 int32) int {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int32 = (-(p0))
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 string = ""
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 bool = false
        var v11 int16 = 466
        v11 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(p0)
    if false {
        switch "foo bar" {
        case "":
            p0 += p0
        }
        p0 -= (p0 / ((p0) | 1))
        fmt.Println(p0)
    } else {
        fmt.Println(p0)
        if (128 == 73) {
            fmt.Println(p0)
        } else {
            v12 := []int{}
            var v13 int32 = p0
            _ = v12
            _ = v13
        }
    }
    v14 := []int{32768, -889}
    _ = p0
    _ = v14
    return 0
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{42, 594}
    var v3 uint64 = 864
    if true {
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 uint64 = 42
            v17 := GenS1{F0: 2147483648, F1: 1, F2: (v1 + v1)}
            _ = fi15
            _ = v16
            _ = v17
        }
        if false {
            var v18 uint32 = 120
            _ = v18
        } else {
            if ((0 != 0) || false) {
                for fi19 := 0; fi19 < 4; fi19++ {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        if false {
                            var v21 bool = (369 <= 909)
                            v3 += v3
                            v21 = v21
                            v22 := GenS2{F0: 32767, F1: 42}
                            _ = v21
                            _ = v22
                        }
                        switch (fi19 << ((255) & 63)) {
                        case -1:
                            v3 = (v3 ^ v3)
                        case -505:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 127:
                            v3 += v3
                            v1 = v1
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = fi20
                    }
                    _ = fi19
                }
                var v23 int8 = 0
                _ = v23
            } else {
                var v24 float64 = ((3.14159 - float64(0.1)) * 0.1)
                _ = v24
            }
        }
    }
    var v25 int8 = -128
    v2 = append(v2, 318)
    for fi26 := 0; fi26 < 14; fi26++ {
        v25 += (v25 / (((v25 + v25)) | 1))
        v27 := GenS1{F0: 434, F1: -1, F2: "a"}
        var v28 uint64 = (v3 % ((v3) | 1))
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        fmt.Println(v1, v25)
        switch v1 {
        case "a":
            var v30 uint16 = 0
            _ = v30
        case "":
            v3 -= v3
        default:
            var v31 uint8 = 127
            _ = v31
        }
        _ = fi29
    }
    if false {
        for fi32 := 0; fi32 < 6; fi32++ {
            v1 = v1
            _ = fi32
        }
    }
    if len(v2) > 0 {
        v2[0] = 545
    }
    v3 += 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
}
