package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }
func genFn0(p0 int) int8 {
    p0 = p0
    switch (p0 - p0) {
    case 2147483647:
        p0 += p0
    case 873:
        v7 := []int{(p0 >> ((18446744073709551615) & 63)), p0, p0}
        _ = v7
    default:
        v8 := []int{4294967295, 32767}
        _ = v8
    }
    if true {
        if (!false) {
            p0 = p0
            for fi9 := 0; fi9 < 18; fi9++ {
                if true {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        fmt.Println(fi9, fi9, p0)
                        _ = fi10
                    }
                } else {
                    var v11 uint8 = 42
                    fmt.Println(fi9)
                    _ = v11
                }
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 float32 = 1e10
                    _ = fi12
                    _ = v13
                }
                _ = fi9
            }
        } else {
            fmt.Println(p0)
            var v14 int32 = -2147483648
            fmt.Println(v14, p0)
            _ = v14
        }
        v15 := GenS1{F0: ("a" != "\n\t"), F1: 1, F2: 856}
        _ = v15
    } else {
        v16 := GenS1{F0: true, F1: 203, F2: 559}
        _ = v16
    }
    _ = p0
    return 42
}

func main() {
    var v1 int = -2147483648
    v2 := []uint32{}
    var v3 uint16 = 32767
    var v4 int16 = 1000
    v5 := make(map[uint32]bool)
    var v6 uint = 878
    v5[65535] = ((v3 >= v3) || (0 > 29))
    switch (v1 - (v1 | v1)) {
    case 32767:
        if false {
            for fi17 := 0; fi17 < 14; fi17++ {
                fmt.Println(v1)
                v4 = v4
                if (3.14159 <= 3.14159) {
                    v18 := v5[2147483647]
                    switch (fi17 / ((v1) | 1)) {
                    case 4294967295:
                        v19 := GenS1{F0: true, F1: 42, F2: 4294967295}
                        _ = v19
                    case 0:
                        for fi20 := 0; fi20 < 17; fi20++ {
                            if v18 {
                                if v18 {
                                    v6 = 1
                                    v3 = (^(v3))
                                } else {
                                    v18 = true
                                }
                            }
                            fmt.Println(fi17, v1)
                            var v21 uint = v6
                            _ = fi20
                            _ = v21
                        }
                    case 574:
                        v5[2] = (0.0 <= 1e10)
                    default:
                        v4 += -1
                    }
                    _ = v18
                }
                _ = fi17
            }
            var v22 uint16 = v3
            v3 = v3
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 8; fi23++ {
            v4 = (v4 << ((v6) & 63))
            _ = fi23
        }
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(v4, v3, v4)
        v25 := GenS1{F0: true, F1: 18446744073709551615, F2: 942}
        _ = fi24
        _ = v25
    }
    v26 := v5[210]
    var v27 int8 = 1
    v28 := GenS1{F0: (v26 && (v26 && (v4 < -32768))), F1: 1, F2: 2147483648}
    delete(v5, 2)
    v3 = v3
    for fi29 := 0; fi29 < 13; fi29++ {
        fmt.Println(v1, v26, v26)
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
