package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint32) uint64 {
    p0 += (p0 >> ((uint(p0)) & 63))
    v5 := []int{0, 32768, 341}
    p0 = p0
    p0 -= (p0 >> ((2) & 63))
    _ = p0
    _ = v5
    return 330
}
func genFn1(p0 uint16, p1 bool) bool {
    v6 := []int{451}
    switch 128 {
    case -866:
        if ((1 > 4294967295) || (genFn0(605) < 178)) {
            p0 = (p0 & p0)
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint = 2
                p0 -= 760
                _ = fi7
                _ = v8
            }
            if (42 <= 2) {
                v9 := GenS1{F0: -9223372036854775808, F1: (4294967295 <= 714)}
                p0 += 42
                v6 = append(v6, 0)
                _ = v9
            } else {
                if (!(-2.25 <= -2.25)) {
                    if (p0 != 255) {
                        var v10 uint = 786
                        if (true && (-1 >= -2147483648)) {
                            if (!(!p1)) {
                                switch 2147483647 {
                                case 986:
                                    var v11 bool = p1
                                    _ = v11
                                case -385:
                                    p1 = true
                                }
                                fmt.Println(p1, v10)
                            }
                            if p1 {
                                for fi12 := 0; fi12 < 2; fi12++ {
                                    for fi13 := 0; fi13 < 18; fi13++ {
                                        p1 = (0.0 != ((-2.25 * (3.14159 * 0.0)) / 1.0))
                                        _ = fi13
                                    }
                                    v14 := GenS1{F0: 4294967295, F1: (p1 || (714 <= 9223372036854775807))}
                                    for fi15 := 0; fi15 < 3; fi15++ {
                                        switch -9223372036854775808 {
                                        case -1:
                                            _ = len(v6)
                                        case 32767:
                                            var v16 uint64 = 0
                                            _ = v16
                                        case 127:
                                            var v17 int = 550
                                            _ = v17
                                        default:
                                            if len(v6) > 0 {
                                                v6[0] = (fi12 << (((v10 << ((0) & 63))) & 63))
                                            }
                                            var v18 float32 = 3.14159
                                            fmt.Println(p1, p1, fi12)
                                            _ = v18
                                        }
                                        _ = fi15
                                    }
                                    _ = fi12
                                    _ = v14
                                }
                                var v19 int32 = 213
                                _ = v19
                            } else {
                                if true {
                                    p0 += p0
                                    var v20 float64 = 1e10
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v20
                                }
                                var v21 float32 = 3.14159
                                v10 += 42
                                _ = v21
                            }
                            v10 = 42
                        } else {
                            v22 := GenS1{F0: -1, F1: p1}
                            _ = v22
                        }
                        v23 := GenS2{F0: (false && (!(-60 >= 42))), F1: -952, F2: (p0 % ((p0) | 1))}
                        _ = v10
                        _ = v23
                    }
                }
            }
        } else {
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 int16 = 1
                var v26 int16 = 140
                var v27 int32 = -118
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        }
    case -777:
        var v28 int16 = 32767
        _ = v28
    default:
        fmt.Println(p0, p0)
    }
    p0 = p0
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    return (!(!(p1 || (-15 <= -334))))
}

func main() {
    var v1 uint64 = 1
    v2 := []string{}
    v3 := make(map[string]int)
    v4 := []uint32{65535, 2147483647, 453, 659}
    var v29 float32 = -2.25
    fmt.Println(v1)
    var v30 string = "hello"
    var v31 int8 = -1
    switch -834 {
    case 0:
        v31 = (^(v31))
    default:
        v1 += 0
    }
    if false {
        var v32 uint64 = (v1 * v1)
        switch -2147483648 {
        case 32768:
            for fi33 := 0; fi33 < 2; fi33++ {
                v30 = (((((v30 + ("a" + "")) + (v30 + "hello")) + ((v30 + (v30 + (v30 + ("" + (v30 + v30))))) + ("foo bar" + (v30 + "hello")))) + ((v30 + v30) + v30)) + "a")
                v32 -= v1
                _ = fi33
            }
        case -9223372036854775808:
            v29 += (1.5 + float32(0.1))
        }
        if true {
            for fi34 := 0; fi34 < 20; fi34++ {
                var v35 int16 = 595
                v36 := GenS1{F0: fi34, F1: false}
                _ = fi34
                _ = v35
                _ = v36
            }
            v37 := GenS2{F0: (!(!true)), F1: -462, F2: 1}
            _ = v37
        }
        _ = v32
    }
    var v38 string = v30
    for fi39 := 0; fi39 < 10; fi39++ {
        v2 = append(v2, "a")
        if true {
            v40 := v3["foo bar"]
            v29 += v29
            var v41 uint8 = 127
            _ = v40
            _ = v41
        } else {
            v42 := GenS3{F0: 0}
            var v43 uint = 853
            _ = v42
            _ = v43
        }
        _ = fi39
    }
    delete(v3, v38)
    v29 = v29
    fmt.Println(v1, v29)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
}
