package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := []uint32{4294967295, 2147483647, 49, 444}
    var v2 int8 = -71
    var v3 uint16 = 42
    for fi4 := 0; fi4 < 17; fi4++ {
        v2 += v2
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(v2, v2, fi5)
            v6 := GenS1{F0: 42, F1: v2, F2: fi4}
            var v7 float32 = (0.1 * (1.5 * 0.0))
            _ = fi5
            _ = v6
            _ = v7
        }
        fmt.Println(v3, v3)
        _ = fi4
    }
    v8 := GenS1{F0: 1, F1: 1, F2: 32767}
    var v9 float32 = 3.14159
    if len(v1) > 0 {
        v1[0] = uint32(v3)
    }
    v10 := GenS1{F0: 864, F1: v2, F2: 448}
    v11 := GenS2{F0: (0 > 1)}
    if true {
        for fi12 := 0; fi12 < 3; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                v3 = v3
                _ = fi13
            }
            switch "hello" {
            case "a":
                for fi14 := 0; fi14 < 7; fi14++ {
                    fmt.Println(v9, v9)
                    fmt.Println(v9, fi14)
                    v15 := GenS3{F0: 0}
                    var v16 uint32 = 0
                    fmt.Println(fi14, fi14, v9)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "foo bar":
                for fi17 := 0; fi17 < 10; fi17++ {
                    v1 = append(v1, 479)
                    v9 = v9
                    switch int(v3) {
                    case 0:
                        switch fi17 {
                        case 128:
                            v9 = v9
                        case 918:
                            if len(v1) > 0 {
                                v1[0] = 358
                            }
                        case 153:
                            var v18 float64 = (3.14159 + ((-2.25 + 0.0) * (3.14159 + 1e10)))
                            _ = v18
                        default:
                            var v19 float64 = (0.1 - -2.25)
                            _ = v19
                        }
                    case 1:
                        if (-419 > 636) {
                            v3 = (v3 - (v3 / ((65535) | 1)))
                            fmt.Println(v3, v2)
                        } else {
                            v9 = ((v9 + v9) + 1e10)
                        }
                    default:
                        switch (-(fi17)) {
                        case 137:
                            var v20 uint8 = 127
                            _ = v20
                        case 32768:
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                        case 9223372036854775807:
                            v9 -= v9
                        default:
                            v1 = append(v1, 2)
                        }
                    }
                    _ = fi17
                }
            case "\n\t":
                for fi21 := 0; fi21 < 7; fi21++ {
                    v9 += -2.25
                    v2 += 0
                    var v22 int64 = 42
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    v22 += (v22 % (((v22 | v22)) | 1))
                    v22 = 42
                    v23 := GenS4{F0: v22}
                    var v24 uint8 = 2
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            default:
                v25 := GenS2{F0: false}
                _ = v25
            }
            v2 = v2
            _ = fi12
        }
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 uint = 1
            _ = fi26
            _ = v27
        }
        if true {
            var v28 int16 = -32768
            var v29 uint32 = 1
            v9 = float32(float64((float32(0.1) - (v9 - (v9 * float32(1.5))))))
            _ = v28
            _ = v29
        } else {
            var v30 uint64 = 341
            _ = v30
        }
    } else {
        v3 += (v3 / (((v3 % ((v3) | 1))) | 1))
        v9 -= v9
    }
    v1 = append(v1, 0)
    v1 = append(v1, 164)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v2)
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 int32 = 42
        for fi33 := 0; fi33 < 6; fi33++ {
            v1 = append(v1, 2147483648)
            if false {
                if len(v1) > 0 {
                    v1[0] = 816
                }
            }
            v34 := GenS4{F0: -1}
            v35 := GenS2{F0: (93 >= 42)}
            var v36 uint32 = 2147483647
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
