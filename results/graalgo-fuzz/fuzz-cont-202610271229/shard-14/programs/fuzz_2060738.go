package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 uint16, p1 uint64) bool {
    var v7 uint8 = 1
    var v8 string = "a"
    switch int(v7) {
    case 0:
        var v9 int64 = 42
        _ = v9
    }
    v10 := []int{128, 0}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    return ((^(v7)) != 1)
}
func genFn1(p0 uint64) int16 {
    v11 := GenS1{F0: 4294967295}
    v12 := GenS1{F0: 255}
    _ = p0
    _ = v11
    _ = v12
    return -821
}

func main() {
    v1 := []string{}
    var v2 uint = 976
    var v3 float64 = float64(-2.25)
    v4 := make(map[int]int)
    var v5 float64 = v3
    var v6 uint32 = 2147483647
    _ = len(v1)
    if ((2147483647 >= -848) && (255 >= 0)) {
        var v13 int64 = -9223372036854775808
        _ = v13
    } else {
        switch -198 {
        case 679:
            if (0.1 > (1.5 * 3.14159)) {
                v5 += v3
                v5 = (-2.25 / v3)
                var v14 int16 = 1000
                _ = v14
            }
        case 571:
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 uint8 = uint8(v2)
                _ = fi15
                _ = v16
            }
        }
        var v17 uint8 = 42
        v18 := GenS1{F0: 348}
        _ = v17
        _ = v18
    }
    v2 -= (v2 * (^(v2)))
    if ((v2 % ((v2) | 1)) < (v2 & v2)) {
        if false {
            if (2147483647 > 42) {
                fmt.Println(v5, v5, v3)
            }
        } else {
            var v19 int16 = -32768
            if ((3.14159 * v3) < 1.5) {
                _ = len(v4)
                fmt.Println(v5, v3)
            }
            v2 += v2
            _ = v19
        }
    } else {
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v20 int8 = 100
            _ = v20
        default:
            var v21 int = -775
            _ = v21
        }
        v3 = (v5 - float64(-2.25))
        if false {
            v3 += (v5 * (((3.14159 * v3) + v3) + ((v3 / (((v3 / (v5 * v3)) - 0.0) / v3)) * float64(((-2.25 + 0.1) + 1e10)))))
            for fi22 := 0; fi22 < 18; fi22++ {
                if false {
                    v2 -= (v2 % ((v2) | 1))
                    for fi23 := 0; fi23 < 8; fi23++ {
                        var v24 int32 = 595
                        v3 = 3.14159
                        if (127 > 127) {
                            v5 -= (float64((0.1 * (((-2.25 - float32(float64((0.0 - -2.25)))) - float32(v3)) * 1e10))) * v3)
                        } else {
                            var v25 uint16 = 255
                            v26 := v4[int(v24)]
                            _ = v25
                            _ = v26
                        }
                        _ = fi23
                        _ = v24
                    }
                } else {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        v5 += (0.1 - v5)
                        if false {
                            v3 += (v5 / v5)
                            var v28 uint = v2
                            var v29 uint8 = 0
                            _ = v28
                            _ = v29
                        }
                        _ = fi27
                    }
                }
                if false {
                    v30 := GenS1{F0: v2}
                    _ = v30
                }
                _ = fi22
            }
        }
    }
    fmt.Println(v5, v5, v6)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        fmt.Println(v3, v5)
        v31 := GenS2{F0: 65535}
        v5 = 1.5
        _ = v31
    } else {
        var v32 int = -9223372036854775808
        _ = len(v1)
        _ = v32
    }
    var v33 uint64 = 18446744073709551615
    switch ("a" + "foo bar") {
    case "hello":
        v34 := GenS3{F0: int(v6), F1: 980}
        _ = v34
    case "foo bar":
        var v35 uint64 = v33
        _ = v35
    case "":
        if true {
            v2 -= uint(v33)
        }
    }
    var v36 uint8 = 127
    if true {
        switch "a" {
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "":
            fmt.Println(v6, v3)
        case "a":
            v2 -= (v2 << (((^(v2))) & 63))
        default:
            v36 = 127
        }
    } else {
        v6 += v6
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
