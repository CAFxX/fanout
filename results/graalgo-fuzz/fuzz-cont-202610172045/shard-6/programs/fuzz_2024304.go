package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int, p1 int) int16 {
    var v4 uint8 = 127
    v5 := GenS1{F0: 42, F1: 1, F2: 2147483647}
    v4 = v4
    var v6 uint16 = 65535
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1000
}
func genFn1(p0 int32, p1 int32, p2 int16) uint8 {
    var v7 string = ""
    var v8 float32 = -2.25
    var v9 uint16 = 65535
    p0 -= (p1 % (((-(p0))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []string{("\n\t" + "\n\t"), "hello", (("a" + (("hello" + "foo bar") + "a")) + "\n\t")}
    var v2 float64 = 3.14159
    v3 := make(map[int]int)
    var v10 int16 = 42
    v10 = (v10 / ((42) | 1))
    v2 += (v2 / v2)
    var v11 int64 = 9223372036854775807
    v3[-75] = 4294967295
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if ((v10 >> ((1) & 63)) != v10) {
        v12 := v3[-767]
        v10 = 32767
        _ = v12
    } else {
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 int = (^(fi13))
            switch (("hello" + "foo bar") + ("\n\t" + ("hello" + ("foo bar" + (("hello" + "hello") + ("a" + (("hello" + "") + (("\n\t" + "\n\t") + "\n\t")))))))) {
            case "\n\t":
                switch (fi13 << ((4294967295) & 63)) {
                case 562:
                    v10 = v10
                default:
                    for fi15 := 0; fi15 < 19; fi15++ {
                        if false {
                            fmt.Println(fi13, v14)
                            if true {
                                v16 := v3[-9223372036854775808]
                                v17 := v3[v14]
                                _ = v16
                                _ = v17
                            }
                            switch "" {
                            case "foo bar":
                                var v18 uint16 = 255
                                _ = v18
                            case "\n\t":
                                if false {
                                    v19 := GenS1{F0: 2, F1: v10, F2: 0}
                                    _ = v19
                                } else {
                                    if false {
                                        switch "hello" {
                                        case "a":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        default:
                                            fmt.Println(v11, v14)
                                        }
                                    }
                                    var v20 int32 = 42
                                    switch (fi15 ^ (fi15 - 0)) {
                                    case 128:
                                        v3[-547] = (fi15 ^ v14)
                                    case 973:
                                        if false {
                                            v21 := GenS2{F0: (v20 & -1), F1: 2}
                                            if false {
                                                var v22 uint8 = 0
                                                var v23 string = (("a" + "foo bar") + "foo bar")
                                                fmt.Println(v22, v20, fi13)
                                                _ = v22
                                                _ = v23
                                            }
                                            _ = v21
                                        }
                                    case 9223372036854775807:
                                        switch fi13 {
                                        case 9223372036854775807:
                                            fmt.Println(v11)
                                            v3[9223372036854775807] = 20
                                            var v24 float32 = 3.14159
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                            _ = v24
                                        case 32768:
                                            var v25 uint32 = 756
                                            v26 := GenS1{F0: 4294967295, F1: 384, F2: 42}
                                            v27 := GenS1{F0: 499, F1: (-(v10)), F2: 0}
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        case 4294967295:
                                            var v28 int32 = (v20 / ((v20) | 1))
                                            _ = v28
                                        default:
                                            fmt.Println(v14)
                                        }
                                    default:
                                        if (!true) {
                                            switch (v14 ^ (fi15 + (^(fi15)))) {
                                            case 0:
                                                v10 = (v10 % ((v10) | 1))
                                            case 9223372036854775807:
                                                var v29 int16 = v10
                                                _ = v29
                                            case 1:
                                                v2 = float64(((0.0 - (0.1 * -2.25)) - 1.5))
                                            }
                                            v2 -= v2
                                            v10 -= v10
                                            switch ("foo bar" + "foo bar") {
                                            case "hello":
                                                v30 := GenS1{F0: 0, F1: 1, F2: (-(v20))}
                                                _ = v30
                                            case "foo bar":
                                                v10 = (v10 & v10)
                                            case "\n\t":
                                                v10 += (v10 | (v10 % ((v10) | 1)))
                                            }
                                        } else {
                                            v31 := GenS3{F0: (v10 << ((42) & 63))}
                                            var v32 string = ("foo bar" + (((("hello" + "a") + "") + "a") + "foo bar"))
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    _ = v20
                                }
                            default:
                                var v33 uint = 450
                                _ = v33
                            }
                            v1 = append(v1, "a")
                        }
                        _ = fi15
                    }
                }
            default:
                if false {
                    v2 -= v2
                } else {
                    var v34 uint = 1
                    _ = v34
                }
            }
            _ = fi13
            _ = v14
        }
        var v35 int8 = -1
        _ = v35
    }
    v36 := GenS3{F0: -32768}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v36
}
