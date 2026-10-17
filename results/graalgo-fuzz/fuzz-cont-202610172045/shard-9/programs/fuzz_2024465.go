package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int32; F1 int16; F2 bool }

func main() {
    var v1 uint8 = 1
    v2 := []uint32{544, 2}
    v3 := []int{127, -134, 2147483647}
    var v4 uint16 = 2
    v5 := GenS1{F0: v1, F1: 1, F2: true}
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 bool = true
        _ = fi6
        _ = v7
    }
    v1 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        switch (fi8 ^ -130) {
        case 250:
            switch fi8 {
            case 416:
                if true {
                    if (!(!true)) {
                        for fi9 := 0; fi9 < 20; fi9++ {
                            var v10 float64 = 0.0
                            var v11 int64 = 581
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    }
                } else {
                    v12 := GenS1{F0: 231, F1: 279, F2: true}
                    _ = v12
                }
            case -2147483648:
                var v13 bool = true
                _ = v13
            default:
                fmt.Println(v1)
            }
        case 127:
            if false {
                v14 := GenS2{F0: (v4 / (((^(v4))) | 1))}
                _ = v14
            }
        case 1:
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 int64 = 1
                var v17 float64 = float64(0.1)
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        v18 := GenS2{F0: (v4 >> ((18446744073709551615) & 63))}
        _ = fi8
        _ = v18
    }
    fmt.Println(v1, v1)
    if (!(6 >= -1)) {
        v1 = (^(v1))
        switch "\n\t" {
        case "hello":
            v19 := GenS1{F0: v1, F1: 981, F2: ((v4 <= v4) || (494 > 65535))}
            _ = v19
        case "a":
            _ = len(v3)
        }
    } else {
        v1 -= uint8(v4)
        if (false || true) {
            switch ("\n\t" + "hello") {
            case "hello":
                if (!true) {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        var v21 uint64 = 130
                        if false {
                            var v22 uint32 = 1
                            var v23 int32 = -2147483648
                            _ = v22
                            _ = v23
                        }
                        _ = fi20
                        _ = v21
                    }
                }
            case "a":
                v24 := GenS3{F0: -894}
                _ = v24
            case "foo bar":
                v25 := GenS4{F0: 274, F1: 1}
                _ = v25
            }
        }
    }
    v26 := GenS4{F0: 1, F1: 42}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = 42
    for fi27 := 0; fi27 < 14; fi27++ {
        if false {
            if true {
                fmt.Println(fi27, fi27)
            }
            var v28 int16 = -1
            var v29 uint32 = 164
            _ = v28
            _ = v29
        } else {
            v30 := GenS5{F0: 505, F1: int16(fi27), F2: false}
            _ = v30
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
