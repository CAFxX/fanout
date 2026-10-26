package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint8; F2 int8 }
type GenS6 struct { F0 int32; F1 uint; F2 int32 }
type GenS7 struct { F0 int8; F1 int16; F2 uint64 }
func genFn0(p0 int32) string {
    fmt.Println(p0)
    v7 := []int{32768, -2147483648}
    if len(v7) > 0 {
        v7[0] = -2147483648
    }
    if false {
        fmt.Println(p0)
    } else {
        v8 := GenS1{F0: p0}
        for fi9 := 0; fi9 < 17; fi9++ {
            switch "\n\t" {
            case "hello":
                fmt.Println(fi9, fi9)
            }
            _ = fi9
        }
        switch -47 {
        case -340:
            if len(v7) > 0 {
                v7[0] = 128
            }
        case 32767:
            var v10 float32 = float32((-2.25 / 3.14159))
            _ = v10
        case 0:
            p0 = (p0 * (p0 & 651))
        default:
            for fi11 := 0; fi11 < 6; fi11++ {
                if false {
                    switch ((("foo bar" + "") + "hello") + "\n\t") {
                    case "foo bar":
                        if ("" >= "") {
                            p0 -= p0
                            if false {
                                v12 := GenS2{F0: -425, F1: 255, F2: 127}
                                v13 := GenS3{F0: -1}
                                var v14 uint64 = 4294967295
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                        }
                    default:
                        for fi15 := 0; fi15 < 19; fi15++ {
                            p0 += p0
                            p0 = 0
                            switch fi11 {
                            case 9223372036854775807:
                                if len(v7) > 0 {
                                    v7[0] = fi11
                                }
                            case -1:
                                var v16 int16 = -730
                                _ = v16
                            case -2147483648:
                                var v17 int64 = 1
                                _ = v17
                            default:
                                p0 = (p0 | p0)
                            }
                            fmt.Println(fi15, fi15, fi11)
                            _ = fi15
                        }
                    }
                } else {
                    p0 = p0
                    var v18 int8 = 126
                    _ = v18
                }
                _ = fi11
            }
        }
        _ = v8
    }
    _ = p0
    _ = v7
    return ("foo bar" + ("hello" + "foo bar"))
}

func main() {
    var v1 int32 = -514
    var v2 uint8 = 42
    var v3 int64 = 9223372036854775807
    var v4 int16 = 0
    v5 := make(map[string]uint32)
    v6 := []int{32768}
    fmt.Println(v1, v4, v4)
    switch 0 {
    case -860:
        v19 := GenS4{F0: (v4 * v4)}
        _ = v19
    case 9223372036854775807:
        v20 := GenS5{F0: "foo bar", F1: 2, F2: int8(v3)}
        _ = v20
    case 42:
        switch 42 {
        case 2147483647:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            v3 += (v3 >> ((2) & 63))
        }
    }
    switch 2147483647 {
    case 4294967295:
        delete(v5, "foo bar")
    case 127:
        if (2147483647 <= 42) {
            v1 += (-(v1))
        } else {
            v4 = (-(v4))
            v3 = (v3 - (-(v3)))
        }
    default:
        v21 := GenS6{F0: (^(v1)), F1: 0, F2: v1}
        _ = v21
    }
    v3 = -517
    var v22 int32 = 0
    fmt.Println(v1, v4, v1)
    v22 += v1
    v23 := GenS1{F0: (v22 * -1)}
    fmt.Println(v1)
    v5[(("a" + "") + "hello")] = 2
    if false {
        v24 := GenS7{F0: 102, F1: 240, F2: 2}
        _ = v24
    } else {
        if false {
            var v25 uint64 = uint64(v2)
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 uint16 = 255
                _ = fi26
                _ = v27
            }
            v28 := v5["\n\t"]
            _ = v25
            _ = v28
        }
        v22 += (v22 & 2147483647)
        for fi29 := 0; fi29 < 7; fi29++ {
            delete(v5, "\n\t")
            switch "" {
            case "hello":
                if ((("foo bar" + "a") + "hello") == "a") {
                    v30 := GenS7{F0: 126, F1: (v4 << ((4294967295) & 63)), F2: 4294967295}
                    _ = v30
                }
            case "foo bar":
                v3 = v3
            }
            var v31 uint = 588
            _ = fi29
            _ = v31
        }
    }
    _ = len(v5)
    v4 += v4
    v2 = (^(v2))
    for fi32 := 0; fi32 < 19; fi32++ {
        fmt.Println(v3)
        fmt.Println(v3, v3)
        switch (("" + "hello") + ("" + "\n\t")) {
        case "foo bar":
            var v33 uint64 = 0
            _ = v33
        }
        _ = fi32
    }
    if true {
        v4 = (v4 % (((v4 & (v4 - v4))) | 1))
    }
    var v34 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
