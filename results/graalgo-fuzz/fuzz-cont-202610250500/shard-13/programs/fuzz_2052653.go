package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint16; F1 int32 }
func genFn0(p0 bool, p1 uint16) uint {
    if ((!p0) && p0) {
        var v7 int = 4294967295
        fmt.Println(p1, p0)
        var v8 float64 = (((3.14159 * -2.25) * 0.1) * float64(((((0.1 / (3.14159 / 1.5)) * (0.0 * float32(float64(1e10)))) - 1e10) - (0.1 + ((3.14159 - -2.25) / (float32(0.0) - -2.25))))))
        _ = v7
        _ = v8
    }
    if (!((-24 != 42) || (p0 || (155 > 128)))) {
        p0 = (!(1 <= 18446744073709551615))
        p1 -= (p1 ^ 343)
    } else {
        v9 := []int{1, 500}
        _ = v9
    }
    v10 := GenS1{F0: 111, F1: 4294967295, F2: 859}
    switch "hello" {
    case "foo bar":
        var v11 int = 80
        _ = v11
    case "a":
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 uint = 255
            p1 = (p1 - (p1 % (((p1 ^ p1)) | 1)))
            if p0 {
                for fi14 := 0; fi14 < 14; fi14++ {
                    if (65535 == 2147483648) {
                        var v15 int64 = 724
                        if (int16(v15) == 1000) {
                            v16 := GenS1{F0: 127, F1: v13, F2: -151}
                            v15 = 9223372036854775807
                            if (2147483648 <= 879) {
                                v17 := GenS1{F0: 1, F1: v13, F2: 1000}
                                _ = v17
                            } else {
                                fmt.Println(p0, fi12, p0)
                                v18 := []int{fi12, -979}
                                _ = v18
                            }
                            _ = v16
                        }
                        v15 += int64(p1)
                        _ = v15
                    } else {
                        var v19 int64 = 503
                        v13 = (v13 >> ((v13) & 63))
                        _ = v19
                    }
                    if false {
                        var v20 bool = p0
                        fmt.Println(v20, v20)
                        _ = v20
                    }
                    var v21 int32 = -2147483648
                    _ = fi14
                    _ = v21
                }
                var v22 int64 = -9223372036854775808
                _ = v22
            }
            _ = fi12
            _ = v13
        }
    case "":
        var v23 int = 623
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v10
    return 4294967295
}
func genFn1() uint64 {
    var v24 uint32 = 1
    var v25 float32 = 1.5
    _ = v24
    _ = v25
    return 18446744073709551615
}

func main() {
    var v1 int32 = 42
    var v2 uint16 = 117
    v3 := make(map[int]uint32)
    var v4 uint64 = 1
    var v5 float64 = (float64(float32((-2.25 / 1.5))) - (1e10 / 0.1))
    v6 := make(map[string]string)
    var v26 uint = 726
    var v27 uint32 = 42
    var v28 int16 = 42
    var v29 bool = true
    if (("foo bar" != "\n\t") || true) {
        if v29 {
            for fi30 := 0; fi30 < 16; fi30++ {
                v27 -= (v27 | 2147483647)
                var v31 bool = ((true || v29) && (216 >= 42))
                _ = fi30
                _ = v31
            }
            v32 := GenS1{F0: 42, F1: (v26 >> (((^(v26))) & 63)), F2: 42}
            _ = v32
        } else {
            delete(v6, "\n\t")
            switch "a" {
            case "a":
                var v33 float64 = (v5 + float64((1.5 + (0.0 + float32(1e10)))))
                _ = v33
            }
            switch 42 {
            case -9223372036854775808:
                v34 := GenS2{F0: 669}
                _ = v34
            default:
                fmt.Println(v28, v28)
            }
        }
    }
    v35 := GenS1{F0: 170, F1: uint(v2), F2: 315}
    var v36 uint8 = uint8(v26)
    v37 := GenS3{F0: 1}
    for fi38 := 0; fi38 < 3; fi38++ {
        v36 = v36
        _ = fi38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        var v40 uint8 = (v36 | uint8(v2))
        _ = fi39
        _ = v40
    }
    var v41 int8 = -1
    var v42 uint8 = 255
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v43 bool = v29
        _ = v43
    case 1:
        v44 := GenS4{F0: v27, F1: 0}
        _ = v44
    case -1:
        v45 := GenS1{F0: v36, F1: (v26 << ((v26) & 63)), F2: -1}
        _ = v45
    }
    v46 := GenS3{F0: -451}
    fmt.Println(v41)
    v47 := GenS5{F0: (v2 + 42), F1: (-(v1))}
    var v48 float64 = ((((float64(float32(0.0)) * 0.1) - v5) * (v5 - (float64(0.1) + (1e10 * v5)))) - 1.5)
    v3[1] = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
