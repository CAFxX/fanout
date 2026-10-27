package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8; F1 int64; F2 string }
type GenS4 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2147483648, 773, 42, 0}
    var v3 float64 = 1e10
    var v4 uint64 = 2
    v4 += v4
    switch -1 {
    case -898:
        _ = len(v2)
    case 42:
        var v5 string = "hello"
        _ = v5
    default:
        var v6 string = "a"
        _ = v6
    }
    if ((1e10 + (1e10 * (((0.0 / 1e10) * 0.1) * (1.5 - 1.5)))) != (0.1 * (0.1 / ((float32(3.14159) + (0.1 - 1.5)) + 0.1)))) {
        v3 = ((v3 - -2.25) / 1.0)
        _ = len(v1)
    }
    v7 := GenS1{F0: 128, F1: 125, F2: 65535}
    var v8 uint64 = (v4 + v4)
    var v9 bool = false
    v10 := GenS1{F0: 255, F1: 0, F2: 1}
    var v11 uint = 2
    v12 := GenS1{F0: 127, F1: 1, F2: uint32(v4)}
    v13 := GenS2{F0: -64, F1: 100}
    v14 := GenS2{F0: 100, F1: 1}
    fmt.Println(v11)
    var v15 uint8 = 0
    v16 := GenS3{F0: v15, F1: -9223372036854775808, F2: ("hello" + "foo bar")}
    switch "" {
    case "foo bar":
        v17 := GenS2{F0: 23, F1: 127}
        _ = v17
    case "":
        v2 = append(v2, 732)
    case "a":
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 int = fi18
            if ((!("" > "foo bar")) && (v9 || (false || v9))) {
                for fi20 := 0; fi20 < 6; fi20++ {
                    if true {
                        v11 -= v11
                        v2 = append(v2, 476)
                    }
                    _ = fi20
                }
                v21 := GenS4{F0: v9, F1: 805, F2: v15}
                if (v11 >= v11) {
                    if (v15 > v15) {
                        switch ("" + "") {
                        case "a":
                            for fi22 := 0; fi22 < 10; fi22++ {
                                fmt.Println(v19)
                                var v23 int64 = -141
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch "hello" {
                                case "hello":
                                    v4 -= 18446744073709551615
                                case "foo bar":
                                    v2 = append(v2, 1)
                                }
                                _ = fi22
                                _ = v23
                            }
                        default:
                            v24 := GenS3{F0: v15, F1: 9223372036854775807, F2: "a"}
                            _ = v24
                        }
                        v19 = 9223372036854775807
                        v25 := GenS5{F0: 42, F1: 621}
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 2; fi26++ {
                        var v27 uint32 = 65535
                        _ = fi26
                        _ = v27
                    }
                }
                _ = v21
            }
            v4 += v4
            _ = fi18
            _ = v19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
