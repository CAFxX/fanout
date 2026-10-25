package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 string; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint32) string {
    switch (("hello" + "") + ("\n\t" + "\n\t")) {
    case "foo bar":
        p1 -= (p1 >> (((^(p0))) & 63))
    case "hello":
        p1 += (p1 >> ((p0) & 63))
    case "a":
        p1 += (p1 / (((p1 << ((p0) & 63))) | 1))
    default:
        v5 := GenS1{F0: 2}
        _ = v5
    }
    v6 := []int{32768, -322}
    _ = p0
    _ = p1
    _ = v6
    return ("\n\t" + "foo bar")
}
func genFn1(p0 string) uint16 {
    v7 := []int{-2147483648}
    switch 555 {
    case 1:
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                _ = len(v7)
                var v10 uint64 = 0
                fmt.Println(fi8, p0, fi8)
                v10 -= 169
                _ = fi9
                _ = v10
            }
            p0 = p0
            var v11 uint64 = 321
            _ = fi8
            _ = v11
        }
    }
    if false {
        var v12 float32 = 0.0
        switch 0 {
        case 42:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case -753:
            p0 = "\n\t"
        default:
            v13 := GenS2{F0: 44, F1: 174}
            _ = v13
        }
        var v14 int = 9223372036854775807
        _ = v12
        _ = v14
    }
    v15 := GenS3{F0: (p0 + ((p0 + ("a" + "hello")) + p0)), F1: 19, F2: 152}
    _ = p0
    _ = v15
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 42
    var v3 uint16 = 1
    v4 := []uint32{}
    v16 := GenS4{F0: 1, F1: 229}
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            if false {
                if (42 <= 338) {
                    switch -2147483648 {
                    case 42:
                        switch -9223372036854775808 {
                        case 128:
                            delete(v1, (fi18 % (((fi17 % ((128) | 1))) | 1)))
                            var v19 uint64 = 42
                            if false {
                                switch fi17 {
                                case -9223372036854775808:
                                    fmt.Println(v2)
                                case 127:
                                    v2 = v2
                                case 1:
                                    if true {
                                        v3 -= 42
                                    }
                                }
                                v4 = append(v4, uint32(v3))
                            }
                            var v20 float32 = 3.14159
                            var v21 int32 = 2147483647
                            _ = v19
                            _ = v20
                            _ = v21
                        default:
                            var v22 int8 = (v2 / (((v2 ^ v2)) | 1))
                            _ = v22
                        }
                    default:
                        v2 = v2
                    }
                    v23 := GenS2{F0: 170, F1: v3}
                    fmt.Println(v3, fi18, v2)
                    _ = v23
                }
                var v24 int = fi17
                _ = v24
            }
            fmt.Println(v2, v2)
            fmt.Println(fi18)
            _ = fi18
        }
        _ = fi17
    }
    var v25 int8 = 127
    if true {
        fmt.Println(v2)
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        for fi27 := 0; fi27 < 4; fi27++ {
            if true {
                switch fi27 {
                case 127:
                    v28 := v1[(fi26 >> ((255) & 63))]
                    _ = v28
                case -739:
                    switch (fi27 >> ((42) & 63)) {
                    case 128:
                        var v29 int = (fi27 & 1)
                        _ = v29
                    case 0:
                        var v30 int = 1
                        _ = v30
                    }
                case 42:
                    _ = len(v4)
                }
                v2 = (v2 * -9)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v2 = 42
                var v31 int = 127
                _ = v31
            }
            v32 := GenS5{F0: (-(v25))}
            _ = fi27
            _ = v32
        }
        v33 := GenS6{F0: 285}
        _ = fi26
        _ = v33
    }
    var v34 int32 = 0
    var v35 float64 = 0.1
    switch "" {
    case "hello":
        v35 += (v35 * float64(((float32((v35 + v35)) / 1.5) * 1.5)))
    default:
        var v36 float32 = ((float32(((v35 / (1e10 - (1.5 / (float64(3.14159) + v35)))) / ((v35 / (v35 * (v35 + v35))) * (v35 + (1.5 * (v35 - (3.14159 - float64(1.5)))))))) * (((1.5 + (float32(float64(-2.25)) - 1e10)) / 3.14159) + 0.0)) / 1e10)
        _ = v36
    }
    var v37 int16 = 720
    var v38 uint32 = 1
    switch "a" {
    case "hello":
        var v39 int8 = v25
        _ = v39
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v40 := GenS3{F0: "\n\t", F1: 1, F2: 127}
    if false {
        v1[int(v25)] = v38
        var v41 int = 804
        var v42 int8 = (v25 | 100)
        _ = v41
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v40
}
