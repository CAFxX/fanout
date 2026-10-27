package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }
func genFn0(p0 uint, p1 uint32, p2 int8) string {
    p0 -= p0
    p1 += (p1 / ((2147483648) | 1))
    if true {
        p1 = p1
        p2 += 124
    }
    v7 := GenS1{F0: "a"}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}
func genFn1(p0 uint32, p1 uint) int8 {
    var v8 uint16 = 42
    for fi9 := 0; fi9 < 8; fi9++ {
        p1 = (p1 + 18446744073709551615)
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: "a"}
            v8 = uint16(p1)
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: ("foo bar" + "hello")}
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v8
    return 100
}

func main() {
    v1 := []string{}
    var v2 string = ("foo bar" + "a")
    v3 := make(map[string]bool)
    v4 := []uint32{42, 42}
    v5 := []string{"foo bar", "hello", (v2 + v2)}
    var v6 int8 = 100
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(fi14)
            fmt.Println(fi14)
            if (365 <= 42) {
                v15 := GenS1{F0: "hello"}
                if (255 == 255) {
                    v2 = (v2 + "foo bar")
                    v16 := GenS2{F0: 720, F1: 252}
                    var v17 uint64 = 1
                    _ = v16
                    _ = v17
                }
                switch (v2 + v2) {
                case "hello":
                    fmt.Println(fi13)
                    v18 := GenS2{F0: 32767, F1: 2}
                    var v19 float64 = -2.25
                    _ = v18
                    _ = v19
                case "foo bar":
                    var v20 int64 = 0
                    _ = v20
                case "\n\t":
                    fmt.Println(fi14)
                default:
                    fmt.Println(fi14, v6, v2)
                }
                _ = v15
            } else {
                v21 := GenS2{F0: 1, F1: 255}
                _ = len(v4)
                v2 = "hello"
                _ = v21
            }
            _ = fi14
        }
        v6 += genFn1(2, 687)
        _ = fi13
    }
    switch -2147483648 {
    case 9223372036854775807:
        var v22 int = 0
        _ = v22
    }
    var v23 int = -2147483648
    v6 = v6
    v24 := GenS1{F0: v2}
    v23 = v23
    _ = len(v4)
    var v25 uint = 255
    fmt.Println(v23, v6)
    switch "foo bar" {
    case "hello":
        if (!(-9223372036854775808 == 861)) {
            fmt.Println(v2, v23)
            for fi26 := 0; fi26 < 7; fi26++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v2 = v2
                _ = fi26
            }
        } else {
            v6 = genFn1(422, 2)
            for fi27 := 0; fi27 < 9; fi27++ {
                if true {
                    var v28 int32 = 0
                    _ = v28
                } else {
                    _ = len(v3)
                    _ = len(v3)
                    v25 += 1
                }
                _ = fi27
            }
            var v29 int8 = (-(v6))
            _ = v29
        }
    }
    if (1 <= 2) {
        _ = len(v5)
        if (v23 >= v23) {
            _ = len(v4)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if len(v5) > 0 {
                v5[0] = ("hello" + v2)
            }
            v30 := GenS2{F0: 1000, F1: v25}
            v31 := GenS2{F0: -583, F1: (^(v25))}
            _ = v30
            _ = v31
        }
    } else {
        var v32 int64 = -1
        _ = v32
    }
    v33 := GenS1{F0: "\n\t"}
    for fi34 := 0; fi34 < 20; fi34++ {
        fmt.Println(v23, v6)
        fmt.Println(v23)
        fmt.Println(v25, v6)
        _ = fi34
    }
    v35 := GenS3{F0: 2, F1: (v23 | 830), F2: 1}
    if false {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        v2 = "\n\t"
    } else {
        v23 = 32768
        v23 = (v23 / (((v23 * v23)) | 1))
        switch 127 {
        case 127:
            v36 := GenS3{F0: 323, F1: (v23 & v23), F2: 1}
            _ = v36
        default:
            if len(v4) > 0 {
                v4[0] = 0
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
