package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 int16; F1 bool }

func main() {
    v1 := []uint32{1, 1, 42, 994}
    v2 := []uint32{615, 592}
    var v3 int64 = 1
    for fi4 := 0; fi4 < 3; fi4++ {
        if (!true) {
            fmt.Println(v3, v3)
            v3 += v3
        } else {
            fmt.Println(fi4, fi4)
            switch "hello" {
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "\n\t":
                var v5 uint64 = 4294967295
                _ = v5
            case "hello":
                _ = len(v2)
            default:
                v3 += v3
            }
        }
        _ = fi4
    }
    v3 = v3
    fmt.Println(v3)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 float32 = float32((1.5 - 0.1))
        v1 = append(v1, 42)
        _ = fi6
        _ = v7
    }
    v3 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    v8 := GenS1{F0: true, F1: false, F2: 2147483647}
    var v9 uint64 = 18446744073709551615
    v9 = (v9 - 0)
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    switch 0 {
    case 374:
        var v10 int16 = 1000
        _ = v10
    case -866:
        var v11 uint = 350
        _ = v11
    case 4294967295:
        _ = len(v1)
    default:
        var v12 uint8 = 171
        _ = v12
    }
    if false {
        if false {
            var v13 int16 = 32767
            _ = len(v1)
            _ = v13
        } else {
            v14 := GenS1{F0: false, F1: true, F2: -1}
            fmt.Println(v3, v3)
            v3 = 118
            _ = v14
        }
        var v15 uint16 = 65535
        _ = v15
    }
    v9 -= (v9 | (v9 & v9))
    v3 -= (v3 << ((4294967295) & 63))
    var v16 int = -18
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS1{F0: true, F1: (((uint8(v9) != 2) && false) && (0 < 42)), F2: -2147483648}
        switch "foo bar" {
        case "\n\t":
            var v19 float32 = (0.1 * ((0.1 - 3.14159) + 0.0))
            _ = v19
        default:
            if (!true) {
                switch ("a" + "foo bar") {
                case "foo bar":
                    switch "" {
                    case "":
                        var v20 uint = 65535
                        _ = v20
                    }
                }
                switch (-(fi17)) {
                case 128:
                    v21 := GenS1{F0: true, F1: (!((float32(0.1) + (1.5 / 1.0)) == 0.1)), F2: 0}
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 9; fi22++ {
                        if false {
                            v1 = append(v1, 2147483647)
                            var v23 uint16 = 975
                            _ = v23
                        } else {
                            var v24 int64 = -9223372036854775808
                            v3 += v24
                            _ = v24
                        }
                        _ = fi22
                    }
                }
                if false {
                    var v25 string = ((((("foo bar" + "a") + (("foo bar" + "\n\t") + "\n\t")) + (("" + "foo bar") + "\n\t")) + "hello") + "\n\t")
                    v26 := GenS2{F0: -32768, F1: false}
                    v27 := GenS1{F0: true, F1: true, F2: 1}
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                    v9 -= 2
                    var v28 uint8 = 255
                    _ = v28
                }
            }
        }
        _ = fi17
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
