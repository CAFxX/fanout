package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int16; F2 string }
type GenS5 struct { F0 int16; F1 int32; F2 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 bool; F2 uint64 }
type GenS8 struct { F0 uint32; F1 int32; F2 int16 }
type GenS9 struct { F0 int16; F1 uint32; F2 uint16 }
func genFn0(p0 uint16, p1 bool) int64 {
    var v5 float64 = 0.0
    v6 := []int{-2147483648, -1, -2147483648, 663}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 int8, p2 uint16) int64 {
    fmt.Println(p1, p0)
    p2 = (p2 ^ p2)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := []int{9223372036854775807, 4294967295, 127, -1}
    var v2 uint64 = 42
    var v3 float64 = -2.25
    var v4 uint = 733
    for fi8 := 0; fi8 < 3; fi8++ {
        _ = len(v1)
        _ = fi8
    }
    v9 := GenS1{F0: 32767, F1: 42}
    v2 += uint64(v4)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint8 = uint8(v4)
        _ = fi10
        _ = v11
    }
    if false {
        v12 := GenS1{F0: 42, F1: -1}
        for fi13 := 0; fi13 < 3; fi13++ {
            v2 -= uint64(v4)
            fmt.Println(fi13, v2)
            _ = fi13
        }
        _ = v12
    } else {
        v14 := GenS2{F0: 128}
        var v15 int16 = 32767
        v15 += v15
        _ = v14
        _ = v15
    }
    switch 2147483647 {
    case -101:
        v4 += v4
    case 2147483647:
        v3 += v3
    case -532:
        var v16 float32 = 0.0
        _ = v16
    }
    switch 4294967295 {
    case 952:
        fmt.Println(v2, v4, v2)
    }
    v17 := GenS3{F0: 42, F1: 2, F2: 1}
    fmt.Println(v2, v3)
    switch "\n\t" {
    case "hello":
        if (v4 <= (v4 & (v4 | 255))) {
            _ = len(v1)
            switch -369 {
            case 128:
                fmt.Println(v3, v2, v2)
            }
        } else {
            if len(v1) > 0 {
                v1[0] = 32768
            }
            v3 -= 3.14159
            var v18 bool = (v4 < v4)
            _ = v18
        }
    case "\n\t":
        v19 := GenS4{F0: 255, F1: 1, F2: "foo bar"}
        _ = v19
    case "":
        var v20 uint = uint(v2)
        _ = v20
    default:
        var v21 float64 = (float64(1.5) - 0.0)
        _ = v21
    }
    v2 += 331
    var v22 int64 = -9223372036854775808
    switch 127 {
    case -1:
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 float32 = -2.25
            if true {
                var v25 uint64 = (v2 >> ((2) & 63))
                v25 = (v2 % ((2) | 1))
                _ = len(v1)
                _ = v25
            }
            if ("hello" > "") {
                if (true && (198 < 42)) {
                    for fi26 := 0; fi26 < 15; fi26++ {
                        var v27 string = "a"
                        var v28 bool = true
                        switch 128 {
                        case 4294967295:
                            if true {
                                v24 = v24
                            } else {
                                v29 := GenS5{F0: 0, F1: -1, F2: (v2 * (v2 ^ 4294967295))}
                                if (((false && v28) && (2147483647 < 2147483647)) && (2147483647 < 2147483648)) {
                                    if v28 {
                                        if v28 {
                                            var v30 string = v27
                                            v24 = ((float32(float64(v24)) / (v24 / (v24 / v24))) + v24)
                                            _ = v30
                                        }
                                    } else {
                                        if v28 {
                                            var v31 uint16 = 32767
                                            _ = v31
                                        } else {
                                            var v32 string = ((v27 + ("" + v27)) + ("" + (v27 + ((v27 + v27) + "\n\t"))))
                                            if ("foo bar" == "") {
                                                var v33 int64 = 42
                                                var v34 int32 = 2147483647
                                                _ = v33
                                                _ = v34
                                            }
                                            _ = v32
                                        }
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        v22 += v22
                                    case "hello":
                                        v27 = ""
                                    case "a":
                                        var v35 uint8 = 2
                                        _ = v35
                                    }
                                }
                                v36 := GenS6{F0: "\n\t"}
                                _ = v29
                                _ = v36
                            }
                            fmt.Println(v28, fi23)
                        case 127:
                            var v37 uint64 = 42
                            _ = v37
                        case 32767:
                            fmt.Println(v28, v27)
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(fi23, v3)
                    if true {
                        if true {
                            switch 713 {
                            case 128:
                                v2 -= v2
                            case -193:
                                v4 += v4
                            default:
                                v22 = -9223372036854775808
                            }
                            switch ("hello" + "\n\t") {
                            case "foo bar":
                                if len(v1) > 0 {
                                    v1[0] = (fi23 | fi23)
                                }
                            case "\n\t":
                                v38 := GenS7{F0: 0, F1: false, F2: (v2 << ((578) & 63))}
                                _ = v38
                            default:
                                var v39 float64 = (v3 / v3)
                                _ = v39
                            }
                        }
                        for fi40 := 0; fi40 < 1; fi40++ {
                            var v41 int = fi40
                            v42 := GenS4{F0: 42, F1: -743, F2: "foo bar"}
                            v43 := GenS8{F0: 42, F1: 2147483647, F2: 1000}
                            _ = fi40
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        if true {
                            var v44 string = "foo bar"
                            if false {
                                var v45 bool = true
                                if true {
                                    var v46 int32 = -2147483648
                                    var v47 int8 = 127
                                    fmt.Println(v44)
                                    _ = v46
                                    _ = v47
                                }
                                var v48 uint32 = 4294967295
                                _ = v45
                                _ = v48
                            }
                            var v49 int16 = 927
                            _ = v44
                            _ = v49
                        }
                    } else {
                        var v50 int8 = 94
                        var v51 uint32 = 282
                        var v52 uint8 = 0
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                }
            }
            _ = fi23
            _ = v24
        }
    case 0:
        v53 := GenS9{F0: 422, F1: 286, F2: 255}
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
