package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
func genFn0(p0 uint, p1 bool) bool {
    if (p1 || (p1 || p1)) {
        p0 += (p0 + 332)
        var v7 uint8 = 1
        _ = v7
    } else {
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    p1 = ((^(p0)) <= (p0 - (p0 ^ p0)))
    if (2 >= 799) {
        for fi9 := 0; fi9 < 8; fi9++ {
            if (128 < 46) {
                p1 = p1
                if p1 {
                    v10 := GenS1{F0: 979}
                    v11 := []int{128, fi9}
                    switch "\n\t" {
                    case "":
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    case "\n\t":
                        fmt.Println(fi9, p1, p0)
                    case "foo bar":
                        p0 = (p0 >> ((p0) & 63))
                    default:
                        p1 = (p1 && true)
                    }
                    _ = v10
                    _ = v11
                }
                v12 := GenS1{F0: 65535}
                _ = v12
            } else {
                var v13 float64 = 0.0
                _ = v13
            }
            _ = fi9
        }
    } else {
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS2{F0: 127, F1: 109, F2: 32767}
            v16 := GenS2{F0: 100, F1: 0, F2: 2}
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 uint16 = 1
        var v18 bool = p1
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    return (!(((p0 | 0) == 255) || (!p1)))
}

func main() {
    v1 := []uint32{895}
    var v2 float64 = 0.0
    var v3 uint32 = 2147483647
    var v4 float32 = (0.0 * (float32(v2) - float32((v2 - (-2.25 + float64((((1.5 * ((((0.1 - (1e10 + 3.14159)) * 1.5) + (3.14159 / (3.14159 / -2.25))) - -2.25)) + (0.1 / float32(float64(1e10)))) * 1.5)))))))
    v5 := []string{("\n\t" + "hello"), "foo bar", "foo bar"}
    v6 := make(map[int]int)
    fmt.Println(v4, v2, v4)
    for fi19 := 0; fi19 < 10; fi19++ {
        if ((3.14159 + 3.14159) != v2) {
            for fi20 := 0; fi20 < 14; fi20++ {
                fmt.Println(fi19, v4)
                _ = fi20
            }
        } else {
            _ = len(v6)
        }
        var v21 uint8 = 68
        _ = fi19
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        fmt.Println(fi22, v3, v2)
        for fi23 := 0; fi23 < 5; fi23++ {
            if false {
                if (fi23 == 0) {
                    var v24 bool = false
                    _ = v24
                }
                if true {
                    switch fi23 {
                    case 4294967295:
                        var v25 int16 = 42
                        _ = v25
                    case -720:
                        v3 = v3
                    }
                } else {
                    var v26 int8 = 100
                    v26 = (v26 ^ (v26 & (v26 & (v26 >> ((39) & 63)))))
                    _ = v26
                }
            } else {
                v27 := v6[128]
                _ = v27
            }
            _ = fi23
        }
        _ = fi22
    }
    if false {
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 string = "foo bar"
            _ = fi28
            _ = v29
        }
    } else {
        var v30 float64 = ((v2 - (v2 / 1.5)) * 1e10)
        v3 += 790
        _ = v30
    }
    v31 := GenS1{F0: v3}
    var v32 uint8 = 117
    var v33 int8 = -128
    var v34 bool = true
    fmt.Println(v33, v33)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
