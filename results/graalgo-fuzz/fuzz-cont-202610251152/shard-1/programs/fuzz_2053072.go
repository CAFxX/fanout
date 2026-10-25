package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int, p1 int16) int {
    var v7 float64 = 1.5
    v8 := GenS1{F0: 2}
    var v9 int32 = -937
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: 1}
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            if (42 != 1) {
                p0 = fi10
            } else {
                v14 := GenS1{F0: 0}
                _ = v14
            }
            _ = fi13
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 string) int64 {
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(fi15, p0)
        _ = fi15
    }
    var v16 bool = false
    _ = p0
    _ = v16
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("" + "a")
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    var v5 int16 = 32767
    var v6 float32 = 0.1
    switch 4294967295 {
    case -191:
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 bool = (65535 >= 65535)
            _ = fi17
            _ = v18
        }
    case 9223372036854775807:
        v5 -= (-(v5))
    case 32768:
        switch 0 {
        case -153:
            if (float64(((-2.25 * v6) / (0.0 * (v6 + (float32(float64(v6)) / ((float32(3.14159) / float32((0.1 + 1.5))) / 0.1)))))) < float64(v6)) {
                if (2147483648 != 2147483647) {
                    v19 := []int{861}
                    var v20 float32 = (float32(0.0) + 1.5)
                    var v21 float32 = (((v20 - 0.0) / (v20 + ((v20 + float32(1e10)) - v6))) / v20)
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    if true {
                        var v22 int = 4294967295
                        _ = v22
                    } else {
                        v23 := []int{}
                        v2 = v2
                        v24 := GenS2{F0: 1, F1: 42, F2: 42}
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS1{F0: 42}
                    _ = v25
                }
                if (18446744073709551615 <= 0) {
                    if (4294967295 == 2147483647) {
                        if ((((v5 << ((243) & 63)) < (v5 >> ((18446744073709551615) & 63))) && false) || (384 > 1)) {
                            for fi26 := 0; fi26 < 4; fi26++ {
                                v6 += float32((-2.25 - (float64(0.0) + 0.1)))
                                v5 += v5
                                fmt.Println(v5, v6, fi26)
                                _ = fi26
                            }
                            v6 += v6
                        }
                        v5 = (v5 >> ((237) & 63))
                        v27 := []int{780, genFn0(genFn0(42, 1), 0)}
                        _ = v27
                    } else {
                        if (((v2 + v2) + "") > v2) {
                            fmt.Println(v5)
                            v28 := GenS1{F0: 42}
                            v1[42] = v2
                            _ = v28
                        } else {
                            fmt.Println(v2, v5, v2)
                        }
                    }
                    switch v2 {
                    case "foo bar":
                        if (486 >= 18446744073709551615) {
                            v5 = -1
                            v29 := GenS2{F0: (v5 ^ 660), F1: 81, F2: 4294967295}
                            _ = v29
                        }
                    case "":
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v31 := []int{genFn0((fi30 + -9223372036854775808), (v5 / (((v5 / (((v5 * int16(fi30))) | 1))) | 1)))}
                            v5 += v5
                            switch "" {
                            case "foo bar":
                                v2 = v2
                            case "a":
                                var v32 uint16 = 1
                                _ = v32
                            }
                            _ = fi30
                            _ = v31
                        }
                    case "\n\t":
                        v6 = ((((v6 + 1e10) * 3.14159) + (-2.25 - (v6 - (1e10 - v6)))) / v6)
                    }
                    var v33 int16 = (v5 * v5)
                    _ = v33
                }
            }
        }
    default:
        switch v2 {
        case "\n\t":
            v5 = v5
        case "foo bar":
            v5 = (-(v5))
        case "a":
            if true {
                v6 -= ((1e10 * (float32(0.0) / (((0.0 - (v6 * float32(3.14159))) / v6) / (3.14159 / v6)))) + 1e10)
            } else {
                v34 := GenS1{F0: 333}
                _ = v34
            }
        default:
            v2 = ("hello" + "a")
        }
    }
    v2 = (v2 + v2)
    switch v2 {
    case "foo bar":
        if false {
            var v35 uint64 = 18446744073709551615
            if false {
                if (!(0 < genFn1((("foo bar" + v2) + v2)))) {
                    v6 = (v6 / v6)
                    if false {
                        v5 += (^(v5))
                    } else {
                        for fi36 := 0; fi36 < 1; fi36++ {
                            v35 = 0
                            _ = fi36
                        }
                        v37 := GenS2{F0: (v5 | (v5 * -833)), F1: 0, F2: 0}
                        _ = v37
                    }
                }
                var v38 int8 = 0
                _ = v38
            }
            _ = v35
        }
    default:
        v39 := GenS3{F0: 825, F1: 2}
        _ = v39
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        v6 = ((v6 - v6) + -2.25)
        _ = fi40
    }
    _ = len(v4)
    var v41 uint8 = 2
    v42 := GenS3{F0: 95, F1: 998}
    for fi43 := 0; fi43 < 6; fi43++ {
        _ = len(v4)
        for fi44 := 0; fi44 < 6; fi44++ {
            var v45 int8 = -1
            _ = fi44
            _ = v45
        }
        var v46 int64 = -901
        _ = fi43
        _ = v46
    }
    fmt.Println(v6)
    v47 := GenS4{F0: v2, F1: true}
    var v48 float32 = (v6 - v6)
    var v49 uint8 = v41
    fmt.Println(v48)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
