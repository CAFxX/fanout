package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 uint64, p2 int16) uint32 {
    p1 += (^(p1))
    switch -1 {
    case 128:
        var v5 int64 = 42
        _ = v5
    case 32768:
        p2 = 32767
    case 9223372036854775807:
        v6 := []int{32767}
        _ = v6
    }
    v7 := []int{-9223372036854775808, 127, 0, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 32768
    var v3 float64 = ((1.5 + 1.5) * (float64(0.1) - (1e10 * 3.14159)))
    var v4 uint8 = 2
    v2 = (v2 >> ((42) & 63))
    var v8 int8 = -128
    fmt.Println(v8, v8)
    var v9 uint16 = 2
    var v10 float64 = (v3 / 1e10)
    var v11 uint64 = 868
    var v12 float32 = 1e10
    switch "foo bar" {
    case "a":
        switch (v2 + (v2 ^ v2)) {
        case 1:
            for fi13 := 0; fi13 < 12; fi13++ {
                v12 = v12
                _ = fi13
            }
        default:
            switch 128 {
            case 42:
                var v14 float64 = float64(v12)
                _ = v14
            case 32767:
                if true {
                    fmt.Println(v2, v3)
                }
            }
        }
    case "":
        if true {
            v15 := GenS1{F0: 2147483647, F1: 4294967295, F2: (v9 >> ((65535) & 63))}
            _ = v15
        } else {
            v16 := GenS2{F0: uint8(v11), F1: (((v9 * 620) > v9) || (111 >= 1)), F2: 2147483647}
            var v17 float32 = (float32((v3 + -2.25)) * float32(0.1))
            _ = v16
            _ = v17
        }
    default:
        var v18 int32 = 238
        _ = v18
    }
    v19 := GenS2{F0: uint8(v11), F1: ((true && (425 <= 18446744073709551615)) && (((-1 < 1) && (v2 != int(v9))) || (true || (1 > int16(v8))))), F2: 0}
    var v20 uint16 = 32767
    v21 := GenS3{F0: -2147483648}
    var v22 int8 = int8(v11)
    var v23 int = v2
    var v24 uint32 = 1
    if true {
        v25 := GenS4{F0: (v20 / (((v20 | v20)) | 1))}
        var v26 uint = 4294967295
        if true {
            if ((v4 - 127) >= 33) {
                var v27 uint8 = v4
                if (((v12 * (float32(((v3 + v10) - v10)) * 0.0)) != 3.14159) || (false || true)) {
                    var v28 float64 = v10
                    for fi29 := 0; fi29 < 15; fi29++ {
                        for fi30 := 0; fi30 < 13; fi30++ {
                            fmt.Println(v4)
                            _ = fi30
                        }
                        var v31 int8 = 10
                        if false {
                            v32 := GenS3{F0: 42}
                            if true {
                                fmt.Println(v28, v2)
                                for fi33 := 0; fi33 < 5; fi33++ {
                                    var v34 float32 = v12
                                    _ = fi33
                                    _ = v34
                                }
                                if (("foo bar" + "hello") == "foo bar") {
                                    fmt.Println(v10, v2)
                                    v12 = v12
                                } else {
                                    for fi35 := 0; fi35 < 1; fi35++ {
                                        v9 -= 0
                                        _ = fi35
                                    }
                                }
                            }
                            _ = v32
                        } else {
                            v2 -= (fi29 ^ (v2 << ((v26) & 63)))
                        }
                        _ = fi29
                        _ = v31
                    }
                    v1[337] = (v23 >> ((uint(v11)) & 63))
                    _ = v28
                }
                var v36 int8 = 1
                _ = v27
                _ = v36
            } else {
                for fi37 := 0; fi37 < 1; fi37++ {
                    var v38 int16 = -1
                    _ = fi37
                    _ = v38
                }
                v20 += v9
                var v39 string = "hello"
                _ = v39
            }
            var v40 float32 = (v12 - ((1.5 - v12) + v12))
            delete(v1, (v24 + v24))
            _ = v40
        }
        _ = v25
        _ = v26
    } else {
        switch 764 {
        case 2147483647:
            v41 := GenS5{F0: ("hello" + "\n\t")}
            _ = v41
        case 127:
            var v42 uint16 = uint16(v4)
            _ = v42
        case 32767:
            for fi43 := 0; fi43 < 3; fi43++ {
                v20 = (v20 & v9)
                v2 += -386
                v44 := GenS4{F0: 2}
                _ = fi43
                _ = v44
            }
        }
        switch 74 {
        case -9223372036854775808:
            fmt.Println(v12, v23)
        case 32768:
            v45 := GenS3{F0: 114}
            _ = v45
        case 128:
            switch v2 {
            case 9223372036854775807:
                v23 = (v23 / ((-1) | 1))
            }
        }
        fmt.Println(v20)
    }
    if false {
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
