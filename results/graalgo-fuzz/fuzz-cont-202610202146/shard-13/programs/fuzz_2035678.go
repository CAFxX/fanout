package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint }
type GenS2 struct { F0 int; F1 int16; F2 uint64 }
type GenS3 struct { F0 string; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int32; F2 int8 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 int32; F1 uint8; F2 uint64 }

func main() {
    var v1 string = ((("hello" + "\n\t") + ("\n\t" + ("hello" + "\n\t"))) + (("hello" + "\n\t") + ""))
    var v2 bool = (-9223372036854775808 != 128)
    v3 := make(map[string]int)
    v4 := GenS1{F0: 461, F1: 42, F2: 0}
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS2{F0: (fi5 + (fi5 & fi5)), F1: -833, F2: 919}
        v7 := GenS2{F0: (fi5 / ((fi5) | 1)), F1: 1, F2: 2}
        _ = fi5
        _ = v6
        _ = v7
    }
    v1 = v1
    switch (v1 + v1) {
    case "":
        v1 = ""
    }
    var v8 float32 = (-2.25 * (0.0 - ((0.1 * 1.5) - (((-2.25 / 1e10) - -2.25) / 0.1))))
    v9 := GenS1{F0: 42, F1: -9223372036854775808, F2: 1}
    v2 = v2
    v10 := []int{-885, 158, 705}
    v1 = v1
    switch "foo bar" {
    case "\n\t":
        v11 := v3[(v1 + (v1 + ""))]
        _ = v11
    case "a":
        fmt.Println(v8, v2, v2)
    case "foo bar":
        _ = len(v3)
    default:
        v1 = v1
    }
    var v12 bool = true
    switch "foo bar" {
    case "a":
        for fi13 := 0; fi13 < 19; fi13++ {
            switch "a" {
            case "foo bar":
                if (42 < 9223372036854775807) {
                    _ = len(v3)
                    var v14 uint16 = 65535
                    v15 := GenS3{F0: v1, F1: 0, F2: 127}
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v12 = (0 != 32767)
                        v8 = v8
                        switch ((v1 + "a") + "a") {
                        case "hello":
                            v17 := v3[(v1 + (v1 + "hello"))]
                            _ = v17
                        default:
                            v8 += (v8 + (((v8 + 1e10) / v8) + (float32(3.14159) / v8)))
                        }
                        v1 = v1
                        var v18 int8 = -1
                        v12 = v12
                        fmt.Println(v12)
                        if v12 {
                            var v19 uint8 = 2
                            _ = v19
                        } else {
                            var v20 int8 = v18
                            _ = len(v3)
                            var v21 int64 = int64(v18)
                            _ = v20
                            _ = v21
                        }
                        _ = fi16
                        _ = v18
                    }
                }
            default:
                if ((909 != -1) && v2) {
                    v22 := GenS1{F0: 97, F1: 1, F2: 255}
                    _ = v22
                } else {
                    v2 = (18446744073709551615 == 499)
                    if v2 {
                        if true {
                            v2 = v12
                        } else {
                            var v23 int32 = -1
                            var v24 int16 = -1
                            fmt.Println(v24, v24)
                            _ = v23
                            _ = v24
                        }
                        var v25 int = fi13
                        v26 := GenS3{F0: v1, F1: 128, F2: 127}
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS3{F0: "", F1: 128, F2: 127}
                        _ = v27
                    }
                    v28 := GenS3{F0: v1, F1: 255, F2: 0}
                    _ = v28
                }
            }
            var v29 bool = v2
            for fi30 := 0; fi30 < 9; fi30++ {
                if (!(910 > 18446744073709551615)) {
                    _ = len(v10)
                    if ((((("a" + v1) + "hello") + "hello") + "hello") >= v1) {
                        if v29 {
                            v29 = v29
                            switch fi30 {
                            case -9223372036854775808:
                                v31 := v3[(v1 + v1)]
                                _ = v31
                            }
                            v10 = append(v10, (fi30 - (-(fi30))))
                            v32 := GenS3{F0: "hello", F1: 32, F2: 31}
                            if (712 != 42) {
                                if false {
                                    switch v1 {
                                    case "hello":
                                        _ = len(v10)
                                    case "foo bar":
                                        v8 = v8
                                    case "a":
                                        v33 := GenS1{F0: 500, F1: 42, F2: 485}
                                        _ = v33
                                    }
                                    v34 := GenS4{F0: 65535}
                                    _ = v34
                                } else {
                                    fmt.Println(v29)
                                    v35 := GenS3{F0: "", F1: 0, F2: -45}
                                    _ = v35
                                }
                                v36 := v3[(("\n\t" + "hello") + (v1 + v1))]
                                fmt.Println(v29)
                                v12 = (651 > 225)
                                switch v1 {
                                case "a":
                                    fmt.Println(v2, v29, v12)
                                case "\n\t":
                                    v10 = append(v10, 0)
                                    v1 = v1
                                    switch v1 {
                                    case "a":
                                        v2 = (-32768 >= 0)
                                    case "":
                                        var v37 int = v36
                                        _ = v37
                                    default:
                                        var v38 uint64 = 42
                                        _ = v38
                                    }
                                default:
                                    v39 := GenS1{F0: 143, F1: 1, F2: 524}
                                    _ = v39
                                }
                                if (-808 >= 9223372036854775807) {
                                    v40 := GenS5{F0: 42, F1: 2147483647, F2: -1}
                                    _ = v40
                                } else {
                                    switch v1 {
                                    case "":
                                        v41 := GenS1{F0: 42, F1: -9223372036854775808, F2: 255}
                                        _ = v41
                                    }
                                }
                                delete(v3, (("" + v1) + "foo bar"))
                                v12 = ((v8 > float32(float64(-2.25))) && (-653 == int64(fi13)))
                                _ = v36
                            } else {
                                if (v12 && v12) {
                                    v12 = v29
                                } else {
                                    _ = len(v10)
                                    v42 := GenS6{F0: 2, F1: v1}
                                    _ = v42
                                }
                            }
                            _ = v32
                        }
                    }
                }
                _ = fi30
            }
            _ = fi13
            _ = v29
        }
    case "":
        var v43 uint8 = 42
        _ = v43
    case "hello":
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    default:
        v44 := GenS7{F0: -1, F1: 35, F2: 42}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
