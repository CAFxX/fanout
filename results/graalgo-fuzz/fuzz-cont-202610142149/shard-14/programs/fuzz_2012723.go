package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint; F2 int64 }

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 42
    v3 := []string{"", ""}
    v4 := make(map[string]uint32)
    var v5 int = -212
    v6 := make(map[string]int)
    v7 := GenS1{F0: 2147483647, F1: 65535, F2: -923}
    v8 := GenS1{F0: 42, F1: 983, F2: -78}
    fmt.Println(v2)
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 float64 = float64((0.0 / 1.0))
        _ = fi9
        _ = v10
    }
    if (255 == 2) {
        if ((^(v2)) > v2) {
            v11 := GenS1{F0: 575, F1: 266, F2: 1000}
            _ = v11
        } else {
            fmt.Println(v2, v5)
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            if false {
                if (((2147483647 == 2147483647) || ((3.14159 - (1.5 / (float32(1.5) / 1e10))) == 1e10)) || true) {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        if true {
                            v1 = "a"
                            if false {
                                v5 = -958
                                v2 += v2
                                fmt.Println(v5, v2)
                                delete(v6, "foo bar")
                                _ = len(v4)
                            }
                        }
                        var v14 float32 = (3.14159 * (-2.25 - ((((-2.25 + float32((float64((0.1 + (0.1 * 0.0))) + float64(1e10)))) - 1e10) - 1.5) / 1.5)))
                        switch (fi12 % (((fi12 << ((uint(v2)) & 63))) | 1)) {
                        case -387:
                            if (!true) {
                                v14 += 1e10
                                var v15 uint = 255
                                var v16 uint16 = (v2 - (^(v2)))
                                switch (v5 % ((fi13) | 1)) {
                                case -871:
                                    fmt.Println(v5)
                                case 1:
                                    v17 := GenS2{F0: 0, F1: 0, F2: (v15 / ((1) | 1))}
                                    _ = v17
                                case -806:
                                    var v18 uint64 = 1
                                    _ = v18
                                }
                                _ = v15
                                _ = v16
                            } else {
                                if (!(127 == 100)) {
                                    v5 -= v5
                                    if false {
                                        var v19 bool = ((4294967295 == -9223372036854775808) || false)
                                        v20 := GenS3{F0: uint(v2)}
                                        switch "\n\t" {
                                        case "foo bar":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case "a":
                                            _ = len(v6)
                                        default:
                                            v1 = v1
                                        }
                                        _ = v19
                                        _ = v20
                                    }
                                }
                                v1 = ("\n\t" + v1)
                            }
                        case 32768:
                            v5 = (fi13 << ((4294967295) & 63))
                        case 128:
                            _ = len(v4)
                        }
                        _ = fi13
                        _ = v14
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v21 := GenS2{F0: 57, F1: 128, F2: 152}
                    _ = v21
                }
            } else {
                v22 := GenS1{F0: 4294967295, F1: 42, F2: 32767}
                var v23 int16 = 42
                _ = v22
                _ = v23
            }
            v24 := GenS1{F0: 790, F1: 2147483647, F2: 0}
            _ = fi12
            _ = v24
        }
        if false {
            if (703 < 1000) {
                v4[v1] = 65535
                v1 = (v1 + (v1 + "foo bar"))
            }
        } else {
            fmt.Println(v2, v1, v1)
            v5 = 4294967295
        }
    } else {
        _ = len(v3)
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    v1 = ("" + "a")
    if (1 > (v2 - v2)) {
        if false {
            fmt.Println(v5, v1, v1)
            var v25 uint16 = v2
            _ = v25
        }
        v1 = (v1 + "foo bar")
        v26 := GenS4{F0: 1000, F1: 203, F2: 9223372036854775807}
        _ = v26
    } else {
        if true {
            if false {
                switch (v5 * (-(v5))) {
                case -308:
                    var v27 uint8 = 127
                    _ = v27
                case 2147483647:
                    if ((true && true) && false) {
                        switch v1 {
                        case "\n\t":
                            var v28 int32 = 2147483647
                            _ = v28
                        case "":
                            v5 = -2147483648
                        case "foo bar":
                            var v29 int32 = 1
                            _ = v29
                        }
                        v1 = "\n\t"
                    } else {
                        v30 := GenS1{F0: 283, F1: 844, F2: 42}
                        _ = v30
                    }
                case 32767:
                    fmt.Println(v1)
                }
            } else {
                v5 += 103
                var v31 bool = true
                _ = v31
            }
            var v32 int = v5
            _ = len(v3)
            _ = v32
        }
    }
    fmt.Println(v2, v2, v5)
    v33 := GenS3{F0: 801}
    for fi34 := 0; fi34 < 5; fi34++ {
        v1 = v1
        delete(v6, "hello")
        fmt.Println(fi34, fi34, v1)
        _ = fi34
    }
    v1 = ("\n\t" + ((v1 + v1) + (("\n\t" + ((v1 + v1) + (v1 + (v1 + v1)))) + v1)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        v1 = v1
        fmt.Println(v5, v2, v5)
        _ = fi35
    }
    var v36 int = v5
    var v37 int16 = 42
    var v38 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
