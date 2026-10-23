package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int8 }
type GenS2 struct { F0 int32; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint8; F2 string }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint64; F1 int; F2 uint8 }
type GenS8 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 2147483647
    var v3 uint64 = 133
    v4 := GenS1{F0: -1, F1: (v3 / (((v3 ^ 651)) | 1)), F2: 127}
    if false {
        var v5 uint8 = 127
        switch "hello" {
        case "":
            for fi6 := 0; fi6 < 10; fi6++ {
                v7 := GenS2{F0: -1, F1: 42, F2: 32767}
                v3 += (v3 - v3)
                _ = fi6
                _ = v7
            }
        case "foo bar":
            fmt.Println(v3, v5)
        case "hello":
            _ = len(v1)
        }
        _ = v5
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := v1["foo bar"]
            _ = fi9
            _ = v10
        }
        var v11 uint32 = 2147483647
        _ = fi8
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        delete(v1, "a")
        v13 := GenS3{F0: 1}
        _ = fi12
        _ = v13
    }
    v1[("a" + "foo bar")] = true
    if false {
        _ = len(v1)
        v3 = v3
        for fi14 := 0; fi14 < 7; fi14++ {
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := GenS4{F0: 106, F1: 42, F2: "hello"}
                if false {
                    v17 := GenS3{F0: -1}
                    switch "\n\t" {
                    case "a":
                        switch ("\n\t" + (("foo bar" + "hello") + "hello")) {
                        case "foo bar":
                            if ((32767 < 42) && true) {
                                v18 := GenS2{F0: 1, F1: 0, F2: 467}
                                _ = v18
                            } else {
                                v3 = v3
                                if true {
                                    v19 := GenS1{F0: -128, F1: (^(v3)), F2: 127}
                                    _ = v19
                                }
                            }
                        case "\n\t":
                            var v20 uint8 = 255
                            _ = v20
                        case "a":
                            var v21 uint = 788
                            _ = v21
                        default:
                            var v22 uint32 = 0
                            _ = v22
                        }
                    default:
                        v23 := []int{fi14}
                        fmt.Println(v2, v3)
                        v3 = (^(v3))
                        _ = v23
                    }
                    v3 -= uint64(v2)
                    switch fi14 {
                    case 280:
                        var v24 int = (fi15 / (((fi14 & (fi15 | (^(fi15))))) | 1))
                        _ = v24
                    case 32768:
                        fmt.Println(fi15)
                    case 42:
                        v2 = (v2 | (v2 ^ 2147483648))
                    }
                    switch (fi14 / ((fi14) | 1)) {
                    case 322:
                        var v25 uint32 = 757
                        _ = v25
                    }
                    _ = v17
                } else {
                    switch (fi15 ^ (fi15 * (^(fi15)))) {
                    case 0:
                        v3 -= v3
                    case 4294967295:
                        if (((-2.25 / -2.25) * 0.0) > 0.0) {
                            var v26 int8 = 64
                            v26 += 88
                            v2 += v2
                            if (false && false) {
                                if (!false) {
                                    v27 := []int{}
                                    v2 = v2
                                    if (!true) {
                                        v28 := GenS2{F0: 42, F1: -1, F2: 957}
                                        v29 := GenS2{F0: -744, F1: 0, F2: 21}
                                        fmt.Println(fi15, v3, v2)
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = v27
                                }
                            } else {
                                v2 = 307
                                var v30 int16 = 0
                                var v31 uint32 = 1
                                _ = v30
                                _ = v31
                            }
                            if true {
                                v32 := GenS5{F0: 9223372036854775807, F1: 255}
                                fmt.Println(v2, fi15)
                                var v33 int32 = 1
                                _ = v32
                                _ = v33
                            }
                            _ = v26
                        } else {
                            v34 := GenS6{F0: (42 > 379), F1: 65535}
                            if false {
                                var v35 float64 = 1e10
                                if ("hello" == "hello") {
                                    v36 := []int{fi14, fi15}
                                    _ = v36
                                }
                                _ = v35
                            }
                            _ = v34
                        }
                    case 1:
                        v2 += (v2 + (^(v2)))
                    default:
                        fmt.Println(fi15, v2, v2)
                    }
                    v3 = (^(v3))
                    var v37 float32 = 1e10
                    _ = v37
                }
                v3 += (v3 - v3)
                _ = fi15
                _ = v16
            }
            if false {
                if false {
                    for fi38 := 0; fi38 < 17; fi38++ {
                        var v39 bool = false
                        _ = fi38
                        _ = v39
                    }
                }
                v3 += 42
                var v40 bool = (-128 < 127)
                _ = v40
            }
            _ = fi14
        }
    }
    var v41 uint16 = 670
    v3 += 751
    v2 = v2
    v42 := GenS7{F0: (v3 >> ((2) & 63)), F1: 0, F2: 211}
    fmt.Println(v2, v41)
    v43 := GenS8{F0: 1, F1: (1.5 >= ((float32(0.1) * 0.0) - 1e10))}
    v44 := []int{64, 603, 42, 32768}
    for fi45 := 0; fi45 < 19; fi45++ {
        v2 = uint32(v41)
        var v46 string = ""
        _ = fi45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
