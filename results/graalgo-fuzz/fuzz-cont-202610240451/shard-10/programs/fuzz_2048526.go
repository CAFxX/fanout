package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint; F1 bool }

func main() {
    var v1 int = -2147483648
    v2 := make(map[string]uint32)
    var v3 int8 = 1
    var v4 uint16 = 32767
    if (!true) {
        var v5 uint = 69
        if (v1 < 32768) {
            v6 := GenS1{F0: 1, F1: v5, F2: 255}
            _ = v6
        }
        v1 = v1
        _ = v5
    } else {
        var v7 uint16 = (v4 & 0)
        fmt.Println(v7, v1, v1)
        _ = len(v2)
        _ = v7
    }
    v8 := []int{}
    fmt.Println(v3)
    fmt.Println(v3)
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v8 = append(v8, v1)
    case "a":
        v9 := GenS2{F0: "\n\t"}
        _ = v9
    default:
        v10 := GenS2{F0: "foo bar"}
        _ = v10
    }
    v11 := GenS2{F0: "\n\t"}
    if (0.1 >= 0.1) {
        for fi12 := 0; fi12 < 20; fi12++ {
            if true {
                v4 = (v4 & (^(v4)))
            }
            v4 = 1
            switch (fi12 >> ((803) & 63)) {
            case -1:
                fmt.Println(v3)
            case 32767:
                v13 := GenS3{F0: -580}
                _ = v13
            case 32768:
                for fi14 := 0; fi14 < 11; fi14++ {
                    v3 -= (v3 | v3)
                    v3 -= v3
                    if false {
                        v8 = append(v8, (v1 >> ((2) & 63)))
                        v3 += (v3 / (((v3 ^ v3)) | 1))
                    } else {
                        v2["a"] = 0
                        if true {
                            v15 := GenS4{F0: 47}
                            v16 := GenS3{F0: 42}
                            fmt.Println(v3)
                            _ = v15
                            _ = v16
                        } else {
                            var v17 int = -9223372036854775808
                            switch ("a" + "hello") {
                            case "\n\t":
                                fmt.Println(fi12, fi14, fi12)
                            default:
                                _ = len(v2)
                            }
                            var v18 int16 = -888
                            var v19 int32 = 42
                            v20 := GenS5{F0: (v4 | v4), F1: -1}
                            _ = v17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        _ = len(v2)
                    }
                    _ = fi14
                }
            }
            _ = fi12
        }
        v4 = 255
        var v21 uint = 691
        _ = v21
    }
    if (!false) {
        var v22 float64 = 0.0
        v4 = v4
        if false {
            for fi23 := 0; fi23 < 18; fi23++ {
                v24 := GenS6{F0: 0, F1: (2147483648 > 2147483647)}
                var v25 uint32 = 65535
                v25 -= (v25 & 358)
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        _ = v22
    } else {
        var v26 int8 = (v3 >> ((2) & 63))
        var v27 uint16 = 42
        var v28 float64 = -2.25
        _ = v26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
