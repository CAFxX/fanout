package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int32; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint16 }

func main() {
    var v1 int32 = 42
    var v2 int16 = 1
    v3 := []uint32{109, 1, 2147483647}
    v4 := make(map[int]uint32)
    v5 := []int{-90, -493, 4294967295}
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: fi6}
        _ = len(v4)
        var v8 uint32 = 65535
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 = append(v3, 840)
    if (3.14159 == (0.1 - -2.25)) {
        if (true && (2 >= 42)) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (false || true) {
                v2 = -625
            } else {
                for fi9 := 0; fi9 < 13; fi9++ {
                    v1 = 2147483647
                    _ = fi9
                }
                var v10 uint32 = 2
                switch 669 {
                case -2147483648:
                    if (-2147483648 > 0) {
                        if (696 >= 696) {
                            _ = len(v4)
                            v11 := GenS1{F0: -602}
                            fmt.Println(v10, v10, v1)
                            _ = v11
                        } else {
                            var v12 int8 = 100
                            v10 = (v10 / (((^(v10))) | 1))
                            _ = v12
                        }
                    }
                default:
                    var v13 uint16 = 130
                    _ = v13
                }
                _ = v10
            }
        } else {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint32 = 0
                switch fi14 {
                case -1:
                    v16 := v4[fi14]
                    _ = v16
                }
                var v17 float32 = 0.1
                _ = fi14
                _ = v15
                _ = v17
            }
        }
    }
    switch 4294967295 {
    case 32768:
        v2 = v2
    case 9223372036854775807:
        if false {
            var v18 float64 = 1e10
            _ = v18
        }
    }
    var v19 uint8 = 127
    v19 -= (v19 ^ 255)
    var v20 uint = 1
    var v21 int32 = 42
    v21 += v21
    if (false || true) {
        var v22 bool = (-1 >= -9223372036854775808)
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint64 = 4294967295
        switch ("a" + "hello") {
        case "foo bar":
            var v25 float32 = ((0.1 * (3.14159 * 0.0)) - 0.0)
            _ = v25
        case "\n\t":
            switch (fi23 % (((^(fi23))) | 1)) {
            case 128:
                v2 -= (v2 | (v2 | v2))
            case -1:
                if (!true) {
                    switch "hello" {
                    case "\n\t":
                        var v26 bool = (458 > 414)
                        _ = v26
                    case "a":
                        v19 += 128
                    case "":
                        if false {
                            v2 = 419
                        }
                    }
                }
            case 0:
                for fi27 := 0; fi27 < 20; fi27++ {
                    var v28 uint = v20
                    _ = fi27
                    _ = v28
                }
            default:
                for fi29 := 0; fi29 < 19; fi29++ {
                    switch "hello" {
                    case "":
                        var v30 string = (("a" + ("\n\t" + "\n\t")) + "a")
                        _ = v30
                    }
                    v19 = (v19 * v19)
                    _ = fi29
                }
            }
        default:
            switch (((("hello" + "foo bar") + "a") + "a") + "hello") {
            case "a":
                var v31 int16 = -800
                _ = v31
            case "":
                var v32 uint8 = 150
                _ = v32
            case "foo bar":
                if (0 > 2) {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v19 -= (^(v19))
                        var v34 uint64 = v24
                        var v35 int8 = -1
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                } else {
                    v20 -= v20
                }
            }
        }
        var v36 int8 = int8(fi23)
        _ = fi23
        _ = v24
        _ = v36
    }
    switch "foo bar" {
    case "":
        fmt.Println(v19)
    default:
        for fi37 := 0; fi37 < 4; fi37++ {
            fmt.Println(v1)
            _ = fi37
        }
    }
    v38 := GenS2{F0: v1, F1: v21, F2: -8}
    v39 := v4[int(v21)]
    for fi40 := 0; fi40 < 1; fi40++ {
        v41 := GenS3{F0: true, F1: 440, F2: 255}
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
