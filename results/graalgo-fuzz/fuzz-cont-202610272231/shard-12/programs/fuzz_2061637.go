package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 bool; F1 int8; F2 uint32 }

func main() {
    var v1 uint8 = 255
    var v2 int32 = 42
    v3 := make(map[string]uint32)
    v4 := []uint32{85, 4294967295, 351, 924}
    v5 := make(map[string]int)
    v6 := make(map[int]bool)
    v7 := GenS1{F0: ("\n\t" + ("\n\t" + "foo bar")), F1: 2}
    var v8 int32 = v2
    var v9 uint16 = 65535
    v6[-2147483648] = (!true)
    switch -2147483648 {
    case 285:
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint64 = 42
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 float64 = 0.0
                var v14 int32 = 966
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    case 711:
        fmt.Println(v1, v9, v8)
    case -9223372036854775808:
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    default:
        v1 = 42
    }
    var v15 int16 = 0
    if (v8 == v8) {
        var v16 int8 = -101
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint64 = 42
        _ = fi17
        _ = v18
    }
    var v19 uint64 = 2
    fmt.Println(v15, v1)
    v20 := GenS1{F0: ("foo bar" + "\n\t"), F1: v1}
    _ = len(v5)
    if true {
        for fi21 := 0; fi21 < 20; fi21++ {
            v19 -= (v19 % ((v19) | 1))
            v19 -= v19
            v22 := GenS1{F0: "", F1: v1}
            _ = fi21
            _ = v22
        }
        fmt.Println(v8)
        v23 := GenS2{F0: 550, F1: (!((v15 % ((int16(v19)) | 1)) != v15))}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := GenS3{F0: 2, F1: 9223372036854775807}
            if true {
                v26 := GenS4{F0: (v1 | v1), F1: 535}
                _ = v26
            } else {
                if (v19 == v19) {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v9 -= v9
                        _ = fi27
                    }
                    fmt.Println(v8, v8)
                    v1 = v1
                } else {
                    fmt.Println(v19)
                    var v28 bool = ((v1 | 2) != 127)
                    _ = v28
                }
                v29 := GenS5{F0: ((^(v15)) < 1), F1: 76, F2: 148}
                switch 32768 {
                case 42:
                    var v30 int8 = -128
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 19; fi31++ {
                        v1 -= 0
                        var v32 string = ("a" + (("a" + ("\n\t" + (("\n\t" + "foo bar") + ""))) + (("foo bar" + "foo bar") + ("\n\t" + ("foo bar" + "\n\t")))))
                        v6[4294967295] = ((v9 ^ (v9 * v9)) == uint16(v1))
                        v1 = (^(v1))
                        fmt.Println(v32, v8)
                        fmt.Println(v8)
                        _ = fi31
                        _ = v32
                    }
                }
                _ = v29
            }
            var v33 string = "a"
            _ = fi24
            _ = v25
            _ = v33
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v2)
    if len(v4) > 0 {
        v4[0] = 871
    }
    v15 += (v15 * v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
