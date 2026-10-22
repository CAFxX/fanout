package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16; F1 int16; F2 bool }
type GenS4 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint8 = 255
    v2 := []int{}
    v3 := make(map[uint32]int)
    var v4 uint32 = 624
    var v5 int = 32768
    var v6 uint16 = 923
    var v7 int32 = 928
    v5 -= (v5 ^ (v5 / ((-436) | 1)))
    fmt.Println(v4)
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v5)
        var v9 string = "hello"
        v10 := GenS1{F0: 1, F1: v4, F2: 481}
        _ = fi8
        _ = v10
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = (v5 << ((2) & 63))
    }
    if (true || true) {
        v11 := GenS2{F0: -69, F1: 100}
        v7 += (v7 ^ v7)
        v1 = 3
        _ = v11
    } else {
        v6 -= v6
        if false {
            for fi12 := 0; fi12 < 2; fi12++ {
                if (true || (((0.1 + 1.5) * 1e10) > -2.25)) {
                    v7 += v7
                } else {
                    if false {
                        if (uint64(v1) >= uint64(v4)) {
                            v7 += 42
                            fmt.Println(v7, v5, v5)
                        } else {
                            v13 := GenS3{F0: 0, F1: 1000, F2: false}
                            _ = v13
                        }
                    }
                }
                _ = fi12
            }
            v7 = v7
        }
    }
    switch 127 {
    case -889:
        for fi14 := 0; fi14 < 7; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                switch (v5 ^ (v5 & -169)) {
                case 4294967295:
                    v16 := GenS2{F0: 2147483647, F1: 1}
                    _ = v16
                case 681:
                    switch (("foo bar" + ("a" + "")) + ("a" + "hello")) {
                    case "a":
                        v17 := GenS2{F0: (v7 % (((v7 | v7)) | 1)), F1: -1}
                        _ = v17
                    case "":
                        v1 += uint8(v6)
                    }
                }
                _ = fi15
            }
            _ = fi14
        }
    case -436:
        v18 := GenS3{F0: 299, F1: 1000, F2: (-9223372036854775808 <= 0)}
        _ = v18
    default:
        var v19 float32 = ((0.1 * 3.14159) + 1.5)
        _ = v19
    }
    fmt.Println(v1, v6, v7)
    var v20 int32 = v7
    fmt.Println(v6, v6, v5)
    _ = len(v2)
    var v21 uint16 = (v6 ^ (v6 >> ((1) & 63)))
    v1 -= v1
    for fi22 := 0; fi22 < 8; fi22++ {
        if false {
            v21 += (v6 & 42)
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 bool = false
                _ = fi23
                _ = v24
            }
            if false {
                fmt.Println(v21)
                v4 = v4
            }
            v25 := GenS4{F0: (((1e10 <= (1.5 * 3.14159)) || true) || false), F1: v1}
            _ = v25
        }
        switch "hello" {
        case "\n\t":
            v26 := GenS3{F0: 32767, F1: 1000, F2: (0 >= 946)}
            _ = v26
        case "foo bar":
            var v27 float32 = (float32(0.0) - ((-2.25 + 1e10) + 0.0))
            _ = v27
        case "":
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 string = ("a" + ("a" + "\n\t"))
                var v30 int8 = -1
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            v21 = v21
        }
        _ = fi22
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
