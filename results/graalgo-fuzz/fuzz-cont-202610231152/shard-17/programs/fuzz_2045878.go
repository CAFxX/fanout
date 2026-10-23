package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 string; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }
type GenS4 struct { F0 uint; F1 bool }

func main() {
    v1 := []int{128}
    v2 := []int{518}
    v3 := []uint32{65535, 631}
    v4 := make(map[int]uint32)
    var v5 uint = 774
    var v6 string = (("\n\t" + "a") + "a")
    var v7 int16 = 1000
    v7 = (-(v7))
    v5 += (v5 + v5)
    var v8 int = 2147483647
    for fi9 := 0; fi9 < 14; fi9++ {
        v5 += 426
        if (false && (0 > 913)) {
            switch ((v6 + "\n\t") + (v6 + "hello")) {
            case "a":
                fmt.Println(v6, v8)
            default:
                v4[(fi9 - fi9)] = 2
            }
            var v10 int8 = 42
            if (!false) {
                if (0 == (v7 * (v7 + (v7 & (-(v7)))))) {
                    switch ("" + v6) {
                    case "\n\t":
                        fmt.Println(v7)
                    case "":
                        v7 = v7
                    case "hello":
                        if true {
                            v11 := GenS1{F0: 5}
                            _ = v11
                        }
                    }
                }
                var v12 bool = (!(2147483647 < -2147483648))
                _ = v12
            } else {
                var v13 int8 = (^(v10))
                v8 = v8
                if false {
                    switch "\n\t" {
                    case "":
                        v14 := GenS1{F0: 109}
                        _ = v14
                    case "a":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        v5 += 18446744073709551615
                    }
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v16 := GenS1{F0: 1}
                        _ = fi15
                        _ = v16
                    }
                    v17 := GenS2{F0: true, F1: v6, F2: 2147483648}
                    _ = v17
                }
                _ = v13
            }
            _ = v10
        } else {
            v7 -= 8
            switch v6 {
            case "a":
                for fi18 := 0; fi18 < 13; fi18++ {
                    fmt.Println(v5)
                    _ = fi18
                }
            }
        }
        _ = fi9
    }
    var v19 uint64 = 2
    if (0.1 > ((3.14159 + 0.1) - float32((float64(3.14159) - float64(((0.1 / 1.0) + 1.5)))))) {
        for fi20 := 0; fi20 < 12; fi20++ {
            delete(v4, int(v19))
            if (-2.25 < (0.1 + ((0.1 * ((0.1 - float32(((float64(0.1) - 1.5) + 0.0))) - float32((1e10 - (-2.25 + 1e10))))) * 1.5))) {
                fmt.Println(v19, v5, fi20)
            }
            var v21 uint32 = 2
            _ = fi20
            _ = v21
        }
    } else {
        _ = len(v4)
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 int32 = int32(v7)
        v8 -= (v8 >> ((v5) & 63))
        _ = fi22
        _ = v23
    }
    v1 = append(v1, v8)
    switch (v8 | (v8 ^ v8)) {
    case -1:
        if len(v1) > 0 {
            v1[0] = -851
        }
    case 0:
        switch 0 {
        case 9223372036854775807:
            switch v8 {
            case 32768:
                for fi24 := 0; fi24 < 18; fi24++ {
                    switch v8 {
                    case -442:
                        var v25 uint32 = 2
                        _ = v25
                    case 42:
                        switch v6 {
                        case "":
                            fmt.Println(v5, fi24)
                        case "foo bar":
                            var v26 bool = (true && (1 < 2147483647))
                            _ = v26
                        case "hello":
                            fmt.Println(fi24, v6, v6)
                        }
                    }
                    var v27 int = fi24
                    v28 := GenS3{F0: 0, F1: 0, F2: -571}
                    _ = fi24
                    _ = v27
                    _ = v28
                }
            case -2147483648:
                v19 = v19
            case 9223372036854775807:
                if false {
                    var v29 uint16 = 511
                    v7 = (v7 << (((v5 << ((v5) & 63))) & 63))
                    _ = v29
                }
            default:
                var v30 int64 = 0
                _ = v30
            }
        case 127:
            var v31 uint32 = 510
            _ = v31
        default:
            switch v8 {
            case 9223372036854775807:
                var v32 uint16 = 2
                _ = v32
            case 698:
                v33 := GenS3{F0: 0, F1: -2147483648, F2: -1}
                _ = v33
            }
        }
    case 32768:
        if (v6 <= v6) {
            v34 := GenS2{F0: (v6 > "hello"), F1: ((v6 + v6) + ("" + v6)), F2: 2147483648}
            var v35 float64 = (0.0 / ((1e10 + -2.25) / 1.0))
            _ = v34
            _ = v35
        }
    }
    if (((1e10 - 3.14159) <= ((0.1 - -2.25) * 0.1)) && true) {
        v36 := GenS4{F0: v5, F1: (2 > 2147483648)}
        v37 := GenS2{F0: false, F1: v6, F2: 2}
        _ = v36
        _ = v37
    }
    if len(v2) > 0 {
        v2[0] = v8
    }
    if (true || (!true)) {
        var v38 int16 = 664
        _ = v38
    } else {
        v7 += v7
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
