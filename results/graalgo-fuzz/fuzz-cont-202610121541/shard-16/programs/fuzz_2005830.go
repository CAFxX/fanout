package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0() int {
    var v6 string = ""
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint = 42
        _ = fi7
        _ = v8
    }
    v9 := []int{-92, 170, 4294967295, 46}
    var v10 float64 = 1.5
    _ = v10
    _ = v6
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 int8 = 100
    v4 := []uint32{2147483648}
    v5 := make(map[int]string)
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            fmt.Println(v3)
        } else {
            if (841 >= 413) {
                fmt.Println(v3)
            }
            var v12 int8 = -1
            v13 := GenS1{F0: genFn0(), F1: -9223372036854775808, F2: (1 != 2147483647)}
            _ = v12
            _ = v13
        }
        v3 += v3
        var v14 float32 = 1.5
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint8 = 0
            var v18 bool = ((v17 * (v17 + (v17 << ((65535) & 63)))) != (v17 ^ 60))
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    v3 = (v3 ^ 1)
    var v19 uint32 = 65535
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS1{F0: -224, F1: 0, F2: ((!false) || ((2 < 0) || (3.14159 < (float64(float32(-2.25)) / 1.5))))}
        var v22 int = genFn0()
        var v23 float64 = float64((float32(1.5) + 3.14159))
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        switch (fi24 * genFn0()) {
        case 2147483647:
            v25 := GenS2{F0: 65535, F1: v19}
            _ = v25
        default:
            var v26 uint32 = 218
            _ = v26
        }
        if (true || true) {
            var v27 uint16 = 32767
            _ = v27
        } else {
            fmt.Println(v3)
            var v28 bool = false
            _ = v28
        }
        v3 -= int8(v19)
        _ = fi24
    }
    var v29 int16 = 1
    for fi30 := 0; fi30 < 4; fi30++ {
        switch "foo bar" {
        case "a":
            if false {
                fmt.Println(v29, fi30, v29)
                var v31 bool = ((!(1 > 1)) || (-2.25 < 0.1))
                _ = v31
            } else {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v19, v3)
                case "hello":
                    v29 = v29
                default:
                    if false {
                        for fi32 := 0; fi32 < 9; fi32++ {
                            var v33 float32 = float32((3.14159 - 0.1))
                            var v34 int = fi32
                            if ((v3 + v3) >= (-(v3))) {
                                var v35 uint64 = 42
                                _ = v35
                            } else {
                                var v36 int32 = 42
                                v37 := GenS2{F0: 42, F1: v19}
                                if true {
                                    var v38 uint = 42
                                    var v39 int = 4294967295
                                    if ((true && false) && true) {
                                        var v40 int = 127
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v40
                                    }
                                    fmt.Println(v19, v29, v38)
                                    v39 += (fi30 / (((-(fi30))) | 1))
                                    _ = v38
                                    _ = v39
                                }
                                _ = v36
                                _ = v37
                            }
                            _ = len(v4)
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        v41 := GenS2{F0: 65535, F1: v19}
                        _ = v41
                    }
                }
                var v42 int64 = -624
                _ = v42
            }
        default:
            v43 := GenS1{F0: (-(fi30)), F1: -1, F2: true}
            _ = v43
        }
        _ = fi30
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        var v45 int8 = v3
        v4 = append(v4, (v19 ^ v19))
        var v46 uint32 = v19
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
