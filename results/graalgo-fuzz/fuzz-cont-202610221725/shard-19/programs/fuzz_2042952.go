package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int32; F1 int64 }
func genFn0(p0 uint8, p1 string, p2 uint) int32 {
    var v5 bool = (!((p0 - (p0 * (p0 / ((p0) | 1)))) != (p0 % ((p0) | 1))))
    var v6 string = (p1 + "foo bar")
    if v5 {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: 4294967295}
            var v9 int8 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int64 = -780
        if v5 {
            v11 := []int{-9223372036854775808, 9223372036854775807, 32768}
            switch p1 {
            case "\n\t":
                if (!v5) {
                    p0 += (p0 >> ((p2) & 63))
                }
            case "hello":
                fmt.Println(p0)
            default:
                v12 := GenS2{F0: false, F1: (v5 || v5)}
                _ = v12
            }
            var v13 int64 = v10
            _ = v11
            _ = v13
        }
        _ = v10
    } else {
        var v14 uint16 = 32767
        _ = v14
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 638
}

func main() {
    v1 := []int{}
    v2 := []int{-1}
    v3 := make(map[uint32]int)
    var v4 float32 = (((-2.25 - 0.1) - (float32(float64(1e10)) * 1e10)) + 0.0)
    var v15 uint32 = 2147483647
    fmt.Println(v4, v15)
    v16 := v3[v15]
    if true {
        for fi17 := 0; fi17 < 20; fi17++ {
            for fi18 := 0; fi18 < 19; fi18++ {
                if (127 <= 255) {
                    v19 := GenS3{F0: false, F1: 42, F2: 1}
                    var v20 float64 = (1e10 + 3.14159)
                    if false {
                        v15 -= v15
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v20)
                    _ = v19
                    _ = v20
                } else {
                    v16 = (fi18 * fi18)
                    var v21 int8 = 127
                    switch 947 {
                    case -9223372036854775808:
                        switch "foo bar" {
                        case "a":
                            var v22 int32 = -2147483648
                            var v23 bool = true
                            switch 128 {
                            case 984:
                                _ = len(v1)
                            case 4294967295:
                                if (239 <= 2) {
                                    switch (fi17 >> ((734) & 63)) {
                                    case 32768:
                                        var v24 int64 = -9223372036854775808
                                        _ = v24
                                    case -210:
                                        if (!(189 == 127)) {
                                            if (0 == 32767) {
                                                var v25 float64 = 0.1
                                                fmt.Println(v25, v16)
                                                if (-1 == 1) {
                                                    var v26 int16 = 327
                                                    switch "a" {
                                                    case "a":
                                                        v4 += (v4 - v4)
                                                    }
                                                    v21 -= -128
                                                    _ = v26
                                                } else {
                                                    var v27 uint32 = (v15 * (^(v15)))
                                                    v15 += 65535
                                                    v28 := GenS3{F0: (false || v23), F1: 206, F2: 1}
                                                    _ = v27
                                                    _ = v28
                                                }
                                                _ = v25
                                            }
                                            var v29 bool = (v23 || v23)
                                            v3[v15] = (fi17 << ((18446744073709551615) & 63))
                                            _ = v29
                                        } else {
                                            v30 := GenS2{F0: true, F1: (v23 || v23)}
                                            _ = v30
                                        }
                                    case -92:
                                        delete(v3, v15)
                                    default:
                                        v4 = v4
                                    }
                                }
                            }
                            _ = v22
                            _ = v23
                        }
                    case 210:
                        var v31 float32 = (float32(-2.25) - -2.25)
                        _ = v31
                    case 128:
                        switch "a" {
                        case "\n\t":
                            switch v16 {
                            case 4294967295:
                                fmt.Println(v15, fi17, v16)
                            case 32767:
                                v15 -= (v15 ^ 42)
                            }
                        case "a":
                            var v32 float32 = 1e10
                            _ = v32
                        case "":
                            v33 := GenS4{F0: true, F1: (v21 - -4)}
                            _ = v33
                        default:
                            var v34 float32 = float32(float64(v4))
                            _ = v34
                        }
                    }
                    _ = v21
                }
                _ = fi18
            }
            _ = fi17
        }
        switch (-(v16)) {
        case 0:
            fmt.Println(v15, v15, v4)
        case 935:
            switch "foo bar" {
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        default:
            if false {
                v35 := GenS5{F0: -2147483648, F1: 0}
                _ = v35
            }
        }
    }
    var v36 int32 = genFn0(127, "", 255)
    var v37 int64 = 592
    var v38 uint = 4294967295
    _ = len(v2)
    v39 := v3[(v15 % ((0) | 1))]
    var v40 uint16 = 65535
    v16 = (-(v16))
    var v41 uint8 = 2
    var v42 int32 = int32(v16)
    fmt.Println(v16, v38, v39)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
