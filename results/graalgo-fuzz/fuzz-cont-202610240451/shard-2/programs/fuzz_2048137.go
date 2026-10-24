package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int; F1 int; F2 uint8 }
func genFn0() uint32 {
    var v4 int = -95
    var v5 int64 = 0
    v6 := GenS1{F0: 164, F1: 128}
    switch "foo bar" {
    case "":
        v4 -= v4
    default:
        for fi7 := 0; fi7 < 3; fi7++ {
            if false {
                v5 = v5
                v4 = v4
            }
            var v8 uint32 = 663
            _ = fi7
            _ = v8
        }
    }
    var v9 int8 = 0
    _ = v4
    _ = v5
    _ = v6
    _ = v9
    return 2
}
func genFn1(p0 uint16, p1 int64) uint {
    var v10 int16 = 433
    v10 -= (v10 << ((537) & 63))
    fmt.Println(p1, v10)
    switch "" {
    case "":
        var v11 int32 = 0
        _ = v11
    case "a":
        p1 -= p1
    case "hello":
        v12 := GenS1{F0: 0, F1: 238}
        _ = v12
    default:
        fmt.Println(v10)
    }
    _ = p0
    _ = p1
    _ = v10
    return 92
}

func main() {
    var v1 uint = 2
    var v2 bool = (!true)
    v3 := make(map[uint32]int)
    if false {
        for fi13 := 0; fi13 < 14; fi13++ {
            fmt.Println(v1, fi13)
            if true {
                for fi14 := 0; fi14 < 14; fi14++ {
                    switch 2147483647 {
                    case -601:
                        v15 := GenS2{F0: 1}
                        _ = v15
                    }
                    v2 = v2
                    _ = len(v3)
                    fmt.Println(v1, v2, fi14)
                    _ = fi14
                }
            }
            fmt.Println(v2)
            _ = fi13
        }
        v1 -= v1
    }
    if ((v2 || v2) && false) {
        var v16 int8 = -128
        v1 += (v1 << ((453) & 63))
        _ = v16
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS1{F0: -422, F1: 200}
            v1 = v1
            _ = fi17
            _ = v18
        }
        switch -314 {
        case 19:
            fmt.Println(v1, v1)
        }
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(v1, v2)
            _ = fi20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        if (!false) {
            fmt.Println(fi21, fi21, v1)
            var v22 uint8 = 222
            _ = v22
        }
        v23 := GenS3{F0: fi21, F1: 127}
        _ = fi21
        _ = v23
    }
    var v24 uint64 = 0
    if (3.14159 <= 3.14159) {
        v2 = (v2 || true)
        v1 += (v1 - 255)
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 int64 = 0
        var v27 uint64 = (v24 << ((811) & 63))
        if v2 {
            switch fi25 {
            case -829:
                var v28 string = ""
                _ = v28
            case 1:
                var v29 uint32 = 2147483647
                _ = v29
            default:
                var v30 int64 = v26
                _ = v30
            }
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    if v2 {
        for fi31 := 0; fi31 < 9; fi31++ {
            if true {
                v32 := GenS2{F0: -813}
                switch 128 {
                case 127:
                    v33 := []int{-409, 820, 128, (fi31 ^ (fi31 ^ -426))}
                    _ = v33
                case -9223372036854775808:
                    v34 := v3[2147483648]
                    _ = v34
                default:
                    if (!(110 <= 255)) {
                        switch "hello" {
                        case "a":
                            v35 := GenS4{F0: fi31, F1: -676, F2: 128}
                            _ = v35
                        }
                        switch (fi31 << ((genFn1(65535, 42)) & 63)) {
                        case 2147483647:
                            var v36 float64 = -2.25
                            _ = v36
                        }
                    }
                }
                _ = v32
            } else {
                var v37 bool = v2
                var v38 int32 = -359
                _ = v37
                _ = v38
            }
            _ = fi31
        }
        v2 = (((1e10 == 1e10) && (631 < 994)) || (0.0 == 1e10))
        var v39 float32 = 0.0
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
