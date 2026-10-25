package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int; F2 int64 }

func main() {
    var v1 uint8 = 127
    var v2 int16 = 1
    v3 := []uint32{42, 627, 1, 2147483648}
    var v4 uint8 = (v1 >> ((255) & 63))
    var v5 int64 = 1
    fmt.Println(v2)
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint32 = 65535
        _ = fi6
        _ = v7
    }
    var v8 int16 = v2
    v8 += (v8 ^ v8)
    fmt.Println(v1)
    v2 += (v8 - v8)
    v8 = (v2 % (((v8 % (((v8 / ((-522) | 1))) | 1))) | 1))
    switch ("foo bar" + "hello") {
    case "foo bar":
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int16 = v2
            _ = fi9
            _ = v10
        }
    case "\n\t":
        var v11 bool = false
        _ = v11
    case "":
        var v12 bool = ((!false) && ((-(v8)) <= v8))
        _ = v12
    default:
        for fi13 := 0; fi13 < 8; fi13++ {
            fmt.Println(v4, v8)
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        switch fi14 {
        case -498:
            switch fi14 {
            case 128:
                var v15 bool = false
                _ = v15
            case 1:
                var v16 uint16 = 676
                _ = v16
            default:
                if ((!false) || false) {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v1 = (^(v1))
                        v18 := GenS1{F0: 4294967295, F1: uint32(v4), F2: 82}
                        _ = fi17
                        _ = v18
                    }
                    fmt.Println(v8, v2)
                } else {
                    v19 := GenS1{F0: 119, F1: 728, F2: 2}
                    switch (fi14 | (fi14 ^ fi14)) {
                    case -2147483648:
                        fmt.Println(v2)
                    case 128:
                        if len(v3) > 0 {
                            v3[0] = 132
                        }
                    case 86:
                        var v20 int8 = 2
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 6; fi21++ {
                        if false {
                            _ = len(v3)
                            v22 := GenS2{F0: -2147483648}
                            v8 = v2
                            switch ("foo bar" + "hello") {
                            case "hello":
                                fmt.Println(v1, fi21)
                            }
                            _ = v22
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v1)
                        _ = fi21
                    }
                    _ = v19
                }
            }
        default:
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint16 = 386
                _ = fi23
                _ = v24
            }
        }
        _ = fi14
    }
    v2 = int16(v4)
    v25 := GenS2{F0: 42}
    switch 128 {
    case -301:
        v26 := GenS3{F0: ("\n\t" + "foo bar"), F1: 1, F2: 9223372036854775807}
        _ = v26
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 uint8 = 33
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        for fi30 := 0; fi30 < 8; fi30++ {
            if false {
                var v31 float32 = 1e10
                _ = v31
            }
            _ = fi30
        }
        _ = len(v3)
        var v32 int32 = 0
        _ = fi29
        _ = v32
    }
    var v33 int16 = (v8 - v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v8
}
