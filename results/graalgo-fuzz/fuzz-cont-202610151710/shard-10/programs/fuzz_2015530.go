package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint8; F1 int16; F2 uint8 }
func genFn0(p0 uint32, p1 uint64, p2 uint) string {
    v5 := GenS1{F0: 1, F1: 909, F2: 4294967295}
    v6 := GenS1{F0: 32767, F1: 744, F2: (p1 * (p1 ^ 392))}
    var v7 uint64 = (p1 & (p1 - 419))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return ("\n\t" + ("\n\t" + "hello"))
}

func main() {
    var v1 int32 = 1
    var v2 bool = (!(!false))
    var v3 int16 = 32767
    v4 := make(map[string]uint32)
    v8 := GenS2{F0: -1}
    if (!false) {
        var v9 string = ("foo bar" + ((("hello" + "a") + ("hello" + "hello")) + ("" + "foo bar")))
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint16 = 60
        if v2 {
            var v12 float32 = (3.14159 * -2.25)
            v13 := GenS1{F0: v11, F1: 9223372036854775807, F2: 886}
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 int8 = 100
    v15 := []int{-1, int(v1)}
    var v16 int = int(v14)
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS3{F0: 42, F1: 1}
        _ = len(v15)
        _ = fi17
        _ = v18
    }
    var v19 float32 = 0.1
    v4["a"] = 269
    v19 = float32(((0.1 * 0.1) - ((3.14159 + 1.5) * 0.0)))
    var v20 int16 = (^(v3))
    if (-188 < (v3 * 1)) {
        switch v16 {
        case -398:
            v3 = (v20 ^ 1000)
        case 897:
            for fi21 := 0; fi21 < 4; fi21++ {
                if (v3 >= int16(fi21)) {
                    var v22 uint = 600
                    var v23 int64 = 1
                    v19 = v19
                    _ = v22
                    _ = v23
                }
                v24 := GenS2{F0: 2147483647}
                if v2 {
                    switch v16 {
                    case 443:
                        var v25 int = (fi21 * v16)
                        _ = v25
                    case 0:
                        _ = len(v4)
                    case 128:
                        _ = len(v15)
                    }
                } else {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        v16 += fi26
                        var v27 float64 = float64(((v19 - 0.0) - 0.0))
                        if false {
                            if (1 >= 42) {
                                v19 -= ((float32(0.1) / v19) / 0.1)
                                fmt.Println(v3, v27)
                            }
                            if v2 {
                                v28 := GenS1{F0: 154, F1: 9223372036854775807, F2: 219}
                                _ = v28
                            } else {
                                var v29 int = int(v14)
                                var v30 bool = false
                                _ = v29
                                _ = v30
                            }
                            fmt.Println(v14)
                        } else {
                            v1 = v1
                            v2 = v2
                        }
                        _ = fi26
                        _ = v27
                    }
                    var v31 float32 = v19
                    fmt.Println(v31, v16)
                    _ = v31
                }
                _ = fi21
                _ = v24
            }
        default:
            for fi32 := 0; fi32 < 8; fi32++ {
                for fi33 := 0; fi33 < 15; fi33++ {
                    fmt.Println(v20, fi32)
                    fmt.Println(v3, v20, v3)
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = fi33
                }
                _ = fi32
            }
        }
        var v34 uint8 = 127
        for fi35 := 0; fi35 < 20; fi35++ {
            v36 := GenS3{F0: 1, F1: 2}
            for fi37 := 0; fi37 < 16; fi37++ {
                v2 = false
                if (true && (0 == v14)) {
                    switch "hello" {
                    case "":
                        var v38 uint32 = 2147483647
                        _ = v38
                    default:
                        switch v16 {
                        case -719:
                            delete(v4, "")
                        case 127:
                            if len(v15) > 0 {
                                v15[0] = -1
                            }
                        default:
                            if v2 {
                                if len(v15) > 0 {
                                    v15[0] = (-(fi37))
                                }
                                fmt.Println(v16, fi35, v1)
                                switch "a" {
                                case "a":
                                    v39 := GenS1{F0: 1, F1: 1, F2: 18446744073709551615}
                                    _ = v39
                                default:
                                    v4[("hello" + "hello")] = 469
                                }
                            } else {
                                v14 = v14
                            }
                        }
                    }
                    v40 := GenS1{F0: 0, F1: 0, F2: 2}
                    v41 := GenS4{F0: 255, F1: 42}
                    _ = v40
                    _ = v41
                }
                _ = fi37
            }
            v14 -= 1
            _ = fi35
            _ = v36
        }
        _ = v34
    } else {
        if true {
            switch v16 {
            case 127:
                fmt.Println(v16)
            }
            for fi42 := 0; fi42 < 17; fi42++ {
                var v43 bool = true
                for fi44 := 0; fi44 < 18; fi44++ {
                    var v45 uint8 = 0
                    v45 += (v45 >> ((0) & 63))
                    _ = len(v4)
                    _ = fi44
                    _ = v45
                }
                _ = fi42
                _ = v43
            }
        }
        switch "foo bar" {
        case "":
            _ = len(v4)
        default:
            if v2 {
                if len(v15) > 0 {
                    v15[0] = v16
                }
                v46 := GenS5{F0: 42, F1: v3, F2: 128}
                var v47 string = "\n\t"
                _ = v46
                _ = v47
            }
        }
        fmt.Println(v2, v16, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
