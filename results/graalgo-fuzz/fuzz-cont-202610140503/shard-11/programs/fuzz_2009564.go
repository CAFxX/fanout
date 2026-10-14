package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
type GenS2 struct { F0 int32; F1 int32; F2 int32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint8, p1 bool) bool {
    var v5 uint64 = 18446744073709551615
    fmt.Println(v5, p0, v5)
    _ = p0
    _ = p1
    _ = v5
    return (p1 || (!p1))
}
func genFn1(p0 int64, p1 int64) int8 {
    v6 := []int{-829, 9223372036854775807, 0, -241}
    switch 127 {
    case 127:
        for fi7 := 0; fi7 < 17; fi7++ {
            if true {
                if true {
                    if true {
                        if false {
                            v8 := GenS1{F0: 42, F1: -12, F2: fi7}
                            var v9 uint16 = 1
                            _ = v8
                            _ = v9
                        } else {
                            var v10 uint64 = 399
                            if true {
                                switch ("hello" + ("a" + "a")) {
                                case "\n\t":
                                    p0 = 9223372036854775807
                                case "":
                                    var v11 int8 = 42
                                    _ = v11
                                default:
                                    for fi12 := 0; fi12 < 16; fi12++ {
                                        _ = len(v6)
                                        fmt.Println(v10)
                                        fmt.Println(v10, fi7, v10)
                                        _ = fi12
                                    }
                                }
                                if false {
                                    p0 = (-(p1))
                                    for fi13 := 0; fi13 < 12; fi13++ {
                                        _ = len(v6)
                                        p1 = (p0 * -1)
                                        p0 += 9223372036854775807
                                        p0 = 9223372036854775807
                                        var v14 int8 = -7
                                        _ = fi13
                                        _ = v14
                                    }
                                    var v15 uint8 = 127
                                    _ = v15
                                } else {
                                    v16 := GenS2{F0: 2147483647, F1: 0, F2: 730}
                                    v17 := GenS3{F0: "a", F1: 127}
                                    v18 := GenS4{F0: 469, F1: v10, F2: 128}
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                                var v19 float64 = (1.5 - 1e10)
                                _ = v19
                            }
                            v6 = append(v6, -819)
                            _ = v10
                        }
                    } else {
                        v20 := GenS1{F0: 584, F1: -43, F2: 32768}
                        v21 := GenS4{F0: 0, F1: 880, F2: 128}
                        _ = v20
                        _ = v21
                    }
                } else {
                    if ((((p0 & (^(p1))) == (p0 >> ((18446744073709551615) & 63))) && (false && false)) || ((375 <= 2147483647) && true)) {
                        p1 += (p1 + (-(p1)))
                    } else {
                        v22 := GenS2{F0: 0, F1: -242, F2: 498}
                        var v23 int16 = 0
                        fmt.Println(p1)
                        _ = v22
                        _ = v23
                    }
                    switch ("\n\t" + "a") {
                    case "":
                        if true {
                            if false {
                                v24 := GenS5{F0: 2, F1: 545}
                                p1 = (p0 & p1)
                                if ("" < ((("hello" + "a") + "\n\t") + "hello")) {
                                    var v25 uint32 = 42
                                    fmt.Println(fi7, p0)
                                    _ = v25
                                }
                                _ = v24
                            } else {
                                p0 = p0
                            }
                        } else {
                            var v26 int64 = -1
                            var v27 int32 = -1
                            v28 := GenS4{F0: 101, F1: 2, F2: 239}
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    case "foo bar":
                        for fi29 := 0; fi29 < 5; fi29++ {
                            fmt.Println(fi7)
                            p1 -= (p1 - p1)
                            v6 = append(v6, (fi7 >> ((255) & 63)))
                            fmt.Println(p1)
                            _ = fi29
                        }
                    default:
                        p1 = -1
                    }
                }
                if true {
                    var v30 string = (("hello" + "hello") + ("foo bar" + ""))
                    _ = v30
                }
                for fi31 := 0; fi31 < 5; fi31++ {
                    var v32 bool = ((-2.25 * 0.1) < 0.0)
                    var v33 bool = (160 < 220)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            } else {
                p0 = p1
            }
            _ = fi7
        }
    case 914:
        p1 -= 0
    default:
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = v6
    return 82
}

func main() {
    v1 := []uint32{0, 0}
    var v2 uint64 = 4294967295
    v3 := []uint32{4294967295, 65535, 592, 1}
    var v4 bool = false
    var v34 uint = 1
    var v35 int32 = -297
    var v36 int8 = 103
    _ = len(v1)
    if len(v3) > 0 {
        v3[0] = uint32(v2)
    }
    v37 := GenS6{F0: v35}
    var v38 uint8 = 217
    var v39 uint32 = 2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
