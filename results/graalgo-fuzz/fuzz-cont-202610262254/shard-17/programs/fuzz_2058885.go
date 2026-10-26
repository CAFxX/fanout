package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int8; F1 uint16; F2 int8 }
func genFn0(p0 uint16, p1 int16, p2 int16) int32 {
    fmt.Println(p1, p0)
    v4 := []int{0, 4294967295, 0, -424}
    if (!false) {
        var v5 int8 = int8(p1)
        _ = v5
    } else {
        var v6 uint8 = 47
        switch 32767 {
        case 730:
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 int8 = -128
                v9 := GenS1{F0: (p1 + p1), F1: v8}
                _ = fi7
                _ = v8
                _ = v9
            }
        case -474:
            _ = len(v4)
        default:
            switch 9223372036854775807 {
            case 32768:
                for fi10 := 0; fi10 < 16; fi10++ {
                    var v11 uint = 312
                    p1 += (p1 & -32768)
                    fmt.Println(fi10, v11, p1)
                    _ = fi10
                    _ = v11
                }
            case 1:
                var v12 uint16 = uint16(v6)
                _ = v12
            case -496:
                if len(v4) > 0 {
                    v4[0] = 42
                }
            default:
                if ((p2 / ((p2) | 1)) > (-(p2))) {
                    if (1e10 > 1.5) {
                        var v13 float32 = (3.14159 - 3.14159)
                        _ = v13
                    }
                    var v14 float32 = 1e10
                    p2 = p1
                    _ = v14
                }
            }
        }
        p1 += (-(p1))
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2147483647
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 100
    var v3 float64 = float64(-2.25)
    var v15 uint = 4294967295
    v3 += float64(1.5)
    switch v1 {
    case "":
        v1 = "a"
    case "foo bar":
        var v16 int16 = 32767
        _ = v16
    case "hello":
        v17 := GenS2{F0: 912, F1: 0}
        _ = v17
    }
    switch "hello" {
    case "foo bar":
        v18 := []int{9223372036854775807, 128, -2147483648, 2147483647}
        _ = v18
    case "\n\t":
        fmt.Println(v3, v2, v1)
    case "hello":
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 float32 = -2.25
            if ((uint8(v15) > 0) && (!true)) {
                fmt.Println(v1)
            } else {
                if (!true) {
                    switch v1 {
                    case "a":
                        if true {
                            switch fi19 {
                            case -905:
                                v20 += -2.25
                            case 32767:
                                v1 = "\n\t"
                            case 127:
                                if (!((!true) && true)) {
                                    v21 := []int{(fi19 << ((0) & 63))}
                                    var v22 uint64 = 4294967295
                                    switch fi19 {
                                    case 1:
                                        _ = len(v21)
                                    }
                                    _ = v21
                                    _ = v22
                                } else {
                                    v23 := GenS3{F0: v2, F1: 1, F2: -61}
                                    v24 := GenS1{F0: 1000, F1: int8(v15)}
                                    _ = v23
                                    _ = v24
                                }
                            }
                        } else {
                            fmt.Println(v15, v1, v1)
                            if (float32(v3) < v20) {
                                var v25 uint64 = 2
                                _ = v25
                            }
                        }
                    case "\n\t":
                        var v26 int32 = -48
                        _ = v26
                    case "hello":
                        var v27 int32 = 1
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 float32 = 3.14159
                        switch fi28 {
                        case 32768:
                            if (int64(fi28) <= -1) {
                                var v30 uint64 = 0
                                _ = v30
                            }
                        case -189:
                            v29 = v20
                        }
                        v31 := GenS2{F0: 4294967295, F1: -1}
                        _ = fi28
                        _ = v29
                        _ = v31
                    }
                    v15 += 0
                } else {
                    v3 -= float64(float32(v3))
                }
            }
            v32 := []int{fi19}
            _ = fi19
            _ = v20
            _ = v32
        }
    }
    if (2147483647 >= -1) {
        var v33 int32 = genFn0(32767, 1000, 592)
        v1 = ((("" + (("hello" + "a") + v1)) + ((v1 + (((v1 + (v1 + ((v1 + v1) + ""))) + "foo bar") + v1)) + v1)) + (v1 + v1))
        var v34 float64 = v3
        _ = v33
        _ = v34
    }
    var v35 uint = 2
    var v36 float64 = (v3 * v3)
    var v37 uint64 = 101
    v38 := []int{}
    _ = len(v38)
    v39 := GenS2{F0: 2147483647, F1: -1}
    if (-9223372036854775808 >= 1) {
        if ((-2.25 - 3.14159) > ((float32(1e10) - float32((1.5 / 0.1))) - (0.1 - 3.14159))) {
            if (2 == 32767) {
                if len(v38) > 0 {
                    v38[0] = 388
                }
                for fi40 := 0; fi40 < 4; fi40++ {
                    fmt.Println(v15)
                    if len(v38) >= 2 {
                        v38 = v38[1:len(v38)]
                    }
                    _ = fi40
                }
            }
        }
        if false {
            var v41 bool = false
            _ = v41
        }
        v35 = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
