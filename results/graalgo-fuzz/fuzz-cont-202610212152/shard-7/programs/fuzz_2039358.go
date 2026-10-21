package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
func genFn0(p0 int16, p1 int16, p2 int) int16 {
    fmt.Println(p1, p0, p0)
    v6 := GenS1{F0: (!(1 >= 2)), F1: 4294967295, F2: 802}
    for fi7 := 0; fi7 < 16; fi7++ {
        if (!true) {
            if (float64(0.1) >= 3.14159) {
                var v8 uint8 = 1
                _ = v8
            }
            v9 := []int{170, fi7, (p2 << ((447) & 63))}
            _ = v9
        } else {
            var v10 int = 128
            _ = v10
        }
        p1 = p1
        if false {
            p2 -= (p2 & fi7)
            var v11 int16 = (p0 & p1)
            _ = v11
        }
        _ = fi7
    }
    if true {
        var v12 uint = 0
        var v13 uint16 = 65535
        var v14 float64 = float64(0.1)
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 & p1)
}
func genFn1(p0 uint16, p1 uint16, p2 uint32) string {
    var v15 int32 = 1
    p2 += (p2 & p2)
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 bool = true
        v18 := []int{fi16}
        for fi19 := 0; fi19 < 19; fi19++ {
            v17 = (1e10 != -2.25)
            _ = fi19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    if (true && (326 == 942)) {
        var v20 uint8 = 225
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return ""
}

func main() {
    var v1 uint8 = 42
    v2 := []int{0}
    var v3 float32 = 1e10
    var v4 string = ("foo bar" + ("a" + "hello"))
    var v5 uint16 = 2
    var v21 float32 = -2.25
    v2 = append(v2, 127)
    _ = len(v2)
    v22 := GenS1{F0: true, F1: 4294967295, F2: (v5 / (((v5 - 255)) | 1))}
    v4 = v4
    v21 -= (((-2.25 / float32(0.1)) * v21) / v21)
    var v23 float64 = 1e10
    v24 := GenS1{F0: (v4 < v4), F1: 122, F2: v5}
    _ = len(v2)
    v2 = append(v2, 32768)
    v5 = 0
    v2 = append(v2, int(v1))
    switch "a" {
    case "hello":
        if (18446744073709551615 == 581) {
            if ((-2147483648 != 0) && false) {
                if (-1 >= 0) {
                    if (!(-2147483648 > 1)) {
                        var v25 uint64 = 116
                        _ = v25
                    } else {
                        if false {
                            _ = len(v2)
                        } else {
                            var v26 bool = true
                            v2 = append(v2, 368)
                            var v27 string = v4
                            _ = v26
                            _ = v27
                        }
                    }
                    var v28 float64 = ((-2.25 - v23) / v23)
                    _ = v28
                } else {
                    fmt.Println(v4, v3, v23)
                    var v29 bool = true
                    _ = v29
                }
            }
        }
    }
    v30 := GenS1{F0: true, F1: 2147483648, F2: v5}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
