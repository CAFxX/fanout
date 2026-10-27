package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS3 struct { F0 string; F1 string; F2 uint32 }
type GenS4 struct { F0 int16; F1 int8; F2 int8 }
type GenS5 struct { F0 uint; F1 uint; F2 int16 }
type GenS6 struct { F0 uint64; F1 uint32; F2 int32 }
func genFn0(p0 uint16, p1 int, p2 int32) int32 {
    v4 := []int{}
    v5 := GenS1{F0: true, F1: (42 < 4294967295)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}
func genFn1(p0 int8) string {
    p0 += (-(p0))
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 += (p0 ^ p0)
        var v7 int32 = 1
        p0 = p0
        _ = fi6
        _ = v7
    }
    _ = p0
    return ("foo bar" + ("foo bar" + "foo bar"))
}

func main() {
    v1 := []string{}
    var v2 int64 = 42
    v3 := []uint32{417, 309, 172}
    switch -1 {
    case 32767:
        fmt.Println(v2)
    case -9223372036854775808:
        if true {
            v2 = v2
            v2 = -9223372036854775808
        }
    }
    _ = len(v1)
    fmt.Println(v2)
    fmt.Println(v2)
    v2 = (v2 ^ v2)
    switch "hello" {
    case "a":
        if true {
            switch -9223372036854775808 {
            case 128:
                v2 -= (v2 - (v2 * v2))
            case 127:
                v8 := GenS2{F0: 240, F1: 0, F2: 352}
                _ = v8
            default:
                if (255 > 500) {
                    switch "a" {
                    case "":
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    }
                    var v9 uint8 = 42
                    fmt.Println(v9, v2)
                    _ = v9
                }
            }
            switch "foo bar" {
            case "hello":
                v2 = (-(v2))
            case "\n\t":
                fmt.Println(v2)
            default:
                if true {
                    switch -749 {
                    case 32768:
                        var v10 int = 1
                        _ = v10
                    default:
                        v1 = append(v1, ("hello" + "a"))
                    }
                    v2 -= (v2 - (v2 | v2))
                    switch 4294967295 {
                    case -119:
                        v2 += v2
                    case 2147483647:
                        v11 := GenS2{F0: 4294967295, F1: 2147483648, F2: 0}
                        _ = v11
                    default:
                        v1 = append(v1, "a")
                    }
                } else {
                    v12 := GenS3{F0: (("\n\t" + "hello") + "\n\t"), F1: "", F2: 0}
                    _ = v12
                }
            }
        } else {
            v13 := GenS4{F0: 1000, F1: -128, F2: 100}
            fmt.Println(v2)
            _ = v13
        }
    case "\n\t":
        v2 = 9223372036854775807
    case "":
        var v14 uint8 = 127
        _ = v14
    default:
        switch -1 {
        case 32768:
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 float64 = 1.5
                _ = fi15
                _ = v16
            }
        case 128:
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v2)
            case "a":
                if (!(42 >= 0)) {
                    fmt.Println(v2)
                } else {
                    if false {
                        var v17 uint32 = 330
                        var v18 bool = false
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v17
                        _ = v18
                    } else {
                        if true {
                            if (!true) {
                                if true {
                                    var v19 uint16 = 2
                                    _ = v19
                                } else {
                                    var v20 uint16 = 32767
                                    var v21 uint64 = 2
                                    _ = v20
                                    _ = v21
                                }
                                var v22 int8 = -47
                                _ = v22
                            } else {
                                if (v2 == (v2 & v2)) {
                                    v2 = (v2 ^ (v2 / (((v2 ^ v2)) | 1)))
                                    var v23 int16 = -32768
                                    _ = v23
                                }
                                var v24 int = -696
                                _ = v24
                            }
                        } else {
                            v25 := GenS2{F0: 255, F1: 505, F2: 65535}
                            v2 += v2
                            var v26 int = -9223372036854775808
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS5{F0: 479, F1: 416, F2: -455}
                        _ = v27
                    }
                    switch 9223372036854775807 {
                    case -1:
                        v28 := GenS6{F0: 2, F1: 2147483648, F2: 42}
                        _ = v28
                    case 39:
                        v2 -= v2
                    default:
                        v29 := GenS6{F0: 42, F1: 3, F2: 2147483647}
                        _ = v29
                    }
                    v2 += (v2 ^ (^(v2)))
                }
            }
        case -2147483648:
            fmt.Println(v2)
        }
    }
    v3 = append(v3, 2147483647)
    for fi30 := 0; fi30 < 9; fi30++ {
        fmt.Println(fi30, v2)
        var v31 uint64 = 0
        switch (fi30 | fi30) {
        case 128:
            v32 := GenS4{F0: 1, F1: 127, F2: -1}
            _ = v32
        default:
            for fi33 := 0; fi33 < 6; fi33++ {
                v2 = 1
                var v34 int16 = 1000
                var v35 uint32 = 0
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        _ = fi30
        _ = v31
    }
    fmt.Println(v2)
    v1 = append(v1, (("\n\t" + (("hello" + "\n\t") + "a")) + ("\n\t" + "a")))
    if (840 >= 386) {
        switch 128 {
        case -506:
            v2 = -583
        case 0:
            v2 -= 589
        }
    }
    v36 := GenS6{F0: 2, F1: 373, F2: genFn0(65535, 741, 2147483647)}
    v2 = 714
    var v37 int64 = v2
    fmt.Println(v2, v2)
    v37 -= v37
    switch 622 {
    case 1:
        for fi38 := 0; fi38 < 20; fi38++ {
            _ = len(v1)
            var v39 uint64 = 0
            _ = fi38
            _ = v39
        }
    case 2147483647:
        v37 = (-(v2))
    case 128:
        if len(v3) > 0 {
            v3[0] = 543
        }
    default:
        v1 = append(v1, "foo bar")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
