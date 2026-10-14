package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint32; F2 int64 }
func genFn0(p0 int32) int {
    if (-9223372036854775808 != -1) {
        p0 = p0
        if (true || (p0 < -762)) {
            for fi7 := 0; fi7 < 5; fi7++ {
                if (-1 > -38) {
                    v8 := GenS1{F0: (fi7 | (fi7 % ((fi7) | 1))), F1: 1, F2: ""}
                    _ = v8
                } else {
                    fmt.Println(fi7, fi7)
                    p0 = p0
                }
                _ = fi7
            }
            var v9 int8 = 42
            _ = v9
        } else {
            if (-1 != 8) {
                v10 := GenS1{F0: 2147483647, F1: 554, F2: ("a" + "\n\t")}
                var v11 int8 = 27
                if (1 < 32767) {
                    if ((0.0 + float32(1.5)) != ((0.1 * (-2.25 - (-2.25 - ((-2.25 * (1e10 + 1e10)) / float32(((0.1 * 0.1) / 3.14159)))))) / 1.5)) {
                        v12 := GenS2{F0: 42}
                        if true {
                            v13 := GenS2{F0: 0}
                            var v14 float64 = -2.25
                            for fi15 := 0; fi15 < 19; fi15++ {
                                var v16 int8 = (^(v11))
                                _ = fi15
                                _ = v16
                            }
                            _ = v13
                            _ = v14
                        } else {
                            for fi17 := 0; fi17 < 18; fi17++ {
                                switch -387 {
                                case -363:
                                    v18 := []int{(fi17 ^ 9223372036854775807), fi17}
                                    _ = v18
                                case 32768:
                                    switch "a" {
                                    case "\n\t":
                                        p0 -= p0
                                    }
                                }
                                fmt.Println(v11)
                                _ = fi17
                            }
                        }
                        _ = v12
                    } else {
                        var v19 int8 = v11
                        v19 -= v11
                        v19 -= (v19 / ((v11) | 1))
                        _ = v19
                    }
                } else {
                    v20 := GenS2{F0: 158}
                    if true {
                        p0 -= (^(p0))
                        switch "" {
                        case "\n\t":
                            var v21 int8 = (v11 * (v11 >> ((676) & 63)))
                            _ = v21
                        case "":
                            if (42 < 299) {
                                p0 += -2147483648
                                var v22 string = "hello"
                                _ = v22
                            }
                        }
                        p0 = p0
                    } else {
                        for fi23 := 0; fi23 < 10; fi23++ {
                            v24 := GenS2{F0: 2}
                            _ = fi23
                            _ = v24
                        }
                    }
                    _ = v20
                }
                _ = v10
                _ = v11
            } else {
                v25 := GenS3{F0: -1, F1: 4294967295, F2: 237}
                fmt.Println(p0)
                _ = v25
            }
            p0 = -142
        }
    } else {
        p0 += (p0 + (-(p0)))
        p0 = 2147483647
    }
    var v26 int16 = -1
    _ = p0
    _ = v26
    return 4294967295
}
func genFn1() int32 {
    var v27 int8 = 100
    var v28 int8 = v27
    var v29 uint8 = 0
    v30 := GenS4{F0: -9223372036854775808}
    _ = v27
    _ = v28
    _ = v29
    _ = v30
    return 0
}

func main() {
    var v1 int16 = 32767
    var v2 uint64 = 42
    v3 := make(map[uint32]bool)
    var v4 int = int(v2)
    v5 := []int{(^(v4)), v4, (v4 % ((v4) | 1)), (v4 ^ (v4 >> ((295) & 63)))}
    v6 := make(map[int]uint32)
    v4 -= v4
    var v31 uint = 0
    var v32 bool = true
    v4 -= v4
    delete(v3, 2147483648)
    switch ("hello" + "foo bar") {
    case "foo bar":
        v33 := GenS2{F0: uint8(v2)}
        _ = v33
    case "a":
        v34 := GenS5{F0: v32, F1: 4294967295, F2: -241}
        _ = v34
    case "\n\t":
        fmt.Println(v31)
    default:
        var v35 int32 = 1
        _ = v35
    }
    if (2 <= 42) {
        for fi36 := 0; fi36 < 15; fi36++ {
            v1 = int16(v31)
            _ = fi36
        }
    } else {
        v1 = (v1 | -104)
        v4 += (v4 << ((4294967295) & 63))
        fmt.Println(v2)
    }
    var v37 uint64 = (v2 / ((v2) | 1))
    fmt.Println(v31, v4)
    var v38 int = 2147483647
    fmt.Println(v31, v31, v32)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, -1)
    var v39 int = (^(v38))
    for fi40 := 0; fi40 < 13; fi40++ {
        _ = len(v6)
        v41 := GenS5{F0: false, F1: 735, F2: 0}
        var v42 int = (fi40 << ((18446744073709551615) & 63))
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
