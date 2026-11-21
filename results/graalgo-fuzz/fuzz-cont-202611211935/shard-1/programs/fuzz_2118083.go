package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int8; F2 int64 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []string{"foo bar", "hello"}
    var v2 uint64 = 275
    v3 := make(map[int]uint32)
    var v4 bool = true
    var v5 int32 = -2147483648
    v6 := GenS1{F0: 42, F1: 42}
    if len(v1) > 0 {
        v1[0] = ("" + "a")
    }
    v5 -= (v5 & (v5 >> ((702) & 63)))
    v2 = (v2 % ((42) | 1))
    v3[0] = 857
    if false {
        v7 := v3[401]
        v1 = append(v1, "foo bar")
        v8 := GenS1{F0: 1, F1: 9223372036854775807}
        _ = v7
        _ = v8
    }
    v2 -= 863
    v9 := GenS2{F0: uint32(v2)}
    v10 := GenS1{F0: -1, F1: -1}
    switch "hello" {
    case "":
        if (0 >= 65535) {
            v11 := GenS3{F0: 580}
            switch -9223372036854775808 {
            case -730:
                for fi12 := 0; fi12 < 9; fi12++ {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    v4 = false
                    _ = fi12
                }
            }
            v13 := GenS2{F0: 2147483647}
            _ = v11
            _ = v13
        } else {
            v4 = (!v4)
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 string = ""
                v16 := GenS4{F0: v5, F1: 1, F2: 1}
                if len(v1) > 0 {
                    v1[0] = v15
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    case "hello":
        v17 := GenS5{F0: 42}
        _ = v17
    case "\n\t":
        _ = len(v3)
    default:
        v5 -= (v5 & v5)
    }
    v1 = append(v1, "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
