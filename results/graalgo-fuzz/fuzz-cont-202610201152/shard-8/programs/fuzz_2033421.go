package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int, p1 uint16) uint {
    for fi4 := 0; fi4 < 6; fi4++ {
        if (65535 >= 42) {
            var v5 int64 = 0
            if true {
                var v6 int = (fi4 >> ((664) & 63))
                fmt.Println(v5)
                for fi7 := 0; fi7 < 13; fi7++ {
                    var v8 int16 = -791
                    v8 = 1000
                    _ = fi7
                    _ = v8
                }
                _ = v6
            } else {
                fmt.Println(p1, v5, v5)
                var v9 uint = 65535
                v9 -= v9
                _ = v9
            }
            _ = v5
        }
        _ = fi4
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if false {
            switch (((("hello" + "hello") + ("a" + "hello")) + "\n\t") + "hello") {
            case "\n\t":
                var v11 float64 = float64(((1e10 + (((0.0 * 1e10) - float32(1.5)) / float32((1e10 / 1.0)))) / float32(1e10)))
                _ = v11
            default:
                v12 := []int{p0}
                _ = v12
            }
        }
        _ = fi10
    }
    v13 := []int{(-(p0))}
    p1 = p1
    var v14 int64 = 266
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 42
}
func genFn1(p0 uint64) uint32 {
    v15 := []int{2147483647, 32768}
    for fi16 := 0; fi16 < 7; fi16++ {
        switch (fi16 / ((198) | 1)) {
        case 128:
            v17 := GenS1{F0: 100, F1: 585, F2: 0}
            _ = v17
        case 825:
            var v18 string = "foo bar"
            _ = v18
        case 9223372036854775807:
            switch "\n\t" {
            case "\n\t":
                p0 += 1
            case "a":
                if (!false) {
                    v15 = append(v15, fi16)
                    _ = len(v15)
                    p0 = p0
                } else {
                    if ((float32((((1.5 / (-2.25 + ((float64((3.14159 / 1.0)) / (-2.25 - (3.14159 - -2.25))) / (-2.25 - (0.0 + (-2.25 * float64(0.1))))))) / (1.5 + 1e10)) / -2.25)) + float32(3.14159)) <= 1.5) {
                        var v19 float32 = ((1e10 + float32(float64((0.0 + 3.14159)))) / ((-2.25 - 1e10) + float32(0.0)))
                        _ = v19
                    } else {
                        _ = len(v15)
                        var v20 float32 = ((1e10 / 1.5) - (((1e10 * 0.1) - 3.14159) + 1.5))
                        var v21 int8 = 0
                        _ = v20
                        _ = v21
                    }
                }
            case "":
                p0 -= (p0 % (((p0 / ((p0) | 1))) | 1))
            default:
                if true {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        switch (((("a" + "\n\t") + "") + "") + (((("" + "foo bar") + "a") + "\n\t") + "hello")) {
                        case "foo bar":
                            p0 += (p0 | (p0 | (p0 & p0)))
                        case "":
                            p0 = (p0 * p0)
                        case "hello":
                            p0 += p0
                        }
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = fi22
                    }
                    var v23 int32 = -1
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v23 = (v23 - v23)
                        fmt.Println(fi16, p0, p0)
                        v25 := GenS2{F0: 42, F1: -290}
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                } else {
                    var v26 int16 = 42
                    _ = v26
                }
            }
        default:
            if true {
                var v27 int16 = 32767
                var v28 bool = ((p0 < 18446744073709551615) && false)
                _ = v27
                _ = v28
            } else {
                var v29 uint8 = 2
                v15 = append(v15, 128)
                for fi30 := 0; fi30 < 7; fi30++ {
                    v29 = v29
                    _ = fi30
                }
                _ = v29
            }
        }
        fmt.Println(p0, fi16)
        _ = fi16
    }
    _ = p0
    _ = v15
    return 223
}

func main() {
    var v1 float64 = 0.1
    var v2 bool = (((1 != -32768) || (315 >= 855)) || (2 <= 88))
    var v3 uint64 = 301
    v3 += 0
    v2 = v2
    v31 := GenS2{F0: 1, F1: 42}
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := []int{fi32}
        v34 := GenS1{F0: 100, F1: 2147483647, F2: 4294967295}
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 float32 = 1e10
            var v37 uint64 = 42
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    var v38 bool = (-998 == -9223372036854775808)
    v39 := GenS1{F0: int8(v3), F1: 0, F2: 309}
    v3 -= 18446744073709551615
    v38 = (v2 && (((127 < 2) && (true && (!v2))) && ((float64(1.5) < v1) && v2)))
    var v40 uint8 = 128
    fmt.Println(v38, v40, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v39
    _ = v40
}
