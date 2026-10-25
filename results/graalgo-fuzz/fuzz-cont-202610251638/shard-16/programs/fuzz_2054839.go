package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint; F1 uint; F2 uint32 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint8, p1 int32) uint {
    p1 += p1
    p1 = (^(p1))
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 int16) int8 {
    if true {
        p0 += p0
        v5 := []int{}
        for fi6 := 0; fi6 < 6; fi6++ {
            if (0 > 0) {
                if true {
                    fmt.Println(p0)
                } else {
                    for fi7 := 0; fi7 < 9; fi7++ {
                        v8 := GenS1{F0: 1, F1: (2147483647 < -192), F2: -128}
                        _ = len(v5)
                        switch fi7 {
                        case 32767:
                            switch fi6 {
                            case -136:
                                var v9 uint32 = 96
                                _ = v9
                            case -2147483648:
                                fmt.Println(p0, fi7)
                            }
                        case 9223372036854775807:
                            if (true || (!true)) {
                                v10 := GenS1{F0: -1, F1: (true || ((^(fi7)) >= 0)), F2: 0}
                                _ = v10
                            } else {
                                fmt.Println(p0, fi6, p0)
                                switch (fi7 >> ((0) & 63)) {
                                case 32768:
                                    var v11 int32 = -382
                                    _ = v11
                                case 1:
                                    var v12 uint8 = 243
                                    _ = v12
                                case 302:
                                    p0 -= -818
                                }
                            }
                            var v13 bool = (-32768 <= p0)
                            _ = v13
                        }
                        _ = fi7
                        _ = v8
                    }
                    var v14 uint = 2
                    _ = v14
                }
                p0 = (^(p0))
            }
            _ = fi6
        }
        _ = v5
    } else {
        v15 := []int{42, 791, 2147483647}
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        fmt.Println(p0)
        _ = v15
    }
    fmt.Println(p0)
    v16 := GenS2{F0: -1, F1: false}
    var v17 int16 = (p0 ^ (p0 | 32767))
    var v18 float32 = float32(-2.25)
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return -36
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -1
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    var v19 string = "foo bar"
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            v2 = -78
            _ = len(v4)
            if true {
                v19 = v19
                switch v19 {
                case "hello":
                    v19 = v19
                case "":
                    v22 := []int{}
                    _ = v22
                default:
                    var v23 int = 1
                    _ = v23
                }
                v2 = (^(v2))
            }
            _ = fi21
        }
        v24 := GenS3{F0: 255, F1: 1, F2: 2147483647}
        _ = fi20
        _ = v24
    }
    var v25 float64 = 0.0
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint64 = 97
        switch "foo bar" {
        case "\n\t":
            if true {
                v19 = ("hello" + v19)
                var v28 string = v19
                fmt.Println(fi26, v25, fi26)
                _ = v28
            } else {
                fmt.Println(v2, v25, v25)
                v27 += (v27 % (((v27 * v27)) | 1))
            }
        case "hello":
            v29 := GenS4{F0: fi26, F1: 127}
            _ = v29
        case "a":
            delete(v1, -9223372036854775808)
        default:
            if (true && false) {
                var v30 bool = true
                if false {
                    var v31 uint16 = 31
                    switch fi26 {
                    case 128:
                        var v32 uint64 = 4294967295
                        _ = v32
                    case 42:
                        v33 := GenS5{F0: 32767}
                        _ = v33
                    case 32767:
                        fmt.Println(v2, v31, v19)
                    default:
                        for fi34 := 0; fi34 < 3; fi34++ {
                            v4[fi26] = (v30 && v30)
                            var v35 float32 = 0.1
                            v27 += 18446744073709551615
                            _ = fi34
                            _ = v35
                        }
                    }
                    _ = v31
                } else {
                    v19 = v19
                }
                for fi36 := 0; fi36 < 14; fi36++ {
                    if (v19 <= "") {
                        v27 = 42
                        var v37 uint8 = 255
                        _ = v37
                    }
                    fmt.Println(fi36, v25, v25)
                    fmt.Println(v19, v19, v30)
                    _ = fi36
                }
                _ = v30
            }
        }
        v38 := GenS6{F0: v2}
        _ = fi26
        _ = v27
        _ = v38
    }
    v25 -= ((((v25 + (v25 + (v25 - (v25 + v25)))) + v25) - v25) * v25)
    fmt.Println(v25)
    v19 = ""
    fmt.Println(v19)
    v39 := GenS6{F0: v2}
    var v40 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
