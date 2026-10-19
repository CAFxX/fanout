package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint32 }
func genFn0(p0 int64, p1 int, p2 uint) uint32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        p1 = 0
        _ = fi7
    }
    var v8 float64 = float64(float32(float64(1e10)))
    switch (p1 - (p1 * (p1 + (-(p1))))) {
    case 128:
        fmt.Println(p0, p0, p2)
    case 42:
        if false {
            v9 := []int{p1, p1, 42}
            _ = v9
        } else {
            var v10 int64 = (p0 & -1)
            v10 = -9223372036854775808
            var v11 int = (p1 + 720)
            _ = v10
            _ = v11
        }
    case 9223372036854775807:
        v12 := []int{(p1 - (p1 >> ((0) & 63))), 42, p1}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 4294967295
}
func genFn1() int8 {
    var v13 int8 = 123
    fmt.Println(v13)
    fmt.Println(v13)
    switch 32767 {
    case 32768:
        v14 := []int{}
        _ = v14
    case 139:
        var v15 uint = 764
        _ = v15
    }
    fmt.Println(v13)
    _ = v13
    return v13
}

func main() {
    var v1 uint = 255
    v2 := []string{("" + "hello"), "foo bar", (("hello" + "foo bar") + "foo bar")}
    v3 := []string{"foo bar"}
    var v4 int64 = 1
    v5 := make(map[int]string)
    v6 := make(map[string]bool)
    v16 := GenS1{F0: 395}
    var v17 int64 = v4
    for fi18 := 0; fi18 < 11; fi18++ {
        v4 = (v17 + (v17 & v4))
        _ = fi18
    }
    v19 := GenS2{F0: (^(v1)), F1: 1, F2: 680}
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            v2 = append(v2, "hello")
            v17 = v17
            _ = fi21
        }
        _ = fi20
    }
    fmt.Println(v1, v17)
    for fi22 := 0; fi22 < 8; fi22++ {
        v1 = v1
        if true {
            if true {
                v17 = v4
                switch "a" {
                case "foo bar":
                    v3 = append(v3, "hello")
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                default:
                    v23 := GenS1{F0: 0}
                    _ = v23
                }
            } else {
                var v24 float32 = (1.5 / 1.0)
                _ = v24
            }
            var v25 int16 = -1
            fmt.Println(fi22, fi22)
            _ = v25
        } else {
            if (true && false) {
                v1 -= (v1 ^ (v1 / ((779) | 1)))
                if false {
                    v26 := GenS1{F0: 817}
                    fmt.Println(v1)
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v1 += (v1 / (((v1 * 42)) | 1))
                        fmt.Println(fi27)
                        if (128 < fi22) {
                            v17 = (-(v17))
                            var v28 int16 = 0
                            _ = v28
                        }
                        _ = fi27
                    }
                    v4 = v17
                }
                v1 = v1
            } else {
                var v29 uint8 = 2
                _ = len(v5)
                fmt.Println(v17, v29)
                _ = v29
            }
            for fi30 := 0; fi30 < 5; fi30++ {
                switch -957 {
                case -9223372036854775808:
                    var v31 uint32 = genFn0((v17 >> (((v1 * (v1 + (v1 % ((3) | 1))))) & 63)), (fi22 % ((fi30) | 1)), v1)
                    _ = v31
                default:
                    if (255 != 46) {
                        switch (fi30 - (^(fi22))) {
                        case 2147483647:
                            v1 += 255
                        }
                        fmt.Println(v17)
                    }
                }
                var v32 int8 = -128
                _ = fi30
                _ = v32
            }
        }
        delete(v5, (^(fi22)))
        _ = fi22
    }
    v4 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
