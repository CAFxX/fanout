package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint8; F2 int32 }
func genFn0(p0 uint64, p1 uint8) uint16 {
    v4 := GenS1{F0: false}
    if false {
        p1 -= 0
        var v5 uint64 = (^(p0))
        v5 += (v5 ^ (p0 % ((p0) | 1)))
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 823
}

func main() {
    var v1 int8 = 42
    var v2 float64 = ((0.1 + (1.5 * 0.1)) - ((0.0 + (1.5 + 1e10)) + 1.5))
    var v3 uint64 = 1
    v6 := GenS1{F0: (v3 > 0)}
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{}
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint64 = 781
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v1 += v1
    v3 -= 468
    v11 := GenS1{F0: (255 <= 18446744073709551615)}
    v12 := []int{281, -1}
    fmt.Println(v1, v2)
    fmt.Println(v2, v3)
    if (-1 >= -1) {
        var v13 bool = false
        _ = v13
    }
    if false {
        if (-1 > 2147483647) {
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 int32 = 42
                v2 += -2.25
                _ = fi14
                _ = v15
            }
        } else {
            for fi16 := 0; fi16 < 2; fi16++ {
                if false {
                    if len(v12) > 0 {
                        v12[0] = (-(fi16))
                    }
                }
                _ = fi16
            }
            v2 -= v2
            if len(v12) > 0 {
                v12[0] = 128
            }
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            for fi18 := 0; fi18 < 14; fi18++ {
                switch 0 {
                case -2147483648:
                    v3 -= (v3 % ((42) | 1))
                case -757:
                    var v19 int = (^(fi18))
                    _ = v19
                default:
                    v20 := GenS2{F0: 0}
                    _ = v20
                }
                v21 := GenS2{F0: 91}
                _ = len(v12)
                _ = fi18
                _ = v21
            }
            _ = fi17
        }
    }
    if (!(v1 >= 42)) {
        v22 := GenS3{F0: 32767, F1: 36, F2: 1}
        switch 42 {
        case 4294967295:
            if true {
                if ((59 == 4294967295) && (0 > 100)) {
                    if (false || true) {
                        var v23 float32 = 0.0
                        switch 0 {
                        case 4294967295:
                            var v24 uint = 42
                            _ = v24
                        }
                        _ = v23
                    }
                    switch "" {
                    case "foo bar":
                        if true {
                            var v25 int16 = 0
                            _ = v25
                        } else {
                            if false {
                                v26 := GenS3{F0: 0, F1: 255, F2: 2147483647}
                                if ((1 >= 819) && true) {
                                    for fi27 := 0; fi27 < 2; fi27++ {
                                        var v28 bool = (v3 >= (v3 + 1))
                                        v2 = v2
                                        v2 += (v2 * v2)
                                        _ = fi27
                                        _ = v28
                                    }
                                }
                                _ = v26
                            }
                        }
                    case "\n\t":
                        v1 = (v1 / ((v1) | 1))
                    case "":
                        fmt.Println(v1, v3)
                    }
                } else {
                    var v29 uint32 = 795
                    _ = v29
                }
            } else {
                var v30 float64 = 0.0
                if (0 < genFn0(v3, 128)) {
                    _ = len(v12)
                } else {
                    v31 := GenS3{F0: 0, F1: 238, F2: -502}
                    _ = v31
                }
                _ = v30
            }
        case 718:
            var v32 int32 = 42
            _ = v32
        }
        fmt.Println(v2)
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
}
