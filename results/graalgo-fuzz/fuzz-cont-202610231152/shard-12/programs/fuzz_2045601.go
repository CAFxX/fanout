package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int64; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 string; F2 uint16 }

func main() {
    v1 := []string{("a" + ("\n\t" + "\n\t")), ("foo bar" + ""), ("" + "\n\t"), "hello"}
    v2 := make(map[int]uint32)
    var v3 string = "foo bar"
    v4 := make(map[string]string)
    var v5 uint = 65535
    if false {
        fmt.Println(v3)
    } else {
        v5 = v5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint64 = 42
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: (2 > 127), F1: 65535, F2: 18446744073709551615}
        if (false && ((127 >= 42) && (true || true))) {
            switch -405 {
            case 128:
                v3 = v3
            case -752:
                v10 := GenS2{F0: 323}
                _ = v10
            case -302:
                if (!false) {
                    v11 := GenS3{F0: 255, F1: 0, F2: 32767}
                    v5 = v5
                    v5 += (v5 >> (((v5 - 603)) & 63))
                    _ = v11
                }
            }
        } else {
            v3 = (((v3 + v3) + v3) + v3)
            var v12 uint16 = 226
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS3{F0: 27, F1: -1, F2: (fi13 & fi13)}
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v3)
    v5 = (^(v5))
    switch v3 {
    case "":
        switch "foo bar" {
        case "":
            fmt.Println(v3, v5)
        case "a":
            fmt.Println(v3, v5)
        }
    case "a":
        for fi15 := 0; fi15 < 16; fi15++ {
            _ = len(v4)
            _ = fi15
        }
    }
    v5 = v5
    if (true || true) {
        v16 := GenS4{F0: 1}
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := GenS1{F0: true, F1: 2147483647, F2: 42}
            switch v3 {
            case "hello":
                v2[(^(fi17))] = 42
            default:
                for fi19 := 0; fi19 < 18; fi19++ {
                    v3 = (v3 + ("\n\t" + "foo bar"))
                    fmt.Println(fi19, fi19, v3)
                    _ = fi19
                }
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        if false {
            fmt.Println(v3, v3)
            var v20 bool = true
            _ = v20
        } else {
            var v21 uint = (v5 / ((v5) | 1))
            var v22 bool = false
            _ = v21
            _ = v22
        }
    }
    var v23 float32 = 1.5
    v24 := GenS5{F0: 42, F1: v3, F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
