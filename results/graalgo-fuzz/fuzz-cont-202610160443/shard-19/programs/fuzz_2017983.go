package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 bool }
type GenS3 struct { F0 uint32; F1 int64; F2 uint8 }
func genFn0(p0 uint16, p1 uint) uint8 {
    var v6 float32 = (((float32(0.0) * 0.0) - 1e10) - 0.1)
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 += (v6 * (float32(0.0) - (0.0 * v6)))
        var v8 uint16 = 787
        if (-32768 == -542) {
            p1 = (p1 * (p1 >> ((70) & 63)))
            if (!true) {
                v9 := GenS1{F0: 65535}
                for fi10 := 0; fi10 < 6; fi10++ {
                    v8 = (^(v8))
                    p1 -= p1
                    _ = fi10
                }
                p1 = p1
                _ = v9
            } else {
                v8 += uint16(p1)
            }
            v6 = (v6 / v6)
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return 0
}

func main() {
    var v1 uint32 = 405
    v2 := []string{""}
    v3 := []string{("a" + ("\n\t" + ("" + "hello"))), ""}
    v4 := make(map[uint32]string)
    var v5 uint32 = (v1 & (v1 % ((v1) | 1)))
    if (false && false) {
        fmt.Println(v1, v1)
        fmt.Println(v5, v5)
        v5 += v5
    }
    if (0.0 <= ((0.0 - (3.14159 + (-2.25 * 0.0))) / 1.0)) {
        v12 := GenS1{F0: v5}
        for fi13 := 0; fi13 < 20; fi13++ {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            _ = fi13
        }
        v14 := GenS1{F0: (^(v1))}
        _ = v12
        _ = v14
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                switch ("hello" + (((("foo bar" + ("hello" + ("foo bar" + ("a" + (("" + "foo bar") + "a"))))) + "foo bar") + (("foo bar" + "") + "\n\t")) + "hello")) {
                case "foo bar":
                    var v17 float32 = float32(1.5)
                    _ = v17
                default:
                    v1 = 4294967295
                }
                v18 := GenS1{F0: 65535}
                if false {
                    v19 := v4[(v1 & (v1 << ((18446744073709551615) & 63)))]
                    if (1e10 == float64(1e10)) {
                        fmt.Println(v19, v1)
                        v20 := GenS1{F0: v1}
                        if (127 == (-(fi16))) {
                            v1 = (v5 << ((781) & 63))
                            fmt.Println(v19, v1)
                            v21 := GenS2{F0: 212, F1: (int64(v1) <= 2), F2: false}
                            _ = v21
                        } else {
                            fmt.Println(v5)
                            switch fi16 {
                            case 618:
                                if (1 > -1) {
                                    var v22 int32 = -2147483648
                                    var v23 int8 = int8(v5)
                                    var v24 float64 = ((float64(1.5) - 3.14159) + (1e10 / 1.0))
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                            }
                        }
                        v19 = (("\n\t" + v19) + "hello")
                        var v25 int8 = 127
                        if false {
                            v3 = append(v3, "hello")
                            if (((!(v19 != "\n\t")) && false) || false) {
                                var v26 uint64 = 4294967295
                                _ = v26
                            } else {
                                v25 = 100
                            }
                            var v27 int64 = 42
                            _ = v27
                        }
                        var v28 int32 = -2147483648
                        _ = v20
                        _ = v25
                        _ = v28
                    }
                    _ = v19
                } else {
                    var v29 uint64 = 157
                    var v30 float32 = 1.5
                    var v31 int32 = -765
                    _ = v29
                    _ = v30
                    _ = v31
                }
                var v32 float32 = 1.5
                _ = fi16
                _ = v18
                _ = v32
            }
            var v33 uint8 = 1
            _ = fi15
            _ = v33
        }
        v1 = 4294967295
    }
    if true {
        v1 -= 0
    }
    v1 -= 1
    var v34 float32 = -2.25
    var v35 int64 = 0
    v36 := GenS3{F0: 2, F1: (v35 - (v35 % (((-(v35))) | 1))), F2: 104}
    fmt.Println(v34, v34)
    var v37 int16 = 1000
    fmt.Println(v35)
    v34 = (float32(1.5) * (float32(0.1) * v34))
    fmt.Println(v1)
    for fi38 := 0; fi38 < 8; fi38++ {
        delete(v4, (v5 - v5))
        v2 = append(v2, ("" + "foo bar"))
        _ = fi38
    }
    var v39 uint32 = (v5 / ((2147483648) | 1))
    var v40 int8 = 100
    var v41 int = 239
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
