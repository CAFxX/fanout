package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 bool, p1 uint16) int16 {
    v6 := GenS1{F0: p1, F1: 224, F2: (4294967295 <= 0)}
    v7 := []int{127, -999, 1, 32767}
    switch "hello" {
    case "":
        var v8 uint16 = (p1 << ((1) & 63))
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -84
}

func main() {
    var v1 float64 = float64((((3.14159 * float32(float64(3.14159))) * 0.0) - (float32(float64(1e10)) - float32(0.0))))
    var v2 int8 = -49
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{42, 2147483648}
    v9 := GenS2{F0: 1, F1: 0}
    var v10 int16 = -875
    if (true && false) {
        if ((float64(-2.25) / 3.14159) != v1) {
            var v11 bool = ((1.5 + (3.14159 * -2.25)) > 0.0)
            for fi12 := 0; fi12 < 14; fi12++ {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v10 = v10
                    _ = len(v3)
                    v10 -= (v10 - 824)
                    _ = fi13
                }
                var v14 uint8 = 1
                v14 -= (v14 / ((v14) | 1))
                _ = fi12
                _ = v14
            }
            switch ("a" + "a") {
            case "a":
                var v15 int = -852
                _ = v15
            case "":
                v16 := GenS2{F0: 128, F1: 114}
                _ = v16
            case "hello":
                var v17 float64 = ((v1 / v1) - (v1 * v1))
                _ = v17
            default:
                v11 = false
            }
            _ = v11
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            if false {
                if (!(false || false)) {
                    v10 -= v10
                }
                var v19 int32 = 2147483647
                _ = v19
            }
            v2 += v2
            v20 := GenS3{F0: 22, F1: (fi18 / (((fi18 >> ((42) & 63))) | 1)), F2: 2}
            _ = fi18
            _ = v20
        }
    } else {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 int16 = (v10 - genFn0((-1 != -2147483648), 2))
            for fi23 := 0; fi23 < 6; fi23++ {
                v24 := v4[65535]
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        if len(v5) > 0 {
            v5[0] = 78
        }
        var v25 int64 = 10
        _ = v25
    }
    v26 := GenS3{F0: 27, F1: 0, F2: 0}
    var v27 uint8 = 128
    v1 -= ((v1 * (0.1 * v1)) / (v1 / v1))
    var v28 uint16 = 42
    v29 := v3[int(v10)]
    for fi30 := 0; fi30 < 18; fi30++ {
        if (-293 <= fi30) {
            v2 -= (-(v2))
            for fi31 := 0; fi31 < 5; fi31++ {
                v2 -= v2
                var v32 uint64 = 42
                _ = fi31
                _ = v32
            }
        } else {
            for fi33 := 0; fi33 < 19; fi33++ {
                var v34 uint64 = 836
                switch "hello" {
                case "foo bar":
                    switch int(v10) {
                    case 127:
                        v34 = (^(v34))
                    default:
                        if false {
                            v35 := GenS2{F0: (-(fi33)), F1: (v27 + v27)}
                            v29 = (v29 + ((("a" + v29) + v29) + v29))
                            _ = v35
                        }
                    }
                case "\n\t":
                    if false {
                        fmt.Println(fi30)
                        v36 := GenS2{F0: (fi30 % ((fi30) | 1)), F1: (^(v27))}
                        v37 := GenS2{F0: fi30, F1: (v27 * v27)}
                        _ = v36
                        _ = v37
                    }
                case "hello":
                    delete(v4, 65535)
                    if len(v5) > 0 {
                        v5[0] = 0
                    }
                    var v38 float32 = (((-2.25 / float32(v1)) * 1e10) + 3.14159)
                    _ = v38
                default:
                    v28 = 643
                }
                v29 = (v29 + (("\n\t" + ((v29 + ((v29 + (((((v29 + (v29 + v29)) + (("" + (("hello" + ("a" + (("" + v29) + (((v29 + (("foo bar" + (v29 + "")) + "\n\t")) + "hello") + v29)))) + v29)) + ("\n\t" + ("" + v29)))) + v29) + "a") + v29)) + v29)) + "hello")) + "foo bar"))
                if (-2.25 > -2.25) {
                    v1 = v1
                    v39 := GenS4{F0: 42}
                    v5 = append(v5, 42)
                    fmt.Println(fi33, v27)
                    _ = v39
                } else {
                    v40 := GenS2{F0: (fi33 % ((932) | 1)), F1: v27}
                    _ = v40
                }
                v41 := v4[65535]
                _ = fi33
                _ = v34
                _ = v41
            }
        }
        _ = fi30
    }
    var v42 int = 1
    switch int(v10) {
    case 4294967295:
        if true {
            v5 = append(v5, 42)
            v43 := GenS1{F0: (v28 * v28), F1: 419, F2: (100 > v2)}
            for fi44 := 0; fi44 < 13; fi44++ {
                fmt.Println(v28, fi44)
                if true {
                    v45 := GenS3{F0: 93, F1: int(v28), F2: 4294967295}
                    fmt.Println(fi44, v1)
                    for fi46 := 0; fi46 < 15; fi46++ {
                        var v47 int16 = v10
                        _ = fi46
                        _ = v47
                    }
                    _ = v45
                }
                _ = fi44
            }
            _ = v43
        }
    case 32768:
        var v48 float32 = float32(0.0)
        _ = v48
    case -9223372036854775808:
        v2 = -26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v9
}
