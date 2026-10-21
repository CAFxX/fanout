package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 int8; F1 bool; F2 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int8; F2 int16 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int }

func main() {
    var v1 bool = false
    var v2 uint64 = 0
    var v3 float32 = (((3.14159 - (0.1 / -2.25)) + 3.14159) / 1.0)
    var v4 uint16 = 0
    v5 := []uint32{}
    fmt.Println(v4)
    _ = len(v5)
    v5 = append(v5, 446)
    v6 := GenS1{F0: 343}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int32 = -1
        switch 2147483647 {
        case -9223372036854775808:
            v9 := GenS1{F0: 20}
            _ = v9
        case -740:
            switch (("hello" + ("a" + "")) + ("a" + "a")) {
            case "a":
                v10 := GenS2{F0: -1, F1: (!(v1 || (!v1))), F2: v4}
                _ = v10
            case "foo bar":
                v11 := GenS3{F0: 127, F1: ((uint(v4) <= uint(v2)) || v1), F2: 127}
                _ = v11
            default:
                v12 := GenS4{F0: (^(v2))}
                _ = v12
            }
        case 128:
            for fi13 := 0; fi13 < 14; fi13++ {
                v4 = (^(v4))
                if true {
                    if (!v1) {
                        if len(v5) > 0 {
                            v5[0] = 2147483647
                        }
                        v14 := GenS2{F0: 282, F1: (("foo bar" + "a") >= ""), F2: (v4 ^ v4)}
                        if len(v5) > 0 {
                            v5[0] = 65535
                        }
                        _ = v14
                    } else {
                        fmt.Println(v1, v8, v1)
                    }
                }
                _ = fi13
            }
        }
        if len(v5) > 0 {
            v5[0] = 32
        }
        _ = fi7
        _ = v8
    }
    var v15 uint8 = 255
    v3 = (3.14159 - 1.5)
    switch ((("a" + ("" + "\n\t")) + (("foo bar" + "\n\t") + "hello")) + "foo bar") {
    case "a":
        var v16 bool = ((float32((1e10 + (0.1 * (3.14159 + 1e10)))) * (((v3 * v3) / v3) / (v3 / 1.5))) != v3)
        _ = v16
    case "foo bar":
        v4 += v4
    }
    if (v3 >= v3) {
        if v1 {
            v3 -= v3
            v1 = (2147483647 == 2147483648)
        } else {
            var v17 bool = false
            switch 42 {
            case 9223372036854775807:
                v3 += ((float32((float64(3.14159) * 0.1)) + float32(-2.25)) - 3.14159)
            case 1:
                v4 -= v4
            default:
                switch 1 {
                case 9223372036854775807:
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                case 0:
                    _ = len(v5)
                default:
                    switch 32768 {
                    case 4294967295:
                        if (!(!(2 >= 1))) {
                            v18 := GenS5{F0: 2, F1: 42, F2: 0}
                            for fi19 := 0; fi19 < 20; fi19++ {
                                _ = len(v5)
                                _ = fi19
                            }
                            _ = v18
                        }
                    case -453:
                        v17 = (-2.25 < v3)
                    case 127:
                        if v1 {
                            v20 := GenS6{F0: 192, F1: -32768}
                            v1 = (-181 <= -1)
                            _ = v20
                        } else {
                            v17 = false
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                switch -465 {
                                case 32768:
                                    v21 := GenS7{F0: 128}
                                    _ = v21
                                case 2147483647:
                                    fmt.Println(v3, v3, v15)
                                case -65:
                                    v2 -= v2
                                }
                            default:
                                switch "" {
                                case "":
                                    v22 := GenS6{F0: (v15 >> ((867) & 63)), F1: 1000}
                                    _ = v22
                                case "foo bar":
                                    v1 = (v15 == 23)
                                case "\n\t":
                                    for fi23 := 0; fi23 < 16; fi23++ {
                                        fmt.Println(v4)
                                        for fi24 := 0; fi24 < 3; fi24++ {
                                            v25 := GenS4{F0: v2}
                                            v17 = (42 < -9223372036854775808)
                                            var v26 int16 = -823
                                            v15 = v15
                                            _ = fi24
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = fi23
                                    }
                                }
                            }
                        }
                    }
                }
            }
            v2 = (^(v2))
            _ = v17
        }
        v27 := GenS8{F0: -9223372036854775808}
        _ = v27
    }
    var v28 int32 = -98
    v29 := GenS4{F0: uint64(v4)}
    var v30 uint = 65535
    if (!true) {
        v5 = append(v5, 2)
        var v31 int16 = int16(v28)
        _ = v31
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        v33 := GenS5{F0: v30, F1: 1, F2: 1000}
        _ = fi32
        _ = v33
    }
    var v34 uint64 = v2
    v35 := GenS5{F0: v30, F1: 46, F2: 32767}
    var v36 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
