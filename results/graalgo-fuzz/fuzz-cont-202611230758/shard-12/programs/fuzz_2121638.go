package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = -2.25
    v3 := []int{997, -1, 2147483647, -1}
    v4 := make(map[int]uint32)
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    v7 := GenS1{F0: 42, F1: 0}
    if true {
        var v8 uint = 811
        _ = v8
    } else {
        v2 -= 1e10
        var v9 uint16 = 163
        _ = v9
    }
    v2 -= 1.5
    var v10 bool = true
    for fi11 := 0; fi11 < 5; fi11++ {
        switch -2147483648 {
        case 32768:
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int8 = -1
                v10 = (!true)
                v14 := GenS1{F0: 1, F1: 1}
                _ = fi12
                _ = v13
                _ = v14
            }
        case 42:
            var v15 uint64 = 89
            _ = v15
        case -9223372036854775808:
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 int32 = 1
                var v18 int64 = 979
                v19 := GenS2{F0: "\n\t", F1: "a"}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        }
        var v20 bool = (!v10)
        _ = fi11
        _ = v20
    }
    v21 := GenS3{F0: 534, F1: 255}
    v22 := GenS3{F0: 18446744073709551615, F1: 2}
    fmt.Println(v2, v10)
    _ = len(v6)
    var v23 int = -484
    var v24 float64 = 0.1
    if false {
        var v25 uint32 = 1
        for fi26 := 0; fi26 < 5; fi26++ {
            for fi27 := 0; fi27 < 16; fi27++ {
                fmt.Println(fi26, v10, fi26)
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v25)
                case "hello":
                    if v10 {
                        var v28 int32 = -581
                        _ = v28
                    } else {
                        var v29 int = fi26
                        switch fi27 {
                        case 7:
                            if len(v3) > 0 {
                                v3[0] = (v29 >> ((uint(v25)) & 63))
                            }
                        default:
                            v24 -= v24
                        }
                        _ = v29
                    }
                case "":
                    v30 := GenS4{F0: 45, F1: 18446744073709551615}
                    _ = v30
                }
                switch ("foo bar" + "foo bar") {
                case "hello":
                    var v31 bool = (v25 > 0)
                    _ = v31
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi27
            }
            if v10 {
                var v32 uint = uint(v25)
                var v33 int32 = 2147483647
                _ = v32
                _ = v33
            }
            _ = fi26
        }
        _ = v25
    }
    var v34 uint = 283
    switch "a" {
    case "a":
        v35 := GenS5{F0: -1}
        _ = v35
    default:
        v4[(v23 | v23)] = 65535
    }
    if v10 {
        v23 = (v23 << ((2) & 63))
    }
    var v36 int64 = 613
    v23 -= -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
