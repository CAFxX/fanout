package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint16; F1 bool; F2 int8 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 int8; F2 int16 }

func main() {
    var v1 int8 = -1
    var v2 uint32 = 0
    var v3 uint16 = 88
    var v4 int64 = 42
    if true {
        if false {
            if (("hello" + "hello") != "hello") {
                var v5 float32 = 0.0
                switch 4294967295 {
                case -762:
                    if true {
                        if (true || (true && true)) {
                            v6 := GenS1{F0: 982, F1: v1, F2: 0}
                            var v7 int32 = 42
                            _ = v6
                            _ = v7
                        }
                        if (!true) {
                            v2 = v2
                            v8 := []int{128, -2147483648}
                            _ = v8
                        }
                        for fi9 := 0; fi9 < 11; fi9++ {
                            v4 += (v4 & (v4 | v4))
                            v2 += v2
                            fmt.Println(fi9, v2)
                            _ = fi9
                        }
                    } else {
                        v3 -= (v3 % ((v3) | 1))
                        switch -1 {
                        case 4294967295:
                            v5 = (1.5 + v5)
                        case 0:
                            for fi10 := 0; fi10 < 9; fi10++ {
                                var v11 bool = false
                                _ = fi10
                                _ = v11
                            }
                        default:
                            switch -80 {
                            case 0:
                                v12 := []int{127}
                                _ = v12
                            case 741:
                                var v13 int32 = 2147483647
                                _ = v13
                            }
                        }
                        v14 := GenS2{F0: v3, F1: true, F2: 1}
                        _ = v14
                    }
                }
                _ = v5
            } else {
                v15 := []int{969, 129}
                _ = v15
            }
        } else {
            v3 += (v3 % ((255) | 1))
            if false {
                var v16 uint64 = 18446744073709551615
                switch -5 {
                case 128:
                    var v17 bool = (!true)
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 20; fi18++ {
                        switch int(v4) {
                        case 9223372036854775807:
                            for fi19 := 0; fi19 < 9; fi19++ {
                                fmt.Println(v16, v16)
                                var v20 int64 = 1
                                _ = fi19
                                _ = v20
                            }
                        }
                        _ = fi18
                    }
                }
                _ = v16
            } else {
                var v21 uint = 1
                _ = v21
            }
            fmt.Println(v4, v1, v3)
        }
        v22 := GenS3{F0: -58, F1: (v3 & 2)}
        _ = v22
    }
    if (42 < 42) {
        v23 := GenS4{F0: (-(v1))}
        fmt.Println(v4, v4, v3)
        _ = v23
    }
    var v24 uint32 = 769
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS5{F0: 127, F1: v1, F2: 912}
        if (v1 == 42) {
            v27 := []int{-673}
            if false {
                if false {
                    v28 := GenS4{F0: -12}
                    if true {
                        if (!(0 < 255)) {
                            for fi29 := 0; fi29 < 13; fi29++ {
                                fmt.Println(v24, v2)
                                v4 = (v4 % (((v4 >> ((1) & 63))) | 1))
                                _ = fi29
                            }
                            v2 -= v24
                        }
                    }
                    _ = v28
                }
                fmt.Println(v3)
                if len(v27) > 0 {
                    v27[0] = (fi25 & fi25)
                }
            }
            v1 -= 100
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        v1 = v1
        for fi31 := 0; fi31 < 2; fi31++ {
            fmt.Println(fi31)
            var v32 int64 = (v4 & (v4 - v4))
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    v24 -= v24
    for fi33 := 0; fi33 < 12; fi33++ {
        fmt.Println(v4, v1)
        var v34 int8 = 0
        _ = fi33
        _ = v34
    }
    var v35 int64 = v4
    var v36 uint = 487
    v37 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
