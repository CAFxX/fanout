package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 int32 }
func genFn0(p0 uint8, p1 int32) uint32 {
    var v5 uint64 = 1
    v6 := []int{-9223372036854775808}
    switch "foo bar" {
    case "hello":
        v7 := GenS1{F0: 42, F1: (p1 / (((p1 & 78)) | 1))}
        _ = v7
    case "foo bar":
        for fi8 := 0; fi8 < 12; fi8++ {
            _ = len(v6)
            var v9 int = (fi8 ^ 199)
            v5 = v5
            _ = fi8
            _ = v9
        }
    default:
        var v10 bool = false
        _ = v10
    }
    switch (("hello" + "a") + "hello") {
    case "\n\t":
        v11 := GenS1{F0: 917, F1: p1}
        _ = v11
    default:
        for fi12 := 0; fi12 < 4; fi12++ {
            if false {
                v5 -= v5
                if ((0.0 + (3.14159 - 3.14159)) < (3.14159 * ((0.0 * 3.14159) - (1.5 * (0.1 - ((1e10 + (-2.25 / (-2.25 * -2.25))) * (float32(1.5) * (3.14159 * ((-2.25 / -2.25) + 1e10))))))))) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v13 uint32 = 332
                    _ = v13
                }
            }
            switch (fi12 - (fi12 | (fi12 >> ((171) & 63)))) {
            case 0:
                fmt.Println(fi12)
            default:
                switch "foo bar" {
                case "\n\t":
                    v14 := GenS1{F0: 317, F1: 2147483647}
                    _ = v14
                }
            }
            switch 1 {
            case 127:
                var v15 uint = 2
                _ = v15
            case -251:
                var v16 uint8 = (^(p0))
                _ = v16
            case 705:
                if (850 == 1) {
                    v6 = append(v6, fi12)
                }
            }
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 386
}

func main() {
    v1 := []uint32{2147483647, 4294967295}
    v2 := make(map[string]bool)
    var v3 int32 = 1
    var v4 int8 = -1
    if (((0 < 1) || (true && (42 >= 894))) || true) {
        v3 += v3
        switch "hello" {
        case "":
            var v17 float32 = 3.14159
            _ = v17
        case "foo bar":
            v18 := GenS2{F0: 42}
            _ = v18
        }
        var v19 int = -9223372036854775808
        _ = v19
    }
    var v20 bool = false
    var v21 uint32 = genFn0(255, v3)
    var v22 int8 = v4
    v20 = (1 == 4294967295)
    for fi23 := 0; fi23 < 18; fi23++ {
        if ("hello" >= "\n\t") {
            var v24 int16 = 32767
            _ = v24
        }
        _ = fi23
    }
    v21 += v21
    v4 -= 1
    for fi25 := 0; fi25 < 20; fi25++ {
        if v20 {
            v26 := GenS3{F0: 0}
            if v20 {
                v3 -= (v3 * 1)
            } else {
                v4 += -128
                for fi27 := 0; fi27 < 1; fi27++ {
                    v28 := GenS3{F0: 64}
                    v4 = v4
                    v29 := GenS2{F0: 32767}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                v20 = (!(42 > (fi25 | fi25)))
            }
            _ = v26
        }
        v2[((("foo bar" + "a") + "hello") + "")] = true
        _ = fi25
    }
    if v20 {
        switch 128 {
        case 9223372036854775807:
            v21 = 156
        case 32768:
            var v30 uint = 2
            _ = v30
        case 128:
            switch 32768 {
            case 856:
                v31 := GenS1{F0: -1, F1: (v3 ^ v3)}
                _ = v31
            }
        default:
            fmt.Println(v21, v21)
        }
        v32 := GenS4{F0: 120}
        _ = v32
    }
    v33 := GenS5{F0: (v4 ^ v4)}
    for fi34 := 0; fi34 < 1; fi34++ {
        for fi35 := 0; fi35 < 1; fi35++ {
            if (0 < 1) {
                v36 := GenS2{F0: 42}
                _ = v36
            } else {
                if (-217 < -167) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v3, fi34)
                } else {
                    if false {
                        switch ("hello" + "\n\t") {
                        case "hello":
                            fmt.Println(v22, v4)
                        case "":
                            v20 = v20
                        case "foo bar":
                            v37 := GenS1{F0: 594, F1: 2147483647}
                            _ = v37
                        }
                    } else {
                        v20 = v20
                        if v20 {
                            fmt.Println(v4)
                            v38 := GenS6{F0: 32767, F1: v3}
                            _ = v38
                        } else {
                            v39 := GenS4{F0: (v22 * v22)}
                            if (!(-837 > 207)) {
                                v40 := GenS1{F0: 399, F1: 1}
                                var v41 int64 = 1
                                switch "hello" {
                                case "\n\t":
                                    v42 := GenS3{F0: 195}
                                    _ = v42
                                }
                                _ = v40
                                _ = v41
                            }
                            fmt.Println(v3, fi35)
                            _ = v39
                        }
                    }
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v2)
            _ = fi35
        }
        _ = fi34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v4
}
