package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint }
type GenS3 struct { F0 bool; F1 int8; F2 uint }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 int8 }
func genFn0() int64 {
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 float64 = ((float64(-2.25) / -2.25) * 0.1)
        _ = fi6
        _ = v7
    }
    var v8 int8 = -6
    fmt.Println(v8)
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 string, p1 int) int {
    var v9 string = p0
    v10 := []int{p1}
    var v11 uint32 = 964
    switch (p1 | p1) {
    case 42:
        if (true && true) {
            fmt.Println(v11)
            var v12 float64 = 3.14159
            _ = v12
        }
    case 128:
        var v13 float64 = (float64(0.0) + 3.14159)
        _ = v13
    case -245:
        v14 := GenS1{F0: 0, F1: false, F2: 455}
        _ = v14
    }
    v15 := GenS2{F0: 0, F1: 51, F2: 660}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v15
    _ = v9
    return (p1 & 42)
}

func main() {
    var v1 int8 = 42
    var v2 uint8 = 106
    v3 := make(map[string]int)
    var v4 string = "a"
    var v5 int32 = -1
    v16 := GenS3{F0: (!true), F1: (v1 | (v1 + v1)), F2: 42}
    if true {
        v2 += (v2 % ((v2) | 1))
        switch "a" {
        case "foo bar":
            v17 := GenS4{F0: (v4 + ""), F1: (v4 + (v4 + v4))}
            _ = v17
        case "a":
            if true {
                if false {
                    var v18 uint32 = 1
                    fmt.Println(v18)
                    v19 := GenS1{F0: 1, F1: true, F2: 0}
                    _ = v18
                    _ = v19
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        switch v4 {
                        case "a":
                            for fi20 := 0; fi20 < 17; fi20++ {
                                if false {
                                    switch (fi20 ^ 698) {
                                    case 128:
                                        var v21 uint = 0
                                        _ = v21
                                    case -166:
                                        v1 += v1
                                    default:
                                        v22 := GenS2{F0: (v2 / ((v2) | 1)), F1: v5, F2: 819}
                                        _ = v22
                                    }
                                }
                                _ = fi20
                            }
                        case "hello":
                            v23 := v3[v4]
                            _ = v23
                        case "":
                            fmt.Println(v2, v2)
                        }
                    case "a":
                        v4 = ("a" + v4)
                    }
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        v2 += 128
                    default:
                        var v24 float64 = 3.14159
                        _ = v24
                    }
                    var v25 int16 = 188
                    _ = v25
                }
            } else {
                var v26 float32 = 1e10
                var v27 int64 = -9223372036854775808
                _ = v26
                _ = v27
            }
        case "hello":
            v28 := v3["a"]
            _ = v28
        }
    }
    var v29 uint8 = v2
    var v30 uint8 = v2
    v31 := []int{127, 9223372036854775807}
    if (!(4294967295 != 549)) {
        var v32 uint16 = 32767
        if true {
            fmt.Println(v30)
        }
        _ = v32
    } else {
        var v33 int64 = 42
        v34 := GenS5{F0: 907}
        v3[v4] = 0
        _ = v33
        _ = v34
    }
    fmt.Println(v1, v29)
    v35 := GenS4{F0: v4, F1: v4}
    for fi36 := 0; fi36 < 1; fi36++ {
        fmt.Println(v2)
        var v37 float32 = float32(3.14159)
        for fi38 := 0; fi38 < 8; fi38++ {
            fmt.Println(v2, v2, fi36)
            _ = fi38
        }
        _ = fi36
        _ = v37
    }
    if (((1e10 * (0.1 / 3.14159)) + float32(1e10)) <= 0.1) {
        v1 += (v1 / ((v1) | 1))
        var v39 int16 = 13
        if (true && false) {
            var v40 uint8 = (v2 / (((^(v2))) | 1))
            if false {
                fmt.Println(v2, v1, v40)
                fmt.Println(v5, v40)
            }
            _ = v40
        } else {
            for fi41 := 0; fi41 < 7; fi41++ {
                for fi42 := 0; fi42 < 16; fi42++ {
                    switch 127 {
                    case 832:
                        var v43 int8 = (^(v1))
                        _ = v43
                    }
                    fmt.Println(v2)
                    v2 = 42
                    switch fi42 {
                    case 32767:
                        v4 = "a"
                    case 42:
                        if false {
                            v44 := GenS2{F0: 73, F1: (v5 >> ((204) & 63)), F2: 4294967295}
                            if (-9223372036854775808 != genFn0()) {
                                v45 := GenS6{F0: 32767, F1: v1}
                                var v46 string = v4
                                _ = v45
                                _ = v46
                            }
                            _ = v44
                        }
                    case -9223372036854775808:
                        v30 += (v30 << ((4294967295) & 63))
                    }
                    _ = fi42
                }
                v30 -= v2
                _ = fi41
            }
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
        }
        _ = v39
    }
    fmt.Println(v30, v30)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
}
