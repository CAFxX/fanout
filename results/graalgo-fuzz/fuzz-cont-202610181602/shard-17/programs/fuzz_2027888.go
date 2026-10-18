package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int64; F1 int32; F2 int64 }
type GenS4 struct { F0 int; F1 uint }
func genFn0(p0 uint64) int8 {
    v7 := []int{-843, 9223372036854775807, 9223372036854775807}
    _ = len(v7)
    var v8 int16 = 0
    v9 := GenS1{F0: ((-1 < 127) || (1 == 362)), F1: -123, F2: 0}
    var v10 uint64 = (^(p0))
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -100
}
func genFn1(p0 int16, p1 uint8, p2 uint16) uint64 {
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 string = ""
        var v13 int64 = -1
        _ = fi11
        _ = v12
        _ = v13
    }
    if (false && true) {
        var v14 int8 = 42
        var v15 bool = ((((1.5 + (3.14159 + 3.14159)) * 0.0) / ((3.14159 - -2.25) - 0.1)) != (3.14159 * -2.25))
        p0 = (p0 - (p0 << ((4294967295) & 63)))
        _ = v14
        _ = v15
    }
    if (633 < 1) {
        if (((1e10 + 1.5) == float32(1e10)) || false) {
            v16 := []int{-9223372036854775808, 127}
            _ = v16
        } else {
            for fi17 := 0; fi17 < 6; fi17++ {
                fmt.Println(p0)
                v18 := GenS1{F0: (320 < 18446744073709551615), F1: 364, F2: (p1 << ((0) & 63))}
                var v19 uint32 = 2147483647
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    }
    var v20 int = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 1
}

func main() {
    var v1 string = ("hello" + "a")
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    v4 := []int{-1, 127}
    var v5 int = 2147483647
    var v6 bool = ((779 <= 2147483648) && false)
    switch v1 {
    case "":
        _ = len(v3)
    default:
        v6 = false
    }
    var v21 int32 = 2147483647
    fmt.Println(v21, v21, v1)
    var v22 int = 127
    v23 := GenS1{F0: ((!((-32768 == 1000) && v6)) && v6), F1: 42, F2: 128}
    v1 = v1
    if true {
        v24 := GenS1{F0: (v6 || v6), F1: -1, F2: 255}
        v25 := GenS2{F0: -1, F1: true}
        _ = v24
        _ = v25
    }
    fmt.Println(v6, v1)
    if false {
        fmt.Println(v5, v1, v21)
    }
    fmt.Println(v6)
    fmt.Println(v6)
    var v26 uint64 = 4294967295
    if v6 {
        v3[1] = (v5 + (-(v22)))
        var v27 float32 = 1.5
        _ = v27
    }
    fmt.Println(v1, v21, v6)
    var v28 string = "hello"
    for fi29 := 0; fi29 < 11; fi29++ {
        if (0 < 135) {
            _ = len(v3)
            _ = len(v3)
            v6 = (true || (v6 && true))
        } else {
            switch v1 {
            case "a":
                v4 = append(v4, fi29)
            case "hello":
                var v30 int = (v22 * -1)
                _ = v30
            case "foo bar":
                for fi31 := 0; fi31 < 9; fi31++ {
                    if false {
                        var v32 int32 = v21
                        fmt.Println(fi29, v26, fi29)
                        _ = v32
                    }
                    _ = fi31
                }
            default:
                for fi33 := 0; fi33 < 4; fi33++ {
                    v21 += v21
                    var v34 uint32 = 2147483647
                    if ((v1 + v28) >= "") {
                        delete(v3, v34)
                        v3[(v34 >> ((2) & 63))] = fi29
                    } else {
                        _ = len(v4)
                        var v35 int32 = v21
                        v36 := GenS2{F0: fi29, F1: v6}
                        _ = v35
                        _ = v36
                    }
                    _ = fi33
                    _ = v34
                }
            }
            var v37 int32 = 79
            var v38 int = (v5 + (v22 ^ 1))
            _ = v37
            _ = v38
        }
        _ = fi29
    }
    v4 = append(v4, 32767)
    switch v5 {
    case 127:
        for fi39 := 0; fi39 < 3; fi39++ {
            if true {
                for fi40 := 0; fi40 < 19; fi40++ {
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                    var v41 uint8 = 0
                    if (v21 == (^(v21))) {
                        if (0 > 42) {
                            v42 := v3[42]
                            fmt.Println(v1, v22, fi40)
                            var v43 uint32 = 817
                            v21 += (^(v21))
                            _ = v42
                            _ = v43
                        } else {
                            v44 := GenS3{F0: -9223372036854775808, F1: (v21 % ((int32(v22)) | 1)), F2: 0}
                            var v45 bool = true
                            _ = v44
                            _ = v45
                        }
                        v3[42] = v5
                    } else {
                        v46 := GenS3{F0: 9223372036854775807, F1: (v21 << ((2) & 63)), F2: 1}
                        if v6 {
                            v47 := GenS3{F0: -592, F1: v21, F2: 0}
                            v1 = v28
                            _ = v47
                        } else {
                            v48 := GenS1{F0: v6, F1: 1, F2: v41}
                            _ = v48
                        }
                        v49 := GenS4{F0: (fi40 * (v5 - 42)), F1: 216}
                        _ = v46
                        _ = v49
                    }
                    _ = fi40
                    _ = v41
                }
            }
            _ = fi39
        }
    case 128:
        v22 = 4294967295
    default:
        var v50 float64 = (3.14159 - (-2.25 - ((1e10 - -2.25) + (0.1 - (3.14159 * 0.0)))))
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
