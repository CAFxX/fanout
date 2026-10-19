package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 string }
func genFn0(p0 uint, p1 uint16, p2 uint32) uint16 {
    v4 := []int{-2147483648}
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return uint16(p2)
}

func main() {
    v1 := []uint32{62, 42}
    var v2 float64 = 3.14159
    var v3 int8 = 24
    v5 := GenS1{F0: 89}
    if true {
        v3 -= v3
        v3 -= -1
        v6 := GenS2{F0: 32768}
        _ = v6
    }
    var v7 int64 = 9223372036854775807
    v1 = append(v1, 623)
    if ((^(v3)) != -128) {
        if (!false) {
            fmt.Println(v3)
        }
        for fi8 := 0; fi8 < 12; fi8++ {
            switch (("hello" + "hello") + "\n\t") {
            case "foo bar":
                var v9 float64 = v2
                _ = v9
            case "a":
                if false {
                    switch ((("hello" + "hello") + ("a" + "a")) + ("foo bar" + "foo bar")) {
                    case "":
                        if (853 > 4294967295) {
                            v10 := GenS3{F0: 1}
                            var v11 int16 = 1000
                            _ = v10
                            _ = v11
                        }
                    case "\n\t":
                        var v12 float32 = 1.5
                        _ = v12
                    default:
                        v13 := GenS4{F0: 2}
                        _ = v13
                    }
                    fmt.Println(fi8)
                } else {
                    fmt.Println(fi8, v3, fi8)
                }
            case "":
                v1 = append(v1, 65535)
            default:
                switch ((("foo bar" + ("" + "hello")) + "\n\t") + (("a" + "hello") + "hello")) {
                case "foo bar":
                    fmt.Println(v2)
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    switch 9223372036854775807 {
                    case 42:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case -128:
                        if len(v1) > 0 {
                            v1[0] = 255
                        }
                    case 0:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (65 != 65535) {
                var v14 uint64 = 4294967295
                var v15 int = (fi8 << ((207) & 63))
                _ = v14
                _ = v15
            }
            _ = fi8
        }
        _ = len(v1)
    } else {
        var v16 float64 = ((v2 - 3.14159) + (float64(3.14159) / 0.1))
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v7 += (v7 & v7)
        fmt.Println(v3, v3)
        _ = fi17
    }
    v7 = v7
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v18 := GenS5{F0: v3, F1: ("\n\t" + (("" + (("hello" + (("hello" + "a") + ("a" + "\n\t"))) + "foo bar")) + "a"))}
        switch 127 {
        case -9223372036854775808:
            _ = len(v1)
        case 783:
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        case 2147483647:
            var v19 uint8 = 255
            _ = v19
        default:
            switch 42 {
            case 1:
                if (!true) {
                    v20 := GenS5{F0: v3, F1: ""}
                    v2 += v2
                    _ = v20
                }
            case -183:
                if false {
                    v2 -= ((v2 - (0.0 / (float64(0.1) * v2))) + 0.0)
                    var v21 string = "\n\t"
                    var v22 int32 = 0
                    _ = v21
                    _ = v22
                }
            default:
                _ = len(v1)
            }
        }
        _ = v18
    }
    v3 -= 1
    v2 += v2
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v7
}
