package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
type GenS2 struct { F0 int8; F1 int; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 string = ("\n\t" + "hello")
    var v2 int8 = 127
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    if false {
        v5 := GenS1{F0: 2147483647, F1: 2147483648, F2: -1}
        v1 = ((v1 + v1) + v1)
        _ = v5
    }
    var v6 uint = 154
    var v7 int = 1
    v7 = (v7 - 127)
    if true {
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 int32 = 2147483647
                fmt.Println(v7, fi8, v7)
                v11 := GenS2{F0: 20, F1: -873, F2: 9223372036854775807}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = len(v4)
            switch "hello" {
            case "\n\t":
                fmt.Println(fi8, v1)
            case "a":
                v12 := GenS3{F0: 2147483648}
                _ = v12
            }
            _ = fi8
        }
        v1 = (v1 + (v1 + v1))
        var v13 uint64 = 18446744073709551615
        _ = v13
    }
    v14 := GenS4{F0: 663}
    delete(v3, "hello")
    if (46 > 2) {
        v15 := GenS3{F0: 0}
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        switch "a" {
        case "\n\t":
            var v17 uint = 1
            _ = v17
        }
        _ = fi16
    }
    v3[v1] = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
