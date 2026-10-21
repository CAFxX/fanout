package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int16; F2 int64 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint8, p1 uint32) int32 {
    fmt.Println(p1, p0)
    p0 = (p0 | (p0 - (^(p0))))
    v7 := []int{}
    var v8 uint = 378
    if (true || false) {
        var v9 int8 = 24
        var v10 float32 = ((-2.25 * 1.5) + 1e10)
        var v11 int16 = 172
        _ = v10
        _ = v11
        _ = v9
    } else {
        p1 = (^(p1))
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -283
}
func genFn1(p0 int64) uint32 {
    fmt.Println(p0)
    p0 = 688
    var v12 uint = 4294967295
    var v13 uint8 = 1
    _ = p0
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 42
    var v3 uint8 = 16
    var v4 bool = (!(-897 > 42))
    var v5 int = 128
    var v6 uint64 = 919
    if (!v4) {
        v14 := GenS1{F0: 127, F1: (v3 | (v3 ^ v3)), F2: ((400 != 4294967295) || (!true))}
        _ = v14
    } else {
        v3 = v3
        switch v5 {
        case 0:
            if v4 {
                v15 := GenS2{F0: (v6 * v6), F1: 2147483647}
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 int16 = 42
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
        case 2147483647:
            v4 = v4
        case -842:
            v18 := GenS3{F0: (v6 % ((880) | 1))}
            _ = v18
        default:
            var v19 float32 = (3.14159 / 1.0)
            _ = v19
        }
        v20 := GenS4{F0: (v5 != (v5 - v5)), F1: 42, F2: 1}
        _ = v20
    }
    v2 += int32(v5)
    var v21 int = v5
    var v22 uint8 = 18
    var v23 float32 = 0.0
    var v24 uint64 = (^(v6))
    fmt.Println(v4, v22, v2)
    for fi25 := 0; fi25 < 20; fi25++ {
        if (2147483648 <= genFn1(454)) {
            v26 := GenS1{F0: 42, F1: (v3 - v22), F2: v4}
            v6 = 18446744073709551615
            _ = v26
        } else {
            v27 := GenS3{F0: (v6 % ((890) | 1))}
            _ = v27
        }
        var v28 int8 = 1
        _ = fi25
        _ = v28
    }
    var v29 int32 = (v2 / ((v2) | 1))
    v6 += 383
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2, v6)
    default:
        fmt.Println(v24, v29, v6)
    }
    var v30 bool = false
    if v4 {
        switch v21 {
        case -127:
            v31 := GenS5{F0: v6, F1: 1}
            _ = v31
        case -419:
            v32 := GenS1{F0: v22, F1: v3, F2: v4}
            _ = v32
        case -513:
            switch 2147483647 {
            case 32767:
                v33 := v1[(-(v21))]
                _ = v33
            }
        default:
            for fi34 := 0; fi34 < 15; fi34++ {
                v3 = (v3 ^ (v3 + 128))
                v35 := GenS5{F0: (v6 >> ((42) & 63)), F1: -1}
                var v36 int64 = 9223372036854775807
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        if v30 {
            var v37 int16 = 989
            _ = v37
        } else {
            v38 := []int{v5, v5, (v5 / (((v5 % ((int(v3)) | 1))) | 1))}
            _ = v38
        }
        v2 += v29
    } else {
        v6 -= (v24 << ((255) & 63))
        v2 -= v29
        switch ("" + "") {
        case "a":
            fmt.Println(v23, v30)
        case "hello":
            v4 = (935 < 42)
        case "foo bar":
            for fi39 := 0; fi39 < 7; fi39++ {
                v40 := GenS6{F0: 255}
                _ = fi39
                _ = v40
            }
        }
    }
    v29 = (v2 | -65)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
