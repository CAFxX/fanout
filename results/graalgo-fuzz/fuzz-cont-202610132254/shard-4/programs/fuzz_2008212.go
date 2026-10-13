package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint8) uint16 {
    fmt.Println(p0)
    v7 := GenS1{F0: 913, F1: 1}
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 uint64) int {
    p0 = (^(p0))
    if true {
        var v8 string = "\n\t"
        v8 = "hello"
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 int = fi9
            if false {
                v11 := GenS2{F0: 0, F1: 65535}
                var v12 int16 = -1
                fmt.Println(v12)
                _ = v11
                _ = v12
            }
            var v13 bool = true
            _ = fi9
            _ = v10
            _ = v13
        }
        _ = v8
    } else {
        p0 = p0
    }
    v14 := []int{32767, 946}
    p0 += (p0 + (p0 | p0))
    p0 -= (p0 ^ (p0 << ((1) & 63)))
    _ = p0
    _ = v14
    return 14
}

func main() {
    var v1 uint16 = 1
    var v2 string = ("\n\t" + "hello")
    var v3 uint32 = 966
    var v4 float32 = -2.25
    var v5 int64 = -9223372036854775808
    var v6 float32 = float32(float64(v4))
    if (0 > 128) {
        var v15 uint32 = 859
        if (2 == v1) {
            if (!(255 > genFn0(79))) {
                var v16 string = ("" + "a")
                _ = v16
            }
        } else {
            if false {
                if (!false) {
                    v5 -= (^(v5))
                    var v17 uint16 = 65535
                    var v18 uint32 = (^(v3))
                    _ = v17
                    _ = v18
                }
                var v19 float64 = 1e10
                _ = v19
            } else {
                fmt.Println(v3, v2, v3)
            }
            switch v2 {
            case "foo bar":
                switch 4294967295 {
                case -601:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v21 := GenS3{F0: -32768}
                        var v22 uint64 = 348
                        v23 := GenS3{F0: -500}
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                case 128:
                    v3 = v3
                }
            case "a":
                for fi24 := 0; fi24 < 19; fi24++ {
                    v25 := []int{}
                    v25 = append(v25, genFn1(317))
                    _ = fi24
                    _ = v25
                }
            case "hello":
                fmt.Println(v6)
            }
            var v26 string = ""
            _ = v26
        }
        _ = v15
    }
    v27 := GenS3{F0: -1}
    v28 := []int{9223372036854775807, -2147483648}
    _ = len(v28)
    v29 := GenS1{F0: 507, F1: genFn0(203)}
    for fi30 := 0; fi30 < 16; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            v32 := GenS1{F0: 255, F1: (v1 * 2)}
            switch "hello" {
            case "foo bar":
                if (4294967295 >= 667) {
                    fmt.Println(v6, v1, fi30)
                    var v33 uint32 = (v3 * v3)
                    _ = v33
                } else {
                    v1 = (v1 | (v1 ^ v1))
                    var v34 uint64 = 2
                    _ = v34
                }
            case "a":
                var v35 float32 = ((-2.25 - v6) - v4)
                _ = v35
            default:
                fmt.Println(v5, v1, v3)
            }
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        fmt.Println(fi36, fi36)
        var v37 uint32 = 42
        v5 -= v5
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        v39 := GenS4{F0: -44, F1: v3, F2: 163}
        v40 := GenS1{F0: 65535, F1: 65535}
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := GenS1{F0: 4294967295, F1: 93}
    v4 = ((v4 / 0.1) / v4)
    var v42 float32 = v4
    v43 := GenS5{F0: 87}
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
