package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []string{(("foo bar" + ("" + "")) + (("foo bar" + "hello") + "foo bar")), "a", ""}
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    var v5 int8 = 127
    v6 := []int{}
    switch "" {
    case "":
        v5 -= v5
    case "foo bar":
        if false {
            v7 := GenS1{F0: -72, F1: (!(42 > -2147483648))}
            v2[""] = 42
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 bool = false
                if true {
                    if v9 {
                        var v10 int8 = v5
                        v3[fi8] = -623
                        _ = len(v2)
                        _ = v10
                    }
                    if ((!(3.14159 >= ((0.0 + 0.0) * float32(3.14159)))) || (42 == 2)) {
                        var v11 int8 = (^(v5))
                        _ = v11
                    }
                } else {
                    for fi12 := 0; fi12 < 17; fi12++ {
                        switch "foo bar" {
                        case "hello":
                            var v13 int64 = 1
                            _ = v13
                        case "a":
                            var v14 uint16 = 753
                            _ = v14
                        default:
                            switch fi12 {
                            case 4294967295:
                                var v15 int32 = 448
                                _ = v15
                            }
                        }
                        _ = fi12
                    }
                    var v16 string = "a"
                    fmt.Println(v9, v9, fi8)
                    _ = v16
                }
                v6 = append(v6, int(v5))
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            var v17 int64 = -9223372036854775808
            var v18 uint64 = 537
            _ = v17
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 uint64 = 4294967295
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    }
    switch -1 {
    case 9223372036854775807:
        for fi22 := 0; fi22 < 3; fi22++ {
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 uint16 = 4
                _ = fi23
                _ = v24
            }
            switch ("\n\t" + "a") {
            case "":
                v25 := GenS2{F0: 948, F1: 42, F2: 1}
                _ = v25
            case "hello":
                if false {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        var v27 int16 = 114
                        var v28 int32 = -88
                        v29 := GenS1{F0: 9223372036854775807, F1: (42 > 2147483647)}
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    switch "hello" {
                    case "a":
                        switch ("" + "\n\t") {
                        case "\n\t":
                            var v30 int8 = (v5 << ((2) & 63))
                            _ = v30
                        case "foo bar":
                            for fi31 := 0; fi31 < 19; fi31++ {
                                v5 = -1
                                _ = fi31
                            }
                        }
                    case "":
                        v5 = (v5 & 100)
                    }
                } else {
                    var v32 int8 = v5
                    _ = v32
                }
            default:
                var v33 int8 = -128
                _ = v33
            }
            _ = fi22
        }
    case 32768:
        fmt.Println(v5)
    case 127:
        _ = len(v2)
    default:
        fmt.Println(v5)
    }
    v5 = v5
    v5 = (v5 % ((42) | 1))
    var v34 float64 = 1.5
    var v35 uint8 = 82
    var v36 bool = true
    v37 := GenS3{F0: 0}
    v36 = false
    v34 += v34
    switch ("\n\t" + "hello") {
    case "":
        var v38 uint = 1
        _ = v38
    case "hello":
        v34 -= (-2.25 + v34)
    case "\n\t":
        v35 -= 128
    default:
        v34 = ((float64(0.0) / 1e10) / (v34 + (v34 / v34)))
    }
    v36 = v36
    _ = len(v1)
    v35 = (v35 / (((v35 - (v35 & v35))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
