package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 string; F1 string; F2 uint64 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 uint; F1 uint }
type GenS11 struct { F0 uint; F1 uint32; F2 string }
func genFn0(p0 uint64) bool {
    var v4 uint32 = 150
    v5 := []int{-319}
    p0 = p0
    var v6 uint = 4294967295
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return ((0 == -9223372036854775808) && (!(false && false)))
}

func main() {
    var v1 uint32 = 65535
    var v2 uint32 = 65535
    var v3 int = -9223372036854775808
    v7 := GenS1{F0: 1, F1: 2}
    if true {
        var v8 int16 = -330
        var v9 string = "hello"
        _ = v8
        _ = v9
    } else {
        var v10 int64 = 42
        v11 := GenS1{F0: 4294967295, F1: 4294967295}
        if (572 == 755) {
            v12 := GenS2{F0: 40}
            _ = v12
        } else {
            v10 += (-(v10))
            v13 := GenS3{F0: (v2 % ((v1) | 1)), F1: (-(v3)), F2: 42}
            switch v3 {
            case 4294967295:
                if true {
                    v14 := GenS4{F0: 1}
                    v1 = (v1 - v1)
                    var v15 int = (v3 >> ((42) & 63))
                    _ = v14
                    _ = v15
                }
            case 32767:
                if false {
                    fmt.Println(v3, v2)
                    fmt.Println(v2)
                }
            }
            _ = v13
        }
        _ = v10
        _ = v11
    }
    switch v3 {
    case 0:
        fmt.Println(v1, v1, v2)
    case 127:
        v16 := []int{-1, v3, 128, (v3 % ((v3) | 1))}
        _ = v16
    }
    var v17 float64 = 0.1
    v18 := GenS5{F0: v3}
    if (!true) {
        var v19 uint = 1
        v20 := GenS6{F0: -527}
        _ = v19
        _ = v20
    }
    if (!true) {
        v3 -= v3
        v3 = v3
    } else {
        v21 := GenS3{F0: (v1 | (v2 % (((v1 + (^(v2)))) | 1))), F1: (-(v3)), F2: 18446744073709551615}
        var v22 int64 = 42
        switch v3 {
        case -489:
            var v23 uint8 = 1
            _ = v23
        case 32767:
            v3 = v3
        case 9223372036854775807:
            switch (v3 ^ (v3 ^ (v3 * v3))) {
            case -368:
                if (!false) {
                    v2 = v1
                    v24 := GenS7{F0: (v1 << ((18446744073709551615) & 63)), F1: 42}
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 12; fi25++ {
                        v3 = 9223372036854775807
                        var v26 int64 = (-(v22))
                        _ = fi25
                        _ = v26
                    }
                }
            case 2147483647:
                v1 = 209
            case 4294967295:
                switch "" {
                case "hello":
                    v27 := GenS8{F0: "hello", F1: "", F2: 18446744073709551615}
                    _ = v27
                case "foo bar":
                    for fi28 := 0; fi28 < 9; fi28++ {
                        v22 = v22
                        _ = fi28
                    }
                }
            default:
                v2 += v1
            }
        }
        _ = v21
        _ = v22
    }
    v29 := []int{v3, -247}
    for fi30 := 0; fi30 < 11; fi30++ {
        for fi31 := 0; fi31 < 17; fi31++ {
            if len(v29) > 0 {
                v29[0] = fi31
            }
            _ = fi31
        }
        v1 = 65535
        _ = fi30
    }
    v32 := GenS9{F0: 1000, F1: 1}
    v1 = 702
    fmt.Println(v2, v17)
    for fi33 := 0; fi33 < 13; fi33++ {
        if ((v2 ^ v1) < 762) {
            var v34 string = ("\n\t" + ("" + "hello"))
            _ = v34
        } else {
            for fi35 := 0; fi35 < 20; fi35++ {
                switch -9223372036854775808 {
                case 32768:
                    _ = len(v29)
                    v2 -= 2147483647
                    fmt.Println(v2, fi33)
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                    fmt.Println(v1)
                    v36 := GenS10{F0: 300, F1: 2}
                    v37 := GenS11{F0: 678, F1: (v2 / ((v2) | 1)), F2: (("hello" + "a") + (("a" + "a") + "\n\t"))}
                    _ = v36
                    _ = v37
                default:
                    v3 -= (fi35 & (v3 / ((fi35) | 1)))
                }
                var v38 uint32 = 0
                _ = fi35
                _ = v38
            }
        }
        _ = fi33
    }
    var v39 uint8 = 61
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v39
    _ = v7
}
