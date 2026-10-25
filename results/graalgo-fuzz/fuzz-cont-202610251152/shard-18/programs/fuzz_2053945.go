package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint; F1 uint; F2 int32 }
type GenS3 struct { F0 int; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int32; F2 int32 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint32, p1 int8) uint {
    v5 := GenS1{F0: (^(p0)), F1: 0}
    switch "foo bar" {
    case "\n\t":
        v6 := []int{4294967295, -219, 1, -577}
        _ = v6
    case "":
        p0 += 2147483648
    default:
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint64, p1 uint8, p2 uint) bool {
    switch ("hello" + "hello") {
    case "foo bar":
        var v7 int64 = 42
        _ = v7
    default:
        v8 := GenS2{F0: (p2 << (((p2 * (p2 ^ p2))) & 63)), F1: 0, F2: -431}
        _ = v8
    }
    if false {
        p2 = p2
        v9 := GenS3{F0: 491, F1: 1, F2: (^(p1))}
        _ = v9
    } else {
        if true {
            v10 := GenS4{F0: 2147483647, F1: 42, F2: 0}
            var v11 uint32 = 65535
            _ = v10
            _ = v11
        }
        fmt.Println(p1)
        v12 := GenS5{F0: (p0 * p0), F1: 0}
        _ = v12
    }
    var v13 int8 = 109
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return false
}

func main() {
    var v1 uint = 230
    var v2 int = 42
    v3 := []string{"foo bar", "a", "foo bar", "\n\t"}
    var v4 int = (v2 << (((v1 >> ((v1) & 63))) & 63))
    v1 += v1
    v4 -= (v4 << ((v1) & 63))
    v4 += -2147483648
    for fi14 := 0; fi14 < 17; fi14++ {
        switch (fi14 << (((^(v1))) & 63)) {
        case 128:
            for fi15 := 0; fi15 < 6; fi15++ {
                v4 = (v4 - fi15)
                switch fi14 {
                case 32767:
                    var v16 string = "a"
                    _ = v16
                default:
                    fmt.Println(v2, fi14, v2)
                }
                if true {
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 -= v1
                        var v17 int16 = 36
                        _ = v17
                    } else {
                        _ = len(v3)
                        switch "" {
                        case "":
                            if ((v1 << ((646) & 63)) != genFn0(2147483648, 100)) {
                                v4 -= (^(v4))
                                switch (-(fi14)) {
                                case 395:
                                    if false {
                                        switch 4294967295 {
                                        case 484:
                                            if false {
                                                var v18 uint = v1
                                                var v19 uint32 = 638
                                                _ = v18
                                                _ = v19
                                            } else {
                                                _ = len(v3)
                                            }
                                        case 32767:
                                            v3 = append(v3, "foo bar")
                                        default:
                                            v1 = 42
                                        }
                                        var v20 int32 = 1
                                        _ = len(v3)
                                        _ = v20
                                    } else {
                                        fmt.Println(fi15)
                                        v21 := GenS5{F0: uint64(v1), F1: 32767}
                                        _ = v21
                                    }
                                case 140:
                                    var v22 bool = true
                                    _ = v22
                                }
                                var v23 uint = (v1 % (((v1 * genFn0(767, -2))) | 1))
                                _ = v23
                            }
                        }
                    }
                    if false {
                        var v24 int64 = -983
                        _ = v24
                    } else {
                        v25 := GenS4{F0: 390, F1: 42, F2: 2147483647}
                        var v26 uint32 = 1
                        if (-63 >= 0) {
                            v27 := GenS6{F0: 2}
                            v28 := GenS3{F0: v2, F1: uint16(v1), F2: 173}
                            if false {
                                if (true || (-2.25 >= 0.1)) {
                                    v29 := GenS7{F0: -1, F1: 0}
                                    _ = v29
                                } else {
                                    v26 -= v26
                                    v3 = append(v3, "a")
                                }
                                _ = len(v3)
                                var v30 int64 = 73
                                _ = v30
                            } else {
                                var v31 int16 = -32768
                                var v32 string = "\n\t"
                                _ = v31
                                _ = v32
                            }
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(v4, fi15, fi15)
                    var v33 int8 = 1
                    _ = v33
                }
                _ = fi15
            }
        default:
            if (-1 <= 0) {
                fmt.Println(v1, fi14)
                var v34 int32 = -2147483648
                var v35 float32 = 1.5
                _ = v34
                _ = v35
            }
        }
        _ = fi14
    }
    v2 = v4
    fmt.Println(v2, v1, v2)
    for fi36 := 0; fi36 < 7; fi36++ {
        v1 -= (v1 ^ (v1 % (((v1 | (v1 << ((42) & 63)))) | 1)))
        _ = fi36
    }
    if (true || (true || (!false))) {
        fmt.Println(v4, v4)
        var v37 float32 = 0.1
        _ = v37
    }
    var v38 string = "a"
    v1 = (v1 >> (((v1 * 65535)) & 63))
    var v39 uint8 = 255
    var v40 int32 = 489
    v3 = append(v3, v38)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
