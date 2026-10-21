package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }

func main() {
    var v1 uint64 = 111
    v2 := make(map[uint32]string)
    var v3 uint32 = 1
    var v4 int16 = 401
    v5 := []string{"foo bar"}
    switch int(v4) {
    case -1:
        switch 127 {
        case 600:
            var v6 float32 = -2.25
            _ = v6
        case 42:
            var v7 uint8 = 0
            _ = v7
        default:
            v4 -= v4
        }
    default:
        switch -1 {
        case -1:
            v1 -= (v1 % (((v1 - (^(v1)))) | 1))
        case -9223372036854775808:
            var v8 uint64 = (v1 - v1)
            _ = v8
        case 327:
            v4 = (^(v4))
        }
    }
    if (331 <= 255) {
        if (1e10 != (1e10 * (-2.25 * -2.25))) {
            v9 := GenS1{F0: 1}
            _ = v9
        }
        var v10 int8 = 1
        v11 := v2[(v3 - (v3 ^ (v3 * v3)))]
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: 128}
        v3 = (^(v3))
        if (!(((0.1 + -2.25) - float64(1.5)) == 3.14159)) {
            for fi13 := 0; fi13 < 18; fi13++ {
                v1 = v1
                var v14 uint16 = 992
                if (false && ("foo bar" < ("" + ""))) {
                    fmt.Println(v1)
                }
                _ = fi13
                _ = v14
            }
            if true {
                v4 -= v4
                var v15 int = -9223372036854775808
                _ = v15
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                v1 = v1
                for fi17 := 0; fi17 < 12; fi17++ {
                    switch fi17 {
                    case -377:
                        _ = len(v2)
                    case 32768:
                        v18 := GenS1{F0: -1}
                        _ = v18
                    }
                    var v19 uint32 = (v3 << ((4294967295) & 63))
                    _ = fi17
                    _ = v19
                }
                var v20 string = "a"
                _ = fi16
                _ = v20
            }
        }
        _ = v12
    }
    v21 := GenS1{F0: -2147483648}
    switch 128 {
    case 722:
        var v22 uint64 = (v1 / ((v1) | 1))
        _ = v22
    case -9223372036854775808:
        v4 = v4
    default:
        fmt.Println(v3)
    }
    var v23 float32 = 1.5
    var v24 uint64 = 42
    for fi25 := 0; fi25 < 19; fi25++ {
        for fi26 := 0; fi26 < 6; fi26++ {
            fmt.Println(v23, v3)
            _ = fi26
        }
        v3 = 173
        fmt.Println(v24, fi25, v1)
        _ = fi25
    }
    v4 = -399
    var v27 int64 = 1
    v23 = 1.5
    if true {
        v28 := GenS2{F0: 1, F1: -1}
        var v29 int = 9223372036854775807
        if (!(!true)) {
            if true {
                v30 := GenS1{F0: -781}
                for fi31 := 0; fi31 < 10; fi31++ {
                    var v32 float32 = (float32((((0.1 / float64(float32(0.1))) * 0.0) * (0.0 * -2.25))) * v23)
                    var v33 string = "hello"
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                v23 = v23
                _ = v30
            }
        } else {
            switch v29 {
            case 585:
                delete(v2, v3)
            case 32767:
                v34 := GenS3{F0: 65535}
                _ = v34
            default:
                v3 -= (v3 - (v3 * v3))
            }
            v24 += (^(v24))
        }
        _ = v28
        _ = v29
    } else {
        var v35 int64 = (v27 % (((v27 ^ 42)) | 1))
        v36 := GenS4{F0: v4, F1: 0, F2: 856}
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
