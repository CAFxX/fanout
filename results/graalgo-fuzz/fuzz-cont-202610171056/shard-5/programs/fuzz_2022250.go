package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 int8; F2 bool }
func genFn0(p0 string, p1 int16, p2 int16) int8 {
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            p1 -= p2
            v9 := GenS1{F0: 306}
            _ = fi8
            _ = v9
        }
        v10 := []int{(fi7 % ((fi7) | 1))}
        _ = fi7
        _ = v10
    }
    fmt.Println(p2, p1)
    p2 = -32768
    if true {
        var v11 uint8 = 42
        fmt.Println(p0, p0, p1)
        v12 := GenS2{F0: (p1 | p1), F1: -2147483648}
        _ = v11
        _ = v12
    } else {
        p2 = p2
        p2 = p2
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 float64 = ((float64((-2.25 * 1e10)) / 1.0) + (1e10 - -2.25))
        p2 -= (p1 * -1)
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 107
}
func genFn1(p0 int16, p1 int64) int32 {
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := []int{(fi15 ^ (^(fi15))), (fi15 & 4294967295), fi15}
        fmt.Println(fi15)
        if false {
            p0 -= int16(p1)
        }
        _ = fi15
        _ = v16
    }
    p1 -= p1
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(p0)
        var v18 int64 = (p1 * (p1 % ((9223372036854775807) | 1)))
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    return -863
}

func main() {
    var v1 string = (("foo bar" + (("hello" + "hello") + "hello")) + (("foo bar" + "a") + "\n\t"))
    v2 := make(map[int]string)
    var v3 int64 = -9223372036854775808
    var v4 int16 = -1
    var v5 uint8 = 210
    var v6 uint8 = (v5 & 1)
    switch v1 {
    case "foo bar":
        var v19 int16 = (^(v4))
        _ = v19
    case "hello":
        if (255 == 42) {
            var v20 float64 = (1.5 + -2.25)
            switch v1 {
            case "":
                v21 := GenS1{F0: 32767}
                _ = v21
            default:
                v22 := GenS3{F0: "a"}
                _ = v22
            }
            fmt.Println(v1, v20, v6)
            _ = v20
        } else {
            v1 = v1
            var v23 int16 = (v4 - (v4 / ((v4) | 1)))
            _ = v23
        }
    case "a":
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 uint16 = 0
            _ = fi24
            _ = v25
        }
    default:
        fmt.Println(v1)
    }
    var v26 int16 = v4
    v6 = v6
    v27 := GenS4{F0: (v1 + (v1 + ""))}
    v28 := []int{130, -9223372036854775808}
    fmt.Println(v26)
    var v29 uint8 = 0
    for fi30 := 0; fi30 < 16; fi30++ {
        if true {
            for fi31 := 0; fi31 < 7; fi31++ {
                _ = len(v2)
                _ = fi31
            }
            switch ((v1 + v1) + (v1 + v1)) {
            case "a":
                var v32 uint16 = 2
                _ = v32
            case "":
                v4 = (v26 / (((-(v4))) | 1))
            default:
                var v33 uint = 375
                _ = v33
            }
            fmt.Println(v5, v26, fi30)
        } else {
            v5 -= (v6 ^ (^(v29)))
            var v34 int = fi30
            for fi35 := 0; fi35 < 19; fi35++ {
                v36 := GenS4{F0: ("hello" + (((v1 + "foo bar") + v1) + (("" + ("hello" + "a")) + "a")))}
                _ = fi35
                _ = v36
            }
            _ = v34
        }
        v29 += 127
        v4 += 763
        _ = fi30
    }
    v1 = "a"
    var v37 int64 = v3
    v28 = append(v28, -324)
    v38 := GenS5{F0: 65535, F1: 1, F2: false}
    var v39 string = (v1 + "hello")
    var v40 float64 = 1.5
    var v41 uint = 255
    v3 += -777
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    fmt.Println(v39, v6, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
