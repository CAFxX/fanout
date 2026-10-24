package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint64 }
func genFn0() int16 {
    var v6 int64 = 42
    fmt.Println(v6)
    _ = v6
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v4 bool = false
    var v5 int8 = 100
    var v7 string = ""
    fmt.Println(v7, v5)
    var v8 uint8 = 255
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v7)
            v11 := GenS1{F0: 42, F1: 178, F2: -1}
            v5 += int8(v1)
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    if (65535 < 2) {
        v1 = (v1 ^ v1)
        switch 2147483647 {
        case -234:
            v4 = false
        case 2147483647:
            fmt.Println(v5, v8)
        case -863:
            v12 := v2[823]
            _ = v12
        }
    } else {
        var v13 uint16 = 800
        for fi14 := 0; fi14 < 7; fi14++ {
            if len(v3) > 0 {
                v3[0] = (((v7 + v7) + v7) + (v7 + v7))
            }
            for fi15 := 0; fi15 < 10; fi15++ {
                v2[fi14] = 924
                v16 := GenS2{F0: genFn0(), F1: 316, F2: uint(v8)}
                fmt.Println(v8, fi14, fi14)
                fmt.Println(v4)
                var v17 int8 = (-(v5))
                if v4 {
                    if false {
                        v13 -= v13
                        v18 := GenS3{F0: 18446744073709551615}
                        var v19 string = "foo bar"
                        _ = v18
                        _ = v19
                    } else {
                        v7 = ("" + v7)
                        v7 = v7
                    }
                    if false {
                        var v20 int16 = genFn0()
                        _ = v20
                    } else {
                        v21 := GenS4{F0: (v17 + int8(v8))}
                        fmt.Println(v8, fi14, v8)
                        _ = v21
                    }
                    if (!(!(v8 >= v8))) {
                        var v22 uint32 = 362
                        v23 := v2[32768]
                        v8 = 1
                        var v24 string = "\n\t"
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS5{F0: 255, F1: 42, F2: 2147483648}
                        v26 := GenS6{F0: 32767, F1: 18446744073709551615}
                        _ = v25
                        _ = v26
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (9223372036854775807 > -9223372036854775808) {
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                    } else {
                        switch fi14 {
                        case 1:
                            if v4 {
                                v1 = 0
                                var v27 float64 = (((1.5 - (3.14159 - ((1e10 * 0.0) + (1e10 / 0.1)))) - 0.1) - 1e10)
                                switch fi15 {
                                case -1:
                                    var v28 float64 = v27
                                    _ = v28
                                case 1:
                                    v17 = 100
                                }
                                _ = v27
                            }
                        case 32768:
                            v29 := v2[fi15]
                            _ = v29
                        case -920:
                            var v30 uint8 = (v8 / ((v8) | 1))
                            _ = v30
                        }
                        v4 = ((636 > 848) || v4)
                        v17 = (-(v5))
                        v31 := GenS4{F0: (v5 + v5)}
                        var v32 float32 = float32(float64(0.0))
                        _ = v31
                        _ = v32
                    }
                }
                switch "hello" {
                case "hello":
                    var v33 float64 = 1e10
                    _ = v33
                case "\n\t":
                    var v34 int64 = -4
                    _ = v34
                }
                var v35 uint = 65535
                v7 = (("foo bar" + "") + "")
                _ = fi15
                _ = v16
                _ = v17
                _ = v35
            }
            _ = fi14
        }
        _ = v13
    }
    v8 = v8
    var v36 uint = 2
    var v37 int = 784
    if (uint16(v36) >= 65535) {
        var v38 uint32 = 0
        _ = v38
    } else {
        var v39 float64 = 1.5
        v7 = v7
        v5 += (-(v5))
        _ = v39
    }
    var v40 int64 = -566
    v4 = v4
    fmt.Println(v8, v5)
    v1 = (v1 / ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v7
    _ = v8
}
