package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 bool }
func genFn0() string {
    var v7 uint32 = 1
    v7 = (^(v7))
    fmt.Println(v7)
    v8 := GenS1{F0: 95, F1: 74}
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 int16, p1 uint64, p2 uint16) uint16 {
    if true {
        var v9 uint16 = (p2 % ((p2) | 1))
        _ = v9
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 int32 = -802
    var v2 bool = false
    v3 := make(map[int]int)
    var v4 string = ""
    v5 := make(map[uint32]int)
    var v6 uint32 = 2
    v10 := GenS1{F0: 247, F1: 147}
    v1 = v1
    switch v4 {
    case "\n\t":
        var v11 uint32 = 42
        _ = v11
    case "":
        if v2 {
            for fi12 := 0; fi12 < 19; fi12++ {
                fmt.Println(v4, fi12)
                switch ((("hello" + v4) + v4) + v4) {
                case "foo bar":
                    var v13 int64 = -189
                    _ = v13
                case "\n\t":
                    v14 := GenS2{F0: 304, F1: 253}
                    _ = v14
                case "a":
                    v15 := []int{4294967295, -537}
                    _ = v15
                }
                switch v4 {
                case "\n\t":
                    fmt.Println(v4, v4)
                case "hello":
                    var v16 string = ((v4 + (v4 + v4)) + "")
                    _ = v16
                default:
                    var v17 string = (v4 + "foo bar")
                    _ = v17
                }
                _ = fi12
            }
            var v18 int64 = 1
            _ = v18
        } else {
            v19 := GenS3{F0: 65535, F1: -147}
            if (!(true && v2)) {
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 float32 = float32(1e10)
                    if ((fi20 * (fi20 - (fi20 << ((255) & 63)))) != fi20) {
                        v4 = v4
                        v22 := GenS3{F0: 0, F1: -254}
                        _ = v22
                    }
                    _ = fi20
                    _ = v21
                }
                switch 128 {
                case 1:
                    for fi23 := 0; fi23 < 20; fi23++ {
                        if (!(42 <= 230)) {
                            v24 := []int{-289}
                            v4 = ("foo bar" + (v4 + "foo bar"))
                            _ = v24
                        } else {
                            v4 = "\n\t"
                        }
                        _ = fi23
                    }
                case 32767:
                    v25 := GenS4{F0: v1}
                    _ = v25
                case 32768:
                    fmt.Println(v4, v2, v4)
                }
                if (-874 != -984) {
                    v1 -= v1
                    v26 := v3[733]
                    _ = v26
                }
            }
            _ = v19
        }
    case "a":
        fmt.Println(v6, v1)
    default:
        switch 0 {
        case 127:
            var v27 uint = 951
            _ = v27
        }
    }
    if (2 != 4294967295) {
        fmt.Println(v4)
    }
    v28 := []int{}
    var v29 uint8 = 128
    if v2 {
        v29 = (v29 >> ((65535) & 63))
    }
    if false {
        var v30 int16 = 42
        var v31 float32 = (1.5 * (-2.25 + -2.25))
        _ = v30
        _ = v31
    }
    v2 = v2
    if ((1e10 - 1.5) > 0.0) {
        var v32 uint32 = 1
        for fi33 := 0; fi33 < 1; fi33++ {
            if v2 {
                v29 -= v29
            }
            v34 := GenS5{F0: 9223372036854775807, F1: (int16(v1) != -32768)}
            if len(v28) > 0 {
                v28[0] = fi33
            }
            _ = fi33
            _ = v34
        }
        var v35 uint16 = 65535
        _ = v32
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
