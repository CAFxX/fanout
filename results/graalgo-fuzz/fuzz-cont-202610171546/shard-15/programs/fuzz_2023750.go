package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 int32 }
type GenS4 struct { F0 int; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int32; F1 int32 }
func genFn0() uint {
    if true {
        switch 2147483647 {
        case 2147483647:
            var v5 int16 = 42
            _ = v5
        default:
            var v6 float32 = 3.14159
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int8 = -128
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "":
        if false {
            v9 := []int{-1, -9223372036854775808, 32768}
            for fi10 := 0; fi10 < 9; fi10++ {
                fmt.Println(fi10)
                v11 := GenS1{F0: 1000, F1: 981, F2: 218}
                v12 := GenS2{F0: (!(!false)), F1: 682}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        } else {
            var v13 int8 = 127
            v13 -= (v13 >> ((431) & 63))
            for fi14 := 0; fi14 < 19; fi14++ {
                v13 += (v13 << ((2) & 63))
                v15 := GenS2{F0: (0 > 364), F1: 639}
                _ = fi14
                _ = v15
            }
            _ = v13
        }
    case "hello":
        switch -801 {
        case -2147483648:
            if (3.14159 == (0.1 * 1e10)) {
                if false {
                    var v16 float64 = 0.0
                    if false {
                        for fi17 := 0; fi17 < 17; fi17++ {
                            v18 := GenS2{F0: false, F1: 42}
                            if false {
                                for fi19 := 0; fi19 < 16; fi19++ {
                                    fmt.Println(v16, v16, fi17)
                                    fmt.Println(fi17)
                                    v16 = (float64((0.0 * 0.1)) / (((1e10 + (float64((0.0 / ((-2.25 - 1e10) / -2.25))) - (float64((1e10 * -2.25)) - v16))) + float64(float32(v16))) * v16))
                                    v20 := []int{fi17, (fi19 - fi17)}
                                    _ = fi19
                                    _ = v20
                                }
                                if (753 != 590) {
                                    v16 = float64(1e10)
                                }
                                for fi21 := 0; fi21 < 12; fi21++ {
                                    var v22 uint = 1
                                    if (!false) {
                                        v22 -= v22
                                    }
                                    _ = fi21
                                    _ = v22
                                }
                            } else {
                                if (1 <= 436) {
                                    var v23 int = fi17
                                    _ = v23
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        v16 = -2.25
                                    }
                                    v24 := GenS3{F0: 42, F1: -32768, F2: 2147483647}
                                    _ = v24
                                }
                                v25 := []int{1, 0, (fi17 * (fi17 & fi17)), 32767}
                                v26 := GenS4{F0: fi17, F1: 140, F2: 4294967295}
                                _ = v25
                                _ = v26
                            }
                            v27 := []int{(-(fi17)), fi17, fi17, (-(fi17))}
                            _ = fi17
                            _ = v18
                            _ = v27
                        }
                    } else {
                        v16 += ((v16 * (v16 - v16)) / v16)
                    }
                    if true {
                        var v28 uint16 = 2
                        _ = v28
                    } else {
                        v16 += v16
                    }
                    _ = v16
                } else {
                    var v29 uint32 = 2
                    _ = v29
                }
                v30 := []int{-545}
                var v31 int16 = 470
                _ = v30
                _ = v31
            } else {
                v32 := GenS5{F0: 65535, F1: 128}
                var v33 uint64 = 1
                _ = v32
                _ = v33
            }
        default:
            var v34 uint8 = 255
            _ = v34
        }
    case "foo bar":
        var v35 uint8 = 1
        _ = v35
    default:
        var v36 uint8 = 128
        _ = v36
    }
    if false {
        v37 := GenS1{F0: 1, F1: 42, F2: 1000}
        v38 := GenS6{F0: 1, F1: -712}
        _ = v37
        _ = v38
    } else {
        v39 := GenS4{F0: -1, F1: 1, F2: 382}
        var v40 int16 = 0
        _ = v39
        _ = v40
    }
    return 5
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2147483647
    var v3 uint64 = 0
    v4 := make(map[uint32]string)
    var v41 int8 = -128
    fmt.Println(v2)
    v42 := []int{0, -262, 4294967295, 2147483647}
    v43 := v4[118]
    var v44 uint = 2
    var v45 uint8 = 0
    for fi46 := 0; fi46 < 5; fi46++ {
        fmt.Println(v2, v44, v43)
        var v47 int8 = 42
        _ = fi46
        _ = v47
    }
    v48 := GenS2{F0: (!false), F1: v2}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v48
}
