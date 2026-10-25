package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 bool = true
    var v3 uint16 = 907
    v4 := make(map[uint32]int)
    var v5 int64 = 42
    v6 := []int{}
    fmt.Println(v1, v2)
    switch "hello" {
    case "":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println(v2)
    var v7 string = ("" + "foo bar")
    v8 := GenS1{F0: v3, F1: 195, F2: -2147483648}
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 float64 = 0.1
        switch v7 {
        case "\n\t":
            var v11 uint64 = 18446744073709551615
            _ = v11
        case "foo bar":
            v12 := GenS1{F0: 32767, F1: 0, F2: 42}
            _ = v12
        case "a":
            v7 = v7
        default:
            if v2 {
                for fi13 := 0; fi13 < 6; fi13++ {
                    switch fi9 {
                    case 9223372036854775807:
                        v14 := GenS1{F0: (v3 << (((v1 % ((v1) | 1))) & 63)), F1: -1, F2: 2147483647}
                        fmt.Println(v3, v7, fi13)
                        _ = v14
                    case 362:
                        switch "a" {
                        case "foo bar":
                            var v15 int = fi13
                            _ = v15
                        case "hello":
                            switch (fi9 ^ fi9) {
                            case 2147483647:
                                _ = len(v4)
                            case 0:
                                v16 := GenS1{F0: (v3 - 42), F1: 42, F2: -1}
                                _ = v16
                            default:
                                fmt.Println(v10)
                            }
                        default:
                            if true {
                                if (418 <= 18446744073709551615) {
                                    v3 = v3
                                    v7 = v7
                                    switch (fi9 >> ((v1) & 63)) {
                                    case 0:
                                        fmt.Println(v7, v10)
                                    case 306:
                                        v17 := v4[0]
                                        _ = v17
                                    case -921:
                                        if (!v2) {
                                            fmt.Println(v3, fi13)
                                            var v18 bool = false
                                            if true {
                                                v5 = (v5 / ((42) | 1))
                                                if (!(2147483647 != -147)) {
                                                    if len(v6) > 0 {
                                                        v6[0] = int(v1)
                                                    }
                                                    v19 := v4[1]
                                                    var v20 int32 = 0
                                                    _ = v19
                                                    _ = v20
                                                }
                                            }
                                            var v21 uint64 = 677
                                            v22 := GenS1{F0: uint16(v21), F1: -1, F2: 0}
                                            _ = v18
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                } else {
                                    var v23 float32 = float32(1.5)
                                    v24 := GenS2{F0: v1}
                                    v25 := GenS2{F0: 2}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                var v26 bool = v2
                                _ = v26
                            } else {
                                v10 -= v10
                            }
                        }
                        v7 = v7
                    default:
                        if false {
                            v10 = (((v10 / v10) - (1.5 / (1.5 + 1e10))) - ((v10 - (3.14159 / 3.14159)) / ((v10 * (float64(1.5) / 1.0)) + 0.0)))
                            var v27 uint8 = 127
                            _ = v27
                        } else {
                            var v28 int32 = 1
                            v5 = 0
                            if ((float64((-2.25 * 0.1)) - (v10 - v10)) < (v10 / v10)) {
                                if v2 {
                                    if v2 {
                                        v28 = v28
                                        fmt.Println(fi13)
                                        switch v7 {
                                        case "\n\t":
                                            v10 = v10
                                        case "hello":
                                            v5 += 1
                                        }
                                    }
                                    var v29 int32 = (v28 % ((v28) | 1))
                                    v3 -= v3
                                    _ = v29
                                } else {
                                    v6 = append(v6, (fi13 & 42))
                                    v4[0] = 0
                                }
                                fmt.Println(v5, v1, v2)
                                var v30 string = v7
                                _ = v30
                            } else {
                                v2 = false
                                if (1 != 0) {
                                    if (float32(v10) < 1.5) {
                                        switch v7 {
                                        case "foo bar":
                                            var v31 int8 = 100
                                            _ = v31
                                        default:
                                            v28 += v28
                                        }
                                    } else {
                                        var v32 bool = ((!((0.1 - float64((3.14159 * 3.14159))) != 3.14159)) || ((!(v2 || (!((v1 ^ (v1 >> (((v1 ^ (v1 >> (((v1 % ((466) | 1))) & 63)))) & 63))) <= 42)))) && false))
                                        _ = v32
                                    }
                                    _ = len(v4)
                                }
                                var v33 int32 = 881
                                _ = v33
                            }
                            _ = v28
                        }
                    }
                    fmt.Println(v3, v7, fi13)
                    var v34 float32 = (1e10 * 1.5)
                    _ = fi13
                    _ = v34
                }
                var v35 int16 = 42
                var v36 string = ("foo bar" + "foo bar")
                _ = v35
                _ = v36
            } else {
                var v37 uint32 = 2
                _ = v37
            }
        }
        _ = fi9
        _ = v10
    }
    var v38 int = 32768
    var v39 uint = (v1 - 2)
    v40 := GenS3{F0: 89}
    v2 = v2
    fmt.Println(v2, v39)
    switch "hello" {
    case "":
        var v41 float64 = (1.5 - -2.25)
        _ = v41
    case "hello":
        v39 = (v39 / ((39) | 1))
    default:
        v42 := GenS3{F0: -19}
        _ = v42
    }
    var v43 int = 127
    for fi44 := 0; fi44 < 5; fi44++ {
        if false {
            v39 = (v1 | (v39 | v1))
            v4[65535] = v43
            fmt.Println(v38, v38)
        }
        var v45 uint8 = 127
        _ = fi44
        _ = v45
    }
    var v46 bool = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
