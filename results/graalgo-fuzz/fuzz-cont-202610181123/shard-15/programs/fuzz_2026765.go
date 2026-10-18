package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint32; F2 string }
type GenS6 struct { F0 uint16; F1 string; F2 int32 }
type GenS7 struct { F0 int32; F1 uint8 }
func genFn0() uint64 {
    var v5 uint8 = 255
    v6 := []int{42, 32768, -2147483648, 2147483647}
    var v7 int32 = 2147483647
    var v8 int32 = v7
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 int16, p1 uint32, p2 uint8) int8 {
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int16 = 0
        fmt.Println(p0, fi9)
        if true {
            v11 := GenS1{F0: 115, F1: 207, F2: -695}
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        switch -275 {
        case 32767:
            fmt.Println(p1, p2, p1)
        case 1:
            v12 := GenS1{F0: -1, F1: -1, F2: 1}
            _ = v12
        case 42:
            var v13 int = 32768
            _ = v13
        }
    default:
        fmt.Println(p2, p2, p0)
    }
    v14 := GenS1{F0: 113, F1: 0, F2: 0}
    if (((1e10 - 3.14159) + -2.25) != ((((((3.14159 / 1.0) / 1.0) * 0.0) / 1e10) - (1e10 + ((-2.25 + ((-2.25 / (3.14159 / float32(3.14159))) + 1.5)) + 1.5))) / float32(float64((1e10 / (1.5 + ((float32(0.1) - 1.5) - -2.25))))))) {
        p2 += (p2 + p2)
        switch -341 {
        case 4294967295:
            var v15 int8 = 0
            _ = v15
        default:
            p0 -= p0
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 127
}

func main() {
    var v1 uint16 = 42
    var v2 int32 = -1
    var v3 float64 = 0.0
    v4 := []uint32{206}
    for fi16 := 0; fi16 < 18; fi16++ {
        if (false || true) {
            v17 := GenS2{F0: 420, F1: 65535, F2: 42}
            _ = v17
        } else {
            _ = len(v4)
            v18 := GenS3{F0: (0 <= 42)}
            _ = v18
        }
        v2 += -561
        var v19 uint16 = 65535
        _ = fi16
        _ = v19
    }
    fmt.Println(v3, v2, v2)
    v20 := GenS3{F0: (v2 > 63)}
    var v21 int64 = -1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        v2 = (-(v2))
        v1 += v1
        fmt.Println(v3, v1)
        _ = fi22
    }
    var v23 float64 = v3
    _ = len(v4)
    var v24 string = "hello"
    if false {
        for fi25 := 0; fi25 < 12; fi25++ {
            switch fi25 {
            case 4294967295:
                v21 = 9223372036854775807
            default:
                var v26 uint8 = 255
                _ = v26
            }
            switch "a" {
            case "hello":
                var v27 float32 = 1e10
                _ = v27
            default:
                fmt.Println(v2, fi25, v24)
            }
            var v28 float64 = v23
            _ = fi25
            _ = v28
        }
        v29 := GenS1{F0: 100, F1: (-(v2)), F2: v21}
        _ = v29
    } else {
        v30 := GenS4{F0: -723}
        var v31 uint16 = v1
        if (!false) {
            v32 := GenS5{F0: 42, F1: 2, F2: v24}
            var v33 string = ((v24 + v24) + v24)
            switch 32768 {
            case -172:
                for fi34 := 0; fi34 < 3; fi34++ {
                    var v35 float64 = (v3 + (0.0 / v3))
                    _ = fi34
                    _ = v35
                }
            case 128:
                if (false && (4294967295 >= 360)) {
                    v36 := GenS6{F0: v1, F1: (v24 + ((("hello" + (((v33 + "") + "foo bar") + "")) + ((v24 + "\n\t") + v24)) + "\n\t")), F2: 2147483647}
                    _ = v36
                } else {
                    v1 = 32767
                }
            case 32768:
                if len(v4) > 0 {
                    v4[0] = 2
                }
            default:
                var v37 bool = (v2 > v2)
                _ = v37
            }
            _ = v32
            _ = v33
        } else {
            var v38 float64 = float64(0.1)
            _ = v38
        }
        _ = v30
        _ = v31
    }
    v39 := GenS7{F0: (v2 ^ int32(v1)), F1: 255}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        v1 -= (^(v1))
        for fi41 := 0; fi41 < 10; fi41++ {
            v3 = v3
            _ = fi41
        }
        if true {
            for fi42 := 0; fi42 < 9; fi42++ {
                v21 -= v21
                var v43 int16 = 42
                _ = fi42
                _ = v43
            }
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
}
