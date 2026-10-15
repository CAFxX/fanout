package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 bool; F2 int64 }
type GenS3 struct { F0 uint16; F1 int; F2 int8 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int64; F1 uint64; F2 int16 }
func genFn0(p0 int, p1 int, p2 bool) int64 {
    var v6 uint32 = 943
    if (true || p2) {
        for fi7 := 0; fi7 < 4; fi7++ {
            fmt.Println(fi7)
            p0 -= -969
            v8 := GenS1{F0: -952}
            _ = fi7
            _ = v8
        }
        var v9 int16 = -32768
        if (!(int8(p1) > 1)) {
            fmt.Println(p1)
        }
        _ = v9
    } else {
        var v10 bool = true
        fmt.Println(p0, p1, p0)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 uint8, p1 int64) uint64 {
    v11 := GenS2{F0: -1, F1: true, F2: (p1 | p1)}
    var v12 string = "foo bar"
    switch v12 {
    case "\n\t":
        p1 += p1
    case "a":
        for fi13 := 0; fi13 < 14; fi13++ {
            switch v12 {
            case "\n\t":
                p0 -= (p0 | 128)
            default:
                if (937 <= 4294967295) {
                    if true {
                        fmt.Println(v12)
                    } else {
                        switch v12 {
                        case "hello":
                            if false {
                                fmt.Println(fi13)
                                if (((2 < 1) || false) && (0.1 <= float32(0.0))) {
                                    v14 := []int{(fi13 + -130), fi13, (fi13 & (fi13 | 4294967295))}
                                    for fi15 := 0; fi15 < 10; fi15++ {
                                        var v16 uint16 = 0
                                        _ = fi15
                                        _ = v16
                                    }
                                    _ = v14
                                }
                                v17 := []int{(fi13 / ((-2147483648) | 1)), 32768, -617}
                                _ = v17
                            }
                        case "\n\t":
                            p1 -= p1
                        default:
                            p1 = (p1 % (((p1 - p1)) | 1))
                        }
                    }
                    if true {
                        fmt.Println(fi13)
                        fmt.Println(fi13)
                        v18 := GenS2{F0: 32767, F1: false, F2: -678}
                        _ = v18
                    } else {
                        if false {
                            var v19 float64 = (3.14159 / ((1e10 + 0.0) - (3.14159 - 1e10)))
                            for fi20 := 0; fi20 < 6; fi20++ {
                                switch -102 {
                                case -120:
                                    var v21 bool = (0.1 == (float32((v19 / (1.5 + v19))) - 1e10))
                                    _ = v21
                                default:
                                    var v22 bool = (!false)
                                    _ = v22
                                }
                                var v23 bool = (true && true)
                                _ = fi20
                                _ = v23
                            }
                            p1 += (^(p1))
                            _ = v19
                        } else {
                            p0 -= (^(p0))
                            var v24 int32 = -607
                            _ = v24
                        }
                        var v25 float64 = (0.1 * (1e10 / float64(3.14159)))
                        v25 -= v25
                        _ = v25
                    }
                } else {
                    p1 += (p1 >> ((346) & 63))
                    switch (((v12 + "hello") + v12) + "hello") {
                    case "a":
                        p1 = -1
                    }
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v12 = v12
                        if (((p1 > (p1 >> ((4294967295) & 63))) || (!(42 <= 996))) || (902 <= 989)) {
                            p0 = (p0 + (p0 / ((p0) | 1)))
                        }
                        switch v12 {
                        case "\n\t":
                            if false {
                                v12 = ""
                            }
                        default:
                            v27 := []int{1}
                            _ = v27
                        }
                        _ = fi26
                    }
                }
            }
            if false {
                v28 := GenS1{F0: 42}
                _ = v28
            }
            _ = fi13
        }
    default:
        v29 := GenS3{F0: 65535, F1: -316, F2: 127}
        _ = v29
    }
    p1 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 465
    var v3 int = 128
    var v4 int32 = 42
    v5 := make(map[int]uint32)
    v4 += (v4 >> ((65535) & 63))
    fmt.Println(v3)
    v30 := GenS4{F0: -1, F1: (^(v3))}
    fmt.Println(v1)
    v3 = v3
    switch ((("foo bar" + ("" + "foo bar")) + (("a" + "hello") + ((("hello" + (("\n\t" + ("foo bar" + "a")) + "\n\t")) + "hello") + "\n\t"))) + "foo bar") {
    case "":
        fmt.Println(v1, v1, v1)
    }
    var v31 uint32 = 1
    v31 -= (v31 ^ v31)
    v4 += (-(v4))
    v1 = v1
    switch "" {
    case "\n\t":
        var v32 int32 = (-(v4))
        _ = v32
    default:
        v33 := []int{}
        _ = v33
    }
    v34 := GenS5{F0: -195, F1: 42, F2: -32768}
    var v35 int = v3
    fmt.Println(v3, v31, v1)
    var v36 int16 = -898
    if false {
        var v37 float32 = -2.25
        delete(v5, (v35 & int(v31)))
        v38 := GenS4{F0: v36, F1: v35}
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
