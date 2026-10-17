package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int32; F2 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    var v3 int64 = -486
    var v4 int8 = 6
    v5 := make(map[int]string)
    v4 = v4
    var v6 int8 = (-(v4))
    var v7 uint16 = 255
    v7 = (v7 | 1)
    var v8 uint = 0
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int8 = (v6 * v6)
        if (v3 > v3) {
            fmt.Println(v6, fi9)
            if (1 == 2147483648) {
                delete(v5, 32767)
                var v11 uint64 = 159
                var v12 uint16 = v7
                _ = v11
                _ = v12
            }
            var v13 bool = true
            _ = v13
        } else {
            v7 -= (v7 | uint16(v8))
            var v14 uint8 = 0
            _ = v14
        }
        if true {
            for fi15 := 0; fi15 < 20; fi15++ {
                switch "" {
                case "a":
                    v16 := []int{127}
                    _ = v16
                }
                if false {
                    v3 -= (v3 >> ((v8) & 63))
                } else {
                    v17 := GenS1{F0: 65535, F1: 819, F2: 4294967295}
                    v3 = -1
                    _ = v17
                }
                fmt.Println(v7, fi15)
                fmt.Println(v3, v3)
                _ = fi15
            }
            if (2147483647 > -92) {
                for fi18 := 0; fi18 < 20; fi18++ {
                    if false {
                        fmt.Println(v6)
                        var v19 int8 = (v10 - v10)
                        _ = v19
                    }
                    var v20 int64 = v3
                    if true {
                        var v21 int16 = 32767
                        _ = v21
                    }
                    _ = fi18
                    _ = v20
                }
                v6 = v4
                var v22 uint8 = 132
                _ = v22
            }
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := []int{}
                var v25 uint32 = 1
                if true {
                    v26 := GenS2{F0: v6, F1: ("foo bar" + ("" + "hello"))}
                    v27 := GenS3{F0: "\n\t"}
                    _ = v26
                    _ = v27
                }
                switch fi9 {
                case 4294967295:
                    var v28 int = (fi9 | 127)
                    _ = v28
                case -168:
                    var v29 uint = (v8 + 0)
                    _ = v29
                case 9223372036854775807:
                    v6 = (v10 * 1)
                }
                switch ("a" + ("foo bar" + "a")) {
                case "a":
                    var v30 uint32 = v25
                    _ = v30
                case "foo bar":
                    var v31 float64 = 1.5
                    _ = v31
                default:
                    v2[(fi9 << ((v8) & 63))] = 1
                }
                v8 += v8
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v6, v4)
    for fi32 := 0; fi32 < 3; fi32++ {
        v6 = (v4 % ((v4) | 1))
        var v33 uint64 = 18446744073709551615
        _ = len(v5)
        _ = fi32
        _ = v33
    }
    var v34 uint = (v8 ^ v8)
    switch (("hello" + "") + "") {
    case "":
        v35 := GenS4{F0: 2, F1: -84, F2: 0}
        _ = v35
    case "foo bar":
        var v36 float64 = ((3.14159 - (((float64(0.0) / (3.14159 - 1e10)) + (0.0 - -2.25)) - 3.14159)) * float64(-2.25))
        _ = v36
    case "\n\t":
        var v37 uint = (^(v34))
        _ = v37
    default:
        var v38 uint16 = (^(v7))
        _ = v38
    }
    v39 := []int{1, 599, -9223372036854775808}
    v39 = append(v39, 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
