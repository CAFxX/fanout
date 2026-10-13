package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int16; F2 int }
type GenS3 struct { F0 uint; F1 bool; F2 uint }

func main() {
    var v1 int8 = 127
    var v2 uint = 2
    v3 := make(map[int]uint32)
    var v4 uint16 = 2
    v5 := make(map[string]bool)
    switch "foo bar" {
    case "a":
        var v6 float32 = 0.1
        _ = v6
    case "":
        fmt.Println(v1, v4)
    }
    var v7 float32 = (-2.25 * (((0.1 * 1.5) / float32(float64(float32((1.5 - (0.0 * 0.1)))))) / (float32((3.14159 - float64(0.0))) * (1e10 / (((0.0 + 1.5) + 3.14159) * (1.5 - 1e10))))))
    if false {
        switch ("hello" + "hello") {
        case "\n\t":
            var v8 int64 = 42
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint = 42
            var v12 int = fi10
            if (9223372036854775807 != 1) {
                if (false || (1000 != int16(v2))) {
                    v4 += (v4 % ((0) | 1))
                }
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v7 += -2.25
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int = -981
            fmt.Println(v2)
            v7 = v7
            _ = fi13
            _ = v14
        }
        _ = fi9
    }
    v2 = v2
    v15 := GenS1{F0: -797}
    if (574 >= 324) {
        v16 := []int{0, 1, -9223372036854775808, 0}
        _ = v16
    } else {
        v17 := GenS1{F0: 472}
        v5["\n\t"] = false
        _ = v17
    }
    fmt.Println(v7)
    var v18 int64 = 1
    var v19 uint64 = uint64(v4)
    if false {
        if (0 != 0) {
            for fi20 := 0; fi20 < 17; fi20++ {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v7 = 1e10
                    _ = fi21
                }
                var v22 uint8 = 42
                if (!false) {
                    v23 := GenS1{F0: v18}
                    v1 = (-(v1))
                    _ = v23
                }
                _ = fi20
                _ = v22
            }
        } else {
            if true {
                v4 += (v4 >> ((v2) & 63))
                v24 := GenS2{F0: uint8(v19), F1: -1, F2: 127}
                _ = v24
            }
            if true {
                v2 = (^(v2))
                switch 1 {
                case 9223372036854775807:
                    for fi25 := 0; fi25 < 8; fi25++ {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v27 := []int{}
                            v19 -= 511
                            _ = fi26
                            _ = v27
                        }
                        _ = fi25
                    }
                case 128:
                    var v28 uint64 = (^(v19))
                    _ = v28
                case -340:
                    fmt.Println(v4, v1, v7)
                default:
                    v7 = 1.5
                }
            }
        }
    } else {
        if (false && false) {
            fmt.Println(v2)
        }
        v29 := GenS3{F0: 693, F1: ("foo bar" >= "\n\t"), F2: (v2 - v2)}
        var v30 int = 1
        _ = v29
        _ = v30
    }
    v31 := []int{42, 42}
    switch 849 {
    case 296:
        var v32 string = ""
        _ = v32
    default:
        var v33 uint32 = 2147483647
        _ = v33
    }
    v31 = append(v31, 4294967295)
    var v34 int16 = 1000
    if false {
        var v35 uint64 = 846
        if (0 < 2147483647) {
            var v36 int64 = v18
            if (-877 == 606) {
                _ = len(v3)
                v35 = (^(v35))
            }
            v2 = (v2 ^ (v2 / (((v2 / (((v2 ^ v2)) | 1))) | 1)))
            _ = v36
        }
        _ = v35
    } else {
        var v37 int8 = int8(v34)
        var v38 int64 = 1
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 7; fi39++ {
        v4 = 1
        for fi40 := 0; fi40 < 9; fi40++ {
            v41 := GenS2{F0: 2, F1: (v34 / ((-830) | 1)), F2: fi40}
            v19 += 18446744073709551615
            _ = fi40
            _ = v41
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            v43 := GenS3{F0: v2, F1: true, F2: uint(v4)}
            var v44 uint8 = uint8(v19)
            fmt.Println(v44, v4, v19)
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi39
    }
    v4 += (v4 ^ (v4 + v4))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v7
}
