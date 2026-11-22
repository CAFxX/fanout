package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0() uint16 {
    v6 := GenS1{F0: ((0.0 * (((-2.25 - 3.14159) / (0.1 / 1.5)) * (3.14159 * (3.14159 / (float64(0.0) - (-2.25 + (0.1 - (1e10 + 0.1)))))))) != ((-2.25 + (0.0 / -2.25)) * 1e10)), F1: 65535, F2: (1.5 == (0.1 / (1.5 - (-2.25 + (float64(float32(1e10)) / 0.1)))))}
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int16 = 42
        v8 -= 1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := []int{fi9, (fi9 | fi9), (fi9 - fi9), 404}
        _ = fi9
        _ = v10
    }
    _ = v6
    return 255
}
func genFn1() int {
    var v11 uint64 = 662
    v12 := []int{778, 931, -634, 127}
    for fi13 := 0; fi13 < 11; fi13++ {
        v11 -= 512
        _ = fi13
    }
    switch 991 {
    case 9223372036854775807:
        if true {
            if true {
                var v14 int64 = 1
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                v15 := GenS1{F0: false, F1: 42, F2: true}
                _ = v14
                _ = v15
            } else {
                fmt.Println(v11)
            }
            var v16 int64 = -1
            var v17 float32 = ((3.14159 - -2.25) * 1.5)
            _ = v16
            _ = v17
        } else {
            v11 = v11
            fmt.Println(v11)
        }
    case 672:
        v11 = (v11 * (v11 + (v11 / (((v11 & (v11 + (v11 % (((v11 ^ v11)) | 1))))) | 1))))
    case -1:
        var v18 int16 = -1
        _ = v18
    default:
        fmt.Println(v11)
    }
    if false {
        var v19 int = -1
        _ = v19
    }
    _ = v11
    _ = v12
    return 4294967295
}

func main() {
    var v1 uint64 = 775
    var v2 string = ""
    v3 := make(map[int]bool)
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 42
    var v20 uint16 = genFn0()
    fmt.Println(v1, v5)
    var v21 int8 = 100
    v5 -= v5
    var v22 string = "foo bar"
    v23 := []int{451, -560}
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            v21 = 1
            v26 := GenS1{F0: ((v1 - (v1 * 18446744073709551615)) < (v1 >> ((0) & 63))), F1: 0, F2: true}
            switch (v22 + v22) {
            case "\n\t":
                v2 = v2
            case "foo bar":
                var v27 int64 = -1
                v28 := GenS1{F0: ((42 <= 0) && (float64((3.14159 / 1.0)) > -2.25)), F1: 2, F2: ((^(v1)) < (v1 + (v1 << ((42) & 63))))}
                _ = v27
                _ = v28
            case "":
                v29 := GenS1{F0: (-491 < 271), F1: 2147483648, F2: ((false && true) && true)}
                _ = v29
            default:
                v22 = v2
            }
            _ = fi25
            _ = v26
        }
        var v30 string = v2
        switch v30 {
        case "hello":
            if true {
                v20 = (^(v20))
                if len(v23) > 0 {
                    v23[0] = fi24
                }
                v2 = ("a" + v22)
            } else {
                var v31 int = (-(fi24))
                if true {
                    var v32 uint = 420
                    v33 := GenS2{F0: 4294967295, F1: 42}
                    _ = len(v4)
                    _ = v32
                    _ = v33
                }
                _ = v31
            }
        case "":
            v4[4294967295] = 751
        default:
            var v34 float32 = (0.1 - 3.14159)
            _ = v34
        }
        _ = fi24
        _ = v30
    }
    v35 := GenS1{F0: false, F1: 4294967295, F2: false}
    if len(v23) > 0 {
        v23[0] = 0
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 int64 = 1
            v2 = v22
            _ = fi37
            _ = v38
        }
        var v39 int64 = 0
        _ = fi36
        _ = v39
    }
    v40 := GenS3{F0: 1, F1: 767, F2: (v1 << ((255) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v5
}
