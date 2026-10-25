package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint) int {
    v5 := []int{0, -109}
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(fi6, p0)
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(p0, fi6, fi6)
            p0 -= (p0 & p0)
            p0 += 65535
            p0 = p0
            _ = fi7
        }
        p0 = (p0 / ((p0) | 1))
        _ = fi6
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        p0 += (p0 << ((p0) & 63))
        p0 -= (p0 << ((p0) & 63))
        p0 = 688
        _ = fi8
    }
    if (65535 != 4294967295) {
        var v9 string = "\n\t"
        p0 = p0
        fmt.Println(v9, v9)
        _ = v9
    } else {
        p0 -= 65535
        fmt.Println(p0)
    }
    p0 -= p0
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 string, p1 uint, p2 uint16) bool {
    v10 := []int{127, 4294967295}
    for fi11 := 0; fi11 < 13; fi11++ {
        if false {
            fmt.Println(p2)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        } else {
            if false {
                v10 = append(v10, fi11)
                var v12 int16 = 711
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 int8 = 68
                    switch fi11 {
                    case 481:
                        if (2 == (p1 << ((p1) & 63))) {
                            var v15 int = (fi13 | genFn0(p1))
                            if len(v10) > 0 {
                                v10[0] = fi13
                            }
                            _ = v15
                        } else {
                            fmt.Println(v14, fi13, fi13)
                            v12 = 411
                            var v16 int8 = (-(v14))
                            _ = v16
                        }
                    case 4294967295:
                        v17 := GenS1{F0: 62, F1: 42}
                        _ = v17
                    case 0:
                        v14 -= v14
                    default:
                        var v18 int = fi13
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                }
                _ = v12
            } else {
                if false {
                    p1 -= p1
                    var v19 uint64 = 4294967295
                    _ = v19
                }
            }
            var v20 int32 = -658
            _ = v20
        }
        _ = fi11
    }
    p2 -= (p2 & (p2 ^ (p2 * 609)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return false
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 1
    var v3 uint16 = 65535
    var v4 int32 = 1
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 string = "hello"
        _ = fi21
        _ = v22
    }
    fmt.Println(v2)
    if (-1 >= genFn0(uint(v2))) {
        if false {
            switch 32767 {
            case 32768:
                v3 = (v3 / ((1) | 1))
            default:
                var v23 bool = ("hello" > "a")
                _ = v23
            }
        } else {
            switch 32767 {
            case 2147483647:
                _ = len(v1)
            case 600:
                fmt.Println(v4)
            case 1:
                delete(v1, genFn0(984))
            default:
                v3 += 221
            }
            v2 = 65535
        }
    } else {
        var v24 uint8 = 127
        if (!(((0.0 - 1.5) - 1e10) == -2.25)) {
            fmt.Println(v4, v24, v2)
        }
        _ = v24
    }
    v25 := GenS2{F0: 912}
    v26 := []int{2147483647, genFn0(65535), int(v3), 687}
    fmt.Println(v3, v3, v2)
    switch "hello" {
    case "a":
        var v27 float64 = ((0.0 + (((float64(1.5) * 1e10) - ((1e10 - 3.14159) * 0.0)) * (0.1 / 1.0))) + 1e10)
        _ = v27
    case "":
        v3 += (v3 >> ((255) & 63))
    case "\n\t":
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    default:
        switch 32767 {
        case 127:
            if (1 < v3) {
                for fi28 := 0; fi28 < 8; fi28++ {
                    v3 = 2
                    v29 := GenS2{F0: 65535}
                    _ = fi28
                    _ = v29
                }
            } else {
                var v30 uint32 = 772
                _ = v30
            }
        case 32768:
            switch 4294967295 {
            case 618:
                if true {
                    v31 := GenS1{F0: 19, F1: 4294967295}
                    _ = v31
                }
            case 1:
                v2 -= v2
            case 2147483647:
                var v32 uint = 255
                _ = v32
            }
        case 942:
            var v33 bool = (18446744073709551615 > 2)
            _ = v33
        }
    }
    v34 := GenS2{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
}
