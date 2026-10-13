package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 2147483647
    v3 := []uint32{2147483647}
    var v4 int32 = v2
    v5 := GenS1{F0: -32768}
    if false {
        v6 := GenS2{F0: 966}
        var v7 uint64 = 4294967295
        var v8 float64 = (-2.25 + 0.1)
        _ = v6
        _ = v7
        _ = v8
    } else {
        v4 -= (v4 << ((1) & 63))
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v2 = (v2 ^ -2147483648)
        _ = fi9
    }
    v2 = v4
    v2 = v4
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS3{F0: 141}
        _ = fi10
        _ = v11
    }
    switch -492 {
    case 42:
        if (v4 == (v4 % (((v2 / ((1) | 1))) | 1))) {
            var v12 float32 = (0.0 / ((1e10 / 0.1) - 1e10))
            _ = v12
        }
    case 0:
        var v13 uint32 = 2147483647
        _ = v13
    }
    switch 252 {
    case 128:
        var v14 int64 = -1
        _ = v14
    case 32768:
        for fi15 := 0; fi15 < 14; fi15++ {
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := GenS2{F0: fi15}
                _ = len(v3)
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    case 604:
        v2 += (v2 >> ((871) & 63))
    default:
        fmt.Println(v2, v4)
    }
    switch -2147483648 {
    case 83:
        fmt.Println(v2)
    default:
        var v18 uint = 255
        _ = v18
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    for fi19 := 0; fi19 < 17; fi19++ {
        if true {
            switch (("a" + "hello") + "foo bar") {
            case "foo bar":
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 int64 = 205
                    switch "hello" {
                    case "":
                        v3 = append(v3, 21)
                    case "hello":
                        var v22 int16 = 32767
                        _ = v22
                    default:
                        v2 += v4
                    }
                    v1[65535] = ""
                    var v23 int8 = 100
                    _ = fi20
                    _ = v21
                    _ = v23
                }
            case "a":
                if false {
                    if ((1e10 * (((float64(1.5) / 0.1) - (1.5 * 3.14159)) + float64((((float32(float64(1.5)) * 1e10) - ((float32(0.0) + (-2.25 * (1e10 * (0.1 / 3.14159)))) - 3.14159)) + 1.5)))) >= float64((1e10 + float32(-2.25)))) {
                        v2 -= (-(v4))
                        v2 = v2
                    }
                    fmt.Println(v2)
                    if true {
                        v24 := v1[4294967295]
                        _ = v24
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                    }
                } else {
                    switch fi19 {
                    case 1:
                        v4 = (v2 | 465)
                    case -1:
                        var v25 float64 = 1.5
                        _ = v25
                    }
                }
            case "\n\t":
                fmt.Println(v4)
            }
            v26 := GenS4{F0: 2147483648, F1: 1}
            _ = v26
        }
        v1[765] = ("\n\t" + ("" + (("foo bar" + "a") + "hello")))
        _ = fi19
    }
    switch 4294967295 {
    case 4294967295:
        v4 = v2
    }
    var v27 int8 = -1
    var v28 float64 = (float64(((0.0 + (((3.14159 / 0.1) * (((0.1 / 1.0) / 1e10) - -2.25)) + 0.1)) + 1.5)) * float64((3.14159 * 3.14159)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
