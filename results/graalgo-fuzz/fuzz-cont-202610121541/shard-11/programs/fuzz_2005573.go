package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint; F2 int32 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 string; F1 int; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = -611
    v3 := make(map[int]bool)
    v4 := GenS1{F0: 0, F1: 0, F2: -1}
    if (0 > 0) {
        switch 1 {
        case 4294967295:
            v5 := []int{32768, 672}
            _ = v5
        case 347:
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v6 uint32 = 42
                _ = v6
            case 673:
                v7 := GenS2{F0: 0}
                _ = v7
            case 42:
                var v8 uint32 = 0
                _ = v8
            }
        case -898:
            var v9 uint8 = 2
            _ = v9
        default:
            var v10 bool = ("foo bar" != (("a" + "foo bar") + "a"))
            _ = v10
        }
    }
    var v11 int8 = 11
    switch "hello" {
    case "":
        for fi12 := 0; fi12 < 15; fi12++ {
            if ((!((0 >= 1) && true)) || (v11 != v11)) {
                if (true || false) {
                    var v13 bool = true
                    var v14 float32 = 0.0
                    _ = v13
                    _ = v14
                } else {
                    if true {
                        var v15 float64 = 0.0
                        _ = v15
                    } else {
                        var v16 int64 = -669
                        _ = v16
                    }
                    var v17 string = "hello"
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 10; fi18++ {
                fmt.Println(v11)
                var v19 uint8 = 255
                if (v2 >= (v2 ^ (v2 * (v2 + (-(v2)))))) {
                    var v20 float32 = 0.0
                    var v21 int32 = 0
                    _ = v20
                    _ = v21
                } else {
                    if (42 == 0) {
                        fmt.Println(v11, fi18)
                        if true {
                            var v22 uint8 = 42
                            v19 += (v22 * v22)
                            _ = v22
                        }
                    }
                }
                _ = fi18
                _ = v19
            }
            v11 = v11
            _ = fi12
        }
    case "foo bar":
        v11 = (v11 + 127)
    default:
        for fi23 := 0; fi23 < 16; fi23++ {
            switch "" {
            case "":
                fmt.Println(fi23, fi23, v11)
            case "a":
                if false {
                    var v24 int64 = -569
                    if (!(v2 < int16(v24))) {
                        v25 := GenS2{F0: 128}
                        v26 := []int{0}
                        _ = v25
                        _ = v26
                    }
                    var v27 string = "\n\t"
                    _ = v24
                    _ = v27
                }
            default:
                if true {
                    v28 := []int{(fi23 / (((^(fi23))) | 1)), fi23, (fi23 * (fi23 / ((fi23) | 1))), fi23}
                    _ = v28
                }
            }
            _ = fi23
        }
    }
    v29 := GenS3{F0: -967, F1: 42, F2: 2147483647}
    var v30 uint64 = 988
    v11 += v11
    var v31 float64 = 3.14159
    for fi32 := 0; fi32 < 2; fi32++ {
        v11 -= (v11 & v11)
        fmt.Println(v30, v2, v2)
        _ = fi32
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        v11 = 42
        _ = fi33
    }
    var v34 float64 = ((float64(0.1) - v31) - 1e10)
    for fi35 := 0; fi35 < 2; fi35++ {
        v36 := v1[2]
        v37 := GenS4{F0: (("\n\t" + "") + ""), F1: v36}
        _ = fi35
        _ = v36
        _ = v37
    }
    v38 := GenS3{F0: -538, F1: 0, F2: -1}
    var v39 int32 = 0
    v39 = v39
    v40 := GenS3{F0: 1, F1: uint(v30), F2: v39}
    for fi41 := 0; fi41 < 17; fi41++ {
        switch fi41 {
        case -212:
            for fi42 := 0; fi42 < 9; fi42++ {
                v43 := []int{int(v11)}
                _ = fi42
                _ = v43
            }
        default:
            if (2147483647 < 65535) {
                v44 := GenS5{F0: "\n\t", F1: (fi41 - (-(fi41))), F2: (-(v2))}
                _ = v44
            }
        }
        if ((v11 / (((v11 & (v11 ^ v11))) | 1)) > (v11 | int8(v30))) {
            fmt.Println(v31, v34, v2)
        } else {
            fmt.Println(fi41, v31, v39)
            v30 = 4294967295
        }
        _ = fi41
    }
    v45 := v3[32768]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v45
}
