package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint; F2 int16 }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 uint16; F1 int; F2 string }

func main() {
    v1 := []uint32{748, 382}
    var v2 int8 = 42
    v3 := make(map[string]uint32)
    var v4 int32 = 1
    v5 := GenS1{F0: -1}
    v6 := GenS2{F0: (v2 ^ (v2 << ((65535) & 63)))}
    var v7 int64 = -9223372036854775808
    for fi8 := 0; fi8 < 6; fi8++ {
        if (4294967295 > 4294967295) {
            var v9 uint32 = 1
            if (3.14159 < 0.1) {
                var v10 bool = (v2 <= 42)
                for fi11 := 0; fi11 < 2; fi11++ {
                    switch fi8 {
                    case -202:
                        _ = len(v1)
                    case 968:
                        var v12 int64 = v7
                        _ = v12
                    default:
                        var v13 int8 = (v2 ^ (v2 ^ (v2 + (-(v2)))))
                        _ = v13
                    }
                    _ = fi11
                }
                v14 := GenS3{F0: fi8}
                _ = v10
                _ = v14
            } else {
                fmt.Println(v2, v9, v9)
            }
            v9 = v9
            _ = v9
        } else {
            for fi15 := 0; fi15 < 3; fi15++ {
                v4 = (v4 / ((int32(v2)) | 1))
                var v16 int = int(v7)
                var v17 uint64 = 1
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 string = ((("" + "hello") + "\n\t") + "a")
            if (127 != 2) {
                for fi19 := 0; fi19 < 20; fi19++ {
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                    _ = fi19
                }
                v18 = v18
            } else {
                switch "foo bar" {
                case "foo bar":
                    var v20 uint16 = 42
                    _ = v20
                default:
                    var v21 float32 = 0.1
                    _ = v21
                }
                v18 = "hello"
            }
            _ = v18
        }
        var v22 float64 = 0.1
        for fi23 := 0; fi23 < 1; fi23++ {
            fmt.Println(v2, v22, v2)
            v4 += (v4 ^ 1)
            v4 -= (v4 * (v4 >> ((65535) & 63)))
            _ = fi23
        }
        _ = fi8
        _ = v22
    }
    var v24 int64 = (v7 & (v7 & (v7 | (v7 << ((867) & 63)))))
    var v25 float64 = 1e10
    if (2 < 0) {
        v25 = float64(1.5)
        v24 += v24
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 uint8 = 24
            var v28 uint16 = 123
            v24 = -876
            _ = fi26
            _ = v27
            _ = v28
        }
    } else {
        if (!(v25 < v25)) {
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 int64 = v24
                v24 -= (v7 + v30)
                v31 := GenS4{F0: v2, F1: 195}
                _ = fi29
                _ = v30
                _ = v31
            }
            if (1e10 < 1e10) {
                var v32 uint = 594
                _ = v32
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v33 int16 = 42
        _ = len(v1)
        _ = v33
    }
    if true {
        v34 := GenS5{F0: 668, F1: 2, F2: 1000}
        var v35 uint = 4294967295
        _ = len(v1)
        _ = v34
        _ = v35
    }
    switch -705 {
    case -1:
        v36 := GenS6{F0: 4294967295, F1: 0}
        _ = v36
    case 4294967295:
        var v37 string = "hello"
        _ = v37
    }
    if false {
        v38 := GenS3{F0: 301}
        if false {
            v25 += (-2.25 * (v25 - 0.0))
        } else {
            if (-128 == (-(v2))) {
                var v39 float32 = (-2.25 - 0.0)
                _ = v39
            }
        }
        v40 := GenS7{F0: 42, F1: 32768, F2: (("hello" + ("a" + "a")) + "hello")}
        _ = v38
        _ = v40
    } else {
        fmt.Println(v4, v7, v24)
    }
    if false {
        if (18446744073709551615 != 25) {
            v41 := GenS7{F0: 65535, F1: 127, F2: (("\n\t" + "") + ("\n\t" + "hello"))}
            _ = v41
        } else {
            fmt.Println(v7)
            v24 -= v24
        }
    }
    if false {
        v3["a"] = 224
    } else {
        var v42 int64 = v7
        fmt.Println(v24)
        v1 = append(v1, 553)
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
