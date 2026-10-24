package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 string; F1 uint16; F2 int64 }
type GenS4 struct { F0 bool; F1 string; F2 int16 }

func main() {
    var v1 bool = (true && false)
    var v2 uint32 = 801
    v3 := []string{"hello", ("a" + "\n\t")}
    var v4 uint8 = 127
    var v5 uint32 = uint32(v4)
    fmt.Println(v5)
    switch "hello" {
    case "hello":
        if (false && true) {
            var v6 int64 = -9223372036854775808
            v1 = false
            _ = v6
        } else {
            v1 = v1
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 uint64 = 774
                _ = fi7
                _ = v8
            }
        }
    case "":
        v9 := GenS1{F0: 127}
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 6; fi10++ {
            v1 = v1
            var v11 int16 = 0
            _ = fi10
            _ = v11
        }
    }
    if ("\n\t" < "foo bar") {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    } else {
        if (!v1) {
            var v12 uint8 = v4
            _ = v12
        } else {
            var v13 int8 = -109
            if ((!true) && (!v1)) {
                var v14 uint64 = 18446744073709551615
                v15 := GenS1{F0: 654}
                if true {
                    v2 += 907
                } else {
                    v1 = v1
                    switch "\n\t" {
                    case "\n\t":
                        v16 := GenS1{F0: -414}
                        _ = v16
                    case "foo bar":
                        v14 = v14
                    default:
                        v2 = (v5 / ((v2) | 1))
                    }
                }
                _ = v14
                _ = v15
            } else {
                v4 += (v4 - v4)
            }
            _ = v13
        }
        var v17 float64 = 3.14159
        _ = v17
    }
    switch "" {
    case "foo bar":
        if (0.1 <= 0.1) {
            v18 := GenS2{F0: 2, F1: 848, F2: v1}
            v5 = 1
            _ = v18
        }
    default:
        var v19 uint8 = (v4 >> ((1) & 63))
        _ = v19
    }
    v2 = (v2 & 2147483648)
    v20 := GenS3{F0: "hello", F1: 0, F2: 42}
    _ = len(v3)
    v1 = false
    v21 := GenS1{F0: 0}
    v4 = v4
    fmt.Println(v1, v1, v4)
    v4 += (v4 - 128)
    if (-2.25 < 1e10) {
        v22 := GenS4{F0: v1, F1: "foo bar", F2: 487}
        _ = v22
    } else {
        v1 = (!v1)
        _ = len(v3)
    }
    v5 = (v2 + (v2 >> ((65535) & 63)))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
