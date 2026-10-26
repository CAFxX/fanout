package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int8; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
func genFn0(p0 uint, p1 int32, p2 uint32) string {
    v6 := []int{777, 32767}
    v7 := GenS1{F0: (127 >= 32768), F1: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = (0.0 + 3.14159)
    var v3 string = "foo bar"
    var v4 int = 42
    v5 := make(map[int]string)
    v1[870] = 1
    switch v3 {
    case "a":
        v2 -= v2
    case "\n\t":
        for fi8 := 0; fi8 < 12; fi8++ {
            if false {
                var v9 int32 = -2147483648
                var v10 int64 = 0
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 13; fi11++ {
                    if false {
                        var v12 uint32 = 2147483648
                        switch v4 {
                        case -9223372036854775808:
                            v13 := v5[fi11]
                            _ = v13
                        case -747:
                            if true {
                                var v14 bool = true
                                v15 := []int{fi11, (fi11 & (fi11 / ((fi11) | 1)))}
                                _ = v14
                                _ = v15
                            } else {
                                v16 := []int{(fi8 + fi8), fi8}
                                v17 := GenS1{F0: false, F1: 255}
                                var v18 bool = true
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        default:
                            delete(v1, v4)
                        }
                        if (1 > 2) {
                            if true {
                                fmt.Println(v12)
                                var v19 int64 = -1
                                v2 -= (v2 - v2)
                                if (false || false) {
                                    v12 += (v12 ^ v12)
                                    v20 := GenS2{F0: (v3 + (v3 + (v3 + ""))), F1: (^(v12))}
                                    _ = v20
                                } else {
                                    var v21 uint16 = 885
                                    v22 := []int{(fi11 + fi8)}
                                    _ = v21
                                    _ = v22
                                }
                                var v23 float64 = (1e10 + (float64(1e10) * float64(-2.25)))
                                _ = v19
                                _ = v23
                            } else {
                                var v24 uint16 = 32767
                                _ = v24
                            }
                            var v25 int64 = 0
                            _ = v25
                        }
                        _ = v12
                    }
                    var v26 float32 = float32((3.14159 + v2))
                    v27 := []int{}
                    _ = fi11
                    _ = v26
                    _ = v27
                }
            }
            var v28 uint32 = 65535
            for fi29 := 0; fi29 < 19; fi29++ {
                fmt.Println(fi29, v28, v2)
                v2 -= (v2 - float64(((1e10 * 1e10) - (-2.25 + 1e10))))
                v3 = v3
                fmt.Println(fi8, v2, fi29)
                var v30 uint32 = (v28 << ((4294967295) & 63))
                _ = fi29
                _ = v30
            }
            _ = fi8
            _ = v28
        }
    default:
        var v31 float32 = float32(((v2 + (v2 * float64((-2.25 + 3.14159)))) / (v2 / 0.1)))
        _ = v31
    }
    if false {
        v32 := []int{}
        fmt.Println(v2)
        for fi33 := 0; fi33 < 2; fi33++ {
            switch 127 {
            case 423:
                v2 = 3.14159
            case 42:
                v4 -= (^(v4))
            case 1:
                var v34 int = v4
                _ = v34
            default:
                if (1000 <= 897) {
                    v3 = (("" + v3) + (v3 + (v3 + v3)))
                } else {
                    if true {
                        var v35 float64 = (1e10 - v2)
                        v36 := GenS3{F0: 2, F1: false, F2: 526}
                        _ = v35
                        _ = v36
                    }
                    fmt.Println(v4, v2)
                    fmt.Println(v2)
                }
            }
            _ = fi33
        }
        _ = v32
    } else {
        v3 = ("" + "\n\t")
        var v37 int32 = 1
        var v38 uint16 = 632
        _ = v37
        _ = v38
    }
    fmt.Println(v4)
    var v39 int = v4
    var v40 uint32 = 4294967295
    v39 -= (v4 | (v4 << ((65535) & 63)))
    v1[(v4 % ((v39) | 1))] = v40
    v41 := GenS2{F0: v3, F1: (v40 % ((864) | 1))}
    v42 := GenS4{F0: 471, F1: v40, F2: ""}
    v43 := GenS4{F0: 107, F1: (v40 + 0), F2: (v3 + v3)}
    switch v39 {
    case 32767:
        for fi44 := 0; fi44 < 9; fi44++ {
            v45 := GenS2{F0: v3, F1: 1}
            _ = fi44
            _ = v45
        }
    }
    if (false || false) {
        _ = len(v5)
    }
    fmt.Println(v4, v4, v40)
    var v46 int16 = 42
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v46
    _ = v5
}
