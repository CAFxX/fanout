package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int }

func main() {
    var v1 float32 = 0.0
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]bool)
    var v4 int16 = 266
    if false {
        var v5 float32 = (v1 - (((-2.25 / 1.5) - 1.5) * v1))
        for fi6 := 0; fi6 < 8; fi6++ {
            v2 = (v2 >> ((18446744073709551615) & 63))
            _ = fi6
        }
        _ = v5
    }
    v7 := []int{-149, 0, -979}
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: (-(fi8)), F1: 14}
            var v11 int = -306
            _ = len(v7)
            _ = fi9
            _ = v10
            _ = v11
        }
        if (2147483648 == 757) {
            v3[(fi8 >> ((65535) & 63))] = (375 < 65535)
        }
        var v12 string = ("foo bar" + (("a" + "\n\t") + (("" + "\n\t") + "\n\t")))
        _ = fi8
        _ = v12
    }
    v13 := v3[-929]
    var v14 uint = 61
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 float64 = -2.25
        if true {
            var v17 uint8 = 128
            fmt.Println(v16)
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    var v18 float64 = (-2.25 / 1.0)
    switch ((("hello" + ("hello" + "hello")) + "") + ("a" + "foo bar")) {
    case "foo bar":
        if true {
            if (43 <= 127) {
                var v19 bool = (v13 || (int(v14) < 127))
                v19 = (!v19)
                _ = v19
            } else {
                v18 = (0.1 / ((v18 - -2.25) - v18))
                var v20 int32 = 0
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if false {
            for fi22 := 0; fi22 < 13; fi22++ {
                _ = len(v7)
                _ = fi22
            }
            var v23 uint64 = (v2 | (v2 ^ (^(v2))))
            _ = v23
        } else {
            if (!v13) {
                for fi24 := 0; fi24 < 20; fi24++ {
                    v7 = append(v7, (^(fi21)))
                    var v25 float32 = 0.0
                    if len(v7) > 0 {
                        v7[0] = (fi21 / ((fi24) | 1))
                    }
                    _ = fi24
                    _ = v25
                }
            }
            v26 := GenS2{F0: "\n\t", F1: 275, F2: 255}
            if v13 {
                fmt.Println(v18, v2)
                v4 += v4
                fmt.Println(v1)
            }
            _ = v26
        }
        v1 += float32(v18)
        var v27 bool = (65535 <= 1)
        _ = fi21
        _ = v27
    }
    var v28 int = 4294967295
    v14 = v14
    if (42 != 718) {
        v13 = (!false)
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 float32 = (-2.25 + float32(v18))
            _ = fi29
            _ = v30
        }
    } else {
        for fi31 := 0; fi31 < 15; fi31++ {
            v2 = (v2 / ((4294967295) | 1))
            switch 4294967295 {
            case 42:
                v32 := GenS3{F0: int(v2)}
                _ = v32
            case 0:
                for fi33 := 0; fi33 < 8; fi33++ {
                    var v34 int8 = 112
                    v35 := GenS2{F0: "foo bar", F1: 652, F2: 697}
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            default:
                v18 = -2.25
            }
            v1 = float32(float64(v1))
            _ = fi31
        }
        var v36 int64 = -1
        fmt.Println(v1)
        _ = v36
    }
    if len(v7) > 0 {
        v7[0] = (v28 ^ v28)
    }
    v37 := GenS2{F0: "hello", F1: 0, F2: 1}
    for fi38 := 0; fi38 < 19; fi38++ {
        v39 := GenS3{F0: (fi38 ^ -262)}
        _ = fi38
        _ = v39
    }
    v4 += -661
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v7
}
