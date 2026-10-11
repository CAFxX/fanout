package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int16; F2 string }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
type GenS4 struct { F0 int16; F1 int16; F2 int8 }
func genFn0() uint {
    var v4 uint32 = 2147483648
    if (!false) {
        switch "" {
        case "hello":
            var v5 bool = (true && (false || true))
            _ = v5
        default:
            v6 := GenS1{F0: (1000 <= 633)}
            _ = v6
        }
    }
    if false {
        var v7 uint16 = 42
        _ = v7
    }
    if (!(false && false)) {
        v8 := []int{128, 32768, 2147483647, 128}
        var v9 uint8 = 128
        v9 -= v9
        _ = v8
        _ = v9
    } else {
        v10 := []int{-537, -2147483648}
        if false {
            for fi11 := 0; fi11 < 14; fi11++ {
                if ((65535 != 42) && true) {
                    var v12 bool = false
                    _ = v12
                }
                _ = fi11
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                for fi14 := 0; fi14 < 5; fi14++ {
                    v4 -= v4
                    switch fi14 {
                    case 32768:
                        fmt.Println(fi14, v4, fi13)
                    case -909:
                        v15 := GenS2{F0: 19, F1: -908, F2: "hello"}
                        _ = v15
                    default:
                        if (1 < 0) {
                            var v16 bool = (-128 < -1)
                            v4 += (v4 | (^(v4)))
                            _ = v16
                        } else {
                            var v17 int8 = 23
                            if (886 >= 1) {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                v4 = v4
                                v4 = (v4 ^ (v4 ^ (v4 ^ (v4 ^ (v4 & (v4 << ((2) & 63)))))))
                            }
                            fmt.Println(v4)
                            _ = v17
                        }
                    }
                    switch "hello" {
                    case "a":
                        fmt.Println(fi14)
                    }
                    _ = fi14
                }
                _ = fi13
            }
            var v18 int64 = 0
            _ = v18
        } else {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            fmt.Println(v4)
            switch 988 {
            case -475:
                v4 += v4
            case 128:
                for fi19 := 0; fi19 < 2; fi19++ {
                    v4 -= v4
                    _ = fi19
                }
            case 9223372036854775807:
                v4 -= (^(v4))
            }
        }
        var v20 float64 = 0.0
        _ = v10
        _ = v20
    }
    _ = v4
    return 0
}
func genFn1() uint {
    switch "\n\t" {
    case "a":
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 uint16 = 32767
            v23 := GenS2{F0: 778, F1: 1, F2: ("" + (("a" + ("foo bar" + "a")) + "\n\t"))}
            var v24 bool = false
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    case "hello":
        var v25 float32 = (0.1 - 1e10)
        _ = v25
    case "":
        var v26 string = "foo bar"
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        fmt.Println(fi27)
        switch 1 {
        case 128:
            v28 := []int{(fi27 >> ((987) & 63)), fi27, fi27, 518}
            _ = v28
        case 2147483647:
            var v29 uint64 = 1
            _ = v29
        case -906:
            for fi30 := 0; fi30 < 10; fi30++ {
                var v31 int64 = 147
                _ = fi30
                _ = v31
            }
        }
        _ = fi27
    }
    var v32 int32 = -1
    v33 := GenS3{F0: (v32 & v32), F1: 2, F2: 1}
    _ = v32
    _ = v33
    return 42
}

func main() {
    v1 := []int{-848, 781, -9223372036854775808}
    v2 := make(map[uint32]bool)
    var v3 uint = 0
    _ = len(v1)
    if (!(1 != 100)) {
        if true {
            v34 := GenS1{F0: (!true)}
            v3 += v3
            _ = v34
        }
        if false {
            v35 := GenS4{F0: 0, F1: 0, F2: -128}
            v3 = 661
            var v36 int = 599
            _ = v35
            _ = v36
        }
    } else {
        v37 := GenS3{F0: 1, F1: 1, F2: 559}
        var v38 int = -2147483648
        _ = v37
        _ = v38
    }
    v3 += (v3 ^ (v3 << ((v3) & 63)))
    fmt.Println(v3)
    var v39 int32 = 560
    var v40 uint32 = 1
    if len(v1) > 0 {
        v1[0] = -297
    }
    var v41 int = -700
    var v42 float64 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
