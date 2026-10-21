package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := []string{}
    var v2 uint8 = 2
    v3 := make(map[int]bool)
    var v4 uint64 = 42
    var v5 int64 = -9223372036854775808
    if false {
        var v6 uint = 0
        var v7 int32 = 42
        v6 = v6
        _ = v6
        _ = v7
    } else {
        fmt.Println(v2, v2, v5)
        if true {
            switch 32767 {
            case 128:
                v5 -= 9223372036854775807
            }
        } else {
            v8 := GenS1{F0: 103, F1: 255, F2: (^(v2))}
            v2 += (v2 - (v2 << ((42) & 63)))
            if false {
                v2 = v2
            } else {
                v9 := GenS1{F0: 127, F1: 1, F2: (^(v2))}
                _ = v9
            }
            _ = v8
        }
    }
    if (false && false) {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: 51, F1: 1, F2: (v2 | v2)}
            v2 -= v2
            switch "foo bar" {
            case "\n\t":
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 uint32 = uint32(v2)
                    v14 := GenS1{F0: 0, F1: 59, F2: v2}
                    v15 := GenS1{F0: 100, F1: 1, F2: (v2 % ((v2) | 1))}
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            default:
                var v16 uint32 = 1
                _ = v16
            }
            _ = fi10
            _ = v11
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 uint64 = 18446744073709551615
            _ = fi17
            _ = v18
        }
        switch "foo bar" {
        case "":
            var v19 int = 1
            _ = v19
        case "hello":
            switch 4294967295 {
            case 42:
                fmt.Println(v4, v4)
            case 857:
                for fi20 := 0; fi20 < 14; fi20++ {
                    delete(v3, fi20)
                    _ = fi20
                }
            }
        }
    }
    v4 -= 42
    fmt.Println(v4)
    if false {
        var v21 int8 = 1
        if false {
            v22 := GenS1{F0: 100, F1: 42, F2: v2}
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 int16 = 124
                _ = fi23
                _ = v24
            }
            delete(v3, 960)
            _ = v22
        } else {
            v25 := GenS1{F0: 100, F1: 184, F2: v2}
            v26 := GenS1{F0: (v21 << ((2) & 63)), F1: 701, F2: 128}
            _ = v25
            _ = v26
        }
        v27 := GenS2{F0: (v5 % ((v5) | 1)), F1: false, F2: 255}
        _ = v21
        _ = v27
    } else {
        delete(v3, 0)
    }
    if ("\n\t" <= "foo bar") {
        v28 := GenS1{F0: 1, F1: 706, F2: v2}
        _ = v28
    }
    v3[32767] = (!(-2147483648 != 809))
    v4 -= (^(v4))
    if true {
        v5 = (v5 / ((-158) | 1))
        v2 = uint8(v4)
        v2 -= v2
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 8; fi30++ {
            v2 = (v2 ^ v2)
            v31 := GenS3{F0: 127}
            fmt.Println(fi29, v5, fi30)
            var v32 int16 = 32767
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
