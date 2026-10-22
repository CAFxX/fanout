package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
func genFn0() uint64 {
    var v7 uint = 0
    v8 := GenS1{F0: 0, F1: 0}
    _ = v7
    _ = v8
    return 384
}
func genFn1() string {
    v9 := GenS1{F0: -128, F1: 19}
    v10 := GenS1{F0: 0, F1: -128}
    if false {
        var v11 uint32 = 27
        if (!true) {
            var v12 int32 = 1
            _ = v12
        }
        _ = v11
    } else {
        v13 := GenS1{F0: -1, F1: 100}
        _ = v13
    }
    _ = v10
    _ = v9
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -128
    v3 := make(map[string]int)
    var v4 uint = 42
    v5 := []string{"foo bar", "foo bar", "a", "\n\t"}
    var v6 uint = (v4 / ((v4) | 1))
    _ = len(v5)
    for fi14 := 0; fi14 < 20; fi14++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi14
    }
    var v15 int64 = -9223372036854775808
    v4 = (v6 / (((v4 >> (((v6 - (v4 * (^(v4))))) & 63))) | 1))
    var v16 bool = ("a" != "")
    if v16 {
        v15 -= (v15 & (v15 - v15))
    }
    switch "\n\t" {
    case "\n\t":
        v1[2147483648] = 42
    case "hello":
        if (!((4294967295 < 2) && (true && v16))) {
            v4 = v6
            v17 := GenS1{F0: v2, F1: 66}
            _ = v17
        }
    case "a":
        if (true && v16) {
            v18 := GenS2{F0: 887}
            switch 127 {
            case 414:
                v19 := GenS2{F0: uint64(v6)}
                _ = v19
            case 0:
                v6 = 667
            }
            fmt.Println(v2, v15)
            _ = v18
        } else {
            v20 := GenS1{F0: v2, F1: (^(v2))}
            v21 := GenS1{F0: (v2 / (((v2 ^ v2)) | 1)), F1: v2}
            if (v16 || (!(("\n\t" + "foo bar") >= ""))) {
                v22 := GenS3{F0: (v4 ^ v4)}
                var v23 uint = v4
                _ = v22
                _ = v23
            }
            _ = v20
            _ = v21
        }
    }
    v6 = (^(v6))
    delete(v3, "")
    v24 := GenS2{F0: 0}
    v1[0] = 152
    var v25 uint16 = 65535
    v6 = (v6 >> (((v4 | v4)) & 63))
    v25 -= (v25 + 32767)
    var v26 int = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
