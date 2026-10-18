package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []uint32{42, 42, 1, 0}
    var v2 uint8 = 2
    v3 := make(map[int]bool)
    var v4 bool = false
    var v5 uint16 = 914
    var v6 float64 = (1e10 / float64(0.1))
    if v4 {
        v2 -= (v2 << ((42) & 63))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!(1 >= 11)) {
        v3[128] = (4294967295 <= 375)
        if (!true) {
            for fi7 := 0; fi7 < 11; fi7++ {
                if v4 {
                    var v8 int64 = 42
                    var v9 uint64 = 42
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 float64 = v6
                        if ((-(v8)) > v8) {
                            if v4 {
                                switch "" {
                                case "":
                                    v9 = 431
                                    if (!v4) {
                                        if true {
                                            v12 := GenS1{F0: (1.5 > 1e10), F1: 632}
                                            v11 -= (float64(-2.25) / float64(0.1))
                                            v11 += 3.14159
                                            _ = v12
                                        }
                                        v13 := GenS1{F0: (false || (127 <= (v2 << ((510) & 63)))), F1: v5}
                                        _ = v13
                                    }
                                case "a":
                                    fmt.Println(v11)
                                case "hello":
                                    v11 = v11
                                }
                            } else {
                                v2 = v2
                                if true {
                                    fmt.Println(v9)
                                } else {
                                    v14 := v3[-322]
                                    fmt.Println(v11, v2, v6)
                                    var v15 int8 = -128
                                    var v16 float32 = 3.14159
                                    switch ("a" + "") {
                                    case "a":
                                        switch fi7 {
                                        case 4294967295:
                                            var v17 uint16 = (v5 % (((v5 + uint16(v9))) | 1))
                                            _ = v17
                                        }
                                    }
                                    fmt.Println(v14, v5, v15)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (v5 != v5) {
                                        v16 = v16
                                        v18 := GenS1{F0: true, F1: v5}
                                        v19 := GenS2{F0: v14}
                                        fmt.Println(v5)
                                        if (v8 <= (-(v8))) {
                                            if v4 {
                                                v20 := GenS2{F0: (v16 < v16)}
                                                v1 = append(v1, 2)
                                                v21 := GenS1{F0: (42 <= 430), F1: v5}
                                                _ = v20
                                                _ = v21
                                            }
                                            v11 = (v11 - v6)
                                        } else {
                                            if (v14 || (391 >= 2)) {
                                                var v22 uint32 = 0
                                                v5 -= (v5 ^ 42)
                                                _ = v22
                                            }
                                        }
                                        _ = v18
                                        _ = v19
                                    } else {
                                        var v23 int16 = 0
                                        v24 := GenS3{F0: 0, F1: 1}
                                        switch 0 {
                                        case 443:
                                            if v14 {
                                                v25 := GenS3{F0: 1, F1: 42}
                                                _ = v25
                                            } else {
                                                if v14 {
                                                    var v26 int8 = (v15 >> ((113) & 63))
                                                    _ = v26
                                                }
                                            }
                                        case 4294967295:
                                            var v27 int = fi10
                                            _ = v27
                                        case -191:
                                            v28 := GenS4{F0: (fi10 - fi7)}
                                            _ = v28
                                        }
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                                fmt.Println(v4, v8, v9)
                            }
                            fmt.Println(v4, fi7)
                            v29 := GenS2{F0: (164 == 4294967295)}
                            _ = v29
                        }
                        _ = fi10
                        _ = v11
                    }
                    _ = v8
                    _ = v9
                }
                var v30 float32 = 3.14159
                fmt.Println(v6)
                _ = fi7
                _ = v30
            }
            v31 := GenS3{F0: -1, F1: 2}
            if (258 == 18446744073709551615) {
                var v32 uint64 = 463
                _ = v32
            } else {
                if len(v1) > 0 {
                    v1[0] = 799
                }
                var v33 uint32 = 0
                _ = v33
            }
            _ = v31
        } else {
            v6 += v6
            v2 += (v2 | (v2 - (v2 >> ((864) & 63))))
        }
    }
    v2 = 255
    v34 := GenS5{F0: int64(v5)}
    v35 := GenS5{F0: -1}
    fmt.Println(v2, v5, v2)
    v36 := GenS6{F0: 2, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
