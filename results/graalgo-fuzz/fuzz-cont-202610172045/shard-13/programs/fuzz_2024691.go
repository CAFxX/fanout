package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint16; F2 bool }
type GenS4 struct { F0 int }

func main() {
    v1 := []int{0, -9223372036854775808, -1, -1}
    v2 := make(map[int]string)
    var v3 string = ("foo bar" + ((("a" + "") + ("a" + "foo bar")) + "a"))
    var v4 uint16 = 823
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (18446744073709551615 == 77) {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 int64 = -250
            _ = fi5
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            switch (fi7 / ((fi7) | 1)) {
            case -2147483648:
                var v8 uint32 = 134
                _ = v8
            case 9223372036854775807:
                for fi9 := 0; fi9 < 5; fi9++ {
                    v4 = (v4 ^ v4)
                    switch (fi7 | fi9) {
                    case -1:
                        fmt.Println(fi9, fi9)
                    }
                    _ = fi9
                }
            }
            fmt.Println(v4, fi7)
            _ = fi7
        }
    }
    fmt.Println(v4, v4)
    v10 := v2[594]
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint = 18446744073709551615
        _ = len(v1)
        v13 := GenS1{F0: v12, F1: 28}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int32 = -994
    if true {
        fmt.Println(v4, v3)
        if false {
            var v15 int32 = (v14 ^ (-(v14)))
            switch "hello" {
            case "hello":
                var v16 uint16 = (^(v4))
                _ = v16
            case "\n\t":
                var v17 uint = 872
                _ = v17
            case "foo bar":
                if false {
                    v18 := v2[795]
                    v1 = append(v1, 128)
                    if false {
                        var v19 float64 = (-2.25 * (3.14159 / 1e10))
                        v20 := GenS2{F0: ""}
                        v10 = (v18 + "a")
                        _ = v19
                        _ = v20
                    }
                    _ = v18
                }
            default:
                v15 += (v15 / ((912) | 1))
            }
            fmt.Println(v15, v3)
            _ = v15
        }
    }
    var v21 uint64 = 18446744073709551615
    v3 = "\n\t"
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 uint = 0
        fmt.Println(v10, v10)
        fmt.Println(v14, v23, v21)
        _ = fi22
        _ = v23
    }
    var v24 int16 = -17
    v24 = v24
    for fi25 := 0; fi25 < 15; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            v4 += (v4 & v4)
            fmt.Println(fi25, v24)
            var v27 uint32 = 1
            var v28 int = (fi26 << ((177) & 63))
            switch "\n\t" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v29 := v2[fi25]
                if true {
                    if (false || false) {
                        _ = len(v1)
                        v2[(v28 - -422)] = v3
                    } else {
                        if (-226 <= -149) {
                            v14 -= (v14 / ((-1) | 1))
                            fmt.Println(v24, v28, v4)
                        }
                        if true {
                            var v30 int8 = 127
                            var v31 int = v28
                            _ = v30
                            _ = v31
                        } else {
                            v2[fi26] = ((("foo bar" + v10) + "a") + "a")
                            v32 := GenS3{F0: 0, F1: 1, F2: ((v24 ^ v24) <= (v24 + v24))}
                            v21 += v21
                            _ = v32
                        }
                        switch (fi26 * (fi25 | -732)) {
                        case 127:
                            fmt.Println(v10, v27, v27)
                        case 4294967295:
                            v33 := GenS4{F0: (fi26 >> ((719) & 63))}
                            _ = v33
                        default:
                            var v34 uint32 = v27
                            _ = v34
                        }
                    }
                } else {
                    switch fi25 {
                    case 9223372036854775807:
                        _ = len(v2)
                    case -436:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case -1:
                        v28 = v28
                    default:
                        var v35 int = (fi25 >> ((2) & 63))
                        _ = v35
                    }
                }
                _ = v29
            default:
                _ = len(v2)
            }
            if true {
                var v36 int64 = 9223372036854775807
                v37 := GenS1{F0: 116, F1: v14}
                _ = v36
                _ = v37
            }
            v27 += (^(v27))
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    v1 = append(v1, 32767)
    v14 -= (v14 - v14)
    v38 := GenS3{F0: 127, F1: 908, F2: (false || (0.1 > ((1e10 + (float32(-2.25) * 0.1)) + ((0.1 * 0.1) + 1.5))))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v38
    _ = v4
}
