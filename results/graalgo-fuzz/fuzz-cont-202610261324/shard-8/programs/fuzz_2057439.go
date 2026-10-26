package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint16; F2 int32 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint64, p1 bool, p2 uint8) bool {
    p0 = (p0 + p0)
    v6 := []int{127, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}
func genFn1(p0 int16, p1 uint32) int {
    if (88 >= 112) {
        if false {
            v7 := GenS1{F0: 1}
            if true {
                for fi8 := 0; fi8 < 11; fi8++ {
                    fmt.Println(p0)
                    var v9 bool = false
                    _ = fi8
                    _ = v9
                }
                if (!false) {
                    var v10 int16 = 32767
                    _ = v10
                } else {
                    switch 4294967295 {
                    case 32768:
                        for fi11 := 0; fi11 < 13; fi11++ {
                            v12 := GenS2{F0: p0, F1: 0, F2: -728}
                            p1 = 2147483647
                            _ = fi11
                            _ = v12
                        }
                    case 32767:
                        p0 = (p0 >> ((255) & 63))
                    case 4294967295:
                        p1 = 418
                    }
                    v13 := []int{127, -281, 42, 32767}
                    p1 -= (p1 & p1)
                    _ = v13
                }
                if ((true || true) && (127 < 42)) {
                    v14 := []int{-9223372036854775808, 1, -1}
                    _ = v14
                } else {
                    var v15 uint16 = 255
                    var v16 int = 127
                    var v17 int16 = -32768
                    _ = v15
                    _ = v16
                    _ = v17
                }
            } else {
                for fi18 := 0; fi18 < 4; fi18++ {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        p0 -= p0
                        v20 := []int{(fi19 % (((fi19 / (((fi19 % ((fi18) | 1))) | 1))) | 1)), (fi18 + (fi18 - fi18)), fi18, fi19}
                        switch -343 {
                        case 1:
                            fmt.Println(p1, fi18)
                        case -2147483648:
                            if false {
                                v21 := GenS2{F0: (p0 >> ((4294967295) & 63)), F1: 459, F2: -1}
                                fmt.Println(fi19, p0, fi18)
                                _ = v21
                            }
                        case 4294967295:
                            fmt.Println(fi18)
                        default:
                            var v22 uint64 = 4294967295
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
                v23 := []int{-1}
                _ = v23
            }
            var v24 int32 = -2147483648
            _ = v24
            _ = v7
        }
    } else {
        p1 += p1
    }
    v25 := GenS1{F0: 131}
    p0 -= 806
    _ = p0
    _ = p1
    _ = v25
    return 32768
}

func main() {
    var v1 int16 = 42
    v2 := make(map[string]string)
    var v3 int8 = -128
    var v4 uint32 = 42
    var v5 int32 = -1
    v26 := []int{-1}
    v27 := v2[("a" + "a")]
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    v3 -= v3
    fmt.Println(v3)
    if false {
        v3 += (v3 & 100)
    } else {
        v4 -= v4
        v28 := GenS3{F0: (^(v5)), F1: -32768}
        v1 -= (v1 % (((v1 * (v1 >> ((71) & 63)))) | 1))
        _ = v28
    }
    v5 = (v5 ^ v5)
    v5 += (-(v5))
    v29 := GenS3{F0: (v5 << ((42) & 63)), F1: (v1 % (((v1 + v1)) | 1))}
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 uint32 = v4
        v32 := GenS3{F0: v5, F1: (-(v1))}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 uint64 = 257
    if ((0.1 - 0.1) < 1.5) {
        for fi34 := 0; fi34 < 18; fi34++ {
            var v35 int32 = (v5 * v5)
            _ = fi34
            _ = v35
        }
        if true {
            if true {
                v1 -= v1
            }
            var v36 uint16 = 65535
            _ = v36
        }
        switch "" {
        case "hello":
            if (!true) {
                for fi37 := 0; fi37 < 7; fi37++ {
                    v4 -= v4
                    var v38 int16 = v1
                    _ = fi37
                    _ = v38
                }
            }
        case "foo bar":
            var v39 int8 = int8(v5)
            _ = v39
        default:
            v40 := GenS4{F0: 18446744073709551615}
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
