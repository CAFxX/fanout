package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int64; F2 string }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 255
    var v4 float64 = float64(3.14159)
    var v5 float64 = ((v4 * 1e10) / (float64(3.14159) - 0.0))
    var v6 uint16 = 65535
    var v7 bool = ((float32(((0.0 + 0.0) / v5)) * -2.25) <= 0.0)
    v8 := GenS1{F0: 18446744073709551615}
    v9 := GenS2{F0: (^(v3)), F1: 1, F2: "a"}
    v7 = (-9223372036854775808 > -9223372036854775808)
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 float64 = 0.1
            var v13 int8 = 127
            var v14 uint64 = 1000
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        v15 := GenS3{F0: ("a" + "foo bar"), F1: 4294967295}
        _ = fi10
        _ = v15
    }
    var v16 float64 = v4
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v4)
        var v18 float64 = v16
        v19 := GenS2{F0: (v3 & 127), F1: -155, F2: ""}
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
