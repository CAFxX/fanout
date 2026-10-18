package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint64; F2 int64 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint16; F1 int8; F2 int32 }
func genFn0() uint64 {
    if (true && ((-2147483648 <= 42) && (!false))) {
        var v4 uint8 = 255
        v4 = v4
        var v5 uint32 = 971
        _ = v4
        _ = v5
    }
    var v6 bool = (18446744073709551615 <= 2)
    _ = v6
    return 2
}
func genFn1(p0 int64) uint {
    p0 = (p0 >> ((741) & 63))
    v7 := []int{-1}
    for fi8 := 0; fi8 < 13; fi8++ {
        p0 += (p0 / ((p0) | 1))
        var v9 float64 = 0.1
        v7 = append(v7, -1)
        _ = fi8
        _ = v9
    }
    switch -2147483648 {
    case -802:
        fmt.Println(p0)
    case 4294967295:
        v10 := GenS1{F0: 42}
        _ = v10
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint16 = 272
    v2 := []int{-2147483648}
    var v3 int64 = 0
    fmt.Println(v3, v3)
    var v11 int64 = (^(v3))
    v12 := GenS2{F0: 0, F1: 420, F2: (v11 % ((-1) | 1))}
    var v13 int64 = 0
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if true {
        if true {
            var v14 int16 = 1
            fmt.Println(v1, v3)
            _ = v14
        } else {
            for fi15 := 0; fi15 < 15; fi15++ {
                fmt.Println(v13)
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint = 794
                    _ = fi16
                    _ = v17
                }
                v18 := GenS3{F0: 127, F1: -32768}
                _ = fi15
                _ = v18
            }
            v13 -= v11
        }
    } else {
        fmt.Println(v1)
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS4{F0: (^(v1)), F1: -128, F2: -2147483648}
                var v22 float32 = ((3.14159 - ((0.0 * (1.5 / float32((1.5 - float64((1e10 * (1.5 * -2.25))))))) + float32(float64(((1e10 - 0.0) * float32((-2.25 * (float64(0.0) / (1e10 * 0.1))))))))) * float32(0.1))
                _ = fi20
                _ = v21
                _ = v22
            }
            v13 -= v3
            _ = fi19
        }
        var v23 int = 659
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        for fi25 := 0; fi25 < 13; fi25++ {
            switch "\n\t" {
            case "a":
                if (((42 > 42) && (!true)) || true) {
                    v1 = (v1 | v1)
                } else {
                    switch ("foo bar" + "") {
                    case "hello":
                        var v26 int64 = v13
                        _ = v26
                    case "foo bar":
                        var v27 uint8 = 0
                        _ = v27
                    case "":
                        if true {
                            v1 += 65535
                            _ = len(v2)
                            fmt.Println(fi25)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    default:
                        switch -308 {
                        case 4294967295:
                            v28 := GenS2{F0: 590, F1: 631, F2: v13}
                            _ = v28
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    fmt.Println(v3, v1)
                    v29 := GenS3{F0: -1, F1: -32768}
                    _ = v29
                }
            case "foo bar":
                fmt.Println(fi24, v13)
            default:
                if (18446744073709551615 == 42) {
                    fmt.Println(fi24, v11, v13)
                    var v30 int = fi25
                    _ = v30
                }
            }
            _ = fi25
        }
        _ = fi24
    }
    var v31 int = -1
    if (true && true) {
        if len(v2) > 0 {
            v2[0] = (v31 ^ 0)
        }
        for fi32 := 0; fi32 < 5; fi32++ {
            v31 = 42
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 int8 = -128
                fmt.Println(fi33, fi33, fi33)
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
    } else {
        var v35 uint = 255
        v3 = v3
        _ = v35
    }
    var v36 uint = 4294967295
    var v37 uint8 = 52
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
}
