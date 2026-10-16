package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint32, p1 bool, p2 int8) int64 {
    switch 32768 {
    case 127:
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{(^(fi6)), (fi6 % ((fi6) | 1)), (fi6 * 32768)}
            if len(v7) > 0 {
                v7[0] = (fi6 & -1)
            }
            _ = fi6
            _ = v7
        }
    case 527:
        fmt.Println(p2)
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        if p1 {
            p1 = true
            if p1 {
                var v9 int32 = -535
                fmt.Println(p0, v9, v9)
                p0 += p0
                _ = v9
            } else {
                p1 = p1
                v10 := []int{(fi8 & -713), (fi8 + fi8)}
                _ = v10
            }
        }
        _ = fi8
    }
    v11 := GenS1{F0: (p0 << ((65535) & 63)), F1: (p0 ^ 42)}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -1
}
func genFn1() string {
    if true {
        if false {
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS1{F0: 54, F1: 472}
                _ = fi12
                _ = v13
            }
        } else {
            var v14 bool = true
            for fi15 := 0; fi15 < 18; fi15++ {
                for fi16 := 0; fi16 < 8; fi16++ {
                    fmt.Println(fi16)
                    if false {
                        v14 = (v14 && (-639 >= 1))
                    }
                    v14 = (((0.1 + float64(0.0)) + 0.1) != (-2.25 * 3.14159))
                    v17 := []int{(fi16 * -1), fi15, (fi15 ^ fi15)}
                    v14 = ((v14 || (!(2 == 32767))) || (9223372036854775807 <= genFn0(65535, (18446744073709551615 > 42), -52)))
                    _ = fi16
                    _ = v17
                }
                var v18 float64 = -2.25
                v19 := []int{(fi15 << ((250) & 63))}
                _ = fi15
                _ = v18
                _ = v19
            }
            _ = v14
        }
    }
    var v20 int16 = 1
    _ = v20
    return "foo bar"
}

func main() {
    var v1 int16 = -874
    v2 := make(map[string]string)
    v3 := []uint32{2147483647, 1, 2, 358}
    var v4 bool = false
    v5 := make(map[uint32]string)
    var v21 string = (("" + "\n\t") + "foo bar")
    var v22 uint64 = 18446744073709551615
    fmt.Println(v21, v4)
    v21 = (v21 + v21)
    v23 := GenS1{F0: 0, F1: 490}
    if v4 {
        for fi24 := 0; fi24 < 10; fi24++ {
            if v4 {
                if v4 {
                    if ((1e10 - 0.0) < -2.25) {
                        var v25 uint32 = 86
                        v26 := GenS2{F0: (v25 ^ (v25 % ((2147483647) | 1))), F1: 0}
                        fmt.Println(v4, v22)
                        _ = v25
                        _ = v26
                    }
                } else {
                    var v27 int16 = (v1 ^ v1)
                    _ = v27
                }
            }
            v21 = (v21 + (v21 + (v21 + "foo bar")))
            fmt.Println(v22, v4, v22)
            _ = fi24
        }
        v21 = (v21 + v21)
    } else {
        v28 := GenS3{F0: ("foo bar" + ("foo bar" + v21)), F1: (-(v1))}
        _ = len(v2)
        if (132 <= 751) {
            v22 -= v22
            v21 = "\n\t"
        }
        _ = v28
    }
    switch 4294967295 {
    case 855:
        if (!v4) {
            for fi29 := 0; fi29 < 2; fi29++ {
                switch (v21 + "foo bar") {
                case "\n\t":
                    v3 = append(v3, uint32(v22))
                default:
                    var v30 uint32 = 235
                    _ = v30
                }
                var v31 uint64 = 975
                _ = fi29
                _ = v31
            }
        } else {
            if (2147483647 == -1) {
                if (670 < 2) {
                    delete(v5, 1)
                    fmt.Println(v22)
                    fmt.Println(v1, v4)
                } else {
                    switch 42 {
                    case 0:
                        if len(v3) > 0 {
                            v3[0] = 92
                        }
                    }
                    var v32 uint8 = 0
                    v22 -= 1
                    _ = v32
                }
            }
            fmt.Println(v22)
        }
    case -2147483648:
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 bool = v4
            v35 := GenS4{F0: 65535}
            _ = fi33
            _ = v34
            _ = v35
        }
    case -200:
        var v36 int32 = 2147483647
        _ = v36
    default:
        v22 += v22
    }
    switch 9223372036854775807 {
    case 4294967295:
        v4 = true
    }
    fmt.Println(v1, v22, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
