package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 uint; F1 int64; F2 int64 }
func genFn0(p0 int, p1 int64, p2 int) int64 {
    p0 = (-(p0))
    p1 = p1
    var v5 uint8 = 101
    var v6 uint64 = 18446744073709551615
    for fi7 := 0; fi7 < 16; fi7++ {
        if (p1 <= p1) {
            var v8 int = fi7
            if false {
                v9 := []int{}
                fmt.Println(p0, p2, v8)
                _ = v9
            }
            _ = v8
        } else {
            v10 := GenS1{F0: v5}
            var v11 uint32 = 2
            _ = v10
            _ = v11
        }
        if true {
            var v12 int8 = -128
            _ = v12
        } else {
            if true {
                p2 = p2
                fmt.Println(v5, v5)
                if false {
                    v13 := []int{int(v6), (fi7 % ((p2) | 1))}
                    for fi14 := 0; fi14 < 4; fi14++ {
                        _ = len(v13)
                        _ = fi14
                    }
                    var v15 string = ("a" + (("" + ("hello" + "\n\t")) + "foo bar"))
                    _ = v13
                    _ = v15
                }
            }
            if (219 != 874) {
                v16 := GenS2{F0: 65535, F1: fi7, F2: 0}
                _ = v16
            } else {
                v17 := []int{(p0 - 2147483647)}
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                if (p0 == int(p1)) {
                    if ((true || (1 < 432)) && (0.1 <= 0.0)) {
                        switch ("\n\t" + "") {
                        case "hello":
                            if ((-(p1)) > (^(p1))) {
                                v18 := GenS3{F0: -1, F1: (p1 % ((p1) | 1))}
                                switch "hello" {
                                case "":
                                    for fi19 := 0; fi19 < 8; fi19++ {
                                        v17 = append(v17, p2)
                                        v20 := GenS1{F0: v5}
                                        p2 += p0
                                        _ = fi19
                                        _ = v20
                                    }
                                }
                                if len(v17) > 0 {
                                    v17[0] = fi7
                                }
                                _ = v18
                            } else {
                                v21 := GenS4{F0: 2147483648, F1: (("hello" + (("a" + ("\n\t" + "foo bar")) + ((("hello" + "") + "") + ("a" + "foo bar")))) + "hello")}
                                var v22 int16 = -1
                                _ = v21
                                _ = v22
                            }
                        }
                        if (351 < 32767) {
                            switch (fi7 % ((p0) | 1)) {
                            case 32768:
                                fmt.Println(p0, p0, v5)
                            case -794:
                                var v23 uint16 = 394
                                _ = v23
                            default:
                                switch p0 {
                                case -596:
                                    p0 += fi7
                                case -9223372036854775808:
                                    if (true || (v6 == v6)) {
                                        v24 := GenS5{F0: 0}
                                        _ = v24
                                    }
                                }
                            }
                            v25 := GenS1{F0: v5}
                            var v26 int32 = 2147483647
                            _ = v25
                            _ = v26
                        }
                    }
                }
                _ = v17
            }
        }
        var v27 float32 = float32(3.14159)
        _ = fi7
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 + -347)
}
func genFn1(p0 int64, p1 int8) uint {
    var v28 int16 = 438
    var v29 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v28
    _ = v29
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = (1e10 * 1.5)
    var v3 int64 = 42
    var v4 string = "hello"
    if false {
        var v30 int64 = 810
        _ = v30
    }
    var v31 int32 = -803
    v32 := []int{int(v31), -2147483648}
    v33 := GenS1{F0: 224}
    delete(v1, 709)
    var v34 uint8 = 127
    if ((v3 / ((-795) | 1)) != (v3 << ((0) & 63))) {
        if (!true) {
            v35 := GenS6{F0: genFn1((v3 + (v3 | (v3 ^ 9223372036854775807))), -10)}
            _ = v35
        } else {
            var v36 float32 = 1.5
            var v37 int8 = 127
            for fi38 := 0; fi38 < 7; fi38++ {
                var v39 int = -795
                _ = fi38
                _ = v39
            }
            _ = v36
            _ = v37
        }
    } else {
        for fi40 := 0; fi40 < 16; fi40++ {
            v4 = "hello"
            switch v4 {
            case "\n\t":
                v34 = v34
            case "hello":
                v41 := GenS7{F0: 1, F1: 42}
                _ = v41
            default:
                v32 = append(v32, fi40)
            }
            _ = fi40
        }
        if ((!(1 == 1)) && ((v34 / (((v34 + (v34 - 13))) | 1)) != 0)) {
            if (0.1 > 1.5) {
                var v42 int = -9223372036854775808
                v43 := GenS8{F0: 65535, F1: v3, F2: genFn0(v42, (v3 | v3), 32767)}
                _ = len(v1)
                _ = v42
                _ = v43
            } else {
                v44 := GenS7{F0: 127, F1: -1}
                if (!(v31 != 0)) {
                    if len(v32) > 0 {
                        v32[0] = 191
                    }
                    v32 = append(v32, 4294967295)
                } else {
                    if len(v32) >= 2 {
                        v32 = v32[1:len(v32)]
                    }
                }
                v4 = ("a" + v4)
                _ = v44
            }
        } else {
            v34 = v34
            if len(v32) > 0 {
                v32[0] = -2147483648
            }
        }
    }
    var v45 uint16 = 42
    v46 := v1[306]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v45
    _ = v46
}
