package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 int64, p1 int8, p2 uint16) uint8 {
    p0 = (-(p0))
    for fi6 := 0; fi6 < 2; fi6++ {
        p2 = p2
        switch "\n\t" {
        case "hello":
            var v7 uint32 = 2
            _ = v7
        case "\n\t":
            if true {
                p1 -= 0
                var v8 int8 = (p1 << ((255) & 63))
                _ = v8
            } else {
                switch (fi6 >> ((255) & 63)) {
                case 32767:
                    fmt.Println(p1, p2, p0)
                default:
                    var v9 float64 = 0.1
                    _ = v9
                }
                v10 := []int{}
                if true {
                    if (-2147483648 >= -1) {
                        v11 := GenS1{F0: 863, F1: (p2 % ((p2) | 1)), F2: (-(p1))}
                        fmt.Println(p2)
                        _ = v11
                    }
                } else {
                    var v12 float64 = 3.14159
                    var v13 uint8 = 0
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
        default:
            for fi14 := 0; fi14 < 4; fi14++ {
                if (859 != 189) {
                    v15 := []int{(fi14 + (fi6 % (((fi14 | (fi14 & 32768))) | 1)))}
                    p2 -= (p2 / ((0) | 1))
                    _ = v15
                } else {
                    switch fi6 {
                    case 573:
                        v16 := []int{}
                        _ = v16
                    case 4294967295:
                        v17 := []int{fi6, fi14, (fi14 * fi14), 127}
                        _ = v17
                    }
                    if (463 >= 42) {
                        p0 = p0
                        v18 := GenS1{F0: p2, F1: (p2 & p2), F2: 127}
                        v19 := []int{fi14}
                        _ = v18
                        _ = v19
                    } else {
                        p0 = -626
                    }
                }
                p0 += p0
                if true {
                    p0 += 946
                    var v20 int8 = p1
                    p2 = (p2 << ((0) & 63))
                    _ = v20
                }
                _ = fi14
            }
        }
        _ = fi6
    }
    var v21 int64 = (p0 ^ (p0 / (((p0 | 9223372036854775807)) | 1)))
    for fi22 := 0; fi22 < 8; fi22++ {
        for fi23 := 0; fi23 < 12; fi23++ {
            fmt.Println(p2, p0)
            v24 := GenS1{F0: (^(p2)), F1: (p2 / ((p2) | 1)), F2: (p1 | 0)}
            var v25 uint32 = uint32(p2)
            var v26 bool = false
            p0 += 9223372036854775807
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi22
    }
    v21 = (p0 / (((v21 - -9223372036854775808)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 128
}
func genFn1(p0 int8, p1 uint8) uint {
    var v27 uint32 = 612
    if false {
        v28 := []int{32767, -2147483648, 32767, -9223372036854775808}
        p1 -= (p1 & (p1 | (^(p1))))
        v27 = v27
        _ = v28
    } else {
        p0 = (-(p0))
        var v29 uint32 = (v27 & 869)
        if true {
            var v30 int = 1
            v31 := []int{-1, v30, (v30 - (v30 % (((v30 ^ v30)) | 1))), v30}
            _ = v30
            _ = v31
        } else {
            if true {
                var v32 uint = 1
                _ = v32
            } else {
                v29 = v29
                v33 := GenS1{F0: 32767, F1: 818, F2: (p0 * (p0 << ((255) & 63)))}
                v29 = 1
                _ = v33
            }
            if true {
                var v34 uint64 = 286
                if (!(true && (-1 > -1))) {
                    v34 = 4294967295
                    if false {
                        v27 = 1
                    }
                    v27 += v27
                }
                _ = v34
            } else {
                for fi35 := 0; fi35 < 17; fi35++ {
                    v36 := GenS2{F0: 154, F1: -1}
                    if ((p0 < p0) && false) {
                        p0 = (p0 ^ p0)
                        var v37 uint16 = 361
                        var v38 int16 = int16(v27)
                        _ = v37
                        _ = v38
                    } else {
                        var v39 float32 = float32((0.1 - 1.5))
                        var v40 uint64 = 2
                        _ = v39
                        _ = v40
                    }
                    _ = fi35
                    _ = v36
                }
                for fi41 := 0; fi41 < 8; fi41++ {
                    v27 = v29
                    v42 := []int{(fi41 % (((-(fi41))) | 1)), -9223372036854775808, -2147483648}
                    var v43 uint16 = 2
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
            var v44 int32 = -740
            _ = v44
        }
        _ = v29
    }
    _ = p0
    _ = p1
    _ = v27
    return 255
}

func main() {
    var v1 uint8 = 2
    var v2 int = 9223372036854775807
    v3 := make(map[string]uint32)
    var v4 uint64 = 713
    v5 := make(map[string]uint32)
    v2 = (v2 - 128)
    fmt.Println(v2, v2)
    v2 = v2
    v1 = genFn0(9223372036854775807, 127, 42)
    for fi45 := 0; fi45 < 18; fi45++ {
        v46 := []int{}
        _ = fi45
        _ = v46
    }
    fmt.Println(v2, v4)
    v47 := GenS2{F0: 0, F1: 731}
    var v48 uint64 = (v4 / ((981) | 1))
    v48 = 88
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
