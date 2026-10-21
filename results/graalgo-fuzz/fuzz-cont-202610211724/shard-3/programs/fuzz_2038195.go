package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]uint32)
    var v3 uint = 226
    v4 := []uint32{495}
    v5 := []string{"foo bar"}
    v6 := make(map[uint32]string)
    v1 = (-(v1))
    switch "a" {
    case "a":
        var v7 int64 = 1
        _ = v7
    case "\n\t":
        fmt.Println(v3)
    }
    switch "hello" {
    case "":
        var v8 uint32 = 4294967295
        _ = v8
    case "foo bar":
        var v9 int64 = 1
        _ = v9
    default:
        var v10 bool = true
        _ = v10
    }
    var v11 uint8 = 0
    var v12 uint32 = 4294967295
    fmt.Println(v11, v11, v11)
    v4 = append(v4, (v12 % ((uint32(v11)) | 1)))
    fmt.Println(v1, v11)
    v11 += 44
    switch 9223372036854775807 {
    case 42:
        if false {
            for fi13 := 0; fi13 < 18; fi13++ {
                v3 = 0
                var v14 int32 = 1
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                if (v3 == v3) {
                    v16 := GenS1{F0: int16(v3), F1: 100, F2: -32768}
                    _ = v16
                }
                _ = fi15
            }
            var v17 uint64 = 2
            _ = v17
        } else {
            if ((1.5 == (3.14159 / ((0.1 + (-2.25 / -2.25)) * 1.5))) && (true && true)) {
                delete(v6, 1)
            } else {
                var v18 float64 = ((0.1 / 1e10) + 3.14159)
                _ = v18
            }
            v19 := GenS2{F0: (v11 % (((v11 - (v11 - (v11 << (((v3 & v3)) & 63))))) | 1)), F1: v3}
            _ = v19
        }
    case -9223372036854775808:
        var v20 int8 = 100
        _ = v20
    case 4294967295:
        for fi21 := 0; fi21 < 2; fi21++ {
            switch fi21 {
            case 9223372036854775807:
                v1 = (v1 + v1)
            case -62:
                if (false || (1 < 32767)) {
                    var v22 float64 = 0.0
                    _ = len(v2)
                    _ = v22
                }
            case -560:
                switch -9223372036854775808 {
                case 128:
                    v11 = 49
                case 0:
                    _ = len(v2)
                default:
                    v12 = (^(v12))
                }
            }
            _ = fi21
        }
    }
    v23 := GenS3{F0: (v11 & 21)}
    for fi24 := 0; fi24 < 9; fi24++ {
        v12 = v12
        if (-589 != -2147483648) {
            var v25 uint64 = 42
            switch (("" + ("" + "foo bar")) + ("hello" + (("hello" + ("\n\t" + "a")) + "hello"))) {
            case "hello":
                v3 += (v3 + v3)
            }
            v26 := GenS2{F0: v11, F1: 42}
            _ = v25
            _ = v26
        }
        _ = fi24
    }
    v27 := GenS3{F0: (^(v11))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
