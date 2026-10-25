package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    v3 := []string{"\n\t", "\n\t", "\n\t", "hello"}
    v4 := []uint32{1, 111, 0, 65535}
    var v5 uint64 = 18446744073709551615
    _ = len(v4)
    v5 -= (v5 << ((4294967295) & 63))
    if (0 <= 42) {
        fmt.Println(v5)
        fmt.Println(v5)
    }
    v5 += (^(v5))
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 string = ("\n\t" + "a")
            var v9 string = "hello"
            var v10 string = ((v8 + v8) + v8)
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        if false {
            for fi11 := 0; fi11 < 13; fi11++ {
                fmt.Println(fi6)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(fi11)
                _ = fi11
            }
            var v12 uint8 = 32
            _ = v12
        } else {
            if ((false && (516 < 1)) && false) {
                v5 -= 2
                for fi13 := 0; fi13 < 14; fi13++ {
                    var v14 uint = 853
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + "a")
                    }
                    _ = fi13
                    _ = v14
                }
            } else {
                v5 += 18446744073709551615
                v2 = append(v2, 42)
                v5 = (v5 << ((0) & 63))
            }
        }
        if true {
            if (510 != 219) {
                v5 += (^(v5))
                fmt.Println(fi6, v5)
                if (!false) {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v5 = v5
                        _ = fi15
                    }
                } else {
                    v5 -= v5
                }
            } else {
                var v16 uint = 18446744073709551615
                var v17 int32 = 519
                for fi18 := 0; fi18 < 3; fi18++ {
                    v1[("" + "a")] = 455
                    if true {
                        if (-2.25 <= (float64(float32(0.0)) - 0.0)) {
                            _ = len(v1)
                            fmt.Println(fi6)
                            _ = len(v1)
                        }
                        var v19 uint8 = 123
                        v19 -= v19
                        _ = v19
                    } else {
                        var v20 int16 = 32767
                        var v21 uint16 = 2
                        v22 := GenS1{F0: fi6, F1: (v20 + (v20 << (((^(v16))) & 63)))}
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    var v23 uint8 = 23
                    _ = fi18
                    _ = v23
                }
                _ = v16
                _ = v17
            }
            if (-9223372036854775808 == 1) {
                _ = len(v1)
                if true {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        if ("foo bar" == ("\n\t" + (("\n\t" + ("hello" + ((("a" + "hello") + (((("hello" + "") + ("hello" + "a")) + (("foo bar" + ("\n\t" + "a")) + "hello")) + "")) + "hello"))) + "\n\t"))) {
                            v5 += (v5 | (v5 >> ((4294967295) & 63)))
                            fmt.Println(fi6)
                        }
                        _ = fi24
                    }
                    var v25 int64 = -33
                    _ = v25
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v5 = (v5 - (v5 - v5))
                }
            } else {
                if false {
                    var v26 int8 = 127
                    _ = v26
                } else {
                    if true {
                        var v27 uint16 = uint16(v5)
                        _ = v27
                    } else {
                        fmt.Println(fi6, v5)
                        switch (fi6 % ((fi6) | 1)) {
                        case 1:
                            var v28 int16 = 32767
                            _ = v28
                        default:
                            for fi29 := 0; fi29 < 9; fi29++ {
                                if false {
                                    if true {
                                        var v30 bool = false
                                        _ = v30
                                    }
                                    v5 = v5
                                    v5 = 42
                                }
                                v31 := GenS2{F0: -128}
                                _ = fi29
                                _ = v31
                            }
                        }
                        for fi32 := 0; fi32 < 17; fi32++ {
                            fmt.Println(fi32)
                            _ = fi32
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                for fi33 := 0; fi33 < 7; fi33++ {
                    v34 := GenS3{F0: 2147483647}
                    v5 -= 4294967295
                    v5 += 0
                    _ = fi33
                    _ = v34
                }
            }
        } else {
            v35 := GenS4{F0: 127, F1: (v5 + (v5 / ((v5) | 1)))}
            var v36 float64 = (3.14159 * 0.0)
            _ = v35
            _ = v36
        }
        _ = fi6
    }
    var v37 int8 = -128
    var v38 float64 = -2.25
    if (-817 <= -2147483648) {
        var v39 float32 = 3.14159
        v5 += (v5 ^ 2)
        if false {
            v1["\n\t"] = 42
        }
        _ = v39
    } else {
        var v40 uint8 = 55
        v40 = (v40 & 108)
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
