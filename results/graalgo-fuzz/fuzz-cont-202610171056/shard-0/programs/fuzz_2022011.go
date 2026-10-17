package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 42
    var v3 float32 = 0.0
    v4 := make(map[uint32]uint32)
    switch 4294967295 {
    case -9223372036854775808:
        v2 = v2
    case 2147483647:
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 int32 = -2147483648
            _ = fi5
            _ = v6
        }
    case 4294967295:
        if (0.1 < 1e10) {
            var v7 uint16 = 811
            _ = v7
        } else {
            v8 := GenS1{F0: 945, F1: (0 != 0)}
            var v9 int8 = 42
            v2 = (v2 % ((int16(v9)) | 1))
            _ = v8
            _ = v9
        }
    default:
        var v10 bool = (!false)
        _ = v10
    }
    v2 += (^(v2))
    v4[2] = 481
    v11 := GenS2{F0: 42, F1: "\n\t"}
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int64 = -9223372036854775808
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 int32 = 1
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v16 := GenS3{F0: ("foo bar" + "a")}
    if ("a" > ("a" + "foo bar")) {
        if (2 != 65535) {
            var v17 uint8 = 1
            switch ("foo bar" + ("foo bar" + "a")) {
            case "":
                fmt.Println(v17, v2)
            default:
                var v18 uint8 = v17
                _ = v18
            }
            _ = v17
        }
        var v19 int16 = -32768
        _ = v19
    } else {
        _ = len(v1)
    }
    v20 := GenS2{F0: 1, F1: ("\n\t" + ("\n\t" + "a"))}
    v2 = v2
    for fi21 := 0; fi21 < 1; fi21++ {
        if false {
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 uint64 = 657
                _ = fi22
                _ = v23
            }
            v3 += 0.1
            var v24 uint64 = 18446744073709551615
            _ = v24
        }
        var v25 uint64 = 4294967295
        v26 := GenS4{F0: 42, F1: (!((v25 >> ((255) & 63)) != 0))}
        _ = fi21
        _ = v25
        _ = v26
    }
    if false {
        var v27 float64 = 1.5
        for fi28 := 0; fi28 < 17; fi28++ {
            v29 := GenS3{F0: "foo bar"}
            _ = fi28
            _ = v29
        }
        v2 -= 1
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
