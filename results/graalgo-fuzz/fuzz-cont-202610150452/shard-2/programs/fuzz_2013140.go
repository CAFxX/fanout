package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 string; F1 int8; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 int64; F2 uint32 }

func main() {
    var v1 int = 4294967295
    var v2 uint64 = 247
    v3 := []string{""}
    var v4 int = -291
    for fi5 := 0; fi5 < 13; fi5++ {
        v2 -= (v2 ^ 413)
        _ = fi5
    }
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = (v2 & (^(v2)))
    switch v1 {
    case 42:
        fmt.Println(v4)
    case 727:
        for fi6 := 0; fi6 < 16; fi6++ {
            if true {
                for fi7 := 0; fi7 < 9; fi7++ {
                    var v8 int = (fi7 + (^(fi6)))
                    var v9 int32 = 42
                    var v10 uint64 = 1
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            }
            var v11 int = v4
            if true {
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                v12 := GenS1{F0: (false && false), F1: 481}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 6; fi13++ {
                    v14 := GenS2{F0: "a", F1: 40, F2: 42}
                    v1 -= (-(fi6))
                    fmt.Println(v11, v4, v11)
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi6
            _ = v11
        }
    default:
        if ((((((0.0 * (0.0 + ((1e10 * float32(0.0)) * 3.14159))) * (3.14159 / 1.5)) * (1.5 - 1e10)) - float32((0.1 + (0.1 * 0.0)))) - float32(-2.25)) < float32((0.1 + 3.14159))) {
            v2 = 1
            _ = len(v3)
        } else {
            if (2147483647 < 0) {
                v15 := GenS3{F0: 522}
                var v16 int16 = -1
                _ = v15
                _ = v16
            }
            if ((1 >= 1) && true) {
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 int = v4
                    v18 += 318
                    _ = fi17
                    _ = v18
                }
                v19 := GenS4{F0: 652, F1: true}
                fmt.Println(v1)
                _ = v19
            }
        }
    }
    var v20 int = (v1 * 127)
    fmt.Println(v20)
    var v21 int16 = 1
    if (-1 < 0) {
        fmt.Println(v20)
    } else {
        v22 := GenS5{F0: 2147483647, F1: 65535}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 += (v20 / ((int(v21)) | 1))
        _ = v22
    }
    v1 += v20
    for fi23 := 0; fi23 < 9; fi23++ {
        v24 := GenS6{F0: 11}
        _ = fi23
        _ = v24
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        v21 += (v21 << ((808) & 63))
        if len(v3) > 0 {
            v3[0] = ("hello" + ("hello" + "\n\t"))
        }
        if (((-2.25 / (3.14159 - float32(0.0))) * 1e10) >= ((((float32((-2.25 - 1e10)) * 0.0) * -2.25) / (-2.25 + (1e10 / 1.5))) + 1.5)) {
            v20 += 438
        }
        _ = fi25
    }
    v2 -= v2
    switch "" {
    case "hello":
        if true {
            v26 := GenS4{F0: v21, F1: true}
            v20 = (v4 - (v4 % (((v4 ^ (^(v4)))) | 1)))
            _ = v26
        } else {
            fmt.Println(v20, v21)
        }
    }
    switch v4 {
    case -347:
        if ((2147483647 != -644) && false) {
            _ = len(v3)
        } else {
            if false {
                v27 := GenS7{F0: ("hello" + ((("\n\t" + "foo bar") + "a") + "a")), F1: 42, F2: 1}
                var v28 uint = 906
                _ = v27
                _ = v28
            }
        }
    case 2147483647:
        fmt.Println(v2)
    }
    if len(v3) > 0 {
        v3[0] = ((((((("\n\t" + "foo bar") + "a") + "foo bar") + "hello") + ("" + ("foo bar" + ("a" + ("hello" + "hello"))))) + "") + "")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
