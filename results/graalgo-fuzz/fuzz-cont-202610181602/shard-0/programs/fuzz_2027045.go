package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0(p0 int64, p1 int) int {
    if true {
        var v5 uint = 65535
        _ = v5
    } else {
        v6 := []int{p1}
        var v7 int = (p1 % ((-296) | 1))
        _ = v6
        _ = v7
    }
    if true {
        switch p1 {
        case 32768:
            var v8 uint64 = 858
            _ = v8
        case -2147483648:
            var v9 uint32 = 65535
            _ = v9
        case -989:
            var v10 uint32 = 2147483648
            _ = v10
        default:
            if ((!false) || (-434 > 38)) {
                v11 := []int{}
                if ((("foo bar" + "a") + ("foo bar" + "a")) > "a") {
                    if true {
                        p1 = p1
                        p0 = p0
                        p1 = (p1 >> ((65535) & 63))
                    }
                } else {
                    p0 += p0
                }
                _ = v11
            }
        }
        var v12 uint32 = 236
        var v13 int64 = 231
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            p1 += fi14
            _ = fi15
        }
        if ("foo bar" < "") {
            p0 += 0
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS1{F0: (2 == 0), F1: 144, F2: 4294967295}
                p0 += p0
                fmt.Println(p1, fi16)
                _ = fi16
                _ = v17
            }
            v18 := GenS1{F0: true, F1: 143, F2: 2}
            _ = v18
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int64 = 487
    v2 := []uint32{1, 4294967295, 42}
    v3 := make(map[uint32]int)
    var v4 bool = false
    if (!false) {
        v4 = (0 > 2147483648)
        if (false || (((365 != -1) && v4) && v4)) {
            v19 := GenS2{F0: 0, F1: 746, F2: 1}
            _ = v19
        } else {
            var v20 int8 = 100
            _ = v20
        }
        v4 = true
    }
    v4 = (60 != 2147483648)
    for fi21 := 0; fi21 < 4; fi21++ {
        v1 += v1
        var v22 int32 = 0
        _ = fi21
        _ = v22
    }
    if true {
        var v23 string = "\n\t"
        v3[680] = 32767
        switch 1 {
        case -9223372036854775808:
            var v24 uint64 = 2
            _ = v24
        }
        _ = v23
    } else {
        v25 := GenS1{F0: v4, F1: 255, F2: 42}
        var v26 bool = false
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: 0, F1: 2147483648, F2: 664}
    v2 = append(v2, 1)
    var v28 string = "hello"
    if (!(42 >= 100)) {
        v29 := GenS3{F0: ("\n\t" + (v28 + (("" + (v28 + v28)) + ((("\n\t" + "foo bar") + v28) + v28)))), F1: 978}
        if (-1 < 607) {
            v30 := GenS4{F0: 4294967295}
            _ = v30
        } else {
            var v31 float64 = -2.25
            _ = v31
        }
        var v32 uint = 18446744073709551615
        _ = v29
        _ = v32
    }
    if (v4 && (65535 > 1)) {
        var v33 float32 = (-2.25 * -2.25)
        v33 -= ((float32((-2.25 * (1e10 * 0.0))) - v33) - v33)
        _ = v33
    } else {
        fmt.Println(v28, v4, v4)
        fmt.Println(v4)
    }
    var v34 uint64 = 4294967295
    v35 := GenS5{F0: 128, F1: (genFn0(v1, 2147483647) > int(v1)), F2: 0}
    for fi36 := 0; fi36 < 13; fi36++ {
        if v4 {
            if len(v2) > 0 {
                v2[0] = 289
            }
            fmt.Println(v28, v28, v28)
            v28 = ("hello" + (((v28 + v28) + (v28 + ((v28 + "foo bar") + (("" + (v28 + v28)) + ("a" + ("\n\t" + ((v28 + ("" + v28)) + (v28 + "\n\t")))))))) + v28))
        }
        _ = fi36
    }
    var v37 int16 = -1
    v1 -= v1
    v4 = (255 == 0)
    v4 = (44 > 2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
}
