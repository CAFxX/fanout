package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0(p0 int) uint64 {
    if true {
        p0 -= (p0 & (p0 | (p0 << ((620) & 63))))
    } else {
        switch "" {
        case "hello":
            v7 := []int{}
            _ = v7
        case "":
            var v8 int = 127
            _ = v8
        case "\n\t":
            if false {
                var v9 uint = 468
                var v10 uint64 = 18446744073709551615
                _ = v10
                _ = v9
            } else {
                var v11 int = (-(p0))
                v11 = (v11 | (p0 + (v11 << ((4294967295) & 63))))
                _ = v11
            }
        }
    }
    if (-44 <= -128) {
        p0 += (p0 ^ (-(p0)))
        var v12 int16 = 42
        v13 := GenS1{F0: 1}
        _ = v12
        _ = v13
    }
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float32 = 3.14159
    var v3 int8 = 48
    var v4 float32 = v2
    var v5 uint64 = 42
    var v6 float32 = v4
    v6 += 1.5
    var v14 uint64 = 659
    v15 := []int{4294967295, -1, -241}
    v4 = v6
    if (true || ((v14 | (^(v5))) >= (^(v14)))) {
        v3 -= v3
        var v16 uint = 4294967295
        for fi17 := 0; fi17 < 4; fi17++ {
            if (v3 > (v3 | 100)) {
                if (v14 <= uint64(v1)) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        switch ("a" + "hello") {
                        case "\n\t":
                            if true {
                                v19 := GenS1{F0: 65535}
                                _ = v19
                            }
                        default:
                            var v20 string = ""
                            fmt.Println(v3, v20)
                            _ = v20
                        }
                        _ = fi18
                    }
                    fmt.Println(v16, v16, v1)
                    switch fi17 {
                    case -1:
                        switch fi17 {
                        case 107:
                            fmt.Println(fi17)
                        case 127:
                            v21 := GenS1{F0: 2}
                            _ = v21
                        case -1:
                            if (fi17 <= fi17) {
                                if (true && true) {
                                    if ("" != "a") {
                                        for fi22 := 0; fi22 < 19; fi22++ {
                                            switch "a" {
                                            case "":
                                                if (52 > fi22) {
                                                    switch ((("" + "") + ("" + ((("a" + "") + ("" + "hello")) + "hello"))) + "a") {
                                                    case "foo bar":
                                                        v1 += v1
                                                    }
                                                    var v23 uint8 = 0
                                                    var v24 uint64 = (^(v5))
                                                    _ = v23
                                                    _ = v24
                                                } else {
                                                    var v25 int8 = v3
                                                    switch "a" {
                                                    case "\n\t":
                                                        fmt.Println(v25, v16)
                                                    default:
                                                        v2 = v4
                                                    }
                                                    if len(v15) > 0 {
                                                        v15[0] = (fi22 | fi17)
                                                    }
                                                    _ = v25
                                                }
                                            case "\n\t":
                                                fmt.Println(v16, v14, v2)
                                            case "hello":
                                                switch "\n\t" {
                                                case "hello":
                                                    fmt.Println(v16)
                                                case "\n\t":
                                                    _ = len(v15)
                                                }
                                            default:
                                                var v26 string = (("\n\t" + "") + ("a" + "hello"))
                                                _ = v26
                                            }
                                            fmt.Println(v2, v6, fi17)
                                            v27 := GenS2{F0: 255}
                                            _ = fi22
                                            _ = v27
                                        }
                                    }
                                    var v28 float32 = float32(3.14159)
                                    _ = v28
                                } else {
                                    v2 = 0.1
                                    v16 -= (^(v16))
                                    if false {
                                        v16 += (v16 * v16)
                                        v29 := GenS2{F0: 163}
                                        _ = v29
                                    }
                                }
                            }
                        }
                    }
                }
                v30 := GenS3{F0: uint32(v5), F1: 65535}
                _ = v30
            }
            var v31 float32 = (-2.25 + v6)
            fmt.Println(v16, v2)
            _ = fi17
            _ = v31
        }
        _ = v16
    } else {
        v2 += (0.1 * v4)
        v5 = (v5 / (((^(v14))) | 1))
        var v32 int8 = int8(v14)
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 int = 32767
        switch v34 {
        case 32768:
            for fi35 := 0; fi35 < 9; fi35++ {
                v15 = append(v15, (v34 << ((255) & 63)))
                switch "hello" {
                case "a":
                    var v36 uint64 = genFn0(1)
                    _ = v36
                case "\n\t":
                    v37 := GenS4{F0: 65535, F1: -433}
                    _ = v37
                case "foo bar":
                    v6 = v6
                }
                _ = fi35
            }
        case 720:
            var v38 int32 = 207
            _ = v38
        }
        fmt.Println(v5, v34)
        _ = fi33
        _ = v34
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        var v40 int = fi39
        _ = fi39
        _ = v40
    }
    if true {
        v41 := GenS3{F0: (v1 - v1), F1: 255}
        _ = v41
    } else {
        v14 = (v14 - v14)
    }
    fmt.Println(v5, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
