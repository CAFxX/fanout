package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint8 }
func genFn0(p0 uint16, p1 uint, p2 int16) int8 {
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint32 = 989
        v7 := GenS1{F0: p0}
        v6 += 727
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int16 = 1
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 4294967295, F1: 42, F2: 2}
    p2 -= p2
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -60
}

func main() {
    var v1 int8 = 42
    var v2 int32 = -433
    var v3 int16 = 690
    v4 := []int{}
    switch "" {
    case "a":
        v1 = v1
    case "\n\t":
        if true {
            for fi12 := 0; fi12 < 16; fi12++ {
                if (-18 < v1) {
                    v13 := GenS2{F0: 193, F1: 4294967295, F2: 255}
                    _ = len(v4)
                    _ = v13
                } else {
                    if (false && (!(42 != 416))) {
                        v2 -= 377
                    } else {
                        var v14 bool = true
                        if (-32768 >= v3) {
                            for fi15 := 0; fi15 < 14; fi15++ {
                                var v16 int32 = (v2 / (((v2 << ((65535) & 63))) | 1))
                                _ = fi15
                                _ = v16
                            }
                            fmt.Println(v14, v2, v3)
                        } else {
                            switch 9223372036854775807 {
                            case 448:
                                v3 -= 42
                            case 2147483647:
                                var v17 uint64 = 2
                                _ = v17
                            default:
                                for fi18 := 0; fi18 < 17; fi18++ {
                                    var v19 uint = 255
                                    v19 = v19
                                    if true {
                                        var v20 uint32 = 42
                                        _ = v20
                                    }
                                    v14 = (!(!v14))
                                    _ = fi18
                                    _ = v19
                                }
                            }
                        }
                        _ = v14
                    }
                    for fi21 := 0; fi21 < 12; fi21++ {
                        var v22 int32 = (v2 ^ (v2 % (((v2 & (v2 / ((v2) | 1)))) | 1)))
                        _ = fi21
                        _ = v22
                    }
                }
                v1 += v1
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 float64 = 3.14159
                    var v25 uint = 4294967295
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi12
            }
            switch -2147483648 {
            case 4294967295:
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 bool = false
                    _ = fi26
                    _ = v27
                }
            case 127:
                for fi28 := 0; fi28 < 12; fi28++ {
                    switch 42 {
                    case 42:
                        if true {
                            v4 = append(v4, 32768)
                        }
                    }
                    _ = fi28
                }
            }
            v3 = (^(v3))
        }
    case "":
        if true {
            var v29 uint = 42
            _ = v29
        }
    }
    v1 -= 42
    if true {
        v30 := GenS1{F0: 1}
        _ = v30
    }
    var v31 int = 127
    for fi32 := 0; fi32 < 7; fi32++ {
        if false {
            var v33 uint16 = 32767
            var v34 float32 = 1.5
            _ = v33
            _ = v34
        }
        _ = fi32
    }
    if len(v4) > 0 {
        v4[0] = (v31 | 1)
    }
    var v35 uint = 348
    v36 := GenS1{F0: 2}
    var v37 uint32 = 2147483648
    for fi38 := 0; fi38 < 9; fi38++ {
        var v39 uint64 = 4294967295
        _ = fi38
        _ = v39
    }
    v35 = 255
    var v40 int32 = v2
    var v41 int32 = 42
    var v42 uint32 = 42
    if len(v4) > 0 {
        v4[0] = v31
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
