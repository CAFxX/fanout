package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 int64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int16; F1 bool; F2 bool }
func genFn0(p0 int64) int16 {
    v7 := []int{0, -1}
    var v8 int16 = 32767
    _ = p0
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 int8, p1 uint16, p2 uint8) int8 {
    var v9 int32 = -70
    var v10 string = "foo bar"
    p2 = (p2 | (p2 - (p2 << ((957) & 63))))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (-(p0))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    var v3 float32 = 1e10
    var v4 uint16 = 852
    var v5 int32 = 2147483647
    var v6 float64 = float64(((float32(3.14159) * (3.14159 + v3)) * v3))
    v6 += 3.14159
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(v5)
        fmt.Println(v3, fi11, v3)
        _ = fi11
    }
    if (0.0 >= float64(v3)) {
        var v12 int32 = v5
        v13 := []int{32767, 0}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS1{F0: 1, F1: 1, F2: 785}
            if (!false) {
                v16 := GenS2{F0: -9223372036854775808, F1: 1, F2: -9223372036854775808}
                _ = v16
            } else {
                if (249 > 255) {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        switch ("a" + "foo bar") {
                        case "\n\t":
                            v18 := GenS1{F0: 60, F1: 1, F2: 42}
                            _ = v18
                        case "foo bar":
                            v19 := []int{0, fi14}
                            _ = v19
                        case "":
                            v1[fi17] = (!false)
                        default:
                            v20 := GenS3{F0: "\n\t", F1: 710}
                            _ = v20
                        }
                        if (!(true && false)) {
                            var v21 int8 = 42
                            v22 := []int{(fi14 & fi17)}
                            _ = v21
                            _ = v22
                        }
                        v6 += 0.1
                        fmt.Println(v4, v5)
                        var v23 float64 = (float64((0.0 - (v3 - ((((float32((v6 / -2.25)) * v3) * 0.0) + float32(1e10)) / float32(float64(v3)))))) / (float64(-2.25) * v6))
                        v23 = (0.1 + (v23 / v23))
                        var v24 uint64 = 1
                        if true {
                            var v25 int32 = v5
                            _ = v25
                        }
                        _ = fi17
                        _ = v23
                        _ = v24
                    }
                } else {
                    var v26 int8 = 127
                    v5 = v5
                    v27 := []int{}
                    _ = v26
                    _ = v27
                }
                var v28 uint16 = (v4 * v4)
                _ = v28
            }
            _ = fi14
            _ = v15
        }
        v29 := GenS3{F0: "a", F1: 0}
        _ = v29
    }
    switch -9 {
    case 128:
        if true {
            v30 := GenS4{F0: 42}
            _ = v30
        }
    case 859:
        if false {
            var v31 float64 = (v6 - float64(v3))
            _ = v31
        } else {
            v5 = v5
        }
    case 289:
        v3 = ((1.5 - 0.1) / (v3 + v3))
    }
    v5 = (v5 & v5)
    v32 := GenS5{F0: true, F1: 255, F2: 189}
    var v33 uint16 = v4
    if (!false) {
        fmt.Println(v4, v6, v4)
        v34 := GenS3{F0: ("a" + "a"), F1: 917}
        fmt.Println(v4, v33)
        _ = v34
    }
    v6 -= v6
    v1[32767] = true
    for fi35 := 0; fi35 < 17; fi35++ {
        var v36 uint = 42
        var v37 int8 = genFn1(127, v33, 127)
        switch "" {
        case "hello":
            if ((fi35 & fi35) == (fi35 % ((fi35) | 1))) {
                if true {
                    v37 = (-(v37))
                } else {
                    v38 := GenS6{F0: 9223372036854775807}
                    var v39 uint16 = v33
                    v40 := v2[fi35]
                    _ = v38
                    _ = v39
                    _ = v40
                }
                if false {
                    var v41 uint = uint(v33)
                    v42 := GenS7{F0: false, F1: -1}
                    var v43 float32 = v3
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    v44 := []int{(fi35 + fi35), fi35, (-(fi35))}
                    switch "" {
                    case "hello":
                        if len(v44) > 0 {
                            v44[0] = fi35
                        }
                    case "":
                        if ("a" != "") {
                            v3 = (((v3 - (v3 / 3.14159)) + v3) - v3)
                        }
                    }
                    fmt.Println(v37, fi35)
                    _ = v44
                }
            }
        case "foo bar":
            v4 = v4
        default:
            fmt.Println(v37, v3)
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    v33 = 0
    v45 := GenS1{F0: 72, F1: 67, F2: 42}
    switch -9223372036854775808 {
    case 42:
        v46 := GenS8{F0: 32767, F1: true, F2: true}
        _ = v46
    }
    v33 = (v4 + (^(v33)))
    v5 -= v5
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
