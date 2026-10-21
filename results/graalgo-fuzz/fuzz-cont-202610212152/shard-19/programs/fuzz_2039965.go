package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := []string{}
    var v2 float64 = 3.14159
    v3 := make(map[uint32]bool)
    var v4 int = 128
    var v5 uint16 = 1
    fmt.Println(v2, v5, v4)
    v2 += float64((3.14159 + ((0.1 - ((-2.25 / 1e10) / 1.0)) * 3.14159)))
    v4 += v4
    v4 = (v4 << ((42) & 63))
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 float32 = 3.14159
        for fi8 := 0; fi8 < 7; fi8++ {
            v7 -= float32((v2 / 0.1))
            var v9 uint16 = (v5 % ((v5) | 1))
            var v10 uint32 = 447
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    if (true && false) {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint16 = (v5 | 65535)
            for fi13 := 0; fi13 < 15; fi13++ {
                switch "hello" {
                case "hello":
                    if (false && (1 >= -476)) {
                        v12 = (v12 + 1)
                    }
                case "\n\t":
                    if (false && true) {
                        fmt.Println(v4, fi11)
                    } else {
                        var v14 uint32 = 1
                        v15 := GenS1{F0: true}
                        if false {
                            v12 += 65535
                            var v16 int8 = 1
                            var v17 int16 = -32768
                            var v18 int = fi11
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        _ = v14
                        _ = v15
                    }
                case "":
                    switch fi13 {
                    case 441:
                        v19 := GenS2{F0: -1}
                        _ = v19
                    case 4294967295:
                        fmt.Println(v4)
                    case 127:
                        switch (fi11 / ((v4) | 1)) {
                        case 32767:
                            switch (v4 ^ fi13) {
                            case -300:
                                var v20 float64 = (-2.25 - 1e10)
                                _ = v20
                            case 979:
                                fmt.Println(fi13, v2)
                            default:
                                if false {
                                    var v21 int16 = 42
                                    _ = v21
                                } else {
                                    var v22 int16 = 1000
                                    _ = v22
                                }
                            }
                        case 127:
                            v2 += float64(3.14159)
                        case 642:
                            v23 := GenS3{F0: -1, F1: 2147483647}
                            _ = v23
                        }
                    default:
                        switch "\n\t" {
                        case "hello":
                            if true {
                                v5 = (v5 / ((v12) | 1))
                                v24 := v3[4294967295]
                                var v25 uint64 = 721
                                if v24 {
                                    switch "hello" {
                                    case "a":
                                        v2 -= 1.5
                                    }
                                    delete(v3, 65535)
                                    v5 = 32767
                                } else {
                                    var v26 int64 = 42
                                    switch (-(v4)) {
                                    case -2147483648:
                                        if len(v1) > 0 {
                                            v1[0] = (("hello" + "\n\t") + (("foo bar" + "") + "foo bar"))
                                        }
                                    case 4294967295:
                                        if v24 {
                                            v25 = (^(v25))
                                            v1 = append(v1, "hello")
                                            _ = len(v3)
                                            var v27 uint = 42
                                            switch int(v27) {
                                            case 9223372036854775807:
                                                v28 := GenS4{F0: 550}
                                                _ = v28
                                            case 4294967295:
                                                var v29 int = 9223372036854775807
                                                _ = v29
                                            }
                                            _ = v27
                                        } else {
                                            if (1.5 <= ((1.5 * 1.5) / float32(v2))) {
                                                v5 = (v12 * (^(v12)))
                                            }
                                        }
                                    case 127:
                                        var v30 float64 = ((v2 * (v2 / v2)) * v2)
                                        _ = v30
                                    }
                                    v5 -= (v5 - v5)
                                    _ = v26
                                }
                                _ = v24
                                _ = v25
                            }
                        default:
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch (-(v4)) {
                                case 128:
                                    var v31 uint64 = 4294967295
                                    _ = v31
                                }
                            }
                        }
                    }
                default:
                    fmt.Println(v12, fi13, fi11)
                }
                v32 := GenS5{F0: v5}
                _ = fi13
                _ = v32
            }
            v5 += 1
            _ = fi11
            _ = v12
        }
    }
    v4 += 118
    var v33 uint16 = v5
    switch (-(v4)) {
    case 127:
        if ((-9223372036854775808 <= 0) || false) {
            for fi34 := 0; fi34 < 17; fi34++ {
                var v35 uint64 = 4294967295
                _ = fi34
                _ = v35
            }
            var v36 uint16 = v5
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        fmt.Println(v2)
        fmt.Println(v2, v2)
        var v38 float64 = v2
        _ = fi37
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
