package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }

func main() {
    v1 := []uint32{127, 2147483647}
    v2 := []uint32{42, 939, 748, 4294967295}
    var v3 int64 = -1
    v4 := make(map[int]int)
    var v5 uint = 445
    var v6 uint32 = 65535
    var v7 int64 = 293
    v5 += v5
    fmt.Println(v6, v6, v5)
    v8 := GenS1{F0: "\n\t", F1: 2}
    v5 = v5
    v7 = v7
    v5 = v5
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int = 938
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        switch "a" {
        case "hello":
            switch ("hello" + ("hello" + "")) {
            case "a":
                v1 = append(v1, (v6 - 4294967295))
            case "\n\t":
                v5 += 65535
            case "":
                var v12 float64 = 0.1
                _ = v12
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        case "":
            v13 := GenS1{F0: "hello", F1: 988}
            _ = v13
        case "foo bar":
            v3 -= (v3 % (((v7 & v3)) | 1))
        }
        _ = fi11
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
