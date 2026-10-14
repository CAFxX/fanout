package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int; F1 int64 }
func genFn0() uint32 {
    switch 32768 {
    case 32768:
        var v4 int32 = 0
        _ = v4
    case 133:
        var v5 int8 = 1
        _ = v5
    case 127:
        if true {
            if true {
                var v6 uint16 = 156
                v6 = v6
                _ = v6
            }
        } else {
            var v7 int32 = -1
            _ = v7
        }
    }
    var v8 uint16 = 255
    if true {
        var v9 int8 = 61
        v10 := GenS1{F0: v8, F1: (-(v9)), F2: 2147483647}
        v11 := GenS1{F0: v8, F1: (-(v9)), F2: -2147483648}
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 string = "foo bar"
    v13 := GenS1{F0: v8, F1: -93, F2: int32(v8)}
    _ = v12
    _ = v13
    _ = v8
    return 65535
}
func genFn1(p0 uint, p1 int) int32 {
    if false {
        p0 += (p0 * (^(p0)))
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 uint8 = 2
            _ = fi14
            _ = v15
        }
        switch "a" {
        case "":
            for fi16 := 0; fi16 < 9; fi16++ {
                fmt.Println(p1, p0, fi16)
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 int8 = -43
                    v19 := []int{fi17, (fi16 ^ (fi16 | fi17)), -645, 32768}
                    if ("hello" == "foo bar") {
                        fmt.Println(p0)
                    }
                    switch "hello" {
                    case "hello":
                        if false {
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            if (false || true) {
                                p1 -= (-(fi16))
                                v19 = append(v19, (fi17 ^ fi16))
                            }
                        }
                    default:
                        fmt.Println(fi16, fi16)
                    }
                    if len(v19) > 0 {
                        v19[0] = (p1 ^ fi17)
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 uint64 = 42
                    v22 := GenS1{F0: 2, F1: 0, F2: -223}
                    v23 := []int{p1, -1}
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi16
            }
        case "hello":
            v24 := GenS2{F0: 375, F1: 42}
            _ = v24
        default:
            p0 = p0
        }
    } else {
        if (1e10 <= 1.5) {
            switch p1 {
            case 128:
                v25 := GenS2{F0: 847, F1: 0}
                _ = v25
            case 32767:
                var v26 uint32 = 263
                _ = v26
            case 1:
                if false {
                    switch 32768 {
                    case 32767:
                        fmt.Println(p0)
                    }
                } else {
                    var v27 string = ""
                    switch "hello" {
                    case "foo bar":
                        var v28 uint64 = 2
                        _ = v28
                    case "\n\t":
                        for fi29 := 0; fi29 < 8; fi29++ {
                            for fi30 := 0; fi30 < 16; fi30++ {
                                v27 = v27
                                var v31 int64 = 0
                                var v32 bool = (515 <= 764)
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            switch "\n\t" {
                            case "":
                                var v33 string = v27
                                _ = v33
                            case "a":
                                var v34 string = v27
                                _ = v34
                            default:
                                p1 -= fi29
                            }
                            for fi35 := 0; fi35 < 9; fi35++ {
                                switch "" {
                                case "a":
                                    v27 = v27
                                default:
                                    v27 = (v27 + (("hello" + v27) + "hello"))
                                }
                                var v36 uint = p0
                                var v37 uint64 = 18446744073709551615
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            _ = fi29
                        }
                    }
                    _ = v27
                }
            }
            v38 := GenS2{F0: 709, F1: 2}
            _ = v38
        }
    }
    p1 -= 127
    v39 := GenS2{F0: genFn0(), F1: 27}
    var v40 uint8 = 127
    _ = p0
    _ = p1
    _ = v39
    _ = v40
    return int32(v40)
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    var v3 uint64 = 18446744073709551615
    var v41 float64 = float64(((((1.5 - (-2.25 + 0.0)) - (float32(0.0) + float32(1e10))) * (1.5 - -2.25)) / 1e10))
    if v1 {
        v1 = true
        v42 := GenS2{F0: 1, F1: 22}
        v43 := GenS2{F0: 4294967295, F1: 255}
        _ = v42
        _ = v43
    }
    if (42 >= 0) {
        switch "hello" {
        case "\n\t":
            var v44 int16 = -608
            _ = v44
        case "foo bar":
            fmt.Println(v1, v41)
        default:
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                v45 := GenS2{F0: 42, F1: 0}
                _ = v45
            default:
                switch 9223372036854775807 {
                case 1:
                    v46 := []int{32768, -2147483648}
                    _ = v46
                case 128:
                    switch (("foo bar" + (("\n\t" + "") + "foo bar")) + (("hello" + "foo bar") + "")) {
                    case "foo bar":
                        var v47 int8 = 42
                        _ = v47
                    }
                default:
                    var v48 int8 = -1
                    _ = v48
                }
            }
        }
    }
    delete(v2, 65535)
    v49 := v2[924]
    v50 := GenS3{F0: 65535, F1: 42}
    fmt.Println(v41)
    v41 += (3.14159 * float64((1e10 * -2.25)))
    v1 = (42 >= 209)
    switch "a" {
    case "":
        switch 127 {
        case 32768:
            v51 := GenS4{F0: 853, F1: 1}
            _ = v51
        default:
            fmt.Println(v1, v49)
        }
    case "foo bar":
        var v52 int64 = 149
        _ = v52
    case "\n\t":
        var v53 uint16 = 619
        _ = v53
    }
    v41 = float64(0.1)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v49
    _ = v50
}
