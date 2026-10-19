package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    v1 := []int{127, 42, 307, -2147483648}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 float64 = 1e10
        for fi6 := 0; fi6 < 1; fi6++ {
            v1 = append(v1, fi6)
            _ = fi6
        }
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(fi4, fi4, fi4)
            _ = fi7
        }
        _ = fi4
        _ = v5
    }
    var v8 int = 9223372036854775807
    v8 = v8
    v8 = (v8 - (^(v8)))
    if (1 >= -128) {
        if false {
            fmt.Println(v8)
            if (!false) {
                fmt.Println(v8)
                v9 := GenS1{F0: 42, F1: 1}
                _ = v9
            } else {
                v8 -= (v8 % ((v8) | 1))
            }
            var v10 string = "\n\t"
            _ = v10
        } else {
            var v11 int32 = 768
            _ = v11
        }
    }
    v8 += v8
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint = 286
        _ = fi12
        _ = v13
    }
    fmt.Println(v8)
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int64 = 1
        if (true && true) {
            var v16 uint = 42
            var v17 int32 = 1
            v8 -= v8
            _ = v16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    if false {
        fmt.Println(v8)
        v8 = (v8 & (v8 % (((-(v8))) | 1)))
        v8 += -1
    }
    switch (-(v8)) {
    case 4294967295:
        if false {
            _ = len(v2)
        }
    case -337:
        v8 += (-(v8))
    case 2147483647:
        switch 9223372036854775807 {
        case 2147483647:
            v8 = -394
        case 32768:
            var v18 int64 = -1
            _ = v18
        }
    }
    var v19 string = "foo bar"
    v19 = v19
    v19 = ((((v19 + v19) + v19) + v19) + (v19 + v19))
    for fi20 := 0; fi20 < 19; fi20++ {
        switch (v19 + "hello") {
        case "\n\t":
            if (0 != 492) {
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 float64 = 1e10
                    v22 = v22
                    _ = fi21
                    _ = v22
                }
            }
        case "a":
            v8 = 9223372036854775807
        case "foo bar":
            for fi23 := 0; fi23 < 20; fi23++ {
                if ((!false) && ((81 <= 128) || true)) {
                    if false {
                        v24 := GenS2{F0: 1, F1: false}
                        v1 = append(v1, 0)
                        v19 = ""
                        if len(v1) > 0 {
                            v1[0] = -863
                        }
                        _ = v24
                    } else {
                        fmt.Println(fi23, fi20)
                        var v25 uint8 = 127
                        _ = v25
                    }
                    var v26 float32 = (-2.25 / (-2.25 + 0.0))
                    if (!(int16(v8) != -585)) {
                        if (4294967295 != 255) {
                            _ = len(v1)
                            var v27 int16 = 42
                            _ = v27
                        }
                        switch (((v19 + v19) + "") + v19) {
                        case "":
                            v26 += v26
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    } else {
                        v26 = v26
                        switch fi20 {
                        case 0:
                            v19 = v19
                        case -682:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 127:
                            var v28 uint64 = 2
                            _ = v28
                        }
                        v29 := GenS2{F0: 2147483648, F1: (363 > -2147483648)}
                        _ = v29
                    }
                    _ = v26
                }
                var v30 uint64 = 1
                v19 = ""
                _ = fi23
                _ = v30
            }
        }
        fmt.Println(v8, fi20)
        if (-9223372036854775808 < 127) {
            switch "a" {
            case "hello":
                switch -116 {
                case 9223372036854775807:
                    fmt.Println(v19, fi20, v8)
                case 176:
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v8 = 126
                        v32 := v3[4294967295]
                        v1 = append(v1, v8)
                        _ = fi31
                        _ = v32
                    }
                }
            default:
                var v33 float64 = 1e10
                _ = v33
            }
        } else {
            fmt.Println(v19)
        }
        _ = fi20
    }
    v34 := GenS1{F0: 943, F1: 5}
    for fi35 := 0; fi35 < 2; fi35++ {
        v8 -= (v8 * fi35)
        _ = fi35
    }
    v19 = v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v8
}
