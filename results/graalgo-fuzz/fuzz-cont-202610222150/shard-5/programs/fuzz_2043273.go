package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint; F2 int16 }
func genFn0(p0 uint8, p1 uint16, p2 bool) int {
    var v4 int8 = -128
    v5 := []int{2147483647, 32768}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}
func genFn1(p0 int8, p1 int8) int8 {
    var v6 uint16 = 255
    if (653 > 2147483647) {
        for fi7 := 0; fi7 < 17; fi7++ {
            if (1.5 > ((((0.0 * 3.14159) * float32((3.14159 / 1.0))) / ((0.0 - (-2.25 / 1.0)) + float32(3.14159))) / 0.1)) {
                var v8 int16 = 42
                _ = v8
            } else {
                v6 -= v6
                if false {
                    v9 := []int{(^(fi7)), (fi7 / ((fi7) | 1)), (fi7 - (fi7 - (fi7 ^ fi7))), -388}
                    fmt.Println(p0, v6, p0)
                    _ = v9
                }
                var v10 int8 = p1
                _ = v10
            }
            _ = fi7
        }
        p0 -= 0
    } else {
        switch 0 {
        case 32767:
            if (32767 != -32768) {
                for fi11 := 0; fi11 < 11; fi11++ {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        v13 := []int{fi12, (fi12 - genFn0(255, (v6 - v6), (int32(fi11) <= 1))), (fi12 / (((-(fi11))) | 1)), fi12}
                        v6 = v6
                        var v14 uint8 = 1
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    p0 = 127
                    _ = fi11
                }
                var v15 uint64 = 18446744073709551615
                _ = v15
            } else {
                switch -9223372036854775808 {
                case -966:
                    p1 = (p1 >> ((4294967295) & 63))
                case 9223372036854775807:
                    v16 := []int{}
                    _ = v16
                default:
                    var v17 int32 = 394
                    _ = v17
                }
            }
        case 9223372036854775807:
            for fi18 := 0; fi18 < 12; fi18++ {
                if (!false) {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        var v20 uint8 = 2
                        _ = fi19
                        _ = v20
                    }
                    v21 := GenS1{F0: (v6 % (((v6 + (v6 % ((v6) | 1)))) | 1))}
                    _ = v21
                } else {
                    v22 := []int{(fi18 & (fi18 % ((fi18) | 1))), (fi18 << ((144) & 63)), 9223372036854775807, (-(fi18))}
                    switch genFn0(0, (v6 * (v6 << ((42) & 63))), ((!(false || false)) && true)) {
                    case -9223372036854775808:
                        v23 := GenS2{F0: -1}
                        _ = v23
                    case -2147483648:
                        if ("foo bar" > "hello") {
                            for fi24 := 0; fi24 < 6; fi24++ {
                                p1 = (p0 & (p0 % ((p0) | 1)))
                                v6 = v6
                                p0 = p0
                                _ = fi24
                            }
                        } else {
                            for fi25 := 0; fi25 < 1; fi25++ {
                                if false {
                                    fmt.Println(fi18, p1, fi25)
                                }
                                p1 -= (p1 + p1)
                                v6 = (v6 | v6)
                                _ = fi25
                            }
                        }
                    case 9223372036854775807:
                        var v26 int8 = (-(p1))
                        _ = v26
                    default:
                        var v27 uint16 = v6
                        _ = v27
                    }
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    _ = v22
                }
                var v28 int = (fi18 + (fi18 >> ((2) & 63)))
                _ = fi18
                _ = v28
            }
        }
    }
    v29 := GenS2{F0: -1}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v29
    _ = v6
    return p1
}

func main() {
    v1 := []uint32{42, 2147483648}
    v2 := []uint32{244}
    var v3 uint64 = 1
    var v30 string = "\n\t"
    v30 = "foo bar"
    var v31 bool = (-128 == 1)
    v32 := GenS1{F0: uint16(v3)}
    for fi33 := 0; fi33 < 8; fi33++ {
        for fi34 := 0; fi34 < 13; fi34++ {
            fmt.Println(v31, v30, v3)
            var v35 uint64 = 0
            v1 = append(v1, 0)
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v37 string = v30
        _ = fi36
        _ = v37
    }
    fmt.Println(v31, v3)
    fmt.Println(v31, v30, v3)
    var v38 float64 = 1e10
    v39 := GenS3{F0: v31, F1: 558, F2: 42}
    var v40 uint64 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v40
}
