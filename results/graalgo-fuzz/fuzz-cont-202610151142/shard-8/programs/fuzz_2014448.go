package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
func genFn0(p0 bool, p1 int8, p2 int32) uint64 {
    var v7 int32 = p2
    if p0 {
        var v8 uint64 = 865
        fmt.Println(p0, p1)
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        if (false && (800 > 4294967295)) {
            p1 = 100
        }
        _ = fi9
    }
    p2 = 42
    v10 := GenS1{F0: 9223372036854775807, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"a"}
    var v3 string = "\n\t"
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    var v6 uint8 = 255
    if true {
        for fi11 := 0; fi11 < 4; fi11++ {
            switch fi11 {
            case 32768:
                v12 := GenS1{F0: fi11, F1: (v6 + 128)}
                _ = v12
            case 0:
                v6 = 0
            default:
                v13 := GenS2{F0: -2147483648, F1: 518}
                _ = v13
            }
            _ = fi11
        }
        var v14 uint = 401
        _ = v14
    } else {
        _ = len(v2)
        for fi15 := 0; fi15 < 8; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                switch (fi15 ^ (fi16 & (-(fi16)))) {
                case 127:
                    if (true || false) {
                        v3 = ("\n\t" + v3)
                    } else {
                        fmt.Println(fi15, v3)
                        v6 += v6
                    }
                case -559:
                    var v17 uint8 = v6
                    _ = v17
                }
                _ = len(v5)
                var v18 int32 = -769
                _ = fi16
                _ = v18
            }
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := GenS1{F0: fi19, F1: v6}
                if false {
                    var v21 int32 = -2147483648
                    v22 := GenS1{F0: fi15, F1: 0}
                    v23 := GenS1{F0: 42, F1: 1}
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    v3 = "a"
                }
                if ((!((false || (128 >= 32768)) && true)) || true) {
                    var v24 uint16 = 65535
                    _ = v24
                }
                delete(v5, 1)
                if (0 == 1) {
                    if true {
                        var v25 float64 = 0.0
                        var v26 uint8 = v6
                        _ = v25
                        _ = v26
                    }
                }
                v3 = v3
                if (false || (!(-425 > -127))) {
                    v2 = append(v2, (("a" + (v3 + v3)) + (v3 + v3)))
                } else {
                    fmt.Println(v6)
                }
                _ = fi19
                _ = v20
            }
            _ = fi15
        }
        v6 -= 49
    }
    v27 := GenS3{F0: 2147483647}
    if true {
        var v28 uint16 = 977
        var v29 bool = (float64((float32(3.14159) / 1.5)) > 0.1)
        _ = v28
        _ = v29
    } else {
        fmt.Println(v3)
    }
    v6 -= (v6 | 255)
    fmt.Println(v3, v6)
    var v30 int8 = 127
    if false {
        var v31 float32 = -2.25
        _ = v31
    }
    v5[2147483648] = 32768
    var v32 int16 = 163
    v3 = (((v3 + "\n\t") + v3) + v3)
    if false {
        var v33 int32 = -2147483648
        _ = v33
    }
    var v34 uint8 = 255
    v35 := GenS4{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
