package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 bool; F1 int; F2 uint }
func genFn0(p0 uint, p1 int16) int16 {
    p0 = 18446744073709551615
    v5 := []int{}
    _ = len(v5)
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1(p0 int64, p1 int16) uint32 {
    switch "a" {
    case "foo bar":
        var v6 uint32 = 774
        _ = v6
    case "hello":
        v7 := []int{-275}
        _ = v7
    case "\n\t":
        if (100 < -1) {
            switch "foo bar" {
            case "a":
                fmt.Println(p0, p1)
            default:
                p0 -= p0
            }
        } else {
            for fi8 := 0; fi8 < 1; fi8++ {
                switch (fi8 / ((fi8) | 1)) {
                case -2147483648:
                    p1 += (-(p1))
                case 127:
                    p0 = p0
                case 65:
                    fmt.Println(p1, p0)
                default:
                    v9 := []int{(fi8 % (((-(fi8))) | 1)), fi8, (fi8 * (^(fi8)))}
                    _ = v9
                }
                v10 := GenS1{F0: 255, F1: 414}
                _ = fi8
                _ = v10
            }
            var v11 bool = false
            p0 = -9223372036854775808
            _ = v11
        }
    }
    p1 += genFn0(4294967295, (p1 << ((1) & 63)))
    var v12 uint = 1
    _ = p0
    _ = p1
    _ = v12
    return 2
}

func main() {
    v1 := []string{"hello", ("" + ""), ("" + ("a" + "hello"))}
    v2 := make(map[uint32]int)
    var v3 int8 = 42
    var v4 float64 = (1e10 / (-2.25 / 1.5))
    for fi13 := 0; fi13 < 8; fi13++ {
        switch "\n\t" {
        case "foo bar":
            delete(v2, 2)
        case "":
            v4 += 0.0
        case "\n\t":
            for fi14 := 0; fi14 < 2; fi14++ {
                if true {
                    var v15 int = -675
                    switch v15 {
                    case -2147483648:
                        v3 = v3
                    default:
                        v15 = (v15 & 860)
                    }
                    var v16 int32 = int32(v3)
                    _ = v15
                    _ = v16
                } else {
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v3, v4, v3)
                    case "":
                        v17 := GenS1{F0: 81, F1: 0}
                        _ = v17
                    case "foo bar":
                        v18 := GenS1{F0: 128, F1: 2}
                        _ = v18
                    }
                    delete(v2, 4294967295)
                    var v19 bool = (0 != 878)
                    if (157 != 127) {
                        v2[482] = (^(fi13))
                        v20 := GenS1{F0: 255, F1: 2}
                        _ = v20
                    } else {
                        if v19 {
                            v19 = ("hello" != "a")
                            v21 := GenS1{F0: 127, F1: 255}
                            if (!true) {
                                var v22 int32 = 1
                                fmt.Println(v19, v22, v4)
                                v23 := GenS1{F0: 0, F1: 1}
                                _ = len(v1)
                                v24 := GenS1{F0: 127, F1: 32767}
                                fmt.Println(v22, v19)
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        } else {
                            if (42 <= (v3 >> ((899) & 63))) {
                                v25 := GenS2{F0: 42, F1: 1}
                                v19 = v19
                                _ = v25
                            } else {
                                if v19 {
                                    var v26 int = (fi14 ^ (fi13 / ((fi13) | 1)))
                                    _ = v26
                                }
                                fmt.Println(fi14)
                            }
                        }
                    }
                    _ = v19
                }
                _ = fi14
            }
        default:
            v4 += 0.1
        }
        if (genFn0(113, -200) != -32768) {
            var v27 int = (-(fi13))
            _ = v27
        }
        _ = fi13
    }
    fmt.Println(v4, v3)
    v4 -= float64(float32((1.5 + v4)))
    v28 := GenS3{F0: true, F1: 0, F2: 2}
    if false {
        var v29 uint64 = 2
        _ = v29
    } else {
        fmt.Println(v4, v3)
        v1 = append(v1, "\n\t")
    }
    var v30 uint32 = 2147483647
    var v31 bool = true
    v31 = v31
    if false {
        for fi32 := 0; fi32 < 18; fi32++ {
            for fi33 := 0; fi33 < 16; fi33++ {
                v34 := GenS1{F0: 128, F1: 255}
                var v35 bool = v31
                v36 := GenS3{F0: true, F1: 32768, F2: 18446744073709551615}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi32
        }
    } else {
        var v37 float32 = 0.0
        if v31 {
            v38 := GenS3{F0: (42 != -1), F1: -61, F2: 4294967295}
            var v39 uint32 = v30
            delete(v2, 806)
            _ = v38
            _ = v39
        } else {
            for fi40 := 0; fi40 < 10; fi40++ {
                var v41 uint64 = 2
                _ = fi40
                _ = v41
            }
            v42 := GenS3{F0: v31, F1: -1, F2: 108}
            var v43 int8 = (v3 | 100)
            _ = v42
            _ = v43
        }
        var v44 float32 = v37
        _ = v37
        _ = v44
    }
    v3 += (v3 >> ((955) & 63))
    v31 = (!v31)
    fmt.Println(v30)
    if ((0 >= 2147483647) || (true && v31)) {
        fmt.Println(v30)
        var v45 string = ""
        _ = v45
    }
    switch -2147483648 {
    case 2147483647:
        v46 := GenS1{F0: 0, F1: 32767}
        _ = v46
    }
    if v31 {
        if v31 {
            switch "foo bar" {
            case "foo bar":
                v47 := GenS2{F0: genFn0(421, -638), F1: 179}
                _ = v47
            case "hello":
                v30 += v30
            default:
                switch (("" + "\n\t") + "\n\t") {
                case "hello":
                    v3 += v3
                case "a":
                    var v48 int8 = -1
                    _ = v48
                }
            }
        }
        v49 := GenS3{F0: v31, F1: 263, F2: 42}
        _ = v49
    }
    fmt.Println(v31, v4, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
