package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 string }
func genFn0() uint8 {
    v6 := []int{}
    var v7 uint64 = 874
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = (0.0 * -2.25)
    var v3 float64 = 1e10
    var v4 string = "foo bar"
    var v5 uint = 255
    v8 := GenS1{F0: 255, F1: v5, F2: v5}
    v2 = float64(((-2.25 + float32(-2.25)) + ((-2.25 * -2.25) / float32(1e10))))
    v9 := []int{42, 366}
    if (127 > -1) {
        switch 42 {
        case 4294967295:
            var v10 uint64 = 2
            _ = v10
        case 1:
            v4 = v4
        case 32768:
            fmt.Println(v5, v5)
        default:
            for fi11 := 0; fi11 < 18; fi11++ {
                switch -9223372036854775808 {
                case 934:
                    if false {
                        for fi12 := 0; fi12 < 15; fi12++ {
                            v2 -= (float64(float32(v3)) + v3)
                            var v13 int16 = -1
                            _ = fi12
                            _ = v13
                        }
                        if false {
                            v14 := GenS1{F0: 128, F1: v5, F2: v5}
                            _ = v14
                        } else {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v3)
                            case "hello":
                                v2 += 1.5
                            case "foo bar":
                                if (!false) {
                                    if len(v9) > 0 {
                                        v9[0] = fi11
                                    }
                                    var v15 int32 = -976
                                    _ = v15
                                }
                            default:
                                for fi16 := 0; fi16 < 13; fi16++ {
                                    v1 = -1
                                    _ = fi16
                                }
                            }
                            fmt.Println(v4)
                        }
                        var v17 string = "\n\t"
                        _ = v17
                    } else {
                        var v18 int8 = 27
                        v18 = (v18 + (v18 >> ((v5) & 63)))
                        fmt.Println(v3, v5, fi11)
                        _ = v18
                    }
                case 32767:
                    if (v5 <= 503) {
                        fmt.Println(fi11, fi11)
                    }
                }
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 int = fi11
                    if len(v9) > 0 {
                        v9[0] = (fi11 % ((fi11) | 1))
                    }
                    v2 += -2.25
                    _ = fi19
                    _ = v20
                }
                v5 += v5
                _ = fi11
            }
        }
        if false {
            if len(v9) > 0 {
                v9[0] = 9223372036854775807
            }
        } else {
            for fi21 := 0; fi21 < 4; fi21++ {
                if len(v9) > 0 {
                    v9[0] = fi21
                }
                _ = fi21
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                for fi23 := 0; fi23 < 8; fi23++ {
                    switch fi22 {
                    case 400:
                        v24 := GenS2{F0: 255, F1: -38}
                        _ = v24
                    case 42:
                        switch "hello" {
                        case "\n\t":
                            switch (-(fi22)) {
                            case 218:
                                fmt.Println(v1, v4)
                            }
                        default:
                            v4 = v4
                        }
                    case -1:
                        v25 := GenS3{F0: (v4 + ((v4 + "\n\t") + v4))}
                        _ = v25
                    default:
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    _ = fi23
                }
                _ = fi22
            }
        }
    }
    v3 = v3
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 int64 = -1
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 bool = false
            _ = fi28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v4)
    v30 := GenS2{F0: uint8(v5), F1: -128}
    v2 -= ((((v2 / ((float64(3.14159) * v3) + v2)) * float64(((1.5 - 0.0) + -2.25))) + 0.0) + v3)
    var v31 uint32 = 2147483648
    var v32 uint64 = 0
    v4 = ("foo bar" + ((((v4 + v4) + "\n\t") + v4) + v4))
    var v33 bool = false
    var v34 int16 = -32768
    v31 -= 65535
    for fi35 := 0; fi35 < 16; fi35++ {
        for fi36 := 0; fi36 < 19; fi36++ {
            var v37 uint16 = 1
            var v38 float64 = 3.14159
            fmt.Println(v2, v1, v33)
            _ = fi36
            _ = v37
            _ = v38
        }
        v32 -= (v32 << ((4294967295) & 63))
        if (v34 <= 32767) {
            fmt.Println(v5)
        } else {
            fmt.Println(v34)
            var v39 uint32 = (v31 | v31)
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v39
        }
        _ = fi35
    }
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
