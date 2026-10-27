package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int64; F2 int64 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int8; F2 string }
type GenS6 struct { F0 int16; F1 uint64 }

func main() {
    var v1 float64 = (0.1 * (float64(1e10) + -2.25))
    v2 := make(map[string]uint32)
    var v3 float64 = (((v1 + 1e10) - (v1 + float64((-2.25 + 1e10)))) - ((0.1 + float64(-2.25)) + v1))
    v4 := []uint32{0, 4294967295}
    var v5 uint = 42
    v6 := []int{-2147483648, -268}
    var v7 uint64 = 42
    if (-643 > 32767) {
        v8 := GenS1{F0: 100, F1: v5}
        var v9 uint64 = (v7 % ((4294967295) | 1))
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int64 = -9223372036854775808
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        switch "" {
        case "hello":
            if true {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 string = "\n\t"
                    if true {
                        var v15 uint = 2
                        v15 = 127
                        switch fi12 {
                        case -9223372036854775808:
                            switch fi12 {
                            case 128:
                                var v16 uint16 = 2
                                _ = v16
                            case -505:
                                var v17 uint32 = 4294967295
                                _ = v17
                            case -113:
                                var v18 float32 = 0.1
                                _ = v18
                            default:
                                fmt.Println(v14, fi13)
                            }
                        case 533:
                            switch (fi12 - (fi12 / (((fi13 ^ fi12)) | 1))) {
                            case 128:
                                if false {
                                    if (v14 <= v14) {
                                        v19 := v2[v14]
                                        _ = v19
                                    }
                                } else {
                                    var v20 int = fi12
                                    fmt.Println(v1, fi12, v5)
                                    _ = v20
                                }
                            case 32768:
                                switch (fi13 >> (((v5 % (((v5 & 0)) | 1))) & 63)) {
                                case -948:
                                    var v21 float32 = (float32(v1) * 0.1)
                                    _ = v21
                                case 4294967295:
                                    v22 := GenS1{F0: -23, F1: 0}
                                    _ = v22
                                case 1:
                                    var v23 uint32 = 2147483647
                                    _ = v23
                                }
                            default:
                                if false {
                                    var v24 uint32 = 334
                                    _ = v24
                                }
                            }
                        case 2147483647:
                            v25 := GenS2{F0: -1}
                            _ = v25
                        default:
                            if (1 >= 276) {
                                v7 = (v7 | 824)
                            } else {
                                var v26 int = 173
                                v14 = "a"
                                if false {
                                    switch "\n\t" {
                                    case "a":
                                        var v27 uint8 = 125
                                        _ = v27
                                    case "foo bar":
                                        delete(v2, "foo bar")
                                    }
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v26, v3)
                                    }
                                    v15 = 973
                                }
                                v7 = (v7 << ((18446744073709551615) & 63))
                                _ = v26
                            }
                        }
                        _ = v15
                    } else {
                        if (!(0.0 == 0.0)) {
                            switch (-(fi13)) {
                            case -242:
                                if false {
                                    switch v14 {
                                    case "hello":
                                        v1 = (v3 * v1)
                                    case "a":
                                        var v28 float32 = 0.0
                                        _ = v28
                                    case "":
                                        v29 := GenS2{F0: -833}
                                        switch -895 {
                                        case 42:
                                            v30 := GenS3{F0: 945, F1: 0, F2: -1}
                                            _ = v30
                                        }
                                        var v31 float64 = (((0.0 + 0.1) + v3) - (v1 * -2.25))
                                        var v32 uint16 = 0
                                        _ = v29
                                        _ = v31
                                        _ = v32
                                    }
                                    v33 := GenS4{F0: 42, F1: true, F2: uint64(v5)}
                                    _ = v33
                                }
                            }
                        }
                    }
                    fmt.Println(v7, v5, fi13)
                    _ = fi13
                    _ = v14
                }
                for fi34 := 0; fi34 < 17; fi34++ {
                    if true {
                        fmt.Println(v3, fi12, v1)
                    } else {
                        if false {
                            v1 += v3
                            fmt.Println(fi12, fi34, v5)
                            if false {
                                delete(v2, "a")
                                v35 := GenS5{F0: (v7 << (((v5 & v5)) & 63)), F1: 0, F2: "foo bar"}
                                var v36 float32 = -2.25
                                _ = v35
                                _ = v36
                            }
                        } else {
                            var v37 float32 = float32(v1)
                            _ = v37
                        }
                    }
                    fmt.Println(fi12)
                    _ = fi34
                }
                switch "\n\t" {
                case "hello":
                    v1 += v3
                default:
                    v3 -= v1
                }
            }
        }
        var v38 uint8 = 25
        v6 = append(v6, (fi12 << (((^(v5))) & 63)))
        _ = fi12
        _ = v38
    }
    v7 = (v7 - v7)
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 float64 = 1.5
        var v41 int16 = 1
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 uint16 = 32767
    fmt.Println(v42, v1)
    _ = len(v6)
    v5 -= v5
    for fi43 := 0; fi43 < 11; fi43++ {
        v44 := GenS6{F0: 42, F1: v7}
        fmt.Println(fi43, v42)
        _ = fi43
        _ = v44
    }
    v5 += (v5 >> ((789) & 63))
    for fi45 := 0; fi45 < 5; fi45++ {
        var v46 int8 = 100
        v47 := v2["foo bar"]
        v48 := GenS2{F0: 1000}
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    var v49 int16 = -1
    _ = len(v2)
    for fi50 := 0; fi50 < 17; fi50++ {
        v51 := v2[("" + "hello")]
        var v52 uint16 = (v42 << (((v5 ^ v5)) & 63))
        var v53 float32 = (1.5 * 1.5)
        _ = fi50
        _ = v51
        _ = v52
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
