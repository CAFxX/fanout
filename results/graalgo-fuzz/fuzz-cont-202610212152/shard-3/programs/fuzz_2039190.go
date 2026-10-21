package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int64; F1 int32; F2 uint8 }
type GenS3 struct { F0 int; F1 uint32; F2 int16 }

func main() {
    var v1 uint16 = 228
    var v2 int8 = 100
    var v3 bool = true
    v4 := []int{0, 32767, 9223372036854775807}
    v5 := GenS1{F0: v3, F1: 1}
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int32 = 872
        _ = fi6
        _ = v7
    }
    fmt.Println(v1, v2)
    fmt.Println(v1)
    for fi8 := 0; fi8 < 12; fi8++ {
        v2 += int8(v1)
        _ = fi8
    }
    switch ("\n\t" + "") {
    case "a":
        v2 -= v2
    case "\n\t":
        for fi9 := 0; fi9 < 4; fi9++ {
            v2 = -1
            if ((0.0 - (0.1 - 0.1)) > (1.5 * 0.1)) {
                if false {
                    fmt.Println(v2)
                    v3 = false
                    if len(v4) > 0 {
                        v4[0] = (fi9 / ((-1) | 1))
                    }
                }
                for fi10 := 0; fi10 < 16; fi10++ {
                    _ = len(v4)
                    v11 := GenS2{F0: 1, F1: 2147483647, F2: 255}
                    if (1 <= 32767) {
                        if len(v4) > 0 {
                            v4[0] = 414
                        }
                        var v12 int16 = 316
                        fmt.Println(fi9, v12, v2)
                        if v3 {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            v2 = v2
                            if ((v1 < 255) && false) {
                                v12 -= (v12 << ((255) & 63))
                            }
                        }
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
            } else {
                _ = len(v4)
            }
            _ = fi9
        }
    default:
        var v13 int16 = -1
        _ = v13
    }
    var v14 int = -2147483648
    v15 := GenS2{F0: -1, F1: 1, F2: 128}
    _ = len(v4)
    for fi16 := 0; fi16 < 9; fi16++ {
        v14 += (v14 - -826)
        for fi17 := 0; fi17 < 15; fi17++ {
            v3 = v3
            var v18 uint16 = v1
            v19 := GenS2{F0: -9223372036854775808, F1: 2147483647, F2: 37}
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 float32 = -2.25
            _ = fi20
            _ = v21
        }
        _ = fi16
    }
    v2 += (v2 - v2)
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 uint64 = 4294967295
            var v25 uint64 = v24
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := GenS2{F0: -9223372036854775808, F1: 0, F2: 92}
        switch "foo bar" {
        case "\n\t":
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 uint32 = 821
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v30 int8 = -1
                v31 := GenS3{F0: -41, F1: (v29 - v29), F2: 32767}
                v4 = append(v4, fi26)
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
        case "foo bar":
            var v32 uint16 = (v1 / ((65535) | 1))
            _ = v32
        case "":
            _ = len(v4)
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
