package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int8; F2 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 int64, p1 int8, p2 int) uint {
    p1 += (p1 % (((^(p1))) | 1))
    p1 -= (p1 << ((896) & 63))
    p1 = (p1 + -9)
    switch "hello" {
    case "\n\t":
        var v4 string = "\n\t"
        _ = v4
    case "hello":
        v5 := GenS1{F0: int16(p2), F1: 255}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 0
    var v3 bool = true
    if v3 {
        if v3 {
            v2 -= v2
            var v6 uint64 = 4
            for fi7 := 0; fi7 < 4; fi7++ {
                if (3.14159 >= (float32((-2.25 + (1e10 * 0.1))) * 1e10)) {
                    v2 = (^(v2))
                    var v8 int16 = v1
                    if true {
                        v1 = v8
                        v3 = (!v3)
                        fmt.Println(v6, v2, v3)
                    } else {
                        var v9 int8 = 59
                        var v10 bool = false
                        v9 -= 1
                        _ = v10
                        _ = v9
                    }
                    _ = v8
                } else {
                    var v11 uint64 = 1
                    v2 += 42
                    _ = v11
                }
                _ = fi7
            }
            _ = v6
        } else {
            var v12 int32 = -2147483648
            v13 := GenS2{F0: 674, F1: 0, F2: 100}
            _ = v12
            _ = v13
        }
        v14 := []int{-1, 32768}
        _ = v14
    }
    fmt.Println(v2, v2, v1)
    var v15 int16 = 1
    var v16 int16 = v15
    v1 -= v1
    v15 -= (v16 + v16)
    var v17 int = 4294967295
    v2 = 219
    fmt.Println(v3)
    v18 := []int{v17, 128}
    var v19 uint32 = 2
    for fi20 := 0; fi20 < 17; fi20++ {
        if v3 {
            switch "" {
            case "hello":
                switch v17 {
                case 128:
                    switch fi20 {
                    case 1:
                        for fi21 := 0; fi21 < 20; fi21++ {
                            var v22 int64 = -9223372036854775808
                            var v23 int64 = 1
                            if false {
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                var v24 uint8 = (v2 + v2)
                                if v3 {
                                    v22 = (-(v22))
                                    fmt.Println(v19, v23)
                                } else {
                                    v25 := GenS2{F0: v19, F1: 100, F2: 0}
                                    _ = v25
                                }
                                _ = v24
                            } else {
                                switch ("" + "a") {
                                case "\n\t":
                                    fmt.Println(v15)
                                }
                                if v3 {
                                    if v3 {
                                        var v26 uint8 = v2
                                        var v27 int8 = 0
                                        var v28 int16 = v1
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                } else {
                                    var v29 int8 = -1
                                    var v30 float64 = 0.0
                                    _ = v29
                                    _ = v30
                                }
                                var v31 string = "hello"
                                _ = v31
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    default:
                        _ = len(v18)
                    }
                case -496:
                    for fi32 := 0; fi32 < 3; fi32++ {
                        if len(v18) > 0 {
                            v18[0] = 42
                        }
                        _ = fi32
                    }
                case 127:
                    for fi33 := 0; fi33 < 18; fi33++ {
                        fmt.Println(v16)
                        fmt.Println(v17, v19, v16)
                        _ = fi33
                    }
                }
            case "foo bar":
                switch "a" {
                case "\n\t":
                    switch v17 {
                    case -307:
                        v19 += 741
                    case 639:
                        v16 = (^(v1))
                    case -620:
                        switch (fi20 * (-(fi20))) {
                        case -1:
                            v34 := GenS3{F0: 18446744073709551615}
                            _ = v34
                        default:
                            var v35 uint8 = v2
                            _ = v35
                        }
                    }
                case "hello":
                    fmt.Println(v15, v1)
                case "foo bar":
                    v36 := GenS2{F0: v19, F1: 1, F2: 127}
                    _ = v36
                }
            }
        } else {
            v3 = v3
            var v37 uint64 = 781
            v16 = (-(v16))
            _ = v37
        }
        v38 := GenS2{F0: 411, F1: 127, F2: 24}
        var v39 uint32 = v19
        _ = fi20
        _ = v38
        _ = v39
    }
    v3 = (!(18446744073709551615 != 4294967295))
    v17 = -345
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
