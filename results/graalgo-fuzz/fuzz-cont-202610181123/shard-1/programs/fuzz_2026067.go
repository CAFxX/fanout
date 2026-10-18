package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 string; F1 int32 }
func genFn0(p0 uint64, p1 uint) uint8 {
    p1 -= (p1 << (((^(p1))) & 63))
    v7 := GenS1{F0: uint16(p1), F1: 38}
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1(p0 bool, p1 int32) int32 {
    p0 = (1 <= 32767)
    if false {
        p1 = -1
    } else {
        v8 := GenS2{F0: 620, F1: -57}
        if p0 {
            if (255 < 255) {
                if p0 {
                    for fi9 := 0; fi9 < 20; fi9++ {
                        for fi10 := 0; fi10 < 16; fi10++ {
                            var v11 uint64 = 0
                            fmt.Println(v11)
                            v12 := []int{}
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        fmt.Println(p0)
                        _ = fi9
                    }
                    var v13 int8 = 1
                    _ = v13
                }
            }
        }
        if p0 {
            p0 = false
        } else {
            p0 = p0
        }
        _ = v8
    }
    switch 32767 {
    case -2147483648:
        p1 = p1
    case 32768:
        p0 = true
    default:
        var v14 uint16 = 42
        _ = v14
    }
    var v15 string = "hello"
    var v16 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return (p1 % ((2147483647) | 1))
}

func main() {
    var v1 string = ((("\n\t" + "") + "a") + "foo bar")
    var v2 int64 = -523
    var v3 uint8 = 37
    var v4 uint64 = 42
    v5 := make(map[string]uint32)
    v6 := make(map[string]int)
    var v17 int16 = 267
    v17 -= (-(v17))
    v17 = 42
    var v18 int = 128
    var v19 string = ("hello" + "\n\t")
    var v20 uint8 = v3
    if false {
        var v21 float32 = ((0.1 * (0.1 / 1.0)) + 1.5)
        _ = v21
    }
    fmt.Println(v17, v19, v4)
    switch v19 {
    case "":
        if ((!true) || ((false && (false && false)) || (65535 <= 255))) {
            if (false || (false || (((v20 >> ((0) & 63)) < (v20 - (v20 % ((v3) | 1)))) || true))) {
                fmt.Println(v20, v1)
                fmt.Println(v18)
                v22 := GenS3{F0: v18, F1: -959}
                _ = v22
            } else {
                if false {
                    v23 := GenS4{F0: "", F1: 1}
                    _ = v23
                } else {
                    v2 += (v2 >> ((42) & 63))
                    for fi24 := 0; fi24 < 18; fi24++ {
                        _ = len(v5)
                        switch (("" + ((v1 + ("hello" + (("hello" + v19) + v1))) + ("a" + ""))) + v19) {
                        case "hello":
                            for fi25 := 0; fi25 < 10; fi25++ {
                                var v26 int = 102
                                v26 += v18
                                switch (v19 + (v19 + v19)) {
                                case "foo bar":
                                    v27 := GenS1{F0: 32767, F1: (v3 | v3)}
                                    _ = v27
                                default:
                                    fmt.Println(v18, v19)
                                }
                                _ = fi25
                                _ = v26
                            }
                        case "foo bar":
                            var v28 int8 = 127
                            _ = v28
                        case "a":
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v18, v20, v4)
                            }
                        default:
                            fmt.Println(v3, v20, v3)
                        }
                        _ = fi24
                    }
                    switch v1 {
                    case "a":
                        v29 := GenS4{F0: v19, F1: 0}
                        _ = v29
                    default:
                        for fi30 := 0; fi30 < 3; fi30++ {
                            v18 = 127
                            switch v19 {
                            case "":
                                v3 = 14
                            case "\n\t":
                                var v31 float32 = 1e10
                                _ = v31
                            }
                            _ = fi30
                        }
                    }
                }
                v32 := GenS3{F0: v18, F1: 2147483647}
                var v33 uint8 = (v3 - (v3 - 129))
                _ = v32
                _ = v33
            }
        }
    case "a":
        v18 += (v18 + (v18 + (v18 - 32768)))
    default:
        var v34 float32 = (0.1 * 1.5)
        _ = v34
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 float32 = float32(3.14159)
        _ = fi35
        _ = v36
    }
    var v37 uint16 = 2
    var v38 int32 = -2147483648
    var v39 int16 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
