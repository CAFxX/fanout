package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 bool; F1 uint16; F2 uint16 }

func main() {
    var v1 int32 = 369
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    var v4 int = 128
    _ = len(v3)
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint16 = 32767
        v6 = 42
        if (598 != 497) {
            v7 := GenS1{F0: (^(v6)), F1: 42, F2: 1}
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    v8 := GenS1{F0: 65535, F1: -9223372036854775808, F2: 1}
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            switch "foo bar" {
            case "\n\t":
                var v11 float64 = (1e10 / 1.0)
                _ = v11
            default:
                _ = len(v3)
            }
            v12 := []int{-9223372036854775808}
            _ = fi10
            _ = v12
        }
        fmt.Println(fi9, fi9)
        var v13 float64 = -2.25
        _ = fi9
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        if (9223372036854775807 == -222) {
            v15 := []int{-596, (v4 << ((18446744073709551615) & 63))}
            _ = v15
        } else {
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := GenS1{F0: 32767, F1: 1, F2: 794}
                var v18 int8 = -32
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v1 = (^(v1))
        _ = fi14
    }
    var v19 uint32 = 65535
    for fi20 := 0; fi20 < 19; fi20++ {
        v1 += v1
        v3[v19] = fi20
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int = (fi20 & v4)
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 bool = (65535 >= 896)
        var v25 int32 = (v1 & (v1 - v1))
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        if false {
            var v27 int32 = 374
            _ = v27
        } else {
            v1 += (-(v1))
        }
        fmt.Println(v19, v4, v4)
        for fi28 := 0; fi28 < 2; fi28++ {
            var v29 uint16 = 1
            if (true && false) {
                if true {
                    var v30 string = "foo bar"
                    _ = v30
                }
                if true {
                    if true {
                        v31 := GenS2{F0: (!true), F1: v29}
                        v32 := GenS3{F0: 1}
                        _ = v31
                        _ = v32
                    } else {
                        v33 := GenS4{F0: true, F1: 568, F2: v19}
                        if false {
                            if true {
                                if (!(!((fi28 * (fi28 % ((1) | 1))) > (-(fi26))))) {
                                    v34 := v3[v19]
                                    switch -1 {
                                    case -2147483648:
                                        v35 := GenS5{F0: 1000, F1: -1}
                                        _ = v35
                                    case 42:
                                        fmt.Println(v34)
                                    case 9223372036854775807:
                                        v29 -= v29
                                    default:
                                        v36 := GenS4{F0: true, F1: 0, F2: v19}
                                        _ = v36
                                    }
                                    _ = v34
                                }
                                _ = len(v3)
                            } else {
                                delete(v2, (fi28 >> ((uint(v29)) & 63)))
                            }
                        }
                        _ = v33
                    }
                    var v37 float64 = (3.14159 / 1.0)
                    var v38 int64 = 42
                    if (true && false) {
                        v4 = (-(fi26))
                    }
                    _ = v37
                    _ = v38
                } else {
                    v39 := v3[0]
                    _ = v39
                }
            }
            var v40 uint = 120
            _ = fi28
            _ = v29
            _ = v40
        }
        _ = fi26
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        fmt.Println(v19, v4, v4)
        _ = fi41
    }
    if true {
        if false {
            v19 -= v19
        } else {
            v42 := []int{v4, 32767, 929, v4}
            _ = v42
        }
    }
    v43 := v2[v4]
    var v44 uint32 = v19
    _ = len(v3)
    v45 := GenS6{F0: false, F1: 255, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v8
}
