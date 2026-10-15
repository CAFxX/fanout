package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint32; F1 uint8; F2 int32 }
func genFn0(p0 uint64) string {
    var v4 uint32 = 1
    var v5 int = -808
    fmt.Println(v5)
    _ = p0
    _ = v4
    _ = v5
    return ""
}
func genFn1() int8 {
    var v6 int = 9223372036854775807
    v6 += 128
    v7 := GenS1{F0: 1, F1: "hello"}
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 461
    var v2 uint64 = 4294967295
    v3 := make(map[string]uint32)
    v2 = v2
    if true {
        var v8 float32 = (0.0 / float32(1.5))
        _ = v8
    } else {
        switch ("hello" + "") {
        case "hello":
            var v9 float64 = 0.0
            _ = v9
        case "foo bar":
            if true {
                if true {
                    v10 := GenS2{F0: 4294967295, F1: 42, F2: -919}
                    v11 := GenS2{F0: 2147483647, F1: 1, F2: 779}
                    _ = v10
                    _ = v11
                } else {
                    v12 := []int{42, 127}
                    var v13 uint64 = v2
                    _ = v12
                    _ = v13
                }
                var v14 int16 = 42
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := []int{fi15, fi15, (fi15 % (((fi15 & (fi15 % ((fi15) | 1)))) | 1))}
                    switch int(v2) {
                    case 32768:
                        if (1.5 >= 3.14159) {
                            var v17 int = (fi15 >> ((255) & 63))
                            if (false && ((67 <= -32) || ((v14 / ((848) | 1)) < v14))) {
                                _ = len(v3)
                                var v18 int32 = 2147483647
                                _ = v18
                            }
                            var v19 string = "hello"
                            _ = v17
                            _ = v19
                        }
                    case -57:
                        v14 = v14
                    case 303:
                        for fi20 := 0; fi20 < 4; fi20++ {
                            var v21 int16 = (^(v14))
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
        }
        v2 = (v2 % (((v2 | 680)) | 1))
    }
    var v22 int64 = -713
    if (0.1 > (0.0 * 0.0)) {
        switch -175 {
        case 763:
            v23 := v3[""]
            _ = v23
        default:
            v22 = -753
        }
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        switch "a" {
        case "foo bar":
            fmt.Println(fi24, v2)
        default:
            v25 := GenS1{F0: 448, F1: "\n\t"}
            _ = v25
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            v1 = v1
            _ = fi26
        }
        _ = fi24
    }
    fmt.Println(v22, v1, v1)
    v27 := []int{1000, 9223372036854775807, 128}
    _ = len(v3)
    if ((v2 / ((1) | 1)) >= 2) {
        var v28 int64 = v22
        _ = v28
    }
    switch ((("hello" + "foo bar") + ("hello" + ("hello" + "hello"))) + "hello") {
    case "\n\t":
        v22 += (v22 ^ v22)
    }
    var v29 uint = 255
    var v30 float64 = (3.14159 + ((1.5 / 1.0) / -2.25))
    fmt.Println(v30)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v29
    _ = v3
    _ = v30
}
