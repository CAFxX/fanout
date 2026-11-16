package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint8 }
func genFn0() uint32 {
    var v7 int64 = 9223372036854775807
    v8 := []int{42, -9223372036854775808}
    _ = v7
    _ = v8
    return 726
}
func genFn1(p0 int64, p1 int32, p2 int32) int {
    v9 := GenS1{F0: genFn0(), F1: 0}
    var v10 int8 = 1
    p2 = p1
    if (96 != 65) {
        if false {
            p1 += p1
            if (!(!((p0 >> ((525) & 63)) != 42))) {
                v11 := []int{128, 4294967295}
                if ("\n\t" < ((("a" + "\n\t") + "") + (("a" + "hello") + "\n\t"))) {
                    v12 := GenS2{F0: 4294967295}
                    fmt.Println(p0, p0)
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 3; fi13++ {
                        for fi14 := 0; fi14 < 14; fi14++ {
                            p0 -= (-(p0))
                            var v15 int64 = p0
                            _ = fi14
                            _ = v15
                        }
                        _ = fi13
                    }
                    var v16 uint8 = 255
                    _ = v16
                }
                _ = v11
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -128
    var v3 uint = 18446744073709551615
    var v4 int16 = 24
    var v5 int32 = -2147483648
    v6 := make(map[uint32]int)
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 int16 = (v4 % ((int16(v5)) | 1))
        _ = fi17
        _ = v18
    }
    var v19 uint64 = 460
    switch -10 {
    case -701:
        switch "" {
        case "a":
            v20 := GenS3{F0: (v4 - v4), F1: 128}
            _ = v20
        default:
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 int64 = 1
                fmt.Println(v2)
                _ = fi21
                _ = v22
            }
        }
    case 2147483647:
        if (genFn1(0, (v5 >> ((v3) & 63)), (v5 & 42)) != -1) {
            v2 = v2
            v23 := []int{127, genFn1(-9223372036854775808, (v5 + (v5 % ((v5) | 1))), v5)}
            _ = v23
        } else {
            var v24 float64 = float64(-2.25)
            _ = len(v1)
            _ = v24
        }
    }
    var v25 int8 = v2
    delete(v6, 0)
    var v26 bool = ((2147483648 > 654) && (1 == 1))
    v19 = v19
    var v27 uint32 = 2147483648
    if v26 {
        fmt.Println(v5, v4, v5)
    }
    v28 := GenS3{F0: (v4 - v4), F1: 110}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
