package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 int }
func genFn0(p0 uint8) bool {
    fmt.Println(p0)
    if (((!true) || (3.14159 == 0.1)) && false) {
        p0 = p0
    } else {
        v4 := GenS1{F0: (p0 | p0), F1: 4294967295}
        var v5 int8 = -1
        var v6 float32 = (3.14159 - ((1e10 * ((float32((0.0 + 0.1)) * (0.1 * (0.1 + float32((((float64(1e10) - 1.5) / (-2.25 * 0.1)) / -2.25))))) - 0.0)) + float32(-2.25)))
        _ = v4
        _ = v5
        _ = v6
    }
    _ = p0
    return true
}

func main() {
    var v1 int16 = -1
    var v2 uint8 = 127
    var v3 int32 = -751
    fmt.Println(v3, v1, v2)
    fmt.Println(v2)
    v7 := GenS1{F0: v2, F1: 684}
    var v8 uint16 = 255
    var v9 int32 = (v3 - v3)
    v1 = v1
    switch 213 {
    case 127:
        switch "foo bar" {
        case "\n\t":
            var v10 uint64 = 18446744073709551615
            _ = v10
        }
    case -855:
        fmt.Println(v2)
    }
    var v11 bool = ((-(v1)) > 1)
    fmt.Println(v9, v1)
    for fi12 := 0; fi12 < 1; fi12++ {
        v11 = (32767 < (v1 & v1))
        v1 = (v1 / ((v1) | 1))
        if true {
            v13 := []int{(fi12 | 32767), -16, 0, fi12}
            v1 -= 32767
            _ = v13
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int64 = -9223372036854775808
        v9 += (v3 - v3)
        if v11 {
            v15 = (-(v15))
            var v16 int32 = v3
            v17 := []int{}
            _ = v16
            _ = v17
        } else {
            v18 := []int{436, fi14}
            if (!false) {
                v19 := GenS1{F0: (v2 ^ v2), F1: 422}
                _ = len(v18)
                v1 = v1
                _ = v19
            }
            var v20 int32 = 42
            _ = v18
            _ = v20
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v9)
    var v21 int8 = 1
    switch -713 {
    case 32767:
        if v11 {
            var v22 uint8 = v2
            switch ("" + "a") {
            case "hello":
                if ((("\n\t" + "a") + "foo bar") > "a") {
                    v23 := GenS1{F0: (v2 >> ((4294967295) & 63)), F1: 42}
                    _ = v23
                } else {
                    if (!(269 > 1)) {
                        for fi24 := 0; fi24 < 5; fi24++ {
                            v25 := []int{fi24}
                            v2 = (v2 % ((1) | 1))
                            _ = fi24
                            _ = v25
                        }
                    } else {
                        v26 := []int{-260, 2147483647, 1, 0}
                        _ = v26
                    }
                    if false {
                        v27 := GenS2{F0: 65535, F1: 4294967295}
                        _ = v27
                    } else {
                        var v28 int32 = int32(v2)
                        v11 = (-280 != 4294967295)
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 3; fi29++ {
                        fmt.Println(v2, v11, fi29)
                        var v30 uint32 = 0
                        _ = fi29
                        _ = v30
                    }
                }
            default:
                if v11 {
                    v31 := []int{}
                    _ = v31
                } else {
                    v1 = (-(v1))
                    for fi32 := 0; fi32 < 9; fi32++ {
                        var v33 uint16 = v8
                        _ = fi32
                        _ = v33
                    }
                    if (v2 > (v22 >> ((42) & 63))) {
                        if ((1 > -172) && v11) {
                            var v34 float64 = (1e10 * (-2.25 / 1e10))
                            v35 := []int{-916}
                            _ = v34
                            _ = v35
                        }
                        v36 := []int{956, 128, -922}
                        var v37 int64 = int64(v21)
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int16 = (^(v1))
                        _ = v38
                    }
                }
            }
            _ = v22
        }
    case 42:
        switch "hello" {
        case "a":
            v39 := []int{32768, 42, 2147483647, 981}
            _ = v39
        case "hello":
            v3 -= 2147483647
        case "foo bar":
            var v40 int64 = 42
            _ = v40
        }
    default:
        var v41 int32 = (v9 % (((v9 >> ((373) & 63))) | 1))
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
