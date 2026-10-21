package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int8 = 100
    var v2 uint = 255
    v3 := []string{}
    var v4 uint16 = 633
    v5 := []uint32{485, 42}
    v6 := make(map[uint32]uint32)
    var v7 uint8 = 0
    var v8 float64 = (float64(((-2.25 - 0.1) / 3.14159)) - 1.5)
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int16 = 42
        _ = fi9
        _ = v10
    }
    if false {
        var v11 uint16 = (v4 & (v4 % (((^(v4))) | 1)))
        v8 -= v8
        _ = v11
    }
    v12 := GenS1{F0: v7, F1: 2147483647, F2: -1}
    var v13 int16 = 345
    v14 := GenS1{F0: (v7 & v7), F1: 667, F2: v13}
    v1 += (v1 % (((v1 - (^(v1)))) | 1))
    _ = len(v3)
    var v15 uint16 = (v4 / (((v4 ^ 0)) | 1))
    v15 += (v4 + (^(v4)))
    v8 = 1.5
    v16 := GenS1{F0: (^(v7)), F1: -1, F2: v13}
    v17 := GenS1{F0: (v7 % ((v7) | 1)), F1: 1, F2: v13}
    if true {
        v7 += (^(v7))
        v1 -= v1
    } else {
        fmt.Println(v15)
        v18 := GenS2{F0: "a", F1: 1, F2: 42}
        _ = v18
    }
    if (true || true) {
        switch "foo bar" {
        case "":
            delete(v6, 4294967295)
        default:
            var v19 uint16 = (^(v15))
            _ = v19
        }
    }
    if (!(v15 <= v15)) {
        var v20 uint8 = v7
        if false {
            switch -1 {
            case 883:
                for fi21 := 0; fi21 < 16; fi21++ {
                    if true {
                        var v22 uint8 = v20
                        _ = len(v6)
                        for fi23 := 0; fi23 < 6; fi23++ {
                            var v24 int8 = v1
                            v1 = v24
                            delete(v6, 797)
                            var v25 int32 = -249
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    }
                    delete(v6, 2147483647)
                    if (!(!(1 < (v2 & 614)))) {
                        var v26 uint8 = v20
                        fmt.Println(v13)
                        v20 += 127
                        _ = v26
                    } else {
                        v27 := GenS3{F0: 0}
                        _ = v27
                    }
                    _ = fi21
                }
            case -771:
                for fi28 := 0; fi28 < 12; fi28++ {
                    v29 := GenS3{F0: 235}
                    var v30 int = (fi28 / ((fi28) | 1))
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
        } else {
            for fi31 := 0; fi31 < 10; fi31++ {
                v32 := v6[42]
                _ = fi31
                _ = v32
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
