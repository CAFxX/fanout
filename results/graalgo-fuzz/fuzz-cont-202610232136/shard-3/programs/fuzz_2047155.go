package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int8, p1 uint16) string {
    v7 := []int{int(p1)}
    p1 = 768
    for fi8 := 0; fi8 < 17; fi8++ {
        switch "a" {
        case "a":
            if true {
                var v9 string = "foo bar"
                v10 := GenS1{F0: 242, F1: -2147483648}
                var v11 int16 = -1
                _ = v10
                _ = v11
                _ = v9
            } else {
                v12 := GenS1{F0: 929, F1: 42}
                v13 := GenS1{F0: 988, F1: -1}
                _ = v12
                _ = v13
            }
        }
        p1 = (^(p1))
        _ = fi8
    }
    switch "foo bar" {
    case "hello":
        _ = len(v7)
    case "\n\t":
        if (0 < 0) {
            if false {
                p1 = (^(p1))
                if (924 < 9223372036854775807) {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            }
        }
    case "foo bar":
        p1 += p1
    default:
        var v14 uint8 = 111
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    return ("\n\t" + "\n\t")
}
func genFn1(p0 bool, p1 int16) bool {
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint64 = 1
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 int8 = 100
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := []string{("hello" + "a"), ("\n\t" + (("" + "hello") + ("hello" + "\n\t"))), ""}
    var v2 float32 = float32(0.0)
    var v3 bool = (963 <= 32767)
    v4 := make(map[int]string)
    var v5 uint8 = 0
    v6 := make(map[string]int)
    for fi19 := 0; fi19 < 20; fi19++ {
        switch "foo bar" {
        case "hello":
            var v20 int = fi19
            _ = v20
        case "foo bar":
            for fi21 := 0; fi21 < 19; fi21++ {
                v4[(fi19 * (fi21 % ((0) | 1)))] = ("" + "hello")
                v1 = append(v1, "foo bar")
                v2 = v2
                _ = fi21
            }
        }
        v1 = append(v1, "")
        var v22 int32 = -1
        _ = fi19
        _ = v22
    }
    var v23 uint16 = 298
    var v24 uint64 = 18446744073709551615
    _ = len(v4)
    for fi25 := 0; fi25 < 9; fi25++ {
        v23 = v23
        _ = fi25
    }
    var v26 int8 = 42
    v27 := GenS1{F0: 255, F1: -2147483648}
    switch "foo bar" {
    case "foo bar":
        if true {
            v28 := GenS2{F0: v26, F1: 9223372036854775807, F2: (v26 >> ((255) & 63))}
            _ = v28
        } else {
            if v3 {
                switch "a" {
                case "a":
                    for fi29 := 0; fi29 < 12; fi29++ {
                        for fi30 := 0; fi30 < 19; fi30++ {
                            var v31 uint32 = 61
                            v32 := GenS3{F0: -456}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        var v33 uint = 255
                        _ = fi29
                        _ = v33
                    }
                }
                if len(v1) > 0 {
                    v1[0] = ("" + "")
                }
            } else {
                switch "" {
                case "a":
                    fmt.Println(v23, v24)
                case "\n\t":
                    var v34 bool = v3
                    _ = v34
                case "hello":
                    v35 := GenS3{F0: 42}
                    _ = v35
                }
                _ = len(v6)
                var v36 bool = true
                _ = v36
            }
        }
    case "hello":
        var v37 int = -460
        _ = v37
    case "a":
        v23 = 255
    default:
        var v38 float32 = ((1.5 + (0.1 * 0.0)) * v2)
        _ = v38
    }
    v26 -= (v26 >> ((106) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
