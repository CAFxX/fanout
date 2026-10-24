package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{(("hello" + "a") + "a")}
    var v3 float64 = 3.14159
    var v4 bool = (-2147483648 == -2147483648)
    var v5 bool = v4
    v6 := []uint32{2147483647}
    fmt.Println(v4)
    var v7 int32 = -2147483648
    var v8 int32 = (v7 >> ((255) & 63))
    var v9 bool = v5
    for fi10 := 0; fi10 < 11; fi10++ {
        if v5 {
            if (111 > 0) {
                if (!(0 <= 2)) {
                    var v11 int = fi10
                    v11 += (fi10 - 191)
                    _ = v11
                } else {
                    var v12 uint = 255
                    if v4 {
                        v4 = (float64(float32(((v3 * 3.14159) + ((v3 - (-2.25 / v3)) * 3.14159)))) == (v3 / v3))
                        v5 = v4
                    }
                    _ = v12
                }
                var v13 string = (("\n\t" + "\n\t") + (("foo bar" + "\n\t") + "\n\t"))
                _ = v13
            } else {
                var v14 int8 = -128
                _ = v14
            }
            v15 := GenS1{F0: 4294967295, F1: 658}
            _ = v15
        } else {
            v7 = v8
        }
        var v16 uint32 = 42
        var v17 int64 = -9223372036854775808
        _ = fi10
        _ = v16
        _ = v17
    }
    switch "hello" {
    case "":
        v1[769] = -1
    case "hello":
        if v5 {
            if (v4 || (float32(v3) <= -2.25)) {
                for fi18 := 0; fi18 < 17; fi18++ {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v4 = ((("\n\t" + (("foo bar" + ("hello" + ("\n\t" + "a"))) + ("\n\t" + "foo bar"))) + ("" + ("a" + ""))) >= "foo bar")
                        _ = fi19
                    }
                    v7 = 2147483647
                    var v20 int16 = 88
                    _ = fi18
                    _ = v20
                }
                var v21 uint16 = 65535
                for fi22 := 0; fi22 < 13; fi22++ {
                    if true {
                        v23 := v1[705]
                        if v5 {
                            fmt.Println(v23, v23)
                        }
                        _ = v23
                    } else {
                        var v24 float64 = (float64(1e10) - (v3 + ((v3 + v3) / (v3 + (v3 * v3)))))
                        _ = v24
                    }
                    switch fi22 {
                    case 42:
                        switch ("a" + "a") {
                        case "hello":
                            v5 = true
                        }
                    case 1:
                        if (636 < 42) {
                            for fi25 := 0; fi25 < 19; fi25++ {
                                v5 = (-2.25 > -2.25)
                                _ = fi25
                            }
                        } else {
                            var v26 float32 = 0.1
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v9 = v9
                            _ = v26
                        }
                    case -7:
                        var v27 float64 = (float64(float32((0.1 + (3.14159 + -2.25)))) * v3)
                        _ = v27
                    default:
                        if (!(false && ((float32((float64(0.0) - v3)) + 1.5) <= -2.25))) {
                            for fi28 := 0; fi28 < 7; fi28++ {
                                if (1000 <= 32767) {
                                    switch "" {
                                    case "foo bar":
                                        v3 = v3
                                    }
                                    _ = len(v6)
                                }
                                v4 = (("\n\t" + "\n\t") <= "hello")
                                var v29 uint = 1
                                _ = fi28
                                _ = v29
                            }
                        }
                    }
                    for fi30 := 0; fi30 < 1; fi30++ {
                        if (((42 != 1000) || v4) && false) {
                            if len(v6) > 0 {
                                v6[0] = 2147483647
                            }
                        }
                        var v31 uint16 = 467
                        var v32 int64 = 175
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    _ = fi22
                }
                _ = v21
            }
        }
    case "a":
        fmt.Println(v5)
    }
    v3 -= (-2.25 * (v3 - 1.5))
    for fi33 := 0; fi33 < 16; fi33++ {
        v8 = v7
        var v34 int = (^(fi33))
        _ = fi33
        _ = v34
    }
    v8 -= (v7 + (v7 | -2147483648))
    v8 -= (-(v8))
    v8 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
