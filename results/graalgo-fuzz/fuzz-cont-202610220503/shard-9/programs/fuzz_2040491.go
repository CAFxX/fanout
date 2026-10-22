package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint8 = 0
    var v2 uint8 = v1
    var v3 uint64 = 1
    v4 := []string{}
    var v5 float32 = (-2.25 / 0.1)
    v6 := GenS1{F0: 0}
    v3 = (^(v3))
    var v7 uint = 255
    v5 += v5
    var v8 int32 = 0
    var v9 uint16 = 723
    var v10 uint = (v7 >> (((v7 * (v7 ^ v7))) & 63))
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v4 = append(v4, "a")
        switch (fi11 ^ -231) {
        case 128:
            for fi12 := 0; fi12 < 12; fi12++ {
                if (1 != 1000) {
                    var v13 int8 = -128
                    _ = v13
                } else {
                    switch ("" + ("hello" + "\n\t")) {
                    case "foo bar":
                        v14 := GenS1{F0: 1}
                        v1 = (v1 / ((42) | 1))
                        v15 := GenS1{F0: -1}
                        _ = v14
                        _ = v15
                    case "":
                        var v16 uint8 = 42
                        _ = v16
                    default:
                        fmt.Println(v5)
                    }
                    v1 += (^(v1))
                    var v17 int64 = 42
                    v8 -= v8
                    v10 = (^(v10))
                    switch ("" + (("a" + ("hello" + ("" + "foo bar"))) + "\n\t")) {
                    case "":
                        v9 = 42
                    case "a":
                        _ = len(v4)
                        var v18 int16 = 42
                        v17 -= 817
                        _ = v18
                    case "\n\t":
                        v19 := GenS2{F0: 0, F1: true}
                        _ = v19
                    default:
                        if (true && (true || true)) {
                            if ((v1 * v1) == v2) {
                                var v20 int64 = (v17 / (((v17 | v17)) | 1))
                                _ = v20
                            }
                        } else {
                            var v21 uint64 = v3
                            var v22 bool = (!false)
                            _ = v21
                            _ = v22
                        }
                        var v23 int16 = -32768
                        fmt.Println(v23, v17, v8)
                        _ = v23
                    }
                    v3 = 4294967295
                    _ = v17
                }
                v2 -= v2
                _ = fi12
            }
        case 384:
            v9 -= v9
        default:
            switch "" {
            case "\n\t":
                var v24 uint8 = (v1 ^ v1)
                _ = v24
            case "a":
                fmt.Println(fi11)
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi11
    }
    v7 -= 805
    for fi25 := 0; fi25 < 16; fi25++ {
        v9 = (v9 / (((v9 * 0)) | 1))
        switch fi25 {
        case -200:
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 uint = v7
                _ = fi26
                _ = v27
            }
        case -726:
            v28 := GenS3{F0: v1, F1: (!false)}
            _ = v28
        }
        var v29 int16 = -815
        _ = fi25
        _ = v29
    }
    v8 = -1
    fmt.Println(v2)
    fmt.Println(v9, v9, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
