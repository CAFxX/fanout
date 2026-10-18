package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 int; F1 int8 }
func genFn0(p0 uint16, p1 int32, p2 uint) int16 {
    p1 = (p1 >> ((p2) & 63))
    var v7 uint16 = p0
    if true {
        for fi8 := 0; fi8 < 17; fi8++ {
            p1 = p1
            p1 -= (p1 + 1)
            _ = fi8
        }
        var v9 uint8 = 1
        _ = v9
    }
    if (4294967295 <= uint64(p2)) {
        switch ("" + "\n\t") {
        case "a":
            switch "hello" {
            case "hello":
                var v10 uint16 = 307
                _ = v10
            case "":
                fmt.Println(p0, v7, v7)
            default:
                var v11 int16 = 1
                _ = v11
            }
        default:
            switch -9223372036854775808 {
            case 2147483647:
                var v12 uint32 = 2147483648
                _ = v12
            case 167:
                switch "foo bar" {
                case "hello":
                    p1 = p1
                case "\n\t":
                    var v13 int16 = -1
                    _ = v13
                default:
                    p1 += p1
                }
            default:
                if ((-2147483648 < 128) || true) {
                    var v14 int8 = -1
                    v15 := GenS1{F0: 0, F1: p2}
                    p1 += (p1 % ((334) | 1))
                    _ = v14
                    _ = v15
                } else {
                    v7 = uint16(p2)
                }
            }
        }
        var v16 float64 = ((3.14159 * (1e10 * 3.14159)) / (((1e10 - 1.5) / ((3.14159 + (1e10 - ((float64((1.5 - float32(3.14159))) * ((3.14159 - (3.14159 / -2.25)) + 1e10)) * float64(0.0)))) + 0.1)) / float64((float32(1.5) * 1.5))))
        _ = v16
    } else {
        v7 = 255
        for fi17 := 0; fi17 < 19; fi17++ {
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 uint16 = v7
                _ = fi18
                _ = v19
            }
            fmt.Println(v7)
            _ = fi17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}
func genFn1() int8 {
    if false {
        var v20 int64 = 9223372036854775807
        _ = v20
    }
    var v21 int8 = -128
    v22 := []int{32767, 127}
    var v23 uint16 = 936
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    _ = v21
    _ = v22
    _ = v23
    return -33
}

func main() {
    v1 := []int{4294967295}
    v2 := make(map[int]uint32)
    v3 := []uint32{691, 65535, 402}
    var v4 float32 = 1e10
    v5 := make(map[uint32]bool)
    var v6 float32 = (v4 * ((v4 + 0.0) - ((v4 * v4) - float32(float64(v4)))))
    switch 644 {
    case -2147483648:
        var v24 uint = 0
        _ = v24
    default:
        if true {
            _ = len(v1)
            if ((32767 >= 1) || (!(!(0 != 100)))) {
                if false {
                    var v25 uint8 = 2
                    _ = v25
                }
                v4 = v4
                var v26 uint = 110
                _ = v26
            } else {
                var v27 uint16 = 255
                if false {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        for fi29 := 0; fi29 < 1; fi29++ {
                            v6 = v4
                            if len(v1) > 0 {
                                v1[0] = fi28
                            }
                            _ = fi29
                        }
                        v4 = 3.14159
                        _ = fi28
                    }
                } else {
                    fmt.Println(v27, v27)
                    var v30 uint64 = 1
                    _ = v30
                }
                v31 := GenS2{F0: 255}
                _ = v27
                _ = v31
            }
        } else {
            v32 := GenS3{F0: 42}
            v33 := GenS4{F0: -663}
            if false {
                switch (("foo bar" + "hello") + ("\n\t" + "foo bar")) {
                case "a":
                    if (0 != -9223372036854775808) {
                        var v34 int8 = -128
                        v35 := GenS1{F0: 128, F1: 42}
                        _ = v34
                        _ = v35
                    } else {
                        var v36 int8 = 127
                        var v37 bool = false
                        _ = v36
                        _ = v37
                    }
                case "":
                    var v38 uint16 = 53
                    _ = v38
                default:
                    v5[65535] = false
                }
            } else {
                v4 -= 1e10
                fmt.Println(v4)
                for fi39 := 0; fi39 < 20; fi39++ {
                    var v40 float32 = (v6 - (v6 + (v6 + float32(0.1))))
                    v6 -= 1.5
                    for fi41 := 0; fi41 < 15; fi41++ {
                        _ = len(v1)
                        _ = fi41
                    }
                    _ = fi39
                    _ = v40
                }
            }
            _ = v32
            _ = v33
        }
    }
    v6 = v4
    for fi42 := 0; fi42 < 9; fi42++ {
        v43 := GenS4{F0: -32768}
        _ = fi42
        _ = v43
    }
    var v44 float64 = 1e10
    if (-2.25 != (1e10 / v6)) {
        v45 := v2[4294967295]
        _ = v45
    }
    fmt.Println(v4)
    var v46 uint8 = 255
    v44 = float64(float32(v44))
    v4 = v6
    if false {
        var v47 uint = 621
        v48 := GenS5{F0: 349, F1: 2}
        switch "hello" {
        case "foo bar":
            v44 = 3.14159
        default:
            fmt.Println(v46, v6)
        }
        _ = v47
        _ = v48
    }
    v3 = append(v3, 329)
    if (!false) {
        v6 = (1e10 - (v4 * (0.0 * (1e10 * v4))))
        v49 := GenS6{F0: 258, F1: 91}
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v5
    _ = v6
}
