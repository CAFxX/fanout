package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint64; F1 string; F2 uint64 }

func main() {
    var v1 float32 = (((0.1 / 0.1) * 0.0) / 1.5)
    var v2 uint16 = 1
    var v3 int64 = 0
    if false {
        switch 0 {
        case 119:
            switch 127 {
            case -58:
                var v4 uint32 = 65535
                _ = v4
            case -660:
                v1 += v1
            }
        case 9223372036854775807:
            switch 2147483647 {
            case 32768:
                fmt.Println(v2, v2, v3)
            case 1:
                var v5 uint8 = 1
                _ = v5
            case 4294967295:
                v6 := GenS1{F0: -1, F1: 42}
                _ = v6
            }
        case -32:
            v1 -= v1
        }
        v7 := GenS1{F0: 51, F1: 42}
        fmt.Println(v3, v2)
        _ = v7
    } else {
        v8 := []int{32767, 847}
        _ = v8
    }
    var v9 int32 = 1
    v9 = v9
    v10 := GenS2{F0: "hello", F1: 1}
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int16 = 572
        fmt.Println(v3, v3, v12)
        v1 -= v1
        _ = fi11
        _ = v12
    }
    var v13 uint = 1
    if (!(0 < 42)) {
        if (true && (226 != 42)) {
            var v14 uint32 = 317
            _ = v14
        } else {
            v15 := GenS2{F0: ("\n\t" + "a"), F1: v9}
            v16 := []int{127, 127, 42, 42}
            v1 += 0.1
            _ = v15
            _ = v16
        }
    } else {
        v17 := GenS3{F0: 126}
        _ = v17
    }
    v13 -= v13
    if (-128 >= -128) {
        fmt.Println(v13, v9, v1)
        fmt.Println(v3, v1)
        v18 := []int{-375}
        _ = v18
    } else {
        v2 += (v2 * v2)
        v19 := []int{32768, int(v13)}
        _ = v19
    }
    fmt.Println(v9, v3)
    v20 := GenS4{F0: 2, F1: -1, F2: 1}
    v9 += -153
    var v21 float64 = 0.1
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int8 = 111
        for fi24 := 0; fi24 < 4; fi24++ {
            switch fi22 {
            case 2147483647:
                if (8 > uint32(v2)) {
                    v1 = 1.5
                    v2 -= (v2 ^ (v2 + v2))
                    var v25 float64 = -2.25
                    _ = v25
                } else {
                    if (((!false) || false) || true) {
                        v26 := GenS5{F0: 18446744073709551615, F1: ("hello" + ""), F2: 4294967295}
                        _ = v26
                    }
                    var v27 float32 = float32(v21)
                    v28 := []int{(fi24 / ((fi22) | 1)), (fi24 - fi24), fi22, fi22}
                    _ = v27
                    _ = v28
                }
            }
            switch int(v23) {
            case 4294967295:
                v13 -= v13
            case -1:
                v29 := []int{(fi22 + -9223372036854775808), (fi22 >> ((v13) & 63))}
                _ = v29
            case -458:
                v13 = (v13 << (((v13 | v13)) & 63))
            }
            _ = fi24
        }
        var v30 float64 = v21
        _ = fi22
        _ = v23
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v9
}
