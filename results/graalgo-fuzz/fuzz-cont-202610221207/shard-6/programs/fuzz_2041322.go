package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 bool }

func main() {
    v1 := []uint32{640}
    var v2 int64 = 9223372036854775807
    var v3 string = ("" + "hello")
    var v4 bool = false
    v4 = v4
    var v5 bool = v4
    v6 := GenS1{F0: (v2 | (^(v2)))}
    if len(v1) > 0 {
        v1[0] = 65535
    }
    var v7 string = "foo bar"
    var v8 int = -2147483648
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int = (fi9 / ((v8) | 1))
        var v11 int8 = -38
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint64 = 18446744073709551615
    v2 += -1
    fmt.Println(v12, v7, v7)
    switch (v8 >> ((2) & 63)) {
    case -1:
        var v13 uint16 = 0
        _ = v13
    case -9223372036854775808:
        _ = len(v1)
    default:
        _ = len(v1)
    }
    v14 := GenS2{F0: -2147483648}
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS2{F0: 42}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        if (2 < uint(v12)) {
            var v18 float64 = ((0.0 + (-2.25 - (1e10 / 0.1))) + 1.5)
            var v19 float64 = v18
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 16; fi20++ {
                if v5 {
                    fmt.Println(fi20, v4, v4)
                    var v21 bool = v5
                    if v5 {
                        fmt.Println(v5, v4)
                    } else {
                        var v22 int64 = (v2 ^ -1)
                        if false {
                            v23 := GenS1{F0: -502}
                            v24 := GenS1{F0: v22}
                            _ = v23
                            _ = v24
                        }
                        var v25 float64 = (float64(((((0.1 / 1.0) / (1.5 / float32(float64(-2.25)))) + 0.1) + -2.25)) + (3.14159 + 0.1))
                        _ = v22
                        _ = v25
                    }
                    _ = v21
                } else {
                    _ = len(v1)
                    var v26 uint64 = (v12 << ((65535) & 63))
                    switch v3 {
                    case "a":
                        fmt.Println(v26)
                    case "":
                        if (!(!true)) {
                            v26 += 18446744073709551615
                            fmt.Println(v5, v5)
                        }
                    case "\n\t":
                        _ = len(v1)
                    default:
                        switch 803 {
                        case 132:
                            var v27 uint = 4294967295
                            _ = v27
                        case -535:
                            fmt.Println(v8, fi17, fi17)
                        }
                        var v28 int8 = -1
                        _ = v28
                    }
                    _ = v26
                }
                _ = fi20
            }
        }
        v3 = v7
        _ = fi17
    }
    switch v8 {
    case -2147483648:
        v29 := GenS1{F0: (v2 - v2)}
        _ = v29
    case -685:
        if len(v1) > 0 {
            v1[0] = 363
        }
    default:
        if (-9223372036854775808 <= (v8 * v8)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        if true {
            var v31 uint8 = 42
            switch (fi30 / (((v8 << ((65535) & 63))) | 1)) {
            case -2147483648:
                v32 := GenS3{F0: 161, F1: 42, F2: v4}
                _ = v32
            }
            var v33 uint32 = 590
            _ = v31
            _ = v33
        } else {
            v34 := GenS3{F0: 42, F1: 393, F2: (!v5)}
            v7 = "foo bar"
            var v35 string = v3
            _ = v34
            _ = v35
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi30
    }
    v8 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
