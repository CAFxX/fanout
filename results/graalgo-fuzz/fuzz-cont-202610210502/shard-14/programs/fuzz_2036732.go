package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int8; F1 string; F2 int }

func main() {
    v1 := []string{}
    v2 := make(map[int]uint32)
    v3 := []uint32{4294967295}
    var v4 float32 = -2.25
    fmt.Println(v4)
    if true {
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := GenS1{F0: 9223372036854775807, F1: 32767}
            switch ("a" + "\n\t") {
            case "hello":
                for fi7 := 0; fi7 < 16; fi7++ {
                    if true {
                        v8 := GenS2{F0: 0, F1: 2}
                        v9 := GenS3{F0: ("a" + "hello"), F1: 0}
                        v10 := GenS4{F0: -1}
                        v4 = float32(((1.5 / (float64(v4) / 1e10)) - 0.1))
                        _ = v10
                        _ = v8
                        _ = v9
                    } else {
                        fmt.Println(v4, fi5)
                    }
                    v4 = v4
                    if (!false) {
                        v11 := GenS4{F0: 42}
                        _ = v11
                    } else {
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            if false {
                                var v12 float64 = float64(float32(0.1))
                                v13 := GenS5{F0: 4294967295, F1: 1}
                                _ = v12
                                _ = v13
                            } else {
                                var v14 uint64 = 285
                                v14 += v14
                                v4 = float32((1e10 / 3.14159))
                                _ = v14
                            }
                        default:
                            switch ("hello" + "hello") {
                            case "":
                                v4 = v4
                            case "hello":
                                fmt.Println(fi5)
                            }
                        }
                        var v15 float32 = v4
                        v16 := GenS3{F0: ("\n\t" + ("\n\t" + ((("hello" + "a") + "\n\t") + ""))), F1: 32767}
                        _ = v15
                        _ = v16
                    }
                    _ = fi7
                }
            case "a":
                v4 = ((float32(1.5) * v4) / v4)
            case "\n\t":
                fmt.Println(v4, v4)
            default:
                if false {
                    fmt.Println(v4, fi5)
                }
            }
            var v17 uint8 = 127
            _ = fi5
            _ = v17
            _ = v6
        }
    }
    fmt.Println(v4)
    delete(v2, 42)
    v4 = 1e10
    var v18 uint = 2
    fmt.Println(v4, v4)
    v4 += ((v4 / ((float32((1.5 * 1e10)) * (v4 * 0.1)) / -2.25)) - v4)
    fmt.Println(v4, v4)
    v19 := GenS6{F0: 1, F1: ("\n\t" + ""), F2: -9223372036854775808}
    v20 := GenS4{F0: 1}
    v4 -= v4
    var v21 uint8 = uint8(v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
