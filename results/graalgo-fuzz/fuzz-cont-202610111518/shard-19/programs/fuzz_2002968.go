package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 int64; F1 int8; F2 uint16 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]string)
    var v3 uint32 = 42
    var v4 uint32 = (v3 + 2147483648)
    var v5 float32 = -2.25
    if (("foo bar" + (("" + ("\n\t" + ("\n\t" + "foo bar"))) + "\n\t")) <= (("\n\t" + "\n\t") + "")) {
        v3 = v4
    } else {
        if true {
            for fi6 := 0; fi6 < 19; fi6++ {
                if true {
                    v4 = v3
                } else {
                    v3 -= v3
                    v5 += v5
                }
                _ = len(v2)
                _ = fi6
            }
            switch 42 {
            case 328:
                var v7 int8 = 0
                _ = v7
            default:
                for fi8 := 0; fi8 < 4; fi8++ {
                    if false {
                        fmt.Println(v5, fi8)
                    }
                    v4 = 583
                    switch 42 {
                    case 9223372036854775807:
                        for fi9 := 0; fi9 < 15; fi9++ {
                            v10 := []int{-877, (-(fi9)), -2147483648, int(v1)}
                            switch -9223372036854775808 {
                            case 0:
                                v3 = 251
                            case 335:
                                v4 = (v4 - v4)
                            }
                            fmt.Println(v1, v4)
                            if false {
                                if false {
                                    var v11 int = (fi9 / (((fi9 + 32767)) | 1))
                                    _ = v11
                                } else {
                                    v3 = (v4 % ((v3) | 1))
                                    v12 := GenS1{F0: ((!true) && false), F1: (fi8 % ((fi9) | 1)), F2: (v1 * (v1 / (((v1 << ((18446744073709551615) & 63))) | 1)))}
                                    _ = v12
                                }
                                var v13 uint32 = (^(v4))
                                _ = v13
                            }
                            _ = fi9
                            _ = v10
                        }
                    case 1:
                        v5 = (((v5 - v5) / v5) * (v5 - ((float32(float64(1.5)) * float32(1.5)) / float32((((1e10 - (3.14159 - 1e10)) / float64((3.14159 + 3.14159))) / 1.0)))))
                    case -1:
                        if (125 < (v3 ^ v3)) {
                            v3 = (^(v3))
                            v1 = (v1 << ((255) & 63))
                            var v14 uint64 = 2
                            _ = v14
                        }
                    }
                    _ = fi8
                }
            }
        } else {
            for fi15 := 0; fi15 < 2; fi15++ {
                fmt.Println(v3)
                var v16 int16 = 1000
                var v17 uint64 = v1
                _ = fi15
                _ = v16
                _ = v17
            }
            if (v4 == (v3 >> ((65535) & 63))) {
                v3 = (v4 >> ((2) & 63))
            }
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := GenS1{F0: true, F1: (fi18 + (^(fi18))), F2: v1}
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 float64 = float64((v5 * v5))
            var v22 uint8 = 2
            v1 = v1
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v3 = (v3 * (v4 - v3))
    for fi23 := 0; fi23 < 9; fi23++ {
        switch "" {
        case "a":
            v24 := GenS1{F0: false, F1: fi23, F2: v1}
            _ = v24
        default:
            if (false && (float64(v5) <= -2.25)) {
                var v25 string = "hello"
                var v26 string = "hello"
                v27 := GenS1{F0: (1e10 >= v5), F1: fi23, F2: v1}
                _ = v25
                _ = v26
                _ = v27
            }
        }
        v1 += (v1 << ((0) & 63))
        _ = fi23
    }
    switch 925 {
    case 32767:
        var v28 uint16 = 65535
        _ = v28
    default:
        var v29 int = 95
        _ = v29
    }
    v5 = (0.1 * (v5 / (v5 - float32(0.1))))
    for fi30 := 0; fi30 < 1; fi30++ {
        if (2147483647 <= 42) {
            fmt.Println(v1)
            v31 := GenS2{F0: 1, F1: -128, F2: 2}
            var v32 uint16 = 0
            _ = v31
            _ = v32
        }
        if (v5 == v5) {
            delete(v2, fi30)
            for fi33 := 0; fi33 < 14; fi33++ {
                v1 = v1
                fmt.Println(fi30)
                v4 -= v3
                _ = fi33
            }
        }
        _ = fi30
    }
    v1 += v1
    v34 := GenS2{F0: 0, F1: 1, F2: 65535}
    v35 := GenS2{F0: 42, F1: -39, F2: 1}
    v4 = 0
    var v36 uint = 255
    var v37 int16 = 42
    v38 := []int{42, int(v37), -547, 4294967295}
    var v39 int16 = (-(v37))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
