package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 4294967295
    var v3 uint16 = 255
    var v4 int32 = -1
    switch 2147483647 {
    case -306:
        var v5 int64 = -530
        _ = v5
    case 32768:
        if false {
            var v6 int32 = v4
            _ = v6
        }
    case 679:
        if (2 > 127) {
            v7 := GenS1{F0: 0, F1: -687}
            v8 := GenS2{F0: -1}
            _ = v7
            _ = v8
        }
    default:
        if (409 != 0) {
            fmt.Println(v3, v3)
            v3 -= (v3 << ((516) & 63))
            _ = len(v1)
        } else {
            fmt.Println(v2, v3)
            fmt.Println(v3)
            v4 = 42
        }
    }
    if true {
        fmt.Println(v2, v2)
    } else {
        switch 32767 {
        case 32767:
            var v9 uint16 = (v3 >> ((uint(v3)) & 63))
            _ = v9
        case -1:
            v4 += -2147483648
        case 0:
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 float32 = (float32((1.5 / (1e10 - (float64(-2.25) * 3.14159)))) * 1.5)
                if (1 >= 1) {
                    v4 -= -1
                }
                _ = fi10
                _ = v11
            }
        default:
            if (false || true) {
                fmt.Println(v3)
                fmt.Println(v3, v4)
            } else {
                fmt.Println(v4, v4)
                v3 = (v3 ^ 0)
                if ((-1 <= 0) && false) {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        if ((1e10 * 1e10) >= -2.25) {
                            for fi13 := 0; fi13 < 15; fi13++ {
                                v14 := GenS3{F0: (^(fi12)), F1: 601}
                                fmt.Println(fi13)
                                v15 := []int{fi13, (fi12 << ((255) & 63)), 592, -412}
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            switch fi12 {
                            case -286:
                                v4 = int32(v3)
                            case 9223372036854775807:
                                fmt.Println(v4, v3)
                            case 0:
                                var v16 float32 = float32(float64(1.5))
                                _ = v16
                            }
                        }
                        v17 := GenS4{F0: ("hello" + ("foo bar" + "foo bar"))}
                        switch 128 {
                        case -385:
                            var v18 int16 = 0
                            _ = v18
                        case -1:
                            v19 := []int{32767, fi12, fi12}
                            _ = v19
                        case -924:
                            if false {
                                v1[uint32(v3)] = false
                                v20 := GenS3{F0: (fi12 & fi12), F1: v3}
                                _ = v20
                            }
                        default:
                            var v21 float64 = 3.14159
                            _ = v21
                        }
                        _ = fi12
                        _ = v17
                    }
                }
            }
        }
        v22 := []int{-2147483648, 32768, 9223372036854775807, 128}
        _ = v22
    }
    v4 -= v4
    switch 450 {
    case 8:
        if (3.14159 <= 1e10) {
            for fi23 := 0; fi23 < 20; fi23++ {
                v24 := GenS5{F0: (fi23 | fi23)}
                _ = fi23
                _ = v24
            }
            v25 := GenS1{F0: 1, F1: 1000}
            for fi26 := 0; fi26 < 10; fi26++ {
                v27 := []int{fi26, 412}
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            v2 = v2
            _ = len(v1)
        }
    case -170:
        for fi28 := 0; fi28 < 10; fi28++ {
            v29 := GenS4{F0: ""}
            v1[2147483648] = (!true)
            v2 = 34
            _ = fi28
            _ = v29
        }
    case 128:
        v30 := GenS2{F0: -110}
        _ = v30
    default:
        for fi31 := 0; fi31 < 1; fi31++ {
            var v32 uint32 = 88
            _ = fi31
            _ = v32
        }
    }
    var v33 string = "hello"
    v34 := GenS5{F0: -196}
    var v35 string = "foo bar"
    var v36 uint16 = (v3 >> ((1) & 63))
    v36 -= 811
    var v37 uint16 = 751
    v37 -= v36
    _ = len(v1)
    var v38 uint32 = 1
    for fi39 := 0; fi39 < 20; fi39++ {
        if false {
            v40 := GenS6{F0: uint8(v2)}
            _ = v40
        } else {
            v37 = 1
            v33 = (v35 + v33)
            v33 = ("hello" + v35)
        }
        _ = fi39
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        var v42 int64 = 589
        var v43 float64 = (0.1 + ((1e10 + 1.5) / 1.0))
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
