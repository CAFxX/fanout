package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0(p0 uint32) int8 {
    if (145 != 1) {
        fmt.Println(p0)
        p0 += 1
    } else {
        p0 -= (p0 & 718)
        v4 := GenS1{F0: 818, F1: -70, F2: 0}
        if (111 >= 1) {
            fmt.Println(p0)
            var v5 int16 = -378
            v6 := []int{478}
            _ = v5
            _ = v6
        }
        _ = v4
    }
    v7 := []int{-1, -2147483648, 42, 1}
    _ = p0
    _ = v7
    return 127
}

func main() {
    var v1 bool = (9223372036854775807 < 9223372036854775807)
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    fmt.Println(v1)
    v1 = (float64(3.14159) != -2.25)
    var v8 float32 = 1e10
    var v9 uint8 = 2
    v10 := GenS2{F0: 516, F1: 65535, F2: 42}
    switch "\n\t" {
    case "foo bar":
        var v11 int = 917
        _ = v11
    }
    v12 := GenS2{F0: 8, F1: 42, F2: 1}
    if v1 {
        v13 := GenS2{F0: 141, F1: 2, F2: 768}
        switch 519 {
        case 985:
            if v1 {
                for fi14 := 0; fi14 < 20; fi14++ {
                    v9 = v9
                    if v1 {
                        if (!v1) {
                            var v15 uint8 = v9
                            v16 := GenS1{F0: 32767, F1: genFn0(2), F2: v15}
                            _ = v15
                            _ = v16
                        }
                        v17 := []int{fi14}
                        var v18 uint8 = (v9 & 42)
                        _ = v17
                        _ = v18
                    } else {
                        var v19 int64 = -9223372036854775808
                        fmt.Println(fi14, v8, v8)
                        if (4294967295 >= fi14) {
                            switch "a" {
                            case "a":
                                var v20 uint = 4294967295
                                _ = v20
                            default:
                                var v21 int64 = v19
                                _ = v21
                            }
                            for fi22 := 0; fi22 < 7; fi22++ {
                                var v23 float64 = -2.25
                                v24 := []int{682, fi22, fi22}
                                if false {
                                    var v25 float32 = (v8 - v8)
                                    var v26 uint32 = 178
                                    if len(v24) > 0 {
                                        v24[0] = (fi22 + 386)
                                    }
                                    _ = v25
                                    _ = v26
                                } else {
                                    v27 := GenS2{F0: 2147483647, F1: 255, F2: 32767}
                                    v1 = (221 == v9)
                                    _ = v27
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            fmt.Println(v8)
                        } else {
                            v1 = ((v19 * (v19 - (v19 - 42))) < v19)
                        }
                        _ = v19
                    }
                    _ = fi14
                }
                var v28 bool = (681 == 946)
                _ = v28
            } else {
                var v29 uint64 = 18446744073709551615
                fmt.Println(v1, v1, v9)
                _ = v29
            }
        }
        v30 := []int{42, 32767, 0}
        _ = v13
        _ = v30
    }
    v9 = v9
    v31 := GenS1{F0: -1, F1: -128, F2: v9}
    for fi32 := 0; fi32 < 9; fi32++ {
        for fi33 := 0; fi33 < 3; fi33++ {
            var v34 string = "a"
            if (0.0 <= ((((float64((((-2.25 + v8) * (float32((3.14159 * ((-2.25 / ((0.0 * 3.14159) - float64(1e10))) * 1.5))) / ((v8 * (v8 / ((0.0 + 0.1) * 3.14159))) / v8))) * float32(0.1))) + 1e10) - 3.14159) + -2.25) - 1e10)) {
                if v1 {
                    var v35 uint32 = 1
                    var v36 int32 = 2147483647
                    _ = v35
                    _ = v36
                } else {
                    var v37 uint = 255
                    _ = v37
                }
            }
            _ = fi33
            _ = v34
        }
        for fi38 := 0; fi38 < 3; fi38++ {
            switch ("a" + "a") {
            case "foo bar":
                fmt.Println(v1)
            case "\n\t":
                fmt.Println(v9, v9, v9)
            }
            _ = fi38
        }
        _ = fi32
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        v9 += 2
        _ = fi39
    }
    v40 := v3[32767]
    var v41 string = "\n\t"
    var v42 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v40
    _ = v41
    _ = v42
    _ = v8
    _ = v9
}
