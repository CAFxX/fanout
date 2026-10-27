package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint16, p1 int32) uint8 {
    var v4 bool = true
    p0 += 2
    _ = p0
    _ = p1
    _ = v4
    return 128
}
func genFn1(p0 uint, p1 int16, p2 string) int16 {
    var v5 uint64 = 2
    switch -9223372036854775808 {
    case 42:
        v5 += (^(v5))
    case 127:
        v6 := GenS1{F0: (v5 << (((p0 ^ p0)) & 63)), F1: -1, F2: 65}
        _ = v6
    }
    v5 = 1
    v7 := GenS1{F0: (^(v5)), F1: -128, F2: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return -32768
}

func main() {
    var v1 int16 = 621
    v2 := make(map[int]bool)
    var v3 int8 = 127
    var v8 int64 = 0
    v9 := GenS2{F0: -2147483648, F1: genFn1(952, (v1 >> ((1) & 63)), "")}
    if ((!true) && ((9223372036854775807 < v8) || (127 > 127))) {
        var v10 int = int(v3)
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        if true {
            var v12 float32 = ((1.5 * 1.5) - 1e10)
            _ = v12
        } else {
            v13 := GenS1{F0: 665, F1: v3, F2: 42}
            _ = v13
        }
        v14 := GenS3{F0: 165}
        _ = fi11
        _ = v14
    }
    if true {
        v3 = v3
        fmt.Println(v3, v3)
        v8 -= (v8 % ((-9223372036854775808) | 1))
    }
    if false {
        v8 -= v8
        v15 := []int{4294967295, -289, 9223372036854775807, 127}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 9; fi16++ {
            if ("\n\t" < "") {
                v1 = (-(v1))
                fmt.Println(v8)
                var v17 uint = 2
                _ = v17
            } else {
                switch fi16 {
                case -9223372036854775808:
                    v18 := GenS1{F0: 4294967295, F1: v3, F2: 241}
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 1; fi19++ {
                        v8 += (v8 >> ((0) & 63))
                        _ = fi19
                    }
                }
                switch (fi16 << ((4294967295) & 63)) {
                case 127:
                    _ = len(v2)
                }
                v20 := []int{}
                _ = v20
            }
            _ = fi16
        }
    }
    if false {
        var v21 float64 = (((0.0 - (0.0 + -2.25)) * (1e10 * 0.0)) + float64(1.5))
        if false {
            var v22 uint32 = 495
            fmt.Println(v3)
            fmt.Println(v22, v21, v1)
            _ = v22
        } else {
            v21 = v21
            if (!(!(!((-2.25 / (1.5 * 1.5)) > ((-2.25 * ((((((((((1e10 / 1.5) * (1.5 * -2.25)) / 3.14159) * -2.25) * 3.14159) / float32((((1.5 / (v21 - v21)) - (v21 / 1.0)) + v21))) / -2.25) * 0.0) + 0.1) * 0.1)) / 1.0))))) {
                if false {
                    var v23 float64 = (((0.0 / v21) + 0.0) + v21)
                    if (!true) {
                        if true {
                            v21 -= v21
                        }
                        switch 32767 {
                        case 32768:
                            v24 := []int{}
                            _ = v24
                        case 2147483647:
                            var v25 int64 = (v8 >> ((18446744073709551615) & 63))
                            _ = v25
                        default:
                            _ = len(v2)
                        }
                        switch -1 {
                        case 4294967295:
                            var v26 int32 = -631
                            _ = v26
                        case 9223372036854775807:
                            if false {
                                v2[128] = true
                                var v27 uint16 = 65535
                                _ = v27
                            } else {
                                v28 := []int{1, -1}
                                v8 += (v8 & v8)
                                v29 := GenS3{F0: 216}
                                _ = v28
                                _ = v29
                            }
                        case 128:
                            switch 42 {
                            case 4294967295:
                                var v30 int16 = 32767
                                _ = v30
                            default:
                                var v31 int = 200
                                _ = v31
                            }
                        }
                    } else {
                        for fi32 := 0; fi32 < 2; fi32++ {
                            v33 := []int{649, 9223372036854775807, (fi32 ^ (fi32 + -1))}
                            v8 = 9223372036854775807
                            fmt.Println(v8, v3)
                            _ = fi32
                            _ = v33
                        }
                        v34 := GenS1{F0: 1, F1: v3, F2: 30}
                        v3 -= -11
                        _ = v34
                    }
                    _ = v23
                } else {
                    var v35 float64 = -2.25
                    _ = v35
                }
                var v36 uint16 = 255
                _ = v36
            } else {
                fmt.Println(v8)
                v8 = (v8 << ((42) & 63))
                v8 -= (v8 - v8)
            }
            v37 := GenS3{F0: 1}
            _ = v37
        }
        v38 := GenS4{F0: 395}
        _ = v21
        _ = v38
    }
    v8 -= (v8 << ((255) & 63))
    v39 := GenS2{F0: 42, F1: genFn1(65535, 0, "\n\t")}
    v1 -= v1
    switch "\n\t" {
    case "":
        var v40 uint16 = 32767
        _ = v40
    case "a":
        var v41 int64 = (v8 - (v8 << ((1) & 63)))
        _ = v41
    case "foo bar":
        v8 -= v8
    default:
        v1 += (v1 * -32768)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v8
    _ = v9
}
