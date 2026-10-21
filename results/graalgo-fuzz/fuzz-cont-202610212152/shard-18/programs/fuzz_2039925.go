package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int }
type GenS2 struct { F0 int64; F1 int; F2 int }

func main() {
    var v1 uint16 = 32767
    v2 := []string{"foo bar"}
    var v3 uint = 42
    v4 := []uint32{2147483648, 2147483647, 0}
    var v5 uint64 = 80
    v5 = uint64(v3)
    v1 = (v1 ^ 2)
    v3 = (v3 % (((v3 & 4294967295)) | 1))
    var v6 int8 = 127
    v7 := GenS1{F0: 4294967295, F1: v1, F2: -665}
    v8 := GenS2{F0: 9223372036854775807, F1: 986, F2: 2147483647}
    for fi9 := 0; fi9 < 16; fi9++ {
        switch ("\n\t" + "a") {
        case "":
            fmt.Println(v6, v1)
        case "foo bar":
            fmt.Println(v3, v3)
        case "\n\t":
            var v10 int16 = -32768
            _ = v10
        }
        _ = fi9
    }
    var v11 int = -647
    v1 += (v1 + v1)
    var v12 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
