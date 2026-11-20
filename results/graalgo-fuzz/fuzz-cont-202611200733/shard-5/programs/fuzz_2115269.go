package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 int8, p1 int32) uint16 {
    p1 = p1
    p0 -= p0
    var v7 bool = true
    if false {
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 bool = true
    var v2 int8 = 0
    var v3 uint64 = 1
    v4 := []int{}
    v5 := []uint32{0, 42, 983}
    var v6 uint8 = 255
    v8 := GenS1{F0: 1}
    v9 := GenS1{F0: 342}
    fmt.Println(v6, v6)
    var v10 int64 = 284
    v10 -= (^(v10))
    for fi11 := 0; fi11 < 19; fi11++ {
        v1 = ((567 >= 1) || v1)
        if ((-2.25 - (1e10 / 1.0)) >= float64(float32((1e10 - (0.1 / 1e10))))) {
            v12 := GenS1{F0: 0}
            var v13 bool = v1
            v10 = 714
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        if (834 > 65535) {
            v3 = v3
            var v15 uint = 175
            if true {
                switch (fi14 - fi14) {
                case 1:
                    v6 -= (^(v6))
                case 4294967295:
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v17 := GenS1{F0: v15}
                        if (v6 < v6) {
                            v2 = v2
                        }
                        v18 := GenS1{F0: (v15 + v15)}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                case 42:
                    if false {
                        switch (-(fi14)) {
                        case 0:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case 127:
                            fmt.Println(v3, v6)
                        default:
                            for fi19 := 0; fi19 < 17; fi19++ {
                                var v20 int32 = 1
                                switch "hello" {
                                case "foo bar":
                                    if v1 {
                                        var v21 uint64 = 2
                                        if ((fi19 >= fi14) && (0 == (fi14 % ((int(v6)) | 1)))) {
                                            var v22 int64 = (^(v10))
                                            _ = v22
                                        }
                                        fmt.Println(fi14, fi19)
                                        _ = v21
                                    }
                                case "hello":
                                    v23 := GenS1{F0: 1}
                                    _ = v23
                                }
                                if (41 < 2) {
                                    fmt.Println(v10, v1, v6)
                                }
                                if false {
                                    switch "\n\t" {
                                    case "hello":
                                        var v24 uint = (v15 * (v15 & (v15 | v15)))
                                        _ = v24
                                    case "a":
                                        fmt.Println(v10)
                                    case "\n\t":
                                        v25 := GenS2{F0: 2, F1: 158}
                                        _ = v25
                                    }
                                }
                                v26 := GenS2{F0: 65535, F1: 0}
                                _ = fi19
                                _ = v20
                                _ = v26
                            }
                        }
                    } else {
                        v6 = v6
                        var v27 int8 = 0
                        v28 := GenS3{F0: (^(v6)), F1: 1}
                        _ = v27
                        _ = v28
                    }
                default:
                    v10 += -1
                }
            }
            _ = v15
        }
        var v29 float64 = 0.1
        _ = fi14
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS2{F0: 0, F1: 400}
        _ = fi30
        _ = v31
    }
    var v32 bool = (v1 && v1)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
