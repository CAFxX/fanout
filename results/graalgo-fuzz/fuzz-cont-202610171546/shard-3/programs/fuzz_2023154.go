package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int64; F1 int; F2 int64 }
type GenS6 struct { F0 uint16; F1 bool }
func genFn0(p0 int) bool {
    var v4 uint = 42
    v5 := GenS1{F0: 0, F1: 42}
    var v6 uint64 = 1
    var v7 int16 = 49
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 int) int32 {
    p0 = p0
    p0 = (p0 << ((303) & 63))
    if (!(false || false)) {
        switch (p0 >> ((858) & 63)) {
        case 4294967295:
            p0 = (p0 >> ((65535) & 63))
        }
        v8 := GenS1{F0: 1, F1: -32768}
        _ = v8
    } else {
        p0 = p0
        v9 := GenS2{F0: (701 == 4294967295)}
        v10 := GenS1{F0: 2147483647, F1: -32768}
        _ = v10
        _ = v9
    }
    if false {
        fmt.Println(p0)
        v11 := GenS3{F0: 1, F1: 65535, F2: 2}
        var v12 uint8 = 159
        _ = v11
        _ = v12
    }
    _ = p0
    return 756
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 42
    fmt.Println(v3)
    v3 -= v3
    fmt.Println(v3)
    var v13 int8 = 0
    v14 := []int{-837, -9223372036854775808, 42, 128}
    var v15 int64 = -1
    if (((1 != 280) || (!true)) || true) {
        var v16 int32 = -2147483648
        var v17 uint16 = 465
        _ = v16
        _ = v17
    }
    v3 = 42
    var v18 int32 = 36
    var v19 uint = 378
    var v20 int = 4294967295
    for fi21 := 0; fi21 < 6; fi21++ {
        if (1e10 == ((0.1 - float32(0.0)) + (1.5 - 3.14159))) {
            v22 := GenS4{F0: v3, F1: 255, F2: 193}
            v13 += v13
            _ = v22
        } else {
            if ((899 <= (fi21 | 1)) && false) {
                v15 = 0
                var v23 int8 = v13
                _ = v23
            } else {
                v18 += genFn1(fi21)
                fmt.Println(v18)
                delete(v2, 2147483647)
            }
            v3 -= v3
            if true {
                v18 = v18
                var v24 int = (-(fi21))
                v15 -= int64(v13)
                _ = v24
            }
        }
        v25 := GenS5{F0: (v15 << ((2) & 63)), F1: (v20 & (v20 | (fi21 * fi21))), F2: 883}
        var v26 uint = (v19 % ((v19) | 1))
        _ = fi21
        _ = v25
        _ = v26
    }
    var v27 uint = (v19 >> (((v19 + (v19 & v19))) & 63))
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 int32 = (v18 | (v18 << (((^(v19))) & 63)))
        v27 -= v19
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        v20 = fi30
        var v31 int16 = 1000
        v19 -= v19
        _ = fi30
        _ = v31
    }
    switch ("hello" + ("\n\t" + ("" + ""))) {
    case "a":
        if (false && false) {
            v15 = (v15 % (((v15 << ((v27) & 63))) | 1))
            switch ("a" + "a") {
            case "a":
                switch v20 {
                case 122:
                    v20 = v20
                case 1:
                    if (((^(v18)) != genFn1(v20)) || (false && ((!true) && true))) {
                        if false {
                            v19 += (v27 | (v27 - uint(v3)))
                        } else {
                            v27 -= (v19 << (((v27 + (^(v27)))) & 63))
                            v32 := GenS6{F0: 32767, F1: true}
                            _ = v32
                        }
                    } else {
                        v15 += (v15 * v15)
                        var v33 int = (v20 ^ v20)
                        v27 -= (v27 % ((v27) | 1))
                        _ = v33
                    }
                }
            case "\n\t":
                v18 += 0
            }
            v20 = 9223372036854775807
        }
    case "\n\t":
        var v34 float32 = 1e10
        _ = v34
    default:
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 uint16 = 255
            v19 += uint(v3)
            var v37 uint32 = 1
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    v3 = (v3 - (v3 % (((v3 >> ((uint(v3)) & 63))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
}
