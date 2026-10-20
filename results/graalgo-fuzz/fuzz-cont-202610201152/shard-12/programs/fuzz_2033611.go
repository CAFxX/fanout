package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS4 struct { F0 bool; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8 }
func genFn0() int8 {
    var v6 float64 = 0.1
    if (-2147483648 > 42) {
        var v7 uint16 = 2
        v6 = 0.1
        var v8 bool = (0 >= 95)
        _ = v7
        _ = v8
    } else {
        v6 += float64(1.5)
        v9 := GenS1{F0: 0, F1: 1}
        switch "\n\t" {
        case "\n\t":
            v6 = v6
        case "":
            v10 := []int{-802, 0, 2147483647}
            _ = v10
        }
        _ = v9
    }
    _ = v6
    return -128
}
func genFn1(p0 int64) string {
    p0 = 42
    switch 32768 {
    case 2147483647:
        v11 := []int{325, 32768}
        _ = v11
    case 0:
        v12 := GenS2{F0: 72, F1: 9223372036854775807}
        _ = v12
    }
    _ = p0
    return ("\n\t" + "\n\t")
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = (0.1 + (3.14159 + (float32(((float64(0.1) / 3.14159) + ((3.14159 + (0.0 / (((0.0 - 3.14159) / 1.0) / 0.1))) * 3.14159))) / (float32(1.5) * -2.25))))
    var v3 uint8 = 127
    var v4 string = "foo bar"
    var v5 bool = ((v3 + (v3 * 190)) > v3)
    v3 -= (v3 >> ((255) & 63))
    v13 := GenS3{F0: 2, F1: -1, F2: 127}
    v1["a"] = v4
    v14 := []int{956, 127, 128, 128}
    v15 := GenS3{F0: 65535, F1: -795, F2: 1}
    v5 = v5
    switch v4 {
    case "\n\t":
        switch "" {
        case "hello":
            var v16 uint16 = 255
            _ = v16
        case "\n\t":
            if false {
                v5 = (-280 < 1000)
            } else {
                var v17 int = -338
                if v5 {
                    if v5 {
                        v3 = v3
                    }
                }
                _ = v17
            }
        }
    }
    v2 = (((v2 * (v2 / v2)) * v2) - v2)
    switch (v4 + (((v4 + "a") + v4) + "foo bar")) {
    case "a":
        var v18 uint16 = 255
        _ = v18
    case "":
        var v19 bool = v5
        _ = v19
    case "\n\t":
        fmt.Println(v3)
    default:
        if true {
            v5 = v5
            if v5 {
                v4 = v4
                if (!((!v5) && false)) {
                    if v5 {
                        v1[v4] = "\n\t"
                        delete(v1, v4)
                        fmt.Println(v5)
                    } else {
                        v1[v4] = ("hello" + "")
                        for fi20 := 0; fi20 < 15; fi20++ {
                            for fi21 := 0; fi21 < 6; fi21++ {
                                if v5 {
                                    fmt.Println(fi20, fi21, v4)
                                    if (!v5) {
                                        var v22 bool = v5
                                        v23 := GenS4{F0: (v22 || false), F1: -396, F2: 255}
                                        if v22 {
                                            v3 = (v3 - 255)
                                            if (!((!v22) || ((0 <= 455) || v22))) {
                                                _ = len(v14)
                                            }
                                            if len(v14) >= 2 {
                                                v14 = v14[1:len(v14)]
                                            }
                                        }
                                        v24 := GenS5{F0: 747}
                                        _ = len(v14)
                                        _ = len(v1)
                                        var v25 int8 = 64
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    v5 = (v5 && v5)
                                } else {
                                    var v26 int8 = 0
                                    _ = v26
                                }
                                v27 := v1["foo bar"]
                                _ = fi21
                                _ = v27
                            }
                            var v28 uint = 58
                            if v5 {
                                var v29 uint8 = v3
                                _ = v29
                            } else {
                                v3 = v3
                                switch v4 {
                                case "foo bar":
                                    fmt.Println(v2)
                                case "":
                                    var v30 float64 = -2.25
                                    _ = v30
                                }
                                v31 := GenS6{F0: v3}
                                _ = v31
                            }
                            _ = fi20
                            _ = v28
                        }
                    }
                    if len(v14) > 0 {
                        v14[0] = 4294967295
                    }
                }
            }
            fmt.Println(v2, v2, v4)
        } else {
            v32 := GenS6{F0: v3}
            _ = v32
        }
    }
    if v5 {
        fmt.Println(v3, v2, v2)
        for fi33 := 0; fi33 < 8; fi33++ {
            v5 = v5
            _ = fi33
        }
        v5 = false
    }
    var v34 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
