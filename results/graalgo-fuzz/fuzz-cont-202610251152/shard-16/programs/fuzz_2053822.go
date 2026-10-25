package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int16; F1 uint; F2 int16 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int64; F1 uint64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = 0
    var v3 int64 = 42
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    var v6 uint16 = 32767
    v7 := GenS1{F0: -2147483648, F1: 100}
    v6 += (^(v6))
    if true {
        switch "" {
        case "a":
            var v8 float64 = 1.5
            _ = v8
        default:
            var v9 float64 = 3.14159
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        if (!((("\n\t" + ("" + "a")) == "") && true)) {
            v11 := GenS1{F0: int32(fi10), F1: 0}
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS2{F0: (v2 ^ (v2 >> ((1) & 63))), F1: 18446744073709551615, F2: v2}
                v14 := []int{}
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                fmt.Println(v6)
                v16 := GenS3{F0: int64(fi10), F1: "hello"}
                _ = fi15
                _ = v16
            }
            _ = v11
        } else {
            v1 -= v3
        }
        if (false && true) {
            v3 += v3
            fmt.Println(v1)
        }
        _ = fi10
    }
    var v17 int16 = v2
    v1 -= (v1 % (((-(v3))) | 1))
    v18 := GenS4{F0: "\n\t", F1: ("a" + (("" + "a") + "foo bar"))}
    v17 += (^(v17))
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v2)
        _ = fi19
    }
    v4[284] = false
    v3 += (v1 >> ((42) & 63))
    v20 := GenS5{F0: v3, F1: 5}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
