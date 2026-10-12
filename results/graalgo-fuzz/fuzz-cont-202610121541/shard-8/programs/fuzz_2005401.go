package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int8; F2 int32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint8, p1 int8) uint16 {
    fmt.Println(p1)
    p0 += p0
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 0
    v3 := []string{"a", "", ("foo bar" + ("\n\t" + "")), ""}
    var v4 float64 = float64(((0.1 - -2.25) * 1.5))
    var v5 int = -168
    if (1 >= 95) {
        for fi6 := 0; fi6 < 4; fi6++ {
            v2 = (^(v2))
            var v7 uint8 = 128
            v8 := GenS1{F0: (-(fi6)), F1: v7}
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            v5 = fi9
            _ = fi9
        }
        if (!(-123 > 42)) {
            var v10 float32 = 0.1
            _ = v10
        }
    } else {
        v11 := v1["hello"]
        fmt.Println(v4, v2)
        if false {
            if ((v2 << ((930) & 63)) >= 145) {
                if false {
                    v12 := GenS1{F0: v5, F1: 104}
                    v13 := GenS2{F0: 4294967295, F1: 100, F2: 0}
                    _ = v12
                    _ = v13
                } else {
                    switch (v5 * v5) {
                    case 2147483647:
                        v11 += v11
                    }
                    v14 := GenS3{F0: v2, F1: v11}
                    _ = len(v1)
                    _ = v14
                }
                fmt.Println(v5, v2)
            }
        } else {
            var v15 string = "\n\t"
            var v16 float64 = (v4 + 0.1)
            v16 += (v16 / (((((float64(3.14159) * -2.25) / float64(0.1)) / (((v4 - v16) - 1.5) + (v4 + v16))) + ((v16 * (1.5 - -2.25)) + 0.1)) + v4))
            _ = v15
            _ = v16
        }
        _ = v11
    }
    v3 = append(v3, "a")
    fmt.Println(v4)
    fmt.Println(v4)
    if true {
        v4 += v4
        v2 = v2
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v17 := GenS4{F0: 165, F1: (v5 % ((v5) | 1))}
        v18 := GenS5{F0: 255, F1: 0}
        _ = v17
        _ = v18
    }
    fmt.Println(v2)
    if true {
        fmt.Println(v4, v4, v5)
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 uint32 = 4294967295
            _ = fi19
            _ = v20
        }
    }
    if true {
        v21 := v1["\n\t"]
        _ = v21
    }
    fmt.Println(v2)
    var v22 bool = true
    v23 := GenS3{F0: v2, F1: 684}
    var v24 uint32 = 0
    v25 := GenS4{F0: 2, F1: (v5 + (v5 - v5))}
    fmt.Println(v2, v22, v5)
    var v26 uint32 = v24
    v27 := GenS2{F0: 747, F1: -4, F2: -2147483648}
    var v28 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
