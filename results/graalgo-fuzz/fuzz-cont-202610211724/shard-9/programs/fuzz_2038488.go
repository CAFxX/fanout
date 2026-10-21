package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int }
func genFn0() uint16 {
    var v5 int16 = 100
    if false {
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS1{F0: 1, F1: 18446744073709551615, F2: 255}
            _ = fi6
            _ = v7
        }
    }
    switch "hello" {
    case "a":
        if (42 > -128) {
            switch 553 {
            case 2147483647:
                v5 += (v5 | (v5 + v5))
            case 32767:
                var v8 uint = 4294967295
                _ = v8
            case 205:
                v5 += 938
            default:
                v9 := []int{127}
                _ = v9
            }
            v10 := GenS1{F0: -1, F1: 1, F2: 2}
            _ = v10
        } else {
            v11 := GenS2{F0: 1, F1: 2147483647, F2: v5}
            var v12 float64 = -2.25
            _ = v11
            _ = v12
        }
    case "":
        if true {
            var v13 float64 = -2.25
            _ = v13
        } else {
            for fi14 := 0; fi14 < 11; fi14++ {
                if (0 <= 42) {
                    switch "" {
                    case "\n\t":
                        fmt.Println(v5, v5)
                    default:
                        fmt.Println(v5, fi14)
                    }
                    v15 := []int{fi14, fi14, fi14}
                    for fi16 := 0; fi16 < 12; fi16++ {
                        if false {
                            fmt.Println(v5, fi16, fi16)
                        } else {
                            var v17 int32 = -2147483648
                            fmt.Println(fi14)
                            if len(v15) > 0 {
                                v15[0] = (^(fi16))
                            }
                            _ = v17
                        }
                        var v18 float64 = 1.5
                        _ = fi16
                        _ = v18
                    }
                    _ = v15
                }
                v19 := GenS2{F0: 18446744073709551615, F1: 2147483647, F2: 183}
                if true {
                    if true {
                        v5 = v5
                    }
                    var v20 float32 = 1.5
                    _ = v20
                } else {
                    v21 := []int{fi14}
                    _ = v21
                }
                _ = fi14
                _ = v19
            }
            if true {
                var v22 uint32 = 42
                v23 := []int{128}
                _ = v22
                _ = v23
            }
        }
    case "\n\t":
        v5 = v5
    }
    _ = v5
    return 667
}
func genFn1(p0 int64, p1 int) uint8 {
    var v24 int32 = 42
    var v25 uint = 52
    v26 := GenS1{F0: -69, F1: 1, F2: 127}
    p0 = p0
    p1 += (-(p1))
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    _ = v26
    return 1
}

func main() {
    var v1 float64 = 0.1
    var v2 string = "\n\t"
    var v3 uint32 = 1
    v4 := make(map[string]string)
    v27 := []int{128, -525}
    fmt.Println(v3, v2, v3)
    v27 = append(v27, 685)
    v1 += v1
    var v28 int8 = 1
    for fi29 := 0; fi29 < 11; fi29++ {
        v30 := GenS3{F0: int32(v3)}
        if (!(-2147483648 < -2147483648)) {
            var v31 int8 = (-(v28))
            switch fi29 {
            case 2147483647:
                v1 -= (0.1 + ((v1 - 0.0) - v1))
            case 21:
                var v32 string = ("\n\t" + "hello")
                _ = v32
            case 0:
                var v33 uint16 = genFn0()
                _ = v33
            default:
                fmt.Println(v28, v1)
            }
            for fi34 := 0; fi34 < 6; fi34++ {
                fmt.Println(v28, fi29)
                fmt.Println(fi34, v28, v1)
                var v35 uint8 = 128
                _ = fi34
                _ = v35
            }
            _ = v31
        } else {
            fmt.Println(v3, v28)
            v36 := GenS4{F0: 1, F1: (fi29 | fi29)}
            fmt.Println(fi29, v2, v1)
            _ = v36
        }
        var v37 int32 = -2147483648
        _ = fi29
        _ = v30
        _ = v37
    }
    v2 = v2
    delete(v4, "foo bar")
    v2 = ""
    if false {
        for fi38 := 0; fi38 < 20; fi38++ {
            v3 -= v3
            _ = fi38
        }
    }
    v1 = v1
    v39 := GenS2{F0: 2, F1: 42, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v39
    _ = v4
}
