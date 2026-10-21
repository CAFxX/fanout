package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int8; F2 int16 }
type GenS6 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint, p1 uint16, p2 uint32) uint {
    v4 := GenS1{F0: -18}
    var v5 bool = ("a" != ("foo bar" + ("" + ("" + "hello"))))
    if v5 {
        v5 = (-1 != -523)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}
func genFn1() uint8 {
    var v6 uint64 = 1
    v7 := GenS2{F0: 0, F1: 255}
    v8 := GenS3{F0: 181}
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 uint64 = 42
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint16 = 360
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS4{F0: 42}
            v10 = (v10 * 42)
            switch (fi11 - fi9) {
            case 4294967295:
                if true {
                    v10 -= 1
                    var v13 int = (fi11 * (fi9 << ((18446744073709551615) & 63)))
                    _ = v13
                } else {
                    v3 = v3
                    v10 = 42
                }
            case 1:
                v14 := GenS3{F0: v10}
                var v15 uint8 = 127
                fmt.Println(fi11, fi11, fi9)
                var v16 uint64 = 4294967295
                v3 += v16
                v3 += 354
                if (0 > 2147483647) {
                    var v17 float32 = float32(((float64(float32((float64(((1e10 / 1.0) - 0.0)) + 0.0))) / 0.1) * 1.5))
                    _ = v17
                }
                var v18 uint16 = (^(v10))
                _ = v14
                _ = v15
                _ = v16
                _ = v18
            case 32767:
                switch (fi9 - (fi9 / ((fi11) | 1))) {
                case -924:
                    v1[fi9] = 2
                case 9223372036854775807:
                    var v19 int8 = -128
                    _ = v19
                }
                v20 := []int{669, (fi11 << ((4294967295) & 63)), fi11, fi9}
                _ = v20
            default:
                var v21 uint64 = v3
                _ = v21
            }
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    switch 635 {
    case 386:
        v22 := GenS2{F0: -192, F1: 20}
        _ = v22
    case -2147483648:
        v23 := GenS2{F0: -678, F1: 2}
        _ = v23
    default:
        switch -2147483648 {
        case 634:
            if (((!true) && false) || false) {
                fmt.Println(v3)
            } else {
                var v24 uint8 = 42
                var v25 uint64 = (v3 * v3)
                _ = v24
                _ = v25
            }
        case 127:
            if ((2 >= 4294967295) && false) {
                switch -2147483648 {
                case -137:
                    _ = len(v1)
                case 32767:
                    fmt.Println(v3)
                case 32768:
                    if (!true) {
                        v3 += (v3 - (v3 ^ (v3 ^ (v3 & 463))))
                        var v26 int16 = 1
                        _ = v26
                    }
                default:
                    for fi27 := 0; fi27 < 8; fi27++ {
                        if false {
                            v3 += (v3 >> ((genFn0(255, 749, 0)) & 63))
                            var v28 uint64 = (v3 * (v3 / (((v3 * v3)) | 1)))
                            v3 = v3
                            _ = v28
                        } else {
                            var v29 uint = 65535
                            _ = v29
                        }
                        fmt.Println(fi27, fi27)
                        for fi30 := 0; fi30 < 3; fi30++ {
                            v3 -= 2
                            var v31 int32 = 2147483647
                            _ = fi30
                            _ = v31
                        }
                        _ = fi27
                    }
                }
                var v32 float32 = (((0.1 / 1.0) * 0.1) + (1.5 * 3.14159))
                _ = v32
            } else {
                var v33 int8 = -92
                _ = v33
            }
        }
    }
    _ = len(v2)
    var v34 uint = genFn0(4294967295, 65535, 2147483648)
    if ("hello" < "a") {
        v35 := []int{4294967295}
        v3 = uint64(v34)
        _ = v35
    }
    var v36 int8 = 100
    if false {
        if (0 == -32768) {
            fmt.Println(v36, v34)
            fmt.Println(v3)
        } else {
            for fi37 := 0; fi37 < 2; fi37++ {
                v38 := GenS3{F0: 0}
                v36 = v36
                _ = fi37
                _ = v38
            }
        }
        fmt.Println(v36, v36, v34)
        var v39 uint = (v34 - v34)
        _ = v39
    }
    v40 := GenS3{F0: 0}
    v36 = (v36 & (v36 >> (((v34 * v34)) & 63)))
    v36 = (v36 + (v36 * v36))
    switch -2147483648 {
    case -9223372036854775808:
        fmt.Println(v3)
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        v42 := []int{}
        var v43 float32 = (((0.1 * ((-2.25 / -2.25) * (3.14159 - 0.0))) * (0.1 * (1e10 + 0.0))) * 3.14159)
        _ = fi41
        _ = v42
        _ = v43
    }
    if (!true) {
        v44 := []int{-9223372036854775808}
        _ = v44
    }
    if false {
        v45 := GenS1{F0: (v36 ^ (v36 | v36))}
        v3 += (^(v3))
        v36 = (v36 | v36)
        _ = v45
    } else {
        v1[32767] = 2147483647
        v46 := GenS5{F0: true, F1: v36, F2: -146}
        for fi47 := 0; fi47 < 14; fi47++ {
            for fi48 := 0; fi48 < 10; fi48++ {
                if (-1 < 0) {
                    if ((42 != 42) || true) {
                        v49 := GenS6{F0: v34, F1: 0}
                        _ = v49
                    } else {
                        v34 = 18446744073709551615
                        v50 := GenS6{F0: 1, F1: 127}
                        _ = v50
                    }
                    var v51 uint = v34
                    v34 = (^(v34))
                    _ = v51
                }
                v36 = 100
                _ = fi48
            }
            for fi52 := 0; fi52 < 14; fi52++ {
                v53 := GenS2{F0: 42, F1: 127}
                fmt.Println(fi47)
                switch "foo bar" {
                case "a":
                    v36 = -103
                }
                _ = fi52
                _ = v53
            }
            _ = fi47
        }
        _ = v46
    }
    if ((!(1 <= 779)) && false) {
        var v54 uint32 = uint32(v3)
        var v55 int16 = 36
        fmt.Println(v34)
        _ = v54
        _ = v55
    }
    var v56 bool = (-902 >= 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v40
    _ = v56
}
