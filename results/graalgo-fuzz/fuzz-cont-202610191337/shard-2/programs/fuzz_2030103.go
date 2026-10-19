package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int64 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
func genFn0(p0 uint64) int {
    if false {
        p0 = p0
        switch 32767 {
        case 32768:
            for fi7 := 0; fi7 < 2; fi7++ {
                var v8 uint16 = 255
                _ = fi7
                _ = v8
            }
        case 755:
            var v9 uint8 = 166
            _ = v9
        case 878:
            v10 := []int{2147483647, 42}
            _ = v10
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            p0 = 4294967295
            v12 := GenS1{F0: 18446744073709551615, F1: p0, F2: -9223372036854775808}
            _ = fi11
            _ = v12
        }
    }
    var v13 uint = 42
    fmt.Println(p0)
    _ = p0
    _ = v13
    return -9223372036854775808
}
func genFn1(p0 bool) uint32 {
    v14 := []int{42, 42, 398}
    p0 = ((-2147483648 != 1) || true)
    _ = p0
    _ = v14
    return 668
}

func main() {
    var v1 float32 = 1.5
    var v2 float32 = (v1 * -2.25)
    var v3 bool = (false && true)
    v4 := []int{-1}
    var v5 uint32 = 1
    var v6 int32 = 772
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = ((193 >= 93) && (1.5 <= 3.14159))
    v1 += v1
    v5 += (v5 / ((v5) | 1))
    if v3 {
        if v3 {
            v2 -= 1.5
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v15 := GenS1{F0: 0, F1: 781, F2: -1}
            _ = v15
        } else {
            v16 := GenS1{F0: 0, F1: 18446744073709551615, F2: 882}
            v6 -= 42
            v17 := GenS1{F0: 2, F1: 10, F2: -502}
            _ = v16
            _ = v17
        }
    }
    fmt.Println(v6, v1, v5)
    v2 += float32(float64(1e10))
    if (2 >= 4294967295) {
        v4 = append(v4, 128)
    }
    switch "a" {
    case "foo bar":
        v3 = true
    case "hello":
        _ = len(v4)
    case "\n\t":
        if (-1 != -2147483648) {
            v18 := GenS1{F0: 88, F1: 70, F2: -9223372036854775808}
            for fi19 := 0; fi19 < 13; fi19++ {
                for fi20 := 0; fi20 < 14; fi20++ {
                    v6 = 2147483647
                    if (!v3) {
                        if (!(!(2 > 126))) {
                            if v3 {
                                if true {
                                    switch fi19 {
                                    case 2147483647:
                                        v1 = (1e10 + 1.5)
                                    case 32767:
                                        v21 := GenS1{F0: 0, F1: 42, F2: int64(fi19)}
                                        _ = v21
                                    default:
                                        v22 := GenS1{F0: 65535, F1: 18446744073709551615, F2: -956}
                                        _ = v22
                                    }
                                    if v3 {
                                        _ = len(v4)
                                    } else {
                                        v1 -= (float32((3.14159 - 0.1)) + -2.25)
                                    }
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                fmt.Println(v3, v3)
                            }
                            v23 := GenS1{F0: 805, F1: 42, F2: -9223372036854775808}
                            v24 := GenS2{F0: "\n\t", F1: 1, F2: 2}
                            _ = v23
                            _ = v24
                        }
                        v5 = 0
                        switch ("" + ("" + "")) {
                        case "a":
                            var v25 bool = v3
                            _ = v25
                        default:
                            if (!((1.5 - ((((v1 + v1) / v1) - (v1 - 1.5)) / (float32((1e10 - (1e10 - 0.0))) / (1e10 + v1)))) == v1)) {
                                if (v5 != (v5 * v5)) {
                                    var v26 int8 = 1
                                    v5 += 1
                                    v4 = append(v4, fi20)
                                    _ = v26
                                }
                                v4 = append(v4, fi19)
                                v2 = 3.14159
                            } else {
                                v1 += float32(3.14159)
                            }
                            v2 = v2
                            fmt.Println(fi20)
                        }
                    }
                    _ = fi20
                }
                var v27 uint16 = 32767
                _ = fi19
                _ = v27
            }
            _ = v18
        }
    }
    fmt.Println(v2)
    if ((v5 >> ((18446744073709551615) & 63)) < v5) {
        v3 = (0.1 >= (float64(0.0) * (1.5 + 0.1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
