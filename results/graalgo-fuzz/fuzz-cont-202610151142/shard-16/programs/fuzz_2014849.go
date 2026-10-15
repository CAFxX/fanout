package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint64; F1 int8; F2 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint32 = 472
    v3 := make(map[int]string)
    v2 -= 42
    fmt.Println(v1, v2)
    var v4 int16 = 212
    for fi5 := 0; fi5 < 1; fi5++ {
        switch fi5 {
        case 1:
            v6 := v3[4294967295]
            _ = v6
        case 252:
            if (3.14159 != -2.25) {
                v7 := GenS1{F0: 65535}
                _ = v7
            }
        case 908:
            fmt.Println(fi5)
        }
        var v8 float64 = float64((((3.14159 * 1.5) - 0.1) - 1.5))
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(v4)
            if false {
                var v10 float32 = 1.5
                switch "hello" {
                case "a":
                    var v11 int32 = -1
                    _ = v11
                case "\n\t":
                    var v12 float64 = float64(1e10)
                    _ = v12
                case "hello":
                    if (203 <= v1) {
                        fmt.Println(fi5, v4)
                        v10 += (v10 - 0.1)
                    } else {
                        var v13 int64 = v1
                        switch "" {
                        case "\n\t":
                            if false {
                                if false {
                                    v14 := GenS2{F0: (v13 >> ((912) & 63))}
                                    _ = v14
                                }
                                var v15 int = (^(fi5))
                                v16 := GenS3{F0: (v2 / (((^(v2))) | 1)), F1: -1}
                                _ = v15
                                _ = v16
                            } else {
                                v1 = (v13 * v1)
                                fmt.Println(v2, fi9)
                                v17 := v3[fi5]
                                v10 = (v10 + -2.25)
                                _ = v17
                            }
                        case "foo bar":
                            v18 := []int{(fi9 >> ((65535) & 63)), fi9, (fi9 & 1)}
                            _ = v18
                        case "a":
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v10)
                            default:
                                v19 := GenS1{F0: 831}
                                _ = v19
                            }
                        }
                        var v20 uint64 = 0
                        _ = v13
                        _ = v20
                    }
                default:
                    var v21 float64 = -2.25
                    _ = v21
                }
                v1 -= 1
                _ = v10
            }
            v4 += v4
            var v22 float32 = float32(v8)
            _ = fi9
            _ = v22
        }
        _ = fi5
        _ = v8
    }
    fmt.Println(v2, v2)
    var v23 uint32 = (v2 ^ (^(v2)))
    if (false && (!false)) {
        if (42 <= 245) {
            v23 = (v23 % ((v23) | 1))
            v4 = 32767
            var v24 bool = (false && false)
            _ = v24
        } else {
            var v25 uint16 = 32767
            var v26 uint64 = 4294967295
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 6; fi27++ {
            if ((v4 >> ((129) & 63)) > 1) {
                for fi28 := 0; fi28 < 10; fi28++ {
                    var v29 int64 = (v1 & -434)
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi27
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            v31 := []int{-9223372036854775808, 128, -652, (fi30 & (fi30 % (((fi30 % ((-231) | 1))) | 1)))}
            _ = fi30
            _ = v31
        }
    }
    v32 := GenS4{F0: 42, F1: -25, F2: v4}
    var v33 int16 = 0
    if true {
        v2 = (v2 % ((v23) | 1))
        v34 := []int{}
        switch "foo bar" {
        case "foo bar":
            if (!(true || true)) {
                v1 = (v1 << ((4294967295) & 63))
                v2 += v23
            }
        case "":
            fmt.Println(v1)
        }
        _ = v34
    } else {
        for fi35 := 0; fi35 < 15; fi35++ {
            if (!(!true)) {
                if (0 > 65535) {
                    fmt.Println(v1, fi35)
                }
            }
            fmt.Println(v33, v23)
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 uint = 255
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
