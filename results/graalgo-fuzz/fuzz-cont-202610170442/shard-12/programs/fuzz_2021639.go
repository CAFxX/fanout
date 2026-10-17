package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint; F2 uint8 }
func genFn0(p0 uint64, p1 int) int32 {
    if false {
        v4 := []int{p1, -2147483648}
        p0 -= (p0 / ((p0) | 1))
        if (((0.1 / (1e10 + -2.25)) + (0.1 - 1e10)) <= (float64(float32(1.5)) + 0.0)) {
            for fi5 := 0; fi5 < 8; fi5++ {
                if (-9223372036854775808 == 42) {
                    p0 = (p0 << ((42) & 63))
                    p0 = (p0 % ((p0) | 1))
                } else {
                    switch p1 {
                    case -2147483648:
                        for fi6 := 0; fi6 < 3; fi6++ {
                            var v7 float32 = 0.1
                            p0 -= (^(p0))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi6
                            _ = v7
                        }
                    default:
                        for fi8 := 0; fi8 < 19; fi8++ {
                            fmt.Println(p0)
                            if ("foo bar" == ("\n\t" + "\n\t")) {
                                p1 = 32768
                                var v9 uint16 = 255
                                p0 = (p0 * p0)
                                _ = v9
                            } else {
                                if false {
                                    v10 := GenS1{F0: fi5}
                                    v11 := GenS1{F0: (p1 / (((^(p1))) | 1))}
                                    _ = v10
                                    _ = v11
                                }
                                fmt.Println(fi8, fi5)
                                p1 = 0
                            }
                            _ = fi8
                        }
                    }
                    fmt.Println(fi5)
                }
                _ = fi5
            }
        } else {
            if false {
                p1 = (p1 + p1)
            }
        }
        _ = v4
    }
    var v12 uint8 = 255
    _ = p0
    _ = p1
    _ = v12
    return -2147483648
}

func main() {
    var v1 bool = true
    v2 := []uint32{0, 42, 0}
    var v3 int8 = -8
    v1 = false
    if v1 {
        v1 = v1
    }
    var v13 int16 = 301
    v2 = append(v2, 0)
    var v14 uint64 = 449
    var v15 float32 = (0.0 / (0.1 * 0.1))
    for fi16 := 0; fi16 < 12; fi16++ {
        v2 = append(v2, 42)
        _ = fi16
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 uint16 = 0
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v20 := GenS1{F0: -158}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v21 := GenS1{F0: -9223372036854775808}
    v22 := GenS2{F0: (uint8(v14) >= 255), F1: 4294967295, F2: 246}
    v14 = v14
    for fi23 := 0; fi23 < 1; fi23++ {
        for fi24 := 0; fi24 < 14; fi24++ {
            _ = len(v2)
            var v25 uint16 = 30
            var v26 int = (fi23 + 0)
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 float64 = 0.1
        if false {
            v3 += v3
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v28 uint = 2
            _ = v28
        }
        _ = fi23
        _ = v27
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    v2 = append(v2, 2147483647)
    if (!v1) {
        if false {
            _ = len(v2)
            var v29 int = 287
            _ = v29
        } else {
            switch 32768 {
            case 9223372036854775807:
                var v30 int = 2147483647
                _ = v30
            case 2147483647:
                fmt.Println(v15, v14, v1)
            }
            v2 = append(v2, 42)
        }
        switch 99 {
        case 616:
            var v31 string = "a"
            _ = v31
        case -425:
            for fi32 := 0; fi32 < 7; fi32++ {
                v3 = (v3 | 50)
                var v33 uint64 = v14
                v34 := GenS1{F0: (fi32 << ((1) & 63))}
                _ = fi32
                _ = v33
                _ = v34
            }
        case 487:
            v3 = 0
        }
    } else {
        switch "" {
        case "\n\t":
            if (9223372036854775807 != 0) {
                v15 += v15
                v1 = false
            }
        case "":
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
        }
        v35 := GenS1{F0: 32768}
        for fi36 := 0; fi36 < 9; fi36++ {
            if ((-(fi36)) == -2147483648) {
                v15 = (v15 * -2.25)
            }
            _ = fi36
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
