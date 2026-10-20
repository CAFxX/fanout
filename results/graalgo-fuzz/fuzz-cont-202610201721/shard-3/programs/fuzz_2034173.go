package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    var v1 uint = 164
    v2 := make(map[string]string)
    v3 := []uint32{uint32(v1), 42}
    fmt.Println(v1)
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 int64 = 572
        if (("foo bar" + (("a" + "\n\t") + ("hello" + ("a" + "hello")))) > "foo bar") {
            v6 := v2["foo bar"]
            v3 = append(v3, 2147483647)
            for fi7 := 0; fi7 < 16; fi7++ {
                if true {
                    if false {
                        var v8 int = fi4
                        _ = v8
                    }
                    fmt.Println(fi7)
                }
                _ = fi7
            }
            _ = v6
        } else {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 int16 = -474
                var v11 int64 = (v5 - (v5 % (((v5 % ((v5) | 1))) | 1)))
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = fi4
        _ = v5
    }
    var v12 int32 = -171
    if (3.14159 < (3.14159 + -2.25)) {
        switch "a" {
        case "a":
            v1 += v1
        case "foo bar":
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(v1, v1, v1)
                v14 := GenS1{F0: "foo bar", F1: 2}
                _ = fi13
                _ = v14
            }
        case "hello":
            v15 := v2["hello"]
            _ = v15
        }
        var v16 uint8 = 93
        _ = v16
    }
    if false {
        v17 := GenS1{F0: "", F1: 18446744073709551615}
        v12 += 42
        _ = v17
    } else {
        v18 := GenS2{F0: (("foo bar" + "hello") + "foo bar"), F1: 9223372036854775807}
        _ = v18
    }
    v12 += (v12 >> (((v1 % ((1) | 1))) & 63))
    fmt.Println(v1, v12)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
