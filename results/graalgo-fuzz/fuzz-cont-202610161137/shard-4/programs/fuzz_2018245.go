package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 int32, p1 uint32) uint {
    p1 = 42
    if true {
        p1 += p1
    }
    var v7 string = "hello"
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    var v1 uint8 = 203
    v2 := []int{4294967295}
    var v3 float64 = -2.25
    var v4 int16 = 18
    var v5 int32 = -2147483648
    var v6 float64 = v3
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint32 = 65535
        var v10 bool = (!false)
        var v11 int = 365
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 int32 = v5
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 float64 = v3
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    switch -119 {
    case -1:
        var v16 uint = 2
        _ = v16
    case -996:
        switch 32767 {
        case 9223372036854775807:
            var v17 uint8 = 124
            _ = v17
        case 128:
            v4 = (^(v4))
        default:
            v1 += 0
        }
    case 400:
        v18 := GenS1{F0: (v4 - 32767)}
        _ = v18
    }
    v19 := GenS1{F0: v4}
    if len(v2) > 0 {
        v2[0] = 128
    }
    var v20 uint16 = 67
    v21 := GenS1{F0: 706}
    if ((42 >= -507) && false) {
        fmt.Println(v6, v12, v6)
    }
    v2 = append(v2, 2147483647)
    if ((0.1 < float32(0.1)) || true) {
        fmt.Println(v12, v6, v5)
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        v5 += -1
        _ = fi22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            if len(v2) > 0 {
                v2[0] = 32767
            }
            if len(v2) > 0 {
                v2[0] = 943
            }
            _ = fi24
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v25 int16 = v4
        _ = fi23
        _ = v25
    }
    var v26 int16 = (-(v4))
    v3 = (v6 - v6)
    if false {
        if len(v2) > 0 {
            v2[0] = 1
        }
    } else {
        v27 := GenS2{F0: v1, F1: 1, F2: v5}
        if (!false) {
            if (!(genFn0((v12 / ((v5) | 1)), 969) != 4294967295)) {
                fmt.Println(v26, v20)
                v2 = append(v2, 2147483647)
                if (0 == v4) {
                    v28 := GenS2{F0: v1, F1: 2147483648, F2: (v5 % (((-(v5))) | 1))}
                    if (true && (false || true)) {
                        v29 := GenS2{F0: 40, F1: uint32(v20), F2: v5}
                        fmt.Println(v26, v3, v5)
                        _ = v29
                    }
                    _ = v28
                } else {
                    if true {
                        fmt.Println(v1)
                        v3 = 1.5
                        v26 = v4
                    } else {
                        var v30 int8 = 127
                        _ = v30
                    }
                    v12 = 42
                    if (!(!(true && false))) {
                        for fi31 := 0; fi31 < 19; fi31++ {
                            if len(v2) > 0 {
                                v2[0] = fi31
                            }
                            fmt.Println(v5, v1)
                            _ = fi31
                        }
                        v6 += (0.0 / (v3 - (0.0 * v6)))
                        v32 := GenS3{F0: (42 != (v20 - (^(v20)))), F1: (v12 - (v5 | v12))}
                        _ = v32
                    }
                }
            }
        }
        fmt.Println(v4, v5)
        _ = v27
    }
    var v33 int64 = 152
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
