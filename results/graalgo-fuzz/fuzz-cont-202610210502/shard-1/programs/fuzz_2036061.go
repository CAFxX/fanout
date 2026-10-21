package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 int }
func genFn0(p0 bool) int64 {
    var v5 uint8 = 1
    v6 := GenS1{F0: 788, F1: 2147483647}
    if (18446744073709551615 == 634) {
        var v7 uint = 2
        v8 := []int{9223372036854775807, 127, 127, 2147483647}
        _ = v7
        _ = v8
    }
    if (42 < 779) {
        v9 := GenS2{F0: "foo bar", F1: 923, F2: 0}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint64 = 42
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 uint8) uint16 {
    var v12 bool = ("" <= ("\n\t" + ((("" + "foo bar") + (("foo bar" + "a") + "")) + "\n\t")))
    switch ((("hello" + "a") + ("\n\t" + "a")) + "\n\t") {
    case "hello":
        var v13 int8 = 70
        _ = v13
    default:
        var v14 bool = v12
        _ = v14
    }
    _ = p0
    _ = v12
    return 751
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]string)
    var v3 uint32 = 459
    var v4 uint64 = 458
    var v15 uint64 = (v4 * 42)
    v3 = v3
    v3 = 490
    fmt.Println(v3, v1)
    var v16 int64 = -1
    _ = len(v2)
    for fi17 := 0; fi17 < 12; fi17++ {
        v16 = v16
        for fi18 := 0; fi18 < 6; fi18++ {
            fmt.Println(v3, v15, fi17)
            v3 = (^(v1))
            _ = fi18
        }
        _ = fi17
    }
    v19 := GenS3{F0: (v3 << ((2) & 63)), F1: 1, F2: 128}
    v2[530] = ""
    v15 = (v15 / ((v15) | 1))
    switch 0 {
    case -9223372036854775808:
        v20 := v2[520]
        _ = v20
    case 42:
        fmt.Println(v16, v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
