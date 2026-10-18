package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 1e10
    var v3 uint = 42
    var v4 int8 = -128
    v5 := make(map[int]bool)
    var v6 string = ""
    var v7 int = -869
    v2 -= v2
    fmt.Println(v4, v4)
    var v8 float64 = -2.25
    v4 -= v4
    fmt.Println(v7, v7, v4)
    for fi9 := 0; fi9 < 3; fi9++ {
        v4 = (v4 | v4)
        v2 += float32(float64(float32(0.1)))
        _ = fi9
    }
    switch v6 {
    case "a":
        if (v8 >= ((v8 / 1e10) / (((v8 * v8) + v8) - v8))) {
            if true {
                delete(v1, uint32(v3))
                var v10 int8 = (v4 % (((v4 % (((v4 >> (((v3 ^ (v3 + v3))) & 63))) | 1))) | 1))
                _ = v10
            }
            v11 := GenS1{F0: 53}
            if (true || true) {
                var v12 uint32 = 2
                var v13 int8 = (v4 - v4)
                _ = v12
                _ = v13
            } else {
                if (true || false) {
                    switch (("a" + v6) + v6) {
                    case "foo bar":
                        fmt.Println(v4, v2, v4)
                    case "":
                        switch (v7 << (((^(v3))) & 63)) {
                        case 32767:
                            var v14 int8 = 1
                            _ = v14
                        default:
                            for fi15 := 0; fi15 < 12; fi15++ {
                                if false {
                                    var v16 uint = (^(v3))
                                    v6 = v6
                                    if true {
                                        if true {
                                            v17 := []int{}
                                            _ = v17
                                        }
                                    } else {
                                        switch int(v4) {
                                        case -256:
                                            v18 := []int{(v7 << ((42) & 63))}
                                            _ = v18
                                        default:
                                            v6 = "hello"
                                        }
                                        v8 -= v8
                                        _ = len(v1)
                                    }
                                    _ = v16
                                }
                                _ = fi15
                            }
                        }
                    }
                } else {
                    var v19 int = (v7 << ((255) & 63))
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v21 := GenS2{F0: -2147483648}
                        _ = fi20
                        _ = v21
                    }
                    switch (v19 | -2147483648) {
                    case 32767:
                        var v22 string = "hello"
                        _ = v22
                    }
                    _ = v19
                }
                v23 := GenS1{F0: 474}
                _ = v23
            }
            _ = v11
        } else {
            fmt.Println(v4)
        }
    case "hello":
        for fi24 := 0; fi24 < 16; fi24++ {
            v3 -= v3
            var v25 int16 = -131
            _ = fi24
            _ = v25
        }
    case "\n\t":
        switch v7 {
        case 72:
            v26 := GenS3{F0: -898, F1: -899, F2: 469}
            _ = v26
        case 128:
            fmt.Println(v4)
        default:
            if (2 <= 0) {
                if false {
                    fmt.Println(v6, v6, v4)
                    v27 := GenS3{F0: 2147483647, F1: 332, F2: v3}
                    _ = v27
                }
                var v28 int = v7
                _ = v28
            } else {
                var v29 int = v7
                v30 := []int{}
                _ = v29
                _ = v30
            }
        }
    default:
        delete(v1, 630)
    }
    v31 := []int{v7, (^(v7)), (v7 >> ((v3) & 63)), (v7 - v7)}
    for fi32 := 0; fi32 < 17; fi32++ {
        if (571 > 622) {
            v6 = v6
        } else {
            if false {
                v31 = append(v31, v7)
            } else {
                v6 = v6
                v33 := GenS3{F0: 0, F1: 0, F2: (v3 * (v3 >> ((v3) & 63)))}
                v31 = append(v31, -756)
                _ = v33
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                v7 = (-(fi32))
                fmt.Println(fi34, v6, fi34)
                switch -867 {
                case 128:
                    var v35 uint64 = 1
                    _ = v35
                default:
                    _ = len(v1)
                }
                _ = fi34
            }
            var v36 int32 = -2147483648
            _ = v36
        }
        _ = fi32
    }
    var v37 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
