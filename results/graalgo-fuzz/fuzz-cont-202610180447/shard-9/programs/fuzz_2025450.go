package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
func genFn0() int {
    var v4 int16 = -1
    if false {
        var v5 int64 = -643
        v6 := []int{}
        for fi7 := 0; fi7 < 18; fi7++ {
            switch fi7 {
            case -9223372036854775808:
                var v8 int = (-(fi7))
                _ = v8
            default:
                if ((!true) || (42 <= 42)) {
                    v6 = append(v6, fi7)
                } else {
                    _ = len(v6)
                }
            }
            v6 = append(v6, (fi7 & (fi7 >> ((0) & 63))))
            _ = fi7
        }
        _ = v5
        _ = v6
    } else {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 float64 = ((1e10 / (((3.14159 + -2.25) * 1e10) / 1.5)) - -2.25)
            for fi11 := 0; fi11 < 13; fi11++ {
                fmt.Println(fi9, v10, fi11)
                _ = fi11
            }
            fmt.Println(v4)
            _ = fi9
            _ = v10
        }
        v12 := []int{32767, 123, -9223372036854775808}
        if false {
            v4 -= 0
            var v13 uint16 = 1
            _ = v13
        }
        _ = v12
    }
    fmt.Println(v4)
    _ = v4
    return -2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 255
    v3 := make(map[uint32]string)
    v14 := []int{1, 127, -228, 2147483647}
    delete(v3, 309)
    for fi15 := 0; fi15 < 9; fi15++ {
        switch 2147483647 {
        case -348:
            if (-41 >= -1) {
                fmt.Println(v2, v2)
            } else {
                var v16 int8 = 0
                _ = v16
            }
        case -9223372036854775808:
            var v17 int64 = -254
            _ = v17
        case -2147483648:
            if true {
                v2 += v2
                if false {
                    if (1 == 0) {
                        v2 = (v2 >> ((65535) & 63))
                        v18 := GenS1{F0: 42, F1: 250}
                        _ = v18
                    } else {
                        var v19 uint16 = 393
                        for fi20 := 0; fi20 < 14; fi20++ {
                            var v21 int64 = 42
                            _ = fi20
                            _ = v21
                        }
                        fmt.Println(fi15)
                        _ = v19
                    }
                    var v22 float32 = (float32(3.14159) / 1.5)
                    if true {
                        v23 := GenS1{F0: -9223372036854775808, F1: v2}
                        _ = v23
                    } else {
                        var v24 bool = true
                        var v25 int8 = -1
                        for fi26 := 0; fi26 < 1; fi26++ {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            v27 := GenS1{F0: int64(fi15), F1: (v2 ^ (v2 % ((1) | 1)))}
                            v28 := GenS1{F0: 718, F1: (v2 * 42)}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                } else {
                    if len(v14) > 0 {
                        v14[0] = (fi15 << ((1) & 63))
                    }
                }
            } else {
                for fi29 := 0; fi29 < 20; fi29++ {
                    v2 = 0
                    v2 += v2
                    _ = fi29
                }
                v2 += (v2 - v2)
            }
        }
        _ = fi15
    }
    var v30 int64 = 720
    v30 = 160
    if len(v14) > 0 {
        v14[0] = 2147483647
    }
    v30 = v30
    if false {
        switch "" {
        case "":
            var v31 uint64 = 4294967295
            _ = v31
        default:
            for fi32 := 0; fi32 < 3; fi32++ {
                var v33 bool = true
                _ = fi32
                _ = v33
            }
        }
        v30 -= v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
}
