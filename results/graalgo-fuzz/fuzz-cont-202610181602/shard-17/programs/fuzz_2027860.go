package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 uint }

func main() {
    var v1 string = ("\n\t" + (("hello" + "\n\t") + "\n\t"))
    var v2 int64 = 0
    var v3 int16 = -32768
    var v4 int16 = (v3 << ((42) & 63))
    if false {
        if (-1 < 460) {
            var v5 int8 = -2
            for fi6 := 0; fi6 < 10; fi6++ {
                fmt.Println(v4, fi6, v2)
                _ = fi6
            }
            _ = v5
        } else {
            v7 := []int{-369, 127, 635, 1}
            for fi8 := 0; fi8 < 19; fi8++ {
                v1 = ((("a" + ("hello" + "a")) + (v1 + "foo bar")) + (v1 + ((("" + ((((("" + v1) + "hello") + v1) + ("foo bar" + v1)) + v1)) + ((("\n\t" + ("" + ((v1 + v1) + "hello"))) + (v1 + v1)) + v1)) + ((v1 + ("" + (v1 + ("hello" + v1)))) + (v1 + v1)))))
                _ = fi8
            }
            v2 = 0
            _ = v7
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            if (v1 == "foo bar") {
                v10 := []int{(fi9 * (-(fi9))), (fi9 / (((fi9 - fi9)) | 1)), fi9, (fi9 / ((fi9) | 1))}
                _ = v10
            }
            _ = fi9
        }
        v11 := []int{769}
        _ = v11
    } else {
        var v12 uint = 1
        _ = v12
    }
    var v13 int32 = -573
    v4 += 32767
    fmt.Println(v13)
    switch (v1 + v1) {
    case "hello":
        var v14 uint32 = 817
        _ = v14
    case "a":
        v3 = -738
    case "\n\t":
        var v15 uint64 = 700
        _ = v15
    }
    if (false && ((((false || ((128 == 2) && false)) && false) && false) && (255 != 4294967295))) {
        fmt.Println(v13, v1)
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: 32767, F1: 255}
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 1, F1: 18446744073709551615}
    v19 := []int{}
    var v20 uint64 = 956
    v13 -= 0
    if len(v19) > 0 {
        v19[0] = 128
    }
    v21 := GenS3{F0: 1}
    for fi22 := 0; fi22 < 2; fi22++ {
        if (!true) {
            if (-734 == v13) {
                v19 = append(v19, (fi22 / ((882) | 1)))
                if true {
                    fmt.Println(v13, v13)
                    for fi23 := 0; fi23 < 7; fi23++ {
                        switch ((("a" + ("\n\t" + ("" + v1))) + "\n\t") + (((("hello" + v1) + ((("\n\t" + "foo bar") + ("\n\t" + "")) + v1)) + v1) + "foo bar")) {
                        case "hello":
                            var v24 uint = 15
                            _ = v24
                        case "":
                            var v25 int = (fi22 >> ((42) & 63))
                            _ = v25
                        }
                        _ = fi23
                    }
                    if ((float32(((1e10 / -2.25) + (-2.25 / -2.25))) * float32(0.1)) != 3.14159) {
                        v4 -= -669
                        fmt.Println(v13, fi22)
                        v26 := GenS1{F0: 0, F1: 65535}
                        _ = v26
                    } else {
                        fmt.Println(v20)
                    }
                } else {
                    var v27 uint = 4294967295
                    if false {
                        switch (fi22 >> (((v27 + 255)) & 63)) {
                        case -2147483648:
                            fmt.Println(v2, v13, fi22)
                        case 4294967295:
                            v20 -= v20
                        case 42:
                            v1 = ((v1 + v1) + ("foo bar" + v1))
                        default:
                            fmt.Println(v13)
                        }
                    }
                    v3 = v3
                    _ = v27
                }
                switch fi22 {
                case 729:
                    v1 = (v1 + v1)
                case 32767:
                    v13 = (-(v13))
                case 477:
                    v2 = (v2 & (v2 + 42))
                default:
                    v2 = v2
                }
            }
        }
        _ = fi22
    }
    var v28 uint64 = v20
    if true {
        v13 = (v13 ^ v13)
        v29 := GenS4{F0: -128, F1: 1}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v4
}
