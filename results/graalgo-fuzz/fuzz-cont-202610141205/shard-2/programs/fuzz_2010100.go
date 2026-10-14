package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 bool; F2 string }
type GenS7 struct { F0 int64; F1 int8 }
func genFn0() uint {
    if true {
        var v7 int = 9223372036854775807
        _ = v7
    } else {
        var v8 int8 = 127
        for fi9 := 0; fi9 < 5; fi9++ {
            v8 = 0
            _ = fi9
        }
        _ = v8
    }
    v10 := GenS1{F0: 821, F1: 2}
    switch "\n\t" {
    case "a":
        if true {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int32 = -1
                switch fi11 {
                case -198:
                    v12 = v12
                }
                _ = fi11
                _ = v12
            }
            switch 0 {
            case 0:
                for fi13 := 0; fi13 < 6; fi13++ {
                    fmt.Println(fi13)
                    var v14 uint8 = 2
                    for fi15 := 0; fi15 < 2; fi15++ {
                        v16 := GenS2{F0: 272, F1: 4294967295}
                        v14 = 94
                        _ = fi15
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
            case 32767:
                var v17 uint = 644
                _ = v17
            case 128:
                var v18 uint8 = 2
                _ = v18
            }
        } else {
            var v19 uint64 = 18446744073709551615
            _ = v19
        }
    default:
        var v20 float32 = 0.1
        _ = v20
    }
    var v21 int8 = -128
    v22 := GenS3{F0: -575, F1: 0}
    _ = v10
    _ = v21
    _ = v22
    return 0
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]bool)
    var v3 uint16 = 255
    var v4 bool = false
    v5 := []uint32{4294967295, 0}
    v6 := make(map[int]int)
    var v23 uint16 = (^(v3))
    switch "a" {
    case "":
        v24 := GenS4{F0: 0}
        _ = v24
    }
    switch 1 {
    case -884:
        var v25 uint8 = 0
        _ = v25
    case -700:
        for fi26 := 0; fi26 < 10; fi26++ {
            v27 := GenS5{F0: 255}
            for fi28 := 0; fi28 < 18; fi28++ {
                v3 = v23
                var v29 int = fi28
                v30 := GenS2{F0: 918, F1: 18446744073709551615}
                _ = fi28
                _ = v29
                _ = v30
            }
            var v31 int32 = 920
            _ = fi26
            _ = v27
            _ = v31
        }
    case -1:
        var v32 int = 9223372036854775807
        _ = v32
    default:
        v33 := GenS6{F0: 65535, F1: v4, F2: "a"}
        _ = v33
    }
    v4 = v4
    for fi34 := 0; fi34 < 19; fi34++ {
        for fi35 := 0; fi35 < 3; fi35++ {
            v36 := GenS7{F0: -682, F1: 42}
            v4 = v4
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    v23 -= 619
    v37 := GenS5{F0: 127}
    if (((((((1e10 * -2.25) / -2.25) / 0.1) * -2.25) / 1e10) + (float32(float64((1.5 / (-2.25 / 1.0)))) * 0.0)) < 1.5) {
        v23 = 823
    }
    fmt.Println(v4)
    if true {
        v1 -= (^(v1))
    }
    var v38 int64 = 42
    v39 := GenS3{F0: 127, F1: 651}
    if (!v4) {
        for fi40 := 0; fi40 < 5; fi40++ {
            switch "a" {
            case "\n\t":
                fmt.Println(v1, v3)
            }
            v41 := GenS6{F0: genFn0(), F1: (32767 == (v3 + 1)), F2: "\n\t"}
            _ = fi40
            _ = v41
        }
        var v42 string = "hello"
        _ = v42
    }
    var v43 int32 = 442
    switch 425 {
    case 9223372036854775807:
        fmt.Println(v4, v23, v1)
    case 128:
        fmt.Println(v3)
    case -380:
        fmt.Println(v38, v38, v43)
    }
    _ = len(v2)
    var v44 int32 = v43
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
