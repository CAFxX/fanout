package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint16; F1 int; F2 int8 }
func genFn0() bool {
    v7 := []int{-9223372036854775808, -717, -324, -1}
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int8 = 127
            switch fi8 {
            case 42:
                fmt.Println(fi8, fi9)
            case 32767:
                v11 := GenS1{F0: ("foo bar" + (("foo bar" + ("foo bar" + "hello")) + ("hello" + "\n\t"))), F1: -1}
                _ = v11
            case 4294967295:
                var v12 float32 = ((3.14159 - 1e10) * ((3.14159 + 1.5) + (((-2.25 + 3.14159) + float32(0.0)) * 1e10)))
                _ = v12
            }
            var v13 bool = (!(fi9 <= (fi9 % (((fi8 * fi9)) | 1))))
            _ = fi9
            _ = v10
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int32 = 1
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            _ = len(v7)
            _ = fi16
        }
        _ = fi8
    }
    var v17 uint8 = 128
    _ = v17
    _ = v7
    return false
}
func genFn1(p0 uint32, p1 int32) int {
    for fi18 := 0; fi18 < 3; fi18++ {
        p1 += p1
        v19 := GenS2{F0: true, F1: 100}
        var v20 int64 = 762
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int64 = 42
    v22 := GenS1{F0: ("a" + "\n\t"), F1: p1}
    for fi23 := 0; fi23 < 9; fi23++ {
        switch fi23 {
        case 32767:
            fmt.Println(fi23, p0, fi23)
        }
        _ = fi23
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return 729
}

func main() {
    var v1 int16 = 42
    var v2 float64 = 0.1
    v3 := []int{128}
    v4 := []uint32{4294967295, 1, 65535}
    v5 := []uint32{4294967295, 208, 65535, 899}
    v6 := make(map[string]int)
    if true {
        var v24 float32 = (3.14159 / 1.0)
        if (!true) {
            var v25 uint64 = 18446744073709551615
            v1 = (v1 % ((-942) | 1))
            _ = v25
        } else {
            v4 = append(v4, 494)
        }
        _ = v24
    } else {
        var v26 uint16 = 65535
        var v27 int32 = 1
        v1 -= int16(v26)
        _ = v26
        _ = v27
    }
    v6[("\n\t" + "a")] = 9223372036854775807
    var v28 int64 = 523
    for fi29 := 0; fi29 < 19; fi29++ {
        fmt.Println(v28, fi29)
        _ = len(v6)
        _ = fi29
    }
    var v30 uint16 = 2
    fmt.Println(v28)
    fmt.Println(v2)
    if (0.0 < ((0.1 / 3.14159) / (3.14159 - float32(1.5)))) {
        for fi31 := 0; fi31 < 14; fi31++ {
            for fi32 := 0; fi32 < 17; fi32++ {
                switch fi31 {
                case 127:
                    v33 := v6["a"]
                    _ = v33
                }
                var v34 uint = 1
                v28 = (v28 - v28)
                var v35 int32 = -1
                _ = fi32
                _ = v34
                _ = v35
            }
            if ((^(v1)) == (v1 ^ v1)) {
                for fi36 := 0; fi36 < 14; fi36++ {
                    if true {
                        fmt.Println(v1, fi31)
                        v2 -= v2
                    }
                    v37 := GenS3{F0: 2, F1: fi36, F2: 100}
                    fmt.Println(fi36, v28)
                    _ = fi36
                    _ = v37
                }
                var v38 int16 = -166
                v39 := v6["\n\t"]
                _ = v38
                _ = v39
            }
            _ = fi31
        }
        var v40 uint = 107
        _ = v40
    }
    v1 -= v1
    if (false && true) {
        v28 = (v28 + v28)
    }
    var v41 bool = true
    v42 := GenS2{F0: (1 >= 42), F1: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
