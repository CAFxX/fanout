package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int64; F2 int8 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 bool }

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{2147483648}
    var v3 int16 = -159
    var v4 float64 = (v1 * (((float64(((float32(v1) - 0.0) / 1.0)) - v1) * -2.25) + (0.0 * v1)))
    var v5 string = ("" + "")
    v6 := GenS1{F0: 127, F1: 2147483647, F2: 731}
    switch "" {
    case "hello":
        _ = len(v2)
    case "":
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS2{F0: v3, F1: 0}
            v9 := GenS3{F0: v5}
            _ = fi7
            _ = v8
            _ = v9
        }
    case "a":
        var v10 string = (v5 + "hello")
        _ = v10
    }
    _ = len(v2)
    switch 42 {
    case -479:
        var v11 uint = 18446744073709551615
        _ = v11
    case 0:
        if false {
            fmt.Println(v4, v3)
        }
    }
    var v12 bool = (42 != 0)
    for fi13 := 0; fi13 < 12; fi13++ {
        switch "a" {
        case "a":
            var v14 int = (fi13 + (fi13 + (fi13 + (^(fi13)))))
            _ = v14
        case "":
            _ = len(v2)
        case "\n\t":
            var v15 int64 = 842
            _ = v15
        default:
            if len(v2) > 0 {
                v2[0] = 42
            }
        }
        var v16 int32 = 2147483647
        _ = fi13
        _ = v16
    }
    var v17 int8 = -53
    fmt.Println(v1, v1, v12)
    if v12 {
        v12 = v12
        for fi18 := 0; fi18 < 11; fi18++ {
            if (65535 < 2147483648) {
                v19 := GenS4{F0: 42, F1: 111, F2: 128}
                var v20 uint16 = 2
                for fi21 := 0; fi21 < 19; fi21++ {
                    if (0 < 0) {
                        var v22 uint16 = v20
                        var v23 float64 = v1
                        _ = v22
                        _ = v23
                    } else {
                        var v24 float64 = (v4 + ((float64((1.5 * (((0.1 / float32(-2.25)) / (1e10 - (3.14159 / 0.1))) + -2.25))) * v1) - 1e10))
                        switch 9223372036854775807 {
                        case 128:
                            v4 = v24
                        }
                        v17 = -32
                        _ = v24
                    }
                    _ = fi21
                }
                _ = v19
                _ = v20
            }
            switch fi18 {
            case 687:
                switch -903 {
                case 128:
                    if len(v2) > 0 {
                        v2[0] = 673
                    }
                }
            case -511:
                fmt.Println(v4, v1)
            default:
                v5 = v5
            }
            _ = fi18
        }
    }
    if v12 {
        var v25 float32 = 0.1
        var v26 int32 = -478
        var v27 uint64 = 4294967295
        _ = v25
        _ = v26
        _ = v27
    } else {
        _ = len(v2)
        var v28 int64 = -9223372036854775808
        for fi29 := 0; fi29 < 5; fi29++ {
            fmt.Println(v17, v4)
            _ = fi29
        }
        _ = v28
    }
    switch 921 {
    case 2147483647:
        switch (v5 + "") {
        case "foo bar":
            if v12 {
                for fi30 := 0; fi30 < 5; fi30++ {
                    v3 += v3
                    var v31 int = 127
                    v32 := GenS5{F0: 0, F1: -1, F2: -128}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        }
    case -102:
        v33 := GenS6{F0: 127, F1: 42}
        _ = v33
    default:
        var v34 int32 = -2147483648
        _ = v34
    }
    v35 := GenS7{F0: false}
    v12 = false
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
