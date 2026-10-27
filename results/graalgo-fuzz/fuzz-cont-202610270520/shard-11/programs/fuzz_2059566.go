package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int }
func genFn0(p0 bool, p1 uint64, p2 uint) int64 {
    var v4 int8 = 42
    if true {
        fmt.Println(p1)
        for fi5 := 0; fi5 < 18; fi5++ {
            fmt.Println(fi5)
            var v6 int64 = 802
            _ = fi5
            _ = v6
        }
    }
    v7 := []int{42, -488, -2147483648, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-465, 1, 42, 32767}
    var v3 int64 = -9223372036854775808
    for fi8 := 0; fi8 < 19; fi8++ {
        v2 = append(v2, (fi8 * fi8))
        var v9 uint64 = 162
        _ = fi8
        _ = v9
    }
    var v10 int8 = 42
    if ((("a" < "\n\t") && (48 == v10)) || true) {
        if true {
            v10 += v10
        } else {
            var v11 string = "\n\t"
            if false {
                if true {
                    switch 1 {
                    case 691:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        fmt.Println(v10, v3, v10)
                    }
                }
                var v12 int32 = -16
                fmt.Println(v11)
                _ = v12
            }
            var v13 uint16 = 1
            _ = v11
            _ = v13
        }
        _ = len(v1)
        var v14 int8 = v10
        _ = v14
    } else {
        var v15 uint = 4294967295
        var v16 uint16 = 65535
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 34, F1: 0, F2: 4294967295}
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int32 = 122
        v20 := GenS1{F0: 128, F1: 2, F2: 42}
        if false {
            v2 = append(v2, (fi18 ^ (fi18 & fi18)))
            var v21 int8 = v10
            _ = len(v2)
            _ = v21
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    var v22 bool = (-2147483648 < 4294967295)
    v23 := GenS2{F0: 4294967295, F1: -1, F2: 255}
    fmt.Println(v3)
    var v24 float64 = float64((1.5 + 1e10))
    fmt.Println(v22)
    delete(v1, 4294967295)
    v24 = (((v24 + v24) - v24) * (((v24 * float64((0.1 / 0.1))) + (v24 * (0.0 / (v24 * v24)))) - 1e10))
    for fi25 := 0; fi25 < 6; fi25++ {
        v3 += 1
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        if ((!v22) || ((42 <= 607) || (v22 && (255 == 255)))) {
            if (3.14159 == float32((float64(1.5) * 1.5))) {
                var v26 int32 = 558
                for fi27 := 0; fi27 < 4; fi27++ {
                    var v28 uint32 = 2147483648
                    v10 = v10
                    var v29 int = fi25
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
            v30 := GenS3{F0: 586}
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 uint8 = 2
                var v33 float64 = 1.5
                switch fi25 {
                case 32768:
                    var v34 int8 = v10
                    _ = v34
                case 348:
                    v33 += ((v33 + -2.25) - ((-2.25 * 0.0) * (float64((-2.25 * (0.0 / (1e10 + 1e10)))) - v33)))
                }
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = v30
        } else {
            switch ("" + (("\n\t" + "\n\t") + "a")) {
            case "\n\t":
                if v22 {
                    v22 = false
                    if true {
                        v35 := GenS4{F0: 99, F1: fi25}
                        var v36 uint64 = 363
                        _ = v35
                        _ = v36
                    }
                } else {
                    for fi37 := 0; fi37 < 11; fi37++ {
                        var v38 uint = 630
                        var v39 int = (fi25 / (((fi25 - fi37)) | 1))
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 16; fi40++ {
                        v41 := GenS1{F0: 127, F1: 42, F2: 665}
                        var v42 uint64 = 0
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    var v43 bool = (127 >= 0)
                    _ = v43
                }
            case "":
                switch fi25 {
                case -2147483648:
                    v3 = v3
                default:
                    if ((v22 || true) && v22) {
                        fmt.Println(v24, fi25, v22)
                        var v44 int64 = (v3 + v3)
                        v24 -= ((v24 / v24) / float64(-2.25))
                        _ = v44
                    }
                }
            }
        }
        _ = fi25
    }
    v45 := GenS2{F0: 277, F1: 1, F2: 560}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v45
}
