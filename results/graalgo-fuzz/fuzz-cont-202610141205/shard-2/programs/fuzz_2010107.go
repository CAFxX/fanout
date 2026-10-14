package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 bool }

func main() {
    var v1 uint = 20
    v2 := make(map[int]bool)
    var v3 uint32 = 399
    v4 := []string{}
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 float64 = 0.0
        v7 := GenS1{F0: false}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int8 = 1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            if false {
                v9 = v9
            }
        } else {
            switch "hello" {
            case "":
                switch (-(fi8)) {
                case -294:
                    v1 += v1
                case -503:
                    v4 = append(v4, "foo bar")
                case -436:
                    switch fi8 {
                    case 127:
                        var v10 uint = 350
                        _ = v10
                    }
                default:
                    _ = len(v2)
                }
            default:
                var v11 uint8 = 82
                _ = v11
            }
            v3 = 172
        }
        _ = fi8
        _ = v9
    }
    var v12 uint = (v1 % (((v1 | v1)) | 1))
    v13 := GenS2{F0: v1}
    if len(v4) > 0 {
        v4[0] = ("foo bar" + "foo bar")
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        _ = len(v4)
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 int32 = 2147483647
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS1{F0: true}
    v1 = (v12 >> ((v12) & 63))
    if (((((float32((3.14159 * 0.1)) - 0.1) + 0.1) / 3.14159) / float32((1.5 + 0.1))) < -2.25) {
        v1 = 4294967295
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := GenS2{F0: (v12 + v12)}
            fmt.Println(v1, v1, fi18)
            _ = fi18
            _ = v19
        }
    }
    var v20 uint32 = 2
    _ = len(v4)
    v21 := GenS3{F0: -1, F1: true}
    v20 -= (v3 >> ((1) & 63))
    v1 = v12
    if (uint32(v12) <= (v20 ^ v3)) {
        v12 = (v1 & v1)
    } else {
        delete(v2, -2147483648)
        switch ("" + "foo bar") {
        case "foo bar":
            v22 := GenS1{F0: false}
            _ = v22
        case "\n\t":
            var v23 int8 = -115
            _ = v23
        default:
            _ = len(v4)
        }
        var v24 uint = (v12 & v12)
        _ = v24
    }
    var v25 float64 = (1e10 - 1.5)
    if false {
        if false {
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 int = fi26
                var v28 bool = false
                var v29 uint = 741
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            _ = len(v2)
        }
        fmt.Println(v25)
    } else {
        _ = len(v4)
        var v30 uint32 = v3
        v4 = append(v4, "foo bar")
        _ = v30
    }
    var v31 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v31
    _ = v4
}
