package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint) int {
    if (!true) {
        for fi4 := 0; fi4 < 18; fi4++ {
            p0 = (p0 & p0)
            if false {
                var v5 int64 = (p0 >> ((p1) & 63))
                _ = v5
            }
            if false {
                fmt.Println(p1, p1, fi4)
            }
            _ = fi4
        }
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 int64 = (p0 * p0)
            _ = fi6
            _ = v7
        }
        var v8 int32 = -178
        _ = v8
    } else {
        v9 := GenS1{F0: true, F1: 42}
        p1 += (p1 ^ 368)
        _ = v9
    }
    p0 -= p0
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 string = ""
    var v2 int64 = 9223372036854775807
    var v3 string = v1
    if ((-1 == 42) && (0 < -2147483648)) {
        v1 = (((v3 + v1) + v1) + v3)
        var v10 int64 = -1
        var v11 int64 = v10
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := []int{(fi12 >> ((1) & 63)), fi12, (fi12 / ((127) | 1)), 4294967295}
        _ = fi12
        _ = v13
    }
    v3 = v3
    v3 = "\n\t"
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := []int{(fi14 % ((fi14) | 1)), fi14, genFn0((v2 ^ (v2 / (((v2 | (-(v2)))) | 1))), 35), -117}
        _ = fi14
        _ = v15
    }
    if false {
        v1 = ""
        var v16 int = 32767
        if true {
            fmt.Println(v2, v3, v2)
        }
        _ = v16
    } else {
        fmt.Println(v2, v3, v2)
        v17 := GenS2{F0: -88}
        v3 = (v1 + v3)
        _ = v17
    }
    var v18 int64 = 42
    switch genFn0(v18, 255) {
    case 784:
        switch -9223372036854775808 {
        case 706:
            v19 := []int{32767, 1}
            _ = v19
        }
    default:
        if true {
            v20 := GenS2{F0: 1}
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS3{F0: 0, F1: 15, F2: 255}
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            v2 += (-(v18))
        }
    }
    if (0 <= 1) {
        v23 := []int{128, 482}
        if false {
            if false {
                v24 := GenS1{F0: (-661 < 32767), F1: -239}
                if len(v23) > 0 {
                    v23[0] = 2147483647
                }
                var v25 int32 = 1
                _ = v24
                _ = v25
            }
            if (18446744073709551615 == 0) {
                if (1 >= 932) {
                    v26 := GenS3{F0: 0, F1: -489, F2: 113}
                    var v27 int8 = -128
                    if false {
                        v3 = v1
                        for fi28 := 0; fi28 < 6; fi28++ {
                            v29 := GenS1{F0: ((fi28 / ((int(v27)) | 1)) < (fi28 & fi28)), F1: -1}
                            v27 += v27
                            v18 = v2
                            _ = fi28
                            _ = v29
                        }
                        v3 = (("hello" + "") + v1)
                    } else {
                        for fi30 := 0; fi30 < 8; fi30++ {
                            v2 += 826
                            v31 := GenS4{F0: 1}
                            _ = fi30
                            _ = v31
                        }
                    }
                    _ = v26
                    _ = v27
                } else {
                    var v32 uint = 42
                    _ = v32
                }
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
            }
        } else {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        }
        v3 = ((v3 + v1) + ((v1 + "hello") + v3))
        _ = v23
    } else {
        v3 = ""
    }
    v1 = v1
    v33 := []int{2147483647, 1, 9223372036854775807, 32768}
    if (((1.5 - 0.1) - float32(float64(float32(float64(1e10))))) != ((3.14159 - (1e10 * (((3.14159 - -2.25) * -2.25) * 3.14159))) + ((0.0 / (1.5 + 3.14159)) * ((float32(1e10) - 1e10) / 1.0)))) {
        v33 = append(v33, genFn0(-229, 1))
        v1 = ((v3 + ((v1 + (v3 + v3)) + v3)) + "hello")
        for fi34 := 0; fi34 < 18; fi34++ {
            fmt.Println(fi34, v18, v18)
            v1 = ("foo bar" + v1)
            _ = fi34
        }
    }
    v3 = (v1 + v1)
    var v35 int8 = -128
    var v36 uint = 331
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
}
