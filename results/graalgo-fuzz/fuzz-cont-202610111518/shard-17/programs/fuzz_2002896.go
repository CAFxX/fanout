package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 uint; F2 bool }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint64 = 209
    v2 := []uint32{2}
    v3 := make(map[int]bool)
    v4 := GenS1{F0: -863, F1: 4294967295}
    v1 -= (v1 + v1)
    var v5 string = "\n\t"
    if (260 <= 128) {
        _ = len(v2)
        if (778 != 32767) {
            var v6 bool = true
            v7 := GenS2{F0: 189, F1: 332}
            v8 := GenS2{F0: (^(v1)), F1: 639}
            _ = v6
            _ = v7
            _ = v8
        }
        switch ("\n\t" + ("" + v5)) {
        case "":
            if true {
                for fi9 := 0; fi9 < 13; fi9++ {
                    v5 = ""
                    v1 += (^(v1))
                    if ((255 <= 202) && (-1 != 1)) {
                        var v10 int32 = 2147483647
                        delete(v3, (fi9 >> ((580) & 63)))
                        v11 := GenS3{F0: 87, F1: 1, F2: (!((!(!true)) && false))}
                        _ = v10
                        _ = v11
                    }
                    _ = fi9
                }
                v1 = v1
                fmt.Println(v1)
            }
        default:
            switch 873 {
            case 42:
                v5 = ("foo bar" + v5)
            case -38:
                v12 := GenS4{F0: 866, F1: 595}
                _ = v12
            case -2147483648:
                for fi13 := 0; fi13 < 8; fi13++ {
                    v5 = (v5 + (v5 + v5))
                    _ = fi13
                }
            default:
                _ = len(v3)
            }
        }
    } else {
        v1 = (v1 + (v1 >> ((695) & 63)))
        delete(v3, 42)
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    v1 = 0
    var v14 bool = false
    switch -9223372036854775808 {
    case 865:
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(fi15, v5)
            v14 = true
            _ = fi15
        }
    case 2147483647:
        v16 := GenS5{F0: -379}
        _ = v16
    }
    if ((v14 && (((0.0 <= (((((1e10 - 3.14159) + -2.25) + 0.0) / 3.14159) / 1e10)) || true) && (-97 < 46))) || (true && (1 < (v1 % ((v1) | 1))))) {
        var v17 int32 = -132
        var v18 int16 = -326
        _ = v17
        _ = v18
    }
    var v19 uint64 = (v1 - v1)
    if len(v2) > 0 {
        v2[0] = 42
    }
    _ = len(v2)
    _ = len(v3)
    v2 = append(v2, 65535)
    if true {
        v5 = "foo bar"
        for fi20 := 0; fi20 < 17; fi20++ {
            fmt.Println(v5)
            var v21 uint16 = 255
            _ = fi20
            _ = v21
        }
        v22 := v3[4294967295]
        _ = v22
    } else {
        fmt.Println(v5, v14, v14)
    }
    switch -2147483648 {
    case 32768:
        _ = len(v2)
    }
    v19 -= v19
    var v23 int64 = 42
    var v24 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
