package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0(p0 uint) uint8 {
    var v4 float64 = (0.1 - 0.0)
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 float32 = (0.0 / -2.25)
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 128
}

func main() {
    var v1 bool = (42 > 255)
    var v2 uint8 = 255
    v3 := make(map[string]int)
    v7 := []int{42, 32768}
    fmt.Println(v2, v1)
    switch ("" + "\n\t") {
    case "\n\t":
        fmt.Println(v2, v1)
    default:
        v3["\n\t"] = 4294967295
    }
    v2 = genFn0(255)
    for fi8 := 0; fi8 < 12; fi8++ {
        if (!true) {
            var v9 float64 = (0.0 - 1.5)
            _ = v9
        }
        _ = fi8
    }
    if ((!(2 != 495)) || v1) {
        var v10 float32 = -2.25
        var v11 int8 = 100
        v12 := GenS1{F0: v11}
        _ = v10
        _ = v11
        _ = v12
    } else {
        v1 = true
        var v13 bool = (v1 || (1.5 >= 0.1))
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint64 = 2
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    switch -407 {
    case -4:
        v2 = v2
    case 9223372036854775807:
        switch "hello" {
        case "foo bar":
            for fi16 := 0; fi16 < 7; fi16++ {
                switch ("" + ("a" + "a")) {
                case "\n\t":
                    var v17 uint8 = (v2 + v2)
                    _ = v17
                case "foo bar":
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if v1 {
                            var v19 int = 127
                            var v20 bool = (float32((3.14159 / (float64(-2.25) / -2.25))) != 0.1)
                            _ = v19
                            _ = v20
                        } else {
                            v2 = (^(v2))
                            v21 := GenS1{F0: -128}
                            v1 = (0.1 != ((1e10 + 1.5) * (1e10 + (0.0 / (((0.0 * 1.5) / 0.1) - 1e10)))))
                            _ = v21
                        }
                        _ = fi18
                    }
                }
                _ = fi16
            }
        case "a":
            var v22 uint8 = v2
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        fmt.Println(v2, v1, v1)
        _ = fi23
    }
    var v24 bool = ("hello" == "foo bar")
    switch 32767 {
    case 0:
        var v25 float32 = 3.14159
        _ = v25
    case -841:
        v26 := GenS2{F0: "\n\t", F1: 269}
        _ = v26
    case 886:
        var v27 bool = (602 == 1)
        _ = v27
    default:
        v28 := GenS3{F0: -2147483648, F1: (v2 * (v2 * 128)), F2: 255}
        _ = v28
    }
    var v29 uint32 = 185
    v24 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v7
}
