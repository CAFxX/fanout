package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 string; F1 uint; F2 int32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int64, p1 int64, p2 int64) uint8 {
    switch "a" {
    case "\n\t":
        if (-128 >= -1) {
            var v5 uint16 = 42
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                var v6 uint16 = (v5 << ((806) & 63))
                _ = v6
            case "\n\t":
                var v7 float32 = 0.1
                _ = v7
            case "":
                v8 := GenS1{F0: false, F1: 42}
                _ = v8
            }
            _ = v5
        }
    default:
        var v9 uint16 = 804
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int8 = -1
        switch fi10 {
        case 2147483647:
            var v12 int64 = (p2 ^ 203)
            _ = v12
        case 32768:
            p0 = (p0 | -9223372036854775808)
        case 217:
            var v13 int8 = (v11 << ((2) & 63))
            _ = v13
        default:
            var v14 int8 = v11
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 bool, p1 bool) int16 {
    fmt.Println(p0, p0)
    var v15 int8 = 60
    _ = p0
    _ = p1
    _ = v15
    return 1000
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float32 = (float32(((0.0 + (1.5 / 1.5)) - 3.14159)) * (((1.5 + 0.0) - (float32((1e10 * 0.1)) + 1.5)) + 1.5))
    var v3 string = "hello"
    var v4 float64 = ((-2.25 - 3.14159) / 3.14159)
    v2 = 1e10
    v2 = 1e10
    var v16 uint64 = 2
    v17 := []int{405, -493, -2147483648}
    var v18 bool = true
    fmt.Println(v16, v1)
    for fi19 := 0; fi19 < 2; fi19++ {
        v1 = (^(v1))
        switch (fi19 + (fi19 >> ((2) & 63))) {
        case 567:
            var v20 uint = 4294967295
            _ = v20
        case 1:
            if v18 {
                var v21 uint16 = 2
                _ = v21
            } else {
                var v22 uint8 = 137
                v1 += v1
                _ = v22
            }
        case 128:
            var v23 string = "\n\t"
            _ = v23
        }
        _ = fi19
    }
    if false {
        switch 0 {
        case -9223372036854775808:
            var v24 int = -527
            _ = v24
        case 127:
            v4 += (v4 * 1e10)
        case 2147483647:
            v25 := GenS1{F0: v18, F1: 2147483647}
            _ = v25
        default:
            switch ((("a" + v3) + v3) + "hello") {
            case "a":
                v26 := GenS2{F0: 639, F1: 365}
                _ = v26
            case "hello":
                var v27 float64 = ((1e10 * ((float64(float32((v4 - v4))) - (1e10 * (((-2.25 * -2.25) * v4) * v4))) * (0.1 * ((v4 - 0.1) - (1.5 * (v4 + v4)))))) * 0.1)
                _ = v27
            }
        }
        v28 := GenS3{F0: 346, F1: v18}
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        switch (fi29 * (fi29 ^ (fi29 % ((32767) | 1)))) {
        case 32768:
            switch (fi29 ^ 128) {
            case 2147483647:
                v4 -= -2.25
            case 127:
                if true {
                    v30 := GenS2{F0: 2, F1: 4294967295}
                    if v18 {
                        v18 = (127 > -37)
                        v4 += v4
                        if v18 {
                            var v31 uint = 19
                            if false {
                                for fi32 := 0; fi32 < 12; fi32++ {
                                    switch ("" + (v3 + v3)) {
                                    case "foo bar":
                                        v18 = v18
                                    case "a":
                                        v3 = "foo bar"
                                    }
                                    _ = fi32
                                }
                                fmt.Println(v4, fi29, v3)
                                switch fi29 {
                                case 42:
                                    _ = len(v17)
                                case 387:
                                    v33 := GenS2{F0: 65535, F1: v16}
                                    _ = v33
                                }
                            }
                            _ = v31
                        }
                    }
                    _ = v30
                } else {
                    for fi34 := 0; fi34 < 7; fi34++ {
                        var v35 int8 = -1
                        _ = fi34
                        _ = v35
                    }
                    var v36 uint16 = 255
                    var v37 int8 = 2
                    _ = v36
                    _ = v37
                }
            case -9223372036854775808:
                v17 = append(v17, fi29)
            }
        default:
            v38 := GenS4{F0: (("a" + "a") + v3), F1: 255, F2: 42}
            _ = v38
        }
        if (!v18) {
            v1 += 9223372036854775807
            if (1 != 2) {
                if len(v17) > 0 {
                    v17[0] = (fi29 + fi29)
                }
            } else {
                fmt.Println(v16, v3)
            }
        }
        v39 := GenS5{F0: 865}
        _ = fi29
        _ = v39
    }
    fmt.Println(v18)
    v4 -= float64((float32((((v4 * 1e10) - -2.25) - v4)) + v2))
    v17 = append(v17, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
