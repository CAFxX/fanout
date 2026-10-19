package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 uint64; F1 int; F2 string }
func genFn0(p0 int8) bool {
    v5 := GenS1{F0: "foo bar", F1: (p0 & 72)}
    v6 := GenS2{F0: 1, F1: -1}
    p0 -= (p0 | p0)
    var v7 float32 = ((1e10 * 0.1) / (3.14159 * 1e10))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (float64((((v7 * (v7 - v7)) + ((0.1 * v7) / 1.5)) - -2.25)) <= 0.0)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 int32 = -94
    v4 := make(map[string]int)
    v3 -= v3
    v8 := GenS3{F0: 42}
    if false {
        if (0 >= 2) {
            if (!(!false)) {
                var v9 uint32 = 1
                v9 = v9
                v10 := []int{32767}
                _ = v10
                _ = v9
            }
            v11 := []int{2147483647, 317, 32768, -622}
            v3 = v3
            _ = v11
        }
        if ("a" > "foo bar") {
            switch -2147483648 {
            case -9223372036854775808:
                delete(v4, "hello")
            default:
                v3 += (v3 << ((856) & 63))
            }
            var v12 string = "\n\t"
            _ = v12
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 int16 = 32767
                v15 := GenS1{F0: "foo bar", F1: 42}
                for fi16 := 0; fi16 < 11; fi16++ {
                    v3 += (v3 / ((-726) | 1))
                    v14 -= (v14 - (v14 | v14))
                    v17 := GenS2{F0: v14, F1: v3}
                    _ = fi16
                    _ = v17
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            v18 := GenS4{F0: "a", F1: 4294967295, F2: 550}
            var v19 float64 = ((float64(1.5) / 1.0) + 3.14159)
            _ = v18
            _ = v19
        }
        if false {
            fmt.Println(v3)
            var v20 bool = (22 < 2)
            if v20 {
                v20 = (false && (!(!v20)))
                v21 := GenS5{F0: 0}
                _ = v21
            }
            _ = v20
        }
    }
    if (2 == 1) {
        var v22 uint = 2
        _ = v22
    }
    v23 := GenS6{F0: 78, F1: 233}
    switch -708 {
    case 108:
        v24 := GenS3{F0: 42}
        _ = v24
    }
    v25 := GenS7{F0: 42, F1: 9223372036854775807, F2: ""}
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := v2["foo bar"]
        _ = fi26
        _ = v27
    }
    fmt.Println(v3)
    switch "\n\t" {
    case "":
        var v28 int = -9223372036854775808
        _ = v28
    case "foo bar":
        var v29 int32 = (v3 % (((-(v3))) | 1))
        _ = v29
    case "a":
        v30 := GenS1{F0: ("foo bar" + ("\n\t" + ("\n\t" + ""))), F1: 42}
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v8
}
