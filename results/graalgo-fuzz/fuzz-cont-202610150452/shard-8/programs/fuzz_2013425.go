package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint) uint32 {
    v7 := GenS1{F0: (153 != 128), F1: 184, F2: 255}
    var v8 int16 = 32767
    v8 -= (^(v8))
    _ = p0
    _ = v7
    _ = v8
    return 289
}

func main() {
    var v1 float32 = (1.5 / (3.14159 * ((0.0 - ((((0.1 / (-2.25 * -2.25)) * 3.14159) + ((0.0 + 1e10) + (0.0 + (1e10 / 0.1)))) + 1.5)) - -2.25)))
    var v2 int = 4294967295
    var v3 float32 = 0.0
    var v4 bool = (true && (((v1 / v3) / v3) >= (3.14159 - (v1 / 0.1))))
    v5 := []uint32{2147483648, 42}
    v6 := make(map[string]int)
    if (0 > 2147483648) {
        switch (v2 / ((476) | 1)) {
        case -375:
            var v9 int16 = 1
            _ = v9
        default:
            v4 = (!(42 != 0))
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := GenS1{F0: v4, F1: 2, F2: 127}
                if (4294967295 < 1) {
                    v4 = v4
                } else {
                    var v13 uint32 = genFn0(2)
                    _ = v13
                }
                var v14 uint16 = 1
                _ = fi11
                _ = v12
                _ = v14
            }
            v15 := GenS1{F0: (!(!v4)), F1: 128, F2: 128}
            _ = fi10
            _ = v15
        }
    } else {
        v4 = v4
    }
    v2 += (v2 * (^(v2)))
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            switch 128 {
            case -639:
                if v4 {
                    _ = len(v5)
                    if v4 {
                        if ((42 < 920) && (926 >= 390)) {
                            switch (fi17 << ((0) & 63)) {
                            case -2147483648:
                                fmt.Println(fi17, v3)
                            case 4294967295:
                                var v18 int32 = -274
                                _ = v18
                            case 128:
                                switch "foo bar" {
                                case "\n\t":
                                    v19 := GenS1{F0: false, F1: 127, F2: 170}
                                    _ = v19
                                default:
                                    v2 -= fi17
                                }
                            }
                            var v20 int = fi16
                            _ = v20
                        } else {
                            v21 := GenS2{F0: 159, F1: 1000, F2: 2}
                            fmt.Println(v3)
                            if false {
                                v22 := GenS3{F0: 32767}
                                var v23 bool = (959 >= 2147483647)
                                _ = v22
                                _ = v23
                            }
                            _ = v21
                        }
                        var v24 int64 = -1
                        var v25 int8 = -1
                        _ = v24
                        _ = v25
                    } else {
                        fmt.Println(v4)
                        switch (^(fi17)) {
                        case 4294967295:
                            v26 := GenS4{F0: 0}
                            _ = v26
                        }
                        v1 -= (1e10 + 0.1)
                    }
                    v2 = (fi16 % ((fi17) | 1))
                }
            case -2147483648:
                delete(v6, "foo bar")
            default:
                v4 = (!(true || ((1 < 770) || (42 == 32767))))
                v2 = v2
                fmt.Println(v2)
            }
            fmt.Println(v4, fi16, fi17)
            _ = fi17
        }
        var v27 int32 = 446
        _ = fi16
        _ = v27
    }
    var v28 float32 = v3
    for fi29 := 0; fi29 < 18; fi29++ {
        if false {
            var v30 uint64 = 594
            _ = len(v6)
            _ = v30
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            v32 := GenS2{F0: 255, F1: 1000, F2: 1}
            _ = fi31
            _ = v32
        }
        _ = fi29
    }
    v5 = append(v5, genFn0(18446744073709551615))
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 bool = (0 != 2)
        _ = fi33
        _ = v34
    }
    v1 = -2.25
    fmt.Println(v3)
    v1 += v28
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
