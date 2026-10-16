package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int }

func main() {
    var v1 float32 = (3.14159 - 3.14159)
    var v2 float32 = 0.0
    v3 := []string{""}
    v4 := make(map[uint32]int)
    if (42 != -128) {
        var v5 int32 = 0
        var v6 int = -1
        _ = v5
        _ = v6
    }
    var v7 int64 = 42
    var v8 float32 = ((float32(1.5) / ((float32(0.0) + v1) - v2)) / (-2.25 * 1e10))
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 string = ""
        var v11 int64 = 42
        switch (fi9 ^ (fi9 >> ((255) & 63))) {
        case 1:
            fmt.Println(v1, v11)
        case 42:
            var v12 float64 = float64(1e10)
            _ = v12
        case 785:
            if (false && (!true)) {
                v13 := GenS1{F0: 949, F1: 2147483647, F2: 128}
                _ = v13
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v7 += v7
    if false {
        if (!true) {
            v2 = v8
            switch 32768 {
            case -9223372036854775808:
                v7 = -681
            case 32768:
                if len(v3) > 0 {
                    v3[0] = ""
                }
            default:
                var v14 uint = 65535
                _ = v14
            }
        } else {
            if (42 == 880) {
                if (!(1 <= 297)) {
                    if (255 < 0) {
                        for fi15 := 0; fi15 < 13; fi15++ {
                            fmt.Println(v2)
                            _ = len(v3)
                            _ = fi15
                        }
                        v7 += 9223372036854775807
                    }
                }
            }
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 int32 = -1
            _ = fi16
            _ = v17
        }
    } else {
        fmt.Println(v2, v8)
        switch -2147483648 {
        case 2147483647:
            switch -2147483648 {
            case 1:
                var v18 int16 = -1
                _ = v18
            case 127:
                v19 := GenS1{F0: 255, F1: 986, F2: 4294967295}
                _ = v19
            case 32768:
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi20 := 0; fi20 < 19; fi20++ {
                        switch ("" + "hello") {
                        case "hello":
                            var v21 int16 = 991
                            _ = v21
                        case "a":
                            var v22 float32 = ((v1 * 1e10) * 0.0)
                            _ = v22
                        }
                        _ = len(v3)
                        _ = fi20
                    }
                }
            }
        case -290:
            if false {
                v1 = 0.1
            } else {
                var v23 int8 = 76
                _ = v23
            }
        }
        var v24 int = -1
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        switch (fi25 + 1) {
        case 312:
            for fi26 := 0; fi26 < 17; fi26++ {
                v27 := GenS1{F0: 212, F1: 4294967295, F2: fi26}
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                var v28 bool = true
                var v29 uint = 277
                v30 := v4[42]
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
                _ = v30
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v2 += (0.1 - 1e10)
        _ = fi25
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        v4[857] = fi31
        v32 := GenS1{F0: 780, F1: 42, F2: (fi31 & fi31)}
        _ = fi31
        _ = v32
    }
    v2 += v1
    v8 += ((1e10 + (0.1 / v1)) + v8)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
