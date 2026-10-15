package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint64; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 uint16; F2 bool }
type GenS7 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (2 < 2)
    v3 := []uint32{2147483647, 4294967295, 484, 42}
    var v4 float64 = 1e10
    var v5 uint8 = 2
    v1[uint32(v5)] = 2147483647
    v6 := GenS1{F0: -1, F1: 63, F2: 581}
    var v7 int8 = 100
    fmt.Println(v5, v2)
    switch ("" + "") {
    case "\n\t":
        if (!v2) {
            if (215 >= 672) {
                if v2 {
                    if v2 {
                        var v8 uint8 = v5
                        var v9 int8 = (v7 >> ((1) & 63))
                        _ = v8
                        _ = v9
                    }
                }
            }
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 int16 = 0
                _ = fi10
                _ = v11
            }
            _ = len(v3)
        } else {
            for fi12 := 0; fi12 < 19; fi12++ {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v3 = append(v3, 1)
                    var v14 int32 = -409
                    if v2 {
                        if (-2.25 != -2.25) {
                            v15 := GenS2{F0: 1}
                            _ = v15
                        }
                        fmt.Println(v7, v4, v7)
                    } else {
                        v4 = -2.25
                        switch "hello" {
                        case "\n\t":
                            v2 = true
                        case "hello":
                            fmt.Println(v14, fi12)
                        case "foo bar":
                            if (!(-1 >= v14)) {
                                var v16 uint16 = 32767
                                if v2 {
                                    v1[0] = (fi12 >> ((1) & 63))
                                    v17 := GenS3{F0: 2}
                                    v5 = (v5 + (v5 + (v5 - 255)))
                                    _ = v17
                                } else {
                                    var v18 int64 = 1
                                    _ = v18
                                }
                                v14 = (v14 / ((v14) | 1))
                                _ = v16
                            }
                        default:
                            v2 = false
                        }
                    }
                    switch (((("" + "a") + "foo bar") + "a") + ("\n\t" + ("\n\t" + "\n\t"))) {
                    case "":
                        if (!(!v2)) {
                            v1[2147483647] = (fi12 * 812)
                        } else {
                            v1[2147483648] = 135
                            if v2 {
                                v14 = (v14 ^ (v14 + (v14 * (v14 % ((v14) | 1)))))
                                if v2 {
                                    var v19 bool = ((float64(1e10) * float64(1.5)) == v4)
                                    var v20 float64 = v4
                                    _ = v19
                                    _ = v20
                                }
                            }
                            v21 := GenS3{F0: 1}
                            v14 = (v14 + (-(v14)))
                            v1[462] = (fi13 % (((fi13 + (fi12 << ((691) & 63)))) | 1))
                            _ = v21
                        }
                    case "hello":
                        if (!(v14 > (v14 ^ (v14 ^ v14)))) {
                            var v22 uint8 = (v5 << ((1) & 63))
                            if v2 {
                                v7 -= (v7 % ((127) | 1))
                                var v23 uint32 = 628
                                if v2 {
                                    var v24 float32 = 0.0
                                    _ = v24
                                } else {
                                    if v2 {
                                        if v2 {
                                            if len(v3) > 0 {
                                                v3[0] = 42
                                            }
                                            if true {
                                                v7 = (v7 + (^(v7)))
                                                v25 := GenS2{F0: -1}
                                                if v2 {
                                                    v26 := GenS1{F0: 32767, F1: 0, F2: v23}
                                                    _ = v26
                                                } else {
                                                    v27 := GenS1{F0: fi13, F1: 0, F2: 555}
                                                    switch "a" {
                                                    case "foo bar":
                                                        v3 = append(v3, (v23 ^ v23))
                                                        v28 := GenS4{F0: 9223372036854775807, F1: 283, F2: "\n\t"}
                                                        _ = v28
                                                    case "\n\t":
                                                        switch fi13 {
                                                        case 943:
                                                            var v29 int32 = -1
                                                            _ = v29
                                                        case 0:
                                                            v30 := GenS5{F0: 2147483648}
                                                            _ = v30
                                                        }
                                                    default:
                                                        var v31 int = (fi12 ^ (fi12 * -2147483648))
                                                        _ = v31
                                                    }
                                                    _ = v27
                                                }
                                                _ = v25
                                            }
                                            if (((float64(((((1.5 - 1.5) - ((1.5 / (float32((v4 * v4)) / 1.0)) / 1e10)) - 0.0) / (3.14159 - -2.25))) + v4) * ((v4 / (((float64((float32(((((1e10 + (((1.5 - float64(((0.0 + 0.0) + -2.25))) / (3.14159 * (v4 + v4))) - float64(1e10))) + ((v4 / 1.5) / (1e10 - (1e10 / v4)))) / (v4 - 0.1)) - v4)) + 1.5)) + 3.14159) * float64(0.0)) * v4)) * v4)) > v4) {
                                                fmt.Println(v2)
                                                var v32 int32 = (^(v14))
                                                _ = v32
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    var v33 float32 = -2.25
                                                    _ = v33
                                                }
                                                v34 := GenS3{F0: 65535}
                                                _ = v34
                                            }
                                        } else {
                                            fmt.Println(v23, v5, fi12)
                                            var v35 float64 = v4
                                            if (!v2) {
                                                v3 = append(v3, (v23 / ((v23) | 1)))
                                                v4 -= (v4 - (v35 * ((v4 * ((v35 * float64(0.1)) * v4)) * v4)))
                                            } else {
                                                if false {
                                                    v23 = v23
                                                    v36 := GenS6{F0: (^(v23)), F1: 65535, F2: v2}
                                                    fmt.Println(v22, v4, v14)
                                                    _ = v36
                                                }
                                            }
                                            _ = v35
                                        }
                                        var v37 string = (((("foo bar" + "\n\t") + "") + "a") + "hello")
                                        var v38 float64 = (1e10 / v4)
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v39 := GenS7{F0: v23, F1: v14}
                                        _ = v39
                                    }
                                    var v40 int32 = v14
                                    var v41 string = "foo bar"
                                    _ = v40
                                    _ = v41
                                }
                                var v42 uint32 = v23
                                _ = v23
                                _ = v42
                            } else {
                                var v43 string = ("\n\t" + (((((("foo bar" + "foo bar") + "\n\t") + "") + "foo bar") + ("" + "\n\t")) + "hello"))
                                _ = v43
                            }
                            _ = v22
                        }
                    }
                    if (v2 || v2) {
                        fmt.Println(v7)
                    } else {
                        var v44 int16 = -32768
                        _ = v44
                    }
                    v45 := GenS3{F0: 42}
                    _ = fi13
                    _ = v14
                    _ = v45
                }
                _ = fi12
            }
            var v46 uint64 = uint64(v5)
            _ = v46
        }
    }
    var v47 uint16 = 32767
    v4 += -2.25
    v7 -= (v7 ^ v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
