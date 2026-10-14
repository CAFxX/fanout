package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint64; F2 int16 }
type GenS5 struct { F0 int }
func genFn0(p0 int) string {
    fmt.Println(p0)
    var v4 int16 = 1
    fmt.Println(p0, p0)
    var v5 int8 = 118
    fmt.Println(v4, p0, p0)
    _ = p0
    _ = v4
    _ = v5
    return (("foo bar" + "") + "a")
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 32767
    var v3 string = "a"
    v6 := v1[(v2 >> ((1) & 63))]
    v2 = (v2 % ((54) | 1))
    v7 := GenS1{F0: 4}
    switch -527 {
    case -9223372036854775808:
        for fi8 := 0; fi8 < 18; fi8++ {
            switch -9223372036854775808 {
            case 727:
                fmt.Println(v6)
            case -382:
                v9 := []int{(^(v2))}
                _ = v9
            case 1:
                var v10 int16 = -1
                _ = v10
            default:
                v6 = ((1e10 / 1.5) < (3.14159 * 1e10))
            }
            v6 = false
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 int64 = 326
                switch v3 {
                case "hello":
                    if (!false) {
                        switch "foo bar" {
                        case "a":
                            var v13 int32 = 42
                            _ = v13
                        case "hello":
                            if v6 {
                                if v6 {
                                    fmt.Println(v3, v12)
                                }
                                v2 -= (fi8 * fi11)
                            }
                        case "":
                            fmt.Println(v2, v12)
                        }
                        v14 := v1[1]
                        _ = v14
                    } else {
                        if v6 {
                            var v15 float32 = (-2.25 - float32(0.0))
                            v16 := []int{(v2 >> ((2) & 63)), fi8}
                            switch v2 {
                            case 242:
                                v17 := GenS1{F0: 42}
                                _ = v17
                            case -77:
                                v3 = v3
                            default:
                                if v6 {
                                    v3 = v3
                                } else {
                                    var v18 float32 = v15
                                    v12 -= (v12 - (-(v12)))
                                    _ = v18
                                }
                            }
                            _ = v15
                            _ = v16
                        } else {
                            v19 := []int{v2, (fi8 * (fi11 % (((fi8 | 1)) | 1))), 498}
                            v2 += (fi11 * -1)
                            _ = v19
                        }
                        v20 := GenS2{F0: 65535}
                        _ = v20
                    }
                case "a":
                    _ = len(v1)
                case "\n\t":
                    v2 += (fi8 / (((-(v2))) | 1))
                }
                var v21 uint64 = 0
                _ = fi11
                _ = v12
                _ = v21
            }
            _ = fi8
        }
    case 1:
        var v22 int16 = -1
        _ = v22
    default:
        var v23 int16 = 1000
        _ = v23
    }
    var v24 string = "hello"
    v25 := GenS1{F0: -461}
    v26 := GenS3{F0: v3}
    v27 := GenS4{F0: "a", F1: 1, F2: 1000}
    v28 := GenS3{F0: v3}
    fmt.Println(v24, v2)
    for fi29 := 0; fi29 < 8; fi29++ {
        v24 = "hello"
        v30 := GenS5{F0: v2}
        _ = fi29
        _ = v30
    }
    if true {
        delete(v1, (v2 >> ((255) & 63)))
        fmt.Println(v3, v2)
    } else {
        switch v2 {
        case -2147483648:
            var v31 int64 = -991
            _ = v31
        case 32767:
            v32 := GenS1{F0: 0}
            _ = v32
        case 1:
            var v33 uint8 = 255
            _ = v33
        default:
            v34 := GenS2{F0: 74}
            _ = v34
        }
    }
    switch "hello" {
    case "hello":
        v35 := []int{651, (v2 >> ((65535) & 63)), (v2 * (v2 % ((-386) | 1)))}
        _ = v35
    }
    fmt.Println(v3, v24)
    switch "\n\t" {
    case "":
        fmt.Println(v6)
    case "a":
        if (26 > 127) {
            var v36 int32 = 1
            var v37 int32 = (v36 ^ (v36 >> ((1) & 63)))
            v24 = ("\n\t" + "\n\t")
            _ = v36
            _ = v37
        } else {
            fmt.Println(v24, v3)
            if v6 {
                var v38 int8 = 127
                v39 := GenS1{F0: -196}
                _ = v38
                _ = v39
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v6
    _ = v7
}
