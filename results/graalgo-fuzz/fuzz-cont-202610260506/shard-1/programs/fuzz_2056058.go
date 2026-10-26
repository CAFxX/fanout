package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }

func main() {
    v1 := []uint32{1, 42, 2147483647}
    v2 := []string{"\n\t"}
    var v3 uint32 = 2
    v4 := []string{}
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 int = fi6
            v8 := GenS1{F0: 42}
            var v9 int32 = 2147483647
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        if len(v4) > 0 {
            v4[0] = ((("" + "\n\t") + "") + "foo bar")
        }
        _ = fi5
    }
    v4 = append(v4, "foo bar")
    fmt.Println(v3)
    switch "\n\t" {
    case "hello":
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 string = "foo bar"
            v3 = (v3 % (((v3 - v3)) | 1))
            _ = fi10
            _ = v11
        }
    default:
        v3 = (v3 % ((v3) | 1))
    }
    v3 += (v3 - (v3 / ((v3) | 1)))
    v12 := GenS1{F0: -161}
    v13 := GenS1{F0: -9223372036854775808}
    if (false && (!false)) {
        v3 = 4294967295
        fmt.Println(v3)
        switch "hello" {
        case "a":
            if true {
                var v14 int16 = 32767
                _ = v14
            } else {
                v15 := GenS1{F0: 42}
                v3 += v3
                _ = v15
            }
        case "foo bar":
            fmt.Println(v3)
        case "hello":
            v16 := GenS2{F0: 178, F1: (v3 % ((2147483648) | 1)), F2: -500}
            _ = v16
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
