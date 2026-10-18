package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint64; F1 string }

func main() {
    var v1 float64 = 0.0
    var v2 float32 = 0.1
    v3 := make(map[uint32]bool)
    var v4 uint32 = 0
    v1 -= 1e10
    v5 := GenS1{F0: -1, F1: -1, F2: -650}
    switch 9223372036854775807 {
    case 32768:
        v1 += 0.0
    case 761:
        v6 := GenS2{F0: 121, F1: -1}
        _ = v6
    case 53:
        if (0 == 0) {
            v3[1] = (!true)
            var v7 int16 = 1
            var v8 int16 = v7
            _ = v7
            _ = v8
        }
    }
    v1 = 1.5
    v3[(v4 & v4)] = (65535 <= 255)
    v1 = 1.5
    switch ("\n\t" + "hello") {
    case "\n\t":
        var v9 uint16 = 65535
        _ = v9
    }
    v2 = v2
    v10 := GenS2{F0: 42, F1: 32768}
    if (1 <= 428) {
        for fi11 := 0; fi11 < 16; fi11++ {
            v4 = v4
            _ = fi11
        }
    }
    var v12 int32 = -892
    v13 := GenS3{F0: 4294967295, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
