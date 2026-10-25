package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint8; F1 uint8; F2 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 uint; F2 int32 }
func genFn0(p0 int, p1 uint64) int8 {
    var v4 int8 = 42
    v4 -= v4
    _ = p0
    _ = p1
    _ = v4
    return v4
}
func genFn1() uint {
    var v5 float32 = (1e10 * 3.14159)
    v6 := GenS1{F0: 477, F1: 2147483647}
    fmt.Println(v5)
    v5 -= float32((0.0 + (0.0 + -2.25)))
    var v7 uint16 = 65535
    _ = v5
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 int16 = -426
    v2 := []uint32{65535}
    var v3 uint16 = 287
    v8 := GenS1{F0: int(v3), F1: 182}
    v1 = v1
    if false {
        if false {
            for fi9 := 0; fi9 < 6; fi9++ {
                for fi10 := 0; fi10 < 8; fi10++ {
                    fmt.Println(fi9, v1)
                    switch fi9 {
                    case 32767:
                        v1 -= v1
                    case 0:
                        var v11 uint32 = 2
                        _ = v11
                    case 127:
                        var v12 uint64 = 4294967295
                        _ = v12
                    default:
                        fmt.Println(fi9)
                    }
                    _ = fi10
                }
                if true {
                    v13 := GenS2{F0: 42}
                    var v14 uint16 = v3
                    _ = v13
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v1 -= (v1 / (((v1 >> ((59) & 63))) | 1))
                        _ = fi15
                    }
                    var v16 uint32 = 65535
                    v17 := GenS3{F0: 255, F1: true}
                    _ = v16
                    _ = v17
                }
                var v18 float64 = 3.14159
                _ = fi9
                _ = v18
            }
            var v19 float32 = 3.14159
            _ = v19
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            var v20 uint8 = 23
            _ = v20
        default:
            if false {
                v21 := GenS1{F0: -2147483648, F1: 247}
                var v22 uint64 = 4294967295
                if true {
                    v1 += v1
                    for fi23 := 0; fi23 < 20; fi23++ {
                        if (!(100 == 0)) {
                            for fi24 := 0; fi24 < 16; fi24++ {
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                                switch (^(fi24)) {
                                case -63:
                                    var v25 bool = (true || true)
                                    _ = v25
                                case 32768:
                                    var v26 uint64 = (v22 << ((233) & 63))
                                    _ = v26
                                case 9223372036854775807:
                                    var v27 int64 = 0
                                    _ = v27
                                }
                                _ = fi24
                            }
                            if ((!(9223372036854775807 >= 9223372036854775807)) || (!(((v22 > v22) && false) || false))) {
                                v1 = 0
                            } else {
                                v2 = append(v2, 1)
                                v22 += 18446744073709551615
                                for fi28 := 0; fi28 < 15; fi28++ {
                                    if (42 >= 638) {
                                        var v29 int32 = 1
                                        v30 := GenS4{F0: 206, F1: 1, F2: false}
                                        v31 := GenS5{F0: 255}
                                        var v32 int8 = 100
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    } else {
                                        v33 := GenS5{F0: 632}
                                        var v34 string = ""
                                        if len(v2) > 0 {
                                            v2[0] = 4294967295
                                        }
                                        _ = v33
                                        _ = v34
                                    }
                                    v35 := GenS6{F0: -1, F1: 18446744073709551615, F2: 42}
                                    _ = fi28
                                    _ = v35
                                }
                            }
                        } else {
                            v1 -= (v1 << ((2) & 63))
                            var v36 int16 = (-(v1))
                            _ = v36
                        }
                        v3 = v3
                        _ = fi23
                    }
                }
                _ = v21
                _ = v22
            }
        }
        var v37 int8 = genFn0(32767, 4294967295)
        _ = v37
    }
    if true {
        v1 -= 0
        v3 -= (v3 - 112)
        for fi38 := 0; fi38 < 3; fi38++ {
            v1 += v1
            for fi39 := 0; fi39 < 20; fi39++ {
                if true {
                    fmt.Println(v3)
                }
                v3 = 42
                _ = fi39
            }
            _ = fi38
        }
    }
    var v40 uint32 = 22
    var v41 float64 = (3.14159 - 1e10)
    v2 = append(v2, (v40 | (v40 * (^(v40)))))
    var v42 int8 = 1
    var v43 string = "\n\t"
    v42 += (v42 + (v42 | (v42 << ((42) & 63))))
    for fi44 := 0; fi44 < 19; fi44++ {
        fmt.Println(fi44, v41, v1)
        _ = fi44
    }
    var v45 int8 = 101
    v2 = append(v2, v40)
    for fi46 := 0; fi46 < 15; fi46++ {
        fmt.Println(v40, v1)
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v45
    _ = v8
}
