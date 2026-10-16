package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint; F1 int8; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 2147483647
    v3 := []uint32{544, 42, 0}
    v4 := make(map[uint32]bool)
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: 4294967295, F1: 167}
    v2 -= (v2 + v2)
    switch "" {
    case "a":
        v2 = v2
    case "":
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 int16 = 32767
            var v9 float32 = (-2.25 * (float32(1.5) * -2.25))
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        if (1 < 9223372036854775807) {
            fmt.Println(v2, v2)
            if true {
                v2 -= (^(v2))
            } else {
                v2 += -2147483648
                switch fi10 {
                case 32767:
                    v2 = (v2 / ((v2) | 1))
                case -719:
                    fmt.Println(fi10, fi10)
                case 2147483647:
                    var v11 int16 = 1000
                    _ = v11
                default:
                    if (2 >= 2147483647) {
                        v2 = 2147483647
                    } else {
                        v12 := GenS2{F0: 85, F1: 1, F2: v2}
                        _ = v12
                    }
                }
                v13 := GenS2{F0: 42, F1: 42, F2: v2}
                _ = v13
            }
            var v14 uint16 = 32767
            _ = v14
        }
        switch (fi10 | 9223372036854775807) {
        case 939:
            v15 := GenS1{F0: int(v2), F1: 0}
            _ = v15
        }
        _ = fi10
    }
    v2 += (v2 % ((v2) | 1))
    for fi16 := 0; fi16 < 20; fi16++ {
        if ((65535 > 2) && false) {
            var v17 uint = 2
            v2 -= v2
            _ = v17
        }
        _ = fi16
    }
    if ((false && false) && (66 != 42)) {
        if (!(42 == 0)) {
            var v18 int32 = (-(v2))
            var v19 int8 = -128
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v18
            _ = v19
        }
        switch "foo bar" {
        case "foo bar":
            v2 -= v2
        }
    }
    fmt.Println(v2)
    if false {
        for fi20 := 0; fi20 < 12; fi20++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v21 := GenS3{F0: 65535, F1: 49, F2: -9223372036854775808}
            _ = fi20
            _ = v21
        }
        v2 = v2
        if len(v3) > 0 {
            v3[0] = 2
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 = 42
        v2 = 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
