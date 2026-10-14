package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int32 }

func main() {
    var v1 int32 = 0
    var v2 int32 = v1
    v3 := make(map[uint32]int)
    var v4 int8 = 1
    v5 := []uint32{4294967295, 0}
    var v6 uint8 = 1
    v2 = -2147483648
    var v7 int16 = 32767
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int8 = v4
        _ = fi8
        _ = v9
    }
    v7 -= v7
    if true {
        var v10 int = -9223372036854775808
        v11 := GenS1{F0: 285, F1: v7}
        if true {
            if false {
                v12 := GenS1{F0: 42, F1: v7}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v12
            }
        } else {
            v2 = (v1 + v1)
            v13 := GenS2{F0: 356, F1: "a"}
            _ = v13
        }
        _ = v10
        _ = v11
    }
    var v14 int64 = 9223372036854775807
    switch ("" + ("" + ("foo bar" + ("" + "foo bar")))) {
    case "foo bar":
        fmt.Println(v14)
    default:
        switch "foo bar" {
        case "\n\t":
            var v15 uint16 = 255
            _ = v15
        default:
            v7 += v7
        }
    }
    if (986 >= 2) {
        var v16 float64 = 0.1
        _ = v16
    } else {
        v17 := GenS2{F0: 163, F1: "hello"}
        if ((v4 & -123) < int8(v6)) {
            v18 := GenS3{F0: v6}
            _ = v18
        }
        _ = v17
    }
    switch "foo bar" {
    case "":
        delete(v3, 4294967295)
    case "\n\t":
        for fi19 := 0; fi19 < 18; fi19++ {
            switch "hello" {
            case "":
                switch fi19 {
                case 128:
                    v6 = v6
                case -387:
                    v6 += v6
                default:
                    fmt.Println(v7)
                }
            }
            var v20 float32 = ((3.14159 * ((float32(0.1) + (float32(0.1) - 0.0)) + ((0.1 - (1e10 * 1e10)) / 1.5))) * (-2.25 * 1e10))
            _ = fi19
            _ = v20
        }
    case "hello":
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS4{F0: false, F1: (v1 % ((42) | 1))}
            _ = fi21
            _ = v22
        }
    default:
        if false {
            var v23 uint64 = 0
            fmt.Println(v23, v23)
            for fi24 := 0; fi24 < 12; fi24++ {
                switch "\n\t" {
                case "foo bar":
                    var v25 uint32 = 618
                    _ = v25
                case "hello":
                    if true {
                        var v26 float32 = 1e10
                        _ = v26
                    }
                case "\n\t":
                    fmt.Println(v1, v1, v4)
                default:
                    for fi27 := 0; fi27 < 19; fi27++ {
                        fmt.Println(v7, fi24)
                        _ = fi27
                    }
                }
                _ = fi24
            }
            _ = v23
        } else {
            switch 32767 {
            case -91:
                v14 = (v14 << ((2) & 63))
            default:
                switch "foo bar" {
                case "a":
                    var v28 uint32 = 684
                    _ = v28
                case "foo bar":
                    _ = len(v3)
                case "hello":
                    var v29 uint = 18446744073709551615
                    _ = v29
                }
            }
            if len(v5) > 0 {
                v5[0] = 42
            }
            v5 = append(v5, 65535)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
