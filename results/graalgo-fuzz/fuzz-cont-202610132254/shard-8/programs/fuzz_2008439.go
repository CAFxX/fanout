package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int }
func genFn0(p0 uint32, p1 bool, p2 int64) int {
    var v4 uint32 = 2
    switch ("hello" + "\n\t") {
    case "foo bar":
        if p1 {
            var v5 uint64 = 18446744073709551615
            if (p1 || (251 >= 142)) {
                v6 := []int{-177, -9223372036854775808}
                for fi7 := 0; fi7 < 3; fi7++ {
                    var v8 uint16 = 597
                    _ = fi7
                    _ = v8
                }
                _ = v6
            } else {
                v4 -= (p0 | 2147483647)
            }
            if ((0 <= 32767) && ((!p1) || (0 != 51))) {
                if (!(-38 <= 2147483647)) {
                    if ((p1 && (3.14159 >= -2.25)) || (!p1)) {
                        v4 -= (^(p0))
                    }
                }
            }
            _ = v5
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                p0 = (p0 / ((2147483647) | 1))
                _ = fi9
            }
            v4 -= 2147483647
            p1 = (p1 && (true || p1))
        }
    default:
        var v10 string = ""
        _ = v10
    }
    v11 := []int{0, 127}
    var v12 int = 128
    p1 = (!(677 >= 42))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v4
    return v12
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint32 = v1
    v3 := []uint32{v1}
    var v13 uint = 4294967295
    var v14 string = ""
    switch v14 {
    case "a":
        if false {
            fmt.Println(v2, v13)
            var v15 uint64 = 1
            _ = v15
        } else {
            fmt.Println(v14, v14, v13)
            _ = len(v3)
            var v16 int = 128
            _ = v16
        }
    case "foo bar":
        var v17 uint64 = 0
        _ = v17
    }
    v1 += (v1 * (v2 * (v2 + (v1 ^ (v1 | (v1 % (((v2 & v1)) | 1)))))))
    v18 := GenS1{F0: 674, F1: (^(v1))}
    for fi19 := 0; fi19 < 7; fi19++ {
        v2 = 931
        if len(v3) > 0 {
            v3[0] = (^(v2))
        }
        v3 = append(v3, v1)
        _ = fi19
    }
    v14 = v14
    _ = len(v3)
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            if (!false) {
                var v22 int32 = 1
                if (float64(((-2.25 - 0.0) + -2.25)) > 3.14159) {
                    var v23 uint8 = 227
                    _ = v23
                }
                _ = v22
            } else {
                v2 += v2
                var v24 float32 = 3.14159
                _ = v24
            }
            _ = len(v3)
            if false {
                if false {
                    v25 := GenS1{F0: 32767, F1: v1}
                    if false {
                        v3 = append(v3, 2147483648)
                    }
                    v26 := GenS2{F0: (fi20 - (fi20 ^ (fi21 / ((-2147483648) | 1))))}
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        switch "" {
                        case "foo bar":
                            v2 += v1
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case "hello":
                        v2 += v1
                    default:
                        var v27 float32 = 3.14159
                        _ = v27
                    }
                    _ = v25
                    _ = v26
                } else {
                    var v28 float32 = 1e10
                    v28 -= v28
                    _ = v28
                }
                var v29 string = v14
                _ = v29
            }
            if true {
                v30 := GenS2{F0: genFn0(2147483648, false, 643)}
                _ = v30
            }
            if true {
                if true {
                    if false {
                        v1 -= (^(v1))
                    } else {
                        var v31 int16 = 0
                        var v32 int32 = -2147483648
                        v3 = append(v3, (v2 % (((v1 ^ (v2 * (v1 >> (((^(v13))) & 63))))) | 1)))
                        var v33 uint8 = 128
                        v3 = append(v3, v2)
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                }
            } else {
                if true {
                    var v34 uint64 = 572
                    _ = v34
                } else {
                    fmt.Println(fi20)
                }
                v35 := GenS1{F0: 0, F1: 665}
                _ = v35
            }
            if len(v3) > 0 {
                v3[0] = v2
            }
            _ = fi21
        }
        _ = fi20
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        var v37 uint16 = 32767
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
}
