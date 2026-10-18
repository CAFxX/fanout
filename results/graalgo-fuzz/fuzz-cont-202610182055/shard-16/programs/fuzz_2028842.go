package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint, p1 uint8, p2 int16) uint8 {
    v6 := GenS1{F0: 2147483648, F1: 1}
    v7 := GenS1{F0: 4294967295, F1: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 bool = false
    var v2 uint8 = 127
    v3 := make(map[string]bool)
    var v4 int8 = 42
    var v5 int64 = -762
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := []int{fi8, fi8, -747}
        var v10 uint8 = 2
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float32 = (0.0 + float32(-2.25))
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint8 = v2
        switch ("foo bar" + "") {
        case "foo bar":
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 uint64 = 216
                _ = fi14
                _ = v15
            }
        }
        _ = fi12
        _ = v13
    }
    v16 := []int{1}
    _ = len(v3)
    if (146 < 0) {
        v5 += (-(v5))
        v1 = ((1 >= 4294967295) && true)
    } else {
        fmt.Println(v11)
    }
    v17 := GenS1{F0: 397, F1: (v2 / ((v2) | 1))}
    v5 += 107
    if (133 >= 255) {
        var v18 float32 = v11
        var v19 bool = (255 < 1)
        v4 = v4
        _ = v18
        _ = v19
    }
    switch "\n\t" {
    case "hello":
        switch 400 {
        case 2147483647:
            var v20 int16 = 61
            _ = v20
        case 825:
            v2 = 2
        case 42:
            var v21 uint = 18446744073709551615
            _ = v21
        }
    case "":
        for fi22 := 0; fi22 < 19; fi22++ {
            v5 = 42
            for fi23 := 0; fi23 < 12; fi23++ {
                v24 := GenS2{F0: v2, F1: v5}
                var v25 uint32 = 2147483648
                fmt.Println(v5, fi22)
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
    case "\n\t":
        var v26 int = -747
        _ = v26
    }
    var v27 uint64 = 868
    var v28 uint16 = 1
    v29 := GenS1{F0: 0, F1: genFn0(4294967295, (v2 | genFn0(1, v2, -187)), 1)}
    if v1 {
        if v1 {
            v30 := GenS2{F0: (v2 % ((uint8(v28)) | 1)), F1: 0}
            if (0 < 490) {
                v31 := GenS3{F0: (v5 << ((1) & 63))}
                _ = v31
            }
            _ = v30
        }
        v11 = 1e10
        if v1 {
            var v32 uint8 = v2
            for fi33 := 0; fi33 < 11; fi33++ {
                for fi34 := 0; fi34 < 12; fi34++ {
                    var v35 int32 = 804
                    if v1 {
                        v36 := v3["a"]
                        switch ("" + "foo bar") {
                        case "a":
                            var v37 bool = ((v35 + int32(v27)) <= (v35 ^ (v35 & v35)))
                            _ = v37
                        case "foo bar":
                            var v38 uint = 436
                            _ = v38
                        default:
                            switch "a" {
                            case "a":
                                v39 := GenS1{F0: 74, F1: v32}
                                _ = v39
                            case "hello":
                                var v40 uint16 = (v28 % (((v28 >> ((0) & 63))) | 1))
                                _ = v40
                            case "":
                                if v1 {
                                    var v41 int = fi34
                                    _ = v41
                                }
                            default:
                                switch (fi34 ^ (-(fi34))) {
                                case 32767:
                                    v32 = uint8(v28)
                                default:
                                    if (false && false) {
                                        v42 := GenS3{F0: (v5 | (^(v5)))}
                                        fmt.Println(v35, fi33, v27)
                                        _ = v42
                                    }
                                }
                            }
                        }
                        v5 = (v5 | -9223372036854775808)
                        _ = v36
                    } else {
                        v11 -= (v11 + (v11 * v11))
                    }
                    var v43 int16 = -1
                    _ = fi34
                    _ = v35
                    _ = v43
                }
                if (-2147483648 != 2147483647) {
                    fmt.Println(v32, v1)
                    fmt.Println(v11, v32)
                }
                switch "a" {
                case "foo bar":
                    var v44 uint32 = 42
                    _ = v44
                default:
                    fmt.Println(v4, v28)
                }
                _ = fi33
            }
            _ = v32
        } else {
            switch "\n\t" {
            case "":
                var v45 uint16 = v28
                _ = v45
            }
            v4 -= (v4 << ((255) & 63))
        }
    }
    v4 = v4
    v11 -= v11
    for fi46 := 0; fi46 < 9; fi46++ {
        v11 = v11
        _ = fi46
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        fmt.Println(v1, v2, v5)
        var v48 int = (^(fi47))
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
