package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint32, p2 uint32) uint32 {
    p2 += p2
    var v5 float64 = ((0.1 / 3.14159) - 1.5)
    var v6 uint8 = 105
    v6 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p2
}

func main() {
    var v1 float64 = 0.0
    v2 := []string{("a" + ("" + "foo bar")), "foo bar", "foo bar"}
    v3 := make(map[string]int)
    v4 := []uint32{377, 42}
    v1 = (v1 - 0.0)
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := v3["\n\t"]
        v1 = (v1 + 1e10)
        if (true && false) {
            switch -735 {
            case -2147483648:
                fmt.Println(v1)
            case 32767:
                v1 = v1
            case 127:
                fmt.Println(v1, v8)
            }
        } else {
            if false {
                if (!(("" + "hello") == ((("hello" + "a") + ("hello" + "a")) + ("hello" + "\n\t")))) {
                    var v9 uint32 = 2
                    if true {
                        switch v8 {
                        case 2147483647:
                            v10 := GenS1{F0: 1, F1: 32767, F2: 411}
                            _ = v10
                        }
                    }
                    _ = v9
                }
                v4 = append(v4, 65535)
            } else {
                if (!false) {
                    v1 = (1.5 - v1)
                    if (-1 == -1) {
                        v11 := GenS1{F0: 42, F1: 1, F2: 1}
                        _ = v11
                    }
                    v12 := GenS1{F0: 409, F1: -32768, F2: 18446744073709551615}
                    _ = v12
                }
            }
            var v13 uint32 = 0
            switch v8 {
            case 1:
                fmt.Println(fi7, fi7, v8)
            }
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS2{F0: "\n\t", F1: 1}
        switch fi14 {
        case -803:
            var v16 float64 = (3.14159 * (v1 / (v1 + v1)))
            _ = v16
        default:
            v1 = v1
        }
        fmt.Println(v1)
        _ = fi14
        _ = v15
    }
    if true {
        v1 += v1
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        switch (fi17 | (fi17 * (^(fi17)))) {
        case 128:
            var v18 uint = 952
            _ = v18
        case 574:
            v1 = (v1 - 0.1)
        case 0:
            v19 := v3[(("hello" + (("" + ("" + "\n\t")) + "hello")) + "")]
            _ = v19
        }
        if true {
            v1 = 0.1
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 float64 = 3.14159
        v22 := GenS1{F0: 4294967295, F1: -1, F2: 4294967295}
        _ = fi20
        _ = v21
        _ = v22
    }
    switch 127 {
    case 1:
        v1 = v1
    case 128:
        if false {
            var v23 int = 845
            _ = v23
        }
    }
    switch 2147483647 {
    case -806:
        v1 += (((((((v1 - (v1 * 3.14159)) - v1) + float64(float32(v1))) + (v1 * ((1.5 / 3.14159) + (v1 + (v1 * v1))))) / (v1 + float64(0.0))) + (3.14159 - 1e10)) - v1)
    case -9223372036854775808:
        if true {
            fmt.Println(v1)
            for fi24 := 0; fi24 < 17; fi24++ {
                v1 += ((3.14159 / 3.14159) + v1)
                _ = fi24
            }
        }
    default:
        v1 = (((v1 / float64((0.1 + float32((3.14159 / v1))))) * ((-2.25 - v1) * v1)) + 1e10)
    }
    v25 := GenS3{F0: 1}
    v26 := GenS4{F0: 715}
    var v27 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
