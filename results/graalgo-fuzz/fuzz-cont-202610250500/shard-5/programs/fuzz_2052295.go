package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int8 = 0
    var v2 int16 = -341
    var v3 uint64 = 4294967295
    var v4 uint8 = 2
    var v5 int16 = v2
    v6 := []int{-415, -507, 9223372036854775807, -2147483648}
    v7 := GenS1{F0: v1}
    fmt.Println(v5, v1, v3)
    if false {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: -128}
            _ = fi8
            _ = v9
        }
        if (v5 == -1) {
            var v10 int8 = v1
            switch "" {
            case "foo bar":
                if ("\n\t" >= "hello") {
                    fmt.Println(v4)
                }
            case "\n\t":
                fmt.Println(v5)
            }
            var v11 int16 = (v2 & 0)
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int32 = -473
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 int = fi14
            v16 := GenS1{F0: v1}
            switch fi14 {
            case 993:
                if (!true) {
                    if true {
                        var v17 uint16 = 42
                        v13 = v13
                        if true {
                            _ = len(v6)
                            v1 -= v1
                            var v18 uint = 4294967295
                            _ = v18
                        } else {
                            v19 := GenS1{F0: v1}
                            _ = v19
                        }
                        _ = v17
                    }
                    switch (fi14 >> ((883) & 63)) {
                    case 2147483647:
                        fmt.Println(v2)
                    case -1:
                        fmt.Println(v2)
                        _ = len(v6)
                        v20 := GenS1{F0: v1}
                        _ = v20
                    case 32767:
                        var v21 uint = 255
                        _ = v21
                    default:
                        if false {
                            v15 += 2147483647
                            switch 32768 {
                            case 127:
                                v13 = (v13 >> ((233) & 63))
                            case -1:
                                switch "\n\t" {
                                case "":
                                    var v22 float64 = (0.0 - float64((3.14159 + (((1e10 / -2.25) + float32((3.14159 - 1e10))) * float32((1e10 / (3.14159 * 0.1)))))))
                                    _ = v22
                                case "\n\t":
                                    var v23 float32 = (1e10 * (((-2.25 - -2.25) / 1.5) + 3.14159))
                                    _ = v23
                                default:
                                    v24 := GenS1{F0: (v1 + v1)}
                                    _ = v24
                                }
                            }
                            var v25 uint32 = 850
                            v4 -= 155
                            _ = v25
                        }
                    }
                    if (!false) {
                        fmt.Println(v3, v15, v5)
                        if true {
                            _ = len(v6)
                        }
                    }
                    _ = len(v6)
                    v2 -= (-(v5))
                    if true {
                        var v26 uint64 = uint64(v4)
                        _ = v26
                    } else {
                        switch (v15 << ((1) & 63)) {
                        case -883:
                            if (!true) {
                                if false {
                                    _ = len(v6)
                                    switch (fi12 << ((65535) & 63)) {
                                    case 42:
                                        v27 := GenS1{F0: (v1 >> ((97) & 63))}
                                        _ = v27
                                    case 355:
                                        if true {
                                            var v28 float64 = 0.0
                                            _ = v28
                                        }
                                        fmt.Println(v5, v1, v15)
                                        v29 := GenS1{F0: v1}
                                        _ = v29
                                    case 491:
                                        _ = len(v6)
                                    }
                                } else {
                                    switch "" {
                                    case "hello":
                                        v6 = append(v6, 2147483647)
                                        v30 := GenS1{F0: 1}
                                        if len(v6) > 0 {
                                            v6[0] = (fi12 & (v15 >> ((65535) & 63)))
                                        }
                                        _ = v30
                                    default:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        fmt.Println(v13, v2, v1)
                                    }
                                }
                            }
                        case 42:
                            switch "foo bar" {
                            case "hello":
                                v15 = (-(fi12))
                            case "foo bar":
                                v15 += fi14
                            case "":
                                if ((((0.0 - (((0.0 / 1.5) - -2.25) - 0.0)) - (0.1 * 0.1)) >= 0.0) || true) {
                                    v1 += 127
                                    fmt.Println(v13, v13)
                                    if true {
                                        fmt.Println(fi14, v5, v4)
                                        v4 -= v4
                                    } else {
                                        v2 += v5
                                    }
                                }
                            default:
                                v3 += v3
                            }
                        default:
                            v3 = v3
                        }
                        var v31 float64 = 1e10
                        _ = v31
                    }
                    v32 := GenS2{F0: 434}
                    _ = v32
                }
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v2, v5, v4)
    var v33 uint16 = 455
    var v34 uint16 = 2
    v35 := GenS3{F0: -32768}
    v2 -= (v2 << ((65535) & 63))
    if true {
        var v36 int8 = v1
        var v37 uint8 = (v4 - (^(v4)))
        for fi38 := 0; fi38 < 1; fi38++ {
            v39 := GenS1{F0: int8(v2)}
            _ = fi38
            _ = v39
        }
        _ = v36
        _ = v37
    }
    if (0.0 < 0.1) {
        var v40 int16 = v2
        var v41 float32 = ((((0.0 - float32(float64(-2.25))) - 0.1) - float32((-2.25 * -2.25))) + (-2.25 / 1e10))
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
