package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 string; F2 int64 }
type GenS4 struct { F0 int; F1 uint8 }

func main() {
    var v1 int32 = 2147483647
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]int)
    v4 := []int{-173, -2147483648, -1, 128}
    var v5 uint = 42
    var v6 uint = 2
    var v7 uint = (v5 >> ((uint(v2)) & 63))
    var v8 float32 = 1.5
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: -32768, F1: true}
            for fi11 := 0; fi11 < 10; fi11++ {
                v1 -= (-(v1))
                var v12 int = 2147483647
                fmt.Println(v1, v5, v8)
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        var v13 int64 = 1
        _ = v13
    } else {
        if ((^(v2)) < v2) {
            if true {
                fmt.Println(v5, v8)
                for fi14 := 0; fi14 < 9; fi14++ {
                    for fi15 := 0; fi15 < 5; fi15++ {
                        if false {
                            v16 := GenS1{F0: 1, F1: (0.1 <= (1.5 + ((float32(0.0) + (((-2.25 * ((1.5 / v8) / (((v8 + v8) * ((v8 - (0.1 + (float32(3.14159) / v8))) + -2.25)) / 1.5))) / float32(1.5)) / v8)) - ((((((-2.25 + (v8 * 0.1)) + v8) / v8) - 0.0) - 1e10) * (v8 - v8)))))}
                            _ = v16
                        }
                        var v17 uint = 18446744073709551615
                        switch ("hello" + (("foo bar" + "a") + "\n\t")) {
                        case "foo bar":
                            if false {
                                if true {
                                    var v18 float32 = v8
                                    switch fi15 {
                                    case 32767:
                                        if true {
                                            _ = len(v3)
                                            fmt.Println(v2)
                                            switch ("" + ("hello" + "")) {
                                            case "":
                                                var v19 uint16 = 2
                                                _ = v19
                                            case "a":
                                                var v20 uint64 = 4294967295
                                                _ = v20
                                            case "\n\t":
                                                var v21 int16 = 1
                                                _ = v21
                                            }
                                        }
                                        fmt.Println(v18, v5)
                                    case -393:
                                        v22 := GenS2{F0: v7}
                                        _ = v22
                                    case 127:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        v23 := GenS3{F0: 1, F1: "hello", F2: -9223372036854775808}
                                        _ = v23
                                    }
                                    v18 = float32((0.0 / (3.14159 - 1.5)))
                                    _ = v18
                                } else {
                                    if true {
                                        switch (fi14 % ((-1) | 1)) {
                                        case 524:
                                            v17 -= 42
                                        default:
                                            fmt.Println(v17, v5, v5)
                                        }
                                        switch fi15 {
                                        case -147:
                                            _ = len(v3)
                                        }
                                        if true {
                                            v24 := v3[(v2 & (v2 % (((v2 - (v2 << ((v6) & 63)))) | 1)))]
                                            v2 -= v2
                                            var v25 int = 9223372036854775807
                                            _ = v24
                                            _ = v25
                                        }
                                    } else {
                                        switch "foo bar" {
                                        case "\n\t":
                                            v1 = 42
                                        case "a":
                                            fmt.Println(v7)
                                        }
                                        var v26 int = fi15
                                        var v27 int16 = -1
                                        _ = v26
                                        _ = v27
                                    }
                                    v28 := GenS1{F0: 1000, F1: true}
                                    var v29 uint = v7
                                    fmt.Println(v8, v7)
                                    delete(v3, v2)
                                    v17 = 255
                                    _ = v28
                                    _ = v29
                                }
                            } else {
                                _ = len(v3)
                            }
                        case "":
                            _ = len(v4)
                        }
                        _ = fi15
                        _ = v17
                    }
                    v8 = v8
                    _ = fi14
                }
                fmt.Println(v8, v1)
            }
            v6 -= 0
        }
    }
    if (!(!true)) {
        var v30 uint8 = 154
        if true {
            if false {
                if false {
                    if true {
                        var v31 int16 = -959
                        v3[v2] = 828
                        _ = v31
                    } else {
                        if (((1.5 / v8) / v8) != (float32(1.5) * (v8 - v8))) {
                            var v32 uint32 = (v2 / (((^(v2))) | 1))
                            _ = v32
                        }
                    }
                }
                fmt.Println(v1)
            } else {
                if ((^(v2)) <= v2) {
                    v7 = v6
                } else {
                    switch int(v30) {
                    case 42:
                        for fi33 := 0; fi33 < 5; fi33++ {
                            var v34 uint32 = v2
                            delete(v3, uint32(v30))
                            v35 := v3[v2]
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    case 127:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    switch "a" {
                    case "hello":
                        var v36 bool = ((false || (255 == v6)) || true)
                        _ = v36
                    case "foo bar":
                        if false {
                            v37 := GenS2{F0: v5}
                            v5 -= (v5 - uint(v30))
                            fmt.Println(v8, v5, v6)
                            _ = v37
                        }
                    }
                    v38 := GenS4{F0: 128, F1: v30}
                    _ = v38
                }
                for fi39 := 0; fi39 < 1; fi39++ {
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 int = fi39
                        v4 = append(v4, (v41 | 42))
                        _ = fi40
                        _ = v41
                    }
                    _ = fi39
                }
            }
            v8 = (((0.0 - -2.25) + (v8 + ((v8 + ((1.5 / v8) + v8)) + (-2.25 * (v8 - ((v8 / v8) / ((v8 * ((1.5 / v8) - v8)) * v8))))))) / v8)
        } else {
            var v42 uint32 = v2
            for fi43 := 0; fi43 < 13; fi43++ {
                switch ((((("a" + "foo bar") + "") + (("a" + ("\n\t" + "")) + (("foo bar" + "a") + "a"))) + ("hello" + (("" + ("hello" + "hello")) + "hello"))) + "\n\t") {
                case "hello":
                    v8 = v8
                case "":
                    v44 := GenS1{F0: 640, F1: (!(("" + (("a" + "a") + "foo bar")) > "hello"))}
                    _ = v44
                }
                _ = fi43
            }
            if (0 >= -128) {
                var v45 int = -137
                v30 = v30
                _ = v45
            }
            _ = v42
        }
        v8 = 1e10
        _ = v30
    }
    var v46 int8 = 127
    v7 -= v7
    v47 := GenS1{F0: 32767, F1: (32767 <= 4294967295)}
    for fi48 := 0; fi48 < 16; fi48++ {
        v5 = 255
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
