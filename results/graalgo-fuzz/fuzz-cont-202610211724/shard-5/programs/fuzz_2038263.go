package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 string; F1 uint8; F2 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 string; F2 int64 }
func genFn0(p0 int64, p1 string) int8 {
    p1 = "a"
    p1 = ("hello" + (((p1 + p1) + p1) + ""))
    v6 := GenS1{F0: 189, F1: 0}
    var v7 string = "\n\t"
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 100
}

func main() {
    var v1 bool = false
    var v2 uint = 0
    var v3 uint = v2
    var v4 uint8 = 153
    v5 := make(map[uint32]uint32)
    v8 := []int{139}
    v9 := GenS1{F0: 255, F1: -9223372036854775808}
    var v10 int64 = 42
    switch "foo bar" {
    case "a":
        var v11 uint8 = v4
        _ = v11
    case "":
        var v12 uint8 = 1
        _ = v12
    }
    v13 := GenS1{F0: 0, F1: 356}
    switch "" {
    case "hello":
        if v1 {
            v14 := v5[2147483648]
            _ = v14
        } else {
            v15 := GenS2{F0: -2147483648, F1: 4294967295}
            _ = v15
        }
    default:
        var v16 int16 = 1
        _ = v16
    }
    fmt.Println(v10)
    fmt.Println(v4, v3)
    v10 = v10
    if v1 {
        var v17 bool = v1
        _ = v17
    } else {
        switch "hello" {
        case "a":
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 int64 = (^(v10))
                _ = fi18
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                switch "" {
                case "hello":
                    fmt.Println(v2)
                case "":
                    switch "a" {
                    case "foo bar":
                        var v22 uint16 = 255
                        _ = v22
                    default:
                        if v1 {
                            if v1 {
                                v23 := GenS1{F0: 2, F1: (^(fi21))}
                                _ = v23
                            } else {
                                v3 = 2
                                var v24 float64 = 0.0
                                _ = v24
                            }
                        }
                    }
                case "a":
                    v3 += (v2 / (((v2 >> ((4294967295) & 63))) | 1))
                    switch fi20 {
                    case 42:
                        if v1 {
                            v8 = append(v8, fi20)
                            switch (fi20 << (((v2 << ((57) & 63))) & 63)) {
                            case 346:
                                v25 := GenS3{F0: 42, F1: -32768}
                                if len(v8) > 0 {
                                    v8[0] = fi20
                                }
                                _ = v25
                            default:
                                v26 := GenS1{F0: 2, F1: -836}
                                _ = v26
                            }
                        } else {
                            _ = len(v8)
                        }
                    }
                    var v27 uint8 = (^(v4))
                    _ = v27
                default:
                    switch 406 {
                    case 1:
                        var v28 uint8 = (v4 + (v4 | (v4 >> ((42) & 63))))
                        _ = v28
                    case -9223372036854775808:
                        v10 += (v10 + v10)
                        switch fi20 {
                        case -2147483648:
                            v1 = v1
                        case 127:
                            if (547 < 65535) {
                                switch 32768 {
                                case -9223372036854775808:
                                    fmt.Println(v4, v4)
                                case 4294967295:
                                    var v29 float64 = 0.1
                                    _ = v29
                                }
                                v30 := GenS4{F0: "hello", F1: 0, F2: 2}
                                var v31 int64 = v10
                                _ = v30
                                _ = v31
                            }
                        }
                        switch "foo bar" {
                        case "":
                            v10 = v10
                        case "foo bar":
                            delete(v5, 509)
                        }
                        switch "" {
                        case "a":
                            fmt.Println(fi20)
                        case "":
                            v4 = v4
                        }
                    case 112:
                        v2 = v2
                    }
                }
                v10 += v10
                _ = fi21
            }
            var v32 uint8 = (v4 >> (((v2 | 18446744073709551615)) & 63))
            var v33 float32 = 1e10
            _ = fi20
            _ = v32
            _ = v33
        }
    }
    v3 = v2
    v2 = 850
    v34 := GenS5{F0: v10}
    if v1 {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        var v35 int32 = 1
        _ = v35
    }
    v2 += uint(v4)
    if (!((!(false || (1 != 96))) && (1e10 < 0.0))) {
        v36 := GenS6{F0: true, F1: "\n\t", F2: v10}
        _ = v36
    } else {
        var v37 float64 = 1e10
        v38 := GenS5{F0: (-(v10))}
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
