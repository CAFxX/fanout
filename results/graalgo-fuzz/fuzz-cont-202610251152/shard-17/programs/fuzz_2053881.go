package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint32, p1 int64, p2 int64) uint {
    v5 := []int{1, -1, 42}
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            _ = len(v5)
            _ = fi7
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 4294967295
}
func genFn1(p0 uint) int64 {
    var v8 string = (("" + "") + (("hello" + "foo bar") + "a"))
    switch "\n\t" {
    case "foo bar":
        switch ("\n\t" + ("hello" + v8)) {
        case "a":
            p0 -= p0
        case "\n\t":
            if true {
                var v9 uint = p0
                _ = v9
            }
        case "foo bar":
            var v10 bool = true
            _ = v10
        }
    }
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 float64 = float64(0.0)
    var v4 int16 = 415
    var v11 int8 = -128
    var v12 float64 = (v3 / float64(float32(((float64((-2.25 - -2.25)) + v3) * v3))))
    v12 = v3
    var v13 uint16 = 255
    fmt.Println(v11, v12)
    for fi14 := 0; fi14 < 4; fi14++ {
        _ = len(v2)
        var v15 float64 = v3
        v16 := []int{2147483647, (^(fi14)), (fi14 << ((1) & 63)), (fi14 / ((fi14) | 1))}
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        if false {
            var v17 float64 = 3.14159
            if false {
                for fi18 := 0; fi18 < 13; fi18++ {
                    var v19 int16 = v4
                    _ = fi18
                    _ = v19
                }
                v1[""] = "hello"
                switch 9223372036854775807 {
                case 42:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v11 = 100
                        v11 += (v11 + (v11 & (v11 >> ((0) & 63))))
                        fmt.Println(v3)
                        _ = fi20
                    }
                case 32767:
                    var v21 uint = 1
                    _ = v21
                case 127:
                    switch 9223372036854775807 {
                    case -2147483648:
                        v13 -= v13
                    default:
                        for fi22 := 0; fi22 < 16; fi22++ {
                            switch "" {
                            case "":
                                v23 := v1["a"]
                                _ = v23
                            case "\n\t":
                                v24 := v1["\n\t"]
                                _ = v24
                            case "a":
                                var v25 uint = 103
                                _ = v25
                            }
                            _ = fi22
                        }
                    }
                }
            }
            _ = v17
        } else {
            var v26 uint = 713
            v13 = v13
            _ = v26
        }
        v11 = v11
    }
    v27 := GenS1{F0: 606, F1: 894, F2: (-(v4))}
    if false {
        v28 := GenS2{F0: 4294967295}
        var v29 uint32 = 514
        if (255 == 927) {
            v30 := GenS3{F0: 802, F1: v13}
            _ = v30
        }
        _ = v28
        _ = v29
    }
    fmt.Println(v3, v11)
    v31 := GenS1{F0: 1, F1: 577, F2: int16(v13)}
    v32 := GenS3{F0: 65535, F1: (v13 * v13)}
    var v33 float32 = 0.1
    v12 = 1.5
    fmt.Println(v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
