package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int16; F1 uint64 }
func genFn0(p0 string, p1 int8, p2 int16) int32 {
    v4 := []int{}
    if true {
        if false {
            if false {
                var v5 uint8 = 143
                v4 = append(v4, -2147483648)
                switch 32767 {
                case -1:
                    var v6 bool = (4294967295 == 1)
                    _ = v6
                case 42:
                    v4 = append(v4, -2147483648)
                }
                _ = v5
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if ((1.5 + 3.14159) >= 3.14159) {
                    for fi7 := 0; fi7 < 12; fi7++ {
                        fmt.Println(p1)
                        _ = fi7
                    }
                    switch -2147483648 {
                    case -2147483648:
                        _ = len(v4)
                    case 42:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    p0 = (((p0 + p0) + (((("foo bar" + p0) + p0) + p0) + p0)) + (("foo bar" + "foo bar") + p0))
                } else {
                    p1 += p1
                    p1 -= p1
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                }
            }
            p0 = ("" + p0)
            v8 := GenS1{F0: 2}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 float32 = (3.14159 * -2.25)
                v11 := GenS1{F0: 260}
                fmt.Println(fi9, fi9, p0)
                _ = fi9
                _ = v10
                _ = v11
            }
            p2 = p2
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 int8 = 66
                fmt.Println(v13)
                v14 := GenS2{F0: 42, F1: 255}
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        var v15 uint64 = 4294967295
        switch -847 {
        case 523:
            var v16 float64 = -2.25
            _ = v16
        }
        _ = v15
    } else {
        for fi17 := 0; fi17 < 4; fi17++ {
            p0 = "hello"
            _ = fi17
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v18 := GenS3{F0: p2, F1: 42}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -2147483648
}

func main() {
    var v1 uint = 625
    var v2 uint8 = 199
    var v3 float64 = float64(3.14159)
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := []int{}
        if (2 != 42) {
            var v21 int64 = -1
            _ = v21
        } else {
            v3 -= float64((1e10 * 1e10))
            v2 = (v2 >> (((v1 >> ((2) & 63))) & 63))
            v22 := GenS1{F0: 1}
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    v23 := []int{915, 32768, -9223372036854775808, 618}
    for fi24 := 0; fi24 < 4; fi24++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        switch (fi24 * (fi24 + (fi24 * (fi24 % (((fi24 * -9223372036854775808)) | 1))))) {
        case -9223372036854775808:
            v3 = 1e10
        case -106:
            var v25 string = "\n\t"
            _ = v25
        default:
            var v26 int16 = -398
            _ = v26
        }
        _ = fi24
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 uint32 = 4294967295
        v1 -= 18446744073709551615
        _ = fi27
        _ = v28
    }
    fmt.Println(v3)
    var v29 int = 32767
    var v30 uint8 = (v2 & (v2 * (v2 / (((v2 + 1)) | 1))))
    v29 = (v29 ^ (^(v29)))
    v31 := GenS3{F0: 42, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
