package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int32; F1 uint16; F2 uint32 }

func main() {
    var v1 int16 = 1000
    var v2 int64 = 42
    v3 := make(map[int]uint32)
    var v4 int64 = (v2 + int64(v1))
    if false {
        v5 := GenS1{F0: 65535, F1: (-788 < 0)}
        _ = v5
    }
    v6 := []int{1, -227}
    if (1 != 127) {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS2{F0: "a", F1: 348}
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := GenS3{F0: -128, F1: 2147483647}
                v1 = (v1 >> ((1) & 63))
                v11 := GenS4{F0: 654, F1: 114, F2: -110}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi7
            _ = v8
        }
        v2 += (v4 >> ((18446744073709551615) & 63))
    } else {
        var v12 uint64 = 2
        if (v1 != (v1 ^ (v1 - 0))) {
            if ((false || false) || true) {
                if false {
                    switch "a" {
                    case "":
                        v13 := GenS5{F0: (v1 - 0), F1: -2147483648}
                        _ = v13
                    case "hello":
                        var v14 int16 = v1
                        _ = v14
                    default:
                        v15 := GenS1{F0: 725, F1: false}
                        _ = v15
                    }
                    v16 := GenS2{F0: "hello", F1: 4294967295}
                    _ = v16
                } else {
                    fmt.Println(v4)
                    fmt.Println(v1, v4)
                    v17 := GenS6{F0: int32(v1), F1: 255, F2: 42}
                    _ = v17
                }
                v18 := GenS5{F0: (v1 ^ (v1 >> ((18446744073709551615) & 63))), F1: 2147483647}
                v4 = v2
                _ = v18
            } else {
                fmt.Println(v2, v1)
            }
            if false {
                _ = len(v6)
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 bool = (236 != 1)
                    var v21 float32 = (((float32(-2.25) - 1e10) / ((0.1 * float32(1e10)) - float32((-2.25 * -2.25)))) * 0.0)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            } else {
                switch 1 {
                case 0:
                    if false {
                        var v22 uint = 102
                        var v23 uint = uint(v12)
                        v22 += v23
                        _ = v22
                        _ = v23
                    }
                case 32767:
                    fmt.Println(v2)
                case 4294967295:
                    var v24 uint16 = 1
                    _ = v24
                }
                _ = len(v6)
            }
            for fi25 := 0; fi25 < 11; fi25++ {
                if true {
                    if (-2147483648 < 268) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                } else {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        var v27 float32 = float32(0.1)
                        switch "hello" {
                        case "hello":
                            v28 := GenS6{F0: 1, F1: 506, F2: 35}
                            _ = v28
                        case "":
                            _ = len(v6)
                        default:
                            v27 = v27
                        }
                        _ = fi26
                        _ = v27
                    }
                }
                _ = fi25
            }
        } else {
            _ = len(v6)
            v1 -= -32768
        }
        _ = v12
    }
    var v29 uint = 1
    v30 := GenS2{F0: ("" + "a"), F1: -2147483648}
    v31 := GenS4{F0: 42, F1: 42, F2: 1}
    v32 := v3[4294967295]
    if len(v6) > 0 {
        v6[0] = -618
    }
    switch (("\n\t" + ("" + ("" + "a"))) + "") {
    case "hello":
        for fi33 := 0; fi33 < 20; fi33++ {
            v6 = append(v6, fi33)
            var v34 float64 = -2.25
            _ = fi33
            _ = v34
        }
    default:
        for fi35 := 0; fi35 < 15; fi35++ {
            for fi36 := 0; fi36 < 10; fi36++ {
                v37 := v3[fi35]
                _ = fi36
                _ = v37
            }
            var v38 string = ("foo bar" + ("" + "hello"))
            switch ((v38 + "foo bar") + v38) {
            case "hello":
                v38 = ((v38 + "\n\t") + (v38 + (("hello" + "\n\t") + (v38 + v38))))
            case "":
                var v39 float32 = (3.14159 + (1e10 / 3.14159))
                _ = v39
            default:
                v1 -= v1
            }
            _ = fi35
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v6
}
