package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint8; F2 int32 }
func genFn0(p0 int) string {
    v5 := []int{9223372036854775807, p0, 145, p0}
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            p0 = (fi6 / ((p0) | 1))
            fmt.Println(fi7, p0, fi6)
            switch ("" + "\n\t") {
            case "foo bar":
                _ = len(v5)
                p0 -= -9223372036854775808
            case "\n\t":
                if true {
                    v8 := GenS1{F0: 1, F1: 1, F2: ("" + ((("foo bar" + "hello") + "a") + ("a" + "")))}
                    p0 = (fi7 ^ fi7)
                    _ = v8
                }
            case "hello":
                if true {
                    v9 := GenS2{F0: 42}
                    switch "\n\t" {
                    case "":
                        v10 := GenS2{F0: 0}
                        _ = v10
                    default:
                        fmt.Println(fi6, fi6, fi7)
                    }
                    _ = v9
                } else {
                    p0 = -9223372036854775808
                    v11 := GenS1{F0: -744, F1: 184, F2: "foo bar"}
                    switch (fi7 >> ((18446744073709551615) & 63)) {
                    case 1:
                        switch fi6 {
                        case 32768:
                            var v12 int64 = -9223372036854775808
                            _ = v12
                        }
                    }
                    _ = v11
                }
            default:
                fmt.Println(fi7)
            }
            _ = fi7
        }
        p0 = fi6
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
    }
    p0 = p0
    _ = p0
    _ = v5
    return "hello"
}
func genFn1(p0 bool) int64 {
    fmt.Println(p0)
    switch "hello" {
    case "a":
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS3{F0: 584}
            if p0 {
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := GenS1{F0: -1, F1: 127, F2: ""}
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(fi13, p0, p0)
                    }
                    p0 = p0
                    _ = fi15
                    _ = v16
                }
                var v17 string = ""
                _ = v17
            } else {
                var v18 int64 = -1
                switch "\n\t" {
                case "\n\t":
                    v19 := []int{(fi13 + (fi13 - (fi13 & 454))), 128, fi13, 0}
                    _ = v19
                case "hello":
                    fmt.Println(p0)
                default:
                    p0 = p0
                }
                _ = v18
            }
            _ = fi13
            _ = v14
        }
    }
    switch "a" {
    case "hello":
        v20 := []int{32767, 42}
        _ = v20
    default:
        switch 127 {
        case 9223372036854775807:
            p0 = (0 > -1)
        case -153:
            p0 = (3.14159 >= 3.14159)
        }
    }
    _ = p0
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 32767
    v3 := make(map[string]int)
    v4 := []uint32{3, 2, 2}
    var v21 uint = 4294967295
    v2 = v2
    var v22 int = 361
    var v23 float64 = 3.14159
    if (!(!("\n\t" == ("foo bar" + "foo bar")))) {
        var v24 int8 = -12
        _ = v24
    }
    _ = len(v3)
    var v25 int8 = 127
    v2 = (v2 * (v2 | v2))
    v26 := GenS4{F0: v2, F1: 255, F2: 2147483647}
    fmt.Println(v22, v21, v22)
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 bool = true
        if v28 {
            for fi29 := 0; fi29 < 18; fi29++ {
                v28 = (65535 != 782)
                var v30 uint32 = 243
                _ = fi29
                _ = v30
            }
            for fi31 := 0; fi31 < 16; fi31++ {
                v2 -= 1
                fmt.Println(fi27, v22, fi31)
                v22 += (fi27 << ((4294967295) & 63))
                var v32 int8 = -85
                fmt.Println(v32)
                v32 += 100
                v21 = v21
                fmt.Println(fi31)
                var v33 float32 = (-2.25 * 0.0)
                var v34 int8 = v25
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
        }
        _ = fi27
        _ = v28
    }
    var v35 uint = 944
    v36 := GenS1{F0: 264, F1: 1, F2: "hello"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
