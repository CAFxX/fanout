package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string; F1 int32; F2 int16 }

func main() {
    var v1 uint32 = 732
    var v2 uint = 2
    var v3 float64 = 1.5
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    var v6 string = ("hello" + ((("foo bar" + "foo bar") + "\n\t") + "\n\t"))
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int32 = -2147483648
        for fi9 := 0; fi9 < 6; fi9++ {
            v1 -= (v1 / ((uint32(v2)) | 1))
            switch v6 {
            case "\n\t":
                switch "\n\t" {
                case "a":
                    v2 = 18446744073709551615
                case "":
                    var v10 uint64 = 2
                    _ = v10
                case "\n\t":
                    fmt.Println(fi9, v6)
                default:
                    var v11 int16 = 1
                    _ = v11
                }
            case "hello":
                var v12 uint = (v2 % ((v2) | 1))
                _ = v12
            case "foo bar":
                if true {
                    fmt.Println(v6, v1)
                    v4[(v1 / ((v1) | 1))] = fi9
                    var v13 bool = (v6 <= (("foo bar" + v6) + v6))
                    _ = v13
                } else {
                    v3 += v3
                    fmt.Println(v2)
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    v14 := GenS1{F0: 42, F1: (v1 & v1)}
    var v15 uint8 = 2
    _ = len(v5)
    v3 += (float64(0.0) * ((v3 - 3.14159) / (0.0 / (-2.25 / v3))))
    v16 := GenS1{F0: 4294967295, F1: 0}
    delete(v4, (v1 & (^(v1))))
    if (597 == 4294967295) {
        v2 += v2
        switch v6 {
        case "\n\t":
            v1 += (v1 % (((v1 << ((0) & 63))) | 1))
        default:
            var v17 float64 = (1e10 - v3)
            _ = v17
        }
        v15 -= v15
    }
    switch -685 {
    case 4294967295:
        v2 += (v2 * (^(v2)))
    case 32768:
        v6 = ((v6 + v6) + v6)
    default:
        v18 := v4[v1]
        _ = v18
    }
    switch (v6 + ((v6 + v6) + v6)) {
    case "foo bar":
        v2 += v2
    }
    switch "foo bar" {
    case "\n\t":
        if (true && (3.14159 > 1.5)) {
            for fi19 := 0; fi19 < 4; fi19++ {
                switch (v6 + v6) {
                case "a":
                    v20 := []int{0, (fi19 << ((4294967295) & 63))}
                    _ = v20
                case "":
                    fmt.Println(v6, fi19)
                default:
                    fmt.Println(fi19, v6, v6)
                }
                fmt.Println(v3)
                _ = fi19
            }
        } else {
            if true {
                v21 := []int{}
                for fi22 := 0; fi22 < 9; fi22++ {
                    var v23 uint = (v2 >> ((52) & 63))
                    _ = fi22
                    _ = v23
                }
                v24 := GenS1{F0: 33, F1: (v1 ^ (v1 ^ v1))}
                _ = v21
                _ = v24
            }
        }
    case "":
        if true {
            var v25 uint16 = 42
            switch 4294967295 {
            case 448:
                var v26 int32 = -590
                _ = v26
            case -636:
                fmt.Println(v1, v1)
            default:
                var v27 float64 = 0.1
                _ = v27
            }
            _ = v25
        }
    case "a":
        v28 := GenS1{F0: uint64(v15), F1: (v1 >> (((v2 + 255)) & 63))}
        _ = v28
    default:
        v29 := GenS2{F0: "a", F1: -2147483648, F2: 890}
        _ = v29
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
