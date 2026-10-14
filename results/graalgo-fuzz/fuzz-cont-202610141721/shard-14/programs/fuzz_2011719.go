package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint8; F1 uint; F2 uint32 }

func main() {
    var v1 int16 = -745
    var v2 float64 = 0.0
    var v3 int16 = v1
    var v4 int16 = (v3 % ((v3) | 1))
    v5 := make(map[uint32]int)
    var v6 int32 = -729
    var v7 bool = true
    for fi8 := 0; fi8 < 8; fi8++ {
        v3 -= (v1 ^ v4)
        _ = fi8
    }
    var v9 uint8 = 42
    switch -9223372036854775808 {
    case 128:
        switch "foo bar" {
        case "hello":
            if (!v7) {
                if v7 {
                    v10 := GenS1{F0: 2}
                    v11 := GenS2{F0: -1, F1: (v1 % ((1) | 1)), F2: 769}
                    v12 := GenS1{F0: 255}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS2{F0: 443, F1: (v3 * (v3 * int16(v6))), F2: 1}
                    _ = v13
                }
                v14 := GenS1{F0: 2}
                _ = v14
            } else {
                if (65535 < 974) {
                    v15 := GenS2{F0: -577, F1: v4, F2: 931}
                    switch 4294967295 {
                    case 1:
                        v16 := GenS3{F0: 4294967295, F1: false}
                        _ = v16
                    default:
                        var v17 uint32 = uint32(v9)
                        _ = v17
                    }
                    fmt.Println(v4)
                    _ = v15
                }
                v18 := GenS1{F0: 42}
                _ = v18
            }
        }
    default:
        v19 := GenS1{F0: 0}
        _ = v19
    }
    v20 := GenS1{F0: 0}
    v2 -= 1.5
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 bool = (((-2.25 * ((1.5 - -2.25) + 0.1)) < 1.5) && true)
        v23 := v5[785]
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 uint8 = 255
    if false {
        delete(v5, 42)
        fmt.Println(v7)
    } else {
        v6 = (v6 + v6)
        if (-525 != int64(v4)) {
            for fi25 := 0; fi25 < 15; fi25++ {
                v26 := []int{(fi25 << ((18446744073709551615) & 63)), fi25}
                _ = fi25
                _ = v26
            }
            if false {
                v27 := GenS3{F0: 639, F1: (false && ((^(v9)) >= 123))}
                if (("foo bar" == "\n\t") || v7) {
                    v28 := v5[42]
                    _ = v28
                } else {
                    if ((^(v24)) >= (v24 << ((4294967295) & 63))) {
                        v5[1] = 9223372036854775807
                        var v29 float64 = (v2 * v2)
                        _ = v29
                    }
                    var v30 uint16 = 467
                    v31 := []int{2147483647, -350}
                    _ = v30
                    _ = v31
                }
                _ = v27
            } else {
                var v32 int64 = -9223372036854775808
                var v33 int32 = v6
                var v34 uint16 = 42
                _ = v32
                _ = v33
                _ = v34
            }
        }
        var v35 int32 = (v6 | v6)
        _ = v35
    }
    if (1 != 255) {
        v5[4294967295] = 258
        var v36 int32 = v6
        _ = v36
    }
    var v37 uint64 = 581
    if v7 {
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v6 -= (-(v6))
        }
        v7 = v7
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        v37 = (^(v37))
        _ = fi38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        var v40 int = int(v9)
        _ = fi39
        _ = v40
    }
    v41 := GenS4{F0: 2, F1: 973, F2: 0}
    for fi42 := 0; fi42 < 3; fi42++ {
        v37 -= v37
        switch fi42 {
        case 42:
            v4 = (v1 >> ((343) & 63))
        }
        switch "foo bar" {
        case "":
            var v43 int32 = (-(v6))
            _ = v43
        case "hello":
            if ((fi42 << ((42) & 63)) != fi42) {
                if (((2147483648 >= 489) && ((v37 / ((217) | 1)) >= v37)) || (!v7)) {
                    var v44 int64 = 9223372036854775807
                    _ = v44
                } else {
                    v45 := v5[65535]
                    _ = v45
                }
                var v46 uint64 = (v37 << ((65535) & 63))
                _ = len(v5)
                _ = v46
            }
        default:
            v2 += ((v2 - (v2 * v2)) * float64(1e10))
        }
        _ = fi42
    }
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
