package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8, p1 int16) uint8 {
    p0 = p0
    var v5 uint8 = 2
    var v6 string = (("hello" + "foo bar") + "foo bar")
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (v5 | v5)
}
func genFn1(p0 uint32, p1 bool, p2 int8) uint16 {
    v7 := []int{2147483647, 9223372036854775807}
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            p1 = p1
            if (!(423 <= p0)) {
                var v10 int8 = -110
                switch "\n\t" {
                case "\n\t":
                    if true {
                        v11 := GenS1{F0: (fi8 & (fi8 % ((fi8) | 1)))}
                        if p1 {
                            if p1 {
                                fmt.Println(fi8, v10)
                                _ = len(v7)
                                v10 = p2
                            } else {
                                if (((^(v10)) > p2) && (p1 || p1)) {
                                    v10 = (p2 % ((v10) | 1))
                                } else {
                                    var v12 float64 = float64((3.14159 - 0.1))
                                    _ = v12
                                }
                                v7 = append(v7, fi9)
                                fmt.Println(v10, p2, p2)
                                var v13 uint8 = genFn0(int8(fi9), 543)
                                var v14 int64 = 42
                                _ = v13
                                _ = v14
                            }
                        } else {
                            var v15 uint = 255
                            v16 := GenS2{F0: (p0 + uint32(v15)), F1: ""}
                            _ = v15
                            _ = v16
                        }
                        if p1 {
                            v10 += (v10 ^ p2)
                            var v17 int8 = (p2 << ((117) & 63))
                            v18 := GenS3{F0: 1}
                            switch "\n\t" {
                            case "a":
                                var v19 int64 = 1
                                var v20 int8 = (v10 / ((127) | 1))
                                _ = v19
                                _ = v20
                            case "":
                                if p1 {
                                    v21 := GenS1{F0: fi9}
                                    _ = v21
                                } else {
                                    _ = len(v7)
                                    fmt.Println(v10)
                                }
                            case "foo bar":
                                var v22 uint32 = (^(p0))
                                _ = v22
                            default:
                                v7 = append(v7, (fi9 - 32767))
                            }
                            fmt.Println(v17, v17, p2)
                            var v23 float64 = -2.25
                            _ = v17
                            _ = v18
                            _ = v23
                        }
                        fmt.Println(fi9, fi9)
                        p0 -= p0
                        _ = v11
                    } else {
                        if (uint8(p0) >= 1) {
                            var v24 int = 128
                            _ = v24
                        }
                    }
                case "foo bar":
                    if (48 == 2) {
                        p1 = false
                        var v25 uint = 255
                        _ = len(v7)
                        if (!p1) {
                            var v26 int8 = 42
                            _ = len(v7)
                            _ = v26
                        } else {
                            var v27 int64 = 0
                            _ = v27
                        }
                        _ = v25
                    }
                }
                _ = v10
            }
            _ = fi9
        }
        var v28 float32 = (1e10 + 0.0)
        for fi29 := 0; fi29 < 18; fi29++ {
            if (!(p1 || p1)) {
                p0 = p0
                v30 := GenS2{F0: p0, F1: "hello"}
                _ = v30
            } else {
                if len(v7) > 0 {
                    v7[0] = (^(fi8))
                }
            }
            _ = fi29
        }
        _ = fi8
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := []string{}
    v2 := make(map[int]int)
    var v3 int = -931
    var v4 int16 = 134
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v4)
    v4 += (v4 + v4)
    v31 := GenS1{F0: v3}
    fmt.Println(v3, v4)
    switch v3 {
    case 42:
        for fi32 := 0; fi32 < 14; fi32++ {
            v1 = append(v1, (((("foo bar" + "a") + "") + ("hello" + "\n\t")) + ("hello" + ("hello" + ""))))
            fmt.Println(fi32)
            _ = fi32
        }
    case 4294967295:
        var v33 uint = 1
        _ = v33
    case 32768:
        var v34 uint16 = genFn1(2147483648, false, 127)
        _ = v34
    }
    fmt.Println(v3, v4)
    v4 -= 1
    v3 += 9223372036854775807
    v35 := GenS3{F0: 1}
    var v36 int64 = 107
    v3 = (v3 / (((v3 | (v3 + v3))) | 1))
    fmt.Println(v4, v4, v4)
    switch "" {
    case "\n\t":
        var v37 string = "a"
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
}
