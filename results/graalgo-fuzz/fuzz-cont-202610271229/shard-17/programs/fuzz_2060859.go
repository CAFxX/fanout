package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int = -1
    v2 := []int{}
    v3 := make(map[int]uint32)
    var v4 uint = 65535
    v5 := make(map[string]bool)
    v1 += v1
    var v6 uint64 = 0
    v2 = append(v2, (v1 >> ((v4) & 63)))
    if true {
        v6 += v6
        switch v1 {
        case 9223372036854775807:
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 int64 = 889
                for fi9 := 0; fi9 < 12; fi9++ {
                    v3[fi7] = 65535
                    v3[fi7] = 205
                    var v10 bool = (false && true)
                    _ = fi9
                    _ = v10
                }
                fmt.Println(v4, v6)
                _ = fi7
                _ = v8
            }
        }
    }
    var v11 int8 = int8(v4)
    fmt.Println(v11, v11, v11)
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint16 = 1
        v13 -= (v13 * v13)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if false {
            v11 -= -28
            fmt.Println(v4, fi14, v11)
            var v15 int8 = -1
            _ = v15
        }
        v4 = (v4 % ((v4) | 1))
        _ = fi14
    }
    var v16 int = 1
    var v17 int8 = -1
    v18 := GenS1{F0: 692, F1: -740, F2: 128}
    for fi19 := 0; fi19 < 20; fi19++ {
        v11 = (v11 << ((v4) & 63))
        v4 = 2
        var v20 int64 = 1
        _ = fi19
        _ = v20
    }
    if (v4 > v4) {
        var v21 uint32 = 522
        switch "\n\t" {
        case "\n\t":
            for fi22 := 0; fi22 < 10; fi22++ {
                fmt.Println(fi22, v1, v16)
                _ = fi22
            }
        case "foo bar":
            fmt.Println(v11, v17)
        case "a":
            v21 = 4294967295
        default:
            switch 0 {
            case -236:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -3:
                v11 = 1
            case 2147483647:
                var v23 uint64 = (^(v6))
                _ = v23
            }
        }
        _ = v21
    } else {
        if (true && (421 <= v6)) {
            if (171 == -109) {
                v3[(v1 & v1)] = 2
                if (false && false) {
                    var v24 uint = (^(v4))
                    v6 = 0
                    v24 += v4
                    _ = v24
                }
            } else {
                var v25 uint8 = 255
                switch "a" {
                case "\n\t":
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v16 += v1
                        _ = fi26
                    }
                default:
                    switch "\n\t" {
                    case "hello":
                        _ = len(v5)
                    case "a":
                        v27 := GenS2{F0: (v11 ^ -56)}
                        _ = v27
                    }
                }
                switch v16 {
                case 127:
                    var v28 int8 = (-(v11))
                    _ = v28
                }
                _ = v25
            }
        } else {
            var v29 float64 = 1.5
            var v30 int16 = -267
            var v31 uint8 = 0
            _ = v29
            _ = v30
            _ = v31
        }
        if (1 >= -3) {
            v32 := GenS3{F0: 147}
            _ = v32
        } else {
            v33 := GenS3{F0: 708}
            _ = v33
        }
    }
    switch "a" {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        v34 := GenS1{F0: 1, F1: 617, F2: 142}
        _ = v34
    case "hello":
        fmt.Println(v11, v11)
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 string = "hello"
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
