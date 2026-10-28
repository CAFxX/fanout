package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint }
func genFn0(p0 string, p1 int64) uint32 {
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(fi6)
            var v8 int = fi6
            v9 := GenS1{F0: -1, F1: 42, F2: 128}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    if (1 >= -32768) {
        p1 -= p1
        for fi10 := 0; fi10 < 4; fi10++ {
            switch fi10 {
            case -212:
                var v11 float32 = float32(3.14159)
                _ = v11
            case -9223372036854775808:
                if (p0 == ("" + p0)) {
                    p1 += 875
                } else {
                    p0 = "\n\t"
                }
            }
            var v12 string = p0
            _ = fi10
            _ = v12
        }
    } else {
        var v13 float32 = 1e10
        _ = v13
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := []string{"a", ("\n\t" + "hello"), "\n\t"}
    var v2 uint16 = 1
    var v3 int = 32768
    var v4 int8 = -1
    v5 := []int{}
    v2 -= (v2 / (((v2 << ((42) & 63))) | 1))
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS1{F0: -39, F1: -1, F2: 2}
        v16 := GenS1{F0: 42, F1: v4, F2: 1}
        v4 -= -34
        _ = fi14
        _ = v15
        _ = v16
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v4 -= 72
        fmt.Println(v4, v2, fi17)
        _ = len(v5)
        _ = fi17
    }
    v2 = 255
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 bool = (v2 <= (v2 / (((^(v2))) | 1)))
        var v20 bool = v19
        for fi21 := 0; fi21 < 15; fi21++ {
            switch -2147483648 {
            case 2147483647:
                var v22 uint32 = 4294967295
                _ = v22
            case -554:
                var v23 int8 = (v4 / ((v4) | 1))
                v23 += -81
                v2 += v2
                _ = v23
            }
            _ = fi21
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    var v24 bool = (1 == 0)
    for fi25 := 0; fi25 < 12; fi25++ {
        v2 += v2
        _ = fi25
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        _ = len(v5)
        switch 141 {
        case 1:
            var v27 bool = (false || (v24 && v24))
            _ = v27
        case -504:
            v28 := GenS1{F0: -502, F1: v4, F2: 42}
            _ = v28
        case -9223372036854775808:
            v24 = true
        default:
            var v29 uint32 = 4294967295
            _ = v29
        }
        _ = fi26
    }
    v5 = append(v5, (^(v3)))
    v30 := GenS1{F0: 1, F1: v4, F2: 1}
    var v31 uint16 = v2
    switch 42 {
    case 32768:
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if v24 {
                _ = len(v1)
            }
            for fi32 := 0; fi32 < 4; fi32++ {
                v33 := GenS1{F0: 535, F1: int8(v31), F2: 255}
                if true {
                    if (-563 >= -1) {
                        for fi34 := 0; fi34 < 19; fi34++ {
                            _ = len(v1)
                            if len(v5) > 0 {
                                v5[0] = fi32
                            }
                            _ = fi34
                        }
                    } else {
                        if (false && (2 <= 42)) {
                            fmt.Println(v3)
                        } else {
                            v2 = v2
                            fmt.Println(fi32)
                        }
                        for fi35 := 0; fi35 < 9; fi35++ {
                            v4 = v4
                            v4 = 100
                            _ = fi35
                        }
                    }
                    v2 = (v31 | v31)
                    if (4294967295 == 42) {
                        var v36 int = (v3 ^ v3)
                        if (42 == 99) {
                            v36 += 2147483647
                            v37 := GenS2{F0: 255}
                            fmt.Println(v24)
                            _ = v37
                        }
                        _ = v36
                    }
                } else {
                    _ = len(v5)
                }
                var v38 uint32 = 913
                _ = fi32
                _ = v33
                _ = v38
            }
        }
    }
    v31 += v31
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
