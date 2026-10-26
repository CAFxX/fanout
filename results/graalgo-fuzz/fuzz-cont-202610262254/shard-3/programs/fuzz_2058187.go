package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 uint8; F1 int }
func genFn0(p0 uint8) int64 {
    fmt.Println(p0)
    switch 430 {
    case 9:
        p0 = (^(p0))
    case 0:
        var v7 string = ""
        _ = v7
    case 32768:
        p0 = p0
    default:
        for fi8 := 0; fi8 < 2; fi8++ {
            p0 -= 2
            _ = fi8
        }
    }
    v9 := []int{32768}
    var v10 uint16 = 713
    v9 = append(v9, 4294967295)
    _ = p0
    _ = v10
    _ = v9
    return 332
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := []uint32{42, 105}
    var v4 string = ""
    var v5 int = 137
    var v6 uint16 = 2
    v11 := GenS1{F0: 2, F1: 760, F2: 1}
    v12 := GenS1{F0: 60, F1: 2147483648, F2: 261}
    if (-1 < 2147483647) {
        v6 -= (v6 + v6)
        fmt.Println(v4)
        var v13 uint8 = 255
        _ = v13
    } else {
        v14 := GenS1{F0: 4294967295, F1: 0, F2: 2}
        _ = v14
    }
    var v15 int16 = 182
    v16 := GenS2{F0: (^(v6))}
    v4 = (("\n\t" + "foo bar") + v4)
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            if true {
                var v19 int64 = -1
                v15 -= -1
                if true {
                    v15 += -780
                }
                v4 = ((v4 + (v4 + "hello")) + v4)
                var v20 int32 = 931
                var v21 bool = false
                _ = v19
                _ = v20
                _ = v21
            }
            switch fi17 {
            case 32767:
                fmt.Println(v6, fi18)
            case 128:
                switch v4 {
                case "":
                    if (((fi18 - 4294967295) < v5) || (1 >= 255)) {
                        if false {
                            if ((3.14159 / 1.5) == float64(-2.25)) {
                                fmt.Println(v4, v4)
                                v22 := GenS3{F0: (v6 % ((v6) | 1)), F1: genFn0(255), F2: 1}
                                _ = v22
                            } else {
                                fmt.Println(v15, v15, v5)
                                fmt.Println(fi17, v15)
                            }
                        } else {
                            v5 += -2147483648
                            var v23 uint16 = v6
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch v4 {
                            case "hello":
                                v24 := GenS4{F0: 128, F1: 4294967295}
                                _ = v24
                            case "\n\t":
                                var v25 int8 = 127
                                _ = v25
                            }
                            _ = v23
                        }
                        v15 = v15
                        v4 = v4
                    }
                default:
                    v26 := v2[743]
                    _ = v26
                }
            case 943:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v15 += (v15 / (((v15 % ((32767) | 1))) | 1))
            }
            switch (-(v5)) {
            case 32768:
                v6 -= 42
            case -2147483648:
                var v27 int16 = (v15 * v15)
                _ = v27
            case 601:
                v6 = (v6 | v6)
            default:
                v15 += -687
            }
            _ = fi18
        }
        var v28 uint16 = (v6 ^ (v6 << ((18446744073709551615) & 63)))
        _ = fi17
        _ = v28
    }
    if false {
        fmt.Println(v15)
    }
    if (int16(v6) >= v15) {
        var v29 int = -1
        v4 = "hello"
        _ = v29
    } else {
        v4 = v4
    }
    fmt.Println(v6)
    var v30 int64 = 42
    switch v5 {
    case 128:
        var v31 string = v4
        _ = v31
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    v32 := GenS2{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
