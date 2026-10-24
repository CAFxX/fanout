package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 bool }

func main() {
    v1 := []uint32{}
    var v2 int8 = -1
    var v3 bool = (2 >= 255)
    var v4 uint16 = 1
    var v5 int64 = -9223372036854775808
    var v6 int8 = v2
    switch 1 {
    case 272:
        if ("\n\t" > "a") {
            v4 = v4
        } else {
            v7 := GenS1{F0: -1, F1: (v3 || (false && (("" < "foo bar") && v3)))}
            _ = v7
        }
    case 310:
        switch 42 {
        case 127:
            v8 := GenS1{F0: v2, F1: false}
            _ = v8
        case 584:
            var v9 uint8 = 1
            _ = v9
        }
    default:
        v2 = v2
    }
    v6 = -68
    var v10 int = 127
    _ = len(v1)
    var v11 int16 = 117
    if v3 {
        fmt.Println(v5)
        var v12 bool = v3
        _ = v12
    } else {
        v6 = (v2 & v6)
        for fi13 := 0; fi13 < 13; fi13++ {
            if (false || ((-2.25 - 1e10) < 1e10)) {
                var v14 int16 = (^(v11))
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch ("hello" + "a") {
                    case "hello":
                        v6 = 127
                        v1 = append(v1, 314)
                        v16 := GenS1{F0: v6, F1: v3}
                        _ = v16
                    case "a":
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                    case "\n\t":
                        var v17 string = "hello"
                        _ = v17
                    default:
                        fmt.Println(fi13, fi13)
                    }
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) > 0 {
                            v1[0] = 879
                        }
                        v1 = append(v1, 2147483647)
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        v18 := GenS1{F0: (v6 - v2), F1: v3}
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                        _ = v18
                    }
                    v6 = int8(v4)
                    _ = fi15
                }
                switch (("" + "foo bar") + "foo bar") {
                case "a":
                    v1 = append(v1, 2)
                default:
                    switch 32767 {
                    case 866:
                        v1 = append(v1, 4294967295)
                    case 849:
                        if v3 {
                            v1 = append(v1, 2147483647)
                        } else {
                            v19 := GenS1{F0: (v2 * (v2 / ((v2) | 1))), F1: ((9223372036854775807 >= v5) || (((-2.25 * float64(1.5)) != (float64((-2.25 + float32((((((3.14159 * float64(-2.25)) - ((float64((-2.25 / 1.0)) * 1.5) + 1.5)) * (float64(float32((float64(3.14159) + -2.25))) / 0.1)) + ((3.14159 * ((((1e10 - (0.0 / (1.5 / 0.1))) + (0.1 - 1e10)) * -2.25) * (float64(((0.1 + 0.0) * (1.5 / 0.1))) * float64(0.1)))) / 3.14159)) / 1.0)))) - (0.1 / (1e10 / 3.14159)))) && v3))}
                            var v20 bool = (2147483647 < v10)
                            _ = v19
                            _ = v20
                        }
                    default:
                        var v21 int = -1
                        _ = v21
                    }
                }
                _ = v14
            }
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 int8 = 39
                if ((^(v5)) == v5) {
                    if ((682 < v11) && v3) {
                        v24 := GenS1{F0: v23, F1: v3}
                        switch fi13 {
                        case -2147483648:
                            var v25 float64 = float64(-2.25)
                            _ = v25
                        case 1:
                            if false {
                                var v26 float32 = -2.25
                                _ = v26
                            }
                        }
                        _ = v24
                    } else {
                        var v27 int = fi22
                        _ = v27
                    }
                } else {
                    if (492 > uint32(v4)) {
                        v28 := GenS2{F0: 1, F1: 128}
                        _ = v28
                    } else {
                        v4 = v4
                        _ = len(v1)
                    }
                    if false {
                        v4 = v4
                        v23 += v2
                        var v29 int64 = int64(v2)
                        var v30 int32 = 61
                        _ = v29
                        _ = v30
                    }
                    if true {
                        fmt.Println(v2)
                        switch fi22 {
                        case 0:
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                        default:
                            v3 = true
                        }
                        switch fi22 {
                        case 2147483647:
                            var v31 int64 = (-(v5))
                            _ = v31
                        default:
                            var v32 uint32 = 65535
                            _ = v32
                        }
                    }
                }
                switch (fi22 / ((790) | 1)) {
                case 2147483647:
                    if v3 {
                        v33 := GenS3{F0: false}
                        v11 -= (v11 % ((v11) | 1))
                        _ = v33
                    }
                }
                _ = fi22
                _ = v23
            }
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 int16 = (v11 | -32768)
                _ = fi34
                _ = v35
            }
            _ = fi13
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            v6 += v2
            var v37 uint8 = 168
            _ = fi36
            _ = v37
        }
    }
    var v38 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
