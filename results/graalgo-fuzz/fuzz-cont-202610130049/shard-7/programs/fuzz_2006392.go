package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint }
func genFn0(p0 int64) bool {
    p0 = (p0 + p0)
    var v7 uint64 = 199
    var v8 uint8 = 1
    _ = p0
    _ = v7
    _ = v8
    return ((false && false) || (!false))
}
func genFn1() bool {
    var v9 int32 = 42
    if true {
        switch 127 {
        case 698:
            fmt.Println(v9)
        case 0:
            v9 = v9
        case 1:
            switch -1 {
            case 9223372036854775807:
                for fi10 := 0; fi10 < 13; fi10++ {
                    fmt.Println(v9)
                    _ = fi10
                }
            case 127:
                if false {
                    if true {
                        var v11 string = ("" + ("" + ("\n\t" + "hello")))
                        v9 = (v9 / ((v9) | 1))
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 uint16 = 2
                            v13 -= (v13 + (v13 - v13))
                            v14 := []int{fi12, 248, fi12, fi12}
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = v11
                    } else {
                        fmt.Println(v9)
                        v15 := []int{4294967295}
                        v15 = append(v15, 210)
                        _ = v15
                    }
                    if (-2.25 != (3.14159 / (1.5 / 1.5))) {
                        switch -9223372036854775808 {
                        case 42:
                            fmt.Println(v9)
                        case -43:
                            v9 = (v9 + (v9 >> ((1) & 63)))
                        case 843:
                            v9 += -296
                        default:
                            var v16 uint = 2
                            _ = v16
                        }
                    } else {
                        if true {
                            var v17 uint8 = 120
                            v18 := GenS1{F0: 309, F1: 1000, F2: 54}
                            var v19 uint16 = 42
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        if true {
                            fmt.Println(v9)
                        }
                        var v20 int = 4294967295
                        _ = v20
                    }
                    fmt.Println(v9)
                }
            default:
                v9 -= (v9 % ((v9) | 1))
            }
        default:
            switch 31 {
            case 1:
                var v21 int8 = 127
                _ = v21
            }
        }
    }
    var v22 int64 = 42
    _ = v22
    _ = v9
    return (!false)
}

func main() {
    v1 := []uint32{2147483648, 42, 2147483647, 117}
    var v2 uint8 = 128
    var v3 uint = 1
    v4 := []string{}
    var v5 uint64 = 2
    var v6 uint32 = 65535
    var v23 int64 = 9223372036854775807
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 += 255
    v3 -= v3
    v24 := GenS1{F0: (v3 << (((v3 | (v3 / ((v3) | 1)))) & 63)), F1: 1000, F2: (^(v3))}
    for fi25 := 0; fi25 < 5; fi25++ {
        v1 = append(v1, uint32(v5))
        if true {
            v3 += v3
        }
        var v26 uint8 = (v2 << ((18446744073709551615) & 63))
        _ = fi25
        _ = v26
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v27 int32 = 2147483647
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS1{F0: (v3 / ((v3) | 1)), F1: 246, F2: (^(v3))}
        for fi30 := 0; fi30 < 18; fi30++ {
            v2 = 127
            if false {
                v31 := GenS1{F0: (v3 | (v3 >> (((v3 | v3)) & 63))), F1: 1000, F2: 255}
                var v32 int32 = v27
                _ = v31
                _ = v32
            }
            v27 -= -925
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    var v33 int16 = 0
    switch 0 {
    case 128:
        var v34 int = -338
        _ = v34
    case -1:
        v3 += v3
    default:
        var v35 int16 = -32768
        _ = v35
    }
    v36 := GenS1{F0: (v3 % ((18446744073709551615) | 1)), F1: (-(v33)), F2: 384}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
