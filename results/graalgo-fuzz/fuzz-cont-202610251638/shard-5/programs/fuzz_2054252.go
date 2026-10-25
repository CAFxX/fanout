package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 uint }
type GenS8 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    v3 := make(map[uint32]bool)
    var v4 uint64 = 555
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 string = "hello"
        if false {
            var v7 float64 = 0.1
            _ = v7
        }
        v8 := GenS1{F0: 42, F1: 888, F2: 128}
        _ = fi5
        _ = v6
        _ = v8
    }
    var v9 int32 = 780
    for fi10 := 0; fi10 < 20; fi10++ {
        v4 += (v4 - 769)
        switch (fi10 / ((fi10) | 1)) {
        case 2147483647:
            switch (fi10 << ((255) & 63)) {
            case 9223372036854775807:
                var v11 int32 = (v9 >> ((2) & 63))
                _ = v11
            default:
                switch fi10 {
                case -2147483648:
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        v12 := GenS2{F0: v9, F1: true, F2: 42}
                        _ = v12
                    case "":
                        v9 -= (v9 ^ -2147483648)
                    default:
                        var v13 uint32 = 2147483648
                        _ = v13
                    }
                case 1:
                    var v14 int = (-(fi10))
                    _ = v14
                case -971:
                    v4 += (v4 + v4)
                }
            }
        default:
            switch -1 {
            case -1:
                if true {
                    v2 = append(v2, fi10)
                    switch (fi10 - 0) {
                    case 9223372036854775807:
                        for fi15 := 0; fi15 < 20; fi15++ {
                            if true {
                                fmt.Println(v4, fi10, v9)
                            }
                            _ = fi15
                        }
                    case -9223372036854775808:
                        v16 := GenS1{F0: 255, F1: int64(v9), F2: 2}
                        _ = v16
                    case 4294967295:
                        switch fi10 {
                        case 42:
                            for fi17 := 0; fi17 < 19; fi17++ {
                                v18 := GenS1{F0: 127, F1: 1, F2: 0}
                                v4 = 4294967295
                                var v19 float64 = -2.25
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        default:
                            var v20 uint = 42
                            _ = v20
                        }
                    default:
                        v21 := GenS3{F0: 393}
                        _ = v21
                    }
                } else {
                    v22 := GenS4{F0: -30}
                    v4 = v4
                    v23 := GenS5{F0: int64(v9)}
                    _ = v22
                    _ = v23
                }
            }
        }
        fmt.Println(fi10, v4)
        _ = fi10
    }
    var v24 uint = 832
    fmt.Println(v4, v24)
    _ = len(v3)
    switch -2147483648 {
    case 128:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        if true {
            v25 := v1[("foo bar" + (((("\n\t" + "\n\t") + ("" + "foo bar")) + "") + "hello"))]
            _ = v25
        } else {
            v9 += 1
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 float64 = 3.14159
                v4 -= (^(v4))
                _ = fi26
                _ = v27
            }
        }
    }
    var v28 uint8 = 0
    if (true || (-9223372036854775808 < -691)) {
        v4 += (v4 << ((v24) & 63))
        v9 = v9
        if true {
            delete(v1, "\n\t")
        }
    } else {
        if true {
            if len(v2) > 0 {
                v2[0] = 127
            }
        } else {
            v4 = v4
            for fi29 := 0; fi29 < 5; fi29++ {
                fmt.Println(v4)
                _ = fi29
            }
        }
        if (v9 >= (-(v9))) {
            fmt.Println(v9)
            v24 -= 422
            for fi30 := 0; fi30 < 8; fi30++ {
                switch 127 {
                case -569:
                    var v31 uint16 = 305
                    _ = v31
                case 316:
                    if (0 == (v24 << (((^(v24))) & 63))) {
                        fmt.Println(v28, v4)
                    }
                default:
                    switch (fi30 << ((3) & 63)) {
                    case 32768:
                        v9 = v9
                    case 507:
                        v32 := GenS6{F0: -32768}
                        _ = v32
                    case 0:
                        switch fi30 {
                        case 42:
                            v33 := GenS5{F0: 42}
                            _ = v33
                        default:
                            switch (fi30 >> (((v24 ^ 255)) & 63)) {
                            case 2147483647:
                                v24 = 238
                            }
                        }
                    }
                }
                v34 := GenS7{F0: 260, F1: (^(v24))}
                _ = fi30
                _ = v34
            }
        } else {
            fmt.Println(v9, v24, v4)
        }
    }
    v35 := GenS2{F0: (v9 + v9), F1: (!(1 >= 32767)), F2: 65535}
    v36 := GenS5{F0: 1}
    v37 := GenS2{F0: 42, F1: false, F2: 878}
    var v38 int64 = -161
    switch -474 {
    case 32767:
        switch "hello" {
        case "\n\t":
            v2 = append(v2, 2147483647)
        }
    }
    v4 = (v4 % (((v4 & 2)) | 1))
    v39 := GenS8{F0: 2147483647, F1: (v4 >> ((2) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v9
}
