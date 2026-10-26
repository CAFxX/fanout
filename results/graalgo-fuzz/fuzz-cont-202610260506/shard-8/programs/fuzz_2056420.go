package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
type GenS2 struct { F0 uint; F1 bool; F2 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int64; F1 int; F2 int16 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 bool; F1 int64; F2 int64 }
func genFn0(p0 uint8, p1 uint, p2 uint) bool {
    v4 := []int{-2147483648}
    fmt.Println(p2)
    p1 += 65535
    if (!true) {
        switch -869 {
        case 128:
            if (3.14159 <= 1.5) {
                v5 := GenS1{F0: 4294967295, F1: 1, F2: p1}
                p0 = (p0 + p0)
                _ = v5
            } else {
                v4 = append(v4, 128)
                p2 = p1
            }
        default:
            fmt.Println(p1)
        }
        var v6 int = -117
        _ = v6
    } else {
        p0 += p0
        p2 += p1
    }
    switch ((("a" + "") + "hello") + "a") {
    case "\n\t":
        var v7 int32 = -502
        _ = v7
    case "":
        if len(v4) > 0 {
            v4[0] = 260
        }
    default:
        v8 := GenS2{F0: (p1 | p1), F1: true, F2: -1}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return false
}

func main() {
    v1 := []int{}
    var v2 float32 = (1e10 - -2.25)
    var v3 int = -319
    var v9 int8 = 0
    switch v3 {
    case 42:
        var v10 float32 = (-2.25 - (1.5 * (1.5 + (v2 * (0.1 / -2.25)))))
        _ = v10
    case 9223372036854775807:
        var v11 uint = 689
        _ = v11
    }
    v12 := GenS1{F0: (v3 * (v3 % (((v3 / (((v3 & v3)) | 1))) | 1))), F1: 32767, F2: 1}
    v13 := GenS3{F0: 32767, F1: v9}
    if (1 >= -490) {
        if (447 < 42) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v2)
            v2 -= float32(0.0)
        }
        switch -843 {
        case 1:
            for fi14 := 0; fi14 < 8; fi14++ {
                v3 = (fi14 << ((4294967295) & 63))
                if (v2 == float32((0.0 / 1.0))) {
                    var v15 uint64 = 1
                    if true {
                        switch "hello" {
                        case "hello":
                            v2 = 0.1
                        case "foo bar":
                            if (true && false) {
                                var v16 uint32 = 129
                                _ = v16
                            }
                        }
                        fmt.Println(v9)
                    }
                    switch "hello" {
                    case "":
                        for fi17 := 0; fi17 < 11; fi17++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v18 := GenS3{F0: 402, F1: 42}
                            _ = fi17
                            _ = v18
                        }
                    }
                    _ = v15
                }
                if (0 <= 785) {
                    switch 807 {
                    case 168:
                        if true {
                            switch "foo bar" {
                            case "hello":
                                v2 = 0.0
                            }
                            fmt.Println(v3, v9)
                            v19 := GenS3{F0: 0, F1: v9}
                            _ = v19
                        } else {
                            v3 = v3
                            if len(v1) > 0 {
                                v1[0] = 60
                            }
                            v20 := GenS3{F0: 42, F1: v9}
                            _ = v20
                        }
                    case 771:
                        if (-157 == 1) {
                            v21 := GenS4{F0: -1, F1: fi14, F2: 42}
                            _ = v21
                        }
                    }
                    v2 = v2
                }
                _ = fi14
            }
        default:
            if true {
                var v22 uint32 = 563
                _ = v22
            }
        }
    } else {
        _ = len(v1)
    }
    if false {
        switch "hello" {
        case "a":
            v23 := GenS5{F0: 2, F1: v9}
            _ = v23
        }
        v3 += v3
    }
    if (-822 == 956) {
        fmt.Println(v2, v2)
        var v24 float64 = 3.14159
        v25 := GenS6{F0: true, F1: 872, F2: -9223372036854775808}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v9 = (v9 * v9)
        _ = fi26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
