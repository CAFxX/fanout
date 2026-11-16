package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint32) uint8 {
    if true {
        v4 := []int{168, -648}
        _ = v4
    } else {
        p0 -= (p0 << ((255) & 63))
        fmt.Println(p0)
    }
    p0 -= 722
    v5 := GenS1{F0: 127, F1: true}
    p0 -= p0
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 bool = (true || (3.14159 != 0.1))
            var v9 uint16 = 454
            v10 := GenS1{F0: -1, F1: false}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        v11 := GenS1{F0: -128, F1: ((-285 != 42) && true)}
        _ = fi6
        _ = v11
    }
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 int64) uint32 {
    p0 -= 0
    fmt.Println(p0)
    _ = p0
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    v3 := []string{""}
    var v12 int16 = 0
    fmt.Println(v12)
    delete(v1, "a")
    for fi13 := 0; fi13 < 14; fi13++ {
        _ = len(v3)
        switch "hello" {
        case "hello":
            switch "hello" {
            case "":
                if (-1 != -427) {
                    if (32767 >= 2) {
                        v14 := GenS1{F0: -1, F1: true}
                        if false {
                            switch (("hello" + "\n\t") + ("hello" + "\n\t")) {
                            case "":
                                _ = len(v3)
                            case "hello":
                                var v15 float64 = (1e10 + float64(1.5))
                                _ = v15
                            case "foo bar":
                                var v16 bool = true
                                _ = v16
                            default:
                                fmt.Println(v12, v12)
                            }
                            v2[genFn1(-568)] = ("foo bar" + "")
                        } else {
                            for fi17 := 0; fi17 < 10; fi17++ {
                                switch "a" {
                                case "foo bar":
                                    var v18 int64 = 42
                                    _ = v18
                                case "":
                                    v19 := GenS2{F0: (v12 * (v12 - (v12 | int16(fi17)))), F1: 255, F2: 42}
                                    _ = v19
                                }
                                v20 := GenS2{F0: v12, F1: 128, F2: genFn0(genFn1(-235))}
                                v21 := GenS1{F0: -50, F1: (9223372036854775807 == 9223372036854775807)}
                                _ = fi17
                                _ = v20
                                _ = v21
                            }
                            v22 := GenS3{F0: 228, F1: fi13}
                            _ = v22
                        }
                        _ = v14
                    }
                }
            case "hello":
                var v23 uint8 = 220
                _ = v23
            case "\n\t":
                for fi24 := 0; fi24 < 14; fi24++ {
                    v12 += v12
                    _ = fi24
                }
            }
        case "a":
            var v25 int64 = -1
            _ = v25
        }
        v1[""] = (1 == 255)
        _ = fi13
    }
    var v26 int32 = 42
    v26 -= (v26 * (v26 % ((v26) | 1)))
    for fi27 := 0; fi27 < 20; fi27++ {
        v12 += v12
        var v28 int64 = 42
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS4{F0: 154, F1: genFn0(42)}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (0 < 18446744073709551615) {
                v12 = (v12 >> ((255) & 63))
                v28 += -789
            }
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    var v31 int8 = 127
    for fi32 := 0; fi32 < 11; fi32++ {
        v26 = v26
        var v33 int = (fi32 ^ -245)
        _ = fi32
        _ = v33
    }
    var v34 float64 = 0.1
    if (255 >= 0) {
        var v35 int8 = 0
        var v36 float64 = ((float64(1e10) - (v34 / v34)) * (v34 / float64((float32(v34) - 3.14159))))
        var v37 uint32 = 2147483648
        _ = v35
        _ = v36
        _ = v37
    }
    v38 := GenS4{F0: 418, F1: 164}
    v26 += v26
    var v39 int32 = v26
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v34
    _ = v38
    _ = v39
}
