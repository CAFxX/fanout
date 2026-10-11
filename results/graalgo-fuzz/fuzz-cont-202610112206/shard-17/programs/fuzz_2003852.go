package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 int8; F2 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 int = 555
    var v3 float32 = 0.1
    var v4 uint32 = 0
    var v5 string = "hello"
    delete(v1, v5)
    delete(v1, (v5 + (v5 + "")))
    switch (v5 + v5) {
    case "a":
        v6 := GenS1{F0: v2, F1: "a", F2: (v2 % ((v2) | 1))}
        _ = v6
    case "\n\t":
        v5 = ((v5 + (("\n\t" + v5) + v5)) + ("a" + ("" + v5)))
    default:
        for fi7 := 0; fi7 < 20; fi7++ {
            v2 -= fi7
            var v8 uint16 = 0
            _ = fi7
            _ = v8
        }
    }
    v9 := v1[(v5 + (("foo bar" + "\n\t") + v5))]
    var v10 uint = 18446744073709551615
    fmt.Println(v10)
    var v11 int16 = 1
    fmt.Println(v4)
    v10 = (v10 & v10)
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: (v2 << ((v10) & 63)), F1: v5, F2: 2147483647}
        v14 := []int{(^(v9)), (^(v9)), v9, v9}
        v11 = v11
        _ = fi12
        _ = v13
        _ = v14
    }
    switch "a" {
    case "":
        var v15 int16 = v11
        _ = v15
    }
    var v16 int = (v9 - (v9 - (v2 + v2)))
    v17 := GenS2{F0: 2147483648, F1: (v11 / (((v11 % ((0) | 1))) | 1))}
    v4 = 0
    fmt.Println(v10, v9, v10)
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 bool = false
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            if false {
                switch ("\n\t" + v5) {
                case "foo bar":
                    fmt.Println(v16)
                case "hello":
                    v9 += -2147483648
                case "":
                    switch (v9 * (fi21 / ((32768) | 1))) {
                    case -9223372036854775808:
                        fmt.Println(v3)
                    case -370:
                        if (int8(v4) < 127) {
                            if true {
                                v22 := GenS2{F0: (v4 - uint32(v10)), F1: v11}
                                _ = v22
                            }
                        }
                    case 127:
                        v23 := []int{4294967295}
                        _ = v23
                    default:
                        var v24 float64 = 3.14159
                        _ = v24
                    }
                }
                v16 = (v2 & (v9 - 32767))
                if false {
                    v5 = v5
                }
                if (-9223372036854775808 != -9223372036854775808) {
                    var v25 int8 = 34
                    switch "foo bar" {
                    case "hello":
                        v26 := GenS2{F0: (v4 << (((v10 + v10)) & 63)), F1: v11}
                        _ = v26
                    default:
                        v27 := GenS1{F0: -259, F1: "\n\t", F2: v2}
                        _ = v27
                    }
                    if true {
                        if false {
                            var v28 float64 = 0.0
                            _ = v28
                        } else {
                            fmt.Println(v2, v5)
                            v3 = 1e10
                            v25 -= (v25 + v25)
                        }
                        v11 += v11
                        v29 := GenS2{F0: v4, F1: (v11 + (v11 + (^(v11))))}
                        _ = v29
                    } else {
                        v30 := []int{v16}
                        _ = v30
                    }
                    _ = v25
                }
                switch "hello" {
                case "foo bar":
                    v31 := GenS3{F0: 580}
                    _ = v31
                case "":
                    if ((-(v11)) < (-(v11))) {
                        var v32 bool = (136 <= 2)
                        fmt.Println(v4)
                        v1[v5] = (fi21 & v9)
                        _ = v32
                    } else {
                        v33 := v1[""]
                        v3 += (0.1 - 0.0)
                        _ = v33
                    }
                case "hello":
                    v5 = "\n\t"
                default:
                    v34 := GenS1{F0: fi20, F1: v5, F2: v16}
                    _ = v34
                }
            } else {
                switch v16 {
                case 493:
                    v35 := GenS1{F0: (fi20 ^ v2), F1: (v5 + "hello"), F2: (-(v16))}
                    _ = v35
                case 0:
                    switch fi21 {
                    case -1:
                        v36 := []int{}
                        _ = v36
                    }
                }
                v9 = 0
                var v37 uint64 = 18446744073709551615
                _ = v37
            }
            if (false && (255 < 1)) {
                v38 := GenS4{F0: (v5 + v5), F1: 0, F2: 1}
                _ = v38
            } else {
                if false {
                    if true {
                        v9 -= (v2 >> ((v10) & 63))
                    }
                } else {
                    v9 = int(v4)
                    v1[v5] = 4294967295
                    var v39 int32 = -716
                    v40 := []int{fi20, v16, -474, (v2 / (((v9 % (((fi21 >> (((v10 * 1)) & 63))) | 1))) | 1))}
                    var v41 int32 = 57
                    var v42 uint64 = 361
                    v11 = (v11 / ((1) | 1))
                    var v43 uint8 = 0
                    _ = len(v40)
                    if (!false) {
                        var v44 float32 = -2.25
                        _ = v44
                    } else {
                        v40 = append(v40, (fi21 / ((fi21) | 1)))
                        var v45 string = v5
                        _ = v45
                    }
                    var v46 int8 = 127
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                    _ = v43
                    _ = v46
                }
                v16 = v9
            }
            _ = fi21
        }
        var v47 int16 = int16(v2)
        _ = fi20
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
