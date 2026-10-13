package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int; F1 bool; F2 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint16 }
func genFn0(p0 bool) bool {
    v7 := []int{}
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: ""}
        var v10 uint = 65535
        v11 := GenS2{F0: 560, F1: 0, F2: v10}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    if (!p0) {
        if (((float32(((((1.5 + (1.5 - (-2.25 - (0.1 / 1e10)))) - 1e10) - (0.0 * 0.1)) + (1.5 * (1e10 * 0.1)))) - float32(1.5)) * 0.1) <= ((float32(3.14159) - 1e10) / (float32(((((((0.1 / -2.25) + 3.14159) + 1.5) + 1e10) - float64(1e10)) / 1.5)) - (float32(1e10) / (3.14159 - (0.0 - float32(3.14159))))))) {
            v12 := GenS2{F0: 2147483648, F1: 255, F2: 65535}
            p0 = false
            _ = v12
        } else {
            p0 = true
        }
    }
    if (9223372036854775807 <= 42) {
        switch -9223372036854775808 {
        case 312:
            var v13 uint32 = 900
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 15; fi15++ {
                p0 = (1 <= 1)
                p0 = p0
                p0 = p0
                _ = fi15
            }
            switch "\n\t" {
            case "\n\t":
                v16 := GenS2{F0: 2147483648, F1: 65535, F2: 0}
                _ = v16
            default:
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := GenS3{F0: 255, F1: 42}
                    fmt.Println(fi17)
                    _ = fi17
                    _ = v18
                }
            }
            _ = fi14
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v7
    return p0
}
func genFn1(p0 uint32, p1 uint64, p2 bool) int32 {
    v19 := GenS4{F0: 4294967295, F1: p2, F2: -416}
    p2 = (-2.25 > float32(float64(1.5)))
    switch ("a" + "\n\t") {
    case "foo bar":
        p2 = p2
    default:
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS3{F0: 18446744073709551615, F1: 42}
            _ = fi20
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 0
}

func main() {
    var v1 bool = (0 <= 42)
    v2 := make(map[int]int)
    var v3 int8 = -1
    v4 := []uint32{1, 65535, 2147483648}
    v5 := make(map[string]uint32)
    var v6 int = 1
    v22 := GenS2{F0: 2147483647, F1: 1, F2: 1}
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := GenS5{F0: (v3 << ((0) & 63))}
        _ = fi23
        _ = v24
    }
    fmt.Println(v3, v3)
    v6 = (v6 | v6)
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS2{F0: 495, F1: 2, F2: 354}
        for fi27 := 0; fi27 < 7; fi27++ {
            if (v1 && v1) {
                v28 := GenS5{F0: -1}
                _ = len(v5)
                _ = v28
            } else {
                var v29 float64 = 0.1
                v30 := GenS6{F0: 4294967295, F1: 442}
                _ = v29
                _ = v30
            }
            switch ("hello" + "a") {
            case "":
                var v31 int32 = 42
                _ = v31
            case "hello":
                _ = len(v2)
            }
            if (true || v1) {
                fmt.Println(fi27, fi27)
                v32 := GenS7{F0: 0, F1: 18446744073709551615}
                var v33 uint64 = 1
                _ = v32
                _ = v33
            }
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    v34 := GenS8{F0: 205}
    switch v6 {
    case 664:
        if true {
            for fi35 := 0; fi35 < 7; fi35++ {
                v1 = (float32(0.0) > (3.14159 * (3.14159 * 0.1)))
                delete(v5, "hello")
                var v36 uint = 18446744073709551615
                _ = fi35
                _ = v36
            }
            var v37 float32 = float32(1.5)
            _ = len(v4)
            _ = v37
        }
    default:
        if (255 != 193) {
            switch "\n\t" {
            case "foo bar":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "":
                for fi38 := 0; fi38 < 2; fi38++ {
                    if v1 {
                        v3 = (v3 % ((v3) | 1))
                    } else {
                        var v39 uint8 = 127
                        var v40 uint16 = 1
                        _ = v39
                        _ = v40
                    }
                    v41 := GenS4{F0: (v6 / ((-2147483648) | 1)), F1: (genFn1(4294967295, 874, (((false && v1) || ((42 > 18446744073709551615) || true)) && v1)) <= 0), F2: fi38}
                    _ = fi38
                    _ = v41
                }
            case "a":
                for fi42 := 0; fi42 < 4; fi42++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi42
                }
            }
        }
    }
    var v43 int32 = 20
    fmt.Println(v1, v6)
    var v44 uint32 = 4294967295
    if len(v4) > 0 {
        v4[0] = (v44 + 0)
    }
    v45 := v2[v6]
    v44 -= (v44 % ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
