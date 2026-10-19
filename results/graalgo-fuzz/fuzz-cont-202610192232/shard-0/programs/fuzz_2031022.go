package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (!(4294967295 >= 42))
    v3 := []string{"foo bar", "", "", "hello"}
    var v4 uint64 = 0
    var v5 float32 = 1.5
    if v2 {
        var v6 uint8 = 173
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v2 = v2
        v8 := GenS1{F0: 247, F1: 65535}
        _ = fi7
        _ = v8
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        switch "\n\t" {
        case "a":
            v5 = (v5 - float32(0.0))
        case "hello":
            for fi9 := 0; fi9 < 11; fi9++ {
                v5 -= -2.25
                _ = fi9
            }
        }
    case -92:
        v10 := GenS1{F0: 42, F1: 51}
        _ = v10
    }
    v1[-9223372036854775808] = 9223372036854775807
    var v11 bool = v2
    v12 := GenS1{F0: 42, F1: 1}
    v2 = false
    for fi13 := 0; fi13 < 5; fi13++ {
        if (!(-32768 <= -1)) {
            if v2 {
                for fi14 := 0; fi14 < 6; fi14++ {
                    v4 = v4
                    var v15 float32 = float32(0.1)
                    v4 = v4
                    _ = fi14
                    _ = v15
                }
            } else {
                v16 := GenS2{F0: fi13, F1: 42}
                v2 = v2
                _ = v16
            }
        } else {
            v4 += 18446744073709551615
            v11 = v11
        }
        v17 := GenS1{F0: 128, F1: 68}
        _ = fi13
        _ = v17
    }
    var v18 int64 = 42
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS2{F0: (fi19 * (fi19 % ((fi19) | 1))), F1: 870}
        if (v2 || (-1 != 42)) {
            var v21 int16 = 32767
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        fmt.Println(v4)
        v2 = false
        if true {
            v18 = (^(v18))
            var v23 uint64 = (v4 % ((18446744073709551615) | 1))
            _ = v23
        }
        _ = fi22
    }
    delete(v1, 145)
    for fi24 := 0; fi24 < 3; fi24++ {
        switch "" {
        case "\n\t":
            if v11 {
                v25 := GenS2{F0: (fi24 % ((fi24) | 1)), F1: 2}
                v11 = (32767 >= 255)
                _ = v25
            }
        case "hello":
            if v2 {
                v11 = (-1 != -614)
            }
        default:
            v11 = true
        }
        _ = fi24
    }
    v18 = 9223372036854775807
    var v26 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
