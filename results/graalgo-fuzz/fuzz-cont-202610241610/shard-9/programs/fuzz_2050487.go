package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 string }
func genFn0(p0 int16, p1 bool, p2 int32) int64 {
    p1 = ((0.0 < (1.5 / (0.1 / (0.0 + -2.25)))) || (true || (0 > 872)))
    v5 := GenS1{F0: 982}
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(p1, p0)
        if (p0 <= 1000) {
            var v7 uint = 899
            var v8 int32 = 2147483647
            _ = v7
            _ = v8
        } else {
            p0 = p0
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -566
}

func main() {
    v1 := []int{0, -483}
    var v2 uint64 = 4294967295
    var v3 int64 = 0
    v4 := make(map[int]int)
    for fi9 := 0; fi9 < 5; fi9++ {
        if (54 == -71) {
            v4[32767] = (fi9 - 32768)
            v3 = (^(v3))
        }
        var v10 int8 = 29
        fmt.Println(v2)
        _ = fi9
        _ = v10
    }
    if (0 > 32767) {
        if true {
            fmt.Println(v2, v3)
        }
        v11 := GenS2{F0: 2147483647}
        _ = v11
    }
    fmt.Println(v3)
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS2{F0: 42}
        v14 := GenS2{F0: 0}
        v3 -= v3
        _ = fi12
        _ = v13
        _ = v14
    }
    v4[1] = 128
    v2 += 2
    _ = len(v1)
    fmt.Println(v2)
    if true {
        v4[681] = -855
        switch -2147483648 {
        case -921:
            for fi15 := 0; fi15 < 6; fi15++ {
                switch (((("hello" + "hello") + ((("\n\t" + ("\n\t" + "")) + ("" + ("\n\t" + ((("" + "a") + "") + "\n\t")))) + "foo bar")) + "hello") + "hello") {
                case "hello":
                    _ = len(v4)
                default:
                    switch fi15 {
                    case 691:
                        delete(v4, (fi15 * fi15))
                    case 293:
                        switch fi15 {
                        case -713:
                            var v16 int64 = genFn0(-32768, false, 680)
                            _ = v16
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case 16:
                        v3 = v3
                    default:
                        var v17 string = ("" + "hello")
                        _ = v17
                    }
                }
                fmt.Println(fi15, v2)
                v2 -= (v2 ^ 1)
                _ = fi15
            }
        default:
            v2 = v2
        }
        v18 := GenS3{F0: -1, F1: "hello"}
        _ = v18
    }
    v2 += (v2 % ((v2) | 1))
    var v19 uint = 65535
    if false {
        v19 = v19
    }
    if ((9223372036854775807 >= -1) && true) {
        delete(v4, -1)
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS1{F0: 862}
            for fi22 := 0; fi22 < 13; fi22++ {
                v19 = v19
                v23 := GenS3{F0: 2147483647, F1: "\n\t"}
                v2 += 1
                v4[fi22] = fi22
                _ = len(v1)
                var v24 float32 = 0.0
                _ = fi22
                _ = v23
                _ = v24
            }
            switch "a" {
            case "foo bar":
                v1 = append(v1, (fi20 % (((fi20 / (((fi20 - fi20)) | 1))) | 1)))
            case "\n\t":
                fmt.Println(v19, v3)
            case "":
                var v25 uint16 = 65535
                _ = v25
            }
            _ = fi20
            _ = v21
        }
        var v26 int32 = -736
        _ = v26
    }
    var v27 string = "\n\t"
    var v28 bool = true
    v29 := v4[32767]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
