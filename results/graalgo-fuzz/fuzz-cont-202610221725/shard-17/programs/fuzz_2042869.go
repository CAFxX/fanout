package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64 }
func genFn0(p0 int) int32 {
    p0 -= p0
    if false {
        p0 -= 32768
    }
    p0 = (p0 / (((p0 << ((0) & 63))) | 1))
    _ = p0
    return -1
}
func genFn1() uint64 {
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 13; fi8++ {
        if (2147483647 >= 65535) {
            v7 = (v7 + v7)
        }
        if (true && false) {
            if (!false) {
                v7 = v7
                v7 = v7
            } else {
                v9 := []int{(fi8 + fi8), (fi8 - fi8), fi8, (fi8 / ((fi8) | 1))}
                _ = v9
            }
            fmt.Println(fi8, v7)
        }
        _ = fi8
    }
    v7 = v7
    switch "foo bar" {
    case "a":
        v10 := GenS1{F0: 1, F1: 128, F2: 32767}
        _ = v10
    case "\n\t":
        switch (v7 + v7) {
        case "hello":
            var v11 int8 = 100
            _ = v11
        case "a":
            var v12 int16 = 1000
            _ = v12
        default:
            v13 := []int{32767, 32767, 2147483647}
            _ = v13
        }
    case "hello":
        v14 := []int{32768, 4294967295, -9223372036854775808, 9223372036854775807}
        _ = v14
    default:
        for fi15 := 0; fi15 < 19; fi15++ {
            if (0 > 255) {
                switch (v7 + v7) {
                case "":
                    switch v7 {
                    case "foo bar":
                        v7 = "\n\t"
                    }
                default:
                    v16 := []int{0}
                    _ = v16
                }
            }
            v17 := []int{-15, (fi15 ^ 1000), -418}
            fmt.Println(v7, v7)
            _ = fi15
            _ = v17
        }
    }
    var v18 uint64 = 1
    _ = v18
    _ = v7
    return v18
}

func main() {
    v1 := []uint32{4294967295, 42, 2147483647, 1}
    v2 := make(map[int]int)
    var v3 uint8 = 2
    var v4 float64 = (-2.25 + 0.1)
    var v5 string = "\n\t"
    var v6 int8 = -128
    var v19 string = "a"
    var v20 uint8 = (v3 & 148)
    fmt.Println(v6)
    var v21 uint16 = 741
    switch "\n\t" {
    case "foo bar":
        v22 := GenS1{F0: 2, F1: 4294967295, F2: (^(v21))}
        _ = v22
    case "hello":
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint32 = 4294967295
            v25 := v2[(fi23 & (fi23 * int(v20)))]
            var v26 bool = ((v3 << ((4294967295) & 63)) < uint8(v21))
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    case "a":
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
    default:
        var v27 float32 = ((0.1 / float32((3.14159 / v4))) / float32(1e10))
        _ = v27
    }
    if (true && (42 != int32(v21))) {
        v28 := GenS2{F0: v20, F1: 846}
        for fi29 := 0; fi29 < 13; fi29++ {
            switch (-(fi29)) {
            case 4294967295:
                v6 -= v6
            case -9223372036854775808:
                if (2147483647 >= 42) {
                    switch v19 {
                    case "hello":
                        var v30 float64 = v4
                        _ = v30
                    case "\n\t":
                        v19 = ((v5 + (v19 + (v19 + v19))) + v5)
                    default:
                        v4 -= v4
                    }
                    var v31 uint64 = genFn1()
                    v32 := GenS1{F0: 788, F1: (fi29 / ((4294967295) | 1)), F2: (v21 + v21)}
                    _ = v31
                    _ = v32
                }
            case 9223372036854775807:
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
            default:
                if false {
                    for fi33 := 0; fi33 < 18; fi33++ {
                        if false {
                            v34 := GenS2{F0: 127, F1: 4294967295}
                            v35 := GenS3{F0: 42}
                            _ = v34
                            _ = v35
                        } else {
                            switch v19 {
                            case "":
                                v20 -= 166
                            default:
                                v36 := GenS4{F0: 4294967295}
                                _ = v36
                            }
                            var v37 int16 = 42
                            var v38 string = (v5 + v19)
                            _ = v37
                            _ = v38
                        }
                        _ = fi33
                    }
                    fmt.Println(v5, v6, v20)
                    for fi39 := 0; fi39 < 11; fi39++ {
                        v3 += 127
                        v1 = append(v1, 42)
                        _ = fi39
                    }
                } else {
                    var v40 int = fi29
                    if true {
                        v21 += (v21 ^ v21)
                    } else {
                        if true {
                            if (0.0 != 3.14159) {
                                delete(v2, (v40 - v40))
                                _ = len(v1)
                            }
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            for fi41 := 0; fi41 < 2; fi41++ {
                                _ = len(v1)
                                _ = len(v1)
                                v6 = v6
                                _ = fi41
                            }
                        }
                        var v42 uint32 = 208
                        _ = v42
                    }
                    _ = v40
                }
            }
            for fi43 := 0; fi43 < 11; fi43++ {
                if (993 < int16(v20)) {
                    if (("foo bar" + v5) != v5) {
                        switch v19 {
                        case "":
                            v44 := GenS5{F0: -9223372036854775808}
                            _ = v44
                        default:
                            var v45 int8 = (v6 & int8(v3))
                            _ = v45
                        }
                    } else {
                        if false {
                            switch v5 {
                            case "foo bar":
                                v21 += (v21 % ((v21) | 1))
                            case "\n\t":
                                v46 := GenS5{F0: -9223372036854775808}
                                _ = v46
                            }
                        }
                    }
                    v47 := GenS1{F0: 2, F1: (fi29 - 128), F2: 803}
                    _ = v47
                } else {
                    v21 = v21
                }
                v5 = v5
                if false {
                    _ = len(v2)
                }
                _ = fi43
            }
            _ = fi29
        }
        v20 = v20
        _ = v28
    }
    fmt.Println(v19)
    v19 = (v19 + (v5 + "foo bar"))
    fmt.Println(v20)
    var v48 uint8 = (v20 | 255)
    fmt.Println(v3, v6, v21)
    var v49 int32 = 42
    var v50 bool = (1.5 == 3.14159)
    fmt.Println(v20)
    if false {
        fmt.Println(v49)
        fmt.Println(v48)
        _ = len(v1)
    }
    var v51 float32 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
