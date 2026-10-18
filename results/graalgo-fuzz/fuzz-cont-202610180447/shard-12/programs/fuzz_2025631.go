package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 int }
type GenS3 struct { F0 int; F1 int; F2 uint16 }
func genFn0() int64 {
    v7 := []int{4294967295, -437, 128}
    var v8 uint64 = 410
    v8 = v8
    switch 202 {
    case -1:
        v8 = 42
    case -909:
        if true {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 bool = (true && (0 == -650))
                var v11 uint16 = 278
                var v12 uint = 1
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            v8 -= (v8 >> ((65535) & 63))
            fmt.Println(v8)
        } else {
            switch "\n\t" {
            case "a":
                fmt.Println(v8)
            case "\n\t":
                if len(v7) > 0 {
                    v7[0] = 909
                }
            case "hello":
                var v13 uint8 = 255
                _ = v13
            default:
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS1{F0: 9223372036854775807, F1: ("\n\t" + "\n\t"), F2: 360}
                    for fi16 := 0; fi16 < 10; fi16++ {
                        fmt.Println(fi16, v8)
                        var v17 float32 = (0.0 / ((((-2.25 - (0.1 + ((1e10 + float32(float64(((0.1 / 1.0) / -2.25)))) * ((1.5 - 1e10) * 0.0)))) - 0.1) + (1.5 * (3.14159 * float32(float64((-2.25 * 3.14159)))))) + 0.0))
                        v7 = append(v7, fi14)
                        _ = fi16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
            }
            v8 -= v8
            if len(v7) > 0 {
                v7[0] = 69
            }
        }
    case 2147483647:
        v8 += (v8 | 2)
    default:
        var v18 int8 = 100
        _ = v18
    }
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -906
    var v3 uint = 42
    var v4 uint32 = 65535
    v5 := []int{-2147483648, 0, 32768, 128}
    v6 := make(map[uint32]string)
    var v19 int32 = 42
    fmt.Println(v2)
    switch int(v2) {
    case 32768:
        switch "" {
        case "\n\t":
            if true {
                for fi20 := 0; fi20 < 20; fi20++ {
                    for fi21 := 0; fi21 < 6; fi21++ {
                        switch "\n\t" {
                        case "\n\t":
                            var v22 uint64 = 0
                            _ = v22
                        default:
                            var v23 int32 = v19
                            _ = v23
                        }
                        v24 := GenS2{F0: 42, F1: (v3 * v3), F2: fi20}
                        switch (fi20 & fi20) {
                        case 4294967295:
                            var v25 int = 42
                            _ = v25
                        default:
                            var v26 int = fi21
                            _ = v26
                        }
                        var v27 int8 = 0
                        _ = fi21
                        _ = v24
                        _ = v27
                    }
                    _ = fi20
                }
                for fi28 := 0; fi28 < 20; fi28++ {
                    fmt.Println(fi28)
                    var v29 float32 = 0.0
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi28
                    _ = v29
                }
                var v30 float64 = 0.0
                _ = v30
            } else {
                v31 := GenS2{F0: 314, F1: (v3 * (v3 % (((v3 << ((v3) & 63))) | 1))), F2: 913}
                for fi32 := 0; fi32 < 6; fi32++ {
                    var v33 int16 = -76
                    v3 += 1
                    _ = fi32
                    _ = v33
                }
                _ = v31
            }
        }
    case 42:
        _ = len(v5)
    }
    _ = len(v5)
    v34 := GenS3{F0: 2147483647, F1: 0, F2: 32767}
    fmt.Println(v2)
    v2 = -38
    for fi35 := 0; fi35 < 15; fi35++ {
        var v36 float32 = (3.14159 * 1e10)
        v36 += -2.25
        _ = fi35
        _ = v36
    }
    if false {
        if true {
            fmt.Println(v4)
            var v37 float64 = float64(((1.5 * 0.0) - 1.5))
            if false {
                fmt.Println(v4, v37, v37)
                fmt.Println(v2)
            } else {
                v19 += (v19 % ((v19) | 1))
            }
            _ = v37
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v38 uint32 = (v4 % ((v4) | 1))
    var v39 string = "hello"
    switch 0 {
    case 4294967295:
        v2 = v2
    default:
        var v40 uint = v3
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
