package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint32; F2 int16 }
type GenS6 struct { F0 int64; F1 int64; F2 int16 }
func genFn0(p0 uint64) uint16 {
    v4 := []int{32768, 1}
    v5 := GenS1{F0: 42, F1: "hello"}
    p0 = (p0 * (p0 ^ (^(p0))))
    v6 := GenS2{F0: 0, F1: uint8(p0)}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 272
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    var v3 uint = 255
    if false {
        if (uint16(v3) != 393) {
            _ = len(v2)
            for fi7 := 0; fi7 < 7; fi7++ {
                v3 = 255
                var v8 uint16 = 0
                _ = fi7
                _ = v8
            }
        } else {
            if (785 > 1) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 = (v3 - v3)
                if true {
                    switch ((("" + ("" + "foo bar")) + "\n\t") + (("hello" + "a") + "\n\t")) {
                    case "foo bar":
                        var v9 uint32 = 2147483647
                        _ = v9
                    case "":
                        v10 := GenS3{F0: 255, F1: -122, F2: 2147483648}
                        _ = v10
                    }
                    var v11 uint8 = 0
                    v3 = v3
                    _ = v11
                } else {
                    switch "hello" {
                    case "\n\t":
                        for fi12 := 0; fi12 < 4; fi12++ {
                            switch fi12 {
                            case 317:
                                if (((((-2.25 + (0.1 - (0.0 / float64(1.5)))) + 0.1) * 1e10) / 1.5) != float64(1.5)) {
                                    fmt.Println(fi12)
                                    fmt.Println(v3)
                                    delete(v1, 559)
                                } else {
                                    if (42 >= -1) {
                                        if true {
                                            v13 := GenS1{F0: 42, F1: "\n\t"}
                                            if (genFn0(8) >= 32767) {
                                                for fi14 := 0; fi14 < 1; fi14++ {
                                                    switch (-(fi14)) {
                                                    case 9223372036854775807:
                                                        v15 := GenS1{F0: 0, F1: "foo bar"}
                                                        _ = v15
                                                    }
                                                    _ = fi14
                                                }
                                                var v16 uint32 = 2147483647
                                                _ = v16
                                            } else {
                                                var v17 string = "a"
                                                _ = v17
                                            }
                                            v18 := GenS4{F0: 42}
                                            _ = v13
                                            _ = v18
                                        } else {
                                            v3 -= (^(v3))
                                            v3 += 18446744073709551615
                                        }
                                    }
                                    for fi19 := 0; fi19 < 3; fi19++ {
                                        fmt.Println(fi12)
                                        _ = fi19
                                    }
                                }
                            case -2147483648:
                                var v20 int8 = 42
                                _ = v20
                            default:
                                var v21 bool = false
                                _ = v21
                            }
                            if (!((-2.25 + 1.5) != 1.5)) {
                                var v22 int32 = -2147483648
                                v23 := GenS5{F0: 1, F1: 1, F2: 473}
                                _ = v22
                                _ = v23
                            } else {
                                var v24 int64 = 1
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    v3 = (^(v3))
                                    v26 := GenS5{F0: -116, F1: 65535, F2: 874}
                                    v3 = v3
                                    v2 = append(v2, "foo bar")
                                    v3 = (v3 + 255)
                                    _ = fi25
                                    _ = v26
                                }
                                _ = v24
                            }
                            v3 -= (v3 - (v3 % ((v3) | 1)))
                            _ = fi12
                        }
                    case "":
                        v27 := v1[0]
                        _ = v27
                    default:
                        var v28 float32 = (1e10 * 0.0)
                        _ = v28
                    }
                    var v29 int32 = -750
                    v30 := GenS6{F0: -629, F1: 42, F2: -32768}
                    _ = v29
                    _ = v30
                }
            }
            v3 += (^(v3))
        }
        v1[2147483647] = -9223372036854775808
    }
    if true {
        fmt.Println(v3)
        if (0 <= 804) {
            v3 = 18446744073709551615
            v3 -= (v3 << (((^(v3))) & 63))
            var v31 int8 = -77
            _ = v31
        }
        var v32 int32 = 42
        _ = v32
    }
    v3 -= v3
    var v33 int32 = 376
    fmt.Println(v33, v33)
    v3 = 11
    var v34 string = "a"
    v33 = 42
    var v35 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
