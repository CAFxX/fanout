package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS4 struct { F0 string }
func genFn0(p0 int64, p1 string) int16 {
    p1 = "hello"
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(fi6, p1, p0)
        v7 := []int{fi6, 4294967295, 32768, fi6}
        _ = fi6
        _ = v7
    }
    v8 := []int{1, 1, 42, 243}
    v9 := GenS1{F0: 1000, F1: (p0 >> ((112) & 63)), F2: true}
    var v10 float32 = 0.1
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 611
}

func main() {
    v1 := []int{-927, -9223372036854775808}
    var v2 int8 = 0
    var v3 int16 = 42
    var v4 int = 833
    var v5 uint32 = 976
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(v5, v3)
        if (v3 <= v3) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v1 = append(v1, (-(v4)))
            for fi12 := 0; fi12 < 4; fi12++ {
                v1 = append(v1, 42)
                _ = fi12
            }
        } else {
            var v13 float32 = ((0.1 - 0.0) + ((3.14159 + 0.0) + (((((((((1e10 + ((-2.25 * (0.0 / 1.0)) + 0.1)) + 0.0) - 0.0) * -2.25) - 0.1) - (float32(float64((3.14159 * 0.1))) - 1e10)) - (1.5 / float32(3.14159))) + float32(((((float64((1.5 - 3.14159)) * 1.5) - 1e10) / (1.5 / 3.14159)) * -2.25))) - 3.14159)))
            v13 -= v13
            var v14 uint = 0
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    switch (^(v4)) {
    case -408:
        for fi15 := 0; fi15 < 20; fi15++ {
            if (!(!true)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch "\n\t" {
                case "a":
                    if ((0 != 42) || false) {
                        v5 += 1
                    } else {
                        var v16 int8 = 16
                        _ = v16
                    }
                case "foo bar":
                    if (!false) {
                        switch (v4 << ((65535) & 63)) {
                        case -1:
                            _ = len(v1)
                        case 4294967295:
                            if (int16(v4) <= 32767) {
                                fmt.Println(v3)
                            }
                        case 127:
                            v17 := GenS2{F0: 86, F1: 9223372036854775807, F2: 1}
                            _ = v17
                        }
                    } else {
                        var v18 int = (v4 - (fi15 | fi15))
                        _ = v18
                    }
                case "":
                    if len(v1) > 0 {
                        v1[0] = (fi15 & fi15)
                    }
                }
            } else {
                if (!(!true)) {
                    switch "foo bar" {
                    case "foo bar":
                        v3 += (v3 >> ((65535) & 63))
                    }
                    var v19 string = ("a" + "\n\t")
                    var v20 uint64 = 17
                    _ = v19
                    _ = v20
                }
                switch "hello" {
                case "hello":
                    v4 = -1
                case "foo bar":
                    v21 := GenS3{F0: 255, F1: 0, F2: 521}
                    _ = v21
                case "\n\t":
                    for fi22 := 0; fi22 < 6; fi22++ {
                        _ = len(v1)
                        _ = fi22
                    }
                default:
                    if false {
                        _ = len(v1)
                    } else {
                        v2 = v2
                    }
                }
                for fi23 := 0; fi23 < 11; fi23++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 += (v2 ^ 1)
                    var v24 uint = 42
                    v2 = (v2 - (v2 * v2))
                    _ = fi23
                    _ = v24
                }
            }
            var v25 float64 = (-2.25 * (1e10 - 1e10))
            _ = fi15
            _ = v25
        }
    case 1:
        v26 := GenS4{F0: ("a" + ("hello" + "\n\t"))}
        _ = v26
    case 2147483647:
        if len(v1) > 0 {
            v1[0] = 128
        }
    default:
        for fi27 := 0; fi27 < 12; fi27++ {
            switch 42 {
            case 1:
                _ = len(v1)
            case -886:
                switch fi27 {
                case 945:
                    switch "foo bar" {
                    case "\n\t":
                        switch v4 {
                        case -443:
                            v28 := GenS2{F0: 127, F1: -1, F2: 42}
                            _ = v28
                        default:
                            var v29 int16 = (v3 << ((uint(v5)) & 63))
                            _ = v29
                        }
                    default:
                        v1 = append(v1, 2147483647)
                    }
                }
            default:
                v2 += 59
            }
            _ = fi27
        }
    }
    v4 = (v4 & v4)
    v30 := GenS3{F0: 1, F1: 0, F2: 15}
    var v31 uint = 295
    for fi32 := 0; fi32 < 2; fi32++ {
        fmt.Println(v5, v2)
        v31 -= v31
        v5 = 2
        _ = fi32
    }
    if true {
        if (true && false) {
            v31 -= v31
            for fi33 := 0; fi33 < 15; fi33++ {
                var v34 uint8 = 63
                var v35 string = "\n\t"
                _ = fi33
                _ = v34
                _ = v35
            }
            fmt.Println(v31, v4)
        }
        v31 -= 2
    } else {
        var v36 uint8 = 255
        if len(v1) > 0 {
            v1[0] = (^(v4))
        }
        fmt.Println(v5, v5, v4)
        _ = v36
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        v38 := GenS3{F0: 363, F1: 1, F2: 799}
        _ = fi37
        _ = v38
    }
    var v39 uint32 = (v5 & 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v5
}
