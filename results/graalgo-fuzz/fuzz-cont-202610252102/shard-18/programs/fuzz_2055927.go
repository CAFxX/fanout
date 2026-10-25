package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int16; F2 uint32 }
func genFn0(p0 int8) int32 {
    var v4 uint16 = 32767
    var v5 uint = 255
    p0 -= -1
    var v6 int = -935
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1() int16 {
    var v7 uint8 = 127
    var v8 bool = false
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = ((3.14159 / 1.5) * float32(1.5))
    v3 := make(map[int]string)
    v2 = 1.5
    for fi9 := 0; fi9 < 10; fi9++ {
        if true {
            v2 = (v2 - v2)
            if (!(true && true)) {
                v1[(fi9 | fi9)] = ("foo bar" + "")
                for fi10 := 0; fi10 < 19; fi10++ {
                    if ((false || (("" + ("hello" + ("" + ""))) != ("hello" + "\n\t"))) && (false && (235 == 4294967295))) {
                        v2 -= (v2 / -2.25)
                        v11 := GenS1{F0: 127, F1: 300, F2: 571}
                        if (-2147483648 > 1) {
                            v2 = 1.5
                            v2 = v2
                            if true {
                                fmt.Println(fi9, v2)
                                v2 += v2
                                v12 := GenS1{F0: -1, F1: 2, F2: 2}
                                var v13 float64 = 1e10
                                _ = v12
                                _ = v13
                            }
                        }
                        v14 := GenS2{F0: 1, F1: 923, F2: 1}
                        _ = v11
                        _ = v14
                    }
                    _ = fi10
                }
                v15 := []int{}
                _ = v15
            } else {
                if false {
                    if (-1 == 0) {
                        var v16 uint32 = 166
                        fmt.Println(v16)
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 18; fi17++ {
                            v18 := GenS3{F0: 1}
                            var v19 int16 = 42
                            v20 := []int{}
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    }
                } else {
                    v21 := v1[(fi9 + -9223372036854775808)]
                    _ = v21
                }
            }
            v22 := GenS4{F0: (fi9 & fi9), F1: -1, F2: 547}
            _ = v22
        }
        var v23 int64 = 9223372036854775807
        _ = fi9
        _ = v23
    }
    fmt.Println(v2)
    v2 -= v2
    var v24 uint16 = 488
    var v25 uint64 = 1
    fmt.Println(v25)
    switch 174 {
    case 128:
        for fi26 := 0; fi26 < 13; fi26++ {
            v27 := []int{}
            _ = fi26
            _ = v27
        }
    case 0:
        var v28 uint32 = 42
        _ = v28
    default:
        v25 = v25
    }
    fmt.Println(v24)
    v24 += uint16(v25)
    if (true || ((0 == 4294967295) || (-128 == 42))) {
        if (!(false || false)) {
            var v29 uint8 = 127
            _ = v29
        } else {
            var v30 int = -924
            v31 := []int{}
            _ = v30
            _ = v31
        }
        var v32 int = -9223372036854775808
        _ = v32
    } else {
        if false {
            if (!true) {
                var v33 uint64 = 767
                v1[32767] = ("hello" + "\n\t")
                switch 42 {
                case 32767:
                    v34 := GenS1{F0: 1, F1: 329, F2: 2}
                    _ = v34
                case 32768:
                    var v35 int8 = 1
                    _ = v35
                }
                _ = v33
            }
            fmt.Println(v25, v25, v2)
            for fi36 := 0; fi36 < 10; fi36++ {
                v25 -= v25
                _ = fi36
            }
        }
        v2 += v2
        v25 = 18446744073709551615
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        fmt.Println(v2)
        _ = fi37
    }
    fmt.Println(v24, v2, v25)
    var v38 float32 = v2
    v39 := GenS3{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v39
}
