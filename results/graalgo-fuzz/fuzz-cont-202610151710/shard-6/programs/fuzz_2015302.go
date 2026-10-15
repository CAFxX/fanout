package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int; F1 int8; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int8; F2 string }
func genFn0() uint32 {
    v4 := GenS1{F0: 2147483648, F1: 2}
    var v5 uint = 401
    v5 = (v5 ^ (v5 >> ((v5) & 63)))
    v6 := []int{}
    fmt.Println(v5)
    _ = v4
    _ = v5
    _ = v6
    return 770
}
func genFn1(p0 uint8, p1 int8) int16 {
    v7 := GenS2{F0: -1, F1: p1, F2: 0}
    if (-957 != 1) {
        v8 := []int{32767}
        switch 4294967295 {
        case 32767:
            var v9 uint8 = (p0 / (((p0 & p0)) | 1))
            _ = v9
        case -997:
            var v10 uint16 = 2
            _ = v10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            switch fi11 {
            case 682:
                switch "foo bar" {
                case "hello":
                    p1 += p1
                case "a":
                    p0 -= 128
                }
            case 2147483647:
                if (42 <= 632) {
                    var v12 float32 = 3.14159
                    _ = v12
                } else {
                    v8 = append(v8, fi11)
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v13 string = ("" + "a")
                    _ = v13
                }
            case 0:
                if false {
                    for fi14 := 0; fi14 < 11; fi14++ {
                        fmt.Println(fi11)
                        _ = fi14
                    }
                    v15 := GenS3{F0: -594}
                    _ = v15
                }
            default:
                v16 := GenS2{F0: int(p0), F1: (p1 + int8(p0)), F2: 1}
                _ = v16
            }
            p0 -= (p0 * (p0 ^ 2))
            v17 := GenS2{F0: (fi11 * fi11), F1: (p1 - 1), F2: 42}
            _ = fi11
            _ = v17
        }
        _ = v8
    }
    p0 = p0
    switch -1 {
    case 9223372036854775807:
        fmt.Println(p0, p0)
    case -2147483648:
        p0 += 0
    default:
        v18 := []int{512, 42, 2147483647, -9223372036854775808}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v7
    return -1
}

func main() {
    v1 := []uint32{1, 2, 444, 421}
    var v2 float32 = 0.0
    v3 := make(map[string]bool)
    fmt.Println(v2)
    v1 = append(v1, 0)
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS4{F0: -1, F1: 127, F2: ("\n\t" + "a")}
        v2 -= 0.1
        var v21 int8 = 78
        _ = fi19
        _ = v20
        _ = v21
    }
    v2 = 1e10
    v2 = -2.25
    var v22 uint32 = 253
    v3[""] = (559 >= 884)
    var v23 bool = (-1 == -9223372036854775808)
    for fi24 := 0; fi24 < 14; fi24++ {
        if ((v22 << ((18446744073709551615) & 63)) >= v22) {
            var v25 uint16 = 0
            v23 = ((v2 / (1.5 - float32((1e10 * float64((v2 * 1.5)))))) >= v2)
            _ = v25
        }
        fmt.Println(fi24, v23, v22)
        if len(v1) > 0 {
            v1[0] = v22
        }
        _ = fi24
    }
    var v26 bool = (2 > 65535)
    delete(v3, "hello")
    var v27 int64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
}
