package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint64; F2 string }

func main() {
    var v1 int = 9223372036854775807
    v2 := []int{}
    v3 := []uint32{0}
    var v4 uint = 42
    for fi5 := 0; fi5 < 1; fi5++ {
        v4 = (^(v4))
        if (!false) {
            fmt.Println(v1)
            if len(v3) > 0 {
                v3[0] = 42
            }
        } else {
            v1 += (v1 << ((42) & 63))
        }
        _ = fi5
    }
    v6 := GenS1{F0: -2147483648}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v7 := GenS1{F0: -2147483648}
    if false {
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            v2 = append(v2, int(v4))
        case "":
            fmt.Println(v1, v1)
        default:
            v1 += 424
        }
    } else {
        var v8 int16 = 1
        v9 := GenS2{F0: 4294967295, F1: 42, F2: 0}
        _ = v8
        _ = v9
    }
    var v10 float32 = (float32(3.14159) - 0.1)
    v4 = (^(v4))
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 uint32 = 1
            var v14 float64 = (((float64(float32(1.5)) / ((0.1 - 3.14159) + 1.5)) - 1e10) / 1.5)
            if (128 < 156) {
                var v15 int64 = -263
                if false {
                    _ = len(v3)
                    var v16 int32 = int32(v4)
                    var v17 int8 = -1
                    _ = v16
                    _ = v17
                } else {
                    if (("\n\t" + "a") != "\n\t") {
                        fmt.Println(fi11, v10)
                    } else {
                        v18 := GenS3{F0: -1, F1: 128, F2: 1}
                        _ = v18
                    }
                    var v19 int8 = -25
                    _ = v19
                }
                switch (("a" + "a") + "") {
                case "":
                    if len(v3) > 0 {
                        v3[0] = (v13 << (((^(v4))) & 63))
                    }
                case "\n\t":
                    var v20 bool = true
                    _ = v20
                case "a":
                    v21 := GenS4{F0: 127}
                    _ = v21
                }
                _ = v15
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    if ((1e10 * float32(((0.1 / (((0.0 - ((-2.25 - (1.5 / 3.14159)) / 3.14159)) + 0.1) * 1.5)) / 1e10))) >= v10) {
        v1 = (v1 ^ v1)
        for fi22 := 0; fi22 < 12; fi22++ {
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS5{F0: -1, F1: 700, F2: ("a" + "foo bar")}
                _ = fi23
                _ = v24
            }
            var v25 uint16 = 235
            v10 += v10
            _ = fi22
            _ = v25
        }
        _ = len(v2)
    } else {
        v26 := GenS3{F0: 795, F1: 145, F2: (^(v4))}
        if (0 <= 42) {
            for fi27 := 0; fi27 < 4; fi27++ {
                switch (fi27 + (v1 >> ((2) & 63))) {
                case -1:
                    var v28 uint = 65535
                    _ = v28
                case -9223372036854775808:
                    switch (fi27 - 2147483647) {
                    case -274:
                        v4 -= 4294967295
                    case 1:
                        var v29 int64 = 9223372036854775807
                        _ = v29
                    default:
                        var v30 uint16 = 255
                        _ = v30
                    }
                }
                if ((!((fi27 << (((v4 * v4)) & 63)) > fi27)) && false) {
                    v1 = (fi27 + fi27)
                    v31 := GenS3{F0: -22, F1: 255, F2: v4}
                    _ = v31
                }
                v10 -= float32(1e10)
                _ = fi27
            }
            if true {
                v4 -= 65535
                v1 += (-(v1))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v32 uint = (v4 >> (((v4 & 65535)) & 63))
                _ = v32
            }
            v33 := GenS4{F0: 223}
            _ = v33
        } else {
            if false {
                v34 := GenS3{F0: -224, F1: 0, F2: (v4 - 790)}
                _ = v34
            }
            v1 -= (v1 / (((v1 / (((v1 & v1)) | 1))) | 1))
            if true {
                var v35 int = v1
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                v10 = v10
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            v4 -= (v4 * v4)
            v4 = (^(v4))
            _ = fi36
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
