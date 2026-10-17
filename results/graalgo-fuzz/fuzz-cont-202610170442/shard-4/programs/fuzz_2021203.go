package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    v4 := []string{"foo bar", "foo bar", (("" + ("" + (("a" + "\n\t") + ("hello" + "a")))) + (("" + "") + ""))}
    var v5 int64 = 42
    v6 := make(map[uint32]string)
    v5 += v5
    v7 := GenS1{F0: "", F1: 1, F2: 128}
    v8 := GenS2{F0: 568, F1: 65535, F2: 1}
    if false {
        fmt.Println(v5)
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int8 = 13
            switch (fi9 % (((fi9 >> ((65535) & 63))) | 1)) {
            case 9223372036854775807:
                var v11 int32 = 686
                _ = v11
            case 4294967295:
                v12 := GenS3{F0: true, F1: 77}
                _ = v12
            default:
                v13 := GenS4{F0: 255}
                _ = v13
            }
            var v14 uint8 = 128
            _ = fi9
            _ = v10
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS5{F0: 56, F1: 266, F2: 255}
            v5 -= v5
            _ = fi15
            _ = v16
        }
    }
    var v17 int32 = -957
    v5 = v5
    var v18 int16 = -250
    for fi19 := 0; fi19 < 2; fi19++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v20 uint16 = 164
        var v21 int8 = -3
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint64 = 4294967295
    v22 = (v22 % ((v22) | 1))
    fmt.Println(v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
