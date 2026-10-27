package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint8 }
func genFn0() int16 {
    v7 := GenS1{F0: 65535}
    var v8 uint8 = 255
    v9 := []int{}
    v10 := GenS1{F0: 0}
    if (-126 < -827) {
        v8 += v8
        for fi11 := 0; fi11 < 7; fi11++ {
            if true {
                var v12 int32 = -336
                v8 += v8
                _ = v12
            } else {
                v8 += v8
                v13 := GenS2{F0: 4294967295, F1: 18446744073709551615, F2: false}
                _ = v13
            }
            _ = fi11
        }
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 760
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    var v5 uint64 = 18446744073709551615
    var v6 float32 = float32((((float64(0.0) * 1e10) * (1e10 - (1.5 / 1.0))) - -2.25))
    v1 = ((("a" + (((("\n\t" + v1) + "") + "foo bar") + (("hello" + "a") + "a"))) + v1) + v1)
    switch 15 {
    case -931:
        v14 := GenS1{F0: 1}
        _ = v14
    }
    v15 := GenS2{F0: 47, F1: v5, F2: false}
    var v16 int16 = -729
    v1 = v1
    v3[42] = -2147483648
    fmt.Println(v1, v16)
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 int64 = -33
        var v19 int16 = (v16 + genFn0())
        _ = fi17
        _ = v18
        _ = v19
    }
    if (((v6 / (v6 - v6)) - 3.14159) > (v6 - (v6 * v6))) {
        fmt.Println(v16, v5, v5)
        switch 128 {
        case -1:
            var v20 int8 = 1
            _ = v20
        case 32767:
            v6 = (float32((float64(v6) / 1.0)) - (-2.25 / (-2.25 * v6)))
        default:
            fmt.Println(v16)
        }
    }
    var v21 int = int(v5)
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            fmt.Println(v16)
            _ = fi23
        }
        var v24 int16 = v16
        _ = fi22
        _ = v24
    }
    switch v21 {
    case 127:
        fmt.Println(v16, v1)
    case 2147483647:
        v25 := GenS3{F0: 108}
        _ = v25
    case 1:
        var v26 int16 = (v16 & 0)
        _ = v26
    }
    var v27 uint = 1
    for fi28 := 0; fi28 < 20; fi28++ {
        fmt.Println(v1)
        v29 := GenS2{F0: (v5 << (((^(v27))) & 63)), F1: (v5 + 42), F2: false}
        fmt.Println(v21, v16, v16)
        _ = fi28
        _ = v29
    }
    v30 := GenS1{F0: 4294967295}
    for fi31 := 0; fi31 < 17; fi31++ {
        if (!(!(v27 >= v27))) {
            var v32 float64 = 3.14159
            var v33 uint32 = 318
            fmt.Println(v5, v27)
            _ = v32
            _ = v33
        }
        switch ((v1 + v1) + ((v1 + v1) + ("foo bar" + ((v1 + "foo bar") + ("" + v1))))) {
        case "foo bar":
            for fi34 := 0; fi34 < 6; fi34++ {
                v21 = (v21 & int(v27))
                if false {
                    switch -863 {
                    case -632:
                        var v35 uint16 = 65535
                        _ = v35
                    case 1:
                        if (false && ((v1 != (v1 + (("foo bar" + "foo bar") + "hello"))) || (112 > 0))) {
                            v27 -= (^(v27))
                            v16 = v16
                        } else {
                            switch v1 {
                            case "foo bar":
                                if ((v16 - v16) <= (v16 >> (((v27 | v27)) & 63))) {
                                    fmt.Println(v16)
                                }
                            }
                            _ = len(v4)
                            var v36 int16 = genFn0()
                            _ = v36
                        }
                    default:
                        fmt.Println(v5)
                    }
                } else {
                    var v37 uint64 = (^(v5))
                    v38 := GenS4{F0: 2, F1: -128, F2: uint8(v27)}
                    var v39 uint32 = 2147483647
                    _ = v37
                    _ = v38
                    _ = v39
                }
                var v40 uint16 = 1
                _ = fi34
                _ = v40
            }
        case "\n\t":
            v41 := GenS3{F0: -2147483648}
            _ = v41
        default:
            v5 = v5
        }
        var v42 bool = true
        _ = fi31
        _ = v42
    }
    _ = len(v4)
    v21 = (v21 & (v21 << (((^(v27))) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
