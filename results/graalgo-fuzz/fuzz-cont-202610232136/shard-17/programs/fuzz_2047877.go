package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64 }
func genFn0(p0 int, p1 int64, p2 int) int8 {
    fmt.Println(p2, p1)
    var v6 float32 = float32(0.1)
    switch "\n\t" {
    case "a":
        v7 := GenS1{F0: -359, F1: -1}
        _ = v7
    default:
        p2 += -155
    }
    if false {
        var v8 string = "\n\t"
        var v9 uint8 = 127
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := []int{p0, (fi10 >> ((0) & 63)), (fi10 * fi10), int(v9)}
            v6 = v6
            v12 := GenS1{F0: p1, F1: 96}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v8
        _ = v9
    }
    p0 -= (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 100
}

func main() {
    var v1 int = 42
    var v2 uint32 = 2147483647
    var v3 string = "foo bar"
    v4 := []int{(v1 & -9223372036854775808)}
    var v5 int16 = 143
    if (false || (1 > 701)) {
        var v13 int16 = v5
        for fi14 := 0; fi14 < 10; fi14++ {
            if false {
                var v15 uint8 = 0
                _ = v15
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                if false {
                    v5 = v13
                    if ("\n\t" >= "") {
                        switch (^(fi16)) {
                        case 0:
                            if true {
                                v17 := GenS1{F0: 7, F1: 1}
                                if false {
                                    var v18 uint = 2
                                    if true {
                                        v5 = v5
                                    } else {
                                        var v19 int32 = -1
                                        switch v3 {
                                        case "a":
                                            var v20 float64 = 1.5
                                            _ = v20
                                        default:
                                            v21 := GenS2{F0: (v19 & -665)}
                                            _ = v21
                                        }
                                        v22 := GenS3{F0: (fi16 >> ((v18) & 63)), F1: 32767}
                                        _ = v19
                                        _ = v22
                                    }
                                    var v23 uint32 = (v2 ^ v2)
                                    v24 := GenS4{F0: false}
                                    _ = v18
                                    _ = v23
                                    _ = v24
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = (fi16 * 32768)
                                    }
                                    if true {
                                        fmt.Println(fi16, v13)
                                    }
                                    v3 = "\n\t"
                                }
                                fmt.Println(fi14, v2, v3)
                                _ = v17
                            } else {
                                v2 = v2
                                if (!true) {
                                    var v25 bool = false
                                    var v26 int8 = 127
                                    v27 := GenS4{F0: v25}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                v1 = v1
                            }
                        case 267:
                            v28 := GenS4{F0: (0.0 <= 0.0)}
                            _ = v28
                        }
                        var v29 int8 = -128
                        _ = v29
                    } else {
                        v30 := GenS3{F0: fi14, F1: 1}
                        var v31 uint64 = 18446744073709551615
                        var v32 uint64 = 42
                        v33 := GenS5{F0: 255, F1: v2}
                        _ = v30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    if true {
                        var v34 uint = 798
                        v13 = (v5 | -32768)
                        _ = v34
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = -820
                    }
                }
                if true {
                    if true {
                        v1 = fi16
                    }
                    var v35 int = (v1 | (fi14 & 656))
                    _ = v35
                } else {
                    fmt.Println(fi14, fi16, fi14)
                    var v36 uint32 = v2
                    v2 += 2147483647
                    _ = v36
                }
                v3 = v3
                _ = fi16
            }
            v4 = append(v4, v1)
            _ = fi14
        }
        var v37 string = ("a" + v3)
        _ = v13
        _ = v37
    } else {
        v3 = (v3 + v3)
    }
    v38 := GenS6{F0: 42}
    v1 = (^(v1))
    var v39 int32 = 1
    var v40 int = -590
    v1 = v1
    if false {
        if false {
            var v41 uint8 = 128
            _ = v41
        } else {
            if false {
                v42 := GenS3{F0: v1, F1: 2}
                _ = v42
            }
            var v43 float32 = (float32(1e10) / (float32(1.5) * (-2.25 + 1.5)))
            switch v1 {
            case 128:
                var v44 uint = 151
                _ = v44
            case 695:
                var v45 uint = 2
                _ = v45
            }
            _ = v43
        }
        fmt.Println(v5, v1)
    }
    v46 := GenS7{F0: 1}
    fmt.Println(v39)
    _ = len(v4)
    v4 = append(v4, (v1 ^ (v1 & v40)))
    v4 = append(v4, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v46
    _ = v5
}
