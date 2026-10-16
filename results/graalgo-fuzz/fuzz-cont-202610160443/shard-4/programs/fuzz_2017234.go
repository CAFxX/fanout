package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2147483648}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    v1 -= (-(v1))
    var v6 uint = 1
    v7 := GenS1{F0: 0, F1: 2147483647, F2: -22}
    switch ("hello" + "hello") {
    case "hello":
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: 42, F1: -1, F2: -116}
            _ = fi8
            _ = v9
        }
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = 0
        }
    case "\n\t":
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int16 = (v1 ^ v1)
            var v12 uint32 = 636
            if false {
                switch (("foo bar" + "\n\t") + ("" + ("" + ""))) {
                case "\n\t":
                    var v13 uint = (v6 | 374)
                    _ = v13
                }
                v11 -= (v11 / ((v1) | 1))
                v12 += 0
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    if true {
        var v14 int32 = -2147483648
        v15 := GenS2{F0: -1}
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: -1}
    v6 = v6
    v6 += v6
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := GenS3{F0: 42}
        _ = fi17
        _ = v18
    }
    fmt.Println(v6)
    var v19 int64 = -9223372036854775808
    for fi20 := 0; fi20 < 14; fi20++ {
        v5[65535] = fi20
        if (0 < uint8(v6)) {
            switch fi20 {
            case 42:
                switch (fi20 & 2147483647) {
                case 4294967295:
                    for fi21 := 0; fi21 < 18; fi21++ {
                        fmt.Println(fi20, v1, v6)
                        var v22 float32 = (((float32(float64((float32((0.0 - 1e10)) / -2.25))) / 0.1) / 3.14159) / -2.25)
                        var v23 int32 = 1
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                default:
                    fmt.Println(v1, v1)
                }
            case -2147483648:
                var v24 uint = v6
                _ = v24
            case 697:
                var v25 int64 = v19
                _ = v25
            default:
                for fi26 := 0; fi26 < 16; fi26++ {
                    if (1 >= -754) {
                        v1 -= v1
                    }
                    _ = fi26
                }
            }
            v2[230] = 48
        } else {
            if false {
                var v27 uint = v6
                _ = v27
            }
            switch fi20 {
            case -478:
                for fi28 := 0; fi28 < 5; fi28++ {
                    fmt.Println(fi28, v6)
                    _ = fi28
                }
            case -1:
                v29 := GenS3{F0: v19}
                _ = v29
            case -274:
                if (!((0 < 2) && true)) {
                    if (65535 != 1) {
                        switch (fi20 * 32768) {
                        case -536:
                            var v30 int32 = 42
                            _ = v30
                        case -9223372036854775808:
                            var v31 int64 = v19
                            _ = v31
                        case -1:
                            v32 := GenS4{F0: 127}
                            _ = v32
                        }
                        if false {
                            var v33 int16 = 1
                            _ = v33
                        } else {
                            var v34 int64 = 9223372036854775807
                            _ = v34
                        }
                        switch fi20 {
                        case 2147483647:
                            v19 -= (v19 & 9223372036854775807)
                        case -9223372036854775808:
                            fmt.Println(v6)
                        case -502:
                            if len(v3) > 0 {
                                v3[0] = 414
                            }
                        }
                    }
                }
            default:
                var v35 uint8 = 128
                _ = v35
            }
        }
        var v36 float32 = 1e10
        _ = fi20
        _ = v36
    }
    delete(v5, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
